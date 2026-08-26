package agent

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetchChangelogFromLocalServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/repos/nlf/ncode/releases/latest" {
			t.Errorf("path = %q", req.URL.Path)
		}
		if req.UserAgent() != "ncode-changelog" {
			t.Errorf("User-Agent = %q", req.UserAgent())
		}
		_, _ = io.WriteString(w, `{"tag_name":"v2.0.0","html_url":"https://github.com/nlf/ncode/releases/tag/v2.0.0","body":"## Changelog\n- local"}`)
	}))
	defer server.Close()

	info, err := fetchChangelogWithClient(context.Background(), "0.0.0", server.URL+"/repos/nlf/ncode", server.Client())
	if err != nil {
		t.Fatalf("fetch local changelog: %v", err)
	}
	if info.Version != "2.0.0" || info.Body != "- local" {
		t.Fatalf("changelog = %#v", info)
	}
}

func TestFetchChangelogUsesNcodeEndpointAndUserAgent(t *testing.T) {
	original := http.DefaultTransport
	t.Cleanup(func() { http.DefaultTransport = original })

	var gotURL, gotAgent string
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		gotURL = req.URL.String()
		gotAgent = req.UserAgent()
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body: io.NopCloser(strings.NewReader(`{
				"tag_name":"v1.2.3",
				"html_url":"https://github.com/nlf/ncode/releases/tag/v1.2.3",
				"body":"## Changelog\n### Fixed\n- updater"
			}`)),
			Request: req,
		}, nil
	})

	info, err := FetchChangelog(context.Background(), "1.2.3")
	if err != nil {
		t.Fatalf("fetch changelog: %v", err)
	}
	if gotURL != "https://api.github.com/repos/nlf/ncode/releases/tags/v1.2.3" {
		t.Errorf("request URL = %q", gotURL)
	}
	if gotAgent != "ncode-changelog" {
		t.Errorf("User-Agent = %q, want ncode-changelog", gotAgent)
	}
	if info.Version != "1.2.3" || info.Body != "\x00H:Fixed\n- updater" {
		t.Fatalf("changelog = %#v", info)
	}
}
