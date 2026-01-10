# Oracle scratch playbook (NOT a pack format)

This document is for **manual debugging / one-off oracle runs**. It is **not** runner-ingestible oraclepack pack format.

## When to use scratch vs pack

Use the **pack** (`references/oracle-pack-template.md`) when:
- You need a strict 20-step Stage-1 pack for oraclepack ingestion.
- You want deterministic execution and validation via `scripts/validate_pack.py`.

Use **scratch** when:
- You need a single oracle call to explore something quickly.
- You are iterating on prompt wording before committing it into the pack.
- You want to test attachment choices with `--dry-run`.

## Scratch workflow

1) Start with one focused question.
2) Attach 0–2 high-signal files.
3) Use a quoted heredoc prompt to avoid shell-escaping issues.
4) If results are weak, add *one* more attachment (or refine the question).

## Pack-adjacent scratch example (single run)

Example pattern (edit paths/flags to match your environment):

- Uses the quoted heredoc prompt style.
- Shows the optional `--dry-run summary` style (if supported).

```bash
# (optional) preview:
# oracle --dry-run summary --files-report -f "README.md" -p "$(cat <<'PROMPT'
# Explain the repo’s main entrypoint and how requests flow through the system.
# If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
# PROMPT
# )"

oracle \
  --files-report \
  -f "README.md" \
  -p "$(cat <<'PROMPT'
Goal: Understand the repo’s main entrypoints and primary request flow.

Answer format:
1) Direct answer (bullets, evidence-cited)
2) Risks/unknowns
3) Next smallest concrete experiment (one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"
```

## Promoting scratch into the pack

When a scratch run looks good:
- Convert it into a numbered step in the pack.
- Add the strict header tokens (`ROI= impact= confidence= effort= horizon= category= reference=`).
- Add `--write-output "{{out_dir}}/NN-<slug>.md"`.
- Ensure category is one of the fixed 10 and update Coverage check accordingly.

```
