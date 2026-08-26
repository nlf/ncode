package tui

import (
	"testing"
	"time"
)

func TestNcodeRenderingEnvironmentControls(t *testing.T) {
	tests := []struct {
		name  string
		env   string
		set   string
		setup func(*testing.T)
		ok    func() bool
	}{
		{name: "inline images", env: "NCODE_INLINE_IMAGES", set: "off", setup: func(t *testing.T) {
			t.Setenv("TERM_PROGRAM", "ghostty")
		}, ok: func() bool { return DetectImageProtocol() == ImageProtocolNone }},
		{name: "cell aspect", env: "NCODE_CELL_ASPECT", set: "4", ok: func() bool { return CellAspectRatio() == 4 }},
		{name: "tool argument width", env: "NCODE_TOOL_ARG_WIDTH", set: "120", ok: func() bool { return toolArgWidth() == 120 }},
		{name: "theme", env: "NCODE_THEME", set: "light", ok: func() bool { return DetectThemeFromBackground(time.Millisecond).FG == Light.FG }},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(t)
			}
			t.Setenv(tc.env, tc.set)
			if !tc.ok() {
				t.Fatalf("%s did not control rendering", tc.env)
			}
		})
	}
}
