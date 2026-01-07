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
