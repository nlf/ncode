package agent

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// These fixtures prove explicit legacy Zot RPC inputs are ignored or rejected;
// they are not supported aliases, token sources, or product negotiation fields.
func TestLegacyZotRPCTokenAloneDoesNotGateNeutralPromptFirst(t *testing.T) {
	t.Setenv("ZOTCORE_RPC_TOKEN", "legacy-secret")
	t.Setenv("NCODE_RPC_TOKEN", "")
	var out bytes.Buffer
	server, _ := newRPCProtocolServer(&out)

	if err := server.run(strings.NewReader(`{"id":"1","type":"prompt","message":"hello"}` + "\n")); err != nil {
		t.Fatalf("legacy token gated neutral prompt: %v; output=%s", err, out.String())
	}
	frames := decodeRPCFrames(t, out.String())
	if frames[0]["command"] != "prompt" || frames[0]["success"] != true {
		t.Fatalf("legacy token altered prompt-first response: %#v", frames[0])
	}
}

func TestLegacyZotExtraFieldsCannotAuthorizeNcodeRPC(t *testing.T) {
	t.Setenv("ZOTCORE_RPC_TOKEN", "legacy-secret")
	t.Setenv("NCODE_RPC_TOKEN", "ncode-secret")
	var out bytes.Buffer
	server, _ := newRPCProtocolServer(&out)

	err := server.run(strings.NewReader(`{"id":"1","type":"hello","zot_token":"ncode-secret","zot_product":"zot"}` + "\n"))
	if err == nil {
		t.Fatalf("unknown Zot-branded fields authorized RPC; output=%s", out.String())
	}
	frames := decodeRPCFrames(t, out.String())
	if len(frames) != 1 || frames[0]["command"] != "hello" || frames[0]["success"] != false {
		t.Fatalf("legacy extra-field rejection changed: %#v", frames)
	}
}

func TestLegacyZotExtraFieldHasNoConfigurationMeaningWithoutToken(t *testing.T) {
	t.Setenv("ZOTCORE_RPC_TOKEN", "")
	t.Setenv("NCODE_RPC_TOKEN", "")
	var out bytes.Buffer
	server, _ := newRPCProtocolServer(&out)

	input := `{"id":"1","type":"prompt","message":"hello","zot_product":"zot"}` + "\n"
	if err := server.run(strings.NewReader(input)); err != nil {
		t.Fatalf("unknown extra field changed neutral decoding: %v; output=%s", err, out.String())
	}
	frames := decodeRPCFrames(t, out.String())
	if frames[0]["command"] != "prompt" || frames[0]["success"] != true {
		t.Fatalf("unknown extra field altered neutral prompt: %#v", frames[0])
	}
}

func TestLegacyZotExtraFieldCannotOverrideValidNcodeToken(t *testing.T) {
	t.Setenv("ZOTCORE_RPC_TOKEN", "legacy-secret")
	t.Setenv("NCODE_RPC_TOKEN", "ncode-secret")
	var out bytes.Buffer
	server, _ := newRPCProtocolServer(&out)

	input := `{"id":"1","type":"hello","token":"ncode-secret","zot_token":"wrong","zot_product":"zot"}` + "\n"
	if err := server.run(strings.NewReader(input)); err != nil {
		t.Fatalf("legacy extra field overrode valid ncode token: %v; output=%s", err, out.String())
	}
	frames := decodeRPCFrames(t, out.String())
	if len(frames) != 1 || frames[0]["command"] != "hello" || frames[0]["success"] != true {
		t.Fatalf("ncode token did not remain authoritative: %#v", frames)
	}
}

func TestLegacyZotRPCClientSourceNamesAreNotPublished(t *testing.T) {
	root := "../../examples/rpc"
	for _, path := range []string{
		root + "/python/zot_client.py",
		root + "/node/zot-client.js",
	} {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Fatalf("legacy RPC client source still published at %s (stat error %v)", path, err)
		}
	}
}
