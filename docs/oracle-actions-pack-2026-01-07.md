# Oraclepack Stage 3 — Action Pack (Taskify)

Generated: 2026-01-07

Parsed args (resolved):
- stage2_path: auto
- out_dir: auto
- pack_path: docs/oracle-actions-pack-2026-01-07.md
- actions_json: auto/_actions.json
- actions_md: auto/_actions.md
- prd_path: .taskmaster/docs/oracle-actions-prd.md
- tag: oraclepack
- mode: autopilot
- top_n: 10
- oracle_cmd: oracle
- task_master_cmd: task-master
- tm_cmd: tm
- extra_files: (embedded where applicable)

This document must contain exactly one bash code fence, and no other code fences.

```bash
# 01) Preflight: resolve Stage 2 inputs and verify tools (fail fast)
set -euo pipefail

STAGE2_PATH="auto"
OUT_DIR="auto"
MODE="autopilot"
TASK_MASTER_CMD="task-master"
ORACLE_CMD="oracle"
TM_CMD="tm"
ACTIONS_JSON="auto/_actions.json"
ACTIONS_MD="auto/_actions.md"
PRD_PATH=".taskmaster/docs/oracle-actions-prd.md"

check_dir_complete() {
  local d="$1"
  [ -d "$d" ] || return 1
  local n
  for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
    shopt -s nullglob
    local matches=( "$d/${n}-"*.md )
    shopt -u nullglob
    if [ "${#matches[@]}" -ne 1 ]; then
      return 1
    fi
  done
  return 0
}

resolve_stage2_auto() {
  local candidate

  candidate="oracle-out"
  if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi

  candidate="docs/oracle-out"
  if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi

  for candidate in $(ls -1d docs/oracle-questions-*/oracle-out/ 2>/dev/null | sort -r); do
    if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi
  done

  for candidate in $(ls -1d docs/oracle-questions-*/ 2>/dev/null | sort -r); do
    if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi
  done

  candidate=$(ls -1 docs/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi

  candidate=$(ls -1 docs/oraclepacks/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi

  candidate=$(ls -1 docs/oracle-questions-*/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi

  candidate=$(ls -1 docs/oracle-questions-*/oraclepacks/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi

  return 1
}

resolve_out_dir() {
  local stage2="$1"
  if [ "$OUT_DIR" != "auto" ]; then
    echo "$OUT_DIR"
    return 0
  fi
  if [ -d "$stage2" ]; then
    echo "$stage2"
    return 0
  fi
  if [[ "$stage2" == docs/oracle-questions-*/* ]]; then
    local base
    base=$(echo "$stage2" | awk -F/ '{print $1"/"$2}')
    echo "$base/oracle-out"
    return 0
  fi
  echo "oracle-out"
}

STAGE2_SOURCE=""
if [ "$STAGE2_PATH" != "auto" ]; then
  if [ -d "$STAGE2_PATH" ] || [ -f "$STAGE2_PATH" ]; then
    STAGE2_SOURCE="$STAGE2_PATH"
  else
    echo "ERROR: stage2_path does not exist: $STAGE2_PATH" >&2
    exit 2
  fi
else
  if ! STAGE2_SOURCE=$(resolve_stage2_auto); then
    echo "ERROR: could not resolve Stage 2 outputs via auto search." >&2
    echo "Searched in order: oracle-out/, docs/oracle-out/, docs/oracle-questions-*/oracle-out/, docs/oracle-questions-*/, docs/oracle-pack-*.md, docs/oraclepacks/oracle-pack-*.md, docs/oracle-questions-*/oracle-pack-*.md, docs/oracle-questions-*/oraclepacks/oracle-pack-*.md" >&2
    exit 3
  fi
fi

if [ -d "$STAGE2_SOURCE" ]; then
  for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
    shopt -s nullglob
    matches=( "$STAGE2_SOURCE/${n}-"*.md )
    shopt -u nullglob
    if [ "${#matches[@]}" -eq 0 ]; then
      echo "ERROR: missing oracle output for prefix ${n}: expected ${STAGE2_SOURCE}/${n}-*.md" >&2
      exit 4
    fi
    if [ "${#matches[@]}" -gt 1 ]; then
      echo "ERROR: multiple oracle outputs for prefix ${n}; expected exactly one file." >&2
      printf '%s\n' "${matches[@]}" >&2
      exit 5
    fi
  done
fi

OUT_DIR_RESOLVED=$(resolve_out_dir "$STAGE2_SOURCE")

if [[ "$ACTIONS_JSON" == auto/* ]]; then
  if [ -d "$OUT_DIR_RESOLVED" ]; then
    ACTIONS_JSON="$OUT_DIR_RESOLVED/_actions.json"
    ACTIONS_MD="$OUT_DIR_RESOLVED/_actions.md"
  else
    ACTIONS_JSON="docs/_actions.json"
    ACTIONS_MD="docs/_actions.md"
  fi
fi

if ! command -v "${TASK_MASTER_CMD%% *}" >/dev/null 2>&1; then
  echo "ERROR: required tool missing: ${TASK_MASTER_CMD%% *} (from task_master_cmd='${TASK_MASTER_CMD}')" >&2
  exit 10
fi

if ! command -v "${ORACLE_CMD%% *}" >/dev/null 2>&1; then
  echo "ERROR: required tool missing: ${ORACLE_CMD%% *} (from oracle_cmd='${ORACLE_CMD}')" >&2
  exit 11
fi

if [ "$MODE" = "autopilot" ]; then
  if ! command -v "${TM_CMD%% *}" >/dev/null 2>&1; then
    echo "ERROR: autopilot mode requires tm_cmd, but tool missing: ${TM_CMD%% *} (from tm_cmd='${TM_CMD}')" >&2
    echo "HINT: rerun Stage 3 with mode=backlog if you only want tasks generated." >&2
    exit 12
  fi
fi

mkdir -p "$(dirname "$ACTIONS_JSON")" "$(dirname "$ACTIONS_MD")" "$(dirname "$PRD_PATH")" "docs"
if [ -d "$OUT_DIR_RESOLVED" ]; then mkdir -p "$OUT_DIR_RESOLVED"; fi

cat > "_actions.schema.md" <<'SCHEMA'
# Canonical Actions JSON Schema (human-readable)
[... schema details omitted for brevity, see asset ...]
SCHEMA
mv "_actions.schema.md" "docs/_actions.schema.md"

echo "OK: Stage2 source: ${STAGE2_SOURCE}"
echo "OK: Out dir: ${OUT_DIR_RESOLVED}"
echo "OK: Mode: ${MODE}"


# 02) Synthesize canonical actions JSON + summary MD
set -euo pipefail

STAGE2_PATH="auto"
OUT_DIR="auto"
ACTIONS_JSON="auto/_actions.json"
ACTIONS_MD="auto/_actions.md"

check_dir_complete() {
  local d="$1"
  [ -d "$d" ] || return 1
  local n
  for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
    shopt -s nullglob
    local matches=( "$d/${n}-"*.md )
    shopt -u nullglob
    if [ "${#matches[@]}" -ne 1 ]; then
      return 1
    fi
  done
  return 0
}

resolve_stage2_auto() {
  local candidate
  candidate="oracle-out"
  if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi
  candidate="docs/oracle-out"
  if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi
  for candidate in $(ls -1d docs/oracle-questions-*/oracle-out/ 2>/dev/null | sort -r); do
    if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi
  done
  for candidate in $(ls -1d docs/oracle-questions-*/ 2>/dev/null | sort -r); do
    if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi
  done
  candidate=$(ls -1 docs/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi
  candidate=$(ls -1 docs/oraclepacks/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi
  candidate=$(ls -1 docs/oracle-questions-*/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi
  candidate=$(ls -1 docs/oracle-questions-*/oraclepacks/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi
  return 1
}

resolve_out_dir() {
  local stage2="$1"
  if [ "$OUT_DIR" != "auto" ]; then
    echo "$OUT_DIR"
    return 0
  fi
  if [ -d "$stage2" ]; then
    echo "$stage2"
    return 0
  fi
  if [[ "$stage2" == docs/oracle-questions-*/* ]]; then
    local base
    base=$(echo "$stage2" | awk -F/ '{print $1"/"$2}')
    echo "$base/oracle-out"
    return 0
  fi
  echo "oracle-out"
}

STAGE2_SOURCE=""
if [ "$STAGE2_PATH" != "auto" ]; then
  if [ -d "$STAGE2_PATH" ] || [ -f "$STAGE2_PATH" ]; then
    STAGE2_SOURCE="$STAGE2_PATH"
  else
    echo "ERROR: stage2_path does not exist: $STAGE2_PATH" >&2
    exit 20
  fi
else
  if ! STAGE2_SOURCE=$(resolve_stage2_auto); then
    echo "ERROR: could not resolve Stage 2 outputs via auto search." >&2
    exit 21
  fi
fi

OUT_DIR_RESOLVED=$(resolve_out_dir "$STAGE2_SOURCE")

if [[ "$ACTIONS_JSON" == auto/* ]]; then
  if [ -d "$OUT_DIR_RESOLVED" ]; then
    ACTIONS_JSON="$OUT_DIR_RESOLVED/_actions.json"
    ACTIONS_MD="$OUT_DIR_RESOLVED/_actions.md"
  else
    ACTIONS_JSON="docs/_actions.json"
    ACTIONS_MD="docs/_actions.md"
  fi
fi

oracle_file_flags=()
if [ -f "$STAGE2_SOURCE" ]; then
  oracle_file_flags+=( -f "$STAGE2_SOURCE" )
else
  shopt -s nullglob
  for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
    matches=( "$STAGE2_SOURCE/${n}-"*.md )
    if [ "${#matches[@]}" -ne 1 ]; then
      echo "ERROR: expected exactly one match for ${STAGE2_SOURCE}/${n}-*.md, got ${#matches[@]}" >&2
      printf '%s\n' "${matches[@]:-}" >&2
      exit 22
    fi
    oracle_file_flags+=( -f "${matches[0]}" )
  done
  shopt -u nullglob
fi

mkdir -p "$(dirname "$ACTIONS_JSON")" "$(dirname "$ACTIONS_MD")"

oracle \
  --write-output "$ACTIONS_JSON" \
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
- pack_date: 2026-01-07
- source_out_dir: auto
- top_n: 10
- tag: oraclepack

Output hygiene:
- Do not include backticks or fenced code blocks anywhere.
- Keep strings concise and specific.

Now produce the JSON.
PROMPT
)"

oracle \
  --write-output "$ACTIONS_MD" \
  -f "$ACTIONS_JSON" \
  -p "$(cat <<'PROMPT'
Write a human-readable Markdown summary of the canonical actions JSON.

Hard requirements:
- Output MUST be Markdown text with headings/bullets only (no code fences).
- Keep ordering aligned with items id ascending (01..20).
- Include:
  - short executive summary (5–10 bullets)
  - top 10 prioritized list (with id, title inferred from recommended_next_action, category, and why)
  - per-item: recommended_next_action + acceptance_criteria bullets + missing_artifacts bullets
- Do not invent facts; reflect only what is present in the JSON.

Now write the summary Markdown.
PROMPT
)"

echo "OK: Wrote $ACTIONS_JSON"
echo "OK: Wrote $ACTIONS_MD"


# 03) Generate PRD for Task Master
set -euo pipefail

ACTIONS_JSON="auto/_actions.json"
PRD_PATH=".taskmaster/docs/oracle-actions-prd.md"

if [[ "$ACTIONS_JSON" == auto/* ]]; then
  if [ -f "docs/_actions.json" ]; then
    ACTIONS_JSON="docs/_actions.json"
  else
    found=$(find docs -name "_actions.json" | head -n 1 || true)
    if [ -n "$found" ]; then ACTIONS_JSON="$found"; fi
  fi
fi

mkdir -p "$(dirname "$PRD_PATH")"

oracle \
  --write-output "$PRD_PATH" \
  -f "$ACTIONS_JSON" \
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
- Use the tag value "oraclepack" in the PRD text where helpful for grouping.

Constants:
- TOP_N=10
- TAG=oraclepack

Now produce the PRD.
PROMPT
)"

echo "OK: Wrote $PRD_PATH"


# 04) Task Master: parse PRD into tasks
set -euo pipefail

PRD_PATH=".taskmaster/docs/oracle-actions-prd.md"
TASK_MASTER_CMD="task-master"
TAG="oraclepack"

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

TASK_MASTER_CMD="task-master"
OUT_DIR="auto"

if [[ "$OUT_DIR" == "auto" ]] || [[ ! -e "$OUT_DIR" ]]; then
  OUT_DIR="docs"
fi
if [ -f "$OUT_DIR" ]; then OUT_DIR="$(dirname "$OUT_DIR")"; fi

mkdir -p "$OUT_DIR"

"${TASK_MASTER_CMD}" analyze-complexity --output "$OUT_DIR/tm-complexity.json"
echo "OK: Wrote $OUT_DIR/tm-complexity.json"


# 06) Task Master: expand tasks
set -euo pipefail

TASK_MASTER_CMD="task-master"

"${TASK_MASTER_CMD}" expand --all
echo "OK: Expanded tasks."


# 07) Pipelines (pipelines mode only): generate deterministic pipelines from tasks.json
set -euo pipefail

MODE="autopilot"

if [ "$MODE" != "pipelines" ]; then
  echo "SKIP: mode=$MODE (pipelines step runs only when mode=pipelines)."
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

  oracle \
    --write-output "docs/oracle-actions-pipelines.md" \
    -f "$tasks_path" \
    -p "$(cat <<'PROMPT'
Generate deterministic command pipelines from tasks.json.

Hard requirements:
- Output MUST be Markdown (no code fences).
- Include 3–6 pipelines, each a numbered list of shell commands.
- Each pipeline must be tests-first and avoid destructive operations.
- Commands should be generic and repo-agnostic (no invented scripts).
- Include a short resume strategy section explaining how to re-run pipelines safely.

Now write docs/oracle-actions-pipelines.md content.
PROMPT
)"

  echo "OK: Wrote docs/oracle-actions-pipelines.md"
fi


# 08) Autopilot (autopilot mode only): branch safety + guarded entrypoint
set -euo pipefail

MODE="autopilot"
TM_CMD="tm"
OUT_DIR="auto"
PACK_DATE="2026-01-07"
TAG="oraclepack"

if [ "$MODE" != "autopilot" ]; then
  echo "SKIP: mode=$MODE (autopilot step runs only when mode=autopilot)."
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
  if [ "$current_branch" = "main" ] || [ "$current_branch" = "master" ]; then
    new_branch="oraclepack/${PACK_DATE}-${TAG}"
    echo "INFO: on default-like branch (${current_branch}); creating work branch: ${new_branch}"
    git checkout -b "$new_branch"
  else
    echo "OK: current branch is ${current_branch}"
  fi

  if [[ "$OUT_DIR" == "auto" ]] || [[ ! -d "$OUT_DIR" ]]; then OUT_DIR="docs"; fi
  mkdir -p "$OUT_DIR"

  cat > "$OUT_DIR/tm-autopilot.state.json" <<STATE
{"pack_date":"${PACK_DATE}","tag":"${TAG}","mode":"autopilot","notes":"State file created by Stage-3 Action Pack. Autopilot tooling should resume from this file if supported."}
STATE

  echo "OK: Wrote $OUT_DIR/tm-autopilot.state.json"
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
