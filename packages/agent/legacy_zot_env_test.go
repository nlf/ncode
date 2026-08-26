package agent

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestLegacyZotEnvironmentInventoryHasExactDecisions(t *testing.T) {
	mapping := []struct{ legacy, ncode string }{
		{"ZOT_HOME", "NCODE_HOME"},
		{"ZOT_FLAT_TOOLS", "NCODE_FLAT_TOOLS"},
		{"ZOT_COMPACT_INPUT", "NCODE_COMPACT_INPUT"},
		{"ZOT_INLINE_IMAGES", "NCODE_INLINE_IMAGES"},
		{"ZOT_CELL_ASPECT", "NCODE_CELL_ASPECT"},
		{"ZOT_TOOL_ARG_WIDTH", "NCODE_TOOL_ARG_WIDTH"},
		{"ZOT_THEME", "NCODE_THEME"},
		{"ZOT_NO_BROWSER", "NCODE_NO_BROWSER"},
		{"ZOT_FORCE_BROWSER", "NCODE_FORCE_BROWSER"},
		{"ZOT_DEBUG_ANTHROPIC", "NCODE_DEBUG_ANTHROPIC"},
		{"ZOT_AGENT_SKILLS", "NCODE_AGENT_SKILLS"},
		{"ZOT_AGENT_CONSENT", ""},
		{"ZOTCORE_RPC_TOKEN", "NCODE_RPC_TOKEN"},
		{"ZOT_SWARM_AGENT_ID", "NCODE_SWARM_AGENT_ID"},
		{"ZOT_SWARM_EVENT_LOG", "NCODE_SWARM_EVENT_LOG"},
		{"ZOT_SWARM_CREDENTIAL_STDIN", "NCODE_SWARM_CREDENTIAL_STDIN"},
		{"ZOT_VERSION", "NCODE_VERSION"},
		{"ZOT_PREFIX", "NCODE_PREFIX"},
		{"ZOT_AGENT_API_KEY_COMMAND_HELPER", "NCODE_AGENT_API_KEY_COMMAND_HELPER"},
		{"ZOT_HELP_HELPER", "NCODE_HELP_HELPER"},
		{"ZOT_SWARM_CREDENTIAL_HELPER", "NCODE_SWARM_CREDENTIAL_HELPER"},
		{"ZOT_API_KEY_COMMAND_HELPER", "NCODE_API_KEY_COMMAND_HELPER"},
		{"ZOT_API_KEY_COMMAND_VALUE", "NCODE_API_KEY_COMMAND_VALUE"},
	}
	if len(mapping) != 23 {
		t.Fatalf("environment decision count = %d, want 23", len(mapping))
	}
	legacySeen, ncodeSeen := map[string]bool{}, map[string]bool{}
	for _, entry := range mapping {
		if entry.legacy == "" || legacySeen[entry.legacy] {
			t.Fatalf("invalid or duplicate legacy mapping %q", entry.legacy)
		}
		legacySeen[entry.legacy] = true
		if entry.legacy == "ZOT_AGENT_CONSENT" {
			if entry.ncode != "" {
				t.Fatalf("removed consent mapped to %q", entry.ncode)
			}
			continue
		}
		if entry.ncode == "" || ncodeSeen[entry.ncode] {
			t.Fatalf("invalid or duplicate ncode mapping %q", entry.ncode)
		}
		ncodeSeen[entry.ncode] = true
	}
	if _, ok := ncodeSeen["NCODE_AGENT_CONSENT"]; ok {
		t.Fatal("deleted consent must not have an NCODE replacement")
	}
}

func TestLegacyZotStateAndRenderingConfigEnvironmentIsIgnored(t *testing.T) {
	trueValue := true
	tests := []struct {
		name  string
		setup func(*testing.T)
		got   func() any
		want  any
	}{
		{name: "legacy home only", setup: func(t *testing.T) {
			t.Setenv("NCODE_HOME", "")
			t.Setenv("ZOT_HOME", filepath.Join(t.TempDir(), "legacy"))
		}, got: func() any { return NcodeHome() }, want: "not-legacy"},
		{name: "ncode home wins", setup: func(t *testing.T) {
			ncode := filepath.Join(t.TempDir(), "ncode")
			t.Setenv("ZOT_HOME", filepath.Join(t.TempDir(), "legacy"))
			t.Setenv("NCODE_HOME", ncode)
		}, got: func() any { return NcodeHome() }, want: "ncode-home"},
		{name: "legacy flat tools only", setup: func(t *testing.T) {
			t.Setenv("NCODE_FLAT_TOOLS", "")
			t.Setenv("ZOT_FLAT_TOOLS", "1")
		}, got: func() any { return Config{}.FlatToolRender() }, want: false},
		{name: "ncode flat tools wins", setup: func(t *testing.T) {
			t.Setenv("ZOT_FLAT_TOOLS", "0")
			t.Setenv("NCODE_FLAT_TOOLS", "1")
		}, got: func() any { return Config{}.FlatToolRender() }, want: true},
		{name: "legacy compact input only", setup: func(t *testing.T) {
			t.Setenv("NCODE_COMPACT_INPUT", "")
			t.Setenv("ZOT_COMPACT_INPUT", "0")
		}, got: func() any { return Config{CompactInput: &trueValue}.CompactUserInput() }, want: true},
		{name: "ncode compact input wins", setup: func(t *testing.T) {
			t.Setenv("ZOT_COMPACT_INPUT", "0")
			t.Setenv("NCODE_COMPACT_INPUT", "1")
		}, got: func() any { return Config{}.CompactUserInput() }, want: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t)
			got := tc.got()
			switch tc.want {
			case "not-legacy":
				if strings.Contains(got.(string), "legacy") {
					t.Fatalf("legacy home selected: %q", got)
				}
			case "ncode-home":
				if !strings.HasSuffix(got.(string), "ncode") {
					t.Fatalf("ncode home did not win: %q", got)
				}
			default:
				if got != tc.want {
					t.Fatalf("result = %v, want %v", got, tc.want)
				}
			}
		})
	}
}

func TestLegacyZotRPCTokenCannotGateOrAuthorize(t *testing.T) {
	tests := []struct {
		name, legacy, ncode, supplied string
		wantErr                       bool
	}{
		{name: "legacy only is ignored", legacy: "legacy-secret", supplied: "anything"},
		{name: "ncode wins conflict", legacy: "legacy-secret", ncode: "ncode-secret", supplied: "ncode-secret"},
		{name: "legacy cannot authorize conflict", legacy: "legacy-secret", ncode: "ncode-secret", supplied: "legacy-secret", wantErr: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("ZOTCORE_RPC_TOKEN", tc.legacy)
			t.Setenv("NCODE_RPC_TOKEN", tc.ncode)
			var out bytes.Buffer
			s := &rpcServer{out: &out, version: "test"}
			err := s.run(strings.NewReader(`{"id":"1","type":"hello","token":"` + tc.supplied + `"}` + "\n"))
			if (err != nil) != tc.wantErr {
				t.Fatalf("rpc error = %v, wantErr %v; output=%s", err, tc.wantErr, out.String())
			}
		})
	}
}

func TestLegacyZotSwarmChildEnvironmentIsIgnoredAndNcodeWins(t *testing.T) {
	tests := []struct {
		name                              string
		legacyCredential, ncodeCredential string
		legacyLog, ncodeLog               string
		wantCredential                    bool
		wantLog                           string
	}{
		{name: "legacy only ignored", legacyCredential: "1", legacyLog: "/tmp/legacy-events.jsonl"},
		{name: "ncode wins conflict", legacyCredential: "0", ncodeCredential: "1", legacyLog: "/tmp/legacy-events.jsonl", ncodeLog: "/tmp/ncode-events.jsonl", wantCredential: true, wantLog: "/tmp/ncode-events.jsonl"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("ZOT_SWARM_CREDENTIAL_STDIN", tc.legacyCredential)
			t.Setenv("NCODE_SWARM_CREDENTIAL_STDIN", tc.ncodeCredential)
			t.Setenv("ZOT_SWARM_EVENT_LOG", tc.legacyLog)
			t.Setenv("NCODE_SWARM_EVENT_LOG", tc.ncodeLog)
			if got := swarmCredentialStdinEnabled(); got != tc.wantCredential {
				t.Fatalf("credential stdin = %v, want %v", got, tc.wantCredential)
			}
			if got := swarmEventLogPath(); got != tc.wantLog {
				t.Fatalf("event log = %q, want %q", got, tc.wantLog)
			}
		})
	}
}

func TestLegacyZotAgentAPIKeyHelperIsIgnoredAndNcodeWins(t *testing.T) {
	tests := []struct {
		name, legacy, ncode string
		wantExecuted        bool
	}{
		{name: "legacy only ignored", legacy: "1"},
		{name: "ncode wins conflict", legacy: "0", ncode: "1", wantExecuted: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			marker := filepath.Join(t.TempDir(), "marker")
			cmd := exec.Command(os.Args[0], "-test.run=^TestAgentAPIKeyCommandHelperProcess$", "--", marker)
			cmd.Env = append(os.Environ(), "NCODE_AGENT_API_KEY_COMMAND_HELPER="+tc.ncode, "ZOT_AGENT_API_KEY_COMMAND_HELPER="+tc.legacy)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("helper subprocess: %v: %s", err, out)
			}
			executed := bytes.Contains(out, []byte("resolved-secret"))
			if executed != tc.wantExecuted {
				t.Fatalf("helper executed = %v, want %v; output=%q", executed, tc.wantExecuted, out)
			}
			_, statErr := os.Stat(marker)
			if tc.wantExecuted && statErr != nil {
				t.Fatalf("ncode helper marker missing: %v", statErr)
			}
			if !tc.wantExecuted && !os.IsNotExist(statErr) {
				t.Fatalf("legacy-only helper wrote marker: %v", statErr)
			}
		})
	}
}
