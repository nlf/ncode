# Inherited Zot capabilities in ncode

Today, ncode is the inherited Zot Go application: one `zot` executable, a distributed host under `packages/agent`, a provider-neutral loop under `packages/core`, a broad provider catalog, terminal and headless modes, and active swarm and extension subsystems. This inventory is anchored to commit `82191b33a9d54993ce9c85988dc250421623b75b`; the local annotated tag `zot-baseline` dereferences to that commit. The implementation paths described below were spot-checked against that baseline.

> [!IMPORTANT]
> **Scope rule:** the baseline inventory is factual evidence about the inherited implementation. The decision register is canonical for capability disposition. It records what ncode will retain, reshape, replace, or remove, but does not choose implementation sequencing.

The [ncode architecture and roadmap](ncode-architecture.md) supplies the broader architecture and roadmap context. For capability disposition, this register is authoritative; implementation design and sequencing remain open where identified below.

## Baseline reality

| Fact | Observed evidence | Qualification |
|---|---|---|
| Product implementation | `cmd/zot/main.go:main` invokes `agent.Run`; module imports remain `github.com/patriceckhart/zot/...`. | There is not yet a new ncode package or composition root. |
| Branding and state | User-facing strings, binary naming, package comments, module paths, and environment/state identifiers still use Zot, including `ZOT_HOME`, `ZOT_SWARM_AGENT_ID`, `ZOT_SWARM_EVENT_LOG`, `ZOT_AGENT_SKILLS`, and `ZOTCORE_RPC_TOKEN`. | ncode is the project name; code identifiers are recorded unchanged. |
| Composition | `packages/agent/cli.go` and `packages/agent/build.go` jointly assemble and run the application. | The inherited composition is distributed rather than a single constructor. |
| Core boundary | `packages/core/agent.go:Agent` depends on `packages/provider/provider.go:Client`. | This is a clean provider-neutral model/tool-loop seam. |
| Default tools | `packages/agent/build.go:buildToolRegistry` registers `read`, `write`, `edit`, and `bash` unless flags narrow or disable them. | Skills, extensions, and auto-swarm can add more tools after initial registry construction. |
| Providers | `packages/agent/build.go:Resolved.NewClient` constructs many direct, cloud, gateway, local, and custom adapters. | The comment in `packages/provider/provider.go` saying the package supports “exactly two providers” is stale. |
| Swarm and extensions | Both are wired into normal execution and have broad focused test suites. | They are not dead code or isolated experiments. |

## Decision register

The tables below record the completed feature walkthrough. They decide capability disposition and constraints, not delivery order or detailed implementation.

### Interfaces and UI

| Capability | Disposition | Decision / constraint |
|---|---|---|
| Interactive TUI | **Replace implementation** | Fully replace the inherited TUI with Bubble Tea. |
| Print, stream, JSON event, and RPC modes | **Retain** | Keep all four headless interfaces. |
| Extension-management and session-pruning CLIs | **Retain** | Keep both management command surfaces. |
| Slash commands | **Replace implementation** | Retain the command capability while replacing its UI and dispatch integration. |
| Multiline editor, history, and file/command completion | **Replace implementation** | Rebuild with Bubble Tea/Bubbles while retaining the user capabilities. |
| Themes, Markdown/syntax rendering, images, and clipboard | **Replace implementation** | Replace the inherited UI integrations and renderers while retaining their user-facing and data capabilities. |
| One-shot session persistence | **Retain and harden** | Keep persistence configurable and provide an explicit ephemeral/no-session mode. |
| RPC session persistence | **Reshape/extend** | Add optional persistence while preserving ephemeral RPC operation. |
| RPC token authentication | **Retain and harden** | Keep token authentication optional and strengthen its contract. |
| Extensions across headless, RPC, and swarm modes | **Reshape/extend** | Retain extension capability in every mode and document a mode capability matrix with explicit supported, translated, rejected, and no-op outcomes. |

### Core, sessions, and prompts

| Capability | Disposition | Decision / constraint |
|---|---|---|
| Provider-neutral `core.Agent` / `provider.Client` boundary | **Retain and harden** | Preserve the neutral seam and strengthen its contracts. |
| Mutable host callbacks | **Reshape/extend** | Replace ad hoc mutable callbacks with explicit event and middleware contracts. |
| Normalized agent event schema | **Retain and harden** | Preserve and version the schema. |
| Retry and recovery | **Retain and harden** | Research external implementations; require typed, replay-safe recovery with no duplicate tool execution. |
| Queueing and cancellation | **Retain and harden** | Preserve both behaviors and strengthen their contracts. |
| Transcript integrity repair | **Retain and harden** | Preserve repair and make integrity guarantees explicit. |
| Session persistence, resume, and branching | **Retain** | Keep durable resume and session branching. |
| Append-oriented JSONL sessions | **Retain and harden** | Keep JSONL initially and version the format. |
| Compaction | **Retain and harden** | Keep compaction configurable; automatic compaction must be visible to the user. |
| System prompt override/append and `AGENTS.md` discovery | **Retain and harden** | Preserve both prompt controls and context discovery. |
| Default system prompt | **Retain** | Use the inherited prompt as the initial baseline and evolve it only from evidence. |

### Providers, authentication, configuration, and accounting

| Capability | Disposition | Decision / constraint |
|---|---|---|
| Inherited provider implementations | **Retain and harden** | Do not prune providers. Keep standard OpenAI and Anthropic API paths, Codex and Claude subscription paths, every cloud/gateway/vendor adapter, custom providers, and custom/local OpenAI-compatible endpoints. |
| Inherited authentication methods | **Retain and harden** | Preserve every inherited authentication method. |
| OAuth refresh and rotated-token persistence | **Retain and harden** | Refresh and persist rotated tokens automatically without prompting on every refresh. |
| Model catalog and discovery | **Reshape/extend** | Improve static, cached/live, custom, and local model handling; research external implementations and define precedence and refresh behavior clearly. |
| Runtime model controls | **Retain and harden** | Keep model, reasoning, temperature, output, and max-step controls with capability validation. |
| Usage, cost, quota, and limits | **Reshape/extend** | Preserve usage/cost tracking and add quota, rate-limit, and subscription-limit visibility when providers expose it. |
| Proxy and insecure TLS configuration | **Retain and harden** | Preserve proxy support and keep insecure TLS explicitly endpoint-scoped. |
| Configuration validation | **Retain and harden** | Never repair config silently: interactive mode shows the exact proposal and asks; noninteractive modes exit with the config unchanged. |

### Tools and safety

| Capability | Disposition | Decision / constraint |
|---|---|---|
| Built-in `read`, `write`, `edit`, and `bash` tools | **Retain and harden** | Preserve the built-ins; additional built-ins may be considered later. |
| Gitignore-aware behavior | **Retain** | Preserve ignore-aware discovery and traversal. |
| Sandbox and containment | **Retain capability; implementation undecided** | Research harden-versus-replace. Do not describe the current containment as a full OS sandbox. |
| Tool confirmation and permissions | **Reshape/extend** | Redesign interactive confirmation for Bubble Tea. Headless modes require an explicit deny, allow, allowlist, or external-approval policy rather than implicit approval. |

### Swarm

| Capability | Disposition | Decision / constraint |
|---|---|---|
| Swarm | **Retain** | Keep swarm as a core ncode capability. |
| Model-initiated `swarm_spawn` | **Reshape/extend** | Keep explicit enablement and add future concurrency, permission, and cost controls. |
| Persisted, detached, resumable swarm state | **Retain and harden** | Preserve and strengthen durable detach/resume behavior. |
| Shared repository/CWD | **Retain** | Keep shared repository/CWD as the default. Worktree isolation is optional research, not a default, and may become a general agent tool rather than swarm-owned automation. |

### Extensions and skills

| Capability | Disposition | Decision / constraint |
|---|---|---|
| Extensions | **Reshape/extend** | Preserve and extend extensions, including additional Bubble Tea-relevant UI hooks. |
| Subprocess extensions | **Retain and harden** | Keep them with explicit full-user-trust disclosure; investigate optional isolation later. |
| Interceptor failure policy | **Reshape/extend** | Make policy explicit and configurable: ordinary interceptors may fail open, while declared security guards fail closed. |
| Zot extension compatibility | **Reshape/extend** | Do not guarantee compatibility. ncode owns and evolves its protocol, SDK, paths, manifests, and hooks. |
| Skills | **Retain and harden** | Revisit every metadata field during the skills phase; currently parsed but unenforced metadata is not automatically canonical. |

### Maintenance and identity

| Capability | Disposition | Decision / constraint |
|---|---|---|
| Self-update | **Replace implementation** | Retain the capability but rebuild it for ncode releases, checksums, changelogs, and configurable background checks. |
| Telegram bot | **Retain; defer investment** | Keep the inherited implementation as-is and defer further investment. |
| Zotfiles | **Remove** | Remove Zotfiles only; protect the general tools, skills, permissions, and sessions they reuse. |
| Zot naming compatibility | **Remove** | Drop compatibility naming across the binary, module/imports, state, environment, release assets, and internal protocol names. |
| Zot state migration compatibility | **Remove** | Build no Zot state import, migration, or fallback path. |

## User-facing entry points and modes

`packages/agent/cli.go:Run` is the top-level router after `cmd/zot/main.go:main`. Subcommands are checked before generic flags, and `packages/agent/args.go:ParseArgs` selects the normal run mode.

| Entry point | Router / executor | Observable role |
|---|---|---|
| `zot` or `zot "prompt"` | `runInteractive` → `modes.NewInteractive` → `(*modes.Interactive).Run` | Full-screen TUI; can start without credentials, then login, choose models, resume sessions, use dialogs, extensions, and swarm. |
| `zot -p "prompt"`; piped stdin with no explicit mode | `runPrintMode` → `modes.RunPrint` | Runs to completion and writes only the final assistant text; optional `--stats`. |
| `zot --stream "prompt"` | `runStreamMode` → `modes.RunStream` | Streams assistant text to stdout and tool diagnostics to stderr. |
| `zot --json "prompt"` | `runJSONMode` → `modes.RunJSON` | Emits newline-delimited JSON agent events. |
| `zot rpc` or `zot --rpc` | `runRPCMode` in `packages/agent/rpc.go` | Long-running JSONL command/event interface over stdin/stdout; supports prompt, abort, compact, state/messages/models, model/reasoning changes, and clear. Optional first-frame auth uses `ZOTCORE_RPC_TOKEN`. |
| Internal `--swarm-agent <socket>` | `runSwarmAgentMode` in `packages/agent/swarm_agent.go` | Long-lived swarm child: persistent session, JSONL stdout events, and a Unix-socket inbox for `user`, `cancel`, and `shutdown`. |
| `zot ext ...` | `runExtCommand` in `packages/agent/extcmd.go` | `list`, `doctor`, `logs`, `enable`, `disable`, `remove`, `install`, and help. |
| `zot sessions prune ...` | `runSessionsCommand` in `packages/agent/sessionscmd.go` | Selects missing-directory or old session groups; supports dry-run, age, CWD, and explicit bulk confirmation. |
| `zot update [--check]` | `runUpdateCommand` in `packages/agent/updatecmd.go` | Checks releases or downloads, verifies, extracts, and replaces the current executable. Interactive startup also performs an async update check. |
| `zot telegram-bot ...` / `zot tg ...` | `runBotCommand` → `telegramSpec` / generic `bot.Runner` | Setup, status, foreground/background run, stop, logs, and reset for a Telegram bridge. |
| `zot pack`, `inspect`, `verify`, `run` | `runZotfileCommand` in `packages/agent/zotfile.go` | Packages, inspects, verifies, and launches local/archive/GitHub Zotfile agents with requirements, consent, scoped data, prompt, skills, and permissions. |
| `-h` / `--help`; `-v` / `--version` | `printHelp`; version branch in `runWithArgsRaw` | Prints the inherited Zot command catalog and build version. |
| `--list-models` | `prepareRuntimeCatalog` → `printModels` | Loads cached/user model data, starts async discovery, prints the active catalog, and exits. |

## Composition map

### Compact flow

```text
cmd/zot/main.go:main
  -> packages/agent/cli.go:Run
       -> subcommand routers (bot | ext | update | sessions | Zotfile)
       -> runWithArgsRaw
            -> ParseArgs
            -> prepareRuntimeCatalog
            -> runWithArgs
                 |
                 +-- interactive: runInteractive
                 |     -> Resolve(false) -----------------------------+
                 |     -> auth.Manager                                |
                 |     -> extensions.Manager -> MergeExtensionTools   |
                 |     -> swarm.Swarm -> optional swarm_spawn tool    |
                 |     -> Resolved.NewAgent -> core.NewAgent          |
                 |     -> session/persistence callbacks               |
                 |     -> tui.NewProcTerm + modes.NewInteractive      |
                 |     -> (*modes.Interactive).Run                    |
                 |
                 +-- headless: runPrintMode | runStreamMode |
                       runJSONMode | runRPCMode | runSwarmAgentMode
                       -> Resolve(true)
                       -> extensions.Manager -> MergeExtensionTools
                       -> Resolved.NewAgent -> core.NewAgent
                       -> mode sink / JSONL protocol / swarm inbox

Resolve (packages/agent/build.go)
  -> LoadConfig + credential resolution + provider/model catalog
  -> tools.Sandbox + read/write/edit/bash registry
  -> skills.Discover + AGENTS.md context + BuildSystemPrompt
  -> Resolved.NewClient
       -> packages/provider/* adapter
       -> packages/provider/provider.go:Client.Stream

core.Agent.Prompt / Continue / Compact
  -> provider.Request -> Client.Stream -> normalized provider.Event values
  -> transcript + usage + retries + tool execution + AgentEvent values
  -> TUI, print/stream/JSON sink, RPC client, extension fanout, or swarm log
```

### Normal interactive composition

1. `cmd/zot/main.go:main` resolves build metadata and calls `packages/agent/cli.go:Run`.
2. `Run` applies proxy configuration and tries `runBotCommand`, `runExtCommand`, `runUpdateCommand`, `runSessionsCommand`, and `runZotfileCommand` before falling through to `runWithArgsRaw`.
3. `runWithArgsRaw` handles piped-stdin mode selection, `ParseArgs`, help/version, and `prepareRuntimeCatalog`; `runWithArgs` dispatches to `runInteractive`.
4. `runInteractive` calls `packages/agent/build.go:Resolve(args, false)`. `Resolve` merges CLI/config/defaults, resolves credentials without requiring one, resolves the model, creates `tools.Sandbox`, builds the baseline registry, discovers skills and `AGENTS.md`, and builds the system prompt.
5. `runInteractive` separately creates `auth.NewManager`, `extensions.New`, and `swarm.New`. Extension tools are merged through `Resolved.MergeExtensionTools`; auto-swarm may inject `tools.SwarmSpawnTool`.
6. `Resolved.NewClient` selects the provider adapter; `Resolved.NewAgent` calls `core.NewAgent` and applies model limits, reasoning, temperature, and max steps.
7. Host callbacks couple the neutral agent to inherited policy: extension interceptors, confirmation, event fanout, per-message usage/session persistence, transcript compaction checkpoints, login/model rebuilds, CWD changes, and session selection.
8. `tui.NewProcTerm`, theme detection, update/changelog channels, and a large `modes.InteractiveConfig` are passed to `modes.NewInteractive`; `(*Interactive).Run` owns the interactive event loop and dialogs.
9. User prompts reach `core.Agent.Prompt`. `core.Agent.runLoop` requests a normalized stream through `provider.Client`, emits `core.AgentEvent` values, executes tools through `core.Registry`, appends tool results, retries selected transient failures, and stops or continues according to the neutral stop reason.

### Headless composition

The print, stream, and JSON paths call `Resolve(args, true)`, so missing credentials fail before execution. Each path starts extensions via `setupNonInteractiveExtensions`, merges registered tools, builds a `core.Agent`, wires extension interceptors/events, opens or resumes a core JSONL session, runs Zotfile `entry.pre` if present, and delegates output policy to `modes.RunPrint`, `modes.RunStream`, or `modes.RunJSON`.

RPC follows the same resolution, extension, provider, and core loop but owns a long-running JSONL command server in `packages/agent/rpc.go:rpcServer`; unlike print/stream/JSON, the observed RPC composition does not call `openOrCreateSession`. Swarm child mode follows the headless path, including extension loading, then adds a durable child session, Unix-socket listener, and JSONL lifecycle/event emitter. In non-interactive modes, `--no-yolo` is reported as ineffective and does not provide an interactive confirmation gate.

## Capability inventory

### Runtime, configuration, providers, and tools

| Capability | Current implementation and composition | Evidence / boundary |
|---|---|---|
| Neutral turn/tool loop | Stateful transcript, queued prompts, normalized streaming, retries, usage, stop reasons, tool execution, cancellation, and host callbacks. | `packages/core/agent.go:Agent`, `Prompt`, `Continue`, `runLoop`, `oneTurn`, `executeTools`. |
| Neutral provider contract | Messages/content, tool specs, request options, usage, stop reasons, stream events, and `Client.Stream`. | `packages/provider/provider.go`; wire types stay in adapters. |
| Configuration and state roots | JSON config, auth, sessions, logs, model caches, settings, proxy, and `ZOT_HOME` fallback. | `packages/agent/config.go:Config`, `ZotHome`, `LoadConfig`, `SaveConfig`. |
| Model catalog | Static catalog plus built-in additions, cached discovery, user models/custom providers, managed llama.cpp models, aliases, defaults, capability/cost metadata, validation, and repair. | `packages/provider/models.go`, `catalog_builtin.go`, `extra_models.go`, `extra_providers.go`; `packages/agent/modelsync.go`. |
| Async catalog refresh | Credential-gated discovery updates a mutex-protected live overlay and cache in background goroutines. | `prepareRuntimeCatalog`, `RefreshModelsAsync`, `refreshModels`, `provider.SetLiveModels`, `SetManagedModels`. |
| Provider construction | Anthropic; OpenAI Chat Completions, Responses, and Codex; Kimi/DeepSeek/Moonshot; Gemini; Ollama/llama.cpp; Bedrock, Vertex, Azure, Copilot; xAI, Groq, Cerebras, Together, Hugging Face, OpenRouter, Gondola, Mistral, ZAI, Xiaomi, MiniMax, Fireworks, Vercel AI Gateway, OpenCode, Cloudflare; and custom OpenAI/Responses/Anthropic-compatible providers. | `packages/agent/build.go:Resolved.NewClient`; adapter files under `packages/provider`. |
| Authentication | Explicit API keys, env/config lookup, stored endpoint credentials, command-backed keys, OAuth/PKCE and manual/device flows, refresh wrappers, account metadata, and cloud-specific ambient credentials. | `packages/agent/config.go:ResolveCredentialFullContext`; `packages/provider/auth`; `Resolved.wrapWithRefresh`. |
| Default resolver tools | `read`, `write`, `edit`, `bash`; `--no-tools` disables all and `--tools` selects a subset. | `packages/agent/build.go:buildToolRegistry`. |
| Filesystem/bash safety | CWD-root sandbox/jail, per-tool confirmation in interactive `--no-yolo`, and Zotfile filesystem/bash permissions. | `packages/agent/tools/sandbox.go`, `permissions.go`; `core.ConfirmGate`; tool implementations. |
| Extension/skill/swarm tools | Extension tools are merged after startup registration; `skill` is added when skills exist; `swarm_spawn` is conditional on auto-swarm. | `Resolved.MergeExtensionTools`, `skills.NewTool`, `runInteractive:injectSwarmSpawn`. |

### Sessions, interfaces, and host capabilities

| Capability | Current implementation and composition | Evidence / boundary |
|---|---|---|
| Session format and durability | Core-owned JSONL metadata, messages, usage, compaction checkpoints, repair, descriptions, rename/delete/list, and portable session handling. | `packages/core/session.go`, `session_portable.go`, `session_prune.go`. |
| Open/resume/continue | Host selects roots and paths, trims the live resume window, seeds usage, updates models, swaps sessions, and persists each appended message. | `packages/agent/cli.go:openOrCreateSession`, `loadSession`, `persistMessage`, `WriteNewTranscript`; TUI session dialogs. |
| Compaction | `core.Agent.Compact` makes a separate provider summarization request, replaces the transcript with a synthetic summary plus repaired tail, and notifies host persistence. Interactive mode adds manual/automatic policy and queuing. | `packages/core/compact.go`; `packages/agent/modes/interactive.go:runCompact` and auto-compaction tests. |
| TUI primitives | Terminal input, editor, layout, rendering, markdown/highlighting, image/clipboard support, view/status bar, themes, wrapping, and resize handling. | `packages/tui`. |
| Interactive policy | Login/model/session/settings dialogs, slash commands, confirmations, compaction, swarm dashboard, extension panels/notes, update/changelog UX, and orchestration closures. | `packages/agent/modes`, especially `interactive.go`. |
| Print/stream/JSON | Final-text, live-text, and full JSON event projections over the same `core.Agent`. | `packages/agent/modes/print.go`, `stream.go`, `json.go`. |
| RPC | JSONL commands/events, optional token gate, concurrent abortable prompts/compaction, state/model/reasoning operations, and extension notifications. | `packages/agent/rpc.go`. |
| Swarm | Multi-agent supervisor and child daemon using the same executable and same repository/CWD. | `packages/agent/swarm`, `packages/agent/swarm_agent.go`, TUI swarm dialogs/slash commands, `tools/swarm_spawn.go`. |
| Extensions | Installed/project/ad-hoc subprocesses with command/tool registration, events/interceptors, TUI hooks, themes, reload, diagnostics, and CLI management. | `packages/agent/extensions`, `extproto`, `ext`, `extcmd.go`, TUI extension integration. |
| Skills | Built-in and user/project `SKILL.md` discovery, prompt manifest, on-demand `skill` tool, and TUI picker/explicit invocation. | `packages/agent/skills`; `Resolve`; `packages/agent/modes/skill_command.go`. |
| Zotfiles | Local/archive/GitHub loading, canonical pack/verify, requirements, consent receipts, agent prompt/default prompt/`entry.pre`, bundled skills, agent-specific data/session roots, and filesystem/bash permissions. | `packages/agent/zotfile.go`; `packages/agent/tools/permissions.go`. |
| Bots / remote bridge | Generic daemon command/runner boundary with a registered Telegram adapter, pairing/config, polling, remote commands, session reuse, PID/log management, and foreground/background lifecycle. | `packages/agent/botcmd.go`, `botspec.go`, `packages/agent/modes/bot`, `modes/telegram`. |
| Updates/changelog | Cached release checks, async interactive banner/changelog, and explicit self-update command with archive/checksum handling. | `packages/agent/update.go`, `updatecmd.go`, `changelog.go`, TUI update/changelog modes. |
| Themes | Dark/light defaults, terminal detection, custom theme loading, extension-provided themes, and configurable render choices. | `packages/tui/theme.go`, `theme_loader.go`; extension `ThemeOptions`; `runInteractive`. |
| Ignore support | Root and nested `.gitignore` matching used by recursive host behavior, with configurable respect for ignore rules. | `packages/ignore`; interactive file suggestion/settings wiring. |

### Important qualifications

- The provider set is materially broader than the initial ncode provider priorities. The register nevertheless retains and hardens every inherited implementation; this inventory records current evidence rather than proof of equal maturity.
- `provider.Client` remains clean even though `Resolved.NewClient` and credential construction are broad and host-specific.
- Skills parse `allowed-tools` and permission metadata for compatibility, but `packages/agent/skills/skills.go` explicitly says those fields are not enforced in this version.
- Zotfile filesystem and bash permissions are implemented. Zotfile manifests that request `permissions.net` or `permissions.env` are currently rejected by `readZotManifest`; executable bundled extensions are also rejected because they cannot yet be confined.
- The Zotfile tests prove rejection of currently unenforced Net/Env declarations, but there is no clear end-to-end proof of a future full Net/Env enforcement path.
- Interactive `--no-yolo` confirmation is host policy around tool calls. Headless modes warn that the flag is ineffective and execute without an interactive prompt.
- Model discovery mutates the active catalog asynchronously after startup. The provider package protects snapshots with `activeMu`, but UI/config behavior can still observe catalog changes over time.

## Swarm: lifecycle, boundaries, protocol, and risks

### Lifecycle and wiring

1. `runInteractive` constructs `swarm.New` with durable root `filepath.Join(ZotHome(), "swarm")` and `RepoRoot: r.CWD`, then calls `Swarm.Reload`.
2. A user action or `tools.SwarmSpawnTool` calls `Swarm.SpawnReq`. It creates `<root>/agents/<id>/meta.json`, `events.jsonl`, and `session.json`, plus a transient Unix-socket inbox path.
3. The production `execRunner` locates the currently running executable with `os.Executable` and launches it as `--swarm-agent <socket> --session <path> --cwd <RepoRoot>`, optionally with provider/model and the initial task.
4. Resolved credentials are JSON-encoded to child stdin when needed, gated by `ZOT_SWARM_CREDENTIAL_STDIN=1`; they are not placed in argv or ordinary environment values. The runner sets `ZOT_SWARM_AGENT_ID` and `ZOT_SWARM_EVENT_LOG`.
5. `runSwarmAgentMode` uses the normal `Resolve` → extension loading → `Resolved.NewAgent` path. Therefore a swarm child loads extensions and runs the same core/provider/tool composition as other headless modes.
6. The child opens `swarm.Listen`, emits `agent_ready`, runs the initial prompt, and accepts line-oriented `user <text>`, `cancel`, and `shutdown` messages. It rejects a concurrent turn as busy rather than queueing it.
7. Child stdout is JSONL. The parent runner decodes events, appends them to `events.jsonl`, updates in-memory activity/transcript, and recognizes prompt-level `turn_end`. If the parent pipe breaks, the child emitter can switch to direct event-log mirroring.
8. `Swarm.Reload` reconstructs durable agents as detached; `Swarm.Resume` starts a fresh child against the same session and event files without replaying the original task. Host session IDs scope dashboard visibility.

### Dependency and persistence boundaries

| Boundary | Current contract |
|---|---|
| Host ↔ supervisor | `packages/agent/cli.go` supplies repo root, credential resolver, active host session, TUI callbacks, and optional `swarm_spawn` registration. |
| Supervisor ↔ child process | CLI flags, credential JSON on stdin, `ZOT_SWARM_*` environment metadata, Unix-socket control lines, and JSONL stdout/stderr events. |
| Child composition | Reuses the same executable, `Resolve`, extension manager, provider adapter construction, core loop, default tools, and session functions. |
| Durable state | Immutable-ish `meta.json`, append-only `events.jsonl`, and core session JSONL at `session.json`; runtime socket is transient. |
| Repository state | Every child uses the parent `RepoRoot` as `cwd`; no worktree, branch, or filesystem isolation is created. |

### Focused test evidence

The focused suites cover supervisor state transitions and shared repo root (`swarm_test.go`), inbox round trips/readiness/concurrency (`inbox_test.go`), event-log append/replay/dedup/following (`event_test.go`), metadata/reload/resume/session scoping (`persist_test.go`), argv and credential-stdin contracts (`runner_test.go`), a stub-child subprocess path (`runner_e2e_test.go`), and child-mode behavior (`packages/agent/swarm_agent_test.go`). TUI slash/dialog and auto-swarm behavior also have focused tests under `packages/agent/modes`.

### Known risks and uncertainty

- **Concurrent edits are real:** parent and children share the same files and CWD without worktree isolation. The code and tests establish this behavior; they do not prove conflict-free multi-agent editing.
- **No live concurrent-edit proof:** the subprocess E2E uses a stub child, not multiple live model-backed agents editing one repository concurrently.
- **Parent/child contract breadth:** flags, stdin credential shape, environment metadata, socket commands, event schema, session format, and executable self-location must remain compatible together.
- **Unix-socket transport:** the observed inbox implementation is Unix-domain-socket based; comments identify Windows named-pipe support as follow-up work.
- **Durability has two writers by design only after orphaning:** normal event-log ownership belongs to the parent runner; the child mirror activates after stdout failure. Tests cover duplicate-log recovery, but crash timing remains a sensitive boundary.

## Extensions: lifecycle, boundaries, protocol, and risks

### Lifecycle and wiring

1. `extensions.New` receives `ZotHome`, CWD, version, provider/model, and a `HostHooks` implementation.
2. Explicit `--ext` directories load first. Unless `--no-ext` is set, `Discover` scans project `.zot/extensions` before global `$ZOT_HOME/extensions`; first name wins.
3. `loadOne` reads `extension.json`. Executable extensions spawn as independent subprocesses in the extension directory; theme-only extensions can load without an executable. Stderr is appended to `$ZOT_HOME/logs/ext-<name>.log`.
4. The extension must send a JSONL `hello`; the host replies with `hello_ack` using `extproto.ProtocolVersion == 1`. The extension can then register slash commands and tools, subscribe to events/interceptors, and send `ready` (with an idle compatibility fallback for older SDKs).
5. `WaitForReady` bounds startup registration. `Resolved.MergeExtensionTools` wraps registered tools as `core.Tool` values and rebuilds the system-prompt tool list.
6. Host agent callbacks route `tool_call`, `turn_start`, and `assistant_message` through serial interceptors and fan normal events to subscribers. Interceptors can block; tool args and visible assistant text can be rewritten.
7. Interactive `HostHooks` support notify, submit, slash submit, editor insert, display/clear notes, and panel open/update/close. Headless hooks degrade these behaviors to stderr, no-ops, or RPC events as appropriate.
8. `Manager.Reload` stops processes, clears registries, reloads explicit/discovered extensions, waits for readiness, and invokes a host callback to rebuild the live tool registry and prompt.
9. Shutdown is graceful with a bounded wait and process fallback. A crashed/malformed extension is removed from command/tool indexes without terminating the host or other extensions.

### Dependency, protocol, and persistence boundaries

| Boundary | Current contract |
|---|---|
| Shared protocol | `packages/agent/extproto` defines one-JSON-object-per-line frames for hello, registration, readiness, subscriptions, commands, tools, events/intercepts, panels, notifications, and shutdown. |
| Host process manager | `packages/agent/extensions.Manager` owns discovery, processes, indexes, correlation IDs, timeouts, diagnostics, reload, and failure isolation. |
| Extension SDK | `packages/agent/ext` shares the protocol types without owning host composition. |
| Core integration | Extension tools implement `core.Tool`; `core.Agent` callbacks expose interception and event observation without importing extension packages. |
| Cycle avoidance | `packages/agent/build.go:ExtensionToolSource` and `ExtensionToolInfo` let `Resolved.MergeExtensionTools` avoid importing `packages/agent/extensions`, which itself depends on core. `cli.go:extToolAdapter` bridges the types. |
| TUI integration | `HostHooks` and `modes.InteractiveConfig.Extensions` expose host actions, commands, notes, panels, keys, and themes. |
| Persistence | Manifests/executables live in project/global extension directories; stderr logs persist under `ZOT_HOME`. Runtime registrations and process state are rebuilt on each load/reload. |

### Focused test evidence

`packages/agent/extensions/manager_test.go` covers spawn/invoke, invalid hello cleanup, pre-hello exits, theme-only loading, spontaneous submit/panels, and diagnostics/conflicts. `tool_test.go` covers tool registration/invocation and result conversion. `intercept_test.go` covers all three intercept surfaces and reload. `packages/agent/ext/ext_test.go`, `packages/agent/extcmd_test.go`, `packages/agent/extupdate_test.go`, and TUI extension/reload tests cover SDK, CLI, update, note, panel, and reload behavior.

### Known risks and uncertainty

- **Executable trust boundary:** extensions are arbitrary subprocesses, not `core.Tool` implementations confined by the Zotfile permission system. Their capability comes from process execution plus the protocol.
- **Host callback coupling:** `HostHooks` reaches deeply into interactive behavior. Non-interactive adapters deliberately discard unsupported actions, so extension behavior is mode-dependent.
- **Timing and fail-open behavior:** readiness has grace/legacy-idle paths; interceptor timeout or write failure allows execution to continue. This isolates failures but is security-relevant for guard-style extensions.
- **Registration conflicts:** command/tool ownership is first-come-first-served, with diagnostics for shadowed registrations; concurrent discovery makes deterministic ownership dependent on the manager's conflict rules rather than an explicit dependency graph.
- **Protocol compatibility:** host, SDK, third-party subprocesses, extension tools, TUI hooks, and reload must evolve together around protocol version 1.

## Ownership seams and dependency hotspots

These are observed seams, not refactoring prescriptions.

| Seam / hotspot | Evidence | Why it matters for later decisions |
|---|---|---|
| Distributed composition | `packages/agent/cli.go:runInteractive` plus `packages/agent/build.go:Resolve`, `Resolved.NewClient`, and `Resolved.NewAgent`. | Provider/model/auth/tool/prompt assembly is split from extensions, swarm, sessions, confirmations, update/theme, and TUI assembly. |
| Clean core → provider boundary | `core.Agent.Client provider.Client`; `provider.Client.Stream(context.Context, Request)`. | Core consumes neutral requests/events while adapters own HTTP/SSE and provider quirks. |
| Host callback coupling | `core.Agent` callbacks (`Before*`, `OnEvent`, `OnMessageAppended`, `OnUsage`, `OnTranscriptCompacted`) and closure-heavy `modes.InteractiveConfig`. | Important host policies are composed through callbacks rather than one explicit host object. |
| Extension cycle avoidance | `ExtensionToolSource` / `ExtensionToolInfo` in `build.go`, bridged by `extToolAdapter` in `cli.go`. | Preserves package acyclicity while coupling mirrored types and adapter code. |
| Swarm parent/child contract | `execRunner`, `swarmAgentArgs`, `runSwarmAgentMode`, `Inbox`/`Listener`, event JSONL, and inherited credential stdin. | A change in executable naming, flags, env, session format, protocol, or CWD behavior can break swarm composition. |
| Host/core session split | Core owns JSONL format/hydration/repair; host owns roots, selection, resume windows, UI, swaps, and persistence callback wiring. | Session behavior spans `packages/core` and `packages/agent/cli.go`. |
| Auth/provider construction | Credential resolution in `packages/agent/config.go`, interactive auth in `packages/provider/auth`, refresh wrappers and adapter switches in `build.go`. | Login UX, stored state, provider selection, and wire client construction cross package boundaries. |
| TUI primitive vs interaction-policy split | Reusable terminal/render/theme primitives in `packages/tui`; dialogs, slash commands, compaction/swarm/extension policy in `packages/agent/modes`. | A TUI change may be a primitive concern, a product-policy concern, or both. |
| Async catalog mutation | `RefreshModelsAsync`/`refreshModels` call `provider.SetLiveModels`; llama.cpp management calls `SetManagedModels`; readers call `Active`. | Model choices can change after initial config validation or UI construction even though data access is mutex-protected. |

## Test coverage and observed evidence gaps

The inherited baseline validation reported by the prior review is `make test` → `go test -race ./...`, passing. This inventory did not replace that run; it maps the tests present at the anchored implementation.

| Area | Existing evidence |
|---|---|
| Core | Scripted provider conversations, context/retry behavior, confirmation, interceptors, queueing, deferred tools, session portability/repair/pruning, and persistence behavior. |
| Host/config/modes | Argument/config/model-sync/auth behavior, prompt/context assembly, print/stream/RPC projections, interactive dialogs and terminal behavior, auto-compaction, tool gates, sessions, updates, Zotfiles, and bots. |
| Providers | Per-adapter request/stream/auth/discovery tests for many adapters, plus shared utility tests for retry, reasoning, HTTP, routing, SSE, models, and user configuration. |
| TUI/ignore | Editor/input/layout/render/markdown/image/theme/view tests and gitignore tests. |
| Swarm | Supervisor, persistence, protocol, socket, event, child-mode, TUI, and stub-subprocess tests described above. |
| Extensions | Manager/process, command/tool, interception/event, diagnostics, reload, SDK, CLI, and TUI hook tests described above. |

Observed gaps to preserve explicitly:

1. **No single full production-composition E2E:** no test was found that starts the real command composition and drives config/auth/session/prompt/tools/provider plus a real TUI or headless mode through one production-shaped fake-provider scenario.
2. **No dedicated core compaction test:** `packages/core/compact.go` has no corresponding focused `compact_test.go`; interactive auto/manual compaction behavior is tested above the core primitive.
3. **No common cross-provider contract suite:** provider adapters have many focused tests, but no one shared scripted neutral conversation suite is run uniformly against every adapter.
4. **No Anthropic signed-thinking replay test:** Anthropic thinking request configuration is tested, but the stream parser currently recognizes `thinking` / `thinking_delta` and does not surface or persist them; no signed-thinking capture-and-replay contract was found.
5. **No clear full Zotfile Net/Env enforcement proof:** current tests prove manifests requesting Net/Env are rejected, not that such permissions are enforced end to end.
6. **No live multi-agent concurrent-edit proof:** swarm tests establish same-CWD behavior and subprocess protocol with fakes/stubs, not safe simultaneous model-backed edits to shared files.

Absence claims above are repository test-inventory observations, not claims that a behavior can never be exercised manually or by an external system.

## Remaining design questions

Capability disposition is settled by the register above. The remaining questions concern implementation contracts, evidence, and sequencing only.

### Interface and composition design

- What Bubble Tea component architecture cleanly separates Bubbles-based primitives, renderers, mode policy, and application composition?
- What exact ncode names should be used for the binary, module/import path, state and environment identifiers, release assets, and internal protocols?
- What event and middleware schemas should replace mutable host callbacks, and how will normalized agent events be versioned?
- How should host/core session ownership work across interactive, print/stream/JSON, optional persistent RPC, bot, and swarm-child lifecycles?
- How should append-oriented session format versions evolve while retaining repair, resume, branching, and compaction behavior?

### Providers, catalog, and accounting

- What is the exact precedence among built-in, cached, live-discovered, custom-provider, custom/local, and managed model entries, and when does each layer refresh?
- What common provider-contract evidence must every retained adapter pass, and which behaviors remain adapter-specific?
- What normalized quota/rate/subscription-limit model can represent the data providers actually expose?

### Safety, extensions, and swarm

- Should containment harden the inherited implementation or replace it, and what verifiable guarantees can ncode accurately claim?
- What should extension protocol v2, SDK evolution, manifests, and Bubble Tea UI hooks expose?
- How are ordinary interceptors and declared security guards identified, configured, and tested across modes?
- Should optional worktree isolation be a general agent tool, and what lifecycle and conflict semantics would it require?
- What concurrency, permission, and cost controls should govern `swarm_spawn` and resumed swarm agents?

### Evidence and sequencing

- Which production-shaped composition E2E, compaction, cross-provider, replay-safety, and live swarm concurrency tests are required for implementation acceptance?
- What implementation sequence delivers these decisions while preserving the anchored baseline and avoiding accidental protocol or state breakage?
