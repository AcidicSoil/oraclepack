<filetree>
Project Structure:
└── skills
    ├── oraclepack-codebase-pack-grouped
    │   ├── references
    │   │   ├── attachment-minimization.md
    │   │   ├── codebase-grouping.md
    │   │   └── codebase-pack-template.md
    │   ├── scripts
    │   │   ├── generate_grouped_packs.py
    │   │   ├── lint_attachments.py
    │   │   └── validate_pack.py
    │   └── SKILL.md
    ├── oraclepack-tickets-pack-common
    │   └── scripts
    │       └── validate_pack.py
    ├── oraclepack-tickets-pack-grouped
    │   ├── references
    │   │   ├── attachment-minimization.md
    │   │   ├── ticket-grouping.md
    │   │   ├── tickets-pack-template-bundle.md
    │   │   └── tickets-pack-template.md
    │   ├── scripts
    │   │   ├── generate_grouped_packs.py
    │   │   ├── lint_attachments.py
    │   │   ├── render_group_packs.py
    │   │   ├── shard_tickets.py
    │   │   ├── validate_pack.py
    │   │   └── validate_shards.py
    │   └── SKILL.md
    ├── oraclepack-codebase-pack-grouped_skill.md
    ├── oraclepack-codebase-pack-grouped_skill_usage_summary.md
    ├── oraclepack-tickets-pack-grouped_skill.md
    └── oraclepack-tickets-pack-grouped_skill_usage_summary.md

</filetree>

<source_code>
skills/oraclepack-codebase-pack-grouped_skill.md
```
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
[TRUNCATED]
```

skills/oraclepack-codebase-pack-grouped_skill_usage_summary.md
```
1) Summary

----------

* Use this skill when you want to extend Codex with a repeatable, task-specific workflow packaged as a “skill” (instructions + optional scripts/resources), rather than re-prompting the same multi-step process each time. [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

* Best fit for generating multiple, deterministic, per-domain/per-topic “mini-packs” for a codebase (grouped output + validation), using a skill folder that includes `SKILL.md` plus optional scripts/references.

    oraclepack-codebase-pack-groupe…

    [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

* Improves reliability via progressive disclosure: Codex loads only skill metadata upfront, then reads full instructions/resources only when invoked (explicitly or implicitly). [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

* Supports portability/versioning: skills are folders of instructions/scripts/resources that can be reused across compatible agents and checked into repo/user/admin scopes. [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

1) Optimal usage moments mapped to lifecycle phases

---------------------------------------------------

Discovery / Scoping: When you need fast, repeatable extraction of “what’s here” across a codebase, but want results split by domain/topic to avoid monolithic context.

Architecture / Design: When you want deterministic grouping and standardized pack structure so downstream analysis is consistent and auditable.

Implementation: When code changes frequently and you want a one-command regeneration of grouped packs with the same schema and constraints.

Testing / Validation: When you need enforceable checks (schema, attachments, size limits) before trusting generated packs or using them in CI.

Release / Rollout: When you want a stable, version-controlled artifact set (per group) to gate releases, document surfaces, or capture invariants.

Operations / Maintenance: When incidents or regressions are localized to a subsystem and you want to quickly re-slice context by directory/domain and re-run analysis consistently.

(Invocation model note: Codex can use a skill by explicit invocation (e.g., selecting from `/skills` or typing `$…`) or by implicitly selecting it when the task matches the skill description.) [OpenAI Developers](https://developers.openai.com/codex/skills/)

1) Example library (12 examples)

--------------------------------

| Phase | Goal | Inputs required (files, constraints, audience, format) | Raw prompt (before) | Optimized prompt (after) | Acceptance criteria |
| --- | --- | --- | --- | --- | --- |
| Discovery / Scoping | Generate grouped packs for an unfamiliar repo | {context} repo path; constraints on size; audience = eng; format = commands + outputs | “Make oracle packs for this repo.” | “Use $oraclepack-codebase-pack-grouped to generate Stage-1 grouped packs for {codebase\_name}. {inputs}: code\_root={code\_root}, out\_dir={out\_dir}, group\_mode=subdir+infer. {constraints}: cap group\_max\_files={group\_max\_files}, group\_max\_chars={group\_max\_chars}, include\_exts={include\_exts}. {acceptance\_criteria}: produce packs/\*.md + \_groups.json and list largest group. {format}: steps + exact commands. {deadline}: {deadline}.” | Mentions explicit skill use; provides concrete KEY=value inputs; defines caps; requests exact artifacts and where they land. |
| Discovery / Scoping | Identify public surface area per subsystem | {context} desired subsystems; constraints; audience = platform team; format = pack set + “how to run” | “Analyze the API surface.” | “Use $oraclepack-codebase-pack-grouped to split the codebase into per-domain packs. Then describe which pack(s) cover public surfaces (CLI/TUI/API). {inputs}: code\_root={code\_root}, exclude\_glob={exclude\_glob}. {format}: (1) generated pack list by group (2) recommended run order (3) what each pack will answer. {acceptance\_criteria}: each group mapped to expected surface area.” | Output ties groups to surfaces, not just “generated packs”; includes exclusions to avoid noise. |
| Architecture / Design | Tune grouping rules to match team domains | {context} desired domain map; constraints on determinism; audience = tech leads; format = updated parameters | “Group these files better.” | “Using the existing grouping approach, propose a deterministic parameter set for $oraclepack-codebase-pack-grouped that aligns with {context} domain boundaries. {inputs}: current repo layout summary, known domain dirs, known ‘loose files’. {constraints}: deterministic ordering; no manual per-file curation unless via code\_paths. {format}: recommended KEY=value args + rationale. {acceptance\_criteria}: explains how loose files resolve; avoids non-deterministic heuristics.” | Provides a concrete args profile and explains deterministic behavior for “loose” files. |
| Architecture / Design | Define skill storage/scope strategy for a mono-repo | {context} mono-repo layout; constraints on precedence; audience = tooling; format = recommended locations | “Where should we put this skill?” | “Recommend where to store this skill so teams can override safely. {inputs}: {context} repo structure + ownership boundaries. {constraints}: follow Codex skill scopes/precedence; prefer version-controlled repo skills. {format}: path recommendations and why (repo root vs nested), plus override plan. {acceptance\_criteria}: uses the documented skill locations and precedence model.” [OpenAI Developers](https://developers.openai.com/codex/skills/) | Uses Codex scope model (repo/user/admin/system) and explains override/precedence. [OpenAI Developers](https://developers.openai.com/codex/skills/) |
| Implementation | Regenerate packs after refactor with minimal noise | {context} refactor areas; constraints on excluding build outputs; audience = eng; format = commands | “Re-run packs after changes.” | “Use $oraclepack-codebase-pack-grouped to regenerate packs for {codebase\_name} focusing on {context}. {inputs}: code\_root={code\_root}, include\_exts={include\_exts}, ignore\_dirs={ignore\_dirs}, exclude\_glob={exclude\_glob}. {constraints}: code\_max\_files={code\_max\_files}; deterministic sort. {format}: exact python command + expected output paths. {acceptance\_criteria}: excludes build/venv artifacts; outputs stable pack slugs.” | Shows exact invocation, includes ignore/exclude controls, and expects stable outputs. |
| Implementation | Create a narrow “single-area” pack for a hotfix | {context} explicit paths; constraints; audience = eng; format = one pack | “Make a pack just for these files.” | “Use $oraclepack-codebase-pack-grouped but force explicit file selection. {inputs}: code\_paths={code\_paths\_csv}, out\_dir={out\_dir}, group\_mode=subdir (or infer off). {constraints}: group\_max\_files={group\_max\_files}. {format}: one generated pack plus how to validate it. {acceptance\_criteria}: pack attaches only the provided files; no extra discovery.” | Uses code\_paths to override globbing; verifies only specified files are included. |
| Testing / Validation | Validate every pack is schema-safe and directly attachable | {context} out\_dir; constraints; audience = CI/tooling; format = command list | “Check these packs.” | “For packs in {out\_dir}/packs, run the skill’s validators and report failures. {inputs}: {out\_dir}. {constraints}: must fail fast; include exact failing step(s). {format}: shell commands + pass/fail summary. {acceptance\_criteria}: every pack checked; failures actionable.” | Includes concrete validation actions and expects actionable failure output. |
| Testing / Validation | Enforce size/token caps by splitting groups | {context} largest group; constraints = token budgets; audience = eng; format = recommended caps + rerun | “This pack is too big.” | “Use $oraclepack-codebase-pack-grouped and adjust caps to meet {constraints} size limits. {inputs}: current group stats, group\_max\_files, group\_max\_chars. {format}: new KEY=value args + rerun command + expected part naming. {acceptance\_criteria}: largest group splits into part 1..N deterministically; no missing files.” | Produces a parameter change and deterministic split strategy, not vague advice. |
| Release / Rollout | Produce release-ready, versioned analysis artifacts | {context} release tag/branch; constraints on reproducibility; audience = release mgr; format = artifact checklist | “Generate release docs.” | “Generate grouped Stage-1 packs for {context} release candidate using $oraclepack-codebase-pack-grouped. {inputs}: code\_root={code\_root}, out\_dir={out\_dir}. {constraints}: reproducible outputs; no env-specific paths beyond out\_dir. {format}: artifact checklist (packs + \_groups.json) + how to run them in order. {acceptance\_criteria}: re-running yields same pack set and slugs for same tree.” | Emphasizes reproducibility and a concrete artifact set. |
| Release / Rollout | Gate release on invariants/contract deltas per domain | {context} “what changed”; constraints; audience = eng/release; format = pack mapping + run plan | “Tell me if contracts changed.” | “Use $oraclepack-codebase-pack-grouped to slice by domain, then outline how each pack will be used to detect contracts/interfaces changes across {context}. {inputs}: baseline vs head refs (described), out\_dir. {format}: (1) which groups matter (2) run sequence (3) what output files to diff. {acceptance\_criteria}: explicit diff targets per group; minimal false positives via excludes.” | Produces a concrete diff plan tied to group outputs. |
| Operations / Maintenance | Incident triage localized to one subsystem | {context} incident area; constraints = speed; audience = oncall; format = narrow regeneration + run | “Help me debug this module.” | “Use $oraclepack-codebase-pack-grouped to generate packs limited to {context} subsystem. {inputs}: code\_root={code\_root}, exclude\_glob={exclude\_glob}, include\_exts={include\_exts}. {constraints}: smaller caps for speed. {format}: exact command + identify which pack(s) to run first. {acceptance\_criteria}: minimal scope; outputs point to likely failure modes and missing signals.” | Keeps scope tight and yields a prioritized run order for the most relevant packs. |
| Operations / Maintenance | Maintain the skill as a reusable capability across teams | {context} ownership model; constraints = portability; audience = platform; format = maintenance playbook | “How do we keep this skill usable?” | “Propose a maintenance playbook for this skill folder so it stays portable and discoverable. {inputs}: {context} team boundaries + repo layout. {constraints}: skill must remain a folder of instructions/scripts/resources; version-controlled; documented invocation. {format}: checklist for updates + review gates. {acceptance\_criteria}: aligns to Agent Skills concept (discoverable folder bundle) and Codex skill loading/scopes.” [Agent Skills+1](https://agentskills.io/home) | Playbook explicitly reflects the “skill folder” model and Codex scope/precedence rules. [Agent Skills+1](https://agentskills.io/home) |

1) Non-goals / anti-patterns (do not use this skill)

----------------------------------------------------

1. Ticket-driven packaging: if the primary inputs are `.tickets/` or issue threads (not repo code), use a ticket-pack skill instead of a codebase-grouped skill.

    oraclepack-codebase-pack-groupe…

2. One-off questions where a full grouped pack set is overkill (e.g., “what does this function do?”).

3. Non-deterministic or manual grouping requirements (e.g., “group by whatever seems interesting today”) where reproducibility is required.

4. Situations where you cannot provide/attach or point to the underlying codebase paths (no evidence to pack).

5. Tasks that primarily require editing/building code rather than generating/validating grouped analysis artifacts (use a coding/implementation workflow instead of a packaging workflow).

---
```

skills/oraclepack-tickets-pack-grouped_skill.md
```
<filetree>
Project Structure:
└── skills
    └── oraclepack-tickets-pack-grouped
        ├── references
        │   ├── attachment-minimization.md
        │   ├── ticket-grouping.md
        │   ├── tickets-pack-template-bundle.md
        │   └── tickets-pack-template.md
        ├── scripts
        │   ├── generate_grouped_packs.py
        │   ├── lint_attachments.py
        │   ├── render_group_packs.py
        │   ├── shard_tickets.py
        │   ├── validate_pack.py
        │   └── validate_shards.py
        └── SKILL.md

</filetree>

<source_code>
skills/oraclepack-tickets-pack-grouped/SKILL.md
```
---
name: oraclepack-tickets-pack-grouped
description: Generate multiple runner-ingestible oraclepack Stage-1 packs grouped by ticket topic/domain (subdir + deterministic inference) with direct ticket attachments. Use when the user wants per-topic/per-domain mini-packs, grouped via subdirectory discovery and inferred assignment of loose tickets, with strict 20-step schema and validation.
---

# oraclepack-tickets-pack-grouped (Stage 1)

## Goal

Produce **multiple** ticket-driven Stage-1 packs, one per inferred topic/domain, with direct ticket attachments. Each pack is schema-safe and self-contained.

## Use this skill

Use when the user wants separate packs per topic/domain, grouped by `.tickets/` subdirectories plus deterministic inference for loose tickets.

## Inputs (parse trailing KEY=value; last-one-wins)

Supported keys (defaults in parentheses):
- `codebase_name` (`Unknown`)
- `out_dir` (`docs/oracle-questions-YYYY-MM-DD`)
- `oracle_cmd` (`oracle`)
- `oracle_flags` (`--files-report`)
- `extra_files` (empty; appended literally)
- `ticket_root` (`.tickets`)
- `ticket_glob` (`**/*.md`)
- `ticket_paths` (empty; comma-separated explicit files; if present, ignore glob)
- `ticket_max_files` (`25`)
- `group_mode` (`subdir+infer`)
- `group_min_score` (`0.08`)
- `group_max_files` (`25`)
- `group_max_chars` (`200000`)
- `dedupe_mode` (`report`)
- `dedupe_jaccard` (`0.55`)
- `dedupe_overlap_hi` (`0.80`)
- `dedupe_overlap_lo` (`0.70`)
- `dedupe_delta_min` (`0.15`)
- `dedupe_body_chars` (`2000`)
- `mode` (`tickets-grouped-direct`)

Notes:
- `YYYY-MM-DD` is computed at pack generation time for default `out_dir`.
- If oracle flag support is uncertain, omit unsupported flags; never invent flags.

## Workflow (deterministic)

1) Read:
- `references/ticket-grouping.md`
- `references/attachment-minimization.md`
- `references/tickets-pack-template.md`

2) Ask user if custom args are needed (numbered picker):

```
1) Use defaults (no args)
2) Provide custom args
```

If `2`, ask for KEY=value args and run with those; otherwise run with defaults.

3) Generate packs (deterministic grouping + per-group pack files):

```bash
python3 /home/user/.codex/skills/oraclepack-tickets-pack-grouped/scripts/generate_grouped_packs.py \
  codebase_name=oraclepack \
  out_dir=docs/oracle-questions-2026-01-08
```

Outputs:
- `{{out_dir}}/packs/*.md` (one pack per group/part)
- `{{out_dir}}/_groups.json` (group -> ticket list)

4) Size control (mandatory; fail fast):
- Run `oracle --dry-run summary --files-report ...` for the **largest** group pack (or each pack if unsure).
- Enforce caps:
  - browser: ≤ 60,000 tokens total input per step
  - api: ≤ 180,000 tokens total input per step
- If exceeded, reduce via `group_max_files` or use explicit `ticket_paths`.

5) Validate every pack (mandatory):

```bash
python3 /home/user/.codex/skills/oraclepack-tickets-pack-grouped/scripts/validate_pack.py <pack.md>
python3 /home/user/.codex/skills/oraclepack-tickets-pack-grouped/scripts/lint_attachments.py <pack.md>
```

## Sharded packs workflow (topic/domain mini-packs)

Use this when you want a manifest-driven, sharded pack per topic/domain with bundle attachments:

First ask the user which args mode to use:

```
1) Use defaults (no args)
2) Provide custom args
```

If `2`, collect args and use them in the commands below.

```bash
python3 /home/user/.codex/skills/oraclepack-tickets-pack-grouped/scripts/shard_tickets.py \\
  --ticket-root .tickets \\
  --out-dir docs/oracle-questions-sharded

python3 /home/user/.codex/skills/oraclepack-tickets-pack-grouped/scripts/render_group_packs.py \\
  --manifest docs/oracle-questions-sharded/manifest.json \\
  --out-dir docs/oracle-questions-sharded

python3 /home/user/.codex/skills/oraclepack-tickets-pack-grouped/scripts/validate_shards.py \\
  --manifest docs/oracle-questions-sharded/manifest.json
```

## Failure behavior

- If no tickets resolve, packs still generate with empty attachments.
- Step 01 prompt must request exact missing ticket file/path pattern(s).

## Output contract

Each pack MUST:
- Have exactly one `bash` fence
- Have exactly 20 steps (01..20)
- Include ROI header tokens
- Include `--write-output` with a group-specific `out_dir`
- Attach tickets directly via `${ticket_args[@]}`
- End with Coverage check outside the bash fence
```

skills/oraclepack-tickets-pack-grouped/references/attachment-minimization.md
```
# Attachment minimization rules (Grouped Tickets Stage 1 — Direct Attach)

Objective: keep each group pack focused and portable.

## Ticket attachments

- Ticket files are attached directly in each step via `${ticket_args[@]}`.
- Use `group_max_files` (default 25) to bound per-pack ticket count.
- If a group is larger than the cap, split into multiple packs (part 1..N).

## Non-ticket attachments (repo evidence)

- Keep explicit non-ticket attachments to **0–1 per step**.
- Prefer a single high-signal file that clarifies contracts or a key code path.

## extra_files (literal append)

- If `extra_files` is provided, append it literally to every oracle command.
- It may include additional `-f/--file` flags.
- Place `extra_files` on its own line with a comment:
  - `# extra_files appended literally`

```

skills/oraclepack-tickets-pack-grouped/references/ticket-grouping.md
```
# Ticket grouping (deterministic, inferred)

Objective: split tickets into focused topic/domain groups and generate one pack per group.

## Inputs

- `ticket_root` (default `.tickets`)
- `ticket_glob` (default `**/*.md`, relative to `ticket_root`)
- `ticket_paths` (optional; comma-separated explicit files; if present, ignore `ticket_glob`)
- `group_mode` (default `subdir+infer`)
- `group_min_score` (default `0.08`)
- `group_max_files` (default `25`; max tickets per pack; >0)
- `group_max_chars` (default `200000`; max total chars per pack; >0)
- `dedupe_mode` (default `report`; one of `off`, `report`, `prune`, `merge`)
- `dedupe_jaccard` (default `0.55`)
- `dedupe_overlap_hi` (default `0.80`)
- `dedupe_overlap_lo` (default `0.70`)
- `dedupe_delta_min` (default `0.15`)
- `dedupe_body_chars` (default `2000`)

## Deterministic grouping rules

1) Collect tickets:
- If `ticket_paths` is non-empty: split on commas, trim whitespace, use exactly that list.
- Else: glob `ticket_root/ticket_glob`.
- Always sort lexicographically by path string.

2) Detect possible duplicates (if `dedupe_mode != off`):
- Signature: filename stem + first heading + first `dedupe_body_chars` chars.
- Compute `jaccard` + `overlap` between tickets.
- Duplicate edge rule:
  - `overlap >= dedupe_overlap_hi` OR (`jaccard >= dedupe_jaccard` AND `overlap >= dedupe_overlap_lo`)
- Connected components become duplicate clusters.
- Canonical: largest content length; tie-break lexicographic.
- Delta vs redundant:
  - delta if unique token ratio >= `dedupe_delta_min` OR heading differs materially.
  - redundant otherwise.

3) Seed groups by subdir:
- For any path under `ticket_root/<group>/...`, assign to group `<group>`.
- Tickets directly under `ticket_root/` are "loose".

4) Infer loose tickets into groups (if any groups exist):
- Build a token set for each group from:
  - group name tokens
  - ticket filenames (stem tokens)
  - first Markdown heading line (if present)
- For each loose ticket, compute Jaccard overlap score with each group token set.
- If `max_score >= group_min_score`, assign to the best group (stable tie-break by group name).
- Otherwise, assign to `misc`.

5) If no groups exist:
- Put all tickets into a single group named `root`.

6) Merge duplicates into primary group:
- `report`: attach all tickets in the cluster to the canonical’s group.
- `prune`: attach canonical + delta only; drop redundant from attachments.
- `merge`: create `out_dir/_ticket_merges/cluster-XXXX.md` and attach only the merged file.
- Emit `_dupes_possible.json`, `_duplicates.json`, and `_dedupe_plan.json`.

7) Split oversized groups:
- If a group exceeds `group_max_files` or `group_max_chars`, split into parts (1..N)
  in sorted order, chunked deterministically.

Hard rule: do not use mtimes, file sizes, or external ML services.

## Required outputs

- `_groups.json`: mapping of group -> list of ticket paths (lexicographic order)
- Pack file per group (and part), each self-contained and direct-attach
- `manifest.json`: groups with pack path + attached vs original ticket lists
```

skills/oraclepack-tickets-pack-grouped/references/tickets-pack-template-bundle.md
```
# Oracle Pack — {{codebase_name}} (Tickets Stage 1)

## Parsed args
- codebase_name: {{codebase_name}}
- out_dir: {{out_dir}}
- oracle_cmd: {{oracle_cmd}}
- oracle_flags: {{oracle_flags}}
- extra_files: {{extra_files}}
- ticket_root: {{ticket_root}}
- ticket_glob: {{ticket_glob}}
- ticket_paths: {{ticket_paths}}
- ticket_bundle_path: {{ticket_bundle_path}}
- mode: {{mode}}

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "{{out_dir}}/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- `## Coverage check` MUST be outside the bash fence (after the closing ```).

```bash
# Prelude (allowed inside the single bash fence)
# - Creates out_dir deterministically
# - Builds ticket_bundle_path deterministically from ticket_root/ticket_glob OR ticket_paths
# - Uses lexicographic ordering only (no mtime/timestamps)

set -euo pipefail

mkdir -p "{{out_dir}}"

python3 - <<'PY'
from __future__ import annotations

import sys
from pathlib import Path

CODEBASE_NAME = "{{codebase_name}}"
OUT_DIR = Path("{{out_dir}}")
TICKET_ROOT = Path("{{ticket_root}}")
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS_RAW = "{{ticket_paths}}".strip()
BUNDLE_PATH = Path("{{ticket_bundle_path}}")

def _read_text(p: Path) -> str:
    return p.read_text(encoding="utf-8", errors="replace")

def _title_from_md(text: str) -> str:
    for ln in text.splitlines():
        s = ln.strip()
        if s.startswith("# "):
            return s[2:].strip() or "Untitled"
    for ln in text.splitlines():
        s = ln.strip()
        if s:
            return s[:80]
    return "Untitled"

def _select_paths() -> list[Path]:
    if TICKET_PATHS_RAW:
        items = [Path(x.strip()) for x in TICKET_PATHS_RAW.split(",") if x.strip()]
        items = sorted(items, key=lambda p: str(p))
        return items

    if not TICKET_ROOT.exists():
        return []

    items = sorted(TICKET_ROOT.glob(TICKET_GLOB), key=lambda p: str(p))
    return items

paths = _select_paths()

BUNDLE_PATH.parent.mkdir(parents=True, exist_ok=True)

lines: list[str] = []
lines.append(f"# Tickets Bundle — {CODEBASE_NAME if CODEBASE_NAME else 'Unknown'}")
lines.append("")
lines.append("## Selection")
lines.append(f"- ticket_root: {TICKET_ROOT}")
lines.append(f"- ticket_glob: {TICKET_GLOB}")
lines.append(f"- ticket_paths: {TICKET_PATHS_RAW if TICKET_PATHS_RAW else '(none)'}")
lines.append("- ordering: lexicographic by path")
lines.append("")

if not paths:
    warn = (
        "## WARNING: No tickets found\n\n"
        "No ticket files were selected.\n\n"
        "What was attempted:\n"
        f"- ticket_root: {TICKET_ROOT}\n"
        f"- ticket_glob: {TICKET_GLOB}\n"
        f"- ticket_paths: {TICKET_PATHS_RAW if TICKET_PATHS_RAW else '(none)'}\n\n"
        "Next: provide explicit ticket_paths or create tickets under ticket_root.\n"
    )
    lines.append(warn)
    print(f"[WARN] No tickets selected; bundle will contain only WARNING.", file=sys.stderr)
else:
    lines.append("## Tickets")
    lines.append("")
    for p in paths:
        lines.append("---")
        lines.append(f"### {_title_from_md(_read_text(p))}")
        lines.append(f"- path: {p}")
        lines.append("")
        try:
            txt = _read_text(p)
        except Exception as e:
            lines.append(f"[ERROR reading file: {e}]")
            lines.append("")
            continue

        # Simple truncation policy: keep first 4000 chars if large.
        if len(txt) > 4000:
            lines.append(txt[:4000])
            lines.append("\n[... truncated ...]\n")
        else:
            lines.append(txt)

        lines.append("")

BUNDLE_PATH.write_text("\n".join(lines).rstrip() + "\n", encoding="utf-8")
print(f"[OK] Wrote ticket bundle: {BUNDLE_PATH}")
PY

# 01) ROI=8.0 impact=9 confidence=0.9 effort=1 horizon=Immediate category=contracts/interfaces reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/01-contracts-interfaces-surface.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #01
Category: contracts/interfaces

Using the attached tickets bundle as the primary evidence, identify the primary public interface(s) implied by the tickets (CLI commands, APIs, file contracts, or user workflows).
For each interface:
- list key inputs/outputs
- list the exact files/modules likely defining it (if unknown, say Unknown)

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=7.8 impact=8 confidence=0.9 effort=1 horizon=Immediate category=contracts/interfaces reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/02-contracts-interfaces-dependencies.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #02
Category: contracts/interfaces

From the attached tickets bundle, infer which external dependencies/services the system must integrate with (CLIs, APIs, SaaS, databases).
For each dependency:
- what contract is required (auth, endpoints, file formats)
- what configuration surface is implied

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=7.6 impact=8 confidence=0.85 effort=2 horizon=Immediate category=invariants reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/03-invariants-must-always-hold.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #03
Category: invariants

Based on the attached tickets bundle, list the invariants that must always hold (data constraints, ordering constraints, security invariants, idempotency).
For each invariant:
- what breaks if violated
- where it should be enforced (layer/module; if unknown, Unknown)

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=7.2 impact=8 confidence=0.8 effort=2 horizon=Immediate category=invariants reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/04-invariants-input-validation.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #04
Category: invariants

Using the attached tickets bundle, identify what inputs must be validated (CLI args, config fields, payloads, file paths).
For each input:
- validation rules implied
- failure message/behavior implied

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=7.0 impact=7 confidence=0.85 effort=2 horizon=Near category=caching/state reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/05-caching-state-state-model.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #05
Category: caching/state

From the attached tickets bundle, infer what state must be persisted or cached (files, DB, in-memory, remote).
For each state item:
- read/write lifecycle
- consistency model implied
- failure recovery requirements

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=6.8 impact=7 confidence=0.8 effort=2 horizon=Near category=caching/state reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/06-caching-state-cache-invalidation.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #06
Category: caching/state

Using the attached tickets bundle, identify caching risks: staleness, invalidation, keying, or race conditions implied by the tickets.
Propose a minimal caching strategy consistent with the tickets.

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 07) ROI=6.9 impact=8 confidence=0.75 effort=3 horizon=Near category=background jobs reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/07-background-jobs-what-runs-async.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #07
Category: background jobs

From the attached tickets bundle, determine what work should run asynchronously/background (schedulers, queues, cron, long-running tasks).
For each job:
- trigger
- inputs/outputs
- retry/backoff requirements

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 08) ROI=6.6 impact=7 confidence=0.75 effort=3 horizon=Near category=background jobs reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/08-background-jobs-idempotency.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #08
Category: background jobs

Using the attached tickets bundle, list the idempotency and concurrency constraints implied for background jobs.
Recommend minimal safeguards (dedupe keys, locks, at-least-once handling) aligned with tickets.

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 09) ROI=7.4 impact=8 confidence=0.8 effort=2 horizon=Immediate category=observability reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/09-observability-logs-metrics-traces.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #09
Category: observability

From the attached tickets bundle, infer required observability: logs, metrics, traces, and user-visible diagnostics.
List:
- what to log/measure
- cardinality risks
- minimal dashboards/alerts implied

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 10) ROI=7.0 impact=7 confidence=0.8 effort=2 horizon=Near category=observability reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/10-observability-error-taxonomy.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #10
Category: observability

Using the attached tickets bundle, define an error taxonomy consistent with ticket failure modes:
- user errors vs system errors
- retryable vs non-retryable
- how errors should surface (CLI exit codes, UI states, logs)

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 11) ROI=7.6 impact=9 confidence=0.75 effort=3 horizon=Immediate category=permissions reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/11-permissions-authz-model.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #11
Category: permissions

From the attached tickets bundle, infer the permissions model (roles, capabilities, scopes).
List:
- what operations require permissions
- how permissions are granted/revoked
- audit requirements implied

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 12) ROI=7.0 impact=8 confidence=0.75 effort=3 horizon=Near category=permissions reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/12-permissions-secret-handling.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #12
Category: permissions

Using the attached tickets bundle, identify sensitive data/secret handling needs.
Recommend:
- where secrets come from (env, files, vault)
- redaction rules
- least-privilege defaults

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

[TRUNCATED]
```

skills/oraclepack-tickets-pack-grouped_skill_usage_summary.md
```
1) Summary

----------

* Use this skill to generate multiple **runner-ingestible “Stage-1” oracle packs** from `.tickets/`, split into focused **topic/domain mini-packs** with **direct ticket attachments** (no bundle dependency).

    oraclepack-tickets-pack-grouped…

* Improves reliability by enforcing a **deterministic workflow**, a strict **20-step pack schema**, and mandatory **pack + attachment lint validation**.

    oraclepack-tickets-pack-grouped…

* Reduces noise by optionally detecting **possible duplicate tickets** (Jaccard + overlap) and applying a chosen dedupe mode (`report/prune/merge`).

    oraclepack-tickets-pack-grouped…

* Fits the Agent Skills model: skills package instructions/resources/scripts so agents can load procedural knowledge “on demand” and execute repeatable workflows consistently. [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

1) Optimal usage moments (mapped to lifecycle phases)

-----------------------------------------------------

Intake / Backlog triage
Use when you have many tickets (often across subdirs) and need them organized into coherent, runnable analysis packs per domain/topic. The skill’s deterministic grouping (`subdir+infer`) and optional dedupe are designed for this.

oraclepack-tickets-pack-grouped…

Requirements / Discovery
Use when you need to derive interfaces, invariants, state, caching, etc. from tickets as primary evidence, but want results separated by domain to avoid cross-contamination. The produced packs standardize the question set across 20 steps/categories.

oraclepack-tickets-pack-grouped…

Architecture / Design
Use when teams want per-area “design brief” outputs (contracts, validation boundaries, rollout flags) generated consistently for each domain group, enabling parallel review.

oraclepack-tickets-pack-grouped…

Implementation planning / Task execution setup
Use when you’re preparing an agent workflow that depends on stable, validated artifacts (pack files + manifests) and needs predictable inputs/outputs. This aligns with “skills as packaged workflows + scripts/resources.” [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

QA / Validation gates
Use when you need a guardrail that packs are schema-safe before execution (exact fence expectations, step count, direct-ticket attachment checks).

oraclepack-tickets-pack-grouped…

Release / Ops / Maintenance
Use when tickets evolve continuously and you need regeneration with deterministic behavior (lexicographic ordering, stable tie-breaks, explicit caps) to keep outputs comparable across runs.

oraclepack-tickets-pack-grouped…

1) Example library

------------------

| Phase | Goal | Inputs required (files, constraints, audience, format) | Raw prompt (before) | Optimized prompt (after) | Acceptance criteria |
| --- | --- | --- | --- | --- | --- |
| Intake / triage | Split `.tickets/` into domain packs automatically | Files: `.tickets/**/*.md`; Constraints: deterministic grouping; Format: packs + manifest | “Group my tickets and make oracle packs.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets ticket\_glob=\*\*/_.md group\_mode=subdir+infer dedupe\_mode=report out\_dir={out\_dir}. {constraints}: group\_max\_files=25 group\_max\_chars=200000. {format}: produce packs/_.md + \_groups.json + manifest.json. {acceptance\_criteria}: deterministic grouping, no bundle attachments, packs validate. {deadline}” | Multiple packs created per group; `_groups.json` + `manifest.json` exist; validation/lint passes. |
| Intake / triage | Detect possible duplicates before grouping | Files: `.tickets/**/*.md`; Constraints: report-only; Audience: maintainers | “Find duplicate tickets.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets dedupe\_mode=report dedupe\_jaccard=0.55 dedupe\_overlap\_hi=0.80 dedupe\_overlap\_lo=0.70 dedupe\_body\_chars=2000. {constraints}: do not drop tickets; emit duplicate reports. {format}: ensure \_dupes\_possible.json + \_duplicates.json + \_dedupe\_plan.json. {acceptance\_criteria}: clusters + canonical + delta/redundant identified. {deadline}” | Duplicate artifacts emitted; clusters include canonical and classification; no tickets lost from assignments. |
| Requirements / discovery | Generate per-domain “public surface changes” outputs | Files: `.tickets/**`; Audience: product/eng; Format: per-pack outputs | “Answer strategist Q1 from these tickets.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets group\_mode=subdir+infer out\_dir={out\_dir}. {constraints}: each pack must be 20 steps; direct `-f` ticket attachments only. {format}: one pack per domain producing 01–20 outputs under each group out\_dir. {acceptance\_criteria}: step 01 answers contracts/interfaces with evidence-cited bullets. {deadline}” | Packs exist per domain; each pack contains step 01 prompt; outputs pathing is group-specific. |
| Requirements / discovery | Keep groups small enough for model limits | Files: `.tickets/**`; Constraints: strict caps | “Make sure it fits context windows.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets group\_max\_files=15 group\_max\_chars=120000. {constraints}: enforce size control; split oversize groups into -part-XX packs. {format}: packs per group/part + manifest updated. {acceptance\_criteria}: no group exceeds caps; parts are deterministically chunked. {deadline}” | Large groups are split; pack names include part suffix; manifest records attached\_paths per part. |
| Architecture / design | Produce validation boundary plan per domain | Files: `.tickets/**`; Audience: platform team | “Propose validation boundaries.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets out\_dir={out\_dir}. {constraints}: keep non-ticket attachments to minimum; direct tickets only. {format}: ensure step 04 (validation boundaries) exists in each domain pack. {acceptance\_criteria}: step 04 returns boundaries + minimal plan + one concrete experiment. {deadline}” | Each pack includes step 04; step output format matches required 4 sections; content is domain-scoped. |
| Architecture / design | Produce caching/state model and cache invalidation per domain | Files: `.tickets/**`; Audience: infra | “Analyze caching and state.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets out\_dir={out\_dir}. {constraints}: preserve determinism; use dedupe\_mode=prune to reduce redundancy. {format}: ensure steps 05 and 06 are present and write outputs. {acceptance\_criteria}: defines cache keys + invalidation + correctness risks; evidence-cited. {deadline}” | Step 05/06 exist for every pack; dedupe reduces redundant attachments; outputs are written per group. |
| Implementation planning | Generate packs for only a subset of tickets (explicit list) | Files: explicit paths list; Constraint: ignore glob | “Only use these tickets.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_paths={ticket\_list\_csv} (comma-separated), ticket\_root=.tickets, ticket\_glob ignored. {constraints}: lexicographic ordering; ticket\_max\_files={n}. {format}: packs generated from explicit list only. {acceptance\_criteria}: no other tickets included; manifest lists only provided paths. {deadline}” | Packs include only the explicit files; order is stable; validation passes. |
| Implementation planning | Append consistent non-ticket evidence files to each oracle call | Files: repo evidence file(s); Constraint: 0–1 extra per step | “Add these extra files everywhere.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets extra\_files={extra\_files\_flags}. {constraints}: extra\_files appended literally, on its own line comment; keep explicit non-ticket attachments minimal. {format}: regenerate packs and confirm extra\_files included in each step invocation. {acceptance\_criteria}: every step includes extra\_files line; direct tickets still attached. {deadline}” | Every step includes extra files exactly as provided; no bundle references introduced; lint passes. |
| QA / validation | Validate every generated pack before running | Files: generated packs; Constraint: fail fast | “Make sure the packs are valid.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: out\_dir={out\_dir}. {constraints}: run validate\_pack.py and lint\_attachments.py on every packs/\*.md; fail on first error. {format}: report pass/fail per pack with file paths. {acceptance\_criteria}: all packs pass schema + direct-attach lint. {deadline}” | Validator and linter run; failures identify exact step and issue; all packs pass or errors are actionable. |
| QA / validation | Enforce “no bundle dependency” in direct-ticket mode | Files: packs/\*.md | “Make sure it doesn’t use a tickets bundle.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: out\_dir={out\_dir}. {constraints}: direct-ticket mode only; reject any `_tickets_bundle` references. {format}: lint output listing any violating steps/lines. {acceptance\_criteria}: zero `_tickets_bundle` matches across packs. {deadline}” | Lint reports zero bundle references; any violation is pinpointed to step and offending line. |
| Release / ops | Regenerate packs deterministically on a schedule (same grouping) | Files: `.tickets/**`; Constraint: stable outputs | “Rebuild the packs after ticket changes.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets out\_dir={out\_dir}. {constraints}: deterministic ordering only; no mtime/size-based heuristics; stable tie-breaks. {format}: regenerate packs + update manifest; keep group slugs stable except when splits change. {acceptance\_criteria}: repeated runs with unchanged tickets produce identical pack content. {deadline}” | With unchanged inputs, outputs are byte-stable (or explainable diffs only); grouping remains consistent. |
| Release / ops | Switch dedupe strategy to reduce attachment load | Files: `.tickets/**`; Constraint: minimize attachments | “Reduce duplicates in packs.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets dedupe\_mode=merge dedupe\_body\_chars=2000. {constraints}: create \_ticket\_merges/cluster-XXXX.md and attach merged file only. {format}: manifest must show attached\_paths reflect merge outputs. {acceptance\_criteria}: merged clusters created; redundant tickets no longer attached directly. {deadline}” | Merge files exist; packs attach merge outputs; dedupe plan reflects merge mode; validation passes. |

1) Non-goals / anti-patterns (do NOT use this skill)

----------------------------------------------------

1. You want a single monolithic pack across all tickets (use a non-grouped tickets pack skill instead).

    oraclepack-tickets-pack-grouped…

2. You need non-deterministic or ML/service-based clustering (this skill explicitly requires deterministic rules and no external ML services).

    oraclepack-tickets-pack-grouped…

3. You want to rewrite ticket contents or author new tickets as the primary output (this skill’s contract is pack generation + validation).

    oraclepack-tickets-pack-grouped…

4. You can’t tolerate the fixed “Stage-1 / 20-step” schema (this skill enforces a strict 20-step contract and validation).

    oraclepack-tickets-pack-grouped…

5. You need interactive, ad-hoc investigation rather than a repeatable workflow package (skills are meant to standardize repeatable workflows via instructions/scripts/resources). [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

---
```

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

skills/oraclepack-tickets-pack-grouped/SKILL.md
```
---
name: oraclepack-tickets-pack-grouped
description: Generate multiple runner-ingestible oraclepack Stage-1 packs grouped by ticket topic/domain (subdir + deterministic inference) with direct ticket attachments. Use when the user wants per-topic/per-domain mini-packs, grouped via subdirectory discovery and inferred assignment of loose tickets, with strict 20-step schema and validation.
---

# oraclepack-tickets-pack-grouped (Stage 1)

## Goal

Produce **multiple** ticket-driven Stage-1 packs, one per inferred topic/domain, with direct ticket attachments. Each pack is schema-safe and self-contained.

## Use this skill

Use when the user wants separate packs per topic/domain, grouped by `.tickets/` subdirectories plus deterministic inference for loose tickets.

## Inputs (parse trailing KEY=value; last-one-wins)

Supported keys (defaults in parentheses):
- `codebase_name` (`Unknown`)
- `out_dir` (`docs/oracle-questions-YYYY-MM-DD`)
- `oracle_cmd` (`oracle`)
- `oracle_flags` (`--files-report`)
- `extra_files` (empty; appended literally)
- `ticket_root` (`.tickets`)
- `ticket_glob` (`**/*.md`)
- `ticket_paths` (empty; comma-separated explicit files; if present, ignore glob)
- `ticket_max_files` (`25`)
- `group_mode` (`subdir+infer`)
- `group_min_score` (`0.08`)
- `group_max_files` (`25`)
- `group_max_chars` (`200000`)
- `dedupe_mode` (`report`)
- `dedupe_jaccard` (`0.55`)
- `dedupe_overlap_hi` (`0.80`)
- `dedupe_overlap_lo` (`0.70`)
- `dedupe_delta_min` (`0.15`)
- `dedupe_body_chars` (`2000`)
- `mode` (`tickets-grouped-direct`)

Notes:
- `YYYY-MM-DD` is computed at pack generation time for default `out_dir`.
- If oracle flag support is uncertain, omit unsupported flags; never invent flags.

## Workflow (deterministic)

1) Read:
- `references/ticket-grouping.md`
- `references/attachment-minimization.md`
- `references/tickets-pack-template.md`

2) Ask user if custom args are needed (numbered picker):

```
1) Use defaults (no args)
2) Provide custom args
```

If `2`, ask for KEY=value args and run with those; otherwise run with defaults.

3) Generate packs (deterministic grouping + per-group pack files):

```bash
python3 /home/user/.codex/skills/oraclepack-tickets-pack-grouped/scripts/generate_grouped_packs.py \
  codebase_name=oraclepack \
  out_dir=docs/oracle-questions-2026-01-08
```

Outputs:
- `{{out_dir}}/packs/*.md` (one pack per group/part)
- `{{out_dir}}/_groups.json` (group -> ticket list)

4) Size control (mandatory; fail fast):
- Run `oracle --dry-run summary --files-report ...` for the **largest** group pack (or each pack if unsure).
- Enforce caps:
  - browser: ≤ 60,000 tokens total input per step
  - api: ≤ 180,000 tokens total input per step
- If exceeded, reduce via `group_max_files` or use explicit `ticket_paths`.

5) Validate every pack (mandatory):

```bash
python3 /home/user/.codex/skills/oraclepack-tickets-pack-grouped/scripts/validate_pack.py <pack.md>
python3 /home/user/.codex/skills/oraclepack-tickets-pack-grouped/scripts/lint_attachments.py <pack.md>
```

## Sharded packs workflow (topic/domain mini-packs)

Use this when you want a manifest-driven, sharded pack per topic/domain with bundle attachments:

First ask the user which args mode to use:

```
1) Use defaults (no args)
2) Provide custom args
```

If `2`, collect args and use them in the commands below.

```bash
python3 /home/user/.codex/skills/oraclepack-tickets-pack-grouped/scripts/shard_tickets.py \\
  --ticket-root .tickets \\
  --out-dir docs/oracle-questions-sharded

python3 /home/user/.codex/skills/oraclepack-tickets-pack-grouped/scripts/render_group_packs.py \\
  --manifest docs/oracle-questions-sharded/manifest.json \\
  --out-dir docs/oracle-questions-sharded

python3 /home/user/.codex/skills/oraclepack-tickets-pack-grouped/scripts/validate_shards.py \\
  --manifest docs/oracle-questions-sharded/manifest.json
```

## Failure behavior

- If no tickets resolve, packs still generate with empty attachments.
- Step 01 prompt must request exact missing ticket file/path pattern(s).

## Output contract

Each pack MUST:
- Have exactly one `bash` fence
- Have exactly 20 steps (01..20)
- Include ROI header tokens
- Include `--write-output` with a group-specific `out_dir`
- Attach tickets directly via `${ticket_args[@]}`
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
[TRUNCATED]
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

skills/oraclepack-tickets-pack-common/scripts/validate_pack.py
```
import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path, PurePosixPath
from typing import Dict, List, Tuple

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

REQUIRED_HEADER_KEYS = [
    "ROI",
    "impact",
    "confidence",
    "effort",
    "horizon",
    "category",
    "reference",
]


@dataclass(frozen=True)
class Step:
    n: str
    header_line: str
    block_lines: List[str]


def _fail(errors: List[str]) -> None:
    for e in errors:
        print(f"[ERROR] {e}", file=sys.stderr)
    sys.exit(1)


def _read_text(path: Path) -> str:
    return path.read_text(encoding="utf-8", errors="replace")


def _extract_single_bash_fence(lines: List[str]) -> Tuple[int, int, List[str], List[str]]:
    start = None
    for i, ln in enumerate(lines):
        if ln.strip().lower() == "```bash":
            if start is not None:
                raise ValueError("Multiple ```bash fences found; expected exactly one.")
            start = i

    if start is None:
        raise ValueError("No ```bash fence found; expected exactly one.")

    end = None
    for i in range(start + 1, len(lines)):
        if lines[i].strip() == "```":
            end = i
            break

    if end is None:
        raise ValueError("No closing ``` found for the ```bash fence.")

    fence_lines = [ln.rstrip("\n") for ln in lines[start + 1 : end]]
    outside_lines = [ln.rstrip("\n") for i, ln in enumerate(lines) if i < start or i > end]
    return start, end, fence_lines, outside_lines


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
        steps.append(Step(n=n, header_line=block[0], block_lines=block))
    return steps


def _parse_header_kv(header_line: str) -> Dict[str, str]:
    out: Dict[str, str] = {}
    tokens = header_line.strip().split()
    i = 0
    while i < len(tokens):
        tok = tokens[i]
        if "=" not in tok:
            i += 1
            continue
        key, val = tok.split("=", 1)
        if key in REQUIRED_HEADER_KEYS:
            if key == "category" and i + 1 < len(tokens):
                # Allow two-word categories like "background jobs" in headers.
                nxt = tokens[i + 1]
                if (val, nxt) in {
                    ("background", "jobs"),
                    ("UX", "flows"),
                    ("failure", "modes"),
                    ("feature", "flags"),
                }:
                    val = f"{val} {nxt}"
            out[key] = val
        i += 1
    return out


def _validate_header(step: Step, errors: List[str]) -> None:
    m = re.match(r"^#\s*(\d{2})\)\s+", step.header_line)
    if not m:
        errors.append(f"Step {step.n}: invalid header format (expected '# NN) ...').")
        return
    if m.group(1) != step.n:
        errors.append(f"Step {step.n}: header number mismatch (found {m.group(1)}).")

    kv = _parse_header_kv(step.header_line)
    for req in REQUIRED_HEADER_KEYS:
        if req not in kv:
            errors.append(f"Step {step.n}: header missing required token {req}=...")

    cat = kv.get("category")
    if cat is not None and cat not in ALLOWED_CATEGORIES:
        errors.append(
            f"Step {step.n}: category must be one of {ALLOWED_CATEGORIES}; got '{cat}'."
        )


def _validate_write_output(step: Step, errors: List[str]) -> None:
    joined = "\n".join(step.block_lines)

    m = re.search(r'(?<!\S)--write-output(?!\S)\s+"([^"]+)"', joined)
    if not m:
        errors.append(f"Step {step.n}: missing --write-output \"...\".")
        return

    out_path = m.group(1)
    if out_path.startswith("/") or out_path.startswith("~"):
        errors.append(f"Step {step.n}: --write-output must not be absolute: {out_path}")
        return

    p = PurePosixPath(out_path)
    if any(part == ".." for part in p.parts):
        errors.append(f"Step {step.n}: --write-output must not contain '..': {out_path}")

    if re.search(rf"(^|/){re.escape(step.n)}-", out_path) is None:
        errors.append(f"Step {step.n}: --write-output must include '{step.n}-' in filename: {out_path}")

    if not out_path.endswith(".md"):
        errors.append(f"Step {step.n}: --write-output must end with .md: {out_path}")


def _validate_ticket_bundle_reference(step: Step, errors: List[str]) -> None:
    joined = "\n".join(step.block_lines)

    if "_tickets_bundle" not in joined:
        errors.append(
            f"Step {step.n}: must reference the ticket bundle (expected '_tickets_bundle' in step block)."
        )

    if re.search(r'(?<!\S)(-f|--file)(?!\S)\s+"[^"\n]*_tickets_bundle[^"\n]*"', joined) is None:
        errors.append(
            f"Step {step.n}: must attach the ticket bundle via -f/--file \"..._tickets_bundle...\"."
        )


def _validate_answer_format(step: Step, errors: List[str]) -> None:
    hay = "\n".join(step.block_lines).lower()
    required = [
        "answer format:",
        "direct answer",
        "risks/unknowns",
        "next smallest concrete experiment",
        "if evidence is insufficient",
        "missing file/path pattern",
    ]
    missing = [s for s in required if s not in hay]
    if missing:
        errors.append(f"Step {step.n}: prompt missing required Answer format components: {missing}")


def _validate_category_counts(steps: List[Step], errors: List[str]) -> None:
    counts: Dict[str, List[str]] = {c: [] for c in ALLOWED_CATEGORIES}
    for st in steps:
        kv = _parse_header_kv(st.header_line)
        cat = kv.get("category")
        if cat in counts:
            counts[cat].append(st.n)

    bad = []
    for cat, ids in counts.items():
        if len(ids) != 2:
            bad.append(f"{cat}={len(ids)} (steps={ids})")
    if bad:
        errors.append(
            "Category distribution must be exactly 2 steps per category (20 total). Problems: "
            + ", ".join(bad)
        )


def _validate_step_numbers(steps: List[Step], errors: List[str]) -> None:
    nums = [st.n for st in steps]
    if nums != [f"{i:02d}" for i in range(1, 21)]:
        errors.append(f"Step numbering must be 01..20 in order; got {nums}.")


def _validate_coverage_check(outside_lines: List[str], errors: List[str]) -> None:
    text = "\n".join(outside_lines)
    m = re.search(r"^##\s+Coverage check\s*$", text, flags=re.IGNORECASE | re.MULTILINE)
    if m is None:
        errors.append('Missing "## Coverage check" section (must be outside the bash fence).')
        return

    after = text[m.end() :]
    for cat in ALLOWED_CATEGORIES:
        pat = rf"^\s*[-*]\s+{re.escape(cat)}\s*:\s*(OK|Missing\([^)]*\))\s*$"
        if re.search(pat, after, flags=re.MULTILINE) is None:
            errors.append(f'Coverage check missing/invalid line for category: "{cat}"')


def _validate_bash_hazards(step: Step, errors: List[str]) -> None:
    lines = step.block_lines[1:]
    for i, ln in enumerate(lines):
        s = ln.strip()

        if s == "\\":
            errors.append(f"Step {step.n}: contains a bare '\\\\' line (orphan backslash).")

        if s.startswith("#"):
            j = i - 1
            while j >= 0 and not lines[j].strip():
                j -= 1
            if j >= 0 and lines[j].rstrip().endswith("\\"):
                errors.append(
                    f"Step {step.n}: comment line appears immediately after a '\\'-continued line (comment-in-continuation hazard)."
                )

        if s.startswith("-p "):
            j = i - 1
            while j >= 0 and not lines[j].strip():
                j -= 1
            if j < 0 or not lines[j].rstrip().endswith("\\"):
                errors.append(
                    f"Step {step.n}: '-p ...' line is not attached to a continued command (detached -p hazard)."
                )


def validate_pack(path: Path, require_bundle: bool) -> None:
    raw = _read_text(path)
    lines = raw.splitlines(True)

    errors: List[str] = []

    try:
        _start, _end, fence, outside = _extract_single_bash_fence(lines)
    except Exception as e:
        _fail([str(e)])

    steps = _parse_steps(fence)

    if len(steps) != 20:
        errors.append(f"Expected exactly 20 steps; found {len(steps)}.")
    else:
        _validate_step_numbers(steps, errors)

    for step in steps:
        _validate_header(step, errors)
        _validate_write_output(step, errors)
        _validate_answer_format(step, errors)
        _validate_bash_hazards(step, errors)
        if require_bundle:
            _validate_ticket_bundle_reference(step, errors)

    _validate_category_counts(steps, errors)
    _validate_coverage_check(outside, errors)

    if errors:
        _fail(errors)

    print("[OK] Pack validates against tickets Stage-1 contract.")


def main() -> None:
    p = argparse.ArgumentParser(description="Validate oraclepack Stage-1 ticket packs.")
    p.add_argument("pack_path", help="Path to the Markdown pack file")
    p.add_argument(
        "--mode",
        choices=["bundle", "direct"],
        default="direct",
        help="Validation mode: bundle requires _tickets_bundle attachments",
    )
    args = p.parse_args()

    path = Path(args.pack_path)
    if not path.exists():
        _fail([f"File not found: {path}"])

    validate_pack(path, require_bundle=args.mode == "bundle")


if __name__ == "__main__":
    main()
```

skills/oraclepack-tickets-pack-grouped/references/attachment-minimization.md
```
# Attachment minimization rules (Grouped Tickets Stage 1 — Direct Attach)

Objective: keep each group pack focused and portable.

## Ticket attachments

- Ticket files are attached directly in each step via `${ticket_args[@]}`.
- Use `group_max_files` (default 25) to bound per-pack ticket count.
- If a group is larger than the cap, split into multiple packs (part 1..N).

## Non-ticket attachments (repo evidence)

- Keep explicit non-ticket attachments to **0–1 per step**.
- Prefer a single high-signal file that clarifies contracts or a key code path.

## extra_files (literal append)

- If `extra_files` is provided, append it literally to every oracle command.
- It may include additional `-f/--file` flags.
- Place `extra_files` on its own line with a comment:
  - `# extra_files appended literally`

```

skills/oraclepack-tickets-pack-grouped/references/ticket-grouping.md
```
# Ticket grouping (deterministic, inferred)

Objective: split tickets into focused topic/domain groups and generate one pack per group.

## Inputs

- `ticket_root` (default `.tickets`)
- `ticket_glob` (default `**/*.md`, relative to `ticket_root`)
- `ticket_paths` (optional; comma-separated explicit files; if present, ignore `ticket_glob`)
- `group_mode` (default `subdir+infer`)
- `group_min_score` (default `0.08`)
- `group_max_files` (default `25`; max tickets per pack; >0)
- `group_max_chars` (default `200000`; max total chars per pack; >0)
- `dedupe_mode` (default `report`; one of `off`, `report`, `prune`, `merge`)
- `dedupe_jaccard` (default `0.55`)
- `dedupe_overlap_hi` (default `0.80`)
- `dedupe_overlap_lo` (default `0.70`)
- `dedupe_delta_min` (default `0.15`)
- `dedupe_body_chars` (default `2000`)

## Deterministic grouping rules

1) Collect tickets:
- If `ticket_paths` is non-empty: split on commas, trim whitespace, use exactly that list.
- Else: glob `ticket_root/ticket_glob`.
- Always sort lexicographically by path string.

2) Detect possible duplicates (if `dedupe_mode != off`):
- Signature: filename stem + first heading + first `dedupe_body_chars` chars.
- Compute `jaccard` + `overlap` between tickets.
- Duplicate edge rule:
  - `overlap >= dedupe_overlap_hi` OR (`jaccard >= dedupe_jaccard` AND `overlap >= dedupe_overlap_lo`)
- Connected components become duplicate clusters.
- Canonical: largest content length; tie-break lexicographic.
- Delta vs redundant:
  - delta if unique token ratio >= `dedupe_delta_min` OR heading differs materially.
  - redundant otherwise.

3) Seed groups by subdir:
- For any path under `ticket_root/<group>/...`, assign to group `<group>`.
- Tickets directly under `ticket_root/` are "loose".

4) Infer loose tickets into groups (if any groups exist):
- Build a token set for each group from:
  - group name tokens
  - ticket filenames (stem tokens)
  - first Markdown heading line (if present)
- For each loose ticket, compute Jaccard overlap score with each group token set.
- If `max_score >= group_min_score`, assign to the best group (stable tie-break by group name).
- Otherwise, assign to `misc`.

5) If no groups exist:
- Put all tickets into a single group named `root`.

6) Merge duplicates into primary group:
- `report`: attach all tickets in the cluster to the canonical’s group.
- `prune`: attach canonical + delta only; drop redundant from attachments.
- `merge`: create `out_dir/_ticket_merges/cluster-XXXX.md` and attach only the merged file.
- Emit `_dupes_possible.json`, `_duplicates.json`, and `_dedupe_plan.json`.

7) Split oversized groups:
- If a group exceeds `group_max_files` or `group_max_chars`, split into parts (1..N)
  in sorted order, chunked deterministically.

Hard rule: do not use mtimes, file sizes, or external ML services.

## Required outputs

- `_groups.json`: mapping of group -> list of ticket paths (lexicographic order)
- Pack file per group (and part), each self-contained and direct-attach
- `manifest.json`: groups with pack path + attached vs original ticket lists
```

skills/oraclepack-tickets-pack-grouped/references/tickets-pack-template-bundle.md
```
# Oracle Pack — {{codebase_name}} (Tickets Stage 1)

## Parsed args
- codebase_name: {{codebase_name}}
- out_dir: {{out_dir}}
- oracle_cmd: {{oracle_cmd}}
- oracle_flags: {{oracle_flags}}
- extra_files: {{extra_files}}
- ticket_root: {{ticket_root}}
- ticket_glob: {{ticket_glob}}
- ticket_paths: {{ticket_paths}}
- ticket_bundle_path: {{ticket_bundle_path}}
- mode: {{mode}}

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "{{out_dir}}/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- `## Coverage check` MUST be outside the bash fence (after the closing ```).

```bash
# Prelude (allowed inside the single bash fence)
# - Creates out_dir deterministically
# - Builds ticket_bundle_path deterministically from ticket_root/ticket_glob OR ticket_paths
# - Uses lexicographic ordering only (no mtime/timestamps)

set -euo pipefail

mkdir -p "{{out_dir}}"

python3 - <<'PY'
from __future__ import annotations

import sys
from pathlib import Path

CODEBASE_NAME = "{{codebase_name}}"
OUT_DIR = Path("{{out_dir}}")
TICKET_ROOT = Path("{{ticket_root}}")
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS_RAW = "{{ticket_paths}}".strip()
BUNDLE_PATH = Path("{{ticket_bundle_path}}")

def _read_text(p: Path) -> str:
    return p.read_text(encoding="utf-8", errors="replace")

def _title_from_md(text: str) -> str:
    for ln in text.splitlines():
        s = ln.strip()
        if s.startswith("# "):
            return s[2:].strip() or "Untitled"
    for ln in text.splitlines():
        s = ln.strip()
        if s:
            return s[:80]
    return "Untitled"

def _select_paths() -> list[Path]:
    if TICKET_PATHS_RAW:
        items = [Path(x.strip()) for x in TICKET_PATHS_RAW.split(",") if x.strip()]
        items = sorted(items, key=lambda p: str(p))
        return items

    if not TICKET_ROOT.exists():
        return []

    items = sorted(TICKET_ROOT.glob(TICKET_GLOB), key=lambda p: str(p))
    return items

paths = _select_paths()

BUNDLE_PATH.parent.mkdir(parents=True, exist_ok=True)

lines: list[str] = []
lines.append(f"# Tickets Bundle — {CODEBASE_NAME if CODEBASE_NAME else 'Unknown'}")
lines.append("")
lines.append("## Selection")
lines.append(f"- ticket_root: {TICKET_ROOT}")
lines.append(f"- ticket_glob: {TICKET_GLOB}")
lines.append(f"- ticket_paths: {TICKET_PATHS_RAW if TICKET_PATHS_RAW else '(none)'}")
lines.append("- ordering: lexicographic by path")
lines.append("")

if not paths:
    warn = (
        "## WARNING: No tickets found\n\n"
        "No ticket files were selected.\n\n"
        "What was attempted:\n"
        f"- ticket_root: {TICKET_ROOT}\n"
        f"- ticket_glob: {TICKET_GLOB}\n"
        f"- ticket_paths: {TICKET_PATHS_RAW if TICKET_PATHS_RAW else '(none)'}\n\n"
        "Next: provide explicit ticket_paths or create tickets under ticket_root.\n"
    )
    lines.append(warn)
    print(f"[WARN] No tickets selected; bundle will contain only WARNING.", file=sys.stderr)
else:
    lines.append("## Tickets")
    lines.append("")
    for p in paths:
        lines.append("---")
        lines.append(f"### {_title_from_md(_read_text(p))}")
        lines.append(f"- path: {p}")
        lines.append("")
        try:
            txt = _read_text(p)
        except Exception as e:
            lines.append(f"[ERROR reading file: {e}]")
            lines.append("")
            continue

        # Simple truncation policy: keep first 4000 chars if large.
        if len(txt) > 4000:
            lines.append(txt[:4000])
            lines.append("\n[... truncated ...]\n")
        else:
            lines.append(txt)

        lines.append("")

BUNDLE_PATH.write_text("\n".join(lines).rstrip() + "\n", encoding="utf-8")
print(f"[OK] Wrote ticket bundle: {BUNDLE_PATH}")
PY

# 01) ROI=8.0 impact=9 confidence=0.9 effort=1 horizon=Immediate category=contracts/interfaces reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/01-contracts-interfaces-surface.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #01
Category: contracts/interfaces

Using the attached tickets bundle as the primary evidence, identify the primary public interface(s) implied by the tickets (CLI commands, APIs, file contracts, or user workflows).
For each interface:
- list key inputs/outputs
- list the exact files/modules likely defining it (if unknown, say Unknown)

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=7.8 impact=8 confidence=0.9 effort=1 horizon=Immediate category=contracts/interfaces reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/02-contracts-interfaces-dependencies.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #02
Category: contracts/interfaces

From the attached tickets bundle, infer which external dependencies/services the system must integrate with (CLIs, APIs, SaaS, databases).
For each dependency:
- what contract is required (auth, endpoints, file formats)
- what configuration surface is implied

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=7.6 impact=8 confidence=0.85 effort=2 horizon=Immediate category=invariants reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/03-invariants-must-always-hold.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #03
Category: invariants

Based on the attached tickets bundle, list the invariants that must always hold (data constraints, ordering constraints, security invariants, idempotency).
For each invariant:
- what breaks if violated
- where it should be enforced (layer/module; if unknown, Unknown)

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=7.2 impact=8 confidence=0.8 effort=2 horizon=Immediate category=invariants reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/04-invariants-input-validation.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #04
Category: invariants

Using the attached tickets bundle, identify what inputs must be validated (CLI args, config fields, payloads, file paths).
For each input:
- validation rules implied
- failure message/behavior implied

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=7.0 impact=7 confidence=0.85 effort=2 horizon=Near category=caching/state reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/05-caching-state-state-model.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #05
Category: caching/state

From the attached tickets bundle, infer what state must be persisted or cached (files, DB, in-memory, remote).
For each state item:
- read/write lifecycle
- consistency model implied
- failure recovery requirements

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=6.8 impact=7 confidence=0.8 effort=2 horizon=Near category=caching/state reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/06-caching-state-cache-invalidation.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #06
Category: caching/state

Using the attached tickets bundle, identify caching risks: staleness, invalidation, keying, or race conditions implied by the tickets.
Propose a minimal caching strategy consistent with the tickets.

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 07) ROI=6.9 impact=8 confidence=0.75 effort=3 horizon=Near category=background jobs reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/07-background-jobs-what-runs-async.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #07
Category: background jobs

From the attached tickets bundle, determine what work should run asynchronously/background (schedulers, queues, cron, long-running tasks).
For each job:
- trigger
- inputs/outputs
- retry/backoff requirements

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 08) ROI=6.6 impact=7 confidence=0.75 effort=3 horizon=Near category=background jobs reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/08-background-jobs-idempotency.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #08
Category: background jobs

Using the attached tickets bundle, list the idempotency and concurrency constraints implied for background jobs.
Recommend minimal safeguards (dedupe keys, locks, at-least-once handling) aligned with tickets.

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 09) ROI=7.4 impact=8 confidence=0.8 effort=2 horizon=Immediate category=observability reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/09-observability-logs-metrics-traces.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #09
Category: observability

From the attached tickets bundle, infer required observability: logs, metrics, traces, and user-visible diagnostics.
List:
- what to log/measure
- cardinality risks
- minimal dashboards/alerts implied

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 10) ROI=7.0 impact=7 confidence=0.8 effort=2 horizon=Near category=observability reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/10-observability-error-taxonomy.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #10
Category: observability

Using the attached tickets bundle, define an error taxonomy consistent with ticket failure modes:
- user errors vs system errors
- retryable vs non-retryable
- how errors should surface (CLI exit codes, UI states, logs)

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 11) ROI=7.6 impact=9 confidence=0.75 effort=3 horizon=Immediate category=permissions reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/11-permissions-authz-model.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #11
Category: permissions

From the attached tickets bundle, infer the permissions model (roles, capabilities, scopes).
List:
- what operations require permissions
- how permissions are granted/revoked
- audit requirements implied

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 12) ROI=7.0 impact=8 confidence=0.75 effort=3 horizon=Near category=permissions reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/12-permissions-secret-handling.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #12
Category: permissions

Using the attached tickets bundle, identify sensitive data/secret handling needs.
Recommend:
- where secrets come from (env, files, vault)
- redaction rules
- least-privilege defaults

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 13) ROI=7.2 impact=8 confidence=0.8 effort=2 horizon=Near category=migrations reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/13-migrations-data-changes.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #13
Category: migrations

From the attached tickets bundle, infer any data/schema/config migrations needed.
For each migration:
- trigger/versioning
- rollout plan
- rollback strategy

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 14) ROI=6.8 impact=7 confidence=0.8 effort=2 horizon=Near category=migrations reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/14-migrations-compatibility.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #14
Category: migrations

Using the attached tickets bundle, identify backwards/forwards compatibility requirements during migration windows.
Recommend minimal compatibility shims or staged rollout steps.

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 15) ROI=7.4 impact=8 confidence=0.8 effort=2 horizon=Immediate category=UX flows reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/15-ux-flows-primary-journeys.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #15
Category: UX flows

From the attached tickets bundle, map the primary user journeys implied by tickets.
For each journey:
- entry points
- steps/screens/commands
- success criteria and user feedback

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 16) ROI=6.9 impact=7 confidence=0.8 effort=2 horizon=Near category=UX flows reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/16-ux-flows-edge-cases.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #16
Category: UX flows

Using the attached tickets bundle, list UX edge cases and failure UX:
- partial completion
- retries
- cancellation
- timeouts
- conflict resolution

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 17) ROI=7.8 impact=9 confidence=0.8 effort=2 horizon=Immediate category=failure modes reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/17-failure-modes-top-risks.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #17
Category: failure modes

From the attached tickets bundle, enumerate the most likely failure modes.
For each failure mode:
- detection signal
- mitigation
- user-visible behavior

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 18) ROI=7.0 impact=8 confidence=0.75 effort=3 horizon=Near category=failure modes reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/18-failure-modes-test-plan.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #18
Category: failure modes

Using the attached tickets bundle, propose a minimal test plan that covers the highest-risk failure modes.
Include:
- unit vs integration coverage split
- fixtures/mocks needed
- one smallest test to write first

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 19) ROI=7.3 impact=8 confidence=0.8 effort=2 horizon=Near category=feature flags reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/19-feature-flags-needed.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #19
Category: feature flags

From the attached tickets bundle, infer where feature flags or staged rollouts are needed.
For each flag:
- what it gates
- default value
- sunset plan

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 20) ROI=6.8 impact=7 confidence=0.8 effort=2 horizon=Near category=feature flags reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/20-feature-flags-observability.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #20
Category: feature flags

Using the attached tickets bundle, propose how to observe/validate a flagged rollout:
- success metrics
- rollback triggers
- logging/alert changes while enabled

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"
```

## Coverage check

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

skills/oraclepack-tickets-pack-grouped/references/tickets-pack-template.md
```
# Oracle Pack — {{codebase_name}} (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: {{codebase_name}}
- out_dir: {{out_dir}}
- oracle_cmd: {{oracle_cmd}}
- oracle_flags: {{oracle_flags}}
- extra_files: {{extra_files}}
- ticket_root: {{ticket_root}}
- ticket_glob: {{ticket_glob}}
- ticket_paths: {{ticket_paths}}
- ticket_max_files: {{ticket_max_files}}
- group_name: {{group_name}}
- group_slug: {{group_slug}}
- mode: {{mode}}

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "{{out_dir}}/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "{{out_dir}}"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference={{group_slug}}

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = "{{ticket_root}}"
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS = "{{ticket_paths}}".strip()
MAX = int("{{ticket_max_files}}")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/01-contracts-interfaces-ticket-surface.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

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

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = "{{ticket_root}}"
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS = "{{ticket_paths}}".strip()
MAX = int("{{ticket_max_files}}")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/02-contracts-interfaces-integration-points.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

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

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = "{{ticket_root}}"
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS = "{{ticket_paths}}".strip()
MAX = int("{{ticket_max_files}}")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/03-invariants-invariant-map.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference={{group_slug}}

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = "{{ticket_root}}"
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS = "{{ticket_paths}}".strip()
MAX = int("{{ticket_max_files}}")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/04-invariants-validation-boundaries.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation); propose minimal validation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=4.4 impact=6 confidence=0.78 effort=2 horizon=NearTerm category=caching/state reference={{group_slug}}

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = "{{ticket_root}}"
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS = "{{ticket_paths}}".strip()
MAX = int("{{ticket_max_files}}")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/05-caching-state-state-artifacts.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #05  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.78, effort=2)

Question:
Using the attached tickets as the primary context, identify state/artifacts that must be produced and preserved; schema/format expectations; stability/back-compat requirements.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=4.2 impact=6 confidence=0.75 effort=2 horizon=NearTerm category=caching/state reference={{group_slug}}

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = "{{ticket_root}}"
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS = "{{ticket_paths}}".strip()
MAX = int("{{ticket_max_files}}")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/06-caching-state-cache-keys.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #06  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: caching/state
Horizon: NearTerm
ROI: 4.2 (impact=6, confidence=0.75, effort=2)

Question:
Using the attached tickets as the primary context, identify any caching opportunities/risks (discovery caches, pack outputs, oracle outputs); define cache keys, invalidation, and correctness risks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 07) ROI=4.3 impact=6 confidence=0.70 effort=2 horizon=MidTerm category=background jobs reference={{group_slug}}

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = "{{ticket_root}}"
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS = "{{ticket_paths}}".strip()
MAX = int("{{ticket_max_files}}")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/07-background-jobs-job-model.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #07  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: background jobs
Horizon: MidTerm
ROI: 4.3 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, identify any background/async work implied (jobs, queues, long-running operations); define responsibilities and interfaces.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 08) ROI=4.0 impact=6 confidence=0.68 effort=3 horizon=MidTerm category=background jobs reference={{group_slug}}

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = "{{ticket_root}}"
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS = "{{ticket_paths}}".strip()
MAX = int("{{ticket_max_files}}")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/08-background-jobs-queue-failure.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #08  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: background jobs
Horizon: MidTerm
ROI: 4.0 (impact=6, confidence=0.68, effort=3)

Question:
Using the attached tickets as the primary context, define how background failures are handled (retries, idempotency, poison messages); define observability hooks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 09) ROI=4.7 impact=7 confidence=0.76 effort=1 horizon=Immediate category=observability reference={{group_slug}}

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = "{{ticket_root}}"
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS = "{{ticket_paths}}".strip()
MAX = int("{{ticket_max_files}}")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/09-observability-logging-metrics.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #09  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: observability
Horizon: Immediate
ROI: 4.7 (impact=7, confidence=0.76, effort=1)

Question:
Using the attached tickets as the primary context, define what logging/metrics must exist to debug pack generation + step execution; propose minimal instrumentation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 10) ROI=4.5 impact=7 confidence=0.74 effort=2 horizon=Immediate category=observability reference={{group_slug}}

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = "{{ticket_root}}"
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS = "{{ticket_paths}}".strip()
MAX = int("{{ticket_max_files}}")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/10-observability-tracing.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #10  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: observability
Horizon: Immediate
ROI: 4.5 (impact=7, confidence=0.74, effort=2)

Question:
Using the attached tickets as the primary context, define tracing/correlation strategy across pack steps and downstream tools; identify required IDs and propagation.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 11) ROI=4.1 impact=6 confidence=0.70 effort=2 horizon=NearTerm category=permissions reference={{group_slug}}

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = "{{ticket_root}}"
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS = "{{ticket_paths}}".strip()
MAX = int("{{ticket_max_files}}")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
[TRUNCATED]
```

skills/oraclepack-tickets-pack-grouped/scripts/generate_grouped_packs.py
```
#!/usr/bin/env python3
from __future__ import annotations

import datetime as _dt
import math
import json
import re
import sys
from pathlib import Path
from typing import Dict, Iterable, List, Tuple

STOPWORDS = {
    "the", "and", "for", "with", "from", "this", "that", "into", "over", "under", "when",
    "then", "than", "else", "only", "must", "should", "could", "would", "will", "shall",
    "ticket", "tickets", "oraclepack", "oracle", "pack", "packs",
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


def _normalize_title(text: str) -> str:
    text = text.strip().lower()
    text = re.sub(r"[^a-z0-9]+", " ", text)
    text = re.sub(r"\s+", " ", text).strip()
    return text


def _read_heading(path: Path) -> str:
    try:
        for line in path.read_text(encoding="utf-8", errors="replace").splitlines():
            if line.startswith("#"):
                return line.lstrip("#").strip()
    except FileNotFoundError:
        return ""
    return ""


def _collect_ticket_paths(ticket_root: str, ticket_glob: str, ticket_paths: str) -> List[Path]:
    if ticket_paths:
        parts = [p.strip() for p in ticket_paths.split(",") if p.strip()]
        return [Path(p) for p in parts]
    root = Path(ticket_root)
    if not root.exists():
        return []
    return [Path(p) for p in root.glob(ticket_glob)]


def _read_signature(path: Path, max_lines: int = 40) -> Tuple[str, str]:
    heading = ""
    lines: List[str] = []
    try:
        for line in path.read_text(encoding="utf-8", errors="replace").splitlines():
            if not heading and line.startswith("#"):
                heading = line.lstrip("#").strip()
            if line.strip():
                lines.append(line.strip())
            if len(lines) >= max_lines:
                break
    except FileNotFoundError:
        pass
    return heading, " ".join(lines)


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8", errors="replace")
    except FileNotFoundError:
        return ""


def _group_by_subdir(paths: Iterable[Path], ticket_root: str) -> Tuple[Dict[str, List[Path]], List[Path]]:
    root = Path(ticket_root)
    groups: Dict[str, List[Path]] = {}
    loose: List[Path] = []
    for p in paths:
        try:
            rel = p.relative_to(root)
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
        tokens.update(_tokenize(_read_heading(p)))
    return tokens


def _ticket_tokens(p: Path) -> set:
    toks = set(_tokenize(p.stem))
    heading, snippet = _read_signature(p)
    toks.update(_tokenize(heading))
    toks.update(_tokenize(snippet))
    return toks


def _signature_tokens(p: Path, body_chars: int) -> set:
    heading = _read_heading(p)
    body = _read_text(p)
    body = body[:body_chars]
    toks = set(_tokenize(p.stem))
    toks.update(_tokenize(heading))
    toks.update(_tokenize(body))
    return toks


def _jaccard(a: set, b: set) -> float:
    if not a or not b:
        return 0.0
    inter = a.intersection(b)
    union = a.union(b)
    return float(len(inter)) / float(len(union))


def _overlap(a: set, b: set) -> float:
    if not a or not b:
        return 0.0
    inter = a.intersection(b)
    denom = min(len(a), len(b))
    if denom == 0:
        return 0.0
    return float(len(inter)) / float(denom)


def _clusters_from_edges(nodes: List[str], edges: Dict[str, List[str]]) -> List[List[str]]:
    seen = set()
    clusters: List[List[str]] = []
    for n in nodes:
        if n in seen:
            continue
        stack = [n]
        comp = []
        seen.add(n)
        while stack:
            cur = stack.pop()
            comp.append(cur)
            for nxt in edges.get(cur, []):
                if nxt not in seen:
                    seen.add(nxt)
                    stack.append(nxt)
        clusters.append(sorted(comp))
    return clusters


def _dedupe_clusters(
    paths: List[Path],
    body_chars: int,
    jaccard_hi: float,
    overlap_hi: float,
    overlap_lo: float,
    delta_min: float,
) -> Tuple[List[List[str]], Dict[str, str], Dict[str, Dict[str, object]], Dict[Tuple[str, str], Dict[str, float]]]:
    tokens: Dict[str, set] = {}
    sizes: Dict[str, int] = {}
    titles: Dict[str, str] = {}
    for p in paths:
        key = str(p)
        tokens[key] = _signature_tokens(p, body_chars)
        sizes[key] = len(_read_text(p))
        titles[key] = _normalize_title(_read_heading(p))

    nodes = sorted(tokens.keys())
    edges: Dict[str, List[str]] = {n: [] for n in nodes}
    pair_scores: Dict[Tuple[str, str], Dict[str, float]] = {}

    for i, a in enumerate(nodes):
        for b in nodes[i + 1 :]:
            jac = _jaccard(tokens[a], tokens[b])
            ov = _overlap(tokens[a], tokens[b])
            pair_scores[(a, b)] = {"jaccard": jac, "overlap": ov}
            if ov >= overlap_hi or (jac >= jaccard_hi and ov >= overlap_lo):
                edges[a].append(b)
                edges[b].append(a)

    clusters = _clusters_from_edges(nodes, edges)
    cluster_meta: Dict[str, Dict[str, object]] = {}
    dup_map: Dict[str, str] = {}

    for idx, members in enumerate(clusters, start=1):
        if len(members) == 1:
            continue
        # canonical: largest content length, then lexicographic
        canon = sorted(
            members,
            key=lambda m: (-sizes.get(m, 0), m),
        )[0]
        deltas: List[str] = []
        redundant: List[str] = []
        for m in members:
            if m == canon:
                continue
            unique = tokens[m] - tokens[canon]
            unique_ratio = float(len(unique)) / float(max(1, len(tokens[m])))
            heading_diff = titles.get(m, "") != titles.get(canon, "")
            if unique_ratio >= delta_min or heading_diff:
                deltas.append(m)
            else:
                redundant.append(m)
            dup_map[m] = canon

        cluster_meta[str(idx)] = {
            "canonical": canon,
            "members": members,
            "deltas": sorted(deltas),
            "redundant": sorted(redundant),
        }

    return clusters, dup_map, cluster_meta, pair_scores


def _infer_groups(
    groups: Dict[str, List[Path]],
    loose: List[Path],
    min_score: float,
) -> Dict[str, List[Path]]:
    if not groups:
        return {"root": list(loose)}

    group_tokens = {k: _group_tokens(k, v) for k, v in groups.items()}
    for p in loose:
        tokens = _ticket_tokens(p)
        best = None
        best_score = -1.0
        for name in sorted(group_tokens.keys()):
            score = _jaccard(tokens, group_tokens[name])
            if score > best_score:
                best_score = score
                best = name
        if best is not None and best_score >= min_score:
            groups.setdefault(best, []).append(p)
        else:
            groups.setdefault("misc", []).append(p)
    return groups


def _chunk(paths: List[Path], size: int) -> List[List[Path]]:
    if size <= 0:
        return [paths]
    return [paths[i : i + size] for i in range(0, len(paths), size)]


def _chunk_by_limits(
    paths: List[Path],
    max_files: int,
    max_chars: int,
) -> List[List[Path]]:
    if max_files <= 0 and max_chars <= 0:
        return [paths]
    chunks: List[List[Path]] = []
    cur: List[Path] = []
    cur_chars = 0
    for p in paths:
        size = len(_read_text(p))
        if cur:
            if (max_files > 0 and len(cur) >= max_files) or (
                max_chars > 0 and cur_chars + size > max_chars
            ):
                chunks.append(cur)
                cur = []
                cur_chars = 0
        cur.append(p)
        cur_chars += size
    if cur:
        chunks.append(cur)
    return chunks


def _render_template(template: str, mapping: Dict[str, str]) -> str:
    out = template
    for key, val in mapping.items():
        out = out.replace("{{" + key + "}}", val)
    unresolved = sorted(set(re.findall(r"\{\{([^}]+)\}\}", out)))
    if unresolved:
        raise ValueError(f"Unresolved template placeholders: {unresolved}")
    return out


def _write_merge_file(
    out_dir: Path,
    cluster_id: str,
    canonical: str,
    deltas: List[str],
    redundant: List[str],
    body_chars: int,
) -> Path:
    merge_dir = out_dir / "_ticket_merges"
    merge_dir.mkdir(parents=True, exist_ok=True)
    path = merge_dir / f"cluster-{int(cluster_id):04d}.md"

    def _cap(text: str) -> str:
        if len(text) <= body_chars:
            return text
        return text[:body_chars] + "\n[... truncated ...]\n"

    lines: List[str] = []
    lines.append(f"# Ticket Merge Cluster {cluster_id}")
    lines.append("")
    lines.append("## Canonical")
    lines.append(f"- path: {canonical}")
    lines.append("")
    lines.append(_cap(_read_text(Path(canonical))))
    lines.append("")

    members = deltas + redundant
    if members:
        lines.append("## Also reported in")
        for m in members:
            lines.append(f"- {m}")
        lines.append("")

    if deltas:
        lines.append("## Unique details from related tickets")
        for m in deltas:
            text = _read_text(Path(m))
            toks = _signature_tokens(Path(m), body_chars)
            canon_toks = _signature_tokens(Path(canonical), body_chars)
            unique = toks - canon_toks
            sel: List[str] = []
            for ln in text.splitlines():
                lnt = _tokenize(ln)
                if any(t in unique for t in lnt):
                    sel.append(ln)
                if len(sel) >= 60:
                    break
            lines.append(f"### {m}")
            if sel:
                lines.extend(sel)
            else:
                lines.append("(no unique lines detected within cap)")
            lines.append("")

    path.write_text("\n".join(lines), encoding="utf-8")
    return path


def main() -> int:
    if len(sys.argv) == 1:
        print("Select how to run:")
        print("1) Use defaults (no args)")
        print("2) Provide custom args (show usage)")
        choice = input("Enter choice [1-2]: ").strip() or "1"
        if choice == "2":
            print("Usage: generate_grouped_packs.py key=value [key=value ...]")
            return 0

    args = _parse_kv_args(sys.argv[1:])
    codebase_name = args.get("codebase_name", "Unknown")
    out_dir = args.get("out_dir", f"docs/oracle-questions-{_today()}")
    oracle_cmd = args.get("oracle_cmd", "oracle")
    oracle_flags = args.get("oracle_flags", "--files-report")
    extra_files = args.get("extra_files", "")
    ticket_root = args.get("ticket_root", ".tickets")
    ticket_glob = args.get("ticket_glob", "**/*.md")
    ticket_paths = args.get("ticket_paths", "")
    ticket_max_files = args.get("ticket_max_files", "25")
    group_mode = args.get("group_mode", "subdir+infer")
    group_min_score = float(args.get("group_min_score", "0.08"))
    group_max_files = int(args.get("group_max_files", "25"))
    group_max_chars = int(args.get("group_max_chars", "200000"))
    dedupe_mode = args.get("dedupe_mode", "report")
    dedupe_jaccard = float(args.get("dedupe_jaccard", "0.55"))
    dedupe_overlap_hi = float(args.get("dedupe_overlap_hi", "0.80"))
    dedupe_overlap_lo = float(args.get("dedupe_overlap_lo", "0.70"))
    dedupe_delta_min = float(args.get("dedupe_delta_min", "0.15"))
    dedupe_body_chars = int(args.get("dedupe_body_chars", "2000"))
    mode = args.get("mode", "tickets-grouped-direct")

    template_path = Path(__file__).resolve().parent.parent / "references" / "tickets-pack-template.md"
    template = template_path.read_text(encoding="utf-8")

    paths = _collect_ticket_paths(ticket_root, ticket_glob, ticket_paths)
    paths = sorted((str(p) for p in paths))
    paths = [Path(p) for p in paths]

    original_paths = list(paths)
    dup_map: Dict[str, str] = {}
    cluster_meta: Dict[str, Dict[str, object]] = {}
    dup_pairs: Dict[Tuple[str, str], Dict[str, float]] = {}
    if dedupe_mode != "off":
        _clusters, dup_map, cluster_meta, dup_pairs = _dedupe_clusters(
            paths,
            body_chars=dedupe_body_chars,
            jaccard_hi=dedupe_jaccard,
            overlap_hi=dedupe_overlap_hi,
            overlap_lo=dedupe_overlap_lo,
            delta_min=dedupe_delta_min,
        )

    # Build grouping base: canonical tickets + singletons
    canonical_set = {meta["canonical"] for meta in cluster_meta.values()}
    dup_set = set(dup_map.keys())
    base_paths: List[Path] = []
    for p in paths:
        sp = str(p)
        if sp in dup_set:
            continue
        base_paths.append(p)

    groups: Dict[str, List[Path]] = {}
    loose: List[Path] = []
    if "subdir" in group_mode:
        groups, loose = _group_by_subdir(base_paths, ticket_root)
    else:
        loose = list(base_paths)

    if "infer" in group_mode:
        groups = _infer_groups(groups, loose, group_min_score)
    else:
        groups.setdefault("misc", []).extend(loose)

    dedupe_plan: Dict[str, Dict[str, object]] = {}
    merge_files: Dict[str, str] = {}
    if cluster_meta:
        primary_to_group: Dict[str, str] = {}
        for gname in groups:
            for p in groups[gname]:
                primary_to_group[str(p)] = gname

        for cluster_id, meta in sorted(cluster_meta.items(), key=lambda x: int(x[0])):
            canonical = meta["canonical"]
            deltas = list(meta["deltas"])
            redundant = list(meta["redundant"])
            gname = primary_to_group.get(canonical, "misc")

            if dedupe_mode == "merge":
                merge_path = _write_merge_file(
                    Path(out_dir),
                    cluster_id=cluster_id,
                    canonical=canonical,
                    deltas=deltas,
                    redundant=redundant,
                    body_chars=dedupe_body_chars,
                )
                merge_files[canonical] = str(merge_path)
                # Replace canonical in group with merge file
                groups[gname] = [p for p in groups[gname] if str(p) != canonical]
                groups[gname].append(merge_path)
            else:
                # report/prune: append related tickets to canonical group
                keep = deltas if dedupe_mode == "prune" else deltas + redundant
                for p in keep:
                    groups.setdefault(gname, []).append(Path(p))

            dedupe_plan[cluster_id] = {
                "canonical": canonical,
                "group": gname,
                "deltas": sorted(deltas),
                "redundant": sorted(redundant),
                "mode": dedupe_mode,
            }

    # Ensure stable order
    for k in sorted(groups.keys()):
        groups[k] = sorted((str(p) for p in groups[k]))
        groups[k] = [Path(p) for p in groups[k]]

    original_set = {str(p) for p in original_paths}
    assignment: Dict[str, str] = {}
    for gname, gpaths in groups.items():
        for p in gpaths:
            sp = str(p)
            if sp in original_set:
                if sp in assignment:
                    raise SystemExit(f"[ERROR] Ticket assigned to multiple groups: {sp}")
                assignment[sp] = gname

    for meta in dedupe_plan.values():
        gname = meta["group"]
        for sp in [meta["canonical"]] + meta["deltas"] + meta["redundant"]:
            if sp not in assignment:
                assignment[sp] = gname

    missing = sorted(original_set - set(assignment.keys()))
    if missing:
        raise SystemExit(f"[ERROR] Tickets missing group assignment: {missing}")

    base_out = Path(out_dir)
    packs_dir = base_out / "packs"
    packs_dir.mkdir(parents=True, exist_ok=True)

    grouping_report: Dict[str, List[str]] = {}
    manifest_groups: List[Dict[str, object]] = []
    group_originals: Dict[str, List[str]] = {g: [] for g in groups.keys()}
    for ticket, gname in assignment.items():
        group_originals.setdefault(gname, []).append(ticket)
    for group_name in sorted(groups.keys()):
        group_paths = groups[group_name]
        grouping_report[group_name] = [str(p) for p in group_paths]

        parts = _chunk_by_limits(group_paths, group_max_files, group_max_chars)
        for idx, part in enumerate(parts, start=1):
            part_suffix = f"-part-{idx:02d}" if len(parts) > 1 else ""
            group_slug = _slugify(group_name + part_suffix)

            pack_out_dir = str(base_out / group_slug)
            pack_file = packs_dir / f"{group_slug}.md"

            mapping = {
                "codebase_name": codebase_name,
                "out_dir": pack_out_dir,
                "oracle_cmd": oracle_cmd,
                "oracle_flags": oracle_flags,
                "extra_files": extra_files,
                "ticket_root": ticket_root,
                "ticket_glob": ticket_glob,
                "ticket_paths": ",".join(str(p) for p in part),
                "ticket_max_files": str(min(len(part), max(1, group_max_files))),
                "group_name": group_name,
                "group_slug": group_slug,
                "mode": mode,
            }

            content = _render_template(template, mapping)
            pack_file.write_text(content, encoding="utf-8")

            manifest_groups.append(
                {
                    "group": group_name,
                    "slug": group_slug,
                    "part": idx,
                    "pack_path": str(pack_file),
                    "out_dir": pack_out_dir,
                    "attached_paths": [str(p) for p in part],
                    "original_tickets": sorted(group_originals.get(group_name, [])),
                }
            )

    (base_out / "_groups.json").write_text(
        json.dumps(grouping_report, indent=2, sort_keys=True),
        encoding="utf-8",
    )

    if dup_map:
        (base_out / "_duplicates.json").write_text(
            json.dumps(dup_map, indent=2, sort_keys=True),
            encoding="utf-8",
        )

    if dedupe_plan:
        (base_out / "_dedupe_plan.json").write_text(
            json.dumps(dedupe_plan, indent=2, sort_keys=True),
            encoding="utf-8",
        )

    if cluster_meta:
        pairs_out = [
            {"a": a, "b": b, **scores} for (a, b), scores in sorted(dup_pairs.items())
        ]
        (base_out / "_dupes_possible.json").write_text(
            json.dumps({"clusters": cluster_meta, "pairs": pairs_out}, indent=2, sort_keys=True),
            encoding="utf-8",
        )

    (base_out / "manifest.json").write_text(
        json.dumps({"groups": manifest_groups}, indent=2, sort_keys=True),
        encoding="utf-8",
    )

    print(f"[OK] wrote packs to: {packs_dir}")
    print(f"[OK] wrote grouping map: {base_out / '_groups.json'}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
```

skills/oraclepack-tickets-pack-grouped/scripts/lint_attachments.py
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
            errors.append(f"Step {step.n}: found '_tickets_bundle' reference (direct-ticket packs must not use bundle).")

        if re.search(r"mapfile\s+-t\s+__tickets\s+<\s+<\(", joined) is None:
            errors.append(f"Step {step.n}: missing mapfile ticket discovery stanza.")

        if re.search(r"ticket_args=\(\)", joined) is None or re.search(r"ticket_args\+\=\(\s*(-f|--file)\b", joined) is None:
            errors.append(f"Step {step.n}: missing ticket_args builder (ticket_args+=(-f \"$p\")).")

        if re.search(r"\$\{ticket_args\[@\]\}", joined) is None:
            errors.append(f"Step {step.n}: missing ${'{'}ticket_args[@]{'}'} usage in oracle invocation.")

        # Heuristic: ensure we did not hardcode a non-existent bundle path.
        if re.search(r'(?<!\S)(-f|--file)(?!\S)\s+"[^"\n]*_tickets_bundle', joined):
            errors.append(f"Step {step.n}: contains a hardcoded _tickets_bundle attachment.")

    if errors:
        for e in errors:
            print(f"[ERROR] {e}", file=sys.stderr)
        sys.exit(1)

    print("[OK] Direct-ticket lint passed.")


def main() -> None:
    p = argparse.ArgumentParser(description="Lint ticket-driven Stage-1 packs (direct-ticket mode).")
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

skills/oraclepack-tickets-pack-grouped/scripts/render_group_packs.py
```
#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import re
from pathlib import Path
from typing import Dict


def _render_template(template: str, mapping: Dict[str, str]) -> str:
    out = template
    for key, val in mapping.items():
        out = out.replace("{{" + key + "}}", val)
    unresolved = sorted(set(re.findall(r"\{\{([^}]+)\}\}", out)))
    if unresolved:
        raise ValueError(f"Unresolved template placeholders: {unresolved}")
    return out


def main() -> int:
    if len(sys.argv) == 1:
        print("Select how to run:")
        print("1) Use defaults (no args)")
        print("2) Provide custom args (show usage)")
        choice = input("Enter choice [1-2]: ").strip() or "1"
        if choice == "2":
            print("Usage: render_group_packs.py --manifest manifest.json --out-dir out")
            return 0

    p = argparse.ArgumentParser(description="Render group-specific bundle packs from manifest.")
    p.add_argument("--manifest", default="manifest.json")
    p.add_argument("--out-dir", default="docs/oracle-questions-sharded")
    p.add_argument("--template", default="/home/user/.codex/skills/oraclepack-tickets-pack-grouped/references/tickets-pack-template-bundle.md")
    p.add_argument("--codebase-name", default="Unknown")
    p.add_argument("--oracle-cmd", default="oracle")
    p.add_argument("--oracle-flags", default="--files-report")
    p.add_argument("--extra-files", default="")
    p.add_argument("--ticket-root", default=".tickets")
    p.add_argument("--ticket-glob", default="**/*.md")
    p.add_argument("--mode", default="tickets-bundle")
    args = p.parse_args()

    manifest_path = Path(args.manifest)
    if not manifest_path.exists():
        raise SystemExit(f"[ERROR] manifest not found: {manifest_path}")

    manifest = json.loads(manifest_path.read_text(encoding="utf-8"))
    template = Path(args.template).read_text(encoding="utf-8")

    out_dir = Path(args.out_dir)
    packs_dir = out_dir / "packs"
    packs_dir.mkdir(parents=True, exist_ok=True)

    for group in manifest.get("groups", []):
        slug = group["slug"]
        tickets = group["tickets"]
        pack_dir = packs_dir / slug
        pack_dir.mkdir(parents=True, exist_ok=True)

        pack_path = pack_dir / f"oracle-pack_{slug}.md"
        bundle_path = pack_dir / f"tickets_bundle_{slug}.md"
        out_run_dir = pack_dir / "out"

        mapping = {
            "codebase_name": args.codebase_name,
            "out_dir": str(out_run_dir),
            "oracle_cmd": args.oracle_cmd,
            "oracle_flags": args.oracle_flags,
            "extra_files": args.extra_files,
            "ticket_root": args.ticket_root,
            "ticket_glob": args.ticket_glob,
            "ticket_paths": ",".join(tickets),
            "ticket_bundle_path": str(bundle_path),
            "mode": args.mode,
        }

        content = _render_template(template, mapping)
        pack_path.write_text(content, encoding="utf-8")
        group["pack_path"] = str(pack_path)

    manifest_path.write_text(json.dumps(manifest, indent=2, sort_keys=True), encoding="utf-8")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
```

skills/oraclepack-tickets-pack-grouped/scripts/shard_tickets.py
```
#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import math
import re
from dataclasses import dataclass
from pathlib import Path
from typing import Dict, List, Tuple

STOPWORDS = {
    "the", "and", "for", "with", "from", "this", "that", "into", "over", "under", "when",
    "then", "than", "else", "only", "must", "should", "could", "would", "will", "shall",
    "ticket", "tickets", "oraclepack", "oracle", "pack", "packs",
}

SECTION_KEYS = {"summary", "acceptance", "criteria", "background", "context"}


@dataclass
class Ticket:
    path: Path
    text: str
    tokens: List[str]
    vector: List[float]


def _tokenize(text: str) -> List[str]:
    text = text.lower()
    text = re.sub(r"[^a-z0-9]+", " ", text)
    return [t for t in text.split() if len(t) >= 3 and t not in STOPWORDS]


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8", errors="replace")
    except FileNotFoundError:
        return ""


def _extract_repr(text: str, stem: str, max_chars: int) -> str:
    lines = text.splitlines()
    heading = ""
    sections: List[str] = []
    capture = False
    for line in lines:
        s = line.strip()
        if not heading and s.startswith("#"):
            heading = s.lstrip("#").strip()
        if s.startswith("#"):
            key = s.lstrip("#").strip().lower()
            capture = any(k in key for k in SECTION_KEYS)
            continue
        if capture and s:
            sections.append(s)
        if len(" ".join(sections)) >= max_chars:
            break
    body = " ".join(sections)
    base = " ".join([stem, heading, body])
    return base[:max_chars]


def _tfidf_vectors(texts: List[str]) -> Tuple[List[List[float]], List[str]]:
    docs = [
        [tok for tok in _tokenize(t)]
        for t in texts
    ]
    vocab: Dict[str, int] = {}
    df: Dict[str, int] = {}
    for toks in docs:
        seen = set()
        for tok in toks:
            if tok not in vocab:
                vocab[tok] = len(vocab)
            if tok not in seen:
                df[tok] = df.get(tok, 0) + 1
                seen.add(tok)

    n_docs = len(docs)
    idf = [0.0] * len(vocab)
    for tok, idx in vocab.items():
        idf[idx] = math.log((1 + n_docs) / (1 + df.get(tok, 1))) + 1.0

    vectors: List[List[float]] = []
    for toks in docs:
        tf: Dict[int, float] = {}
        for tok in toks:
            tf[vocab[tok]] = tf.get(vocab[tok], 0.0) + 1.0
        vec = [0.0] * len(vocab)
        for idx, count in tf.items():
            vec[idx] = count * idf[idx]
        # L2 normalize
        norm = math.sqrt(sum(v * v for v in vec)) or 1.0
        vec = [v / norm for v in vec]
        vectors.append(vec)

    inv_vocab = [None] * len(vocab)
    for tok, idx in vocab.items():
        inv_vocab[idx] = tok
    return vectors, inv_vocab


def _cosine(a: List[float], b: List[float]) -> float:
    return sum(x * y for x, y in zip(a, b))


def _centroid(vectors: List[List[float]]) -> List[float]:
    if not vectors:
        return []
    dim = len(vectors[0])
    out = [0.0] * dim
    for v in vectors:
        for i, val in enumerate(v):
            out[i] += val
    n = float(len(vectors)) or 1.0
    out = [v / n for v in out]
    norm = math.sqrt(sum(v * v for v in out)) or 1.0
    return [v / norm for v in out]


def _kmeans_split(vectors: List[List[float]], k: int, iters: int = 10) -> List[List[int]]:
    if k <= 1:
        return [list(range(len(vectors)))]
    # deterministic init: first k vectors
    centroids = [vectors[i][:] for i in range(k)]
    for _ in range(iters):
        clusters = [[] for _ in range(k)]
        for idx, v in enumerate(vectors):
            best = 0
            best_score = -1.0
            for c_idx, c in enumerate(centroids):
                score = _cosine(v, c)
                if score > best_score:
                    best_score = score
                    best = c_idx
            clusters[best].append(idx)
        new_centroids = []
        for cluster in clusters:
            if cluster:
                new_centroids.append(_centroid([vectors[i] for i in cluster]))
            else:
                new_centroids.append(centroids[len(new_centroids)])
        centroids = new_centroids
    return clusters


def main() -> int:
    if len(sys.argv) == 1:
        print("Select how to run:")
        print("1) Use defaults (no args)")
        print("2) Provide custom args (show usage)")
        choice = input("Enter choice [1-2]: ").strip() or "1"
        if choice == "2":
            print("Usage: shard_tickets.py --ticket-root .tickets --out-dir out")
            return 0

    p = argparse.ArgumentParser(description="Shard tickets into topic/domain groups.")
    p.add_argument("--ticket-root", default=".tickets")
    p.add_argument("--ticket-glob", default="**/*.md")
    p.add_argument("--ticket-paths", default="")
    p.add_argument("--out-dir", default="docs/oracle-questions-sharded")
    p.add_argument("--min-sim", type=float, default=0.15)
    p.add_argument("--max-group-size", type=int, default=25)
    p.add_argument("--min-group-size", type=int, default=1)
    p.add_argument("--max-bundle-chars", type=int, default=200000)
    p.add_argument("--repr-chars", type=int, default=2000)
    p.add_argument("--use-llm-for-ambiguous", action="store_true")
    args = p.parse_args()

    ticket_root = Path(args.ticket_root)
    if args.ticket_paths:
        paths = [Path(p.strip()) for p in args.ticket_paths.split(",") if p.strip()]
    else:
        paths = sorted(ticket_root.glob(args.ticket_glob), key=lambda p: str(p)) if ticket_root.exists() else []

    texts: List[str] = []
    tickets: List[Ticket] = []
    for pth in paths:
        txt = _read_text(pth)
        rep = _extract_repr(txt, pth.stem, args.repr_chars)
        texts.append(rep)

    vectors, vocab = _tfidf_vectors(texts)
    for pth, txt, vec in zip(paths, texts, vectors):
        tickets.append(Ticket(path=pth, text=txt, tokens=_tokenize(txt), vector=vec))

    groups: Dict[str, List[int]] = {}
    loose: List[int] = []
    for idx, t in enumerate(tickets):
        try:
            rel = t.path.relative_to(ticket_root)
        except ValueError:
            loose.append(idx)
            continue
        if len(rel.parts) >= 2:
            g = rel.parts[0]
            groups.setdefault(g, []).append(idx)
        else:
            loose.append(idx)

    # Compute centroids for subdir groups
    centroids: Dict[str, List[float]] = {}
    for g, idxs in groups.items():
        centroids[g] = _centroid([tickets[i].vector for i in idxs])

    # Assign loose tickets by similarity
    reasons: Dict[int, Dict[str, object]] = {}
    for idx in loose:
        best_g = None
        best_sim = -1.0
        for g, c in centroids.items():
            sim = _cosine(tickets[idx].vector, c)
            if sim > best_sim:
                best_sim = sim
                best_g = g
        if best_g is not None and best_sim >= args.min_sim:
            groups.setdefault(best_g, []).append(idx)
            reasons[idx] = {"assigned_to": best_g, "sim": best_sim, "reason": "tfidf"}
        else:
            groups.setdefault("misc", []).append(idx)
            reasons[idx] = {
                "assigned_to": "misc",
                "sim": best_sim,
                "reason": "ambiguous" if not args.use_llm_for_ambiguous else "ambiguous_llm_needed",
            }

    # Merge small groups
    if args.min_group_size > 1 and len(groups) > 1:
        for g in sorted(list(groups.keys())):
            if g == "misc":
                continue
            if len(groups[g]) < args.min_group_size:
                # merge into nearest group
                g_centroid = _centroid([tickets[i].vector for i in groups[g]])
                best_g = None
                best_sim = -1.0
                for og, c in centroids.items():
                    if og == g:
                        continue
                    sim = _cosine(g_centroid, c)
                    if sim > best_sim:
                        best_sim = sim
                        best_g = og
                if best_g:
                    groups.setdefault(best_g, []).extend(groups[g])
                    del groups[g]

    # Split large groups using deterministic kmeans
    final_groups: Dict[str, List[int]] = {}
    for g in sorted(groups.keys()):
        idxs = groups[g]
        idxs_sorted = sorted(idxs, key=lambda i: str(tickets[i].path))
        total_chars = sum(len(_read_text(tickets[i].path)) for i in idxs_sorted)
        k = max(
            1,
            math.ceil(len(idxs_sorted) / max(1, args.max_group_size)),
            math.ceil(total_chars / max(1, args.max_bundle_chars)),
        )
        if k <= 1:
            final_groups[g] = idxs_sorted
            continue
        clusters = _kmeans_split([tickets[i].vector for i in idxs_sorted], k)
        part = 1
        for cluster in clusters:
            if not cluster:
                continue
            slug = f"{g}-part-{part:02d}"
            final_groups[slug] = [idxs_sorted[i] for i in cluster]
            part += 1

    # Build manifest
    out_dir = Path(args.out_dir)
    out_dir.mkdir(parents=True, exist_ok=True)
    manifest_groups = []
    for g in sorted(final_groups.keys()):
        idxs = final_groups[g]
        vecs = [tickets[i].vector for i in idxs]
        centroid = _centroid(vecs)
        top_terms = []
        for i, score in sorted(enumerate(centroid), key=lambda x: -x[1])[:8]:
            if score <= 0:
                continue
            top_terms.append(vocab[i])
        sims = []
        for i in idxs:
            sims.append(_cosine(tickets[i].vector, centroid))
        conf = sum(sims) / float(len(sims)) if sims else 0.0

        manifest_groups.append(
            {
                "slug": g,
                "tickets": [str(tickets[i].path) for i in idxs],
                "keywords": top_terms,
                "confidence": conf,
            }
        )

    manifest = {"groups": manifest_groups}
    (out_dir / "manifest.json").write_text(json.dumps(manifest, indent=2, sort_keys=True), encoding="utf-8")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
```

skills/oraclepack-tickets-pack-grouped/scripts/validate_pack.py
```
from pathlib import Path
import runpy

COMMON = Path(__file__).resolve().parents[2] / "oraclepack-tickets-pack-common" / "scripts" / "validate_pack.py"
if not COMMON.exists():
    raise SystemExit(f"[ERROR] Shared validator not found: {COMMON}")

runpy.run_path(str(COMMON), run_name="__main__")
```

skills/oraclepack-tickets-pack-grouped/scripts/validate_shards.py
```
#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import subprocess
from pathlib import Path
from typing import Dict


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8", errors="replace")
    except FileNotFoundError:
        return ""


def main() -> int:
    if len(sys.argv) == 1:
        print("Select how to run:")
        print("1) Use defaults (no args)")
        print("2) Provide custom args (show usage)")
        choice = input("Enter choice [1-2]: ").strip() or "1"
        if choice == "2":
            print("Usage: validate_shards.py --manifest manifest.json")
            return 0

    p = argparse.ArgumentParser(description="Validate sharded packs manifest.")
    p.add_argument("--manifest", default="manifest.json")
    p.add_argument("--max-bundle-chars", type=int, default=200000)
    p.add_argument(
        "--validator",
        default="/home/user/.codex/skills/oraclepack-tickets-pack-common/scripts/validate_pack.py",
    )
    args = p.parse_args()

    manifest_path = Path(args.manifest)
    if not manifest_path.exists():
        raise SystemExit(f"[ERROR] manifest not found: {manifest_path}")

    manifest = json.loads(manifest_path.read_text(encoding="utf-8"))
    counts: Dict[str, int] = {}

    for group in manifest.get("groups", []):
        for t in group.get("tickets", []):
            counts[t] = counts.get(t, 0) + 1

    bad = [t for t, c in counts.items() if c != 1]
    if bad:
        raise SystemExit(f"[ERROR] Tickets assigned !=1 times: {bad}")

    for group in manifest.get("groups", []):
        pack_path = Path(group.get("pack_path", ""))
        if not pack_path.exists():
            raise SystemExit(f"[ERROR] pack missing: {pack_path}")

        # validate pack
        subprocess.run(
            [
                "python3",
                args.validator,
                "--mode",
                "bundle",
                str(pack_path),
            ],
            check=True,
        )

        # size check
        total = 0
        for t in group.get("tickets", []):
            total += len(_read_text(Path(t)))
        if total > args.max_bundle_chars:
            raise SystemExit(
                f"[ERROR] group '{group.get('slug')}' exceeds max bundle chars: {total} > {args.max_bundle_chars}"
            )

    print("[OK] Sharded packs manifest validated.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
```

</source_code>