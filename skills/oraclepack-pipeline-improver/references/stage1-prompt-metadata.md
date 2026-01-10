<!-- # path: oraclepack-pipeline-improver/references/stage1-prompt-metadata.md -->
# Stage 1 prompt-embedded metadata (proposed)

Goal: improve downstream parsing for Stage 2/3 without changing the oracle pack schema.

## Constraints

- Do not change the pack’s structural schema unless an explicit migration path is provided.
- Embed metadata inside the prompt text (the `-p` payload) in a parseable, consistent format.

## Recommended metadata block (inside the prompt)

Add at the *very top* of each prompt:

```

[oraclepack-meta]
QuestionId: 07
Category: permissions
Reference: src/server/auth/middleware.ts
ExpectedArtifacts:

* src/server/auth/**
* docs/plans/...
  [/oraclepack-meta]

```

## Parsing rules (deterministic)

- The block starts with `[oraclepack-meta]` and ends with `[/oraclepack-meta]`.
- Keys are case-sensitive as shown.
- Multi-line lists are allowed for `ExpectedArtifacts`.
- If any key is missing, treat it as `null` and continue.

## Minimal required keys (recommended)

- QuestionId
- Category
- Reference
- ExpectedArtifacts (optional but recommended)

If the Stage 1 generator cannot produce these, it should write `Unknown` values explicitly rather than omitting keys.
