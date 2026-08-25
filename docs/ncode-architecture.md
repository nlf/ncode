# ncode architecture and roadmap

ncode will be a **product-owned fork of Zot** whose inherited capabilities are evaluated deliberately rather than kept or removed by default. It starts with Zot's Go runtime, provider boundary, TUI, tools, and host capabilities; connects directly to Claude, Codex, and local OpenAI-compatible endpoints; and evolves behind explicit decisions and regression tests. ncode owns the prompt and agent loop. Provider-specific compatibility stays behind provider adapters.

> [!WARNING]
> Claude Pro/Max and ChatGPT Plus/Pro subscription login is an unofficial compatibility path. It reuses public OAuth client identities from first-party CLIs and may violate provider terms, stop working without notice, or cause tokens or accounts to be revoked. API keys and documented provider APIs are the safe default. Subscription OAuth must remain explicit, removable, and clearly labeled as unsupported by the providers.

## Decisions at a glance

| Area | Decision |
|---|---|
| Product shape | Fork Zot, establish ncode's ownership boundaries, and evaluate inherited capabilities individually before deciding what to retain, reshape, defer, or remove. |
| Runtime | Keep the provider-neutral types and tool loop in [`packages/core`](../packages/core) and [`packages/provider/provider.go`](../packages/provider/provider.go). |
| Host | Establish a clear ncode-owned composition layer while evaluating inherited [`packages/agent`](../packages/agent) capabilities individually. |
| Interface | Keep the reusable terminal primitives in [`packages/tui`](../packages/tui); ncode owns the interaction policy above them. |
| Providers | Initially support Anthropic, OpenAI Codex subscription access, and local OpenAI-compatible endpoints. Retain normal API-key access where it shares those adapters. |
| Authentication | Use direct PKCE OAuth and self-managed refresh for subscription access. Do not delegate the agent loop to ACP or a first-party CLI subprocess. |
| Behavioral reference | Treat [oh-my-pi (OMP)](https://github.com/can1357/oh-my-pi) as the Anthropic protocol oracle. Port observed behavior and contracts, not TypeScript structure. |
| Prompt ownership | ncode owns its system prompt, context assembly, tool policy, compaction policy, and multi-turn loop. Transport-required provider identity is a compatibility envelope, not the product prompt. |
| Upstream | Watch [Zot upstream](https://github.com/patriceckhart/zot) for narrow core/TUI/security fixes; do not routinely merge its growing host layer. |

## Goals and non-goals

### Goals

- Ship a small, dependency-light Go coding agent with one understandable execution path.
- Preserve a provider-neutral conversation, tool, event, usage, and stop-reason model.
- Make Claude and Codex subscription access work directly, including token refresh and provider-required request shapes.
- Make local OpenAI-compatible inference a first-class path rather than an afterthought.
- Match current Anthropic behavior using reproducible contracts derived from OMP.
- Preserve read, write, edit, and bash as baseline tools while evaluating the broader default tool surface.
- Keep protocol risk visible to users and isolate volatile compatibility code.
- Make upstream intake selective and reviewable.

### Non-goals

- Predetermine the final feature or provider set before evaluating inherited behavior and dependencies.
- Treat inherited features as automatically retained or removed without an explicit decision; swarm and extensions are the first stated retention requirements.
- Build an ACP client or use ACP as the model execution layer.
- Shell out to Claude Code or Codex CLI to run the agent on ncode's behalf.
- Reproduce a first-party CLI UI or claim provider endorsement.
- Put Anthropic or OpenAI headers, token claims, or stream event names into the core loop.
- Copy OMP wholesale or mirror its TypeScript package architecture.
- Guarantee that unofficial subscription OAuth will remain available.

## Why start from Zot

At the fork point, Zot is MIT-licensed, dependency-light, and approximately 82k lines of Go. Its layers include core (about 4.7k lines), providers (about 17k), TUI (about 10.8k), and a substantial host layer (about 49k). The fork provides mature streaming, terminal, session, tool-loop, swarm, extension, and other host behavior that ncode can evaluate from a working baseline.

Starting over would recreate low-level terminal and stream handling and discard capabilities that may belong in ncode. Forking preserves known-good mechanics while allowing ownership boundaries and product scope to be decided incrementally.

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
HTTP/SSE wire protocols (never exposed to the TUI or core)
```

### Current ownership and evaluation status

This table is a starting inventory, not a final keep/remove list. No inherited capability is removed until it has been evaluated, an explicit decision is recorded, and focused regression tests protect the behavior ncode still needs.

| Current path | Disposition | ncode responsibility |
|---|---|---|
| [`packages/core`](../packages/core) | **Retain** | Provider-neutral transcript, tool loop, events, retries, compaction primitive, and session mechanics. Simplify only when behavior is covered. |
| [`packages/provider/provider.go`](../packages/provider/provider.go) | **Retain and tighten** | Neutral message/content/request/client contract. Opaque replay metadata is allowed; provider wire types are not. |
| [`packages/provider/anthropic.go`](../packages/provider/anthropic.go) | **Retain, then port behavior** | Anthropic Messages serialization, OAuth compatibility envelope, signed thinking replay, stream parsing, and Anthropic-only quirks. |
| [`packages/provider/openai_codex.go`](../packages/provider/openai_codex.go) | **Retain** | ChatGPT Codex private Responses endpoint, account header, reasoning replay, and Codex stream translation. |
| [`packages/provider/openai.go`](../packages/provider/openai.go), [`packages/provider/openai_responses.go`](../packages/provider/openai_responses.go), [`packages/provider/usermodels.go`](../packages/provider/usermodels.go) | **Retain and evaluate** | Preserve API-key OpenAI and configurable local OpenAI-compatible behavior; evaluate adjacent catalog capabilities individually. |
| [`packages/provider/auth`](../packages/provider/auth) | **Retain and evaluate** | Preserve PKCE, loopback/manual callback, credential storage, refresh, and logout; add explicit risk UX for Claude and Codex while evaluating other auth paths separately. |
| [`packages/tui`](../packages/tui) | **Retain** | Terminal input, layout, rendering, markdown, image, and theme primitives. Product dialogs and policy belong to the host. |
| [`packages/agent/tools`](../packages/agent/tools) | **Retain and evaluate** | Preserve read, write, edit, bash, permissions, and their filesystem helpers as the baseline; evaluate the broader tool surface separately. |
| [`packages/ignore`](../packages/ignore) | **Retain if required** | Gitignore-aware file behavior used by the retained tools. |
| [`packages/agent/build.go`](../packages/agent/build.go) and interactive/session code | **Evaluate and reshape** | Establish a clear ncode composition root while preserving useful host behavior until replacement decisions are tested. |
| [`packages/agent/swarm`](../packages/agent/swarm), [`packages/agent/extensions`](../packages/agent/extensions) | **Expected to retain** | Swarm and extensions are current ncode requirements; evaluate their boundaries and dependencies before reshaping them. |
| Other modes and host capabilities under [`packages/agent`](../packages/agent) | **Undecided** | Review capabilities one by one. Document evidence and dependencies before choosing to retain, reshape, defer, or remove anything. |
| Provider catalog and cloud adapters under [`packages/provider`](../packages/provider) | **Undecided beyond initial priorities** | Anthropic, Codex/OpenAI, and local-compatible paths are initial priorities; evaluate other adapters individually rather than deleting them as a group. |
| [`cmd/zot`](../cmd/zot) | **Reshape late** | Keep the build working while product boundaries are established; switch product naming/module paths only after the baseline remains stable. |

Removal requires an explicit capability decision, no unresolved retained dependency, and passing replacement/regression contracts. Line count alone is never a removal criterion.

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

1. the ncode system prompt and standing instructions;
2. context-file discovery and ordering;
3. the enabled tools and confirmation policy;
4. transcript persistence and compaction policy;
5. turn limits, retries, cancellation, and tool execution;
6. all user-visible events and TUI behavior.

The core loop in [`packages/core/agent.go`](../packages/core/agent.go) remains the single model/tool loop. A provider adapter may prepend a transport-required identity block—for example, the Claude Code identity required by Anthropic OAuth—but that block must be assembled inside the adapter and must not replace, rewrite, or leak into ncode's product prompt.

[Jcode](https://github.com/1jehuang/jcode) remains useful as corroborating prior art for direct-provider agent behavior, but it is not the protocol authority. Before a Jcode behavior is used as implementation evidence, record its commit in the relevant fixture or port note; do not rely on an unpinned recollection.

## Authentication and direct subscription access

### Shared mechanism

Subscription login is direct OAuth 2.0 with PKCE:

1. ncode generates a verifier and S256 challenge.
2. It opens the provider authorization URL using the public client identity and registered redirect shape of the first-party CLI.
3. A fixed loopback callback receives the code; a manual/headless flow is used where supported.
4. ncode exchanges the code itself and stores access token, refresh token, expiry, and only the provider metadata needed for inference.
5. Before credential use, ncode refreshes an expired token with a safety margin and persists rotated refresh tokens.
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

Do not resolve long-lived divergence by merging upstream `main`. ncode intentionally diverges in host scope, prompts, provider set, and risk UX.

## Staged execution plan

### Stage 0 — Freeze the baseline

- [x] Record the frozen Zot fork baseline commit: `82191b33a9d54993ce9c85988dc250421623b75b`.
- [ ] Inventory imports, binary entry points, package sizes, and tests across inherited layers.
- [ ] Add the provider contract matrix and fixture provenance format.
- [x] Use lowercase `ncode` as the canonical project spelling in all product-facing artifacts.
- [ ] Confirm the state directory, binary/module rename sequence, and migration policy before renaming paths.

**Exit:** Current behavior is reproducible, the inherited capability inventory is complete, and no scope decision relies on an untested assumption.

### Stage 1 — Establish the ncode composition root

- [ ] Create one ncode composition path for config → credentials → provider → core → TUI.
- [ ] Make ncode's prompt and context assembly explicit and snapshot-tested.
- [ ] Preserve read/write/edit/bash and required permission/filesystem helpers while inventorying other tools separately.
- [ ] Keep the existing command building throughout the transition.

**Exit:** A scripted fake provider can complete a multi-turn tool call through the TUI/print host while inherited capability boundaries remain understood and testable.

### Stage 2 — Evaluate inherited capabilities

- [ ] Inventory each host mode, swarm, extensions, updater, remote-agent, skills, Zotfiles, and their dependency boundaries.
- [ ] Inventory provider adapters and catalog behavior without assuming the final provider set.
- [ ] For each capability, record a retain, reshape, defer, or remove decision with rationale and test evidence.
- [ ] Keep any resulting changes capability-scoped and independently reviewable.

**Exit:** Every inherited capability has a documented disposition; swarm and extensions have supported ncode boundaries; any removal is justified independently and protected by regression tests.

### Stage 3 — Stabilize provider and local contracts

- [ ] Lock the neutral `provider.Client` contract.
- [ ] Add common adapter tests and local OpenAI-compatible fixtures.
- [ ] Verify API-key and local paths contain no subscription-only behavior.
- [ ] Keep Codex encrypted-reasoning replay green while simplifying surrounding catalog code.

**Exit:** Anthropic API-key, Codex, and a fake local OpenAI-compatible endpoint pass the same neutral conversation contract where capabilities overlap.

### Stage 4 — Harden direct OAuth

- [ ] Make unofficial status and ToS/revocation risk explicit in login UX.
- [ ] Test PKCE state/verifier, callbacks, manual flow, token permissions, refresh rotation, logout, and revoked grants.
- [ ] Centralize volatile client/fingerprint/account metadata per provider.
- [ ] Ensure API key remains the safe, straightforward alternative.

**Exit:** OAuth can be enabled, refreshed, revoked, and completely removed without touching the agent loop or local provider path.

### Stage 5 — Port Anthropic behavior from OMP

Execute the ordered Anthropic roadmap above, one contract at a time.

**Exit:** The pinned OMP behavior matrix is either implemented and tested or explicitly marked out of scope with rationale; signed-thinking replay and tool continuations pass.

### Stage 6 — Productize and maintain

- [ ] Complete binary/module/state-directory branding with migration tests.
- [ ] Document supported providers, risk policy, local endpoint setup, and troubleshooting.
- [ ] Add release checks that run unit, contract, race/static, and packaging validations selected for the repository.
- [ ] Begin the upstream watch cadence and record compatibility incidents.

**Exit:** A new user can install ncode, choose API key/unofficial OAuth/local inference, complete a tool-using task, resume safely, and understand the support boundaries.

## Final acceptance checklist

### Architecture

- [ ] ncode has one product-owned prompt and one core model/tool loop.
- [ ] Core and TUI contain no provider OAuth headers, endpoints, or account bootstrap logic.
- [ ] Provider adapters contain all wire-specific behavior and expose normalized events.
- [ ] Every inherited subsystem and provider has a documented, evidence-based disposition.
- [ ] Read/write/edit/bash remain available while the default tool surface is decided explicitly.

### Authentication and providers

- [ ] Claude and Codex OAuth are direct PKCE flows with self-managed refresh and complete logout.
- [ ] Login visibly states unofficial status, possible ToS conflict, and revocation/account risk.
- [ ] API-key operation remains available and does not inherit OAuth compatibility behavior.
- [ ] Codex uses the private Responses route and required account identity without leaking it into transcripts.
- [ ] Local OpenAI-compatible endpoints work with explicit model/capability configuration and receive no hosted-provider identity.

### Anthropic parity

- [ ] OMP source is pinned and fixture provenance is complete.
- [ ] Current fingerprint/billing attestation and account bootstrap are covered by tests.
- [ ] Built-in and custom tool names round-trip correctly.
- [ ] Signed thinking is preserved and replayed without being exposed or altered.
- [ ] Context management, server tools, model quirks, and watchdog behavior are implemented or explicitly deferred.
- [ ] OAuth and API-key request paths remain separately tested.

### Quality and maintenance

- [ ] Pure, golden, fake-HTTP, common-provider, negative, and interruption tests pass offline.
- [ ] Live tests are opt-in and no secret or personal data exists in repository fixtures/logs.
- [ ] Every retained upstream change is reviewed by capability and covered by focused tests.
- [ ] Broad upstream merges are not part of the maintenance process.
- [ ] Documentation identifies the baseline Zot commit, pinned OMP commit, and any Jcode commit actually used as evidence.

## Source map

| Role | Source |
|---|---|
| ncode inherited core loop | [`packages/core/agent.go`](../packages/core/agent.go) |
| Neutral provider contract | [`packages/provider/provider.go`](../packages/provider/provider.go) |
| Zot Anthropic baseline | [`packages/provider/anthropic.go`](../packages/provider/anthropic.go) |
| Zot Codex baseline | [`packages/provider/openai_codex.go`](../packages/provider/openai_codex.go) |
| OAuth and refresh baseline | [`packages/provider/auth/oauth.go`](../packages/provider/auth/oauth.go), [`packages/provider/auth/manager.go`](../packages/provider/auth/manager.go) |
| Host composition to evaluate and reshape | [`packages/agent/build.go`](../packages/agent/build.go) |
| Upstream Zot | [github.com/patriceckhart/zot](https://github.com/patriceckhart/zot) |
| Anthropic behavioral oracle | [github.com/can1357/oh-my-pi](https://github.com/can1357/oh-my-pi) |
| Corroborating direct-provider implementation | [github.com/1jehuang/jcode](https://github.com/1jehuang/jcode) |

This document is the decision record for scope and sequencing. Implementation plans may add detail, but changing provider ownership, the OMP-oracle policy, the unofficial-OAuth warning, or the retained/removed layer boundary requires updating this document first.
