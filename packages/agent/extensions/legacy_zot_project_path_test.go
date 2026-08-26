package extensions

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestLegacyZotProjectExtensionIsIgnoredAndUnchanged(t *testing.T) {
	home := t.TempDir()
	cwd := t.TempDir()
	legacyManifest := filepath.Join(cwd, ".zot", "extensions", "poison", "extension.json")
	if err := os.MkdirAll(filepath.Dir(legacyManifest), 0o755); err != nil {
		t.Fatal(err)
	}
	legacyData := []byte(`{"name":"legacy-zot-poison","exec":"./must-not-run"}`)
	if err := os.WriteFile(legacyManifest, legacyData, 0o640); err != nil {
		t.Fatal(err)
	}
	before, err := os.Stat(legacyManifest)
	if err != nil {
		t.Fatal(err)
	}

	mgr := New(home, cwd, "0.0.0-test", "", "", nil)
	if errs := mgr.Discover(context.Background()); len(errs) != 0 {
		t.Fatalf("Discover read legacy Zot extension: %v", errs)
	}
	afterData, err := os.ReadFile(legacyManifest)
	if err != nil {
		t.Fatal(err)
	}
	after, err := os.Stat(legacyManifest)
	if err != nil {
		t.Fatal(err)
	}
	if string(afterData) != string(legacyData) || after.Mode() != before.Mode() || !after.ModTime().Equal(before.ModTime()) {
		t.Fatalf("legacy Zot extension changed: before=%v after=%v data=%q", before, after, afterData)
	}
}
