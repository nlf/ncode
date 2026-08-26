package auth

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestLegacyZotAPIKeyHelperEnvironmentIsIgnoredAndNcodeWins(t *testing.T) {
	tests := []struct {
		name                                string
		legacyHelper, legacyValue           string
		ncodeHelper, ncodeValue, wantSecret string
	}{
		{name: "legacy only ignored", legacyHelper: "1", legacyValue: "legacy-secret"},
		{name: "ncode wins conflict", legacyHelper: "1", legacyValue: "legacy-secret", ncodeHelper: "1", ncodeValue: "ncode-secret", wantSecret: "ncode-secret"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command(os.Args[0], "-test.run=^TestAPIKeyCommandHelperProcess$", "--", "print")
			cmd.Env = append(os.Environ(),
				"ZOT_API_KEY_COMMAND_HELPER="+tc.legacyHelper,
				"ZOT_API_KEY_COMMAND_VALUE="+tc.legacyValue,
				"NCODE_API_KEY_COMMAND_HELPER="+tc.ncodeHelper,
				"NCODE_API_KEY_COMMAND_VALUE="+tc.ncodeValue,
			)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("helper subprocess: %v: %s", err, out)
			}
			if bytes.Contains(out, []byte(tc.legacyValue)) {
				t.Fatalf("legacy helper value was emitted: %q", out)
			}
			if tc.wantSecret != "" && !bytes.Contains(out, []byte(tc.wantSecret)) {
				t.Fatalf("ncode helper value missing: %q", out)
			}
		})
	}
}

func TestLegacyZotBrowserEnvironmentIsIgnoredAndNcodeWins(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		want bool
	}{
		{name: "legacy no browser only is ignored", env: map[string]string{"ZOT_NO_BROWSER": "1", "DISPLAY": "1"}, want: true},
		{name: "legacy force browser only is ignored", env: map[string]string{"ZOT_FORCE_BROWSER": "1"}, want: false},
		{name: "ncode no browser wins conflict", env: map[string]string{"ZOT_FORCE_BROWSER": "1", "NCODE_NO_BROWSER": "1", "DISPLAY": "1"}, want: false},
		{name: "ncode force browser wins conflict", env: map[string]string{"ZOT_NO_BROWSER": "1", "NCODE_FORCE_BROWSER": "1"}, want: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, key := range []string{"ZOT_NO_BROWSER", "ZOT_FORCE_BROWSER", "NCODE_NO_BROWSER", "NCODE_FORCE_BROWSER", "DISPLAY", "WAYLAND_DISPLAY"} {
				t.Setenv(key, "")
			}
			for key, value := range tc.env {
				t.Setenv(key, value)
			}
			if got := hasBrowser("linux", os.Getenv, false); got != tc.want {
				t.Fatalf("hasBrowser() = %v, want %v", got, tc.want)
			}
		})
	}
}
