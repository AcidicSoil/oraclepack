# Oraclepack Issues Overview.prd

1. Overview

Problem statement. Oraclepack is a TUI/CLI wrapper that executes “Oracle Packs” (multi-step bash workflows embedded in Markdown) and records state/report outputs for repeatability.  Today, operational runs can fail (or silently skip later validation) due to (a) output verification requiring section tokens that don’t match what pack prompts actually instruct the assistant to write, (b) verification being tightly coupled to execution, and (c) defaults being hard to control at scale (no `.env`/env-var defaults).  Additionally, generated packs can contain bash syntax hazards where flags (notably `-p/--prompt`) appear orphaned on a new line, causing immediate bash failures (e.g., `-p: command not found`).

Who has it. Teams using oraclepack for runbooks, audits, migrations, or LLM evaluation flows, especially those running in CI and/or using oracle “browser mode.”  ([GitHub][1])

Why current solutions fail.

* Current verifier behavior can be triggered by the presence of “Answer format” and then requires four tokens in a single output file, even when the step explicitly asks for “Direct answer only,” producing false failures.
* Disabling verification allows execution to proceed but loses the ability to “verify later” without re-running steps.
* No standardized env-var defaults exist for verification/retries, making it cumbersome to apply consistent behavior across environments.
* Pack generation can emit syntactically invalid bash (or structurally risky line breaks), causing deterministic failures before the underlying tool (oracle) is even invoked correctly.

Target users.

* Pack authors: write/maintain packs; need templates/docs that reliably satisfy verifier contracts.
* Pack runners (ops/dev): execute packs interactively or in CI; need predictable verification defaults and “verify later.”
* CI maintainers: need non-interactive validation and fast-fail checks.

Success metrics (measurable).

* Reduce “output verification failed … missing tokens” false-fail rate to near-zero for packs that follow documented templates (track via report/state aggregation in CI).
* `oraclepack verify-outputs` provides deterministic, non-executing validation and returns non-zero when expectations are not met.
* Env-var defaults reduce per-invocation flag duplication in CI (measured by removal of repeated CLI flags in pipelines).
* Eliminate “orphaned -p/--prompt line” bash failures in generated packs; validator catches regressions before execution.

Assumptions (explicit).

* Oracle CLI is a third-party dependency; oraclepack will not change oracle’s behavior, only adjust oraclepack’s expectations/configuration and improve pack authoring guidance.
* Existing expectation/token normalization logic should be reused where possible.

1. Capability Tree (Functional Decomposition)

Capability: Output expectation inference
Description: Derive required output tokens per step based on step text, declared constraints, and write-output patterns.

Feature (MVP): Contract-aware single-output expectations

* Description: If a step both includes “Answer format” and explicitly constrains output to “Direct answer only,” require only the “direct answer” token for the single write-output file; otherwise preserve the “all four tokens” behavior.
* Inputs: Step code text; detected `--write-output` paths; token normalization rules; contract detector rules.
* Outputs: Map[path]requiredTokens for the step (or nil if no validation applies).
* Behavior: Detect “direct answer only” constraint; branch expectation set accordingly; do not change multi-output (chunked) behavior.

Feature (MVP): Chunked output expectations by suffix

* Description: For multiple `--write-output` paths, assign required tokens based on documented suffixes (`-direct-answer`, `-risks-unknowns`, `-next-experiment`, `-missing-evidence`).
* Inputs: Extracted write-output paths; suffix rules; token normalization.
* Outputs: Map[path]requiredTokens.
* Behavior: For each path, pick token set by suffix; ignore paths not matching any supported suffix.

Capability: Output file validation + reporting
Description: Validate that each output file exists/readable and contains required tokens, then provide actionable failure reports.

Feature (MVP): Validate output files and produce structured failures

* Description: Validate content for each expected file; return missing token sets per file; propagate IO errors distinctly.
* Inputs: Path; requiredTokens; normalization rules.
* Outputs: ok boolean; missingTokens; error; plus aggregated per-step report.
* Behavior: Read file; normalize; check token containment; record missing tokens and/or IO errors.

Feature (MVP): Human-readable CLI failure output

* Description: Print failures as step + file + missing tokens; exit non-zero if any validation fails.
* Inputs: Aggregated validation results.
* Outputs: Console output; process exit code.
* Behavior: Stable formatting for CI parsing; summarize totals.

Capability: Verification workflows
Description: Allow verification both during execution and independently after execution.

Feature (MVP): `oraclepack verify-outputs <pack.md>`

* Description: Validate outputs for all steps without executing step commands.
* Inputs: Pack file path; parsed steps; expectation inference; file validation.
* Outputs: Exit code; printed report.
* Behavior: Parse pack; compute expectations per step; validate referenced outputs; never spawn bash/oracle.

Capability: Verification configuration and defaults
Description: Support repeatable configuration for verification and retries using env vars as defaults while preserving CLI precedence.

Feature (MVP): Env-var defaults for output verification and retries

* Description: Support `ORACLEPACK_OUTPUT_VERIFY` and `ORACLEPACK_OUTPUT_RETRIES` as defaults when flags are not provided.
* Inputs: Environment variables; CLI flags; default values.
* Outputs: Resolved runtime config for `run` and `verify-outputs`.
* Behavior: Precedence: flags override environment variables; env vars override built-in defaults.  ([GitHub][2])

Capability: Pack authoring guidance (docs/templates)
Description: Ensure pack templates and docs align with verification behavior and oracle execution modes.

Feature (MVP): Templates for single-file and chunked outputs

* Description: Provide two canonical patterns: (a) single output file including all required sections/tokens when “Answer format” is used, and (b) chunked outputs using the suffix scheme.
* Inputs: Documentation source files; example packs/templates.
* Outputs: Updated docs/templates.
* Behavior: Explicitly document literal-token requirements (including “If evidence is insufficient” vs “Missing evidence” mismatch risk).

Feature (Non-MVP but recommended): Browser-mode reliability guidance

* Description: Document oracle browser-mode fragility and recommend mitigations (timeouts and API-engine preference).
* Inputs: Oracle CLI docs; current pack guidance.
* Outputs: Updated guidance section.
* Behavior: Recommend adjusting `--browser-timeout` / `--browser-input-timeout`; note defaults and duration formats.  ([GitHub][1])

Capability: Bash syntax safety for generated packs
Description: Prevent and detect unsafe multi-line command emission (notably orphaned flags) and fail fast in validation.

Feature (MVP): Orphaned-flag structural rule enforcement

* Description: Detect lines where `-p/--prompt` (or other flags) appears as a standalone command line (no continuation), and fail validation.
* Inputs: Extracted bash fence content; rule set; exceptions list (explicit continuations).
* Outputs: Validation errors with line numbers and rule IDs.
* Behavior: Regex/lexer checks for “flag-only line”; disallow inline end-of-line comments on `oracle ...` lines when wrapping is used.

Feature (MVP): Bash syntax check integration

* Description: Run `bash -n` on extracted bash content during `oraclepack validate`.
* Inputs: Bash text; bash binary availability.
* Outputs: Pass/fail; diagnostics.
* Behavior: If bash not present, skip with explicit notice (consistent with current “missing tools are skipped” behavior).

Feature (Non-MVP): Shellcheck integration (optional)

* Description: Run `shellcheck` when available to catch common pitfalls.
* Inputs: Bash text; shellcheck binary availability.
* Outputs: Lint findings.
* Behavior: Non-blocking by default unless a “strict” flag is set (decision point).

1. Repository Structure + Module Definitions (Structural Decomposition)

Proposed structure (Go; aligned to existing internal paths referenced in code snippets).

cmd/oraclepack/

* run.go (existing): wire flags and execution.
* validate.go (existing): pack-level validation command.
* verify_outputs.go (new): implements `oraclepack verify-outputs`.

internal/pack/

* parse.go (existing assumed): parse Markdown bash fence into Pack/Step.
* output_check.go (existing): contains ExtractWriteOutputPaths, StepOutputExpectations, ValidateOutputFile today.
* output_expectations.go (new): single responsibility: expectation inference (contract-aware + suffix mapping).

  * Exports:

    * `InferOutputExpectations(step Step) (map[string][]string, bool /*enabled*/)`
    * `DetectOutputContract(step Step) OutputContract`
* output_validator.go (new): single responsibility: validate files given expectations; aggregate results.

  * Exports:

    * `ValidateOutputs(expectations map[string][]string) []OutputFailure`
* verify_report.go (new): single responsibility: formatting + summary for CLI output.

  * Exports:

    * `RenderFailures(failures []OutputFailure) string`
* bash_syntax_validator.go (new): single responsibility: structural checks on bash content (orphaned flag rules).

  * Exports:

    * `ValidateBashStructure(bash string) []SyntaxFinding`
* bash_tooling_checks.go (new): single responsibility: optional external tool checks (`bash -n`, `shellcheck`).

  * Exports:

    * `RunBashNoexec(bash string) []SyntaxFinding`
    * `RunShellcheck(bash string) []SyntaxFinding`

internal/config/

* defaults.go (new): single responsibility: define default values and env var names.

  * Exports:

    * `DefaultOutputVerify() bool`
    * `DefaultOutputRetries() int`
* resolve.go (new): single responsibility: resolve runtime config from flags/env/defaults.

  * Exports:

    * `ResolveRunConfig(flags RunFlags, env Env) RunConfig`
    * `ResolveVerifyConfig(flags VerifyFlags, env Env) VerifyConfig`

internal/cli/

* flags.go (new or existing assumed): single responsibility: define CLI flag structs consumed by config resolver.

  * Exports:

    * `type RunFlags struct { … }`
    * `type VerifyFlags struct { … }`

Data types (single-source-of-truth types file).
internal/types/

* pack.go (new if not existing): Pack, Step.
* verification.go (new): OutputContract, OutputFailure, SyntaxFinding, severity enums.

Notes on aligning with existing code:

* Current `StepOutputExpectations` includes a placeholder token for `-missing-evidence` paths (“missing file path pattern”), which should be replaced with the literal required token behavior described in the tickets.
* Overrides types exist in `internal/overrides/types.go` and should remain independent of verification logic.

1. Dependency Chain (layers, explicit “Depends on: […]”)

Foundation layer (Phase 0)

* internal/types: core structs/enums. Depends on: []
* internal/config/defaults: env var names + defaults. Depends on: [internal/types]
* internal/pack/parse (existing or formalized): parse Pack/Step from Markdown. Depends on: [internal/types]

Verification layer (Phase 1)

* internal/pack/output_expectations: expectation inference + contract detection. Depends on: [internal/types, internal/pack/parse]
* internal/pack/output_validator: file IO + token validation. Depends on: [internal/types, internal/pack/output_expectations]
* internal/pack/verify_report: render failures. Depends on: [internal/types]

Config resolution layer (Phase 1)

* internal/config/resolve: resolve precedence flags/env/defaults. Depends on: [internal/types, internal/config/defaults]
* internal/cli/flags: typed view of CLI flags. Depends on: [internal/types]

CLI commands layer (Phase 2)

* cmd/oraclepack/verify_outputs: command wiring; no execution of steps. Depends on: [internal/config/resolve, internal/pack/parse, internal/pack/output_expectations, internal/pack/output_validator, internal/pack/verify_report]
* cmd/oraclepack/run (existing integration): uses config resolver for defaults. Depends on: [internal/config/resolve, internal/cli/flags, existing run modules]

Bash safety layer (Phase 2)

* internal/pack/bash_syntax_validator: orphaned-flag checks. Depends on: [internal/types]
* internal/pack/bash_tooling_checks: bash -n / shellcheck integration. Depends on: [internal/types]
* cmd/oraclepack/validate (integration changes): ensure structural checks are run during validate and (optionally) enforced before run paths. Depends on: [internal/pack/parse, internal/pack/bash_syntax_validator, internal/pack/bash_tooling_checks]

1. Development Phases (Phase 0…N; entry/exit criteria; tasks with dependencies + acceptance criteria + test strategy)

Phase 0: Foundations
Entry criteria: Current build passes; baseline parsing and existing verification functions compile.

Tasks:

* internal/types: define Pack/Step minimal structs (or adapt existing types into a single exported location).

  * Depends on: none
  * Acceptance criteria: All downstream modules compile importing these types.
  * Test strategy: Unit compile checks via `go test ./...`.
* internal/config/defaults: define env var names and defaults.

  * Depends on: internal/types
  * Acceptance criteria: Defaults accessible from resolver; no behavior changes yet.
  * Test strategy: Unit tests for default values and parsing helpers.

Exit criteria: Foundation modules compile and can be imported without cycles.

Phase 1: Output verification correctness + configuration defaults (MVP core)
Entry criteria: Phase 0 complete.

Tasks (parallelizable where possible):

* Contract-aware expectations for single-output steps.

  * Depends on: internal/pack/parse, internal/types
  * Acceptance criteria: A step containing “Answer format” plus “Return only: Direct answer” yields requiredTokens = [“direct answer”] for the single output; non-constrained steps preserve “all four tokens” requirement.
  * Test strategy: Unit tests over step text fixtures; table-driven tests for contract detection.
* Fix chunked suffix expectation mapping, including “missing evidence” token literal handling.

  * Depends on: internal/types
  * Acceptance criteria: `-direct-answer`, `-risks-unknowns`, `-next-experiment`, `-missing-evidence` map to the documented tokens; no placeholder token remains.
  * Test strategy: Unit tests for suffix-to-token mapping.
* Env-var default resolver for output verification and retries with flag precedence.

  * Depends on: internal/config/defaults, internal/types
  * Acceptance criteria: `ORACLEPACK_OUTPUT_VERIFY` and `ORACLEPACK_OUTPUT_RETRIES` set defaults only when CLI flags are absent; flags override env vars.  ([GitHub][2])
  * Test strategy: Unit tests over precedence matrix (flags present/absent; env present/absent).

Exit criteria: Running existing `oraclepack run` with no explicit flags can be controlled via env vars without breaking existing CLI behavior; expectation inference matches the contract rules.

Phase 2: `verify-outputs` workflow (MVP deliverable)
Entry criteria: Phase 1 complete.

Tasks:

* Implement `oraclepack verify-outputs <pack.md>` command.

  * Depends on: output_expectations, output_validator, verify_report, config/resolve, pack/parse
  * Acceptance criteria: Command does not execute step commands; validates referenced `--write-output` files; prints failures with step + file + missing token set; exits non-zero on any failure.
  * Test strategy: Integration tests using fixture packs and fixture output files; assert exit code and key output lines.

Exit criteria: CI can run `oraclepack verify-outputs` on a pack directory and get deterministic pass/fail without execution.

Phase 3: Pack authoring guidance updates (docs/templates)
Entry criteria: Phase 1 complete (so docs match behavior).

Tasks:

* Update templates/examples for single-file and chunked-output patterns and token literal warnings.

  * Depends on: none (docs), but should be consistent with Phase 1 behavior
  * Acceptance criteria: Docs explicitly state the “Answer format” trigger behavior, required tokens for single-file outputs, suffix scheme for chunked outputs, and the literal token mismatch warning.
  * Test strategy: Doc review + optional doc lint; include example packs used by tests.

* Add oracle browser-mode mitigation guidance (timeouts, API preference).

  * Depends on: none
  * Acceptance criteria: Docs reference oracle’s `--browser-timeout` / `--browser-input-timeout` knobs and duration formats. ([GitHub][1])
  * Test strategy: None (docs), but validate linked examples in a smoke run.

Exit criteria: A pack author can follow docs to produce outputs that pass verification reliably.

Phase 4: Bash syntax safety (MVP for bash failure class)
Entry criteria: Phase 0 complete.

Tasks:

* Add orphaned-flag structural validator and integrate into `oraclepack validate`.

  * Depends on: internal/pack/parse, internal/pack/bash_syntax_validator
  * Acceptance criteria: Packs that contain standalone `-p` lines (or similar orphaned flags) fail validation with clear diagnostics and line numbers.
  * Test strategy: Unit tests on extracted bash fixtures; integration test calling validate on a known-bad pack.
* Add `bash -n` check (when available) to validation.

  * Depends on: internal/pack/bash_tooling_checks
  * Acceptance criteria: Syntax-invalid bash is caught before run; tool absence is reported as skipped (consistent behavior).
  * Test strategy: Integration test gated by bash availability in CI.

Exit criteria: The cited `-p: command not found` failure mode is prevented by validation, and regressions are caught pre-execution.

Phase 5: Enforcement and CI guardrails (recommended)
Entry criteria: Phase 4 complete.

Tasks:

* Make `run` paths call `validate` (or embed equivalent checks) before executing steps.

  * Depends on: validate integration tasks
  * Acceptance criteria: Interactive/TUI run/rerun cannot bypass the new structural checks.
  * Test strategy: Integration test ensuring run refuses invalid packs.

* Add CI/pre-commit guidance to run `oraclepack validate` on modified packs.

  * Depends on: none (docs/CI config)
  * Acceptance criteria: CI fails fast on invalid generated packs.

1. User Experience

Personas.

* Pack author: wants “copy/paste safe” templates that satisfy output verification.
* Runner (interactive): wants errors that explain exactly what’s missing and how to fix.
* Runner (CI): wants deterministic, machine-friendly exit codes and stable text output.

Key flows.

* Verify later flow: user runs `oraclepack run … --output-verify=false` to unblock execution, then runs `oraclepack verify-outputs <pack.md>` after outputs are generated to validate without re-running steps.
* Environment defaulting: CI sets `ORACLEPACK_OUTPUT_VERIFY=0` or `1` and `ORACLEPACK_OUTPUT_RETRIES=N` and omits repetitive flags; a one-off local run can override via flags.
* Authoring guidance: docs provide two patterns (single-file with all required tokens vs chunked outputs with suffixes) and call out literal-token pitfalls.
* Validation fast-fail: `oraclepack validate` reports orphaned flag lines and bash syntax issues before a run starts.

UI/UX notes.

* Failure output should always include step ID (e.g., “01”), file path, and missing tokens.
* For IO errors, distinguish “file missing/unreadable” from “missing tokens” to guide remediation.

1. Technical Architecture

System components.

* Pack parser: extracts bash fence, splits steps, provides Step.Code used for expectation inference and bash validation.
* Expectation inference: determines required tokens per output file (contract-aware).
* Output validator: reads files and checks token inclusion using normalization.
* CLI config resolver: merges flags/env/defaults with explicit precedence.
* Docs/templates: non-code artifacts that define the contract pack authors follow.

Data models.

* OutputContract: enum (AllSections, DirectAnswerOnly, ChunkedBySuffix, None).
* OutputExpectation: {StepID, Path, RequiredTokens, Contract}.
* OutputFailure: {StepID, Path, MissingTokens, IOError}.
* SyntaxFinding: {RuleID, Line, Message, Severity}.

Integrations and constraints.

* Oracle CLI “browser mode” uses timeouts that can be tuned via flags like `--browser-timeout` and `--browser-input-timeout` (durations support h/m/s/ms). ([GitHub][1])
* Config precedence should match established Go CLI patterns (flags override env vars). ([GitHub][2])
* Windows Git Bash path conversion can be controlled with `MSYS_NO_PATHCONV=1`; document this as an execution environment caveat where relevant. ([CloudBees Documentation][3])

Key decisions.

* Keep expectation inference deterministic and derived from step text + write-output paths (no timing-based “wait longer” fixes for content/contract mismatch).
* Add “verify later” as a first-class workflow via a separate command rather than overloading execution paths.

1. Test Strategy

Test pyramid targets.

* Unit: ~70% (expectation inference, contract detection, suffix mapping, env precedence, orphaned-flag detector).
* Integration: ~25% (parse pack → expectations → validate fixture files; validate command on fixture packs; verify-outputs command end-to-end without execution).
* E2E: ~5% (optional; run a minimal pack in CI with controlled outputs, only if stable).

Coverage minimums.

* Unit modules: ≥85% line coverage for expectation inference and config resolver.
* Critical paths: contract-aware inference and verify-outputs command must have branch coverage for the key conditionals.

Critical scenarios (by module).

* output_expectations:

  * Happy path: “Answer format” + no “direct answer only” → four tokens required.
  * Happy path: “Answer format” + “Return only: Direct answer” → only “direct answer” required.
  * Chunked outputs: suffix mapping produces correct tokens, including missing-evidence literal handling.
* output_validator:

  * Missing file → IO error surfaced with non-zero exit code.
  * Missing tokens → missing list is stable and complete.
* config/resolve:

  * Flags override env vars; env vars apply only when flags absent.
* bash_syntax_validator:

  * Orphaned `-p` line triggers finding with correct line mapping.

Integration points.

* verify-outputs command: parse pack, infer expectations, validate output files, render summary; ensure no step execution occurs (mock/guard against spawning bash/oracle).

1. Risks and Mitigations

Risk: Ambiguity in contract detection phrases (“Return only: Direct answer” variants).

* Impact: Medium; false passes/fails.
* Likelihood: Medium.
* Mitigation: Use conservative detection with a small, documented set of accepted phrases; document exact patterns; add tests for known variants; allow future extension via explicit marker token in packs.
* Fallback: Provide an explicit per-step override comment marker in pack syntax (future enhancement).

Risk: Existing users rely on current behavior (four-token requirement) and would be surprised by relaxed verification.

* Impact: Medium.
* Likelihood: Low–Medium.
* Mitigation: Only relax when an explicit “direct answer only” constraint is detected; preserve behavior otherwise.
* Fallback: Add a strict mode flag to force legacy behavior.

Risk: External tool variability (oracle browser mode produces partial outputs).

* Impact: Medium–High.
* Likelihood: Medium.
* Mitigation: Document mitigations (timeouts, API preference). ([GitHub][1])
* Fallback: Encourage chunked outputs to reduce truncation risk; rely on verify-outputs reruns post-hoc.

Risk: Shellcheck/bashtool availability differs across environments.

* Impact: Low–Medium.
* Likelihood: High.
* Mitigation: Treat as optional; keep `bash -n` as primary; clear “skipped” output when tools absent.
* Fallback: Provide a `--strict` mode for environments that guarantee tooling.

Risk: Validation added to run paths increases friction for advanced users.

* Impact: Medium.
* Likelihood: Medium.
* Mitigation: Make the new checks fast; provide `--no-validate` escape hatch only if absolutely necessary (not MVP).
* Fallback: Keep enforcement limited to CI via recommended pipeline steps.

1. Appendix

Primary internal references (uploaded).

* Output verification epic and split tickets (T1–T4, constraints, acceptance criteria).
* Bash orphaned `-p/--prompt` line failure details and proposed validation/generation guardrails.
* Existing output verification code paths and current suffix/token handling snippet (notably missing-evidence placeholder).
* Oraclepack product scope and operating model (packs as Markdown bash workflows; state/report files).

External references (web).

* Oracle CLI browser-mode flags and timeout behavior (`--browser-timeout`, `--browser-input-timeout`, duration format). ([GitHub][1])
* Viper precedence model supporting “flags > env vars > config > defaults” patterns (useful for env-var defaulting design). ([GitHub][2])
* MSYS/Git Bash path conversion and `MSYS_NO_PATHCONV=1` workaround documentation. ([CloudBees Documentation][3])

Open questions (to resolve during implementation; not blockers for Phase 0).

* Exact current CLI framework in oraclepack (cobra/viper vs custom) to implement env-var defaults cleanly.
* Exact locations of pack templates/docs to update for the “Gold Stage packs” guidance.

[1]: https://github.com/steipete/oracle/blob/main/docs/browser-mode.md?utm_source=chatgpt.com "oracle/docs/browser-mode.md at main"
[2]: https://github.com/spf13/viper?utm_source=chatgpt.com "spf13/viper: Go configuration with fangs"
[3]: https://docs.cloudbees.com/docs/cloudbees-cd-kb/latest/cloudbees-cd-kb/kbec-00420-stopping-path-conversion-for-git-bash?utm_source=chatgpt.com "KBEC-00420 - Stopping Path Conversion for Git Bash"
