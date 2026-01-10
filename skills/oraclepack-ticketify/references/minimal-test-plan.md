# Minimal test plan

## Fixture

Create a tiny repo fixture:

.tickets/
  001-login.md
  002-payments.md

Example `001-login.md`:
- Contains `# Login` heading
- Contains at least one `- [ ]` task item
- Contains `## Acceptance Criteria` with bullet(s)

Example `002-payments.md`:
- Contains frontmatter with `id: PAY-2` and `title: Payments`
- Contains at least one `- [ ]` task item

## Expected deterministic outputs

Run the generator (via the produced Action Pack steps 01–03) twice without changing files.

Expect byte-identical:
- `<out_dir>/_tickets_index.json`
- `<out_dir>/_actions.json`
- `<out_dir>/_actions.md`
- `<prd_path>`

And pack validation should pass:
- `<pack_path>` has exactly one `bash` fence and 20 steps.

## Failure-mode checks

1) Missing tickets dir
- Ensure neither `.tickets/` nor `tickets/` exists.
- Expect fail-fast with:
  - stderr includes `ERROR: tickets_dir not found`
  - non-zero exit code

2) Empty match set
- Create `.tickets/` but no `*.md` matches (or filters exclude all).
- Expect fail-fast `ERROR: no tickets remain after include/exclude filtering` (or similar).

3) Duplicate IDs
- Two files with same `id:` in frontmatter (or same filename-derived ID).
- Expect fail-fast `ERROR: duplicate derived ticket_id`.

4) Unreadable file
- Make a ticket unreadable (permissions).
- Expect fail-fast `ERROR: unreadable ticket`.

5) Invalid pack formatting
- Add a second code fence to the pack.
- `scripts/validate_ticket_action_pack.py` must fail non-zero.
