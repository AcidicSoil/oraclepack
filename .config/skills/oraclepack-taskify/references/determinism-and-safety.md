# Determinism and safety guardrails

## Determinism rules

- Always select inputs by prefix ordering:
  - exactly one match for each: `01-*.md` … `20-*.md`
- If any prefix is missing or has multiple matches, exit non-zero with a precise error.
- Keep generated JSON normalized:
  - items sorted by id ascending (01..20)
  - stable ordering for arrays
- Keep output paths explicit and stable:
  - do not rely on shared environment variables across steps
  - each step re-declares its constants

## Safety rules

- No interactive prompts in the Action Pack.
- Fail fast when prerequisites are missing:
  - `task-master` (or override)
  - `oracle` (or override)
  - `tm` only in autopilot mode (default)
- Always `mkdir -p` parent directories before writing files.
- Avoid destructive operations:
  - do not delete
  - do not force push
  - do not commit to main/master in autopilot mode
- Autopilot mode:
  - require a clean working tree
  - if on main/master, create a new work branch before starting autopilot
  - write a state file to support resumption

## Failure behavior

- Prefer explicit, early errors over partial or ambiguous outputs.
- If Task Master output paths differ from defaults, print warnings but keep the pack deterministic.
