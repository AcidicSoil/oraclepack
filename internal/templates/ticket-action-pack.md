# Ticket Action Pack

```bash
out_dir=".oraclepack/ticketify"
--write-output

# 01)
echo "Build tickets index"

# 02)
echo "Generate actions json"

# 03)
echo "Generate tickets PRD"

# 04)
echo "Prep taskmaster inputs"

# 05)
task-master parse-prd .taskmaster/docs/prd.md

# 06)
task-master analyze-complexity --research

# 07)
task-master expand --all --research

# 08)
echo "Prepare headless automation"

# 09)
if command -v gemini >/dev/null 2>&1; then
  gemini run "Select next tasks" --write-output ".oraclepack/ticketify/next.json"
else
  echo "Skipped: gemini missing"
fi

# 10)
if command -v codex >/dev/null 2>&1; then
  codex exec "Implement tasks" --write-output ".oraclepack/ticketify/codex-implement.md"
else
  echo "Skipped: codex missing"
fi

# 11)
if command -v codex >/dev/null 2>&1; then
  codex exec "Verify changes" --write-output ".oraclepack/ticketify/codex-verify.md"
else
  echo "Skipped: codex missing"
fi

# 12)
if command -v gemini >/dev/null 2>&1; then
  gemini run "Review outputs" --write-output ".oraclepack/ticketify/gemini-review.json"
else
  echo "Skipped: gemini missing"
fi

# 13)
if command -v codex >/dev/null 2>&1; then
  codex exec "Prepare fixes" --write-output ".oraclepack/ticketify/codex-fixes.md"
else
  echo "Skipped: codex missing"
fi

# 14)
echo "Summarize results"

# 15)
echo "Prepare release notes"

# 16)
if command -v codex >/dev/null 2>&1; then
  codex exec "Draft PR description" --write-output ".oraclepack/ticketify/PR.md"
else
  echo "Skipped: codex missing"
fi

# 17)
echo "Finalize checklist"

# 18)
echo "Post-run cleanup"

# 19)
echo "Audit artifacts"

# 20)
echo "Done"
```
