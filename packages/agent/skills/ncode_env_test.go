package skills

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNcodeAgentSkillsEnvironmentAddsSearchPath(t *testing.T) {
	extra := t.TempDir()
	dir := filepath.Join(extra, "from-env")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte("---\nname: from-env\ndescription: ncode env test\n---\nbody\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("NCODE_AGENT_SKILLS", extra)

	got, errs := Discover(t.TempDir(), t.TempDir(), "", true)
	if len(errs) != 0 {
		t.Fatalf("Discover errors: %v", errs)
	}
	if FindByName(got, "from-env") == nil {
		t.Fatal("NCODE_AGENT_SKILLS path was not discovered")
	}
}
