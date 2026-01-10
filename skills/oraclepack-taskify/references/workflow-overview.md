# Stage 3 (oraclepack-taskify) — Workflow overview

## What this stage solves

Stage 1 produces a 20-question oracle pack.
Stage 2 runs oraclepack and produces 20 answer files.

Stage 2 outputs are answers, not work. Stage 3 creates the deterministic bridge from answers to executable planning artifacts and (by default) starts a guarded autopilot to begin implementation.

## Inputs

- A completed oraclepack output directory containing exactly:
  - `01-*.md` … `20-*.md` (one file per prefix)
- Optional additional files to improve synthesis fidelity (extra attachments)

## Primary output (this skill generates)

- An “Action Pack” markdown file at:
  - default: `docs/oracle-actions-pack-YYYY-MM-DD.md`
  - override: `pack_path=...`

The Action Pack is designed to be executed as a deterministic pipeline.

## Artifacts the Action Pack produces when executed

- Canonical actions:
  - `<out_dir>/_actions.json` (machine-consumable)
  - `<out_dir>/_actions.md` (human summary)
- PRD/spec suitable for Task Master:
  - `.taskmaster/docs/oracle-actions-prd.md`
- Task Master outputs:
  - tasks created/expanded by `task-master`
  - complexity report: `<out_dir>/tm-complexity.json`
- Optional:
  - pipelines doc: `docs/oracle-actions-pipelines.md` (pipelines mode)
  - autopilot entrypoint + state file (autopilot mode, default)

## Execution modes

- backlog: actions → PRD → tasks
- pipelines: backlog + pipelines generation
- autopilot (default): backlog + guarded autopilot entrypoint
