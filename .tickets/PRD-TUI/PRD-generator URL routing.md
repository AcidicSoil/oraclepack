Title:

* Add PRD-generator project URL routing + ticketfy→PRD “micro-pack” generation; avoid `tickets_prd.md` pack-parse failure

Summary:

* `oraclepack-ticketfy` generates `tickets_prd.md`, but oraclepack cannot execute it as a pack and fails with `Error: invalid pack structure: no bash code block found`.
* Need an automated way (via TUI option) to route a PRD-generation run to a specific “PRD generator” ChatGPT Project URL (one of many project URLs), without sending entire packs for simple calls.

Background / Context:

* User wants oraclepack (TUI) to support selecting/storing multiple ChatGPT Project URLs and sending the chosen URL “with oraclepack” for PRD generation.
* Current understanding (per assistant): runner injects selected URL into `oracle …` invocations via `--chatgpt-url <selected>`, and there is a URL picker/store concept with defaults.
* User proposes: parse/transform `tickets_prd.md` (artifact) into a valid oraclepack that calls `oracle` against the PRD-generator project URL, and include additional context missing from `tickets_prd.md` as part of that stage.

Current Behavior (Actual):

* Running/using `tickets_prd.md` as a pack fails validation with: `Error: invalid pack structure: no bash code block found`.
* PRD generation flow is not currently integrated as a first-class TUI option that:

  * selects a PRD-generator ChatGPT Project URL “from one of many project urls”, and
  * runs a lightweight, single-purpose call without requiring a full pack workflow.
* Missing UI support (per assistant) for per-step ChatGPT URL selection in Overrides Wizard; URL selection appears global for the run today.

Expected Behavior:

* `tickets_prd.md` should be usable as input to PRD generation without being treated as a runnable pack.
* TUI should offer an option to run PRD generation routed to a specific PRD-generator ChatGPT Project URL (selectable from stored URLs), optionally scoped to only the PRD step(s).
* Workflow should support “simple oracle calls” that do not require sending entire packs.

Requirements:

* Do not attempt to execute `tickets_prd.md` directly as an oraclepack pack; instead, generate a valid runnable pack that calls `oracle` and attaches `tickets_prd.md`.
* Add an automated way (TUI option) to select and apply a PRD-generator ChatGPT Project URL for the PRD generation flow.
* Support “one of many project urls” (multiple ChatGPT Project URLs) and the ability to target PRD generation steps specifically (per-step URL override).
* Include “missing context” needed by the PRD generator as part of the stage (ticket-derived context bundle alongside `tickets_prd.md`).
* Provide a lightweight path for single-shot oracle calls (CLI/TUI) that does not require pack parsing (per assistant suggestion: new `oraclepack call` mode).

Out of Scope:

* Not provided.

Reproduction Steps:

1. Generate `tickets_prd.md` via `oraclepack-ticketfy`.
2. Attempt to run `tickets_prd.md` through oraclepack as if it were a runnable pack.
3. Observe error: `Error: invalid pack structure: no bash code block found`.

Environment:

* OS: Unknown
* oraclepack version/commit: Unknown
* oracle CLI version: Unknown
* ticketfy skill version: Unknown
* TUI vs no-TUI: Unknown

Evidence:

* Error message: `Error: invalid pack structure: no bash code block found`.
* Proposed workaround “micro-pack” example (per assistant) that wraps an `oracle` call and attaches `.taskmaster/docs/tickets_prd.md`.
* Suggested architecture (per assistant):

  * Add Overrides Wizard step “ChatGPT URL” writing to `pendingOverrides.ChatGPTURL`.
  * Generate `.oraclepack/ticketify/prd_context.md` and `.oraclepack/ticketify/prd-generator.pack.md`.
  * Optional new command path bypassing `internal/pack/parser.go` (pack parsing) via `oraclepack call ...`.

Decisions / Agreements:

* Constraint acknowledged: `tickets_prd.md` is an artifact and not runnable as a pack; running it directly will fail due to missing ` ```bash … ``` ` fenced block (per assistant).
* Preferred direction (user): convert artifact → valid oraclepack routed to PRD-generator project URL with added context for higher-quality PRD generation.

Open Items / Unknowns:

* Exact location/path of generated `tickets_prd.md` in the repo (example paths were suggested, but actual is not provided).
* How ChatGPT Project URLs are stored/serialized (format, file path, scope rules) in the current implementation: Not provided.
* Whether headless/CI runs need a CLI override flag (`--chatgpt-url` or `--chatgpt-url-name`): implied as needed, but exact requirement not confirmed.

Risks / Dependencies:

* Must preserve strict pack schema expectations (bash fenced block requirement); any solution that weakens validation could impact runner ingest reliability.
* Per-step URL routing depends on runtime overrides and a TUI flow to author/apply them (missing UI support noted).
* PRD generator quality depends on providing adequate context (new context bundle artifact needed).

Acceptance Criteria:

* [ ] Attempting to run `tickets_prd.md` directly is no longer part of the intended flow; documentation/UX guides user to PRD-generation pathway instead.
* [ ] Ticketfy stage outputs (or an adjacent stage) include:

  * [ ] a deterministic PRD context bundle artifact (e.g., `prd_context.md`), and
  * [ ] a valid runnable PRD generator pack (single `bash` fence) that attaches `tickets_prd.md` (+ context bundle) and invokes `oracle`.
* [ ] TUI exposes an option like “Generate enhanced PRD via PRD Generator Project” that:

  * [ ] lets user select a stored ChatGPT Project URL (PRD Generator),
  * [ ] applies it to the PRD generation run (optionally step-targeted).
* [ ] (If implemented) `oraclepack call` (single-shot) can run an `oracle` invocation with selected ChatGPT URL + attachments, without requiring pack parsing.

Priority & Severity (if inferable from text):

* Not provided.

Labels (optional):

* enhancement
* tui
* workflow
* prd
* url-routing
* pack-validation
