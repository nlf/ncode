package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func writeJSON(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}

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

func TestLoadConfigUsesNcodeGlobalAndProjectPaths(t *testing.T) {
	home := t.TempDir()
	project := t.TempDir()
	t.Setenv("NCODE_HOME", home)
	writeJSON(t, filepath.Join(home, "mcp.json"), `{"mcpServers":{"global":{"command":"global"}}}`)
	writeJSON(t, filepath.Join(project, ".ncode", "mcp.json"), `{"mcpServers":{"project":{"command":"project"}}}`)

	cfg, err := loadConfig(project)
	if err != nil {
		t.Fatalf("loadConfig: %v", err)
	}
	for _, name := range []string{"global", "project"} {
		if _, ok := cfg.MCPServers[name]; !ok {
			t.Fatalf("server %q missing from ncode config paths: %+v", name, cfg.MCPServers)
		}
	}
}

func TestLoadConfigProjectOverridesGlobal(t *testing.T) {
	home := t.TempDir()
	proj := t.TempDir()
	t.Setenv("NCODE_HOME", home)

	writeJSON(t, filepath.Join(home, "mcp.json"), `{
		"mcpServers": {
			"shared": {"command": "global-cmd"},
			"global-only": {"command": "g"}
		}
	}`)
	writeJSON(t, filepath.Join(proj, ".ncode", "mcp.json"), `{
		"mcpServers": {
			"shared": {"command": "project-cmd"}
		}
	}`)

	cfg, err := loadConfig(proj)
	if err != nil {
		t.Fatalf("loadConfig: %v", err)
	}
	if got := cfg.MCPServers["shared"].Command; got != "project-cmd" {
		t.Fatalf("project must override global: shared.command = %q", got)
	}
	if _, ok := cfg.MCPServers["global-only"]; !ok {
		t.Fatal("global-only server lost in merge")
	}
	// Defaults applied during merge:
	if got := cfg.MCPServers["shared"].RequestTimeout; got != 60 {
		t.Fatalf("default requestTimeout = %d, want 60", got)
	}
}

func TestLoadConfigMissingFilesIsNotAnError(t *testing.T) {
	t.Setenv("NCODE_HOME", t.TempDir())
	cfg, err := loadConfig(t.TempDir())
	if err != nil {
		t.Fatalf("loadConfig with no files: %v", err)
	}
	if len(cfg.MCPServers) != 0 {
		t.Fatalf("expected empty config, got %+v", cfg.MCPServers)
	}
}

func TestHandleSetupProjectRequiresCwd(t *testing.T) {
	t.Setenv("NCODE_HOME", t.TempDir())
	if _, err := handleSetup([]string{"add", "grep", "--project"}, ""); err == nil {
		t.Fatal("expected error for --project with unknown working directory")
	}
}
