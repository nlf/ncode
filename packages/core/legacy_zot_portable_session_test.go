package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nlf/ncode/packages/provider"
)

func TestLegacyZotPortableSessionIsNotAcceptedAndRemainsUnchanged(t *testing.T) {
	root := t.TempDir()
	session, err := NewSession(root, "/legacy/project", "openai", "test-model", "test-version")
	if err != nil {
		t.Fatal(err)
	}
	if err := session.AppendMessage(provider.Message{Role: provider.RoleUser, Content: []provider.Content{provider.TextBlock{Text: "legacy Zot portable sentinel"}}}); err != nil {
		t.Fatal(err)
	}
	if err := session.Close(); err != nil {
		t.Fatal(err)
	}
	legacyPath := filepath.Join(t.TempDir(), "legacy.zotsession")
	data, err := os.ReadFile(session.Path)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(legacyPath, data, 0o640); err != nil {
		t.Fatal(err)
	}
	before, err := os.Stat(legacyPath)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := ImportSession(legacyPath, t.TempDir(), "/ncode/project", "test-version"); err == nil {
		t.Fatal("ImportSession accepted legacy .zotsession as a ncode portable session")
	}
	afterData, err := os.ReadFile(legacyPath)
	if err != nil {
		t.Fatal(err)
	}
	after, err := os.Stat(legacyPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(afterData) != string(data) || after.Mode() != before.Mode() || !after.ModTime().Equal(before.ModTime()) {
		t.Fatalf("legacy .zotsession changed: before=%v after=%v", before, after)
	}
}
