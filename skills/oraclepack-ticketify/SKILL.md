---
name: oraclepack-ticketify
description: Generate a deterministic, runner-ingestible Ticket Action Pack from `.tickets/` (or `tickets/`) plus canonical planning artifacts (_tickets_index.json, _actions.json/_actions.md, and a Task Master PRD). Use when planning/execution should be driven by ticket Markdown files.
metadata:
  short-description: Turn `.tickets/` into a deterministic oraclepack Action Pack + Task Master PRD
---

# oraclepack-ticketify

## Decision (implementation direction)

**Option A (new skill `oraclepack-ticketify`)** is the intended design:
- Preserves existing `oraclepack-taskify` Stage-3 contract and responsibilities.
- Keeps ticket discovery + ticket normalization isolated from oracle-out ingestion.
- Produces a deterministic **Ticket Action Pack** (single `bash` fence) aligned to oraclepack runner constraints.

## Quick start

When invoked, this skill produces:
- A deterministic ticket inventory index.
- A deterministic action plan derived from tickets.
- A Task Master-compatible PRD document at a deterministic path.
- A **Ticket Action Pack** markdown file containing **exactly one** `bash` fence with **exactly 20** numbered steps.

Default inputs assume tickets live in `.tickets/` (fallback to `tickets/` if `.tickets/` is absent).

## Interface

#
## Setup wizard (always run)

For each available argument listed in this skill, ask the user:

```
1) Use default
2) Enter custom value
```

Collect a choice for every argument (no skips). Then output a summary table of chosen values before running.

For each available argument listed in this skill, ask the user:

```
1) Use default
2) Enter custom value
```

Collect a choice for every argument (no skips). Then run the workflow using the resolved values.

## Inputs (KEY=value args)

Parse args from the user’s invocation as `KEY=value` pairs:
- **Last-one-wins** if repeated.
- **Unknown keys ignored**.
- Values may be unquoted or quoted; treat them as raw strings (no shell expansion).

Supported keys (all optional):

- `tickets_dir`
  Default: auto:
  1) if `.tickets/` exists use `.tickets`
  2) else if `tickets/` exists use `tickets`
  3) else fail-fast (see Failure modes)

- `ticket_glob`
  Default: `**/*.md` (recursive).
  Use `*.md` explicitly if non-recursive behavior is desired.


- `include` / `exclude`
  Optional comma-separated list of globs/paths.
  - Apply `include` first (keep matches), then `exclude` (remove matches).
  - Matching is against repo-relative, forward-slashed paths.

- `top_n`
  Default: `25`
  Clamp: `1..200` (fail if not an integer)

- `out_dir`
  Default: `.oraclepack/ticketify`
  Deterministic output directory for generated artifacts.

- `pack_path`
  Default: `<out_dir>/ticket-action-pack.md`

- `prd_path`
  Default: `.taskmaster/docs/tickets_prd.md`

- `mode`
  Default: `backlog`
  Supported: `backlog|pipelines|autopilot`
  - `backlog`: generate PRD + run Task Master parse/analyze/expand steps in the pack.
  - `pipelines`: additionally generate a pipeline sketch doc at `<out_dir>/pipelines.md` (still deterministic, minimal).
  - `autopilot`: additionally generate a guarded autopilot entry doc at `<out_dir>/autopilot.md` (does not execute by default).

### Examples (invocations)

1) Defaults (prefer `.tickets/`, else `tickets/`)
- `oraclepack-ticketify`

2) Explicit tickets dir + recursion + tighter selection
- `oraclepack-ticketify tickets_dir=.tickets ticket_glob=**/*.md include=backend/**,frontend/** exclude=**/wip/** top_n=40`

3) Custom output paths
- `oraclepack-ticketify out_dir=.oraclepack/ticketify-demo pack_path=.oraclepack/ticketify-demo/pack.md prd_path=.taskmaster/docs/demo_prd.md`

4) Pipelines mode
- `oraclepack-ticketify mode=pipelines`

5) Autopilot mode (guarded)
- `oraclepack-ticketify mode=autopilot`

## Deterministic workflow

Follow this procedure exactly:

1) **Locate tickets directory**
   - If `tickets_dir` is provided, use it.
   - Else prefer `.tickets/` if it exists; else `tickets/` if it exists; else fail-fast.

2) **Enumerate candidate ticket files**
   - Build the candidate set via `ticket_glob` within `tickets_dir`.
   - Normalize each candidate to a repo-relative path with forward slashes.
   - Apply `include` then `exclude` filters if provided.
   - Sort lexicographically by normalized path (stable ordering).

3) **Select top N**
   - Clamp `top_n` to `1..200`.
   - Take the first N tickets from the sorted list.

4) **Normalize tickets → `_tickets_index.json`**
   - For each selected ticket:
     - Derive `ticket_id` deterministically (see `references/ticket-normalization.md`).
     - Extract a stable `title` and `summary`.
     - Compute a content hash (sha256) for determinism checks.
   - Fail-fast on duplicate `ticket_id`.
   - Write to `<out_dir>/_tickets_index.json` with stable JSON formatting:
     - UTF-8
     - `indent=2`
     - `sort_keys=true`
     - trailing newline

5) **Derive actions → `_actions.json` + `_actions.md`**
   - Derive candidate actions from Markdown task-list items and “Acceptance Criteria” bullets (heuristics defined in `references/ticket-normalization.md`).
   - Order actions deterministically:
     - by `ticket_path` (already sorted), then by first appearance line number within the file.
   - Assign stable `action_id` values:
     - `<ticket_id>-A01`, `<ticket_id>-A02`, ...
   - Write `<out_dir>/_actions.json` with stable JSON formatting and `<out_dir>/_actions.md` as a human-readable summary.

6) **Generate Task Master PRD**
   - Write a PRD to `prd_path` (default `.taskmaster/docs/tickets_prd.md`) using a deterministic template and ticket-derived content.
   - Do not include timestamps or machine-specific data.

7) **Generate the Ticket Action Pack**
   - Create `pack_path` as a Markdown document that contains:
     - No code fences except **one** fenced block labeled `bash`.
     - Inside the `bash` fence: **exactly 20** steps, numbered `# 01)` .. `# 20)`.
     - Each step must be shell-self-contained (no reliance on variables exported by previous steps).
     - Each step must write outputs to explicit deterministic paths.

8) **Validate pack**
   - Run `python3 scripts/validate_ticket_action_pack.py <pack_path>` and fail if invalid.

9) **Write manifest**
   - Run `python3 scripts/write_manifest.py --out-dir <out_dir> --pack-path <pack_path> --tickets-index <out_dir>/_tickets_index.json --actions-json <out_dir>/_actions.json --actions-md <out_dir>/_actions.md --prd-path <prd_path>`

## Output contract

All outputs must be deterministic and written to these exact paths:

- `<out_dir>/_tickets_index.json`
- `<out_dir>/_actions.json`
- `<out_dir>/_actions.md`
- `<prd_path>` (default `.taskmaster/docs/tickets_prd.md`)
- `<pack_path>` (default `<out_dir>/ticket-action-pack.md`)
- `<out_dir>/manifest.json` (pack_path + inputs + derived outputs)
- If `mode=pipelines`: `<out_dir>/pipelines.md`
- If `mode=autopilot`: `<out_dir>/autopilot.md`

See `references/output-contract.md` for required JSON fields.

## Task Master integration (in the generated Action Pack)

If Task Master integration is enabled (default behavior), the Action Pack must include steps that run:

- `task-master parse-prd <prd_path>`
- `task-master analyze-complexity --output <out_dir>/tm-complexity.json`
  - Fallback if `--output` is unsupported: redirect stdout to the same file path.
- `task-master expand --all`

## Failure modes

Fail-fast means:
- Do not “guess” missing paths or contents.
- Emit a clear `ERROR:` line to stderr and exit non-zero.

Common failures:
- Tickets directory missing:
  - If `.tickets/` and `tickets/` are absent → fail.
- No ticket files matched after filters → fail.
- Any unreadable ticket file → fail.
- Duplicate derived IDs → fail.
- Generated pack violates single-`bash`-fence or 20-step constraints → fail.

If evidence is insufficient (missing repo files required to mirror an existing style), emit a `missing_artifacts` list in your final response (paths/patterns only; no open-ended questions).

## Repository evidence to consult (if present)

Attempt to read these files if they exist; if missing, proceed with the defaults above and report them under `missing_artifacts`:
- `oraclepack-goldpack-taskify_SKILL.md` (TODO: attach if not available)
- `oraclepack-tui.md` (TODO: attach if not available)
- Any sample tickets under `.tickets/` or `tickets/` (if present in the working repo)

## References

- `assets/ticket-action-pack.template.md`: Single-fence 20-step pack template (placeholders must be resolved).
- `assets/prd.template.md`: Deterministic PRD template.
- `references/determinism.md`: Ordering, hashing, and stable formatting rules.
- `references/ticket-normalization.md`: How to derive IDs, titles, summaries, and actions.
- `references/output-contract.md`: Required output schemas.
- `references/minimal-test-plan.md`: Fixture + expected outputs + failure checks.
