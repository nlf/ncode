package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLegacyZotMCPConfigsAreIgnoredAndUnchanged(t *testing.T) {
	root := t.TempDir()
	ncodeHome := filepath.Join(root, "ncode-home")
	zotHome := filepath.Join(root, "zot-home")
	project := filepath.Join(root, "project")
	t.Setenv("NCODE_HOME", ncodeHome)
	t.Setenv("ZOT_HOME", zotHome)

	legacyGlobal := filepath.Join(zotHome, "mcp.json")
	legacyProject := filepath.Join(project, ".zot", "mcp.json")
	legacyData := []byte(`{"mcpServers":{"legacy-zot-poison":{"command":"must-not-run"}}}`)
	for _, path := range []string{legacyGlobal, legacyProject} {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, legacyData, 0o640); err != nil {
			t.Fatal(err)
		}
	}
	writeJSON(t, filepath.Join(ncodeHome, "mcp.json"), `{"mcpServers":{"ncode":{"command":"ok"}}}`)

	cfg, err := loadConfig(project)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := cfg.MCPServers["legacy-zot-poison"]; ok {
		t.Fatal("legacy Zot MCP config was loaded")
	}
	if _, ok := cfg.MCPServers["ncode"]; !ok {
		t.Fatal("ncode MCP config was not loaded")
	}
	for _, path := range []string{legacyGlobal, legacyProject} {
		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != string(legacyData) {
			t.Fatalf("legacy Zot MCP config %s changed: %q", path, got)
		}
	}
}
