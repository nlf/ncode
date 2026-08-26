# Proposal: Clean-break ncode identity

## Intent

Establish lowercase `ncode` as the sole product identity for the inherited Go application. A new installation and every supported runtime, build, extension, and integration surface must use `ncode`, not Zot. This is a deliberate clean break: old Zot commands, imports, state, environment variables, and protocol names are not supported.

The change addresses the present mismatch between the product-owned ncode fork and an implementation that still ships and exposes Zot identity (`go.mod`, `cmd/zot`, `bin/zot`, `ZOT_HOME`, release/install assets, user documentation, and integration surfaces). It makes the product boundary understandable to users, embedders, extension authors, and maintainers without changing the settled retained-capability decisions.

## Product outcome

A user can install and run lowercase `ncode`; it writes and reads only ncode-owned state, advertises ncode in help and diagnostics, and uses ncode names in subprocess, RPC, swarm, extension, SDK, and release interactions. Existing Zot state and integrations are intentionally not recognized. Historical documentation and legal/upstream attribution remain accurate where those references identify the inherited Zot baseline, upstream repository, or applicable license.

## Scope

### Exhaustive identity inventory to map in design

Design must produce a repository-wide, evidence-backed mapping for every category below, with each occurrence classified as **rename/remove**, **retain as provenance**, or **not applicable**. Text search alone is insufficient: generated/release paths, tests, fixtures, example consumers, package metadata, and runtime-composed names must be included.

| Category | Current evidence and required ncode outcome |
|---|---|
| Go module, imports, and package comments | `go.mod` is `github.com/patriceckhart/zot`; production and test imports use that path; package comments describe Zot (for example `packages/agent/config.go`, `extensions/manager.go`, and `skills/skills.go`). Establish `github.com/nlf/ncode` as the canonical ncode module/import path and update package/product comments; no Zot compatibility import path or alias remains. |
| Command directory, binary, build/install targets, scripts, release/update assets | `cmd/zot`, Makefile `bin/zot` targets, `.goreleaser.yaml` project/build/archive/release owner/name, `install.sh`, `install.ps1`, README install instructions, and GitHub release workflow are Zot-named. Move to ncode command and binary names, ncode release repository/assets/checksums/update endpoints, and ncode installer variables; do not ship a `zot` binary alias. |
| State, config, cache, session, auth, model, extension, skill, swarm, and bot paths | `ZotHome()` owns `config.json`, `auth.json`, sessions, logs, model cache/models, docs, themes, update cache, bot files, agents/consents/data, and swarm persistence; native project discovery uses `.zot` for extensions and skills. Create only ncode roots and ncode project paths. Preserve non-branded state-file basenames such as `config.json`, `auth.json`, and session JSONL unless a basename itself contains Zot identity; this is not a state-schema redesign. Record and preserve existing lazy/eager state-creation timing so identity work does not create state incidentally. Do not read, copy, import, migrate, or fall back to `$ZOT_HOME`, OS Zot default directories, `.zot`, Zot session/auth/config/cache files, or Zot-named temporary/persistence files. |
| Environment variables | Rename product variables including `ZOT_HOME`, UI/image/theme controls, browser/debug controls, skill and consent controls, API-key command test controls, and all swarm metadata to ncode names. Old variables must be ignored rather than accepted as fallback aliases. Standard provider or OS variables not branded Zot are unchanged. |
| RPC, swarm, subprocess, and internal protocol identifiers | Rename `ZOTCORE_RPC_TOKEN`, swarm child environment metadata, executable invocation assumptions, socket/log prefixes, JSON/protocol branding, extension handshake product/version fields, and other internal identifiers. `ncode` is the only accepted product protocol identity; no old token variable, subprocess command, protocol-name fallback, or dual handshake is offered. |
| Product-facing help, errors, logs, docs, examples, tests, and fixtures | Rename user-visible command/help/version/error/log wording and product-owned docs, examples, issue templates, example extension clients, RPC clients, fixtures, test names/data, and snapshots where they present product identity. Preserve behavioral assertions while changing identity expectations. |
| Extensions, SDK, manifests, and protocol names | Rename native extension discovery paths, SDK/module imports, manifest/protocol documentation, handshake fields and example extensions. Existing Zot extensions, manifests, SDK imports, project directories, and protocol clients have no compatibility guarantee and must not be loaded through an alias or fallback. The retained extension capability remains; only its Zot identity compatibility is removed. |
| Provenance, baseline, upstream, and license references | Retain factual Zot references that identify the inherited baseline/tag/commit, upstream remote/source (`patriceckhart/zot`), upstream-watch policy, historical capability evidence, and MIT LICENSE/notices for retained code. Such references must be explicit provenance, never live product branding, commands, paths, or compatibility promises. |

### Boundaries and non-goals

- This proposal covers identity ownership only. It must preserve the settled retained capabilities (providers/authentication, core loop, headless modes, sessions, tools, swarm, extensions, and skills) and must not use a rename to prune or redesign them.
- Zotfiles remain removed according to the architecture decision; this change does not revive, rename, or migrate Zotfiles as an identity-compatibility mechanism.
- This change does not replace the inherited interactive UI, redesign composition, alter provider wire compatibility, or change general state schemas beyond relocating/reidentifying state roots.
- No runtime/build compatibility alias is allowed: no `zot` executable shim, Go import alias/module replacement, duplicate release asset, legacy command dispatch, or build/install alias.
- No old-state support is allowed: no import, migration, copying, probing, fallback, auto-discovery, warning-assisted conversion, or dual-read behavior for Zot config, credentials, sessions, caches, logs, extension/skill directories, swarm data, bots, or agent data.
- No old environment or protocol fallback is allowed: `ZOT_*`, `ZOTCORE_*`, legacy RPC/swarm/subprocess names, and legacy extension protocol/SDK identities must not activate ncode behavior.
- This proposal does not erase factual historical/upstream/license references. It explicitly distinguishes those references from active product identity.

## Affected areas

- **Build and distribution:** module metadata, command location, Makefile, GoReleaser, CI/release configuration, installers, update metadata and release lookup.
- **Host identity and durable state:** `packages/agent/config.go`, model/update/session/auth/bot/theme paths, project-local discovery, state-related tests and fixtures.
- **Runtime contracts:** CLI help/router text, RPC token and examples, swarm child runner/environment/socket/log naming, subprocess invocation, extension manager/SDK/extproto handshake and manifests.
- **Product communication:** README, product docs, release notes, installer/install documentation, first-run/setup documentation, examples, issue templates, test data, error/help/log strings, extension and RPC example consumers.
- **Attribution:** `LICENSE`, architecture records, inherited-capability evidence, Git remote/upstream-watch documentation, and any retained source notices.

## Proposal acceptance outcomes

1. A repository inventory classifies 100% of identity-bearing occurrences in the listed categories, including generated/build/release inputs and example consumers, with an evidence location and a rename/remove/provenance decision.
2. The built and installed command is `ncode`; Go builds, `go install`, release archives, installers, update discovery, checksums, and documentation reference only ncode product assets, with no distributed `zot` alias.
3. The module/import graph and package-facing documentation use the canonical `github.com/nlf/ncode` module path and product spelling; repository-owned code and example SDK consumers contain no active Zot import or product identity.
4. A fresh ncode run resolves only ncode state roots and ncode project paths. With both ncode and Zot roots present, it neither reads nor modifies the Zot root; no migration/import/copy/fallback behavior exists; non-branded basenames remain unchanged and state-creation timing remains current behavior.
5. Only ncode-branded environment variables and internal identifiers influence runtime behavior. Supplying a legacy Zot variable/token/protocol identifier alone has no effect and yields no implicit compatibility behavior.
6. CLI/script, Go embedder, extension, RPC, and swarm contracts are all documented and use ncode names end to end. Zot-named extensions, SDK imports, manifests, paths, and protocol clients are unsupported and receive no compatibility guarantee.
7. Help, errors, logs, docs, examples, tests, fixtures, issue templates, and release/update messaging identify the active product as lowercase `ncode`; release notes, installer/install documentation, and first-run/setup documentation prominently state the clean break and absence of Zot state/integration reuse; no runtime compatibility prompt is added; retained Zot references are limited to factual provenance or legal/upstream attribution and are labeled/contextualized accordingly.
8. Focused negative tests demonstrate the absence of Zot binary/import/state/environment/protocol acceptance, while existing retained-capability behavior continues to pass under ncode identity.

## Risks and mitigations

| Risk | Why it matters | Proposal control |
|---|---|---|
| Repository-wide mechanical rename misses a runtime-composed or non-Go surface | A stale path, variable, release URL, or subprocess name can create broken installs or inconsistent identity. | Require the exhaustive classified inventory, category-level negative tests, and searches plus behavioral checks before declaring the change complete. |
| State clean break appears as data loss to existing Zot users | Credentials, sessions, caches, extensions, and swarm state will intentionally not carry forward. | State the break prominently in release/install/product communication; create no importer, migration, automatic copying, or fallback that weakens the policy. |
| External integrations break | Embedders, scripts, extension authors, swarm launchers, and release automation may use Zot names. | Treat as an intentional major integration break, publish ncode contract/docs/examples, and explicitly make no Zot extension/SDK/protocol compatibility guarantee. |
| Over-broad replacement destroys factual attribution | Replacing every `zot` string could corrupt baseline evidence, upstream attribution, remote topology, or MIT notices. | Classify provenance separately and preserve factual baseline/upstream/LICENSE references in context. |
| Review size obscures errors | A repository-wide rename can exceed the 400 changed-line review budget and hide functional changes. | Keep implementation in buildable, identity-category slices within the review budget where practical; isolate generated/mechanical moves from behavioral changes and review provenance exceptions explicitly. Detailed work-unit sequencing belongs in design/tasks. |

## Rollback

Before release, rollback is a normal revert of the ncode identity change and its ncode-only artifacts. After a clean-break release, rollback means shipping a corrective ncode release; it must not restore Zot aliases, consume Zot state, or add legacy environment/protocol fallbacks. Users who need historical Zot data continue using the prior Zot release separately; ncode never imports it.

## Outcome-level implementation sequence

1. Freeze and classify identity surfaces and provenance exceptions, including every CLI/script, Go embedder, extension, RPC, and swarm consumer contract.
2. Establish the canonical `github.com/nlf/ncode` module, ncode command, and build/distribution identity while retaining factual provenance.
3. Establish ncode-only state/path/environment/runtime-protocol identity, preserving non-branded state basenames and existing state-creation timing, and prove legacy inputs are ignored.
4. Publish and validate ncode contracts for CLI/scripts, Go embedders, extensions, RPC, and swarm in dependency and reviewability order; no audience is silently dropped.
5. Publish the clean-break notice in release notes, installer/install documentation, and first-run/setup documentation without a runtime compatibility prompt.
6. Complete category-level negative and retained-behavior verification, then audit the inventory for unmapped active Zot identity.

## Resolved product decisions

The following decisions resolve the automatic-mode product questions and are binding for design and implementation.

- The clean break is intentional: release notes, installer/install documentation, and first-run/setup documentation must state that Zot credentials, sessions, extensions, settings, and other integrations are not reused. No runtime compatibility prompt is added.
- CLI/script users, Go embedders, extension authors, RPC clients, and swarm integrators are all in scope. Their contracts must be documented and implemented in dependency and reviewability order, not prioritized by dropping any audience.
- Non-branded state-file basenames, including `config.json`, `auth.json`, and session JSONL, remain unchanged unless the basename itself contains Zot identity. This change relocates/reidentifies state roots without redesigning state schemas.
- Existing lazy/eager state-creation timing is an invariant: design must record current behavior and implementation must not alter it incidentally.
- The canonical Go module path is `github.com/nlf/ncode`.
- `ncode` remains lowercase in every active product identifier, while Go package identifiers may follow language conventions where required.
- Provenance references remain factual and contextual, including the frozen Zot baseline, `upstream` remote/source attribution, and applicable MIT LICENSE notices.
