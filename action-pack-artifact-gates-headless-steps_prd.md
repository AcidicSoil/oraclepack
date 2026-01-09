1. Overview

Problem statement. Oraclepack Action Packs are executed as literal shell steps (documented by the tickets as `bash -lc ...` in the repo root), but oraclepack’s current “special handling” (command detection, flag injection/overrides, and validation via `oracle --dry-run summary`) is anchored to commands that begin with `oracle`, causing steps that use other CLIs (`tm`/`task-master`, `codex`, `gemini`) to be silently excluded from override/validation pipelines. This breaks Action Pack workflows that rely on non-oracle tools and placeholder steps in `ticket-action-pack.md` (notably steps 09–13 and 16) that must be replaced by headless/non-interactive automation while remaining ingestible as a 20-step single-`bash`-fence pack.

Who has the problem. Developers using oraclepack (CLI/TUI and CI) to run Action Packs for ticket-to-PR workflows, especially in environments where codex/gemini may be absent or interactive behavior can stall runs.

Why current solutions fail.

* Detection/validation is regex-anchored to `oracle` (`^(\\s*)(oracle)\\b` noted in tickets), so non-oracle steps can be skipped rather than validated.
* Placeholder steps in the action pack template prevent end-to-end execution with expected artifacts, and missing binaries on PATH can fail runs unless explicitly guarded.

Success metrics (measurable).

* Validation coverage: 0 “silently excluded” steps due solely to lacking an `oracle ...` prefix; steps beginning with `tm`/`task-master`, `codex`, or `gemini` must be detected and included in validation output.
* Artifact gates: for the Action Pack workflow, required artifacts are produced or an explicit “skipped due to missing binary” record is emitted:

  * Step 09: `.oraclepack/ticketify/next.json`
  * Step 10: `.oraclepack/ticketify/codex-implement.md`
  * Step 11: `.oraclepack/ticketify/codex-verify.md` OR `.oraclepack/ticketify/gemini-review.json`
  * Step 16: `.oraclepack/ticketify/PR.md`
* Headless execution: new steps do not block on interactivity in CI (codex uses non-interactive `codex exec` for automation). ([OpenAI Developers][1])
* Backwards compatibility: existing behavior for `oracle ...` steps remains unchanged; steps 01–07 semantics in `ticket-action-pack.md` remain unchanged.

Constraints, integrations, assumptions.

* Execution semantics remain “literal shell via `bash -lc ...`” (so environment/login-shell behaviors remain relevant).  ([gnu.org][2])
* The Action Pack “pack contract” remains: exactly one fenced `bash` block; exactly 20 steps; steps are self-contained; output files are written as shown; artifact paths are stable unless versioned.
* Tooling integration targets: `tm`/`task-master`, OpenAI Codex CLI (`codex exec`), and `gemini` CLI (exact flags/overrides for non-oracle tools are not fully specified by the provided evidence; treat as an explicit open question and implement detection + guardrails first).  ([OpenAI Developers][1])

1. Capability Tree (Functional Decomposition)

Capability: Pack specification and template management

* Feature: Action Pack execution semantics documentation (MVP)

  * Description: Publish user-facing docs for how Action Pack steps execute and what oraclepack “special handling” applies to.
  * Inputs: Pack file content; current execution implementation behavior.
  * Outputs: Documentation pages and CLI/TUI help text updates.
  * Behavior: Document that steps execute as `bash -lc ...` in repo root and that special handling currently applies by command classification; explicitly list supported command prefixes and how they’re handled.
* Feature: `ticket-action-pack.md` template replacement (MVP)

  * Description: Replace placeholder steps 09–13 and 16 with headless `gemini` + non-interactive `codex exec` automation while keeping the pack contract.
  * Inputs: Template source; required artifact contract.
  * Outputs: Updated `ticket-action-pack.md` content that remains ingestible.
  * Behavior: Keep steps 01–07 unchanged; ensure steps produce required artifacts; add `command -v` guards to degrade gracefully when binaries are missing.

Capability: Step classification and dispatch

* Feature: Multi-tool command detection (MVP)

  * Description: Expand step command detection beyond `oracle` to include `tm`/`task-master`, `codex`, and `gemini`.
  * Inputs: Step shell text.
  * Outputs: Classified command prefix + normalized “tool kind” for each step.
  * Behavior: Replace oracle-anchored regex with a multi-prefix detector; preserve existing classification for `oracle ...` commands.
* Feature: Tool-specific dispatch policy (non-MVP, gated by open questions)

  * Description: Define what override injection means per tool and how it is applied.
  * Inputs: Tool kind; user config; step command.
  * Outputs: Possibly transformed command + metadata describing applied overrides.
  * Behavior: For `oracle`, preserve existing injection/behavior; for other tools, initially perform detection-only unless explicit policy is defined.

Capability: Execution engine and environment guards

* Feature: Headless execution mode selection (MVP)

  * Description: Ensure steps use non-interactive invocations when specified by the pack (e.g., `codex exec` for Codex automation).
  * Inputs: Step definition; tool kind; environment.
  * Outputs: Executed process result and logs.
  * Behavior: Run steps as `bash -lc "<step>"`; when using codex, prefer `codex exec` semantics suitable for scripts/CI.  ([OpenAI Developers][1])
* Feature: Missing-binary graceful skip (MVP)

  * Description: When `codex`/`gemini`/`tm` is missing, record an explicit skip instead of failing the whole run (where the step is designed to be optional under missing tools).
  * Inputs: Tool kind; PATH lookup result.
  * Outputs: Step result marked Skipped with reason.
  * Behavior: Enforce “skip vs block” rules per step; do not hide the skip; write into state/report.

Capability: Validation and override pipeline

* Feature: Validation inclusion for non-oracle tools (MVP)

  * Description: Ensure steps using `tm`/`codex`/`gemini` are included in validation output (not skipped because they are non-oracle).
  * Inputs: Pack; classified steps; installed tool availability.
  * Outputs: Validation report listing status per step and per-tool checks.
  * Behavior: Keep current oracle-only validation intact; add additional validators that (at minimum) check tool presence + required artifact gates for relevant steps.
* Feature: Artifact gate validation (MVP)

  * Description: Validate that specific steps produce required files under `.oraclepack/ticketify/` (or emit explicit skip records).
  * Inputs: Step results; filesystem state.
  * Outputs: Pass/fail/skip gate status per required artifact.
  * Behavior: Treat artifact existence requirements as acceptance gates; align with CI/canary gating.

Capability: Artifact contract and run state/reporting

* Feature: Standardized `.oraclepack/ticketify/` outputs (MVP)

  * Description: Standardize and document expected output artifact names and locations for the Action Pack workflow.
  * Inputs: Step definitions; pack run.
  * Outputs: `.oraclepack/ticketify/next.json`, `codex-implement.md`, `codex-verify.md`/`gemini-review.json`, `PR.md`.
  * Behavior: Do not move or rename without versioning; enforce stability via validation gates.
* Feature: Persistent run state + report (MVP)

  * Description: Persist per-step results and a summary report to support resumability and UX flows.
  * Inputs: Pack run events.
  * Outputs: `ticket-action-pack.state.json` and `ticket-action-pack.report.json` (names referenced by the tickets/UX notes).
  * Behavior: Track step status (Succeeded/Failed/Skipped/Blocked), logs pointers, timestamps, and artifact gate outcomes.

Capability: UX and operational observability

* Feature: Run journey states in TUI/CLI (MVP)

  * Description: Represent the “Selected → Parsed/Validated → Ready → Running(step N) → {Succeeded|Failed|Skipped|Blocked} → Completed/Resumable” state machine.
  * Inputs: State/report files; runner events.
  * Outputs: CLI/TUI rendering and exit codes suitable for CI.
  * Behavior: Display step-level status and explicit skip reasons (missing tool, ROI filtering, etc.).
* Feature: Operational alerts/telemetry hooks (non-MVP)

  * Description: Track known failure modes (missing PATH binaries, interactivity stalls, ROI-filter skips).
  * Inputs: Run events.
  * Outputs: Metrics/log events.
  * Behavior: Emit structured events suitable for CI parsing; highlight ROI-filter behavior explicitly.

1. Repository Structure + Module Definitions (Structural Decomposition)

Assumption: oraclepack is implemented in Go (ticket hints reference `**/*report*.go`, `**/*state*.go`). If the repo language differs, keep module boundaries and public exports, but translate file layout accordingly.

Proposed structure (new or refactored packages only; avoid “utils” buckets):

oraclepack/

* cmd/oraclepack/

  * main.go
  * cli.go (wire commands)
* internal/foundation/

  * errors.go
  * config.go
  * fs.go
  * clock.go
* internal/shell/

  * runner.go
  * env.go
* internal/pack/

  * types.go
  * parser.go
  * linter.go
* internal/dispatch/

  * detector.go
  * policy.go
* internal/tools/

  * types.go
  * oracle.go
  * codex.go
  * gemini.go
  * taskmaster.go
* internal/artifacts/

  * contract.go
  * gates.go
* internal/state/

  * models.go
  * writer.go
  * reader.go
  * report.go
* internal/validate/

  * validate.go
  * oracle_validator.go
  * tool_presence_validator.go
  * artifact_gate_validator.go
* internal/templates/

  * ticket_action_pack.go (template source or embed logic)
  * render.go
* internal/tui/ (if applicable)

  * screens.go
  * run_flow.go

Module definitions (single responsibility + exports):

Module: internal/foundation

* Responsibility: Common primitives (errors, config, filesystem IO abstraction, clock).
* Exports:

  * type `Config`
  * `LoadConfig() (Config, error)`
  * `ReadFile(path) ([]byte, error)`, `WriteFileAtomic(path, data) error`
  * `Now() time.Time`
  * typed errors (e.g., `ErrMissingBinary`)

Module: internal/shell

* Responsibility: Execute shell steps using the oraclepack contract (`bash -lc ...`) and capture outputs.
* Exports:

  * `RunBashLoginCommand(cmd string, env EnvSpec) (ExecResult, error)`
  * `DetectBinary(name string) (path string, ok bool)`
    Notes: `bash -lc` semantics rely on login-shell startup behavior; document and test accordingly. ([gnu.org][2])

Module: internal/pack

* Responsibility: Parse pack documents into typed steps and enforce pack-level contracts where applicable.
* Exports:

  * `ParsePack(content []byte) (Pack, error)`
  * `LintPack(pack Pack) ([]LintIssue, error)`
  * types: `Pack`, `Step`, `StepID`, `StepMeta`

Module: internal/dispatch

* Responsibility: Classify a step’s primary tool and apply dispatch policy decisions (detection-first).
* Exports:

  * `DetectTool(step Step) ToolKind`
  * `BuildExecutionPlan(pack Pack) ExecutionPlan`
  * `ApplyPolicy(step Step, cfg Config) (Step, PolicyNotes)` (oracle-preserving; other tools may be no-op initially)

Module: internal/tools

* Responsibility: Tool-specific knowledge that is safe to centralize (command builders, presence checks, recommended non-interactive forms).
* Exports:

  * `NonInteractiveHint(kind ToolKind) string`
  * `PresenceCheck(kind ToolKind) (ok bool, details string)`
  * `BuildCodexExec(prompt string, flags CodexFlags) string` (explicitly uses `codex exec` for automation). ([OpenAI Developers][1])

Module: internal/artifacts

* Responsibility: Define and evaluate artifact contracts (expected outputs, gating rules).
* Exports:

  * `ActionPackArtifactContract() ArtifactContract`
  * `EvaluateGates(contract ArtifactContract, fs FS) GateResults`
  * types: `ArtifactContract`, `ArtifactGate`, `GateResult`

Module: internal/state

* Responsibility: Persist and load run state and a summarized report.
* Exports:

  * `WriteState(path string, state RunState) error`
  * `ReadState(path string) (RunState, error)`
  * `WriteReport(path string, report RunReport) error`
  * types: `RunState`, `RunReport`, `StepResult`, `StepStatus`

Module: internal/validate

* Responsibility: Compose validators and produce a validation report.
* Exports:

  * `ValidatePack(pack Pack, cfg Config) (ValidationReport, error)`
  * validators:

    * `ValidateOracleDryRun(...)`
    * `ValidateToolPresence(...)`
    * `ValidateArtifactGates(...)`
      Back-compat: oracle dry-run behavior remains intact; new validators add coverage for non-oracle steps.

Module: internal/templates

* Responsibility: Render official pack templates (including updated `ticket-action-pack`).
* Exports:

  * `RenderTicketActionPack(params TemplateParams) ([]byte, error)`

1. Dependency Chain (layers, explicit “Depends on: […]”)

Foundation layer (no dependencies)

* internal/foundation: Depends on: []
* internal/shell: Depends on: [internal/foundation] (uses config/env/error/fs abstractions)

Pack layer

* internal/pack: Depends on: [internal/foundation]
* internal/state: Depends on: [internal/foundation]

Tool knowledge layer

* internal/tools: Depends on: [internal/foundation]

Dispatch layer

* internal/dispatch: Depends on: [internal/pack, internal/tools, internal/foundation]

Artifacts layer

* internal/artifacts: Depends on: [internal/foundation]

Validation layer

* internal/validate: Depends on: [internal/pack, internal/dispatch, internal/tools, internal/artifacts, internal/foundation]

Template layer

* internal/templates: Depends on: [internal/foundation, internal/artifacts] (template aligns to artifact contract)

Orchestration/CLI layer

* cmd/oraclepack: Depends on: [internal/pack, internal/dispatch, internal/shell, internal/state, internal/validate, internal/templates, internal/foundation]

TUI layer (optional)

* internal/tui: Depends on: [cmd/oraclepack wiring or internal/* interfaces, internal/state, internal/validate]

1. Development Phases (Phase 0…N; entry/exit criteria; tasks with dependencies + acceptance criteria + test strategy)

Phase 0: Foundation
Entry criteria: repo builds/tests run.
Tasks:

* [ ] Implement internal/foundation (depends on: none)

  * Acceptance criteria: config + fs helpers compile; typed errors exist.
  * Test strategy: unit tests for config parsing; atomic write semantics.
* [ ] Implement internal/shell runner wrapper for `bash -lc` (depends on: internal/foundation)

  * Acceptance criteria: can execute a trivial command and capture stdout/stderr/exit code deterministically.
  * Test strategy: unit tests using temp dirs; integration test executing `bash -lc "echo ok"`; document login-shell implications. ([gnu.org][2])
    Exit criteria: other modules can execute shell commands via a stable interface.

Phase 1: Pack parse + state/report models
Entry criteria: Phase 0 complete.
Tasks:

* [ ] Implement internal/pack ParsePack + step model (depends on: internal/foundation)

  * Acceptance criteria: parses step IDs and metadata; detects malformed packs.
  * Test strategy: unit tests with fixtures including 20-step contract variants.
* [ ] Implement internal/state RunState/RunReport + JSON read/write (depends on: internal/foundation)

  * Acceptance criteria: can write/read `*.state.json` and `*.report.json`; preserves step status transitions.
  * Test strategy: round-trip tests; schema validation tests.
    Exit criteria: CLI can parse a pack and write an empty initialized state/report file set.

Phase 2: Multi-tool detection and tool presence checks
Entry criteria: Phase 1 complete.
Tasks:

* [ ] Implement internal/tools ToolKind + presence checks (depends on: internal/foundation)

  * Acceptance criteria: `oracle`, `tm`/`task-master`, `codex`, `gemini` kinds; presence results include details.
  * Test strategy: unit tests with PATH manipulation via test harness.
* [ ] Implement internal/dispatch DetectTool (depends on: internal/pack, internal/tools)

  * Acceptance criteria: steps are classified even when not starting with `oracle`; preserves `oracle` behavior.
  * Test strategy: table-driven tests with step strings that begin with whitespace and each prefix.
    Exit criteria: validation/execution can identify non-oracle steps reliably.

Phase 3: Artifact contract and Action Pack template updates
Entry criteria: Phase 2 complete.
Tasks:

* [ ] Implement internal/artifacts Action Pack contract + gate evaluation (depends on: internal/foundation)

  * Acceptance criteria: gates cover `.oraclepack/ticketify/{next.json,codex-implement.md,codex-verify.md|gemini-review.json,PR.md}`.
  * Test strategy: unit tests that create/miss files and verify gate statuses.
* [ ] Update internal/templates to render updated `ticket-action-pack` steps 09–13 and 16 (depends on: internal/foundation, internal/artifacts)

  * Acceptance criteria: template preserves 20-step single-fence contract; steps 01–07 unchanged; new steps use `command -v` guards and include non-interactive codex invocation guidance (`codex exec`).  ([OpenAI Developers][1])
  * Test strategy: golden-file tests for rendered template; lint tests for contract invariants.
    Exit criteria: a “golden” `ticket-action-pack.md` fixture exists and matches required artifacts and contract.

Phase 4: Validation pipeline expansion (compat-safe)
Entry criteria: Phase 3 complete.
Tasks:

* [ ] Extend internal/validate to include non-oracle steps (depends on: internal/pack, internal/dispatch, internal/tools, internal/artifacts)

  * Acceptance criteria: validation output includes `tm`/`codex`/`gemini` steps; oracle dry-run validation remains unchanged.
  * Test strategy: integration tests for (a) oracle-only pack, (b) pack with only `codex` prefix, ensuring no silent exclusion.
* [ ] Implement artifact gate validation for relevant steps (depends on: internal/artifacts, internal/state)

  * Acceptance criteria: artifact gates pass/fail/skip recorded; explicit skip records when binaries missing.
  * Test strategy: integration tests in two environments: with and without `codex`/`gemini` on PATH.
    Exit criteria: `oraclepack validate` can be used as a CI gate for Action Pack artifacts.

Phase 5: Execution UX and documentation surfaces
Entry criteria: Phase 4 complete.
Tasks:

* [ ] Implement explicit step statuses and journey rendering (depends on: internal/state, internal/validate, internal/shell)

  * Acceptance criteria: statuses include Succeeded/Failed/Skipped/Blocked; reason strings are shown; resume boundary is represented.
  * Test strategy: end-to-end test running subset steps (09–11, 16) with `--no-tui` and verifying state/report outputs.
* [ ] Publish docs for execution semantics and tool handling (depends on: internal/dispatch behavior finalized)

  * Acceptance criteria: docs clearly state `bash -lc` execution semantics and command classification rules; includes how to run codex in automation using `codex exec`.  ([OpenAI Developers][1])
    Exit criteria: end-to-end usable MVP: updated Action Pack template + validation gates + explicit skips + stable artifacts.

1. User Experience

Personas.

* CI maintainer: wants deterministic validation gates and non-interactive runs.
* Developer running local TUI/CLI: wants clear step-by-step progress, resumability, and artifact locations.

Key flows.

* Validate Action Pack.

  * User runs `oraclepack validate <pack>`.
  * Output includes per-step tool classification and validation coverage; no non-oracle steps are omitted.
  * Artifact gates reported for steps with expected outputs.
* Run Action Pack (headless).

  * User runs `oraclepack run --no-tui <pack>`.
  * State transitions follow: Selected → Parsed/Validated → Ready → Running(step N) → Step status → Completed/Resumable.
  * If `codex`/`gemini` is missing, affected steps are Skipped with explicit reasons; run continues where safe.
* Inspect artifacts.

  * User checks `.oraclepack/ticketify/` for standardized outputs; report links or enumerates them.

UI/UX notes tied to capabilities.

* Always show “why skipped” (missing binary vs ROI filtering vs policy) to avoid the “it didn’t run” ambiguity.
* Treat artifact gates as first-class status indicators in both CLI and TUI.

1. Technical Architecture

System components.

* Pack parser: converts pack markdown to typed steps.
* Dispatcher: classifies steps by tool prefix; preserves existing oracle behavior.
* Shell runner: executes `bash -lc` commands, capturing stdout/stderr and exit codes (login-shell semantics affect startup files and env). ([gnu.org][2])
* Validators:

  * Oracle validator: current `oracle --dry-run summary` behavior for oracle steps (unchanged).
  * Tool presence validator: checks `tm`/`codex`/`gemini` availability.
  * Artifact gate validator: asserts required outputs exist or a skip record exists.
* State/report store: writes `ticket-action-pack.state.json` and `ticket-action-pack.report.json`.

Data models (minimum).

* Pack: `{id, title, steps[]}`
* Step: `{id, meta{ROI, impact, confidence, effort, horizon, category, reference}, shell, toolKind}`
* StepResult: `{stepId, status, exitCode, startedAt, endedAt, toolKind, skipReason?, artifactsProduced[], logsRef?}`
* RunState: `{packId, currentStepId?, resultsByStepId, environment{toolPresence}}`
* RunReport: `{summaryCounts, gates[], failures[], warnings[]}`
* ArtifactContract: list of required outputs per step (Action Pack specific).

External tools / APIs.

* Codex CLI: use `codex exec` for non-interactive automation; it is explicitly positioned for scripts/CI and supports JSONL output for machine consumption. ([OpenAI Developers][1])
* (Out of scope for MVP, but present as tickets): MCP publishing/exposure work exists; not enough evidence included here to define interfaces safely.  Background context on MCP: it is an open standard for connecting assistants to tools/data via MCP servers/clients. ([Anthropic][3])

Key decisions and trade-offs.

* Decision: Implement “detection + validation inclusion” for non-oracle tools before “override injection” for them.

  * Rationale: Tickets explicitly call out silent exclusion as a problem; they also state unknowns about what overrides should mean for `tm`/`codex`/`gemini`.
  * Trade-off: Less automation initially, but avoids incorrect injection behavior and preserves compatibility.
* Decision: Enforce artifact gates via validation rather than relying on tool-specific dry-runs for all tools.

  * Rationale: Required artifacts are explicit rollout gates in the tickets.

1. Test Strategy

Test pyramid targets.

* Unit tests: ~70% (parsers, detectors, gate evaluators, state read/write).
* Integration tests: ~25% (validate/run flows with temp repos and PATH manipulation).
* End-to-end tests: ~5% (run the “golden” Action Pack subset in two environments).

Coverage minimums.

* Line: 85%+
* Branch: 80%+
* Critical modules (dispatch, validate, artifacts, state): 90%+

Critical scenarios per module.

* internal/dispatch

  * Happy path: detect `oracle`, `tm`, `task-master`, `codex`, `gemini` with leading whitespace.
  * Error cases: empty step shell; multiple commands; ambiguous prefixes.
* internal/validate

  * Oracle-only pack: behavior unchanged.
  * Non-oracle-only pack: steps included (no silent exclusion).
  * Mixed pack: oracle validator runs for oracle steps; presence and artifact validators run for others.
* internal/artifacts

  * Gates pass when files exist.
  * Gates fail when files missing and step not skipped.
  * Gates skip when binary missing and skip record written.
* internal/state

  * Round-trip JSON persistence and status transitions.
* E2E (per ticket recommendation)

  * Run `oraclepack validate` then `oraclepack run --no-tui` for steps 09–11 and 16 in environments with/without `codex`/`gemini` on PATH; assert artifact-path acceptance criteria.

1. Risks and Mitigations

Risk: Undefined override semantics for non-oracle tools.

* Impact: High (wrong behavior or broken workflows).
* Likelihood: High (explicitly called out as unknown).
* Mitigation: Detection-first + artifact gates + presence checks; introduce per-tool overrides only when specified by tickets/spec.
* Fallback: Keep non-oracle override injection as no-op with explicit warnings.

Risk: Interactivity stalls in CI even with “headless” flags.

* Impact: High.
* Likelihood: Medium.
* Mitigation: Prefer `codex exec` (non-interactive mode) for Codex automation; add timeouts and “Blocked” status with guidance. ([OpenAI Developers][1])

Risk: Missing binaries on PATH cause failures instead of explicit skips.

* Impact: Medium–High.
* Likelihood: High.
* Mitigation: Standard `command -v` guards in template; presence validator; explicit skip records.

Risk: ROI filtering hides required steps.

* Impact: Medium.
* Likelihood: Medium.
* Mitigation: Ensure required steps have ROI metadata or document running with ROI threshold = 0 during canary; surface ROI-based skips explicitly.

Risk: Artifact path drift breaks consumers.

* Impact: High.
* Likelihood: Medium.
* Mitigation: Treat `.oraclepack/ticketify/*` paths as versioned contract; add validation gate tests and golden fixtures.

1. Appendix

Evidence used (provided).

* Public surface changes and back-compat constraints for Action Packs: execution semantics (`bash -lc`), oracle-only special handling, need multi-tool detection, validation inclusion, template step replacements, and standardized artifacts.
* Validation/monitoring gates and recommended experiment using a golden fixture + PATH/no-PATH environments.
* UX journey states and state/report filenames.

External references (behavioral grounding).

* Bash login-shell startup behaviors relevant to `bash -lc` execution. ([gnu.org][2])
* Codex non-interactive mode via `codex exec` for scripts/CI. ([OpenAI Developers][1])
* MCP high-level definition (not MVP-scoped here, but referenced by tickets as future work). ([Anthropic][3])

Open questions (explicit).

* What exact override/flag injection policy is desired for `tm`/`task-master`, `codex`, and `gemini` (beyond presence checks and artifact gates)?
* Where is the authoritative generator/template source for `ticket-action-pack.md` in the codebase (tickets suggest it is not provided in the current evidence set)?

[1]: https://developers.openai.com/codex/noninteractive/ "Non-interactive mode"
[2]: https://www.gnu.org/s/bash/manual/html_node/Bash-Startup-Files.html "Bash Startup Files (Bash Reference Manual)"
[3]: https://www.anthropic.com/news/model-context-protocol "Introducing the Model Context Protocol \ Anthropic"
