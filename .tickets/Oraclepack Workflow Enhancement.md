Title:

* Stabilize oraclepack “oracle-pack” pipeline and add profile-based context + Stage-3 synthesis for actionable follow-through

Summary:

* The current two-step workflow generates an `oracle-pack` Markdown file (20 `oracle` calls) via Codex skills/Gemini CLI commands, then runs it through the `oraclepack` Go wrapper to produce outputs and state/report artifacts.

    Workflow Optimization for Oracl…

    Workflow Optimization for Oracl…

* A key failure mode is schema/format drift in the pack file (human-doc + machine-ingest combined), which can break ingestion; an example drift is step headers using an em dash (`# 01 — ROI=…`) while the documented contract expects `# NN)`.

    Workflow Optimization for Oracl…

* Requested outcome: improve workflow continuity, enable richer context injection without breaking the strict pack contract, and add an automatic next stage that turns the “final twenty questions/answers” into actionable implementation steps with minimal human interaction.

    Workflow Optimization for Oracl…

    Workflow Optimization for Oracl…

Background / Context:

* Workflow stages:

  * Stage 1: LLM authoring creates `docs/oracle-pack-YYYY-MM-DD.md` containing 20 `oracle` commands (with ROI metadata and per-step output paths).

        Workflow Optimization for Oracl…

  * Stage 2: `oraclepack` executes the pack to produce 20 outputs under `oracle-out/...` plus state/report JSON artifacts.

        Workflow Optimization for Oracl…

* `oraclepack` is a wrapper around `oracle` intended for batched/bulk requests.

    Workflow Optimization for Oracl…

* Core concern: “disconnection” after the 20-question output; desire to chain into a useful, actionable implementation stage.

    Workflow Optimization for Oracl…

    Workflow Optimization for Oracl…

Current Behavior (Actual):

* Pack file acts as both:

  * Human documentation, and

  * A strict machine-ingest contract, making formatting drift a pipeline-breaking event.

        Workflow Optimization for Oracl…

* Documented/expected step-header schema (`# NN)`) can drift to alternative formats (example: `# 01 — ROI=…`), risking parse/validation failures.

    Workflow Optimization for Oracl…

* High-risk edits include adding additional code fences (especially additional bash fences) or introducing lines that accidentally match the step-header pattern.

    Workflow Optimization for Oracl…

Expected Behavior:

* Packs remain schema-stable and reliably parse/validate/run across providers (Codex skills, Gemini CLI commands).

    Workflow Optimization for Oracl…

* Richer “skill context” can be injected without changing the pack’s ingest shape (no added code fences / no header drift).

    Workflow Optimization for Oracl…

* After Stage 2 produces 20 outputs + report JSON, a subsequent stage can automatically convert results into actionable implementation steps.

    Workflow Optimization for Oracl…

Requirements:

* Preserve the non-negotiable pack contract:

  * Pack is Markdown containing exactly one `bash` code block; the first bash block is executed.

        Workflow Optimization for Oracl…

  * Steps are identified via header pattern `# NN)` with sequential numbering starting at `01`.

        Workflow Optimization for Oracl…

  * Prelude content before the first step header executes once.

        Workflow Optimization for Oracl…

* Standardize Stage-1 generation to the strict header form `# NN)` (avoid em dash variants).

    Workflow Optimization for Oracl…

* Add a hard gate between Stage 1 and Stage 2:

  * Make `oraclepack validate` mandatory before `oraclepack run` (prevent schema drift reaching execution).

        Workflow Optimization for Oracl…

* Provide schema-safe extensibility for context injection:

  * Allow context to be injected via `oracle -p` prompt text and/or `oracle -f` file/directory attachments (preferred for larger context).

        Workflow Optimization for Oracl…

  * Use prelude variables and templating only if it does not interfere with header parsing.

        Workflow Optimization for Oracl…

  * Avoid adding extra code fences or lines resembling step headers.

        Workflow Optimization for Oracl…

* Implement “Context Profiles” as file-backed bundles:

  * Add `skills/oracle-pack/references/profiles/<name>.md` and inject via `oracle -f "$profile_file"` plus a short prompt preamble (“Follow the attached profile standards”).

  * Add an optional `profile` input to the Stage-1 skill/command, with backwards-compatible behavior when absent.

        Workflow Optimization for Oracl…

* Add a first-class Stage 3 synthesis step:

  * Provide a command shape such as `oraclepack synthesize --in oracle-out --report pack.report.json --out docs/implementation-pack-YYYY-MM-DD.md` that reads the 20 outputs, extracts proposed changes/file targets/tests, and emits a new validated pack for implementation.

        Workflow Optimization for Oracl…

  * Support minimal-interaction execution for Stage 3 (e.g., headless usage via Codex/Gemini CLI).

        Workflow Optimization for Oracl…

Out of Scope:

* Not provided.

Reproduction Steps:

* Not provided.

Environment:

* `oraclepack` Go program wrapping `oracle` CLI.

* Stage-1 generation tools mentioned: Codex skills, Gemini CLI commands.

    Workflow Optimization for Oracl…

* OS/CI details: Unknown.

Evidence:

* Attachment: “Workflow Optimization for Oraclepack.md”.

    Workflow Optimization for Oracl…

    Workflow Optimization for Oracl…

* Example schema drift called out: step headers using `# 01 — ROI=…` vs documented `# NN)`.

    Workflow Optimization for Oracl…

* Proposed validation/run sequence:

  * `oraclepack validate docs/oracle-pack-YYYY-MM-DD.md`

  * `oraclepack list docs/oracle-pack-YYYY-MM-DD.md`

  * `oraclepack run docs/oracle-pack-YYYY-MM-DD.md --no-tui --run-all --stop-on-fail=true --out-dir oracle-out`

        Workflow Optimization for Oracl…

Decisions / Agreements:

* Treat the pack as a stable intermediate representation (IR) and keep context flowing only through `-p` and `-f` to avoid breaking the ingest contract.

* Prefer “Context Profiles” as a file-backed mechanism located under `skills/oracle-pack/references/profiles/`.

* Add a validation gate (`validate` before `run`) to reduce pipeline breakage from formatting drift.

    Workflow Optimization for Oracl…

Open Items / Unknowns:

* Exact current parser/validator behavior regarding em dash header variants (whether it currently accepts them) is not provided; only that it is avoidable schema drift.

    Workflow Optimization for Oracl…

* Exact filenames/paths for current `SKILL.md` and template files in the repo are referenced conceptually but not provided in full.

    Workflow Optimization for Oracl…

* Whether `oraclepack synthesize` already exists or is a new feature request is not provided; it is described as a proposed product shape.

    Workflow Optimization for Oracl…

Risks / Dependencies:

* Dependency on `oracle` CLI flags and behavior (`-p/--prompt`, `-f/--file`, `--write-output`, `--files-report`, `--dry-run`).

    Workflow Optimization for Oracl…

* Risk of pack invalidation from added code fences, additional bash blocks, or accidental header-like lines.

    Workflow Optimization for Oracl…

* Cross-provider consistency risk (Codex skills vs Gemini CLI commands) unless Stage 1 is standardized around a shared template/profile mechanism.

    Workflow Optimization for Oracl…

Acceptance Criteria:

* Pack schema stability

  * Packs validate when they contain exactly one bash block and step headers are strictly `# NN)` starting at `01`.

  * Stage-1 generation output uses `# NN)` (no em dash variant) across providers.

        Workflow Optimization for Oracl…

* Validation gate

  * Workflow includes a required `oraclepack validate` pass before any `oraclepack run`.

        Workflow Optimization for Oracl…

* Context Profiles

  * A `profile` selection results in `oracle -f "$profile_file"` being attached per step without adding new code fences or breaking parsing.

        Workflow Optimization for Oracl…

  * Absence of `profile` preserves current behavior (backwards compatible).

        Workflow Optimization for Oracl…

* Stage 3 synthesis

  * A synthesis step can consume `oracle-out` outputs + report JSON and emit a follow-on implementation pack intended to be validated and run.

        Workflow Optimization for Oracl…

Priority & Severity (if inferable from text):

* Not provided.

Labels (optional):

* enhancement

* workflow

* cli

* parsing

* validation

* context-bundles

* automation

---
