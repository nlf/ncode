# Apply progress: clean-break-ncode-identity

Updated: 2026-08-25T23:54:09Z

## Attempt summary

- **Status:** WU 1 complete; the broader change remains in progress.
- **Assigned boundary:** WU 1 — Remove Zotfiles without replacement.
- **Delivery:** stacked PRs to `main`; current boundary is WU 1 only.
- **Review budget:** 2,865 changed lines (165 additions + 2,700 deletions), excluding the pre-existing `docs/ncode-architecture.md` worktree edit and OpenSpec artifacts. The user-approved WU 1-only `size:exception` applies because this slice is deletion-heavy and atomic.
- **Commits/branches/PRs:** none created, as instructed.
- **Independent verification:** VERIFIED after correcting the final table-driven test name; focused WU1 tests, `git diff --check`, and `make test` (`go test -race ./...`) passed with no repository mutation.
- **Review gate:** native ordinary review START was attempted after staging the new rejection test so it entered scope, but returned `review-mode-disabled` with `lineage_created:false`. No review receipt exists; WU2 and publication remain blocked pending explicit review-mode enablement and a completed review.
- **Engram mirror:** apply progress saved as observation 121; the shared-helper discovery was saved as bugfix observation 122; tasks observation 117 was updated to 4/44.

## Structured status consumed/produced

```yaml
schemaName: spec-driven
changeName: clean-break-ncode-identity
artifactStore: both
planningHome:
  root: /Users/nlf/Projects/nlf/ncode/openspec
  changesDir: /Users/nlf/Projects/nlf/ncode/openspec/changes
changeRoot: /Users/nlf/Projects/nlf/ncode/openspec/changes/clean-break-ncode-identity
artifacts:
  proposal: done
  specs: done
  design: done
  tasks: done
  applyProgress: done
  verifyReport: missing
  syncReport: missing
taskProgress:
  total: 44
  complete: 4
  remaining: 40
deferredParentActions:
  total: 2
  complete: 1
  remaining: 1
taskArtifactErrors: []
applyState: ready
dependencies:
  apply: ready
  verify: blocked
  sync: blocked
  archive: blocked
actionContext:
  mode: repo-local
  workspaceRoot: /Users/nlf/Projects/nlf/ncode
  allowedEditRoots:
    - /Users/nlf/Projects/nlf/ncode
  warnings:
    - Parent omitted a structured actionContext; repo-local scope was reconstructed from the explicit authoritative repository path.
    - The gentle-ai binary was unavailable, so status used the installed manual fallback contract.
nextRecommended: parent-lifecycle
isNonAuthoritative: false
```

The active change was explicitly selected by the parent and confirmed in the `both` store through complete OpenSpec artifacts plus Engram observations 114, 115, and 117. No malformed task ownership markers were found.

## Completed implementation tasks and persisted checkboxes

- [x] WU 1 RED — persisted at `tasks.md:67`.
- [x] WU 1 GREEN — persisted at `tasks.md:68`.
- [x] WU 1 TRIANGULATE — persisted at `tasks.md:69`.
- [x] WU 1 REFACTOR — persisted at `tasks.md:70`.

## TDD Cycle Evidence

| Task | Test files | Layer | Safety Net | RED | GREEN | TRIANGULATE | REFACTOR |
|---|---|---|---|---|---|---|---|
| WU1 RED | `packages/agent/args_test.go`, `packages/agent/build_test.go` | Package/unit | `go test ./packages/agent/... ./packages/core/... ./packages/tui/...` PASS; baseline `make test` PASS | `go test ./packages/agent -run '^TestRemovedPortableAgentCommandFallsThroughToNormalHelp$' -count=1` FAIL: `stat --help: no such file or directory`; RED `make test` FAIL only on that new assertion | N/A | N/A | N/A |
| WU1 GREEN | same, plus production deletion | Package/integration | Baseline above | RED inherited from WU1 RED | `go test ./packages/agent -run '^(TestRemovedPortableAgentCommandFallsThroughToNormalHelp|TestResolveAppliesArgsPermissionSetToSandbox)$' -count=1` PASS; focused package suite PASS; `make test` PASS | N/A | N/A |
| WU1 TRIANGULATE | `packages/agent/args_test.go`, `packages/agent/legacy_zotfile_absence_test.go` | Package/negative | GREEN above | Initial command probe failed before deletion | Four removed command probes, two removed consent-flag probes, three legacy archive/path probes, and help absence probe PASS; retained tools/skills/sessions/extensions/confirmation/headless checks PASS; `make test` PASS | 11 behavioral cases across 5 test functions | N/A |
| WU1 REFACTOR | retained tests above | Package/regression | TRIANGULATE green | N/A | N/A | N/A | Capability-dead startup-pre/display/session helpers and comments cleaned; `go test ./packages/agent/... ./packages/core/... -count=1` PASS; final `make test` PASS |

### Test summary

- **New tests:** 5 test functions with 11 behavioral cases.
- **Layers:** package-level unit, negative routing/help, and retained integration evidence.
- **Approval/characterization evidence:** direct permission wiring and existing retained-capability suites.
- **Pure functions created:** none; this unit removes a capability.
- **Live providers/credentials:** none used.

## Exact verification commands and outcomes

### Safety net

- `go test ./packages/agent/... ./packages/core/... ./packages/tui/...` — PASS.
- `make test` — PASS (`go test -race ./...`).

### RED

- `go test ./packages/agent -run '^TestRemovedPortableAgentCommandFallsThroughToNormalHelp$' -count=1` — expected FAIL: removed `inspect` was still routed and attempted `stat --help`.
- `make test` — expected FAIL only at the new removed-command assertion.

### GREEN

- `go test ./packages/agent -run '^(TestRemovedPortableAgentCommandFallsThroughToNormalHelp|TestResolveAppliesArgsPermissionSetToSandbox)$' -count=1` — PASS after deletion and moving the shared `prepareRuntimeCatalog` helper out of the deleted file.
- `go test ./packages/agent/... ./packages/core/... ./packages/tui/...` — PASS.
- `make test` — PASS.

### TRIANGULATE

- `go test ./packages/agent -run '^(TestRemovedPortableAgentCommandsFallThroughToNormalHelp|TestRemovedPortableAgentConsentFlagsAreRejected|TestLegacyZotfileArchiveInputsAreNotRouted|TestLegacyZotfileHelpAndReplacementAreAbsent|TestResolveAppliesArgsPermissionSetToSandbox|TestRunPrintModeComposesResolvedProviderCoreToolsAndSession)$' -count=1` — PASS.
- `go test ./packages/agent/tools ./packages/agent/skills ./packages/core ./packages/agent/extensions -count=1` — PASS.
- `go test ./packages/agent/modes -run '^(TestConfirm|TestStartupResources)' -count=1` — PASS.
- `git grep -nI -E 'Zotfile|zotfile|ncodefile|ZOT_AGENT_CONSENT|\.zot' -- packages/agent docs README.md` — executed; remaining output is factual provenance or identity/path work explicitly assigned to later WUs, not a surviving Zotfile capability.
- `make test` — PASS.

### REFACTOR/final

- `rg -n 'Zotfile|zotfile|ncodefile|ZOT_AGENT_CONSENT|StartupPre|StartupAgentName|AgentName|AgentDataDir|runZotfileCommand|ZotfileManifest' packages/agent packages/core packages/tui README.md AGENTS.md CONTRIBUTING.md --glob '!legacy_zotfile_absence_test.go'` — no matches.
- `rg -n -A2 -B2 'sandbox\.SetPermissions\(args\.PermissionSet\)' packages/agent/build.go` — exact wiring remains at lines 573–574.
- File absence assertions for `packages/agent/zotfile.go`, `packages/agent/zotfile_test.go`, `packages/agent/zotfile_pre_test.go`, and `docs/zotfiles.md` — PASS.
- `go test ./packages/agent/... ./packages/core/... -count=1` — PASS.
- `go test ./packages/agent -run '^TestResolveAppliesArgsPermissionSetToSandbox$' -count=1` — PASS with isolated state.
- `git diff --check` — PASS.
- Final `make test` — PASS.

Two initial REFACTOR search attempts used ineffective `rg` glob exclusions and stopped before package verification; the corrected exact basename exclusion produced the recorded zero-match result. This was a search-command issue, not a product or test failure.

### Independent parent verification and review attempt

- An initial independent verifier invocation used the historical singular RED/GREEN test name after TRIANGULATE had expanded and renamed it; Go reported `[no tests to run]`, so that procedural attempt failed closed.
- The corrected final command `go test ./packages/agent -run '^TestRemovedPortableAgentCommandsFallThroughToNormalHelp$' -count=1` passed.
- The direct-permission test, complete focused WU1 expression, `git diff --check`, and final `make test` all passed. Absence and retained-capability findings remained valid, and before/after status was unchanged.
- Native `gentle_review inspect` admitted the tracked WU1 scope after `packages/agent/legacy_zotfile_absence_test.go` was staged. `gentle_review start` then skipped because receipt-driven review mode is globally disabled; it created no lineage and performed no mutation. Review enablement is a parent-controlled blocker, not verification evidence.

## Retained-capability evidence

- **Permission wiring:** new `TestResolveAppliesArgsPermissionSetToSandbox` proves the exact `Args.PermissionSet` pointer reaches `Sandbox.SetPermissions` and enforces allowed/denied read scopes.
- **Normal tools + headless composition:** existing `TestRunPrintModeComposesResolvedProviderCoreToolsAndSession` proves resolved provider/core/read-tool/session behavior with a local `httptest` server.
- **Tools and sandbox:** all tests in `packages/agent/tools` pass.
- **Skills:** all tests in `packages/agent/skills` pass; project/global discovery and the `skill` tool remain.
- **Sessions:** all tests in `packages/core` pass; normal JSONL, portable, prune, repair, and confirmation behavior remain.
- **Extensions:** all tests in `packages/agent/extensions` pass, including subprocess tool and interceptor behavior.
- **Confirmation:** focused `packages/agent/modes` confirmation tests and all `packages/core` confirmation tests pass.
- **Print/stream/JSON/RPC/swarm composition:** final `make test` passes all repository package tests; no neutral RPC, swarm, provider, or extension protocol identity was renamed in WU 1.

## Files changed by WU 1

### Added

- `packages/agent/legacy_zotfile_absence_test.go`
- `openspec/changes/clean-break-ncode-identity/apply-progress.md`

### Modified

- `AGENTS.md`
- `CONTRIBUTING.md`
- `README.md`
- `packages/agent/args.go`
- `packages/agent/args_test.go`
- `packages/agent/build_test.go`
- `packages/agent/cli.go`
- `packages/agent/modes/interactive.go`
- `packages/agent/modes/startup_instructions_test.go`
- `packages/agent/tools/permissions.go`
- `packages/tui/view.go`
- `packages/tui/view_tool_overlay_test.go`
- `openspec/changes/clean-break-ncode-identity/tasks.md`

### Deleted

- `docs/zotfiles.md`
- `packages/agent/zotfile.go`
- `packages/agent/zotfile_test.go`
- `packages/agent/zotfile_pre_test.go`
- `packages/agent/modes/startup_pre_test.go`

`docs/ncode-architecture.md` was already modified before this apply attempt and was not edited by this executor.

## Design deviations and discoveries

- **Evidence placement deviation:** the task forecast named many retained test files, but the user explicitly required the smallest missing behavioral evidence and allowed existing focused tests to count. New tests were limited to command/help/archive absence and direct permission wiring; existing focused suites supplied tools, skills, sessions, extensions, confirmation, and headless evidence.
- **Shared-helper discovery:** `prepareRuntimeCatalog` was located in the deleted Zotfile implementation even though normal CLI startup uses it. It was moved unchanged to `packages/agent/cli.go`; deleting it would have broken all ordinary runtime catalog setup.
- **No identity work:** module/import/product/state/protocol names were not renamed. Only Zotfile-specific `ZOT_AGENT_CONSENT` disappeared with the deleted capability.
- **Provenance:** `docs/inherited-capabilities.md`, `docs/ncode-architecture.md`, and OpenSpec planning artifacts retain factual Zot history.

## Remaining implementation tasks (exact unchecked lines)

- [ ] Characterize before the mechanical move without manufacturing RED: capture `go list ./...`, `go test ./...`, `go mod tidy` dry diff, and `git grep -nI -F 'github.com/patriceckhart/zot' -- '*.go' go.mod examples`; identify all repository-owned import sites, `go.mod`, `docs.go`, `cmd/zot`, `examples/sdk`, `examples/extensions/mcp-bridge/go.mod`, and `examples/extensions/todo/go.mod`. <!-- sdd-owner: implementation -->
- [ ] Atomically change `go.mod` and every repository-owned Go/nested-example import/module declaration from `github.com/patriceckhart/zot` to `github.com/nlf/ncode`; remove the absolute `/Users/pat/Developer/zot` nested replace and do not add a `replace`, compatibility package, or import alias. <!-- sdd-owner: implementation -->
- [ ] Run post-move evidence: `go mod tidy`, tidy each nested example module, `go list ./...`, `go test ./...`, `make test`, and `git grep -nI -F 'github.com/patriceckhart/zot' -- .` restricted to the inventory’s exact reviewed provenance files; prove no active import or `replace` remains. <!-- sdd-owner: implementation -->
- [ ] Characterize the mechanical symbol surface without manufacturing RED by recording focused `go test` results for `./packages/agent/... ./packages/core/... ./packages/provider/... ./packages/tui/...` and locating `ZotHome`, `ZotDocsDir`, `ZotVersion`, `zotHome`, `zotVersion`, and package/product comments in `packages/agent/{config,build,cli,extensions,ext,extproto,skills,swarm,sdk}`, `packages/{core,provider,tui}`, and `docs.go`. <!-- sdd-owner: implementation -->
- [ ] Rename the symbol groups to `NcodeHome`, `NcodeDocsDir`, `NcodeVersion`, and ncode local identifiers at their declarations and all callers; reword active package comments and inherited-parser comments to lowercase ncode or neutral historical wording, preserving exported API shape except the intentional product-symbol rename. <!-- sdd-owner: implementation -->
- [ ] Run post-move evidence: focused package tests, `go vet ./...`, `make test`, and forbidden-symbol searches for `ZotHome|ZotDocsDir|ZotVersion` outside exact provenance/soon-to-be-dedicated rejection fixtures; inspect imports to preserve the `ExtensionToolSource` and SDK/agent cycle seams. <!-- sdd-owner: implementation -->
- [ ] Characterize the mechanical command/build surface without manufacturing RED: record `make build`, `go build -o /tmp/zot ./cmd/zot`, command help/version output, `Makefile` binary targets, `cmd/zot/{main.go,main_test.go}`, and `docs.go` embed references before the move. <!-- sdd-owner: implementation -->
- [ ] Move `cmd/zot/{main.go,main_test.go}` to `cmd/ncode/`, update `Makefile` and command/test references to `./cmd/ncode` and `bin/ncode`, and rename command/product help/version/error literals to lowercase ncode; leave no `cmd/zot`, `bin/zot`, dispatch branch, or alias target. <!-- sdd-owner: implementation -->
- [ ] Run post-move evidence: `go build -o /tmp/ncode ./cmd/ncode`, isolated-`GOBIN` `go install github.com/nlf/ncode/cmd/ncode`, `make build`, command help/version tests, `make test`, and searches proving `/tmp/zot`, `cmd/zot`, and `bin/zot` are absent from active build inputs. <!-- sdd-owner: implementation -->
- [ ] RED — add failing table-driven tests in `packages/agent/config_home_test.go`, `packages/agent/build_test.go`, `packages/agent/extcmd_test.go`, `packages/agent/extensions/manager_test.go`, `packages/agent/skills/skills_test.go`, `packages/core/session_portable_test.go`, and `examples/extensions/mcp-bridge/*_test.go` for the exact `NCODE_HOME`/XDG/macOS/Windows `%LOCALAPPDATA%`/fallback `.ncode` order, `.ncode/{extensions,skills,mcp.json}`, `.ncodesession`, no `.zot` scan, and host/MCP equivalence; run focused tests and `make test` expecting the new assertions to fail. <!-- sdd-owner: implementation -->
- [ ] GREEN — implement the exact `NcodeHome` order in `packages/agent/config.go`; update `packages/agent/{extcmd.go,extensions/manager.go,skills/skills.go}`, `packages/core/session_portable.go`, relevant `packages/agent/modes/*` session help/import UI, and `examples/extensions/mcp-bridge/config.go` to use only ncode roots/project paths and `.ncodesession`; preserve non-branded basenames and make the MCP bridge use `%LOCALAPPDATA%`, not `%APPDATA%`; run focused tests and `make test`. <!-- sdd-owner: implementation -->
- [ ] TRIANGULATE — add `legacy_zot_home_test.go`, `legacy_zot_project_path_test.go`, and `legacy_zot_portable_session_test.go` at the owning package locations with coexistence sentinels/snapshots proving `$ZOT_HOME`, OS Zot defaults, `.zot`, and `.zotsession` are ignored and unchanged; add filesystem timing coverage for path resolution, eager ignored-error docs installation in `Resolve`, and existing lazy triggers, then run focused tests and `make test`. <!-- sdd-owner: implementation -->
- [ ] REFACTOR — consolidate only duplicated test fixtures/path helpers while retaining independent host and MCP assertions; rerun state/project/portable focused tests, `go test ./examples/extensions/mcp-bridge/...`, the no-legacy-path search, and `make test`. <!-- sdd-owner: implementation -->
- [ ] RED — add failing table-driven tests near the owning readers (`packages/agent/{config,build,args,rpc,swarm_agent}_test.go`, `packages/agent/modes/*_test.go`, and swarm tests) for every inventory mapping: `NCODE_HOME`, UI/image/theme, browser/debug, skills, API-key helpers, `NCODE_RPC_TOKEN`, and all `NCODE_SWARM_*`; assert ncode-only success, legacy-only non-use, and ncode-wins conflicts, then run focused tests and `make test` expecting those assertions to fail. <!-- sdd-owner: implementation -->
- [ ] GREEN — rename each `os.Getenv`/`LookupEnv` product control discovered by `git grep -nI -E 'ZOTCORE_|ZOT_[A-Z0-9_]+' -- '*.go'` according to the 23-entry inventory table, with `ZOTCORE_RPC_TOKEN` becoming exactly `NCODE_RPC_TOKEN`; delete the already-removed consent control rather than mapping it, preserve provider/OS controls, and run focused tests plus `make test`. <!-- sdd-owner: implementation -->
- [ ] TRIANGULATE — add dedicated `legacy_zot_env_test.go` fixtures for legacy-only and conflicting values across state, rendering, consent removal, API-key helpers, RPC authorization, and swarm metadata; prove old values neither configure nor authorize behavior, then run focused tests and `make test`. <!-- sdd-owner: implementation -->
- [ ] REFACTOR — table-drive repeated environment setup without hiding the 23-name coverage; rerun `git grep -nI -E 'ZOTCORE_|ZOT_[A-Z0-9_]+' -- .`, focused tests, and `make test`. <!-- sdd-owner: implementation -->
- [ ] RED — add failing RPC tests in `packages/agent/rpc_reasoning_test.go` and adjacent RPC test targets for token-unset prompt-first, `NCODE_RPC_TOKEN` first-frame hello success/failure, `ZOTCORE_RPC_TOKEN` ignored, and unchanged v1 neutral prompt/response/event/optional-hello behavior; run focused tests and `make test` expecting new cases to fail. <!-- sdd-owner: implementation -->
- [ ] GREEN — update `packages/agent/rpc.go`, command invocation/help, `docs/rpc.md`, and `examples/rpc/{shell,go,python,node}` (renaming `zot-client.js`/`zot_client.py` to ncode names) to invoke `ncode rpc` and use `NCODE_RPC_TOKEN`; do not add RPC v2, product fields, mandatory hello, legacy token read, or frame-shape change; run focused tests, client syntax/build checks, and `make test`. <!-- sdd-owner: implementation -->
- [ ] TRIANGULATE — add dedicated legacy RPC input coverage proving `ZOTCORE_RPC_TOKEN` alone neither enables nor gates RPC and an unknown explicit Zot-branded extra field has no authorization/configuration meaning; use fakes/local processes only, then run neutral-frame regressions and `make test`. <!-- sdd-owner: implementation -->
- [ ] REFACTOR — share RPC test transport/setup helpers only if v1 prompt-first and token-required paths remain explicit; rerun RPC tests, all reference-client syntax/build checks, forbidden RPC identity searches, and `make test`. <!-- sdd-owner: implementation -->
- [ ] RED — add failing acknowledgement and lifecycle tests in `packages/agent/extproto/*_test.go`, `packages/agent/ext/ext_test.go`, `packages/agent/extensions/{manager,intercept,tool}_test.go`, and `packages/agent/sdk/*_test.go` requiring `product:"ncode"`, `protocol_version:2`, `ncode_version`, `HostInfo.NcodeVersion`, old `zot_version` acknowledgement rejection, and unchanged neutral hello/frame plus idle auto-ready behavior; run focused tests and `make test` expecting failures. <!-- sdd-owner: implementation -->
- [ ] GREEN — change `packages/agent/extproto/extproto.go` `ProtocolVersion` to 2 and `HelloAckFromHost`; update `packages/agent/extensions/manager.go`, `packages/agent/ext/ext.go`, and `packages/agent/sdk` to emit/validate the ncode acknowledgement and rename the SDK field; retain `ExtensionToolSource`, readiness channels/timeouts, neutral manifest/frame names, and auto-ready logic while changing only its `[zot]` diagnostic to `[ncode]`; update `docs/extensions.md`, `packages/agent/skills/builtin/write-zot-extension/`, and raw extension examples together, then run focused tests and `make test`. <!-- sdd-owner: implementation -->
- [ ] TRIANGULATE — rename the authoring skill directory to `write-ncode-extension`, add dedicated legacy acknowledgement fixtures, build/test `examples/extensions/mcp-bridge` and `examples/extensions/todo` from their own modules, and verify an idle extension still becomes ready without a ready frame; run all focused tests and `make test`. <!-- sdd-owner: implementation -->
- [ ] REFACTOR — remove transitional field names/test duplication without accepting a dual ack or weakening old-ack rejection; rerun extproto/manager/ext/SDK tests, nested extension module builds, protocol searches, and `make test`. <!-- sdd-owner: implementation -->
- [ ] RED — add failing focused tests in `packages/agent/{swarm_agent,extupdate,updatecmd}_test.go`, `packages/agent/swarm/{runner,inbox,socketpath,event}_test.go`, `packages/agent/tools/*_test.go`, `packages/agent/modes/timeline_view_test.go`, `packages/provider/{gemini,provider}_test.go`, and `packages/tui/*_test.go` for `NCODE_SWARM_*`, `ncode-swarm-*`, ncode temp/log/request names, and ncode product headers; run focused tests and `make test` expecting failures. <!-- sdd-owner: implementation -->
- [ ] GREEN — update `packages/agent/{swarm_agent.go,extupdate.go,updatecmd.go,build.go}`, `packages/agent/swarm/{socketpath.go,runner.go}`, `packages/agent/tools/bash.go`, `packages/agent/modes/timeline_view.go`, `packages/provider/{gemini.go,openai_codex.go}`, and `packages/tui/clipboard_darwin.go` to the inventory’s exact `NCODE_*`/`ncode-*` values; retain `os.Executable` child location, neutral swarm event/control shapes, and provider-mandated external fingerprints; run focused tests and `make test`. <!-- sdd-owner: implementation -->
- [ ] TRIANGULATE — add `legacy_zot_swarm_test.go` and provider/internal negative cases proving Zot swarm metadata cannot supply credentials/event logs or create `zot-swarm-*`, while retained child execution, JSONL events, and neutral controls work; use local stubs/`httptest`, never live providers, then run focused tests and `make test`. <!-- sdd-owner: implementation -->
- [ ] REFACTOR — centralize only repeated ncode prefix assertions while preserving exact output/path assertions; rerun swarm/provider/temp focused tests, composed-name searches, and `make test`. <!-- sdd-owner: implementation -->
- [ ] RED — add failing static/unit tests or snapshots for `.goreleaser.yaml`, `install.sh`, `install.ps1`, `packages/agent/{update,updatecmd,changelog}_test.go`, and `.github/workflows/release.yml` requiring `nlf/ncode`, `ncode_<version>_<os>_<arch>`, `ncode[.exe]`, `NCODE_VERSION`, `NCODE_PREFIX`, ncode checksums/install URLs, and no legacy endpoint; run local focused tests and `make test` expecting failures. <!-- sdd-owner: implementation -->
- [ ] GREEN — update `.goreleaser.yaml`, `install.sh`, `install.ps1`, `.github/workflows/release.yml`, and `packages/agent/{update.go,updatecmd.go,changelog.go}` to ncode-only owner/repository/assets/extraction/user-agent/temp identity; delete `examples/extensions/todo/zot-todo-extension` and `examples/rpc/python/__pycache__/zot_client.cpython-314.pyc` without replacements, preserve neutral `checksums.txt`, and add the required clean-break installer/release notice; run tests and `make test`. <!-- sdd-owner: implementation -->
- [ ] TRIANGULATE — validate GoReleaser config/snapshot and archive contents when tooling is available, installer shell syntax with mocked HTTP/filesystems, PowerShell parser/static checks, updater `httptest` responses, and checksum/asset-name agreement; assert no release call reaches a network provider or live GitHub release, then run `make test`. <!-- sdd-owner: implementation -->
- [ ] REFACTOR — derive repeated asset-name expectations from one local test helper/config fixture without adding an updater fallback; rerun distribution checks, `make build`, focused updater tests, forbidden asset/endpoint searches, and `make test`. <!-- sdd-owner: implementation -->
- [ ] Characterize the mechanical communication surface without manufacturing RED: record link/command checks and nested-example builds for `README.md`, `CONTRIBUTING.md`, `AGENTS.md`, `docs/{providers,extensions,rpc,skills,themes}.md`, `.github/ISSUE_TEMPLATE/{bug_report,feature_request}.yml`, `examples/**`, `packages/agent/skills/builtin/write-zot-themes/SKILL.md`, and active help/error/log expectations; distinguish exact provenance files from active materials. <!-- sdd-owner: implementation -->
- [ ] Rename active product prose, commands, imports, state/env/SDK/extension references, test names/expectations, manifests, and example consumers to lowercase ncode; move `write-zot-themes` to `write-ncode-themes`; ensure README Install, installer/install output, and README first-run/authentication prominently state the clean break and no Zot credentials/settings/sessions/caches/extensions/SDK-RPC-swarm reuse, with no migration/prompt instruction. <!-- sdd-owner: implementation -->
- [ ] Run post-edit evidence: nested example module tests/builds, example shell/Python/Node syntax checks, markdown link/command searches, installer/help snippet checks, `make test`, and a forbidden-active-identity scan; retain Zot wording only in the exact provenance manifest or dedicated `legacy_zot_*` rejection files. <!-- sdd-owner: implementation -->
- [ ] Add or complete focused no-live-provider retained-capability coverage for provider/auth resolution, print/stream/JSON/RPC v1, sessions, direct permissions/tools, swarm, extensions including idle auto-ready v2 ack, skills, themes, updater via `httptest`, and retained Telegram tests; run the affected packages, `go vet ./...`, and `make test`. <!-- sdd-owner: implementation -->
- [ ] Update `openspec/changes/clean-break-ncode-identity/identity-inventory.md` before the final search to add this exact `tasks.md` path to the reviewed provenance manifest only if it contains legacy terms solely as historical planning/rejection context; then verify every provenance line and every `legacy_zot_*` line is factual or asserts rejection/non-use, not active support. <!-- sdd-owner: implementation -->
- [ ] Run the exact final allowlist gates from `identity-inventory.md`: construct `audit_pathspecs` with each exact provenance file plus `openspec/changes/clean-break-ncode-identity/tasks.md` when admitted and only `:(exclude,glob)**/legacy_zot_*`; require zero output from `git grep -nI -i -e zot -- "${audit_pathspecs[@]}"`, old-module, `ZOTCORE_|ZOT_[A-Z0-9_]+`, `.zot(session)?`, and `zot[-_][[:alnum:]_-]+` searches, and from `git ls-files | grep -i zot | grep -vE '(^|/)legacy_zot_[^/]*$'`. <!-- sdd-owner: implementation -->
- [ ] Run the exact workable-file and mandatory-line reviews from `identity-inventory.md`: build `/tmp/ncode-zot-provenance-files.txt`, `/tmp/ncode-zot-rejection-files.txt`, and `/tmp/ncode-zot-allowed-files.txt`; require `comm -23` against `git grep -Il -i -e zot -- .` to emit no paths; emit and review every line into `/tmp/ncode-zot-provenance-lines.txt` and `/tmp/ncode-zot-rejection-lines.txt`; then run `go mod tidy`, `go list ./...`, `go vet ./...`, `make build`, `make test`, nested example builds, installer/package checks, and the inventory reproduction commands against planning SHA `18325b75cc89c75b5f4842924cb377aa5bef5c4b`. <!-- sdd-owner: implementation -->

## Deferred parent lifecycle action (exact unchecked line)

- [ ] Start or reuse bounded review after each applied work-unit commit, reviewing its stated boundary, focused evidence, rollback isolation, and changed-line count before the next dependent unit proceeds. <!-- sdd-owner: parent -->

## Risks and next boundary

- WU 1 is intentionally above 400 lines under its narrow approved deletion-heavy exception; that approval does not apply to WU 2 or later.
- The worktree contains a pre-existing modification to `docs/ncode-architecture.md`; parent delivery must keep ownership clear and avoid accidentally including unrelated changes.
- The repository is on `main` with no branch or commit created. Parent owns commit/review/PR lifecycle.
- Do not begin WU 2 until the parent completes the WU 1 lifecycle boundary and explicitly delegates the next stacked slice.
