package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func runAnthropicDebugRequest(t *testing.T) {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("content-type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	stream, err := NewAnthropic("local-test-key", server.URL).Stream(context.Background(), Request{Model: "claude-sonnet-4-5", MaxTokens: 1})
	if err != nil {
		t.Fatalf("Stream: %v", err)
	}
	for range stream {
	}
}

func TestLegacyZotAnthropicDebugEnvironmentIsIgnoredAndNcodeWins(t *testing.T) {
	legacyDump := filepath.Join(t.TempDir(), "legacy.jsonl")
	t.Setenv("NCODE_DEBUG_ANTHROPIC", "")
	t.Setenv("ZOT_DEBUG_ANTHROPIC", legacyDump)
	runAnthropicDebugRequest(t)
	if _, err := os.Stat(legacyDump); !os.IsNotExist(err) {
		t.Fatalf("legacy debug path was written: %v", err)
	}

	ncodeDump := filepath.Join(t.TempDir(), "ncode.jsonl")
	t.Setenv("ZOT_DEBUG_ANTHROPIC", legacyDump)
	t.Setenv("NCODE_DEBUG_ANTHROPIC", ncodeDump)
	runAnthropicDebugRequest(t)
	if _, err := os.Stat(ncodeDump); err != nil {
		t.Fatalf("ncode debug path was not written: %v", err)
	}
	if _, err := os.Stat(legacyDump); !os.IsNotExist(err) {
		t.Fatalf("legacy debug path won conflict: %v", err)
	}
}
