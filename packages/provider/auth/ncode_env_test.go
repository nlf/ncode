package auth

import "testing"

func TestNcodeBrowserEnvironmentControls(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want bool
	}{
		{name: "force browser", env: "NCODE_FORCE_BROWSER", want: true},
		{name: "disable browser", env: "NCODE_NO_BROWSER", want: false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("NCODE_FORCE_BROWSER", "")
			t.Setenv("NCODE_NO_BROWSER", "")
			t.Setenv(tc.env, "1")
			if got := HasBrowser(); got != tc.want {
				t.Fatalf("HasBrowser() = %v, want %v", got, tc.want)
			}
		})
	}
}
