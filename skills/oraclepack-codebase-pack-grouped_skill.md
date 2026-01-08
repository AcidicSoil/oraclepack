<filetree>
Project Structure:
└── skills
    └── oraclepack-codebase-pack-grouped
        ├── references
        │   ├── attachment-minimization.md
        │   ├── codebase-grouping.md
        │   └── codebase-pack-template.md
        ├── scripts
        │   ├── generate_grouped_packs.py
        │   ├── lint_attachments.py
        │   └── validate_pack.py
        └── SKILL.md

</filetree>

<source_code>
skills/oraclepack-codebase-pack-grouped/SKILL.md
```
---
name: oraclepack-codebase-pack-grouped
description: Generate multiple runner-ingestible oraclepack Stage-1 packs grouped by codebase topic/domain (subdir + deterministic inference) with direct code attachments. Use when the user wants per-topic/per-domain mini-packs for a target repo/project/codebase instead of ticket folders, with strict 20-step schema and validation.
---

# oraclepack-codebase-pack-grouped (Stage 1)

## Goal

Produce **multiple** codebase-driven Stage-1 packs, one per inferred topic/domain, with direct code attachments. Each pack is schema-safe and self-contained.

## Use this skill

Use when the user wants separate packs per topic/domain grouped by a target repo/project/codebase, not a `.tickets/` folder.

## Inputs (parse trailing KEY=value; last-one-wins)

Supported keys (defaults in parentheses):
- `codebase_name` (`Unknown`)
- `out_dir` (`docs/oracle-questions-YYYY-MM-DD`)
- `oracle_cmd` (`oracle`)
- `oracle_flags` (`--files-report`)
- `extra_files` (empty; appended literally)
- `code_root` (`.`)
- `code_glob` (`**/*`)
- `code_paths` (empty; comma-separated explicit files; if present, ignore glob)
- `code_max_files` (`200`)
- `group_mode` (`subdir+infer`)
- `group_min_score` (`0.10`)
- `group_max_files` (`200`)
- `group_max_chars` (`200000`)
- `ignore_dirs` (empty; comma-separated; merged with defaults)
- `include_exts` (empty; uses default extension allowlist)
- `exclude_glob` (empty; comma-separated glob patterns)
- `mode` (`codebase-grouped-direct`)

Notes:
- `YYYY-MM-DD` is computed at pack generation time for default `out_dir`.
- If oracle flag support is uncertain, omit unsupported flags; never invent flags.

## Workflow (deterministic)

1) Read:
- `references/codebase-grouping.md`
- `references/attachment-minimization.md`
- `references/codebase-pack-template.md`

2) Ask user if custom args are needed (numbered picker):

```
1) Use defaults (no args)
2) Provide custom args
```

If `2`, ask for KEY=value args and run with those; otherwise run with defaults.

3) Generate packs (deterministic grouping + per-group pack files):

```bash
python3 /home/user/.codex/skills/oraclepack-codebase-pack-grouped/scripts/generate_grouped_packs.py \
  codebase_name=oraclepack \
  out_dir=docs/oracle-questions-2026-01-08
```

Outputs:
- `{{out_dir}}/packs/*.md` (one pack per group/part)
- `{{out_dir}}/_groups.json` (group -> file list)

4) Size control (mandatory; fail fast):
- Run `oracle --dry-run summary --files-report ...` for the **largest** group pack (or each pack if unsure).
- Enforce caps:
  - browser: ≤ 60,000 tokens total input per step
  - api: ≤ 180,000 tokens total input per step
- If exceeded, reduce via `group_max_files`, `code_max_files`, or `include_exts`.

5) Validate every pack (mandatory):

```bash
python3 /home/user/.codex/skills/oraclepack-codebase-pack-grouped/scripts/validate_pack.py <pack.md>
python3 /home/user/.codex/skills/oraclepack-codebase-pack-grouped/scripts/lint_attachments.py <pack.md>
```

## Failure behavior

- If no files resolve, packs still generate with empty attachments.
- Step 01 prompt must request exact missing file/path pattern(s).

## Output contract

Each pack MUST:
- Have exactly one `bash` fence
- Have exactly 20 steps (01..20)
- Include ROI header tokens
- Include `--write-output` with a group-specific `out_dir`
- Attach code files directly via `${code_args[@]}`
- End with Coverage check outside the bash fence
```

skills/oraclepack-codebase-pack-grouped/references/attachment-minimization.md
```
# Attachment minimization rules (Codebase Stage 1 — Direct Attach)

Objective: keep each group pack focused and portable.

## Code attachments

- Code files are attached directly in each step via `${code_args[@]}`.
- Use `group_max_files` (default 200) to bound per-pack file count.
- If a group is larger than the cap, split into multiple packs (part 1..N).
- Prefer code_glob + include_exts to avoid irrelevant files.

## Non-code attachments (extra_files)

- Keep explicit non-code attachments to **0–1 per step**.
- Prefer a single high-signal file (e.g., README, architecture doc).

## extra_files (literal append)

- If `extra_files` is provided, append it literally to every oracle command.
- It may include additional `-f/--file` flags.
- Place `extra_files` on its own line with a comment:
  - `# extra_files appended literally`
```

skills/oraclepack-codebase-pack-grouped/references/codebase-grouping.md
```
# Codebase grouping rules (Stage 1 — Direct Attach)

Objective: deterministically split a target codebase into topic/domain groups and produce one Stage-1 pack per group.

## Grouping behavior

- Primary grouping: by top-level subdirectory under `code_root`.
- Loose files (root-level or outside `code_root`) are assigned via token overlap (Jaccard) against existing groups.
- If no group scores above `group_min_score`, loose files fall into a `root` group.

## Determinism

- File discovery is lexicographically sorted.
- Group names are derived from directory names; sharded parts are `group_name part N`.
- Group slug is a normalized lowercase `a-z0-9-` token.

## Limits

- `code_max_files` caps total discovered files before grouping.
- `group_max_files` and `group_max_chars` cap each group pack; groups split into part 1..N.

## Exclusions

- Ignore directories include `.git`, `node_modules`, `dist`, `build`, `.venv`, and other common build outputs.
- Additional ignore names can be provided via `ignore_dirs` (comma-separated).
- Use `exclude_glob` to drop specific paths.
```

skills/oraclepack-codebase-pack-grouped/references/codebase-pack-template.md
```
# Oracle Pack — {{codebase_name}} (Grouped Codebase Stage 1 — Direct Attach)

## Parsed args
- codebase_name: {{codebase_name}}
- out_dir: {{out_dir}}
- oracle_cmd: {{oracle_cmd}}
- oracle_flags: {{oracle_flags}}
- extra_files: {{extra_files}}
- code_root: {{code_root}}
- code_glob: {{code_glob}}
- code_paths: {{code_paths}}
- code_max_files: {{code_max_files}}
- group_name: {{group_name}}
- group_slug: {{group_slug}}
- group_mode: {{group_mode}}
- group_min_score: {{group_min_score}}
- group_max_files: {{group_max_files}}
- group_max_chars: {{group_max_chars}}
- ignore_dirs: {{ignore_dirs}}
- include_exts: {{include_exts}}
- exclude_glob: {{exclude_glob}}
- mode: {{mode}}

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "{{out_dir}}/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach code files directly (no bundle dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "{{out_dir}}"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/01-contracts-interfaces-public-surface.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #01  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached code files as primary evidence, map the public surface area (CLI/TUI/API/interfaces/contracts). Call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/02-contracts-interfaces-integrations.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #02  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached code files as primary evidence, identify external integrations implied by this area; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/03-invariants-invariant-map.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #03  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached code files, map invariants and critical assumptions (data shape, ordering, idempotency, contracts). Identify the weakest or least-tested invariant.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=4.7 impact=6 confidence=0.76 effort=1 horizon=NearTerm category=caching/state reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/04-caching-state-reads-writes.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #04  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: caching/state
Horizon: NearTerm
ROI: 4.7 (impact=6, confidence=0.76, effort=1)

Question:
Using the attached code files, identify stateful reads/writes and any caches (in-memory, disk, external). Note invalidation boundaries and any silent staleness risks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=4.4 impact=6 confidence=0.70 effort=1 horizon=NearTerm category=background jobs reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/05-background-jobs-queues.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #05  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: background jobs
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.70, effort=1)

Question:
Using the attached code files, list any background jobs/queues/cron tasks. Note retries, idempotency, and failure modes.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=4.5 impact=6 confidence=0.75 effort=1 horizon=Immediate category=observability reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/06-observability-logging-metrics.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #06  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: observability
Horizon: Immediate
ROI: 4.5 (impact=6, confidence=0.75, effort=1)

Question:
Using the attached code files, identify logging/metrics/tracing in this area. Call out missing signals for debugging incidents.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 07) ROI=4.3 impact=6 confidence=0.68 effort=1 horizon=NearTerm category=permissions reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/07-permissions-authz.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #07  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: permissions
Horizon: NearTerm
ROI: 4.3 (impact=6, confidence=0.68, effort=1)

Question:
Using the attached code files, identify authorization and permission checks. Note any missing checks or implicit trust boundaries.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 08) ROI=4.2 impact=6 confidence=0.66 effort=1 horizon=NearTerm category=migrations reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/08-migrations-backfills.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #08  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: migrations
Horizon: NearTerm
ROI: 4.2 (impact=6, confidence=0.66, effort=1)

Question:
Using the attached code files, identify migrations/backfills/data-shape changes implied in this area. Note rollout risks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 09) ROI=4.1 impact=6 confidence=0.64 effort=1 horizon=NearTerm category=UX flows reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/09-ux-flows.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #09  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: UX flows
Horizon: NearTerm
ROI: 4.1 (impact=6, confidence=0.64, effort=1)

Question:
Using the attached code files, describe the main user flows in this area. Note any fragile or confusing steps.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 10) ROI=4.0 impact=6 confidence=0.62 effort=1 horizon=NearTerm category=failure modes reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/10-failure-modes.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #10  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: failure modes
Horizon: NearTerm
ROI: 4.0 (impact=6, confidence=0.62, effort=1)

Question:
Using the attached code files, enumerate likely failure modes (network, data, validation, retries). Note missing handling.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 11) ROI=3.8 impact=6 confidence=0.60 effort=1 horizon=NearTerm category=feature flags reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/11-feature-flags.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #11  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: feature flags
Horizon: NearTerm
ROI: 3.8 (impact=6, confidence=0.60, effort=1)

Question:
Using the attached code files, identify any feature flags or config toggles. Note rollout/rollback behavior and gaps.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 12) ROI=3.9 impact=6 confidence=0.62 effort=1 horizon=NearTerm category=caching/state reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/12-caching-state-consistency.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #12  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: caching/state
Horizon: NearTerm
ROI: 3.9 (impact=6, confidence=0.62, effort=1)

Question:
Using the attached code files, identify consistency boundaries (read-after-write, eventual vs strong). Note any mismatches across layers.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 13) ROI=3.7 impact=6 confidence=0.58 effort=1 horizon=MidTerm category=observability reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/13-observability-gaps.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #13  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: observability
Horizon: MidTerm
ROI: 3.7 (impact=6, confidence=0.58, effort=1)

Question:
Using the attached code files, identify observability gaps that will block triage or SLA guarantees.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 14) ROI=3.6 impact=6 confidence=0.56 effort=1 horizon=MidTerm category=permissions reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/14-permissions-gaps.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #14  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: permissions
Horizon: MidTerm
ROI: 3.6 (impact=6, confidence=0.56, effort=1)

Question:
Using the attached code files, identify authorization edge cases or privilege escalations to test.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 15) ROI=3.5 impact=6 confidence=0.54 effort=1 horizon=MidTerm category=migrations reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/15-migrations-risk.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #15  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: migrations
Horizon: MidTerm
ROI: 3.5 (impact=6, confidence=0.54, effort=1)

Question:
Using the attached code files, identify any migration risks, data backfill triggers, or state shape changes that require careful sequencing.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 16) ROI=3.4 impact=6 confidence=0.52 effort=1 horizon=MidTerm category=UX flows reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/16-ux-flow-gaps.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #16  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: UX flows
Horizon: MidTerm
ROI: 3.4 (impact=6, confidence=0.52, effort=1)

Question:
Using the attached code files, identify UX or developer flow bottlenecks; propose smallest flow test to validate.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 17) ROI=3.3 impact=6 confidence=0.50 effort=1 horizon=MidTerm category=failure modes reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/17-failure-modes-debt.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #17  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: failure modes
Horizon: MidTerm
ROI: 3.3 (impact=6, confidence=0.50, effort=1)

Question:
Using the attached code files, list failure handling debt or missing retries/rollbacks and rank by user impact.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 18) ROI=3.2 impact=6 confidence=0.48 effort=1 horizon=LongTerm category=feature flags reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/18-feature-flags-roadmap.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #18  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: feature flags
Horizon: LongTerm
ROI: 3.2 (impact=6, confidence=0.48, effort=1)

Question:
Using the attached code files, identify where staged rollouts or flags should exist but do not.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 19) ROI=3.1 impact=6 confidence=0.46 effort=1 horizon=LongTerm category=background jobs reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/19-background-jobs-scale.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #19  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: background jobs
Horizon: LongTerm
ROI: 3.1 (impact=6, confidence=0.46, effort=1)

Question:
Using the attached code files, identify long-term scaling risks in background processing or async pipelines.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 20) ROI=3.0 impact=6 confidence=0.44 effort=1 horizon=LongTerm category=contracts/interfaces reference={{group_slug}}

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''{{group_files_json}}''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/20-contracts-interfaces-roadmap.md"   "${code_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #20  (codebase-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: contracts/interfaces
Horizon: LongTerm
ROI: 3.0 (impact=6, confidence=0.44, effort=1)

Question:
Using the attached code files, identify longer-term public surface changes likely needed in this area.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"
```

Coverage check:
- contracts/interfaces: OK (01,02,20)
- invariants: OK (03)
- caching/state: OK (04,12)
- background jobs: OK (05,19)
- observability: OK (06,13)
- permissions: OK (07,14)
- migrations: OK (08,15)
- UX flows: OK (09,16)
- failure modes: OK (10,17)
- feature flags: OK (11,18)
```

skills/oraclepack-codebase-pack-grouped/scripts/generate_grouped_packs.py
```
#!/usr/bin/env python3
from __future__ import annotations

import datetime as _dt
import fnmatch
import json
import re
import sys
from pathlib import Path
from typing import Dict, Iterable, List, Tuple

STOPWORDS = {
    "the", "and", "for", "with", "from", "this", "that", "into", "over", "under", "when",
    "then", "than", "else", "only", "must", "should", "could", "would", "will", "shall",
    "repo", "repos", "code", "codebase", "project", "oraclepack", "oracle", "pack", "packs",
}

DEFAULT_IGNORE_DIRS = {
    ".git",
    ".hg",
    ".svn",
    "node_modules",
    "dist",
    "build",
    ".next",
    ".venv",
    "venv",
    "coverage",
    "target",
}

DEFAULT_INCLUDE_EXTS = {
    ".py", ".ts", ".tsx", ".js", ".jsx", ".go", ".rs", ".java", ".kt", ".cpp", ".c",
    ".h", ".hpp", ".cs", ".rb", ".php", ".swift", ".scala", ".sql", ".md", ".yaml",
    ".yml", ".json", ".toml", ".ini", ".sh", ".ps1", ".tf", ".proto",
}


def _parse_kv_args(argv: List[str]) -> Dict[str, str]:
    args: Dict[str, str] = {}
    for raw in argv:
        if "=" not in raw:
            continue
        k, v = raw.split("=", 1)
        args[k.strip()] = v.strip()
    return args


def _today() -> str:
    return _dt.date.today().isoformat()


def _slugify(s: str) -> str:
    s = s.strip().lower()
    s = re.sub(r"[^a-z0-9]+", "-", s)
    s = re.sub(r"-+", "-", s).strip("-")
    return s or "group"


def _tokenize(text: str) -> List[str]:
    text = text.lower()
    text = re.sub(r"[^a-z0-9]+", " ", text)
    toks = [t for t in text.split() if len(t) >= 3 and t not in STOPWORDS]
    return toks


def _group_by_subdir(paths: Iterable[Path], code_root: str) -> Tuple[Dict[str, List[Path]], List[Path]]:
    root = Path(code_root).resolve()
    groups: Dict[str, List[Path]] = {}
    loose: List[Path] = []
    for p in paths:
        try:
            rel = p.resolve().relative_to(root)
        except ValueError:
            loose.append(p)
            continue
        if len(rel.parts) >= 2:
            key = rel.parts[0]
            groups.setdefault(key, []).append(p)
        else:
            loose.append(p)
    return groups, loose


def _group_tokens(group_name: str, paths: Iterable[Path]) -> set:
    tokens = set(_tokenize(group_name))
    for p in paths:
        tokens.update(_tokenize(p.stem))
    return tokens


def _file_tokens(p: Path) -> set:
    toks = set(_tokenize(p.stem))
    toks.update(_tokenize(str(p.parent.name)))
    return toks


def _jaccard(a: set, b: set) -> float:
    if not a or not b:
        return 0.0
    inter = a.intersection(b)
    union = a.union(b)
    return float(len(inter)) / float(len(union))


def _collect_paths(
    code_root: str,
    code_glob: str,
    code_paths: str,
    include_exts: str,
    exclude_glob: str,
    ignore_dirs: str,
) -> List[Path]:
    if code_paths:
        parts = [p.strip() for p in code_paths.split(",") if p.strip()]
        return [Path(p) for p in parts]

    root = Path(code_root)
    if not root.exists():
        return []

    ignore = {p.strip() for p in ignore_dirs.split(",") if p.strip()}
    ignore = ignore.union(DEFAULT_IGNORE_DIRS)

    include = {e.strip().lower() for e in include_exts.split(",") if e.strip()}
    if not include_exts.strip():
        include = set(DEFAULT_INCLUDE_EXTS)

    excludes = [g.strip() for g in exclude_glob.split(",") if g.strip()]

    out: List[Path] = []
    for p in root.glob(code_glob):
        if p.is_dir():
            continue
        parts = set(p.parts)
        if parts.intersection(ignore):
            continue
        if excludes and any(fnmatch.fnmatch(str(p), g) for g in excludes):
            continue
        if include:
            ext = p.suffix.lower()
            if ext not in include:
                continue
        out.append(p)

    return out


def _cap_group(paths: List[Path], max_files: int, max_chars: int) -> List[List[Path]]:
    chunks: List[List[Path]] = []
    current: List[Path] = []
    size = 0

    for p in paths:
        p_size = 0
        try:
            p_size = p.stat().st_size
        except FileNotFoundError:
            p_size = 0

        if max_files and len(current) >= max_files:
            chunks.append(current)
            current = []
            size = 0

        if max_chars and current and size + p_size > max_chars:
            chunks.append(current)
            current = []
            size = 0

        current.append(p)
        size += p_size

    if current:
        chunks.append(current)

    return chunks


def main() -> None:
    args = _parse_kv_args(sys.argv[1:])

    codebase_name = args.get("codebase_name", "Unknown")
    out_dir = args.get("out_dir", f"docs/oracle-questions-{_today()}")
    oracle_cmd = args.get("oracle_cmd", "oracle")
    oracle_flags = args.get("oracle_flags", "--files-report")
    extra_files = args.get("extra_files", "")
    code_root = args.get("code_root", ".")
    code_glob = args.get("code_glob", "**/*")
    code_paths = args.get("code_paths", "")
    code_max_files = int(args.get("code_max_files", "200"))
    group_mode = args.get("group_mode", "subdir+infer")
    group_min_score = float(args.get("group_min_score", "0.10"))
    group_max_files = int(args.get("group_max_files", "200"))
    group_max_chars = int(args.get("group_max_chars", "200000"))
    ignore_dirs = args.get("ignore_dirs", "")
    include_exts = args.get("include_exts", "")
    exclude_glob = args.get("exclude_glob", "")
    mode = args.get("mode", "codebase-grouped-direct")

    template_path = Path(__file__).resolve().parents[1] / "references" / "codebase-pack-template.md"
    if not template_path.exists():
        raise SystemExit(f"[ERROR] Template not found: {template_path}")

    paths = _collect_paths(code_root, code_glob, code_paths, include_exts, exclude_glob, ignore_dirs)
    paths = sorted(paths, key=lambda p: str(p))
    if code_max_files and code_max_files > 0:
        paths = paths[:code_max_files]

    groups: Dict[str, List[Path]] = {}
    loose: List[Path] = []

    if "subdir" in group_mode:
        groups, loose = _group_by_subdir(paths, code_root)
    else:
        loose = list(paths)

    if "infer" in group_mode and loose:
        group_tokens = {k: _group_tokens(k, v) for k, v in groups.items()}
        for p in loose:
            best = None
            best_score = 0.0
            pt = _file_tokens(p)
            for g, gt in group_tokens.items():
                score = _jaccard(pt, gt)
                if score > best_score:
                    best_score = score
                    best = g
            if best and best_score >= group_min_score:
                groups.setdefault(best, []).append(p)
            else:
                groups.setdefault("root", []).append(p)
    else:
        if loose:
            groups.setdefault("root", []).extend(loose)

    if not groups:
        groups = {"root": []}

    out_dir_path = Path(out_dir)
    packs_dir = out_dir_path / "packs"
    packs_dir.mkdir(parents=True, exist_ok=True)

    rendered_groups: Dict[str, List[str]] = {}

    template = template_path.read_text(encoding="utf-8")
    for group_name in sorted(groups.keys()):
        files = sorted(groups[group_name], key=lambda p: str(p))
        chunks = _cap_group(files, group_max_files, group_max_chars)
        for idx, chunk in enumerate(chunks, start=1):
            part_suffix = f" part {idx}" if len(chunks) > 1 else ""
            full_group_name = f"{group_name}{part_suffix}"
            group_slug = _slugify(full_group_name)
            pack_path = packs_dir / f"{group_slug}.md"

            rendered = template
            rendered = rendered.replace("{{codebase_name}}", codebase_name)
            rendered = rendered.replace("{{out_dir}}", str(out_dir))
            rendered = rendered.replace("{{oracle_cmd}}", oracle_cmd)
            rendered = rendered.replace("{{oracle_flags}}", oracle_flags)
            rendered = rendered.replace("{{extra_files}}", extra_files)
            rendered = rendered.replace("{{code_root}}", code_root)
            rendered = rendered.replace("{{code_glob}}", code_glob)
            rendered = rendered.replace("{{code_paths}}", code_paths)
            rendered = rendered.replace("{{code_max_files}}", str(code_max_files))
            rendered = rendered.replace("{{group_name}}", full_group_name)
            rendered = rendered.replace("{{group_slug}}", group_slug)
            rendered = rendered.replace("{{group_mode}}", group_mode)
            rendered = rendered.replace("{{group_min_score}}", str(group_min_score))
            rendered = rendered.replace("{{group_max_files}}", str(group_max_files))
            rendered = rendered.replace("{{group_max_chars}}", str(group_max_chars))
            rendered = rendered.replace("{{ignore_dirs}}", ignore_dirs)
            rendered = rendered.replace("{{include_exts}}", include_exts)
            rendered = rendered.replace("{{exclude_glob}}", exclude_glob)
            rendered = rendered.replace("{{mode}}", mode)
            rendered = rendered.replace(
                "{{group_files_json}}",
                json.dumps([str(p) for p in chunk], indent=2),
            )

            pack_path.write_text(rendered, encoding="utf-8")
            rendered_groups.setdefault(full_group_name, []).append(str(pack_path))

    groups_json = {
        "code_root": code_root,
        "groups": {k: [str(p) for p in v] for k, v in groups.items()},
        "packs": rendered_groups,
    }
    (out_dir_path / "_groups.json").write_text(json.dumps(groups_json, indent=2), encoding="utf-8")


if __name__ == "__main__":
    main()
```

skills/oraclepack-codebase-pack-grouped/scripts/lint_attachments.py
```
import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple


@dataclass
class Step:
    n: str
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
        steps.append(Step(n=n, lines=fence_lines[start_i:end_i]))
    return steps


def lint(path: Path) -> None:
    raw = _read_text(path)
    lines = raw.splitlines(True)
    fence = _extract_bash_fence(lines)
    steps = _parse_steps(fence)

    errors: List[str] = []
    for step in steps:
        joined = "\n".join(step.lines)

        if "_tickets_bundle" in joined:
            errors.append(f"Step {step.n}: found '_tickets_bundle' reference (codebase packs must attach code files directly).")

        if re.search(r"mapfile\s+-t\s+__code_files\s+<\s+<\(", joined) is None:
            errors.append(f"Step {step.n}: missing mapfile code discovery stanza.")

        if re.search(r"code_args=\(\)", joined) is None or re.search(r"code_args\+\=\(\s*(-f|--file)\b", joined) is None:
            errors.append(f"Step {step.n}: missing code_args builder (code_args+=(-f \"$p\")).")

        if re.search(r"\$\{code_args\[@\]\}", joined) is None:
            errors.append(f"Step {step.n}: missing ${'{'}code_args[@]{'}'} usage in oracle invocation.")

    if errors:
        for e in errors:
            print(f"[ERROR] {e}", file=sys.stderr)
        sys.exit(1)

    print("[OK] Direct-code lint passed.")


def main() -> None:
    p = argparse.ArgumentParser(description="Lint codebase-driven Stage-1 packs (direct-code mode).")
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

skills/oraclepack-codebase-pack-grouped/scripts/validate_pack.py
```
from pathlib import Path
import runpy

COMMON = Path(__file__).resolve().parents[2] / "oraclepack-tickets-pack-common" / "scripts" / "validate_pack.py"
if not COMMON.exists():
    raise SystemExit(f"[ERROR] Shared validator not found: {COMMON}")

runpy.run_path(str(COMMON), run_name="__main__")
```

</source_code>