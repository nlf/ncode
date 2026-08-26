package agent

import (
	"path/filepath"
	"testing"
)

func TestNcodeHomePrefersExplicitHomeOverXDGStateHome(t *testing.T) {
	t.Setenv("ZOT_HOME", filepath.Join("explicit", "zot"))
	t.Setenv("XDG_STATE_HOME", filepath.Join("xdg", "state"))

	if got, want := NcodeHome(), filepath.Join("explicit", "zot"); got != want {
		t.Fatalf("NcodeHome() = %q, want %q", got, want)
	}
}

func TestNcodeHomeUsesXDGStateHome(t *testing.T) {
	t.Setenv("ZOT_HOME", "")
	t.Setenv("XDG_STATE_HOME", filepath.Join("xdg", "state"))

	if got, want := NcodeHome(), filepath.Join("xdg", "state", "zot"); got != want {
		t.Fatalf("NcodeHome() = %q, want %q", got, want)
	}
}
