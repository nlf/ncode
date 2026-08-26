package agent

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

const cleanBreakNotice = "clean break from Zot"

func distributionFile(t *testing.T, name string) string {
	t.Helper()
	_, testFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve distribution test path")
	}
	path := filepath.Join(filepath.Dir(testFile), "..", "..", name)
	body, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", name, err)
	}
	return string(body)
}

func requireDistributionStrings(t *testing.T, file string, wants ...string) {
	t.Helper()
	body := distributionFile(t, file)
	for _, want := range wants {
		if !strings.Contains(body, want) {
			t.Errorf("%s missing %q", file, want)
		}
	}
}

func TestGoReleaserPublishesNcodeAssets(t *testing.T) {
	requireDistributionStrings(t, ".goreleaser.yaml",
		"project_name: ncode",
		"id: ncode",
		"main: ./cmd/ncode",
		"binary: ncode",
		"owner: nlf",
		"name: ncode",
		`name_template: "checksums.txt"`,
		"## ncode {{ .Tag }}",
		cleanBreakNotice,
	)
}

func TestInstallersPublishNcodeContract(t *testing.T) {
	for _, tc := range []struct {
		file  string
		wants []string
	}{
		{"install.sh", []string{`OWNER="nlf"`, `REPO="ncode"`, `BINARY="ncode"`, "NCODE_VERSION", "NCODE_PREFIX", "ncode-installer", cleanBreakNotice}},
		{"install.ps1", []string{`$owner  = "nlf"`, `$repo   = "ncode"`, `$binary = "ncode"`, "NCODE_VERSION", "NCODE_PREFIX", "ncode-installer", cleanBreakNotice}},
	} {
		t.Run(tc.file, func(t *testing.T) {
			requireDistributionStrings(t, tc.file, tc.wants...)
		})
	}
}

func TestReleaseWorkflowValidatesNcodeSource(t *testing.T) {
	requireDistributionStrings(t, ".github/workflows/release.yml",
		"github.com/nlf/ncode",
		"cmd/ncode",
		"goreleaser",
	)
}

func expectedReleaseAssetName(version, goos, goarch string) string {
	ext := "tar.gz"
	if goos == "windows" {
		ext = "zip"
	}
	return "ncode_" + version + "_" + goos + "_" + goarch + "." + ext
}

func TestDistributionAssetAndChecksumAgreement(t *testing.T) {
	if got := expectedReleaseAssetName("1.2.3", "linux", "amd64"); got != "ncode_1.2.3_linux_amd64.tar.gz" {
		t.Fatalf("linux asset = %q", got)
	}
	if got := expectedReleaseAssetName("1.2.3", "windows", "amd64"); got != "ncode_1.2.3_windows_amd64.zip" {
		t.Fatalf("windows asset = %q", got)
	}
	requireDistributionStrings(t, ".goreleaser.yaml",
		`name_template: "{{ .ProjectName }}_{{ .Version }}_{{ .Os }}_{{ .Arch }}"`,
		"- LICENSE", "- README.md", `name_template: "checksums.txt"`)
	requireDistributionStrings(t, "install.sh",
		`ARCHIVE="${BINARY}_${VER_NUM}_${OS}_${ARCH}.tar.gz"`,
		`CHECKSUMS_URL="${BASE_URL}/checksums.txt"`)
	requireDistributionStrings(t, "install.ps1",
		`$archive     = "${binary}_${verNum}_windows_${arch}.zip"`,
		`$checksumUrl = "$baseUrl/checksums.txt"`)
}

func TestPowerShellInstallerStaticSyntax(t *testing.T) {
	body := distributionFile(t, "install.ps1")
	for _, pair := range [][2]string{{"{", "}"}, {"(", ")"}} {
		if strings.Count(body, pair[0]) != strings.Count(body, pair[1]) {
			t.Errorf("unbalanced %s%s delimiters", pair[0], pair[1])
		}
	}
	for _, required := range []string{"[CmdletBinding()]", "param(", "try {", "finally {", "Invoke-RestMethod", "Invoke-WebRequest", "Get-FileHash", "Expand-Archive"} {
		if !strings.Contains(body, required) {
			t.Errorf("install.ps1 missing static syntax anchor %q", required)
		}
	}
}

func TestShellInstallerUsesMockedDownloadsAndFilesystem(t *testing.T) {
	root := t.TempDir()
	fakeBin := filepath.Join(root, "fake-bin")
	prefix := filepath.Join(root, "prefix")
	logPath := filepath.Join(root, "curl.log")
	if err := os.MkdirAll(fakeBin, 0o755); err != nil {
		t.Fatal(err)
	}
	writeFakeCommand(t, fakeBin, "uname", `#!/bin/sh
if [ "$1" = "-s" ]; then echo Linux; else echo x86_64; fi
`)
	writeFakeCommand(t, fakeBin, "curl", `#!/bin/sh
printf '%s\n' "$*" >> "$NCODE_TEST_CURL_LOG"
out=""
url=""
while [ "$#" -gt 0 ]; do
  if [ "$1" = "-o" ]; then shift; out="$1"; else url="$1"; fi
  shift
done
case "$url" in
  */checksums.txt) printf 'abc  ncode_1.2.3_linux_amd64.tar.gz\n' > "$out" ;;
  *) printf archive > "$out" ;;
esac
`)
	writeFakeCommand(t, fakeBin, "sha256sum", `#!/bin/sh
printf 'abc  %s\n' "$1"
`)
	writeFakeCommand(t, fakeBin, "tar", `#!/bin/sh
dst=""
while [ "$#" -gt 0 ]; do
  if [ "$1" = "-C" ]; then shift; dst="$1"; fi
  shift
done
cat > "$dst/ncode" <<'EOF'
#!/bin/sh
echo 'ncode 1.2.3'
EOF
chmod +x "$dst/ncode"
`)

	cmd := exec.Command("bash", distributionPath(t, "install.sh"), "v1.2.3", prefix)
	cmd.Env = append(os.Environ(),
		"PATH="+fakeBin+":"+os.Getenv("PATH"),
		"NCODE_TEST_CURL_LOG="+logPath,
		"GITHUB_TOKEN=",
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("mocked installer: %v\n%s", err, output)
	}
	if _, err := os.Stat(filepath.Join(prefix, "ncode")); err != nil {
		t.Fatalf("installed binary: %v", err)
	}
	logBody, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatal(err)
	}
	logText := string(logBody)
	for _, want := range []string{"-A ncode-installer", "github.com/nlf/ncode/releases/download/v1.2.3/ncode_1.2.3_linux_amd64.tar.gz", "github.com/nlf/ncode/releases/download/v1.2.3/checksums.txt"} {
		if !strings.Contains(logText, want) {
			t.Errorf("mock curl log missing %q: %s", want, logText)
		}
	}
	if !strings.Contains(string(output), cleanBreakNotice) || !strings.Contains(string(output), "run:  ncode --help") {
		t.Fatalf("installer output missing ncode notice/help:\n%s", output)
	}
}

func distributionPath(t *testing.T, name string) string {
	t.Helper()
	_, testFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve distribution path")
	}
	return filepath.Join(filepath.Dir(testFile), "..", "..", name)
}

func writeFakeCommand(t *testing.T, dir, name, body string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o755); err != nil {
		t.Fatal(err)
	}
}
