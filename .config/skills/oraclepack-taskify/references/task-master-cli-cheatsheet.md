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
