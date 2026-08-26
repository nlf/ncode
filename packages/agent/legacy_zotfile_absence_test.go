package agent

import (
	"os"
	"strings"
	"testing"
)

// These legacy Zotfile literals are rejection evidence only. They prove that
// archived agents and their product-specific help are no longer active inputs.
func TestLegacyZotfileArchiveInputsAreNotRouted(t *testing.T) {
	for _, command := range []string{"inspect", "verify", "run"} {
		t.Run(command, func(t *testing.T) {
			if err := Run([]string{command, "legacy.zot", "--help"}, "test"); err != nil {
				t.Fatalf("legacy Zotfile archive was still routed: %v", err)
			}
		})
	}
}

func TestLegacyZotfileHelpAndReplacementAreAbsent(t *testing.T) {
	path := t.TempDir() + "/help.txt"
	out, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	printHelp(out, "test")
	if err := out.Close(); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	help := strings.ToLower(string(data))
	for _, absent := range []string{"zotfile", ".zot", "ncodefile", "-y, --yes"} {
		if strings.Contains(help, absent) {
			t.Fatalf("help still contains removed capability marker %q:\n%s", absent, help)
		}
	}
}
