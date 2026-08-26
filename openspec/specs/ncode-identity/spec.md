# Ncode Identity Specification

## Purpose

Establish lowercase `ncode` as the sole active product identity while preserving the settled retained capabilities of the inherited application. Zot references are permitted only when they are factual, contextual provenance or legal attribution. Neutral protocol names and frame shapes that carry no Zot identity MAY remain unchanged.

## Requirements

### Requirement: Complete identity inventory

Before the change is accepted, the repository MUST contain an evidence-backed inventory of every identity-bearing occurrence in module/import/package comments; command/build/install/release/update assets; durable and project-local paths; environment variables; RPC, swarm, subprocess, and internal protocol identifiers; product-facing help, errors, logs, documentation, examples, tests, and fixtures; extensions, SDKs, manifests, and handshakes; and provenance, baseline, upstream, and license references. Each occurrence, including generated inputs, runtime-composed names, package metadata, example consumers, and test fixtures, MUST be classified as rename/remove, retain as provenance, or not applicable. Neutral names and frame shapes that carry no Zot identity MUST be recorded as not applicable rather than treated as active Zot branding. The inventory MUST identify its location and evidence, and no active identity-bearing occurrence may remain unclassified.

#### Scenario: Inventory covers active and generated surfaces

- GIVEN the repository contains source, tests, fixtures, examples, release configuration, installers, workflows, and runtime-composed identifiers
- WHEN the identity inventory is reviewed before acceptance
- THEN every occurrence in each required category has an evidence location and one permitted classification
- AND the review records no unclassified active identity surface.

#### Scenario: Neutral protocol material is distinguished from branding

- GIVEN an existing protocol name or frame shape carries no Zot identity
- WHEN the inventory classifies that material
- THEN it is classified as not applicable
- AND it is not treated as a reason to change protocol behavior solely for the identity break.

#### Scenario: Provenance is distinguished from live branding

- GIVEN a Zot reference identifies an inherited baseline, an upstream source, or an applicable license/notice
- WHEN the inventory classifies that reference
- THEN it is classified as retained provenance with its factual context
- AND it is not treated as an active product command, path, protocol, compatibility promise, or brand.

### Requirement: Canonical module and distribution identity

The system MUST use `github.com/nlf/ncode` as its canonical Go module and repository-owned import path. The command directory, built binary, `go install` target, build targets, scripts, installers, release archives, checksums, update discovery, and release metadata MUST identify the product as lowercase `ncode`. Distributed product assets MUST NOT provide a `zot` binary, command dispatch, Go import alias, module replacement, duplicate release asset, installer alias, or update endpoint fallback.

#### Scenario: A user installs and invokes ncode

- GIVEN a user follows a supported installer, release, build, or `go install` instruction
- WHEN installation completes and the user invokes the installed command
- THEN the command and generated/distributed asset are named `ncode`
- AND its module/import documentation identifies `github.com/nlf/ncode`.

#### Scenario: A Go embedder uses the canonical module

- GIVEN a Go consumer imports a repository-owned ncode package using `github.com/nlf/ncode/...`
- WHEN the consumer is built against the released module
- THEN the import resolves through the ncode module identity
- AND package-facing product comments identify ncode where they describe the product.

#### Scenario: A legacy binary or import is requested

- GIVEN a user or build requests a Zot command, Zot build/install target, Zot release asset, or a Zot compatibility import from this product
- WHEN the request is evaluated
- THEN ncode provides no alias, shim, replacement, redirect, or alternate dispatch for that request
- AND no Zot-named product asset is distributed.

### Requirement: Ncode-only durable and project-local state

The system MUST resolve and create durable state only under ncode-owned roots and native project resources only under ncode-owned project paths. It MUST NOT read, probe, import, copy, migrate, fall back to, auto-discover, create, modify, or delete any Zot root, `.zot` project directory, or Zot-named temporary or persistence path. Non-branded state-file basenames, including `config.json`, `auth.json`, and session JSONL filenames, MUST remain unchanged unless the basename itself contains Zot identity. The existing lazy or eager creation timing for every state category MUST be recorded and preserved; identity resolution alone MUST NOT create state that the baseline did not create at the same lifecycle point.

#### Scenario: Fresh ncode state uses the ncode root

- GIVEN a fresh environment with no pre-existing application state
- WHEN ncode executes a lifecycle that creates configuration, credentials, sessions, caches, logs, extensions, skills, swarm data, bots, models, themes, or update data
- THEN any created product root and project-local resource uses ncode identity
- AND retained non-branded filenames retain their baseline names.

#### Scenario: Zot and ncode roots coexist

- GIVEN both a populated Zot root and a populated ncode root are present, with distinguishable configuration, credentials, sessions, caches, logs, extensions, skills, swarm data, and bot data
- WHEN ncode starts, runs a supported mode, and exits
- THEN its behavior is determined only by ncode-owned inputs
- AND the Zot root and `.zot` project resources are neither read nor modified
- AND no Zot data is copied into ncode state.

#### Scenario: State creation timing is unchanged

- GIVEN baseline lifecycle evidence identifies which state categories are created eagerly and which are created only by a triggering operation
- WHEN the equivalent ncode lifecycle runs without that trigger
- THEN it creates no additional state category or file solely because of the identity change
- AND a triggering operation creates the corresponding ncode state at the same lifecycle point as the baseline behavior.

#### Scenario: A legacy state location is the only available location

- GIVEN only a Zot state root, `.zot` project resource, or Zot-named persistence/temporary path exists
- WHEN ncode starts or a feature that would normally consume state runs
- THEN ncode does not reuse, migrate, warn-assisted-convert, or fall back to that location
- AND it follows the normal fresh-ncode behavior or reports the normal ncode failure for missing required state.

### Requirement: Ncode-only environment configuration

The system MUST expose ncode-branded names for every product environment control, including home, UI/image/theme, browser/debug, skills/consent, API-key command test, RPC, and swarm metadata controls. The RPC token environment variable MUST be `NCODE_RPC_TOKEN`. `ZOT_*`, `ZOTCORE_*`, including `ZOTCORE_RPC_TOKEN`, and other legacy Zot-branded variables MUST be ignored and MUST NOT alter ncode configuration, authentication, rendering, consent, process execution, state location, protocol authorization, or swarm behavior. Standard provider and operating-system variables that are not Zot-branded MUST remain unchanged.

#### Scenario: An ncode environment control is supplied

- GIVEN a supported ncode-branded environment variable is set to a valid value
- WHEN the corresponding feature runs
- THEN the feature honors that ncode variable according to its documented contract.

#### Scenario: The ncode RPC token is supplied

- GIVEN `NCODE_RPC_TOKEN` is set to a valid token
- WHEN an RPC interaction requires token authorization
- THEN ncode uses that token according to the existing RPC authorization contract.

#### Scenario: Only a legacy environment control is supplied

- GIVEN a legacy Zot-branded environment variable, including `ZOTCORE_RPC_TOKEN`, is set and its ncode counterpart is unset
- WHEN ncode runs the affected feature
- THEN the legacy value has no effect
- AND ncode does not treat it as a compatibility alias or derive configuration from it.

#### Scenario: Legacy and ncode controls coexist

- GIVEN conflicting legacy Zot-branded and ncode-branded values are both present
- WHEN ncode resolves the affected configuration
- THEN only the ncode-branded value may influence behavior
- AND ncode does not read or report the legacy value as an active setting.

### Requirement: Ncode identity in runtime and integration contracts

The system MUST rename active product identity in RPC executable invocation, documentation, examples, token configuration, and product-bearing diagnostics to lowercase `ncode`; the RPC executable contract MUST be `ncode rpc`, and its token configuration MUST be `NCODE_RPC_TOKEN`. Existing neutral RPC protocol names, prompt and event frame shapes, and optional hello semantics that carry no Zot identity MUST be preserved. The system MUST NOT require an always-on product-bearing RPC hello, introduce RPC protocol version 2 solely for identity, or reject an otherwise valid current neutral prompt or event frame merely because an older client could send it.

The system MUST use ncode identity for swarm child invocation and metadata, subprocess assumptions, product-bearing socket/log labels and diagnostics, extension discovery, SDK imports, manifests, and extension handshake product/version fields. It MUST NOT accept or offer an explicit Zot-branded RPC token variable, product field, path, command, subprocess command, swarm metadata field, extension path or manifest fallback, SDK import compatibility, or legacy product-bearing handshake. Neutral protocol names and shapes MAY remain when they contain no Zot identity.

Because the current extension host acknowledgment explicitly carries `zot_version`, the extension protocol MUST provide a signaled version 2 acknowledgment with ncode product and version fields. Existing neutral extension idle auto-ready behavior MUST be preserved, with only its product-facing diagnostic rebranded to ncode. Unrelated retained extension behavior MUST NOT be removed or changed solely for the identity break.

#### Scenario: Ncode RPC is invoked with its documented identity

- GIVEN an RPC client follows the documented ncode contract
- WHEN it launches `ncode rpc` and supplies `NCODE_RPC_TOKEN` when authorization is required
- THEN the RPC interaction uses the existing neutral protocol frames and optional hello behavior
- AND product-bearing documentation or diagnostics identify ncode.

#### Scenario: A current neutral RPC frame is received

- GIVEN an RPC client sends an otherwise valid current prompt or event frame with no Zot-branded field
- WHEN ncode processes the frame without a hello or with the existing optional hello behavior
- THEN ncode preserves the current frame-handling behavior
- AND it does not require a product-bearing hello or protocol version 2 solely because of the identity change.

#### Scenario: An explicit Zot-branded RPC identity is presented

- GIVEN an RPC client supplies `ZOTCORE_RPC_TOKEN` or an explicit Zot-branded RPC product field
- WHEN ncode processes the interaction
- THEN the Zot-branded identity does not authorize or alter the interaction
- AND ncode does not offer an alias, fallback, or dual product identity.

#### Scenario: Ncode swarm and extension integrations interoperate

- GIVEN a swarm parent or child, subprocess extension, or SDK consumer follows the documented ncode contract
- WHEN it connects or launches under ncode
- THEN the interaction uses ncode command, environment, path, log/socket label, import, manifest, and product-bearing handshake identities as applicable
- AND the retained swarm, subprocess-extension, and SDK capabilities remain available.

#### Scenario: Extension host acknowledgment is rebranded through version 2

- GIVEN an extension host and extension negotiate the signaled extension protocol version 2 acknowledgment
- WHEN the host sends its product/version acknowledgment
- THEN the acknowledgment contains ncode product and version fields instead of `zot_version`
- AND the extension does not receive a Zot-branded product field or dual handshake.

#### Scenario: Neutral extension idle auto-ready behavior is retained

- GIVEN an extension is idle under the existing auto-ready conditions
- WHEN the host reports or performs that auto-ready behavior
- THEN the readiness behavior remains unchanged
- AND any product-facing diagnostic identifies ncode.

#### Scenario: A legacy integration identity is presented

- GIVEN a swarm launcher requests a Zot child command or metadata, or an extension presents a Zot path, manifest, SDK identity, or product-bearing handshake field
- WHEN ncode processes the request
- THEN it does not launch, load, or negotiate through that explicit Zot-branded identity
- AND it does not retry with, infer, or advertise a Zot compatibility identity.

#### Scenario: A protocol failure is reported under ncode identity

- GIVEN a ncode RPC, swarm, subprocess, or extension interaction fails validation or startup
- WHEN ncode reports the failure through its documented channel
- THEN any product-facing diagnostic identifies ncode
- AND the failure does not trigger a legacy fallback or dual product handshake.

### Requirement: Active product communication uses lowercase ncode

All active product-facing command help, version output, errors, logs, documentation, examples, issue templates, test names and expectations, fixtures, package comments, installer text, release/update messaging, and example consumer contracts MUST identify the product as lowercase `ncode`. Product-owned active materials MUST NOT present Zot as a current command, state path, environment variable, release source, SDK/import, extension identity, or explicit product-bearing protocol integration. Neutral protocol names and shapes that contain no Zot identity MAY remain. Negative-test data MAY contain a Zot identifier only when it explicitly verifies rejection or non-use of the legacy identity.

#### Scenario: A user reads active product material

- GIVEN a user reads current help, an error or log, a product document, an installer, a release/update message, an example, or an issue template
- WHEN the material identifies the product or a product-owned integration surface
- THEN it uses lowercase `ncode`
- AND it does not instruct the user to invoke, install, configure, or integrate with Zot.

#### Scenario: A neutral protocol is documented

- GIVEN a product document or example describes a neutral RPC frame or optional hello behavior
- WHEN that material contains no Zot identity
- THEN it MAY preserve the existing neutral name and shape
- AND any product-bearing invocation, token name, or diagnostic uses ncode.

#### Scenario: A regression fixture verifies legacy rejection

- GIVEN a test fixture includes a Zot-named command, path, variable, or explicit product-bearing protocol input
- WHEN the test suite evaluates that fixture
- THEN the fixture is explicitly scoped to proving legacy rejection or non-use
- AND it does not represent Zot as an active supported product identity.

### Requirement: Clean-break communication without runtime prompting

Release notes, installer/install documentation, and first-run/setup documentation MUST prominently state that the ncode release is a clean break and does not reuse Zot credentials, settings, sessions, caches, extensions, SDK/protocol integrations, or other Zot state. These materials MUST direct users to the ncode contract rather than a migration path. The runtime MUST NOT display, require, or depend on a compatibility, migration, or conversion prompt for Zot state or integrations.

#### Scenario: A user encounters each required communication channel

- GIVEN a user reads the release notes, supported install documentation or installer output, and first-run/setup documentation
- WHEN the material explains identity and state behavior
- THEN each channel prominently states the clean break and absence of Zot state/integration reuse
- AND none offers a migration, importer, fallback, or Zot compatibility instruction.

#### Scenario: Ncode starts with legacy data present

- GIVEN Zot state or integrations are present when ncode starts
- WHEN ncode initializes or executes a supported runtime mode
- THEN it does not display a migration or compatibility prompt
- AND it does not pause or alter normal runtime behavior to solicit legacy-data conversion.

### Requirement: Contextual provenance and legal attribution

The system MUST retain factual Zot references needed to identify the inherited baseline/tag/commit, the `patriceckhart/zot` upstream repository or upstream-watch policy, historical capability evidence, and applicable MIT LICENSE or source notices. Each retained active-document reference MUST be contextualized as provenance, upstream attribution, historical evidence, or legal attribution and MUST NOT imply active Zot branding, an executable command, a state location, a supported product-bearing protocol, or compatibility.

#### Scenario: A maintainer reviews attribution material

- GIVEN a maintainer reads architecture records, inherited-capability evidence, upstream-watch material, or LICENSE/notices
- WHEN a Zot reference appears
- THEN the reference accurately identifies provenance, upstream, history, or legal attribution
- AND surrounding context makes clear that ncode is the active product identity.

### Requirement: Retained capabilities survive the identity break

The identity change MUST preserve the settled retained capabilities and their documented behavior under ncode identity: provider-neutral core operation, inherited providers and authentication paths, headless print/stream/JSON/RPC modes, sessions, tools and permissions, swarm, extensions, skills, themes, update capability, and the retained Telegram support level. It MUST NOT revive Zotfiles, remove a retained capability, redesign provider or neutral RPC wire compatibility, alter general state schemas beyond root identity, or change unrelated interactive, composition, or extension behavior solely to provide or remove Zot compatibility.

#### Scenario: Retained workflows run as ncode

- GIVEN representative existing capability tests and documented ncode workflows for provider/authentication, core turns, headless modes, sessions, tools, swarm, extensions, skills, themes, updates, and Telegram support
- WHEN they run using ncode identities and isolated ncode state
- THEN their retained behavioral assertions continue to pass
- AND their product-facing identities are ncode.

#### Scenario: Removing legacy identity does not remove a capability

- GIVEN a retained workflow previously used a Zot-named command, path, variable, or explicit product-bearing protocol identifier
- WHEN its caller is updated to the documented ncode contract
- THEN the workflow remains available with its retained behavior
- AND no Zotfiles feature, Zot compatibility alias, or Zot state migration is introduced to make it work.

#### Scenario: Focused negative and retained-behavior verification is performed

- GIVEN the completed identity change is validated
- WHEN focused negative tests exercise Zot binary/import, state/project-path, environment, explicit product-bearing RPC identity, swarm, subprocess, extension, and SDK inputs
- THEN each explicit Zot-branded input is rejected or ignored without fallback as specified
- AND retained-capability verification, including neutral current RPC frames and extension idle auto-ready behavior, remains green under ncode identity.
