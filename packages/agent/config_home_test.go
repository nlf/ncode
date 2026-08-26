package agent

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestResolveNcodeHomeOrder(t *testing.T) {
	errNoHome := errors.New("home unavailable")
	tests := []struct {
		name string
		goos string
		env  map[string]string
		home string
		err  error
		want string
	}{
		{name: "explicit home wins", goos: "linux", env: map[string]string{"NCODE_HOME": filepath.Join("explicit", "home"), "XDG_STATE_HOME": filepath.Join("xdg", "state")}, home: filepath.Join("users", "me"), want: filepath.Join("explicit", "home")},
		{name: "xdg applies on macOS", goos: "darwin", env: map[string]string{"XDG_STATE_HOME": filepath.Join("xdg", "state")}, home: filepath.Join("users", "me"), want: filepath.Join("xdg", "state", "ncode")},
		{name: "xdg applies on Windows", goos: "windows", env: map[string]string{"XDG_STATE_HOME": filepath.Join("xdg", "state"), "LOCALAPPDATA": filepath.Join("local", "appdata")}, home: filepath.Join("users", "me"), want: filepath.Join("xdg", "state", "ncode")},
		{name: "macOS application support", goos: "darwin", home: filepath.Join("users", "me"), want: filepath.Join("users", "me", "Library", "Application Support", "ncode")},
		{name: "Windows local app data", goos: "windows", env: map[string]string{"LOCALAPPDATA": filepath.Join("local", "appdata"), "APPDATA": filepath.Join("roaming", "appdata")}, home: filepath.Join("users", "me"), want: filepath.Join("local", "appdata", "ncode")},
		{name: "home state fallback", goos: "linux", home: filepath.Join("users", "me"), want: filepath.Join("users", "me", ".local", "state", "ncode")},
		{name: "relative fallback", goos: "linux", err: errNoHome, want: ".ncode"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			getenv := func(key string) string { return tc.env[key] }
			userHome := func() (string, error) { return tc.home, tc.err }
			if got := resolveNcodeHome(tc.goos, getenv, userHome); got != tc.want {
				t.Fatalf("resolveNcodeHome() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestNcodeHomeResolutionAndMissingConfigCreateNothing(t *testing.T) {
	root := filepath.Join(t.TempDir(), "not-created")
	t.Setenv("NCODE_HOME", root)

	_ = NcodeHome()
	_ = ConfigPath()
	_ = AuthPath()
	_ = SessionsPath()
	_ = LogsPath()
	if _, err := LoadConfig(); err != nil {
		t.Fatalf("LoadConfig missing defaults: %v", err)
	}
	if _, err := os.Stat(root); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("path resolution or missing config created state: %v", err)
	}

	if err := SaveConfig(Config{}); err != nil {
		t.Fatalf("SaveConfig trigger: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "config.json")); err != nil {
		t.Fatalf("SaveConfig did not create config lazily: %v", err)
	}
}

func TestNcodeHomeUsesNcodeEnvironment(t *testing.T) {
	t.Setenv("NCODE_HOME", filepath.Join("explicit", "ncode"))
	t.Setenv("XDG_STATE_HOME", filepath.Join("xdg", "state"))

	if got, want := NcodeHome(), filepath.Join("explicit", "ncode"); got != want {
		t.Fatalf("NcodeHome() = %q, want %q", got, want)
	}
}
