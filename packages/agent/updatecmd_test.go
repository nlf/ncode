package agent

import "testing"

func TestUpdateTempPatternUsesNcodePrefix(t *testing.T) {
	if got := updateTempPattern(); got != "ncode-update-" {
		t.Fatalf("update temp pattern = %q, want ncode-update-", got)
	}
}
