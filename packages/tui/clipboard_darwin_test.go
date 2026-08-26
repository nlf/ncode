//go:build darwin

package tui

import (
	"path/filepath"
	"testing"
)

func TestClipboardImageDirUsesNcodeRuntimeName(t *testing.T) {
	if got := filepath.Base(clipboardImageDir()); got != "ncode-clipboard-images" {
		t.Fatalf("clipboard image directory = %q, want ncode-clipboard-images", got)
	}
}
