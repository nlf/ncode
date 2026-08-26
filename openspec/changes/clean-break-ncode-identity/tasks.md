# Implementation tasks: clean-break ncode identity

Implementation targets the pre-rename planning SHA `18325b75cc89c75b5f4842924cb377aa5bef5c4b`. The frozen `zot-baseline` is comparison-only, never an implementation target. Use `make test` as the authoritative runner (`go test -race ./...`); use fakes, `httptest`, local files, and local subprocess stubs only—never live providers or credentials in fixtures.

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 4,550–7,450 total; estimates by work unit below |
| 400-line budget risk | High |
| Chained PRs recommended | No; the user explicitly accepted the combined final-PR size exception |
| Approved split | WU 1 is merged; preserve WU 2 → 3 → 4 → 5 → 6 → 7 → 8 → 9 → 10 → 11 → 12 as independently revertible commits on `feat/clean-break-ncode-identity-02` |
| Delivery strategy | One long-lived feature branch with checkpoint pushes as appropriate, no intermediate PRs, and one final PR to `main` |
| Final PR | Link issue #1 and apply exactly `type:breaking-change` |

Decision needed before apply: No. After WU 2 apply, the user chose single-feature-branch delivery and explicitly accepted the combined final-PR size exception because they are not reviewing intermediate PRs and see no material benefit in PR overhead. Complete focused verification and `make test` for each independently revertible WU commit before beginning the next WU.
Chained PRs recommended: No — superseded by the approved final-PR delivery
Branch strategy: `feat/clean-break-ncode-identity-02` through WU 12
400-line budget risk: High — accepted for the combined final PR

| WU | Reviewable work unit | Estimated changed lines | Over 400? |
|---:|---|---:|---|
| 1 | Remove Zotfiles without replacement | 450–650 | Yes |
| 2 | Atomic module/import graph move | 600–900 | Yes |
| 3 | Product Go symbols and package comments | 300–500 | Likely |
| 4 | Command and local build identity | 180–320 | No |
| 5 | State, project, and portable paths | 650–950 | Yes |
| 6 | NCODE environment namespace | 300–480 | Likely |
| 7 | RPC identity edges and clients | 350–550 | Likely |
| 8 | Extension protocol-v2 acknowledgement and SDK | 650–900 | Yes |
| 9 | Swarm/subprocess/provider/temp internal identity | 300–480 | Likely |
| 10 | Release, updater, and installers | 450–700 | Yes |
| 11 | Active docs, examples, and issue communication | 500–800 | Yes |
| 12 | Retained-capability and identity audit | 120–220 | No |

## Chain Context

| Field | Value |
|---|---|
| Change | `clean-break-ncode-identity` |
| Strategy | WU 2–WU 12 remain on `feat/clean-break-ncode-identity-02`; checkpoint pushes are allowed, but no intermediate PRs |
| Final PR | One PR from `feat/clean-break-ncode-identity-02` to `main`, linked to issue #1 with exactly `type:breaking-change` |
| Order | WU 1 merged → WU 2 complete → WU 3 complete → WU 4 complete → WU 5 complete → WU 6 complete → WU 7 complete → WU 8 complete → WU 9 next/unstarted → WU 10 → WU 11 → WU 12 |
| Current boundary | WU 8 is complete; WU 9 is next and unstarted |
| Review budget | The user explicitly accepts the combined final-PR size exception; no intermediate PR review budget applies |
| Verification | Preserve each WU as an independently revertible commit; run focused verification and `make test` before beginning the next WU |
| Gates | Clone-local bounded review remains disabled; strict TDD, native SDD attempt authority, issue linkage, verification, CI, and final merge gates remain |
| Rollback | Revert one work-unit commit without reverting unrelated work units; after release, use an ncode-only corrective release |

```text
main: WU 1 merged
 └── feat/clean-break-ncode-identity-02
      ├── WU 2 complete
      ├── WU 3 complete
      ├── WU 4 complete
      ├── WU 5 complete
      ├── WU 6 complete
      ├── WU 7 complete
      ├── WU 8 complete
     └── 📍 WU 9 next (unstarted) → WU 10 … WU 12 → one final PR to main
```

## Constraints and common evidence

- Preserve the `Resolve` → `Resolved.NewClient` → `Resolved.NewAgent` spine, explicit host composers, `ExtensionToolSource` adapter seam, and package-cycle boundaries from `design.md`.
- Do not create a `ncodefile`, state migration, Zot alias, module `replace`, runtime prompt, legacy fallback, or tracked generated binary/bytecode replacement.
- RPC stays neutral v1: preserve prompt/event frame shapes and optional hello; rename only explicit product edges. Extension acknowledgement alone becomes signaled protocol v2 because it contains `zot_version`.
- For each unit, record the stated verification in the commit message/body and commit only that unit with its tests and applicable docs. Rollback is the named unit revert before release; after release, remediation is a corrective ncode-only release.

## WU 1 — Remove Zotfiles without replacement

**Boundary:** Start with the planning-SHA Zotfile capability intact; finish with no Zotfile command, archive, consent, agent-data, or named-agent-session behavior and no ncodefile substitute. **Commit:** `feat(agent): remove Zotfile capability`. **Verification/rollback:** run the unit checks below and `make test`; revert this commit alone before release. **Review budget:** user-approved WU 1-only `size:exception` because deleting the capability's code, tests, and documentation already exceeds 2,100 changed lines and an artificial split would leave misleading or non-deliverable intermediate states; this exception does not carry forward.

- [x] RED — add failing characterization/absence tests in `packages/agent/args_test.go`, `packages/agent/build_test.go`, `packages/agent/cli_headless_test.go`, `packages/agent/cli_session_test.go`, `packages/agent/tools/permissions_test.go`, `packages/agent/tools/sandbox_test.go`, `packages/agent/skills/skills_test.go`, and `packages/core/session*_test.go` proving `pack`/`inspect`/`verify`/`run` are not routed while direct `Args.PermissionSet` still reaches `Sandbox.SetPermissions`, normal tools, skills, sessions, extensions, and confirmation remain available; run the focused tests and `make test` expecting only these new assertions to fail. <!-- sdd-owner: implementation -->
- [x] GREEN — delete `packages/agent/zotfile.go`, `packages/agent/zotfile_test.go`, `packages/agent/zotfile_pre_test.go`, and `docs/zotfiles.md`; remove `runZotfileCommand`, `ZotfileManifest`, `Args.AgentName`, `Args.AgentDataDir`, `Args.StartupPre`, Zotfile-only `-y/--yes`, consent/agent-data/startup-display/named-session wiring, router/help/embed/README references, and `ZOT_AGENT_CONSENT` without adding any `NCODE_AGENT_CONSENT` or ncodefile symbol; preserve the exact `args.PermissionSet` → `sandbox.SetPermissions(args.PermissionSet)` call in `packages/agent/build.go`, then run focused tests and `make test`. <!-- sdd-owner: implementation -->
- [x] TRIANGULATE — extend absence tests with legacy archive/path inputs and command/help probes, and run `git grep -nI -E 'Zotfile|zotfile|ncodefile|ZOT_AGENT_CONSENT|\.zot' -- packages/agent docs README.md` plus focused permission, tools, skills, session, extension, confirmation, and headless tests followed by `make test`. <!-- sdd-owner: implementation -->
- [x] REFACTOR — remove only capability-dead helpers/imports and clarify remaining permission/session comments without changing retained behavior; rerun the WU 1 searches, `go test ./packages/agent/... ./packages/core/...`, and `make test`. <!-- sdd-owner: implementation -->

## WU 2 — Move the canonical module and import graph atomically

**Boundary:** Start after WU 1 with all imports at the old canonical path; finish with one buildable `github.com/nlf/ncode` graph and no compatibility import/replace. **Commit:** `refactor(module): move canonical imports to ncode`. **Verification/rollback:** atomic graph move; revert the entire commit, never a subset. **Review budget:** user-approved WU 2-only implementation `size:exception` because the build-preserving module/import graph move is atomic and cannot be split honestly below 400 changed lines; the separate combined final-PR size exception covers delivery of WU 2–WU 12 on the long-lived feature branch.

- [x] Characterize before the mechanical move without manufacturing RED: capture `go list ./...`, `go test ./...`, `go mod tidy` dry diff, and `git grep -nI -F 'github.com/patriceckhart/zot' -- '*.go' go.mod examples`; identify all repository-owned import sites, `go.mod`, `docs.go`, `cmd/zot`, `examples/sdk`, `examples/extensions/mcp-bridge/go.mod`, and `examples/extensions/todo/go.mod`. <!-- sdd-owner: implementation -->
- [x] Atomically change `go.mod` and every repository-owned Go/nested-example import/module declaration from `github.com/patriceckhart/zot` to `github.com/nlf/ncode`; remove the absolute `/Users/pat/Developer/zot` nested replace and do not add a `replace`, compatibility package, or import alias. <!-- sdd-owner: implementation -->
- [x] Run post-move evidence: `go mod tidy`, tidy each nested example module, `go list ./...`, `go test ./...`, `make test`, and `git grep -nI -F 'github.com/patriceckhart/zot' -- .` restricted to the inventory’s exact reviewed provenance files; prove no active import or `replace` remains. <!-- sdd-owner: implementation -->

## WU 3 — Rename Go product symbols and package comments

**Boundary:** Start from the ncode module graph; finish with renamed public product symbols but no state-root behavior redesign. **Commit:** `refactor(identity): rename product symbols`. **Verification/rollback:** focused package builds and tests remain green; revert independently while retaining WU 2 imports.

- [x] Characterize the mechanical symbol surface without manufacturing RED by recording focused `go test` results for `./packages/agent/... ./packages/core/... ./packages/provider/... ./packages/tui/...` and locating `ZotHome`, `ZotDocsDir`, `ZotVersion`, `zotHome`, `zotVersion`, and package/product comments in `packages/agent/{config,build,cli,extensions,ext,extproto,skills,swarm,sdk}`, `packages/{core,provider,tui}`, and `docs.go`. <!-- sdd-owner: implementation -->
- [x] Rename the symbol groups to `NcodeHome`, `NcodeDocsDir`, `NcodeVersion`, and ncode local identifiers at their declarations and all callers; reword active package comments and inherited-parser comments to lowercase ncode or neutral historical wording, preserving exported API shape except the intentional product-symbol rename. <!-- sdd-owner: implementation -->
- [x] Run post-move evidence: focused package tests, `go vet ./...`, `make test`, and forbidden-symbol searches for `ZotHome|ZotDocsDir|ZotVersion` outside exact provenance/soon-to-be-dedicated rejection fixtures; inspect imports to preserve the `ExtensionToolSource` and SDK/agent cycle seams. <!-- sdd-owner: implementation -->

## WU 4 — Move command and local build identity

**Boundary:** Start with the renamed import graph; finish with `cmd/ncode` and local build/install outputs only. Release/updater/installers remain WU 10. **Commit:** `refactor(cli): build ncode command`. **Verification/rollback:** a locally built/installed binary is ncode-only; revert this slice as one unit.

- [x] Characterize the mechanical command/build surface without manufacturing RED: record `make build`, `go build -o /tmp/zot ./cmd/zot`, command help/version output, `Makefile` binary targets, `cmd/zot/{main.go,main_test.go}`, and `docs.go` embed references before the move. <!-- sdd-owner: implementation -->
- [x] Move `cmd/zot/{main.go,main_test.go}` to `cmd/ncode/`, update `Makefile` and command/test references to `./cmd/ncode` and `bin/ncode`, and rename command/product help/version/error literals to lowercase ncode; leave no `cmd/zot`, `bin/zot`, dispatch branch, or alias target. <!-- sdd-owner: implementation -->
- [x] Run post-move evidence: `go build -o /tmp/ncode ./cmd/ncode`, isolated-`GOBIN` `go install github.com/nlf/ncode/cmd/ncode`, `make build`, command help/version tests, `make test`, and searches proving `/tmp/zot`, `cmd/zot`, and `bin/zot` are absent from active build inputs. <!-- sdd-owner: implementation -->

## WU 5 — Establish ncode-only state, project, and portable paths

**Boundary:** This unit owns home resolution, `.ncode` project discovery, MCP bridge parity, and `.ncodesession` only; WU 9 owns branded temp/log/socket names. **Commit:** `feat(state): isolate ncode paths`. **Verification/rollback:** state timing and no-legacy-read behavior travel with this unit; revert it whole before release.

- [x] RED — add failing table-driven tests in `packages/agent/config_home_test.go`, `packages/agent/build_test.go`, `packages/agent/extcmd_test.go`, `packages/agent/extensions/manager_test.go`, `packages/agent/skills/skills_test.go`, `packages/core/session_portable_test.go`, and `examples/extensions/mcp-bridge/*_test.go` for the exact `NCODE_HOME`/XDG/macOS/Windows `%LOCALAPPDATA%`/fallback `.ncode` order, `.ncode/{extensions,skills,mcp.json}`, `.ncodesession`, no `.zot` scan, and host/MCP equivalence; run focused tests and `make test` expecting the new assertions to fail. <!-- sdd-owner: implementation -->
- [x] GREEN — implement the exact `NcodeHome` order in `packages/agent/config.go`; update `packages/agent/{extcmd.go,extensions/manager.go,skills/skills.go}`, `packages/core/session_portable.go`, relevant `packages/agent/modes/*` session help/import UI, and `examples/extensions/mcp-bridge/config.go` to use only ncode roots/project paths and `.ncodesession`; preserve non-branded basenames and make the MCP bridge use `%LOCALAPPDATA%`, not `%APPDATA%`; run focused tests and `make test`. <!-- sdd-owner: implementation -->
- [x] TRIANGULATE — add `legacy_zot_home_test.go`, `legacy_zot_project_path_test.go`, and `legacy_zot_portable_session_test.go` at the owning package locations with coexistence sentinels/snapshots proving `$ZOT_HOME`, OS Zot defaults, `.zot`, and `.zotsession` are ignored and unchanged; add filesystem timing coverage for path resolution, eager ignored-error docs installation in `Resolve`, and existing lazy triggers, then run focused tests and `make test`. <!-- sdd-owner: implementation -->
- [x] REFACTOR — consolidate only duplicated test fixtures/path helpers while retaining independent host and MCP assertions; rerun state/project/portable focused tests, `go test ./examples/extensions/mcp-bridge/...`, the no-legacy-path search, and `make test`. <!-- sdd-owner: implementation -->

## WU 6 — Establish the coherent NCODE environment namespace

**Boundary:** Rename all 23 product environment controls except deleted Zotfile consent; provider/OS variables remain untouched. **Commit:** `feat(config): use ncode environment controls`. **Verification/rollback:** the complete positive/legacy/conflict table is in this commit; revert it together.

- [x] RED — add failing table-driven tests near the owning readers (`packages/agent/{config,build,args,rpc,swarm_agent}_test.go`, `packages/agent/modes/*_test.go`, and swarm tests) for every inventory mapping: `NCODE_HOME`, UI/image/theme, browser/debug, skills, API-key helpers, `NCODE_RPC_TOKEN`, and all `NCODE_SWARM_*`; assert ncode-only success, legacy-only non-use, and ncode-wins conflicts, then run focused tests and `make test` expecting those assertions to fail. <!-- sdd-owner: implementation -->
- [x] GREEN — rename each `os.Getenv`/`LookupEnv` product control discovered by `git grep -nI -E 'ZOTCORE_|ZOT_[A-Z0-9_]+' -- '*.go'` according to the 23-entry inventory table, with `ZOTCORE_RPC_TOKEN` becoming exactly `NCODE_RPC_TOKEN`; delete the already-removed consent control rather than mapping it, preserve provider/OS controls, and run focused tests plus `make test`. <!-- sdd-owner: implementation -->
- [x] TRIANGULATE — add dedicated `legacy_zot_env_test.go` fixtures for legacy-only and conflicting values across state, rendering, consent removal, API-key helpers, RPC authorization, and swarm metadata; prove old values neither configure nor authorize behavior, then run focused tests and `make test`. <!-- sdd-owner: implementation -->
- [x] REFACTOR — table-drive repeated environment setup without hiding the 23-name coverage; rerun `git grep -nI -E 'ZOTCORE_|ZOT_[A-Z0-9_]+' -- .`, focused tests, and `make test`. <!-- sdd-owner: implementation -->

## WU 7 — Rebrand RPC edges without changing neutral RPC v1

**Boundary:** Own `ncode rpc`, `NCODE_RPC_TOKEN`, RPC diagnostics, documentation, and reference clients; keep v1 frames and optional hello exactly neutral. **Commit:** `feat(rpc): publish ncode invocation contract`. **Verification/rollback:** server, clients, and docs are rolled back together before release.

- [x] RED — add failing RPC tests in `packages/agent/rpc_reasoning_test.go` and adjacent RPC test targets for token-unset prompt-first, `NCODE_RPC_TOKEN` first-frame hello success/failure, `ZOTCORE_RPC_TOKEN` ignored, and unchanged v1 neutral prompt/response/event/optional-hello behavior; run focused tests and `make test` expecting new cases to fail. <!-- sdd-owner: implementation -->
- [x] GREEN — update `packages/agent/rpc.go`, command invocation/help, `docs/rpc.md`, and `examples/rpc/{shell,go,python,node}` (renaming `zot-client.js`/`zot_client.py` to ncode names) to invoke `ncode rpc` and use `NCODE_RPC_TOKEN`; do not add RPC v2, product fields, mandatory hello, legacy token read, or frame-shape change; run focused tests, client syntax/build checks, and `make test`. <!-- sdd-owner: implementation -->
- [x] TRIANGULATE — add dedicated legacy RPC input coverage proving `ZOTCORE_RPC_TOKEN` alone neither enables nor gates RPC and an unknown explicit Zot-branded extra field has no authorization/configuration meaning; use fakes/local processes only, then run neutral-frame regressions and `make test`. <!-- sdd-owner: implementation -->
- [x] REFACTOR — share RPC test transport/setup helpers only if v1 prompt-first and token-required paths remain explicit; rerun RPC tests, all reference-client syntax/build checks, forbidden RPC identity searches, and `make test`. <!-- sdd-owner: implementation -->

## WU 8 — Ship the signaled extension protocol-v2 ncode acknowledgement

**Boundary:** Own host, `extproto`, Go SDK, raw examples, authoring skill, and extension docs so no buildable revision mixes acknowledgement identities. Preserve neutral extension hello/frames and idle auto-ready. **Commit:** `feat(extensions): acknowledge ncode protocol v2`. **Verification/rollback:** revert the entire acknowledgement contract slice before release.

- [x] RED — add failing acknowledgement and lifecycle tests in `packages/agent/extproto/*_test.go`, `packages/agent/ext/ext_test.go`, `packages/agent/extensions/{manager,intercept,tool}_test.go`, and `packages/agent/sdk/*_test.go` requiring `product:"ncode"`, `protocol_version:2`, `ncode_version`, `HostInfo.NcodeVersion`, old `zot_version` acknowledgement rejection, and unchanged neutral hello/frame plus idle auto-ready behavior; run focused tests and `make test` expecting failures. <!-- sdd-owner: implementation -->
- [x] GREEN — change `packages/agent/extproto/extproto.go` `ProtocolVersion` to 2 and `HelloAckFromHost`; update `packages/agent/extensions/manager.go`, `packages/agent/ext/ext.go`, and `packages/agent/sdk` to emit/validate the ncode acknowledgement and rename the SDK field; retain `ExtensionToolSource`, readiness channels/timeouts, neutral manifest/frame names, and auto-ready logic while changing only its `[zot]` diagnostic to `[ncode]`; update `docs/extensions.md`, `packages/agent/skills/builtin/write-zot-extension/`, and raw extension examples together, then run focused tests and `make test`. <!-- sdd-owner: implementation -->
- [x] TRIANGULATE — rename the authoring skill directory to `write-ncode-extension`, add dedicated legacy acknowledgement fixtures, build/test `examples/extensions/mcp-bridge` and `examples/extensions/todo` from their own modules, and verify an idle extension still becomes ready without a ready frame; run all focused tests and `make test`. <!-- sdd-owner: implementation -->
- [x] REFACTOR — remove transitional field names/test duplication without accepting a dual ack or weakening old-ack rejection; rerun extproto/manager/ext/SDK tests, nested extension module builds, protocol searches, and `make test`. <!-- sdd-owner: implementation -->

## WU 9 — Rename swarm, subprocess, provider, and internal composed identity

**Boundary:** Sole owner of branded temp, log, socket, request, and swarm metadata names; do not revisit WU 5 home/project/portable paths. **Commit:** `feat(runtime): rebrand internal identity`. **Verification/rollback:** parent/child and composed-name contract reverts as one unit.

- [ ] RED — add failing focused tests in `packages/agent/{swarm_agent,extupdate,updatecmd}_test.go`, `packages/agent/swarm/{runner,inbox,socketpath,event}_test.go`, `packages/agent/tools/*_test.go`, `packages/agent/modes/timeline_view_test.go`, `packages/provider/{gemini,provider}_test.go`, and `packages/tui/*_test.go` for `NCODE_SWARM_*`, `ncode-swarm-*`, ncode temp/log/request names, and ncode product headers; run focused tests and `make test` expecting failures. <!-- sdd-owner: implementation -->
- [ ] GREEN — update `packages/agent/{swarm_agent.go,extupdate.go,updatecmd.go,build.go}`, `packages/agent/swarm/{socketpath.go,runner.go}`, `packages/agent/tools/bash.go`, `packages/agent/modes/timeline_view.go`, `packages/provider/{gemini.go,openai_codex.go}`, and `packages/tui/clipboard_darwin.go` to the inventory’s exact `NCODE_*`/`ncode-*` values; retain `os.Executable` child location, neutral swarm event/control shapes, and provider-mandated external fingerprints; run focused tests and `make test`. <!-- sdd-owner: implementation -->
- [ ] TRIANGULATE — add `legacy_zot_swarm_test.go` and provider/internal negative cases proving Zot swarm metadata cannot supply credentials/event logs or create `zot-swarm-*`, while retained child execution, JSONL events, and neutral controls work; use local stubs/`httptest`, never live providers, then run focused tests and `make test`. <!-- sdd-owner: implementation -->
- [ ] REFACTOR — centralize only repeated ncode prefix assertions while preserving exact output/path assertions; rerun swarm/provider/temp focused tests, composed-name searches, and `make test`. <!-- sdd-owner: implementation -->

## WU 10 — Replace release, updater, and installer identity

**Boundary:** Own release owner/repo/assets, updater/changelog lookup, installers, and tracked generated-artifact deletion; command source remains WU 4. **Commit:** `feat(release): distribute ncode only`. **Verification/rollback:** distribution inputs and their tests revert together before release.

- [ ] RED — add failing static/unit tests or snapshots for `.goreleaser.yaml`, `install.sh`, `install.ps1`, `packages/agent/{update,updatecmd,changelog}_test.go`, and `.github/workflows/release.yml` requiring `nlf/ncode`, `ncode_<version>_<os>_<arch>`, `ncode[.exe]`, `NCODE_VERSION`, `NCODE_PREFIX`, ncode checksums/install URLs, and no legacy endpoint; run local focused tests and `make test` expecting failures. <!-- sdd-owner: implementation -->
- [ ] GREEN — update `.goreleaser.yaml`, `install.sh`, `install.ps1`, `.github/workflows/release.yml`, and `packages/agent/{update.go,updatecmd.go,changelog.go}` to ncode-only owner/repository/assets/extraction/user-agent/temp identity; delete `examples/extensions/todo/zot-todo-extension` and `examples/rpc/python/__pycache__/zot_client.cpython-314.pyc` without replacements, preserve neutral `checksums.txt`, and add the required clean-break installer/release notice; run tests and `make test`. <!-- sdd-owner: implementation -->
- [ ] TRIANGULATE — validate GoReleaser config/snapshot and archive contents when tooling is available, installer shell syntax with mocked HTTP/filesystems, PowerShell parser/static checks, updater `httptest` responses, and checksum/asset-name agreement; assert no release call reaches a network provider or live GitHub release, then run `make test`. <!-- sdd-owner: implementation -->
- [ ] REFACTOR — derive repeated asset-name expectations from one local test helper/config fixture without adding an updater fallback; rerun distribution checks, `make build`, focused updater tests, forbidden asset/endpoint searches, and `make test`. <!-- sdd-owner: implementation -->

## WU 11 — Rewrite active product communication and examples

**Boundary:** After contracts settle, rewrite active docs/examples/issues and authoring themes; review provenance records separately. **Commit:** `docs(identity): publish ncode clean break`. **Verification/rollback:** active communication can revert independently without altering runtime contracts.

- [ ] Characterize the mechanical communication surface without manufacturing RED: record link/command checks and nested-example builds for `README.md`, `CONTRIBUTING.md`, `AGENTS.md`, `docs/{providers,extensions,rpc,skills,themes}.md`, `.github/ISSUE_TEMPLATE/{bug_report,feature_request}.yml`, `examples/**`, `packages/agent/skills/builtin/write-zot-themes/SKILL.md`, and active help/error/log expectations; distinguish exact provenance files from active materials. <!-- sdd-owner: implementation -->
- [ ] Rename active product prose, commands, imports, state/env/SDK/extension references, test names/expectations, manifests, and example consumers to lowercase ncode; move `write-zot-themes` to `write-ncode-themes`; ensure README Install, installer/install output, and README first-run/authentication prominently state the clean break and no Zot credentials/settings/sessions/caches/extensions/SDK-RPC-swarm reuse, with no migration/prompt instruction. <!-- sdd-owner: implementation -->
- [ ] Run post-edit evidence: nested example module tests/builds, example shell/Python/Node syntax checks, markdown link/command searches, installer/help snippet checks, `make test`, and a forbidden-active-identity scan; retain Zot wording only in the exact provenance manifest or dedicated `legacy_zot_*` rejection files. <!-- sdd-owner: implementation -->

## WU 12 — Final retained-capability and identity audit

**Boundary:** Verification-only and narrowly scoped corrections assigned back to their owning WU; do not create apply/verify artifacts or broaden behavior. **Commit:** `test(identity): audit ncode clean break` only if dedicated audit tests/manifest edits are needed. **Verification/rollback:** all gates pass from a clean tree; corrective changes revert to their owner unit before release.

Use this exact audit after the inventory manifest has admitted this task file as reviewed planning provenance:

```sh
audit_pathspecs=(
  .
  ':(exclude)docs/ncode-architecture.md'
  ':(exclude)docs/inherited-capabilities.md'
  ':(exclude)docs/superpowers/specs/2026-07-28-configurable-auto-compact-threshold-design.md'
  ':(exclude)docs/superpowers/plans/2026-07-28-configurable-auto-compact-threshold.md'
  ':(exclude)openspec/project.md'
  ':(exclude)openspec/config.yaml'
  ':(exclude)openspec/changes/clean-break-ncode-identity/proposal.md'
  ':(exclude)openspec/changes/clean-break-ncode-identity/specs/ncode-identity/spec.md'
  ':(exclude)openspec/changes/clean-break-ncode-identity/design.md'
  ':(exclude)openspec/changes/clean-break-ncode-identity/identity-inventory.md'
  ':(exclude)openspec/changes/clean-break-ncode-identity/tasks.md'
  ':(exclude,glob)**/legacy_zot_*'
)
git grep -nI -i -e zot -- "${audit_pathspecs[@]}"
git grep -nI -F 'github.com/patriceckhart/zot' -- "${audit_pathspecs[@]}"
git grep -nI -E 'ZOTCORE_|ZOT_[A-Z0-9_]+' -- "${audit_pathspecs[@]}"
git grep -nI -E '\.zot(session)?([^[:alnum:]_]|$)' -- "${audit_pathspecs[@]}"
git grep -nI -E 'zot[-_][[:alnum:]_-]+' -- "${audit_pathspecs[@]}"
git ls-files | grep -i zot | grep -vE '(^|/)legacy_zot_[^/]*$'

printf '%s\n' \
  docs/ncode-architecture.md \
  docs/inherited-capabilities.md \
  docs/superpowers/specs/2026-07-28-configurable-auto-compact-threshold-design.md \
  docs/superpowers/plans/2026-07-28-configurable-auto-compact-threshold.md \
  openspec/project.md \
  openspec/config.yaml \
  openspec/changes/clean-break-ncode-identity/proposal.md \
  openspec/changes/clean-break-ncode-identity/specs/ncode-identity/spec.md \
  openspec/changes/clean-break-ncode-identity/design.md \
  openspec/changes/clean-break-ncode-identity/identity-inventory.md \
  openspec/changes/clean-break-ncode-identity/tasks.md \
  > /tmp/ncode-zot-provenance-files.txt
git ls-files -- ':(glob)**/legacy_zot_*' > /tmp/ncode-zot-rejection-files.txt
cat /tmp/ncode-zot-provenance-files.txt /tmp/ncode-zot-rejection-files.txt \
  | LC_ALL=C sort -u > /tmp/ncode-zot-allowed-files.txt
git grep -Il -i -e zot -- . | LC_ALL=C sort -u > /tmp/ncode-zot-affected-files.txt
comm -23 /tmp/ncode-zot-affected-files.txt /tmp/ncode-zot-allowed-files.txt
git grep -nI -i -e zot -- $(cat /tmp/ncode-zot-provenance-files.txt) \
  | tee /tmp/ncode-zot-provenance-lines.txt
git grep -nI -i -e zot -- ':(glob)**/legacy_zot_*' \
  | tee /tmp/ncode-zot-rejection-lines.txt
```

- [ ] Add or complete focused no-live-provider retained-capability coverage for provider/auth resolution, print/stream/JSON/RPC v1, sessions, direct permissions/tools, swarm, extensions including idle auto-ready v2 ack, skills, themes, updater via `httptest`, and retained Telegram tests; run the affected packages, `go vet ./...`, and `make test`. <!-- sdd-owner: implementation -->
- [ ] Update `openspec/changes/clean-break-ncode-identity/identity-inventory.md` before the final search to add this exact `tasks.md` path to the reviewed provenance manifest only if it contains legacy terms solely as historical planning/rejection context; then verify every provenance line and every `legacy_zot_*` line is factual or asserts rejection/non-use, not active support. <!-- sdd-owner: implementation -->
- [ ] Run the exact final allowlist gates from `identity-inventory.md`: construct `audit_pathspecs` with each exact provenance file plus `openspec/changes/clean-break-ncode-identity/tasks.md` when admitted and only `:(exclude,glob)**/legacy_zot_*`; require zero output from `git grep -nI -i -e zot -- "${audit_pathspecs[@]}"`, old-module, `ZOTCORE_|ZOT_[A-Z0-9_]+`, `.zot(session)?`, and `zot[-_][[:alnum:]_-]+` searches, and from `git ls-files | grep -i zot | grep -vE '(^|/)legacy_zot_[^/]*$'`. <!-- sdd-owner: implementation -->
- [ ] Run the exact workable-file and mandatory-line reviews from `identity-inventory.md`: build `/tmp/ncode-zot-provenance-files.txt`, `/tmp/ncode-zot-rejection-files.txt`, and `/tmp/ncode-zot-allowed-files.txt`; require `comm -23` against `git grep -Il -i -e zot -- .` to emit no paths; emit and review every line into `/tmp/ncode-zot-provenance-lines.txt` and `/tmp/ncode-zot-rejection-lines.txt`; then run `go mod tidy`, `go list ./...`, `go vet ./...`, `make build`, `make test`, nested example builds, installer/package checks, and the inventory reproduction commands against planning SHA `18325b75cc89c75b5f4842924cb377aa5bef5c4b`. <!-- sdd-owner: implementation -->

## Parent delivery and review gates

- [x] Record the user's approved single-feature-branch delivery strategy: WU 1 is merged; keep WU 2–WU 12 as independently revertible commits on `feat/clean-break-ncode-identity-02`, push checkpoints as appropriate without intermediate PRs, and open one final PR to `main` linked to issue #1 with exactly `type:breaking-change`; the user accepts the combined final-PR size exception. <!-- sdd-owner: parent -->
- [x] Bounded review was explicitly disabled for this clone by the user after a committed-only controller defect; continue each work unit under strict TDD and native SDD attempt authority through focused verification and `make test`, retaining issue linkage, CI, final PR, and merge gates without starting review transactions unless the user re-enables them. <!-- sdd-owner: parent -->
