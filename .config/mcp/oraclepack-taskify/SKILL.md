---
name: oraclepack-taskify
description: Generate a Stage-3 Action Pack from oraclepack output (oracle-out 01–20 .md answers) that synthesizes a canonical actions plan + Task Master PRD, runs Task Master to create/expand tasks, and by default starts a guarded autopilot to begin implementation.
metadata:
  short-description: Stage 3:answers → tasks → pipelines/autopilot
---

# oraclepack-taskify

## Use when

Use this skill only after **oraclepack Stage 2** has produced **20 answer files** under an output directory (default `oracle-out/`), and the user wants Stage 3 work products such as:

- “Stage 3” / “taskify” / “actionize” / “turn oracle outputs into tasks”
- “Task Master follow-up” / “PRD from oracle-out” / “implementation plan”
- “Start work automatically from oracle-out” / “autopilot top tasks”

## Deliverable

When invoked, produce exactly one primary deliverable:

- A single Markdown doc at `pack_path` (default `docs/oracle-actions-pack-YYYY-MM-DD.md`)

The generated Action Pack MUST be oraclepack-ingestible and MUST obey:

- Exactly **one** fenced code block labeled `bash` in the entire document.
- No other code fences anywhere in the document.
- Inside the bash fence, include sequential step headers exactly like: `# NN)` where `NN` is `01, 02, 03...`
- Each step is self-contained and must **not rely on shell variables created in previous steps**.
- Each step writes artifacts to explicit, deterministic paths.

## Skill interface (args)

Parse user trailing text as whitespace-separated `KEY=value` tokens (last-one-wins). Do not ask follow-ups.

Supported args (all optional):

- `out_dir` (default `oracle-out`)
- `pack_path` (default `docs/oracle-actions-pack-YYYY-MM-DD.md`)
- `actions_json` (default `<out_dir>/_actions.json`)
- `actions_md` (default `<out_dir>/_actions.md`)
- `prd_path` (default `.taskmaster/docs/oracle-actions-prd.md`)
- `tag` (default `oraclepack`)
- `mode` (default `autopilot`; allowed `backlog|pipelines|autopilot`)
- `top_n` (default `10`; clamp to `1..20`)
- `oracle_cmd` (default `oracle`; allow a multi-word command like `npx -y @steipete/oracle` only if the user requests it)
- `task_master_cmd` (default `task-master`)
- `tm_cmd` (default `tm`; used only in autopilot mode)
- `extra_files` (default empty; if provided, treat as a comma-separated list of file paths; attach them wherever the synthesis step accepts file inputs)

If any referenced file/path does not exist, the skill still outputs the Action Pack, but includes an early step that prints a clear error and exits non-zero.

## Workflow (what to do when invoked)

1) Parse args from `KEY=value` tokens (no follow-ups; last-one-wins; unknown keys ignored).

2) Resolve defaults deterministically:
   - Compute `YYYY-MM-DD` using the local date at generation time.
   - Apply defaults and clamp `top_n` to `1..20`.
   - Default `mode=autopilot` unless explicitly overridden.
   - Derive:
     - `actions_json = <out_dir>/_actions.json` unless user overrides
     - `actions_md = <out_dir>/_actions.md` unless user overrides

3) Render the Action Pack:
   - Start from `assets/action-pack-template.md`.
   - Substitute placeholders:
     - `{{pack_date}}`, `{{out_dir}}`, `{{pack_path}}`, `{{actions_json}}`, `{{actions_md}}`, `{{prd_path}}`, `{{tag}}`, `{{mode}}`, `{{top_n}}`, `{{oracle_cmd}}`, `{{task_master_cmd}}`, `{{tm_cmd}}`
   - Expand `extra_files` into literal bash lines of the form:
     - `oracle_file_flags+=( -f "<path>" )`
     - and insert them only where the template indicates “Extra attachments”.

4) Ensure Action Pack contract:
   - Exactly one `bash` code fence in the document.
   - No other code fences.
   - Step headers `# 01)`.. are sequential and stable.
   - Each step re-declares its variables and does not depend on variables from earlier steps.

5) Write the final pack to `pack_path` (create parent dirs).

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
- Stable ordering: select exactly the 20 outputs by filename prefix ordering `01`..`20`.
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

- Missing `out_dir` → pack Step 01 exits non-zero with a clear message.
- Missing any of `01-*.md`..`20-*.md` → Step 01 exits non-zero (and prints which one is missing).
- Missing required tools → Step 01 exits non-zero (and prints which command is missing).
- Unknown `mode` → treat as `autopilot` (clamp at generation time) and render `mode=autopilot` in the pack.

## Resources

- `assets/action-pack-template.md`: base template for the Action Pack deliverable.
- `assets/actions-json-schema.md`: canonical schema spec for `_actions.json` normalization.
- `assets/prd-synthesis-prompt.md`: exact prompt text embedded into the Action Pack for synthesis.
- `references/*`: Stage 3 overview and guardrails for future maintenance.
- `scripts/*`: optional helpers (standalone); may also be embedded verbatim into packs if desired.

## Invocation examples

- `$oraclepack-taskify out_dir=oracle-out` (defaults to mode=autopilot)
- `$oraclepack-taskify out_dir=oracle-out mode=backlog`
- `$oraclepack-taskify out_dir=oracle-out mode=pipelines tag=oraclepack-top top_n=10`
- `$oraclepack-taskify out_dir=oracle-out mode=autopilot tag=oraclepack-top pack_path=docs/oracle-actions-pack-2026-01-05.md`
- `$oraclepack-taskify out_dir=oracle-out extra_files=README.md,package.json`
