package skills

import (
	"os"
	"path/filepath"
	"testing"
)

func writeLegacyEnvSkill(t *testing.T, root, name string) {
	t.Helper()
	dir := filepath.Join(root, name)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	body := "---\nname: " + name + "\ndescription: rejection fixture\n---\nbody\n"
	if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestLegacyZotAgentSkillsEnvironmentIsIgnoredAndNcodeWins(t *testing.T) {
	legacy := t.TempDir()
	ncode := t.TempDir()
	writeLegacyEnvSkill(t, legacy, "legacy-poison")
	writeLegacyEnvSkill(t, ncode, "ncode-extra")

	t.Setenv("ZOT_AGENT_SKILLS", legacy)
	t.Setenv("NCODE_AGENT_SKILLS", "")
	got, errs := Discover(t.TempDir(), t.TempDir(), "", true)
	if len(errs) != 0 {
		t.Fatalf("legacy-only Discover errors: %v", errs)
	}
	if FindByName(got, "legacy-poison") != nil {
		t.Fatal("legacy skill path was discovered")
	}

	t.Setenv("NCODE_AGENT_SKILLS", ncode)
	got, errs = Discover(t.TempDir(), t.TempDir(), "", true)
	if len(errs) != 0 {
		t.Fatalf("conflict Discover errors: %v", errs)
	}
	if FindByName(got, "ncode-extra") == nil || FindByName(got, "legacy-poison") != nil {
		t.Fatal("ncode skill path did not exclusively win conflict")
	}
}
