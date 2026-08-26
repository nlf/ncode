package agent

import (
	"bytes"
	"testing"

	"github.com/nlf/ncode/packages/core"
)

func TestRPCSetReasoningMax(t *testing.T) {
	var out bytes.Buffer
	s := &rpcServer{agent: &core.Agent{}, out: &out}
	s.dispatch("set_reasoning", "1", []byte(`{"reasoning":"max"}`))

	if s.agent.Reasoning != "max" {
		t.Fatalf("reasoning = %q, want max", s.agent.Reasoning)
	}
	frames := decodeRPCFrames(t, out.String())
	if len(frames) != 1 || frames[0]["type"] != "response" || frames[0]["command"] != "set_reasoning" || frames[0]["success"] != true {
		t.Fatalf("reasoning response shape changed: %#v", frames)
	}
	data, ok := frames[0]["data"].(map[string]any)
	if !ok || data["reasoning"] != "max" {
		t.Fatalf("reasoning response data = %#v", frames[0]["data"])
	}
}

func TestRPCSetReasoningRejectsUnknownLevel(t *testing.T) {
	var out bytes.Buffer
	s := &rpcServer{agent: &core.Agent{}, out: &out}
	s.dispatch("set_reasoning", "1", []byte(`{"reasoning":"extreme"}`))

	if s.agent.Reasoning != "" {
		t.Fatalf("reasoning = %q, want unchanged", s.agent.Reasoning)
	}
	frames := decodeRPCFrames(t, out.String())
	if len(frames) != 1 || frames[0]["type"] != "response" || frames[0]["command"] != "set_reasoning" || frames[0]["success"] != false {
		t.Fatalf("reasoning error response shape changed: %#v", frames)
	}
}
