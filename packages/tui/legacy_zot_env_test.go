package tui

import (
	"testing"
	"time"
)

func TestLegacyZotRenderingEnvironmentIsIgnoredAndNcodeWins(t *testing.T) {
	tests := []struct {
		name  string
		setup func(*testing.T)
		ok    func() bool
	}{
		{name: "legacy inline only", setup: func(t *testing.T) {
			t.Setenv("TERM_PROGRAM", "ghostty")
			t.Setenv("NCODE_INLINE_IMAGES", "")
			t.Setenv("ZOT_INLINE_IMAGES", "off")
		}, ok: func() bool { return DetectImageProtocol() == ImageProtocolKitty }},
		{name: "ncode inline wins", setup: func(t *testing.T) {
			t.Setenv("TERM_PROGRAM", "ghostty")
			t.Setenv("ZOT_INLINE_IMAGES", "kitty")
			t.Setenv("NCODE_INLINE_IMAGES", "off")
		}, ok: func() bool { return DetectImageProtocol() == ImageProtocolNone }},
		{name: "legacy cell aspect only", setup: func(t *testing.T) {
			t.Setenv("NCODE_CELL_ASPECT", "")
			t.Setenv("ZOT_CELL_ASPECT", "4")
		}, ok: func() bool { return CellAspectRatio() == defaultCellAspectRatio }},
		{name: "ncode cell aspect wins", setup: func(t *testing.T) {
			t.Setenv("ZOT_CELL_ASPECT", "1")
			t.Setenv("NCODE_CELL_ASPECT", "4")
		}, ok: func() bool { return CellAspectRatio() == 4 }},
		{name: "legacy tool width only", setup: func(t *testing.T) {
			t.Setenv("NCODE_TOOL_ARG_WIDTH", "")
			t.Setenv("ZOT_TOOL_ARG_WIDTH", "120")
		}, ok: func() bool { return toolArgWidth() == defaultToolArgWidth }},
		{name: "ncode tool width wins", setup: func(t *testing.T) {
			t.Setenv("ZOT_TOOL_ARG_WIDTH", "20")
			t.Setenv("NCODE_TOOL_ARG_WIDTH", "120")
		}, ok: func() bool { return toolArgWidth() == 120 }},
		{name: "legacy theme only", setup: func(t *testing.T) {
			t.Setenv("NCODE_THEME", "")
			t.Setenv("ZOT_THEME", "light")
		}, ok: func() bool { return DetectThemeFromBackground(time.Millisecond).FG == Dark.FG }},
		{name: "ncode theme wins", setup: func(t *testing.T) {
			t.Setenv("ZOT_THEME", "dark")
			t.Setenv("NCODE_THEME", "light")
		}, ok: func() bool { return DetectThemeFromBackground(time.Millisecond).FG == Light.FG }},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t)
			if !tc.ok() {
				t.Fatal("legacy value influenced rendering or ncode failed to win")
			}
		})
	}
}
