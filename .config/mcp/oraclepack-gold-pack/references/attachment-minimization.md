# Attachment minimization rules (Stage 1 packs)

Objective: keep oracle calls fast, portable, and deterministic by attaching the minimum evidence per step.

## Hard limits

- Default: **0–2 attachments per step** (`-f/--file`).
- If you need more than 2, the step is not scoped tightly enough: split or reduce.
- Any `extra_files` the user provides must be appended **literally** (do not reinterpret), but you should still keep the step’s own attachments ≤2.

## What to attach (rule of thumb)

For a given step, prefer:
1) One file that *defines* the concept (contract/schema/config/type)
2) One file that *enforces/uses* the concept (handler/service/policy)

If you can’t find both confidently, attach only the “definition” file.

## Common attachment choices by category (patterns, not requirements)

Use these as **patterns** to recognize likely candidates in a repo; do not assume these paths exist.

- contracts/interfaces:
  - route registration, API schema/spec, public type definitions, CLI command registry

- invariants:
  - domain model definitions, validation layer, schema types, critical service functions that enforce rules

- caching/state:
  - cache client config, state container/store, session manager, any TTL/invalidation logic

- background jobs:
  - worker entrypoint, job registry, scheduler configuration, queue client config

- observability:
  - logger initialization/config, metrics/tracing setup, middleware that injects correlation ids

- permissions:
  - authn/authz middleware, policy definitions, role/scope mapping, guard functions

- migrations:
  - migrations folder index, migration runner config, schema definition file (if present)

- UX flows:
  - key UI/router flow code, top-level workflow orchestrator, controller/handler representing the flow

- failure modes:
  - error handling utilities, retry/backoff config, boundary middleware, circuit breaker wrappers (if any)

- feature flags:
  - flag config/registry, evaluation hook, rollout/targeting logic

## If you cannot find good attachments

- Attach nothing or only 1 file.
- Set `reference=Unknown`.
- Make the prompt request “exact missing file/path pattern(s) to attach next.”

## Avoid these attachment anti-patterns

- Attaching entire directories when one file is enough.
- Attaching multiple duplicates (e.g., the same config in three copies).
- Attaching generated/lock files unless the question is explicitly about them.
- Attaching secrets.
