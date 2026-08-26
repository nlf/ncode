package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/nlf/ncode/packages/core"
	"github.com/nlf/ncode/packages/provider"
)

type headlessOpenAIRequest struct {
	Model    string                  `json:"model"`
	Messages []headlessOpenAIMessage `json:"messages"`
	Tools    []headlessOpenAITool    `json:"tools"`
}

type headlessOpenAIMessage struct {
	Role       string                   `json:"role"`
	Content    json.RawMessage          `json:"content"`
	ToolCalls  []headlessOpenAIToolCall `json:"tool_calls"`
	ToolCallID string                   `json:"tool_call_id"`
}

type headlessOpenAIToolCall struct {
	ID       string `json:"id"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function"`
}

type headlessOpenAITool struct {
	Function struct {
		Name string `json:"name"`
	} `json:"function"`
}

type recordedHeadlessRequest struct {
	Authorization string
	Body          string
	Request       headlessOpenAIRequest
}

func TestRunPrintModeComposesResolvedProviderCoreToolsAndSession(t *testing.T) {
	const (
		credential      = "headless-test-credential"
		model           = "gpt-4.1-mini"
		inheritedSystem = "persistent headless system instruction"
		appendSystem    = "append-only headless instruction"
		fixtureName     = "fixture.txt"
		fixtureContent  = "isolated read result sentinel"
		firstPrompt     = "first headless composition prompt"
		firstFinal      = "first headless final answer"
		secondPrompt    = "second resumed composition prompt"
		secondFinal     = "resumed headless final answer"
		toolCallID      = "call-read-fixture"
	)

	zotHome := t.TempDir()
	cwd := t.TempDir()
	t.Setenv("ZOT_HOME", zotHome)
	t.Setenv("OPENAI_API_KEY", "")
	if err := os.WriteFile(filepath.Join(zotHome, "SYSTEM.md"), []byte(inheritedSystem), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(cwd, fixtureName), []byte(fixtureContent+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	oldCWD, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(cwd); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(oldCWD); err != nil {
			t.Errorf("restore cwd: %v", err)
		}
	})

	emptyStdin, err := os.Create(filepath.Join(cwd, "stdin.txt"))
	if err != nil {
		t.Fatal(err)
	}
	oldStdin := os.Stdin
	os.Stdin = emptyStdin
	t.Cleanup(func() {
		os.Stdin = oldStdin
		_ = emptyStdin.Close()
	})

	stdout, err := os.Create(filepath.Join(cwd, "stdout.txt"))
	if err != nil {
		t.Fatal(err)
	}
	oldStdout := os.Stdout
	os.Stdout = stdout
	t.Cleanup(func() {
		os.Stdout = oldStdout
		_ = stdout.Close()
	})

	var (
		requestsMu sync.Mutex
		requests   []recordedHeadlessRequest
	)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, readErr := io.ReadAll(r.Body)
		if readErr != nil {
			http.Error(w, readErr.Error(), http.StatusBadRequest)
			return
		}
		var request headlessOpenAIRequest
		if err := json.Unmarshal(body, &request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		requestsMu.Lock()
		requests = append(requests, recordedHeadlessRequest{
			Authorization: r.Header.Get("Authorization"),
			Body:          string(body),
			Request:       request,
		})
		call := len(requests)
		requestsMu.Unlock()

		w.Header().Set("Content-Type", "text/event-stream")
		switch call {
		case 1:
			fmt.Fprintf(w, "data: %s\n\n", `{"choices":[{"index":0,"delta":{"tool_calls":[{"index":0,"id":"`+toolCallID+`","type":"function","function":{"name":"read","arguments":"{\"path\":\"`+fixtureName+`\"}"}}]},"finish_reason":"tool_calls"}]}`)
		case 2:
			fmt.Fprintf(w, "data: %s\n\n", `{"choices":[{"index":0,"delta":{"content":"`+firstFinal+`"},"finish_reason":"stop"}]}`)
		default:
			fmt.Fprintf(w, "data: %s\n\n", `{"choices":[{"index":0,"delta":{"content":"`+secondFinal+`"},"finish_reason":"stop"}]}`)
		}
		fmt.Fprint(w, "data: [DONE]\n\n")
	}))
	t.Cleanup(server.Close)

	sessionPath := filepath.Join(cwd, "fixed-session.jsonl")
	args := Args{
		Provider:           "openai",
		Model:              model,
		APIKey:             credential,
		BaseURL:            server.URL,
		AppendSystemPrompt: []string{appendSystem},
		Session:            sessionPath,
		CWD:                cwd,
		Tools:              []string{"read"},
		NoExt:              true,
		NoSkill:            true,
		NoContextFiles:     true,
		Prompt:             firstPrompt,
	}

	if err := runPrintMode(context.Background(), args, "headless-test"); err != nil {
		t.Fatalf("first runPrintMode: %v", err)
	}
	firstRunRequests := snapshotHeadlessRequests(&requestsMu, &requests)
	if len(firstRunRequests) != 2 {
		t.Fatalf("first run provider calls = %d, want exactly 2", len(firstRunRequests))
	}
	assertHeadlessRequestComposition(t, firstRunRequests[0], credential, model, inheritedSystem, appendSystem, firstPrompt)
	if !headlessRequestHasToolCall(firstRunRequests[1].Request, toolCallID, "read", fixtureName) {
		t.Fatalf("second provider request does not contain read tool call for %q: %+v", fixtureName, firstRunRequests[1].Request.Messages)
	}
	if !headlessRequestHasText(firstRunRequests[1].Request, "tool", fixtureContent) {
		t.Fatalf("second provider request does not contain read result %q", fixtureContent)
	}

	firstTranscript := openHeadlessSession(t, sessionPath)
	assertFirstHeadlessTranscript(t, firstTranscript, firstPrompt, toolCallID, fixtureName, fixtureContent, firstFinal)
	assertCredentialAbsentFromFile(t, sessionPath, credential)

	args.Prompt = secondPrompt
	if err := runPrintMode(context.Background(), args, "headless-test"); err != nil {
		t.Fatalf("resumed runPrintMode: %v", err)
	}
	allRequests := snapshotHeadlessRequests(&requestsMu, &requests)
	if len(allRequests) != 3 {
		t.Fatalf("provider calls after resumed invocation = %d, want 3 total", len(allRequests))
	}
	resumed := allRequests[2]
	assertHeadlessCredential(t, resumed, credential)
	if resumed.Request.Model != model {
		t.Fatalf("resumed request model = %q, want %q", resumed.Request.Model, model)
	}
	for role, text := range map[string]string{
		"user":      firstPrompt,
		"tool":      fixtureContent,
		"assistant": firstFinal,
	} {
		if !headlessRequestHasText(resumed.Request, role, text) {
			t.Errorf("resumed request missing prior %s message containing %q", role, text)
		}
	}
	if !headlessRequestHasToolCall(resumed.Request, toolCallID, "read", fixtureName) {
		t.Error("resumed request missing prior assistant read tool call")
	}
	if !headlessRequestHasText(resumed.Request, "user", secondPrompt) {
		t.Errorf("resumed request missing new user prompt %q", secondPrompt)
	}

	resumedTranscript := openHeadlessSession(t, sessionPath)
	if !transcriptHasText(resumedTranscript, provider.RoleAssistant, secondFinal) {
		t.Fatalf("resumed session missing distinct final assistant response %q", secondFinal)
	}
	assertCredentialAbsentFromFile(t, sessionPath, credential)
	if err := stdout.Sync(); err != nil {
		t.Fatal(err)
	}
	output, err := os.ReadFile(stdout.Name())
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{firstFinal, secondFinal} {
		if !strings.Contains(string(output), want) {
			t.Errorf("print output missing %q: %q", want, output)
		}
	}
	if strings.Contains(string(output), credential) {
		t.Fatal("credential leaked into print output")
	}
}

func snapshotHeadlessRequests(mu *sync.Mutex, requests *[]recordedHeadlessRequest) []recordedHeadlessRequest {
	mu.Lock()
	defer mu.Unlock()
	return append([]recordedHeadlessRequest(nil), (*requests)...)
}

func assertHeadlessRequestComposition(t *testing.T, got recordedHeadlessRequest, credential, model, inheritedSystem, appendSystem, prompt string) {
	t.Helper()
	assertHeadlessCredential(t, got, credential)
	if got.Request.Model != model {
		t.Fatalf("request model = %q, want %q", got.Request.Model, model)
	}
	if !headlessRequestHasText(got.Request, "system", inheritedSystem) || !headlessRequestHasText(got.Request, "system", appendSystem) {
		t.Fatalf("system request does not contain inherited and appended instructions")
	}
	if !headlessRequestHasText(got.Request, "user", prompt) {
		t.Fatalf("request missing user prompt %q", prompt)
	}
	for _, tool := range got.Request.Tools {
		if tool.Function.Name == "read" {
			return
		}
	}
	t.Fatal("resolved built-in read tool was not attached to provider request")
}

func assertHeadlessCredential(t *testing.T, got recordedHeadlessRequest, credential string) {
	t.Helper()
	if got.Authorization != "Bearer "+credential {
		t.Fatalf("authorization header = %q, want explicit bearer credential", got.Authorization)
	}
	if strings.Contains(got.Body, credential) {
		t.Fatal("credential leaked into provider request body")
	}
}

func headlessRequestHasText(request headlessOpenAIRequest, role, want string) bool {
	for _, message := range request.Messages {
		if message.Role != role {
			continue
		}
		var text string
		if json.Unmarshal(message.Content, &text) == nil && strings.Contains(text, want) {
			return true
		}
	}
	return false
}

func headlessRequestHasToolCall(request headlessOpenAIRequest, id, name, argumentText string) bool {
	for _, message := range request.Messages {
		if message.Role != "assistant" {
			continue
		}
		for _, call := range message.ToolCalls {
			if call.ID == id && call.Function.Name == name && strings.Contains(call.Function.Arguments, argumentText) {
				return true
			}
		}
	}
	return false
}

func openHeadlessSession(t *testing.T, path string) []provider.Message {
	t.Helper()
	session, messages, err := core.OpenSession(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := session.Close(); err != nil {
		t.Fatal(err)
	}
	return messages
}

func assertFirstHeadlessTranscript(t *testing.T, messages []provider.Message, firstPrompt, callID, fixtureName, fixtureContent, final string) {
	t.Helper()
	if !transcriptHasText(messages, provider.RoleUser, firstPrompt) {
		t.Fatalf("session user message missing %q", firstPrompt)
	}
	var foundCall bool
	for _, message := range messages {
		if message.Role != provider.RoleAssistant {
			continue
		}
		for _, content := range message.Content {
			if call, ok := content.(provider.ToolCallBlock); ok && call.ID == callID && call.Name == "read" && strings.Contains(string(call.Arguments), fixtureName) {
				foundCall = true
			}
		}
	}
	if !foundCall {
		t.Fatalf("session assistant message missing read call for %q", fixtureName)
	}
	var foundResult bool
	for _, message := range messages {
		if message.Role != provider.RoleTool {
			continue
		}
		for _, content := range message.Content {
			result, ok := content.(provider.ToolResultBlock)
			if !ok || result.CallID != callID {
				continue
			}
			for _, resultContent := range result.Content {
				if text, ok := resultContent.(provider.TextBlock); ok && strings.Contains(text.Text, fixtureContent) {
					foundResult = true
				}
			}
		}
	}
	if !foundResult {
		t.Fatalf("session tool message missing result %q", fixtureContent)
	}
	if !transcriptHasText(messages, provider.RoleAssistant, final) {
		t.Fatalf("session final assistant message missing %q", final)
	}
}

func transcriptHasText(messages []provider.Message, role provider.Role, want string) bool {
	for _, message := range messages {
		if message.Role == role && messageHasText(message, want) {
			return true
		}
	}
	return false
}

func messageHasText(message provider.Message, want string) bool {
	for _, content := range message.Content {
		if text, ok := content.(provider.TextBlock); ok && strings.Contains(text.Text, want) {
			return true
		}
	}
	return false
}

func assertCredentialAbsentFromFile(t *testing.T, path, credential string) {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), credential) {
		t.Fatalf("credential leaked into session file %q", path)
	}
}
