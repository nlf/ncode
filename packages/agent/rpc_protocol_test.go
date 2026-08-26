package agent

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/nlf/ncode/packages/core"
	"github.com/nlf/ncode/packages/provider"
)

type rpcProtocolClient struct {
	request provider.Request
}

func (c *rpcProtocolClient) Name() string { return "rpc-protocol-test" }

func (c *rpcProtocolClient) Stream(_ context.Context, req provider.Request) (<-chan provider.Event, error) {
	c.request = req
	events := make(chan provider.Event, 3)
	events <- provider.EventStart{Provider: c.Name(), Model: req.Model}
	events <- provider.EventTextDelta{Delta: "neutral reply"}
	events <- provider.EventDone{
		Stop: provider.StopEnd,
		Message: provider.Message{
			Role:    provider.RoleAssistant,
			Content: []provider.Content{provider.TextBlock{Text: "neutral reply"}},
		},
	}
	close(events)
	return events, nil
}

func newRPCProtocolServer(out *bytes.Buffer) (*rpcServer, *rpcProtocolClient) {
	client := &rpcProtocolClient{}
	agent := core.NewAgent(client, "test-model", "", core.NewRegistry())
	return &rpcServer{
		ctx:      context.Background(),
		agent:    agent,
		provider: "test-provider",
		model:    "test-model",
		out:      out,
		version:  "test-version",
	}, client
}

func decodeRPCFrames(t *testing.T, output string) []map[string]any {
	t.Helper()
	var frames []map[string]any
	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		var frame map[string]any
		if err := json.Unmarshal(scanner.Bytes(), &frame); err != nil {
			t.Fatalf("decode RPC frame %q: %v", scanner.Text(), err)
		}
		frames = append(frames, frame)
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("scan RPC frames: %v", err)
	}
	if len(frames) == 0 {
		t.Fatal("RPC emitted no frames")
	}
	return frames
}

func TestRPCWithoutTokenAcceptsNeutralPromptAsFirstFrame(t *testing.T) {
	t.Setenv("NCODE_RPC_TOKEN", "")
	var out bytes.Buffer
	server, client := newRPCProtocolServer(&out)

	err := server.run(strings.NewReader(`{"id":"prompt-1","type":"prompt","message":"hello"}` + "\n"))
	if err != nil {
		t.Fatalf("prompt-first RPC run: %v; output=%s", err, out.String())
	}
	if got := client.request.Messages[0].Content[0].(provider.TextBlock).Text; got != "hello" {
		t.Fatalf("provider prompt = %q, want hello", got)
	}

	frames := decodeRPCFrames(t, out.String())
	response := frames[0]
	if response["type"] != "response" || response["id"] != "prompt-1" || response["command"] != "prompt" || response["success"] != true {
		t.Fatalf("prompt response shape changed: %#v", response)
	}
	data, ok := response["data"].(map[string]any)
	if !ok || data["started"] != true {
		t.Fatalf("prompt response data = %#v, want started=true", response["data"])
	}

	wantEvents := map[string]bool{"turn_start": false, "user_message": false, "text_delta": false, "assistant_message": false, "turn_end": false, "done": false}
	for _, frame := range frames[1:] {
		if typ, ok := frame["type"].(string); ok {
			if _, tracked := wantEvents[typ]; tracked {
				wantEvents[typ] = true
			}
		}
		if _, exists := frame["product"]; exists {
			t.Fatalf("neutral RPC event gained product field: %#v", frame)
		}
		if _, exists := frame["ncode_version"]; exists {
			t.Fatalf("neutral RPC event gained ncode_version field: %#v", frame)
		}
	}
	for typ, seen := range wantEvents {
		if !seen {
			t.Fatalf("missing neutral %q event in %#v", typ, frames)
		}
	}
}

func TestRPCOptionalHelloRemainsNeutralProtocolVersionOne(t *testing.T) {
	t.Setenv("NCODE_RPC_TOKEN", "")
	var out bytes.Buffer
	server, _ := newRPCProtocolServer(&out)

	if err := server.run(strings.NewReader(`{"id":"hello-1","type":"hello"}` + "\n")); err != nil {
		t.Fatalf("optional hello: %v", err)
	}
	frames := decodeRPCFrames(t, out.String())
	if len(frames) != 1 {
		t.Fatalf("hello emitted %d frames, want 1: %#v", len(frames), frames)
	}
	response := frames[0]
	if response["type"] != "response" || response["id"] != "hello-1" || response["command"] != "hello" || response["success"] != true {
		t.Fatalf("hello response shape changed: %#v", response)
	}
	data, ok := response["data"].(map[string]any)
	if !ok {
		t.Fatalf("hello data type = %T", response["data"])
	}
	if data["protocol_version"] != float64(1) || data["version"] != "test-version" || data["provider"] != "test-provider" || data["model"] != "test-model" {
		t.Fatalf("hello data changed: %#v", data)
	}
	if len(data) != 4 {
		t.Fatalf("hello data gained fields: %#v", data)
	}
}
