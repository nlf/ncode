package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLegacyZotProjectExtensionsAreIgnoredAndUnchanged(t *testing.T) {
	home := t.TempDir()
	cwd := t.TempDir()
	t.Setenv("NCODE_HOME", home)
	legacyManifest := filepath.Join(cwd, ".zot", "extensions", "poison", "extension.json")
	if err := os.MkdirAll(filepath.Dir(legacyManifest), 0o755); err != nil {
		t.Fatal(err)
	}
	legacyData := []byte(`{"name":"legacy-zot-poison","exec":"./must-not-run"}`)
	if err := os.WriteFile(legacyManifest, legacyData, 0o640); err != nil {
		t.Fatal(err)
	}
	before := snapshotPathFile(t, legacyManifest)
	oldCWD, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(oldCWD) })
	if err := os.Chdir(cwd); err != nil {
		t.Fatal(err)
	}

	for _, dir := range extensionDirs() {
		if dir == filepath.Join(cwd, ".zot", "extensions") {
			t.Fatalf("extensionDirs scanned legacy Zot project path %q", dir)
		}
	}
	for _, row := range scanExtDoctorStatic() {
		if row.Name == "legacy-zot-poison" {
			t.Fatal("extension doctor read legacy Zot project manifest")
		}
	}
	assertPathFileSnapshot(t, legacyManifest, before)
}
