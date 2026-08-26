package provider

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestLegacyZotProductHeadersAreNotEmitted(t *testing.T) {
	named := NewOpenAIResponsesNamed("token", "https://example.test/v1/responses", "openai").(*renamedClient)
	client := named.inner.(*codexClient)
	var got http.Header
	client.http.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		got = r.Header.Clone()
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"text/event-stream"}},
			Body:       io.NopCloser(strings.NewReader("")),
		}, nil
	})

	events, err := client.Stream(context.Background(), Request{
		Model:    "gpt-5.5",
		Messages: []Message{{Role: RoleUser, Content: []Content{TextBlock{Text: "hi"}}}},
	})
	if err != nil {
		t.Fatal(err)
	}
	for range events {
	}
	if got.Get("originator") != "ncode" || !strings.HasPrefix(got.Get("user-agent"), "ncode (") {
		t.Fatalf("product headers = %#v, want ncode identity", got)
	}
	if strings.Contains(strings.ToLower(got.Get("originator")+" "+got.Get("user-agent")), "zot") {
		t.Fatalf("legacy product header emitted: %#v", got)
	}
}

func TestLegacyZotGeminiImageNameIsNotCreated(t *testing.T) {
	dir := t.TempDir()
	previous, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(previous) })

	path, err := saveGeminiImageToWorkingDir("image/png", []byte("png"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(path, "ncode-gemini-image-") {
		t.Fatalf("image path = %q, want ncode prefix", path)
	}
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "zot-gemini-image-") {
			t.Fatalf("legacy image name was created: %s", entry.Name())
		}
	}
}
