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

# Runtime discovery if out_dir is 'auto' or missing
if [[ "${OUT_DIR}" == "auto" ]] || [[ ! -e "${OUT_DIR}" ]]; then
  echo "INFO: OUT_DIR is '${OUT_DIR}' or missing. Attempting auto-discovery of Stage 2 outputs..."
  found_dir=""
  # Search order:
  # 1. Local oracle-out
  # 2. docs/oracle-out
  # 3. docs/oracle-questions-* (newest)
  candidates=( "oracle-out" "docs/oracle-out" )
  # Add glob matches if they exist
  for d in docs/oracle-questions-*/oracle-out/ docs/oracle-questions-*/; do
    [ -d "$d" ] && candidates+=("$d")
  done
  
  # Check candidates (iterate in order; sort by date descending for docs/oracle-questions)
  # Simple check: take the first one that has 01-*.md content
  for d in "${candidates[@]}"; do
    if [ -d "$d" ]; then
      count=$(ls "$d"/01-*.md 2>/dev/null | wc -l)
      if [ "$count" -gt 0 ]; then
        found_dir="$d"
        echo "INFO: Discovered Stage 2 outputs in: ${found_dir}"
        OUT_DIR="${found_dir}"
        break
      fi
    fi
  done
  
  if [ -z "${found_dir}" ]; then
    echo "ERROR: Could not resolve Stage 2 output directory. Expected 01-*.md in oracle-out/, docs/oracle-out/, or docs/oracle-questions-*/" >&2
    exit 2
  fi
fi

if [ -d "${OUT_DIR}" ]; then
  # Directory mode check
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
elif [ -f "${OUT_DIR}" ]; then
  # Single-pack file mode check
  echo "INFO: Input is a single file: ${OUT_DIR}"
else
  echo "ERROR: Input path '${OUT_DIR}' is neither a directory nor a file." >&2
  exit 5
fi

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

mkdir -p "$(dirname "{{actions_json}}")" "$(dirname "{{actions_md}}")" "$(dirname "{{prd_path}}")" "docs"
# Note: If OUT_DIR is a file, we shouldn't try to mkdir -p it. 
if [ -d "${OUT_DIR}" ]; then mkdir -p "${OUT_DIR}"; fi

cat > "_actions.schema.md" <<'SCHEMA'
# Canonical Actions JSON Schema (human-readable)
[... schema details omitted for brevity, see asset ...]
SCHEMA
# Move schema to a safe place (docs or beside output)
mv "_actions.schema.md" "docs/_actions.schema.md"

echo "OK: Preflight passed."
if [ -d "${OUT_DIR}" ]; then
  echo "OK: Inputs: ${OUT_DIR}/01-*.md .. ${OUT_DIR}/20-*.md"
else
  echo "OK: Input: ${OUT_DIR}"
fi
echo "OK: Mode: ${MODE}"


# 02) Synthesize canonical actions JSON + summary MD
set -euo pipefail

OUT_DIR="{{out_dir}}"
ACTIONS_JSON="{{actions_json}}"
ACTIONS_MD="{{actions_md}}"

# Re-resolve if auto (duplicated for self-containment)
if [[ "${OUT_DIR}" == "auto" ]] || [[ ! -e "${OUT_DIR}" ]]; then
  candidates=( "oracle-out" "docs/oracle-out" )
  for d in docs/oracle-questions-*/oracle-out/ docs/oracle-questions-*/; do [ -d "$d" ] && candidates+=("$d"); done
  for d in "${candidates[@]}"; do
    if [ -d "$d" ] && [ "$(ls "$d"/01-*.md 2>/dev/null | wc -l)" -gt 0 ]; then
      OUT_DIR="$d"
      break
    fi
  done
fi

# Fix derived paths if OUT_DIR changed from auto
if [[ "{{actions_json}}" == *"auto"* ]]; then
  # If actions_json was derived from auto, re-base it on OUT_DIR if OUT_DIR is a dir
  if [ -d "${OUT_DIR}" ]; then
    ACTIONS_JSON="${OUT_DIR}/_actions.json"
    ACTIONS_MD="${OUT_DIR}/_actions.md"
  else
    # If OUT_DIR is a file (pack), put actions next to it or in docs/
    ACTIONS_JSON="docs/_actions.json"
    ACTIONS_MD="docs/_actions.md"
  fi
fi

oracle_file_flags=()
if [ -f "${OUT_DIR}" ]; then
  # Single-pack mode
  oracle_file_flags+=( -f "${OUT_DIR}" )
else
  # Directory mode
  shopt -s nullglob
  for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
    matches=( "${OUT_DIR}/${n}-"*.md )
    if [ "${#matches[@]}" -ne 1 ]; then
      echo "ERROR: expected exactly one match for ${OUT_DIR}/${n}-*.md, got ${#matches[@]}" >&2
      printf '%s\n' "${matches[@]:-}" >&2
      exit 20
    fi
    oracle_file_flags+=( -f "${matches[0]}" )
  done
fi

# Extra attachments (auto-expanded at pack generation time)
{{EXTRA_FILES_LINES}}

mkdir -p "$(dirname "${ACTIONS_JSON}")" "$(dirname "${ACTIONS_MD}")"

{{oracle_cmd}} \
  --write-output "${ACTIONS_JSON}" \
  "${oracle_file_flags[@]}" \
  -p "$(cat <<'PROMPT'
You are producing a SINGLE JSON document and nothing else.

Task: Normalize oraclepack answers into a canonical actionable plan.

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
# Re-resolve if actions_json path is generic
if [[ "{{actions_json}}" == *"auto"* ]]; then
  # Fallback logic mirroring Step 02
  if [ -f "docs/_actions.json" ]; then ACTIONS_JSON="docs/_actions.json"; fi
  # Try to find it in likely places
  found=$(find docs -name "_actions.json" | head -n 1)
  if [ -n "$found" ]; then ACTIONS_JSON="$found"; fi
fi

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
# Re-resolve if auto (light version for report output)
if [[ "${OUT_DIR}" == "auto" ]] || [[ ! -e "${OUT_DIR}" ]]; then
  # Default to docs/ if we can't find source
  OUT_DIR="docs"
fi
if [ -f "${OUT_DIR}" ]; then OUT_DIR="$(dirname "${OUT_DIR}")"; fi

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

  # Ensure dir exists (re-resolve safe fallback)
  if [[ "${OUT_DIR}" == "auto" ]] || [[ ! -d "${OUT_DIR}" ]]; then OUT_DIR="docs"; fi
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