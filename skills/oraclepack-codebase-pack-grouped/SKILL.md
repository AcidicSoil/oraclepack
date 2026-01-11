---
name: oraclepack-codebase-pack-grouped
description: Generate multiple runner-ingestible oraclepack Stage-1 packs grouped by codebase topic/domain (subdir + deterministic inference) with direct code attachments. Use when the user wants per-topic/per-domain mini-packs for a target repo/project/codebase instead of ticket folders, with strict 20-step schema and validation.
---

# oraclepack-codebase-pack-grouped (Stage 1)

## Goal

Produce **multiple** codebase-driven Stage-1 packs, one per inferred topic/domain, with direct code attachments. Each pack is schema-safe and self-contained.

## Use this skill

Use when the user wants separate packs per topic/domain grouped by a target repo/project/codebase, not a `.tickets/` folder.


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
- `code_root` (`.`)
- `code_glob` (`**/*`)
- `code_paths` (empty; comma-separated explicit files; if present, ignore glob)
- `code_max_files` (`200`)
- `group_mode` (`subdir+infer`)
- `group_min_score` (`0.10`)
- `group_max_files` (`200`)
- `group_max_chars` (`200000`)
- `ignore_dirs` (empty; comma-separated; merged with defaults)
- `.gitignore` is respected by default; patterns are applied first, then `ignore_dirs`.
- `include_exts` (empty; uses default extension allowlist)
- `exclude_glob` (empty; comma-separated glob patterns)
- `mode` (`codebase-grouped-direct`)
- `command_mode` (`inline`)
- `inline_max_tokens` (`60000`)
- `allow_overwrite` (`false`)

Notes:
- `YYYY-MM-DD-HHMMSS` is computed at pack generation time for default `out_dir`.
- If oracle flag support is uncertain, omit unsupported flags; never invent flags.

## Workflow (deterministic)

1) Read:
- `references/codebase-grouping.md`
- `references/attachment-minimization.md`
- `references/codebase-pack-template.md`

If `2`, ask for KEY=value args and run with those.
If `3`, propose 2–4 concrete KEY=value bundles tailored to the target codebase (focus on `code_root`, `include_exts`, `exclude_glob`, `group_max_files`, `code_max_files`), then let the user pick one or edit; run with the chosen bundle.
Otherwise run with defaults.

3) Generate packs (deterministic grouping + per-group pack files):

```bash
python3 /home/user/.codex/skills/oraclepack-codebase-pack-grouped/scripts/generate_grouped_packs.py \
  codebase_name=oraclepack \
  out_dir=docs/oracle-questions-2026-01-08
```

Outputs:
- `{{out_dir}}/packs/*.md` (one pack per group/part)
- `{{out_dir}}/_groups.json` (group -> file list)
- `{{out_dir}}/manifest.json` (group -> pack paths + files)

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
