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
