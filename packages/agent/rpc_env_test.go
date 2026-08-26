package agent

import (
	"bytes"
	"strings"
	"testing"
)

func TestNcodeRPCTokenAuthorizesExistingHelloContract(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		wantErr bool
	}{
		{name: "matching token", token: "ncode-secret"},
		{name: "wrong token", token: "wrong", wantErr: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("NCODE_RPC_TOKEN", "ncode-secret")
			var out bytes.Buffer
			s := &rpcServer{out: &out, version: "test"}
			err := s.run(strings.NewReader(`{"id":"1","type":"hello","token":"` + tc.token + `"}` + "\n"))
			if (err != nil) != tc.wantErr {
				t.Fatalf("rpc hello error = %v, wantErr %v; output=%s", err, tc.wantErr, out.String())
			}
		})
	}
}
