package ext

import (
	"strings"
	"testing"
	"time"
)

// TestLegacyZotVersionAcknowledgementIsRejected proves the extension SDK
// does not negotiate the product-bearing version-1 acknowledgement.
func TestLegacyZotVersionAcknowledgementIsRejected(t *testing.T) {
	h := newHarness("legacy-rejection")
	runErr := make(chan error, 1)
	go func() { runErr <- h.ext.Run() }()

	if frame := h.next(t); frame.hdr.Type != "hello" {
		t.Fatalf("first extension frame = %q, want hello", frame.hdr.Type)
	}
	legacy := `{"type":"hello_ack","protocol_version":1,"zot_version":"0.3.51","provider":"anthropic","model":"claude-test","cwd":"/work"}` + "\n"
	if _, err := h.hostW.Write([]byte(legacy)); err != nil {
		t.Fatal(err)
	}

	select {
	case err := <-runErr:
		if err == nil || !strings.Contains(err.Error(), "ncode protocol v2") {
			t.Fatalf("Run error = %v, want ncode protocol v2 rejection", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("legacy acknowledgement was accepted")
	}
	h.hostW.Close()
}

// TestDualNcodeAndZotAcknowledgementIsRejected proves adding the old field to
// an otherwise valid v2 acknowledgement cannot create a dual handshake.
func TestDualNcodeAndZotAcknowledgementIsRejected(t *testing.T) {
	h := newHarness("dual-rejection")
	runErr := make(chan error, 1)
	go func() { runErr <- h.ext.Run() }()

	if frame := h.next(t); frame.hdr.Type != "hello" {
		t.Fatalf("first extension frame = %q, want hello", frame.hdr.Type)
	}
	dual := `{"type":"hello_ack","product":"ncode","protocol_version":2,"ncode_version":"0.4.0","zot_version":"0.3.51","provider":"anthropic","model":"claude-test","cwd":"/work"}` + "\n"
	if _, err := h.hostW.Write([]byte(dual)); err != nil {
		t.Fatal(err)
	}

	select {
	case err := <-runErr:
		if err == nil || !strings.Contains(err.Error(), "hello_ack") {
			t.Fatalf("Run error = %v, want dual acknowledgement rejection", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("dual acknowledgement was accepted")
	}
	h.hostW.Close()
}
