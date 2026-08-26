# Exhaustive identity inventory: clean-break ncode identity

This inventory is the review index for replacing active Zot identity with lowercase `ncode`. Its classification rule is exhaustive: only the exact reviewed provenance manifest below is **retain as provenance**; the Zotfile capability and every other current active match are **rename/remove**; identity-free surfaces and dedicated `legacy_zot_*` fixtures whose sole purpose is proving rejection/non-use are **not applicable**. No fourth classification exists.

## Reproduce the inventory

The implementation inventory is pinned to the current pre-rename planning snapshot `18325b75cc89c75b5f4842924cb377aa5bef5c4b`. That snapshot includes the completed architecture, characterization, composition, and `.codegraph`-ignore commits, but predates tracking these OpenSpec change artifacts. The older `zot-baseline` tag remains a separate historical comparison; it is not the implementation inventory target.

Run from the repository root against the pinned refs so dotfiles, workflows, nested example modules, fixtures, and tracked generated/release inputs are included while `.git/`, ignored build output, local state, and later planning artifacts are excluded.

```sh
PLANNING_REF=18325b75cc89c75b5f4842924cb377aa5bef5c4b
FROZEN_ZOT_REF=zot-baseline

# Tracked path universe and identity-bearing path count for implementation planning.
git ls-tree -r --name-only "$PLANNING_REF" | LC_ALL=C sort
git ls-tree -r --name-only "$PLANNING_REF" \
  | grep -i 'zot' \
  | tee /tmp/ncode-identity-paths.txt \
  | awk 'END { print NR }'                              # 13

# Every textual occurrence and affected tracked text file in the planning snapshot.
git grep -nI -i -e 'zot' "$PLANNING_REF" -- . \
  | tee /tmp/ncode-identity-lines.txt \
  | awk 'END { print NR }'                              # 2181
git grep -Il -i -e 'zot' "$PLANNING_REF" -- . \
  | LC_ALL=C sort -u \
  | tee /tmp/ncode-identity-files.txt \
  | awk 'END { print NR }'                              # 282

# Canonical module/import, environment, dot-path, and composed-name line counts.
git grep -nI -F 'github.com/patriceckhart/zot' "$PLANNING_REF" -- . | awk 'END { print NR }' # 293
git grep -nI -E 'ZOTCORE_|ZOT_[A-Z0-9_]+' "$PLANNING_REF" -- . | awk 'END { print NR }'    # 310
git grep -nI -E '\.zot(session)?([^[:alnum:]_]|$)' "$PLANNING_REF" -- . | awk 'END { print NR }' # 106
git grep -nI -E 'zot[-_][[:alnum:]_-]+' "$PLANNING_REF" -- . | awk 'END { print NR }'      # 80

# Runtime composition and creation-site audit, including names not caught by a plain text replacement.
git grep -nI -E 'Resolve\(|NewClient\(|NewAgent\(|composeHeadlessAgent' "$PLANNING_REF" -- '*.go'
git grep -nI -E 'Mkdir(All|Temp)|OpenFile|Create(Temp)?|WriteFile|Rename|Remove(All)?' "$PLANNING_REF" -- '*.go'
git grep -nI -E 'filepath\.Join|os\.(Getenv|Setenv)|exec\.(Command|CommandContext)|go:embed' "$PLANNING_REF" -- '*.go'
```

| Measure | Planning snapshot | Frozen `zot-baseline` |
|---|---:|---:|
| Zot-bearing filenames | 13 | 13 |
| Case-insensitive Zot-bearing content lines | 2181 | 2098 |
| Affected tracked text files | 282 | 278 |
| Old-module lines | 293 | 285 |
| Zot-environment lines | 310 | 301 |
| Dot-path lines | 106 | 105 |
| Composed-name lines | 80 | 79 |

The binary `.pyc` path contributes to both filename counts but is excluded from textual counts by `git grep -I`. The planning snapshot contains 23 unique Zot-prefixed environment names; 310 is its matching line count, not its unique-name count. Future tracked OpenSpec artifacts are reviewed through the exact provenance manifest below rather than folded retroactively into either pinned snapshot.

## Classification precedence and exact manifests

| Classification | Exact rule | Evidence |
|---|---|---|
| retain as provenance | A match is retained only in an exact file in the reviewed provenance manifest below, and every remaining line in that file must identify the inherited baseline, upstream/source policy, historical evidence, clean-break planning/rejection language, or legal attribution. | Exact manifest below; there is no directory wildcard. |
| rename/remove | The Zotfile capability and every tracked path/content match not admitted by the exact manifest are removed or renamed to the mapping below. Existing tests that currently assert active Zot behavior are in this class, not negative fixtures. | All active code, tests, docs, examples, installers, release inputs, issue templates, and package metadata listed below. |
| not applicable | Identity-free tracked surfaces remain unchanged. After implementation, legacy literals used to prove rejection/non-use are permitted only in a dedicated tracked test/fixture whose basename starts `legacy_zot_`; every such line is emitted for review. | Provider/OS variables, neutral state basenames and protocol shapes, plus dedicated `legacy_zot_*` test/fixture files only. |

### Exact reviewed provenance manifest

| Exact tracked file | Reason Zot literals may remain |
|---|---|
| `docs/ncode-architecture.md` | Product-owned architecture records the inherited fork baseline, settled no-compatibility disposition, and `patriceckhart/zot` upstream-watch policy. |
| `docs/inherited-capabilities.md` | Frozen capability evidence records exact Zot-era commands, paths, variables, composition, and the baseline commit for comparison; it is not current user guidance. |
| `docs/superpowers/specs/2026-07-28-configurable-auto-compact-threshold-design.md` | Historical Zot-era implementation specification with original `$ZOT_HOME` and product wording; add an explicit historical-record header and do not present it as current setup. |
| `docs/superpowers/plans/2026-07-28-configurable-auto-compact-threshold.md` | Historical Zot-era implementation plan with original test commands and values; add the same explicit historical-record header. |
| `openspec/project.md` | Project context identifies the inherited Zot fork and settled removal/clean-break decisions. |
| `openspec/config.yaml` | Change-planning metadata names the Zot-to-ncode identity goal. |
| `openspec/changes/clean-break-ncode-identity/proposal.md` | Approved change proposal records baseline evidence, unsupported legacy identities, and provenance policy. |
| `openspec/changes/clean-break-ncode-identity/specs/ncode-identity/spec.md` | Corrected normative specification names legacy inputs only to require rejection/non-use and names factual provenance. |
| `openspec/changes/clean-break-ncode-identity/design.md` | Technical design records baseline mappings, removals, negative cases, and exact rollout constraints. |
| `openspec/changes/clean-break-ncode-identity/identity-inventory.md` | This reviewed inventory necessarily records baseline literals, mappings, audit expressions, and allowed-evidence rules. |
| `openspec/changes/clean-break-ncode-identity/tasks.md` | Approved implementation planning records old-path characterization commands, clean-break rejection constraints, and final audit expressions; it does not document active compatibility. |

`LICENSE` remains byte-for-byte MIT legal attribution and currently contains no Zot literal, so it is not excluded from the final search. Any future source notice containing a Zot literal must be added as an exact reviewed manifest entry rather than admitted through a wildcard. Local `.git/refs/tags/zot-baseline` is untracked repository metadata identifying the frozen baseline; `.git/` is outside tracked-file audits and is not rewritten.

Historical wording in active user guides is not automatically provenance. `README.md`, `CONTRIBUTING.md`, `AGENTS.md`, `docs/providers.md`, `docs/extensions.md`, `docs/rpc.md`, `docs/skills.md`, and `docs/themes.md` describe supported behavior and are therefore rename/remove. No `openspec/**`, `docs/superpowers/**`, test-directory, or fixture-directory wildcard is a provenance exception.

## Identity-bearing tracked filenames

The frozen baseline has 13 tracked Zot-bearing paths. Every row is **rename/remove**; the untracked local Git tag is separately retained as repository provenance.

| Baseline path | Classification | Target |
|---|---|---|
| `cmd/zot/main.go`, `cmd/zot/main_test.go` | rename/remove | Move to `cmd/ncode/`; no `cmd/zot` directory or dispatch remains. |
| `docs/zotfiles.md` | rename/remove | Delete with the Zotfile capability; do not create `ncodefiles.md`. |
| `packages/agent/zotfile.go`, `zotfile_test.go`, `zotfile_pre_test.go` | rename/remove | Delete capability implementation and capability-only tests. |
| `packages/agent/skills/builtin/write-zot-extension/SKILL.md` | rename/remove | Move to `write-ncode-extension/SKILL.md` and rewrite the retained extension-authoring capability. |
| `packages/agent/skills/builtin/write-zot-themes/SKILL.md` | rename/remove | Move to `write-ncode-themes/SKILL.md` and rewrite active identity. |
| `examples/rpc/node/zot-client.js` | rename/remove | `examples/rpc/node/ncode-client.js`. |
| `examples/rpc/python/zot_client.py` | rename/remove | `examples/rpc/python/ncode_client.py`. |
| `examples/rpc/python/__pycache__/zot_client.cpython-314.pyc` | rename/remove | Delete the tracked generated bytecode; do not track a renamed `.pyc`. |
| `examples/extensions/todo/zot-todo-extension` | rename/remove | Delete the tracked generated binary; the source-run example remains. |
| `packages/provider/auth/assets/zot-logo.png` | rename/remove | Replace with `ncode-logo.png`, update `go:embed`, alt text, auth pages, and README; the asset must not retain a Zot `z` mark as active art. |


## Canonical module, imports, command, and Go symbols

| Occurrence/pattern group | Classification | Evidence paths | Exact target |
|---|---|---|---|
| Root module and all repository-owned imports | rename/remove | `go.mod`; every `*.go`; root `docs.go`; `examples/sdk`; `examples/extensions/*`; nested `go.mod` files; authoring skill snippets | `github.com/patriceckhart/zot` → `github.com/nlf/ncode`, including `github.com/nlf/ncode/cmd/ncode`. No `replace` or compatibility module remains. |
| Nested example modules | rename/remove | `examples/extensions/mcp-bridge/go.mod`; `examples/extensions/todo/go.mod` | Module names become `github.com/nlf/ncode/examples/extensions/{mcp-bridge,todo}`; dependency/replace paths use `github.com/nlf/ncode`; remove the absolute `/Users/pat/Developer/zot` replace. |
| Command directory/binary | rename/remove | `cmd/zot`; `Makefile`; `.goreleaser.yaml`; installers; updater; examples; docs | `cmd/ncode`, `ncode`, `bin/ncode`, platform artifacts `ncode-<os>-<arch>` / `ncode.exe`; no alias or shim. |
| Public product symbols and fields | rename/remove | `packages/agent/config.go`; `build.go`; `cli.go`; `modes.InteractiveConfig`; extensions/extproto/ext SDK; tests | `ZotHome`→`NcodeHome`; `ZotDocsDir`→`NcodeDocsDir`; `ZotVersion`→`NcodeVersion`; local `zotHome/zotVersion/zotdocs` identifiers→`ncodeHome/ncodeVersion/ncodedocs`. |
| Zotfile symbols | rename/remove | `ZotfileManifest`, `runZotfileCommand`, `runZotfileStartupPre`, `AgentName`, consent/pack/load helpers | Remove with capability; do not rename to Ncodefile. |
| Package/product comments | rename/remove | `cmd/zot/main.go`; `packages/provider/provider.go`; `packages/provider/auth/assets/assets.go`; `packages/agent/{config,extensions,ext,extproto,skills,swarm,sdk}`; `docs.go`; comments found by the full scan | Describe lowercase ncode. Go exported identifier capitalization follows Go, but prose/product literals remain lowercase `ncode`. |
| Historical parser comments in active source | rename/remove | `packages/core/session.go` and any other production file | Reword without a Zot literal (for example, “older inherited builds”) or move detailed evidence into `docs/inherited-capabilities.md`; active source is not a provenance exception. |

The module/import move is one atomic build-preserving mechanical unit because changing `go.mod` before or after only part of the import graph breaks `go list` and tests.

## Durable, project-local, portable, and temporary paths

### Root resolution

| Baseline | Classification | Exact ncode result |
|---|---|---|
| `$ZOT_HOME` | rename/remove | `$NCODE_HOME`; `$ZOT_HOME` is never read. |
| `$XDG_STATE_HOME/zot` | rename/remove | `$XDG_STATE_HOME/ncode`. |
| macOS `~/Library/Application Support/zot` | rename/remove | `~/Library/Application Support/ncode`. |
| Linux/other-home `~/.local/state/zot` | rename/remove | `~/.local/state/ncode`. |
| Windows `%LOCALAPPDATA%\zot` | rename/remove | `%LOCALAPPDATA%\ncode`. |
| MCP bridge's independent Windows `%APPDATA%\zot` default | rename/remove | `examples/extensions/mcp-bridge/config.go` must use the exact host contract and `%LOCALAPPDATA%\ncode`; do not preserve `%APPDATA%` divergence. |
| last-resort `.zot` from `ZotHome()` | rename/remove | `.ncode`. |
| project `.zot/extensions`, `.zot/skills`, `.zot/mcp.json` | rename/remove | `.ncode/extensions`, `.ncode/skills`, `.ncode/mcp.json`; no scan/probe of `.zot`. The MCP bridge global and project resolver mirrors the host order. |
| Zotfile `.zot` archives | rename/remove | Remove capability and archive handling; no `.ncode` archive and no “ncodefile”. |
| `.zotsession` | rename/remove | `.ncodesession`; importer rejects/does not infer `.zotsession`. |

### Basenames retained unchanged

These are **not applicable** because they carry no Zot identity and this change is not a schema redesign: `config.json`, `auth.json`, `models.json`, `models-cache.json`, `update-check.json`, `AGENTS.md`, `SYSTEM.md`, `sessions/`, timestamped session `.jsonl`, `docs/`, `skills/`, `themes/`, `extensions/`, `logs/`, `extension.json`, `theme.json`, `bot.json`, `bot.pid`, `bot.log`, `swarm/agents/<id>/meta.json`, `events.jsonl`, `session.json`, and `checksums.txt`.

The Zotfile-only `agents/<name>/data`, `agents/<name>/consents`, bundled `AGENT.md`, manifest fields, and agent-scoped session roots are **rename/remove by deletion**, not migrated into ncode.

### Runtime-composed names

Ownership is non-overlapping: the state/path work unit owns only home, project-local, and portable-session paths; the swarm/subprocess/provider/temp/internal-identity work unit owns every branded temp, log, request, and socket name in the table below.

| Baseline composer | Classification | Exact target | Evidence |
|---|---|---|---|
| `.zotsession` portable export | rename/remove | `.ncodesession` | `packages/core/session_portable.go`; modes/docs/tests. |
| `zot-swarm-<root-hash>` | rename/remove | `ncode-swarm-<root-hash>` | `packages/agent/swarm/socketpath.go`. |
| `zot-bash-<hash>.log` | rename/remove | `ncode-bash-<hash>.log` | `packages/agent/tools/bash.go`. |
| `zot-gemini-image-<uuid>` | rename/remove | `ncode-gemini-image-<uuid>` | `packages/provider/gemini.go`. |
| `zot-clipboard-images` | rename/remove | `ncode-clipboard-images` | `packages/tui/clipboard_darwin.go`. |
| `zot-timeline-<time>.json` | rename/remove | `ncode-timeline-<time>.json` | `packages/agent/modes/timeline_view.go`. |
| `zot-update-*` temp/stash names | rename/remove | `ncode-update-*` | `packages/agent/updatecmd.go`; `extupdate.go`. |
| example `/tmp/zot-guard-audit.log` | rename/remove | `/tmp/ncode-guard-audit.log` | `examples/extensions/guard`. |
| provider request IDs `zot-<n>` | rename/remove | `ncode-<n>` | `packages/provider/openai_codex.go`. |
| default Kimi device id `zot` | rename/remove | `ncode` | `packages/agent/build.go`. |

## Environment variables

The source/test inventory has 23 unique Zot-prefixed names. Product variables use one `NCODE_*` namespace; the anomalous RPC name explicitly drops `CORE`.

| Baseline | Classification | Exact target/decision |
|---|---|---|
| `ZOT_HOME` | rename/remove | `NCODE_HOME`. |
| `ZOT_FLAT_TOOLS` | rename/remove | `NCODE_FLAT_TOOLS`. |
| `ZOT_COMPACT_INPUT` | rename/remove | `NCODE_COMPACT_INPUT`. |
| `ZOT_INLINE_IMAGES` | rename/remove | `NCODE_INLINE_IMAGES`. |
| `ZOT_CELL_ASPECT` | rename/remove | `NCODE_CELL_ASPECT`. |
| `ZOT_TOOL_ARG_WIDTH` | rename/remove | `NCODE_TOOL_ARG_WIDTH`. |
| `ZOT_THEME` | rename/remove | `NCODE_THEME`. |
| `ZOT_NO_BROWSER` | rename/remove | `NCODE_NO_BROWSER`. |
| `ZOT_FORCE_BROWSER` | rename/remove | `NCODE_FORCE_BROWSER`. |
| `ZOT_DEBUG_ANTHROPIC` | rename/remove | `NCODE_DEBUG_ANTHROPIC`. |
| `ZOT_AGENT_SKILLS` | rename/remove | `NCODE_AGENT_SKILLS`; retained as the explicit extra-skill search path. |
| `ZOT_AGENT_CONSENT` | rename/remove | Delete with Zotfiles; no ncode replacement. |
| `ZOTCORE_RPC_TOKEN` | rename/remove | `NCODE_RPC_TOKEN`; no `CORE`, fallback, or dual read. |
| `ZOT_SWARM_AGENT_ID` | rename/remove | `NCODE_SWARM_AGENT_ID`. |
| `ZOT_SWARM_EVENT_LOG` | rename/remove | `NCODE_SWARM_EVENT_LOG`. |
| `ZOT_SWARM_CREDENTIAL_STDIN` | rename/remove | `NCODE_SWARM_CREDENTIAL_STDIN`. |
| `ZOT_VERSION`, `ZOT_PREFIX` | rename/remove | Installer overrides `NCODE_VERSION`, `NCODE_PREFIX`. |
| `ZOT_AGENT_API_KEY_COMMAND_HELPER` | rename/remove | Test-only `NCODE_AGENT_API_KEY_COMMAND_HELPER`. |
| `ZOT_HELP_HELPER` | rename/remove | Test-only `NCODE_HELP_HELPER`. |
| `ZOT_SWARM_CREDENTIAL_HELPER` | rename/remove | Test-only `NCODE_SWARM_CREDENTIAL_HELPER`. |
| `ZOT_API_KEY_COMMAND_HELPER`, `ZOT_API_KEY_COMMAND_VALUE` | rename/remove | Test-only `NCODE_API_KEY_COMMAND_HELPER`, `NCODE_API_KEY_COMMAND_VALUE`. |

Standard/provider variables are **not applicable** and remain unchanged: `XDG_STATE_HOME`, `LOCALAPPDATA`, `HOME`, `GITHUB_TOKEN`, proxy variables, `TERM_PROGRAM`, all provider API/OAuth/cloud variables, `XDG_RUNTIME_DIR`, and MCP child environment values.

## RPC, extension, swarm, subprocess, and provider-internal identity

| Surface | Classification | Baseline evidence | Target contract |
|---|---|---|---|
| RPC process and token | rename/remove for product edges; neutral protocol is not applicable | `packages/agent/rpc.go`; `docs/rpc.md`; all `examples/rpc` | Spawn `ncode rpc`; use only `NCODE_RPC_TOKEN`; ignore `ZOTCORE_RPC_TOKEN`. Preserve current RPC protocol version 1, prompt/response/event shapes, and optional hello. With no ncode token, prompt-first remains valid; with a ncode token, the existing first-frame token hello remains required. Add no product field or mandatory hello, and give any unknown Zot-branded extra field no authorization/configuration meaning. |
| Extension protocol identity | rename/remove for product-bearing ack; neutral frames are not applicable | `extproto.HelloAckFromHost.ZotVersion/json:"zot_version"`; manager; Go SDK; docs/examples | Increment shared protocol to 2; keep extension-originated hello neutral; host ack becomes `product:"ncode", protocol_version:2, ncode_version:<version>`. Rename SDK field to `NcodeVersion`; reject the old product-bearing ack in clients while preserving idle auto-ready semantics and changing only its product diagnostic. |
| Extension manifests/discovery | rename/remove for branding; not applicable for neutral schema | `.zot/extensions`; `extension.json`; `exec`, `name`, `version` | Discover only `.ncode/extensions` and `$NCODE_HOME/extensions`; keep neutral `extension.json` schema/basename. |
| Swarm child | rename/remove | `os.Executable`; `--swarm-agent`; `ZOT_SWARM_*`; socket/log names; comments/tests | Keep self-executable construction and existing flags/JSON event/control lines, but all env/temp/socket/log identity is `NCODE_*`/`ncode-*`; old env alone has no effect. |
| Provider-originated product headers | rename/remove | Codex `originator`/user-agent/request id; Kimi auth user-agent; Zot debug name | Use lowercase `ncode` where the value represents this product. Provider-mandated first-party compatibility headers and standard provider identities are not applicable and remain unchanged. |
| Extension log markers | rename/remove | `[zot]` in `extensions/manager.go` | `[ncode]`. |
| Generic frame/event names | not applicable | RPC `hello`/`prompt`/response/events; extension `hello`/`hello_ack`/`ready`/`tool_call`; swarm `events.jsonl` and inbox `user/cancel/shutdown` | Retain neutral names and payload semantics. RPC version/optional-hello behavior is unchanged; only the extension host's already product-bearing acknowledgement advances to signaled v2. |

## Build, release, update, and installer inventory

| Surface | Classification | Evidence | Exact target |
|---|---|---|---|
| Make targets/output | rename/remove | `Makefile` | Build/install/run/release `./cmd/ncode` and `bin/ncode*`; target names `build/install/run/release` remain neutral. |
| GoReleaser | rename/remove | `.goreleaser.yaml` | `project_name`, build/archive IDs, binary, comments/header → `ncode`; GitHub owner/name `nlf/ncode`; archive `ncode_<version>_<os>_<arch>`; `checksums.txt` retained. |
| Release workflow | not applicable for generic orchestration; rename/remove for identity-coupled comments/checks | `.github/workflows/release.yml` | Keep CI-gated tag mechanics; ensure package verification expects ncode assets and repository. No Zot asset is uploaded. |
| CI workflow | not applicable | `.github/workflows/ci.yml` | Keep platform matrix and no-live-provider tests; module rename naturally changes build graph. |
| Unix/PowerShell installers | rename/remove | `install.sh`, `install.ps1` | Owner `nlf`, repo/binary `ncode`, `NCODE_VERSION/PREFIX`, ncode user-agent/temp prefix/help, ncode archive URLs. Emit the clean-break notice; no Zot lookup or alias. |
| Self-update/changelog | rename/remove | `packages/agent/update.go`, `updatecmd.go`, `changelog.go`, tests | GitHub API `nlf/ncode`; assets `ncode_*`; extracted binary `ncode[.exe]`; ncode temp/log/help; checksum remains `checksums.txt`. |
| Tracked generated artifacts | rename/remove | `examples/extensions/todo/zot-todo-extension`; `examples/rpc/python/__pycache__/zot_client.cpython-314.pyc` | Delete both; do not track renamed binary/bytecode counterparts. Generated `bin/`, `dist/`, GoReleaser output, and Python caches remain ignored/not applicable. |

## User-facing code, docs, examples, tests, and fixtures

| Exhaustive group | Classification | Evidence/target rule |
|---|---|---|
| CLI help/version/errors/logs/system prompt/auth pages/TUI labels | rename/remove | `cmd/ncode`; `packages/agent/{args,cli,config,systemprompt,update*,extcmd,botcmd,modelsync}`; `packages/agent/modes`; `packages/provider/auth`; `packages/tui`; provider comments/errors. Replace active identity with lowercase ncode and update expectations. |
| Root active documentation and issue templates | rename/remove | `README.md`, `CONTRIBUTING.md`, `AGENTS.md`, `.github/ISSUE_TEMPLATE/{bug_report,feature_request}.yml`. Commands, state, imports, links, examples, and product prose become ncode. |
| Active subsystem docs | rename/remove | `docs/{providers,extensions,rpc,skills,themes}.md`. Rewrite contracts to exact ncode mappings; embedded copies in `$NCODE_HOME/docs` follow automatically. |
| Zotfile guide and references | rename/remove | Delete `docs/zotfiles.md`; remove README/help/router/embed references and all pack/inspect/verify/run instructions. Do not substitute a ncodefile capability. |
| Examples and manifests | rename/remove | All `examples/rpc`, `examples/sdk`, `examples/extensions`, `examples/themes`; nested modules/Makefiles; manifest descriptions. Use ncode imports/process/env/project paths/handshakes; neutral manifest fields remain. |
| Current production/test source matches | rename/remove | Every current `git grep -nI -i zot -- packages cmd docs.go` result. Existing behavior test helpers/data are renamed; Zotfile tests are deleted; production source has no provenance exception. |
| New rejection fixtures | not applicable | Keep legacy literals only in dedicated test/fixture files whose basename begins `legacy_zot_`, such as `legacy_zot_home_test.go` or `legacy_zot_hello.json`. Each file must state and assert rejection/non-use and never document support; a merely Zot-named test function inside a generic file is insufficient. |
| Exact frozen architecture/planning records | retain as provenance | Only the eleven exact tracked files in the reviewed provenance manifest may retain contextual Zot lines. No `openspec/**` or `docs/superpowers/**` wildcard is allowed. |
| MIT attribution | retain as provenance but not a search exception | `LICENSE` remains byte-for-byte unless legal review requires an additive notice; it currently has no Zot literal and is included in the zero-active scan. |

## Zotfile capability removal boundary

The entire capability is **rename/remove by deletion**, never rename:

- delete `packages/agent/zotfile.go`, its two dedicated test files, and `docs/zotfiles.md`;
- remove router calls for `pack`, `inspect`, `verify`, and `run`, Zotfile-only `-y/--yes`, and README/help/embed references;
- delete only Zotfile-specific `Args.AgentName`, `Args.AgentDataDir`, `Args.StartupPre`, entry-pre execution/reload hooks, startup-agent display, consent/agent-data wiring, and named-agent default session-root behavior without touching existing Zot data on disk;
- remove Zotfile-only `agents/<name>/{data,consents}` active contracts and remove `ZOT_AGENT_CONSENT` rather than creating `NCODE_AGENT_CONSENT`;
- do not create `ncodefile`, `.ncode` archives, `NcodefileManifest`, or migration tooling.

The following shared capabilities are protected and therefore **not applicable to capability deletion**, though their active comments are debranded: built-in read/write/edit/bash tools; `tools.Sandbox`; `Args.PermissionSet` in `packages/agent/args.go`; the exact `args.PermissionSet` → `sandbox.SetPermissions(args.PermissionSet)` wiring in `packages/agent/build.go`; `tools.PermissionSet`, `Sandbox.SetPermissions`, and focused direct-wiring tests; normal/global/project skills; normal sessions/import/export/fork/tree; extension subprocesses; confirmation; and provider composition.

## Final reviewed manifest and audit

At final audit, every remaining case-insensitive Zot line must be either (a) contextual evidence in one of the eleven exact provenance files or (b) rejection/non-use evidence in a dedicated `legacy_zot_*` test/fixture. The audit is intentionally two-part: the active-surface commands must emit no lines, then reviewers inspect **every** allowed line emitted by the provenance and rejection-evidence commands.

```sh
# Exact exclusions; do not replace any entry with openspec/**, docs/superpowers/**,
# a test directory, or a fixture directory.
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

# Zero-active gates: each command must print no matches.
git grep -nI -i -e zot -- "${audit_pathspecs[@]}"
git grep -nI -F 'github.com/patriceckhart/zot' -- "${audit_pathspecs[@]}"
git grep -nI -E 'ZOTCORE_|ZOT_[A-Z0-9_]+' -- "${audit_pathspecs[@]}"
git grep -nI -E '\.zot(session)?([^[:alnum:]_]|$)' -- "${audit_pathspecs[@]}"
git grep -nI -E 'zot[-_][[:alnum:]_-]+' -- "${audit_pathspecs[@]}"
git ls-files | grep -i zot | grep -vE '(^|/)legacy_zot_[^/]*$'

# Workable file-set check: every affected file must be exact provenance or a
# dedicated rejection/non-use file. `comm` must emit no paths.
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

# Mandatory line review: save and review every remaining provenance and
# rejection/non-use line; no allowed line is accepted merely by directory.
git grep -nI -i -e zot -- \
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
  | tee /tmp/ncode-zot-provenance-lines.txt
git grep -nI -i -e zot -- ':(glob)**/legacy_zot_*' \
  | tee /tmp/ncode-zot-rejection-lines.txt
```

Review sign-off records that every provenance line has its listed factual reason and every `legacy_zot_*` line participates in an assertion of rejection, ignore, or non-use. Neutral protocol names and frame shapes are not exceptions because they contain no Zot identity and remain unchanged.
