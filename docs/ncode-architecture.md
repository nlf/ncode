# ncode architecture and roadmap

ncode is a **product-owned fork of Zot** with settled capability dispositions and an intentionally incremental implementation path. The canonical [inherited capability decision register](inherited-capabilities.md) retains the Go runtime, provider boundary, every inherited provider and authentication path, headless modes, sessions, tools, swarm, extensions, and skills; replaces the interactive UI with Bubble Tea; and removes Zotfiles and Zot compatibility. ncode owns the prompt and agent loop. Provider-specific compatibility stays behind provider adapters.

> [!WARNING]
> Claude Pro/Max and ChatGPT Plus/Pro subscription login is an unofficial compatibility path. It reuses public OAuth client identities from first-party CLIs and may violate provider terms, stop working without notice, or cause tokens or accounts to be revoked. API keys and documented provider APIs are the safe default. Subscription OAuth must remain explicit, removable, and clearly labeled as unsupported by the providers.

## Decisions at a glance

| Area | Decision |
|---|---|
| Product shape | Fork Zot and implement the settled dispositions in the [capability decision register](inherited-capabilities.md) as reviewable changes. Implementation design and sequencing remain open where the register says so. |
| Runtime | Keep and harden the provider-neutral types and tool loop in [`packages/core`](../packages/core) and [`packages/provider/provider.go`](../packages/provider/provider.go). |
| Host | Preserve one ncode-owned `Resolve` → `Resolved.NewClient` → `Resolved.NewAgent` construction spine with explicit host composers; do not force distinct lifecycle policy through one generic all-mode constructor. |
| Interface | Fully replace the inherited interactive TUI and its UI integrations with Bubble Tea/Bubbles; retain the user-facing capabilities being rebuilt. |
| Providers | Retain and harden every inherited provider, custom/local adapter, and authentication method. Anthropic, Codex, and local OpenAI-compatible contract work remains an early hardening priority, not a scope boundary. |
| Authentication | Use direct PKCE OAuth and self-managed refresh for subscription access. Refresh and rotated-token persistence remain automatic. Do not delegate the agent loop to ACP or a first-party CLI subprocess. |
| Identity and state | Use lowercase `ncode` with no Zot naming compatibility and no Zot state import, migration, or fallback path. |
| Behavioral reference | Treat [oh-my-pi (OMP)](https://github.com/can1357/oh-my-pi) as the Anthropic protocol oracle. Port observed behavior and contracts, not TypeScript structure. |
| Prompt ownership | Start from the inherited system prompt baseline. ncode owns its later evolution, context assembly, tool policy, compaction policy, and multi-turn loop; transport-required provider identity remains a separate compatibility envelope. |
| Upstream | Watch [Zot upstream](https://github.com/patriceckhart/zot) for narrow core, security, and relevant terminal fixes; do not routinely merge its growing host layer. |

## Goals and non-goals

### Goals

- Ship a dependency-light Go coding agent with one understandable construction spine, explicit host composers, and one core model/tool loop.
- Preserve a provider-neutral conversation, tool, event, usage, and stop-reason model.
- Retain and harden all inherited providers and authentication methods, including direct Claude/Codex subscription access and custom/local OpenAI-compatible endpoints.
- Preserve headless modes, sessions, tools, swarm, extensions, and skills while giving them explicit ncode-owned boundaries.
- Replace the inherited interactive TUI and UI integrations with Bubble Tea without discarding the retained user capabilities.
- Match current Anthropic behavior using reproducible contracts derived from OMP.
- Keep protocol and containment risk visible to users and isolate volatile compatibility code.
- Make upstream intake selective and reviewable.

### Non-goals

- Reopen settled capability dispositions during implementation; changes belong first in the [decision register](inherited-capabilities.md).
- Build a Zot compatibility layer or any Zot state import, migration, or fallback path.
- Claim that the inherited containment is a full OS sandbox or promise sandbox guarantees before the implementation is decided and verified.
- Build an ACP client or use ACP as the model execution layer.
- Shell out to Claude Code or Codex CLI to run the agent on ncode's behalf.
- Reproduce a first-party CLI UI or claim provider endorsement.
- Put Anthropic or OpenAI headers, token claims, or stream event names into the core loop.
- Copy OMP wholesale or mirror its TypeScript package architecture.
- Guarantee that unofficial subscription OAuth will remain available.

## Why start from Zot

At the fork point, Zot is MIT-licensed, dependency-light, and approximately 82k lines of Go. Its layers include core (about 4.7k lines), providers (about 17k), TUI (about 10.8k), and a substantial host layer (about 49k). The fork provides mature streaming, terminal, session, tool-loop, swarm, extension, and other host behavior that ncode inherited from a working baseline.

Starting over would recreate low-level terminal and stream handling and discard retained capabilities. Forking preserves known-good mechanics while allowing ncode-owned boundaries and implementations to be introduced incrementally.

The existing MIT license in [`LICENSE`](../LICENSE) remains applicable; preserve required notices when code is retained or redistributed.

### Research reference points

These commits anchor the source review that produced this decision record. Implementation work must pin fresh commits when deriving fixtures or porting behavior.

> **Frozen Zot baseline:** `82191b33a9d54993ce9c85988dc250421623b75b` is the inherited ncode baseline against which tests, capability evaluations, and future changes are compared.

| Source | Commit reviewed | Role |
|---|---|---|
| [ncode fork baseline](https://github.com/nlf/ncode/commit/82191b33a9d54993ce9c85988dc250421623b75b) | `82191b33a9d54993ce9c85988dc250421623b75b` | Inherited Zot implementation and package boundaries. |
| [oh-my-pi](https://github.com/can1357/oh-my-pi/commit/eab72e88e447a4be45bea2bc302995844c0c51a2) | `eab72e88e447a4be45bea2bc302995844c0c51a2` | Anthropic OAuth and Messages behavioral oracle reviewed during architecture research. |
| [Jcode](https://github.com/1jehuang/jcode/commit/d37ffd3dc238bd794dc36d0a72793307cc4d89f4) | `d37ffd3dc238bd794dc36d0a72793307cc4d89f4` | Corroborating direct-provider OAuth and request behavior. |

## Target layers

```text
cmd entry point
      |
      v
ncode host: config, prompt, session, login UX, tool policy
      |                         |
      |                         +--> provider/auth: credentials + refresh
      v
packages/core: transcript + model/tool loop + neutral events
      |
      v
packages/provider.Client
      |------------------|-----------------------|
      v                  v                       v
Anthropic adapter   Codex Responses adapter   OpenAI-compatible local adapter
      |
      v
HTTP/SSE wire protocols (never exposed to the UI or core)
```

The three adapters shown are representative early hardening paths, not the complete retained provider set. Every inherited adapter remains in scope.

### Current ownership and settled disposition

The [inherited capability decision register](inherited-capabilities.md) is authoritative for the full capability-by-capability disposition. This summary maps those settled choices onto the major ownership seams without repeating all 47 decisions. Detailed implementation design and sequencing remain open where noted.

| Current path or capability | Disposition | ncode responsibility |
|---|---|---|
| [`packages/core`](../packages/core) | **Retain and harden** | Preserve the provider-neutral transcript, tool loop, events, retries, compaction, and session mechanics. Characterize current behavior before reshaping it. |
| [`packages/provider/provider.go`](../packages/provider/provider.go) | **Retain and harden** | Preserve the neutral message/content/request/client contract. Opaque replay metadata is allowed; provider wire types are not. |
| Provider adapters under [`packages/provider`](../packages/provider) | **Retain and harden** | Keep every inherited direct, cloud, gateway, vendor, custom, and local provider path. Initial hardening priorities do not limit retained scope. |
| Model catalog and discovery under [`packages/provider`](../packages/provider) | **Reshape and extend** | Preserve built-in, cached, live-discovered, custom, and local model handling while making precedence, refresh behavior, and capability data explicit. |
| [`packages/provider/anthropic.go`](../packages/provider/anthropic.go) | **Retain and harden** | Preserve Anthropic Messages and API-key behavior, then port verified OAuth compatibility, signed-thinking replay, streaming, and model quirks from pinned evidence. |
| [`packages/provider/openai_codex.go`](../packages/provider/openai_codex.go) | **Retain and harden** | Preserve the private Responses route, account header, reasoning replay, and stream translation behind the neutral contract. |
| [`packages/provider/auth`](../packages/provider/auth) | **Retain and harden** | Preserve every inherited authentication method. Subscription OAuth refresh and rotated-token persistence remain automatic; login, logout, and failure UX remain explicit. |
| Inherited interactive TUI and UI integrations | **Replace implementation** | Fully replace the inherited TUI, editor, commands, dialogs, renderers, themes, images, and clipboard integrations with Bubble Tea/Bubbles while retaining their user-facing capabilities. |
| Headless modes and sessions under [`packages/agent`](../packages/agent) | **Retain and harden** | Keep print, stream, JSON event, RPC, session persistence/resume/branching, and related management surfaces. |
| [`packages/agent/tools`](../packages/agent/tools) and [`packages/ignore`](../packages/ignore) | **Retain and harden** | Preserve tools, permissions, and Gitignore-aware behavior. Keep sandbox capability, but do not choose or claim a full sandbox implementation until research settles it. |
| [`packages/agent/build.go`](../packages/agent/build.go) and host callbacks | **Reshape** | Preserve the shared `Resolve` → `Resolved.NewClient` → `Resolved.NewAgent` spine, use explicit host composers, and later replace mutable callbacks with event/middleware contracts. This is composition work, not a second core implementation. |
| [`packages/agent/swarm`](../packages/agent/swarm) | **Retain and harden** | Keep durable swarm with shared repository/CWD as the default. Optional worktree isolation remains research and may become a general agent tool. |
| [`packages/agent/extensions`](../packages/agent/extensions) and skills | **Retain and reshape** | Preserve extensions and skills across interactive and headless modes; evolve trust, protocol, metadata, and Bubble Tea integration explicitly. |
| Telegram bridge | **Retain as-is; defer investment** | Keep the inherited implementation and defer further product investment. |
| Zotfiles | **Remove** | Remove only Zotfiles while protecting the general tools, skills, permissions, and sessions they reuse. |
| [`cmd/zot`](../cmd/zot), Zot naming, and Zot state | **Replace without compatibility** | Move to lowercase `ncode` naming with no Zot naming compatibility and no Zot state import, migration, or fallback. Keep each rename step buildable and reviewable. |

Any future disposition change must update the canonical register first. Removal work must still prove that retained capabilities do not depend on the removed path.

## Provider-neutral contract

[`provider.Client`](../packages/provider/provider.go) is the architectural seam. The core should know only:

- user, assistant, and tool messages;
- text, image, tool-call, tool-result, and opaque reasoning/replay blocks;
- model, system text, tools, output limits, temperature, and a neutral reasoning level;
- normalized stream events, usage, stop reasons, and errors.

Provider adapters own:

- URLs, authentication headers, beta headers, and first-party client fingerprints;
- role/content conversion and provider tool naming;
- cache controls, request IDs, billing attestations, and account bootstrap;
- signed/encrypted thinking payloads and provider-specific replay rules;
- SSE event names, malformed-event tolerance, watchdogs, and error decoding;
- model aliases and provider-only capability quirks.

When a provider requires opaque data to be replayed, the neutral transcript may preserve it as an opaque block with a provider discriminator. Core may store and order that block but must not interpret its fields. This keeps replay possible without turning core into an Anthropic or Codex implementation.

### Prompt and loop ownership

ncode, not OMP, Jcode, Zot upstream, ACP, Claude Code, or Codex CLI, owns:

1. the system prompt and standing instructions;
2. context-file discovery and ordering;
3. the enabled tools and confirmation policy;
4. transcript persistence and compaction policy;
5. turn limits, retries, cancellation, and tool execution;
6. all user-visible events and interactive behavior.

The inherited system prompt is the initial ncode baseline. It already exists in the inherited composition, so establishing ncode ownership does not require a redundant prompt-extraction change; later prompt evolution should be evidence-driven and snapshot-tested.

The core loop in [`packages/core/agent.go`](../packages/core/agent.go) remains the single model/tool loop. A provider adapter may prepend a transport-required identity block—for example, the Claude Code identity required by Anthropic OAuth—but that block must be assembled inside the adapter and must not replace, rewrite, or leak into ncode's product prompt.

[Jcode](https://github.com/1jehuang/jcode) remains useful as corroborating prior art for direct-provider agent behavior, but it is not the protocol authority. Before a Jcode behavior is used as implementation evidence, record its commit in the relevant fixture or port note; do not rely on an unpinned recollection.

## Authentication and direct subscription access

### Shared mechanism

Subscription login is direct OAuth 2.0 with PKCE:

1. ncode generates a verifier and S256 challenge.
2. It opens the provider authorization URL using the public client identity and registered redirect shape of the first-party CLI.
3. A fixed loopback callback receives the code; a manual/headless flow is used where supported.
4. ncode exchanges the code itself and stores access token, refresh token, expiry, and only the provider metadata needed for inference.
5. Before credential use, ncode automatically refreshes an expired token with a safety margin and automatically persists rotated refresh tokens; routine refresh does not prompt the user.
6. Logout removes the complete provider credential record.
7. The provider adapter sends the token directly to the provider; no first-party CLI process or ACP server is involved.

The inherited implementation anchors are [`packages/provider/auth/oauth.go`](../packages/provider/auth/oauth.go), [`packages/provider/auth/manager.go`](../packages/provider/auth/manager.go), and [`packages/provider/auth/store.go`](../packages/provider/auth/store.go). Keep credentials user-only on disk, never print tokens, and never place real credentials in fixtures.

### Claude Pro/Max

| Step | Required behavior |
|---|---|
| OAuth | Mirror Claude Code's PKCE authorization/token exchange, callback constraints, scopes, and refresh semantics. |
| Request | Send bearer auth to the Anthropic Messages API with the current accepted first-party compatibility fingerprint. |
| Prompt envelope | Prefix the transport-required Claude identity while preserving the ncode system prompt as a separate block. |
| Tools | Map built-in tools to accepted Claude names and namespace custom tools as required by current behavior. Reverse-map tool calls to ncode's registered name. |
| Conversation | Preserve and replay signed thinking blocks exactly when required. Keep tool-use/result ordering stable. |
| Accounting | Send any current OAuth billing/attestation fields required by the service and normalize usage back into the neutral contract. |

The current Zot adapter already has a minimal OAuth shape—bearer auth, Claude identity prefix, beta headers, and casing for known tools—in [`packages/provider/anthropic.go`](../packages/provider/anthropic.go). It is a baseline, not the target protocol specification.

### ChatGPT Codex

| Step | Required behavior |
|---|---|
| OAuth | Mirror Codex CLI's PKCE flow and request offline access. |
| Account bootstrap | Extract the ChatGPT account ID from the returned ID token and retain only the metadata required by inference. |
| Request | POST Responses-shaped requests to `https://chatgpt.com/backend-api/codex/responses`. |
| Headers | Send bearer auth and `chatgpt-account-id`, plus the current request identity/session headers required by the backend. |
| Conversation | Preserve encrypted reasoning items and replay them before related assistant output/tool calls. |
| Refresh | Refresh directly, preserve a prior refresh token if rotation omits a replacement, and require re-login after revocation. |

[`packages/provider/openai_codex.go`](../packages/provider/openai_codex.go) already implements the private Responses route, account header, Responses translation, and encrypted reasoning replay. ncode should narrow and test this behavior rather than redesign it in the core.

### Risk policy

- Label the feature **Unofficial subscription OAuth** at login and in documentation.
- Require an explicit user choice; never silently prefer subscription OAuth over an API key.
- State that client reuse may conflict with provider terms and can be revoked without a migration period.
- Treat HTTP 401/403, invalid grant, missing account, or changed fingerprint as actionable compatibility failures—not reasons to retry indefinitely.
- Make logout complete and make switching to an API key straightforward.
- Keep protocol fingerprints and billing fields centralized so emergency removal is a small patch.
- Never imply that a successful login means the integration is supported by Anthropic or OpenAI.

### Configuration integrity

Configuration is never silently repaired. Interactive mode must show the exact proposed change and ask before writing it. Print, stream, JSON, RPC, swarm-child, bot, and other noninteractive paths must exit with an actionable error and leave the configuration unchanged.

## Local OpenAI-compatible endpoints

Local inference uses the generic OpenAI-compatible adapter, not the Codex subscription adapter. The minimum contract is:

- explicit base URL, model ID, and optional API key;
- OpenAI chat-completions support first, with Responses support only when the endpoint advertises or is configured for it;
- streaming text and function/tool calls normalized into the same provider events as hosted models;
- configurable context/output limits when model discovery is absent or unreliable;
- clear errors for unsupported images, tools, reasoning fields, or response shapes;
- no Claude/Codex OAuth headers, identity prefixes, or account metadata;
- no assumption that `localhost` means safe—TLS bypass and remote binding remain explicit choices.

Local servers vary more than the OpenAI schema suggests. Capability flags belong in provider/model configuration, while the core continues to see one `provider.Client`. Contract tests must include at least one chat-completions fixture and one intentionally unsupported capability case.

## Anthropic: Zot baseline versus OMP target

OMP is the **behavioral protocol oracle** because it tracks the request and replay behavior currently accepted by Anthropic's coding endpoints. Zot's implementation is older and intentionally smaller. OMP is not an architectural oracle: its TypeScript organization, UI, and agent loop are not to be copied into ncode.

| Capability | Zot fork baseline | OMP behavior to verify and port |
|---|---|---|
| OAuth identity | Fixed Claude Code version, identity prefix, beta headers, and bearer token. | Current Cowork/Claude-compatible fingerprint and any required billing attestation. |
| OAuth bootstrap | PKCE, loopback/manual callback, token store, and refresh. | Expanded account/profile/organization bootstrap and current failure handling. |
| Tool names | Canonical casing for a small known set; unknown custom names pass through. | Current built-in naming plus required prefixes/namespacing for custom tools. |
| Thinking | Requests extended/adaptive thinking, but streamed `thinking` and `thinking_delta` are discarded. | Capture signed thinking, preserve block order, and replay signatures without exposing or editing opaque data. |
| Context management | ncode/Zot host compaction and cache breakpoints. | Anthropic-native context-management fields/events where they improve correctness; keep product compaction policy in ncode. |
| Server tools | Client function tools only. | Supported server tools and their stream/result events behind explicit capabilities. |
| Model quirks | Some adaptive-thinking and effort handling. | Current per-model restrictions, token limits, beta requirements, aliases, and sampling rules. |
| Stream resilience | Generic SSE parsing and retries; HTTP client has no overall timeout. | Idle watchdogs, event-specific validation, recoverable stream failures, and safe retry boundaries. |

The most important baseline gap is visible in [`packages/provider/anthropic.go`](../packages/provider/anthropic.go): thinking blocks are currently recognized and discarded. By contrast, [`packages/provider/openai_codex.go`](../packages/provider/openai_codex.go) already demonstrates the neutral pattern of storing opaque reasoning and replaying it on a later request.

### Why OMP is the oracle

- Its Anthropic path is exercised against the current behavior ncode needs to match.
- It covers the complete conversation lifecycle, not only initial authentication.
- It captures details that API documentation does not promise for unofficial OAuth: fingerprint, attestation, custom-tool naming, signed-thinking replay, and model exceptions.
- It provides executable behavior to compare against fixtures.
- It is still only evidence. Every port must be pinned to an OMP commit, translated into an ncode contract, and tested independently.

Useful starting points are the [OMP repository](https://github.com/can1357/oh-my-pi), its [`packages/ai` provider layer](https://github.com/can1357/oh-my-pi/tree/main/packages/ai), and its [`packages/coding-agent` host layer](https://github.com/can1357/oh-my-pi/tree/main/packages/coding-agent). Pin a commit before extracting behavior because `main` is a moving target. The Zot baseline is pinned by this fork's initial commit lineage; upstream source is available at [patriceckhart/zot](https://github.com/patriceckhart/zot).

## Ordered Anthropic port roadmap

Port in this order so each step creates the contract required by the next:

1. **Pin and inventory the oracle.** Record the OMP commit, locate its Anthropic/auth implementations and tests, and write a behavior matrix. No production port begins from an unpinned branch.
2. **Build sanitized wire fixtures.** Capture request headers/body and representative SSE events for API-key and OAuth modes. Redact tokens, account data, signatures that are user-derived, and prompt content.
3. **Update OAuth fingerprint and billing attestation.** Centralize the current compatibility identity and prove exact request shape with an HTTP recorder test.
4. **Port account bootstrap and refresh failures.** Add only metadata required for inference; test rotation, expiry, revoked grant, missing account, and re-login guidance.
5. **Port tool naming.** Cover built-ins, custom tools, collisions, reverse mapping, invalid names, and multiple calls in one turn.
6. **Port signed thinking end to end.** Parse blocks/deltas, retain opaque signatures in order, persist them safely, replay them verbatim, and test tool-call continuations. Never log or render hidden thinking by default.
7. **Complete the model-quirk matrix.** Extended versus adaptive thinking, effort, temperature restrictions, output budgets, beta flags, and aliases are table-driven provider behavior.
8. **Add context management.** Translate supported server context-management fields/events while keeping ncode's decision to compact in the host/core boundary.
9. **Add server tools selectively.** Introduce explicit capabilities and neutral result events only for tools ncode intends to expose. Do not let server-tool wire types leak into core.
10. **Harden streaming.** Add first-event and idle watchdogs, malformed-event tests, disconnect handling, and retries only where duplicate tool execution cannot occur.
11. **Run opt-in live smoke tests.** Validate API-key and OAuth flows manually against a disposable conversation, then update the pinned compatibility record. Live tests never become credentialed default CI.

A later step may not weaken contracts from an earlier step. In particular, context management and watchdog retries must preserve signed-thinking and tool-result ordering.

## Testing and contract fixtures

### Test pyramid

| Level | Purpose | Default CI? |
|---|---|---|
| Pure conversion tests | Neutral messages ↔ provider request blocks; stream events ↔ neutral events. | Yes |
| Golden request fixtures | Exact method, path, safe headers, JSON body, tool names, block ordering, and capability fields. | Yes |
| Recorded SSE fixtures | Text, tools, usage, thinking, errors, malformed events, and interrupted streams. | Yes |
| Fake HTTP end-to-end | PKCE exchange/refresh, request construction, stream loop, retries, cancellation, and watchdog timing. | Yes |
| Core contract suite | Run one common scripted conversation against every adapter and assert neutral events/transcript. | Yes |
| Live provider smoke | Detect drift in unofficial behavior. Uses user-supplied credentials and a low-cost prompt. | No; explicit opt-in only |

### Fixture rules

Each protocol fixture records:

- source project and exact commit (OMP, Jcode when used, or Zot);
- source file/test path or capture method;
- capture date and provider/model family;
- API-key, OAuth, or local-endpoint mode;
- which fields were redacted or normalized;
- the ncode contract the fixture protects.

Fixtures contain no access tokens, refresh tokens, authorization codes, cookies, account IDs, email addresses, private prompts, or user-derived signed content. Replace volatile request IDs and timestamps with deterministic placeholders. Opaque replay values may use synthetic test values only.

### Required contract cases

- [ ] API-key Anthropic request does not receive OAuth-only identity or headers.
- [ ] OAuth Anthropic request has the exact compatibility envelope and keeps the ncode prompt separate.
- [ ] Custom tool names round-trip without collisions.
- [ ] Signed thinking survives stream parsing, persistence, and replay byte-for-byte.
- [ ] Tool-use and tool-result ordering remains valid after resume and compaction.
- [ ] Codex sends the account header and replays encrypted reasoning in order.
- [ ] Expired tokens refresh once; revoked or invalid grants request re-login without a retry loop.
- [ ] Local endpoints never receive subscription headers.
- [ ] Unsupported local capabilities fail clearly rather than emitting a malformed request.
- [ ] Stream watchdog cancellation cannot duplicate an already-started tool execution.
- [ ] Core contract tests pass without knowing the selected provider.

Tests should use `httptest` or an equivalent in-process server and bind only to loopback. Assertions compare semantic JSON plus an explicit allowlist of order-sensitive blocks and headers; avoid brittle full-byte snapshots for harmless map ordering.

## Upstream-watch strategy

The fork already distinguishes `origin` (`nlf/ncode`) from `upstream` (`patriceckhart/zot`). Keep that topology.

| Source | Watch for | Intake rule |
|---|---|---|
| Zot upstream | Core correctness, TUI terminal fixes, shared SSE/HTTP hardening, security fixes, Go/toolchain changes. | Prefer a small cherry-pick or manual port with focused tests. Do not merge the host wholesale. |
| OMP | Anthropic auth/fingerprint, signed thinking, model quirks, context management, server tools, and stream resilience. | Re-derive behavior from a pinned commit, update fixtures, then implement in Go. |
| First-party Claude Code/Codex releases | OAuth redirect/client changes, required headers, endpoint changes, model behavior, revocation signals. | Treat as volatile external evidence; never silently change risk posture. |
| Jcode | Corroborating prompt/loop or direct-provider behavior discussed during design. | Use only with a canonical URL and pinned commit recorded beside the fixture; it does not override OMP for Anthropic protocol behavior. |

### Review cadence

1. Check Zot and OMP releases/commits on a regular maintenance cadence and immediately after a provider compatibility incident.
2. Record candidate changes by retained path/capability, not as a generic “sync upstream” task.
3. Classify each candidate: security, correctness, compatibility, feature, or out of scope.
4. Port one category at a time with the smallest relevant contract test.
5. Update the source commit in fixture provenance and this roadmap if ownership boundaries change.
6. Revisit prior scope decisions only with an explicit ncode requirement and updated evidence.

Do not resolve long-lived divergence by merging upstream `main`. ncode intentionally diverges in host ownership, prompts, UI, identity, and risk UX.

## Staged execution plan

These stages describe dependency and delivery order, not the chronology of discovery or documentation. Stage 2 capability evaluation was completed before Stage 1 implementation and remains satisfied.

The first implementation work proceeds as ordered, separately reviewable units:

1. **Characterize only composition behavior at risk — complete.** Existing tests were audited, and one production-shaped print fixture closed the missing credential/provider/prompt/tool/session wiring gap without starting an exhaustive core test campaign.
2. **Make the shared construction spine explicit — complete for the inherited hosts.** `composeHeadlessAgent` now centralizes print, stream, JSON, and swarm-child assembly over `Resolve` and `Resolved.NewAgent`; RPC, interactive, bot, and SDK remain explicit composers because their lifecycle policies differ.
3. **Apply ncode identity in planned slices.** The [`clean-break-ncode-identity`](../openspec/changes/clean-break-ncode-identity/) OpenSpec change defines the exact inventory, specification, design, and dependency-ordered work units before production names change.

Keep characterization, explicit host composition, each identity slice, the Bubble Tea replacement, and provider hardening separate. Do not introduce a generic callback/options container merely to route unlike hosts through one constructor.

### Stage 0 — Freeze and inventory the baseline

- [x] Record the frozen Zot fork baseline commit: `82191b33a9d54993ce9c85988dc250421623b75b`.
- [x] Inventory inherited entry points, package composition, capabilities, dependency seams, and existing tests in [`inherited-capabilities.md`](inherited-capabilities.md).
- [x] Record the inherited baseline validation evidence and explicit test gaps in that inventory.
- [ ] Complete a package-size inventory across inherited layers.
- [ ] Add the provider contract matrix and fixture provenance format.
- [x] Use lowercase `ncode` as the canonical product spelling.
- [x] Set the identity policy: no Zot naming compatibility and no Zot state import, migration, or fallback.
- [ ] Define the build-preserving binary/module/state-directory rename sequence.

**Exit status:** Baseline anchoring and capability inventory are complete. Package-size evidence, provider contract fixtures/matrix, and rename sequencing remain before the full Stage 0 exit is satisfied.

### Stage 1 — Characterize composition behavior at risk

- [x] Inventory existing tests across only the config → credentials → provider → core → host seams that Stage 3 restructures.
- [x] Identify the concrete retained behavior the composition change could break and existing tests would not detect.
- [x] Add one focused fake-provider headless fixture for the proven assembled-path gap; add no speculative retry, cancellation, compaction, provider, extension, swarm, or unrelated session coverage.
- [x] Protect inherited credential/model, prompt/tool attachment, tool continuation, output, session persistence/resume, and credential non-leakage through the real print path.
- [x] Keep the existing command building throughout characterization.

**Exit status:** Complete. Existing tests plus `TestRunPrintModeComposesResolvedProviderCoreToolsAndSession` protect the composition behavior that moved, and the inherited command remains buildable.

### Stage 2 — Evaluate inherited capabilities

- [x] Inventory host modes, sessions, tools, swarm, extensions, updater, Telegram, skills, Zotfiles, and their dependency boundaries.
- [x] Inventory inherited provider adapters, authentication paths, and catalog behavior.
- [x] Record a disposition, rationale, and constraint for every inherited capability.
- [x] Publish the canonical register in [`inherited-capabilities.md`](inherited-capabilities.md).

**Exit status:** Complete. The decision register settles every inherited capability disposition, including retained swarm/extensions boundaries and the isolated removal of Zotfiles. Implementation evidence is added capability by capability as work proceeds.

### Stage 3 — Establish the ncode construction spine and host composers

- [x] Preserve `Resolve` → `Resolved.NewClient` → `Resolved.NewAgent` as the one shared construction spine without creating another core loop.
- [x] Add `composeHeadlessAgent` for the identical print/stream/JSON/swarm-child assembly while leaving session and output policy with each host.
- [x] Keep RPC, interactive, bot, and SDK as explicit composers rather than hiding their distinct extension, auth, rebuild, daemon, or persistence policy in a generic options container.
- [x] Preserve all retained tools and required permission/filesystem helpers and keep the inherited command building.
- [x] Keep the Stage 1 characterization and full race-enabled suite green through the composition refactor.

**Exit status:** Complete. The shared construction spine and explicit host composers are established; focused characterization and `make test` remain green.

### Stage 4 — Apply ncode identity as a clean break

- [ ] Execute the build-preserving sequence from [`clean-break-ncode-identity`](../openspec/changes/clean-break-ncode-identity/) as separately reviewable functional and mechanical slices before the Bubble Tea replacement.
- [ ] Rename the binary, Go module and imports, state directory, environment variables, and internal-protocol identifiers from Zot identity to lowercase `ncode` identity.
- [ ] Provide no compatibility alias, compatibility import, migration, or fallback for the old Zot identity on any renamed surface.
- [ ] Keep the command building after each slice and add tests proving Zot binary/module/import/state/environment/internal-protocol identity is not accepted or reused.

**Exit:** The binary, module/import graph, state, environment, and internal protocols use only lowercase `ncode` identity, with no Zot compatibility alias/import, migration, or fallback.

### Stage 5 — Replace the interactive UI with Bubble Tea

- [ ] Replace the inherited interactive event loop, editor, commands, dialogs, renderers, themes, image, and clipboard integrations with Bubble Tea/Bubbles.
- [ ] Preserve retained interactive capabilities, including sessions, tool confirmation, swarm, extensions, skills, login, models, and compaction.
- [ ] Define explicit behavior for extension UI hooks and headless-mode equivalents.
- [ ] Remove inherited TUI implementation only after focused replacement contracts pass.

**Exit:** Bubble Tea fully owns the interactive experience, retained capabilities remain available, and headless modes continue to pass their independent contracts.

### Stage 6 — Stabilize all-provider contracts

- [ ] Lock the neutral `provider.Client` contract with a common scripted conversation suite.
- [ ] Add adapter-family fixtures in reviewable increments for every retained provider path; do not prune providers because they are not an initial hardening priority.
- [ ] Add local OpenAI-compatible chat-completions and unsupported-capability fixtures.
- [ ] Verify API-key and local paths contain no subscription-only behavior.
- [ ] Keep Codex encrypted-reasoning replay green while hardening surrounding catalog code.

**Exit:** Every retained adapter is represented in the provider contract matrix, with shared behavior tested uniformly and adapter-specific behavior tested separately. Anthropic API-key, Codex, and fake local paths provide the first complete vertical fixtures.

### Stage 7 — Harden direct OAuth

- [ ] Make unofficial status and ToS/revocation risk explicit in login UX.
- [ ] Test PKCE state/verifier, callbacks, manual flow, token permissions, automatic refresh and rotated-token persistence, logout, and revoked grants.
- [ ] Centralize volatile client/fingerprint/account metadata per provider.
- [ ] Ensure API key remains the safe, straightforward alternative.

**Exit:** OAuth refreshes and persists tokens automatically, can be revoked or completely removed without touching the agent loop, and leaves local/API-key paths isolated.

### Stage 8 — Port Anthropic behavior from OMP

Execute the ordered Anthropic roadmap above, one contract at a time.

**Exit:** The pinned OMP behavior matrix is either implemented and tested or explicitly marked out of scope with rationale; signed-thinking replay and tool continuations pass.

### Stage 9 — Productize and maintain

- [ ] Enforce configuration integrity: interactive repair is proposed and confirmed; noninteractive validation exits without mutation.
- [ ] Replace inherited self-update integration with ncode release, asset, checksum, changelog, and configurable background-check behavior.
- [ ] Document every retained provider, risk policy, local endpoint setup, sandbox/containment limits, and troubleshooting.
- [ ] Add release checks that run unit, contract, race/static, and packaging validations selected for the repository.
- [ ] Begin the upstream watch cadence and record compatibility incidents.

**Exit:** A new user can install and update ncode, choose any retained provider/auth path, complete a tool-using task in Bubble Tea or a headless mode, resume safely, and understand the support and containment boundaries.

## Final acceptance checklist

### Product and architecture

- [x] ncode has one shared construction spine, explicit host composers, the inherited prompt as its explicit initial baseline, and one core model/tool loop.
- [ ] Bubble Tea/Bubbles fully replaces the inherited interactive TUI and UI integrations while preserving the retained user capabilities.
- [ ] Print, stream, JSON event, RPC, and management/headless interfaces remain available and independently tested; one-shot persistence is configurable, and RPC supports explicit ephemeral/persistent operation plus its optional token gate.
- [ ] Core and UI code contain no provider OAuth headers, endpoints, or account bootstrap logic.
- [ ] Provider adapters contain wire-specific behavior and expose normalized events.
- [ ] Sessions, tools, skills, swarm, and extensions remain supported across their documented modes.
- [ ] Zotfiles are removed without removing the general tools, skills, permissions, or sessions they reused.
- [ ] Lowercase `ncode` binary, module/import, state, environment, and internal-protocol naming has no Zot compatibility alias/import, migration, or fallback path.

### Configuration, safety, swarm, and extensions

- [ ] Interactive config repair shows the exact proposal and writes only after confirmation.
- [ ] Noninteractive config validation exits with the config unchanged.
- [ ] Read/write/edit/bash and Gitignore-aware behavior remain available; retained additional tools are changed only through explicit decisions.
- [ ] Sandbox capability remains available, but documentation and UX claim only guarantees proven by the selected implementation.
- [ ] Shared repository/CWD remains the swarm default; optional worktree isolation is not required for acceptance and is not presented as automatic safety.
- [ ] Durable swarm state, model-initiated spawning controls, extensions, subprocess trust disclosure, and skills remain supported.
- [ ] Telegram remains operational at its inherited support level without implying new investment.

### Authentication and providers

- [ ] Every inherited provider and authentication method remains available and has appropriate focused contract coverage.
- [ ] Claude and Codex OAuth are direct PKCE flows with automatic refresh, automatic rotated-token persistence, and complete logout.
- [ ] Login visibly states unofficial status, possible ToS conflict, and revocation/account risk.
- [ ] API-key operation remains available and does not inherit OAuth compatibility behavior.
- [ ] Codex uses the private Responses route and required account identity without leaking it into transcripts.
- [ ] Custom and local OpenAI-compatible endpoints work with explicit model/capability configuration and receive no hosted-provider identity.
- [ ] Model catalog/discovery precedence and refresh behavior are explicit, and usage accounting surfaces quota, rate-limit, or subscription-limit data where providers expose it.

### Anthropic parity

- [ ] OMP source is pinned and fixture provenance is complete.
- [ ] Current fingerprint/billing attestation and account bootstrap are covered by tests.
- [ ] Built-in and custom tool names round-trip correctly.
- [ ] Signed thinking is preserved and replayed without being exposed or altered.
- [ ] Context management, server tools, model quirks, and watchdog behavior are implemented or explicitly deferred.
- [ ] OAuth and API-key request paths remain separately tested.

### Quality and maintenance

- [ ] Existing coverage plus the smallest change-justified characterization fixtures pass; production-shaped headless E2E, pure, golden, fake-HTTP, common-provider, negative, and interruption tests are added and run when their corresponding work changes those behaviors.
- [ ] Live tests are opt-in and no secret or personal data exists in repository fixtures/logs.
- [ ] Every retained upstream change is reviewed by capability and covered by focused tests.
- [ ] Broad upstream merges are not part of the maintenance process.
- [ ] Self-update uses only ncode releases, assets, checksums, and changelogs, with configurable background checks.
- [ ] Documentation identifies the baseline Zot commit, pinned OMP commit, and any Jcode commit actually used as evidence.

## Source map

| Role | Source |
|---|---|
| ncode inherited core loop | [`packages/core/agent.go`](../packages/core/agent.go) |
| Neutral provider contract | [`packages/provider/provider.go`](../packages/provider/provider.go) |
| Zot Anthropic baseline | [`packages/provider/anthropic.go`](../packages/provider/anthropic.go) |
| Zot Codex baseline | [`packages/provider/openai_codex.go`](../packages/provider/openai_codex.go) |
| OAuth and refresh baseline | [`packages/provider/auth/oauth.go`](../packages/provider/auth/oauth.go), [`packages/provider/auth/manager.go`](../packages/provider/auth/manager.go) |
| Inherited host composition to reshape | [`packages/agent/build.go`](../packages/agent/build.go) |
| Upstream Zot | [github.com/patriceckhart/zot](https://github.com/patriceckhart/zot) |
| Anthropic behavioral oracle | [github.com/can1357/oh-my-pi](https://github.com/can1357/oh-my-pi) |
| Corroborating direct-provider implementation | [github.com/1jehuang/jcode](https://github.com/1jehuang/jcode) |

This document records architecture and sequencing. The [inherited capability decision register](inherited-capabilities.md) is authoritative for capability disposition; change that register first when retained, replaced, deferred, or removed scope changes. Changes to provider ownership, the OMP-oracle policy, or the unofficial-OAuth warning also require this architecture document to be updated.
