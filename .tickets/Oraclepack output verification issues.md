Parent Ticket:

* Title: oraclepack output verification failures and verification configurability gaps
* Summary: Runs fail with `output verification failed for step 01` due to oraclepack’s post-step validator requiring section tokens (triggered by “Answer format”) that are not present in the `--write-output` file, plus cases where the output file is not created/readable on retry. The ticket also includes requests to make verification easier to disable/configure (including via `.env`) and to support “verify later” workflows.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “output verification failed for step 01 … verifier is expecting more sections than you are actually emitting.”
* Global Constraints:

  * Do not add “wait time” as a validator fix when the failure is due to content/contract mismatch.
* Global Environment:

  * Unknown
* Global Evidence:

  * Error context: “output verification failed for step 01: docs/…” and validator behavior around “Answer format” token requirements.
  * Noted token set required when “Answer format” is present: “direct answer”, “risks unknowns”, “next smallest concrete experiment”, “if evidence is insufficient”.
  * Suggested multi-file suffix scheme: `-direct-answer`, `-risks-unknowns`, `-next-experiment`, `-missing-evidence`.

Split Plan:

* Coverage Map:

  * “oraclepack’s ‘output verification’ path … expects the file to exist and … contain certain section tokens” → T1
  * “single `--write-output` + phrase ‘Answer format’ causes verifier to require all four section tokens” → T1
  * “If your step instructs ‘Return only: Direct answer’ … verification fails” → T1
  * “On the re-run, the output file isn’t being written at all” → T4
  * “Don’t add a longer wait time … step command already blocks until CLI exits” → Info-only
  * “If oracle browser run is incomplete/truncated … increase oracle `--browser-timeout` / `--browser-input-timeout` or switch to API” → T4
  * “How to make the validator optional … `--output-verify=false` and `--output-retries`” → T4
  * “Align chunking with verifier … split outputs into multiple `--write-output` files using suffixes” → T4
  * “Edge: ‘Missing evidence’ heading may fail if required token is literally ‘If evidence is insufficient’” → T4
  * “Add ‘verify later’ workflow: new `oraclepack verify-outputs <pack.md>` command” → T3
  * “No built-in env var toggle for output verification; toggle is CLI flag today” → T2
  * “Option: wrapper script to inject `--output-verify=false`” → T2
  * “Option: add explicit env var support like `ORACLEPACK_OUTPUT_VERIFY` / `ORACLEPACK_OUTPUT_RETRIES`” → T2
  * “Which skill produced the pack … matches oraclepack-gold-pack (Gold Stage 1)” → Info-only
* Dependencies:

  * Not provided
* Split Tickets:

```ticket T1
T1 Title: Make output verification contract-aware for “Direct answer only” single-output steps
Type: enhancement
Target Area: oraclepack run-step output verification (single `--write-output` expectations)
Summary:
  The validator currently requires four section tokens whenever the step text includes “Answer format”, which fails when the produced output contains only “Direct answer” (e.g., when the prompt says “Return only: Direct answer”). Adjust validation so a single-output step can be treated as “direct answer only” when the step explicitly declares that constraint, avoiding false failures.
In Scope:
  - Update expectation logic so “Answer format” does not always imply all four tokens for single-output steps.
  - Support a “Direct answer only” contract when the step text explicitly indicates that output constraint (e.g., “Return only: Direct answer”).
Out of Scope:
  - Not provided
Current Behavior (Actual):
  - If a step contains “Answer format” and has exactly one `--write-output`, the output file is required to contain all of:
    - direct answer
    - risks unknowns
    - next smallest concrete experiment
    - if evidence is insufficient
  - Steps that output only “Direct answer …” fail verification.
Expected Behavior:
  - If a step explicitly constrains output to “Direct answer only”, the validator should not require the other three section tokens for that step’s single `--write-output` file.
  - Steps that truly require all four sections should continue to be validated against all four tokens.
Reproduction Steps:
  1) Run a step with a prompt that includes “Answer format:” and also includes “Return only: Direct answer”.
  2) Produce an output file containing only the “Direct answer” section.
  3) Observe `output verification failed for step 01` due to missing tokens.
Requirements / Constraints:
  - Must preserve existing “all four tokens required” behavior when the step is not explicitly “Direct answer only”.
Evidence:
  - Error symptom: `output verification failed for step 01: docs/...`
  - Required tokens set when “Answer format” is present: “direct answer”, “risks unknowns”, “next smallest concrete experiment”, “if evidence is insufficient”.
Open Items / Unknowns:
  - Exact mechanism used today to detect “Answer format” and to compute required tokens (unknown in provided text).
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - A single-output step that includes “Answer format” but explicitly states “Return only: Direct answer” passes verification when the file contains “Direct answer” and omits the other sections.
  - A single-output step that includes “Answer format” without “Direct answer only” constraint continues to fail verification if any of the other required tokens are missing.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “single `--write-output ...` + a prompt that includes the phrase ‘Answer format’ causes the verifier to require all four section tokens”
  - “If your step instructs ‘Return only: Direct answer’ … verification fails”
  - “change … verification logic to accept a ‘direct answer only’ contract”
```

```ticket T2
T2 Title: Add `.env` / environment-variable control for output verification defaults
Type: enhancement
Target Area: oraclepack CLI configuration (env var support for `--output-verify` / `--output-retries`)
Summary:
  Output verification can be toggled via CLI flags, but there is no built-in environment variable that can be set in `.env` to control verification on/off. Add explicit env-var support (as suggested) to allow `.env` driven defaults, while keeping CLI flags available.
In Scope:
  - Implement an environment variable toggle for output verification (suggested name: `ORACLEPACK_OUTPUT_VERIFY`).
  - Implement an environment variable for output retries (suggested name: `ORACLEPACK_OUTPUT_RETRIES`).
  - Define precedence so the env vars provide defaults without removing CLI flag control (exact precedence not specified in provided text; implement per existing CLI config patterns).
Out of Scope:
  - Not provided
Current Behavior (Actual):
  - Output verification is toggled by CLI flag (`--output-verify`), not an env var.
  - No built-in `.env` variable exists to turn verification on/off.
Expected Behavior:
  - Users can set a `.env` environment variable to control output verification behavior without modifying command lines everywhere.
Reproduction Steps:
  1) Attempt to set an env var in `.env` to disable output verification.
  2) Observe there is “no built-in environment variable in the Go `oraclepack` CLI to toggle output verification.”
Requirements / Constraints:
  - Must not remove existing CLI flag support for `--output-verify` / `--output-retries`.
Evidence:
  - “There is no built-in environment variable … The toggle is currently a CLI flag (`--output-verify`), not an env var.”
  - Suggested addition: “Implement something like `ORACLEPACK_OUTPUT_VERIFY=0/1` (and optionally `ORACLEPACK_OUTPUT_RETRIES`).”
Open Items / Unknowns:
  - Current CLI config/env binding conventions in the Go CLI (unknown in provided text).
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - Setting `ORACLEPACK_OUTPUT_VERIFY=0` results in verification being disabled by default for `oraclepack run` when the user does not pass an explicit `--output-verify` flag.
  - Setting `ORACLEPACK_OUTPUT_VERIFY=1` results in verification being enabled by default under the same conditions.
  - Setting `ORACLEPACK_OUTPUT_RETRIES=<N>` results in the default retries being `<N>` when the user does not pass an explicit `--output-retries` flag.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “what environment variable can we set in our .env to turn the verify output on/off?”
  - “There is no built-in environment variable … The toggle is currently a CLI flag (`--output-verify`), not an env var.”
  - “Implement something like `ORACLEPACK_OUTPUT_VERIFY=0/1` (and optionally `ORACLEPACK_OUTPUT_RETRIES`).”
```

```ticket T3
T3 Title: Add a standalone `verify-outputs` command to validate pack outputs without executing steps
Type: enhancement
Target Area: oraclepack CLI commands (output verification workflow)
Summary:
  Disabling output verification allows runs to proceed, but then verification is tied to execution and won’t be re-applied later (e.g., with `--resume`). Add a command that only validates outputs using the same expectation logic, enabling “run now, verify later” workflows.
In Scope:
  - Add a command: `oraclepack verify-outputs <pack.md>`.
  - For each step, compute output expectations and validate the referenced `--write-output` files against required tokens.
  - Exit non-zero when validation fails and print failures (step + file + missing token set).
Out of Scope:
  - Not provided
Current Behavior (Actual):
  - Verification is performed as part of step execution; disabling verification allows success without later validation.
Expected Behavior:
  - `verify-outputs` validates outputs for a pack without executing any step commands.
Reproduction Steps:
  1) Run `oraclepack run <pack.md> --output-verify=false`.
  2) Attempt to validate outputs later without re-running steps.
  3) Observe verification is tied to execution in the current workflow.
Requirements / Constraints:
  - Must reuse the existing expectation/token normalization logic already used during execution (as described).
Evidence:
  - Suggested command: “`oraclepack verify-outputs <pack.md>` … compute `pack.StepOutputExpectations(step)` and run `pack.ValidateOutputFile(path, requiredTokens)`.”
Open Items / Unknowns:
  - Exact function names/locations of expectation and validation logic in the codebase (unknown in provided text).
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - Running `oraclepack verify-outputs <pack.md>` does not execute step commands.
  - The command validates each `--write-output` file against the expected token set and reports missing tokens.
  - The command exits with a non-zero status when any step output fails validation.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “If you want ‘run now, verify later’ … add a new command that only validates outputs without executing steps.”
  - “`oraclepack verify-outputs <pack.md>` … reuses the same expectation logic and token normalization.”
```

```ticket T4
T4 Title: Update pack/templates and documentation to avoid validator mismatches and improve reliability in browser mode
Type: docs
Target Area: pack authoring guidance (Gold Stage packs), output chunking conventions, oracle CLI invocation guidance
Summary:
  The failures are driven by mismatch between what the validator requires (triggered by “Answer format”) and what is actually written to `--write-output`, plus fragility in oracle browser automation runs. Update templates and documentation to (a) ensure outputs include required section tokens when “Answer format” is used, or (b) use multi-file chunking with the documented filename suffixes, and (c) document browser-timeout mitigations and path preconditions for `--write-output`.
In Scope:
  - Document that when “Answer format” is present, the validator requires all four tokens in a single-output file unless using chunked outputs.
  - Provide template guidance for two supported patterns:
    - Single file per step: output must include headings/phrases that satisfy all required tokens (“Direct answer”, “Risks/unknowns”, “Next smallest concrete experiment”, “If evidence is insufficient”).
    - Chunked outputs: use multiple `--write-output` files with suffixes:
      - `-direct-answer`
      - `-risks-unknowns`
      - `-next-experiment`
      - `-missing-evidence`
  - Document the “Missing evidence” vs “If evidence is insufficient” token mismatch risk (token is literal).
  - Document write-output path preconditions and caveats:
    - Ensure output directory exists (e.g., `mkdir -p docs`).
    - Prefer absolute paths if CWD issues arise.
    - Oracle “Saved assistant output to …” does not guarantee `--write-output` was created.
  - Document oracle browser-mode fragility and mitigations:
    - Increase `--browser-timeout` / `--browser-input-timeout`.
    - Prefer API engine where possible.
Out of Scope:
  - Changing oracle CLI implementation (third-party).
  - Adding new validator logic beyond the contract-awareness changes covered in T1.
Current Behavior (Actual):
  - Validator requires four tokens when “Answer format” is present; single-file outputs that omit any section fail verification.
  - Browser automation mode is described as experimental/fragile; outputs may be partial/truncated.
Expected Behavior:
  - Pack templates and docs lead authors to produce outputs that satisfy validation (either all-in-one or chunked-by-suffix) and reduce intermittent browser-mode failures via recommended settings.
Reproduction Steps:
  1) Use a pack step with “Answer format” and a single `--write-output` file.
  2) Produce output missing one or more required section tokens.
  3) Observe verification failure.
  4) Optionally, observe incomplete outputs in browser automation runs.
Requirements / Constraints:
  - Must preserve the stated suffix names and required token phrases as documented.
Evidence:
  - Required token set: “direct answer”, “risks unknowns”, “next smallest concrete experiment”, “if evidence is insufficient”.
  - Multi-file suffix validation scheme: `-direct-answer`, `-risks-unknowns`, `-next-experiment`, `-missing-evidence`.
  - Browser-mode mitigation knobs: `--browser-timeout`, `--browser-input-timeout`; suggestion to switch to API.
Open Items / Unknowns:
  - Exact pack/template files that should be updated (unknown in provided text).
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - Documentation explicitly states the validator trigger (“Answer format”) and the required token behavior for single-file outputs.
  - Templates/examples demonstrate both:
    - A single-file step that includes all required section headings/phrases.
    - A chunked-output step that writes to the four suffixed files and requests only the matching section per file.
  - Documentation includes the literal-token warning about “If evidence is insufficient” vs “Missing evidence”.
  - Documentation includes guidance on output directory existence and the oracle log caveat.
  - Documentation includes browser-timeout mitigation guidance and notes browser automation fragility.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “remove the ‘Answer format’ block … or disable oraclepack’s output verification”
  - “make the output actually include all required sections … or split into multiple outputs … suffixes”
  - “Increase oracle browser timeouts … switch to API engine”
```
