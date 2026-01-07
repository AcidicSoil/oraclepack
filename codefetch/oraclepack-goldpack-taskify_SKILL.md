<filetree>
Project Structure:
└── .config
    └── skills
        ├── oraclepack-gold-pack
        │   ├── references
        │   │   ├── attachment-minimization.md
        │   │   ├── inference-first-discovery.md
        │   │   ├── oracle-pack-template.md
        │   │   └── oracle-scratch-format.md
        │   └── SKILL.md
        └── oraclepack-taskify
            ├── assets
            │   ├── action-pack-template.md
            │   ├── actions-json-schema.md
            │   └── prd-synthesis-prompt.md
            ├── references
            │   ├── determinism-and-safety.md
            │   ├── task-master-cli-cheatsheet.md
            │   └── workflow-overview.md
            └── SKILL.md

</filetree>

<source_code>
.config/skills/oraclepack-taskify/SKILL.md
```
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
```

.config/skills/oraclepack-taskify/assets/action-pack-template.md
```
# Oraclepack Stage 3 — Action Pack (Taskify)

Generated: {{pack_date}}

Parsed args (resolved):
- out_dir: {{out_dir}}
- pack_path: {{pack_path}}
- actions_json: {{actions_json}}
- actions_md: {{actions_md}}
- prd_path: {{prd_path}}
- tag: {{tag}}
- mode: {{mode}}
- top_n: {{top_n}}
- oracle_cmd: {{oracle_cmd}}
- task_master_cmd: {{task_master_cmd}}
- tm_cmd: {{tm_cmd}}
- extra_files: (embedded where applicable)

This document must contain exactly one bash code fence, and no other code fences.

```bash
# 01) Preflight: verify inputs and tools (fail fast)
set -euo pipefail

OUT_DIR="{{out_dir}}"
MODE="{{mode}}"
TASK_MASTER_CMD="{{task_master_cmd}}"
ORACLE_CMD="{{oracle_cmd}}"
TM_CMD="{{tm_cmd}}"

if [ ! -d "${OUT_DIR}" ]; then
  echo "ERROR: out_dir does not exist: ${OUT_DIR}" >&2
  exit 2
fi

shopt -s nullglob
for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
  matches=( "${OUT_DIR}/${n}-"*.md )
  if [ "${#matches[@]}" -eq 0 ]; then
    echo "ERROR: missing oracle output for prefix ${n}: expected ${OUT_DIR}/${n}-*.md" >&2
    exit 3
  fi
  if [ "${#matches[@]}" -gt 1 ]; then
    echo "ERROR: multiple oracle outputs for prefix ${n}; expected exactly one file." >&2
    printf '%s\n' "${matches[@]}" >&2
    exit 4
  fi
done

if ! command -v "${TASK_MASTER_CMD%% *}" >/dev/null 2>&1; then
  echo "ERROR: required tool missing: ${TASK_MASTER_CMD%% *} (from task_master_cmd='${TASK_MASTER_CMD}')" >&2
  exit 10
fi

if ! command -v "${ORACLE_CMD%% *}" >/dev/null 2>&1; then
  echo "ERROR: required tool missing: ${ORACLE_CMD%% *} (from oracle_cmd='${ORACLE_CMD}')" >&2
  exit 11
fi

if [ "${MODE}" = "autopilot" ]; then
  if ! command -v "${TM_CMD%% *}" >/dev/null 2>&1; then
    echo "ERROR: autopilot mode requires tm_cmd, but tool missing: ${TM_CMD%% *} (from tm_cmd='${TM_CMD}')" >&2
    echo "HINT: rerun Stage 3 with mode=backlog if you only want tasks generated." >&2
    exit 12
  fi
fi

mkdir -p "$(dirname "{{actions_json}}")" "$(dirname "{{actions_md}}")" "$(dirname "{{prd_path}}")" "docs" "${OUT_DIR}"

cat > "${OUT_DIR}/_actions.schema.md" <<'SCHEMA'
# Canonical Actions JSON Schema (human-readable)

This file describes the required structure of `_actions.json` produced in Step 02.

## Root object

- `metadata` (object, required)
  - `generated_at` (string, ISO-8601 recommended)
  - `pack_date` (string, YYYY-MM-DD)
  - `source_out_dir` (string)
  - `repo` (object)
    - `name` (string, optional)
    - `root` (string, optional)
    - `head_sha` (string, optional)
  - `tooling` (object)
    - `oracle_cmd` (string)
    - `task_master_cmd` (string)
  - `top_n` (number)

- `items` (array, required; max 20)
  - Each item is normalized and must include:

## Item fields (required unless marked optional)

- `id` (string): "01".."20"
- `source_file` (string): the exact answer file path used for this item
- `category` (string): a stable label (e.g., `contracts/interfaces`, `permissions`, `observability`, etc.)
- `priority_score` (number): higher means more important (can be ROI if available)
- `recommended_next_action` (string): a single imperative sentence
- `missing_artifacts` (array of strings): file/path globs to locate or create
- `acceptance_criteria` (array of strings): testable, objective conditions
- `risk_notes` (array of strings): specific risks/unknowns, no generic filler
- `estimated_effort` (string): one of `XS|S|M|L|XL` (or a short consistent scale)

## Optional item fields

- `dependencies` (array of strings): other `id` values that should precede this item

## Normalization rules

- Keep `items` sorted by `id` ascending (01..20), regardless of priority.
- Keep all arrays stably ordered (most important first; ties by lexical order).
- Do not include code fences in any string values.
SCHEMA

cat > "${OUT_DIR}/_prd_synthesis_prompt.md" <<'PROMPT'
See the canonical prompt text in the skill asset: assets/prd-synthesis-prompt.md.

This repo-local copy exists for traceability and to keep the Action Pack portable.
PROMPT

echo "OK: Preflight passed."
echo "OK: Inputs: ${OUT_DIR}/01-*.md .. ${OUT_DIR}/20-*.md"
echo "OK: Mode: ${MODE}"


# 02) Synthesize canonical actions JSON + summary MD
set -euo pipefail

OUT_DIR="{{out_dir}}"
ACTIONS_JSON="{{actions_json}}"
ACTIONS_MD="{{actions_md}}"

shopt -s nullglob
oracle_file_flags=()
for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
  matches=( "${OUT_DIR}/${n}-"*.md )
  if [ "${#matches[@]}" -ne 1 ]; then
    echo "ERROR: expected exactly one match for ${OUT_DIR}/${n}-*.md, got ${#matches[@]}" >&2
    printf '%s\n' "${matches[@]:-}" >&2
    exit 20
  fi
  oracle_file_flags+=( -f "${matches[0]}" )
done

# Extra attachments (auto-expanded at pack generation time)
# (If none, this section is empty)
{{EXTRA_FILES_LINES}}

mkdir -p "$(dirname "${ACTIONS_JSON}")" "$(dirname "${ACTIONS_MD}")"

{{oracle_cmd}} \
  --write-output "${ACTIONS_JSON}" \
  "${oracle_file_flags[@]}" \
  -p "$(cat <<'PROMPT'
You are producing a SINGLE JSON document and nothing else.

Task: Normalize 20 oraclepack answer files into a canonical actionable plan.

Hard requirements:
- Output MUST be valid JSON (no markdown, no prose, no code fences).
- Output MUST follow the schema described below.
- Output MUST be deterministic in ordering:
  - items sorted by id ascending: 01..20
  - arrays use stable ordering (highest priority first; ties lexical)
- Extract only actionable work. Do not invent repo facts. If evidence is missing, record it in missing_artifacts/risk_notes explicitly.

Schema (summarized):
Root object:
- metadata: { generated_at, pack_date, source_out_dir, repo?, tooling?, top_n }
- items: array (max 20)
Each item:
- id: "01".."20"
- source_file: string
- category: string
- priority_score: number
- recommended_next_action: string (single imperative sentence)
- missing_artifacts: string[]
- acceptance_criteria: string[] (testable)
- risk_notes: string[]
- estimated_effort: "XS"|"S"|"M"|"L"|"XL"
- dependencies?: string[] of ids

Repo/run context:
- pack_date: {{pack_date}}
- source_out_dir: {{out_dir}}
- top_n: {{top_n}}
- tag: {{tag}}

Output hygiene:
- Do not include backticks or fenced code blocks anywhere.
- Keep strings concise and specific.

Now produce the JSON.
PROMPT
)"

{{oracle_cmd}} \
  --write-output "${ACTIONS_MD}" \
  -f "${ACTIONS_JSON}" \
  -p "$(cat <<'PROMPT'
Write a human-readable Markdown summary of the canonical actions JSON.

Hard requirements:
- Output MUST be Markdown text with headings/bullets only (no code fences).
- Keep ordering aligned with items id ascending (01..20).
- Include:
  - short executive summary (5–10 bullets)
  - top {{top_n}} prioritized list (with id, title inferred from recommended_next_action, category, and why)
  - per-item: recommended_next_action + acceptance_criteria bullets + missing_artifacts bullets
- Do not invent facts; reflect only what is present in the JSON.

Now write the summary Markdown.
PROMPT
)"

echo "OK: Wrote ${ACTIONS_JSON}"
echo "OK: Wrote ${ACTIONS_MD}"


# 03) Generate PRD for Task Master
set -euo pipefail

ACTIONS_JSON="{{actions_json}}"
PRD_PATH="{{prd_path}}"

mkdir -p "$(dirname "${PRD_PATH}")"

{{oracle_cmd}} \
  --write-output "${PRD_PATH}" \
  -f "${ACTIONS_JSON}" \
  -p "$(cat <<'PROMPT'
Write a Task Master-compatible PRD (Markdown) derived from the canonical actions JSON.

Hard requirements:
- Output MUST be Markdown (no code fences).
- Be dependency-aware (use dependencies if present; otherwise infer minimal dependencies cautiously).
- Prioritize focus: select the top N items by priority_score (N = TOP_N), but keep a traceability appendix mapping all ids 01..20.
- Every selected item must become an implementation-ready PRD section with:
  - Goal
  - Scope
  - Non-goals
  - Constraints
  - Acceptance criteria (testable)
  - Risks/unknowns
  - Dependencies (explicit)
- Use the tag value "{{tag}}" in the PRD text where helpful for grouping.

Constants:
- TOP_N={{top_n}}
- TAG={{tag}}

Now produce the PRD.
PROMPT
)"

echo "OK: Wrote ${PRD_PATH}"


# 04) Task Master: parse PRD into tasks
set -euo pipefail

PRD_PATH="{{prd_path}}"
TASK_MASTER_CMD="{{task_master_cmd}}"
TAG="{{tag}}"

if "${TASK_MASTER_CMD}" parse-prd "${PRD_PATH}" --tag "${TAG}" 2>/dev/null; then
  echo "OK: task-master parse-prd (tagged) succeeded."
else
  echo "INFO: task-master parse-prd did not accept --tag; retrying without tag."
  "${TASK_MASTER_CMD}" parse-prd "${PRD_PATH}"
fi

if [ -f ".taskmaster/tasks.json" ]; then
  echo "OK: Found .taskmaster/tasks.json"
elif [ -f "tasks.json" ]; then
  echo "OK: Found tasks.json"
else
  echo "WARN: tasks.json not found at .taskmaster/tasks.json or tasks.json. Check your Task Master configuration/output path."
fi


# 05) Task Master: analyze complexity and save report
set -euo pipefail

TASK_MASTER_CMD="{{task_master_cmd}}"
OUT_DIR="{{out_dir}}"

mkdir -p "${OUT_DIR}"

"${TASK_MASTER_CMD}" analyze-complexity --output "${OUT_DIR}/tm-complexity.json"
echo "OK: Wrote ${OUT_DIR}/tm-complexity.json"


# 06) Task Master: expand tasks
set -euo pipefail

TASK_MASTER_CMD="{{task_master_cmd}}"

"${TASK_MASTER_CMD}" expand --all
echo "OK: Expanded tasks."


# 07) Pipelines (pipelines mode only): generate deterministic pipelines from tasks.json
set -euo pipefail

MODE="{{mode}}"

if [ "${MODE}" != "pipelines" ]; then
  echo "SKIP: mode=${MODE} (pipelines step runs only when mode=pipelines)."
else
  tasks_path=""
  if [ -f ".taskmaster/tasks.json" ]; then
    tasks_path=".taskmaster/tasks.json"
  elif [ -f "tasks.json" ]; then
    tasks_path="tasks.json"
  else
    echo "ERROR: tasks.json not found; cannot generate pipelines." >&2
    exit 70
  fi

  mkdir -p "docs"

  {{oracle_cmd}} \
    --write-output "docs/oracle-actions-pipelines.md" \
    -f "${tasks_path}" \
    -p "$(cat <<'PROMPT'
Generate deterministic command pipelines from tasks.json.

Hard requirements:
- Output MUST be Markdown (no code fences).
- Include 3–6 pipelines, each a numbered list of shell commands.
- Each pipeline must be tests-first and avoid destructive operations.
- Commands should be generic and repo-agnostic (no invented scripts).
- Include a short “resume strategy” section explaining how to re-run pipelines safely.

Now write docs/oracle-actions-pipelines.md content.
PROMPT
)"

  echo "OK: Wrote docs/oracle-actions-pipelines.md"
fi


# 08) Autopilot (autopilot mode only): branch safety + guarded entrypoint
set -euo pipefail

MODE="{{mode}}"
TM_CMD="{{tm_cmd}}"
OUT_DIR="{{out_dir}}"
PACK_DATE="{{pack_date}}"
TAG="{{tag}}"

if [ "${MODE}" != "autopilot" ]; then
  echo "SKIP: mode=${MODE} (autopilot step runs only when mode=autopilot)."
else
  if ! command -v git >/dev/null 2>&1; then
    echo "ERROR: autopilot mode requires git on PATH." >&2
    exit 80
  fi

  if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    echo "ERROR: not inside a git work tree; autopilot mode requires a git repo." >&2
    exit 81
  fi

  if ! git diff --quiet || ! git diff --cached --quiet; then
    echo "ERROR: working tree not clean. Commit/stash before autopilot." >&2
    exit 82
  fi

  current_branch="$(git rev-parse --abbrev-ref HEAD)"
  if [ "${current_branch}" = "main" ] || [ "${current_branch}" = "master" ]; then
    new_branch="oraclepack/${PACK_DATE}-${TAG}"
    echo "INFO: on default-like branch (${current_branch}); creating work branch: ${new_branch}"
    git checkout -b "${new_branch}"
  else
    echo "OK: current branch is ${current_branch}"
  fi

  mkdir -p "${OUT_DIR}"
  cat > "${OUT_DIR}/tm-autopilot.state.json" <<STATE
{"pack_date":"${PACK_DATE}","tag":"${TAG}","mode":"autopilot","notes":"State file created by Stage-3 Action Pack. Autopilot tooling should resume from this file if supported."}
STATE

  echo "OK: Wrote ${OUT_DIR}/tm-autopilot.state.json"
  echo "INFO: Starting autopilot via: ${TM_CMD} autopilot"
  echo "INFO: If your tm tool uses a different subcommand, edit this step accordingly."

  if ! "${TM_CMD}" --help 2>&1 | grep -qi "autopilot"; then
    echo "ERROR: '${TM_CMD}' does not advertise 'autopilot' in --help output." >&2
    echo "HINT: rerun Stage 3 with mode=backlog if you only want tasks generated." >&2
    exit 83
  fi

  "${TM_CMD}" autopilot
fi
```
```

.config/skills/oraclepack-taskify/assets/actions-json-schema.md
```
# Canonical Actions JSON Schema (human-readable)

This schema defines the required structure of the canonical actions file:

- Default path: `<out_dir>/_actions.json`

## Root object

The root MUST be a JSON object with:

### metadata (required)

- `generated_at` (string): generation timestamp (ISO-8601 recommended)
- `pack_date` (string): `YYYY-MM-DD`
- `source_out_dir` (string): the oraclepack output dir used (e.g., `oracle-out`)
- `repo` (object, optional):
  - `name` (string, optional)
  - `root` (string, optional)
  - `head_sha` (string, optional)
- `tooling` (object, optional):
  - `oracle_cmd` (string)
  - `task_master_cmd` (string)
- `top_n` (number): the top-N focus value used to build PRD/pipelines

### items (required; max 20)

`items` MUST be an array with up to 20 objects. Each item MUST include:

- `id` (string): `"01"`..`"20"`
- `source_file` (string): the specific answer file used for this item
- `category` (string): stable label describing the domain area
- `priority_score` (number): higher means higher priority (can be ROI if present)
- `recommended_next_action` (string): single imperative sentence
- `missing_artifacts` (array of strings): file/path globs or concrete paths
- `acceptance_criteria` (array of strings): testable, objective conditions
- `risk_notes` (array of strings): risks/unknowns grounded in evidence gaps
- `estimated_effort` (string): use a consistent scale such as `XS|S|M|L|XL`

Optional:

- `dependencies` (array of strings): other ids (e.g., `["03","07"]`) that should precede this item

## Normalization rules

- Items MUST be sorted by `id` ascending (`01..20`) for machine stability.
- Within each item:
  - `missing_artifacts`, `acceptance_criteria`, and `risk_notes` MUST be stably ordered.
- Do not include fenced code blocks in any string values.
```

.config/skills/oraclepack-taskify/assets/prd-synthesis-prompt.md
```
# Stage 3 Synthesis Prompts (exact text)

Use these prompts verbatim in the Action Pack.

## Prompt A — Canonical actions JSON (_actions.json)

You are producing a SINGLE JSON document and nothing else.

Task: Normalize 20 oraclepack answer files into a canonical actionable plan.

Hard requirements:
- Output MUST be valid JSON (no markdown, no prose, no code fences).
- Output MUST follow the schema described below.
- Output MUST be deterministic in ordering:
  - items sorted by id ascending: 01..20
  - arrays use stable ordering (highest priority first; ties lexical)
- Extract only actionable work. Do not invent repo facts. If evidence is missing, record it in missing_artifacts/risk_notes explicitly.

Schema (summarized):
Root object:
- metadata: { generated_at, pack_date, source_out_dir, repo?, tooling?, top_n }
- items: array (max 20)
Each item:
- id: "01".."20"
- source_file: string
- category: string
- priority_score: number
- recommended_next_action: string (single imperative sentence)
- missing_artifacts: string[]
- acceptance_criteria: string[] (testable)
- risk_notes: string[]
- estimated_effort: "XS"|"S"|"M"|"L"|"XL"
- dependencies?: string[] of ids

Output hygiene:
- Do not include backticks or fenced code blocks anywhere.
- Keep strings concise and specific.

## Prompt B — Task Master PRD (oracle-actions-prd.md)

Write a Task Master-compatible PRD (Markdown) derived from the canonical actions JSON.

Hard requirements:
- Output MUST be Markdown (no code fences).
- Be dependency-aware (use dependencies if present; otherwise infer minimal dependencies cautiously).
- Prioritize focus: select the top N items by priority_score (N = TOP_N), but keep a traceability appendix mapping all ids 01..20.
- Every selected item must become an implementation-ready PRD section with:
  - Goal
  - Scope
  - Non-goals
  - Constraints
  - Acceptance criteria (testable)
  - Risks/unknowns
  - Dependencies (explicit)

Hygiene:
- Do not invent scripts/paths; use missing_artifacts when you need the repo to supply something.
- Keep acceptance criteria objective and testable.
```

.config/skills/oraclepack-taskify/references/determinism-and-safety.md
```
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
```

.config/skills/oraclepack-taskify/references/task-master-cli-cheatsheet.md
```
# Task Master CLI cheatsheet (minimal)

This skill assumes only these Task Master commands:

## Parse PRD into tasks

- `task-master parse-prd <prd_path>`
- If tag scoping is supported in your setup, the Action Pack attempts:
  - `task-master parse-prd <prd_path> --tag <tag>`
  - and falls back to the untagged command if the flag is not accepted.

## Analyze complexity

- `task-master analyze-complexity --output <out_dir>/tm-complexity.json`

## Expand tasks

- `task-master expand --all`

## Autopilot (default mode behavior)

- The Action Pack attempts: `tm autopilot` (via `tm_cmd`, default `tm`)
- If your `tm` tool does not support autopilot, run Stage 3 with `mode=backlog` or `mode=pipelines`.

Notes:
- The Action Pack checks for `.taskmaster/tasks.json` or `tasks.json` after parsing, but Task Master may be configured differently. If neither file exists, the pack prints a warning.
```

.config/skills/oraclepack-taskify/references/workflow-overview.md
```
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
```

.config/skills/oraclepack-gold-pack/SKILL.md
```
---
name: oraclepack-gold-pack
description: Generate a single canonical Stage-1 oraclepack question pack as Markdown:exactly one ```bash fence containing exactly 20 steps (01..20) with strict ROI header tokens, per-step --write-output, fixed categories + coverage check. Use when you need a gold, runner-ingestible pack template (Stage 1 only; not Stage 3 taskify).
metadata:
  short-description: Gold Stage-1 oraclepack pack generator + validators
---

# Oraclepack Gold Pack (Stage 1)

This skill produces the **canonical Stage-1** oraclepack question pack (20 Oracle CLI calls). It is intentionally strict to prevent schema drift.

**Non-negotiable contract (pack output):**
- Exactly **one** fenced code block: starts with exactly ` ```bash` on its own line, and ends with exactly ` ````
- No other fenced code blocks anywhere in the pack.
- Exactly **20** steps, numbered **01..20** in order.
- Every step has a header line matching:
  - `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes `--write-output "<out_dir>/<nn>-<slug>.md"`.
- Categories are fixed to this exact set (no additions/renames):
  - `contracts/interfaces`
  - `invariants`
  - `caching/state`
  - `background jobs`
  - `observability`
  - `permissions`
  - `migrations`
  - `UX flows`
  - `failure modes`
  - `feature flags`
- Pack ends with a **Coverage check** section listing all 10 categories as `OK` or `Missing(<step ids>)`.

The pack template is the contract. The scratch doc is **not** a pack format.

References:
- Contract template: `references/oracle-pack-template.md`
- Repo discovery: `references/inference-first-discovery.md`
- Attachment rules: `references/attachment-minimization.md`
- Scratch playbook (not pack): `references/oracle-scratch-format.md`

## Quick start

1) Generate a pack file (intended path):
- `docs/oracle-pack-YYYY-MM-DD.md`

1) Validate it (recommended before running oraclepack):
- `python3 scripts/validate_pack.py docs/oracle-pack-YYYY-MM-DD.md`

1) Optional attachment lint:
- `python3 scripts/lint_attachments.py docs/oracle-pack-YYYY-MM-DD.md`

## Inputs (do not ask follow-ups)

Interpret the user’s trailing text as conceptual `{{args}}`. Extract:

- `codebase_name` (default `Unknown`)
- `constraints` (default `None`)
- `non_goals` (default `None`)
- `team_size` (default `Unknown`)
- `deadline` (default `Unknown`)
- `out_dir` (default `docs/oracle-questions-YYYY-MM-DD`)
- `oracle_cmd` (default `oracle`)
- `oracle_flags` (default `--files-report`)
- `engine` (`api|browser`; optional; if provided, append to flags *only if your oracle CLI supports it*, else omit and record `TODO(engine flag unknown)`)
- `model` (optional; if provided, append to flags *only if your oracle CLI supports it*, else omit and record `TODO(model flag unknown)`)
- `extra_files` (default empty; if provided, append **literally** to every command)

If any value is missing: use defaults and proceed.

## Workflow (deterministic)

### 1) Read the contract template first
Open `references/oracle-pack-template.md` and treat it as the **single source of truth** for formatting.

### 2) Repo discovery (inference-first)
Follow `references/inference-first-discovery.md`:
- Read a small set of “anchors” first.
- Infer what’s present in the repo.
- Only then choose the best 1–2 attachments per step.

### 3) Plan the 20 probes (2 per category)
Use the fixed categories and produce **exactly 2 steps per category** (20 total). Keep each step’s prompt focused and non-overlapping.

For each step:
- Pick a **reference anchor** (`reference=` token): `{path}:{symbol}` OR `{path}` OR `Unknown`.
- Pick ≤2 attachments (or fewer) using `references/attachment-minimization.md`.
- Ensure the prompt asks for:
  - Direct answer (bullets)
  - Risks/unknowns
  - Next smallest concrete experiment
  - Missing artifact patterns to request if evidence is insufficient

### 4) Emit the pack (single file)
Produce exactly one Markdown document with:
- Title + parsed args section (plain markdown; no code fences)
- Exactly one ` ```bash` fence containing the 20 steps
- A Coverage check section after the fence

### 5) Validate and correct drift
Run:
- `python3 scripts/validate_pack.py <pack.md>`
If it fails, fix the pack until it passes.

Optionally run:
- `python3 scripts/lint_attachments.py <pack.md>`
If it fails, reduce attachments to ≤2 per step (before any literal `extra_files`).

## Output contract

When invoked, you produce:
- One runner-ingestible Markdown pack (intended filename: `docs/oracle-pack-YYYY-MM-DD.md`)

You do **not**:
- Run oraclepack
- Generate Stage-3 “action packs” (that is `oraclepack-taskify`)

## Failure modes (do not guess)

- Missing repo evidence → set `reference=Unknown`, attach fewer files, and explicitly request missing file/path patterns in the prompt.
- Uncertain CLI flag support (`engine`, `model`) → omit flags and write `TODO(engine/model flag unknown)` in the pack’s parsed args notes (do not invent flags).
- Any schema drift → fix until `scripts/validate_pack.py` passes.

## Invocation examples

1) “Generate a gold oraclepack Stage-1 pack for this repo. out_dir=docs/oracle-questions-2026-01-06”
2) “Make the strict 20-step pack for codebase_name=AcmeAPI constraints=‘no DB changes’”
3) “Create the canonical pack; engine=browser model=gpt-5.2-pro (if supported)”
4) “Produce the gold pack; add extra_files='-f docs/ARCHITECTURE.md -f docs/API.md'”
5) “Regenerate this pack but fix headers and coverage check so it validates.”
```

.config/skills/oraclepack-gold-pack/references/attachment-minimization.md
```
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
```

.config/skills/oraclepack-gold-pack/references/inference-first-discovery.md
```
# Inference-first discovery (Stage 1 packs)

Goal: pick the *right* 1–2 attachments per step without over-attaching, by inferring repo shape from a small set of anchors.

## Principles

- Prefer **evidence** (actual files) over assumptions.
- Start broad with **cheap, high-signal** files; only then zoom in.
- If a file/path doesn’t exist: record `Unknown` and continue.
- Keep steps self-contained: each step should succeed without relying on shared shell setup.

## Deterministic discovery order

1) **Repo identity + entrypoints**
- `README*` (first ~200 lines)
- top-level manifests (language/framework/package)
- main entrypoints (server start, CLI main, app bootstrap) if obvious from tree

2) **Configuration + environment**
- example config files
- `.env.example` or equivalent (if present)
- CI config files (to infer build/test and deploy steps)

3) **Public surface**
- routing tables / controllers / handlers
- schema/contract definitions (API specs, message schemas, type definitions)
- CLI command registration (if applicable)

4) **Data + jobs + operations**
- data models and storage adapters
- migrations directory (if present)
- background job definitions and worker entrypoints (if present)
- logging/metrics/tracing configuration (if present)

## Turning discovery into step-specific attachments

For each planned step:
- Choose 1 “definition” file (where the thing is declared).
- Optionally choose 1 “use-site” file (where the thing is used/enforced).
- If you can’t confidently pick: attach fewer files and use `reference=Unknown`.

Then write the prompt so the oracle can request missing artifact patterns when needed.

## What to do when evidence is insufficient

- Set `reference=Unknown` in the step header.
- In the prompt, explicitly ask for:
  - “the exact missing file/path pattern(s) to attach next”
- Keep attachments minimal; do not guess file paths.
```

.config/skills/oraclepack-gold-pack/references/oracle-pack-template.md
```
# Oracle Pack — {{codebase_name}} (Gold Stage 1)

## Parsed args
- codebase_name: {{codebase_name}}
- constraints: {{constraints}}
- non_goals: {{non_goals}}
- team_size: {{team_size}}
- deadline: {{deadline}}
- out_dir: {{out_dir}}
- oracle_cmd: {{oracle_cmd}}
- oracle_flags: {{oracle_flags}}
- engine: {{engine}}
- model: {{model}}
- extra_files: {{extra_files}}

Notes:
- Template is the **contract**. Keep the pack runner-ingestible.
- Exactly one fenced `bash` block in this whole document.
- Exactly 20 steps, numbered `01..20`.
- Each step includes: `ROI= impact= confidence= effort= horizon= category= reference=`
- Categories must be exactly the fixed set used in Coverage check.

## Commands
```bash
# Optional preflight pattern:
# - Add `--dry-run summary` to preview what will be sent, and keep `--files-report` enabled when available.

# 01) ROI=... impact=... confidence=... effort=... horizon=Immediate category=contracts/interfaces reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/01-contracts-interfaces-surface.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #01

Reference: Unknown
Category: contracts/interfaces
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Identify the primary public interface(s) of this system (API endpoints, CLI commands, public SDK surface, event contracts). For each, list the key request/response (or input/output) shapes and where they are defined in the code.

Rationale (one sentence):
We need a trustworthy map of the system’s “outside-facing contract” before deeper planning.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=... impact=... confidence=... effort=... horizon=Immediate category=contracts/interfaces reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/02-contracts-interfaces-integration.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #02

Reference: Unknown
Category: contracts/interfaces
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What are the top integration points with external systems (databases, queues, third-party APIs, SSO, storage), and what contract(s) or config declare them? Provide the minimal list of files/locations that define each integration.

Rationale (one sentence):
Integration boundaries drive risk, deployment needs, and test strategy.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=invariants reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/03-invariants-domain.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #03

Reference: Unknown
Category: invariants
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
List the system’s most important invariants (business rules, correctness properties, “must always be true” conditions). For each, show where it is enforced (or where it should be enforced but currently is not).

Rationale (one sentence):
Invariants define correctness and are the backbone of reliable changes.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=invariants reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/04-invariants-validation.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #04

Reference: Unknown
Category: invariants
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Where does validation happen (input validation, schema validation, domain validation)? Identify the validation boundaries and the most likely gaps that could cause inconsistent state.

Rationale (one sentence):
Knowing validation boundaries prevents regressions and reduces security/correctness risk.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=caching/state reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/05-caching-state-layers.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #05

Reference: Unknown
Category: caching/state
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What stateful components exist (in-memory state, caches, sessions, client-side state, persisted state)? For each, describe lifecycle, invalidation/expiry, and where it is implemented.

Rationale (one sentence):
State and caching are common sources of subtle bugs and performance issues.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=caching/state reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/06-caching-state-consistency.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #06

Reference: Unknown
Category: caching/state
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Identify the top consistency risks between caches/state layers and the source of truth (stale reads, write skew, missing invalidation). Where are the knobs/configs for cache behavior?

Rationale (one sentence):
Consistency failure modes often surface as “random bugs” and are expensive to debug.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 07) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=background jobs reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/07-background-jobs-discovery.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #07

Reference: Unknown
Category: background jobs
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What background jobs/workers/scheduled tasks exist? For each, identify trigger mechanism, payload, retries, idempotency, and where it is defined.

Rationale (one sentence):
Background work affects reliability, cost, and operational complexity.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 08) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=background jobs reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/08-background-jobs-reliability.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #08

Reference: Unknown
Category: background jobs
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Where are the main reliability controls for background work (dead-lettering, backoff, concurrency limits, reprocessing), and what is missing or inconsistent?

Rationale (one sentence):
Reliability controls prevent incident loops and data corruption.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 09) ROI=... impact=... confidence=... effort=... horizon=Immediate category=observability reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/09-observability-signals.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #09

Reference: Unknown
Category: observability
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What observability signals exist (logs/metrics/traces/events), and what are the primary identifiers for correlating a request/job across components? Point to the code/config that defines them.

Rationale (one sentence):
You can’t operate or improve what you can’t measure or debug quickly.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 10) ROI=... impact=... confidence=... effort=... horizon=Immediate category=observability reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/10-observability-gaps.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #10

Reference: Unknown
Category: observability
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Where are the biggest observability gaps (missing logs around key decisions, missing metrics for SLOs, missing trace spans)? Recommend the smallest additions that would most improve debugging.

Rationale (one sentence):
Targeted observability improvements compound across all future changes.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 11) ROI=... impact=... confidence=... effort=... horizon=Immediate category=permissions reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/11-permissions-model.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #11

Reference: Unknown
Category: permissions
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What is the permission model (roles/scopes/claims/ACLs), and where is it defined? Provide the minimal set of files that encode “who can do what.”

Rationale (one sentence):
Permission rules are a high-risk area with security and product impact.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 12) ROI=... impact=... confidence=... effort=... horizon=Immediate category=permissions reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/12-permissions-enforcement.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #12

Reference: Unknown
Category: permissions
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Where are permissions enforced (middleware/guards/policies/service-layer checks), and where are likely bypass risks? Identify the enforcement chokepoints and any inconsistent patterns.

Rationale (one sentence):
Enforcement consistency prevents privilege escalation and policy drift.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 13) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=migrations reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/13-migrations-schema.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #13

Reference: Unknown
Category: migrations
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
How are schema/config migrations handled (DB migrations, data backfills, versioned configs)? Identify the tooling, directories, and how migrations are applied in CI/deploy.

Rationale (one sentence):
Migration mechanics are critical for safe releases and rollbacks.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 14) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=migrations reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/14-migrations-compat.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #14

Reference: Unknown
Category: migrations
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What are the backward/forward compatibility expectations during migrations (rolling deploys, dual-read/dual-write, feature-flagged schema use)? Identify where compatibility is ensured or currently risky.

Rationale (one sentence):
Compatibility strategy prevents outages during deployments.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 15) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=UX flows reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/15-ux-flows-primary.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #15

Reference: Unknown
Category: UX flows
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What are the primary user flows (or primary operator workflows) and their steps? Map each to the main components/modules involved, and note the key state transitions.

Rationale (one sentence):
Flow maps reveal critical paths and help prioritize work with user impact.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 16) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=UX flows reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/16-ux-flows-edgecases.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #16

Reference: Unknown
Category: UX flows
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
For the primary flows, what are the top edge cases and “gotchas” (validation failures, partial completion, retries, timeouts)? Identify where these cases are handled and where they are missing.

Rationale (one sentence):
Edge-case handling is where many UX and reliability issues originate.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 17) ROI=... impact=... confidence=... effort=... horizon=Immediate category=failure modes reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/17-failure-modes-taxonomy.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #17

Reference: Unknown
Category: failure modes
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What is the failure-mode taxonomy of this system (timeouts, retries, partial failures, inconsistent state, dependency failures)? Identify where failures are classified/handled and what is surfaced to users/operators.

Rationale (one sentence):
Explicit failure handling prevents incidents and reduces user-facing errors.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 18) ROI=... impact=... confidence=... effort=... horizon=Immediate category=failure modes reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/18-failure-modes-resilience.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #18

Reference: Unknown
Category: failure modes
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What resilience mechanisms exist (circuit breakers, bulkheads, retries with jitter, rate limiting, graceful degradation)? Where are they configured, and which critical path lacks them?

Rationale (one sentence):
Resilience patterns determine real-world reliability under stress.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 19) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=feature flags reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/19-feature-flags-inventory.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #19

Reference: Unknown
Category: feature flags
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What feature-flag system exists (or how are conditional rollouts handled)? Inventory the flags (or equivalents) and identify where flags are defined, evaluated, and documented.

Rationale (one sentence):
Flags enable safe rollout and experimentation and reduce release risk.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 20) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=feature flags reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/20-feature-flags-rollout.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #20

Reference: Unknown
Category: feature flags
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Describe the flag/rollout lifecycle (how flags are created, tested, ramped, monitored, and retired). Identify the minimum guardrails needed to prevent “flag debt.”

Rationale (one sentence):
A disciplined rollout lifecycle reduces long-term complexity and operational risk.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"
```

Coverage check
--------------

Mark each as `OK` or `Missing(<which step ids>)`:

*   contracts/interfaces: OK

*   invariants: OK

*   caching/state: OK

*   background jobs: OK

*   observability: OK

*   permissions: OK

*   migrations: OK

*   UX flows: OK

*   failure modes: OK

*   feature flags: OK
```

.config/skills/oraclepack-gold-pack/references/oracle-scratch-format.md
```
# Oracle scratch playbook (NOT a pack format)

This document is for **manual debugging / one-off oracle runs**. It is **not** runner-ingestible oraclepack pack format.

## When to use scratch vs pack

Use the **pack** (`references/oracle-pack-template.md`) when:
- You need a strict 20-step Stage-1 pack for oraclepack ingestion.
- You want deterministic execution and validation via `scripts/validate_pack.py`.

Use **scratch** when:
- You need a single oracle call to explore something quickly.
- You are iterating on prompt wording before committing it into the pack.
- You want to test attachment choices with `--dry-run`.

## Scratch workflow

1) Start with one focused question.
2) Attach 0–2 high-signal files.
3) Use a quoted heredoc prompt to avoid shell-escaping issues.
4) If results are weak, add *one* more attachment (or refine the question).

## Pack-adjacent scratch example (single run)

Example pattern (edit paths/flags to match your environment):

- Uses the quoted heredoc prompt style.
- Shows the optional `--dry-run summary` style (if supported).

```bash
# (optional) preview:
# oracle --dry-run summary --files-report -f "README.md" -p "$(cat <<'PROMPT'
# Explain the repo’s main entrypoint and how requests flow through the system.
# If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
# PROMPT
# )"

oracle \
  --files-report \
  -f "README.md" \
  -p "$(cat <<'PROMPT'
Goal: Understand the repo’s main entrypoints and primary request flow.

Answer format:
1) Direct answer (bullets, evidence-cited)
2) Risks/unknowns
3) Next smallest concrete experiment (one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"
```

## Promoting scratch into the pack

When a scratch run looks good:
- Convert it into a numbered step in the pack.
- Add the strict header tokens (`ROI= impact= confidence= effort= horizon= category= reference=`).
- Add `--write-output "{{out_dir}}/NN-<slug>.md"`.
- Ensure category is one of the fixed 10 and update Coverage check accordingly.

```
```

</source_code>