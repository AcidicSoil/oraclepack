# Inference-first discovery (Stage 1 packs)

Goal: pick the *right* 1–2 attachments per step without over-attaching, by inferring repo shape from a small set of anchors.

## Principles

- Prefer **evidence** (actual files) over assumptions.
- Start broad with **cheap, high-signal** files; only then zoom in.
- If a file/path doesn’t exist: record `Unknown` and continue.
- Keep steps self-contained: each step should succeed without relying on shared shell setup.

## Deterministic discovery order

1) **Repo identity + entrypoints**
- `README*` (first ~200 lines)
- top-level manifests (language/framework/package)
- main entrypoints (server start, CLI main, app bootstrap) if obvious from tree

2) **Configuration + environment**
- example config files
- `.env.example` or equivalent (if present)
- CI config files (to infer build/test and deploy steps)

3) **Public surface**
- routing tables / controllers / handlers
- schema/contract definitions (API specs, message schemas, type definitions)
- CLI command registration (if applicable)

4) **Data + jobs + operations**
- data models and storage adapters
- migrations directory (if present)
- background job definitions and worker entrypoints (if present)
- logging/metrics/tracing configuration (if present)

## Turning discovery into step-specific attachments

For each planned step:
- Choose 1 “definition” file (where the thing is declared).
- Optionally choose 1 “use-site” file (where the thing is used/enforced).
- If you can’t confidently pick: attach fewer files and use `reference=Unknown`.

Then write the prompt so the oracle can request missing artifact patterns when needed.

## What to do when evidence is insufficient

- Set `reference=Unknown` in the step header.
- In the prompt, explicitly ask for:
  - “the exact missing file/path pattern(s) to attach next”
- Keep attachments minimal; do not guess file paths.
