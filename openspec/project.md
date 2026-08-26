# ncode project context

- **Project:** ncode, a product-owned Go fork of Zot becoming a personal coding agent.
- **Canonical spelling:** lowercase `ncode`.
- **Repository:** current workspace; inherited implementation remains Zot-named at initialization.
- **Architecture records:** `docs/ncode-architecture.md` and `docs/inherited-capabilities.md`.
- **Validation:** `make test`, running `go test -race ./...`.
- **Working-tree baseline:** clean at HEAD `480a2c8` at initialization.
- **Current goal:** plan a clean-break rename of Zot identity surfaces to ncode.
- **Identity constraints:** no compatibility aliases, state import, migration, or fallback.
- **Scope constraint:** initialization only; do not create the identity-rename change.
- **Testing policy:** preserve existing conventions; live provider tests are opt-in and repository fixtures contain no credentials.

## Architecture summary

ncode retains and hardens the provider-neutral Go core, provider adapters, authentication paths, headless modes, sessions, tools, swarm, extensions, and skills. The inherited interactive UI will be replaced with Bubble Tea. ncode owns the prompt and agent loop while provider-specific wire behavior remains behind adapters. Zotfiles and Zot compatibility are removed.

## Sequencing context

The architecture roadmap protects the shared `Resolve` → `Resolved.NewClient` → `Resolved.NewAgent` construction spine and uses explicit host composers rather than one generic all-mode constructor before the clean-break identity rename. Future SDD work must keep changes incremental, buildable, and separately reviewable, without modifying product code during initialization.
