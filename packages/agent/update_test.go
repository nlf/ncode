package agent

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

func TestVersionLess(t *testing.T) {
	pseudo := "0.3.35-0.20260809165004-9e3c28d4b65"
	cases := []struct {
		name string
		a, b string
		want bool
	}{
		{name: "older release", a: "0.3.33", b: "0.3.34", want: true},
		{name: "newer release", a: "0.3.35", b: "0.3.34", want: false},
		{name: "pseudo after prior release", a: pseudo, b: "0.3.34", want: false},
		{name: "pseudo before final release", a: pseudo, b: "0.3.35", want: true},
		{name: "prior release before pseudo", a: "0.3.34", b: pseudo, want: true},
		{name: "decorated build", a: "0.3.33 (abc1234, 2026-08-09)", b: "0.3.34", want: true},
		{name: "invalid current", a: "not-a-version", b: "0.3.34", want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := versionLess(tc.a, tc.b); got != tc.want {
				t.Fatalf("versionLess(%q, %q) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestBuildInfoDoesNotOfferPseudoVersionDowngrade(t *testing.T) {
	current := "0.3.35-0.20260809165004-9e3c28d4b65"
	info := buildInfo(current, "v0.3.34", "https://example.com/release")
	if info.Available {
		t.Fatalf("buildInfo(%q, %q) offered a downgrade", current, info.Latest)
	}
}

func TestFetchLatestReleaseFromLocalServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/releases/latest" {
			t.Errorf("path = %q", req.URL.Path)
		}
		if req.UserAgent() != "ncode-updater" {
			t.Errorf("User-Agent = %q", req.UserAgent())
		}
		_, _ = io.WriteString(w, `{"tag_name":"v2.0.0","html_url":"https://github.com/nlf/ncode/releases/tag/v2.0.0"}`)
	}))
	defer server.Close()

	tag, releaseURL, err := fetchLatestReleaseWithClient(context.Background(), server.URL+"/releases/latest", server.Client())
	if err != nil {
		t.Fatalf("fetch local latest release: %v", err)
	}
	if tag != "v2.0.0" || releaseURL != "https://github.com/nlf/ncode/releases/tag/v2.0.0" {
		t.Fatalf("release = (%q, %q)", tag, releaseURL)
	}
}

func TestFetchLatestReleaseUsesNcodeEndpointAndUserAgent(t *testing.T) {
	original := http.DefaultTransport
	t.Cleanup(func() { http.DefaultTransport = original })

	var gotURL, gotAgent string
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		gotURL = req.URL.String()
		gotAgent = req.UserAgent()
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader(`{"tag_name":"v1.2.3","html_url":"https://github.com/nlf/ncode/releases/tag/v1.2.3"}`)),
			Request:    req,
		}, nil
	})

	tag, releaseURL, err := fetchLatestRelease(context.Background())
	if err != nil {
		t.Fatalf("fetch latest release: %v", err)
	}
	if gotURL != "https://api.github.com/repos/nlf/ncode/releases/latest" {
		t.Errorf("request URL = %q", gotURL)
	}
	if gotAgent != "ncode-updater" {
		t.Errorf("User-Agent = %q, want ncode-updater", gotAgent)
	}
	if tag != "v1.2.3" || releaseURL != "https://github.com/nlf/ncode/releases/tag/v1.2.3" {
		t.Fatalf("release = (%q, %q)", tag, releaseURL)
	}
}
