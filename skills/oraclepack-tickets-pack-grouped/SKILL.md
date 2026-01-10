---
name: oraclepack-tickets-pack-grouped
description: Generate multiple runner-ingestible oraclepack Stage-1 packs grouped by ticket topic/domain (subdir + deterministic inference) with direct ticket attachments. Use when the user wants per-topic/per-domain mini-packs, grouped via subdirectory discovery and inferred assignment of loose tickets, with strict 20-step schema and validation.
---

# oraclepack-tickets-pack-grouped (Stage 1)

## Goal

Produce **multiple** ticket-driven Stage-1 packs, one per inferred topic/domain, with direct ticket attachments. Each pack is schema-safe and self-contained.

## Use this skill

Use when the user wants separate packs per topic/domain, grouped by `.tickets/` subdirectories plus deterministic inference for loose tickets.


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

## Inputs (parse trailing KEY=value; last-one-wins)

Supported keys (defaults in parentheses):
- `codebase_name` (`Unknown`)
- `out_dir` (`docs/oracle-questions-YYYY-MM-DD-HHMMSS`)
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
- `allow_overwrite` (`false`)

Notes:
- `YYYY-MM-DD-HHMMSS` is computed at pack generation time for default `out_dir`.
- If oracle flag support is uncertain, omit unsupported flags; never invent flags.
- `.gitignore` is respected by default when selecting tickets (primary), before any explicit filters.

## Workflow (deterministic)

1) Read:
- `references/ticket-grouping.md`
- `references/attachment-minimization.md`
- `references/tickets-pack-template.md`

If `2`, ask for KEY=value args and run with those.
If `3`, propose 2–4 concrete KEY=value bundles tailored to the ticket set (focus on `ticket_root`, `ticket_glob`, `ticket_max_files`, `group_max_files`, `group_max_chars`), then let the user pick one or edit; run with the chosen bundle.
Otherwise run with defaults.

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
3) Offer suggestions
```

If `2`, collect args and use them in the commands below.
If `3`, propose 2–4 concrete KEY=value bundles tailored to the ticket set, let the user pick or edit, then use the chosen bundle in the commands below.

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
