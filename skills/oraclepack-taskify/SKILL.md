---
name: oraclepack-taskify
description: Generate a Stage-3 Action Pack from oraclepack output (Stage-2 answers) that synthesizes a canonical actions plan + Task Master PRD, runs Task Master to create/expand tasks, and by default starts a guarded autopilot to begin implementation.
metadata:
  short-description: Stage 3:answers → tasks → pipelines/autopilot
---

# oraclepack-taskify

## Use when

Use this skill only after **oraclepack Stage 2** has produced its outputs, and the user wants Stage 3 work products such as:

- “Stage 3” / “taskify” / “actionize” / “turn oracle outputs into tasks”
- “Task Master follow-up” / “PRD from Stage 2 output” / “implementation plan”
- “Start work automatically from oracle outputs” / “autopilot top tasks”

Stage 2 outputs may be in **either** form:

1) **Directory form (legacy):** a folder containing **20 answer files** named `01-*.md` … `20-*.md`
   - Common locations:
     - `oracle-out/` (legacy default)
     - `docs/**/oracle-out/` (newer oraclepack variants)
     - `docs/oracle-questions-YYYY-MM-DD/**` (date-stamped runs)

2) **Single-pack form (newer):** a consolidated Stage-2 markdown pack file (e.g. `oracle-pack-YYYY-MM-DD.md`)
   - Common locations:
     - `docs/oracle-pack-YYYY-MM-DD.md` (default gold-pack location)
     - `docs/oraclepacks/oracle-pack-YYYY-MM-DD.md`
     - `docs/oracle-questions-YYYY-MM-DD/oraclepacks/oracle-pack-YYYY-MM-DD.md`

If you know where Stage 2 landed, pass `stage2_path=<dir-or-file>`.
If you don’t, omit it (or set `stage2_path=auto`) and the generated Action Pack must resolve it deterministically at runtime using the ordered search rules defined in this skill.

## Deliverable

When invoked, produce exactly one primary deliverable:

- A single Markdown doc at `pack_path` (default `docs/oracle-actions-pack-YYYY-MM-DD-HHMMSS.md`)

The generated Action Pack MUST be oraclepack-ingestible and MUST obey:

- Exactly **one** fenced code block labeled `bash` in the entire document.
- No other code fences anywhere in the document.
- Inside the bash fence, include sequential step headers exactly like: `# NN)` where `NN` is `01, 02, 03...`
- Each step is self-contained and must **not rely on shell variables created in previous steps**.
- Each step writes artifacts to explicit, deterministic paths.

## Skill interface (args)

Parse user trailing text as whitespace-separated `KEY=value` tokens (last-one-wins). Do not ask follow-ups.

Supported args (all optional):

### Stage-2 input selection

- `stage2_path` (default `auto`)
  - May be either:
    - A directory containing `01-*.md`..`20-*.md`, or
    - A single consolidated Stage-2 pack markdown file (e.g. `docs/oraclepacks/oracle-pack-YYYY-MM-DD.md` or `docs/oracle-pack-*.md`)
  - If `auto`, the Action Pack MUST resolve a Stage-2 source at runtime using the ordered resolver rules in this skill.

### Stage-3 output locations

- `out_dir` (default `auto`)
  - Directory where Stage-3 artifacts are written (`_actions.json`, `_actions.md`, Task Master outputs).
  - If `auto`:
    - If Stage-2 source is a **directory**, set `out_dir=<that directory>`.
    - If Stage-2 source is a **file**, set:
      - `out_dir=<parent docs run folder>/oracle-out` when the file is under a `docs/oracle-questions-YYYY-MM-DD/` run folder, else
      - `out_dir=oracle-out`.

### Primary deliverable path

- `pack_path` (default `docs/oracle-actions-pack-YYYY-MM-DD-HHMMSS.md`)

### Derived artifact paths (based on out_dir unless overridden)

- `actions_json` (default `<out_dir>/_actions.json`)
- `actions_md` (default `<out_dir>/_actions.md`)
- `prd_path` (default `.taskmaster/docs/oracle-actions-prd.md`)

### Behavior / tagging

- `tag` (default `oraclepack`)
- `mode` (default `autopilot`; allowed `backlog|pipelines|autopilot`)
- `top_n` (default `10`; clamp to `1..20`)

### Tool commands

- `oracle_cmd` (default `oracle`; allow a multi-word command like `npx -y @steipete/oracle` only if the user requests it)
- `task_master_cmd` (default `task-master`)
- `tm_cmd` (default `tm`; used only in autopilot mode)

### Optional attachments

- `extra_files` (default empty; if provided, treat as a comma-separated list of file paths; attach them wherever the synthesis step accepts file inputs)

If any referenced file/path does not exist, the skill still outputs the Action Pack, but includes an early step that prints a clear error and exits non-zero.


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

## Workflow (what to do when invoked)

1) Parse args from `KEY=value` tokens (no follow-ups; last-one-wins; unknown keys ignored).

If `2`, collect overrides and apply them; otherwise proceed with defaults.

2) Resolve defaults deterministically:
   - Compute `YYYY-MM-DD-HHMMSS` using the local date/time at generation time.
   - Clamp `top_n` to `1..20`.
   - Default `mode=autopilot` unless explicitly overridden.
   - Apply `stage2_path` and `out_dir` defaults to `auto` if not provided.
   - Derive:
     - `actions_json = <out_dir>/_actions.json` unless user overrides
     - `actions_md = <out_dir>/_actions.md` unless user overrides

3) Render the Action Pack:
   - Start from `assets/action-pack-template.md`.
   - Substitute placeholders:
     - `{{pack_date}}`, `{{pack_path}}`, `{{tag}}`, `{{mode}}`, `{{top_n}}`, `{{oracle_cmd}}`, `{{task_master_cmd}}`, `{{tm_cmd}}`
     - `{{stage2_path}}`, `{{out_dir}}`, `{{actions_json}}`, `{{actions_md}}`, `{{prd_path}}`
   - Expand `extra_files` into literal bash lines of the form:
     - `oracle_file_flags+=( -f "<path>" )`
     - and insert them only where the template indicates “Extra attachments”.
   - IMPORTANT: The generated Action Pack MUST implement Stage-2 resolution at runtime (Step 01) so it works when Stage 2 outputs are under `docs/` instead of `oracle-out/`.

4) Ensure Action Pack contract:
   - Exactly one `bash` code fence in the document.
   - No other code fences.
   - Step headers `# 01)`.. are sequential and stable.
   - Each step re-declares its variables and does not depend on variables from earlier steps.

5) Write the final pack to `pack_path` (create parent dirs).

6) Emit a minimal manifest (recommended):
   - `python3 scripts/write_manifest.py --pack-path <pack_path> --out-dir <out_dir> --actions-json <actions_json> --actions-md <actions_md> --prd-path <prd_path>`

## Stage-2 resolver rules (Action Pack MUST implement in Step 01)

The Action Pack must resolve a Stage-2 source **deterministically**.

Inputs:

- `stage2_path` (may be `auto`)
- `out_dir` (may be `auto`)

Resolution:

1) If `stage2_path != auto`:
   - If it’s a directory: treat as **directory form** input.
   - If it’s a file: treat as **single-pack form** input.
   - Else: fail fast with a clear error.

2) If `stage2_path == auto`, try in this exact order (first match wins):
   - Directory form:
     1. `oracle-out/` if it contains exactly one match for each of `01-*.md`..`20-*.md`
     2. `docs/oracle-out/` with the same rule
     3. Newest (lexicographically greatest) `docs/oracle-questions-*/oracle-out/` that satisfies the same rule
     4. Newest (lexicographically greatest) `docs/oracle-questions-*/` that directly contains the required `01-*.md`..`20-*.md`
   - Single-pack form:
     5. Newest (lexicographically greatest) `docs/oracle-pack-*.md`
     6. Newest (lexicographically greatest) `docs/oraclepacks/oracle-pack-*.md`
     7. Newest (lexicographically greatest) `docs/oracle-questions-*/oracle-pack-*.md`
     8. Newest (lexicographically greatest) `docs/oracle-questions-*/oraclepacks/oracle-pack-*.md`

3) If `out_dir == auto`:
   - If Stage-2 source is a directory: `out_dir=<stage2 directory>`
   - If Stage-2 source is a file:
     - If the file path matches `docs/oracle-questions-YYYY-MM-DD/**`: set `out_dir` to `docs/oracle-questions-YYYY-MM-DD/oracle-out`
     - Else: set `out_dir=oracle-out`

Ingestion behavior:

- If Stage-2 source is **directory form**, the pack attaches exactly the 20 files `01-*.md`..`20-*.md` (and errors on missing or ambiguous matches).
- If Stage-2 source is **single-pack form**, the pack attaches exactly that one file as the Stage-2 input (and MUST NOT try to glob `01-*.md`..`20-*.md`).

## Modes

### mode=backlog

Action Pack should:

- Synthesize canonical `_actions.json` + `_actions.md`
- Write PRD to `prd_path`
- Run:
  - `task-master parse-prd <prd_path>` (attempt with tag scoping if supported)
  - `task-master analyze-complexity --output <out_dir>/tm-complexity.json`
  - `task-master expand --all`

### mode=pipelines

Do everything in `backlog`, then:

- Generate deterministic pipelines from tasks.json
- Write: `docs/oracle-actions-pipelines.md`

### mode=autopilot (default)

Do everything in `backlog`, then:

- Enforce branch safety and tests-first guardrails
- Start a guarded autopilot entrypoint (expects `tm`-style autopilot tooling)
- Never commit to the default branch (do not run `git commit` on main/master; create a work branch first)

Important: if the environment does not provide a compatible `tm autopilot`, Step 08 will fail fast with a clear error. To avoid that, run Stage 3 with `mode=backlog` or `mode=pipelines`.

## Determinism + safety rules

- No interactive prompts in the generated pack.
- Stage-2 auto-discovery MUST be deterministic:
  - Only use lexicographic sorting over ISO-date directory/file names when selecting “newest”.
  - Do not use `mtime`/timestamps for selection.
- Stable ordering:
  - Directory form: select exactly the 20 outputs by filename prefix ordering `01`..`20`.
  - Single-pack form: select exactly one pack file.
- Fail fast when required tools are missing:
  - `command -v <task_master_cmd first word>`
  - `command -v <oracle_cmd first word>`
  - `command -v <tm_cmd first word>` only in autopilot mode
- Always create directories before writing files.
- Avoid destructive commands:
  - Do not delete files.
  - Do not force-push.
  - Do not commit to main/master (autopilot mode creates a new branch).
- If multiple outputs exist for a prefix (e.g., `01-*.md` expands to more than one file), exit non-zero with an explicit error listing the matches.

## Failure modes (handle without questions)

- Stage-2 cannot be resolved (`stage2_path=auto` finds nothing) → Step 01 exits non-zero with a clear message listing searched locations.
- Stage-2 directory form:
  - Missing `01-*.md`..`20-*.md` → Step 01 exits non-zero (prints which is missing).
  - Ambiguous match for any prefix → Step 01 exits non-zero (prints matches).
- Stage-2 single-pack form:
  - Missing file → Step 01 exits non-zero (prints expected path).
- Missing required tools → Step 01 exits non-zero (prints which command is missing).
- Unknown `mode` → treat as `autopilot` (clamp at generation time) and render `mode=autopilot` in the pack.

## Resources

- `assets/action-pack-template.md`: base template for the Action Pack deliverable.
- `assets/actions-json-schema.md`: canonical schema spec for `_actions.json` normalization.
- `assets/prd-synthesis-prompt.md`: exact prompt text embedded into the Action Pack for synthesis.
- `references/*`: Stage 3 overview and guardrails for future maintenance.
- `scripts/*`: optional helpers (standalone); may also be embedded verbatim into packs if desired.

## Invocation examples

Legacy directory-form (Stage-2 answers in `oracle-out/`):

- `$oraclepack-taskify out_dir=oracle-out` (defaults to mode=autopilot)
- `$oraclepack-taskify out_dir=oracle-out mode=backlog`
- `$oraclepack-taskify out_dir=oracle-out mode=pipelines tag=oraclepack-top top_n=10`

Directory-form under docs (explicit):

- `$oraclepack-taskify stage2_path=docs/oracle-questions-2026-01-07/oracle-out out_dir=auto`
- `$oraclepack-taskify stage2_path=docs/oracle-out out_dir=docs/oracle-out mode=backlog`

Single-pack form (explicit):

- `$oraclepack-taskify stage2_path=docs/oraclepacks/oracle-pack-2026-01-07.md out_dir=docs/oracle-out`
- `$oraclepack-taskify stage2_path=docs/oracle-questions-2026-01-07/oraclepacks/oracle-pack-2026-01-07.md out_dir=auto`

Let the pack resolve Stage-2 automatically:

- `$oraclepack-taskify stage2_path=auto out_dir=auto`

Attachments:

- `$oraclepack-taskify stage2_path=auto extra_files=README.md,package.json`
