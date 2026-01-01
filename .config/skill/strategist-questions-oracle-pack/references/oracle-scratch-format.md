<!-- path: ~/.codex/skills/strategist-questions-oracle-pack/references/oracle-scratch-format.md -->
# oracle usage examples

---

## add attachments

```bash
oracle \
  --browser-attachments always \
  --browser-input-timeout 5s \
  -p "Run the UI smoke test." \
  -f "packages/"
```

## copy to clipboard

```bash
oracle --render --copy \
  -p "Explain the CLI architecture and control flow: index.ts -> registry init -> list/run commands. Identify error-handling gaps, async pitfalls, and how you'd improve UX (help text, exit codes, structured output). Provide proposed code changes." \
  -f "codefetch/packages.md" \
  -f "packages/cli/src/index.ts" \
  -f "packages/cli/src/commands/list.ts" \
  -f "packages/cli/src/commands/run.ts"
```

---

## oracle/codefetch

```bash
oracle \
  --browser-attachments always \
  -p "Using packages.md + attached files, give me a dependency graph and a build order. Then propose 5 repo hygiene improvements." \
  -f "codefetch/packages.md" \
  -f "packages/**/package.json" \
  -f "packages/**/tsconfig.json"
```

---

## other scratch examples

```bash
oracle \
  -p "Walk through the UI smoke test" \
  --file "src/**/*.ts"
```

```bash
oracle \
  -p "Review these changes and propose fixes" \
  -f "src/**/*.ts,!**/*.test.ts"
```

```bash
oracle \
  -p "Using packages.md + attached files, give me a dependency graph and a build order. Then propose 5 repo hygiene improvements." \
  -f "codefetch/packages.md" \
  -f "packages/**/package.json" \
  -f "packages/**/tsconfig.json"
```

```bash
oracle --render --copy \
  -p "Explain the CLI architecture and control flow: index.ts -> registry init -> list/run commands. Identify error-handling gaps, async pitfalls, and how you'd improve UX (help text, exit codes, structured output). Provide proposed code changes." \
  -f "codefetch/packages.md" \
  -f "packages/cli/src/index.ts" \
  -f "packages/cli/src/commands/list.ts" \
  -f "packages/cli/src/commands/run.ts"
```

```bash
oracle  \
  --prompt "Read packages.md and explain: (1) what each package does, (2) how workflows are executed end-to-end, (3) where to add a new workflow, and (4) the top 5 refactors to reduce coupling." \
  --file "codefetch/packages.md"
```

