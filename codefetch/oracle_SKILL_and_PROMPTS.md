<filetree>
Project Structure:
├── .config
│   ├── commands
│   │   └── oracle-pack_v2.toml
│   ├── completion
│   │   └── oraclepack.completion.sh
│   ├── mcp
│   │   ├── oraclepack-gold-pack
│   │   │   ├── references
│   │   │   │   ├── attachment-minimization.md
│   │   │   │   ├── inference-first-discovery.md
│   │   │   │   ├── oracle-pack-template.md
│   │   │   │   └── oracle-scratch-format.md
│   │   │   ├── scripts
│   │   │   │   ├── lint_attachments.py
│   │   │   │   └── validate_pack.py
│   │   │   └── SKILL.md
│   │   └── oraclepack-taskify
│   │       ├── assets
│   │       │   ├── action-pack-template.md
│   │       │   ├── actions-json-schema.md
│   │       │   └── prd-synthesis-prompt.md
│   │       ├── references
│   │       │   ├── determinism-and-safety.md
│   │       │   ├── task-master-cli-cheatsheet.md
│   │       │   └── workflow-overview.md
│   │       ├── scripts
│   │       │   ├── detect-oracle-outputs.sh
│   │       │   └── validate-action-pack.sh
│   │       └── SKILL.md
│   └── skills
│       ├── oracle-pack
│       │   ├── assets
│       │   │   └── oracle-pack-template.md
│       │   ├── references
│       │   │   ├── attachment-minimization.md
│       │   │   ├── inference-first-discovery.md
│       │   │   └── oracle-scratch-format.md
│       │   └── SKILL.md
│       ├── oracle-pack-rpg-infused
│       │   ├── assets
│       │   │   └── oracle-pack-template.md
│       │   ├── references
│       │   │   └── oracle-pack-spec.md
│       │   ├── scripts
│       │   │   └── validate_oracle_pack.py
│       │   └── SKILL.md
│       ├── oracle-strategist-commands
│       │   └── SKILL.md
│       ├── oraclepack-gold-pack
│       │   ├── references
│       │   │   ├── attachment-minimization.md
│       │   │   ├── inference-first-discovery.md
│       │   │   ├── oracle-pack-template.md
│       │   │   └── oracle-scratch-format.md
│       │   ├── scripts
│       │   │   ├── lint_attachments.py
│       │   │   └── validate_pack.py
│       │   └── SKILL.md
│       ├── oraclepack-taskify
│       │   ├── assets
│       │   │   ├── action-pack-template.md
│       │   │   ├── actions-json-schema.md
│       │   │   └── prd-synthesis-prompt.md
│       │   ├── references
│       │   │   ├── determinism-and-safety.md
│       │   │   ├── task-master-cli-cheatsheet.md
│       │   │   └── workflow-overview.md
│       │   ├── scripts
│       │   │   ├── detect-oracle-outputs.sh
│       │   │   └── validate-action-pack.sh
│       │   └── SKILL.md
│       ├── oraclepack-usecase-rpg
│       │   ├── references
│       │   │   └── oracle-pack-template.md
│       │   ├── scripts
│       │   │   └── validate_oraclepack_pack.py
│       │   └── SKILL.md
│       └── strategist-questions-oracle-pack
│           ├── assets
│           │   └── oracle-pack-template.md
│           └── references
│               └── oracle-scratch-format.md
├── .tickets
│   └── mcp
│       ├── oraclepack-MCP.md
│       └── oraclepack_mcp_server.md
├── docs
│   ├── oracle-questions-2026-01-07
│   │   ├── 01-contracts-interfaces-surface.md
│   │   ├── 02-contracts-interfaces-integration.md
│   │   ├── 03-invariants-domain.md
│   │   ├── 04-invariants-validation.md
│   │   ├── 05-caching-state-layers.md
│   │   ├── 06-caching-state-consistency.md
│   │   ├── 07-background-jobs-discovery.md
│   │   ├── 08-background-jobs-reliability.md
│   │   ├── 09-observability-signals.md
│   │   ├── 10-observability-gaps.md
│   │   ├── 11-permissions-model.md
│   │   ├── 12-permissions-enforcement.md
│   │   ├── 13-migrations-schema.md
│   │   ├── 14-migrations-compat.md
│   │   ├── 15-ux-flows-primary.md
│   │   ├── 16-ux-flows-edgecases.md
│   │   ├── 17-failure-modes-taxonomy.md
│   │   ├── 18-failure-modes-resilience.md
│   │   ├── 19-feature-flags-inventory.md
│   │   └── 20-feature-flags-rollout.md
│   ├── oracle-actions-pack-2026-01-07.md
│   └── oracle-pack-2026-01-07.md
├── internal
│   └── exec
│       ├── oracle_scan.go
│       ├── oracle_scan_test.go
│       ├── oracle_validate.go
│       └── oracle_validate_test.go
├── oracle-pack-2026-01-02.chatgpt-urls.json
└── oracle-pack-2026-01-02.state.json

</filetree>

<source_code>
oracle-pack-2026-01-02.chatgpt-urls.json
```
{
  "default": "",
  "items": [
    {
      "name": "oracle",
      "url": "https://chatgpt.com/g/g-p-694ed925bab08191acaf80aefaf27dfc-oracle/project",
      "lastUsed": "2026-01-02T22:40:05Z"
    }
  ]
}
```

oracle-pack-2026-01-02.state.json
```
{
  "schema_version": 1,
  "pack_hash": "",
  "start_time": "0001-01-01T00:00:00Z",
  "step_statuses": {},
  "roi_threshold": 1.6,
  "roi_mode": "over"
}
```

docs/oracle-actions-pack-2026-01-07.md
```
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
[TRUNCATED]
```

docs/oracle-pack-2026-01-07.md
```
# Oracle Pack — Unknown (Gold Stage 1)

## Parsed args
- codebase_name: Unknown
- constraints: None
- non_goals: None
- team_size: Unknown
- deadline: Unknown
- out_dir: docs/oracle-questions-2026-01-07
- oracle_cmd: oracle
- oracle_flags: --files-report
- engine: None
- model: None
- extra_files: None

Notes:
- Template is the contract. Keep the pack runner-ingestible.
- Exactly one fenced bash block in this whole document.
- Exactly 20 steps, numbered 01..20.
- Each step includes: ROI= impact= confidence= effort= horizon= category= reference=
- Categories must be exactly the fixed set used in Coverage check.

## Commands
```bash
# Optional preflight pattern:
# - Add `--dry-run summary` to preview what will be sent, and keep `--files-report` enabled when available.

# 01) ROI=8.2 impact=9 confidence=0.9 effort=1 horizon=Immediate category=contracts/interfaces reference=internal/cli/root.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/01-contracts-interfaces-surface.md" \
  --file "internal/cli/root.go" \
  --file "internal/cli/cmds.go" \
  -p "$(cat <<'PROMPT'
Strategist question #01

Reference: internal/cli/root.go
Category: contracts/interfaces
Horizon: Immediate
ROI: 8.2 (impact=9, confidence=0.9, effort=1)

Question:
Identify the primary public interface(s) of this system (CLI commands and the oracle-pack Markdown contract). For each, list the key inputs/outputs and where they are defined in code.

Rationale (one sentence):
We need a trustworthy map of the system’s “outside-facing contract” before deeper planning.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=7.5 impact=8 confidence=0.9 effort=1 horizon=Immediate category=contracts/interfaces reference=internal/exec/runner.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/02-contracts-interfaces-integration.md" \
  --file "internal/exec/runner.go" \
  --file "internal/exec/oracle_validate.go" \
  -p "$(cat <<'PROMPT'
Strategist question #02

Reference: internal/exec/runner.go
Category: contracts/interfaces
Horizon: Immediate
ROI: 7.5 (impact=8, confidence=0.9, effort=1)

Question:
What are the top integration points with external systems (oracle CLI invocation, filesystem outputs, environment variables), and what contract(s) or config declare them? Provide the minimal list of files/locations that define each integration.

Rationale (one sentence):
Integration boundaries drive risk, deployment needs, and test strategy.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=6.8 impact=8 confidence=0.85 effort=1 horizon=NearTerm category=invariants reference=internal/pack/parser.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/03-invariants-domain.md" \
  --file "internal/pack/parser.go" \
  --file "internal/pack/types.go" \
  -p "$(cat <<'PROMPT'
Strategist question #03

Reference: internal/pack/parser.go
Category: invariants
Horizon: NearTerm
ROI: 6.8 (impact=8, confidence=0.85, effort=1)

Question:
List the system’s most important invariants about pack structure and step ordering. For each, show where it is enforced (or where it should be enforced but currently is not).

Rationale (one sentence):
Invariants define correctness and are the backbone of reliable changes.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=6.2 impact=7 confidence=0.8 effort=1 horizon=NearTerm category=invariants reference=internal/app/app.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/04-invariants-validation.md" \
  --file "internal/app/app.go" \
  --file "internal/pack/parser.go" \
  -p "$(cat <<'PROMPT'
Strategist question #04

Reference: internal/app/app.go
Category: invariants
Horizon: NearTerm
ROI: 6.2 (impact=7, confidence=0.8, effort=1)

Question:
Where does validation happen (pack parse/validate, config resolution, runtime checks)? Identify validation boundaries and the most likely gaps that could cause inconsistent state.

Rationale (one sentence):
Knowing validation boundaries prevents regressions and reduces correctness risk.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=6.0 impact=7 confidence=0.8 effort=1 horizon=NearTerm category=caching/state reference=internal/state/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/05-caching-state-layers.md" \
  --file "internal/state/types.go" \
  --file "internal/state/persist.go" \
  -p "$(cat <<'PROMPT'
Strategist question #05

Reference: internal/state/types.go
Category: caching/state
Horizon: NearTerm
ROI: 6.0 (impact=7, confidence=0.8, effort=1)

Question:
What stateful components exist (run state, status tracking, persisted state files)? For each, describe lifecycle, read/write points, and where it is implemented.

Rationale (one sentence):
State handling is a common source of subtle bugs and operational issues.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=5.8 impact=7 confidence=0.75 effort=1 horizon=NearTerm category=caching/state reference=internal/app/run.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/06-caching-state-consistency.md" \
  --file "internal/app/run.go" \
  --file "internal/state/types.go" \
  -p "$(cat <<'PROMPT'
Strategist question #06

Reference: internal/app/run.go
Category: caching/state
Horizon: NearTerm
[TRUNCATED]
```

.config/commands/oracle-pack_v2.toml
```
# path: commands/oraclepack/oracle-pack_v2.toml
# invoked via: /oraclepack:oracle-pack_v2 {{args}}

description = "Generate a strategist-questions Oracle pack that validates/runs under oraclepack (```bash fence + # NN) headers), writes oracle outputs under docs/, and preserves the Codex-style template (parsed args + ROI ordering + coverage check)."

prompt = """
You are generating output for a STRICT downstream CLI parser (oraclepack).
Your output MUST be parseable by oraclepack’s Markdown parser:
- The commands MUST be inside a ```bash fenced code block (triple backticks), not ~~~.
- Step headers MUST be exactly: # NN) ... (e.g., # 01) ...). Do NOT use em dashes (—) in place of the ')'.

User args (raw): {{args}}

------------------------------------------------------------
A) ARG PARSING (no follow-ups; apply defaults)
Extract from {{args}} if present (key=value or similar), else default:

- codebase_name: Unknown
- constraints: None
- non_goals: None
- team_size: Unknown
- deadline: Unknown

Oracle pack controls:
- out_dir: docs/oracle/strategist-questions/$(date +%F)
- oracle_cmd: oracle
- oracle_flags: --browser-attachments always --files-report
- extra_files: (empty; if present, treat as comma-separated file paths/globs to include as additional -f entries in EVERY command)

You MUST render these values into the final output under the “parsed args” section.

------------------------------------------------------------
B) INFERENCE-FIRST DISCOVERY (adaptive; evidence-driven; deterministic)
Goal: infer what exists before searching broadly.

1) Prefer cheap “index” artifacts first:
   - README / docs index (if present)
   - primary manifests (package.json, pyproject.toml, go.mod, Cargo.toml, etc.)
   - obvious entrypoints referenced by scripts/manifests

2) Derive a search plan from evidence:
   - follow imports/registrations from entrypoints into:
     routing/handlers, auth/permissions, jobs/queues, data layer/migrations,
     feature flags, observability/logging, caching/state, failure modes.

3) Harvest >= 20 candidate anchors, each as ONE of:
   - {path}:{symbol} (preferred)
   - {endpoint}
   - {event}
If a category cannot be evidenced, keep its anchor as Unknown AND record the missing artifact pattern.

Required coverage categories (must appear across your 20 questions):
- contracts/interfaces
- invariants
- caching/state
- background jobs
- observability
- permissions
- migrations
- UX flows
- failure modes
- feature flags

------------------------------------------------------------
C) QUESTION GENERATION (exactly 20; 12 Immediate + 8 Strategic)
For each question, produce:
- Reference: {path}:{symbol} OR {endpoint} OR {event} OR Unknown
- Category: one of the required categories (or a sensible subcategory that still maps to them)
- Horizon: Immediate or Strategic
- Question: focused, feasibility-first
- Rationale: EXACTLY one sentence
- Smallest experiment today: EXACTLY one concrete action

No duplicates by intent or reference.

------------------------------------------------------------
D) ROI SCORING + ORDERING (required)
For each question, choose:
- impact, confidence, effort ∈ {0.1..1.0} with one decimal
Compute:
- ROI = (impact * confidence) / effort

Sort ALL 20 commands by:
1) ROI descending
2) tie-breaker: lower effort first

Numbering remains 01..20 in this final sorted order.

------------------------------------------------------------
E) ORACLE COMMAND EMISSION + ATTACHMENTS (minimal; evidence-driven)
For each command:
- Use exactly:
  <oracle_cmd> \
    <oracle_flags> \
    --write-output "<out_dir>/<nn>-<slug>.md" \
    -p "<prompt>" \
    -f "<file 1>" \
    -f "<file 2>" \
    ... \
    <extra_files...>

Attachment minimization:
- Prefer 1–3 attachments per command.
- If Reference is {path}:{symbol}: attach that file; optionally ONE more upstream entrypoint/router/config if needed.
- If Reference is {endpoint}: attach route map + handler file (if different).
- If Reference is {event}: attach job registration + worker implementation (if different).
- If Reference is Unknown: attach only “index” files (README + primary manifest + 1 best-guess entrypoint if found).
- NEVER attach files you did not discover or cannot justify.

extra_files behavior:
- If extra_files was provided, append those as additional -f entries at the end of EVERY command (after the minimal evidence attachments).

Slug rules:
- <slug> lowercase a-z0-9 and hyphens only.
- Derive from category + a hint of reference (fallback: category only).

------------------------------------------------------------
F) SAVE / WRITE SEMANTICS (required; docs/)
- The oracle outputs are written by oracle itself via --write-output "<out_dir>/...".
  Because out_dir defaults under docs/, oracle outputs will land under docs/.
- This slash-command response (the pack Markdown) is printed to stdout.
  The caller MUST save it to the path printed on the first line:
  Output file: docs/strategist-questions-oracle-pack-YYYY-MM-DD.md
  (If you cannot determine date, use: docs/strategist-questions-oracle-pack.md)

------------------------------------------------------------
G) OUTPUT FORMAT (STRICT; MUST MATCH THIS SHAPE)
1) First line of the assistant response MUST be:
   Output file: docs/strategist-questions-oracle-pack-YYYY-MM-DD.md
   (Use today’s date if known; otherwise: docs/strategist-questions-oracle-pack.md)

2) After that line, output EXACTLY ONE Markdown document matching the template below.
3) Do NOT add any extra prose, headings, or fences beyond what the template requires.
4) The ```bash code fence MUST be triple-backtick fenced, and MUST close with triple-backticks.
5) Step headers inside the bash fence MUST be: # NN) ... (to satisfy oraclepack).

TEMPLATE (MUST MATCH):
<!-- begin template -->
# oracle strategist question pack

---

## parsed args

- codebase_name: <Unknown|value>
- constraints: <None|value>
- non_goals: <None|value>
- team_size: <Unknown|value>
- deadline: <Unknown|value>
- out_dir: <docs/oracle/strategist-questions/$(date +%F)|value>
- oracle_cmd: <oracle|value>
- oracle_flags: <--browser-attachments always --files-report|value>
- extra_files: <empty|value>

---

## commands (exactly 20; sorted by ROI desc; ties by lower effort)

```bash
out_dir="<out_dir>"
mkdir -p "$out_dir"

# 01) ROI=<..> impact=<..> confidence=<..> effort=<..> horizon=<Immediate|Strategic> category=<...> reference=<...>
<oracle_cmd> \
  <oracle_flags> \
  --write-output "<out_dir>/01-<slug>.md" \
  -p "Strategist question #01
Reference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>
Category: <one of required categories>
Horizon: <Immediate|Strategic>
ROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)
Question: <question text>
Rationale: <exactly one sentence>
Smallest experiment today: <exactly one action>
Constraints: <constraints or None>
Non-goals: <non_goals or None>

Answer format:
[TRUNCATED]
```

.config/completion/oraclepack.completion.sh
```
# bash completion V2 for oraclepack                           -*- shell-script -*-

__oraclepack_debug()
{
    if [[ -n ${BASH_COMP_DEBUG_FILE-} ]]; then
        echo "$*" >> "${BASH_COMP_DEBUG_FILE}"
    fi
}

# Macs have bash3 for which the bash-completion package doesn't include
# _init_completion. This is a minimal version of that function.
__oraclepack_init_completion()
{
    COMPREPLY=()
    _get_comp_words_by_ref "$@" cur prev words cword
}

# This function calls the oraclepack program to obtain the completion
# results and the directive.  It fills the 'out' and 'directive' vars.
__oraclepack_get_completion_results() {
    local requestComp lastParam lastChar args

    # Prepare the command to request completions for the program.
    # Calling ${words[0]} instead of directly oraclepack allows handling aliases
    args=("${words[@]:1}")
    requestComp="${words[0]} __complete ${args[*]}"

    lastParam=${words[$((${#words[@]}-1))]}
    lastChar=${lastParam:$((${#lastParam}-1)):1}
    __oraclepack_debug "lastParam ${lastParam}, lastChar ${lastChar}"

    if [[ -z ${cur} && ${lastChar} != = ]]; then
        # If the last parameter is complete (there is a space following it)
        # We add an extra empty parameter so we can indicate this to the go method.
        __oraclepack_debug "Adding extra empty parameter"
        requestComp="${requestComp} ''"
    fi

    # When completing a flag with an = (e.g., oraclepack -n=<TAB>)
    # bash focuses on the part after the =, so we need to remove
    # the flag part from $cur
    if [[ ${cur} == -*=* ]]; then
        cur="${cur#*=}"
    fi

    __oraclepack_debug "Calling ${requestComp}"
    # Use eval to handle any environment variables and such
    out=$(eval "${requestComp}" 2>/dev/null)

    # Extract the directive integer at the very end of the output following a colon (:)
    directive=${out##*:}
    # Remove the directive
    out=${out%:*}
    if [[ ${directive} == "${out}" ]]; then
        # There is not directive specified
        directive=0
    fi
    __oraclepack_debug "The completion directive is: ${directive}"
    __oraclepack_debug "The completions are: ${out}"
}

__oraclepack_process_completion_results() {
    local shellCompDirectiveError=1
    local shellCompDirectiveNoSpace=2
    local shellCompDirectiveNoFileComp=4
    local shellCompDirectiveFilterFileExt=8
    local shellCompDirectiveFilterDirs=16
    local shellCompDirectiveKeepOrder=32

    if (((directive & shellCompDirectiveError) != 0)); then
        # Error code.  No completion.
        __oraclepack_debug "Received error from custom completion go code"
        return
    else
        if (((directive & shellCompDirectiveNoSpace) != 0)); then
            if [[ $(type -t compopt) == builtin ]]; then
                __oraclepack_debug "Activating no space"
                compopt -o nospace
            else
                __oraclepack_debug "No space directive not supported in this version of bash"
            fi
        fi
        if (((directive & shellCompDirectiveKeepOrder) != 0)); then
            if [[ $(type -t compopt) == builtin ]]; then
                # no sort isn't supported for bash less than < 4.4
                if [[ ${BASH_VERSINFO[0]} -lt 4 || ( ${BASH_VERSINFO[0]} -eq 4 && ${BASH_VERSINFO[1]} -lt 4 ) ]]; then
                    __oraclepack_debug "No sort directive not supported in this version of bash"
                else
                    __oraclepack_debug "Activating keep order"
                    compopt -o nosort
                fi
            else
                __oraclepack_debug "No sort directive not supported in this version of bash"
            fi
        fi
        if (((directive & shellCompDirectiveNoFileComp) != 0)); then
            if [[ $(type -t compopt) == builtin ]]; then
                __oraclepack_debug "Activating no file completion"
                compopt +o default
            else
                __oraclepack_debug "No file completion directive not supported in this version of bash"
            fi
        fi
    fi

    # Separate activeHelp from normal completions
    local completions=()
    local activeHelp=()
    __oraclepack_extract_activeHelp

    if (((directive & shellCompDirectiveFilterFileExt) != 0)); then
        # File extension filtering
        local fullFilter="" filter filteringCmd

        # Do not use quotes around the $completions variable or else newline
        # characters will be kept.
        for filter in ${completions[*]}; do
            fullFilter+="$filter|"
        done

        filteringCmd="_filedir $fullFilter"
        __oraclepack_debug "File filtering command: $filteringCmd"
        $filteringCmd
    elif (((directive & shellCompDirectiveFilterDirs) != 0)); then
        # File completion for directories only

        local subdir
        subdir=${completions[0]}
        if [[ -n $subdir ]]; then
            __oraclepack_debug "Listing directories in $subdir"
            pushd "$subdir" >/dev/null 2>&1 && _filedir -d && popd >/dev/null 2>&1 || return
        else
            __oraclepack_debug "Listing directories in ."
            _filedir -d
        fi
    else
        __oraclepack_handle_completion_types
    fi

    __oraclepack_handle_special_char "$cur" :
    __oraclepack_handle_special_char "$cur" =

    # Print the activeHelp statements before we finish
    __oraclepack_handle_activeHelp
}

__oraclepack_handle_activeHelp() {
    # Print the activeHelp statements
    if ((${#activeHelp[*]} != 0)); then
        if [ -z $COMP_TYPE ]; then
            # Bash v3 does not set the COMP_TYPE variable.
            printf "\n";
            printf "%s\n" "${activeHelp[@]}"
            printf "\n"
            __oraclepack_reprint_commandLine
            return
        fi

        # Only print ActiveHelp on the second TAB press
        if [ $COMP_TYPE -eq 63 ]; then
            printf "\n"
            printf "%s\n" "${activeHelp[@]}"

            if ((${#COMPREPLY[*]} == 0)); then
                # When there are no completion choices from the program, file completion
                # may kick in if the program has not disabled it; in such a case, we want
                # to know if any files will match what the user typed, so that we know if
                # there will be completions presented, so that we know how to handle ActiveHelp.
                # To find out, we actually trigger the file completion ourselves;
                # the call to _filedir will fill COMPREPLY if files match.
                if (((directive & shellCompDirectiveNoFileComp) == 0)); then
                    __oraclepack_debug "Listing files"
                    _filedir
                fi
            fi

            if ((${#COMPREPLY[*]} != 0)); then
                # If there are completion choices to be shown, print a delimiter.
                # Re-printing the command-line will automatically be done
                # by the shell when it prints the completion choices.
                printf -- "--"
            else
                # When there are no completion choices at all, we need
                # to re-print the command-line since the shell will
                # not be doing it itself.
                __oraclepack_reprint_commandLine
            fi
        elif [ $COMP_TYPE -eq 37 ] || [ $COMP_TYPE -eq 42 ]; then
            # For completion type: menu-complete/menu-complete-backward and insert-completions
            # the completions are immediately inserted into the command-line, so we first
            # print the activeHelp message and reprint the command-line since the shell won't.
            printf "\n"
            printf "%s\n" "${activeHelp[@]}"

            __oraclepack_reprint_commandLine
        fi
    fi
}

__oraclepack_reprint_commandLine() {
[TRUNCATED]
```

.tickets/mcp/oraclepack-MCP.md
```
## MCP surface for `oraclepack` (so agents can act on Taskify artifacts)

### What to expose as MCP tools

Map the existing `oraclepack` CLI capabilities (validate/list/run + flags like `--no-tui`, `--out-dir`, `--oracle-bin`) into MCP tools so an agent can run packs non-interactively and then inspect the resulting state/report/artifacts.

Add a small “taskify helper” layer to make the **Stage-2 → Stage-3** workflow deterministic for agents:

* **Detect Stage-2 outputs** (dir-form `01-*.md..20-*.md` OR single-pack form) using the ordered resolver rules described in your skill.
* **Validate Stage-2 outputs** (exactly one match per prefix 01..20).
* **Validate Stage-3 Action Pack** structure constraints (single ```bash fence, step headers `# NN)`, etc.) before executing anything.
* **Summarize Stage-3 artifacts** (e.g., `_actions.json`, `_actions.md`, PRD path, `tm-complexity.json`, pipelines doc) so agents can immediately consume them.

### Transport choices (real-time vs local)

* **stdio** is simplest for local agent runtimes; it requires MCP messages only on stdout (logs must go to stderr). ([Model Context Protocol][1])
* **Streamable HTTP** is better for “real-time” multi-client usage; implement Origin validation and bind to localhost + auth to avoid DNS rebinding and local-network abuse. ([Model Context Protocol][1])

### Tool UX metadata (important for agents)

Use MCP **tool annotations** so clients can correctly present approval UX:

* mark validate/list/read tools as `readOnlyHint: true`
* mark run tools as `destructiveHint: true`, `openWorldHint: true` (they can touch filesystem, run processes, etc.) ([Model Context Protocol][2])

### Security defaults (recommended)

* Hard **deny-by-default execution**: require an env flag (e.g. `ORACLEPACK_ENABLE_EXEC=1`) to enable “run” tools.
* Restrict filesystem access to **allowlisted roots** to avoid path traversal and accidental exfiltration of secrets.
* Enforce stdout/stderr truncation and timeouts so a pack can’t wedge the server process.

## Reference implementation

I provided a complete, drop-in **Python MCP server** (FastMCP from the official MCP Python SDK) that exposes:

* `oraclepack_validate_pack`, `oraclepack_list_steps`, `oraclepack_run_pack`
* `oraclepack_taskify_detect_stage2`, `oraclepack_taskify_validate_stage2`
* `oraclepack_taskify_validate_action_pack`, `oraclepack_taskify_artifacts_summary`
* `oraclepack_taskify_run_action_pack`

It supports **stdio** and **streamable-http** transports per the official SDK. ([GitHub][3])

The full multi-file project is in the canvas titled **“oraclepack-mcp-server”**.

[1]: https://modelcontextprotocol.io/specification/2025-03-26/basic/transports "Transports - Model Context Protocol"
[2]: https://modelcontextprotocol.io/legacy/concepts/tools "Tools - Model Context Protocol"
[3]: https://github.com/modelcontextprotocol/python-sdk "GitHub - modelcontextprotocol/python-sdk: The official Python SDK for Model Context Protocol servers and clients"
```

.tickets/mcp/oraclepack_mcp_server.md
```
# oraclepack-mcp-server (MCP wrapper for oraclepack + taskify helpers)

## tree

```text
oraclepack-mcp-server/
  README.md
  requirements.txt
  oraclepack_mcp_server/
    __init__.py
    __main__.py
    config.py
    security.py
    oraclepack_cli.py
    taskify.py
    server.py
```

---

```md
<!-- path: oraclepack-mcp-server/README.md -->
# oraclepack-mcp-server

Expose `oraclepack` (validate/list/run) plus a few Stage-3 (oraclepack-taskify) helpers as **MCP tools**, so an agent can:
- validate and inspect packs
- run packs non-interactively (`--no-tui --yes --run-all`) to produce artifacts
- validate Stage-2 outputs (01-*.md..20-*.md)
- validate Stage-3 Action Packs (single ```bash fence, step headers, etc.)
- summarize Stage-3 artifacts (`_actions.json`, PRD, Task Master outputs, etc.)

## Install

```bash
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```

## Configure (recommended)

Environment variables:

- `ORACLEPACK_ALLOWED_ROOTS` (default: current working directory)
  - Colon-separated list of allowed filesystem roots the server may read from.
  - Example: `ORACLEPACK_ALLOWED_ROOTS=/repo:/tmp/oracle-out`
- `ORACLEPACK_BIN` (default: `oraclepack`) – path to the oraclepack CLI
- `ORACLEPACK_WORKDIR` (default: current working directory)
  - Where packs are executed from (important for relative paths).
- `ORACLEPACK_ENABLE_EXEC` (default: `0`)
  - Must be `1` to enable `oraclepack_run_pack` and `oraclepack_taskify_run_action_pack`.
- `ORACLEPACK_CHARACTER_LIMIT` (default: `25000`) – truncate large stdout/stderr
- `ORACLEPACK_MAX_READ_BYTES` (default: `500000`) – max bytes read from a file

## Run (stdio)

```bash
# Stdio transport is the simplest local integration.
python -m oraclepack_mcp_server --transport stdio
```

## Run (Streamable HTTP)

```bash
python -m oraclepack_mcp_server --transport streamable-http
```

## Tools

- `oraclepack_validate_pack`
- `oraclepack_list_steps`
- `oraclepack_run_pack` (requires `ORACLEPACK_ENABLE_EXEC=1`)
- `oraclepack_read_file`
- `oraclepack_taskify_detect_stage2`
- `oraclepack_taskify_validate_stage2`
- `oraclepack_taskify_validate_action_pack`
- `oraclepack_taskify_artifacts_summary`
- `oraclepack_taskify_run_action_pack` (requires `ORACLEPACK_ENABLE_EXEC=1`)

## Typical Stage-3 usage

1) Detect/validate Stage-2 outputs (directory or single-pack)
2) Validate the Action Pack markdown
3) Run the Action Pack via `oraclepack run ...`
4) Summarize produced artifacts
```

```txt
# path: oraclepack-mcp-server/requirements.txt
mcp>=1.0.0
pydantic>=2.0.0
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/__init__.py
__all__ = []
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/config.py
from __future__ import annotations

import os
from dataclasses import dataclass
from pathlib import Path


def _truthy(value: str | None) -> bool:
    if value is None:
        return False
    return value.strip().lower() in {"1", "true", "yes", "y", "on"}


@dataclass(frozen=True)
class ServerConfig:
    allowed_roots: tuple[Path, ...]
    oraclepack_bin: str
    workdir: Path
    enable_exec: bool
    character_limit: int
    max_read_bytes: int


def load_config() -> ServerConfig:
    # Allowed roots: colon-separated. Default to CWD.
    roots_raw = os.environ.get("ORACLEPACK_ALLOWED_ROOTS")
    if roots_raw:
        roots = tuple(Path(p).expanduser().resolve() for p in roots_raw.split(":") if p.strip())
    else:
        roots = (Path.cwd().resolve(),)

    oraclepack_bin = os.environ.get("ORACLEPACK_BIN", "oraclepack")
    workdir = Path(os.environ.get("ORACLEPACK_WORKDIR", str(Path.cwd()))).expanduser().resolve()

    enable_exec = _truthy(os.environ.get("ORACLEPACK_ENABLE_EXEC", "0"))

    character_limit = int(os.environ.get("ORACLEPACK_CHARACTER_LIMIT", "25000"))
    max_read_bytes = int(os.environ.get("ORACLEPACK_MAX_READ_BYTES", "500000"))

    return ServerConfig(
        allowed_roots=roots,
        oraclepack_bin=oraclepack_bin,
        workdir=workdir,
        enable_exec=enable_exec,
        character_limit=character_limit,
        max_read_bytes=max_read_bytes,
    )
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/security.py
from __future__ import annotations

from pathlib import Path


class PathNotAllowedError(ValueError):
    pass


def resolve_under_roots(path: Path, allowed_roots: tuple[Path, ...]) -> Path:
    """Resolve a path and enforce it lives under at least one allowed root."""
    p = path.expanduser().resolve()

    for root in allowed_roots:
        r = root.expanduser().resolve()
        try:
            p.relative_to(r)
            return p
        except ValueError:
            continue

    raise PathNotAllowedError(
        f"Path not allowed (outside allowed roots). path={p} roots={[str(r) for r in allowed_roots]}"
    )


def safe_read_text(path: Path, max_bytes: int) -> tuple[str, bool]:
    """Read up to max_bytes from a file as UTF-8 (replace errors)."""
    data = path.read_bytes()
    truncated = False
    if len(data) > max_bytes:
        data = data[:max_bytes]
        truncated = True
    return data.decode("utf-8", errors="replace"), truncated


def safe_read_bytes(path: Path, max_bytes: int) -> tuple[bytes, bool]:
    data = path.read_bytes()
    truncated = False
    if len(data) > max_bytes:
        data = data[:max_bytes]
        truncated = True
    return data, truncated
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/oraclepack_cli.py
from __future__ import annotations

import asyncio
import os
import time
from dataclasses import dataclass
from pathlib import Path


@dataclass
class CmdResult:
    ok: bool
    exit_code: int
    duration_s: float
    stdout: str
    stderr: str
    stdout_truncated: bool
    stderr_truncated: bool


def _truncate(s: str, limit: int) -> tuple[str, bool]:
    if limit <= 0:
        return s, False
    if len(s) <= limit:
        return s, False
    return s[:limit], True


async def run_cmd(
    argv: list[str],
    cwd: Path,
    timeout_s: int,
    env: dict[str, str] | None,
    character_limit: int,
) -> CmdResult:
    start = time.time()

    proc = await asyncio.create_subprocess_exec(
        *argv,
        cwd=str(cwd),
        env=(os.environ | (env or {})),
        stdout=asyncio.subprocess.PIPE,
        stderr=asyncio.subprocess.PIPE,
    )

    try:
        out_b, err_b = await asyncio.wait_for(proc.communicate(), timeout=timeout_s)
    except asyncio.TimeoutError:
        proc.kill()
        await proc.communicate()
        duration = time.time() - start
        return CmdResult(
            ok=False,
            exit_code=124,
            duration_s=duration,
            stdout="",
            stderr=f"Timed out after {timeout_s}s: {' '.join(argv)}",
            stdout_truncated=False,
            stderr_truncated=False,
        )

    duration = time.time() - start
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/01-contracts-interfaces-surface.md
```
1) Direct answer (1–6 bullets, evidence-cited)

Primary public interface #1: CLI root command oraclepack (Cobra)

Inputs: persistent flags

--no-tui (bool, default false) to disable TUI

--oracle-bin (string, default "oracle") path to oracle binary

--out-dir / -o (string, default "") output directory for step execution

Outputs: process exit code is mapped via errors.ExitCode(err) when rootCmd.Execute() fails.

Defined in: internal/cli/root.go (root command + persistent flags + Execute error handling).

Primary public interface #2: CLI command validate [pack.md]

Inputs: one positional arg pack.md (cobra.ExactArgs(1)), passed as app.Config{PackPath: args[0]}

Outputs: on success prints Pack is valid.; on failure returns an error from a.LoadPack() (then surfaced by root Execute path).

Defined in: internal/cli/cmds.go (validate command + wiring into root).

Primary public interface #3: CLI command list [pack.md]

Inputs: one positional arg pack.md (cobra.ExactArgs(1)), passed as app.Config{PackPath: args[0]}

Outputs: prints one line per step: "<step.ID>: <step.OriginalLine>", iterating a.Pack.Steps.

Defined in: internal/cli/cmds.go (list command + wiring into root).

Primary public interface #4: “oracle-pack” Markdown contract (pack file)

Inputs (implied by CLI): a Markdown file path (pack.md) that can be loaded by a.LoadPack() from app.New(cfg) where cfg.PackPath is set.

Outputs (observable via list): the parsed pack must produce a.Pack.Steps, and each step must at least have fields ID and OriginalLine (since they’re printed).

Where defined in code: only the usage is shown in internal/cli/cmds.go; the actual contract/schema + parser are in the internal/app package (not included in the provided files).

2) Risks/unknowns (bullets)

Additional CLI commands/flags may exist in other internal/cli/*.go files; only validate and list are visible here.

The actual oracle-pack Markdown schema (required headings, step syntax, validation rules, parsing behavior, versioning) is not visible; we only know LoadPack() yields Pack.Steps[] with ID and OriginalLine.

Exit-code mapping and error taxonomy depend on internal/errors (only referenced, not shown).

--oracle-bin / --out-dir are declared at root-level but not used in the shown commands; actual execution interface likely exists in other commands not included.

3) Next smallest concrete experiment (exactly one action)

Run a repo-wide search for the contract types and parser entrypoints: rg -n "type Pack|type Step|LoadPack\\(|PackPath|OriginalLine" internal/ to locate the canonical oracle-pack Markdown schema definitions.

4) If evidence is insufficient, missing file/path pattern(s) to attach next

internal/app/*.go (especially files defining Config, New, LoadPack, and any Pack / Step structs)

internal/errors/*.go (to confirm public error*
```

docs/oracle-questions-2026-01-07/02-contracts-interfaces-integration.md
```
1) Direct answer (top integration points, minimal defining files/locations)

Shell execution boundary (runs arbitrary scripts via OS process execution)

Integration: os/exec launches a shell with -lc (login shell semantics), sets working dir, env, and streams stdout/stderr. This is the core “external system” boundary for running everything.

Contract/config: RunnerOptions{Shell, WorkDir, Env, OracleFlags, Overrides, ChatGPTURL} → Runner (constructed in NewRunner).

Defined in: internal/exec/runner.go (NewRunner, RunPrelude, RunStep, run). Evidence: exec.CommandContext(ctx, r.Shell, "-lc", script), cmd.Dir = r.WorkDir, cmd.Env = r.Env, cmd.Stdout/Stderr = logWriter.

Oracle CLI invocation boundary (flag injection + per-step overrides + optional dry-run validation)

Integration: Oracle commands are expected to appear inside step code; runner injects flags into those commands before execution, and validation can run oracle --dry-run summary for overridden steps.

Contract/config: baseline Runner.OracleFlags + optional overrides.RuntimeOverrides (scoped by step.ID) + optional ChatGPTURL mapped into flags (ApplyChatGPTURL); validation is scoped by over.ApplyToSteps[step.ID].

Defined in:

internal/exec/runner.go (RunStep: ApplyChatGPTURL, Overrides.EffectiveFlags, InjectFlags).

internal/exec/oracle_validate.go (ValidateOverrides: ExtractOracleInvocations, --dry-run summary, execDryRun).

Note: The actual parsing/injection logic lives elsewhere (functions referenced but not shown): ExtractOracleInvocations, InjectFlags, ApplyChatGPTURL.

Environment-variable boundary (inherit + append + PATH propagation during validation)

Integration: Environment is inherited from the parent process and extended with caller-provided entries; validation explicitly re-exports PATH into the command string to ensure the shell sees it.

Contract/config: RunnerOptions.Env merged with os.Environ(); findEnvValue(env, "PATH") drives PATH export behavior in dry-run validation.

Defined in:

internal/exec/runner.go (NewRunner: Env: append(os.Environ(), opts.Env...)).

internal/exec/oracle_validate.go (env := append(os.Environ(), opts.Env...), execDryRun PATH export + quoting via shellQuote).

Filesystem outputs boundary (working directory as the implicit artifact root)

Integration: All script execution happens with cmd.Dir = r.WorkDir, so any relative file reads/writes done by steps land under that directory (and any output directories are created by the scripts themselves).

Contract/config: RunnerOptions.WorkDir / Runner.WorkDir.

Defined in: internal/exec/runner.go (run: cmd.Dir = r.WorkDir), plus va
```

docs/oracle-questions-2026-01-07/03-invariants-domain.md
```
1) Direct answer (1–8 bullets, evidence-cited)

Pack must contain a fenced ```bash ... ``` code block (the pack is parsed only from that block).
Enforced: Parse() searches bashFenceRegex and errors if none found (no bash code block found).
Not enforced / edge: if multiple bash blocks exist, only the first match is used (FindSubmatch), others are ignored.
Evidence: bashFenceRegex = ... "```bash\n(.*?)\n```" and if match == nil ... return ... ErrInvalidPack in internal/pack/parser.go.

Steps are defined only by Markdown-style headers inside the bash block, and must start with # + a 2-digit step id plus an accepted separator.
Enforced (implicitly by parsing): step boundaries are recognized only when stepHeaderRegex matches a line; everything else is either prelude (before first step) or step body.
Invariants implied: step IDs must be exactly two digits (\d{2}) and separator must be one of ), —, or - (with whitespace variants).
Evidence: stepHeaderRegex = regexp.MustCompile(^#\s*(\d{2})(?:\)|[\s]+[—-])) and the headerMatch := ...; if len(headerMatch) > 1 { ... } logic in internal/pack/parser.go.

All lines before the first step header become “Prelude”; all lines after step parsing begins are assigned to the current step’s Code.
Enforced: inSteps toggles true at first header; before that lines go to preludeLines; after that lines append to currentStep.Code.
Should be considered: any configuration intended for metadata extraction must appear in the prelude, or it won’t be discovered.
Evidence: var preludeLines []string; var inSteps bool + if inSteps { currentStep.Code += ... } else { preludeLines = append(...) } and pack.Prelude.Code = strings.Join(preludeLines, "\n") in internal/pack/parser.go.

At least one step is required for a valid pack.
Enforced: Validate() rejects packs with len(p.Steps) == 0. Note: Parse() itself does not call Validate()—caller must.
Evidence: if len(p.Steps) == 0 { return ... ErrInvalidPack } in internal/pack/parser.go.

Step numbers must be positive, unique, and sequential starting at 1 in the encountered order (i.e., the file order is the order).
Enforced: Validate() checks step.Number > 0, no duplicates (seen map), and step.Number == i+1. This implies the first step must be 01 (because 00 parses to 0 and fails).
Evidence: seen := make(map[int]bool) + duplicate check + if step.Number != i+1 { ... sequential starting from 1 ... } in internal/pack/parser.go.

Metadata invariants: out_dir=... and --write-output are derived only from the Prelude, with strict-ish patterns.
Enforced (extraction behavior): DeriveMetadata() looks only at p.Prelude.Code; out_dir must match ^out_dir=... with no spaces around =, and --write-output is a simple substring/regex presence test.
Not enforced: packs can place these in steps and still “work” at runtime, but metadata will be wrong; multiple out_dir= lines are not resolved explicitly (first match_
```

docs/oracle-questions-2026-01-07/04-invariants-validation.md
```
1) Direct answer (validation boundaries + likely gaps)

Pack parse/structural validation happens in App.LoadPack(): it reads the pack file, calls pack.Parse(data), then p.Validate() and fails fast on any error (internal/app/app.go:LoadPack).

Boundary: “bytes on disk” → “in-memory Pack that is considered valid enough to run.”

Parse-time validation is minimal and mostly syntactic: pack.Parse() rejects packs without a fenced ```bash block (bashFenceRegex) and only recognizes steps whose headers match stepHeaderRegex (internal/pack/parser.go:Parse).

Gap: a Markdown pack can be “accepted” but silently miss intended steps if headers don’t match the regex (e.g., # 1 - ... vs # 01 - ...), leading to Validate() failing later with “at least one step is required” or producing an unexpected step set.

Pack semantic validation is narrow (step numbering only): Pack.Validate() enforces: at least one step, positive numbers, no duplicates, and sequential numbering starting from 1 in parse order (internal/pack/parser.go:Validate).

Gap: it does not validate step code presence/emptiness, command safety, required prelude directives, or cross-step dependencies—so many “bad” packs pass validation and fail at runtime.

Metadata extraction happens without validation: after parsing, DeriveMetadata() extracts out_dir=... and --write-output from the prelude using regexes (internal/pack/parser.go:DeriveMetadata).

Gap: extracted OutDir is not validated for emptiness, path traversal, or consistency with CLI overrides (it’s just a string captured by regex).

Config resolution validation is effectively “filesystem-only”: App.Prepare() resolves outDir by precedence CLI cfg.OutDir > pack p.OutDir > ".", then only validates by attempting os.MkdirAll(outDir, 0755) (internal/app/app.go:Prepare).

Gap: there’s no validation of other config fields here (PackPath, StatePath, ROIMode values, ROIThreshold range, etc.), so invalid config can survive until deeper runtime code.

State validation boundary is weak and can mask corruption: App.LoadState() on Resume tries state.LoadState(...); if it errors, it silently falls back to a fresh RunState{SchemaVersion:1, StepStatuses: map[...]} (internal/app/app.go:LoadState).

Most likely inconsistent-state gap: partial outputs on disk + corrupted/old state file ⇒ the system “resumes” from an empty state and may rerun steps unexpectedly or skip intended recovery paths.

Runtime checks shown are limited to environment shaping: Prepare() appends out_dir=<resolved> into Runner.Env but does not persist the resolved outDir back into Config/State, and does not de-dup/override existing out_dir entries (internal/app/app.go:Prepare).

Gap: scripts may see multiple out_dir values depending on how Runner.Env is constructed elsewhere; plus “where outputs went” is not recorded in state, increasing resume/report inconsistency risk.

Parser robustness gap that can cause silent truncation: pack.Parse() uses bufio.Scanner but does not call scanner.Err() after the scan loop, nor increase the scanner buffer (internal/pack/parser.go:Parse).

Likely inconsistent-state failure mode: long lines in bash blocks can exceed Scanner’s token lim
```

docs/oracle-questions-2026-01-07/05-caching-state-layers.md
```
Direct answer (1–8 bullets, evidence-cited)

Persisted run-state file (RunState JSON)

What it is: A single JSON document representing an oracle-pack run: schema version, pack hash, start time, per-step status map, ROI settings, and warnings. (internal/state/types.go: RunState{SchemaVersion, PackHash, StartTime, StepStatuses, ROIThreshold, ROIMode, Warnings})

Lifecycle (as implemented in this package): In-memory RunState is serialized to disk and later deserialized back into memory. (internal/state/persist.go: SaveStateAtomic, LoadState)

Read/write points: Writes via SaveStateAtomic(path, state); reads via LoadState(path). (internal/state/persist.go)

Per-step status tracking (StepStatus)

What it is: A per-step record with state machine values pending|running|success|failed|skipped, plus exit code, timestamps, and error string. (internal/state/types.go: type Status, constants; StepStatus{Status, ExitCode, StartedAt, EndedAt, Error})

Lifecycle (represented, not enforced here): A step typically transitions pending → running → (success|failed) or pending → skipped; timestamps support “running window” and duration computations. The code shown defines the model but not the transition logic. (internal/state/types.go)

Run warnings (non-fatal notes)

What it is: Structured warnings attached to run state, scoped and optionally tied to a specific step and line/token. (internal/state/types.go: Warning{Scope, StepID, Line, Token, Message}; RunState.Warnings)

Lifecycle / read/write: Persisted as part of RunState JSON alongside statuses. (internal/state/types.go; internal/state/persist.go)

Atomic persistence mechanism

What it does: Writes path + ".tmp" then renames to path, removing the temp file on rename failure. (internal/state/persist.go: tempPath := path + ".tmp", os.WriteFile, os.Rename, os.Remove)

Operational behavior: Output is pretty-printed JSON (json.MarshalIndent) which aids debugging but can be larger/slower than compact JSON. (internal/state/persist.go)

State “identity” / cache key fields

What it is: PackHash suggests the run state is intended to be tied to a specific pack content hash, likely to detect “pack changed vs. old state”. (internal/state/types.go: RunState.PackHash)

Lifecycle: The code provided stores the value but does not validate it during load; any reconciliation logic must be elsewhere. (internal/state/types.go; internal/state/persist.go)

Risks/unknowns (bullets)

Unknown where state is created/updated. The snippets define the data model and persistence, but not when StepStatuses are initialized, advanced, or saved (e.g., after each step, only at end, on error). That logic is likely in the executor/runner layer, not shown.

Missing “no state yet” behavior. LoadState returns an error when the state file does not exist ("state file not found"), which may force callers to treat “first run” as an error unless they special-case it. (internal/state/persist.go: os.IsNotExist(err))

Map key ambiguity. StepStatuses map[string]StepStatus does not guarantee ordering; the key format (step ID? label?) and collision rules are unknown without the parser/runner code that populates it. (internal/state/types.go)

Cross-platform atomicity. os.Rename(tempPath, path) may not overwrite an existing destination file on some p
```

docs/oracle-questions-2026-01-07/06-caching-state-consistency.md
```
1) Direct answer (top consistency risks + knobs)

Silent state-write failures → persisted state can lag in-memory indefinitely. saveState() discards errors (_ = state.SaveStateAtomic(...)), so the run can proceed while the on-disk state remains stale. Evidence: internal/app/run.go: saveState().

Non-transactional step transitions (RUNNING → SUCCESS/FAILED) can leave “stuck RUNNING” on disk. Each step does: set StatusRunning + saveState(), then later set EndedAt + final status + saveState(); any crash or write failure between these creates divergence (disk shows RUNNING; memory/report show final). Evidence: internal/app/run.go: RunPlain() loop setting StatusRunning then later StatusFailed/StatusSuccess.

Warnings can be persisted out of sync with step status. recordWarnings() drains and appends warnings and immediately calls saveState(); it runs after RunStep but before the final status save, so the state file can contain warnings from a step while still showing that step as RUNNING if the later status save fails. Evidence: internal/app/run.go: recordWarnings() + call sites around RunStep.

ROI-based skipping is not reflected in persisted state/report as “skipped”. ROI filter continues before writing any StepStatus; yet the state model has StatusSkipped available. This makes outputs inconsistent with console logs (“Skipping step…”) and loses auditability. Evidence: internal/app/run.go: ROIThreshold/ROIMode filter continues + internal/state/types.go: StatusSkipped exists.

StartTime is not guaranteed to flush. StartTime is set in-memory (if IsZero() { StartTime = time.Now() }) but there’s no immediate saveState(); if there are zero steps (and no warnings) you can generate a report with the updated StartTime while the persisted state file never updates. Evidence: internal/app/run.go: StartTime set; finalize() can run without any prior save.

Report write success is claimed even if the write fails. finalize() ignores json.MarshalIndent and os.WriteFile errors and prints “Report written to …” unconditionally, so report output can be missing/partial while stdout says it exists. Evidence: internal/app/run.go: finalize().

Resume behavior depends on persisted success markers and can misbehave with stale state. Resume skips only when StepStatuses[step.ID].Status == success; if state saves fail, already-completed steps can rerun (or previously-successful steps might not be recognized). Evidence: internal/app/run.go: resume check.

State behavior knobs/configs (observable here):

Config.StatePath: enables/disables persisted run state (saveState() is a no-op if empty). Evidence: internal/app/run.go: saveState().

Config.ReportPath: enables/disables report writing. Evidence: internal/app/run.go: finalize().

Config.StopOnFail: controls early termination path (still calls finalize() on failure). Evidence: internal/app/run.go: StopOnFail branch.

Config.ROIThreshold + Config.ROIMode: controls skip filter semantics (“under” is strict <; otherwise >=). Evidence: internal/app/run.go: ROI filter.

2) Risks/unknowns

Atomicity guarantees are unknown. state.SaveStateAtomic is assumed to be atomic, but wi
```

docs/oracle-questions-2026-01-07/07-background-jobs-discovery.md
```
1) Direct answer (1–8 bullets, evidence-cited)

No explicit background jobs/workers/scheduled tasks are defined in internal/app/run.go. The execution model is a single synchronous control flow in (*App).RunPlain: optional prelude, then a for-loop over steps, each executed via a.Runner.RunStep(...), followed by state updates and writes. Evidence: RunPlain(...) calls RunPrelude(...) and RunStep(...) inline (no goroutines, tickers, or async constructs shown).

Triggers: Only direct, immediate calls from RunPlain (prelude runs if a.Pack.Prelude.Code != ""; steps run when they pass ROI/resume filters). Evidence: prelude conditional and for _, step := range a.Pack.Steps { ... err := a.Runner.RunStep(...) ... }.

Payloads: The “work” payload is the current step (&step) and out writer passed into Runner.RunStep, plus the prelude (&a.Pack.Prelude). Evidence: a.Runner.RunPrelude(ctx, &a.Pack.Prelude, out) and a.Runner.RunStep(ctx, &step, out).

Retries: No retry loop/backoff exists here. The only “try again” behavior is manual restart/resume: successful steps are skipped on subsequent runs. Evidence: resume check if s, ok := a.State.StepStatuses[step.ID]; ok && s.Status == state.StatusSuccess { ... continue }.

Idempotency: “At-least-once with resume-skip on success” at the step level. Steps are marked running, persisted, then marked success/failed and persisted; rerun skips StatusSuccess steps. Evidence: status set to StatusRunning + a.saveState(), then on success sets StatusSuccess + a.saveState(), and skip logic above.

Scheduled behavior: None (no timers/cron). The only time usage is stamping StartTime, StartedAt, EndedAt via time.Now(). Evidence: a.State.StartTime = time.Now(), StartedAt: time.Now(), status.EndedAt = time.Now().

2) Risks/unknowns

a.Runner.RunPrelude, a.Runner.RunStep, and a.Runner.DrainWarnings() could hide concurrency (goroutines, worker pools, background process supervision) inside internal/exec or other packages; this file alone cannot confirm that.

state.SaveStateAtomic(...) could include locking/FS semantics or side effects relevant to “background-like” behavior (e.g., fsync patterns), though it is still invoked synchronously here.

CLI/TUI entrypoints may spawn background refreshers (status polling, UI redraw loops) outside run.go.

3) Next smallest concrete experiment (exactly one action)

Run a repo-wide search for background/scheduling primitives:
rg -n --hidden "time\.NewTicker|time\.AfterFunc|cron|scheduler|worker|goroutine|go\s+func|errgroup|WaitGroup|queue|async" internal/

4) If evidence is insufficient, missing file/path pattern(s) to attach next

internal/exec/**/*.go (especially internal/exec/runner.go and any process-management helpers)

internal/cli/**/*.go (entrypoints/TUI loops can introduce background refresh/polling)

internal/state/**/*.go (e.g., SaveStateAtomic implementation)
```

docs/oracle-questions-2026-01-07/08-background-jobs-reliability.md
```
1) Direct answer (1–8 bullets, evidence-cited)

No dedicated background job system is evident in the Go code shown (no queue/worker process, scheduler, or job runner). The primary execution path appears to be a synchronous CLI (Cobra) that runs work inline rather than enqueuing background tasks. Evidence: CLI entry is Cobra-based with flags (internal/cli/root.go), and execution uses context + os/exec to run an external command (internal/exec/runner.go).

Because there’s no background job substrate, classic reliability controls (DLQ/dead-lettering, queue backoff policies, concurrency limits, reprocessing UI) are not implemented as first-class “background work” features—failures would surface as immediate CLI errors rather than being quarantined/retried by a worker. Evidence: inline execution model suggested by os/exec runner (internal/exec/runner.go) and lack of worker constructs in the CLI surface (internal/cli/root.go).

What does exist is “run reliability” via status/state tracking, which is the closest analogue to reprocessing controls: the system models step/run outcomes with explicit statuses (pending/running/success/failed/skipped). Evidence: Status* constants in internal/state/types.go.

State/report paths are configurable at the app level, which suggests persistence of outcomes to disk (useful for restart/manual re-run), but this is not the same as automated retries/backoff. Evidence: app-wide config includes StatePath and Report... fields (internal/app/app.go), and run orchestration imports internal/state + internal/report (internal/app/run.go snippet).

Concurrency limiting is implicitly “1 run = 1 process” (single CLI invocation) unless internal/app/run.go spawns goroutines (not shown). No explicit worker pool/limit controls are visible in the provided excerpts. Evidence: only synchronous CLI + runner excerpts provided (internal/cli/root.go, internal/exec/runner.go, partial internal/app/run.go).

2) Risks/unknowns (bullets)

internal/app/run.go is truncated; it could still spawn goroutines, use tickers, or implement internal retry/backoff that isn’t visible in the snippet.

The persistence semantics (when state is flushed, atomicity, resume behavior) can’t be confirmed without the full implementations of internal/state/* and internal/report/*.

There may be background work outside the shown files (e.g., separate cmd/* binaries, build tags, or a different package) that wasn’t included in the excerpts.

3) Next smallest concrete experiment (exactly one action)

Run a repo-wide search for background-work primitives and libraries, e.g. rg -n "(time\\.NewTicker|time\\.Tick|cron|scheduler|queue|worker|goroutine|go\\s+func|asynq|machinery|temporal|celery|bull)" .

4) If evidence is insufficient, exact missing file/path pattern(s) to attach next

internal/app/run.go (full file)

internal/state/*.go

internal/report/*.go

cmd/** and internal/cli/*.go (full set, not snippets)

Any directories named or resembling jobs/, worker/, background/, scheduler/, queue/, cron/ (if present)
```

docs/oracle-questions-2026-01-07/09-observability-signals.md
```
1) Direct answer (observability + correlation identifiers)

Primary structured observability output is the run “report”: internal/report/GenerateReport(s *state.RunState, packName string) *ReportV1 builds a ReportV1 containing PackInfo{Name, Hash}, GeneratedAt, per-step results (Steps: []StepReport), and a Summary with counts + total duration. (Evidence: internal/report/generate.go, GenerateReport populating PackInfo, Steps, Summary.)

Per-step signals captured (and thus observable via the report) include: StepReport.ID (step identifier), Status, ExitCode, Duration/DurationMs, and Error. These are derived from RunState.StepStatuses map[string]StepStatus and each StepStatus’s timestamps/exit code/error. (Evidence: internal/report/generate.go builds StepReport from the for id, status := range s.StepStatuses loop; internal/state/types.go defines RunState.StepStatuses and StepStatus fields.)

Warnings are first-class, structured observability signals: RunState.Warnings []Warning are copied into report.Warnings with fields {Scope, StepID, Line, Token, Message}. (Evidence: internal/state/types.go defines Warning; internal/report/generate.go copies s.Warnings into report.Warnings.)

Primary correlation identifiers across components (based on what’s present in state/report):

pack_hash: RunState.PackHash and ReportV1.PackInfo.Hash (ties a run to a specific pack content). (Evidence: internal/state/types.go PackHash; internal/report/generate.go assigns Hash: s.PackHash.)

step_id: the map key in RunState.StepStatuses and StepReport.ID; also optionally present on warnings (Warning.StepID). This is the main join key between step execution, step status, and warnings. (Evidence: internal/state/types.go StepStatuses map[string]StepStatus, Warning.StepID; internal/report/generate.go uses loop variable id → StepReport.ID.)

start_time / generated_at: RunState.StartTime and ReportV1.GeneratedAt provide coarse run-level temporal correlation (e.g., match logs/events by time window). (Evidence: internal/state/types.go StartTime; internal/report/generate.go sets GeneratedAt: time.Now().)

There is no explicit “run_id”/trace id in the provided state/report schema; correlation appears to rely on (pack_hash, start_time) + step IDs, rather than a stable UUID. (Evidence: internal/state/types.go RunState fields shown do not include a run identifier other than PackHash/StartTime.)

Log lines: no logging primitives (log.*, fmt.Printf, etc.) are shown in the provided generate.go or state/types.go, so log-based observability and log correlation keys cannot be confirmed from these files alone. (Evidence: the provided snippets contain only struct composition / time math.)

2) Risks / unknowns

Run correlation fragility: without a dedicated run_id, correlating across components (state file writes, executor output, CLI output, report artifacts) will be brittle when multiple runs share the same pack hash or occur close in time.

Non-deterministic step ordering in reports: iterating for id, status := range s.StepStatuses over a Go map yields randomized order, which can create noisy diffs and make “compare reports across runs” harder.

Missing visibility into logging: cannot verify
```

docs/oracle-questions-2026-01-07/10-observability-gaps.md
```
1) Direct answer (1–8 bullets, evidence-cited)

No correlation IDs in execution logs, so step output cannot be tied back to a step/run reliably. run() wires cmd.Stdout/cmd.Stderr directly to logWriter with no prefix/header/trailer, and it receives no stepID/runID to stamp lines (run(ctx, script, logWriter) only gets the script + writer). Evidence: RunPrelude/RunStep call r.run(ctx, script, logWriter); run() sets cmd.Stdout = logWriter and cmd.Stderr = logWriter and emits no logs of its own. (internal/exec/runner.go: RunPrelude, RunStep, run)

Key decision points are silent (missing “what config did we actually use?” logs). Shell defaulting (/bin/bash), override application, ChatGPTURL flag application, and flag injection happen without any structured record, making misconfig/debugging difficult. Evidence: NewRunner defaults Shell; RunStep conditionally calls Overrides.EffectiveFlags(...) and ApplyChatGPTURL(...), then InjectFlags(...), with no logging. (internal/exec/runner.go: NewRunner, RunStep)

Errors lose high-value context (step, command, exit code, duration). On failure, run() returns fmt.Errorf("%w: %v", errors.ErrExecutionFailed, err) (or ctx.Err()), which preserves a category but omits step ID, elapsed time, and process exit code/signal—often the first things needed to diagnose. Evidence: run() error path and wrapping; error taxonomy only maps categories to exit codes. (internal/exec/runner.go: run; internal/errors/errors.go: ExitCode)

Sanitizer warnings are emitted as human text only and are not emitted at “decision boundaries.” Warnings are printed as a single formatted line and accumulated in-memory, but there’s no summary count per step/run, no machine-parseable form, and no “this step was sanitized” header at step start/end. Evidence: recordWarnings prints fmt.Fprintf(... "⚠ oraclepack: sanitized label ...") and appends to r.warnings; DrainWarnings() returns them but nothing here enforces they’re logged/serialized elsewhere. (internal/exec/runner.go: recordWarnings, DrainWarnings)

No run/step outcome metrics exist at the point where the outcome is known. The runner is the natural place to emit: step_success/step_failure counts, durations, and warning counts; currently it emits none. Evidence: run() measures no time and publishes no counters; it only runs the command and returns error/nil. (internal/exec/runner.go: run)

Smallest, highest-leverage additions (minimal surface area):

Add a tiny correlation context passed into run() (at least step_id, optionally run_id) and emit two boundary lines: step.start and step.end (include duration_ms, status, exit_code if available). This alone makes raw stdout/stderr interpretable. (Hook point: RunPrelude/RunStep just before/after r.run(...); data already present: s.ID in RunStep.) (internal/exec/runner.go: RunStep, run)

On error, detect *exec.ExitError and include exit_code/`s_*
```

docs/oracle-questions-2026-01-07/11-permissions-model.md
```
Direct answer (1–8 bullets, evidence-cited)

No first-class permission model is evident in the provided code: none of the shown entrypoints/config/state types include a notion of user identity, roles, scopes, ACLs, or authorization checks (e.g., the CLI root command only defines local flags like noTUI, oracleBin, outDir). Evidence: internal/cli/root.go (flag surface only; no auth/identity inputs shown).

Execution appears to be “whoever can run the CLI can do everything the CLI does.” The runner shells out to an external binary and interacts with the local filesystem; no gating by role/scope is shown. Evidence: internal/exec/runner.go (uses os/exec + os without any authz hooks shown).

“Who can do what” is therefore effectively determined by OS-level permissions and environment constraints, not by application-level RBAC/ACL logic. Evidence: internal/app/app.go (config is paths like PackPath/StatePath/Report…; no policy/role configuration shown) + internal/app/run.go (reads/writes JSON/state/report to local paths; no authz shown).

Any real access control is likely delegated to external dependencies (e.g., the invoked oracle binary and filesystem permissions), not implemented in oraclepack itself, based on the visible integration boundary. Evidence: internal/exec/runner.go (external command invocation) + internal/report/generate.go (report generation from run state; no authz metadata).

Minimal set of files that encode “who can do what” (based on what’s been provided)

None (no app-level permission model is shown in the provided files). Practically, behavior is governed by:

OS user permissions (read/write/execute) for the invoked binary and directories (PackPath, StatePath, report/output dirs). Evidence: internal/app/app.go, internal/app/run.go.

External tool behavior (the oracle binary). Evidence: internal/exec/runner.go.

Risks/unknowns (bullets)

The repository may contain an unshared permissions/auth subsystem (e.g., internal/auth/*, HTTP server middleware, policy engine) that wasn’t included in the snippets.

The external oracle binary may implement its own authn/authz, which would materially change the true permission model even if oraclepack itself has none.

There may be implicit “permissions” via overrides/config resolution (e.g., internal/overrides) that constrain execution; not enough code was provided to verify.

Next smallest concrete experiment (exactly one action)

Run a repo-wide search for authorization primitives: rg -n "(rbac|role|roles|scope|scopes|acl|permission|authorize|authn|authz|oauth|jwt|policy|casbin)" .

If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

internal/**/auth*.go

internal/**/perm*.go

internal/**/{rbac,acl,policy}*.go

internal/**/middleware*.go

internal/**/http*/**/*.go and/or internal/**/server*/**/*.go

cmd/** (any additional CLIs/entrypoints beyond internal/cli/root.go)

Any config/schema files that might define roles/scopes (e.g., **/*config*.*, **/*schema*.*, **/*policy*.*)
```

docs/oracle-questions-2026-01-07/12-permissions-enforcement.md
```
1) Direct answer (1–8 bullets, evidence-cited)

There is no explicit permission enforcement layer in the surfaced code (no authn/authz context, no roles/scopes/ACL objects, no guards/middleware imports). The CLI entrypoint is Cobra-based with flag globals (noTUI, oracleBin, outDir) and does not introduce any identity/permission concept. Evidence: internal/cli/root.go snippet shows only Cobra + basic flags; no auth-related imports/types.

Effective “permissions” are the OS user’s privileges at runtime. The system executes external commands via os/exec, so whatever the current process user can do (filesystem, network, process execution) the tool can do. Evidence: internal/exec/runner.go imports and uses os/exec and os.

Primary enforcement chokepoint is the execution boundary (Runner) — but it enforces execution, not authorization. If you need “who can run what,” the only practical choke is where commands are spawned (and where env/args are assembled). Evidence: internal/exec/runner.go is the dedicated Runner module and pulls overrides.

Pack parsing enables embedding runnable shell content; there is no shown policy layer restricting it. The parser extracts bash fenced blocks (regex), which strongly implies step payloads can contain shell scripts. **Evidence:** `internal/pack/parser.go` defines `bashFenceRegex = ... bash\n(.*?)\n``` ...`.

State/report generation has no permission gates; it simply materializes outputs. So any caller who can run the CLI can cause state/report writes wherever configured. Evidence: internal/report/generate.go constructs a report from RunState; internal/app/app.go config includes StatePath/Report… fields (from snippet) with no auth checks.

Likely bypass risks are “configuration-as-authority” rather than missing middleware. If a caller can control oracleBin, pack content, overrides, output/state paths, they effectively control what runs and where it writes—without an authorization decision point. Evidence: internal/cli/root.go exposes oracleBin/outDir; internal/exec/runner.go consumes overrides + executes commands; internal/pack/parser.go extracts bash payloads.

Inconsistent pattern: business logic lives in internal/app/*, but enforcement is absent at service-layer boundaries. There’s no “service-layer check” because there’s no identity/claims passed through app -> exec -> filesystem; everything is implicitly trusted input. Evidence: internal/app/app.go wires exec, pack, state with no auth context types.

2) Risks/unknowns (bullets)

Unknown whether any permission checks exist elsewhere (e.g., an HTTP server, MCP server wrapper, or API layer not included in the snippets). If those exist, they may be the real enforcement point.

Unknown how overrides are applied (env var allow/deny lists, argument escaping, path validation). internal/overrides/* was referenced but not provided.

Unknown whether pack execution is sandboxed (chroot, containers, seccomp, allowlisted commands). Nothing in the provided snippets suggests sandboxing, but absence of evidence isn’t proof.

Unknown whether state/output paths are validated (path traversal, writing outside workspace, overwriting). Config shows paths, but validation code wasn’t included.

3) Next smallest concrete experiment (exactly one action)

Run a repo-wide search to locate any authz/authn/permission enforcement and all chokepoints where requests/commands enter:

bash
Copy code
rg -n --hidden --no-ignore-vcs "(auth|authz|authorize|authorization|permission|rbac|acl|role|scope|policy|guard|middleware|claims|principal|subject|tenant)" .

4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

Attach any of the following if they exist (these are the most likely places where enforcement would live, and we haven’
```

docs/oracle-questions-2026-01-07/13-migrations-schema.md
```
1) Direct answer (1–8 bullets, evidence-cited)

State schema has an explicit version field but no migration logic is implemented. RunState includes SchemaVersion int \json:"schema_version"`ininternal/state/types.go, but internal/state/persist.go::LoadStatesimplyjson.Unmarshals into RunStateand returns it without checking or upgradingSchemaVersion`.

Evidence: internal/state/types.go (RunState.SchemaVersion) and internal/state/persist.go (LoadState unmarshal + return; no branching on version).

State evolution currently relies on Go JSON’s “additive compatibility” behavior. New optional fields like ROIThreshold, ROIMode, and Warnings are tagged omitempty in internal/state/types.go, meaning:

older state files missing these fields will still load (they default to zero values),

newer state files with extra fields would be ignored by older binaries (Go encoding/json ignores unknown fields by default).

Evidence: internal/state/types.go JSON tags on ROIThreshold, ROIMode, Warnings.

There is no tooling or “migration runner” in the shown code paths. The persistence layer only provides atomic write + read (SaveStateAtomic, LoadState) and does not do version stamping, backfills, or transforms.

Evidence: internal/state/persist.go contains only marshal/write/rename and read/unmarshal.

Practical implication: only additive changes are safe; renames/removals are breaking. Because no upgrade step exists, any future schema changes that rename/remove fields, change meaning, or change requiredness will either silently degrade (defaults) or break logic downstream.

Where migrations would be applied (natural chokepoint): internal/state/persist.go::LoadState (or the caller immediately after it) is the obvious place to:

read → unmarshal

check SchemaVersion against a CurrentSchemaVersion constant

run sequential migrations migrateV1toV2, migrateV2toV3, etc.

set SchemaVersion = CurrentSchemaVersion

optionally rewrite the upgraded file via SaveStateAtomic

Evidence for chokepoint existence: LoadState/SaveStateAtomic are the only provided persistence entry points in internal/state/persist.go.

Report format evolution cannot be confirmed from the provided snippets. Earlier context shows GenerateReport returns a *ReportV1, which suggests intent to version report structures by type name, but there’s no evidence (in the snippets here) of report persistence, report version fields, or report migrations.

Evidence: internal/report/generate.go shows GenerateReport(...) *ReportV1 (from prior snippet), but no internal/report/*.go definitions/persistence were provided here.

2) Risks/unknowns

Unknown whether SchemaVersion is ever set/non-zero (e.g., on initial state creation), and what the “current” version is meant to be (no constant or assignment shown).

Unknown whether there is any hidden migration handling in the caller (likely internal/app/run.go or similar) before/after LoadState.

Unknown report persistence/read path (where report files are written/read, whether a version field exists, whether multiple report versions are supported).

Unknown whether pack/config/overrides have versioning (separate from state schema; not shown in the provided files).

3) Next smallest concrete experiment (exactly one action)

Run and paste the output of:
rg -n "SchemaVersion|schema_version|CurrentSchema|migrat(e|ion)|ReportV[0-9]|report_version" internal/

4) If evidence is insufficient, exact missing file/path pattern(s) to attach next

internal/report/*.go (especially the file defining ReportV1 and any report write/read functions)

internal/app/run.go (or whichever file calls state.LoadState / state.SaveStateAtomic, and whi
```

docs/oracle-questions-2026-01-07/14-migrations-compat.md
```
1) Direct answer (backward/forward compatibility expectations; where ensured vs risky)

Run state is explicitly versioned; expectation is “migrate-by-version when needed.” RunState carries SchemaVersion int \json:"schema_version"`which is the primary hook for backward/forward compatibility across persisted state formats. Evidence:internal/state/types.go (RunState.SchemaVersion`).

State schema is designed to be additive (backward-compatible) via optional fields. Fields like ROIThreshold, ROIMode, and Warnings are tagged omitempty, implying they can be absent in older saved state and should default safely in newer code. Evidence: internal/state/types.go (ROIThreshold ...omitempty, ROIMode ...omitempty, Warnings ...omitempty).

Forward compatibility for state/report within Go likely relies on default encoding/json behavior (unknown fields ignored). The structs are plain JSON-tagged types with no visible strict decoding controls here; in Go, older binaries typically ignore new fields when unmarshalling (unless DisallowUnknownFields is used elsewhere). Evidence: presence of standard JSON tags across structs in internal/state/types.go and internal/report/types.go (no strictness shown).

Report schema is implicitly versioned by type name, not by an on-the-wire version field (riskier for evolvability). The report root type is ReportV1, but there is no schema_version/report_version JSON field on ReportV1. This suggests compatibility depends on file naming/conventions or external knowledge, not self-describing payloads. Evidence: internal/report/types.go (type ReportV1 struct { ... } with no version field).

Report is partially “consumer-friendly” via redundant duration fields, suggesting an evolution/compat strategy. Summary.TotalDuration (a time.Duration) is paired with TotalDurationMs int64, and StepReport.Duration is paired with DurationMs int64. This looks like an intentional compatibility measure for consumers that can’t/won’t parse Go time.Duration encoding. Evidence: internal/report/types.go (TotalDuration + TotalDurationMs, Duration + DurationMs).

Two distinct Warning types exist (state vs report), duplicating schema; compatibility depends on keeping them in lockstep. Both packages define Warning with the same JSON fields; this is currently compatible, but any future divergence becomes a schema drift risk. Evidence: internal/state/types.go type Warning struct { Scope, StepID, Line, Token, Message } and internal/report/types.go type Warning struct { ...same fields... }.

Map-based step status storage is tolerant of step additions but sensitive to step ID churn. RunState.StepStatuses map[string]StepStatus supports new step IDs appearing over time, but renames/removals can break continuity of resumed/inspected runs unless migration logic remaps keys. Evidence: internal/state/types.go (StepStatuses map[string]StepStatus).

2) Risks/unknowns

Unknown whether reads/writes are strict or tolerant. If any loader uses json.Decoder.DisallowUnknownFields(), forward compatibility (new fields read by old binaries) will break. Missing evidence: state/report load code.

No visible migration dispatcher. SchemaVersion exists, but we don’t see code that applies migrations on read (or stamps versions on write), so rollback/upgrade safety is unproven from the provided files alone.

Report payload is not self-describing. Without a report_version, consumers can’t robustly branch on schema, making forward/backward compatibility brittle as ReportV2 emerges.

G
```

docs/oracle-questions-2026-01-07/15-ux-flows-primary.md
```
1) Direct answer (primary TUI user/operator flows, mapped to components + key state transitions)

App entry → TUI boot (CLI “run” command): internal/cli/run.go loads/initializes app (pack, runner, state), then constructs tui.Model via tui.NewModel(a.Pack, a.Runner, a.State, cfg.StatePath, cfg.ROIThreshold, cfg.ROIMode, runAll) and runs Bubble Tea (tea.NewProgram(..., tea.WithAltScreen())). The --run-all flag is passed as autoRun into the TUI model. (internal/cli/run.go: runCmd.RunE; internal/tui/tui.go: NewModel, Init)

Browse/select steps (default “Steps” screen): Primary UI is list.Model (left) + viewport.Model (right preview). Selection changes update preview content via setListPreviewContent(newID) when not running and viewState == ViewSteps. (internal/tui/tui.go: Model.View, Update list handling, setListPreviewContent, selectedItemID)

Run a single selected step (“Enter”): From ViewSteps, pressing Enter sets running=true, viewState=ViewRunning, runAll=false, clears logLines, and triggers runStep(stepID) + waitForLogs() + spinner.Tick. Completion arrives as FinishedMsg{Err,...}; success sets running=false, viewState=ViewDone; failure sets err, running=false, runAll=false, viewState=ViewDone. (internal/tui/tui.go: Update key "enter", runStep, waitForLogs, FinishedMsg handler)

Run all filtered steps sequentially (“a”) + auto-run on init: Pressing a in ViewSteps (when not running) sets running=true, viewState=ViewRunning, runAll=true, currentIdx=0, selects first filtered list item, then executes runStep(firstID). On each FinishedMsg success while runAll=true, it increments currentIdx, selects next item, and calls runStep(nextID) until exhausted, then sets running=false, runAll=false, viewState=ViewDone. Auto-run uses the same path via StartAutoRunMsg emitted from Init() when autoRun is true. (internal/tui/tui.go: Update key "a", StartAutoRunMsg handling, FinishedMsg sequential logic; Init)

ROI filtering flow (“f” and “m”):

f enters filter input modal: isFiltering=true, filterInput.Focus(), view switches to centered modal via View() (still viewState remains Steps, but UI is driven by isFiltering). Enter parses float, updates roiThreshold, calls refreshList(), and persists into state (m.state.ROIThreshold, m.state.ROIMode) via state.SaveStateAtomic. Esc cancels.

m toggles roiMode between "under" and "over", refreshes list, persists state.
This directly changes what “run all” iterates over because it uses m.list.Items() (filtered). (internal/tui/tui.go: Update keys "f", "m", refreshList, persistFilterState, View modal branches)

Step preview (“v”) with wrap/copy/back: v sets previewID, viewState=ViewStepPreview, resets viewport offset, and renders Markdown (render.RenderMarkdown) into the viewport. In preview: t toggles wrap (previewWrap) and re-renders; c copies step plain text (clipboard fallback writes to file); b/esc returns to ViewSteps and restores list preview. (internal/tui/tui.go: Update key "v", `ViewSte
```

docs/oracle-questions-2026-01-07/16-ux-flows-edgecases.md
```
1) Direct answer (top edge cases / gotchas, where handled vs missing)

Cancel is “hard cancel” everywhere, with no confirm and no “go back one step” on Esc: OverridesFlowModel.Update maps "esc" directly to OverridesCancelledMsg{} regardless of wizard step; there’s no “Are you sure?” or “Back” on Esc. (Handled: cancel exit path exists; Missing: confirm/discard + step-back on Esc.) (OverridesFlowModel.Update: case "esc": ... OverridesCancelledMsg{})

Back navigation conflicts with normal text editing / list filtering: the parent model binds "backspace" (and "shift+tab") to step-back (m.step--) whenever m.step > OverridesFlags. If StepsPickerModel (and/or confirm UI) uses inputs or list filtering, Backspace is commonly needed for editing—this will unexpectedly navigate back instead. (Missing: scoped key handling / focus-aware routing.) (OverridesFlowModel.Update: case "shift+tab", "backspace": if m.step > OverridesFlags { m.step-- })

Next navigation conflicts with child controls: the parent captures "enter" and "tab" to advance steps (or trigger validation on confirm). If the list widgets use Enter/Tab for selection, filtering, or focus, those interactions can be overridden by step progression. (Missing: “only advance when not filtering / not editing” guard, or explicit Next button/key distinct from list controls.) (OverridesFlowModel.Update: case "enter", "tab": ... m.step++)

Validation failures are displayed, but recovery UX is minimal: confirm step handles ValidationResultMsg by setting confirm.validating=false, storing confirm.errors, and setting a short errMsg (“N validation errors detected.”). It does not automatically move the user back to the step that likely caused the errors, nor does it highlight which selection caused which validation error (depends on OverridesConfirmModel.View, not shown). (Handled: error surfaced; Missing: guided remediation / jump-back.) (OverridesFlowModel.Update: case ValidationResultMsg: ... m.confirm.errors = v.Errors ... errMsg ... return m, nil)

Validation can hang indefinitely (no timeout/cancel), making the UI feel frozen: validateCmd uses context.Background() and calls exec.ValidateOverrides(...) with no deadline and no cancel path. While confirm.validating prevents duplicate submissions, there’s no way to abort a stuck validation. (Missing: context timeout + Esc to cancel validation.) (validateCmd: exec.ValidateOverrides(context.Background(), ...))

Baseline flags appear incorrectly wired into the flags picker (likely a functional bug): NewOverridesFlowModel constructs flags: NewFlagsPickerModel(nil) even though it has a baseline argument and separately stores baseline: exec.ApplyChatGPTURL(baseline, ...). This means flags that are already in baseline won’t be marked (base) / [*], won’t auto-check, and won’t be protected from toggling. (Missing: pass the baseline into NewFlagsPickerModel(baseline).) (NewOverridesFlowModel: NewFlagsPickerModel(nil) vs stored baseline: ...)

Users cannot remove baseline flags (only add): currentOverrides() sets RemovedFlags: nil always, and FlagsPickerModel.SelectedFlags() only returns selected flags where !fi.IsBaseline. Even if you wanted removal, there’s no path in the
```

docs/oracle-questions-2026-01-07/17-failure-modes-taxonomy.md
```
1) Direct answer (failure-mode taxonomy; where classified/handled; what’s surfaced)

Precondition / lifecycle failures (internal “programmer error” / wrong call order): RunPlain hard-fails if a.Pack == nil or a.State == nil with plain fmt.Errorf("pack not loaded") / "state not loaded". These are not classified into internal/errors sentinel types, so callers will likely treat them as generic failures. (Evidence: internal/app/run.go:RunPlain)

Prelude execution failures (treated as fatal): If Pack.Prelude.Code != "", it runs a.Runner.RunPrelude(...); on error it records warnings, then returns fmt.Errorf("prelude failed: %w", err) without writing a report (no finalize() call on this path). (Evidence: internal/app/run.go:RunPlain)

Step execution failures (classified at the step-status layer, optionally fatal at the run layer):

Before executing a step, state is set to StatusRunning with StartedAt, then persisted via a.saveState().

On RunStep error: step status becomes StatusFailed, status.Error = err.Error(), state is persisted, and:

if StopOnFail is true: finalize() runs and RunPlain returns the underlying err;

otherwise it continues to the next step.
This is the core “execution failed” mode in practice. (Evidence: internal/app/run.go:RunPlain)

“Stop-on-fail” semantics create two user-visible failure shapes: with StopOnFail=true, the run returns immediately on the first failed step (after writing the report); with StopOnFail=false, the run returns nil even if one or more steps failed, relying on state/report for operators to discover failures. (Evidence: internal/app/run.go:RunPlain)

Skips are treated as non-fail outcomes, but are not consistently represented in state:

ROI-based skips print Skipping step... and continue without writing any step status (no StatusSkipped assignment here).

Resume-based skips require an existing StatusSuccess entry; they print and continue leaving the prior success intact.
Result: reports/operators may not be able to distinguish “intentionally skipped by ROI filter” vs “never attempted / missing status” unless report generation infers it elsewhere. (Evidence: internal/app/run.go:RunPlain)

Warnings are a parallel channel, not failures: recordWarnings() drains runner warnings and appends them to a.State.Warnings, then persists state. They are not printed directly by App (beyond whatever the runner prints), so visibility is primarily via persisted state/report. (Evidence: internal/app/run.go:recordWarnings, saveState)

State persistence failures are explicitly ignored (silent data-loss risk): saveState() discards errors from state.SaveStateAtomic(...) via _ = .... If state writes fail, the run proceeds and operators may later see stale/missing progress. (Evidence: internal/app/run.go:saveState)

Report generation / write failures are explicitly ignored, but success is always announced: finalize() discards errors from json.MarshalIndent(...) and os.WriteFile(...), yet always prints Report written to ... when ReportPath != "". This can falsely signal success to operators even when no report was produced. (Evidence: internal/app/run.go:finalize)

System-wide error classification exists, but RunPlain largely relies on upstream to use it: internal/errors/errors.go defines sentinel errors (ErrInvalidPack, ErrExecutionFailed, ErrConfigInvalid) and an ExitCode(err) mapping using `errors.Is(
```

docs/oracle-questions-2026-01-07/18-failure-modes-resilience.md
```
1) Direct answer (resilience mechanisms; gaps; where configured)

Script “label line” sanitization exists and is applied to both preludes and steps. SanitizeScript converts single-token, label-shaped lines into echo "<token>" to avoid accidental command execution (e.g., a bare Setup line). This is invoked from Runner.RunPrelude and Runner.RunStep before execution. (internal/exec/sanitize.go: SanitizeScript; internal/exec/runner.go: RunPrelude, RunStep)

Sanitization is deliberately scoped to reduce false positives. It skips blank lines and comments (#...), skips anything inside heredocs (tracks start via heredocStartToken and ignores lines until the terminator), and only rewrites lines where strings.Fields(trimmed) yields exactly 1 field. (internal/exec/sanitize.go: SanitizeScript, heredocStartToken)

Sanitization avoids rewriting real shell constructs and real executables. Tokens are rewritten only if they match labelTokenRegex (^[A-Za-z][A-Za-z0-9_-]*$), are not shell builtins/keywords, and are not found on $PATH (osexec.LookPath(token) must fail). (internal/exec/sanitize.go: labelTokenRegex, shellBuiltins, shellKeywords, LookPath check)

Sanitization produces explicit warnings and they are surfaced at least to the run log stream. Warnings are stored on the runner and also written to logWriter as ⚠ oraclepack: sanitized label ..., and can be retrieved via DrainWarnings(). (internal/exec/runner.go: recordWarnings, DrainWarnings)

Execution is cancellation-aware via context. exec.CommandContext(ctx, ...) is used; on failure, if ctx.Err() != nil it returns the context error directly. This is a resilience mechanism for user cancel / deadline propagation. (internal/exec/runner.go: run)

Execution errors are wrapped with a sentinel classification. Non-context execution errors return fmt.Errorf("%w: %v", errors.ErrExecutionFailed, err), enabling upstream code to recognize “execution failed” as a class. (internal/exec/runner.go: run)

Configurable execution knobs exist, but they’re mostly “wiring,” not resilience policies. RunnerOptions configures Shell (default /bin/bash), WorkDir, Env (appended to os.Environ()), OracleFlags, Overrides, and ChatGPTURL. These influence behavior but do not add retries/limits/strict-mode by themselves. (internal/exec/runner.go: RunnerOptions, NewRunner)

Critical path gap: “stop-on-fail” within a script is not enforced. The runner executes bash -lc <script> without forcing set -e / -euo pipefail. So a failing command earlier in a script may not fail the step unless the script author ensured it (depends on the script). This is a major resilience gap for “partial failure looks like success.” (internal/exec/runner.go: run uses r.Shell, "-lc", script)

Critical path gap: sanitization misses common “label + trailing comment” patterns. Because only len(fields) == 1 is sanitized, a line like Setup # do thing becomes two fields and is not rewritten; bash will attempt to execute Setup and fail (or worse, execute if Setup exists on PATH). (internal/exec/sanitize.go: Fields/len(fields) != 1 early-continue)

Critical path gap: environment
```

docs/oracle-questions-2026-01-07/19-feature-flags-inventory.md
```
Direct answer (inventory + where defined/evaluated/documented)

Runtime “flag override” mechanism exists via RuntimeOverrides: you can append flags, remove flags, and inject a ChatGPT URL (as --chatgpt-url) at runtime. Defined in internal/overrides/types.go as AddedFlags, RemovedFlags, and ChatGPTURL. (internal/overrides/types.go: RuntimeOverrides)

Per-step targeting exists via ApplyToSteps: a set keyed by step ID that determines which steps the overrides apply to. Defined in internal/overrides/types.go as ApplyToSteps map[string]bool with comment indicating “If empty, applies to none.” (internal/overrides/types.go: RuntimeOverrides.ApplyToSteps)

TUI exposes a curated list of toggleable “Oracle Flags” (feature-flag-like user controls):
--files-report, --render, --render-plain, --copy, --wait. Defined in internal/tui/overrides_flags.go in curated := []FlagItem{...}. (internal/tui/overrides_flags.go: curated)

Baseline vs. user-selected behavior is enforced: baseline flags are pre-selected and cannot be toggled off; only non-baseline flags can be toggled with spacebar. Defined/evaluated in NewFlagsPickerModel (baseline set, IsBaseline, initial Selected) and Update (toggle only if !item.IsBaseline). (internal/tui/overrides_flags.go: NewFlagsPickerModel, Update)

UI-level “documentation” exists for these flags via the Desc field displayed alongside each curated flag (e.g., “Show per-file token usage”, “Print assembled markdown bundle”). Implemented in flagsDelegate.Render. (internal/tui/overrides_flags.go: FlagItem.Desc, Render)

Discoverability control exists in the TUI via filtering: the list enables filtering (l.SetFilteringEnabled(true)), which is a UX control for narrowing/selecting flags (not ROI, but “runtime control”). (internal/tui/overrides_flags.go: SetFilteringEnabled(true))

Selection output is explicitly “delta-only”: SelectedFlags() returns only flags that are selected and not baseline, i.e., the override payload is “changes from baseline,” which is consistent with safe rollout semantics. (internal/tui/overrides_flags.go: SelectedFlags)

Risks/unknowns

Where overrides are actually applied/evaluated is not shown in the provided snippets (i.e., where AddedFlags/RemovedFlags/ChatGPTURL/ApplyToSteps mutate the final oracle CLI invocation per step).

ROI filtering is not evidenced in the provided files; if it exists, it’s elsewhere (selection of steps by ROI, filtering packs/steps, etc.).

Documentation scope is currently “in-TUI only” for curated flags; no evidence here of CLI --help docs, README references, or pack schema docs for these controls.

Next smallest concrete experiment (exactly one action)

Run: rg -n "RuntimeOverrides|AddedFlags|RemovedFlags|ChatGPTURL|ApplyToSteps|--chatgpt-url|ROI|roi" internal/ and attach the matching snippets (especially any code that builds the exec command / selects steps).

If evidence is insufficient, attach these missing file/path pattern(s) next

internal/overrides/*.go (anything beyond types.go, especially functions that apply overrides)

The portion(s) of internal/exec/runner.go (or related exec builder) that incorporate overrides into the command/args

Any step-selection or filtering logic files, likely under: internal/app/*.go, internal/pack/*.go, internal/tui/*flow*.go (for ROI or step filtering)
```

docs/oracle-questions-2026-01-07/20-feature-flags-rollout.md
```
## 1) Direct answer (flag/rollout lifecycle, evidence-cited)

* **Selection is an explicit, user-supplied “diff” over baseline flags**: runtime overrides are modeled as `AddedFlags`, `RemovedFlags`, and an optional `ChatGPTURL` injection (`--chatgpt-url <url>`). Evidence: `internal/overrides/types.go` (`RuntimeOverrides` fields) and `internal/overrides/merge.go` (injection logic in `EffectiveFlags`).
* **Targeting is allowlist-by-step and defaults to “off”**: overrides apply only when `ApplyToSteps[stepID] == true`; if `ApplyToSteps` is `nil` (or the step key is absent/false), the baseline flags are returned unchanged. Evidence: `internal/overrides/merge.go` early-return guard in `EffectiveFlags`; `internal/overrides/types.go` comment “If empty, applies to none.”
* **Application happens at flag-materialization time, per step**: the lifecycle is “start from baseline → remove exact-match flags → append added flags → optionally append `--chatgpt-url <ChatGPTURL>`”. Evidence: `internal/overrides/merge.go` (`removed` map filter, then `append(r.AddedFlags...)`, then URL injection).
* **Validation is not implemented in the shown override module**: there is no validation method or error path in `RuntimeOverrides` or `EffectiveFlags`; invalid flags, conflicting flags, malformed URLs, or nonexistent step IDs are not rejected here. Evidence: `internal/overrides/types.go` contains only fields; `internal/overrides/merge.go` is pure transformation with no errors/validation.
* **Retirement is purely process-driven with current code**: there is no TTL/expiry, metadata, audit, or “unused override” detection in the override types/merge logic. Evidence: `internal/overrides/types.go` has no lifecycle metadata; `internal/overrides/merge.go` has no tracking/telemetry hooks.

**Minimum guardrails to prevent “flag debt” (smallest set that actually closes the main failure modes):**

* **Central validation gate** before any run: validate (a) every `ApplyToSteps` key exists in the pack’s step IDs, (b) flags conform to an allowlisted pattern (e.g., `^--[a-z0-9-]+(=.*)?$`), (c) no duplicates/conflicts after merge (e.g., two `--model` variants), and (d) `ChatGPTURL` is a valid URL. (Missing today in `types.go`/`merge.go`.)
* **Add lifecycle metadata** to each override bundle: `owner`, `reason`, `created_at`, `expires_at` (or `ttl_days`), and `issue/link`. Enforce “expired overrides fail closed.” (Not present today in `RuntimeOverrides`.)
* **Auditability**: emit a structured record per step of `{baseline_flags, removed, added, url_injected, effective_flags}` so operators can see what actually happened and spot stale overrides. (Nothing in `EffectiveFlags` records decisions.)
* **Retirement enforcement**: CI/test-time lints that fail if (a) overrides have expired, (b) an override hasn’t been applied to any step in N runs, or (c) overrides reference removed steps/flags. (No mechanism exists in the shown code; must be added around it.)
* **Prefer semantic removal** (optional but high leverage): today removal is exact-string match (`removed[flag]`), which doesn’t catch `--flag=value` vs `--flag` variants and encourages drift/debt. Either normalize flags or support key-based removal. Evidence of exact-match removal: `internal/overrides/merge.go` baseline filter uses `removed[flag]` only.

## 2) Risks/unknowns

* **Where overrides are created/loaded/persisted is not shown**, so it’s unclear whether any validation exists upstream (TUI wizard, config loader, etc.) or whether overrides can be injected via environment/CLI.
* **No evidence of observability/logging around overrides**: if something breaks due to an override, operators may not be able to quickly attribute the effective flag set to the override bundle.
* **Exact-match removal is brittle** and can lead to long-lived “ineffective” `RemovedFlags` entries (flag debt), because they silently don’t remove anything when baseline formatting changes. Evidence: `internal/overrides/merge.go`.
* **No conflict resolution**: if baseline already includes a flag that `AddedFlags` adds again, the effective list can contain duplicates, and downstream flag parsing semantics may become ambiguous. Evidence: `internal/overrides/merge.go` simply appends.

## 3) Next smallest concrete experiment (exactly one action)

* **Add `(*RuntimeOverrides) Validate(stepIDs []string) error` and call it once before execution begins**, rejecting unknown step IDs, malformed flags, duplicates/conflicts, and invalid `ChatGPTURL`; include a unit test table for the edge cases above.

## 4) If evidence is insufficient, exact missing file/path pattern(s) to attach next

* `internal/tui/overrides_flow.go` (how users select/serialize overrides)
* `internal/exec/runner.go` (where `EffectiveFlags` is invoked and how effective flags are executed/logged)
* Any override config IO paths, e.g. `internal/overrides/*.go` beyond `types.go`/`merge.go`, and any config structs/files such as `internal/app/*.go`, `internal/pack/*.go`, or `internal/config/*.go` that load/store runtime overrides
```

internal/exec/oracle_scan.go
```
package exec

import (
	"regexp"
	"strings"
)

var oracleCmdRegex = regexp.MustCompile(`^(\s*)(oracle)\b`)

// OracleInvocation represents a detected oracle command in a script.
type OracleInvocation struct {
	StartLine   int    // 0-based start line index
	EndLine     int    // 0-based end line index (inclusive)
	Raw         string // The full command string (joined if multi-line)
	Display     string // A trimmed version for UI display
	Indentation string // The leading whitespace
}

// ExtractOracleInvocations extracts oracle invocations from a script.
func ExtractOracleInvocations(script string) []OracleInvocation {
	var invocations []OracleInvocation
	lines := strings.Split(script, "\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)

		// Skip comments
		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		// Check for oracle command
		loc := oracleCmdRegex.FindStringSubmatchIndex(line)
		if loc != nil {
			startLine := i
			// Group 1 is the indentation
			indentation := line[loc[2]:loc[3]]

			var cmdBuilder strings.Builder
			cmdBuilder.WriteString(line)

			endLine := i
			// Handle line continuations
			// Check if line ends with backslash (ignoring trailing whitespace)
			for {
				if endLine+1 >= len(lines) {
					break
				}

				// Check current line for continuation
				currTrimmed := strings.TrimRight(lines[endLine], " \t")
				if !strings.HasSuffix(currTrimmed, "\\") {
					break
				}

				endLine++
				cmdBuilder.WriteString("\n")
				cmdBuilder.WriteString(lines[endLine])
			}

			raw := cmdBuilder.String()
			invocations = append(invocations, OracleInvocation{
				StartLine:   startLine,
				EndLine:     endLine,
				Raw:         raw,
				Display:     strings.TrimSpace(raw),
				Indentation: indentation,
			})

			i = endLine // Advance loop
		}
	}
	return invocations
}
```

internal/exec/oracle_scan_test.go
```
package exec

import (
	"reflect"
	"testing"
)

func TestExtractOracleInvocations(t *testing.T) {
	tests := []struct {
		name   string
		script string
		want   []OracleInvocation
	}{
		{
			name:   "Simple command",
			script: "oracle --json",
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 0, Raw: "oracle --json", Display: "oracle --json", Indentation: ""},
			},
		},
		{
			name:   "Indented command",
			script: "  oracle --json",
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 0, Raw: "  oracle --json", Display: "oracle --json", Indentation: "  "},
			},
		},
		{
			name: "Multiline command",
			script: `oracle \
  --json \
  --files`,
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 2, Raw: `oracle \
  --json \
  --files`, Display: `oracle \
  --json \
  --files`, Indentation: ""},
			},
		},
		{
			name: "Commented command",
			script: `# oracle --json
oracle --real`,
			want: []OracleInvocation{
				{StartLine: 1, EndLine: 1, Raw: "oracle --real", Display: "oracle --real", Indentation: ""},
			},
		},
		{
			name: "Multiple commands",
			script: `
echo start
oracle --one
echo mid
oracle --two
echo end
`,
			want: []OracleInvocation{
				{StartLine: 2, EndLine: 2, Raw: "oracle --one", Display: "oracle --one", Indentation: ""},
				{StartLine: 4, EndLine: 4, Raw: "oracle --two", Display: "oracle --two", Indentation: ""},
			},
		},
		{
			name:   "Oraclepack prefix (should not match)",
			script: "oraclepack run",
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractOracleInvocations(tt.script)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractOracleInvocations() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
```

internal/exec/oracle_validate.go
```
package exec

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/pack"
)

// ValidationError captures a failed oracle validation for a step.
type ValidationError struct {
	StepID       string
	Command      string
	ErrorMessage string
}

// ValidateOverrides runs oracle --dry-run summary for targeted steps.
func ValidateOverrides(
	ctx context.Context,
	steps []pack.Step,
	over *overrides.RuntimeOverrides,
	baseline []string,
	opts RunnerOptions,
) ([]ValidationError, error) {
	if over == nil || over.ApplyToSteps == nil {
		return nil, nil
	}

	shell := opts.Shell
	if shell == "" {
		shell = "/bin/bash"
	}
	env := append(os.Environ(), opts.Env...)

	var results []ValidationError
	for _, step := range steps {
		if !over.ApplyToSteps[step.ID] {
			continue
		}

		invocations := ExtractOracleInvocations(step.Code)
		if len(invocations) == 0 {
			continue
		}

		flags := over.EffectiveFlags(step.ID, baseline)
		flags = append(flags, "--dry-run", "summary")

		for _, inv := range invocations {
			cmdStr := InjectFlags(inv.Raw, flags)
			msg, err := execDryRun(ctx, shell, opts.WorkDir, env, cmdStr)
			if err == nil {
				continue
			}

			results = append(results, ValidationError{
				StepID:       step.ID,
				Command:      cmdStr,
				ErrorMessage: msg,
			})
		}
	}

	return results, nil
}

func execDryRun(ctx context.Context, shell, workDir string, env []string, command string) (string, error) {
	if pathVal := findEnvValue(env, "PATH"); pathVal != "" {
		command = "export PATH=" + shellQuote(pathVal) + "; " + command
	}

	cmd := exec.CommandContext(ctx, shell, "-lc", command)
	if workDir != "" {
		cmd.Dir = workDir
	}
	cmd.Env = env

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err == nil {
		return stdout.String(), nil
	}
	if stderr.Len() > 0 {
		return strings.TrimSpace(stderr.String()), err
	}
	if stdout.Len() > 0 {
		return strings.TrimSpace(stdout.String()), err
	}
	return err.Error(), err
}

func findEnvValue(env []string, key string) string {
	prefix := key + "="
	for _, entry := range env {
		if strings.HasPrefix(entry, prefix) {
			return strings.TrimPrefix(entry, prefix)
		}
	}
	return ""
}

func shellQuote(value string) string {
	if value == "" {
		return "''"
	}
	return "'" + strings.ReplaceAll(value, "'", "'\\''") + "'"
}
```

internal/exec/oracle_validate_test.go
```
package exec

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/pack"
)

func TestValidateOverrides_Success(t *testing.T) {
	dir := t.TempDir()
	writeOracleStub(t, dir)

	steps := []pack.Step{
		{ID: "01", Code: "oracle --ok"},
	}
	over := &overrides.RuntimeOverrides{
		ApplyToSteps: map[string]bool{"01": true},
	}

	_, err := ValidateOverrides(
		context.Background(),
		steps,
		over,
		[]string{"--base"},
		RunnerOptions{
			WorkDir: dir,
			Env:     []string{"PATH=" + dir + string(os.PathListSeparator) + os.Getenv("PATH")},
		},
	)
	if err != nil {
		t.Fatalf("ValidateOverrides failed: %v", err)
	}
}

func TestValidateOverrides_Error(t *testing.T) {
	dir := t.TempDir()
	writeOracleStub(t, dir)

	steps := []pack.Step{
		{ID: "01", Code: "oracle --bad"},
	}
	over := &overrides.RuntimeOverrides{
		ApplyToSteps: map[string]bool{"01": true},
	}

	errs, err := ValidateOverrides(
		context.Background(),
		steps,
		over,
		nil,
		RunnerOptions{
			WorkDir: dir,
			Env:     []string{"PATH=" + dir + string(os.PathListSeparator) + os.Getenv("PATH")},
		},
	)
	if err != nil {
		t.Fatalf("ValidateOverrides failed: %v", err)
	}
	if len(errs) != 1 {
		t.Fatalf("expected 1 validation error, got %d", len(errs))
	}
	msg := errs[0].ErrorMessage
	if !strings.Contains(msg, "invalid flag") && !strings.Contains(msg, "unknown option") {
		t.Fatalf("unexpected error message: %q", msg)
	}
	if !strings.Contains(errs[0].Command, "--dry-run summary") {
		t.Fatalf("expected command to include --dry-run summary, got %q", errs[0].Command)
	}
}

func writeOracleStub(t *testing.T, dir string) {
	t.Helper()
	stub := `#!/bin/sh
has_dry=0
has_summary=0
for arg in "$@"; do
  if [ "$arg" = "--dry-run" ]; then has_dry=1; fi
  if [ "$arg" = "summary" ]; then has_summary=1; fi
  if [ "$arg" = "--bad" ]; then echo "invalid flag" 1>&2; exit 1; fi
done
if [ $has_dry -eq 0 ] || [ $has_summary -eq 0 ]; then
  echo "missing dry run" 1>&2
  exit 1
fi
exit 0
`
	path := filepath.Join(dir, "oracle")
	if err := os.WriteFile(path, []byte(stub), 0o755); err != nil {
		t.Fatalf("write oracle stub: %v", err)
	}
}
```

.config/mcp/oraclepack-gold-pack/SKILL.md
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

.config/mcp/oraclepack-taskify/SKILL.md
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
[TRUNCATED]
```

.config/skills/oracle-pack/SKILL.md
```
---
name: oracle-pack
description: Generate exactly 20 evidence-cited, ROI-ranked strategist questions for the current repository (12 immediate + 8 strategic), then emit them as runnable @steipete/oracle CLI commands with minimal targeted file attachments and deterministic per-question output paths (a copy/paste-ready “oracle question pack” in the same scratch-style format as oracle-scratch.md).
---

## Quick start

Use this skill when the user wants strategist questions **and** wants to run each question through Oracle as a separate command (so a second model can answer with repo context).

Invocation:

- `$oracle-pack <free-text args>`

Primary output: a Markdown doc whose main content is a **single** fenced `bash` block containing **exactly 20** `oracle ...` commands (copy/paste-ready), in the same “scratch-style” format shown in `references/oracle-scratch-format.md`.

## Inputs

- Repository contents (current working directory).
- Free-text args (may include): `codebase_name`, `constraints`, `non_goals`, `team_size`, `deadline`, plus Oracle pack controls:
  - `out_dir` (default: `oracle-out`)
  - `oracle_cmd` (default: `oracle`)
  - `oracle_flags` (default: `--files-report`)
  - `extra_files` (optional: comma-separated extra `-f/--file` entries to include in *every* command; default: none)

## Deterministic search/tooling rules (must follow)

- Prefer deterministic, non-interactive commands so runs are reproducible.
- Before using `ck` in a session, **always run**: `ck --help` and only rely on flags that `ck --help` confirms.
- Default search tool is `ck`:
  - Conceptual: `ck --sem "<query>"`
  - Literal: `ck --regex "<pattern>"`
  - Best of both: `ck --hybrid "<query>"`
  - For post-processing: prefer `ck --jsonl | jq ... | sort | head`
- Structural search: use `ast-grep` when syntax/shape matters (see `references/inference-first-discovery.md`).
- File finding: prefer `fd` (filename/path/extension filters).

## Args parsing (no follow-ups)

Extract values if present; otherwise set defaults:

- `codebase_name`: `Unknown`
- `constraints`: `None`
- `non_goals`: `None`
- `team_size`: `Unknown`
- `deadline`: `Unknown`
- `out_dir`: `oracle-out`
- `oracle_cmd`: `oracle`
- `oracle_flags`: `--files-report`
- `extra_files`: (empty)

Honor `constraints`. Explicitly exclude `non_goals`.

## Workflow (inference-first discovery → evidence → questions → oracle pack)

### 1) Inference-first discovery (adaptive, codebase-search compliant)

Goal: infer *what this repo is* and *where the truth likely lives* before broad sweeps.

Do this in order:

1. **Discover the search interface (required)**
   - Run: `ck --help`
   - If `ck` indicates it needs setup/indexing, follow the instructions printed by `ck`.

2. **Map the repo surface (cheap, high-signal)**
   - Use `fd` to list likely roots deterministically:
     - `fd . -d 2`
     - `fd -p README.md`
     - `fd -p package.json`
     - `fd -p pyproject.toml`
     - `fd -p go.mod`
     - `fd -p Cargo.toml`
   - Read the smallest set of “index” files that describe structure:
     - README / docs index if present
     - primary manifests/config
     - obvious entrypoint referenced by scripts/manifests

3. **Infer stack + conventions from evidence**
   - From manifests/config, infer runtime + framework signals (dependencies, filenames, directory conventions).
   - From entrypoints, infer wiring patterns (router, DI/container, job scheduler, migration tool).

4. **Derive a “search plan” from inferred signals**
   - Prefer following imports/references from entrypoints into:
     - routing/handlers
     - auth middleware/policies
     - job/queue registration
     - data layer/migrations
     - feature flags config
     - observability bootstrap
   - Use `ck` in the smallest inferred app roots first:
     - Conceptual: `ck --sem "<subsystem intent>"`
     - Hybrid when unsure: `ck --hybrid "<subsystem intent>"`
   - If results are broad/noisy:
     - Narrow to directories/files from the top hits, then re-run `ck`.
     - Use `ast-grep` when you need structure, not text matches.

This improves efficiency vs. hard-targeting a long list of patterns up front: you start with the repo’s own “self-description” and expand only as needed.

### 2) Evidence harvesting (collect anchors before writing questions)

Collect **at least 20 candidate anchors** as one of:

- `{path}:{symbol}` (preferred: nearest function/class/type/handler)
- `{endpoint}` (literal endpoint string)
- `{event}` (job/queue/event name)

For each required category, attempt to identify at least one anchor:

- contracts/interfaces
- invariants
- caching/state
- background jobs
- observability
- permissions
- migrations
- UX flows
- failure modes
- feature flags

Use deterministic selection of candidates:
- If using `ck --jsonl`, post-process to stable top-N:
  - `ck --jsonl ... | jq ... | sort | head -n 20`
- If using plain output, cap deterministically with `head`.

If evidence for a category cannot be found:

- still generate a question for it with reference `Unknown`
- explicitly state which missing artifact pattern would likely provide it (e.g., `**/migrations/**` not found)

### 3) Generate exactly 20 strategist questions (evidence-cited + minimal experiments)

Produce exactly:

- **12** Immediate Exploration
- **8** Strategic Planning

Each question must include:

- **Reference**: `{path}:{symbol}` OR `{endpoint}` OR `{event}` OR `Unknown`
- **Question**: incisive, feasibility-focused, no scope creep
- **Rationale**: exactly one sentence
- **Smallest experiment today**: exactly one concrete action runnable today (targeted `ck` query, targeted `ast-grep` pattern, trace a codepath, add a log line, minimal unit/integration test, run migration dry-run, execute a single endpoint, etc.)

No duplicates (by intent or reference).

### 4) Score and order by near-term ROI (required math + sorting)

For each question compute:

- `ROI = (impact * confidence) / effort`
- `impact`, `confidence`, `effort` ∈ `{0.1..1.0}` (one decimal)

Sort all 20 descending by ROI; break ties by lower effort.

### 5) Convert the 20 questions into an Oracle “question pack” (final deliverable)

For each question (in final sorted order), emit a runnable command using:

- command: `{{oracle_cmd}}` (default `oracle`)
- include `{{oracle_flags}}` (default `--files-report`)
- deterministic output file: `--write-output "<out_dir>/<nn>-<slug>.md"`
  - `<nn>` = zero-padded 2-digit index (01..20)
[TRUNCATED]
```

.config/skills/oracle-pack-rpg-infused/SKILL.md
```
---
name: oracle-pack-rpg-infused
description: Generate a strictly oraclepack-ingestible “oracle strategist question pack” Markdown file under docs/oracle-pack-YYYY-MM-DD.md containing exactly one fenced bash block with exactly 20 numbered oracle commands (01..20), ROI-scored/sorted and category-covered, with an RPG block (WHAT/HOW + dependency edges) embedded in each -p prompt for graph extraction.
---

# Oracle Pack (RPG-infused)

## Quick start

1) Read `assets/oracle-pack-template.md` and keep its structure exactly.
2) Produce exactly one Markdown deliverable and output it verbatim in your final response:
   - First line: `Output file: docs/oracle-pack-YYYY-MM-DD.md`
   - Then the exact Markdown content (no extra commentary, no additional code fences).
3) Ensure the bash block contains exactly 20 steps, numbered `01..20`, each with:
   - One header comment line in the required format
   - Exactly one runnable oracle command line

For the full spec, follow `references/oracle-pack-spec.md`.

## Inputs and defaults

Interpret any user-provided details as “args” and apply these defaults when missing:

- `codebase_name=Unknown`
- `constraints=None`
- `non_goals=None`
- `team_size=Unknown`
- `deadline=Unknown`
- `out_dir=oracle-out`
- `oracle_cmd=oracle`
- `oracle_flags=--files-report`
- `extra_files=empty` (only use if the user provides an explicit mechanism/flag; otherwise omit and keep evidence handling inside the prompt body as text)

If the user supplies their own values, use them verbatim unless they would break the strict output contract.

## Workflow

1) Establish minimal repo understanding (inference-first)
   - Prefer reading index artifacts (e.g., README, top-level config/manifests) before proposing file sweeps.
   - If you cannot access the repo contents, keep `reference=Unknown` and make the “missing artifact pattern(s)” explicit in the prompt body.

2) Plan the 20 questions (before writing any command lines)
   - Use only the required categories (no additions, no removals).
   - Ensure horizon mix: exactly 12 `Immediate` and 8 `Strategic`.
   - Ensure RPG infusion in every step’s `-p` prompt (WHAT/HOW + DependsOn + Phase).
   - Ensure dependencies only point backward (DependsOn may reference only earlier step IDs).
   - Ensure coverage: across all 20 steps, each required category appears at least once.

3) Assign scoring and compute ROI
   - Pick `impact`, `confidence`, `effort` each as one decimal in `[0.1..1.0]`.
   - Compute `ROI = (impact * confidence) / effort` (one decimal).
   - Sort all 20 steps by ROI descending; tie-break by lower effort.

4) Emit the pack Markdown
   - Copy `assets/oracle-pack-template.md` structure exactly.
   - Fill in metadata values and the 20 steps.
   - Keep exactly ONE fenced bash block.
   - Keep exactly 20 steps (01..20) inside that block.

5) Emit the coverage check section
   - List the 10 required categories exactly once each with `OK` or `Missing(...)`.

6) Validate (optional but recommended when execution is available)
   - Run `scripts/validate_oracle_pack.py docs/oracle-pack-YYYY-MM-DD.md` and fix any reported issues.

## Output contract (must hold)

- Exactly one Markdown deliverable.
- Matches `assets/oracle-pack-template.md` structure exactly.
- Exactly one fenced bash block.
- Bash block contains exactly 20 steps, numbered 01..20.
- Each step has:
  - Header line with: `ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
  - Exactly one runnable oracle command including:
    - `${oracle_cmd}` (default `oracle`)
    - `${oracle_flags}` (default `--files-report`)
    - `--write-output "<out_dir>/<nn>-<slug>.md"`
- Steps sorted by ROI desc; tie-break by lower effort.
- Coverage check section present and correctly formatted.

## Failure modes (do not guess)

- Unknown subsystem/file references: set `reference=Unknown` and include explicit missing file/path patterns in section (4) of the prompt body.
- Missing or unclear user args: apply defaults; do not invent tool flags or repo-specific details.
- Potential contract violations: prioritize strict compliance over richer prose.

## Invocation examples

1) “Generate an RPG-infused oracle strategist pack for this repo. out_dir=oracle-out.”
2) “Create the oracle-pack for a monorepo; prioritize caching/state, observability, and permissions. Use oracle_flags=--files-report.”
3) “Make a pack for codebase_name=AcmeWeb, constraints=‘No DB migrations this sprint’, non_goals=‘No UI redesign’.”
4) “Produce a pack that emphasizes feature flags and invariants, but still covers all required categories.”
5) “Generate the pack with conservative confidence values (0.2–0.6) and document missing artifact patterns explicitly.”
```

.config/skills/oracle-strategist-commands/SKILL.md
```
---

name: oracle-strategist-commands
description: Generate exactly 20 evidence-anchored strategist questions as runnable @steipete/oracle bash commands (12 immediate + 8 strategic), using inference-first repo discovery to avoid wasted searching.
metadata:
  short-description: 20 ROI-ranked Oracle CLI question commands for a repo (inference-first discovery)
  output: bash
  count: 20
---

## Use when

Use this skill when the user wants incisive, feasibility-focused strategist questions, but **as an executable pack** of **exactly 20** `oracle` CLI commands.

## Deliverable (strict)

Output MUST be:

- A single fenced code block: ```bash
- Optional header lines (env vars, `mkdir -p`, shared preamble)
- Then **exactly 20** Oracle invocations (lines beginning with `oracle` or `npx -y @steipete/oracle`)
- Counts:
  - **12** commands that are “Immediate” (01–12)
  - **8** commands that are “Strategic” (13–20)
- No prose outside the bash code fence.

Each command MUST include:

- `-p/--prompt`
- at least one `-f/--file`
- `--files-report`
- `--write-output "<out_dir>/<nn>-<slug>.md"`
- `--wait` (unless user explicitly requests omitting it)

Optional but recommended:

- a commented preflight line using `--dry-run summary` before each real command.

## Inputs (do not ask follow-ups)

Interpret free-text args as optionally including:

- `codebase_name` (default `Unknown`)
- `constraints` (default `None`)
- `non_goals` (default `None`)
- `team_size` (default `Unknown`)
- `deadline` (default `Unknown`)
- `engine` (`api|browser`; if unspecified, omit)
- `model` (if unspecified, omit)
- `out_dir` (default `docs/oracle-questions-YYYY-MM-DD`)

If missing, apply defaults and proceed.

## Core principle: inference-first discovery

Avoid “hard-targeting” specific artifact types up front (e.g., hunting for migrations/auth/jobs) because many repos won’t have them. Instead:

1) Gather cheap high-signal evidence (anchors).
2) Infer the stack and architecture hypotheses.
3) Run small, evidence-gated probes to confirm.
4) Only search/read deeper for categories that now have supporting evidence.

This reduces wasted probing and keeps attachments minimal.

## Workflow

### 1) Build a repo fingerprint (anchors only)

Read the smallest set of files that often exist and immediately reveal stack/shape:

- `README*` (first ~200 lines)
- root manifests that exist (one or more): `package.json`, `pyproject.toml`, `go.mod`, `Cargo.toml`, `pom.xml`, `build.gradle`, `composer.json`, `Gemfile`
- `Dockerfile`, `docker-compose.*` (if present)
- `.github/workflows/*` (names + one representative workflow file)
- `Makefile` / `Taskfile.*` (if present)

Also capture a lightweight file index summary (without reading everything):

- top-level directories
- dominant file extensions
- any “obvious entrypoint” indicators from scripts/config (e.g., package.json scripts, go `cmd/`, rust `src/main.rs`)

### 2) Form hypotheses (stack + entrypoints)

From the fingerprint, infer:

- primary language/runtime
- framework candidates (e.g., Express/Nest/Next, FastAPI/Django, chi/gin, Rails, Spring)
- likely entrypoints and routing surfaces
- where “truth” lives (schema, API contracts, config)

Write down 3–6 candidate entrypoints (paths) you will validate.

### 3) Evidence-gated probing loop (bounded)

Run a small loop of probes; each probe must have a clear purpose and a stop condition.

**Budgets (defaults):**

- max probes: 12
- max file reads: 25
- max search passes: 6

**Allowed probes (choose the cheapest next step):**

- Open an inferred entrypoint and follow imports to one routing/handler layer.
- Read a config file referenced by an entrypoint.
- Perform a targeted search ONLY for confirmed-framework idioms (e.g., “router”, “app.get”, “FastAPI()”, “chi.NewRouter”) after framework is evidenced.
- If you see a strong signal (e.g., `prisma` in deps), then and only then probe its likely files (`prisma/schema.prisma`, migrations).

**Do NOT** scan for categories (auth/jobs/migrations/feature flags/observability) unless you have at least one of:

- a dependency/config signal (library name, config key)
- an import path signal
- a directory naming signal found naturally via the fingerprint or followed imports

### 4) Build the evidence index (mandatory)

Create an internal evidence index with **>= 20** references that actually exist, as one of:

- `path:symbol` (function/class/type/handler name)
- explicit endpoint strings (e.g., `/api/users`)
- job/event names (queue/topic/schedule identifiers)
- config keys/flags (feature gates, env var keys)

For each reference, also record a “best attachment set” (1–6 files) that makes it intelligible.

If the repo is too small to reach 20, pad with `Unknown` entries; those questions must explicitly call out missing evidence.

### 5) Draft exactly 20 strategist questions (12 immediate + 8 strategic)

Each question must:

- be feasibility-focused and grounded in repo reality
- respect `constraints`
- explicitly exclude `non_goals`
- be actionable for `team_size` and `deadline`
- cite at least one evidence reference (or `Unknown` with a clear gap)
- fall into one of these areas (only when evidenced; otherwise treat as a “risk/unknowns” question):
  - contracts/interfaces
  - invariants/data consistency
  - auth/permissions
  - migrations/schema
  - request/UX flows (if applicable)
  - failure modes & reliability
  - feature flags/config gating
  - observability (logs/metrics/tracing)

Guidance:

- Immediate (01–12): unblockers, correctness, runtime risk, deployment friction, quick wins.
- Strategic (13–20): architectural constraints, roadmap leverage, risk retirement, scalability, maintainability.

### 6) Emit the question pack as bash commands (final output)

#### Common command shape

Each command MUST:

- call Oracle with `-p/--prompt` and `-f/--file` (repeat `-f` as needed)
- include `--files-report`
- include `--write-output "<out_dir>/<nn>-<slug>.md"`
- include `--wait` unless user opted out
- attach minimal files:
  - small base context: 1–3 anchor files (entrypoint + config)
  - focus context: files containing the cited evidence reference(s)

Avoid broad `**/*` globs. Prefer explicit files and small globs with negation excludes.

#### Prompt micro-format (required in every `-p`)

Each `-p` must require Oracle to respond using exactly:

- **Answer**
- **Evidence** (bullets; cite concrete refs)
- **Risks / unknowns**
- **Smallest experiment** (runnable today)

[TRUNCATED]
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
[TRUNCATED]
```

.config/skills/oraclepack-usecase-rpg/SKILL.md
```
---
name: oraclepack-usecase-rpg
description: Generate a strict, oraclepack-ingestible 20-step Oracle strategist pack (single ```bash fence) with ROI-ranked step headers and RPG (Repository Planning Graph) metadata embedded in each -p prompt body.
metadata:
  short-description: Strict oraclepack pack generator (RPG-infused)
---

## Quick start

Use this skill when the user asks for any of the following:
- an “oraclepack pack”, “oracle strategist question pack”, or “20-step ROI-ranked pack”
- “oraclepack use cases” that must keep the strict runner-ingestible output
- the same, but with RPG/RPG graph/Repository Planning Graph infusion

What you produce:
- Exactly one Markdown deliverable (intended path: `docs/oracle-pack-YYYY-MM-DD.md`)
- The deliverable contains **exactly one** fenced code block that starts with **exactly** ` ```bash` then a newline (no extra tokens after `bash`)
- The fenced `bash` block contains **exactly 20** sequential steps, `01..20`

Reference template: `references/oracle-pack-template.md`

Optional validator:
- `python3 scripts/validate_oraclepack_pack.py docs/oracle-pack-YYYY-MM-DD.md`

## Workflow

### 1) Parse inputs (free-text args)

Interpret the user’s trailing text as conceptual `{{args}}` and extract:
- `codebase_name` (default: `Unknown`)
- `constraints` (default: `None`)
- `non_goals` (default: `None`)
- `team_size` (default: `Unknown`)
- `deadline` (default: `Unknown`)
- `out_dir` (default: `oracle-out`)
- `oracle_cmd` (default: `oracle`)
- `oracle_flags` (default: `--files-report`)
- `extra_files` (default: empty; append **literally** to every command if provided)

If any value is missing or ambiguous, use the defaults above (do not ask questions).

### 2) Repo discovery (inference-first, deterministic)

Read small, high-signal “index” artifacts first (if they exist):
- README / overview docs
- manifests / dependency lockfiles
- entrypoints (main binaries, CLI, server startup)
- config examples
- CI configuration

Then do targeted discovery (only as needed):
- Prefer deterministic searches (e.g., `ck`, `ast-grep`, `fd`) if available.
- Cap results deterministically (fixed max hits).
- If a file/path does not exist: record `Unknown` and continue.

### 3) Plan the 20 probes (what to cover)

You must keep categories **exactly** to this set (do not add/remove/rename):
- contracts/interfaces
- invariants
- caching/state
- background jobs
- observability
- permissions
- migrations
- UX flows
- failure modes
- feature flags

Across the 20 steps, ensure at least one step explicitly targets each sub-layer below (embed it in the *Question* text; map to the best required category above):
- supply chain/dependencies
- secrets/credentials
- CI/CD or release process
- runtime configuration/IaC (if present)
- privacy/PII/retention (if applicable)
- SLO/alerting/runbooks
- cost/performance hotspots
- testing strategy/coverage

Horizon mix requirement:
- Exactly **12** steps with `horizon=Immediate`
- Exactly **8** steps with `horizon=Strategic`

### 4) RPG infusion (non-breaking)

Inside each step’s `-p` prompt body, include an `RPG:` block with:
- `FunctionalNode: <Capability>::<Feature> (WHAT)`
- `StructuralNode: <module-or-dir> :: <public surface> (HOW)`
- `DependsOn: [<prior step ids, optional>]`
- `Phase: <P0|P1|P2|P3>`

Dependency rule (to preserve ROI-sorted executability):
- `DependsOn` may reference **only earlier step IDs**.

Phase mapping guidance (do not create new categories):
- `P0`: contracts/interfaces, invariants, permissions, observability
- `P1`: caching/state, background jobs, failure modes
- `P2`: migrations, feature flags
- `P3`: UX flows

### 5) ROI scoring and ordering

For each step choose:
- `impact`, `confidence`, `effort` as **one decimal** in `0.1..1.0`

Compute:
- `ROI = (impact * confidence) / effort`

Ordering requirement:
- Sort all 20 steps by ROI descending
- Tie-break by lower effort
- After sorting, assign IDs `01..20` in that sorted order

If ordering conflicts with RPG dependencies:
- Adjust impact/confidence/effort (and thus ROI) so prerequisite/foundation steps float earlier.
- Keep `DependsOn` pointing backwards only.

### 6) Render the pack (strict output)

You must output exactly one Markdown deliverable matching the structure in `references/oracle-pack-template.md`:
- Title section
- Parsed args list
- Commands section with **exactly one** fenced code block:
  - opening fence must be exactly: <code>```bash</code>
  - closing fence must be exactly: <code>```</code>
- Coverage check section

Inside the fenced `bash` block:
- Include a small prelude setting `out_dir="..."` (so pack metadata can be derived).
- Each of the 20 steps must start with a header line matching one of:
  - `# 01) ...`
  - `# 01 — ...`
  - `# 01 - ...`

Each step header line must include all tokens:
- `ROI=<number>`
- `impact=<i>`
- `confidence=<c>`
- `effort=<e>`
- `horizon=<Immediate|Strategic>`
- `category=<one of required categories>`
- `reference=<...>` (use `Unknown` if needed)

Each step must invoke the Oracle CLI once and must include:
- `--write-output "<out_dir>/<nn>-<slug>.md"`
- A `-p` prompt body that includes (in this exact section order):
  - `Strategist question #NN`
  - `RPG:` block (fields listed above)
  - `Reference: ...`
  - `Category: ...`
  - `Horizon: ...`
  - `ROI: ... (impact=..., confidence=..., effort=...)`
  - `Question: ...`
  - `Rationale: ...` (exactly one sentence)
  - `Smallest experiment today: ...` (exactly one action)
  - `Constraints: ...`
  - `Non-goals: ...`
  - `Answer format:` with the 4 required numbered items

Quoting guidance for `-p` (recommended, robust):
- Use a heredoc in command substitution so the prompt stays readable while remaining a single oracle invocation:
  - `-p "$(cat <<'PROMPT' ... PROMPT)"`

### 7) Validate before finalizing

Run the included validator if possible:
- `python3 scripts/validate_oraclepack_pack.py docs/oracle-pack-YYYY-MM-DD.md`

Fix any failures by editing the pack until it passes.

## Output contract

When invoked, produce:

1) A single line:
- `Output file: docs/oracle-pack-YYYY-MM-DD.md`

2) Immediately after, the **exact** Markdown content for that file (no extra commentary).

If the current date is not available in the execution context:
- Use `docs/oracle-pack.md`.

## Failure modes (do not guess)

- Missing repo evidence → use `Reference: Unknown` and in the prompt’s item (4) name the exact missing file/path patterns to attach next.
- Missing args → apply defaults; do not ask follow-ups.
[TRUNCATED]
```

.config/mcp/oraclepack-gold-pack/references/attachment-minimization.md
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

.config/mcp/oraclepack-gold-pack/references/inference-first-discovery.md
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

.config/mcp/oraclepack-gold-pack/references/oracle-pack-template.md
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
[TRUNCATED]
```

.config/mcp/oraclepack-gold-pack/references/oracle-scratch-format.md
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

.config/mcp/oraclepack-gold-pack/scripts/lint_attachments.py
```
# path: oraclepack-gold-pack/scripts/lint_attachments.py
import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple


@dataclass
class Step:
    n: str
    header: str
    lines: List[str]


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        return path.read_text(encoding="utf-8", errors="replace")


def _extract_bash_fence(lines: List[str]) -> List[str]:
    fence_idxs = [i for i, ln in enumerate(lines) if ln.startswith("```")]
    if len(fence_idxs) != 2:
        raise ValueError(f"Expected exactly one fenced block (2 fence lines). Found {len(fence_idxs)}.")
    open_i, close_i = fence_idxs
    if lines[open_i].rstrip("\n") != "```bash":
        raise ValueError("Opening fence must be exactly ```bash.")
    if lines[close_i].rstrip("\n") != "```":
        raise ValueError("Closing fence must be exactly ```.")
    return [ln.rstrip("\n") for ln in lines[open_i + 1 : close_i]]


def _parse_steps(fence_lines: List[str]) -> List[Step]:
    header_re = re.compile(r"^#\s*(\d{2})\)\s+")
    header_idxs: List[Tuple[int, str]] = []
    for i, ln in enumerate(fence_lines):
        m = header_re.match(ln)
        if m:
            header_idxs.append((i, m.group(1)))

    if not header_idxs:
        raise ValueError("No step headers found inside bash fence.")

    steps: List[Step] = []
    for idx, (start_i, n) in enumerate(header_idxs):
        end_i = header_idxs[idx + 1][0] if idx + 1 < len(header_idxs) else len(fence_lines)
        block = fence_lines[start_i:end_i]
        steps.append(Step(n=n, header=block[0], lines=block))
    return steps


def _count_attachments(step: Step) -> int:
    count = 0
    for ln in step.lines[1:]:
        s = ln.strip()
        if not s or s.startswith("#"):
            continue
        # Count occurrences in non-comment lines
        count += len(re.findall(r"(?<!\S)(-f|--file)(?!\S)", ln))
    return count


def _is_unknown_reference(step: Step) -> bool:
    # Step header token format contains reference=...
    m = re.search(r"\breference=([^\s]+)", step.header)
    if not m:
        return False
    val = m.group(1).strip()
    return val.lower() == "unknown"


def _has_missing_artifact_request(step: Step) -> bool:
    hay = "\n".join(step.lines).lower()
    # Accept several common phrasings, keep it simple and robust.
    patterns = [
        r"missing file/path pattern",
        r"missing file.*pattern",
        r"attach next",
        r"name the exact missing.*pattern",
    ]
    return any(re.search(p, hay) for p in patterns)

def lint(path: Path) -> None:
    raw = _read_text(path)
    lines = raw.splitlines(True)
    fence = _extract_bash_fence(lines)
    steps = _parse_steps(fence)

    errors: List[str] = []
    for step in steps:
        attachments = _count_attachments(step)
        if attachments > 2:
            errors.append(f"Step {step.n}: has {attachments} attachments; must be <= 2 (minimal attachments rule).")

        if _is_unknown_reference(step) and not _has_missing_artifact_request(step):
            errors.append(
                f"Step {step.n}: reference=Unknown but prompt does not request missing file/path pattern(s) to attach next."
            )

    if errors:
        for e in errors:
            print(f"[ERROR] {e}", file=sys.stderr)
        sys.exit(1)

    print("[OK] Attachment lint passed (<=2 attachments per step; Unknown-reference prompts request missing patterns).")

def main() -> None:
    p = argparse.ArgumentParser(description="Lint oraclepack Stage-1 gold pack attachments (<=2 per step) and Unknown-reference prompt behavior.")
    p.add_argument("pack_path", help="Path to the Markdown pack file")
    args = p.parse_args()

    path = Path(args.pack_path)
    if not path.exists():
        print(f"[ERROR] File not found: {path}", file=sys.stderr)
        sys.exit(1)

    lint(path)


if __name__ == "__main__":
    main()
```

.config/mcp/oraclepack-gold-pack/scripts/validate_pack.py
```
# path: oraclepack-gold-pack/scripts/validate_pack.py
import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple, Optional, Dict


ALLOWED_CATEGORIES = [
    "contracts/interfaces",
    "invariants",
    "caching/state",
    "background jobs",
    "observability",
    "permissions",
    "migrations",
    "UX flows",
    "failure modes",
    "feature flags",
]

REQUIRED_TOKENS = [
    "ROI=",
    "impact=",
    "confidence=",
    "effort=",
    "horizon=",
    "category=",
    "reference=",
]


@dataclass
class Step:
    n: str
    header_line_no: int
    header_line: str
    block_lines: List[str]


def _fail(errors: List[str]) -> None:
    for e in errors:
        print(f"[ERROR] {e}", file=sys.stderr)
    sys.exit(1)


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        # fallback
        return path.read_text(encoding="utf-8", errors="replace")


def _extract_single_bash_fence(lines: List[str]) -> Tuple[int, int, List[str], List[str]]:
    """
    Enforces:
      - exactly two fence lines in entire document
      - first fence line must be exactly ```bash
      - closing fence line must be exactly ```
    Returns: (open_idx, close_idx, fence_lines, outside_lines)
    """
    fence_idxs = [i for i, ln in enumerate(lines) if ln.startswith("```")]

    if len(fence_idxs) != 2:
        raise ValueError(
            f"Expected exactly 1 fenced code block (2 fence lines), found {len(fence_idxs)} fence line(s)."
        )

    open_idx, close_idx = fence_idxs
    if lines[open_idx].rstrip("\n") != "```bash":
        raise ValueError("Opening fence must be exactly ```bash (no extra tokens/spaces).")
    if lines[close_idx].rstrip("\n") != "```":
        raise ValueError("Closing fence must be exactly ``` (no extra tokens/spaces).")
    if close_idx <= open_idx:
        raise ValueError("Closing fence appears before opening fence.")

    fence_lines = lines[open_idx + 1 : close_idx]
    outside_lines = lines[:open_idx] + lines[close_idx + 1 :]
    return open_idx, close_idx, fence_lines, outside_lines


def _parse_steps(fence_lines: List[str]) -> List[Step]:
    header_re = re.compile(r"^#\s*(\d{2})\s+")
    header_idxs: List[Tuple[int, str]] = []
    for i, ln in enumerate(fence_lines):
        m = header_re.match(ln)
        if m:
            header_idxs.append((i, m.group(1)))

    if len(header_idxs) != 20:
        raise ValueError(f"Expected exactly 20 step headers inside bash fence, found {len(header_idxs)}.")

    expected = [f"{i:02d}" for i in range(1, 21)]
    got = [n for _, n in header_idxs]
    if got != expected:
        raise ValueError(f"Step numbering must be sequential 01..20. Got: {got}")

    steps: List[Step] = []
    for idx, (start_i, n) in enumerate(header_idxs):
        end_i = header_idxs[idx + 1][0] if idx + 1 < len(header_idxs) else len(fence_lines)
        block = fence_lines[start_i:end_i]
        steps.append(
            Step(
                n=n,
                header_line_no=start_i + 1,  # 1-based within fence
                header_line=block[0].rstrip("\n"),
                block_lines=[b.rstrip("\n") for b in block],
            )
        )
    return steps


def _validate_header(step: Step, errors: List[str]) -> None:
    header = step.header_line

    for tok in REQUIRED_TOKENS:
        if tok not in header:
            errors.append(f"Step {step.n}: missing required token '{tok}' in header: {header}")

    # ensure each token has a value (non-empty, non-space)
    for key in ["ROI", "impact", "confidence", "effort", "horizon", "category", "reference"]:
        m = re.search(rf"\b{re.escape(key)}=([^\s]+)", header)
        if not m:
            errors.append(f"Step {step.n}: header missing/empty '{key}=' value: {header}")

    cat_m = re.search(r"\bcategory=([^\s]+(?:\s+[^\s]+)?)", header)
    # Category can contain a space (e.g., "background jobs", "UX flows", "failure modes", "feature flags")
    # The regex above captures up to two words; handle 2-word categories explicitly by matching allowed set.
    if cat_m:
        # Extract by checking allowed set presence after 'category='
        after = header.split("category=", 1)[1]
        # category value ends at next token start or end of line
        end = len(after)
        for token in [" reference=", " ROI=", " impact=", " confidence=", " effort=", " horizon="]:
            pos = after.find(token)
            if pos != -1:
                end = min(end, pos)
        cat_val = after[:end].strip()
        if cat_val not in ALLOWED_CATEGORIES:
            errors.append(
                f"Step {step.n}: invalid category='{cat_val}'. Must be one of: {ALLOWED_CATEGORIES}"
            )
    else:
        errors.append(f"Step {step.n}: could not parse category=... from header: {header}")


def _validate_write_output(step: Step, errors: List[str]) -> None:
    # Find first --write-output in the step block
    joined = "\n".join(step.block_lines)
    m = re.search(r'--write-output\s+(["(\S+)"|"(\S+)"|(\S+)])', joined)
    if not m:
        errors.append(f"Step {step.n}: missing --write-output in step block.")
        return

    path = m.group(2) or m.group(3) or m.group(4) or ""
    if "/" not in path:
        errors.append(
            f"Step {step.n}: --write-output path must look like <out_dir>/{step.n}-<slug>.md; got: {path}"
        )
        return

    filename = path.split("/")[-1]
    if not filename.startswith(f"{step.n}-"):
        errors.append(
            f"Step {step.n}: --write-output filename must start with '{step.n}-'; got: {filename}"
        )
    if not filename.endswith(".md"):
        errors.append(f"Step {step.n}: --write-output filename must end with .md; got: {filename}")


def _validate_coverage_check(outside_lines: List[str], errors: List[str]) -> None:
    text = "\n".join(outside_lines)
    # Require a "Coverage check" heading (case-insensitive)
    if re.search(r"^##\s+Coverage check\s*$", text, flags=re.IGNORECASE | re.MULTILINE) is None:
        errors.append('Missing "## Coverage check" section (must be outside the bash fence).')
        return

    # Ensure every category appears as "<category>:` somewhere after the heading
    # Find slice after the heading
[TRUNCATED]
```

.config/mcp/oraclepack-taskify/assets/action-pack-template.md
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
[TRUNCATED]
```

.config/mcp/oraclepack-taskify/assets/actions-json-schema.md
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

.config/mcp/oraclepack-taskify/assets/prd-synthesis-prompt.md
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

.config/mcp/oraclepack-taskify/references/determinism-and-safety.md
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

.config/mcp/oraclepack-taskify/references/task-master-cli-cheatsheet.md
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

.config/mcp/oraclepack-taskify/references/workflow-overview.md
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

.config/mcp/oraclepack-taskify/scripts/detect-oracle-outputs.sh
```
#!/usr/bin/env bash
set -euo pipefail

out_dir="${1:-oracle-out}"

if [[ ! -d "${out_dir}" ]]; then
  echo "ERROR: out_dir does not exist: ${out_dir}" >&2
  exit 2
fi

shopt -s nullglob

for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
  matches=( "${out_dir}/${n}-"*.md )
  if [[ "${#matches[@]}" -eq 0 ]]; then
    echo "ERROR: missing output for prefix ${n}: expected ${out_dir}/${n}-*.md" >&2
    exit 3
  fi
  if [[ "${#matches[@]}" -gt 1 ]]; then
    echo "ERROR: multiple outputs for prefix ${n}; expected exactly one file." >&2
    printf '%s\n' "${matches[@]}" >&2
    exit 4
  fi
  printf '%s\n' "${matches[0]}"
done
```

.config/mcp/oraclepack-taskify/scripts/validate-action-pack.sh
```
#!/usr/bin/env bash
set -euo pipefail

pack_path="${1:-}"
if [[ -z "${pack_path}" ]]; then
  echo "Usage: validate-action-pack.sh <path/to/oracle-actions-pack.md>" >&2
  exit 2
fi

if [[ ! -f "${pack_path}" ]]; then
  echo "ERROR: file not found: ${pack_path}" >&2
  exit 3
fi

# Rule: exactly one bash fence, and no other fences.
bash_fence_count="$(grep -cE '^[[:space:]]*```bash[[:space:]]*$' "${pack_path}" || true)"
any_fence_count="$(grep -cE '^[[:space:]]*```' "${pack_path}" || true)"

if [[ "${bash_fence_count}" -ne 1 ]]; then
  echo "ERROR: expected exactly one '```bash' fence; found ${bash_fence_count}" >&2
  exit 10
fi

if [[ "${any_fence_count}" -ne 2 ]]; then
  echo "ERROR: expected exactly 2 total fences (start+end); found ${any_fence_count}" >&2
  echo "Fences found:" >&2
  grep -nE '^[[:space:]]*```' "${pack_path}" >&2 || true
  exit 11
fi

# Extract lines within the bash fence and validate step headers.
in_bash=0
headers=()
while IFS= read -r line; do
  if [[ "${line}" =~ ^[[:space:]]*```bash[[:space:]]*$ ]]; then
    in_bash=1
    continue
  fi
  if [[ "${line}" =~ ^[[:space:]]*```[[:space:]]*$ ]]; then
    in_bash=0
    continue
  fi
  if [[ "${in_bash}" -eq 1 ]]; then
    if [[ "${line}" =~ ^#\ ([0-9]{2})\) ]]; then
      headers+=( "${BASH_REMATCH[1]}" )
    fi
  fi
done < "${pack_path}"

if [[ "${#headers[@]}" -lt 1 ]]; then
  echo "ERROR: no step headers found inside bash fence (expected '# NN)')" >&2
  exit 20
fi

# Validate strict sequential: 01..N with no gaps and no duplicates.
seen=""
expected=1
for h in "${headers[@]}"; do
  # Check duplicates
  if [[ " ${seen} " == *" ${h} "* ]]; then
    echo "ERROR: duplicate step header: ${h}" >&2
    exit 21
  fi
  seen="${seen} ${h}"

  exp="$(printf '%02d' "${expected}")"
  if [[ "${h}" != "${exp}" ]]; then
    echo "ERROR: non-sequential step header. Expected ${exp}, got ${h}" >&2
    echo "All headers: ${headers[*]}" >&2
    exit 22
  fi
  expected=$((expected + 1))
done

echo "OK: Action Pack validation passed."
```

.config/skills/oracle-pack/assets/oracle-pack-template.md
```
<!-- path: ~/.codex/skills/oracle-pack/assets/oracle-pack-template.md -->
# oracle strategist question pack

---

## parsed args

- codebase_name: <Unknown|value>
- constraints: <None|value>
- non_goals: <None|value>
- team_size: <Unknown|value>
- deadline: <Unknown|value>
- out_dir: <oracle-out|value>
- oracle_cmd: <oracle|value>
- oracle_flags: <--browser-attachments always --files-report|value>
- extra_files: <empty|value>

---

## commands (exactly 20; sorted by ROI desc; ties by lower effort)

```bash
# 01 — ROI=<..> impact=<..> confidence=<..> effort=<..> horizon=<Immediate|Strategic> category=<...> reference=<...>
<oracle_cmd> \
  <oracle_flags> \
  --write-output "<out_dir>/01-<slug>.md" \
  -p "Strategist question #01
Reference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>
Category: <one of required categories>
Horizon: <Immediate|Strategic>
ROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)
Question: <question text>
Rationale: <exactly one sentence>
Smallest experiment today: <exactly one action>
Constraints: <constraints or None>
Non-goals: <non_goals or None>

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "<minimal evidence file 1>" \
  -f "<optional supporting file 2>" \
  <optional extra_files entries...>

# 02 ...
# ...
# 20 ...
```

---

## coverage check (must be satisfied)

*   contracts/interfaces: <OK|Missing (state missing artifact pattern)>

*   invariants: <OK|Missing (...)>

*   caching/state: <OK|Missing (...)>

*   background jobs: <OK|Missing (...)>

*   observability: <OK|Missing (...)>

*   permissions: <OK|Missing (...)>

*   migrations: <OK|Missing (...)>

*   UX flows: <OK|Missing (...)>

*   failure modes: <OK|Missing (...)>

*   feature flags: <OK|Missing (...)>
```
```

.config/skills/oracle-pack/references/attachment-minimization.md
```
<!-- path: ~/.codex/skills/oracle-pack/references/attachment-minimization.md -->
# Attachment minimization rules (evidence-driven)

Objective: keep Oracle inputs small and high-signal. Attach only what Oracle needs to answer the single question.

## Always exclude

- Secrets (`.env`, key files, tokens)
- Large generated dirs unless explicitly needed (`node_modules`, `dist`, `build`, `coverage`, etc.)
- Broad globs like `src/**` unless the question has `Unknown` reference and no better anchor exists

## Core mapping: reference → files to attach

### Reference = `{path}:{symbol}`

Attach:
1) `path` (the file containing the symbol)

Optionally attach **one** more file (only if it materially affects correctness):
- nearest router/entrypoint that calls into `path`
- config file that wires the subsystem (`config/*`, `settings.*`, `env` loader)
- interface/schema referenced by the symbol (e.g., `openapi.*`, `schema.*`)

Avoid attaching more than 2 files unless the repo uses very small modules and it’s clearly required.

### Reference = `{endpoint}`

Attach:
1) the route map / router file where the endpoint is declared
2) the handler file (if different)

Prefer literal files over globs.

### Reference = `{event}`

Attach:
1) the job/worker registration file where the event is scheduled/declared
2) the worker implementation file (if different)

### Reference = `Unknown`

Attach only “index” evidence:
- README (or closest equivalent)
- primary manifest (`package.json`, `pyproject.toml`, etc.)
- one best-guess entrypoint file (if you found one)

In the prompt, explicitly request the missing artifact pattern to attach next (e.g., “attach `**/migrations/**`”).

## Optional global extras (`extra_files`)

If the user provided `extra_files`, append those attachments to every command *after* the minimal evidence attachments. Use this sparingly.
```

.config/skills/oracle-pack/references/inference-first-discovery.md
```
<!-- path: ~/.codex/skills/oracle-pack/references/inference-first-discovery.md -->
# Inference-first discovery (adaptive search ladder, ck-first)

Goal: avoid wasting time on hard-targeted globs/greps that may not exist by deriving the search plan from repo evidence,
using the codebase-search rules (ck → ast-grep when structure matters → deterministic output shaping).

## Step 0 — Discover the search interface (required)

- Always run: `ck --help`
- Use only flags confirmed by `ck --help`.
- If `ck` indicates indexing/setup is needed, follow its printed instructions.

## Step 1 — Read “index” artifacts first (cheap, high-signal)

Use `fd` to locate these quickly:

- Repo docs/index:
  - `fd -p README.md`
  - `fd -e md . docs | head`
- Manifests/config (choose the ones that exist):
  - `fd -p package.json`
  - `fd -p pyproject.toml`
  - `fd -p go.mod`
  - `fd -p Cargo.toml`
  - `fd -p pom.xml`
  - `fd -p build.gradle`

Read the smallest set of files that explain structure and entrypoints.

## Step 2 — Infer subsystem locations from what you actually saw

Build a short “signals” list from evidence:
- Router/framework name (from dependencies and scripts)
- Job system
- Migration tool
- Feature flag system
- Observability stack
- Cache layer

Then: follow imports/registration code rather than searching the whole repo.

## Step 3 — Derive targeted `ck` queries (conceptual-first)

Start conceptual (meaning-based), scoped to inferred roots:

- Conceptual: `ck --sem "<intent phrase>"`
  - Examples of intent phrases (adapt to the repo’s vocabulary):
    - “route registration” / “router setup”
    - “auth middleware” / “permission check”
    - “queue worker registration”
    - “migration runner” / “schema change”
    - “feature flag evaluation”
    - “telemetry initialization” / “logger setup”

When unsure, use hybrid:
- `ck --hybrid "<intent phrase>"`

For literal hits (exact tokens you already saw), use regex:
- `ck --regex "<literal-or-regex>"`

### Deterministic narrowing

If results are broad/noisy, narrow deterministically:

- Use `ck --jsonl` (if available) and cap results:
  - `ck --jsonl ... | jq ... | sort | head -n 20`
- Otherwise cap with `head` and re-run `ck` constrained to the top-hit directories/files.

## Step 4 — Structural search with `ast-grep` (when structure matters)

Use `ast-grep` when you need syntax-aware matching (AST-level), such as:
- finding a specific call pattern (e.g., permission checks around handlers)
- finding registration shapes (e.g., route definitions)
- reducing false positives vs text search

Examples:

- Find code structure:
  - `ast-grep --lang <language> -p '<pattern>'`
- List matching files (cap output):
  - `ast-grep -l --lang <language> -p '<pattern>' | head -n 10`

Use `ck` first to locate candidate files/modules, then `ast-grep` to enforce structure inside those.

## Step 5 — Progressive widening (fallback ladder)

If inference cannot locate a subsystem:

1) Run a small set of conceptual `ck --sem` / `ck --hybrid` queries using generic intent phrases.
2) If still nothing, broaden to repo-wide `ck --hybrid` for that intent.
3) If still nothing, mark evidence for that category as missing and record the most likely missing artifact pattern.

## Step 6 — Early stopping (don’t over-search)

Stop harvesting once:
- you have >= 20 candidate anchors total, and
- every required category has at least one candidate (or a clearly documented “missing artifact pattern”)

Then generate questions; do not keep scanning “just in case”.
```

.config/skills/oracle-pack/references/oracle-scratch-format.md
```
<!-- path: ~/.codex/skills/oracle-pack/references/oracle-scratch-format.md -->
# oracle usage examples

---

## add attachments

```bash
oracle \
  --browser-attachments always \
  --browser-input-timeout 5s \
  -p "Run the UI smoke test." \
  -f "packages/"
```

## copy to clipboard

```bash
oracle --render --copy \
  -p "Explain the CLI architecture and control flow: index.ts -> registry init -> list/run commands. Identify error-handling gaps, async pitfalls, and how you'd improve UX (help text, exit codes, structured output). Provide proposed code changes." \
  -f "codefetch/packages.md" \
  -f "packages/cli/src/index.ts" \
  -f "packages/cli/src/commands/list.ts" \
  -f "packages/cli/src/commands/run.ts"
```

---

## oracle/codefetch

```bash
oracle \
  --browser-attachments always \
  -p "Using packages.md + attached files, give me a dependency graph and a build order. Then propose 5 repo hygiene improvements." \
  -f "codefetch/packages.md" \
  -f "packages/**/package.json" \
  -f "packages/**/tsconfig.json"
```

---

## other scratch examples

```bash
oracle \
  -p "Walk through the UI smoke test" \
  --file "src/**/*.ts"
```

```bash
oracle \
  -p "Review these changes and propose fixes" \
  -f "src/**/*.ts,!**/*.test.ts"
```

```bash
oracle \
  -p "Using packages.md + attached files, give me a dependency graph and a build order. Then propose 5 repo hygiene improvements." \
  -f "codefetch/packages.md" \
  -f "packages/**/package.json" \
  -f "packages/**/tsconfig.json"
```

```bash
oracle --render --copy \
  -p "Explain the CLI architecture and control flow: index.ts -> registry init -> list/run commands. Identify error-handling gaps, async pitfalls, and how you'd improve UX (help text, exit codes, structured output). Provide proposed code changes." \
  -f "codefetch/packages.md" \
  -f "packages/cli/src/index.ts" \
  -f "packages/cli/src/commands/list.ts" \
  -f "packages/cli/src/commands/run.ts"
```

```bash
oracle  \
  --prompt "Read packages.md and explain: (1) what each package does, (2) how workflows are executed end-to-end, (3) where to add a new workflow, and (4) the top 5 refactors to reduce coupling." \
  --file "codefetch/packages.md"
```
```

.config/skills/oracle-pack-rpg-infused/assets/oracle-pack-template.md
```
# oracle-pack (RPG-infused)

PRIMARY GOAL
Generate an oracle strategist question pack that is strictly ingestible by the oracle-pack template and oraclepack runner, while embedding an RPG (Repository Planning Graph) fragment in each prompt for graph extraction.

## Pack metadata

- codebase_name: <fill or Unknown>
- constraints: <fill or None>
- non_goals: <fill or None>
- team_size: <fill or Unknown>
- deadline: <fill or Unknown>
- out_dir: <fill or oracle-out>
- oracle_cmd: <fill or oracle>
- oracle_flags: <fill or --files-report>

## Steps (exactly one fenced bash block; exactly 20 steps inside)

```bash
# 01 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/01-<slug>.md" -p $'Strategist question #01\nRPG:\nFunctionalNode: <Capability>::<Feature> (WHAT)\nStructuralNode: <module-or-dir> :: <public surface> (HOW)\nDependsOn: []\nPhase: <P0|P1|P2|P3>\n\nReference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>\nCategory: <one of required categories>\nHorizon: <Immediate|Strategic>\nROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)\nQuestion: <question text>\nRationale: <exactly one sentence>\nSmallest experiment today: <exactly one action>\n\nConstraints: <constraints or None>\nNon-goals: <non_goals or None>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 02 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/02-<slug>.md" -p $'Strategist question #02\nRPG:\nFunctionalNode: <Capability>::<Feature> (WHAT)\nStructuralNode: <module-or-dir> :: <public surface> (HOW)\nDependsOn: [01]\nPhase: <P0|P1|P2|P3>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 03 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/03-<slug>.md" -p $'Strategist question #03\nRPG:\nFunctionalNode: <Capability>::<Feature> (WHAT)\nStructuralNode: <module-or-dir> :: <public surface> (HOW)\nDependsOn: [01, 02]\nPhase: <P0|P1|P2|P3>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 04 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/04-<slug>.md" -p $'Strategist question #04\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<prior step ids>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 05 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/05-<slug>.md" -p $'Strategist question #05\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 06 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/06-<slug>.md" -p $'Strategist question #06\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 07 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/07-<slug>.md" -p $'Strategist question #07\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 08 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
[TRUNCATED]
```

.config/skills/oracle-pack-rpg-infused/references/oracle-pack-spec.md
```
# Oracle Pack (RPG-infused) — Spec

## Strict output contract

1) Produce exactly one Markdown deliverable matching `assets/oracle-pack-template.md` structure exactly.
2) The deliverable must contain exactly ONE fenced bash block.
3) The bash block must contain EXACTLY 20 steps, numbered 01..20 sequentially.
4) Each step header line must match:

`# NN ROI=<roi> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>`

5) For each step, emit exactly one runnable oracle command:
   - command: `<oracle_cmd>` (default `oracle`)
   - include `<oracle_flags>` (default `--files-report`)
   - include deterministic output file:

`--write-output "<out_dir>/<nn>-<slug>.md"`

6) Compute ROI:

- Pick `impact`, `confidence`, `effort` each as one decimal in `[0.1..1.0]`.
- Compute `ROI = (impact * confidence) / effort` and round to one decimal.

7) Sort all 20 steps by ROI descending; tie-break by lower effort.

8) Include the coverage check section listing the 10 required categories exactly (OK or Missing(...)).

## Required categories (do not add/remove)

- contracts/interfaces
- invariants
- caching/state
- background jobs
- observability
- permissions
- migrations
- UX flows
- failure modes
- feature flags

## RPG infusion requirements (non-breaking)

A) Each strategist prompt (`-p`) MUST include this RPG block:

- FunctionalNode: capability + feature (WHAT)
- StructuralNode: module/file boundary + interface points (HOW)
- DependsOn: a list of prior step IDs, e.g. [01, 03]
- Phase: P0|P1|P2|P3

B) Dependencies MUST ONLY point backwards:
- DependsOn may reference only earlier step IDs.

C) Across the 20 steps, ensure coverage of BOTH semantics:
- At least 10 steps are functional-first (capability/feature discovery).
- At least 10 steps are structural-first (module/file/interface localization).

D) Ensure at least one step explicitly targets each RPG dimension:
- capabilities/features
- file/module boundaries
- public interfaces/contracts
- data flows/state
- dependency edges/topological build order

## Horizon mix (unchanged)

Exactly 12 Immediate and 8 Strategic.

## Attachment minimization (unchanged)

- Read index artifacts first; infer subsystem locations before sweeping.
- Use deterministic ck/ast-grep/fd queries if available; cap results deterministically.
- Attach only minimal evidence files needed for the single question.
- If reference is Unknown: attach only index files (if possible) and state the missing artifact pattern(s).

## Prompt body format inside -p (must include these sections)

Strategist question #NN
RPG:
FunctionalNode: <Capability>::<Feature> (WHAT)
StructuralNode: <module-or-dir> :: <public surface> (HOW)
DependsOn: [<prior step ids, optional>]
Phase: <P0|P1|P2|P3>

Reference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>
Category: <one of required categories>
Horizon: <Immediate|Strategic>
ROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)
Question: <question text>
Rationale: <exactly one sentence>
Smallest experiment today: <exactly one action>

Constraints: <constraints or None>
Non-goals: <non_goals or None>

Answer format:

1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

## Defaults

If args are missing, use:

- codebase_name=Unknown
- constraints=None
- non_goals=None
- team_size=Unknown
- deadline=Unknown
- out_dir=oracle-out
- oracle_cmd=oracle
- oracle_flags=--files-report
- extra_files=empty

## Final response requirement

In the assistant response:

- Print: `Output file: docs/oracle-pack-YYYY-MM-DD.md`
- Then print the exact same Markdown content (no extra commentary).
```

.config/skills/oracle-pack-rpg-infused/scripts/validate_oracle_pack.py
```
#!/usr/bin/env python3
"""
Validate an oracle-pack (RPG-infused) markdown file against the strict contract.

Usage:
  scripts/validate_oracle_pack.py docs/oracle-pack-YYYY-MM-DD.md
"""

from __future__ import annotations

import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Optional, Tuple


REQUIRED_CATEGORIES: List[str] = [
    "contracts/interfaces",
    "invariants",
    "caching/state",
    "background jobs",
    "observability",
    "permissions",
    "migrations",
    "UX flows",
    "failure modes",
    "feature flags",
]


@dataclass(frozen=True)
class Step:
    nn: int
    roi: float
    impact: float
    confidence: float
    effort: float
    horizon: str
    category: str
    reference: str
    cmd_line: str


def die(msg: str) -> None:
    print(f"[ERROR] {msg}", file=sys.stderr)
    raise SystemExit(1)


def ok(msg: str) -> None:
    print(f"[OK] {msg}")


def read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8")
    except Exception as e:
        die(f"Failed to read {path}: {e}")
        raise  # unreachable


def find_single_bash_block(md: str) -> str:
    blocks = re.findall(r"```bash\s*\n(.*?)\n```", md, flags=re.DOTALL)
    if len(blocks) != 1:
        die(f"Expected exactly 1 fenced bash block; found {len(blocks)}.")
    # Ensure no other fenced blocks exist
    fence_count = len(re.findall(r"```", md))
    if fence_count != 2:
        die(f"Expected exactly 1 fenced code block total (2 fences); found {fence_count} fences.")
    return blocks[0]


def parse_out_dir(md: str) -> str:
    m = re.search(r"(?m)^\s*-\s*out_dir:\s*([^\s]+)\s*$", md)
    if not m:
        return "oracle-out"
    return m.group(1).strip()


HEADER_RE = re.compile(
    r"""
    ^\n    #\s*\n    (?P<nn>\d{2})
    \s+
    ROI=(?P<roi>\d+(?:\.\d)?)
    \s+
    impact=(?P<impact>[01]\.\d)
    \s+
    confidence=(?P<confidence>[01]\.\d)
    \s+
    effort=(?P<effort>[01]\.\d)
    \s+
    horizon=(?P<horizon>Immediate|Strategic)
    \s+
    category=(?P<category>.+?)
    \s+
    reference=(?P<reference>.+)
    $
    """,
    flags=re.VERBOSE,
)


def parse_steps(bash: str) -> List[Step]:
    lines = [ln.rstrip("\n") for ln in bash.splitlines()]
    steps: List[Step] = []

    i = 0
    while i < len(lines):
        ln = lines[i].strip()
        if not ln:
            i += 1
            continue

        m = HEADER_RE.match(ln)
        if not m:
            die(f"Unexpected line in bash block (expected step header): {lines[i]}")
        nn = int(m.group("nn"))
        roi = float(m.group("roi"))
        impact = float(m.group("impact"))
        confidence = float(m.group("confidence")),
        effort = float(m.group("effort")),
        horizon = m.group("horizon")
        category = m.group("category").strip()
        reference = m.group("reference").strip()

        # Next non-empty, non-comment line must be the single oracle command
        j = i + 1
        cmd_line: Optional[str] = None
        while j < len(lines):
            nxt = lines[j].strip()
            if not nxt:
                j += 1
                continue
            if nxt.startswith("#"):
                die(f"Found comment line where oracle command line was expected after step {nn:02d}.")
            cmd_line = lines[j].strip()
            j += 1
            break
        if cmd_line is None:
            die(f"Missing oracle command line after step header {nn:02d}.")

        # Ensure there are no extra non-empty, non-comment lines before next header
        k = j
        while k < len(lines):
            peek = lines[k].strip()
            if not peek:
                k += 1
                continue
            if HEADER_RE.match(peek):
                break
            if peek.startswith("#"):
                die(f"Unexpected comment line inside step body for {nn:02d}; each step must be exactly two lines.")
            die(f"Unexpected extra command/content line inside step {nn:02d}: {lines[k]}")
        i = k

        steps.append(
            Step(
                nn=nn,
                roi=roi,
                impact=impact,
                confidence=confidence,
                effort=effort,
                horizon=horizon,
                category=category,
                reference=reference,
                cmd_line=cmd_line,
            )
        )

    return steps


def require_20_steps(steps: List[Step]) -> None:
    if len(steps) != 20:
        die(f"Expected exactly 20 steps; found {len(steps)}.")

    nns = [s.nn for s in steps]
    expected = list(range(1, 21))
    if nns != expected:
        die(f"Step numbers must be 01..20 in order; found: {[f'{n:02d}' for n in nns]}")


def validate_scores(steps: List[Step]) -> None:
    for s in steps:
        for name, v in [("impact", s.impact), ("confidence", s.confidence), ("effort", s.effort)]:
            if v < 0.1 or v > 1.0:
                die(f"Step {s.nn:02d} {name} must be in [0.1..1.0]; got {v}.")
            # One decimal check (string form after rounding to 1 decimal should match)
            if abs(v - round(v, 1)) > 1e-9:
                die(f"Step {s.nn:02d} {name} must have exactly one decimal; got {v}.")

        expected_roi = round((s.impact * s.confidence) / s.effort, 1)
        if abs(s.roi - expected_roi) > 0.05:
            die(
                f"Step {s.nn:02d} ROI mismatch: header ROI={s.roi} but computed ROI={expected_roi} "
                f"from (impact*confidence)/effort."
            )
        if abs(s.roi - round(s.roi, 1)) > 1e-9:
            die(f"Step {s.nn:02d} ROI must have exactly one decimal; got {s.roi}.")


def validate_sort_order(steps: List[Step]) -> None:
    for a, b in zip(steps, steps[1:]):
        if b.roi > a.roi + 1e-9:
            die(
                f"Steps must be sorted by ROI descending; step {a.nn:02d} ROI={a.roi} "
                f"precedes step {b.nn:02d} ROI={b.roi}."
            )
        if abs(b.roi - a.roi) < 1e-9 and b.effort < a.effort - 1e-9:
            die(
                f"ROI tie-break must prefer lower effort first; step {a.nn:02d} effort={a.effort} "
                f"precedes step {b.nn:02d} effort={b.effort} with equal ROI={a.roi}."
            )


def validate_horizon_mix(steps: List[Step]) -> None:
[TRUNCATED]
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
[TRUNCATED]
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

.config/skills/oraclepack-gold-pack/scripts/lint_attachments.py
```
# path: oraclepack-gold-pack/scripts/lint_attachments.py
import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple


@dataclass
class Step:
    n: str
    header: str
    lines: List[str]


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        return path.read_text(encoding="utf-8", errors="replace")


def _extract_bash_fence(lines: List[str]) -> List[str]:
    fence_idxs = [i for i, ln in enumerate(lines) if ln.startswith("```")]
    if len(fence_idxs) != 2:
        raise ValueError(f"Expected exactly one fenced block (2 fence lines). Found {len(fence_idxs)}.")
    open_i, close_i = fence_idxs
    if lines[open_i].rstrip("\n") != "```bash":
        raise ValueError("Opening fence must be exactly ```bash.")
    if lines[close_i].rstrip("\n") != "```":
        raise ValueError("Closing fence must be exactly ```.")
    return [ln.rstrip("\n") for ln in lines[open_i + 1 : close_i]]


def _parse_steps(fence_lines: List[str]) -> List[Step]:
    header_re = re.compile(r"^#\s*(\d{2})\)?\s+")
    header_idxs: List[Tuple[int, str]] = []
    for i, ln in enumerate(fence_lines):
        m = header_re.match(ln)
        if m:
            header_idxs.append((i, m.group(1)))

    if not header_idxs:
        raise ValueError("No step headers found inside bash fence.")

    steps: List[Step] = []
    for idx, (start_i, n) in enumerate(header_idxs):
        end_i = header_idxs[idx + 1][0] if idx + 1 < len(header_idxs) else len(fence_lines)
        block = fence_lines[start_i:end_i]
        steps.append(Step(n=n, header=block[0], lines=block))
    return steps


def _count_attachments(step: Step) -> int:
    count = 0
    for ln in step.lines[1:]:
        s = ln.strip()
        if not s or s.startswith("#"):
            continue
        # Count occurrences in non-comment lines
        count += len(re.findall(r"(?<!\S)(-f|--file)(?!\S)", ln))
    return count


def _is_unknown_reference(step: Step) -> bool:
    # Step header token format contains reference=...
    m = re.search(r"\breference=([^\s]+)", step.header)
    if not m:
        return False
    val = m.group(1).strip()
    return val.lower() == "unknown"


def _has_missing_artifact_request(step: Step) -> bool:
    hay = "\n".join(step.lines).lower()
    # Accept several common phrasings, keep it simple and robust.
    patterns = [
        r"missing file/path pattern",
        r"missing file.*pattern",
        r"attach next",
        r"name the exact missing.*pattern",
    ]
    return any(re.search(p, hay) for p in patterns)

def lint(path: Path) -> None:
    raw = _read_text(path)
    lines = raw.splitlines(True)
    fence = _extract_bash_fence(lines)
    steps = _parse_steps(fence)

    errors: List[str] = []
    for step in steps:
        attachments = _count_attachments(step)
        if attachments > 2:
            errors.append(f"Step {step.n}: has {attachments} attachments; must be <= 2 (minimal attachments rule).")

        if _is_unknown_reference(step) and not _has_missing_artifact_request(step):
            errors.append(
                f"Step {step.n}: reference=Unknown but prompt does not request missing file/path pattern(s) to attach next."
            )

    if errors:
        for e in errors:
            print(f"[ERROR] {e}", file=sys.stderr)
        sys.exit(1)

    print("[OK] Attachment lint passed (<=2 attachments per step; Unknown-reference prompts request missing patterns).")

def main() -> None:
    p = argparse.ArgumentParser(description="Lint oraclepack Stage-1 gold pack attachments (<=2 per step) and Unknown-reference prompt behavior.")
    p.add_argument("pack_path", help="Path to the Markdown pack file")
    args = p.parse_args()

    path = Path(args.pack_path)
    if not path.exists():
        print(f"[ERROR] File not found: {path}", file=sys.stderr)
        sys.exit(1)

    lint(path)


if __name__ == "__main__":
    main()
```

.config/skills/oraclepack-gold-pack/scripts/validate_pack.py
```
# path: oraclepack-gold-pack/scripts/validate_pack.py
import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple, Optional, Dict


ALLOWED_CATEGORIES = [
    "contracts/interfaces",
    "invariants",
    "caching/state",
    "background jobs",
    "observability",
    "permissions",
    "migrations",
    "UX flows",
    "failure modes",
    "feature flags",
]

REQUIRED_TOKENS = [
    "ROI=",
    "impact=",
    "confidence=",
    "effort=",
    "horizon=",
    "category=",
    "reference=",
]


@dataclass
class Step:
    n: str
    header_line_no: int
    header_line: str
    block_lines: List[str]


def _fail(errors: List[str]) -> None:
    for e in errors:
        print(f"[ERROR] {e}", file=sys.stderr)
    sys.exit(1)


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        # fallback
        return path.read_text(encoding="utf-8", errors="replace")


def _extract_single_bash_fence(lines: List[str]) -> Tuple[int, int, List[str], List[str]]:
    """
    Enforces:
      - exactly two fence lines in entire document
      - first fence line must be exactly ```bash
      - closing fence line must be exactly ```
    Returns: (open_idx, close_idx, fence_lines, outside_lines)
    """
    fence_idxs = [i for i, ln in enumerate(lines) if ln.startswith("```")]

    if len(fence_idxs) != 2:
        raise ValueError(
            f"Expected exactly 1 fenced code block (2 fence lines), found {len(fence_idxs)} fence line(s)."
        )

    open_idx, close_idx = fence_idxs
    if lines[open_idx].rstrip("\n") != "```bash":
        raise ValueError("Opening fence must be exactly ```bash (no extra tokens/spaces).")
    if lines[close_idx].rstrip("\n") != "```":
        raise ValueError("Closing fence must be exactly ``` (no extra tokens/spaces).")
    if close_idx <= open_idx:
        raise ValueError("Closing fence appears before opening fence.")

    fence_lines = lines[open_idx + 1 : close_idx]
    outside_lines = lines[:open_idx] + lines[close_idx + 1 :]
    return open_idx, close_idx, fence_lines, outside_lines


def _parse_steps(fence_lines: List[str]) -> List[Step]:
    header_re = re.compile(r"^#\s*(\d{2})\)?\s+")
    header_idxs: List[Tuple[int, str]] = []
    for i, ln in enumerate(fence_lines):
        m = header_re.match(ln)
        if m:
            header_idxs.append((i, m.group(1)))

    if len(header_idxs) != 20:
        raise ValueError(f"Expected exactly 20 step headers inside bash fence, found {len(header_idxs)}.")

    expected = [f"{i:02d}" for i in range(1, 21)]
    got = [n for _, n in header_idxs]
    if got != expected:
        raise ValueError(f"Step numbering must be sequential 01..20. Got: {got}")

    steps: List[Step] = []
    for idx, (start_i, n) in enumerate(header_idxs):
        end_i = header_idxs[idx + 1][0] if idx + 1 < len(header_idxs) else len(fence_lines)
        block = fence_lines[start_i:end_i]
        steps.append(
            Step(
                n=n,
                header_line_no=start_i + 1,  # 1-based within fence
                header_line=block[0].rstrip("\n"),
                block_lines=[b.rstrip("\n") for b in block],
            )
        )
    return steps


def _validate_header(step: Step, errors: List[str]) -> None:
    header = step.header_line

    for tok in REQUIRED_TOKENS:
        if tok not in header:
            errors.append(f"Step {step.n}: missing required token '{tok}' in header: {header}")

    # ensure each token has a value (non-empty, non-space)
    for key in ["ROI", "impact", "confidence", "effort", "horizon", "category", "reference"]:
        m = re.search(rf"\b{re.escape(key)}=([^\s]+)", header)
        if not m:
            errors.append(f"Step {step.n}: header missing/empty '{key}=' value: {header}")

    cat_m = re.search(r"\bcategory=([^\s]+(?:\s+[^\s]+)?)", header)
    # Category can contain a space (e.g., "background jobs", "UX flows", "failure modes", "feature flags")
    # The regex above captures up to two words; handle 2-word categories explicitly by matching allowed set.
    if cat_m:
        # Extract by checking allowed set presence after 'category='
        after = header.split("category=", 1)[1]
        # category value ends at next token start or end of line
        end = len(after)
        for token in [" reference=", " ROI=", " impact=", " confidence=", " effort=", " horizon="]:
            pos = after.find(token)
            if pos != -1:
                end = min(end, pos)
        cat_val = after[:end].strip()
        if cat_val not in ALLOWED_CATEGORIES:
            errors.append(
                f"Step {step.n}: invalid category='{cat_val}'. Must be one of: {ALLOWED_CATEGORIES}"
            )
    else:
        errors.append(f"Step {step.n}: could not parse category=... from header: {header}")


def _validate_write_output(step: Step, errors: List[str]) -> None:
    # Find first --write-output in the step block
    joined = "\n".join(step.block_lines)
    m = re.search(r"--write-output\s+(\"([^\"]+)\"|'([^']+)'|(\S+))", joined)
    if not m:
        errors.append(f"Step {step.n}: missing --write-output in step block.")
        return

    path = m.group(2) or m.group(3) or m.group(4) or ""
    if "/" not in path:
        errors.append(
            f"Step {step.n}: --write-output path must look like <out_dir>/{step.n}-<slug>.md; got: {path}"
        )
        return

    filename = path.split("/")[-1]
    if not filename.startswith(f"{step.n}-"):
        errors.append(
            f"Step {step.n}: --write-output filename must start with '{step.n}-'; got: {filename}"
        )
    if not filename.endswith(".md"):
        errors.append(f"Step {step.n}: --write-output filename must end with .md; got: {filename}")


def _validate_coverage_check(outside_lines: List[str], errors: List[str]) -> None:
    text = "\n".join(outside_lines)
    # Require a "Coverage check" heading (case-insensitive)
    if re.search(r"^##\s+Coverage check\s*$", text, flags=re.IGNORECASE | re.MULTILINE) is None:
        errors.append('Missing "## Coverage check" section (must be outside the bash fence).')
        return

    # Ensure every category appears as "<category>:` somewhere after the heading
[TRUNCATED]
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
[TRUNCATED]
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

.config/skills/oraclepack-taskify/scripts/detect-oracle-outputs.sh
```
#!/usr/bin/env bash
set -euo pipefail

out_dir="${1:-oracle-out}"

if [[ ! -d "${out_dir}" ]]; then
  echo "ERROR: out_dir does not exist: ${out_dir}" >&2
  exit 2
fi

shopt -s nullglob

for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
  matches=( "${out_dir}/${n}-"*.md )
  if [[ "${#matches[@]}" -eq 0 ]]; then
    echo "ERROR: missing output for prefix ${n}: expected ${out_dir}/${n}-*.md" >&2
    exit 3
  fi
  if [[ "${#matches[@]}" -gt 1 ]]; then
    echo "ERROR: multiple outputs for prefix ${n}; expected exactly one file." >&2
    printf '%s\n' "${matches[@]}" >&2
    exit 4
  fi
  printf '%s\n' "${matches[0]}"
done
```

.config/skills/oraclepack-taskify/scripts/validate-action-pack.sh
```
#!/usr/bin/env bash
set -euo pipefail

pack_path="${1:-}"
if [[ -z "${pack_path}" ]]; then
  echo "Usage: validate-action-pack.sh <path/to/oracle-actions-pack.md>" >&2
  exit 2
fi

if [[ ! -f "${pack_path}" ]]; then
  echo "ERROR: file not found: ${pack_path}" >&2
  exit 3
fi

# Rule: exactly one bash fence, and no other fences.
bash_fence_count="$(grep -cE '^[[:space:]]*```bash[[:space:]]*$' "${pack_path}" || true)"
any_fence_count="$(grep -cE '^[[:space:]]*```' "${pack_path}" || true)"

if [[ "${bash_fence_count}" -ne 1 ]]; then
  echo "ERROR: expected exactly one '```bash' fence; found ${bash_fence_count}" >&2
  exit 10
fi

if [[ "${any_fence_count}" -ne 2 ]]; then
  echo "ERROR: expected exactly 2 total fences (start+end); found ${any_fence_count}" >&2
  echo "Fences found:" >&2
  grep -nE '^[[:space:]]*```' "${pack_path}" >&2 || true
  exit 11
fi

# Extract lines within the bash fence and validate step headers.
in_bash=0
headers=()
while IFS= read -r line; do
  if [[ "${line}" =~ ^[[:space:]]*```bash[[:space:]]*$ ]]; then
    in_bash=1
    continue
  fi
  if [[ "${line}" =~ ^[[:space:]]*```[[:space:]]*$ ]]; then
    in_bash=0
    continue
  fi
  if [[ "${in_bash}" -eq 1 ]]; then
    if [[ "${line}" =~ ^#\ ([0-9]{2})\) ]]; then
      headers+=( "${BASH_REMATCH[1]}" )
    fi
  fi
done < "${pack_path}"

if [[ "${#headers[@]}" -lt 1 ]]; then
  echo "ERROR: no step headers found inside bash fence (expected '# NN)')" >&2
  exit 20
fi

# Validate strict sequential: 01..N with no gaps and no duplicates.
seen=""
expected=1
for h in "${headers[@]}"; do
  # Check duplicates
  if [[ " ${seen} " == *" ${h} "* ]]; then
    echo "ERROR: duplicate step header: ${h}" >&2
    exit 21
  fi
  seen="${seen} ${h}"

  exp="$(printf '%02d' "${expected}")"
  if [[ "${h}" != "${exp}" ]]; then
    echo "ERROR: non-sequential step header. Expected ${exp}, got ${h}" >&2
    echo "All headers: ${headers[*]}" >&2
    exit 22
  fi
  expected=$((expected + 1))
done

echo "OK: Action Pack validation passed."
```

.config/skills/oraclepack-usecase-rpg/references/oracle-pack-template.md
```
<!-- # path: oraclepack-usecase-rpg/references/oracle-pack-template.md -->
# Oracle Pack — {{codebase_name}} (RPG-infused)

## Parsed args
- codebase_name: {{codebase_name}}
- constraints: {{constraints}}
- non_goals: {{non_goals}}
- team_size: {{team_size}}
- deadline: {{deadline}}
- out_dir: {{out_dir}}
- oracle_cmd: {{oracle_cmd}}
- oracle_flags: {{oracle_flags}}
- extra_files: {{extra_files}}

## Commands
```bash
out_dir="{{out_dir}}"

# 01) ROI=... impact=... confidence=... effort=... horizon=Immediate category=contracts/interfaces reference=...
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/01-<slug>.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #01
RPG:
FunctionalNode: <Capability>::<Feature> (WHAT)
StructuralNode: <module-or-dir> :: <public surface> (HOW)
DependsOn: []
Phase: P0

Reference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>
Category: contracts/interfaces
Horizon: Immediate
ROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)
Question: <question text>
Rationale: <exactly one sentence>
Smallest experiment today: <exactly one action>

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ...
# ...
# 20) ...
```

## Coverage check

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
```

.config/skills/oraclepack-usecase-rpg/scripts/validate_oraclepack_pack.py
```
# path: oraclepack-usecase-rpg/scripts/validate_oraclepack_pack.py
from __future__ import annotations

import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import Dict, List, Tuple


REQUIRED_CATEGORIES = [
    "contracts/interfaces",
    "invariants",
    "caching/state",
    "background jobs",
    "observability",
    "permissions",
    "migrations",
    "UX flows",
    "failure modes",
    "feature flags",
]

BASH_FENCE_RE = re.compile(r"(?s)```bash\n(.*?)\n```")
STEP_HEADER_RE = re.compile(r"^#\s*(\d{2})(?:\)|[\s]+[—-])")
TOKEN_RE = {
    "roi": re.compile(r"\bROI=(\d+(?:\.\d+)?)\b"),
    "impact": re.compile(r"\bimpact=(\d+(?:\.\d+)?)\b"),
    "confidence": re.compile(r"\bconfidence=(\d+(?:\.\d+)?)\b"),
    "effort": re.compile(r"\beffort=(\d+(?:\.\d+)?)\b"),
    "horizon": re.compile(r"\bhorizon=(Immediate|Strategic)\b"),
    "category": re.compile(r"\bcategory=(contracts/interfaces|invariants|caching/state|background jobs|observability|permissions|migrations|UX flows|failure modes|feature flags)\b"),
    "reference": re.compile(r"\breference=([^\s]+)\b"),
}
ORACLE_INVOCATION_RE = re.compile(r"(?m)^\s*(oracle|\$oracle_cmd|\$\{oracle_cmd\})\b")
WRITE_OUTPUT_RE = re.compile(r'--write-output\s+["\\]?([^"\\]+)["\\]?')
PROMPT_FLAG_RE = re.compile(r"\s-\s*p\s|\s-p\s")


@dataclass
class Step:
    step_id: str  # "01"
    header_line: str
    body: str


def die(msg: str) -> None:
    print(f"ERROR: {msg}", file=sys.stderr)
    sys.exit(1)


def warn(msg: str) -> None:
    print(f"WARNING: {msg}", file=sys.stderr)


def parse_bash_block(md: str) -> str:
    matches = list(BASH_FENCE_RE.finditer(md))
    if len(matches) != 1:
        die(f"Expected exactly one fenced bash block starting with ```bash, found {len(matches)}.")
    return matches[0].group(1)


def split_steps(bash: str) -> List[Step]:
    lines = bash.splitlines(True)  # keep \n
    headers: List[Tuple[int, str, str]] = []  # (line_idx, step_id, header_line)
    for i, line in enumerate(lines):
        m = STEP_HEADER_RE.match(line.strip())
        if m:
            headers.append((i, m.group(1), line.rstrip("\n")))

    if len(headers) != 20:
        die(f"Expected exactly 20 step headers, found {len(headers)}.")

    # Ensure sequential 01..20 by appearance
    expected = [f"{i:02d}" for i in range(1, 21)]
    got = [sid for _, sid, _ in headers]
    if got != expected:
        die(f"Step IDs are not sequential 01..20 in order. Got: {got}")

    steps: List[Step] = []
    for idx, (line_i, sid, header_line) in enumerate(headers):
        start = line_i + 1
        end = headers[idx + 1][0] if idx + 1 < len(headers) else len(lines)
        body = "".join(lines[start:end]).strip("\n")
        steps.append(Step(step_id=sid, header_line=header_line, body=body))
        return steps


def parse_float(token: str, s: str, step_id: str) -> float:
    m = TOKEN_RE[token].search(s)
    if not m:
        die(f"Step {step_id}: missing token {token} in header.")
    return float(m.group(1))


def parse_str(token: str, s: str, step_id: str) -> str:
    m = TOKEN_RE[token].search(s)
    if not m:
        die(f"Step {step_id}: missing token {token} in header.")
    return m.group(1)


def validate_header_tokens(steps: List[Step]) -> Tuple[int, int, Dict[str, int]]:
    immediate = 0
    strategic = 0
    category_counts: Dict[str, int] = {c: 0 for c in REQUIRED_CATEGORIES}

    for st in steps:
        header = st.header_line

        roi = parse_float("roi", header, st.step_id)
        impact = parse_float("impact", header, st.step_id)
        confidence = parse_float("confidence", header, st.step_id)
        effort = parse_float("effort", header, st.step_id)
        horizon = parse_str("horizon", header, st.step_id)
        category = parse_str("category", header, st.step_id)
        _ = parse_str("reference", header, st.step_id)

        for name, val in [("impact", impact), ("confidence", confidence), ("effort", effort)]:
            if not (0.1 <= val <= 1.0):
                die(f"Step {st.step_id}: {name} must be in 0.1..1.0, got {val}.")
            # One decimal requirement (tolerate float formatting issues)
            if round(val, 1) != val:
                die(f"Step {st.step_id}: {name} must have one decimal, got {val}.")

        if effort == 0.0:
            die(f"Step {st.step_id}: effort must be non-zero.")

        expected_roi = (impact * confidence) / effort
        if abs(expected_roi - roi) > 0.02:
            die(
                f"Step {st.step_id}: ROI mismatch. Header ROI={roi}, expected {(impact*confidence)/effort:.2f} "
                f"from impact={impact}, confidence={confidence}, effort={effort}."
            )

        if horizon == "Immediate":
            immediate += 1
        elif horizon == "Strategic":
            strategic += 1
        else:
            die(f"Step {st.step_id}: invalid horizon={horizon}.")

        if category not in category_counts:
            die(f"Step {st.step_id}: category not in required set: {category}")
            category_counts[category] += 1

    return immediate, strategic, category_counts


def validate_step_bodies(steps: List[Step], out_dir_hint: str | None) -> None:
    for st in steps:
        body = st.body

        invocations = ORACLE_INVOCATION_RE.findall(body)
        if len(invocations) < 1:
            die(f"Step {st.step_id}: no oracle invocation detected in body.")
        if len(invocations) > 1:
            warn(f"Step {st.step_id}: multiple oracle-like invocations detected; expected one. Matches={invocations}")

        if "--write-output" not in body:
            die(f"Step {st.step_id}: missing --write-output in oracle invocation.")
        if "-p" not in body:
            die(f"Step {st.step_id}: missing -p prompt body in oracle invocation.")

        m = WRITE_OUTPUT_RE.search(body)
        if not m:
            die(f"Step {st.step_id}: could not parse --write-output path.")
        out_path = m.group(1)

        # Ensure step number is reflected in write-output file path
[TRUNCATED]
```

.config/skills/strategist-questions-oracle-pack/assets/oracle-pack-template.md
```
<!-- path: ~/.codex/skills/strategist-questions-oracle-pack/assets/oracle-pack-template.md -->
# oracle strategist question pack

---

## parsed args

- codebase_name: <Unknown|value>
- constraints: <None|value>
- non_goals: <None|value>
- team_size: <Unknown|value>
- deadline: <Unknown|value>
- out_dir: <oracle-out|value>
- oracle_cmd: <oracle|value>
- oracle_flags: <--browser-attachments always --files-report|value>
- extra_files: <empty|value>

---

## commands (exactly 20; sorted by ROI desc; ties by lower effort)

```bash
# 01 — ROI=<..> impact=<..> confidence=<..> effort=<..> horizon=<Immediate|Strategic> category=<...> reference=<...>
<oracle_cmd> \
  <oracle_flags> \
  --write-output "<out_dir>/01-<slug>.md" \
  -p "Strategist question #01
Reference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>
Category: <one of required categories>
Horizon: <Immediate|Strategic>
ROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)
Question: <question text>
Rationale: <exactly one sentence>
Smallest experiment today: <exactly one action>
Constraints: <constraints or None>
Non-goals: <non_goals or None>

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "<minimal evidence file 1>" \
  -f "<optional supporting file 2>" \
  <optional extra_files entries...>

# 02 ...
# ...
# 20 ...
```

---

## coverage check (must be satisfied)

*   contracts/interfaces: <OK|Missing (state missing artifact pattern)>

*   invariants: <OK|Missing (...)>

*   caching/state: <OK|Missing (...)>

*   background jobs: <OK|Missing (...)>

*   observability: <OK|Missing (...)>

*   permissions: <OK|Missing (...)>

*   migrations: <OK|Missing (...)>

*   UX flows: <OK|Missing (...)>

*   failure modes: <OK|Missing (...)>

*   feature flags: <OK|Missing (...)>
```
```

.config/skills/strategist-questions-oracle-pack/references/oracle-scratch-format.md
```
<!-- path: ~/.codex/skills/strategist-questions-oracle-pack/references/oracle-scratch-format.md -->
# oracle usage examples

---

## add attachments

```bash
oracle \
  --browser-attachments always \
  --browser-input-timeout 5s \
  -p "Run the UI smoke test." \
  -f "packages/"
```

## copy to clipboard

```bash
oracle --render --copy \
  -p "Explain the CLI architecture and control flow: index.ts -> registry init -> list/run commands. Identify error-handling gaps, async pitfalls, and how you'd improve UX (help text, exit codes, structured output). Provide proposed code changes." \
  -f "codefetch/packages.md" \
  -f "packages/cli/src/index.ts" \
  -f "packages/cli/src/commands/list.ts" \
  -f "packages/cli/src/commands/run.ts"
```

---

## oracle/codefetch

```bash
oracle \
  --browser-attachments always \
  -p "Using packages.md + attached files, give me a dependency graph and a build order. Then propose 5 repo hygiene improvements." \
  -f "codefetch/packages.md" \
  -f "packages/**/package.json" \
  -f "packages/**/tsconfig.json"
```

---

## other scratch examples

```bash
oracle \
  -p "Walk through the UI smoke test" \
  --file "src/**/*.ts"
```

```bash
oracle \
  -p "Review these changes and propose fixes" \
  -f "src/**/*.ts,!**/*.test.ts"
```

```bash
oracle \
  -p "Using packages.md + attached files, give me a dependency graph and a build order. Then propose 5 repo hygiene improvements." \
  -f "codefetch/packages.md" \
  -f "packages/**/package.json" \
  -f "packages/**/tsconfig.json"
```

```bash
oracle --render --copy \
  -p "Explain the CLI architecture and control flow: index.ts -> registry init -> list/run commands. Identify error-handling gaps, async pitfalls, and how you'd improve UX (help text, exit codes, structured output). Provide proposed code changes." \
  -f "codefetch/packages.md" \
  -f "packages/cli/src/index.ts" \
  -f "packages/cli/src/commands/list.ts" \
  -f "packages/cli/src/commands/run.ts"
```

```bash
oracle  \
  --prompt "Read packages.md and explain: (1) what each package does, (2) how workflows are executed end-to-end, (3) where to add a new workflow, and (4) the top 5 refactors to reduce coupling." \
  --file "codefetch/packages.md"
```

```

</source_code>