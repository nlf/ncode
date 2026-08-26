package agent

import (
	"bytes"
	"strings"
	"testing"
)

func TestNcodeRPCTokenAuthorizesExistingHelloContract(t *testing.T) {
	tests := []struct {
		name        string
		token       string
		wantErr     bool
		wantSuccess bool
	}{
		{name: "matching token", token: "ncode-secret", wantSuccess: true},
		{name: "wrong token", token: "wrong", wantErr: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("NCODE_RPC_TOKEN", "ncode-secret")
			var out bytes.Buffer
			s := &rpcServer{out: &out, version: "test", provider: "test-provider", model: "test-model"}
			err := s.run(strings.NewReader(`{"id":"1","type":"hello","token":"` + tc.token + `"}` + "\n"))
			if (err != nil) != tc.wantErr {
				t.Fatalf("rpc hello error = %v, wantErr %v; output=%s", err, tc.wantErr, out.String())
			}
			frames := decodeRPCFrames(t, out.String())
			if frames[0]["success"] != tc.wantSuccess {
				t.Fatalf("hello success = %#v, want %v; frame=%#v", frames[0]["success"], tc.wantSuccess, frames[0])
			}
			if tc.wantSuccess {
				data := frames[0]["data"].(map[string]any)
				if data["protocol_version"] != float64(1) || len(data) != 4 {
					t.Fatalf("authorized hello changed neutral v1 data: %#v", data)
				}
			}
		})
	}
}

func TestNcodeRPCTokenRejectsNonHelloFirstFrameAndCloses(t *testing.T) {
	t.Setenv("NCODE_RPC_TOKEN", "ncode-secret")
	var out bytes.Buffer
	s := &rpcServer{out: &out, version: "test"}
	input := strings.Join([]string{
		`{"id":"1","type":"ping"}`,
		`{"id":"2","type":"hello","token":"ncode-secret"}`,
	}, "\n") + "\n"

	err := s.run(strings.NewReader(input))
	if err == nil {
		t.Fatalf("non-hello first frame did not close authenticated RPC; output=%s", out.String())
	}
	frames := decodeRPCFrames(t, out.String())
	if len(frames) != 1 {
		t.Fatalf("authenticated RPC processed frames after first-frame failure: %#v", frames)
	}
	if frames[0]["id"] != "1" || frames[0]["command"] != "ping" || frames[0]["success"] != false {
		t.Fatalf("first-frame failure response changed: %#v", frames[0])
	}
}

func TestNcodeRPCTokenRejectsMalformedFirstFrameAndCloses(t *testing.T) {
	t.Setenv("NCODE_RPC_TOKEN", "ncode-secret")
	var out bytes.Buffer
	s := &rpcServer{out: &out, version: "test"}
	input := strings.Join([]string{
		`{"id":"broken","type":"hello"`,
		`{"id":"2","type":"hello","token":"ncode-secret"}`,
	}, "\n") + "\n"

	err := s.run(strings.NewReader(input))
	if err == nil {
		t.Fatalf("malformed first frame did not close authenticated RPC; output=%s", out.String())
	}
	frames := decodeRPCFrames(t, out.String())
	if len(frames) != 1 {
		t.Fatalf("authenticated RPC processed frames after malformed first frame: %#v", frames)
	}
	if frames[0]["command"] != "" || frames[0]["success"] != false {
		t.Fatalf("malformed first-frame response changed: %#v", frames[0])
	}
	if message, _ := frames[0]["error"].(string); !strings.Contains(message, "malformed json") {
		t.Fatalf("malformed first-frame error = %q, want malformed json", message)
	}
}

func TestRPCWithoutTokenContinuesAfterMalformedFrame(t *testing.T) {
	t.Setenv("NCODE_RPC_TOKEN", "")
	var out bytes.Buffer
	s := &rpcServer{out: &out, version: "test"}
	input := strings.Join([]string{
		`{"id":"broken","type":"ping"`,
		`{"id":"2","type":"ping"}`,
	}, "\n") + "\n"

	if err := s.run(strings.NewReader(input)); err != nil {
		t.Fatalf("token-unset RPC stopped after malformed frame: %v; output=%s", err, out.String())
	}
	frames := decodeRPCFrames(t, out.String())
	if len(frames) != 2 {
		t.Fatalf("token-unset RPC emitted %d frames, want malformed error and ping response: %#v", len(frames), frames)
	}
	if frames[0]["command"] != "" || frames[0]["success"] != false {
		t.Fatalf("malformed neutral response changed: %#v", frames[0])
	}
	data, ok := frames[1]["data"].(map[string]any)
	if frames[1]["id"] != "2" || frames[1]["command"] != "ping" || frames[1]["success"] != true || !ok || data["pong"] != true {
		t.Fatalf("token-unset RPC did not continue to ping: %#v", frames[1])
	}
}
