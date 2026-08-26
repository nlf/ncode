package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestNcodeAnthropicDebugEnvironmentWritesRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("content-type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	dump := filepath.Join(t.TempDir(), "anthropic.jsonl")
	t.Setenv("NCODE_DEBUG_ANTHROPIC", dump)
	stream, err := NewAnthropic("local-test-key", server.URL).Stream(context.Background(), Request{Model: "claude-sonnet-4-5", MaxTokens: 1})
	if err != nil {
		t.Fatalf("Stream: %v", err)
	}
	for range stream {
	}
	body, err := os.ReadFile(dump)
	if err != nil {
		t.Fatalf("read debug dump: %v", err)
	}
	if len(body) == 0 || body[len(body)-1] != '\n' {
		t.Fatalf("debug dump is not a non-empty JSONL row: %q", body)
	}
}
