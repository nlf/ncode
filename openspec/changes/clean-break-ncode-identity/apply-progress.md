# Apply progress: clean-break-ncode-identity

Updated: 2026-08-26T20:59:29Z

## Attempt summary

- **Status:** WU 1 is merged and WU 2–WU 7 are complete; the broader change remains in progress, with WU 8 next and unstarted.
- **Assigned boundary:** exactly WU 7 — RPC identity edges and reference clients — complete. WU 8 was not started.
- **Delivery:** WU 2–WU 12 remain on the single long-lived branch `feat/clean-break-ncode-identity-02`; no intermediate PR or branch change was made.
- **Review budget:** WU 7 implementation and tests use `402` rename-aware changed lines, below the native 1,300-line maximum; the accepted one-final-PR size exception and no-intermediate-review strategy cover the 400-line forecast.
- **Commits/branches/PRs:** no commit, push, PR, merge, branch switch, review, settlement, or WU 8 work occurred. Parent owns the revertible WU 7 commit and attempt settlement.
- **Verification:** genuine RED → GREEN → TRIANGULATE → REFACTOR evidence, focused RPC tests, all four reference-client checks, exact forbidden identity/protocol searches, `git diff --check`, and final `make test` passed.
- **Native authority:** parent supplied proceed token `sha256:b922f6bb7768cbf8d3481e6b3f4c5a15fc56ad88775831cd7079c5299d9b4cc9`; the native CLI was unavailable locally, no extra acquire was possible, and settlement is intentionally deferred to the parent.

## Structured status consumed/produced

```yaml
schemaName: spec-driven
changeName: clean-break-ncode-identity
artifactStore: both
planningHome:
  root: /Users/nlf/Projects/nlf/ncode/openspec
  changesDir: /Users/nlf/Projects/nlf/ncode/openspec/changes
changeRoot: /Users/nlf/Projects/nlf/ncode/openspec/changes/clean-break-ncode-identity
artifactPaths:
  proposal: [openspec/changes/clean-break-ncode-identity/proposal.md]
  specs: [openspec/changes/clean-break-ncode-identity/specs/ncode-identity/spec.md]
  design: [openspec/changes/clean-break-ncode-identity/design.md]
  tasks: [openspec/changes/clean-break-ncode-identity/tasks.md]
  applyProgress: [openspec/changes/clean-break-ncode-identity/apply-progress.md]
artifacts:
  proposal: done
  specs: done
  design: done
  tasks: done
  applyProgress: done
taskProgress:
  total: 44
  complete: 25
  remaining: 19
deferredParentActions:
  total: 2
  complete: 2
  remaining: 0
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
    - native gentle-ai status/attempt CLI unavailable; authoritative OpenSpec state reconstructed from the installed status contract
nextRecommended: parent-lifecycle
isNonAuthoritative: false
```

The active change was explicit and unambiguous. All edits stayed under the allowed repository root. Ownership validation found no malformed markers. The accepted single-feature-branch size exception resolved the workload gate; WU 8 remains outside this attempt.

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

## WU 2 apply evidence

### Completed implementation tasks and persisted checkboxes

- [x] Characterization/RED evidence — persisted in `tasks.md`.
- [x] Atomic canonical module/import graph move — persisted in `tasks.md`.
- [x] Post-move verification — persisted in `tasks.md`.

### TDD Cycle Evidence

| Task | Test layer | Safety Net / RED | GREEN | TRIANGULATE | REFACTOR |
|---|---|---|---|---|---|
| WU2 characterize | Approval/structural | Baseline `go list ./...` listed 26 packages and `go test ./...` passed. No RED was manufactured. The dry tidy check honestly found pre-existing WU1 residue: one unused direct requirement and two checksum lines. | N/A | Inventory found 274 old-path lines across 160 Go/module/example files and identified every required anchor; `docs.go` was inspected but contained no canonical import. | Generated tidy edits were restored before production edits began. |
| WU2 atomic move | Structural/module integration | Characterization above is the approved RED-phase evidence for this mechanical move. | Replaced 281 occurrences across 163 active files; root `go list ./...` passed for 26 packages under the new module. | Root plus both nested module declarations were checked; all tracked `replace` directives and active old imports were absent. | No behavior refactor was appropriate; only exact path substitution and nested module cleanup were performed. |
| WU2 post evidence | Package/integration/race | Root and nested module state characterized above. | Root tidy, `go list ./...`, and `go test ./...` passed. | Both nested modules passed tidy and `go test ./...` against the current root through a transient untracked validation overlay. | `make test` (`go test -race ./...`) and `git diff --check` passed; construction spine and cycle seam declarations remain. |

### Exact commands and outcomes

- `go list ./...` before move — PASS; 26 packages under the old module identity.
- `go test ./...` before move — PASS.
- `go mod tidy` dry-diff characterization — NON-NO-OP: would remove unused `github.com/klauspost/compress v1.19.0` and its two checksums; changes were restored before the move.
- Old-path inventory over `*.go`, root `go.mod`, and `examples` — 274 lines across 160 files.
- Root `go mod tidy` after move — PASS; accepted only the characterized unused dependency/checksum cleanup.
- Plain nested `go mod tidy` — expected bootstrap failure because every currently published `github.com/nlf/ncode` tag still declares the inherited module path.
- `(cd examples/extensions/<module> && go mod edit -replace=github.com/nlf/ncode=../../.. && go mod tidy && go test ./... && go mod edit -dropreplace=github.com/nlf/ncode)` for `mcp-bridge` and `todo` — PASS; the overlay was transient and no tracked `replace` remains.
- `go list ./...` after move — PASS; 26 packages under `github.com/nlf/ncode`.
- `go test ./...` after move — PASS.
- `make test` — PASS (`go test -race ./...`).
- Active Go/module/example old-path search — zero matches.
- Absolute legacy nested path search restricted to module files — zero matches.
- `git grep -nI -E '^replace[[:space:]]' -- go.mod '**/go.mod'` — zero matches.
- `git diff --check` — PASS.
- Spine/seam inspection — `Resolve`, `Resolved.NewClient`, `Resolved.NewAgent`, `composeHeadlessAgent`, and `ExtensionToolSource` remain declared at their existing boundaries.

### Changed-line and path scope

- **Implementation:** 164 files, 281 additions, 288 deletions, **569 changed lines**.
- **SDD artifacts:** `tasks.md`, `apply-progress.md`, and `identity-inventory.md`; `tasks.md` includes the parent’s pre-existing one-line lifecycle update.
- **Total worktree changed lines:** 534 additions + 335 deletions = 869 changed lines.
- **Implementation paths (exact):**

- `README.md`
- `cmd/zot/main.go`
- `docs/extensions.md`
- `examples/extensions/approve/main.go`
- `examples/extensions/guard/main.go`
- `examples/extensions/hello/main.go`
- `examples/extensions/mcp-bridge/bridge.go`
- `examples/extensions/mcp-bridge/deferred_tools_test.go`
- `examples/extensions/mcp-bridge/go.mod`
- `examples/extensions/mcp-bridge/main.go`
- `examples/extensions/secret/main.go`
- `examples/extensions/todo/go.mod`
- `examples/extensions/todo/main.go`
- `examples/extensions/weather/main.go`
- `examples/rpc/go/main.go`
- `examples/sdk/main.go`
- `go.mod`
- `go.sum`
- `packages/agent/args.go`
- `packages/agent/botcmd.go`
- `packages/agent/botspec.go`
- `packages/agent/build.go`
- `packages/agent/build_test.go`
- `packages/agent/cli.go`
- `packages/agent/cli_headless_test.go`
- `packages/agent/cli_session_test.go`
- `packages/agent/config.go`
- `packages/agent/config_command_test.go`
- `packages/agent/config_gcp_test.go`
- `packages/agent/confirmation_events.go`
- `packages/agent/confirmation_events_test.go`
- `packages/agent/ext/ext.go`
- `packages/agent/ext/ext_test.go`
- `packages/agent/extcmd.go`
- `packages/agent/extcmd_test.go`
- `packages/agent/extensions/events.go`
- `packages/agent/extensions/manager.go`
- `packages/agent/extensions/manager_test.go`
- `packages/agent/extensions/tool.go`
- `packages/agent/headless.go`
- `packages/agent/instruction_context.go`
- `packages/agent/instruction_context_test.go`
- `packages/agent/modelsync.go`
- `packages/agent/modelsync_test.go`
- `packages/agent/modes/auto_swarm_test.go`
- `packages/agent/modes/bot/adapter.go`
- `packages/agent/modes/bot/runner.go`
- `packages/agent/modes/bot/runner_test.go`
- `packages/agent/modes/bot/status.go`
- `packages/agent/modes/btw_dialog.go`
- `packages/agent/modes/btw_dialog_test.go`
- `packages/agent/modes/changelog_dialog.go`
- `packages/agent/modes/changelog_dialog_test.go`
- `packages/agent/modes/clipboard_images.go`
- `packages/agent/modes/clipboard_images_test.go`
- `packages/agent/modes/clipboard_inputs_test.go`
- `packages/agent/modes/clipboard_paste.go`
- `packages/agent/modes/clipboard_paste_test.go`
- `packages/agent/modes/compact_queue_test.go`
- `packages/agent/modes/confirm_dialog.go`
- `packages/agent/modes/confirm_dialog_test.go`
- `packages/agent/modes/dialog_frame.go`
- `packages/agent/modes/dialog_text.go`
- `packages/agent/modes/dialog_text_test.go`
- `packages/agent/modes/ext_notes_test.go`
- `packages/agent/modes/ext_panel_dialog.go`
- `packages/agent/modes/file_suggest.go`
- `packages/agent/modes/help.go`
- `packages/agent/modes/help_test.go`
- `packages/agent/modes/input_history_test.go`
- `packages/agent/modes/interactive.go`
- `packages/agent/modes/interactive_model_test.go`
- `packages/agent/modes/interactive_terminal_test.go`
- `packages/agent/modes/jail_setting_test.go`
- `packages/agent/modes/json.go`
- `packages/agent/modes/jump_dialog.go`
- `packages/agent/modes/llama_dialog.go`
- `packages/agent/modes/llama_dialog_test.go`
- `packages/agent/modes/login_dialog.go`
- `packages/agent/modes/login_dialog_test.go`
- `packages/agent/modes/logout_dialog.go`
- `packages/agent/modes/logout_dialog_test.go`
- `packages/agent/modes/model_dialog.go`
- `packages/agent/modes/print.go`
- `packages/agent/modes/print_test.go`
- `packages/agent/modes/reload_ext.go`
- `packages/agent/modes/reload_ext_test.go`
- `packages/agent/modes/rescue_dialog.go`
- `packages/agent/modes/session_agent_root_test.go`
- `packages/agent/modes/session_dialog.go`
- `packages/agent/modes/session_ops_dialog.go`
- `packages/agent/modes/session_tree_dialog.go`
- `packages/agent/modes/settings_dialog.go`
- `packages/agent/modes/shell_escape_test.go`
- `packages/agent/modes/skill_command.go`
- `packages/agent/modes/skill_command_test.go`
- `packages/agent/modes/skills_dialog.go`
- `packages/agent/modes/slash_suggest.go`
- `packages/agent/modes/slash_suggest_test.go`
- `packages/agent/modes/spinner.go`
- `packages/agent/modes/stream.go`
- `packages/agent/modes/stream_test.go`
- `packages/agent/modes/swarm_dialog.go`
- `packages/agent/modes/swarm_dialog_test.go`
- `packages/agent/modes/swarm_slash.go`
- `packages/agent/modes/swarm_slash_test.go`
- `packages/agent/modes/telegram/adapter.go`
- `packages/agent/modes/telegram/api_test.go`
- `packages/agent/modes/telegram/bridge.go`
- `packages/agent/modes/telegram/commands.go`
- `packages/agent/modes/telegram/daemon.go`
- `packages/agent/modes/telegram/status.go`
- `packages/agent/modes/telegram/status_test.go`
- `packages/agent/modes/telegram_dialog.go`
- `packages/agent/modes/timeline_view.go`
- `packages/agent/modes/timeline_view_test.go`
- `packages/agent/modes/update_banner.go`
- `packages/agent/modes/welcome.go`
- `packages/agent/print_stats_test.go`
- `packages/agent/rpc.go`
- `packages/agent/rpc_reasoning_test.go`
- `packages/agent/sdk/reasoning_test.go`
- `packages/agent/sdk/sdk.go`
- `packages/agent/sdk/types.go`
- `packages/agent/sessionscmd.go`
- `packages/agent/sessionscmd_test.go`
- `packages/agent/settings_store.go`
- `packages/agent/skills/builtin/write-zot-extension/SKILL.md`
- `packages/agent/skills/tool.go`
- `packages/agent/swarm_agent.go`
- `packages/agent/swarm_agent_test.go`
- `packages/agent/tools/bash.go`
- `packages/agent/tools/edit.go`
- `packages/agent/tools/read.go`
- `packages/agent/tools/swarm_spawn.go`
- `packages/agent/tools/swarm_spawn_test.go`
- `packages/agent/tools/telegram_send.go`
- `packages/agent/tools/tools_test.go`
- `packages/agent/tools/write.go`
- `packages/agent/updatecmd.go`
- `packages/core/agent.go`
- `packages/core/agent_context_test.go`
- `packages/core/agent_retry_test.go`
- `packages/core/compact.go`
- `packages/core/core_test.go`
- `packages/core/cost.go`
- `packages/core/events.go`
- `packages/core/intercept_test.go`
- `packages/core/queue_test.go`
- `packages/core/session.go`
- `packages/core/session_deferred_tools_test.go`
- `packages/core/session_portable.go`
- `packages/core/session_portable_test.go`
- `packages/core/session_prune_test.go`
- `packages/core/session_repair_test.go`
- `packages/core/tool.go`
- `packages/provider/auth/callback.go`
- `packages/provider/google_vertex.go`
- `packages/tui/statusbar_test.go`
- `packages/tui/view.go`
- `packages/tui/view_compact_mode_test.go`
- `packages/tui/view_compact_user_test.go`
- `packages/tui/view_flat_tools_test.go`
- `packages/tui/view_tool_overlay_test.go`

### Deviations and discoveries

- Published `github.com/nlf/ncode` tags through v0.3.50 still declare the inherited module path, so a plain nested tidy cannot resolve the new path before this canonical move is published. Transient validation overlays proved both nested examples against the current root without persisting a `replace`.
- The root tidy characterization was not a no-op because WU1 left `klauspost/compress` unused after Zotfile deletion. The cleanup was restored during characterization and then accepted by the required post-move tidy.
- Canonical-path references in active README install/repository lines, SDK docs, extension docs, and the extension-authoring skill were updated with WU2 so no active old canonical path remained; broader product prose remains assigned to later WUs.
- `tasks.md` was admitted early to the exact provenance manifest because its checked characterization/audit commands must retain legacy literals; the final WU12 line review remains unchecked and deferred.
- No live providers or credentials were used.

## WU 3 apply evidence
- [x] Characterization, symbol/comment rename, and post-rename verification remain persisted at `tasks.md:85-87`; the surgical correction preserved every task checkbox and WU3-complete/WU4-next marker.
- The authorized corrections cumulatively reworded 20 stale WU3 comments: the prior eight at `sdk.go:94`, `provider.go:70`, `skills.go:57,166,195`, `builtin.go:10`, and `session.go:300,608`, plus the corrective 12 Zot-bearing lines in `ignore/gitignore.go`, `provider/{llamacpp,reasoning}.go`, `tui/render.go`, `agent/build.go`, and `core/session.go`. No behavior, executable, UI, state, environment, path, protocol, task checkbox, or docs literal changed.
### TDD Cycle Evidence

| Task | Layer | Safety Net / RED | GREEN | TRIANGULATE | REFACTOR |
|---|---|---|---|---|---|
| WU3 characterize | Approval/structural | Baseline focused package tests passed; no RED was manufactured for the mechanical rename. | N/A | 190 symbol/local occurrences were classified against later work-unit boundaries. | No edits during characterization. |
| WU3 rename | Structural/API | Existing tests were the approval safety net. | Five symbol/local groups were renamed; focused tests passed. | Uncached direct-package consumers passed. | Go formatting, package docs, and inherited-parser comments were cleaned. |
| WU3 comment corrections | REFACTOR/remediation | Prior core, SDK, skills, provider, and full WU3 package safety nets were green. | N/A; these are comment-only REFACTOR corrections, no behavior changed, and no RED was manufactured. | All 20 corrected comments use lowercase `ncode` or neutral history and contain zero case-insensitive `zot` matches. | Focused package tests, `go vet ./...`, `make test`, comprehensive comment inventory, and `git diff --check` passed. |

### Exact verification and searches

- `go test ./packages/agent/sdk ./packages/agent/skills ./packages/provider -count=1` and `go test ./packages/agent/... ./packages/core/... ./packages/provider/... ./packages/tui/... -count=1` before and after correction — PASS.
- `go test -count=1 ./packages/agent ./packages/agent/ext ./packages/agent/extensions ./packages/agent/modes ./packages/agent/modes/telegram ./packages/agent/sdk ./packages/agent/skills ./packages/agent/tools ./packages/core ./packages/provider ./packages/provider/auth ./packages/tui` — PASS.
- `go vet ./...`; `make test` (`go test -race ./...`); `git diff --check` — PASS; no live providers or credentials were used.
- Exact-line scan over `sdk.go:94`, `provider.go:70`, `skills.go:57,166,195`, `builtin.go:10`, and `session.go:300,608` — six lowercase `ncode` matches, two neutral historical comments, and zero case-insensitive `zot` matches; stale-phrase searches include the two removed parser phrases and produce zero matches.
- `rg -n --glob '*.go' '(ZotHome|ZotDocsDir|ZotVersion|\bzotHome\b|\bzotVersion\b)' .` and the active-Go package-declaration scan for Zot product wording — zero matches; these remain the WU3 symbol/package gates.
- Comprehensive inventory scanned all 350 tracked Go files and found 273 remaining Zot-bearing line comments after correction; every match is classified below, with zero remaining WU3-owned package/exported/product/parser blockers.
- The WU3 symbol gate remains zero-match; `NcodeVersion` intentionally retains `json:"zot_version"` until WU8.

### Final comprehensive-comment corrective REFACTOR

- Attempt `sha256:39875b93e6f8475e355d832717e5de698aff782c10303674aaf27e5c4147cffa` (`WU3-final-comments`) received one corrective rerun. It changed comments only in the six requested Go files plus the cumulative progress artifact; tasks/checkmarks stayed byte-for-byte unchanged and WU4 did not start.
- Strict TDD disposition: comment-only REFACTOR with existing behavior preserved; no RED was manufactured, no tests were added, and the prior green safety net remained the approval baseline.
- WU4 command/build (15): `cmd/zot/main.go:1,15,45`; `packages/agent/botcmd.go:31,73,113,114,279,314,367`; `packages/agent/cli.go:210,216,740,1012`; `packages/tui/clipboard_text_linux.go:11`.
- WU5 state/path (42): `examples/extensions/guard/main.go:18`; `examples/extensions/hello/main.go:9`; `examples/extensions/mcp-bridge/config.go:6,7`; `examples/extensions/mcp-bridge/main.go:8,9`; `examples/extensions/weather/main.go:13`; `packages/agent/args.go:83,89`; `packages/agent/build.go:623,675,694`; `packages/agent/cli.go:1332`; `packages/agent/config.go:123`; `packages/agent/extcmd.go:463`; `packages/agent/extupdate.go:16`; `packages/agent/extupdate_test.go:69,98`; `packages/agent/modelsync.go:37`; `packages/agent/modes/interactive.go:6716`; `packages/agent/sdk/sdk.go:38`; `packages/agent/skills/skills.go:15,16`; `packages/agent/swarm/runner.go:51,134`; `packages/agent/swarm/socketpath.go:32`; `packages/agent/systemprompt.go:49`; `packages/agent/update.go:39`; `packages/core/session_portable.go:20,25,27,67,151,571`; `packages/core/session_portable_test.go:86,100`; `packages/provider/catalog_builtin.go:9`; `packages/provider/usermodels.go:12`; `packages/tui/theme_loader.go:14,25,175,261`.
- WU6 env (26): `packages/agent/build.go:403`; `packages/agent/build_test.go:296`; `packages/agent/config.go:36,42,154,172`; `packages/agent/modes/interactive.go:89,94`; `packages/agent/rpc.go:41`; `packages/agent/swarm/swarm.go:220`; `packages/agent/swarm_agent.go:28`; `packages/provider/anthropic.go:535`; `packages/tui/image.go:31,34,35,36,37,38,61,80,244`; `packages/tui/render.go:34,336`; `packages/tui/view.go:2219,2225,2251`.
- WU7 RPC (7): `examples/rpc/go/main.go:1`; `packages/agent/cli.go:234`; `packages/agent/modes/json.go:38`; `packages/agent/rpc.go:177,185`; `packages/agent/sdk/sdk.go:22`; `packages/agent/sdk/types.go:13`.
- WU8 protocol/SDK/examples (49): `examples/extensions/approve/main.go:19,21`; `examples/extensions/guard/main.go:19`; `examples/extensions/hello/main.go:1`; `examples/extensions/mcp-bridge/bridge.go:1,3,8,51,58,74,85,86,101,411,452,464,495,526`; `examples/extensions/mcp-bridge/main.go:1,5,33,50,108`; `examples/extensions/mcp-bridge/server.go:194`; `examples/extensions/secret/main.go:23,25`; `examples/extensions/weather/main.go:14`; `examples/sdk/main.go:1,2`; `packages/agent/extcmd.go:20,510,560`; `packages/agent/extcmd_test.go:13`; `packages/agent/extensions/manager.go:36,50,55,140,271,312,313,452,482,950,1026,1264`; `packages/agent/extensions/proc_unix.go:10,14,15`; `packages/agent/skills/skills_test.go:76`.
- WU9 runtime/swarm/temp/provider (59): `packages/agent/build.go:476,485`; `packages/agent/build_test.go:133`; `packages/agent/cli.go:1113,1276`; `packages/agent/modes/swarm_dialog.go:120,123`; `packages/agent/swarm/agent.go:24,29,55`; `packages/agent/swarm/event_test.go:83`; `packages/agent/swarm/inbox.go:17,158`; `packages/agent/swarm/inbox_test.go:108`; `packages/agent/swarm/persist.go:8,52,316,340`; `packages/agent/swarm/persist_test.go:19,66,778,822`; `packages/agent/swarm/runner.go:17,19,21,34,40,61`; `packages/agent/swarm/runner_test.go:23`; `packages/agent/swarm/socketpath.go:113`; `packages/agent/swarm/swarm.go:57,62,143,258`; `packages/agent/swarm/swarm_test.go:66`; `packages/agent/swarm/testdata/cmd/stubchild/main.go:2,17`; `packages/agent/swarm_agent.go:20,78,235`; `packages/agent/swarm_agent_test.go:19`; `packages/provider/amazon_bedrock.go:610`; `packages/provider/anthropic_image.go:20`; `packages/provider/auth/oauth.go:23,81`; `packages/provider/extra_models.go:6`; `packages/provider/extra_providers.go:337`; `packages/provider/gemini.go:28,57,409`; `packages/provider/gemini_test.go:540`; `packages/provider/github_copilot.go:8,22`; `packages/provider/models.go:63`; `packages/provider/openai.go:257,342`; `packages/provider/openai_codex.go:24,78,405`.
- WU10 release/updater (9): `packages/agent/extupdate.go:199`; `packages/agent/modes/changelog_dialog.go:11`; `packages/agent/update.go:18`; `packages/agent/updatecmd.go:19,34,110,205,329,389`.
- WU11 UI/help/examples/prose (65): `packages/agent/args.go:78,319`; `packages/agent/build.go:380,693`; `packages/agent/config.go:50,192`; `packages/agent/modelsync.go:51,52,84`; `packages/agent/modes/file_suggest.go:66,132`; `packages/agent/modes/interactive.go:127,150,151,200,671,684,686,1044,1369,1735,2375,3794,4405,5644,5722,6224`; `packages/agent/modes/telegram/bridge.go:40,56,59,164,165,167,171,186`; `packages/agent/modes/telegram/daemon.go:14`; `packages/agent/modes/welcome.go:6,9,10`; `packages/agent/systemprompt.go:37`; `packages/core/session_portable.go:157`; `packages/provider/auth/assets/assets.go:12`; `packages/provider/auth/callback.go:126,128,182`; `packages/tui/highlight.go:123`; `packages/tui/image.go:76`; `packages/tui/render.go:65,67,448,504`; `packages/tui/theme.go:9,36,218,237,278,284`; `packages/tui/view.go:196,404,447,740,866,1757,2497,2532`.
- Dedicated legacy rejection fixture (1): `packages/agent/legacy_zotfile_absence_test.go:9`.
- Required verification passed exactly: `go test ./packages/agent/... ./packages/core/... ./packages/provider/... ./packages/tui/... ./packages/ignore/... -count=1`; `go vet ./...`; `make test`; `git diff --check`; and the tracked-Go `git grep` inventory above.
- Exact cumulative attempt scope: 7 files, 43 additions + 27 deletions = **70 changed lines**, rising by 36 from the prior 34-line checkpoint and remaining below the 100-line maximum (`apply-progress.md` 29+13; `build.go` 2+2; `session.go` 3+3; `gitignore.go` 1+1; `llamacpp.go` 1+1; `reasoning.go` 6+6; `render.go` 1+1). Updated cumulative worktree scope: 51 files, 375 additions + 334 deletions = **709 changed lines**.

### Scope and deviations

- Prior surgical correction scope: 5 files, 22 additions + 21 deletions = **43 changed lines** relative to reset candidate `sha256:cba34272f556f3830479f8e98f2da5d22939bea1de89f35388632bedacd84539`. Final parser-comment attempt statistics and updated cumulative scope are recorded after its exact verification evidence below.
- `NcodeVersion` intentionally preserves `json:"zot_version"`; current executable/UI and state/env/path behavior is unchanged; WU4 was not implemented.
- `ExtensionToolSource`, SDK-to-agent direction, and `Resolve` → `Resolved.NewClient` → `Resolved.NewAgent` remain intact; no commit, push, PR, branch switch, or review occurred.

## WU 4 apply evidence

### Completed implementation tasks and persisted checkboxes

- [x] Characterization without manufactured RED — persisted at `tasks.md:92`.
- [x] Command/local-build move and identity update — persisted at `tasks.md:93`.
- [x] Post-move build/install/help/version/full verification — persisted at `tasks.md:94`.

### TDD Cycle Evidence

| Task | Layer | Safety Net / RED | GREEN | TRIANGULATE | REFACTOR |
|---|---|---|---|---|---|
| Characterize | Approval/structural | `make build`, `go build -o /tmp/zot ./cmd/zot`, `go test ./cmd/zot`, and `make test` passed. Baseline help/version advertised `zot`; no RED was manufactured for the approved mechanical move. | N/A | Makefile targets, both command files, and unchanged `docs.go` embed inputs were inventoried. | Temporary baseline outputs were removed after evidence capture. |
| Move/rebrand | Structural + command integration | Approval baseline above. | `cmd/ncode`, `bin/ncode`, central/subcommand help, version output, and local product errors were updated together; focused command/agent/TUI tests passed. | `TestHelpOutputStreams` now covers help on stdout, argument-error help on stderr, absence of Zot command rows, and exact `ncode test` version output. | Command literals were rebranded without changing routers, state/env/path/protocol behavior, or release/updater internals. |
| Post-move | Build/install/race | Characterization above. | Direct build, isolated-GOBIN install, and `make build` passed with only ncode executable outputs. | Both `--version` and `-v`, command help, no-alias/dispatch searches, and local build-input path searches passed. | `git diff --check` and `make test` passed. |

Triangulation was structural for the move itself because there is one command location/output. Observable help/version behavior was triangulated across long/short version flags and success/error help streams.

### Characterization and verification

- Baseline `make build` produced `bin/zot`; `go build -o /tmp/zot ./cmd/zot`, `go test ./cmd/zot -count=1`, and `make test` passed.
- Baseline `/tmp/zot --help` advertised Zot commands; `--version` and `-v` printed `zot 0.3.51-0.20260826183238-5099c80b49c7`.
- Baseline Makefile had `build`, `install`, `run`, and five cross-build outputs rooted at `cmd/zot`/`bin/zot`; command sources were `main.go`/`main_test.go`; `docs.go` embedded `README.md docs/*.md` and installed five mapped docs.
- `go build -o /tmp/ncode ./cmd/ncode` — PASS; help and both version flags identify lowercase ncode.
- Isolated `GOBIN` `go install github.com/nlf/ncode/cmd/ncode` — PASS; installed only `ncode`.
- `make build` — PASS; produced executable `bin/ncode` and no `bin/zot`.
- `go test ./cmd/ncode -run '^TestResolvedVersionFromBuildInfo$' -count=1` — PASS.
- `go test ./packages/agent -run '^TestHelpOutputStreams$' -count=1` — PASS.
- `go test ./packages/agent ./packages/tui -count=1` — PASS.
- `make test` (`go test -race ./...`) — PASS.
- Local build-input and alias searches over `Makefile`, `docs.go`, `cmd/ncode`, and command routers found no `/tmp/zot`, `cmd/zot`, `bin/zot`, Zot target, or Zot dispatch branch; worktree paths `cmd/zot`, `bin/zot`, and `/tmp/zot` were absent.
- `git diff --check` — PASS; temporary binaries and isolated install directories were removed.

### Files and scope

- Moved: `cmd/zot/{main.go,main_test.go}` → `cmd/ncode/`.
- Build/docs: `Makefile`, `README.md`, `CONTRIBUTING.md`.
- Command/help/error code and test: `packages/agent/{args.go,args_test.go,botcmd.go,build.go,cli.go,extcmd.go,modelsync.go,sessionscmd.go}`, `packages/agent/skills/tool.go`, `packages/tui/clipboard_text_linux.go`.
- SDD artifacts: `tasks.md` and this cumulative `apply-progress.md`.
- Exact WU4 scope: **343 changed lines** (203 additions + 140 deletions, rename-aware); conservative no-rename accounting is 559 (311 + 248), also below the native 700-line maximum.

### Deviations and deferred boundaries

- No design behavior deviation occurred. Mechanical characterization replaced manufactured RED as explicitly authorized; one version assertion was added during triangulation.
- README and CONTRIBUTING command/build paths moved now because they are active embedded/local-build references; broader communication remains WU11.
- `.goreleaser.yaml` still contains its WU10-owned Zot release build id/main/binary, and updater/installer/release asset identity remains deferred to WU10. State variables and paths such as `$ZOT_HOME`/`.zot`, RPC and extension protocol fields, and runtime composed names remain assigned to WU5–WU10.
- No commit, push, PR, merge, branch switch, review transaction, native attempt settlement, or WU5 work occurred.

## WU 5 apply evidence

### Completed implementation tasks and persisted checkboxes

- [x] RED path-contract tests — persisted in `tasks.md`.
- [x] GREEN ncode-only state/project/portable implementation — persisted in `tasks.md`.
- [x] TRIANGULATE legacy coexistence and filesystem-timing tests — persisted in `tasks.md`.
- [x] REFACTOR fixture consolidation and final verification — persisted in `tasks.md`.

### TDD Cycle Evidence

| Task | Layer | RED | GREEN | TRIANGULATE | REFACTOR |
|---|---|---|---|---|---|
| WU5 RED | Package/unit/integration | Host and MCP home tables failed to compile because `resolveNcodeHome` did not exist; manager expected `.ncode/extensions` but got `.zot/extensions`; skills could not find the `.ncode` project skill; portable expected `.ncodesession` but got `.zotsession`. RED `make test` failed only in the newly asserted agent, extension, skills, and portable contracts. | N/A | N/A | N/A |
| WU5 GREEN | State/project/portable behavior | RED above. | Added the exact resolver order independently in host and MCP; changed extension, skill, MCP, and portable paths; updated session UI/help and all `NCODE_HOME` consumers. Focused tests and `make test` passed. | Host and MCP independently exercise the same seven-case OS table, including Windows `LOCALAPPDATA` winning over `APPDATA`. | The first GREEN agent rerun exposed only a macOS `/var` versus `/private/var` test-fixture path normalization issue; the fixture now compares against `os.Getwd()`, after which production behavior passed unchanged. |
| WU5 TRIANGULATE | Negative/coexistence/timing | New legacy portable coverage failed because `ImportSession` still accepted `.zotsession`. | `ImportSession` now requires the case-insensitive `.ncodesession` suffix; the focused portable suite passed. | Dedicated legacy home, project extension/skill/MCP, and portable fixtures prove legacy roots and files are ignored and unchanged. Path resolution and missing config create nothing; `SaveConfig` remains lazy; `Resolve` eagerly attempts docs installation and ignores deterministic installation failure. | N/A |
| WU5 REFACTOR | Regression/race | Prior cycle evidence. | N/A | Independent host and MCP resolver assertions remain. | Consolidated the two agent-package file snapshots in `path_test_helpers_test.go`; all focused suites, nested MCP tests, searches, `git diff --check`, and `make test` passed. |

### Exact RED evidence

- `go test ./packages/agent -run '^(TestResolveNcodeHomeOrder|TestNcodeHomeUsesNcodeEnvironment|TestResolveInstallsDocsEagerlyUnderNcodeHome|TestExtensionDirsUseNcodeProjectPath)$' -count=1` — expected FAIL: `undefined: resolveNcodeHome`.
- `go test ./packages/agent/extensions -run '^TestManagerSearchDirsUseNcodeProjectPath$' -count=1` — expected FAIL: project search returned `.zot/extensions`, not `.ncode/extensions`.
- `go test ./packages/agent/skills -run '^TestDiscoverUsesNcodeProjectAndGlobalPaths$' -count=1` — expected FAIL: `project-ncode` was not discovered.
- `go test ./packages/core -run '^TestExportToFilePath$' -count=1` — expected FAIL: `PortableExt = ".zotsession", want .ncodesession`.
- Nested MCP focused test with a transient `/tmp` modfile and local-root replacement — expected FAIL: `undefined: resolveNcodeHome`.
- `make test` — expected FAIL only on the new agent compile assertion and new manager, skills, and portable assertions; all unaffected packages passed.
- TRIANGULATE RED: `go test ./packages/core -run '^TestLegacyZotPortableSessionIsNotAcceptedAndRemainsUnchanged$' -count=1` — expected FAIL because legacy `.zotsession` was still accepted.

### Final verification commands and outcomes

- State/agent focused expression covering exact home order, no-create timing, eager docs success/failure, extension paths, and legacy home/project coexistence — PASS.
- `go test ./packages/agent/extensions ./packages/agent/skills -count=1` — PASS.
- Portable focused expression covering round trip, explicit export path, large JSONL row, and legacy portable rejection — PASS.
- `go test ./packages/agent/modes -count=1` — PASS.
- `(cd examples/extensions/mcp-bridge && go test -modfile=/tmp/ncode-mcp-refactor.mod ./... -count=1)` with a transient local-root replacement — PASS; tracked nested modules remain replace-free.
- Active constructor search for `Getenv("ZOT_HOME")`, `filepath.Join(..., ".zot", ...)`, and `PortableExt = ".zotsession"`, excluding dedicated rejection fixtures — zero matches.
- MCP search for an `APPDATA` reader — zero matches; the resolver reads `LOCALAPPDATA` only.
- `git diff --check` — PASS.
- Final `make test` — PASS (`go test -race ./...`).
- No live providers, network release calls, credentials, commit, push, branch switch, PR, review, or settlement occurred.

### Files changed by WU 5

- **Home/path implementation and active path communication:** `packages/agent/config.go`, `packages/agent/args.go`, `packages/agent/build.go`, `packages/agent/botcmd.go`, `packages/agent/extcmd.go`, `packages/agent/extensions/manager.go`, `packages/agent/extupdate.go`, `packages/agent/modelsync.go`, `packages/agent/modes/interactive.go`, `packages/agent/modes/skills_dialog.go`, `packages/agent/modes/slash_suggest.go`, `packages/agent/sdk/sdk.go`, `packages/agent/skills/skills.go`, `packages/agent/swarm/runner.go`, `packages/agent/swarm/swarm.go`, `packages/agent/systemprompt.go`, `packages/agent/update.go`, `packages/agent/updatecmd.go`, `packages/core/session_portable.go`, `packages/provider/catalog_builtin.go`, `packages/provider/usermodels.go`, `packages/tui/theme_loader.go`, and extension example comments in `examples/extensions/{guard,hello,weather}/main.go`.
- **MCP bridge parity:** `examples/extensions/mcp-bridge/{config.go,config_test.go,fsutil_test.go,main.go,setup.go,setup_test.go}`.
- **Updated host tests/fixtures:** `packages/agent/{build_test.go,cli_headless_test.go,config_command_test.go,config_gcp_test.go,config_home_test.go,config_proxy_test.go,extcmd_test.go,extupdate_test.go,modelsync_test.go,sessionscmd_test.go,settings_store_test.go}`, `packages/agent/extensions/manager_test.go`, `packages/agent/skills/skills_test.go`, `packages/core/{core_test.go,session_portable_test.go}`.
- **New dedicated tests/helpers:** `packages/agent/{legacy_zot_home_test.go,legacy_zot_project_path_test.go,path_test_helpers_test.go}`, `packages/agent/extensions/legacy_zot_project_path_test.go`, `packages/agent/skills/legacy_zot_project_path_test.go`, `packages/core/legacy_zot_portable_session_test.go`, and `examples/extensions/mcp-bridge/legacy_zot_project_path_test.go`.
- **SDD artifacts:** `openspec/changes/clean-break-ncode-identity/{tasks.md,apply-progress.md}`.

### Scope, line count, and deviations

- Implementation scope is 898 changed lines: 407 additions + 170 deletions in tracked implementation files, plus 321 lines across seven new test/helper files. This is below the native 1,400-line maximum.
- No design deviation occurred. Neutral state basenames and schemas remain unchanged; there is no legacy fallback, import, migration, or scan.
- The MCP bridge preserves an independent resolver while matching host order exactly and using `LOCALAPPDATA`, never `APPDATA`, on Windows.
- Existing tests and active comments that selected the home root were changed from `ZOT_HOME` to `NCODE_HOME`; no other product environment control was renamed, so WU 6 remains unstarted.
- Portable import now validates the ncode portable suffix, as required by the dedicated legacy rejection scenario; JSONL content and session-store behavior remain unchanged.
- The shared `Resolve` → `Resolved.NewClient` → `Resolved.NewAgent` spine, WU4 command identity, RPC v1, extension acknowledgement, swarm/runtime composed identifiers, and release/update endpoints remain outside this unit.

## Remaining implementation tasks (exact unchecked lines)
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

## Deferred parent lifecycle actions

Both parent-owned rows are checked and were preserved. Clone-local bounded review remains disabled. Parent owns the independently revertible WU5 commit, native attempt settlement, checkpoint pushes, issue linkage, CI, one final PR from `feat/clean-break-ncode-identity-02` to `main` with exactly `type:breaking-change`, final merge, and any later WU6 delegation.

## Risks and next boundary

- The user explicitly accepts the combined final-PR size exception; independently revertible WU commits plus focused verification and `make test` before the next WU remain the delivery controls.
- Standalone nested-module tidy remains publication-bootstrapped until a released ncode version declares the new module path; transient local-root validation passed and tracked modules remain replace-free.
- This worktree remains on `feat/clean-break-ncode-identity-02`; WU5 apply performed no commit, push, PR, merge, branch switch, review transaction, settlement, or WU6 work.
- WU6 is next but unstarted. Parent must first preserve WU5 as its own revertible commit and settle the native attempt after its independent gate.

## WU 5 gate correction: external module verification cleanup

- **Boundary:** corrected only the WU5 nested-module gate residue under continuation token `sha256:b4ca7db27ff48bc49314bfb696cd493b4027cfeedf18ad4313930a3fb8665e0d`; WU6 remained unstarted and settlement remains parent-owned.
- **Cleanup:** removed `/tmp/ncode-mcp-refactor.mod` and `/tmp/ncode-mcp-refactor.sum`; explicit final checks proved both paths absent.
- **Fresh nested verification:** copied the complete `examples/extensions/mcp-bridge` module to `/tmp/ncode-mcp-bridge-fresh.MkMsLj`, proved the copy matched before modification, added `replace github.com/nlf/ncode => /Users/nlf/Projects/nlf/ncode` only to the copy, and ran `GOWORK=off GOFLAGS='' go mod tidy` and `GOWORK=off GOFLAGS='' go test ./... -count=1`; both passed. The external copy was then deleted and proved absent.
- **Repository immutability before artifact persistence:** repository status, binary diff, and all modified/untracked content were byte-identical before and after cleanup plus fresh external verification. SHA-256 values were respectively `725a15c140e56f4106bd808235d4c0b01242ca4189040fcf6eed543dac21bccb`, `428fe813c94d736e9d3d4780ea52c738e1d0022f502c7c661f1d496c960d9b88`, and `340cbea1edac7dc41ef06e4f9a6c53e05a1e642f5904bc3ff4d4c236665879f9` both before and after. This progress-only evidence update followed that comparison; no product file or task checkbox changed.
- **Repository overlay checks:** shell and Go environment `GOFLAGS`/`GOWORK` were unset/empty; no `go.work` or `go.work.sum`, no `replace` directive, and no untracked modfile/workspace/overlay existed. The only repository module metadata remained root `go.mod`/`go.sum`, MCP bridge `go.mod`/`go.sum`, and todo `go.mod`.
- **Final hygiene:** `git diff --check` passed after this cumulative progress update. No commit, push, PR, merge, branch switch, review, settlement, or WU6 work occurred.

### TDD Cycle Evidence

| Task | Layer | Safety Net | RED | GREEN | TRIANGULATE | REFACTOR |
|---|---|---|---|---|---|---|
| WU5 nested-module gate correction | Verification-only external module integration | Existing WU5 implementation and task checkboxes were preserved byte-for-byte | N/A — no production behavior or code changed | Fresh copied module passed tidy and `go test ./...` against the local root | Complete-copy check plus repository residue/overlay absence and post-delete checks proved isolation | Removed leaked external artifacts; no product refactor was performed |

### Test summary

- **Tests written:** none; this was a cleanup and fresh nested-module verification correction.
- **Test command:** `GOWORK=off GOFLAGS='' go test ./... -count=1` from the fresh external MCP bridge copy — PASS (`ok github.com/nlf/ncode/examples/extensions/mcp-bridge 0.409s`).
- **Live providers/credentials:** none used.

## WU 6 apply evidence

### Completed implementation tasks and persisted checkboxes

- [x] RED — positive NCODE environment tests and genuine failures, persisted in `tasks.md`.
- [x] GREEN — all Go product readers and test helpers renamed, persisted in `tasks.md`.
- [x] TRIANGULATE — dedicated legacy-only/conflict fixtures, persisted in `tasks.md`.
- [x] REFACTOR — explicit 23-entry table, helper extraction, searches, and final tests, persisted in `tasks.md`.

### Exact 23-entry environment mapping

| Legacy inventory name | Ncode decision |
|---|---|
| `ZOT_HOME` | `NCODE_HOME` |
| `ZOT_FLAT_TOOLS` | `NCODE_FLAT_TOOLS` |
| `ZOT_COMPACT_INPUT` | `NCODE_COMPACT_INPUT` |
| `ZOT_INLINE_IMAGES` | `NCODE_INLINE_IMAGES` |
| `ZOT_CELL_ASPECT` | `NCODE_CELL_ASPECT` |
| `ZOT_TOOL_ARG_WIDTH` | `NCODE_TOOL_ARG_WIDTH` |
| `ZOT_THEME` | `NCODE_THEME` |
| `ZOT_NO_BROWSER` | `NCODE_NO_BROWSER` |
| `ZOT_FORCE_BROWSER` | `NCODE_FORCE_BROWSER` |
| `ZOT_DEBUG_ANTHROPIC` | `NCODE_DEBUG_ANTHROPIC` |
| `ZOT_AGENT_SKILLS` | `NCODE_AGENT_SKILLS` |
| `ZOT_AGENT_CONSENT` | deleted with Zotfiles; no replacement |
| `ZOTCORE_RPC_TOKEN` | `NCODE_RPC_TOKEN` exactly |
| `ZOT_SWARM_AGENT_ID` | `NCODE_SWARM_AGENT_ID` |
| `ZOT_SWARM_EVENT_LOG` | `NCODE_SWARM_EVENT_LOG` |
| `ZOT_SWARM_CREDENTIAL_STDIN` | `NCODE_SWARM_CREDENTIAL_STDIN` |
| `ZOT_VERSION` | `NCODE_VERSION` (installer reader remains WU 10) |
| `ZOT_PREFIX` | `NCODE_PREFIX` (installer reader remains WU 10) |
| `ZOT_AGENT_API_KEY_COMMAND_HELPER` | `NCODE_AGENT_API_KEY_COMMAND_HELPER` |
| `ZOT_HELP_HELPER` | `NCODE_HELP_HELPER` |
| `ZOT_SWARM_CREDENTIAL_HELPER` | `NCODE_SWARM_CREDENTIAL_HELPER` |
| `ZOT_API_KEY_COMMAND_HELPER` | `NCODE_API_KEY_COMMAND_HELPER` |
| `ZOT_API_KEY_COMMAND_VALUE` | `NCODE_API_KEY_COMMAND_VALUE` |

The dedicated inventory test asserts exactly 23 unique legacy decisions, 22 unique ncode targets, and no consent replacement. All actual Go `Getenv` product controls use ncode names; standard provider and OS controls were not renamed. Installer `VERSION`/`PREFIX` readers are intentionally still owned by WU 10, and non-Go RPC clients/docs remain WU 7/WU 11 work.

### TDD Cycle Evidence

| Task | Test files | Layer | Safety Net | RED | GREEN | TRIANGULATE | REFACTOR |
|---|---|---|---|---|---|---|---|
| WU6 RED | `config_env_test.go`, `rpc_env_test.go`, package `ncode_env_test.go` files, `swarm/runner_test.go` | Package/integration | Focused seven-package suite and baseline `make test` passed | Config overrides returned old config values; wrong NCODE RPC token was accepted; four TUI controls ignored NCODE; browser disable was ignored; debug file was absent; extra skills were absent; swarm helper exited 2. Corrected RED focused suites and RED `make test` failed on the new contracts. | N/A | N/A | N/A |
| WU6 GREEN | production readers plus existing helper tests | Package/integration | RED above | RED above | Focused owner tests and `make test` passed after exact namespace replacement. | Existing ncode tests exercised alternate values and retained provider/OS behavior. | N/A |
| WU6 TRIANGULATE | five dedicated `legacy_zot_env_test.go` files | Negative/package/local subprocess | GREEN above | New injectable browser and swarm-reader tests initially failed to compile because `hasBrowser`, `swarmCredentialStdinEnabled`, and `swarmEventLogPath` did not exist. | Minimal helper extraction preserved behavior and made those cases pass. | State/config, rendering, browser/debug, skills, API-key helpers, RPC authorization, and swarm child metadata cover legacy-only and conflicting values with local files, subprocesses, and `httptest`. | N/A |
| WU6 REFACTOR | same tables and owning tests | Regression/race | TRIANGULATE green | N/A | N/A | Exact mapping remains visible as 23 rows. | Environment setup is table-driven; browser and swarm reads are testable without behavior drift; focused tests, searches, `git diff --check`, and final `make test` passed. |

Two initial RED fixtures were corrected before production edits: the theme assertion compared an uncomparable struct, and the debug fixture used an unknown model. The corrected fixtures then failed for the intended missing NCODE behavior.

### Verification and search evidence

- Focused safety net: `go test ./packages/agent ./packages/agent/modes ./packages/agent/skills ./packages/agent/swarm ./packages/provider ./packages/provider/auth ./packages/tui -count=1` — PASS.
- Baseline `make test` (`go test -race ./...`) — PASS.
- Corrected RED focused runs — FAIL on every owning package for missing NCODE behavior; RED `make test` — FAIL on the new contracts.
- GREEN focused tests and `make test` — PASS.
- TRIANGULATE focused package suite and `make test` — PASS.
- Final focused command over `examples/rpc/go` and the six owning packages — PASS.
- Exact tracked-Go search `git grep -nI -E 'ZOTCORE_|ZOT_[A-Z0-9_]+' -- '*.go'` emitted only four pre-existing WU5 rejection lines in `legacy_zot_home_test.go` / `legacy_zot_project_path_test.go` because new files were untracked.
- All-worktree scan found 56 legacy Go lines, all in dedicated `legacy_zot_*` files; active old `os.Getenv`/`LookupEnv` readers were zero.
- Full-repository exact search still reports planned active docs, non-Go RPC clients, installer variables, provenance, and rejection fixtures assigned to WU 7/WU 10/WU 11/WU 12; none is an active Go reader owned by WU 6.
- `NCODE_AGENT_CONSENT` appears only as a negative assertion in `packages/agent/legacy_zot_env_test.go`; no production or active test helper reads it.
- `git diff --check` — PASS.
- Final `make test` — PASS (`go test -race ./...`).
- No live provider, credential, release network, commit, push, PR, branch switch, review, settlement, or WU 7 action occurred.

### Files and scope

- Production/env edges: `examples/rpc/go/main.go`; `packages/agent/{config.go,rpc.go,swarm_agent.go}`; `packages/agent/skills/skills.go`; `packages/agent/swarm/runner.go`; `packages/provider/anthropic.go`; `packages/provider/auth/manager.go`; `packages/tui/{detect_bg.go,image.go,render.go,view.go}`.
- Updated existing tests/helpers: `packages/agent/{args_test.go,config_command_test.go}`; `packages/agent/modes/interactive.go`; `packages/agent/skills/skills_test.go`; `packages/agent/swarm/runner_test.go`; `packages/provider/auth/api_key_command_test.go`; `packages/tui/{image_test.go,view_tool_arg_width_test.go}`.
- New positive tests: `packages/agent/{config_env_test.go,rpc_env_test.go}`, `packages/agent/skills/ncode_env_test.go`, `packages/provider/ncode_env_test.go`, `packages/provider/auth/ncode_env_test.go`, `packages/tui/ncode_env_test.go`.
- New rejection fixtures: `packages/agent/legacy_zot_env_test.go`, `packages/agent/skills/legacy_zot_env_test.go`, `packages/provider/legacy_zot_env_test.go`, `packages/provider/auth/legacy_zot_env_test.go`, `packages/tui/legacy_zot_env_test.go`.
- SDD artifacts: `tasks.md` and this cumulative `apply-progress.md`.
- Implementation/test scope: 775 changed lines before this progress artifact, below the 1,300-line cap.

### Deviations, risks, and next boundary

- No runtime design deviation occurred. RPC v1 frame/hello semantics, WU5 path behavior, provider/OS variables, and swarm event/control schemas remain unchanged.
- The Go RPC example token read moved in WU6 because it is an actual Go `Getenv` product control; broader RPC invocation, clients, docs, and diagnostics remain WU7.
- Full-repository old-environment output is not yet a zero-active audit because WU7 owns remaining non-Go RPC clients/docs, WU10 owns installer `ZOT_VERSION`/`ZOT_PREFIX`, WU11 owns active prose, and WU12 owns the final allowlist gate.
- WU7 is the exact next unchecked implementation boundary. Parent must settle and preserve WU6 as an independently revertible commit before starting it.

## WU 7 apply evidence

### Completed implementation tasks and persisted checkboxes

- [x] RED — RPC identity and neutral-v1 tests with genuine first-frame/client-contract failures, persisted in `tasks.md`.
- [x] GREEN — ncode invocation, token, diagnostics/comments, docs, and four reference clients, persisted in `tasks.md`.
- [x] TRIANGULATE — dedicated legacy-token and Zot-branded unknown-extra-field coverage, persisted in `tasks.md`.
- [x] REFACTOR — explicit prompt-first/token paths, shared test-only frame helpers, client checks, searches, and full tests, persisted in `tasks.md`.

All four rows are visibly `[x]` in the OpenSpec task artifact and in Engram task observation `117`. WU8 remains unchecked and unstarted.

### TDD Cycle Evidence

| Task | Test files | Layer | Safety Net | RED | GREEN | TRIANGULATE | REFACTOR |
|---|---|---|---|---|---|---|---|
| WU7 RED | `rpc_protocol_test.go`, `rpc_env_test.go`, `rpc_identity_test.go`, `legacy_zot_rpc_test.go`, adjacent `rpc_reasoning_test.go` | Package/integration + static contract | Focused RPC/help, complete agent package, and baseline `make test` passed | Focused run and `make test` failed because a non-hello first frame did not close token-gated RPC, legacy Python/Node source filenames remained, and docs/shell/Go/Python/Node still published old invocation/token identity. Neutral prompt-first and optional-hello characterization cases already passed. | N/A | N/A | N/A |
| WU7 GREEN | server edge, docs, command/help comments, shell/Go/Python/Node clients | Package/integration + client contract | RED above | RED above | First-frame auth now closes after one neutral failure response; all active RPC edges use `ncode rpc`/`NCODE_RPC_TOKEN`; focused RPC, agent package, client contract, and `make test` passed. | Existing matching/wrong native-token cases and prompt/hello paths use distinct inputs. | N/A |
| WU7 TRIANGULATE | `legacy_zot_rpc_test.go`, `rpc_protocol_test.go`, `rpc_env_test.go` | Negative/package/local fake | GREEN above | Dedicated legacy tests were written before production edits and joined the RED suite. | Legacy-only prompt-first, unknown-field non-authorization, and neutral unknown-field decoding passed. | A valid ncode token with conflicting Zot-branded extras also passed, proving only the native token is authoritative across four legacy/extra-field paths. | N/A |
| WU7 REFACTOR | all RPC tests, including `rpc_reasoning_test.go` | Regression/race + static/syntax/build | TRIANGULATE green | N/A | N/A | Neutral v1 prompt/response/event/optional-hello cases remain explicit. | Shared only test transport/frame decoding helpers; strengthened reasoning response-shape assertions; focused RPC/help/SDK tests, all clients, searches, `git diff --check`, and final `make test` passed. |

### RED evidence

- Safety net `go test ./packages/agent -run 'RPC|HelpOutputStreams' -count=1` — PASS.
- Safety net `go test ./packages/agent -count=1` — PASS.
- Safety net `make test` (`go test -race ./...`) — PASS.
- RED focused expression covering prompt-first, optional hello, native token, legacy token/extra fields, and published contracts — expected FAIL:
  - `TestNcodeRPCTokenRejectsNonHelloFirstFrameAndCloses` observed the server emit an auth failure and then incorrectly accept the second hello;
  - `TestLegacyZotRPCClientSourceNamesAreNotPublished` found both old source filenames;
  - the five documentation/reference-client subtests found old invocation/token identity or missing ncode client filenames.
- RED `make test` — expected FAIL on the same new WU7 assertions while unaffected packages remained green.

### GREEN and TRIANGULATE evidence

- GREEN focused WU7 expression — PASS.
- GREEN `go test ./packages/agent -count=1` — PASS.
- GREEN `make test` — PASS.
- TRIANGULATE focused neutral/legacy RPC expression — PASS.
- TRIANGULATE `make test` — PASS.
- The only runtime correction makes the documented token requirement truly first-frame: a token-gated non-hello receives one existing-shape `success:false` response and the connection closes. Matching/wrong hello behavior and all token-unset behavior remain otherwise unchanged.

### Reference-client checks

- `bash -n examples/rpc/shell/prompt.sh` — PASS.
- `go test ./examples/rpc/go -count=1` — PASS (`[no test files]`, package builds).
- `go build -o <external-temp> ./examples/rpc/go` and executable assertion — PASS; temp output removed.
- `PYTHONPYCACHEPREFIX=<external-temp> python3 -m py_compile examples/rpc/python/ncode_client.py` — PASS; external cache removed and no repository bytecode created.
- `node --check examples/rpc/node/ncode-client.js` — PASS.

### Neutral RPC v1 protocol proof

- Token unset: a prompt is accepted as the first frame and reaches a local fake provider with the exact message; the first output remains `type:"response"`, the same id/`command:"prompt"`/`success:true`, and `data.started:true`.
- Neutral events: local fake output retains `turn_start`, `user_message`, `text_delta`, `assistant_message`, `turn_end`, and `done`; no product or ncode-version field appears.
- Optional hello: token-unset hello remains optional and returns exactly the four existing data keys `protocol_version`, `version`, `provider`, and `model`, with protocol version `1`.
- Native authorization: matching `NCODE_RPC_TOKEN` succeeds with neutral v1 hello data; wrong token and non-hello first frame fail without accepting a later frame.
- Legacy isolation: `ZOTCORE_RPC_TOKEN` alone does not gate prompt-first RPC; unknown `zot_token`/`zot_product` fields cannot authorize native token-gated RPC, cannot override a valid native token, and have no configuration meaning when authorization is disabled.
- Static protocol search found only protocol version `1` in `rpc.go`/`docs/rpc.md` and found no RPC v2, `product`, or `ncode_version` frame field in server/docs/clients.

### REFACTOR/final verification

- `go test ./packages/agent -run 'RPC|HelpOutputStreams' -count=1` — PASS after the final test-helper/reasoning assertion refactor.
- `go test ./packages/agent/sdk -count=1` — PASS.
- Active forbidden RPC identity search over server, invocation/help comments, neutral JSON/SDK comments, docs, clients, and non-legacy tests — zero matches.
- RPC source-filename search excluding the WU10-owned tracked bytecode artifact — zero Zot-bearing source filenames.
- `git diff --check` — PASS.
- Final `make test` — PASS (`go test -race ./...`).
- No live provider, credential, external network, release API, commit, push, PR, branch switch, review, settlement, or WU8 action occurred.

### Files and scope

- Modified runtime/invocation/comments: `packages/agent/{rpc.go,cli.go}`, `packages/agent/modes/json.go`, `packages/agent/sdk/{sdk.go,types.go}`.
- Modified and added tests: `packages/agent/{rpc_env_test.go,rpc_reasoning_test.go,rpc_protocol_test.go,rpc_identity_test.go,legacy_zot_rpc_test.go}`.
- Modified user contract: `docs/rpc.md`, `examples/rpc/{go/main.go,shell/prompt.sh}`.
- Renamed and updated: `examples/rpc/node/zot-client.js` → `examples/rpc/node/ncode-client.js`; `examples/rpc/python/zot_client.py` → `examples/rpc/python/ncode_client.py`.
- SDD artifacts: `openspec/changes/clean-break-ncode-identity/{tasks.md,apply-progress.md}` plus Engram task/apply-progress observations.
- Rename-aware implementation/test scope: **402 changed lines** (`350` additions + `52` deletions), under the native 1,300-line maximum. The accepted one-final-PR strategy resolves the 400-line review forecast; no intermediate PR is created.

### Deviations, deferred boundaries, and next work unit

- No protocol-design deviation occurred: RPC remains neutral version 1, hello remains optional when no native token is set, and no product field, mandatory always-on hello, legacy read, alias, or frame-shape redesign was introduced.
- Returning after the first token-gated non-hello frame corrects the existing documented “first line must be hello; process exits” behavior rather than adding a new handshake.
- The tracked `examples/rpc/python/__pycache__/zot_client.cpython-314.pyc` artifact remains untouched because WU10 explicitly owns its deletion; syntax checks wrote only to an external temporary cache.
- WU8 extension protocol-v2 acknowledgement and SDK work remains wholly unstarted. The cumulative exact unchecked-task section now begins with all four WU8 rows.
- Parent owns the independently revertible `feat(rpc): publish ncode invocation contract` commit and settlement for token `sha256:b922f6bb7768cbf8d3481e6b3f4c5a15fc56ad88775831cd7079c5299d9b4cc9`.

## WU 7 gate correction: malformed token-gated first frame and stale chain context

### Correction boundary and status

- **Status:** both delegated WU7 gate blockers are corrected; WU1–WU7 remain complete and WU8 remains next/unstarted.
- **Runtime boundary:** when `NCODE_RPC_TOKEN` is set, malformed JSON in the first non-empty frame now emits the existing malformed-JSON failure shape and terminates the connection before a later hello can authorize. Token-unset malformed-frame recovery remains unchanged.
- **Planning boundary:** the tasks order, current boundary, and chain diagram now show WU6 and WU7 complete and WU8 next/unstarted. Every implementation and parent checkbox was preserved.
- **Native authority:** this correction reused parent-supplied token `sha256:b922f6bb7768cbf8d3481e6b3f4c5a15fc56ad88775831cd7079c5299d9b4cc9`; `gentle-ai` remains unavailable locally, so no acquire or settlement was attempted. Parent owns settlement.
- **Workload/PR boundary:** correction only; no WU8, extension protocol, commit, push, PR, branch switch, review, or settlement action occurred.

```yaml
schemaName: spec-driven
changeName: clean-break-ncode-identity
artifactStore: both
artifacts:
  proposal: done
  specs: done
  design: done
  tasks: done
  applyProgress: done
taskProgress:
  total: 44
  complete: 25
  remaining: 19
deferredParentActions:
  total: 2
  complete: 2
  remaining: 0
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
  allowedEditRoots: [/Users/nlf/Projects/nlf/ncode]
  warnings:
    - native gentle-ai status/attempt CLI unavailable; authoritative OpenSpec state reconstructed from the installed status contract
nextRecommended: parent-lifecycle
isNonAuthoritative: false
```

Ownership recount after the tasks context edit found 44 implementation rows (`25` checked, `19` unchecked), 2 parent rows (both checked), and 0 malformed markers. The exact WU7 rows remain `[x]`; the exact WU8 rows remain `[ ]`.

### TDD Cycle Evidence

| Task | Test file | Layer | Safety Net | RED | GREEN | TRIANGULATE | REFACTOR |
|---|---|---|---|---|---|---|---|
| Reject malformed first token-gated frame | `packages/agent/rpc_env_test.go` | Package/integration | Existing native-token hello, non-hello close, token-unset prompt-first, and optional-hello cases passed | `TestNcodeRPCTokenRejectsMalformedFirstFrameAndCloses` genuinely failed: the server emitted malformed JSON, then accepted the later valid hello and returned nil | Minimal auth-order correction returns after the malformed response while unauthenticated; the focused regression passed | `TestRPCWithoutTokenContinuesAfterMalformedFrame` proves token-unset RPC still emits the malformed response and processes a later neutral ping | No extraction was needed; the three-line production guard is the minimal branch and focused/full tests remained green |
| Refresh tasks boundary | `tasks.md` structural context | Planning artifact | Checkbox recount was 25/44 implementation complete plus 2/2 parent complete and 0 malformed markers | N/A — no runtime behavior and no checkbox change | Order, boundary, and diagram now agree on WU7 complete/WU8 next | Table and diagram independently state the same boundary | No further planning refactor was needed |

### Exact correction evidence

- Safety net: `go test ./packages/agent -run '^(TestNcodeRPCTokenAuthorizesExistingHelloContract|TestNcodeRPCTokenRejectsNonHelloFirstFrameAndCloses|TestRPCWithoutTokenAcceptsNeutralPromptAsFirstFrame|TestRPCOptionalHelloRemainsNeutralProtocolVersionOne)$' -count=1` — PASS.
- RED: `go test ./packages/agent -run '^TestNcodeRPCTokenRejectsMalformedFirstFrameAndCloses$' -count=1` — expected FAIL; output contained the malformed failure followed by a successful `hello` response, proving later authorization was possible.
- GREEN: the same single-test command — PASS after the production guard.
- TRIANGULATE: `go test ./packages/agent -run '^(TestNcodeRPCTokenRejectsMalformedFirstFrameAndCloses|TestRPCWithoutTokenContinuesAfterMalformedFrame)$' -count=1` — PASS.
- Focused RPC: `go test ./packages/agent -run 'RPC|HelpOutputStreams' -count=1` — PASS.
- SDK: `go test ./packages/agent/sdk -count=1` — PASS.
- Full runner: `make test` (`go test -race ./...`) — PASS.
- Final hygiene: `git diff --check` — PASS after artifact persistence.
- Task context/count gate — PASS: WU7 complete, WU8 next/unstarted; implementation `25/44` checked with `19` unchecked; parent `2/2` checked; malformed ownership markers `0`.

### Correction files and stats

- `packages/agent/rpc.go` — 3 added production lines.
- `packages/agent/rpc_env_test.go` — 50 added test lines across two focused cases.
- `openspec/changes/clean-break-ncode-identity/tasks.md` — 5 additions and 3 deletions in order/boundary/diagram context; no checkbox changed.
- `openspec/changes/clean-break-ncode-identity/apply-progress.md` — cumulative correction evidence appended and timestamp refreshed.
- Correction scope before this cumulative progress text: **61 changed lines** (`58` additions, `3` deletions) across production, tests, and tasks.

### Test summary

- **Tests written:** 2 focused RPC test functions.
- **Behavioral paths:** token-gated malformed-first-frame close and token-unset malformed-frame continuation.
- **Pure functions created:** none; the scanner/auth ordering required only a minimal conditional return.
- **Protocol preservation:** RPC remains v1; response/event/hello shapes and token-unset prompt-first behavior are unchanged.
- **Live providers/credentials/network:** none used.
- **Remaining implementation tasks:** unchanged exact 19 unchecked lines already listed cumulatively above, beginning with WU8 RED. WU8 was not started.

## WU 8 apply evidence

### Attempt summary

- **Status:** WU 8 is complete; WU 1–WU 8 are complete and WU 9 is next/unstarted.
- **Assigned boundary:** exactly WU 8 — signaled extension protocol-v2 ncode acknowledgement, extension Go SDK, raw clients, authoring skill, and extension documentation.
- **Native authority:** parent supplied active token `sha256:89974a7ebd50ef447647d51d861d710fdb7e50247c25dc79be049b5e2584571f`. The `gentle-ai` CLI was unavailable, so no duplicate acquire or settlement was attempted; parent owns settlement after its independent gate.
- **Delivery boundary:** one long-lived branch `feat/clean-break-ncode-identity-02`, accepted one-final-PR size exception, and one independently revertible WU8 slice. No commit, push, PR, branch switch, review, settlement, or WU9 work occurred.
- **Scope:** rename-aware implementation/test/docs scope is **716 changed lines** (`518` additions, `198` deletions), within the supplied 1,800-line and planned 650–900-line WU8 boundaries.
- **Network/provider boundary:** all Go module validation used `GOPROXY=off`; no live provider, credentials, release API, or other network service was used.

### Structured status consumed/produced

```yaml
schemaName: spec-driven
changeName: clean-break-ncode-identity
artifactStore: both
planningHome:
  root: /Users/nlf/Projects/nlf/ncode/openspec
  changesDir: /Users/nlf/Projects/nlf/ncode/openspec/changes
changeRoot: /Users/nlf/Projects/nlf/ncode/openspec/changes/clean-break-ncode-identity
artifactPaths:
  proposal: [openspec/changes/clean-break-ncode-identity/proposal.md]
  specs: [openspec/changes/clean-break-ncode-identity/specs/ncode-identity/spec.md]
  design: [openspec/changes/clean-break-ncode-identity/design.md]
  tasks: [openspec/changes/clean-break-ncode-identity/tasks.md, sdd/clean-break-ncode-identity/tasks]
  applyProgress: [openspec/changes/clean-break-ncode-identity/apply-progress.md, sdd/clean-break-ncode-identity/apply-progress]
artifacts:
  proposal: done
  specs: done
  design: done
  tasks: done
  applyProgress: done
taskProgress:
  total: 44
  complete: 29
  remaining: 15
deferredParentActions:
  total: 2
  complete: 2
  remaining: 0
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
  allowedEditRoots: [/Users/nlf/Projects/nlf/ncode]
  warnings:
    - native gentle-ai status/attempt CLI unavailable; authoritative OpenSpec status reconstructed from the installed status contract
nextRecommended: parent-lifecycle
isNonAuthoritative: false
```

The active change and branch were explicit and unambiguous. Every edit stayed under the authoritative repository root. The accepted single-feature-branch size exception resolves the workload gate. Ownership recount found 29/44 implementation rows checked, 15 unchecked, 2/2 parent rows checked, and zero malformed markers.

### Completed implementation tasks and persisted checkboxes

- [x] WU8 RED — protocol acknowledgement/lifecycle tests and genuine failures, persisted in OpenSpec tasks and Engram task observation `117`.
- [x] WU8 GREEN — ncode v2 host/SDK/docs/raw-client contract, persisted in both task backends.
- [x] WU8 TRIANGULATE — dual/legacy rejection, authoring-skill rename, idle auto-ready, and copied nested builds, persisted in both task backends.
- [x] WU8 REFACTOR — transitional-name removal, exact searches, focused/full verification, and WU9-next context, persisted in both task backends.

### TDD Cycle Evidence

| Task | Test files | Layer | Safety Net | RED | GREEN | TRIANGULATE | REFACTOR |
|---|---|---|---|---|---|---|---|
| WU8 RED | `packages/agent/extproto/protocol_test.go`, `packages/agent/ext/ext_test.go`, `packages/agent/ext/legacy_zot_ack_test.go`, `packages/agent/extensions/manager_test.go`; existing intercept/tool/SDK tests | Unit + subprocess integration + protocol contract | Focused extproto/ext/extensions/SDK suite PASS; baseline `make test` PASS | Focused run failed in four distinct ways: protocol constant was 1; `HostInfo.NcodeVersion` was empty for `ncode_version`; old v1 product acknowledgement was accepted; manager emitted v1 `zot_version` without `product`. RED `make test` failed on only the new WU8 cases. | N/A | Neutral hello/tool-frame characterization already passed in the RED suite. | N/A |
| WU8 GREEN | same tests plus host/SDK production | Unit + subprocess integration | RED above | RED above | `ProtocolVersion=2`, exact host fields, strict SDK identity/version checks, `HostInfo.NcodeVersion`, ncode diagnostics, docs, skill content, and raw clients passed focused tests and `make test`. | Valid ncode v2 and legacy v1 inputs exercised different paths; idle no-ready behavior remained green. | N/A |
| WU8 TRIANGULATE | dedicated `legacy_zot_ack_test.go` fixtures; manager idle test; nested MCP test | Negative + lifecycle + external module integration | GREEN above | `TestDualNcodeAndZotAcknowledgementIsRejected` genuinely failed because the SDK accepted an otherwise-valid v2 acknowledgement carrying an additional old product field. | Strict unknown-field decoding rejected the dual acknowledgement while the exact ncode v2 case stayed green. | Old-only, dual, exact-native, neutral frame, explicit ready, and idle-without-ready paths are covered. Complete external copies of both nested modules passed offline tidy/test/build. | Skill directory moved to `write-ncode-extension`; copied-module cleanup and repository immutability were proved. |
| WU8 REFACTOR | all WU8 tests | Regression/race + static protocol audit | TRIANGULATE green | N/A | N/A | Existing interceptor, tool, readiness, RPC-v1, and neutral frame behavior remained green. | Old field literals were isolated to dedicated fixtures; transitional active names and extension protocol-v1 searches returned zero; focused tests, nested builds, `git diff --check`, and `make test` passed. |

### RED evidence

- Focused safety net: `go test ./packages/agent/extproto ./packages/agent/ext ./packages/agent/extensions ./packages/agent/sdk -count=1` — PASS.
- Baseline `make test` (`go test -race ./...`) — PASS.
- RED focused command over the five new protocol/lifecycle tests — expected FAIL:
  - `TestHelloAckPublishesExactNcodeProtocolV2Identity`: `ProtocolVersion = 1, want 2`.
  - `TestRunAcceptsNcodeProtocolV2AckAndExposesHostInfo`: `NcodeVersion` was empty because the wire tag remained old.
  - `TestLegacyZotVersionAcknowledgementIsRejected`: old acknowledgement was accepted and the SDK continued.
  - `TestManagerEmitsNcodeProtocolV2AckAndIdleExtensionAutoReadies`: host output was `protocol_version:1` plus `zot_version`, with no `product` or `ncode_version`.
- RED `make test` — expected FAIL on those new extproto/ext/extensions cases while unaffected packages stayed green.
- TRIANGULATE RED: `go test ./packages/agent/ext -run '^TestDualNcodeAndZotAcknowledgementIsRejected$' -count=1` — expected FAIL because an otherwise valid v2 acknowledgement with both version fields was accepted.

### Exact extension protocol proof

- `extproto.Product` is exactly `"ncode"`; `extproto.ProtocolVersion` is exactly `2`.
- Host `hello_ack` serializes exactly the required identity keys `type:"hello_ack"`, `product:"ncode"`, `protocol_version:2`, and `ncode_version:<host version>`, followed by retained provider/model/cwd/directory metadata.
- The old field is absent from active structs, host emission, docs, skill, and raw clients. It exists only in `legacy_zot_ack_test.go` rejection fixtures.
- The Go extension SDK requires exact product `ncode`, protocol version 2, and non-empty `ncode_version`; unknown fields, trailing JSON, old-only acknowledgements, and dual acknowledgements fail closed.
- `HostInfo.NcodeVersion` receives the exact `ncode_version` value before `OnHello` registrations run.
- `TestNeutralHelloAndToolFramesRemainUnchanged` proves the extension-originated `hello` and host `tool_call` JSON shapes byte-for-byte; searches preserve neutral `extension.json`, `hello`, `ready`, `register_tool`, and `tool_call` names.
- `TestManagerEmitsNcodeProtocolV2AckAndIdleExtensionAutoReadies` proves an extension that emits no `ready` frame still reaches `Ready=true`, `AutoReady=true`, `ReadyTimedOut=false`, with `[ncode] no ready frame; auto-readying after idle` in its log.
- Active WU8 search for `zot_version|ZotVersion|write-zot-extension|[zot]`, excluding dedicated `legacy_zot_*` fixtures — zero matches.
- Active WU8 extension-protocol-v1 search — zero matches. Separate RPC proof remains `"protocol_version": 1` at `packages/agent/rpc.go:224,244`; WU8 did not alter RPC v1.
- `ExtensionToolSource` remains declared at `packages/agent/build.go:116`; readiness channels, `readyIdleWindow`, timeouts, manifests, and neutral frames remain in place.

### Cleanup and external nested-module evidence

Two independent complete-copy validations were run; both copied each whole nested module before any replacement and required `diff -qr` equality against its source. Only the external copies received `replace github.com/nlf/ncode => /Users/nlf/Projects/nlf/ncode`.

1. `/tmp/ncode-wu8-extension-copies.g54mgL`
   - `mcp-bridge`: offline tidy, `go test ./... -count=1`, and build — PASS.
   - `todo`: offline tidy, `go test ./... -count=1`, and build — PASS.
   - Repository status SHA-256 before/after: `943804ae0bf4751dac4486b0fc688029cab1fcdf6bf5849f6c068fd231b5ca19`.
   - Repository module-metadata SHA-256 before/after: `c9aa814cb7bab4842d516207d66042919433ba7ca99ec055cbd1a18ef1984c3e`.
   - Temp root deleted and explicitly proved absent.
2. `/tmp/ncode-wu8-refactor-copies.e39cOx`
   - Both modules again passed complete-copy offline tidy/test/build.
   - Repository status SHA-256 before/after: `05e08a198e06195be07fcbb917afbd5cb03b323a21ae0b5389ec004b5c84c35d`.
   - Repository module-metadata SHA-256 before/after: `c9aa814cb7bab4842d516207d66042919433ba7ca99ec055cbd1a18ef1984c3e`.
   - Temp root deleted and explicitly proved absent.

No tracked nested `replace`, local modfile, workspace, generated build output, or temporary copy remains in the repository.

### Final verification

- `go test ./packages/agent/extproto ./packages/agent/ext ./packages/agent/extensions ./packages/agent/sdk ./packages/agent/skills -count=1` — PASS.
- Focused named protocol/neutral/legacy/idle/interceptor/tool expression — PASS.
- `node --check examples/extensions/clock/index.js` — PASS.
- TypeScript compiler was unavailable; an offline static contract check proved the scratchpad raw client contains exact product/protocol/version validation and strict allowed acknowledgement keys.
- Complete-copy nested MCP bridge and todo offline tidy/test/build — PASS twice, with cleanup proof above.
- Transitional protocol and active extension-protocol-v1 searches — zero matches.
- Dedicated rejection-line review — only factual rejection/non-use fixtures.
- `git diff --check` — PASS.
- Final `make test` (`go test -race ./...`) — PASS.

### Files changed by WU 8

- Protocol/SDK: `packages/agent/extproto/{extproto.go,protocol_test.go}`, `packages/agent/ext/{ext.go,ext_test.go,legacy_zot_ack_test.go}`.
- Host/lifecycle: `packages/agent/extensions/{manager.go,manager_test.go,events.go,proc_unix.go,legacy_zot_ack_test.go}`.
- Nested extension test fixture: `examples/extensions/mcp-bridge/deferred_tools_test.go`.
- Raw clients: `examples/extensions/clock/index.js`, `examples/extensions/scratchpad/index.ts`.
- Documentation/authoring: `docs/extensions.md`; `packages/agent/skills/builtin/write-zot-extension/SKILL.md` renamed to `packages/agent/skills/builtin/write-ncode-extension/SKILL.md`; `packages/agent/skills/skills_test.go`.
- SDD artifacts: `openspec/changes/clean-break-ncode-identity/{tasks.md,apply-progress.md}` plus Engram task/apply-progress observations.

### Deviations and discoveries

- The extension-authoring Go SDK is `packages/agent/ext`, where `HostInfo` and acknowledgement validation live. The separate in-process embedding SDK at `packages/agent/sdk` has no extension handshake or `HostInfo`; it was left behaviorally unchanged and its tests were run as a retained-boundary check.
- Strict unknown-field decoding is intentional for the signaled v2 acknowledgement: it rejects a dual acknowledgement without retaining a transitional production field name. Neutral non-acknowledgement frames keep their existing decoding behavior.
- Raw Node/TypeScript clients now wait for and validate the exact v2 acknowledgement before sending registration frames. Their neutral frame names and payloads are unchanged.
- Extension documentation and the authoring skill were fully rebranded in this WU because they are atomic participants in the published acknowledgement contract; broader unrelated docs/example communication remains WU11.
- No design deviation affected RPC v1, the construction spine, extension tool adapter, readiness timing, retained manifests/frames, or later runtime/release boundaries.

### Remaining implementation tasks (exact unchecked lines)

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

### Parent lifecycle and next boundary

Parent owns independent gate, attempt settlement, the independently revertible `feat(extensions): acknowledge ncode protocol v2` commit, any checkpoint push, and later WU9 delegation. WU9 is the exact next implementation boundary; WU9 was not started.
