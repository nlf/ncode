package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLegacyZotHomeAndDefaultsAreIgnoredAndUnchanged(t *testing.T) {
	root := t.TempDir()
	ncodeHome := filepath.Join(root, "ncode-home")
	zotHome := filepath.Join(root, "zot-home")
	if err := os.MkdirAll(zotHome, 0o755); err != nil {
		t.Fatal(err)
	}
	legacyConfig := filepath.Join(zotHome, "config.json")
	legacyData := []byte(`{"provider":"legacy-zot-poison"}`)
	if err := os.WriteFile(legacyConfig, legacyData, 0o640); err != nil {
		t.Fatal(err)
	}
	before := snapshotPathFile(t, legacyConfig)

	t.Setenv("NCODE_HOME", ncodeHome)
	t.Setenv("ZOT_HOME", zotHome)
	if got := NcodeHome(); got != ncodeHome {
		t.Fatalf("NcodeHome() = %q, want %q; ZOT_HOME must be ignored", got, ncodeHome)
	}
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Provider != "" {
		t.Fatalf("legacy Zot config influenced ncode: %+v", cfg)
	}
	assertPathFileSnapshot(t, legacyConfig, before)

	tests := []struct {
		name string
		goos string
		env  map[string]string
		home string
		want string
	}{
		{name: "macOS Zot default ignored", goos: "darwin", home: filepath.Join(root, "user"), want: filepath.Join(root, "user", "Library", "Application Support", "ncode")},
		{name: "Windows Zot default ignored", goos: "windows", env: map[string]string{"LOCALAPPDATA": filepath.Join(root, "local"), "APPDATA": filepath.Join(root, "roaming")}, home: filepath.Join(root, "user"), want: filepath.Join(root, "local", "ncode")},
		{name: "Unix Zot default ignored", goos: "linux", home: filepath.Join(root, "user"), want: filepath.Join(root, "user", ".local", "state", "ncode")},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			getenv := func(key string) string {
				if key == "ZOT_HOME" {
					return zotHome
				}
				return tc.env[key]
			}
			if got := resolveNcodeHome(tc.goos, getenv, func() (string, error) { return tc.home, nil }); got != tc.want {
				t.Fatalf("resolveNcodeHome() = %q, want %q", got, tc.want)
			}
		})
	}
}
