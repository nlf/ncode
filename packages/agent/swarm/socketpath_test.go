package swarm

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestInboxSocketPathUsesNcodeRuntimeDirectory(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("swarm inbox transport uses Unix-domain sockets")
	}

	runtimeDir := shortSocketDir(t)
	stateRoot := filepath.Join(t.TempDir(), "state")
	t.Setenv("XDG_RUNTIME_DIR", runtimeDir)

	path, err := inboxSocketPath(stateRoot, "agent-123")
	if err != nil {
		t.Fatalf("inboxSocketPath: %v", err)
	}
	wantDir := filepath.Join(runtimeDir, "ncode-swarm-"+rootTag(stateRoot))
	if filepath.Dir(path) != wantDir {
		t.Fatalf("socket directory = %q, want exact ncode directory %q", filepath.Dir(path), wantDir)
	}
	if strings.HasPrefix(path, stateRoot+string(filepath.Separator)) {
		t.Fatalf("path = %q; transient socket must not be below state root %q", path, stateRoot)
	}
}

func TestInboxSocketPathFallsBackFromUnusableRuntimeDirectory(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("swarm inbox transport uses Unix-domain sockets")
	}

	badRuntime := filepath.Join(shortSocketDir(t), "not-a-directory")
	if err := os.WriteFile(badRuntime, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	tempDir := shortSocketDir(t)
	t.Setenv("XDG_RUNTIME_DIR", badRuntime)
	t.Setenv("TMPDIR", tempDir)

	path, err := inboxSocketPath(t.TempDir(), "agent-456")
	if err != nil {
		t.Fatalf("inboxSocketPath: %v", err)
	}
	if !strings.HasPrefix(path, tempDir+string(filepath.Separator)) {
		t.Fatalf("path = %q; want fallback below %q", path, tempDir)
	}
}

func TestInboxSocketPathShortensLongAgentID(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("swarm inbox transport uses Unix-domain sockets")
	}

	t.Setenv("XDG_RUNTIME_DIR", shortSocketDir(t))
	path, err := inboxSocketPath(t.TempDir(), strings.Repeat("agent", 100))
	if err != nil {
		t.Fatalf("inboxSocketPath: %v", err)
	}
	if len(path) > maxUnixSocketPath {
		t.Fatalf("socket path length = %d; want <= %d: %q", len(path), maxUnixSocketPath, path)
	}
}
