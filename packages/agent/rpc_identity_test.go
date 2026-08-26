package agent

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNcodeRPCDocumentationAndReferenceClientsPublishOneIdentity(t *testing.T) {
	root := filepath.Clean(filepath.Join("..", ".."))
	contracts := []struct {
		path string
		want []string
	}{
		{path: "docs/rpc.md", want: []string{"# ncode RPC", "`ncode rpc`", "`NCODE_RPC_TOKEN`"}},
		{path: "examples/rpc/shell/prompt.sh", want: []string{"`ncode rpc`", "NCODE_RPC_TOKEN", "| ncode rpc"}},
		{path: "examples/rpc/go/main.go", want: []string{"`ncode rpc`", `exec.Command("ncode", "rpc")`, "NCODE_RPC_TOKEN"}},
		{path: "examples/rpc/python/ncode_client.py", want: []string{"class NcodeClient", `["ncode", "rpc"`, "NCODE_RPC_TOKEN"}},
		{path: "examples/rpc/node/ncode-client.js", want: []string{"class NcodeClient", `spawn("ncode"`, "NCODE_RPC_TOKEN"}},
	}
	for _, contract := range contracts {
		t.Run(contract.path, func(t *testing.T) {
			content, err := os.ReadFile(filepath.Join(root, contract.path))
			if err != nil {
				t.Fatalf("read RPC contract: %v", err)
			}
			for _, want := range contract.want {
				if !strings.Contains(string(content), want) {
					t.Errorf("%s missing %q", contract.path, want)
				}
			}
		})
	}
}
