# Verification report: clean-break ncode identity

## Status

**PASS with non-blocking warnings.** The completed change satisfies all nine normative requirements and all 31 scenarios in the ncode identity specification. No implementation task remains unchecked, no active legacy identity was found outside the exact reviewed evidence classes, the authoritative race suite is green, and no archive blocker was found.

Verification was read-only except for this required report artifact. No implementation correction, commit, push, PR, branch switch, review transaction, sync, archive, or native-attempt settlement was performed.

## Executive summary

- Canonical identity is exactly `github.com/nlf/ncode`, `cmd/ncode`, `bin/ncode`, ncode release/install/update assets, ncode state/project paths, `NCODE_*`, neutral RPC v1 with `NCODE_RPC_TOKEN`, and extension acknowledgement v2 with `product:"ncode"` plus `ncode_version`.
- The clean break is enforced: no executable/import/build/release alias, state import/migration/probe/fallback, legacy environment read, product-bearing RPC fallback, swarm fallback, extension/SDK dual handshake, or Zotfile replacement was found.
- Retained no-live-provider coverage is green for provider/auth resolution, print/stream/JSON/RPC, sessions, permissions/tools, swarm, extensions, skills, themes, updater, and Telegram.
- Exact identity evidence passed: 12 provenance files, 18 dedicated rejection files, 30 allowed files, 537 provenance lines reviewed, 161 rejection lines reviewed, and zero unclassified affected files.
- Required root checks passed: tidy, list, vet, build, isolated install, race tests, examples, nested complete-copy modules, syntax/static installer checks, local package/checksum checks, and cleanup.

## Structured status and action context

| Finding | Result | Evidence |
|---|---|---|
| Active change | PASS | `clean-break-ncode-identity` is explicit and unambiguous. |
| Branch | PASS | `feat/clean-break-ncode-identity-02`, matching the delegated target. |
| Native phase status | PASS | Parent supplied authoritative `apply all_done`, tasks `44/44`, parent rows `2/2`, verify ready. |
| Artifact store | PASS | `both`; required spec, tasks, and apply-progress were read from Engram and OpenSpec. |
| Workspace ownership | PASS | Repository root and allowed edit root are `/Users/nlf/Projects/nlf/ncode`; implementation and report paths are within it. |
| Action mode | PASS | `repo-local`; not workspace-planning. |
| Review gate | PASS / disabled | Review was explicitly disabled by the user; no receipt was required or created. |
| Native attempt | DEFERRED | Proceed token `sha256:2636bf5ee581d10763b50253f5b26956379a718ca259214fa59f7337f9bfc7ac`, work unit `verify-final`; parent retains settlement. |

The cumulative apply artifact contains status snapshots from earlier work-unit boundaries where verify was still blocked. Those historical snapshots are not the final native status; the parent-provided final status and the confirmed 44/44 task state establish readiness.

## Pass/fail matrix

| Area | Result | Verification |
|---|---|---|
| Complete identity inventory | PASS | Inventory covers required categories; exact active gates and workable-file check emitted zero findings. |
| Canonical module/distribution | PASS | `go.mod` is `github.com/nlf/ncode`; build and isolated install produce only `ncode`; release/install/update tests pass. |
| Durable/project-local state | PASS | Exact `NCODE_HOME` resolution, `.ncode` discovery, `.ncodesession`, timing, coexistence, unchanged sentinels, and no legacy reads are covered and green. |
| Environment namespace | PASS | Exact 23-decision inventory is covered; ncode values work, legacy-only values are ignored, and ncode wins conflicts. |
| RPC/integration contracts | PASS | RPC remains neutral v1 and prompt-first without a token; `NCODE_RPC_TOKEN` gates first-frame hello; extension acknowledgement is exact ncode v2; swarm/subprocess/SDK negatives pass. |
| Active communication | PASS | Active-surface scans are empty; help, version, docs, examples, issues, release/update, installers, and diagnostics use lowercase ncode. |
| Clean-break communication | PASS | README install/first-run, both installers, and GoReleaser release header state the clean break, non-reuse categories, and no importer/prompt/fallback. |
| Provenance/legal context | PASS | All 537 emitted lines in the 12 exact provenance files were reviewed as baseline/history/upstream/legal/planning evidence rather than live support. |
| Retained capabilities | PASS | Focused no-live-provider package matrix and full race suite cover providers/auth, all output modes, sessions, tools, swarm, extensions, skills, themes, updater, and Telegram. |
| No aliases/import/migration/fallback | PASS | Six zero-active gates, dedicated negative tests, file/path absence, and distribution checks found none. |
| Cleanup/residue | PASS | No `cmd/zot`, legacy binaries/assets, `bin/ncode`, `bin/zot`, `dist/`, tracked Python cache, nested replacement, or temporary copied module remained. |

## Normative spec coverage

All **9 requirements / 31 scenarios** are covered.

1. **Complete identity inventory — 3/3 scenarios:** exact category inventory, neutral protocol classification, and contextual provenance review pass.
2. **Canonical module and distribution identity — 3/3:** direct build/install, canonical imports, package comments, alias/asset absence, and local archive checks pass.
3. **Ncode-only durable/project state — 4/4:** fresh roots, coexistence/unchanged sentinels, creation timing, and legacy-only fresh behavior pass.
4. **Ncode-only environment configuration — 4/4:** native success, RPC token, legacy-only ignore, and conflict precedence pass.
5. **Runtime and integration contracts — 8/8:** ncode RPC, neutral prompt/event/optional hello, explicit legacy rejection, swarm/extensions/SDK, exact extension v2 ack, idle auto-ready, legacy integration rejection, and ncode diagnostics pass.
6. **Active communication — 3/3:** active materials, neutral protocol docs, and dedicated rejection fixtures pass.
7. **Clean-break communication — 2/2:** all required channels contain the notice, and no runtime conversion prompt exists.
8. **Contextual provenance/legal attribution — 1/1:** exact mandatory line review passes.
9. **Retained capabilities — 3/3:** representative workflows, capability preservation after identity update, and focused negative/retained behavior pass.

## Task completion

- Implementation tasks: **44/44 checked**.
- Parent-owned rows: **2/2 checked**.
- Unchecked implementation lines matching `^\s*- \[ \]`: **none**.
- Malformed ownership markers: **none reported and none observed**.
- Archive completeness blocker from tasks: **none**. Archive was not performed because it is outside this delegation.

## Strict TDD compliance

Strict TDD is active in `openspec/config.yaml` and the parent prompt. Global verification guidance was loaded from `/Users/nlf/.pi/agent/gentle-ai/support/strict-tdd-verify.md`.

| Check | Result | Details |
|---|---|---|
| TDD Cycle Evidence reported | PASS | Apply-progress contains TDD Cycle Evidence tables for WU1–WU12 and focused corrections. |
| Task-to-evidence traceability | PASS | 44/44 task rows trace to RED/GREEN/triangulation or explicitly approved characterization-first mechanical evidence. |
| Reported test files exist | PASS | Current test targets exist; the old WU1 rejection filename is explicitly reconciled to `legacy_zot_zotfile_absence_test.go`, while capability files intentionally deleted by WU1 remain absent. |
| GREEN remains true | PASS | Focused packages, full `make test` race suite, nested modules, and distribution tests pass now. |
| Triangulation | PASS | Native/legacy/conflict, success/error, prompt/hello, old/dual/new acknowledgement, state timing, and retained/negative paths use distinct expectations. |
| Safety nets | PASS | Apply evidence records pre-edit package/full-suite safety nets; mechanical units explicitly avoided manufactured failures. |

The support template's literal emoji wording is not used in apply-progress, but the required semantic evidence—test files, genuine failure output where appropriate, passing results, triangulation, and safety nets—is present and cross-checked. Characterization-first rows are appropriate for approved mechanical/module/docs moves and WU12 verification coverage.

### Test layer distribution

Changed/created Go test files were scanned and executed.

| Layer | Test functions | Files | Tools |
|---|---:|---:|---|
| Unit | 187 | 55 | Go `testing` |
| Integration/local process/filesystem/HTTP | 402 | 71 | Go `testing`, `httptest`, local subprocesses, temporary files |
| E2E/browser | 0 | 0 | Not required; no browser E2E harness is installed |
| **Total** | **589** | **126** | 52 additional `t.Run` call sites |

### Assertion quality

**PASS: 0 CRITICAL, 0 WARNING.** The audit found no tautologies, self-comparisons, type-only assertions used alone, smoke-only assertions, CSS/detail assertions, mock-heavy files, or unguarded ghost loops. Added negative loops either iterate fixed non-empty tables or assert absence after a production call; dynamic event searches include explicit missing-event or shape failures. Key JSON, RPC, extension, distribution, state, swarm, and provider tests were manually reviewed in addition to the repository-wide static scan.

### Changed-file coverage

`go test -coverprofile=/tmp/ncode-verify-cover.out ./...` passed. Statement-weighted coverage across 133 changed Go production/example files is **49.2% (11,919/24,235)**; 99 files are below 80% and 19 entrypoint/example files are at 0%. This is a **non-blocking WARNING** under strict-TDD guidance. The change mechanically touched a broad inherited graph, so whole-file coverage includes substantial pre-existing behavior outside the changed identity lines.

Representative identity-critical files: extension manager 83.9%, swarm socket path 82.5%, extension protocol 75.0%, swarm runner 75.6%, portable sessions 70.3%, Codex integration 61.6%, config 54.8%, extension SDK 51.0%, RPC 46.7%, updater discovery 45.3%, update command 14.5%. Exact per-file data and uncovered ranges were generated in `/tmp/ncode-verify-changed-coverage.tsv` for this verification run.

### Quality metrics

- **Go type/build checks:** PASS via `go list ./...`, `go vet ./...`, `make build`, focused tests, and race tests.
- **Linter:** `go vet ./...` PASS; standalone ShellCheck unavailable.
- **Coverage:** PASS with the informational low-coverage warning above.

## Identity allowlist and mandatory line review

The exact audit was run before creating this report.

| Measure | Result |
|---|---:|
| Exact provenance paths | 12 |
| Dedicated `legacy_zot_*` rejection paths | 18 |
| Combined allowed paths | 30 |
| Affected tracked text files | 30 |
| Unclassified affected files (`comm -23`) | 0 |
| Provenance lines reviewed | 537 |
| Rejection lines reviewed | 161 |

Every rejection line participates in rejection, ignore, unchanged-sentinel, absence, or non-use behavior. Every provenance line is contextual baseline/history/upstream/legal/planning/execution evidence. No broad directory exception was used.

The six exact tracked gates each emitted zero lines:

```sh
git grep -nI -i -e zot -- "${audit_pathspecs[@]}"
git grep -nI -F 'github.com/patriceckhart/zot' -- "${audit_pathspecs[@]}"
git grep -nI -E 'ZOTCORE_|ZOT_[A-Z0-9_]+' -- "${audit_pathspecs[@]}"
git grep -nI -E '\.zot(session)?([^[:alnum:]_]|$)' -- "${audit_pathspecs[@]}"
git grep -nI -E 'zot[-_][[:alnum:]_-]+' -- "${audit_pathspecs[@]}"
git ls-files | grep -i zot | grep -vE '(^|/)legacy_zot_[^/]*$'
```

## Pinned inventory reproduction

- Planning SHA `18325b75cc89c75b5f4842924cb377aa5bef5c4b`: **13/2181/282/293/310/106/80**, exact match.
- Frozen `zot-baseline`: **13/2098/278/285/301/105/79**, exact match.

Order: identity-bearing filenames / content lines / affected text files / old-module lines / old-environment lines / dot-path lines / composed-name lines.

## Validation commands

### Passed

```sh
go test ./packages/provider ./packages/provider/auth ./packages/agent ./packages/agent/modes ./packages/agent/modes/telegram ./packages/core ./packages/agent/tools ./packages/agent/swarm ./packages/agent/extproto ./packages/agent/ext ./packages/agent/extensions ./packages/agent/sdk ./packages/agent/skills ./packages/tui -count=1
go mod tidy
go list ./...
go vet ./...
make build
GOBIN=/tmp/ncode-verify-install.nJTjFT go install github.com/nlf/ncode/cmd/ncode
make test
go test ./examples/... -count=1
bash -n install.sh
go test ./packages/agent -run '^(TestDistributionIdentityIsNcodeOnly|TestPowerShellInstallerStaticSyntax|TestLegacyZotReleaseIdentityIsNotPublished|TestLegacyZotGeneratedArtifactsAreDeleted|TestFetchLatestReleaseFromLocalServer|TestFetchChangelogFromLocalServer|TestDownloadFileVerifiesLocalChecksum|TestShellInstallerInstallsLocalNcodeArchive)$' -count=1
go test -coverprofile=/tmp/ncode-verify-cover.out ./...
go tool cover -func=/tmp/ncode-verify-cover.out
git diff --check
git status --short --branch
```

Nested complete-copy modules each passed these commands with a temporary replacement only in the external copy:

```sh
GOPROXY=off GOWORK=off go mod tidy
GOPROXY=off GOWORK=off go test ./... -count=1
GOPROXY=off GOWORK=off go build ./...
```

Additional checks passed: all example shell scripts with `bash -n`; one Python example via `ast.parse`; two JavaScript examples via `node --check`; local ncode tar archive containing exactly `ncode`, `README.md`, and `LICENSE`; SHA-256 verification; PowerShell delimiter-count/required-anchor static check; help/version identity; generated-output cleanup.

### Verification-tool failures and corrections

These did not expose product defects:

1. An ad hoc Python assertion scanner initially failed with `re.PatternError: invalid group reference 1`; the regex was corrected and the complete scan then reported zero banned assertion patterns.
2. An initial hand-written PowerShell quote-aware delimiter parser reported `unbalanced PowerShell delimiters`. It incorrectly treated PowerShell/comment apostrophes as string delimiters. The repository's delimiter-count/required-anchor check and `TestPowerShellInstallerStaticSyntax` both passed.
3. One compound pinned-count shell attempt was rejected by command-safety as ambiguous before execution. The exact planning and frozen count commands were rerun directly and matched all pinned values.

## Tool availability warnings

The following executables were unavailable and were not downloaded: `goreleaser`, `pwsh`, `powershell`, `shellcheck`, `yq`, and `tsc`. Therefore native GoReleaser snapshot/config execution, native PowerShell parsing/execution, ShellCheck, YAML-tool parsing, and TypeScript compilation were unavailable. Static/unit contracts, mocked local installers, local package/checksum creation, JavaScript/Python syntax checks, and nested Go builds passed. This is a warning, not a blocker.

## Review workload and PR boundary

- Forecast: High aggregate 400-line risk; chained PRs explicitly not recommended after the user's accepted combined final-PR exception.
- Compliance: WU1 is merged; WU2–WU12 appear as independently revertible commits on `feat/clean-break-ncode-identity-02` in the required order.
- `size:exception`: explicitly recorded for WU1, WU2, and the combined final PR strategy.
- Chain strategy: one final PR to `main`, issue #1 linkage, exactly `type:breaking-change`; no intermediate PRs.
- Verification boundary: only verify work was performed. No PR, push, review, sync, archive, or delivery action occurred.
- Scope creep: none found; WU12 contains one retained JSON characterization and no production change.

## Blockers and risks

### Blockers

**None.**

### Non-blocking risks/warnings

1. Native GoReleaser, PowerShell, ShellCheck, YAML, and TypeScript tooling was unavailable; substitutes passed but do not equal native execution.
2. Changed-file statement coverage is 49.2%, below the informational 80% threshold, largely because the identity move touched broad inherited files mechanically.
3. Live provider calls were intentionally not run; retained provider/auth behavior is verified with local fakes, injected transports, and `httptest` as required.
4. This verify report is created after the exact 12+18 tracked allowlist gate; it is a phase artifact and was not retroactively added to the pinned implementation inventory.

## Recommendation

Return to the parent lifecycle to settle the native verify attempt and continue only with authorized CI/final-delivery actions. Do not sync or archive from this verifier.
