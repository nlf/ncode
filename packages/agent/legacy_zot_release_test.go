package agent

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// These fixtures prove the inherited release identity is rejected rather than
// retained as an endpoint, asset, installer variable, or distributed file.
func TestLegacyZotReleaseIdentityIsNotPublished(t *testing.T) {
	files := []string{
		".goreleaser.yaml",
		"install.sh",
		"install.ps1",
		".github/workflows/release.yml",
		"packages/agent/update.go",
		"packages/agent/updatecmd.go",
		"packages/agent/changelog.go",
	}
	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			body := distributionFile(t, file)
			for _, legacy := range []string{
				"api.github.com/repos/patriceckhart/zot",
				"github.com/patriceckhart/zot/releases",
				"raw.githubusercontent.com/patriceckhart/zot",
				"ZOT_VERSION",
				"ZOT_PREFIX",
				"zot_{{",
				"zot.exe",
			} {
				if strings.Contains(body, legacy) {
					t.Errorf("%s still contains legacy distribution identity %q", file, legacy)
				}
			}
		})
	}
}

func TestLegacyZotGeneratedArtifactsAreDeleted(t *testing.T) {
	for _, file := range []string{
		"examples/extensions/todo/zot-todo-extension",
		"examples/rpc/python/__pycache__/zot_client.cpython-314.pyc",
	} {
		if _, err := distributionFileBytes(file); err == nil {
			t.Errorf("legacy generated artifact still exists: %s", file)
		}
	}
}

func distributionFileBytes(name string) ([]byte, error) {
	_, testFile, _, ok := runtime.Caller(0)
	if !ok {
		return nil, os.ErrNotExist
	}
	root := filepath.Join(filepath.Dir(testFile), "..", "..")
	return os.ReadFile(filepath.Join(root, name))
}
