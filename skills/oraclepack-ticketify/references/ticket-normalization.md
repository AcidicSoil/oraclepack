# Ticket normalization

This document defines how tickets are interpreted and converted into canonical artifacts.

## Ticket file format assumptions

- Tickets are Markdown files (`*.md`) under `.tickets/` or `tickets/`.
- YAML frontmatter is optional. If present, it must start at the beginning of the file:
  - First line: `---`
  - Ends at the next `---` on its own line.
- Only simple `key: value` pairs are supported for deterministic parsing (no nested YAML).

If tickets do not follow this, do not guess—fail-fast with an `ERROR:` message.

## Derived fields

### ticket_id
Derivation order:
1) Frontmatter key `id`
2) Frontmatter key `ticket`
3) Frontmatter key `key`
4) Fallback: basename without extension

Fail-fast if any duplicates occur among selected tickets.

### title
Derivation order:
1) Frontmatter key `title`
2) First Markdown H1 line (`# ...`)
3) Fallback: filename

### summary
Deterministic heuristic:
- First non-heading paragraph after the title/frontmatter.
- Concatenate lines until a blank line or ~240 chars (whichever comes first).

## Derived actions

Actions are derived from:

1) Markdown task list items:
- `- [ ] ...`
- `- [x] ...`
- `* [ ] ...`

2) Bullets under a heading named `Acceptance Criteria` (case-insensitive):
- Any `- ...` or `* ...` bullet until the next heading.

Ordering:
- Primary: ticket path order (already stable)
- Then by line number
- Then by text
