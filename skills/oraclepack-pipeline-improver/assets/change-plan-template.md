<!-- # path: oraclepack-pipeline-improver/assets/change-plan-template.md -->
# Oraclepack Change Plan

Run:
- pack_id: TODO
- pack_hash: TODO
- generated_at: TODO

## Principles

- Smallest shippable increments first.
- Every step has an acceptance check.
- Unknowns are explicit; no guessing.

## Phase 0 — Guardrails (validate + safety)

1) Implement/confirm strict validation (validate --strict --json)
- Scope:
  - TODO
- Acceptance:
  - TODO (e.g., rejects non-20 packs; emits JSON summary)
- Tests:
  - TODO (fixtures for invalid packs)

2) Path safety for output writing
- Scope:
  - TODO
- Acceptance:
  - TODO (rejects .. traversal / absolute escape)
- Tests:
  - TODO

## Phase 1 — Deterministic runs (run dir + manifests + resume)

3) Stable run dir + run.json / steps.json
- Scope:
  - TODO
- Acceptance:
  - TODO (creates .oraclepack/runs/<pack_id>/..., stable naming)
- Tests:
  - TODO

4) Resume default + --rerun semantics
- Scope:
  - TODO
- Acceptance:
  - TODO (interrupt + rerun skips completed via hashes)
- Tests:
  - TODO

## Phase 2 — Reliability (concurrency + retries + optional caching)

5) Concurrency cap
- Scope:
  - TODO
- Acceptance:
  - TODO (never exceeds N parallel calls)

6) Retry/backoff on transient errors
- Scope:
  - TODO
- Acceptance:
  - TODO (bounded retries; recorded in steps.json)

7) Optional caching (if enabled)
- Scope:
  - TODO
- Acceptance:
  - TODO (unchanged inputs cause zero provider calls)

## Phase 3 — Actionizer (Stage 3)

8) Implement actionize command and artifacts
- Scope:
  - normalized.jsonl + backlog.md + change-plan.md
- Acceptance:
  - TODO (byte-identical output on rerun with unchanged inputs)

## CI integration (optional)

9) Add CI mode wiring (run --ci --non-interactive --json-log; actionize --ci)
- Policy thresholds:
  - TODO/Unknown
- Acceptance:
  - TODO (exit codes match policy)
