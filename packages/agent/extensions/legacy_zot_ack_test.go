package extensions

import "testing"

// assertNoLegacyAckField keeps the rejected product-bearing field isolated in
// a dedicated fixture while manager tests prove the host never emits it.
func assertNoLegacyAckField(t *testing.T, ack map[string]any, raw []byte) {
	t.Helper()
	if _, ok := ack["zot_version"]; ok {
		t.Fatalf("legacy Zot acknowledgement field was emitted: %s", raw)
	}
}
