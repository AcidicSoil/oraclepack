Parent Ticket:

* Title: Oraclepack CLI/MCP parity with grouped-pack skills and reliable artifact generation
* Summary: Tighten `oraclepack` validation and execution context so generated packs run without common failures (pack shape drift, workdir/path drift, oracle CLI flag drift). Add MCP passthrough for key run flags, add preflight checks, align pack templates with `out_dir`, and add a first-class `oraclepack generate` surface (wrapper first, native generator later) matching grouped-pack skill outputs.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “get our oraclepack CLI/MCP up to date so artifacts … run through oraclepack without any errors”
* Global Constraints:

  * Packs must have exactly one ` ```bash … ``` ` fence and no other code fences.
  * Packs must have exactly 20 steps (01..20) with correct step header format.
  * Generation should be deterministic (e.g., lexicographic discovery) per described grouping rules.
* Global Environment:

  * Unknown
* Global Evidence:

  * References to: `internal/pack/parser.go` / `pack.Validate()`, `app.Config.WorkDir`, `ExtractOracleInvocations`, `ValidateOverrides`, MCP tool `oraclepack_run_pack`.
  * Commands mentioned: `oraclepack validate`, `oraclepack list`, `oraclepack run --no-tui --yes --run-all`, `oracle … --dry-run summary`, `python3 scripts/validate_pack.py`.

Split Plan:

* Coverage Map:

  * “Make `oraclepack validate` enforce … ‘strict pack shape’ …” → T1
  * “In `internal/pack/parser.go` (or `pack.Validate()`), count … fences …” → T1
  * “there isn’t exactly one ` ```bash … ``` ` block …” → T1
  * “any other code fences exist …” → T1
  * “Fix workdir determinism … missing `-f` files …” → T2
  * “Add a persistent `--work-dir` flag … `app.Config.WorkDir`.” → T2
  * “Default behavior when `--work-dir` is omitted:” → T2
  * “auto-detect repo root from `packPath` … set WorkDir …” → T2
  * “In the MCP server, extend `oraclepack_run_pack` to accept …” → T3
  * “Expose `--oracle-bin` (and optionally `--out-dir`) through MCP” → T3
  * “Add MCP tool params: … `oracle_bin` … `out_dir` … `work_dir` …” → T3
  * “Append `--oracle-bin …` / `--out-dir …` / `--work-dir …` …” → T3
  * “Add preflight checks that match how the packs fail …” → T4
  * “Missing attachment file preflight … verify every `-f <path>` exists …” → T4
  * “Oracle flag drift preflight … `ValidateOverrides` …” → T5
  * “Wire this into `oraclepack run` … `--preflight-oracle` …” → T5
  * “Extract invocations (already implemented).” → T5
  * “Inject flags safely (already implemented).” → T5
  * “Execute `--dry-run summary` and block …” → T5
  * “Align pack templates with … `out_dir` model …” → T6
  * “In the pack prelude … include `out_dir="<resolved default>"` …” → T6
  * “Use `--write-output "$out_dir/…"` and `mkdir -p "$out_dir"`.” → T6
  * “Minimum validation pipeline to prevent regressions …” → T10
  * “Generation-time … `python3 scripts/validate_pack.py …`” → T10
  * “Oraclepack-level: `oraclepack validate …`” → T10
  * “Parse sanity: `oraclepack list …` prints 20 …” → T10
  * “Execution smoke … stub `oracle` via `--oracle-bin` …” → T10
  * “Add a new CLI command …” → T7
  * “`oraclepack generate tickets-grouped …`” → T7
  * “`oraclepack generate codebase-grouped …`” → T7
  * “Implementation: … shell out … generate_grouped_packs.py” → T7
  * “After generation, call `oraclepack validate` …” → T7
  * “Embed the exact templates … into the `oraclepack` binary …” → T8
  * “Port the deterministic grouping rules … `internal/generate/*` …” → T8
  * “discover files deterministically …” → T8
  * “group by subdir + infer loose files …” → T8
  * “split oversized groups into ‘part N’ packs” → T8
  * “Render packs … fail on unresolved placeholders …” → T8
  * “Validate with the existing pack parser/validator constraints:” → T8
  * “parser requires a ` ```bash ... ``` ` code block … ROI …” → T8
  * “generator should always emit exactly 20 steps … one bash fence …” → T8
  * “CLI shape: add `oraclepack generate` via Cobra … flags …” → T7
  * “MCP: expose generation as tools …” → T9
  * “`oraclepack_generate_tickets_grouped(…) -> {packs…, manifest…}`” → T9
  * “`oraclepack_generate_codebase_grouped(…) -> {packs…, manifest…}`” → T9
  * “What the generator must output … `out_dir/packs/*.md` … `_groups.json` …” → T7
  * “Each pack must remain runner-ingestible …” → T7
  * “Minimal acceptance test (proves parity) … generate … validate/list … run …” → T10
  * Standalone “...” truncation markers in the source text → Info-only
* Dependencies:

  * T3 depends on T2 because MCP `work_dir` passthrough requires a working CLI `--work-dir` behavior.
  * T4 depends on T2 because attachment existence is checked relative to the resolved WorkDir.
  * T7 depends on T1 because generation “fail fast” relies on strict `oraclepack validate` shape enforcement.
  * T9 depends on T7 because MCP generation can initially delegate to the CLI `oraclepack generate` surface.
  * T10 depends on T7 because the parity acceptance flow starts from `oraclepack generate …`.
* Split Tickets:

````ticket T1
T# Title: Enforce strict pack-shape validation in `oraclepack validate`
Type: chore
Target Area: `oraclepack validate`; `internal/pack/parser.go` and/or `pack.Validate()`
Summary:
  Tighten pack validation to match the grouped-pack skill contract by enforcing code-fence invariants. This prevents “shape drift” (extra fences, wrong fence language) from reaching execution. The goal is early, actionable failure during `oraclepack validate`.
In Scope:
  - Count all triple-backtick fences during validation.
  - Fail validation if there is not exactly one ` ```bash … ``` ` block.
  - Fail validation if any other code fences exist.
Out of Scope:
  - Not provided
Current Behavior (Actual):
  - Parser extracts a bash block via regex and parses steps via header regex, but does not inherently guarantee “exactly one code fence and no other fences.”
Expected Behavior:
  - `oraclepack validate` rejects packs that contain multiple fences, non-bash fences, or missing the single required `bash` fence.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Enforce “exactly one `bash` fence and no other fences.”
Evidence:
  - References: “pack parser extracts the bash block via a regex … doesn’t inherently guarantee ‘exactly one code fence and no other fences.’”
Open Items / Unknowns:
  - Exact location of current validation entrypoint (whether `internal/pack/parser.go` or `pack.Validate()` is canonical) not provided.
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - `oraclepack validate` fails when a pack has 0 ` ```bash ` fences.
  - `oraclepack validate` fails when a pack has 2+ code fences.
  - `oraclepack validate` fails when a non-bash code fence exists alongside the bash fence.
  - `oraclepack validate` passes when exactly one ` ```bash … ``` ` fence exists and no other code fences exist.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Make `oraclepack validate` enforce … ‘strict pack shape’…”
  - “count all triple-backtick fences and fail if …”
  - “there isn’t exactly one ` ```bash … ``` ` block …”
````

```ticket T2
T# Title: Make `oraclepack run` workdir deterministic for relative `-f` paths
Type: bug
Target Area: `oraclepack run`; CLI config `app.Config.WorkDir`; repo-root detection from `packPath`
Summary:
  Fix failures caused by running packs from the wrong working directory, which breaks relative `-f` attachments. Add a persistent `--work-dir` flag and a default workdir inference strategy when the flag is omitted. This makes pack execution consistent regardless of where the pack file lives.
In Scope:
  - Add a persistent CLI `--work-dir` flag.
  - Plumb `--work-dir` into `app.Config.WorkDir`.
  - Default behavior when omitted: auto-detect repo root from `packPath` and set WorkDir to that root (not the pack’s directory).
Out of Scope:
  - MCP server tool signature changes (handled separately).
Current Behavior (Actual):
  - `oraclepack run` hardcodes `WorkDir: "."`, causing missing `-f` files when packs reference repo-root-relative files from other working directories.
Expected Behavior:
  - Running a pack resolves relative paths consistently under the inferred repo root, unless `--work-dir` is explicitly provided.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Default must “auto-detect repo root from `packPath` … and set WorkDir to that root.”
Evidence:
  - “Fix workdir determinism … missing `-f` files”
  - “`oraclepack run` currently hardcodes `WorkDir: "."`…”
Open Items / Unknowns:
  - Exact repo-root detection signals (e.g., `.git/`, `go.mod`, etc.) are referenced but truncated in the source text.
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - CLI exposes `--work-dir` and it sets the runner workdir.
  - When `--work-dir` is omitted, WorkDir is inferred from `packPath` and is not the pack’s directory.
  - Packs referencing repo-root-relative `-f` attachments do not fail solely due to execution from a different current directory.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Fix workdir determinism …”
  - “Add a persistent `--work-dir` flag … `app.Config.WorkDir`.”
  - “auto-detect repo root from `packPath` … set WorkDir …”
```

```ticket T3
T# Title: Expose run flags through MCP: `work_dir`, `oracle_bin`, `out_dir`
Type: enhancement
Target Area: MCP server (FastMCP); tool `oraclepack_run_pack`; CLI invocation wiring
Summary:
  The MCP wrapper must pass through key execution-context flags so packs run reliably in MCP environments (where `oracle` may not be on PATH and workdir may differ). Extend the MCP tool signature to accept `work_dir`, `oracle_bin`, and `out_dir`, and forward them to the CLI invocation.
In Scope:
  - Extend MCP `oraclepack_run_pack` tool to accept `work_dir`.
  - Add MCP tool params: `oracle_bin: Optional[str]`, `out_dir: Optional[str]`, `work_dir: Optional[str]`.
  - Append `--oracle-bin <path>` / `--out-dir <dir>` / `--work-dir <dir>` when calling the CLI.
Out of Scope:
  - Implementing CLI semantics for `--work-dir` (handled in T2).
Current Behavior (Actual):
  - MCP wrapper only passes `run … --no-tui [--yes] [--run-all]` and does not expose `--oracle-bin` / `--out-dir` / `--work-dir`, causing failures when `oracle` is not on PATH or workdir differs.
Expected Behavior:
  - MCP can run packs with explicit oracle binary, output directory, and working directory controls.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - MCP signature changes must follow the FastMCP tool-parameter pattern (per source text).
Evidence:
  - “The MCP wrapper doesn’t expose it today; if `oracle` isn’t on PATH … runs will fail …”
  - “extend `oraclepack_run_pack` to accept `work_dir` …”
Open Items / Unknowns:
  - Exact MCP server file/module containing `oraclepack_run_pack` definition not provided.
Risks / Dependencies:
  - Depends on T2 for `--work-dir` behavior to exist in CLI.
Acceptance Criteria:
  - MCP tool `oraclepack_run_pack` accepts `work_dir`, `oracle_bin`, and `out_dir` parameters.
  - MCP invocation forwards those parameters to the CLI as `--work-dir`, `--oracle-bin`, and `--out-dir`.
  - Running via MCP does not require `oracle` to be on PATH when `oracle_bin` is supplied.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Expose `--oracle-bin` (and optionally `--out-dir`) through MCP”
  - “Add MCP tool params: `oracle_bin … out_dir … work_dir …`”
  - “Append `--oracle-bin …` / `--out-dir …` / `--work-dir …` …”
```

```ticket T4
T# Title: Add preflight to fail fast on missing `-f` attachment files
Type: enhancement
Target Area: `oraclepack run` preflight; oracle invocation scanning via `ExtractOracleInvocations`; path resolution via WorkDir
Summary:
  Add a preflight check that detects missing attachment files referenced by `oracle … -f <path>` before executing steps. This turns late runtime failures into early, actionable errors that identify the step number and missing path. The check must resolve paths relative to the resolved WorkDir.
In Scope:
  - Before execution, scan each step’s code for oracle invocations.
  - For each invocation, verify every `-f <path>` exists relative to the resolved WorkDir.
  - Fail fast with an error like “step XX references missing file Y.”
Out of Scope:
  - Oracle CLI argument validation (handled in T5).
Current Behavior (Actual):
  - Missing attachments surface as runtime failures when `oracle` attempts to open `-f` inputs.
Expected Behavior:
  - Missing `-f` files are detected prior to running the step, with explicit step attribution.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must use the resolved WorkDir for existence checks.
Evidence:
  - “Missing attachment file preflight … verify every `-f <path>` exists … fail fast with ‘step XX references missing file Y.’”
Open Items / Unknowns:
  - Exact set of oracle invocations and how `ExtractOracleInvocations` represents `-f` arguments not provided.
Risks / Dependencies:
  - Depends on T2 for deterministic WorkDir resolution.
Acceptance Criteria:
  - If any step contains an oracle invocation with `-f <nonexistent>`, `oraclepack run` fails before executing steps and identifies the step number and missing path.
  - If all `-f` paths exist, the preflight does not block execution.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “4a) ‘Missing attachment file’ preflight … verify every `-f <path>` exists …”
  - “fail fast with ‘step XX references missing file Y.’”
```

```ticket T5
T# Title: Add optional `--preflight-oracle` gate using `--dry-run summary` to catch oracle flag drift
Type: enhancement
Target Area: `oraclepack run`; override validation via `ValidateOverrides`; invoking `oracle … --dry-run summary`
Summary:
  Add an optional preflight gate (`oraclepack run --preflight-oracle`) that validates extracted oracle invocations against the upstream `oracle` CLI using `--dry-run summary`. This catches flag/argument drift early, before spending time on full runs. The gate uses existing extraction and flag injection capabilities described as “already implemented.”
In Scope:
  - Add `oraclepack run --preflight-oracle` optional gate.
  - Extract invocations (noted as already implemented).
  - Inject flags safely (noted as already implemented).
  - Execute `oracle … --dry-run summary` and block execution if any invocation fails parse/validation.
Out of Scope:
  - Missing attachment existence checks (handled in T4).
Current Behavior (Actual):
  - Oracle CLI rejects invalid args/flags during execution time rather than preflight.
Expected Behavior:
  - When enabled, `--preflight-oracle` prevents execution if any extracted invocation fails `--dry-run summary`.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must block execution on any dry-run parse/validation failure.
Evidence:
  - “Wire this into `oraclepack run` as an optional gate, e.g. `oraclepack run --preflight-oracle`”
  - “Execute `--dry-run summary` and block …”
Open Items / Unknowns:
  - Exact mapping between “overrides” and invocation injection behavior not provided.
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - `oraclepack run --preflight-oracle` runs a dry-run validation over extracted invocations before step execution.
  - If any invocation fails dry-run parse/validation, `oraclepack run` fails without executing steps.
  - If all invocations pass, execution proceeds normally.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “4b) ‘Oracle flag drift’ preflight … `ValidateOverrides`”
  - “`oraclepack run --preflight-oracle`”
  - “Execute `--dry-run summary` and block …”
```

```ticket T6
T# Title: Align pack templates with `out_dir` runtime model using `$out_dir` in `--write-output`
Type: chore
Target Area: Pack templates/prelude; `--write-output` paths; `oraclepack --out-dir` coherence
Summary:
  Update pack templates so they use an `out_dir` variable inside the bash prelude and reference `$out_dir` in `--write-output`. This avoids hard-substituting `{{out_dir}}` everywhere and makes `oraclepack --out-dir` coherent with runtime behavior. Include directory creation via `mkdir -p`.
In Scope:
  - In pack prelude, include `out_dir="<resolved default>"` (or ensure oraclepack can infer it).
  - Use `--write-output "$out_dir/…"` in steps.
  - Add `mkdir -p "$out_dir"`.
Out of Scope:
  - Not provided
Current Behavior (Actual):
  - Templates hard-substitute `{{out_dir}}` everywhere, making `oraclepack --out-dir` unable to reliably “move outputs.”
Expected Behavior:
  - Packs use `$out_dir` for `--write-output` and ensure `$out_dir` exists, so `--out-dir` remains coherent.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Maintain pack runnability and alignment with oraclepack’s designed execution model.
Evidence:
  - “If … hard-substitute `{{out_dir}}` everywhere, `oraclepack --out-dir` can’t reliably ‘move outputs’…”
Open Items / Unknowns:
  - Which specific pack templates are in use by the skills is not provided in this ticket text.
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - Generated/updated packs include an `out_dir=...` definition in the bash prelude (or equivalent inference).
  - `--write-output` usages reference `$out_dir/...`.
  - Packs create the output directory with `mkdir -p "$out_dir"`.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Align pack templates with oraclepack’s `out_dir` model …”
  - “include `out_dir="<resolved default>"` …”
  - “Use `--write-output "$out_dir/…"` and `mkdir -p "$out_dir"`.”
```

```ticket T7
T# Title: Add `oraclepack generate` CLI (wrapper) for grouped tickets/codebase pack generation
Type: enhancement
Target Area: CLI (Cobra) command surface: `oraclepack generate`; generator wrapper invoking existing scripts
Summary:
  Add a first-class `oraclepack generate` command that produces the same Stage-1 artifacts as the grouped skills by delegating to existing generator scripts. The command should support tickets-grouped and codebase-grouped modes and then validate generated packs so drift is caught immediately. This provides immediate parity while keeping a future path to replace the backend with a native generator.
In Scope:
  - Add CLI subcommand `oraclepack generate tickets-grouped --ticket-root .tickets --out-dir <...>`.
  - Add CLI subcommand `oraclepack generate codebase-grouped --code-root . --out-dir <...>`.
  - Implement wrapper behavior: shell out to the described scripts:
    - `.../oraclepack-tickets-pack-grouped/scripts/generate_grouped_packs.py`
    - `.../oraclepack-codebase-pack-grouped/scripts/generate_grouped_packs.py`
  - After generation, call `oraclepack validate` (and optionally the skill’s `validate_pack.py`/lint scripts) to fail fast on shape drift.
  - Ensure outputs match the described contract: `out_dir/packs/*.md` plus `_groups.json`/manifest.
  - Ensure generated packs remain runner-ingestible (one bash fence, 20 steps, correct headers, write-output paths).
Out of Scope:
  - Implementing the native Go generator backend (handled in T8).
Current Behavior (Actual):
  - `oraclepack` does not natively generate the same grouped-pack artifacts; generation exists in skills/scripts.
Expected Behavior:
  - `oraclepack generate …` produces grouped-pack artifacts and immediately validates them.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - CLI shape should follow Cobra patterns and include the described flags (`--ticket-root`, `--code-root`, `--out-dir`).
Evidence:
  - “Add a new CLI command … `oraclepack generate tickets-grouped …`”
  - “Implementation: … shell out to the same scripts …”
Open Items / Unknowns:
  - Exact location/path where the generator scripts live at runtime in the CLI environment is not provided.
Risks / Dependencies:
  - Depends on T1 to ensure `oraclepack validate` enforces strict pack shape.
Acceptance Criteria:
  - `oraclepack generate tickets-grouped …` creates `out_dir/packs/*.md` and `_groups.json`/manifest.
  - `oraclepack generate codebase-grouped …` creates `out_dir/packs/*.md` and `_groups.json`/manifest.
  - After generation, `oraclepack validate` is executed and failures block completion.
  - Generated packs meet the stated ingestibility constraints (single bash fence; 20 steps; correct step headers).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “wrap the existing skill generators behind `oraclepack generate`”
  - “shell out to the same scripts your skills run …”
  - “What the generator must output (match the skills exactly) …”
```

````ticket T8
T# Title: Implement native Go grouped-pack generator with embedded templates and deterministic grouping
Type: enhancement
Target Area: `internal/generate/*`; embedded templates; pack rendering and validation constraints
Summary:
  Replace external script dependency by embedding the canonical templates and implementing deterministic grouping and rendering in Go. The generator must match the grouped-skill output contract and produce packs that satisfy the pack parser/validator constraints (single bash fence, correct headers, 20 steps). It must fail on unresolved placeholders.
In Scope:
  - Embed the exact templates used by the skills into the `oraclepack` binary.
  - Port deterministic grouping rules into `internal/generate/*`:
    - discover files deterministically (lexicographic order)
    - group by subdir + infer loose files
    - split oversized groups into “part N” packs
  - Render packs from embedded templates and fail on unresolved placeholders.
  - Ensure generator output satisfies pack constraints:
    - parser requires a ` ```bash ... ``` ` block and step headers matching the step regex; ROI extracted from header line
    - generator emits exactly 20 steps (01..20) and exactly one bash fence
Out of Scope:
  - MCP tool exposure for generation (handled in T9).
Current Behavior (Actual):
  - Generation relies on external scripts/skills; native generator not present.
Expected Behavior:
  - `oraclepack` can generate grouped packs without external script dependencies while preserving deterministic outputs and strict pack constraints.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must match the described grouping rules and output constraints.
Evidence:
  - “Embed the exact templates used by the skills …”
  - “Port the deterministic grouping rules … `internal/generate/*` …”
Open Items / Unknowns:
  - Exact thresholds/limits used for “infer loose files” are referenced but not fully specified in the provided text.
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - Native generator can produce packs that pass `oraclepack validate` under strict shape rules.
  - Generator output is deterministic for a fixed input set (lexicographic discovery).
  - Generated packs contain exactly 20 steps and exactly one bash fence.
  - Generation fails when placeholders remain unresolved.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Embed the exact templates used by the skills …”
  - “discover files deterministically (lexicographic order) …”
  - “generator should always emit exactly 20 steps … exactly one bash fence …”
````

```ticket T9
T# Title: Add MCP generation tools for grouped packs (tickets and codebase)
Type: enhancement
Target Area: MCP server tool surface; new tools for pack generation returning packs + manifest
Summary:
  Expose grouped-pack generation as MCP tools so agents can request generation directly via MCP rather than only validate/list/run. The tools should return a structured result including generated packs and manifest information, aligned with the described output contract.
In Scope:
  - Add MCP tools:
    - `oraclepack_generate_tickets_grouped(ticket_root, out_dir, ...) -> {packs:[...], manifest:...}`
    - `oraclepack_generate_codebase_grouped(code_root, out_dir, ...) -> {packs:[...], manifest:...}`
  - Ensure MCP can access the same generation backend exposed by `oraclepack generate` (wrapper or native).
Out of Scope:
  - Implementing the generation backend itself (handled in T7/T8).
Current Behavior (Actual):
  - MCP exposes validate/list/run but not “generate packs.”
Expected Behavior:
  - MCP supports generation tools that produce grouped-pack artifacts and return packs + manifest details.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must align with the generator output contract: packs plus manifest (`_groups.json`/manifest).
Evidence:
  - “MCP: expose generation as tools … Add tools like: …”
Open Items / Unknowns:
  - Exact return schema for `{packs:[...], manifest:...}` beyond the example is not provided.
Risks / Dependencies:
  - Depends on T7 for an initial `oraclepack generate` surface to delegate to.
Acceptance Criteria:
  - MCP exposes both generation tools with the specified names/signatures (as provided).
  - Invoking either tool results in generated packs and a manifest being produced/returned.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “MCP: expose generation as tools …”
  - “`oraclepack_generate_tickets_grouped(…) -> {packs:[...], manifest:...}`”
  - “`oraclepack_generate_codebase_grouped(…) -> {packs:[...], manifest:...}`”
```

```ticket T10
T# Title: Codify and automate the “minimum validation pipeline” and parity acceptance flow
Type: tests
Target Area: Validation pipeline (generation-time + oraclepack-level + smoke execution); parity acceptance flow for generated packs
Summary:
  Establish the described minimum pipeline to prevent regressions and to prove parity of generated packs. The flow includes generation-time validation, oraclepack-level validation, parsing sanity checks, and execution smoke using a stub oracle binary. It should operationalize the “minimal acceptance test” sequence described.
In Scope:
  - Generation-time validation: `python3 scripts/validate_pack.py pack.md` (as part of the skill contract).
  - Oraclepack-level validation: `oraclepack validate pack.md` (after strict fence validation).
  - Parse sanity: `oraclepack list pack.md` shows 20 sequential steps.
  - Execution smoke without real oracle calls: run with a stub `oracle` binary via `--oracle-bin` and ensure steps execute and `-f` paths are checked.
  - Parity acceptance sequence:
    1) `oraclepack generate tickets-grouped ...`
    2) For each emitted pack: `oraclepack validate <pack.md>` and `oraclepack list <pack.md>`
    3) `oraclepack run --no-tui --yes --run-all <pack.md>` in a small fixture repo to confirm no missing `-f` paths / bad headers
Out of Scope:
  - Not provided
Current Behavior (Actual):
  - The pipeline is described but not established as an enforced regression-prevention mechanism.
Expected Behavior:
  - The minimum validation pipeline and parity flow are repeatable and used to catch shape/path/oracle-arg regressions.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must use `--oracle-bin` stub for smoke runs as described.
Evidence:
  - “Minimum validation pipeline to prevent regressions (CLI + MCP) …”
  - “Minimal acceptance test (proves parity with the skills) …”
Open Items / Unknowns:
  - Where this pipeline should live (CI job, script, make target, etc.) is not provided.
  - What constitutes the “stub `oracle` binary” behavior is not specified beyond being runnable via `--oracle-bin`.
Risks / Dependencies:
  - Depends on T7 because the parity flow begins with `oraclepack generate …`.
Acceptance Criteria:
  - The described pipeline steps can be executed end-to-end as a single repeatable flow.
  - Generated packs are validated (`oraclepack validate`) and list exactly 20 steps (`oraclepack list`) in the parity flow.
  - Smoke execution runs using `--oracle-bin` stub and fails when `-f` paths are missing.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Minimum validation pipeline to prevent regressions (CLI + MCP) …”
  - “Execution smoke … stub `oracle` binary via `--oracle-bin` …”
  - “Minimal acceptance test (proves parity with the skills) …”
```
