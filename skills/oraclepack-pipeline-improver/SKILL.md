---
name: oraclepack-pipeline-improver
description: Improve an oraclepack (Go wrapper around @steipete/oracle) pipeline by specifying/implementing deterministic validate→run→actionize behavior:strict pack validation, run manifests, stable run directories, resume/rerun semantics, concurrency/backoff, optional caching, and a Stage 3 “Actionizer” that converts 20 oracle outputs into actionable engineering work artifacts.
metadata:
  short-description: Deterministic oraclepack validate/run/actionize pipeline spec + implementation rails
---

## Quick start

Use this skill when the user wants to:

- make oraclepack runs deterministic and resume-safe,
- add a strict validator and machine-readable outputs,
- add Stage 3 “actionize” to convert the 20 question outputs into an actionable backlog/change plan,
- make the pipeline CI-friendly and path-safe.

Interpret the user’s free-text `{{args}}` as the target subset (validate/run/actionize/caching/CI) plus any paths to focus on.

If repo context or current CLI behavior is missing, write **Unknown/TODO** and proceed with a spec-first answer.

## Workflow

### 1) Establish “observed vs proposed”

1. List what inputs are available (repo files, current CLI help text, sample pack md, run output dirs).
2. Split all statements into:
   - **Observed** (backed by provided evidence),
   - **Proposed** (the target contract to implement),
   - **Unknown/TODO** (needs files/flags not provided).

### 2) Define the target pipeline contract (deterministic by default)

Produce a concrete contract for:

- `oraclepack validate` (strict + JSON output),
- `oraclepack run` (stable run dir + `run.json`/`steps.json` + outputs + resume/rerun),
- `oraclepack actionize` (reads run dir and produces `actionizer/` artifacts),
- CI mode behavior (non-interactive, structured logs, policy-driven exit codes),
- Path safety for output writing.

Use:

- references/cli-contract.md
- references/run-manifest-spec.md
- references/actionizer-spec.md

### 3) Map contract → implementation deltas (minimal, additive, backward-compatible)

1. Identify current commands/flags and current on-disk layout (Observed).
2. Propose additive changes:
   - new flags and new subcommands should not break existing pack schema without an explicit migration path,
   - new on-disk outputs should be in `.oraclepack/runs/<pack_id>/...` without removing legacy output locations (unless requested).
3. For each proposed change, specify:
   - code touchpoints (files/modules: **Unknown** if repo not provided),
   - acceptance tests and fixtures,
   - failure modes and user-visible error messages.

### 4) Stage 1 prompt shaping (pack generation) to help Stage 3 parse reliably

If the workflow includes Stage 1 pack generation:

- propose embedding **mini-metadata inside each prompt** (does not change pack schema),
- keep metadata parseable and consistent.

Use references/stage1-prompt-metadata.md.

### 5) Produce final deliverables (spec + plan, optionally code)

Deliverables should be:

1. **Pipeline contract** (validate/run/actionize + CI + safety).
2. **On-disk schemas** (`run.json`, `steps.json`, `normalized.jsonl`, `backlog.md`, `change-plan.md`).
3. **Acceptance criteria** and a minimal test plan.
4. **Implementation plan** (ordered steps, smallest shippable increments).
5. If code context is provided and the user wants implementation: output concrete file edits + new files.

## Output contract

Unless the user asks for something else, output a single Markdown report with:

- **Scope** (what parts of validate/run/actionize/CI/caching are included)
- **Observed current behavior** (or **Unknown**)
- **Proposed contract** (link to reference sections where applicable)
- **Disk layout + schemas**
- **Acceptance criteria**
- **Implementation plan** (phased; smallest first)
- **Risks / unknowns**
- **Missing inputs** (exact paths/flags/help output needed)

If asked to generate templates, use the assets:

- assets/backlog-template.md
- assets/change-plan-template.md
- assets/normalized.example.jsonl

## Failure modes

- Missing repo / CLI help / sample run dirs → mark **Unknown** and provide a spec-first response.
- Missing definitions for CI thresholds / policies → include **TODO** defaults and clearly label them as policy choices.
- Any “current behavior” claim without evidence → downgrade to **Unknown**.

## Invocation examples

1) Add strict validator + JSON output:

- `$oraclepack-pipeline-improver Add oraclepack validate --strict --json; define schema checks and CI gating exit codes`

1) Deterministic run dir + resume/rerun:

- `$oraclepack-pipeline-improver Specify .oraclepack/runs/<pack_id>/ layout, run.json/steps.json, resume default, --rerun failed|all|01,03`

1) Concurrency + backoff policy:

- `$oraclepack-pipeline-improver Add --max-parallel N and transient error retry budget/backoff rules`

1) Stage 3 Actionizer:

- `$oraclepack-pipeline-improver Implement oraclepack actionize; generate normalized.jsonl + backlog.md + change-plan.md with stable IDs`

1) CI mode:

- `$oraclepack-pipeline-improver Provide run --ci --non-interactive --json-log and actionize --ci; policy-driven exit codes`

1) Stage 1 prompt metadata shaping:

- `$oraclepack-pipeline-improver Add prompt-embedded metadata (QuestionId/Category/Reference/ExpectedArtifacts) without changing pack schema`
