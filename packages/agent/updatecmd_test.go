package agent

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestUpdateTempPatternUsesNcodePrefix(t *testing.T) {
	if got := updateTempPattern(); got != "ncode-update-" {
		t.Fatalf("update temp pattern = %q, want ncode-update-", got)
	}
}

func TestReleaseAssetNameUsesNcodeArchiveContract(t *testing.T) {
	name, format, err := releaseAssetName("1.2.3")
	if err != nil {
		t.Skipf("current platform is not released: %v", err)
	}
	ext := "tar.gz"
	if runtime.GOOS == "windows" {
		ext = "zip"
	}
	want := expectedReleaseAssetName("1.2.3", runtime.GOOS, runtime.GOARCH)
	if name != want || format != ext {
		t.Fatalf("releaseAssetName = (%q, %q), want (%q, %q)", name, format, want, ext)
	}
}

func TestDownloadAndChecksumUseLocalNcodeAsset(t *testing.T) {
	asset := expectedReleaseAssetName("1.2.3", "linux", "amd64")
	payload := []byte("local ncode archive")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/"+asset {
			t.Errorf("path = %q", req.URL.Path)
		}
		if req.UserAgent() != "ncode-updater" {
			t.Errorf("User-Agent = %q", req.UserAgent())
		}
		_, _ = w.Write(payload)
	}))
	defer server.Close()

	dir := t.TempDir()
	archivePath := filepath.Join(dir, asset)
	if err := downloadFile(context.Background(), server.URL+"/"+asset, archivePath); err != nil {
		t.Fatalf("download local asset: %v", err)
	}
	sum := fmt.Sprintf("%x", sha256.Sum256(payload))
	checksumsPath := filepath.Join(dir, "checksums.txt")
	if err := os.WriteFile(checksumsPath, []byte(sum+"  "+asset+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	got, err := lookupChecksum(checksumsPath, asset)
	if err != nil {
		t.Fatalf("lookup checksum: %v", err)
	}
	if got != sum {
		t.Fatalf("checksum = %q, want %q", got, sum)
	}
}
