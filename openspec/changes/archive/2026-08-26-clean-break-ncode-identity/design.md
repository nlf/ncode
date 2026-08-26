# Technical design: make ncode the only active identity

## Decision summary

The implementation will make lowercase `ncode` the sole command, module, state, environment, integration, and release identity. This is a clean break: it adds no Zot executable, import, state, environment, protocol, installer, updater, extension, or SDK fallback. Factual Zot baseline/upstream/LICENSE material remains contextual provenance, and Zotfiles are removed in a separate capability-scoped work unit rather than renamed.

The runtime architecture remains the existing single construction spine:

```text
host-specific composer
  -> agent.Resolve(args, requireCredential)
  -> Resolved.NewClient()
  -> Resolved.NewAgent()
  -> host-specific sessions/events/extensions/output policy
```

Print, stream, JSON, and swarm-child already share the explicit `composeHeadlessAgent` composer and its focused production-shaped print test. That completed composition work is retained and renamed in place. RPC, SDK, interactive, bot, and management entry points remain explicit host composers over the same spine. This design does not introduce a generic all-mode constructor, a second core loop, or a provider/UI redesign.

The exhaustive occurrence classification and exact mapping are in [identity-inventory.md](identity-inventory.md). The current pre-rename planning snapshot `18325b75cc89c75b5f4842924cb377aa5bef5c4b` records 13 Zot-bearing filenames, 2181 content lines across 282 text files, 293 old-module lines, 310 Zot-environment lines, 106 dot-path lines, and 80 composed-name lines. The older frozen `zot-baseline` remains a separate historical comparison at 13/2098/278/285/301/105/79. Acceptance targets the planning snapshot plus later reviewed OpenSpec artifacts and uses the inventory's exact manifest rather than broad directory exclusions.

## Scope and boundaries

The behavior-changing center is `packages/agent`: command routing, resolution, state roots, sessions, RPC, extensions, swarm composition, update behavior, SDK construction, and Zotfile removal. Repository-wide mechanical changes are required for the Go module/import graph, command directory, examples, docs, installers, and release inputs. `packages/core`, `packages/provider`, and `packages/tui` change only where they expose product identity, branded portable/temp names, imports, or focused tests.

Retained capabilities remain provider-neutral core operation, all inherited providers/authentication paths, interactive behavior, print/stream/JSON/RPC, sessions, tools/permissions, swarm, extensions, skills, themes, updater behavior, and the inherited Telegram support level. The identity change does not alter provider wire schemas except product-owned header/user-agent/request identifiers, does not change general state schemas, and does not add live-provider tests.

## Architecture and construction contracts

### Shared construction spine

`Resolve` remains the only configuration-to-runtime resolver. It owns effective provider/model/credential/CWD/tool/skill/context/system-prompt state. `Resolved.NewClient` remains the only provider adapter selection point, and `Resolved.NewAgent` remains the only `core.NewAgent` construction point used by product hosts.

Explicit composers remain responsible for host behavior:

| Host | Construction contract after rename |
|---|---|
| Print/stream/JSON | Keep `composeHeadlessAgent`: `Resolve(true)` → extension setup/tool merge → `Resolved.NewAgent` → extension hooks. Each mode separately owns prompt projection and session flush. |
| Swarm child | Keep using `composeHeadlessAgent`, then add explicit child session, ncode swarm inbox, and JSONL emitter. The parent keeps self-location through `os.Executable`; it does not search for a `zot` or `ncode` command on PATH. |
| RPC | Keep an explicit RPC composer over `Resolve(true)` → extension manager/tool merge → `Resolved.NewAgent` → RPC server. Do not route RPC through a generic mode constructor. |
| Go SDK | Keep `sdk.New` as an explicit in-process composer over `agent.Resolve(true)` → `Resolved.NewAgent`; callers own persistence. Imports move to `github.com/nlf/ncode/packages/agent/sdk`. |
| Interactive | Keep explicit auth, extension, swarm, session, update/theme, and TUI composition around the shared spine. Identity work only renames fields/strings/paths. |
| Bot/management commands | Keep their explicit router/composer boundaries and use `NcodeHome`; no new composition abstraction. |

The existing `ExtensionToolSource` adapter remains in `packages/agent/build.go` to avoid a cycle between build composition and `packages/agent/extensions`. The module rename must not collapse this seam.

### Data flow

```text
CLI / SDK / subprocess input
  -> Args and NCODE_* controls
  -> NcodeHome and .ncode project discovery
  -> Resolve
       reads config/auth/model/context/skills
       attempts embedded-doc installation at the existing lifecycle point and ignores installation failure, as the baseline does
       builds tools/system prompt
  -> Resolved.NewClient
  -> Resolved.NewAgent
  -> explicit host composer
       sessions / RPC / extensions / swarm / output
  -> ncode-only state and diagnostics
```

No step probes, copies, imports, migrates, deletes, or reports Zot state as an active input. A Zot path may appear at runtime only as arbitrary user data or in a negative test; it is never constructed as a fallback.

## Exact product contracts

### Module, command, and symbols

The canonical module is `github.com/nlf/ncode`. `go.mod`, all repository imports, nested example modules, source snippets, and Go embedder documentation move atomically. No `replace github.com/patriceckhart/zot`, vanity redirect, compatibility package, or import alias remains.

The only command is `ncode`, built from `cmd/ncode` into `bin/ncode`; Windows uses `ncode.exe`. `cmd/zot`, `bin/zot`, a `zot` dispatch branch, and duplicate release assets do not exist. Product symbols become `NcodeHome`, `NcodeDocsDir`, and `NcodeVersion`; lower-level local identifiers follow the same naming. Product prose remains lowercase `ncode` even where Go identifiers require capitalization.

### State and project paths

`NcodeHome` resolves in this exact order:

1. non-empty `$NCODE_HOME`;
2. `$XDG_STATE_HOME/ncode` on every OS when XDG state is set;
3. macOS `~/Library/Application Support/ncode`;
4. Windows `%LOCALAPPDATA%\ncode`;
5. `~/.local/state/ncode` when a home directory is available;
6. `.ncode` only when no preceding root can be resolved.

`$ZOT_HOME` and all Zot defaults are ignored. Native project discovery uses `.ncode/extensions`, `.ncode/skills`, and the example MCP bridge's `.ncode/mcp.json`. `.zot` is never scanned. Portable session export uses `.ncodesession`; `.zotsession` is not inferred or accepted as the product portable extension.

The MCP bridge has an independent resolver in `examples/extensions/mcp-bridge/config.go`; its Windows baseline uses `%APPDATA%\zot`, unlike the host's `%LOCALAPPDATA%\zot`. It must not preserve that divergence. Rename/refactor the bridge resolver and its tests so it implements the same ordered `NcodeHome` contract above, including `%LOCALAPPDATA%\ncode` on Windows, and reads only global `$NCODE_HOME/mcp.json` plus project `.ncode/mcp.json`.

Non-branded basenames and schemas remain unchanged: `config.json`, `auth.json`, `models.json`, `models-cache.json`, `update-check.json`, `AGENTS.md`, `SYSTEM.md`, sessions JSONL, docs, skills, themes, extensions, logs, bot files, extension manifests, and swarm `meta.json`/`events.jsonl`/`session.json`.

### Environment controls

Every product control uses `NCODE_*`. The full map is in the inventory. In particular, `ZOTCORE_RPC_TOKEN` becomes `NCODE_RPC_TOKEN`, not `NCODE_CORE_RPC_TOKEN`. Zotfile consent is deleted with the capability rather than mapped. Legacy names are not consulted even when their ncode counterpart is absent; when both are set, only ncode can affect behavior.

Provider, OS, proxy, terminal, and cloud variables remain unchanged because they are not product identity. This includes `ANTHROPIC_API_KEY`, `OPENAI_API_KEY`, `GITHUB_TOKEN`, `XDG_STATE_HOME`, `XDG_RUNTIME_DIR`, proxy variables, and equivalent provider controls.

### RPC protocol

RPC keeps its existing neutral version-1 JSONL frame behavior and optional-hello semantics. Identity work is limited to the product-bearing edges:

- process invocation becomes `ncode rpc`;
- the optional token control becomes `NCODE_RPC_TOKEN`;
- when `NCODE_RPC_TOKEN` is unset, an otherwise valid current prompt can remain the first frame;
- when `NCODE_RPC_TOKEN` is set, the existing first-frame `hello` token check remains required and its neutral response remains protocol version 1;
- `ZOTCORE_RPC_TOKEN` is never read, so setting it alone neither enables authorization nor gates an otherwise unauthenticated neutral client;
- no mandatory product field, always-on hello, RPC protocol version 2, product negotiation, or rejection of current neutral prompt/event shapes is introduced;
- command, response, event, optional hello, provider, model, and version fields retain their current neutral names and behavior; only docs, examples, invocation strings, token naming, and product-bearing diagnostics are rebranded.

Unknown extra JSON fields continue to have their normal neutral decoding behavior; an explicit Zot-branded field has no authorization or configuration semantics and is not documented as a supported identity. Shell, Go, Python, and Node reference clients update invocation/token/product prose together without changing their neutral frame contract.

### Extension protocol and SDK

The retained subprocess extension protocol has a justified, signaled version-2 rebrand because the current host acknowledgement is product-bearing through `zot_version`:

- increment shared `extproto.ProtocolVersion` to 2;
- keep the extension-originated `hello` shape neutral and otherwise unchanged;
- host `hello_ack` includes `product:"ncode"`, `protocol_version:2`, and `ncode_version`, and removes `zot_version`;
- Go SDK `HostInfo.NcodeVersion` replaces `ZotVersion`, and SDK/raw clients validate the signaled ncode v2 acknowledgement rather than accepting the old product-bearing acknowledgement;
- preserve the existing idle auto-ready lifecycle, timeout, readiness channel behavior, and diagnostic state; change only its product-facing log prefix/text from Zot to ncode;
- neutral manifest basename/fields and frame names (`extension.json`, `hello`, `ready`, `register_tool`, `tool_call`, and similar) remain unchanged.

Host, shared `extproto`, Go SDK, manager tests, raw-language examples, authoring skill, and docs change in one protocol work unit so there is no mixed product-bearing acknowledgement in a buildable revision. The v2 change is not permission to remove unrelated readiness behavior or redesign neutral frames.

### Swarm and internal identifiers

The parent still starts the current executable returned by `os.Executable`, preserving the existing same-binary composition. Child argv, session ownership, JSONL events, and inbox commands remain neutral. Product metadata becomes `NCODE_SWARM_AGENT_ID`, `NCODE_SWARM_EVENT_LOG`, and `NCODE_SWARM_CREDENTIAL_STDIN`; socket directories become `ncode-swarm-<root-hash>`. A child given only Zot swarm variables follows normal fresh ncode credential/log behavior and does not read or write the Zot-provided event path.

Branded temp/log/request identifiers become `ncode-*` as inventoried. Product-owned Codex/Kimi originator/user-agent/request IDs become ncode. Provider-required first-party compatibility fingerprints remain provider-owned and unchanged.

### Distribution and updates

GoReleaser owns `ncode`, builds `cmd/ncode`, emits `ncode_<version>_<os>_<arch>.{tar.gz,zip}`, and publishes to `nlf/ncode`. `checksums.txt` remains unchanged because it is non-branded. Both installers use owner `nlf`, repository/binary `ncode`, and `NCODE_VERSION`/`NCODE_PREFIX`. The updater and changelog APIs query only `nlf/ncode`, extract only `ncode[.exe]`, and offer no Zot endpoint fallback.

The release workflow retains its CI/tag orchestration but validates only ncode assets. The generic CI workflow remains identity-neutral. The tracked generated example binary and Python bytecode artifact are deleted; ignored `bin/`, `dist/`, Python caches, and local GoReleaser output are regenerated only for validation, never committed.

### User communication

All active help, auth pages, system prompts, labels, logs, errors, README/contributor guidance, subsystem docs, issue templates, examples, and expected test text use lowercase ncode. The clean-break notice appears in all required channels:

| Channel | Location and content |
|---|---|
| Release notes | `.goreleaser.yaml` release header states that ncode does not reuse Zot credentials, settings, sessions, caches, extensions, SDK/RPC/swarm contracts, or other state. |
| Installer/install docs | `install.sh`, `install.ps1`, and README Install state the same break before/at installation and point only to ncode setup. |
| First-run/setup docs | README's first-run/authentication section prominently explains fresh ncode state and integration setup. |

There is no runtime migration, compatibility, or conversion prompt. Auth pages and TUI startup simply present ncode.

## Zotfile removal design

Zotfiles are removed before broad renaming as a standalone functional unit. This prevents a mechanical replacement from accidentally inventing ncodefiles.

The unit deletes the implementation, dedicated tests, guide, command router branches (`pack`, `inspect`, `verify`, `run`), archive/manifest/consent wiring, `-y/--yes` usage that exists only for Zotfiles, and Zotfile-only help/README references. It removes `ZOT_AGENT_CONSENT` with no replacement. Existing Zotfile state and archives are left untouched on disk.

The unit deletes only Zotfile-specific host behavior: `Args.AgentName`, `Args.AgentDataDir`, `Args.StartupPre`, entry-pre execution/reload hooks, consent receipts/prompts, agent-data creation, startup-agent display, and named-agent session-root selection. It does **not** delete the general permission contract.

The unit explicitly preserves:

- read/write/edit/bash tools and `tools.Sandbox`;
- `Args.PermissionSet` in `packages/agent/args.go` and the exact `args.PermissionSet` → `sandbox.SetPermissions(args.PermissionSet)` wiring in `packages/agent/build.go`;
- `tools.PermissionSet`, `Sandbox.SetPermissions`, and focused permission tests that construct `Args.PermissionSet` independently of Zotfiles;
- normal project/global skills and the `skill` tool;
- ordinary sessions, resume, import/export, fork/tree, pruning, and JSONL format;
- extension subprocesses and confirmation behavior.

No file, type, command, artifact extension, environment variable, or documentation named ncodefile is created.

## Baseline state-creation timing

Identity changes preserve when state appears, not only where it appears.

| State category | Baseline creation timing to preserve under ncode | Evidence |
|---|---|---|
| Root/path resolution | `NcodeHome`, `ConfigPath`, `AuthPath`, model paths, and theme/skill/extension search path calculation create nothing. | `packages/agent/config.go`; read-only loaders. |
| Embedded docs | `Resolve` eagerly attempts `docs.EnsureInstalled`, but intentionally ignores its returned path error (`docsDir, _ := ...`). With a writable root this creates the product root, `docs/`, and embedded active docs for SDK/RPC/headless resolution; when installation fails, `Resolve` can still succeed without docs. Identity work preserves both the eager attempt and ignored-error behavior. | `packages/agent/build.go`; `docs.go`. |
| Config | Missing `config.json` is read as defaults. It is created only by an explicit save/repair/settings/auth-related lifecycle; first interactive changelog seeding may save it after startup. Merely calling `NcodeHome` does not create it. | `LoadConfig`, `SaveConfig`, `ValidateAndRepairConfig`, `SeedChangelogVersion`. |
| Auth | Missing `auth.json` is read as empty and created only by login/credential persistence/refresh. | `packages/provider/auth/store.go`; `AuthStoreFor`. |
| Models | `models.json` is read-only user input. `models-cache.json` is written only after successful non-empty background discovery when the cache is stale. | `modelsync.go`; `provider/cache.go`. |
| Normal sessions | With sessions enabled and a resolved agent, interactive and print/stream/JSON call `openOrCreateSession` before the first prompt and write a metadata row. Closing a fresh session with no messages deletes the file but leaves created directories. RPC and SDK do not create normal sessions by default. | `cli.go`; `core/session.go`; `sdk.go`; `rpc.go`. |
| Extensions/logs | Missing discovery directories are only read. Loading an executable extension creates `logs/` and `ext-<name>.log` at spawn; theme-only discovery does not need a process log. | `extensions/manager.go`. |
| Swarm | `swarm.New` and `Reload` are read-only when state is absent. `SpawnReq` creates the per-agent state directory and socket runtime directory/probe; runner opens `events.jsonl`; child opens `session.json`. | `swarm/{swarm,persist,runner,socketpath}.go`; `swarm_agent.go`. |
| Update/changelog | Update cache is written only after a successful release query for a non-development version. Update temp directories exist only during explicit update. Changelog config write follows current first-launch/dismiss timing. | `update.go`; `updatecmd.go`; `changelog.go`. |
| Themes/skills/context | Discovery reads only existing files. No theme, skill, `AGENTS.md`, or `SYSTEM.md` is created by discovery. | `skills/skills.go`; `build.go`; `tui/theme_loader.go`. |
| Telegram | `bot.json` is created by successful setup/persistence; pid/log files appear only during run/start lifecycle. Status-only reads do not create them. | `modes/telegram`; `botcmd.go`. |
| Zotfile data/consent | Baseline Zotfile launch eagerly creates agent data before consent and may write a receipt. The removal unit eliminates those writes entirely and creates no replacement. | `zotfile.go` baseline; removal decision. |
| Portable sessions | Export creates a destination only when requested; import creates a normal session file only when requested. The suffix changes to `.ncodesession`, timing and JSONL shape do not. | `core/session_portable.go`. |

### Timing tests

Focused tests record the filesystem tree before and after each lifecycle. They assert both positive and negative timing:

- path resolution alone creates no directories;
- with a writable root, `Resolve` creates the same embedded-doc category at the same point under ncode;
- with a deterministic install failure (for example, a parent path that is a regular file), `Resolve` preserves the ignored error and can succeed without a docs directory;
- missing config/auth/models/themes/skills remain missing until their existing trigger;
- interactive/headless session metadata is created before first prompt and an empty fresh file is removed on close while RPC/SDK remain sessionless by default;
- extension log, swarm state/socket, update cache, bot files, portable export/import, and config/auth writes occur only on their existing triggers;
- no test performs a real release/provider call: updater/model discovery use local servers or direct write helpers, and provider calls use fakes/`httptest`.

## Coexistence and negative verification

A reusable isolated-state fixture creates distinguishable ncode and Zot roots plus project `.ncode` and `.zot` trees. Before running ncode it snapshots relative paths, content hashes, permissions, sizes, and modification times for the Zot tree; after the lifecycle the snapshot must be identical.

The coexistence matrix covers config, auth, sessions, model cache/models, logs, extensions, skills, themes, docs, update cache, swarm agent data, bot data, project-local resources, portable sessions, temp/event paths, and extension logs. Sentinel data makes accidental reads observable: Zot config selects a different provider/model, Zot auth has a synthetic unique credential, Zot skills/system/context inject poison text, and a Zot extension would create a marker if spawned. On Unix, selected legacy roots may additionally be unreadable to turn probing into a deterministic failure; Windows relies on sentinel behavior and snapshots.

Negative tests prove:

- `$ZOT_HOME` alone does not change `NcodeHome`; conflicting `$NCODE_HOME` wins without reading Zot;
- Zot OS-default roots and `.zot` project resources are neither read nor modified, and no file is copied to ncode;
- every old product environment variable alone has no effect, and conflicts resolve exclusively to `NCODE_*`;
- `ZOTCORE_RPC_TOKEN` does not enable or gate ncode RPC; `NCODE_RPC_TOKEN` retains the existing optional first-frame token hello behavior, and current neutral prompt/event/hello frames remain interoperable at RPC version 1;
- extension clients do not accept the old product-bearing v1/`zot_version` acknowledgement, while neutral extension hello/frame shapes and idle auto-ready behavior remain available under the signaled ncode v2 acknowledgement;
- Zot swarm metadata does not supply credentials or an event log and no `zot-swarm-*` socket path is created;
- `.zotsession` is not treated as `.ncodesession`;
- no `zot` binary, command directory, import compatibility, module replace, release asset, installer path, updater endpoint, or command alias exists;
- Zotfile commands/artifacts are absent while normal tools, permissions, skills, sessions, extensions, and headless modes remain green.

Legacy literals used by tests must be isolated in dedicated, auditable files or fixtures whose basename begins `legacy_zot_` (for Go, for example `legacy_zot_home_test.go`). Ordinary behavior tests are renamed to ncode and contain no Zot literal. Each dedicated file states and asserts rejection, ignore, or non-use; the final audit emits every line from those files for review rather than treating the directory containing them as an exception.

## File-change plan

| Area | Planned change |
|---|---|
| `packages/agent` | Remove Zotfiles; rename state/symbol/env/help/update/RPC/swarm wiring; preserve shared spine and explicit composers; add state/protocol negative tests. |
| `packages/agent/{extensions,extproto,ext}` | Signaled protocol-v2 ncode host acknowledgement and SDK version field; preserve neutral extension hello/frames and idle auto-ready; update manager/SDK tests. |
| `packages/agent/{skills,swarm,tools,modes,sdk}` | `.ncode` discovery, NCODE metadata, ncode temp/log/public comments, `.ncodesession` UI, ncode SDK comments; preserve behavior. Work unit 9 exclusively owns branded temp/log/socket changes even when files overlap state-related packages. |
| `packages/core` | Canonical imports, `.ncodesession`, product comments/negative tests only; no loop/session schema redesign. |
| `packages/provider` and `packages/provider/auth` | Canonical imports, ncode product headers/request/temp names/auth pages/logo/package comments; provider-required external identity unchanged. |
| `packages/tui` | Canonical imports, ncode labels/env/temp names/comments/tests; no UI replacement. |
| `cmd`, root module/build/docs | Move `cmd/ncode`, update `go.mod`, `docs.go`, Makefile, README, AGENTS, CONTRIBUTING, LICENSE handling. |
| `examples` | Ncode process/import/env/handshake/project paths; nested module cleanup; remove the tracked example binary and Python bytecode artifact. |
| `.goreleaser.yaml`, installers, workflows | Ncode owner/repo/binary/assets/update/install/release notice; generic workflow logic retained. |
| architecture/SDD records | Retain and contextualize factual Zot provenance; do not blanket-replace. |

## Dependency-aware review work units

These are design units, not implementation task checkboxes. Each unit must leave the repository buildable and keeps tests/docs with the behavior it changes.

| Order | Work unit | Type and dependency | Focused validation | Review forecast / rollback boundary |
|---:|---|---|---|---|
| 1 | Remove Zotfiles without replacement | Functional, capability-scoped; first so broad rename cannot create ncodefiles. Delete only Zotfile `AgentName`/`AgentDataDir`/`StartupPre`, consent, agent-data, startup-display, and named-session behavior; retain `Args.PermissionSet` → `Sandbox.SetPermissions`. | Agent router/help tests; direct `Args.PermissionSet` sandbox enforcement tests in `args.go`/`build.go`; tools permission/sandbox tests; skills tests; ordinary core/session and mode session tests; searches for Zotfile symbols/docs/commands. | **Expected over 400 changed lines** because implementation, large tests, and guide are deleted. Ask-on-risk must be surfaced by the orchestrator later; this design does not choose. Revert this unit alone before release. |
| 2 | Move canonical module/import graph | Mechanical and atomic; depends on removed files no longer needing edits. | `go mod tidy`; nested-module tidy; `go list ./...`; `go test ./...`; search old module outside provenance; no `replace` compatibility. | **Expected over 400 changed lines** due repository-wide imports. Atomic rollback is the whole unit. |
| 3 | Rename Go product symbols and package comments | Mechanical; depends on canonical imports. No runtime values change yet beyond names. | Focused package builds/tests for agent/core/provider/tui; `go vet ./...`; search public `ZotHome/ZotVersion/ZotDocsDir`. | Likely near/over 400; monitor and flag if forecast crosses budget. Revert independently while imports stay ncode. |
| 4 | Move command/build identity | Mechanical distribution slice: `cmd/ncode`, Makefile, version/error text, no alias. | `go build -o /tmp/ncode ./cmd/ncode`; `go install` with isolated `GOBIN`; assert no `/tmp/zot`/`cmd/zot`/`bin/zot`; command help/version tests. | May fit under 400 if release/docs stay later. Revert command/build slice as a unit. |
| 5 | Establish ncode-only state, project, and portable paths | Functional: `NCODE_HOME`, exact OS defaults, `.ncode` discovery, and `.ncodesession`; depends on symbol/command names. This unit owns home/project/portable paths only, including making the MCP bridge use the exact host resolver contract and `%LOCALAPPDATA%\ncode` on Windows. | Host and MCP table-driven path-resolution tests; writable-root and failed-doc-install timing tests; coexistence snapshot tests; extension/skill/MCP project path tests; portable import/export negative tests; no legacy path probes. | **Expected over 400 changed lines** across code/tests. Keep state behavior and tests together; rollback restores this entire pre-release slice only. |
| 6 | Establish coherent NCODE environment namespace | Functional but narrow; depends on state root. Zotfile consent is already gone. | Table-driven positive/legacy-only/conflict tests for all 23 names; focused auth/TUI/swarm tests; search old env outside negative/provenance. | Likely under 400 if installer vars remain distribution docs; rollback env map/tests together. |
| 7 | Rebrand RPC edges and update all clients | Functional contract slice; depends on command/env. Preserve neutral RPC v1 frames and optional hello. | Token-unset prompt-first test; `NCODE_RPC_TOKEN` hello success/failure tests; `ZOTCORE_RPC_TOKEN` ignored test; unchanged neutral hello/response/event assertions; shell/Go/Python/Node syntax/build checks; no live provider. | Could approach 400 with four examples/docs; flag if over. Rollback server, clients, docs together before release. |
| 8 | Rebrand extension protocol/SDK acknowledgement as signaled v2 | Functional protocol slice; depends on module and project paths. Preserve cycle adapter and idle auto-ready. | `extproto`, manager, SDK, tool/interceptor tests; exact ncode v2 ack tests; old product-bearing ack rejection; unchanged neutral hello/frame and auto-ready tests; build nested extension modules. | **Expected over 400 changed lines** across host, SDK, tests, examples, docs/skill. Atomic product-bearing acknowledgement rollback only before release. |
| 9 | Rename swarm/subprocess/provider/temp/internal identity | Mixed narrow functional/mechanical slice; depends on env/protocol. This is the sole owner of branded temp/log/request/socket names; unit 5 does not own them. | Swarm runner/child/socket/event tests, credential-stdin negative test, provider request/header tests, exact branded temp/log/request/socket path assertions; neutral event/control shapes remain unchanged; no live provider. | Monitor near 400. Revert as one parent/child/internal-name contract slice. |
| 10 | Replace release/update/install identity | Functional distribution slice; depends on command and owner/repo decisions. | GoReleaser config check/snapshot when available; Make release/package asset-name tests; installer static/syntax tests; updater local HTTP/unit tests; checksum verification; no network release call. | Likely over 400 when both installers and tests are included; flag if so. Rollback entire distribution slice before release. |
| 11 | Rewrite active product docs/examples/issues and publish clean-break notices | Mechanical communication slice after contracts settle; provenance docs reviewed separately. | Link/command searches; nested example builds; installer/help snippets; verify three notice channels; forbidden-active-identity audit. | **Expected over 400 changed lines**, especially README/subsystem docs. Split by audience only if each slice stays accurate; do not decide ask-on-risk here. |
| 12 | Final retained-capability and identity audit | Verification-only corrections; depends on every prior unit. | Module/build/install/package checks, all negative matrices, retained mode/session/tool/swarm/extension/skill/theme/update/Telegram tests, `go vet ./...`, final `make test`, final allowlist searches. | Keep corrections scoped to the owning earlier unit where possible. After release, rollback is corrective ncode-only release, never Zot fallback. |

Functional units are deliberately separate from broad import/symbol/docs mechanics. If an over-budget unit is later split, every slice must still compile and its protocol/state tests must travel with the behavior; the `ask-on-risk` choice belongs to orchestration/tasks, not this design.

## Validation strategy

### Per-unit minimum

Every production-code unit runs the smallest owning package tests, `go test` for directly dependent packages, `go vet` when process/concurrency/interfaces change, and the unit's forbidden-identity search. Mechanical module/command units additionally run `go list ./...`, build the command, and inspect `go env GOBIN`/an isolated install destination. Nested example modules are validated from their own directories.

Distribution validation checks:

- `go build` and isolated `go install github.com/nlf/ncode/cmd/ncode` behavior;
- Makefile outputs and cross-compile names;
- GoReleaser configuration/snapshot asset list where the tool is available;
- archive contents include `ncode[.exe]`, README, and LICENSE, never `zot`;
- `checksums.txt` covers each ncode archive and installers/updater derive the same exact name;
- shell syntax (`bash -n install.sh`) and PowerShell parser/static execution with mocked HTTP/filesystem where available;
- workflows refer to the correct build/release inputs.

Protocol/state validation uses fakes, local files, local subprocess stubs, and `httptest`. It never calls a paid or live provider. No credentials appear in fixtures.

### Final gate

The final gate is:

```sh
go mod tidy
go list ./...
go vet ./...
make build
make test
```

It also runs nested example module tests/builds, installer/package checks, state/env/protocol negative suites, the inventory commands, and the final allowlist searches from `identity-inventory.md`. `make test` is authoritative and runs `go test -race ./...`.

## Risks and controls

| Risk | Control |
|---|---|
| Mechanical replacement damages provenance or provider facts | Use tracked-file scan plus the exact reviewed provenance manifest; emit every retained line and review architecture/LICENSE changes separately from active docs. |
| A runtime-composed path/env/header survives text-oriented review | Inventory creation/composer call sites and assert exact composed values in focused tests. |
| State clean break accidentally reads or mutates Zot data | Coexistence sentinels, unreadable roots where portable, before/after snapshots, conflict tests, and no migration/probe code. |
| Creation timing changes because path helpers are centralized | Preserve the eager docs-install attempt and ignored failure separately, plus every lazy lifecycle and pre-prompt session point. |
| RPC clients are broken by an unnecessary identity handshake | Preserve neutral RPC v1 frames and optional hello; rename only invocation, token control, examples/docs, and product-bearing diagnostics. |
| Extension product identity remains ambiguous or readiness regresses | Ship the product-bearing ncode v2 acknowledgement atomically while retaining neutral frame shapes and the existing idle auto-ready semantics. |
| Swarm parent/child mismatch | Change command/env/socket/tests in one contract unit; retain `os.Executable` self-location and explicit child composer. |
| Package cycle introduced while renaming extension composition | Preserve `ExtensionToolSource`/adapter; extproto stays dependency-light; SDK imports agent but agent never imports SDK. |
| Zotfile removal deletes shared safety/session behavior | Dedicated removal assertions keep sandbox, PermissionSet, skills, normal sessions, extensions, and tools green. |
| Release assets/installers/updater disagree | One exact asset template, shared expectations in updater tests, GoReleaser snapshot/archive inspection, installer checksum tests. |
| Review budget hides errors | Mechanical/functional separation, dependency-ordered units, explicit >400 flags, and later ask-on-risk handling by orchestrator. |

## Generated artifacts and package-cycle constraints

Generated/release outputs are `bin/`, `dist/`, GoReleaser archives, checksums, cross-compiled binaries, and temporary updater/install directories. They are regenerated for validation and remain untracked. Two tracked generated artifacts are removed: `examples/extensions/todo/zot-todo-extension` and `examples/rpc/python/__pycache__/zot_client.cpython-314.pyc`; neither receives a renamed tracked binary counterpart. `go.mod` and nested module metadata are source inputs; `go.sum` changes are accepted only when `go mod tidy` requires them, not from unrelated dependency upgrades. The ncode logo is a source embed input and must be committed with its `go:embed` rename.

Cycle constraints remain:

- `packages/core` may import `packages/provider`; provider must not import core;
- `packages/agent/extensions` and `packages/agent/ext` share `extproto` but do not own host composition;
- `packages/agent/build.go` must not import extensions directly, preserving the adapter seam;
- `packages/agent/sdk` imports agent; agent must not import sdk;
- `packages/agent/swarm` must not import agent and continues to cross the boundary through the executable protocol;
- root `docs` may be imported by agent for embedding, but docs must not import agent.

## Rollout and rollback

This is a single clean-break release assembled from buildable pre-release units. Before tagging, any unit can be reverted at its stated boundary in reverse dependency order. State and protocol units are never partially reverted because mixed identities would create ambiguous behavior.

Release notes, installers, and first-run docs ship with the first ncode-only release. No runtime prompt or migration period exists. Existing Zot installations and state remain untouched and independently usable by their historical binary.

After release, rollback means a corrective release under the ncode module/repository/binary/state/protocol identity. It must not restore a Zot command, module replace, state import/probe, old environment fallback, Zot-branded extension acknowledgement/SDK field, duplicate asset, or Zotfile capability. Neutral RPC version-1 frames and optional hello remain part of the retained contract throughout. A serious defect may disable the affected ncode capability temporarily, but remediation remains ncode-only.
