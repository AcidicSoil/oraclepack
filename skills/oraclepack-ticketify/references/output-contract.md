# Output contract (required fields)

## <out_dir>/_tickets_index.json

Required top-level keys:
- `version` (int)
- `tickets_dir` (string)
- `ticket_glob` (string)
- `include` (array of strings)
- `exclude` (array of strings)
- `top_n` (int)
- `selected_count` (int)
- `tickets` (array)

Each `tickets[]` entry must include:
- `ticket_id` (string)
- `path` (string, repo-relative POSIX)
- `title` (string)
- `summary` (string)
- `sha256` (string, hex)

## <out_dir>/_actions.json

Required top-level keys:
- `version` (int)
- `out_dir` (string)
- `count` (int)
- `actions` (array)

Each `actions[]` entry must include:
- `action_id` (string)
- `ticket_id` (string)
- `ticket_path` (string)
- `source_line` (int)
- `kind` (string: `task` or `acceptance`)
- `text` (string)

## <out_dir>/_actions.md

- Human-readable list of actions.
- Must be deterministic and derived only from tickets + the canonical action list.

## <prd_path>

- Markdown PRD suitable for `task-master parse-prd <prd_path>`.
- Must not include timestamps or machine-specific info.

## <pack_path>

- Markdown document with:
  - Exactly one fenced code block labeled `bash`.
  - No other code fences anywhere.
  - Inside the bash block: exactly 20 steps, `# 01)`..`# 20)`.
  - Steps do not rely on shell variables set by prior steps.
