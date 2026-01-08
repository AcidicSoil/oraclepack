<!-- # path: oraclepack-pipeline-improver/references/run-manifest-spec.md -->
# Run manifest spec (proposed)

This defines the minimum content for run artifacts to enable traceability, resume/rerun, and Stage 3 processing.

## `.oraclepack/runs/<pack_id>/run.json`

**Required fields (proposed minimum):**
- `pack_id` (string)
- `pack_path` (string)
- `pack_hash` (string; full hash)
- `created_at` (RFC3339 string)
- `git_sha` (string | null)
- `oraclepack_version` (string | TODO if not available)
- `oracle_version` (string | TODO if not available)
- `max_parallel` (number | null)
- `ci` (boolean)
- `providers` (object | TODO if not available)
- `models` (object | TODO if not available)

**Notes:**
- If any value cannot be derived, set it to `null` and record a `warnings[]` entry rather than inventing it.

## `.oraclepack/runs/<pack_id>/steps.json`

Represent steps as an array of 20 items ordered by `step_id`.

**Per-step fields (proposed minimum):**
- `step_id` (string `"01"`..`"20"`)
- `question_id` (string | null) — if present in the pack/prompt metadata
- `category` (string | null)
- `reference` (string | null)
- `invocation_hash` (string) — hash of canonical invocation inputs (prompt + attachments + provider/model knobs)
- `output_path` (string)
- `output_hash` (string | null)
- `status` (enum: `pending` | `ok` | `failed` | `skipped`)
- `attempts` (number)
- `last_error` (string | null)
- `started_at` (RFC3339 string | null)
- `finished_at` (RFC3339 string | null)

## Hashing (proposed)

### `pack_hash`
- Compute from a canonical representation of the pack file:
  - normalize line endings,
  - remove non-semantic whitespace if safe (TODO/Unknown: exact pack grammar),
  - hash the resulting bytes.

### `invocation_hash`
- Compute from:
  - prompt text (including embedded mini-metadata),
  - attachment file contents (hash of contents, not just paths),
  - provider + model identifiers,
  - deterministic knobs (temperature, etc.) if applicable.

If attachment content hashing cannot be performed, record **Unknown/TODO** and do not enable caching based on incomplete inputs.
