# Archive Report: clean-break-ncode-identity

## Status

**PASS — archived successfully.**

## Executive summary

The completed ncode identity change passed verification (9 requirements / 31 scenarios), all 44 implementation tasks and 2 parent tasks are checked, and canonical synchronization succeeded in commit `ff997e2`. The active OpenSpec change was moved intact to the dated archive. No implementation, test, branch, commit, push, PR, review, or canonical-spec mutation was performed by archive.

## Structured SDD status and action context

- `schemaName`: `spec-driven`
- `changeName`: `clean-break-ncode-identity`
- `artifactStore`: `both`
- `applyState`: `all_done`
- `verifyState`: `all_done` / PASS with non-blocking warnings
- `syncState`: `synced` / canonical spec created in `ff997e2`
- `archiveState`: `all_done`
- `actionContext.mode`: `repo-local`
- `workspaceRoot`: `/Users/nlf/Projects/nlf/ncode`
- `allowedEditRoots`: repository root
- `branch`: `feat/clean-break-ncode-identity-02`
- `review`: explicitly disabled; no receipt required
- `nextRecommended`: final delivery (one PR to main, issue #1, exactly `type:breaking-change`)

## Artifacts read

Proposal, domain spec, design, tasks, apply-progress, verify-report, sync-report, and `openspec/config.yaml` were read. The persisted tasks artifact was re-read immediately before archive actions; no unchecked implementation task markers remain.

## Sync and requirement accounting

- Domain: `ncode-identity`
- Canonical destination preserved: `openspec/specs/ncode-identity/spec.md`
- Sync commit: `ff997e2`
- ADDED: all 9 requirements installed as a new canonical domain specification (not a delta)
- MODIFIED: none
- REMOVED: none
- Destructive merge approval: not applicable
- Active same-domain change warning: none

## Completion and evidence

- Tasks: 44/44 implementation, 2/2 parent; unchecked implementation boxes: none.
- Verification: PASS, all 9 requirements / 31 scenarios; no unresolved FAIL, BLOCKED, CRITICAL, or verification blocker.
- Historical evidence preserved, including proposal, spec, design, tasks, apply-progress, verify-report, sync-report, inventory, and canonical spec.
- Non-blocking warnings carried forward: unavailable native GoReleaser, PowerShell, ShellCheck, yq, and tsc; 49.2% changed-file coverage; no live providers by design.

## Archived path

`openspec/changes/archive/2026-08-26-clean-break-ncode-identity/`

The canonical spec remains at `openspec/specs/ncode-identity/spec.md`.

## Risks

Only the recorded non-blocking verification warnings remain. Archive does not resolve native-tool availability, coverage, or live-provider limitations.
