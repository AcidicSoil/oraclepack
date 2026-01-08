<!-- # path: oraclepack-pipeline-improver/references/cli-contract.md -->
# oraclepack CLI contract (proposed)

This document is the target CLI behavior to implement. If the current repo differs, treat that as **Observed** and this as **Proposed**.

## Commands

### 1) `oraclepack validate`

**Goal:** Deterministically validate a pack file and (optionally) emit machine-readable results.

**Proposed:**
- `oraclepack validate <pack.md> [--strict] [--json]`

**--strict checks (proposed minimum):**
- Exactly **20** oracle invocations.
- No schema drift vs expected pack format (**exact schema rules: TODO/Unknown** unless provided).
- Each question has required fields present (as enforceable in current schema; else **TODO**).
- Stable ordering checks (if applicable): ROI desc, effort asc (only if the pack uses those fields; else skip).

**--json output (proposed):**
- Print a single JSON object to stdout (or to a specified file if supported; **TODO**).
- Include `ok: boolean`, `errors: []`, `warnings: []`, and per-question metadata when parseable.

### 2) `oraclepack run`

**Goal:** Execute the 20 steps into a deterministic run directory with resumable semantics.

**Proposed:**
- `oraclepack run <pack.md> [--max-parallel N] [--resume] [--rerun all|failed|01,03,07] [--ci] [--non-interactive] [--json-log]`

**Run dir (proposed):**
- Create: `.oraclepack/runs/<pack_id>/`
- Emit at least:
  - `.oraclepack/runs/<pack_id>/run.json`
  - `.oraclepack/runs/<pack_id>/steps.json`
  - `.oraclepack/runs/<pack_id>/outputs/` (20 files; naming convention below)
  - `.oraclepack/runs/<pack_id>/logs/` (optional)

**pack_id (proposed):**
- `YYYY-MM-DD__<gitshort>__<packhash8>`
- If git SHA unavailable: use `nogit` for `<gitshort>`.

**Output naming (proposed):**
- Prefer deterministic: `outputs/01.md` ... `outputs/20.md`
- If a stable QuestionId exists: optionally include in filename, but do not break determinism.

**Resume/rerun (proposed):**
- Resume is default if the run dir already exists:
  - skip steps already marked `ok` with matching output hash.
- `--rerun failed` reruns only failed steps.
- `--rerun all` reruns all steps.
- `--rerun 01,03,07` reruns specified step IDs.

**Concurrency (proposed):**
- `--max-parallel N` bounds parallel provider calls.
- Optional: per-provider caps via config (**TODO/Unknown**: config format).

**Transient errors (proposed):**
- Implement exponential backoff + jitter on retryable errors (e.g., 429/503) up to a retry budget.
- Persist retry counts/outcomes into `steps.json`.

### 3) `oraclepack actionize`

**Goal:** Convert run outputs into actionable engineering work artifacts.

**Proposed:**
- `oraclepack actionize --run-dir .oraclepack/runs/<pack_id> [--ci]`

**Inputs:**
- `run.json`, `steps.json`
- `outputs/` (20 outputs)

**Outputs:**
- `.oraclepack/runs/<pack_id>/actionizer/normalized.jsonl`
- `.oraclepack/runs/<pack_id>/actionizer/backlog.md`
- `.oraclepack/runs/<pack_id>/actionizer/change-plan.md`
- Optional:
  - `.oraclepack/runs/<pack_id>/actionizer/github-issues.json`
  - `.oraclepack/runs/<pack_id>/actionizer/taskmaster.json`

## CI mode (proposed)

- `oraclepack run --ci --non-interactive --json-log`
- `oraclepack actionize --ci`

**Behavior (proposed):**
- No TUI interaction.
- Structured logs enabled (JSONL or JSON objects; **TODO/Unknown** exact format).
- Exit codes are policy-driven:
  - validation failures → non-zero
  - run failures exceeding retry budget → non-zero
  - optional policy thresholds (completion rate, action yield) → **TODO/Unknown** threshold values.

## Security / path safety (proposed)

- Prevent any output flag (including legacy `--write-output` if present) from writing outside the intended run directory.
- Reject path traversal (e.g., `..`) and absolute paths when writing within `.oraclepack/runs/<pack_id>/...`.
