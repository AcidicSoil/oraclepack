# Determinism rules

These rules are mandatory for any artifact produced by `oraclepack-ticketify`.

## Stable ordering

1) Ticket file ordering:
- Normalize each ticket file path to repo-relative POSIX form (forward slashes).
- Sort lexicographically by that normalized path.

2) Selection:
- Apply include filters first (keep only matches), then exclude (remove matches).
- Apply `top_n` after sorting and filtering.
- Result is the first `top_n` paths in sorted order.

3) Derived entities:
- Actions are ordered by `(ticket_path, source_line, text)`.

## Stable IDs

- `ticket_id`:
  - Prefer an explicit frontmatter key in this order: `id`, `ticket`, `key`.
  - Else fallback: filename without extension.
- `action_id`: `<ticket_id>-A01`, `<ticket_id>-A02`, ... in stable per-ticket order.

## Stable output formatting

- JSON:
  - UTF-8
  - `indent=2`
  - `sort_keys=true`
  - trailing newline
- Markdown:
  - no timestamps
  - no machine-specific data (absolute paths, usernames, hostnames)
