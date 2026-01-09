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

1) Size control (mandatory; fail fast):

- Run `oracle --dry-run summary --files-report ...` for the **largest** group pack (or each pack if unsure).
- Enforce caps:
  - browser: ≤ 60,000 tokens total input per step
  - api: ≤ 180,000 tokens total input per step
- If exceeded, reduce via `group_max_files` or use explicit `ticket_paths`.

1) Validate every pack (mandatory):

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

1) Detect possible duplicates (if `dedupe_mode != off`):

- Signature: filename stem + first heading + first `dedupe_body_chars` chars.
- Compute `jaccard` + `overlap` between tickets.
- Duplicate edge rule:
  - `overlap >= dedupe_overlap_hi` OR (`jaccard >= dedupe_jaccard` AND `overlap >= dedupe_overlap_lo`)
- Connected components become duplicate clusters.
- Canonical: largest content length; tie-break lexicographic.
- Delta vs redundant:
  - delta if unique token ratio >= `dedupe_delta_min` OR heading differs materially.
  - redundant otherwise.

1) Seed groups by subdir:

- For any path under `ticket_root/<group>/...`, assign to group `<group>`.
- Tickets directly under `ticket_root/` are "loose".

1) Infer loose tickets into groups (if any groups exist):

- Build a token set for each group from:
  - group name tokens
  - ticket filenames (stem tokens)
  - first Markdown heading line (if present)
- For each loose ticket, compute Jaccard overlap score with each group token set.
- If `max_score >= group_min_score`, assign to the best group (stable tie-break by group name).
- Otherwise, assign to `misc`.

1) If no groups exist:

- Put all tickets into a single group named `root`.

1) Merge duplicates into primary group:

- `report`: attach all tickets in the cluster to the canonical’s group.
- `prune`: attach canonical + delta only; drop redundant from attachments.
- `merge`: create `out_dir/_ticket_merges/cluster-XXXX.md` and attach only the merged file.
- Emit `_dupes_possible.json`, `_duplicates.json`, and `_dedupe_plan.json`.

1) Split oversized groups:

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

- contracts/interfaces: OK
- invariants: OK
- caching/state: OK
- background jobs: OK
- observability: OK
- permissions: OK
- migrations: OK
- UX flows: OK
- failure modes: OK
- feature flags: OK

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
  echo "WARNING: no tickets resolved for group '{{group_name}}'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/11-permissions-authz-gaps.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #11  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: permissions
Horizon: NearTerm
ROI: 4.1 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, identify permission/authz boundaries implied by tickets (file access, command execution, network); propose safe defaults.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 12) ROI=3.9 impact=6 confidence=0.68 effort=2 horizon=NearTerm category=permissions reference={{group_slug}}

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
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/12-permissions-secrets-config.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #12  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: permissions
Horizon: NearTerm
ROI: 3.9 (impact=6, confidence=0.68, effort=2)

Question:
Using the attached tickets as the primary context, identify secrets/config handling needs (API keys, tokens); propose secure config discovery and redaction.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 13) ROI=3.8 impact=6 confidence=0.66 effort=3 horizon=MidTerm category=migrations reference={{group_slug}}

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
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/13-migrations-schema-migrations.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #13  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: migrations
Horizon: MidTerm
ROI: 3.8 (impact=6, confidence=0.66, effort=3)

Question:
Using the attached tickets as the primary context, identify any required migrations (schema/format/CLI flags); define migration strategy and compat approach.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 14) ROI=3.7 impact=6 confidence=0.64 effort=3 horizon=MidTerm category=migrations reference={{group_slug}}

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
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/14-migrations-backfill-plan.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #14  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: migrations
Horizon: MidTerm
ROI: 3.7 (impact=6, confidence=0.64, effort=3)

Question:
Using the attached tickets as the primary context, define any needed backfill/one-time transforms; estimate risks; define verification plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 15) ROI=4.6 impact=6 confidence=0.74 effort=1 horizon=Immediate category=UX flows reference={{group_slug}}

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
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/15-ux-flows-user-journeys.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #15  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: UX flows
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, identify UX/TUI workflows implied by tickets; define user journey states and expected outputs.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 16) ROI=4.3 impact=6 confidence=0.72 effort=2 horizon=Immediate category=UX flows reference={{group_slug}}

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
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/16-ux-flows-edge-cases.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #16  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: UX flows
Horizon: Immediate
ROI: 4.3 (impact=6, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify edge cases in UX flows (cancel, resume, partial runs); define minimal UX behavior.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 17) ROI=4.9 impact=7 confidence=0.78 effort=1 horizon=Immediate category=failure modes reference={{group_slug}}

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
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/17-failure-modes-timeouts-retries.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #17  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: failure modes
Horizon: Immediate
ROI: 4.9 (impact=7, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, define timeouts/retries behavior for external calls; define failure classification and operator actions.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 18) ROI=4.4 impact=7 confidence=0.74 effort=2 horizon=Immediate category=failure modes reference={{group_slug}}

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
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/18-failure-modes-rollback-plan.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #18  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: failure modes
Horizon: Immediate
ROI: 4.4 (impact=7, confidence=0.74, effort=2)

Question:
Using the attached tickets as the primary context, define rollback plan for partial runs and how to preserve artifacts; define 'safe to re-run' semantics.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 19) ROI=4.0 impact=6 confidence=0.70 effort=2 horizon=NearTerm category=feature flags reference={{group_slug}}

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
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/19-feature-flags-flag-plan.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #19  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: feature flags
Horizon: NearTerm
ROI: 4.0 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, define feature-flag strategy for rollout (scopes, defaults, telemetry); ensure compat for existing users.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 20) ROI=3.8 impact=6 confidence=0.68 effort=2 horizon=NearTerm category=feature flags reference={{group_slug}}

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
{{oracle_cmd}}   {{oracle_flags}}   --write-output "{{out_dir}}/20-feature-flags-compat-rollout.md"   "${ticket_args[@]}"   {{extra_files}}   -p "$(cat <<'PROMPT'
Strategist question #20  (ticket-driven, group: {{group_name}})

Reference: {{group_slug}}
Category: feature flags
Horizon: NearTerm
ROI: 3.8 (impact=6, confidence=0.68, effort=2)

Question:
Using the attached tickets as the primary context, define minimal compat-safe rollout plan and guardrails; include fallback behavior and monitoring gates.

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

## Coverage check

- contracts/interfaces: OK
- invariants: OK
- caching/state: OK
- background jobs: OK
- observability: OK
- permissions: OK
- migrations: OK
- UX flows: OK
- failure modes: OK
- feature flags: OK

```

skills/oraclepack-tickets-pack-grouped/scripts/generate_grouped_packs.py
```

# !/usr/bin/env python3
from **future** import annotations

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
    return_dt.date.today().isoformat()

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
    heading, snippet =_read_signature(p)
    toks.update(_tokenize(heading))
    toks.update(_tokenize(snippet))
    return toks

def _signature_tokens(p: Path, body_chars: int) -> set:
    heading =_read_heading(p)
    body =_read_text(p)
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
        tokens[key] =_signature_tokens(p, body_chars)
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

if **name** == "**main**":
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
    steps =_parse_steps(fence)

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

if **name** == "**main**":
    main()

```

skills/oraclepack-tickets-pack-grouped/scripts/render_group_packs.py
```

# !/usr/bin/env python3
from **future** import annotations

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

if **name** == "**main**":
    raise SystemExit(main())

```

skills/oraclepack-tickets-pack-grouped/scripts/shard_tickets.py
```

# !/usr/bin/env python3
from **future** import annotations

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
    out = [0.0] *dim
    for v in vectors:
        for i, val in enumerate(v):
            out[i] += val
    n = float(len(vectors)) or 1.0
    out = [v / n for v in out]
    norm = math.sqrt(sum(v* v for v in out)) or 1.0
    return [v / norm for v in out]

def _kmeans_split(vectors: List[List[float]], k: int, iters: int = 10) -> List[List[int]]:
    if k <= 1:
        return [list(range(len(vectors)))]
    # deterministic init: first k vectors
    centroids = [vectors[i][:] for i in range(k)]
    for _in range(iters):
        clusters = [[] for_ in range(k)]
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

if **name** == "**main**":
    raise SystemExit(main())

```

skills/oraclepack-tickets-pack-grouped/scripts/validate_pack.py
```

from pathlib import Path
import runpy

COMMON = Path(**file**).resolve().parents[2] / "oraclepack-tickets-pack-common" / "scripts" / "validate_pack.py"
if not COMMON.exists():
    raise SystemExit(f"[ERROR] Shared validator not found: {COMMON}")

runpy.run_path(str(COMMON), run_name="**main**")

```

skills/oraclepack-tickets-pack-grouped/scripts/validate_shards.py
```

# !/usr/bin/env python3
from **future** import annotations

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

if **name** == "**main**":
    raise SystemExit(main())

```

</source_code>
