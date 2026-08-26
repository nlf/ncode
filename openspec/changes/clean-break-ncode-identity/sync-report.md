# Sync Report: clean-break-ncode-identity

## Status

**synced**

The verified change is PASS with non-blocking warnings. Canonical synchronization was performed without moving the active change to archive.

## Structured status and action context

- `schemaName`: `spec-driven`
- `changeName`: `clean-break-ncode-identity`
- `artifactStore`: `both`
- `applyState`: `all_done`
- `dependencies`: apply `all_done`, verify `all_done`, sync `ready`, archive `blocked until sync`
- `actionContext.mode`: `repo-local`
- `workspaceRoot`: `/Users/nlf/Projects/nlf/ncode`
- `allowedEditRoots`: repository root
- `branch`: `feat/clean-break-ncode-identity-02`
- `review`: disabled; no review receipt required
- `nextRecommended`: `sdd-archive`

## Artifacts and changes

- Copied `openspec/changes/clean-break-ncode-identity/specs/ncode-identity/spec.md` to the previously absent canonical file `openspec/specs/ncode-identity/spec.md`.
- Canonical files updated: `openspec/specs/ncode-identity/spec.md`.
- Change artifact persisted: `openspec/changes/clean-break-ncode-identity/sync-report.md`.
- No production code or tests edited.
- No proposal, design, tasks, verify report, or provenance artifacts changed.

Because the canonical domain spec did not exist, this was a new canonical-spec creation rather than a delta merge. The canonical content is preserved byte-for-byte from the verified domain spec.

## Requirement accounting

- ADDED: none as a delta; all 9 verified requirements were installed as the new canonical domain specification.
- MODIFIED: none.
- REMOVED: none.
- RENAMED: none.

## Collisions and guardrails

- Active same-domain collisions: none found under `openspec/changes/`.
- Legacy flat change spec: none; the domain spec is present.
- Destructive removal or large modification approval: not applicable.
- RENAMED requirement handling: not applicable.
- Archive remains intentionally unperformed.

## Validation and warnings

- Confirmed source domain spec exists and canonical destination was absent before sync.
- Confirmed verification report is clearly passing with no unresolved `FAIL`, `BLOCKED`, `CRITICAL`, or verification blockers.
- Confirmed sync destination is inside the authoritative workspace.
- Verification warnings carried forward honestly: native GoReleaser, PowerShell, ShellCheck, YAML, and TypeScript tooling was unavailable; changed-file statement coverage was 49.2%; live provider calls were not run; the verification report was created after the pinned inventory gate.

## Recommendation

Proceed to `sdd-archive` after the parent lifecycle accepts this sync report. Keep the change folder active until archive executes.
