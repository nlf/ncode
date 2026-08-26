package skills

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLegacyZotProjectSkillIsIgnoredAndUnchanged(t *testing.T) {
	home := t.TempDir()
	cwd := t.TempDir()
	legacySkill := filepath.Join(cwd, ".zot", "skills", "legacy-zot-poison", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(legacySkill), 0o755); err != nil {
		t.Fatal(err)
	}
	legacyData := []byte("---\nname: legacy-zot-poison\ndescription: must not load\n---\npoison\n")
	if err := os.WriteFile(legacySkill, legacyData, 0o640); err != nil {
		t.Fatal(err)
	}
	before, err := os.Stat(legacySkill)
	if err != nil {
		t.Fatal(err)
	}

	got, errs := Discover(home, cwd, "", true)
	if len(errs) != 0 {
		t.Fatalf("Discover read legacy Zot skill: %v", errs)
	}
	if FindByName(got, "legacy-zot-poison") != nil {
		t.Fatal("legacy Zot project skill was loaded")
	}
	afterData, err := os.ReadFile(legacySkill)
	if err != nil {
		t.Fatal(err)
	}
	after, err := os.Stat(legacySkill)
	if err != nil {
		t.Fatal(err)
	}
	if string(afterData) != string(legacyData) || after.Mode() != before.Mode() || !after.ModTime().Equal(before.ModTime()) {
		t.Fatalf("legacy Zot skill changed: before=%v after=%v data=%q", before, after, afterData)
	}
}
