package agent

import (
	"os"
	"testing"
	"time"
)

type pathFileSnapshot struct {
	data    string
	mode    os.FileMode
	size    int64
	modTime time.Time
}

func snapshotPathFile(t *testing.T, path string) pathFileSnapshot {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	return pathFileSnapshot{data: string(data), mode: info.Mode(), size: info.Size(), modTime: info.ModTime()}
}

func assertPathFileSnapshot(t *testing.T, path string, want pathFileSnapshot) {
	t.Helper()
	got := snapshotPathFile(t, path)
	if got != want {
		t.Fatalf("path changed: %s\nbefore: %#v\nafter:  %#v", path, want, got)
	}
}
