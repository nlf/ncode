package swarm

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLegacyZotMetadataHelperProcess(t *testing.T) {
	if os.Getenv("NCODE_LEGACY_SWARM_HELPER") != "1" {
		return
	}
	if os.Getenv("NCODE_SWARM_CREDENTIAL_STDIN") != "" {
		os.Exit(2)
	}
	if os.Getenv("NCODE_SWARM_AGENT_ID") != "ncode-agent" {
		os.Exit(3)
	}
	if !strings.HasSuffix(os.Getenv("NCODE_SWARM_EVENT_LOG"), "events.jsonl") {
		os.Exit(4)
	}
	fmt.Println(`{"type":"agent_stopped","reason":"completed"}`)
	os.Exit(0)
}

func TestLegacyZotSwarmMetadataCannotSupplyCredentialOrEventLog(t *testing.T) {
	root := t.TempDir()
	legacyLog := filepath.Join(root, "legacy-events.jsonl")
	const sentinel = "legacy event log must remain unchanged\n"
	if err := os.WriteFile(legacyLog, []byte(sentinel), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("NCODE_LEGACY_SWARM_HELPER", "1")
	t.Setenv("ZOT_SWARM_AGENT_ID", "legacy-agent")
	t.Setenv("ZOT_SWARM_EVENT_LOG", legacyLog)
	t.Setenv("ZOT_SWARM_CREDENTIAL_STDIN", "1")
	t.Setenv("NCODE_SWARM_CREDENTIAL_STDIN", "")

	ncodeLog := filepath.Join(root, "ncode", "events.jsonl")
	a := &Agent{
		ID:           "ncode-agent",
		Dir:          root,
		SessionPath:  filepath.Join(root, "session.json"),
		InboxPath:    filepath.Join(root, "in.sock"),
		EventLogPath: ncodeLog,
	}
	runner := &execRunner{
		agent:   a,
		Command: []string{os.Args[0], "-test.run=^TestLegacyZotMetadataHelperProcess$"},
	}
	if err := runner.Run(context.Background(), agentSink{a: a}); err != nil {
		t.Fatal(err)
	}

	body, err := os.ReadFile(legacyLog)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != sentinel {
		t.Fatalf("legacy event log changed: %q", body)
	}
	events, err := ReadEventLog(ncodeLog)
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 2 || events[0].Type != "agent_stopped" || events[0].Data["reason"] != "completed" ||
		events[1].Type != "agent_stopped" || events[1].Data["reason"] != "exit" || events[1].Data["code"] != float64(0) {
		t.Fatalf("ncode JSONL events = %+v, want child completion followed by neutral exit lifecycle", events)
	}
}

func TestLegacyZotSocketDirectoryIsNotCreated(t *testing.T) {
	if testing.Short() {
		t.Skip("requires a Unix socket probe")
	}
	base := shortSocketDir(t)
	root := filepath.Join(t.TempDir(), "state")
	path, err := socketPathInBase(base, root, "agent-123")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(filepath.Base(filepath.Dir(path)), "ncode-swarm-") {
		t.Fatalf("socket path = %q, want ncode swarm directory", path)
	}
	entries, err := os.ReadDir(base)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "zot-swarm-") {
			t.Fatalf("legacy socket directory was created: %s", entry.Name())
		}
	}
}
