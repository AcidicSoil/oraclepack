<!-- # path: oraclepack-pipeline-improver/references/actionizer-spec.md -->
# Stage 3 “Actionizer” spec (proposed)

Goal: deterministically convert the 20 outputs of a run into actionable engineering work artifacts, without duplicating work on reruns.

## Inputs

- `.oraclepack/runs/<pack_id>/run.json`
- `.oraclepack/runs/<pack_id>/steps.json`
- `.oraclepack/runs/<pack_id>/outputs/*`

## Processing pipeline (deterministic)

1) **Load** run + steps + outputs
- If any output is missing, mark that step as `missing_output` in normalization.

2) **Normalize** each step output into a stable record
- Extract (when present):
  - question metadata (QuestionId/Category/Reference/ExpectedArtifacts),
  - recommended actions,
  - evidence anchors (paths, symbols, commands),
  - unknowns / missing inputs.
- Classify:
  - `actionable` (clear tasks),
  - `blocked` (missing evidence prevents action),
  - `conflict` (contradictory requirements/answers),
  - `noop` (no action required).

3) **Deduplicate** tasks across steps
- Use stable task IDs derived from `pack_hash` + a stable task key (e.g., normalized title + target path).
- Reruns must produce byte-identical outputs when inputs unchanged.

4) **Generate** three core artifacts:
- `normalized.jsonl` (machine-readable records)
- `backlog.md` (human-prioritized tasks)
- `change-plan.md` (ordered implementation plan, smallest-first)

5) Optional exports:
- `github-issues.json` (issue objects; exact schema TODO/Unknown)
- `taskmaster.json` (taskmaster-style import; exact schema TODO/Unknown)

## Output files

### A) `actionizer/normalized.jsonl`

One JSON object per line.

**Proposed record fields:**
- `pack_id` (string)
- `pack_hash` (string)
- `step_id` (string `"01"`..`"20"`)
- `task_id` (string; stable)
- `title` (string)
- `status` (enum: `actionable` | `blocked` | `conflict` | `noop`)
- `category` (string | null)
- `reference` (string | null)
- `expected_artifacts` (string[] | null)
- `actions` (string[]) — concrete “do X” items
- `evidence` (object)
  - `paths` (string[])
  - `symbols` (string[])
  - `commands` (string[])
- `notes` (string[])
- `missing_inputs` (string[])

### B) `actionizer/backlog.md`

Use assets/backlog-template.md as the base.

Rules:
- Group by category (if present), else by subsystem (inferred from reference paths).
- Include blocked/conflict items explicitly with “what evidence is needed”.

### C) `actionizer/change-plan.md`

Use assets/change-plan-template.md as the base.

Rules:
- Ordered, smallest shippable increments first.
- Each step includes acceptance criteria and “done when…” checks.
- If CI gating thresholds exist, include them; else mark TODO.

## Handling blocked/conflict outputs

- `blocked` must include a **single next smallest experiment** (one action) to obtain missing evidence.
- `conflict` must include:
  - conflicting statements,
  - what file/path/log is needed to resolve,
  - a proposed resolution strategy (flagged as Proposed).

## Idempotency / stability requirements

- Stable IDs (required):
  - deterministic function of `pack_hash` + `step_id` + normalized title (or similar stable key).
- Reruns:
  - must not duplicate tasks,
  - must regenerate byte-identical artifacts if inputs unchanged.
