<filetree>
Project Structure:
└── .tickets
    ├── actions
    │   ├── Enable Action Packs Dispatch.md
    │   ├── Improving Oraclepack Workflow.md
    │   ├── Oraclepack Action Pack Integration.md
    │   ├── Oraclepack Action Pack Issue.md
    │   ├── Oraclepack Action Packs.md
    │   └── Oraclepack Compatibility Issues.md
    ├── mcp
    │   ├── Expose Oraclepack as MCP.md
    │   ├── MCP Server for Oraclepack.md
    │   ├── gaps-still-not-covered.md
    │   ├── gaps_part2-mcp-builder.md
    │   ├── oraclepack-MCP.md
    │   └── oraclepack_mcp_server.md
    ├── other
    │   ├── Oraclepack Pipeline Improvements.md
    │   ├── Oraclepack Prompt Generator.md
    │   ├── Oraclepack Workflow Enhancement.md
    │   └── Verbose Payload Rendering TUI.md
    ├── PRD-TUI
    │   ├── Oraclepack TUI Integration.md
    │   └── PRD-generator URL routing.md
    ├── Formalize LLM Decision Points.md
    ├── Oraclepack CLI MCP Parity.md
    ├── Oraclepack File Storage.md
    ├── Oraclepack Parity Automation.md
    ├── Oraclepack Schema Approach.md
    ├── Oraclepack bash fix.md
    ├── Oraclepack output verification issues.md
    ├── Oraclepack-CLI-agents.md
    ├── Publish OraclePack MCP.md
    └── Skills Integration Status.md

</filetree>

<source_code>
.tickets/Formalize LLM Decision Points.md
```
Parent Ticket:

* Title: Formalize enforceable LLM decision points for oraclepack generation/execution across Skills and MCP
* Summary:

  * The ticket enumerates a comprehensive set of pre-generation and runtime decision points (DP-01..DP-60) intended to improve pack generation and step execution quality. It also notes a preferred way to make these decision points “real, enforceable” by emitting small structured artifacts (JSON/YAML) that deterministic scripts consume, while keeping the emitted pack Markdown schema unchanged.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “an ‘extensive’ set of additional, unique LLM decision points you can introduce (or explicitly formalize)…”
* Global Constraints:

  * “keeping the emitted pack Markdown schema unchanged”
* Global Environment:

  * Unknown
* Global Evidence:

  * OpenAI Codex “Custom Prompts” documentation. ([OpenAI Developers][1])
  * Gemini CLI “Agent Skills” documentation. ([Gemini CLI][2])
  * MCP specification revision 2025-11-25. ([Model Context Protocol][3])
  * MCP tools specification (server/tools) revision 2025-06-18. ([Model Context Protocol][4])

Split Plan:

* Coverage Map:

  * “In Codex/Gemini CLI, the ‘live agent’ inference…” → Info-only
  * “In an MCP setup, the ‘live agent’ inference…” → Info-only
  * “Below is an ‘extensive’ set…” → Info-only
  * DP-01 Choose generator mode (tickets-grouped vs codebase-grouped vs gold) → T2
  * DP-02 Select root(s) to scan (ticket_root/code_root) → T2
  * DP-03 Decide include/exclude glob rules → T2
  * DP-04 Decide whether to treat “loose” items as first-class candidates → T2
  * DP-05 Decide whether to require “strict schema mode” for outputs → T2
  * DP-06 Choose max pack size strategy (by tokens/bytes/files) → T2
  * DP-07 Choose per-pack cap: number of tickets/files → T2
  * DP-08 Decide ticket/file canonical title extraction rule → T3
  * DP-09 Choose text preprocessing (stopwords, stemming, code tokens) → T3
  * DP-10 Decide duplicate threshold policy (strict vs lenient) → T3
  * DP-11 Decide duplicate merge strategy (merge vs link vs pick canonical) → T3
  * DP-12 Decide what “differences” to preserve when merging duplicates → T3
  * DP-13 Decide hierarchical topic taxonomy vs flat groups → T4
  * DP-14 Decide group count target → T4
  * DP-15 Decide grouping algorithm choice (heuristic vs LLM-labeled vs hybrid) → T4
  * DP-16 Decide “ambiguous” assignment policy (multi-home vs best-fit) → T4
  * DP-17 Decide whether to create an “Unsorted / Needs triage” pack → T4
  * DP-18 Generate group names optimized for human scanning → T4
  * DP-19 Decide pack order (dependency, ROI, urgency, confidence) → T4
  * DP-20 Decide ticket order within pack (chronological vs dependency graph) → T4
  * DP-21 Select “context bundle” files to attach per pack → T5
  * DP-22 Decide whether to attach full files vs excerpts/summaries → T5
  * DP-23 Decide whether to generate a synthetic “pack brief” doc → T5
  * DP-24 Decide whether to include prior run artifacts (outputs) as inputs → T5
  * DP-25 Choose template variant (tickets vs codebase vs mixed) → T5
  * DP-26 Decide whether to inject organization standards into prompts → T5
  * DP-27 Decide step allocation across subtopics (steps per subdomain) → T5
  * DP-28 Choose prompt “stance” (audit-first vs implement-first vs design-first) → T5
  * DP-29 Choose “evidence bar” (strict citations vs lightweight) → T5
  * DP-30 Decide per-step required outputs (files changed, diffs, JSON, etc.) → T5
  * DP-31 Decide self-contained steps vs relying on previous outputs → T5
  * DP-32 Decide whether to add “ask-user” gates for missing critical inputs → T5
  * DP-33 Choose which MCP tools to call during generation (list/validate/generate) → T6
  * DP-34 Decide oracle model/engine selection and parameters → T6
  * DP-35 Decide whether to preflight oracle invocations (“--dry-run”) → T6
  * DP-36 Decide redaction policy for sensitive strings → T6
  * DP-37 Decide trace/correlation ID scheme for packs/steps → T6
  * DP-38 Decide which metrics/log events must be emitted per stage → T6
  * DP-39 Decide run strategy (run-all vs selective) → T7
  * DP-40 Decide concurrency / rate-limiting policy → T7
  * DP-41 Decide fail-fast vs continue collecting failures → T7
  * DP-42 Decide retry policy per failure type → T7
  * DP-43 Decide “prompt patch” for retries → T7
  * DP-44 Decide when to escalate to user for clarification → T7
  * DP-45 Decide whether to re-run earlier steps on contradictions → T7
  * DP-46 Decide acceptance criteria for a step output (format, completeness) → T8
  * DP-47 Decide whether to auto-validate produced artifacts (lint/tests/validate) → T8
  * DP-48 Decide how to synthesize step outputs into a final report → T9
  * DP-49 Decide PR/patch vs documentation-only output → T9
  * DP-50 Decide how to resolve conflicting recommendations across steps → T9
  * DP-51 Decide whether outputs meet policy/security standards before writing → T9
  * DP-52 Decide whether to reuse cached groupings/packs vs regenerate → T10
  * DP-53 Decide cache invalidation scope (one pack vs all) → T10
  * DP-54 Decide incident-style annotation on failures → T10
  * DP-55 Decide what artifacts to persist (manifests, intermediate summaries) → T10
  * DP-56 Decide how to shard outputs into “mini-packs” for follow-on runs → T10
  * DP-57 Decide naming/versioning for generated packs → T10
  * DP-58 Decide “next best tool call” in MCP (validate vs list vs run vs regenerate) → T10
  * DP-59 Decide whether to present diffs, file lists, or narrative only → T10
  * DP-60 Decide whether to extract new org heuristics into a reusable profile → T10
  * “make each decision point produce a small structured artifact (JSON/YAML)…” → T1
* Dependencies:

  * T2 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T3 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T4 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T5 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T6 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T7 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T8 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T9 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T10 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.

```ticket T1
T# Title: Make decision points enforceable via structured decision artifacts while keeping pack Markdown schema unchanged
Type: chore
Target Area: Pack generation/execution control plane (decision capture + deterministic consumption)
Summary:
- Establish the “typical pattern” described: each decision point emits a small structured artifact (JSON/YAML) and deterministic scripts consume it. Ensure the emitted pack Markdown schema remains unchanged. This provides a consistent way to formalize the DP-01..DP-60 decisions for both Skills-based flows and MCP tool-calling flows.
In Scope:
- “make each decision point produce a small structured artifact (JSON/YAML)”
- “make the deterministic scripts consume it”
- “keeping the emitted pack Markdown schema unchanged”
- Applicability across “skills” and “MCP”
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Decision points emit a small structured artifact (JSON/YAML).
- Deterministic scripts consume the artifact.
- Emitted pack Markdown schema remains unchanged.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Keep emitted pack Markdown schema unchanged.
Evidence:
- Text: “make each decision point produce a small structured artifact (JSON/YAML)… deterministic scripts consume it… pack Markdown schema unchanged.”
Open Items / Unknowns:
- Exact schema/fields for the structured artifact (JSON vs YAML specifics not provided).
- Where deterministic consumption is implemented (locations not provided).
Risks / Dependencies:
- Not provided
Acceptance Criteria:
- A structured decision artifact format exists that can represent DP outputs.
- Deterministic scripts consume the decision artifact to drive generation/execution flow.
- No change to the emitted pack Markdown schema is required to adopt the pattern.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “make each decision point produce a small structured artifact (JSON/YAML)”
- “make the deterministic scripts consume it”
- “keeping the emitted pack Markdown schema unchanged”
```

```ticket T2
T# Title: Implement pre-generation routing, scope, governance, and budgeting decision points (DP-01..DP-07)
Type: enhancement
Target Area: Pre-gen decision layer (routing/scope/governance/budgeting)
Summary:
- Add explicit, formalized pre-generation decisions for selecting generator mode, scan roots, include/exclude rules, loose-item handling, strict schema mode, and pack sizing/caps. Each decision produces the specified output/action for use by pack generation.
In Scope:
- DP-01 Choose generator mode (tickets-grouped vs codebase-grouped vs gold)
- DP-02 Select root(s) to scan (ticket_root/code_root)
- DP-03 Decide include/exclude glob rules
- DP-04 Decide whether to treat “loose” items as first-class candidates
- DP-05 Decide whether to require “strict schema mode” for outputs
- DP-06 Choose max pack size strategy (by tokens/bytes/files)
- DP-07 Choose per-pack cap: number of tickets/files
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Generator mode can be selected and recorded.
- Scan roots and include/exclude patterns can be selected and recorded.
- Loose-item handling policy and strict schema mode selection can be decided and recorded.
- Pack sizing strategy and per-pack caps can be decided and recorded.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Decisions should produce the “Output / action produced” listed for each DP (e.g., “Selected generator + params”, “Root paths”, “Include/exclude patterns”, “Sharding plan”, “Caps per pack”).
Evidence:
- DP-01..DP-07 rows (routing/scope/governance/budgeting)
Open Items / Unknowns:
- What “gold” mode entails (not provided).
- Default config values and where they are defined (not provided).
- Token/byte/file limits used for sizing (not provided).
Risks / Dependencies:
- Depends on T1 if adopting the structured decision artifact pattern.
Acceptance Criteria:
- DP-01 emits “Selected generator + params”.
- DP-02 emits “Root paths”.
- DP-03 emits “Include/exclude patterns”.
- DP-04 emits a “Loose-ticket policy”.
- DP-05 emits an “Enforce extra validation gates” decision.
- DP-06 emits a “Sharding plan”.
- DP-07 emits “Caps per pack”.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Choose generator mode (tickets-grouped vs codebase-grouped vs gold)”
- “Select root(s) to scan (ticket_root/code_root)”
- “Choose max pack size strategy (by tokens/bytes/files)”
```

```ticket T3
[TRUNCATED]
```

.tickets/Oraclepack CLI MCP Parity.md
```
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
[TRUNCATED]
```

.tickets/Oraclepack File Storage.md
```
Parent Ticket:

* Title: Stop oraclepack from writing run state/config JSON files into the project working directory
* Summary: oraclepack currently writes per-pack `*.state.json`, `*.report.json`, and `*.chatgpt-urls.json` files into the repo/working directory. The requested change is to store these as config/state/cache outside the repo root (prefer XDG base dirs) and/or under a dedicated project-local `.oraclepack/` directory to avoid clutter.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “Move state/report outputs out of CWD by default… Stop producing per-pack `*.chatgpt-urls.json` by default.”
* Global Constraints:

  * Treat outputs as **config/state/cache** and store outside repo root using XDG base dirs (per ticket text).
  * Use Go `os.UserConfigDir()` / `os.UserCacheDir()` for cross-platform defaults (per ticket text).
  * No `UserStateDir()` in Go stdlib; implement `$XDG_STATE_HOME` fallback (per ticket text).
* Global Environment:

  * Unknown
* Global Evidence:

  * Current filenames mentioned: `<packBase>.state.json`, `<packBase>.report.json`, `<sameBase>.chatgpt-urls.json`.
  * XDG Base Directory spec reference (background).
  * Go `os.UserConfigDir` / `os.UserCacheDir` reference (background).

Split Plan:

* Coverage Map:

  * Original item: “`oraclepack run` … derives filenames from the pack basename and writes them to the current working directory: `statePath := <packBase>.state.json`, `reportPath := <packBase>.report.json`”

    * Assigned Ticket ID: T2
  * Original item: “The TUI ‘ChatGPT URL picker’ then creates `<sameBase>.chatgpt-urls.json` next to the state file (or next to the pack file if statePath is empty).”

    * Assigned Ticket ID: T3
  * Original item: “It also defaults edits to **project scope**, so it will keep generating project-scoped stores unless the user explicitly switches to global.”

    * Assigned Ticket ID: T3
  * Original item: “Treat these as **config/state/cache** and store them outside the repo root using standard base dirs: … `$XDG_CONFIG_HOME` … `$XDG_STATE_HOME` … `$XDG_CACHE_HOME`…”

    * Assigned Ticket ID: T1
  * Original item: “In Go, you should use `os.UserConfigDir()` / `os.UserCacheDir()`… (There’s no `UserStateDir()`… implement XDG_STATE_HOME fallback…)”

    * Assigned Ticket ID: T1
  * Original item: “Move state/report outputs out of CWD by default… Update `internal/cli/run.go`… Make the directory overridable with a flag/env (e.g., `--state-dir` / `ORACLEPACK_STATE_DIR`).”

    * Assigned Ticket ID: T2
  * Original item: “Stop producing per-pack `*.chatgpt-urls.json` by default… Best UX default: change … default save scope to **global**…”

    * Assigned Ticket ID: T3
  * Original item: “Keep ‘project scope’ as an opt-in mode, but write it to a single per-project location (e.g., `<repo>/.oraclepack/chatgpt-urls.json`), not `<packName>.chatgpt-urls.json`.”

    * Assigned Ticket ID: T3
  * Original item: “Acceptable alternative (project-local…): `<repo>/.oraclepack/state/*.state.json` … `<repo>/.oraclepack/chatgpt-urls.json` … add `.oraclepack/` to `.gitignore`.”

    * Assigned Ticket ID: T4
  * Original item: “Immediate workaround (no code changes): Add these to `.gitignore`: `*.state.json`, `*.report.json`, `*.chatgpt-urls.json`.”

    * Assigned Ticket ID: T4
* Dependencies:

  * T2 depends on T1 because T2 needs an agreed/default “oraclepack state dir” location strategy (XDG-based) to write into.
  * T3 depends on T1 because T3 needs a global config location strategy (XDG-based) for URL persistence.
* Split Tickets:

```ticket T1
T# Title: Define XDG-based directory strategy for oraclepack config/state/cache
Type: chore
Target Area: Config/state path resolution (shared utility / helpers)
Summary:
- Define the standard locations where oraclepack stores user config, run state, and cache so outputs stop polluting the repo root.
- The ticket requires using XDG base dirs and Go’s cross-platform helpers where applicable, with an explicit fallback for state.
In Scope:
- Adopt XDG directory categories as the guiding model:
  - Config: `$XDG_CONFIG_HOME` (default `~/.config`)
  - State: `$XDG_STATE_HOME` (default `~/.local/state`)
  - Cache: `$XDG_CACHE_HOME` (default `~/.cache`)
- Use Go `os.UserConfigDir()` / `os.UserCacheDir()` for cross-platform defaults (per ticket text).
- Implement a state-dir resolver that honors `$XDG_STATE_HOME` and falls back when not set (since Go stdlib has no `UserStateDir()`).
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- oraclepack has a single, consistent mechanism to determine:
  - “oraclepack config dir” (for user prefs like URL lists)
  - “oraclepack state dir” (for resume/run state)
  - “oraclepack cache dir” (for non-essential cached data)
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Treat outputs as config/state/cache; store outside repo root using standard base dirs (per ticket text).
- Use `os.UserConfigDir()` / `os.UserCacheDir()` where applicable (per ticket text).
- Implement `$XDG_STATE_HOME` fallback logic (per ticket text).
Evidence:
- “Treat these as config/state/cache and store them outside the repo root using standard base dirs…” (parent ticket)
- “In Go, you should use os.UserConfigDir() / os.UserCacheDir()… There’s no UserStateDir()…” (parent ticket)
Open Items / Unknowns:
- Exact package/file locations for where to place the shared directory-resolution logic: Unknown
Risks / Dependencies:
- Not provided
Acceptance Criteria:
- A single directory-resolution mechanism exists for config/state/cache categories as described in scope.
- The state-dir resolution honors `$XDG_STATE_HOME` when set and has a documented fallback when unset.
- The config/cache resolution uses Go’s `os.UserConfigDir()` / `os.UserCacheDir()` (or equivalent wrapper) per ticket text.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Treat these as config/state/cache and store them outside the repo root using standard base dirs…”
- “In Go, you should use os.UserConfigDir() / os.UserCacheDir()… There’s no UserStateDir()…”
```

```ticket T2
T# Title: Move run-generated state/report JSON outputs out of CWD and add state-dir override
Type: enhancement
Target Area: Run command output paths (`internal/cli/run.go`)
Summary:
- oraclepack currently writes `<packBase>.state.json` and `<packBase>.report.json` into the current working directory.
- Update the run pathing so these files are written under a dedicated “oraclepack state dir” by default, with an override via flag/env.
In Scope:
- Change default output location for:
  - `<packBase>.state.json`
  - `<packBase>.report.json`
  from current working directory to a dedicated “oraclepack state dir”.
- Update `internal/cli/run.go` to compute state/report paths under that state dir (per ticket text).
- Add override via flag and env:
  - `--state-dir`
  - `ORACLEPACK_STATE_DIR`
Out of Scope:
- Not provided
Current Behavior (Actual):
- `<packBase>.state.json` and `<packBase>.report.json` are written to the current working directory.
Expected Behavior:
- By default, running oraclepack does not create `*.state.json` / `*.report.json` in the repo root / CWD.
- By default, state/report files are written under the dedicated oraclepack state dir.
- Setting `--state-dir` or `ORACLEPACK_STATE_DIR` writes state/report files under the specified directory.
Reproduction Steps:
1) Run `oraclepack run` from a repo root (or any working directory).
2) Observe creation of `<packBase>.state.json` and `<packBase>.report.json` in the working directory.
Requirements / Constraints:
- Must be overridable by `--state-dir` / `ORACLEPACK_STATE_DIR` (per ticket text).
- Should use the state-dir strategy defined in T1 for the default state dir.
Evidence:
- “`oraclepack run` … writes them to the current working directory: `statePath := <packBase>.state.json`, `reportPath := <packBase>.report.json`”
- “Update `internal/cli/run.go` … Make the directory overridable with a flag/env (e.g., `--state-dir` / `ORACLEPACK_STATE_DIR`).”
Open Items / Unknowns:
- Whether state/report filenames must remain exactly `<packBase>.state.json` / `<packBase>.report.json` or can change: Not provided
Risks / Dependencies:
- Depends on T1 for default state-dir resolution strategy.
Acceptance Criteria:
- Running oraclepack with no overrides does not create `*.state.json` or `*.report.json` in the current working directory.
- With no overrides, state/report files are written under the resolved oraclepack state dir.
- With `--state-dir=<dir>`, state/report files are written under `<dir>`.
- With `ORACLEPACK_STATE_DIR=<dir>`, state/report files are written under `<dir>`.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “`oraclepack run` … writes … `<packBase>.state.json` … `<packBase>.report.json` … to the current working directory”
- “Move state/report outputs out of CWD by default… Update internal/cli/run.go… `--state-dir` / `ORACLEPACK_STATE_DIR`”
```

```ticket T3
T# Title: Stop generating per-pack `*.chatgpt-urls.json`; default URL picker persistence to global store
Type: enhancement
Target Area: TUI “ChatGPT URL picker” persistence
Summary:
- The TUI URL picker currently creates `<sameBase>.chatgpt-urls.json` near the pack/state file and defaults edits to project scope.
- Change it so the default save scope is global (one file), while keeping project scope as an opt-in that writes to a single stable per-project path.
In Scope:
- Remove/avoid creating `<sameBase>.chatgpt-urls.json` “next to the state file (or next to the pack file…)” (per ticket text).
- Change the URL picker default save scope to **global** (per ticket text).
- Keep “project scope” as opt-in, but store at a single stable path:
  - `<repo>/.oraclepack/chatgpt-urls.json`
  rather than `<packName>.chatgpt-urls.json` (per ticket text).
- Persist the global URL store to a single global location:
  - Per ticket text, an existing global store path is referenced: `~/.oraclepack/chatgpt-urls.json`.
Out of Scope:
- Not provided
Current Behavior (Actual):
- URL picker creates `<sameBase>.chatgpt-urls.json` next to the state file (or pack file).
- URL picker defaults edits to project scope.
Expected Behavior:
- Using the URL picker does not create `<packBase>.chatgpt-urls.json` files.
- Default persistence is global (one stable file).
- Project scope, if selected, writes only to `<repo>/.oraclepack/chatgpt-urls.json`.
Reproduction Steps:
1) Use the TUI “ChatGPT URL picker” during a run.
2) Observe `<sameBase>.chatgpt-urls.json` being created near pack/state file.
Requirements / Constraints:
- Default save scope should be global (per ticket text).
- Project scope must be opt-in and must not create per-pack URL json files (per ticket text).
- Global persistence location:
  - Conflicting guidance exists: ticket recommends XDG config dir generally, but also references existing `~/.oraclepack/chatgpt-urls.json` path.
Evidence:
- “The TUI ‘ChatGPT URL picker’ then creates `<sameBase>.chatgpt-urls.json`…”
- “It also defaults edits to project scope…”
- “Stop producing per-pack `*.chatgpt-urls.json` by default… change … default save scope to global…”
- “Keep ‘project scope’ as an opt-in… write it to `<repo>/.oraclepack/chatgpt-urls.json`… not `<packName>.chatgpt-urls.json`.”
Open Items / Unknowns:
- Whether to keep `~/.oraclepack/chatgpt-urls.json` as the global path or migrate to `$XDG_CONFIG_HOME/...` (both appear in the parent ticket guidance).
- Exact file/path of the “URL picker” implementation: Not provided
Risks / Dependencies:
- Depends on T1 if migrating global storage to XDG config dir.
Acceptance Criteria:
- After using the URL picker, no `<packBase>.chatgpt-urls.json` is created near the pack/state/CWD.
[TRUNCATED]
```

.tickets/Oraclepack Parity Automation.md
```
Parent Ticket:

* Title: Keep oraclepack in parity with upstream oracle updates
* Summary: Define and automate a process so oraclepack (wrapper) stays compatible as upstream oracle changes. Approach centers on pinning the upstream oracle version, minimizing wrapper coupling, adding contract tests, and automating upgrade PRs + canary checks.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “keep parity with updates with a target repo as their repo updates/changes… wrapper around oracle… keep… up to date”
* Global Constraints:

  * Not provided
* Global Environment:

  * Node 22+ (per referenced upstream distribution note)
  * GitHub Actions (Dependabot + scheduled workflows)
  * Go toolchain used in CI example (version shown as 1.24.x)
* Global Evidence:

  * Dependabot configuration requires `.github/dependabot.yml` and supports npm + GitHub Actions ecosystems. ([GitHub Docs][1])
  * Renovate npm manager documentation (alternative automation). ([Renovate Docs][2])
  * peter-evans/create-pull-request GitHub Action (PR automation). ([GitHub][3])
  * Upstream Sync marketplace action (fork-sync scenario). ([GitHub][4])
  * Git submodule/subtree background (vendor option scenario). ([Atlassian][5])
  * Full original discussion text.

Split Plan:

* Coverage Map:

  * Original item: “treat upstream as a versioned dependency… automate detection/compat testing/PRs/releases” → Assigned Ticket ID: T1
  * Original item: “pins `@steipete/oracle`… runs deterministically” → Assigned Ticket ID: T1
  * Original item: “log/record the exact `oracle` version used into its `.state.json`/`.report.json`” → Assigned Ticket ID: T1
  * Original item: “Vendor via Node workspace… execute `node_modules/.bin/oracle`” → Assigned Ticket ID: T1
  * Original item: “Or run `npx -y @steipete/oracle@<pinned>`” → Assigned Ticket ID: T1
  * Original item: “Reduce parity surface area… pass-through custom args/flags forwarded unchanged” → Assigned Ticket ID: T2
  * Original item: “Only model/enumerate upstream flags where you must (Overrides Wizard)” → Assigned Ticket ID: T2
  * Original item: “If you do enumerate, auto-discover rather than hardcode” → Assigned Ticket ID: T3
  * Original item: “Add a contract test suite that detects upstream CLI surface changes early” → Assigned Ticket ID: T4
  * Original item: “Snapshot tests… capture `oracle --help` and `oracle <critical-subcommand> --help`… diff snapshots in CI” → Assigned Ticket ID: T4
  * Original item: “Behavioral ‘golden path’ tests… run a small pack fixture… assert only on wrapper-owned artifacts” → Assigned Ticket ID: T4
  * Original item: “Automate updates: Dependabot/Renovate” → Assigned Ticket ID: T5
  * Original item: “Scheduled ‘canary’ workflow… run against `@steipete/oracle@latest`” → Assigned Ticket ID: T7
  * Original item: “Release discipline: compatibility statement… gate releases on contract tests + canary” → Assigned Ticket ID: T8
  * Original item: “Concrete GitHub config examples: `.github/dependabot.yml`” → Assigned Ticket ID: T5
  * Original item: “Concrete workflow example: `.github/workflows/bump-oracle.yml` using `peter-evans/create-pull-request`” → Assigned Ticket ID: T6
  * Original item: “If you want ‘sync the fork with upstream’… use fork-sync/upstream-sync workflow/action” → Assigned Ticket ID: Info-only
  * Original item: “If you ever decide to vendor upstream code… use subtree/submodule… submodules don’t auto-track branches” → Assigned Ticket ID: Info-only
  * Original item: “Minimal next step… pinned version file + CI contract suite + enable Dependabot weekly updates” → Assigned Ticket ID: T1 (pin), T4 (contract suite), T5 (Dependabot)
* Dependencies:

  * T4 depends on T1 because contract tests are described as running against a pinned/deterministic upstream version.
  * T6 depends on T4 because the bump workflow is described as running unit + contract tests before opening a PR.
  * T7 depends on T4 because the canary job is described as running the integration/contract suite.
  * T8 depends on T4 and T7 because release gating is described as using contract tests and the canary signal.
* Split Tickets:

```ticket T1
T# Title: Pin upstream oracle version and record the runtime oracle version used
Type: enhancement
Target Area: Dependency boundary / runtime invocation / run artifacts (.state.json/.report.json)
Summary:
  Establish a deterministic upstream boundary by pinning the upstream `@steipete/oracle` version used by oraclepack runs. Ensure the exact oracle version used is recorded in wrapper-owned artifacts so reports/bugs are attributable to a specific upstream version.
In Scope:
  - Add a single source of truth for the pinned upstream oracle version (example given: `tools/oracle-version.txt`).
  - Ensure oraclepack uses the pinned `@steipete/oracle` version deterministically at runtime.
  - Record/log the exact oracle version used into `.state.json` and/or `.report.json` (as referenced).
  - Support one of the described execution modes:
    - Node workspace vendoring (`node_modules/.bin/oracle`), or
    - `npx -y @steipete/oracle@<pinned>` execution.
Out of Scope:
  - Upstream fork syncing and vendoring upstream source code (subtree/submodule) approaches.
Current Behavior (Actual):
  - Not provided
Expected Behavior:
  - oraclepack runs use a deterministic pinned oracle version rather than an implicitly installed/global/latest version.
  - oraclepack run artifacts include the upstream oracle version used.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Upstream oracle is referenced as an npm package `@steipete/oracle` and expects Node 22+ (as noted).
Evidence:
  - Not provided
Open Items / Unknowns:
  - Whether oraclepack currently shells out to a global oracle, vendors a local dependency, or uses npx (not provided).
  - Whether both `.state.json` and `.report.json` exist and their current schema/paths (not provided).
Risks / Dependencies:
  - Depends on CI/runtime environment having Node available if execution uses node workspace or npx.
Acceptance Criteria:
  - A pinned oracle version is stored in a single source-of-truth file (example: `tools/oracle-version.txt`).
  - oraclepack execution uses the pinned version deterministically (no implicit “latest”).
  - `.state.json` and/or `.report.json` emitted by a run includes the oracle version used.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “pins `@steipete/oracle`… runs that version deterministically”
  - “logs/records the exact `oracle` version used into its `.state.json`/`.report.json`”
  - “Vendor via Node workspace… or run `npx -y @steipete/oracle@<pinned>`”
```

```ticket T2
T# Title: Add pass-through forwarding for upstream oracle args/flags to reduce wrapper coupling
Type: enhancement
Target Area: CLI wrapper interface (oraclepack → oracle invocation)
Summary:
  Reduce breakage risk by defaulting to a “pass-through” model where oraclepack forwards custom args/flags to upstream oracle unchanged. Only model/enumerate upstream flags where oraclepack explicitly needs structured UX (e.g., an overrides wizard).
In Scope:
  - Provide a mechanism for users/config to supply freeform upstream args/flags that oraclepack forwards to oracle unchanged.
  - Define/implement the boundary: which flags are wrapper-owned vs upstream pass-through.
  - Keep modeled/enumerated upstream flags limited to where explicitly required (example given: Overrides Wizard).
Out of Scope:
  - Auto-discovery of upstream flags/commands for modeled UX (handled in T3).
Current Behavior (Actual):
  - Not provided
Expected Behavior:
  - oraclepack can forward arbitrary upstream oracle args/flags without needing wrapper updates for each new upstream option.
  - oraclepack only hard-couples to upstream flag semantics where required for wrapper UX.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Not provided
Evidence:
  - Not provided
Open Items / Unknowns:
  - Current oraclepack surface area that models upstream flags vs forwards arguments (not provided).
  - “Overrides Wizard” location/implementation details in oraclepack (not provided).
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - A user-visible way exists to pass through upstream oracle args/flags unchanged.
  - Wrapper-owned flags are not silently forwarded (clear separation in behavior/docs/help output).
  - Existing wrapper features that depend on specific upstream flags continue to function as before.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Reduce ‘parity surface area’… Prefer: ‘pass-through’ custom args/flags… forwards… unchanged”
  - “Only model/enumerate upstream flags where you must (e.g., your Overrides Wizard)”
```

```ticket T3
T# Title: Auto-discover upstream oracle CLI flags/commands used by modeled wrapper UX (avoid hardcoding)
Type: enhancement
Target Area: Wrapper UX that enumerates upstream options (e.g., Overrides Wizard)
Summary:
  Where oraclepack must enumerate upstream oracle flags/commands for structured UX, implement auto-discovery so upstream additions/changes do not require manual hardcoded updates. This specifically targets the stated risk of wrappers breaking when they re-model upstream too strictly.
In Scope:
  - Implement auto-discovery of upstream CLI surface for any wrapper UX that enumerates upstream flags/commands.
  - Ensure the enumerated list updates when the pinned upstream oracle version changes (ties to T1).
Out of Scope:
  - General pass-through forwarding (handled in T2).
Current Behavior (Actual):
  - Not provided
Expected Behavior:
  - Wrapper UX that enumerates upstream options derives its list from upstream oracle, not a hardcoded list.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Not provided
Evidence:
  - Not provided
Open Items / Unknowns:
  - Which wrapper flows enumerate upstream options (only “Overrides Wizard” is referenced; details not provided).
  - The discovery mechanism source (help output vs another interface) is not specified.
Risks / Dependencies:
  - Depends on T1 for a stable/pinned upstream version target during discovery.
Acceptance Criteria:
  - For any wrapper UX that enumerates upstream flags/commands, the list is derived from upstream oracle rather than hardcoded constants.
  - Updating the pinned upstream oracle version results in the enumerated UX reflecting new/changed upstream options without manual list edits.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “If you do enumerate, auto-discover rather than hardcode”
  - “The main way wrappers break is when they re-model upstream flags/commands too strictly”
```

```ticket T4
T# Title: Add contract test suite for upstream oracle CLI surface changes (snapshot + golden path)
Type: tests
Target Area: CI tests / compatibility checks between oraclepack and upstream oracle
Summary:
  Add tests that fail when upstream oracle changes in ways that could break oraclepack. Include snapshot diffs for help output and a small “golden path” integration fixture that asserts on stable, wrapper-owned artifacts.
In Scope:
  - Snapshot tests capturing `oracle --help` output and `oracle <critical-subcommand> --help` output for a small set of relied-upon commands, then diff snapshots in CI.
  - A “golden path” test that runs a small pack fixture to exercise oraclepack integration points (dry-run path; API path if keys exist; browser path optionally skipped in CI as described).
  - Assertions focus on wrapper-owned artifacts (exit codes, oraclepack report/state schema) rather than upstream human text output.
Out of Scope:
  - Automation that updates the pinned upstream version (handled in T5/T6).
Current Behavior (Actual):
  - Not provided
Expected Behavior:
  - CI provides early signal when upstream oracle help/CLI surface changes.
  - CI validates at least one end-to-end wrapper “golden path” behavior against the pinned upstream version.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Tests should target the pinned upstream oracle version (deterministic) rather than “latest”.
Evidence:
  - Not provided
Open Items / Unknowns:
  - The “critical subcommands” list oraclepack relies on (examples referenced, but exact dependency set not provided).
[TRUNCATED]
```

.tickets/Oraclepack Schema Approach.md
```
Parent Ticket:

* Title: Adopt a schema-driven approach to prevent oraclepack run failures
* Summary:

  * Current runs fail because structure is inferred from Markdown heuristics (e.g., exactly one ```bash fence, sequential step headers, exactly 20 steps).
  * Proposal: generate a machine-validated **manifest** (JSON Schema) and **deterministically render** the Markdown pack; optionally add stricter linting for Markdown-only packs.
* Source:

  * Link/ID (if present) or “Not provided”
  * Original ticket excerpt (≤25 words) capturing the overall theme

    * “separate ‘machine-validated structure’ from ‘human-readable Markdown’ … generate only a JSON manifest … then a deterministic renderer produces the Markdown pack.”
* Global Constraints:

  * Keep existing oraclepack Markdown contract / backward-compatible (“keep the existing Markdown contract for oraclepack execution”).
  * Steps must be exactly 20; step IDs must be sequential 01..20.
* Global Environment:

  * Unknown
* Global Evidence:

  * Error text: “invalid pack structure: no bash code block found”.
  * Pack constraints referenced: “Exactly one ```bash fence”, “Exactly 20 steps”, “sequential step numbers”.

Split Plan:

* Coverage Map:

  * Original item: “separate ‘machine-validated structure’ from ‘human-readable Markdown.’”

    * Assigned Ticket ID: T1
  * Original item: “AI generates only a JSON manifest that must validate against a JSON Schema; renderer produces Markdown pack.”

    * Assigned Ticket ID: T1
  * Original item: “Prevent missing/multiple ```bash fences (root cause of ‘invalid pack structure: no bash code block found’).”

    * Assigned Ticket ID: T3
  * Original item: “Prevent non-sequential steps (Go validator requires sequential step numbers).”

    * Assigned Ticket ID: T1
  * Original item: “Prevent wrong step count (enforce exactly 20 in schema).”

    * Assigned Ticket ID: T1
  * Original item: “Minimal ‘Pack Manifest v1’ JSON Schema (Draft 2020-12) with schema_version/kind/out_dir/write_output/steps; step fields id/title/bash plus roi/impact/confidence/effort/horizon/category/reference.”

    * Assigned Ticket ID: T1
  * Original item: “Rendering rule (deterministic): one ```bash fence; prelude out_dir=…; optional --write-output; each step ‘# 01) …’ with body.”

    * Assigned Ticket ID: T2
  * Original item: “If Markdown-only: add explicit schema/lint mode (exactly one ```bash fence; exactly 20 steps; sequential 01..20; optional header tokens).”

    * Assigned Ticket ID: T3
  * Original item: “Stage-2 directory contract: exactly one file per prefix 01-*.md … 20-*.md.”

    * Assigned Ticket ID: T3
  * Original item: “Action pack lint (Stage 3): one ```bash fence; enforce 01..20 exact count.”

    * Assigned Ticket ID: T3
  * Original item: “CI checks: validate(manifest.json) → render(pack.md) → oraclepack validate pack.md → (optional) dry-run checks.”

    * Assigned Ticket ID: T4
* Dependencies:

  * T2 depends on T1 because the renderer needs the validated “Pack Manifest v1” structure as input.
  * T4 depends on T1 and T2 because CI runs “validate(manifest.json) → render(pack.md)”.
* Split Tickets:

```ticket T1
T# Title: Define and validate “Pack Manifest v1” schema (manifest-first)
Type: chore
Target Area: Pack authoring contract (manifest JSON + JSON Schema validation)
Summary:
- Introduce a manifest-first source of truth: the AI produces a JSON manifest that must validate against a JSON Schema.
- The schema enforces step count (exactly 20) and step IDs (01..20) to prevent structural runner failures.
- This separates machine-validated structure from the Markdown pack to reduce malformed packs.
In Scope:
- Define “Pack Manifest v1” JSON Schema (Draft 2020-12) with required fields: schema_version (const 1), kind (enum), out_dir, steps (min/max 20).
- Define step object constraints: required id/title/bash; id pattern for 01..20; optional roi/impact/confidence/effort/horizon/category/reference.
- Validate manifests against the schema before rendering/using them.
Out of Scope:
- Not provided
Current Behavior (Actual):
- Runner infers structure from Markdown heuristics; malformed structure can cause run-time validation errors.
- Step count and sequential numbering can be violated if not enforced early.
Expected Behavior:
- A manifest that does not conform (wrong count, wrong IDs, missing required fields) is rejected by schema validation.
- Manifests accepted by validation always contain exactly 20 steps with valid IDs and required fields.
Reproduction Steps:
1) Provide a manifest with fewer than 20 steps.
2) Provide a manifest with a non-matching step id (e.g., "21" or "1").
3) Validate manifest and confirm it fails schema validation.
Requirements / Constraints:
- schema_version must be 1.
- steps must be exactly 20 (minItems=20, maxItems=20).
- step id must match 01..20 via pattern.
Evidence:
- “the AI generates only a JSON manifest that must validate against a JSON Schema” (ticket text)
- “Wrong step count (you can enforce exactly 20 in schema rather than ‘hoping’ the model did it).”
Open Items / Unknowns:
- Exact location/path conventions for storing manifest.json are not provided.
- How/where validation is invoked (CLI, CI, library) is not provided.
Risks / Dependencies:
- Risk: keeping backward compatibility requires rendering to the existing Markdown pack contract (handled in T2).
Acceptance Criteria:
- [ ] A JSON Schema exists for “Pack Manifest v1” with required fields and constraints as described.
- [ ] A manifest with != 20 steps fails validation.
- [ ] A manifest with an invalid step id fails validation.
- [ ] A manifest missing required fields (schema_version/kind/out_dir/steps) fails validation.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Manifest-first + JSON Schema (then render Markdown)”
- “minItems: 20, maxItems: 20”
- “id … pattern: ^(0[1-9]|1[0-9]|20)$”
```

````ticket T2
T# Title: Implement deterministic renderer from manifest → oraclepack Markdown pack
Type: enhancement
Target Area: Pack rendering (manifest → Markdown pack)
Summary:
- Add a deterministic rendering rule that converts a validated manifest into a Markdown pack that always satisfies oraclepack’s structural expectations.
- This prevents issues like missing/multiple bash fences and malformed step formatting by making Markdown a compiled artifact.
In Scope:
- Render exactly one fenced code block labeled `bash` in the entire document.
- Render prelude lines including: `out_dir="..."` and optional `--write-output` as described.
- Render each step with header `# NN) ...` and step body from `bash` content.
Out of Scope:
- Not provided
Current Behavior (Actual):
- Markdown is the primary artifact; structure can be malformed by generation, causing downstream failures.
Expected Behavior:
- Renderer output always includes exactly one `bash` fence and emits all 20 steps in order.
- Pack contains the required prelude line(s) described in the ticket text.
Reproduction Steps:
1) Validate a manifest (per T1).
2) Render to Markdown.
3) Confirm output contains exactly one `bash` fence and step headers for 01..20 in sequence.
Requirements / Constraints:
- Output must be runner-ingestible per the described structural rules (single bash fence, 20 steps, sequential).
Evidence:
- “Rendering rule (deterministic) … emits exactly: one ```bash fence … prelude lines … then each step: # 01) … Step body = bash”
Open Items / Unknowns:
- Exact step title formatting beyond “# NN) …” is not provided.
- Whether additional header tokens (ROI=…) are required at render time is not provided.
Risks / Dependencies:
- Depends on T1 (renderer assumes manifest structure/constraints).
Acceptance Criteria:
- [ ] Given a valid manifest, renderer produces a Markdown pack with exactly one ```bash fenced block.
- [ ] Output contains 20 step headers numbered 01..20 in order.
- [ ] Output includes the `out_dir="..."` prelude line.
- [ ] Renderer can conditionally include the optional `--write-output` line when present in the manifest.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “From this manifest, your renderer emits exactly: One ```bash fence”
- “Prelude lines: out_dir="..." … optional --write-output”
- “Then each step: # 01) … Step body = bash”
````

````ticket T3
T# Title: Add stricter lint/validation for Markdown-only packs and Stage-2/Stage-3 outputs
Type: chore
Target Area: Pack linting/validation (Markdown packs + output directory contract)
Summary:
- If the project keeps Markdown-only as a supported input path, add an explicit lint/validation mode that enforces the same structural contract.
- Extend checks to Stage-2 output directory naming expectations and Stage-3 action pack constraints to reduce “runner infers structure” failures.
In Scope:
- Pack-level lint (Stage 1) enforcing:
  - Exactly one ```bash fence.
  - Exactly 20 steps.
  - Step IDs exactly 01..20 and sequential.
  - Optional enforcement of required header tokens (ROI= impact= confidence= … reference=) as described.
- Stage-2 directory contract lint:
  - Exactly one file per prefix 01-*.md … 20-*.md.
- Stage-3 action pack lint:
  - Exactly one ```bash fence.
  - Enforce 01..20 and exact count.
Out of Scope:
- Not provided
Current Behavior (Actual):
- Common failure mode noted: “invalid pack structure: no bash code block found.”
- Existing checks may be incomplete (“your current check only ensures ‘some’ step headers exist” per ticket text).
Expected Behavior:
- Markdown packs that violate the contract are rejected early with lint errors before execution.
- Stage-2 outputs and Stage-3 action packs are validated against exact-count and naming/structure rules.
Reproduction Steps:
1) Create a Markdown pack with no `bash` fenced block → lint should fail.
2) Create a Markdown pack with 19 steps or non-sequential IDs → lint should fail.
3) Create an output directory missing `07-*.md` or containing duplicates for a prefix → lint should fail.
Requirements / Constraints:
- Enforce: one ```bash fence, exactly 20 steps, sequential 01..20.
Evidence:
- “Missing / multiple ```bash fences (common root cause of ‘invalid pack structure: no bash code block found’).”
- “Add an explicit schema/lint mode … Exactly one ```bash fence … Exactly 20 steps … Step IDs exactly 01..20”
- “Stage-2 directory contract … Exactly one file per prefix 01-*.md … 20-*.md”
- “Action pack lint (Stage 3) … Enforce 01..20 and exact count”
Open Items / Unknowns:
- Exact current validator behaviors and what already exists vs missing are not provided.
Risks / Dependencies:
- Risk: enforcing optional header tokens could break existing packs if not already standardized.
Acceptance Criteria:
- [ ] Lint fails when no `bash` fence exists and surfaces a clear error.
- [ ] Lint fails when step count != 20.
- [ ] Lint fails when step IDs are not exactly 01..20 sequential.
- [ ] Stage-2 lint fails when any step output prefix 01..20 is missing or duplicated.
- [ ] Stage-3 lint fails when action pack does not have exactly one `bash` fence or correct 01..20 steps.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “invalid pack structure: no bash code block found”
- “Pack-level lint (Stage 1) … Exactly one ```bash fence … Exactly 20 steps”
- “Stage-2 directory contract … Exactly one file per prefix 01-*.md … 20-*.md”
````

```ticket T4
T# Title: Add CI validation pipeline for manifest-first workflow (validate → render → oraclepack validate → optional dry-run)
Type: chore
Target Area: CI checks / pipeline gating
Summary:
- Add CI checks that gate merges/runs on structural correctness by validating the manifest, rendering Markdown deterministically, and validating the rendered pack with oraclepack tooling.
- This formalizes the “Markdown is compiled artifact” approach and reduces runtime surprises.
In Scope:
- CI sequence as described:
  - validate(manifest.json)
  - render(pack.md)
  - oraclepack validate pack.md
  - optional dry-run checks
Out of Scope:
- Not provided
Current Behavior (Actual):
[TRUNCATED]
```

.tickets/Oraclepack bash fix.md
```
Parent Ticket:

* Title: Prevent oraclepack pack failures caused by orphaned `-p/--prompt` lines in generated bash steps
* Summary: Generated oraclepack markdown packs can emit a multiline `oracle ...` command where `-p "$(cat <<'PROMPT' ...)"` starts on a new line without a continuation, causing Bash to treat `-p` as a standalone command and fail (`exit status 127`). The fix requires making pack generation structurally safe and adding validator guardrails that fail fast on regressions.
* Source:

  * Link/ID: Bash command syntax fix.md
  * Original ticket excerpt (≤25 words) capturing the overall theme: “make the generator/template structurally unable to emit orphaned flags… and make oraclepack validate fail fast”
* Global Constraints:

  * “never put `-p/--prompt` (or any flag) on its own line”
  * “no inline comments at end of an `oracle ...` line”
* Global Environment:

  * Unknown
* Global Evidence:

  * Error: `bash: line 59: -p: command not found` + `exit status 127`
  * Reference pack: `oracle-pack-2026-01-08-tickets-direct.md` (pattern repeated across steps)

Split Plan:

* Coverage Map:

  * “`bash: line 59: -p: command not found` + `exit status 127`” → T1
  * “`-p "$(cat <<'PROMPT' ...)"` is on the next line without `\` … repeated in others” → T1
  * “Minimal fix: add a continuation backslash, or put `-p` on the same line” → T1
  * “Optional comment goes ABOVE the command, not inline” → T1
  * “Wherever you render `oracle ... "${ticket_args[@]}" # extra_files ...` then newline then `-p ...` … change it” → T1
  * “Permanent template rule: never put `-p/--prompt` on its own line; build prompt first, then call oracle on a single command line” → T1
  * “Enforce: no inline comments at end of an `oracle ...` line” → T1
  * “If you must wrap long lines, require explicit `\` continuations and disallow comments on continued lines” → T1
  * “Add checks to `oraclepack validate` after extracting the single `bash` fence” → T2
  * “Add `bash -n` syntax check” → T2
  * “Add `shellcheck` static analysis” → T2
  * “Custom ‘orphaned flag line’ detector (regex + continuation exceptions)” → T2
  * “Ensure `oraclepack run` always calls `validate` first (or at minimum in TUI Run/Rerun paths)” → T3
  * “Add CI/pre-commit: run `oraclepack validate` on any generated/modified pack” → T3
  * “Offer: point at exact rendering pattern + canonical snippet” → Non-actionable / Info-only
* Dependencies:

  * T3 depends on T2 because “Make validate unavoidable” is intended to enforce the added validator checks that catch regressions.
* Split Tickets:

```ticket T1
T# Title: Make pack generation structurally safe (no orphaned `-p/--prompt` lines)
Type: bug
Target Area: Pack generator/template that emits oraclepack Markdown steps (tickets-direct pack generation)
Summary:
- Generated packs can split an `oracle ...` invocation across lines such that `-p "$(cat <<'PROMPT' ...)"` starts on a new line without a continuation.
- Bash then executes `-p` as a standalone command, causing `command not found` and `exit status 127`.
- Update generation patterns so prompts are built safely and the `oracle` command remains a single logical command (or uses correct continuations without inline comment footguns).
In Scope:
- Eliminate multiline `oracle` invocations that place `-p/--prompt` on its own line.
- Apply the “minimal fix” pattern where multiline is unavoidable: add a line-continuation `\` (and ensure comments do not break continuation).
- Enforce generator rule: no inline trailing comments on `oracle ...` lines (comments/newlines can terminate the command unexpectedly).
- Adopt canonical “build prompt first, then call oracle” step shape as the standard emission pattern.
Out of Scope:
- Not provided
Current Behavior (Actual):
- `oracle ...` command is terminated by a newline, then `-p "$(cat <<'PROMPT' ...)"` appears on the next line without `\`, so Bash treats `-p` as a command and fails.
Expected Behavior:
- Generated bash steps never emit orphaned flag lines (e.g., `-p`, `-f`, `--prompt`) that can be interpreted as standalone shell commands.
- Generated `oracle` invocations are either a single logical line or correctly continued (without inline comments breaking continuation).
Reproduction Steps:
1. Run the generated pack `oracle-pack-2026-01-08-tickets-direct.md`.
2. Observe the step where `-p` begins a new line without a continuation and the shell errors.
Requirements / Constraints:
- “never put `-p/--prompt` (or any flag) on its own line”
- “no inline comments at end of an `oracle ...` line”
- If wrapping is necessary: require explicit `\` continuations and disallow comments on continued lines.
Evidence:
- Error: `bash: line 59: -p: command not found` + `exit status 127`
- Pattern described: `-p "$(cat <<'PROMPT' ...)"` on next line without `\` in `oracle-pack-2026-01-08-tickets-direct.md`
Open Items / Unknowns:
- Exact location(s) of the emitting template(s): Unknown / Not provided
- Whether multiple generators/templates emit the pattern beyond tickets-direct: Unknown / Not provided
Risks / Dependencies:
- Risk: Partial fixes (only adding `\`) may regress if inline comments or formatting are reintroduced.
Acceptance Criteria:
- Generated packs do not contain any step where a line begins with `-p` / `--prompt` intended as a continuation of `oracle` without an explicit safe structure.
- Running the regenerated tickets-direct pack no longer produces `-p: command not found` / `exit status 127` for the previously failing steps.
- Generator output follows one of:
  - prompt built first + `oracle ... --prompt "$prompt"` as a single logical command, OR
  - explicit `\` continuation with no inline trailing comments on continued lines.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “`bash: line 59: -p: command not found` + `exit status 127`”
- “the `-p "$(cat <<'PROMPT' ...)"` part is on the next line without a line-continuation (`\`)”
- “Permanent template rule: never put `-p/--prompt` (or any flag) on its own line”
```

```ticket T2
T# Title: Add validator guardrails (bash-lint + orphaned-flag detection) to fail fast
Type: enhancement
Target Area: `oraclepack validate` (after extracting the single `bash` fence)
Summary:
- Even with a safer generator, regressions can reintroduce orphaned flag lines that only fail at runtime.
- Add validation checks that detect bash syntax issues and the specific “orphaned flag line” class before execution.
- Validation should clearly fail on suspicious standalone flag lines unless safely continued.
In Scope:
- Run `bash -n` against the extracted bash step script(s) as a syntax sanity check.
- Run `shellcheck` static analysis on the extracted bash script(s).
- Implement the custom “orphaned flag line” detector:
  - Fail if a non-heredoc line matches `^\s*-(p|f)\b` or `^\s*--(prompt|file|write-output)\b`
  - Unless the previous non-empty line ends with a legal continuation (`\`, `|`, `&&`, `||`, `(`, etc.)
Out of Scope:
- Making `validate` mandatory in all run paths (handled separately)
Current Behavior (Actual):
- Not provided
Expected Behavior:
- `oraclepack validate` fails fast with a clear error when a pack contains likely orphaned flag lines (e.g., `-p ...`) outside permitted continuation contexts.
- `oraclepack validate` reports bash syntax issues before execution.
Reproduction Steps:
1. Create/modify a pack step where `-p` is on its own line without a valid continuation.
2. Run `oraclepack validate`.
Requirements / Constraints:
- Checks are added “after extracting the single `bash` fence”.
- Orphaned-flag detector must ignore heredoc bodies (“non-heredoc line”).
Evidence:
- “Add these checks to `oraclepack validate` after extracting the single `bash` fence”
- Detector specification (regex + continuation exceptions) provided in ticket text.
Open Items / Unknowns:
- Availability/installation method for `shellcheck` in the execution environment: Unknown / Not provided
- Exact current structure of `oraclepack validate` and how it extracts bash fence: Unknown / Not provided
Risks / Dependencies:
- Risk: False positives if continuation heuristics are too strict; must match the specified allowed continuations.
Acceptance Criteria:
- `oraclepack validate` includes `bash -n` and fails on invalid bash syntax in the extracted script(s).
- `oraclepack validate` runs `shellcheck` and surfaces failures per project policy (pass/fail behavior not specified in ticket text).
- `oraclepack validate` fails when a non-heredoc line begins with `-p`, `-f`, `--prompt`, `--file`, or `--write-output` and the previous non-empty line does not end with an allowed continuation token.
- `oraclepack validate` does not falsely flag valid heredoc prompt bodies.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Add these checks to `oraclepack validate` after extracting the single `bash` fence”
- “`bash -n` syntax check (cheap sanity)”
- “Custom ‘orphaned flag line’ detector… `^\s*-(p|f)\b` … unless… ends with… (`\`, `|`, `&&`, `||`, `(`, etc.)”
```

```ticket T3
T# Title: Make validation unavoidable in normal use (run/TUI) and add CI/pre-commit gate
Type: chore
Target Area: `oraclepack run` execution flow, TUI “Run/Rerun” paths, and repo automation (CI/pre-commit)
Summary:
- Validation guardrails are only effective if they run consistently before pack execution.
- Ensure `oraclepack run` calls `validate` first (or at minimum in TUI Run/Rerun paths).
- Add automated gating so modified/generated packs are validated before being executed/merged.
In Scope:
- Ensure `oraclepack run` always calls `validate` first.
- Ensure TUI “Run/Rerun” paths invoke `validate` first (at minimum).
- Add CI/pre-commit step to run `oraclepack validate` on generated/modified packs.
Out of Scope:
- Implementing the validator checks themselves (handled separately)
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Running a pack via CLI or TUI triggers validation first, preventing execution of invalid packs.
- CI/pre-commit blocks changes that introduce invalid pack structure detectable by `oraclepack validate`.
Reproduction Steps:
1. Introduce a known-invalid pattern (e.g., orphaned `-p` line) into a pack.
2. Attempt to run via `oraclepack run` and via TUI Run/Rerun.
3. Attempt to commit/CI-run with the invalid pack present.
Requirements / Constraints:
- “Ensure `oraclepack run` always calls `validate` first (or at minimum in TUI ‘Run/Rerun’ paths).”
- “Add CI/pre-commit: run `oraclepack validate` on any generated/modified pack.”
Evidence:
- The ticket text specifies making validation unavoidable and adding CI/pre-commit gating.
Open Items / Unknowns:
- Existing CI/pre-commit tooling and where to hook validation: Unknown / Not provided
- Exact TUI entrypoints for Run/Rerun: Unknown / Not provided
Risks / Dependencies:
- Depends on `oraclepack validate` providing the intended guardrails to justify making it mandatory.
Acceptance Criteria:
- `oraclepack run` invokes `validate` before executing pack steps.
- TUI Run/Rerun paths invoke `validate` before execution (at minimum).
- CI/pre-commit configuration exists to run `oraclepack validate` on generated/modified packs and fails on validation errors.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Make validate unavoidable in normal use”
- “Ensure `oraclepack run` always calls `validate` first (or at minimum in TUI ‘Run/Rerun’ paths).”
[TRUNCATED]
```

.tickets/Oraclepack output verification issues.md
```
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
[TRUNCATED]
```

.tickets/Oraclepack-CLI-agents.md
```

In Codex/Gemini CLI, the “live agent” inference is the model deciding what actions to take given the skill context (skills are discovered/activated and the model can inspect skill assets). [OpenAI Developers+2OpenAI Developers+2](https://developers.openai.com/codex/custom-prompts/?utm_source=chatgpt.com) In an MCP setup, the “live agent” inference is the model deciding which tools to call (and with what arguments) based on the tools’ schemas. [Model Context Protocol+3Model Context Protocol+3Model Context Protocol+3](https://modelcontextprotocol.io/specification/2025-11-25?utm_source=chatgpt.com)

Below is an “extensive” set of additional, unique LLM decision points you can introduce (or explicitly formalize) beyond “classify tickets/domains then feed into grouping”.

| ID | Stage | Class | LLM decision point | Inputs considered | Output / action produced |
| --- | --- | --- | --- | --- | --- |
| DP-01 | Pre-gen | Routing | Choose generator mode (tickets-grouped vs codebase-grouped vs gold) | User goal, repo state, available skills/tools | Selected generator + params |
| DP-02 | Pre-gen | Scope | Select root(s) to scan (ticket\_root/code\_root) | Working dir tree, config defaults | Root paths |
| DP-03 | Pre-gen | Scope | Decide include/exclude glob rules | Repo conventions, noise directories | Include/exclude patterns |
| DP-04 | Pre-gen | Scope | Decide whether to treat “loose” items as first-class candidates | Ticket placement quality, density | Loose-ticket policy |
| DP-05 | Pre-gen | Governance | Decide whether to require “strict schema mode” for outputs | Target environment strictness | Enforce extra validation gates |
| DP-06 | Pre-gen | Budgeting | Choose max pack size strategy (by tokens/bytes/files) | Model context limits, file sizes | Sharding plan |
| DP-07 | Pre-gen | Budgeting | Choose per-pack cap: number of tickets/files | Expected coherence vs coverage | Caps per pack |
| DP-08 | Pre-gen | Normalization | Decide ticket/file canonical title extraction rule | Noisy headings, filenames | Canonical titles for clustering |
| DP-09 | Pre-gen | Normalization | Choose text preprocessing (stopwords, stemming, code tokens) | Domain vocabulary | Feature extraction policy |
| DP-10 | Pre-gen | Dedup | Decide duplicate threshold policy (strict vs lenient) | Similarity signals | Threshold values |
| DP-11 | Pre-gen | Dedup | Decide duplicate merge strategy (merge vs link vs pick canonical) | Diff size, recency, metadata | Merge plan + canonical selection |
| DP-12 | Pre-gen | Dedup | Decide what “differences” to preserve when merging duplicates | Patch-like deltas | Delta retention format |
| DP-13 | Pre-gen | Clustering | Decide whether to use hierarchical topic taxonomy vs flat groups | Repo size, domain diversity | Taxonomy mode |
| DP-14 | Pre-gen | Clustering | Decide group count target | Ticket count, entropy | K groups target |
| DP-15 | Pre-gen | Clustering | Decide grouping algorithm choice (heuristic vs LLM-labeled vs hybrid) | Determinism needs, accuracy | Algorithm selection |
| DP-16 | Pre-gen | Clustering | Decide “ambiguous” assignment policy (multi-home vs best-fit) | Similarity ties | Assignment rule |
| DP-17 | Pre-gen | Clustering | Decide whether to create an “Unsorted / Needs triage” pack | Low-confidence items | Extra pack creation |
| DP-18 | Pre-gen | Naming | Generate group names optimized for human scanning | Group summaries | Group name strings |
| DP-19 | Pre-gen | Ordering | Decide pack order (by dependency, ROI, urgency, confidence) | Ticket metadata, heuristics | Pack sequence |
| DP-20 | Pre-gen | Ordering | Decide ticket order within pack (chronological vs dependency graph) | References among tickets | Ordered ticket list |
| DP-21 | Pre-gen | Context | Select “context bundle” files to attach per pack | Code references, paths mentioned | Attachment list |
| DP-22 | Pre-gen | Context | Decide whether to attach full files vs excerpts/summaries | Token budget, file size | Attachment granularity |
| DP-23 | Pre-gen | Context | Decide whether to generate a synthetic “pack brief” doc | Need for shared framing | Brief content |
| DP-24 | Pre-gen | Context | Decide whether to include prior run artifacts (outputs) as inputs | Existing out\_dir contents | Reuse plan |
| DP-25 | Pre-gen | Template | Choose template variant (tickets vs codebase vs mixed) | Pack type, audience | Template choice |
| DP-26 | Pre-gen | Template | Decide whether to inject organization standards into prompts (not pack shape) | Policy bundles, style rules | Prompt preamble content |
| DP-27 | Pre-gen | Template | Decide step allocation across subtopics (how many steps per subdomain) | Group composition | Step budget map |
| DP-28 | Pre-gen | Prompting | Choose prompt “stance” (audit-first vs implement-first vs design-first) | Risk, stage | Prompt style per step |
| DP-29 | Pre-gen | Prompting | Choose “evidence bar” (strict citations vs lightweight) | Audience and stakes | Evidence requirements |
| DP-30 | Pre-gen | Prompting | Decide per-step required outputs (files changed, diffs, JSON, etc.) | Tooling integration, downstream parser | Output spec per step |
| DP-31 | Pre-gen | Prompting | Decide if each step should be self-contained or rely on previous step outputs | Runtime environment, caching | Dependency policy |
| DP-32 | Pre-gen | Prompting | Decide whether to add “ask-user” gates for missing critical inputs | Missing file detection | Gate instructions in prompts |
| DP-33 | Pre-gen | Tooling | Choose which MCP tools to call during generation (list/validate/generate) | Tool availability, cost | Tool call plan [Model Context Protocol+1](https://modelcontextprotocol.io/specification/2025-06-18/server/tools?utm_source=chatgpt.com) |
| DP-34 | Pre-gen | Tooling | Decide oracle model/engine selection and parameters | Cost, latency, quality | `oracle` flags (model/temperature) |
| DP-35 | Pre-gen | Tooling | Decide whether to preflight `oracle` invocations (`--dry-run` summary) | Compatibility risk | Preflight on/off |
| DP-36 | Pre-gen | Compliance | Decide redaction policy for sensitive strings in prompts/attachments | Secrets risk | Redaction rules |
| DP-37 | Pre-gen | Observability | Decide trace/correlation ID scheme for packs/steps | Downstream logging needs | ID format + propagation plan |
| DP-38 | Pre-gen | Observability | Decide which metrics/log events must be emitted per stage | Debuggability requirements | Instrumentation checklist |
| DP-39 | Runtime | Execution planning | Decide run strategy (run-all vs selective) | Confidence map, cost | Step subset |
| DP-40 | Runtime | Execution planning | Decide concurrency / rate-limiting policy | Provider limits | Parallelism level |
| DP-41 | Runtime | Execution gating | Decide whether to halt on first failure vs continue collecting failures | Failure criticality | Fail-fast/continue |
| DP-42 | Runtime | Recovery | Decide retry policy per failure type (tool error vs content error) | Error classification | Retry plan |
| DP-43 | Runtime | Recovery | Decide “prompt patch” for retries (tighten constraints vs broaden context) | Failure analysis | Revised prompt text |
| DP-44 | Runtime | Recovery | Decide when to escalate to user for clarification | Low confidence / missing inputs | User question(s) |
| DP-45 | Runtime | Consistency | Decide whether to re-run earlier steps when later steps reveal contradictions | Cross-step inconsistency | Re-run selection |
| DP-46 | Runtime | Quality control | Decide acceptance criteria for a step output (format, completeness) | Output parsing/validation | Pass/fail verdict |
| DP-47 | Runtime | Quality control | Decide whether to auto-validate produced artifacts (lint/tests/validate) | Available checks | Validation commands |
| DP-48 | Runtime | Post-processing | Decide how to synthesize step outputs into a final report | Audience, required structure | Summary artifact |
| DP-49 | Runtime | Post-processing | Decide whether to generate PR/patch vs documentation-only output | Repo permissions, user intent | Output mode |
| DP-50 | Runtime | Post-processing | Decide how to resolve conflicting recommendations across steps | Conflicts detected | Resolution rationale |
| DP-51 | Runtime | Governance | Decide whether outputs meet policy/security standards before writing | Policy bundle | Block/allow + edits |
| DP-52 | Runtime | Caching | Decide whether to reuse cached groupings/packs vs regenerate | Cache keys, repo diffs | Cache decision |
| DP-53 | Runtime | Caching | Decide cache invalidation scope (one pack vs all) | Changed inputs | Invalidation plan |
| DP-54 | Runtime | Observability | Decide incident-style annotation (root cause, repro, next action) on failures | Error logs, outputs | Failure note artifact |
| DP-55 | Runtime | Observability | Decide what artifacts to persist (manifests, intermediate summaries) | Debug needs | Persist list |
| DP-56 | Runtime | Packaging | Decide how to shard outputs into “mini-packs” for follow-on runs | Token limits, dependencies | Mini-pack plan |
| DP-57 | Runtime | Packaging | Decide naming/versioning for generated packs | Date, domain, stability | Pack names + version tags |
| DP-58 | Runtime | UX | Decide “next best tool call” in MCP (validate vs list vs run vs regenerate) | Current state | Tool invocation [Model Context Protocol+1](https://modelcontextprotocol.io/docs/develop/build-server?utm_source=chatgpt.com) |
| DP-59 | Runtime | UX | Decide whether to present diffs, file lists, or narrative only | Reviewer preference | Presentation format |
| DP-60 | Runtime | Learning loop | Decide whether to extract new org heuristics from this run into a reusable profile | Repeated patterns | Proposed profile snippet |

[TRUNCATED]
```

.tickets/Publish OraclePack MCP.md
```
Parent Ticket:

* Title: Publish/distribute `oraclepack-mcp-server` to avoid long-form MCP client configuration
* Summary: Replace the hardcoded venv interpreter path in MCP client configs with a portable, short config, and optionally enable one-click installation for supported clients.
* Source:

  * Link/ID (if present) or “Not provided”: `oraclepack-op-mcp.md`
  * Original ticket excerpt (≤25 words) capturing the overall theme: “publish this so we do not have to use the long form configuration for configuring mcp clients”
* Global Constraints:

  * Eliminate reliance on an absolute venv interpreter path in MCP client configuration.
  * Preserve required env variables (`ORACLEPACK_BIN`, `ORACLEPACK_ALLOWED_ROOTS`, `ORACLEPACK_ENABLE_EXEC`) in examples.
* Global Environment:

  * Unknown
* Global Evidence:

  * Current MCP client config example (shows venv path + args + env).

Split Plan:

* Coverage Map:

  * Original item: “how do I publish this so we do not have to use the long form configuration for configuring mcp clients?”

    * Assigned Ticket ID: T1
  * Original item: Current config uses venv interpreter path: `"command": "/home/user/.../venv/bin/python"`

    * Assigned Ticket ID: T1
  * Original item: Current args: `["-m", "oraclepack_mcp_server", "--transport", "stdio"]`

    * Assigned Ticket ID: T1
  * Original item: Current env vars: `ORACLEPACK_BIN`, `ORACLEPACK_ALLOWED_ROOTS`, `ORACLEPACK_ENABLE_EXEC`

    * Assigned Ticket ID: T1
  * Original item: Option 1: “Publish a Python package so the MCP command is just a PATH executable” + `uv build`, `uv publish`, `uv tool install ...`

    * Assigned Ticket ID: T1
  * Original item: Option 1 config snippet (command becomes `oraclepack-mcp-server`, args `--transport stdio`, env preserved)

    * Assigned Ticket ID: T1
  * Original item: Note: “If you want to reduce env too, prefer absolute paths…”

    * Assigned Ticket ID: T1
  * Original item: Option 2: “No install short config: run via uvx” + config snippet using `"command": "uvx"`

    * Assigned Ticket ID: T2
  * Original item: Option 2 note: aligns with `server.json` PyPI example using `runtimeHint: "uvx"`

    * Assigned Ticket ID: T4
  * Original item: Option 3a: “Claude Desktop: ship a .mcpb bundle” + `mcpb init`, `mcpb pack` + distribute `.mcpb`

    * Assigned Ticket ID: T3
  * Original item: Option 3b: “publish to Official MCP Registry (and GitHub MCP Registry)” via `server.json` and `mcp-publisher` steps

    * Assigned Ticket ID: T4
  * Original item: Recommendation section (choose Option 1 for no venv path; Option 3 for no manual config)

    * Assigned Ticket ID: Info-only
  * Original item: Note: “standardize the executable name to match the PyPI identifier”

    * Assigned Ticket ID: T1
* Dependencies:

  * T2 depends on T1 because the `uvx` approach runs the published package name (`oraclepack-mcp-server`).
  * T4 depends on T1 because the described registry publishing path references a PyPI stdio server example.
* Split Tickets:

```ticket T1
T# Title: Publish `oraclepack-mcp-server` as a PATH executable (PyPI + uv tools) and update config example
Type: enhancement
Target Area: Distribution/packaging for MCP server (`oraclepack-mcp-server`) + MCP client config examples
Summary:
- Publish the MCP server as a Python package so MCP clients can invoke it via a normal command on PATH instead of a venv interpreter path.
- Provide the shorter MCP client configuration example that uses `command: "oraclepack-mcp-server"` and preserves required env vars.
In Scope:
- Publish steps using uv:
  - `uv build`
  - `uv publish`
- Install guidance via uv tools:
  - `uv tool install oraclepack-mcp-server`
- Update the MCP client config example to:
  - `"command": "oraclepack-mcp-server"`
  - `"args": ["--transport", "stdio"]`
  - Keep env: `ORACLEPACK_BIN`, `ORACLEPACK_ALLOWED_ROOTS`, `ORACLEPACK_ENABLE_EXEC`
- Naming guidance: standardize executable name to match the PyPI identifier (e.g., `oraclepack-mcp-server`).
- Guidance note: prefer absolute paths for env values if trying to reduce/env-stabilize in hosts with undefined working directory.
Out of Scope:
- “One-click install” packaging and registry publishing (handled in other tickets).
Current Behavior (Actual):
- MCP client config points at a venv interpreter path and runs `-m oraclepack_mcp_server`:
  - `"command": "/home/user/.../venv/bin/python"`
Expected Behavior:
- MCP client config can run a PATH command directly:
  - `"command": "oraclepack-mcp-server"`
  - No venv absolute path required.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Preserve required env variables in examples:
  - `ORACLEPACK_BIN`
  - `ORACLEPACK_ALLOWED_ROOTS`
  - `ORACLEPACK_ENABLE_EXEC`
Evidence:
- Current config snippet includes venv interpreter path and env vars:
  - `"command": "/home/user/projects/temp/oraclepack/oraclepack-mcp-server/venv/bin/python"`
  - `"args": ["-m", "oraclepack_mcp_server", "--transport", "stdio"]`
  - `ORACLEPACK_BIN`, `ORACLEPACK_ALLOWED_ROOTS`, `ORACLEPACK_ENABLE_EXEC`
Open Items / Unknowns:
- Package metadata and repository details for publishing (Unknown / Not provided).
- Desired final executable name if it differs from `oraclepack-mcp-server` (Unknown / Not provided).
Risks / Dependencies:
- Not provided
Acceptance Criteria:
- A published distribution path exists that does not require MCP clients to reference a venv interpreter path.
- Documentation/config example shows:
  - `"command": "oraclepack-mcp-server"`
  - `"args": ["--transport", "stdio"]`
  - env variables preserved as shown in the source text.
- Executable naming guidance is documented (“match the PyPI identifier”).
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “publish/distribute paths that eliminate the venv absolute path”
- “Publish a Python package so the MCP `command` is just a PATH executable”
- “uv build / uv publish … uv tool install oraclepack-mcp-server”
```

```ticket T2
T# Title: Add “no install” MCP config option using `uvx`
Type: docs
Target Area: MCP client configuration documentation/examples
Summary:
- Provide a short MCP client configuration that runs the server via `uvx` so users don’t need a pre-created local venv path in config.
- Keep required environment variables in the example config.
In Scope:
- Document the `uvx`-based MCP client config example:
  - `"command": "uvx"`
  - `"args": ["oraclepack-mcp-server", "--transport", "stdio"]`
  - env: `ORACLEPACK_BIN`, `ORACLEPACK_ALLOWED_ROOTS`, `ORACLEPACK_ENABLE_EXEC`
Out of Scope:
- Publishing to PyPI (handled in T1).
- Registry publishing via `server.json`/`mcp-publisher` (handled in T4).
Current Behavior (Actual):
- Config is “long-form” due to a venv interpreter path.
Expected Behavior:
- Users can use a short config that invokes `uvx` with the package name and stdio transport args.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Preserve required env variables in the example config.
Evidence:
- Option 2 config snippet:
  - `"command": "uvx"`
  - `"args": ["oraclepack-mcp-server", "--transport", "stdio"]`
Open Items / Unknowns:
- Whether target MCP clients/hosts support `uvx` invocation in their MCP server configuration (Unknown / Not provided).
Risks / Dependencies:
- Depends on T1 (published package name referenced by `uvx`).
Acceptance Criteria:
- Documentation includes the `uvx` config snippet exactly as described in the source text.
- Documentation explicitly retains the required env variable keys used in the source text.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “No install short config: run via `uvx`”
- `"command": "uvx", "args": ["oraclepack-mcp-server", "--transport", "stdio"]`
```

```ticket T3
T# Title: Create and distribute a `.mcpb` bundle for Claude Desktop installation
Type: enhancement
Target Area: MCP Bundle packaging for Claude Desktop distribution
Summary:
- Package the local MCP server as a `.mcpb` bundle so users can install via a UI flow in supported clients (Claude Desktop mentioned).
- Document the bundle creation commands and distribution approach.
In Scope:
- Use MCPB tooling steps as described:
  - `npm install -g @anthropic-ai/mcpb`
  - `mcpb init`
  - `mcpb pack`
- Distribute the resulting `.mcpb` (example given: GitHub Releases).
- Document that users install via Claude Desktop Extensions UI flow (per source text).
Out of Scope:
- PyPI publishing and `uv` tooling approach (handled in T1).
- Official/GitHub MCP registry publishing (handled in T4).
Current Behavior (Actual):
- Users must configure MCP clients manually with JSON.
Expected Behavior:
- Users can install the server via a `.mcpb` bundle in clients that support MCP bundles (Claude Desktop mentioned).
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Follow the described `.mcpb` workflow (init + pack).
Evidence:
- “Claude Desktop: ship a `.mcpb` bundle … mcpb init … mcpb pack … Distribute the resulting `.mcpb`”
Open Items / Unknowns:
- Bundle manifest contents and exact server entrypoints required by MCPB for this server (Unknown / Not provided).
Risks / Dependencies:
- Not provided
Acceptance Criteria:
- A `.mcpb` bundle can be produced using the documented CLI steps.
- Documentation explains how to obtain the `.mcpb` (distribution channel mentioned) and install it in Claude Desktop (Extensions UI flow mentioned).
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Claude Desktop: ship a `.mcpb` bundle”
- “mcpb init … mcpb pack”
- “Distribute the resulting `.mcpb` (e.g., GitHub Releases)”
```

```ticket T4
T# Title: Publish `server.json` via `mcp-publisher` for MCP Registry / GitHub MCP Registry discovery
Type: enhancement
Target Area: Registry publishing metadata (`server.json`) + publishing workflow (`mcp-publisher`)
Summary:
- Enable “one-click install” in supported clients by publishing a `server.json` descriptor via `mcp-publisher`, targeting the Official MCP Registry and GitHub MCP Registry (as described).
- Document the high-level publishing steps and ownership proof requirement as stated.
In Scope:
- Generate `server.json` using `mcp-publisher init`.
- Follow the publishing sequence as described:
  1) Install `mcp-publisher`
  2) `mcp-publisher init` to generate `server.json`
  3) Prove package ownership (PyPI: add `mcp-name: ...` to README)
  4) `mcp-publisher login github`
  5) `mcp-publisher publish`
- Ensure `server.json` aligns with the described PyPI stdio example capabilities (mentions `environmentVariables` and `runtimeHint: "uvx"`).
Out of Scope:
- `.mcpb` bundling (handled in T3).
- Core PyPI package publishing steps (handled in T1).
Current Behavior (Actual):
- Users must manually configure MCP clients using JSON and local paths.
Expected Behavior:
- Server is discoverable/installable via registry mechanisms described (VS Code / ecosystem via registry).
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Include the described package ownership proof metadata for PyPI (README `mcp-name: ...`).
Evidence:
- “publish a `server.json` via `mcp-publisher`”
- Step list including `mcp-publisher init`, README `mcp-name: ...`, `login github`, `publish`
- Note about PyPI example supporting `environmentVariables` and `runtimeHint: "uvx"`
Open Items / Unknowns:
- Final server identifier/name to use for `mcp-name: ...` (Unknown / Not provided).
- Which registries/clients are in scope beyond “VS Code / ecosystem” phrasing (Unknown / Not provided).
Risks / Dependencies:
- Depends on T1 if the published registry entry targets a PyPI package distribution (as described).
Acceptance Criteria:
- A `server.json` exists generated/maintained via `mcp-publisher init` per the described workflow.
- Documentation includes the stated publishing steps and the PyPI ownership proof requirement (`mcp-name: ...` in README).
- Documentation notes the described `runtimeHint: "uvx"` alignment for the PyPI stdio example.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “publish to the Official MCP Registry (and GitHub MCP Registry)”
- “publish a `server.json` via `mcp-publisher`”
- “Prove package ownership … add `mcp-name: ...` to your README”
```
```

.tickets/Skills Integration Status.md
```
Directly “integrating skills into the CLI/MCP” helps, but it does not solve the core failure mode by itself. The real problem is that the CLI/MCP is running against a moving target (skills) with no version pinning, no compatibility gates, and no automated contract testing. The fix is to make skills a versioned, discoverable runtime dependency of oraclepack, with explicit compatibility rules and a lockable install.

What you have today

* Skills already define a stable-ish invocation surface: trailing `KEY=value` args, deterministic scripts like `generate_grouped_packs.py`, and mandatory validators (`validate_pack.py`, `lint_attachments.py`).
* MCP today shells out to the Go CLI (`oraclepack validate`, `oraclepack run`) and has a clear “exec enabled” safety gate (`ORACLEPACK_ENABLE_EXEC`).
* MCP tools are schema-described and discoverable via `tools/list`; that same pattern can be used for skills-as-tools, but you still need versioning/compatibility or you’ll just move the breakage boundary. ([Model Context Protocol][1])

A solution that actually keeps oraclepack in sync with rapidly changing skills

1. Make “skills” a versioned plugin/bundle, not “whatever is on disk”

* Define a skill bundle manifest (per skill, or per bundle) with:

  * `name`, `version`
  * `oraclepack_requires` (a version range)
  * entrypoints to scripts (`generate`, `validate`, `lint`, etc.)
  * an args schema (even if you keep `KEY=value` internally)
  * a content hash (sha256 tree digest) for reproducibility
* This is the same “declare a public API + version it” principle behind SemVer: you need an explicit contract before versions mean anything. ([Semantic Versioning][2])

1. Pin skill versions per run (lockfile), and enforce compatibility at runtime

* Add a lockfile (e.g., `oraclepack.lock.json`) that records:

  * skill name + exact version (or git sha)
  * manifest digest
  * oraclepack CLI version
* On `oraclepack run` / `oraclepack skill run`, oraclepack verifies:

  * installed skill version matches the lock (or fetches it)
  * `oraclepack_requires` matches the current binary
  * otherwise: fail fast with a clear message: “skill X@Y requires oraclepack >=A,<B”
* This eliminates “skills changed under me” failures.

1. Ship an “install/update” mechanism so the CLI can stay stable while skills move fast

* CLI commands:

  * `oraclepack skill list`
  * `oraclepack skill install oraclepack-tickets-pack-grouped@1.4.2` (or `@<gitsha>`)
  * `oraclepack skill run <name> KEY=value...`
  * `oraclepack skill update <name>` (explicit, not automatic)
* Skills can be distributed as:

  * GitHub release tarballs, or
  * a Python package with entry points for discovery (clean plugin model). ([Python Packaging][3])
* Store installed bundles under `~/.oraclepack/skills/<name>/<version>/…` and always run from that immutable path.

1. Unify execution so CLI and MCP cannot drift

* Implement one “skill runner” code path:

  * CLI subcommand uses it
  * MCP tool calls the same runner
* Your MCP server already uses subprocess patterns + safety gates; add “skills” under the same gate as pack execution because skills write files and can invoke tooling.
* Expose via MCP tools:

  * `oraclepack_skill_list`
  * `oraclepack_skill_describe`
  * `oraclepack_skill_run`
    MCP tools are designed to be discoverable with explicit schemas. ([Model Context Protocol][1])

1. Add contract tests that run on every skills change

* When any skill changes, CI runs:

  * `oraclepack skill run oraclepack-tickets-pack-grouped …` against a fixture repo
  * then `oraclepack validate` on every generated pack
* This converts “skills broke runs” into “PR fails before merge.”

Why “just integrate” isn’t sufficient

* If oraclepack always runs “whatever skill checkout you happen to have,” integration simply makes the breakage happen at runtime inside oraclepack instead of outside it.
* What prevents breakage is: versioned artifacts + compatibility gating + lockfiles + contract tests.

Minimal implementation order (fastest path to stability)

1. Add skill manifest + digest + `oraclepack_requires` range.
2. Add `oraclepack skill run` that executes the existing deterministic scripts referenced in the skill (no rewriting).
3. Add lockfile support + install cache.
4. Add MCP `oraclepack_skill_*` tools that call the same runner.
5. Add CI contract tests on skills changes.

If you want, I can draft the exact manifest schema and the CLI/MCP command/tool surfaces to match what your skills already expose (`KEY=value` args, deterministic generators, validators), without forcing you to change how skills are authored.

[1]: https://modelcontextprotocol.io/specification/2025-06-18/server/tools?utm_source=chatgpt.com "Tools"
[2]: https://semver.org/?utm_source=chatgpt.com "Semantic Versioning 2.0.0 | Semantic Versioning"
[3]: https://packaging.python.org/guides/creating-and-discovering-plugins/?utm_source=chatgpt.com "Creating and discovering plugins"
```

.tickets/PRD-TUI/Oraclepack TUI Integration.md
```
Parent Ticket:

* Title: Oraclepack TUI integration for PRD Generator Project URL + ticketify PRD run flow
* Summary: Enable oraclepack to route a ticket-derived PRD artifact (`tickets_prd.md`) to a specific “PRD Generator” ChatGPT Project URL via the TUI and automated runs. Address the current failure when attempting to run `tickets_prd.md` as a pack (“no bash code block found”) by introducing supported execution paths (micro-pack and/or single-shot call), and optionally generating a richer PRD context bundle + dedicated PRD-generator pack from ticketify outputs.
* Source:

  * Link/ID: Oraclepack TUI Integration (1).md
  * Original ticket excerpt (≤25 words) capturing the overall theme: “utilize our specific prd-generator from one of many gpt project urls… add that project url… as an option in the tui.”
* Global Constraints:

  * `tickets_prd.md` is content/artifact and cannot be executed as a pack unless wrapped in a valid oraclepack structure (requires a `bash` fenced block).
  * Pack parsing requires a ` ```bash … ``` ` fenced code block; otherwise error: “invalid pack structure: no bash code block found”.
  * Avoid hardcoding ChatGPT Project URLs into packs/repos; prefer selecting/storing via the TUI URL picker/store.
  * Support “one of many project urls” including per-step targeting for PRD generation steps.
  * “Simple oracle calls” should be possible without sending entire multi-step packs.
* Global Environment:

  * Unknown
* Global Evidence:

  * Error string: `Error: invalid pack structure: no bash code block found`.
  * “Projects in ChatGPT” (Projects retain chats/files within a project). ([OpenAI Help Center][1])
  * CommonMark: fenced code blocks support an “info string” after the opening fence (language identifier). ([CommonMark Spec][2])

Split Plan:

* Coverage Map:

  1. Original item: “add that project url and send it with our oraclepack in an automated manner… option in the tui.”
     Assigned Ticket ID: T1
  2. Original item: “simple way of utilizing oracle for simpler calls… do not require entire packs being sent.”
     Assigned Ticket ID: T4
  3. Original item: “oraclepack will not allow the `tickets_prd.md`… `invalid pack structure: no bash code block found`.”
     Assigned Ticket ID: T3
  4. Original item: “Add a new entry in the URL picker… Name: `PRD Generator`… Scope: `project`… or `global`.”
     Assigned Ticket ID: T1
  5. Original item: “Set it as default (`s`)…”
     Assigned Ticket ID: T1
  6. Original item: “Headless/CI: add a CLI flag… `oraclepack run --chatgpt-url <url>` (or `--chatgpt-url-name <saved-name>`).”
     Assigned Ticket ID: T2
  7. Original item: “Using multiple project URLs… `RuntimeOverrides` supports `ChatGPTURL` and `ApplyToSteps`… missing piece is Overrides Wizard UI.”
     Assigned Ticket ID: T5
  8. Original item: “Add a new wizard step: ‘ChatGPT URL’… reuse URLPickerModel… write to `pendingOverrides.ChatGPTURL`.”
     Assigned Ticket ID: T5
  9. Original item: “Option A: ‘micro-pack’… attach `tickets_prd.md`… run `oracle` once…”
     Assigned Ticket ID: T3
  10. Original item: “Option B: add `oraclepack call`… pick URL preset… pick files… run one `oracle` invocation… bypass `internal/pack/parser.go`…”
      Assigned Ticket ID: T4
  11. Original item: “Better idea: `tickets_prd.md` artifact parsed into a valid oraclepack… sent to PRD-generator project url… add missing context as part of stage.”
      Assigned Ticket ID: T6
  12. Original item: “Generate deterministic `prd_context.md`… feature summary, prioritized requirements, user stories + AC, constraints/deps/out-of-scope/risks/open questions, keep vs rewrite.”
      Assigned Ticket ID: T6
  13. Original item: “Generate `.oraclepack/ticketify/prd-generator.pack.md`… attach `tickets_prd.md` + `prd_context.md`… `--write-output ".taskmaster/docs/final_prd.md"`.”
      Assigned Ticket ID: T6
  14. Original item: “Wire it into the TUI… toggle `[ ] Generate enhanced PRD via PRD Generator Project`… prompt pick URL… run pack… apply `RuntimeOverrides{ChatGPTURL: <picked>, ApplyToSteps: {"01": true}}`.”
      Assigned Ticket ID: T7
  15. Original item: “Static context in ChatGPT Project; dynamic context in `prd_context.md` attachment.”
      Assigned Ticket ID: T6
* Dependencies:

  * T7 depends on T6 because the TUI flow runs the generated `prd-generator.pack.md` and attaches `prd_context.md`.
* Split Tickets:

```ticket T1
T# Title:
- Add/select PRD Generator ChatGPT Project URL in TUI URL picker (project/global scope + default)

Type:
- enhancement

Target Area:
- oraclepack TUI URL picker/store (“ChatGPT Project URLs”)

Summary:
- Provide a standard way to store and select a “PRD Generator” ChatGPT Project URL from the existing URL picker/store, with project vs global scope guidance and the ability to set it as the default. This enables routing PRD-generation runs through the intended ChatGPT Project without hardcoding URLs into packs.

In Scope:
- Ensure the TUI supports creating/selecting an entry named `PRD Generator` with a ChatGPT Project URL.
- Support scope selection: `project` (repo-specific) vs `global` (shared), as described.
- Support setting the selected entry as the default (`s`) for the relevant scope.

Out of Scope:
- Not provided

Current Behavior (Actual):
- Not provided

Expected Behavior:
- User can add/select `PRD Generator` as a named URL entry.
- User can set it as default for the current workflow/repo.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Do not hardcode URLs into packs/repos (prefer picker/store selection).
- Preserve support for project vs global URL scoping.

Evidence:
- “Add a new entry in the URL picker… Name: `PRD Generator`… Scope: `project`… `global`…”
- “Set it as default (`s`)…”

Open Items / Unknowns:
- Actual PRD Generator ChatGPT Project URL value(s) (not provided).
- Whether a predefined/seeded entry is required vs user-created entry (not provided).

Risks / Dependencies:
- Not provided

Acceptance Criteria:
- [ ] TUI allows creating/selecting a `PRD Generator` URL entry.
- [ ] TUI allows choosing `project` vs `global` scope for the entry.
- [ ] TUI allows setting the entry as default (`s`) and that default is subsequently used when selected.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Add a new entry in the URL picker: … Name: `PRD Generator`”
- “Scope: `project`… `global`…”
- “Set it as default (`s`)…”
```

```ticket T2
T# Title:
- Add headless/CI CLI override for ChatGPT URL selection during runs

Type:
- enhancement

Target Area:
- oraclepack CLI run command (headless / `--no-tui` runs)

Summary:
- Add a CLI override so headless/CI executions can force a ChatGPT URL (or saved URL name) rather than relying on interactive selection. This supports automated routing to the PRD Generator project URL.

In Scope:
- Add CLI support equivalent to: `oraclepack run --chatgpt-url <url>` and/or `--chatgpt-url-name <saved-name>`.
- Ensure the provided value sets the runner’s ChatGPT URL for the run.

Out of Scope:
- Not provided

Current Behavior (Actual):
- “Right now, the TUI resolves a URL; but there isn’t a CLI flag … that forces it in `--no-tui` runs.”

Expected Behavior:
- Headless/CI runs can specify a ChatGPT URL (or saved name) and have it applied for the run.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must work without requiring the interactive TUI URL picker.
- Must not require editing packs to include URLs.

Evidence:
- “Headless/CI: add a CLI flag to override the picker”
- “Add `oraclepack run --chatgpt-url <url>` (or `--chatgpt-url-name <saved-name>`)...”

Open Items / Unknowns:
- Exact CLI flag naming/shape (both `--chatgpt-url` and `--chatgpt-url-name` are suggested; final selection not provided).

Risks / Dependencies:
- Not provided

Acceptance Criteria:
- [ ] `oraclepack run --chatgpt-url <url>` applies the provided URL for the run.
- [ ] If `--chatgpt-url-name <saved-name>` is implemented, it resolves to a stored URL and applies it for the run.
- [ ] Works in non-interactive/headless mode.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Headless/CI: add a CLI flag to override the picker”
- “Add `oraclepack run --chatgpt-url <url>` (or `--chatgpt-url-name <saved-name>`)...”
```

````ticket T3
T# Title:
- Provide a micro-pack wrapper to run PRD generation using tickets_prd.md (valid bash fence pack)

Type:
- enhancement

Target Area:
- oraclepack pack inputs (micro-pack file used to call `oracle` with `tickets_prd.md`)

Summary:
- Introduce a minimal, valid oraclepack “micro-pack” that wraps PRD generation as an executable pack step. This avoids attempting to run `tickets_prd.md` directly (which fails pack validation) while enabling attaching `tickets_prd.md` to a single `oracle` call.

In Scope:
- Provide a one-step (or 2–3 step) pack that:
  - Is valid for oraclepack parsing (contains a `bash` fenced code block).
  - Attaches `tickets_prd.md`.
  - Runs `oracle` once to generate a PRD output.
- Ensure this flow can be routed to the selected PRD Generator ChatGPT Project URL (via existing URL selection/overrides mechanism).

Out of Scope:
- Full “single-shot call” mode that bypasses pack parsing (see T4).

Current Behavior (Actual):
- Running `tickets_prd.md` as a pack fails with: “Error: invalid pack structure: no bash code block found”.

Expected Behavior:
- A micro-pack wrapper can be run successfully by oraclepack and performs the PRD-generation oracle call using `tickets_prd.md` as an attachment.

Reproduction Steps:
1) Attempt to run `tickets_prd.md` as a pack and observe “no bash code block found”.
2) Run the micro-pack wrapper pack and observe it parses and executes.

Requirements / Constraints:
- Pack must include a ` ```bash … ``` ` fenced block to be parseable.
- Must attach `tickets_prd.md` to the oracle invocation.

Evidence:
- Error: “invalid pack structure: no bash code block found”
- “Generate a 1-step pack… attach `tickets_prd.md`… run `oracle` once…”

Open Items / Unknowns:
- Final location/name of the micro-pack file (example name provided: `prd-generator-call.pack.md`).
- Exact prompt text and output path conventions for the oracle call (not fully specified).

Risks / Dependencies:
- Not provided

Acceptance Criteria:
- [ ] A micro-pack file exists and is parseable by oraclepack (contains a `bash` fenced block).
- [ ] Running the micro-pack executes a single `oracle` call that includes `tickets_prd.md` as an attachment.
- [ ] The flow does not require running `tickets_prd.md` directly as a pack.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Generate a 1-step pack… attach `tickets_prd.md`…”
- “Create a new file… valid pack… run `oracle` while attaching `tickets_prd.md`.”
- “Error: invalid pack structure: no bash code block found”
````

```ticket T4
T# Title:
- Add single-shot oracle invocation mode (CLI/TUI) that bypasses pack parsing

Type:
- enhancement

Target Area:
- oraclepack CLI + TUI (“single-call” mode)

Summary:
- Add a new execution path for “simple oracle calls” that does not require a full pack file or pack parsing. The user can select a ChatGPT URL preset, attach files such as `tickets_prd.md`, provide a prompt/template, and run exactly one `oracle` invocation.

In Scope:
- Implement a new subcommand such as `oraclepack call` (or `oraclepack oracle`) that:
  - Lets the user pick a ChatGPT URL preset.
  - Lets the user specify attachments (e.g., `tickets_prd.md`).
[TRUNCATED]
```

.tickets/PRD-TUI/PRD-generator URL routing.md
```
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
```

.tickets/actions/Enable Action Packs Dispatch.md
```
Title:

* Enable oraclepack Action Packs to dispatch to non-oracle executors (codex/gemini/task-master/tm)

Summary:

* Current oraclepack usage feels “oracle-only” because certain UX features (flag injection and overrides validation) are hard-coded to `oracle` invocations, while the user needs Action Packs that actually execute work via other agents/tools (e.g., `codex exec`, `gemini`, `task-master`, `tm`). Update the Stage-3 Action Pack generation and/or oraclepack runner logic so Action Packs can deterministically run the correct executor commands for each action item.

Background / Context:

* User concern: “oraclepack is a wrapper around `oracle`” and adding more `oracle` calls won’t implement tasks; Action Packs must run the real tools/agents used in their workflow (examples: `codex exec ...`, `tm ...`, `gemini ...`).
* Current behavior explanation: oraclepack executes shell steps but has special handling only for lines starting with `oracle` (detection regex, flag injection, and overrides validation via `oracle --dry-run summary`).
* Stage-3 Action Pack template already runs non-oracle tools (`task-master …` and `tm autopilot`) and performs guarded checks for autopilot.
* Referenced repos/tools: Gemini CLI, Claude Task Master, OpenAI Codex, steipete/oracle.
* Referenced code/assets: `oraclepack-tui.md`, `oracle_pack_and_taskify-skills.md` (not included in transcript).

Current Behavior (Actual):

* Oraclepack “nice UX features” are oracle-specific:

  * Detects invocations using a regex anchored to literal `oracle`.
  * Injects selected flags only into `oracle …` lines.
  * Validates overrides by running `oracle --dry-run summary` only for detected oracle invocations.
* Action Packs can run arbitrary shell commands, but oraclepack’s overrides/validation UX does not generalize to other tools (codex/gemini/task-master/tm).

Expected Behavior:

* Generated Action Packs can deterministically dispatch execution to the intended executor per action item (e.g., `codex exec …`, `gemini -p …`, `task-master …`, `tm …`) rather than relying on more `oracle` calls.
* Optional: oraclepack’s overrides/flag-injection UX can recognize additional command prefixes beyond `oracle` (if desired).

Requirements:

* Update the Stage-3 generator (“oraclepack-taskify” / Stage-3 Action Pack template) to support configurable tool command strings beyond existing `oracle_cmd`, `task_master_cmd`, `tm_cmd`:

  * Add `codex_cmd` (default `codex`)
  * Add `gemini_cmd` (default `gemini`)
  * Optional: add `autopilot_cmd` (default `${tm_cmd} autopilot`).
* Extend `_actions.json` schema to include an executor plan:

  * `tooling`: include `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`
  * Per action item: `executor` (`"codex" | "gemini" | "tm" | "manual"`), `exec_prompt` (deterministic; no code fences), optional `inputs` (paths/globs).
* Add a new Action Pack execution path for implementation:

  * Either add `mode=implement`, or add a new “Step 09” gated by `MODE == implement`.
  * Step reads `<out_dir>/_actions.json`, selects top N items (using existing `top_n`), then dispatches:

    * `codex exec …` when `executor == codex`
    * `gemini -p …` when `executor == gemini`.
* Safety constraint:

  * Keep defaults strict (avoid “yolo” execution by default); Gemini approval/tool execution should remain conservative unless explicitly opted in.
* Optional (nice-to-have): generalize oraclepack’s oracle-specific detection/injection:

  * Generalize `ExtractOracleInvocations` / `InjectFlags` to recognize a registry of prefixes (`oracle`, `codex`, `gemini`, `task-master`, `tm`).
  * Add per-tool override sets (so “Oracle Flags” aren’t incorrectly applied to other tools).

Out of Scope:

* Not provided.

Reproduction Steps:

* Not provided.

Environment:

* OS: Unknown
* oraclepack version/commit: Unknown
* Shell/runner context: Unknown
* Stage-3 generator version/commit: Unknown

Evidence:

* User statement of need: actionpacks must call their agents/tools (examples: `codex exec ...`, `tm ...`, `gemini ...`).
* Oraclepack oracle-specific UX behavior (regex detection, flag injection, `oracle --dry-run summary` validation).
* Proposed schema + implement mode/Step 09 dispatcher design.
* Attachment: Oraclepack Action Pack Integration.md

Decisions / Agreements:

* “Fastest path” identified in transcript: upgrade Stage-3 Action Packs to include an executor dispatch step and extend `_actions.json` with `executor` metadata; modifying oraclepack core is optional.

Open Items / Unknowns:

* Current `_actions.json` schema (exact fields/types) is not provided.
* Current Stage-3 Action Pack template structure (exact steps and modes) is not provided beyond `backlog|pipelines|autopilot`.
* Exact locations/implementations of `ExtractOracleInvocations` / `InjectFlags` in `oraclepack-tui.md` are not provided in this transcript.
* Expected non-interactive invocation patterns/flags for each tool in this project (codex/gemini/task-master/tm) beyond the examples are not provided.

Risks / Dependencies:

* Dependency on external CLIs and their execution/approval modes (especially Gemini CLI) with safety implications; defaults must remain conservative.
* If oraclepack overrides/validation stays oracle-only, users may expect those UX features to apply to non-oracle commands; requires clear separation or per-tool override support.

Acceptance Criteria:

* `_actions.json` produced by the Stage-3 generator includes:

  * `tooling` with `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`
  * Per-item `executor` and `exec_prompt` (and optional `inputs`).
* Action Pack supports `implement` execution (via `mode=implement` or Step 09) that:

  * Reads `<out_dir>/_actions.json`
  * Selects top N actions
  * Runs `codex exec …` for `executor=codex` and `gemini -p …` for `executor=gemini` deterministically.
* Defaults do not enable unrestricted/automatic tool execution without opt-in (conservative approval/safety posture).
* (Optional) oraclepack runner recognizes non-oracle prefixes for invocation detection and does not incorrectly apply oracle-specific overrides to other tools.

Priority & Severity (if inferable from text):

* Priority: Not provided
* Severity: Not provided

Labels (optional):

* enhancement
* action-pack
* executor-dispatch
* cli-integration
* oraclepack
* taskify
```

.tickets/actions/Improving Oraclepack Workflow.md
```
Title:

* Add deterministic chaining + structured outputs to oraclepack (fix prelude semantics and enable stage-3 “actions” workflow)

Summary:

* The current oraclepack workflow is a 2-stage pipeline (pack generation → oraclepack execution) but has a “disconnect” after execution: the 20 Markdown outputs are not machine-consumable for automated follow-on work, blocking a seamless next stage for actionable implementation.

    Workflow Improvement Suggestions

* Additionally, the runner executes the pack prelude and each step in separate `bash -lc` processes, so prelude-defined variables do not persist into steps, creating a mismatch between template expectations and runtime behavior.

    Workflow Improvement Suggestions

Background / Context:

* Stage 1: a Codex “skill” or Gemini CLI “command” generates a single Markdown “oracle-pack” document whose machine-critical content is a single fenced `bash` block containing 20 numbered steps that call `oracle` with `--write-output ...`.

    Workflow Improvement Suggestions

* Stage 2: oraclepack parses that Markdown by extracting the first \`\`\`bash fence and splitting steps by a step-header pattern, then runs a “prelude” and each step in separate shells, producing 20 output files.

    Workflow Improvement Suggestions

* Goal: improve workflow productivity and longer runtimes with minimal human interaction, including automatically passing the final 20 results into a next stage that yields actionable implementation steps.

    Workflow Improvement Suggestions

Current Behavior (Actual):

* Prelude variables do not persist into step execution because prelude and steps run in separate `bash -lc` processes; packs must use explicit paths instead of relying on prelude variables like `$out_dir`.

    Workflow Improvement Suggestions

* The 20 `oracle` outputs are human-readable Markdown but lack a machine-friendly handoff artifact (e.g., JSON/YAML), making automated downstream processing brittle.

    Workflow Improvement Suggestions

* Parser/run is vulnerable to “format drift” from pack generators (extra code fences, missing/incorrect headers, multiple `bash` fences—parser grabs the first).

    Workflow Improvement Suggestions

Expected Behavior:

* Pack prelude semantics should match user/template expectations (either reliably shared across steps or explicitly non-shared with enforced guidance).

    Workflow Improvement Suggestions

* After a run, oraclepack should produce a deterministic, machine-readable “handoff” that enables an immediate next stage without manual intervention (“without missing a beat”).

    Workflow Improvement Suggestions

* Pack ingestion should be resilient and self-validating (fail fast on drift and contract violations).

    Workflow Improvement Suggestions

Requirements:

* Prelude semantics (choose and make official):

    Workflow Improvement Suggestions

  * Option A: “Prelude is prep-only” (no shared vars): update pack template guidance accordingly.

  * Option B: “Prelude is sourced into every step”: implement by prefixing each step script with prelude content (or run all steps in a single long-lived shell session).

* Stage-1 → Stage-2 contract hardening (“self-healing”):

    Workflow Improvement Suggestions

  * Standardize step headers to the conservative form `# NN)`.

  * Enforce exactly one fenced `bash` block per pack (or validation error).

  * Run `oraclepack validate` immediately after pack generation (wrapper/scriptable convention).

* Add a machine-readable “run index” artifact per run:

    Workflow Improvement Suggestions

  * Must include per-step `--write-output` paths, exit codes, timestamps.

  * Include parsed metadata when available (ROI/category/reference from step header line).

* Add a first-class “chain” capability to generate an “actions” next stage:

    Workflow Improvement Suggestions

  * Proposed interface: `oraclepack chain <pack.md> --mode actions`.

  * Must synthesize: `oracle-out/_summary.md` (human) and `oracle-out/_actions.json` (machine).

  * `_actions.json` should normalize each item with at least: `category`, `roi`, `reference`, `recommended_action` (“Next smallest concrete experiment”), `missing_artifacts[]`, `risk_notes[]`.

        Workflow Improvement Suggestions

  * Emit a follow-on pack: `docs/oracle-actions-YYYY-MM-DD.md`.

        Workflow Improvement Suggestions

* Execution/runtime considerations:

  * Keep compatibility with non-interactive operation (`--no-tui`, `--run-all`) and stop-on-fail behavior so chaining can run in CI.

        Workflow Improvement Suggestions

  * Optional: opt-in parallel execution with a small concurrency cap if provider/rate limits allow.

        Workflow Improvement Suggestions

* Improve pack input/context usage:

  * Support “focus + exclusion” inputs (e.g., `focus_categories=permissions,migrations`, `exclude_paths=docs,node_modules,dist`) without changing pack shape.

        Workflow Improvement Suggestions

  * Treat “extra\_files” as a deliberate context bundle (org standards/ADRs/threat model/coding conventions) appended to commands.

        Workflow Improvement Suggestions

Out of Scope:

* Not provided

Reproduction Steps:

1. Generate an oracle-pack whose prelude defines `out_dir="..."` and steps reference `$out_dir/...`.

2. Run oraclepack on the pack.

3. Observe that step shells do not see prelude-defined variables (because each step runs in a separate `bash -lc`), requiring explicit paths instead.

    Workflow Improvement Suggestions

Environment:

* OS: Unknown

* oraclepack: Unknown version (Go wrapper around `oracle`)

    Workflow Improvement Suggestions

* Shell execution model: `bash -lc` per step (per assistant analysis)

    Workflow Improvement Suggestions

* Pack generators: Codex skill and Gemini CLI command workflows

    Workflow Improvement Suggestions

Evidence:

* “oraclepack executes **prelude and steps in separate `bash -lc` processes**, so variables defined in the prelude do **not** persist into the step shells.”

    Workflow Improvement Suggestions

* Format drift risks listed: extra code fences, missing/incorrect step headers, multiple `bash` fences (parser grabs the first).

    Workflow Improvement Suggestions

* Proposed chain command + structured artifacts: `_summary.md`, `_actions.json`, and `docs/oracle-actions-YYYY-MM-DD.md`.

    Workflow Improvement Suggestions

Decisions / Agreements:

* None explicitly finalized; two alternative fixes for prelude semantics were presented (runner fix vs template fix), but no selection recorded.

    Workflow Improvement Suggestions

Open Items / Unknowns:

* Which prelude semantic is intended as the official contract (“prep-only” vs “sourced into every step”).

    Workflow Improvement Suggestions

* Exact current parser rules (regex specifics), validation behavior, and current report/state outputs (whether a run index already exists).

    Workflow Improvement Suggestions

* Whether Task Master integration is desired as a required dependency or just an optional downstream consumer.

    Workflow Improvement Suggestions

Risks / Dependencies:

* Dependency on consistent pack formatting across multiple generators (Codex/Gemini); drift can break parsing/validation.

    Workflow Improvement Suggestions

* If parallelism is added, provider rate limits and error handling may complicate deterministic runs.

    Workflow Improvement Suggestions

* Downstream automation quality depends on having a structured index/JSON handoff rather than parsing free-form Markdown answers.

    Workflow Improvement Suggestions

Acceptance Criteria:

* Running a pack that defines variables in the prelude and references them in steps behaves according to the selected official contract:

  * If “sourced prelude”: steps can reference prelude-defined variables successfully.

  * If “prep-only”: validation or guidance prevents packs from relying on prelude vars (or template guidance is updated and enforced).

        Workflow Improvement Suggestions

* `oraclepack validate` reliably fails on packs with:

  * multiple `bash` fences,

  * missing/incorrect `# NN)` step headers (per enforced convention),

  * other contract-breaking drift conditions.

        Workflow Improvement Suggestions

* After a run, oraclepack emits a machine-readable run index that includes per-step output paths, exit codes, and timestamps.

    Workflow Improvement Suggestions

* `oraclepack chain <pack.md> --mode actions` produces:

  * `oracle-out/_summary.md`,

  * `oracle-out/_actions.json` with the specified normalized fields,

  * `docs/oracle-actions-YYYY-MM-DD.md` suitable for immediate stage-2 execution.

        Workflow Improvement Suggestions

* Chaining works in non-interactive mode and respects stop-on-fail behavior to support CI usage.

    Workflow Improvement Suggestions

Priority & Severity (if inferable from text):

* Priority: Not provided

* Severity: Not provided

Labels (optional):

* enhancement

* workflow

* automation

* cli

* parsing

* schema

* ci

---
```

.tickets/actions/Oraclepack Action Pack Integration.md
```
Parent Ticket:

* Title: Oraclepack Action Pack Integration: dispatch Action Packs to external agents/tools (Codex/Gemini/Task Master)
* Summary:

  * Current concern: oraclepack is perceived as “oracle-only,” and adding more `oracle` calls won’t implement tasks.
  * Desired outcome: Action Packs should run the correct agent/tool commands (e.g., `codex exec …`, `tm …`, `gemini …`) and be tool-agnostic in how they dispatch work.
  * Optional scope: extend oraclepack’s oracle-specific UX (flag injection / validation) to support non-`oracle` commands.
* Source:

  * Link/ID (if present) or “Not provided”: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “make it so our actionpacks will perform the correct calls to the agents… Example ‘codex exec …’, ‘tm …’, ‘gemini …’”
* Global Constraints:

  * Not provided
* Global Environment:

  * Tools referenced as available/on PATH in discussion: `oracle`, `codex`, `gemini`, `task-master`, `tm`
  * Action Pack modes referenced: `backlog|pipelines|autopilot` (and proposed `implement`)
  * Runner behavior referenced: steps execute as `bash`; stdin/TTY not attached (impacting interactive CLIs)
* Global Evidence:

  * Referenced files: `oracle_pack_and_taskify-skills.md`, `oraclepack-tui.md`
  * Referenced tool repos: `https://github.com/google-gemini/gemini-cli`, `https://github.com/eyaltoledano/claude-task-master`, `https://github.com/openai/codex`, `https://github.com/steipete/oracle`
  * Mentioned artifacts/screens: “screenshot of oraclepack consuming oraclepack-taskify artifacts”; “oracle-actions-pack-2026-01-07.md”

Split Plan:

* Coverage Map:

  * “oraclepack is a wrapper around `oracle`… can not see how… more oracle calls will help us implement” → Info-only
  * “make it so our actionpacks will perform the correct calls… ‘codex exec …’, ‘tm …’, ‘gemini …’” → T3
  * “oraclepack… special logic only for lines that start with `oracle`… detect… inject flags… overrides validation” → T5
  * “Stage-3 skill already supports… `oracle_cmd`, `task_master_cmd`, `tm_cmd`… Extend… `codex_cmd`, `gemini_cmd`, optionally `autopilot_cmd`” → T1
  * “Extend `_actions.json`… include… `tooling`… per item `executor`, `exec_prompt`, `inputs`” → T2
  * “Add an ‘implement’ mode (or a Step 09)… reads… `_actions.json`… selects top N… dispatches… `codex exec…` / `gemini -p…`” → T3
  * “Keep safety defaults strict (do not ‘yolo’ by default)” → T4
  * “interactive CLIs… may fail/hang because oraclepack doesn’t provide stdin/TTY” → T4
  * “This particular Action Pack does not call `codex`, `gemini`…” → Info-only
  * “Optional: improve oraclepack UX to ‘understand’ non-oracle commands… registry of prefixes… per-tool override sets” → T5
* Dependencies:

  * T3 depends on T2 because the proposed dispatcher reads `_actions.json` fields like `executor` / `exec_prompt`.
  * T2 depends on T1 because the proposed `_actions.json.tooling` includes `codex_cmd` / `gemini_cmd` command strings.
* Split Tickets:

```ticket T1
T# Title:
- Add Codex/Gemini command configuration to Stage-3 generator (alongside oracle/task-master/tm)

Type:
- enhancement

Target Area:
- Stage-3 generator / skill template that currently supports `oracle_cmd`, `task_master_cmd`, `tm_cmd`

Summary:
- Extend the Stage-3 generation inputs to support additional tool command strings so Action Packs can invoke Codex and Gemini explicitly.
- This is intended to mirror the existing configurable command pattern already used for `oracle` and Task Master tools.

In Scope:
- Add `codex_cmd` (default `codex`) to the generator inputs.
- Add `gemini_cmd` (default `gemini`) to the generator inputs.
- Add optional `autopilot_cmd` (default `${tm_cmd} autopilot`) to the generator inputs.
- Ensure generated Action Pack uses these command variables where relevant.

Out of Scope:
- Not provided

Current Behavior (Actual):
- Generator supports configurable tool command strings: `oracle_cmd`, `task_master_cmd`, `tm_cmd`.

Expected Behavior:
- Generator also supports configurable tool command strings: `codex_cmd`, `gemini_cmd` (and optional `autopilot_cmd`), enabling Action Packs to call these tools directly.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Preserve the existing configurable command pattern already used for `oracle_cmd`, `task_master_cmd`, `tm_cmd`.
- Defaults as stated in the ticket text (`codex`, `gemini`, `${tm_cmd} autopilot`).

Evidence:
- References to existing pattern: `oracle_cmd`, `task_master_cmd`, `tm_cmd`
- Proposed additions: `codex_cmd`, `gemini_cmd`, `autopilot_cmd`

Open Items / Unknowns:
- Where the Stage-3 generator defines/declares these args (exact file/path not provided).
- How generated artifacts currently surface tool command strings (exact schema not provided).

Risks / Dependencies:
- Depends on T2 if these command strings must also be emitted into `_actions.json.tooling`.

Acceptance Criteria:
- Stage-3 generator accepts `codex_cmd` and `gemini_cmd` (and optional `autopilot_cmd`) as inputs.
- Defaults match the ticket text when values are not provided.
- Existing inputs (`oracle_cmd`, `task_master_cmd`, `tm_cmd`) remain supported and unchanged.
- Generated Action Pack artifacts include/use these variables (where the template expects tool invocations).

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Your Stage-3 skill already supports… `oracle_cmd`, `task_master_cmd`, `tm_cmd`”
- “Extend that same pattern for: `codex_cmd`… `gemini_cmd`… optionally `autopilot_cmd`”
```

```ticket T2
T# Title:
- Extend `_actions.json` to include per-item executor metadata and tool command strings

Type:
- enhancement

Target Area:
- Canonical actions artifacts (`_actions.json` / `_actions.md`) generated from Stage-2 outputs

Summary:
- Add explicit executor metadata to each action item so downstream Action Pack steps can deterministically decide which tool to run.
- Add a `tooling` object to carry tool command strings (including Codex/Gemini) for use by the dispatcher.

In Scope:
- Extend `_actions.json` to include `tooling` with: `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`.
- For each action item, add:
  - `executor`: `"codex" | "gemini" | "tm" | "manual"`
  - `exec_prompt`: short instruction string (no code fences; deterministic)
  - `inputs`: optional list of paths/globs (as referenced by ticket text)
- Ensure `_actions.md` can be produced from the JSON after the schema extension (format details not provided in ticket).

Out of Scope:
- Not provided

Current Behavior (Actual):
- `_actions.json` exists and includes `tooling` and per-item fields such as `recommended_next_action`, `acceptance_criteria` (as described in ticket text).

Expected Behavior:
- `_actions.json` includes the new `tooling` and per-item executor fields so execution can be dispatched without guessing.
- `_actions.md` can still be generated from `_actions.json`.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- `exec_prompt` must be short and deterministic (“no code fences” stated in ticket text).
- `inputs` is optional and can be derived from “missing_artifacts / repo anchors” (source details not provided).

Evidence:
- Proposed schema fields: `tooling`, `executor`, `exec_prompt`, `inputs`
- Existing per-item fields referenced: `recommended_next_action`, `acceptance_criteria`

Open Items / Unknowns:
- Exact current `_actions.json` schema and where it is defined (not provided).
- How `_actions.md` is rendered from `_actions.json` (not provided).

Risks / Dependencies:
- T3 depends on these fields being present to implement the dispatcher logic.

Acceptance Criteria:
- `_actions.json` includes `tooling` with the listed command keys.
- Each action item includes `executor` and `exec_prompt`; `inputs` is present when available and omitted/empty when not.
- The executor enum matches the ticket text: `codex|gemini|tm|manual`.
- `_actions.md` generation continues to work with the extended JSON structure.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Extend `_actions.json` to include an executor plan per item”
- “Add fields like… `tooling`: `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`”
- “per item… `executor`… `exec_prompt`… `inputs`”
```

```ticket T3
T# Title:
- Add `implement` dispatch mode (or Step 09) to Action Pack template to run Codex/Gemini based on `_actions.json`

Type:
- enhancement

Target Area:
- Stage-3 Action Pack template (modes: `backlog|pipelines|autopilot`) and step execution flow

Summary:
- Introduce an execution path that reads `_actions.json`, selects the top N actions, and dispatches to the appropriate tool based on `executor`.
- This is intended to make Action Packs perform actual implementation calls (e.g., `codex exec …`, `gemini -p …`) instead of only producing planning artifacts.

In Scope:
- Add either:
  - `mode=implement`, or
  - a new Step 09 gated by `MODE == implement`.
- In that mode/step:
  - Read `<out_dir>/_actions.json`.
  - Select the top N items (based on existing `top_n` concept referenced in ticket text).
  - Dispatch:
    - `codex exec …` for items with executor `codex`
    - `gemini -p …` for items with executor `gemini`
  - (Other executor values referenced in ticket text: `tm`, `manual` — behavior not specified beyond inclusion in enum.)

Out of Scope:
- Not provided

Current Behavior (Actual):
- Action Pack modes exist: `backlog|pipelines|autopilot`.
- Current example pack “does not call `codex`, `gemini`, etc. at all” (per ticket text).

Expected Behavior:
- When `MODE=implement`, the Action Pack executes tool commands for actions according to `_actions.json.executor`.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Use `_actions.json` as the source of truth for executor decisions (as described).
- Top-N selection should use the existing `top_n` referenced in ticket text.

Evidence:
- Proposed mode/step: “Add an ‘implement’ mode (or a Step 09)”
- Proposed dispatch: “`codex exec …`… `gemini -p …`”
- Top-N: “selects the top N items (you already have `top_n`)”

Open Items / Unknowns:
- Where `top_n` is currently defined and how ranking is determined (not provided).
- Exact location/structure for `<out_dir>/_actions.json` (not provided).

Risks / Dependencies:
- Depends on T2 because dispatcher requires `executor`/`exec_prompt` data in `_actions.json`.
- Depends on T1 if dispatcher uses `codex_cmd` / `gemini_cmd` variables.

Acceptance Criteria:
- Action Pack supports a distinct execution path for `implement` (as a mode or gated step).
- In `implement`, the Action Pack reads `_actions.json` and dispatches at least `codex` and `gemini` executors using the stated command forms.
- Existing modes (`backlog|pipelines|autopilot`) remain unchanged in behavior.
- If `_actions.json` is missing/unreadable, the implement path fails clearly (exact messaging not provided).

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Add an ‘implement’ mode (or a Step 09)… reads `<out_dir>/_actions.json`”
- “dispatches deterministically: `codex exec …`… `gemini -p …`”
- “This particular Action Pack does not call `codex`, `gemini`, etc. at all”
```

```ticket T4
T# Title:
- Add non-interactive and safety guardrails for Codex/Gemini execution in Action Packs

Type:
- enhancement

Target Area:
- Action Pack execution behavior when running non-`oracle` tools (Codex/Gemini), including runner constraints around stdin/TTY

Summary:
- The ticket notes that oraclepack runs step scripts without attaching stdin/TTY, which can break or hang interactive CLIs.
- Add guardrails so Action Packs use non-interactive invocation patterns and keep defaults safe (explicitly: do not “yolo” by default).

In Scope:
- Ensure Codex invocation is non-interactive (ticket references `codex exec …` as the intended entrypoint).
- Ensure Gemini invocation is non-interactive (ticket references `gemini -p/--prompt` as intended).
- Apply safety defaults: “Keep safety defaults strict (do not ‘yolo’ by default)” (flag specifics not mandated beyond this phrase).
[TRUNCATED]
```

.tickets/actions/Oraclepack Action Pack Issue.md
```
Parent Ticket:

* Title: Oraclepack Action Pack compatibility and non-`oracle` command handling clarity
* Summary: Clarify what happens when running `oraclepack-taskify`-generated Action Packs in current oraclepack, especially the difference between “executing the pack” vs “dispatching/wrapping non-`oracle` commands.” Capture current limitations: oraclepack’s special handling (flag injection + override validation) targets `oracle` invocations only; other CLIs run as plain shell commands and may fail/block depending on PATH and interactivity.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “Two meanings of ‘run these actionpack artifacts’… execute the pack file vs dispatch non-`oracle` commands.”
* Global Constraints:

  * Action Pack is described as “oraclepack-ingestible” with “single ```bash fence” and “# NN) step headers.”
* Global Environment:

  * Unknown
* Global Evidence:

  * docs.task-master.dev link(s) referenced in ticket text
  * developers.openai.com Codex CLI link(s) referenced in ticket text
  * docs.cloud.google.com / google-gemini.github.io Gemini CLI link(s) referenced in ticket text

Split Plan:

* Coverage Map:

  * Original item: “Can our current oraclepack run these actionpack artifacts… generated from the oraclepack-taskify skill?” → Assigned Ticket ID: T1
  * Original item: “Action Pack… ‘oraclepack-ingestible’ (single ```bash fence, # NN) step headers, deterministic paths…)” → Assigned Ticket ID: T1
  * Original item: “How you’d run it… `oraclepack validate …` / `oraclepack run …`” → Assigned Ticket ID: T1
  * Original item: “Actionpacks calling other agents/tools (codex exec, tm autopilot, gemini)… will run them as long as installed and usable non-interactively” → Assigned Ticket ID: T1
  * Original item: “Why did another reviewer say otherwise?” → Assigned Ticket ID: T1
  * Original item: “oraclepack runs each step as a bash script (bash -lc <step script>)” → Assigned Ticket ID: T1
  * Original item: “injects flags into commands that begin with oracle… does not match tm/task-master/codex/gemini” → Assigned Ticket ID: T3
  * Original item: “TUI ‘override validation’ only targets oracle… steps without oracle invocations are skipped” → Assigned Ticket ID: T2
  * Original item: “If binary isn’t installed/on PATH → step fails; if CLI is interactive → it will block” → Assigned Ticket ID: T1
  * Original item: “Only way it ‘gets confused’ is if you expect oracle output text to magically invoke Codex/Gemini” → Assigned Ticket ID: T1
  * Original item: “Two meanings… execute vs dispatch/apply oraclepack overrides to non-oracle commands” → Assigned Ticket ID: T1
  * Original item: “Reviewer answered ‘No’ because broader goal is dispatcher behavior… not interpret actions” → Assigned Ticket ID: Info-only
  * Original item: “Reconciliation: both statements can be true” → Assigned Ticket ID: Info-only
* Dependencies:

  * Not provided
* Split Tickets:

````ticket T1
T# Title: Document current Action Pack execution semantics and operator expectations
Type: docs
Target Area: oraclepack documentation / runbook for Action Packs
Summary:
  Clarify what “running an Action Pack” means in current oraclepack: steps execute as shell via `bash -lc`, and any `tm`/`codex`/`gemini` lines run as plain shell commands. Document the two meanings that caused reviewer disagreement: executing the pack vs dispatching/applying oracle-specific override logic to non-`oracle` commands. Include operator-facing notes on common failure modes already described (PATH missing binaries, interactive CLIs blocking, environment guardrails for autopilot).
In Scope:
  - Explain: “oraclepack executes each step’s body as shell via `bash -lc <command>`.”
  - Explain: non-`oracle` commands (`tm`, `task-master`, `codex`, `gemini`) are executed “directly as normal shell commands.”
  - Provide the exact run examples already given (`oraclepack validate …` and `oraclepack run …`).
  - Capture the limitation: oraclepack’s oracle-specific overrides/validation do not apply to non-`oracle` commands.
  - Document noted failure/blocking conditions:
    - Missing binaries not on PATH.
    - Interactive CLIs blocking waiting for input.
    - Autopilot “fail fast” environment issues (e.g., not in git repo, dirty tree) as stated.
Out of Scope:
  - Implementing dispatcher behavior for non-`oracle` tools (not specified beyond “broader goal” mention).
Current Behavior (Actual):
  - Steps are run as `bash -lc <step script>`.
  - Non-`oracle` commands run directly; no automatic dispatch/wrapping is applied.
Expected Behavior:
  - Operators can correctly predict:
    - What will execute (literal commands in the pack).
    - What will not happen automatically (no “magic” invocation of Codex/Gemini from oracle output text).
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Preserve stated Action Pack constraints: “single ```bash fence” and “# NN) step headers” (as described).
Evidence:
  - Included links mentioned in ticket text:
    - docs.task-master.dev references
    - developers.openai.com Codex CLI references
    - Gemini CLI references (docs.cloud.google.com / google-gemini.github.io)
Open Items / Unknowns:
  - Exact location(s) where this documentation should live (Not provided).
  - Whether this should be surfaced in CLI help text vs README vs TUI (Not provided).
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - Documentation explicitly states:
    - Packs execute step bodies via `bash -lc`.
    - Non-`oracle` commands run as-is and are not routed through oracle-specific logic.
    - Common failure modes listed in the ticket text (PATH missing, interactive blocking, autopilot environment guards).
  - Documentation includes the exact run commands already provided (`oraclepack validate …`, `oraclepack run …`).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “oraclepack executes each step’s body as shell via `bash -lc <command>`.”
  - “Will oraclepack dispatch non-`oracle` commands…? No… only targets commands that start with `oracle`.”
  - “If the CLI is interactive → it will block waiting for input.”
````

```ticket T2
T# Title: Make TUI override validation behavior explicit for steps without `oracle` invocations
Type: enhancement
Target Area: TUI overrides flow / override validation messaging
Summary:
  The ticket text states that TUI “override validation” runs `oracle --dry-run summary` on detected `oracle` invocations and skips steps without `oracle` calls. Make this behavior explicit in the TUI so operators do not misinterpret skipped steps as “validated,” especially when packs include `tm`/`codex`/`gemini` commands.
In Scope:
  - Surface an explicit note/state in the overrides validation flow indicating:
    - Validation applies only to detected `oracle` invocations.
    - Steps without `oracle` invocations are skipped by this validator.
Out of Scope:
  - Adding validation implementations for `tm`/`codex`/`gemini` (not described in ticket text).
Current Behavior (Actual):
  - “Override validation… only targets `oracle` commands… Steps without oracle invocations are skipped.”
Expected Behavior:
  - TUI clearly communicates when steps are skipped (and why), avoiding operator confusion.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must preserve current behavior described (validate only `oracle` invocations) unless changed elsewhere (Not provided).
Evidence:
  - “The TUI ‘override validation’ also only targets `oracle` commands… Steps without oracle invocations are skipped…”
Open Items / Unknowns:
  - Exact UI copy/placement and which screen(s) in the TUI should show this (Not provided).
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - When override validation runs, the UI explicitly indicates:
    - It validates only `oracle` invocations (via `oracle --dry-run summary`, as stated).
    - Steps without `oracle` invocations are skipped (and shown as skipped).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “The TUI ‘override validation’ also only targets `oracle` commands… runs `oracle --dry-run summary`…”
  - “Steps without oracle invocations are skipped by that validator.”
```

```ticket T3
T# Title: Extend runtime flag-injection matching beyond `oracle` invocations (configurable prefixes)
Type: enhancement
Target Area: command rewriting / runtime override injection
Summary:
  The ticket text states oraclepack injects flags only into commands that begin with `oracle` (regex anchored to `^(\s*)(oracle)\b`) and does not match `tm`, `task-master`, `codex`, or `gemini`. Add support for matching additional command prefixes (or a configurable list) so override injection is not limited to `oracle` only, aligning with packs that include other CLIs.
In Scope:
  - Expand the “inject flags” behavior beyond `oracle`-only matching, as motivated by:
    - “does not match `tm`, `task-master`, `codex`, `gemini`, etc.”
  - Preserve anchored/prefix-based matching semantics as described (no broad substring matching implied).
Out of Scope:
  - Defining tool-specific semantics for what flags should be injected for each CLI (Not provided).
  - Implementing dispatcher logic that changes execution from “literal shell command” to “interpreted actions” (Not provided).
Current Behavior (Actual):
  - “Injects flags into commands that begin with `oracle`… regex anchored to `^(\s*)(oracle)\b`.”
  - “It does not match `tm`, `task-master`, `codex`, `gemini`, etc.”
Expected Behavior:
  - Flag injection can apply to non-`oracle` command prefixes as configured/defined (details not provided in ticket text).
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must not break execution of packs that rely on current `oracle`-only behavior (Not provided).
Evidence:
  - “injects flags into commands that begin with `oracle` (regex anchored to `^(\s*)(oracle)\b`). It does not match `tm`, `task-master`, `codex`, `gemini`, etc.”
Open Items / Unknowns:
  - Which non-`oracle` commands should be included first (list not provided beyond examples).
  - Where configuration should live (Not provided).
  - Whether injection should be opt-in per pack/step or global (Not provided).
Risks / Dependencies:
  - Risk: unintended rewriting of commands if prefix matching is overly broad (mitigate via anchored matching; exact approach not provided).
Acceptance Criteria:
  - There is a documented/configured way to include additional command prefixes for injection beyond `oracle`.
  - Existing `oracle` prefix injection continues to work unchanged.
  - Demonstrably, a command beginning with one added prefix is recognized for injection (exact flags and CLI semantics: Not provided).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “injects flags into commands that begin with `oracle` (regex is anchored to `^(\\s*)(oracle)\\b`).”
  - “It does not match `tm`, `task-master`, `codex`, `gemini`, etc.”
```
```

.tickets/actions/Oraclepack Action Packs.md
```
Parent Ticket:

* Title: Oraclepack Action Packs: tool-agnostic execution (Codex/Gemini/Task Master) instead of oracle-only flows
* Summary:

  * Current pain point: oraclepack is a wrapper around `oracle`, and “more oracle calls” won’t implement taskified work; Action Packs must run real tool commands (e.g., `codex`, `gemini`, `tm`, `task-master`).
  * Key idea: keep oracle as “planner” and make Action Packs do deterministic executor dispatch via `_actions.json` metadata and a new implement mode/step.
* Source:

  * Link/ID (if present) or “Not provided”: Uploaded file: Oraclepack Action Pack Integration.md
  * Original ticket excerpt (≤25 words) capturing the overall theme: “make it so our actionpacks will perform the correct calls… Example ‘codex exec …’, ‘tm …’, ‘gemini …’”.
* Global Constraints:

  * Action Packs should be “tool-agnostic” (dispatch to `codex`, `gemini`, `tm`, etc.).
  * `exec_prompt` should be short and deterministic (explicitly “no code fences”).
  * Safety defaults must be strict (“do not yolo by default”); conservative approval/tool execution unless opted-in.
  * Optional/“nice-to-have”: generalize oraclepack’s oracle-specific overrides UX for other command prefixes.
* Global Environment:

  * Unknown
* Global Evidence:

  * Mentioned repos/tools:

    * `https://github.com/google-gemini/gemini-cli`
    * `https://github.com/eyaltoledano/claude-task-master`
    * `https://github.com/openai/codex`
    * `https://github.com/steipete/oracle`
  * Reference docs cited/used in the ticket text:

    * Codex CLI reference (supports “codex exec”). ([OpenAI Developers][1])
    * Gemini CLI docs (tools/approval concepts referenced by the ticket). ([Gemini CLI][2])
    * Claude Task Master repository. ([GitHub][3])

Split Plan:

* Coverage Map:

  * “make it so our actionpacks will perform the correct calls… ‘codex exec …’, ‘tm …’, ‘gemini …’” → T3
  * Oraclepack has special logic only for lines starting with `oracle` (regex detection / flag injection / validation) → T5
  * “Stage-3 Action Pack template already runs non-oracle tools (`task-master …` and `tm autopilot`) … guarded branch checks” → T3
  * Stage-3 skill supports `oracle_cmd`, `task_master_cmd`, `tm_cmd` → T1
  * “Extend that same pattern for `codex_cmd`, `gemini_cmd`, optionally `autopilot_cmd`” → T1
  * “Extend `_actions.json` … `executor`, `exec_prompt` (no code fences), `inputs`, plus expanded `tooling`” → T2
  * “Add an ‘implement’ mode (or Step 09)… reads `<out_dir>/_actions.json`… selects top N (`top_n`)… dispatches” → T3
  * “Keep safety defaults strict (do not ‘yolo’ by default)” → T4
  * “Minimal changes… add args… update Prompt A… update Action Pack template” → T1, T2, T3
  * “Optional: improve oraclepack UX… registry of command prefixes… per-tool override sets” → T5
  * “Concrete command patterns… `codex exec`, `gemini -p`, Task Master pipeline commands” → Info-only
  * “If you want, I can propose the exact schema delta… Step 09 bash logic” → Info-only
* Dependencies:

  * T3 depends on T2 because the implement/dispatcher step reads `_actions.json` and needs `executor`/`exec_prompt` metadata.
  * T2 depends on T1 because `_actions.json.tooling` expansion references new tool command fields (`codex_cmd`, `gemini_cmd`, optional `autopilot_cmd`).
  * T4 depends on T3 because safety defaults/opt-ins apply to the implement/dispatcher execution path.
  * T5 is independent (optional) but may follow T3 if you want the TUI/overrides UX to apply to non-oracle commands.
* Split Tickets:

```ticket T1
T1 Title:
- Extend oraclepack-taskify Stage-3 generator to accept and propagate executor CLI command configs (codex/gemini/autopilot)

Type:
- enhancement

Target Area:
- oraclepack-taskify (Stage-3 generator inputs/args and emitted configs)

Summary:
- The Stage-3 generator already supports configurable command strings for `oracle_cmd`, `task_master_cmd`, and `tm_cmd`.
- Extend the same configuration pattern to include `codex_cmd` and `gemini_cmd`, and optionally `autopilot_cmd`, so Action Packs can invoke the intended executors without hard-coding tool names.

In Scope:
- Add generator inputs/args for:
  - `codex_cmd` (default `codex`)
  - `gemini_cmd` (default `gemini`)
  - optional `autopilot_cmd` (default `${tm_cmd} autopilot`)
- Ensure generated artifacts carry these command strings for later use by the Action Pack execution steps.

Out of Scope:
- Modifying oraclepack core/TUI behavior (handled in T5)
- Implement-mode dispatcher logic (handled in T3)

Current Behavior (Actual):
- Stage-3 skill supports configurable tool command strings:
  - `oracle_cmd`, `task_master_cmd`, `tm_cmd`

Expected Behavior:
- Stage-3 generator also supports `codex_cmd`, `gemini_cmd`, and optionally `autopilot_cmd`, using the same configuration pattern.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Preserve the existing “configurable tool command strings” pattern already used by the Stage-3 generator.
- Defaults as stated in the ticket text.

Evidence:
- References: “Your Stage-3 skill already supports… `oracle_cmd`, `task_master_cmd`, `tm_cmd`… Extend… `codex_cmd`… `gemini_cmd`… optionally `autopilot_cmd`…”

Open Items / Unknowns:
- Where these args are defined/passed in the current generator (file paths not provided).
- Whether additional executors beyond codex/gemini/tm are needed (not provided).

Risks / Dependencies:
- Depends on T2 if `_actions.json.tooling` is expanded to include these command strings.

Acceptance Criteria:
- Stage-3 generator accepts `codex_cmd` and `gemini_cmd` (and optional `autopilot_cmd`) inputs.
- Defaults match: `codex`, `gemini`, and `${tm_cmd} autopilot` (optional).
- Generated outputs expose these command strings for downstream Action Pack steps.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Your Stage-3 skill already supports configurable tool command strings: `oracle_cmd`, `task_master_cmd`, `tm_cmd`”
- “Extend that same pattern for: `codex_cmd`… `gemini_cmd`… optionally `autopilot_cmd`”
```

```ticket T2
T2 Title:
- Extend `_actions.json` schema + Prompt A to emit per-item executor metadata (executor/exec_prompt/inputs) and expanded tooling

Type:
- enhancement

Target Area:
- Canonical actions schema (`_actions.json`) and “Prompt A” (canonical actions synthesis)

Summary:
- The current actions schema has `tooling` (oracle/task-master) and per-item fields like `recommended_next_action` and `acceptance_criteria`.
- Add executor planning fields per action item so Action Packs can deterministically select and run the correct executor (`codex`, `gemini`, `tm`, or manual), and include relevant inputs.

In Scope:
- Update `_actions.json` to add:
  - `tooling`: `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`
  - per-item fields:
    - `executor`: `"codex" | "gemini" | "tm" | "manual"`
    - `exec_prompt`: short instruction string (explicitly “no code fences; deterministic”)
    - `inputs`: optional list of paths/globs (from `missing_artifacts` / repo anchors)
- Update “Prompt A” to emit the above fields for each action item.

Out of Scope:
- Implement-mode dispatcher logic that consumes `_actions.json` (handled in T3)
- oraclepack core UX changes (handled in T5)

Current Behavior (Actual):
- `_actions.json` has `tooling` (oracle/task-master) and per-item fields like `recommended_next_action`, `acceptance_criteria`.

Expected Behavior:
- `_actions.json` includes explicit executor plan per item (`executor`, `exec_prompt`, optional `inputs`) and expanded `tooling` including codex/gemini command strings.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- `exec_prompt` must be short and deterministic; explicitly “no code fences”.
- Executor enum values and tooling fields as written in the ticket text.

Evidence:
- References: “Extend `_actions.json` to include an executor plan per item… `executor`… `exec_prompt`… `inputs`…”

Open Items / Unknowns:
- The exact existing `_actions.json` schema shape and where it’s validated (not provided).
- Whether `missing_artifacts` / repo anchors already exist in the schema (not provided).

Risks / Dependencies:
- Depends on T1 if `tooling` is expected to include the new `codex_cmd`/`gemini_cmd` command strings.
- Required by T3 since the dispatcher reads `_actions.json`.

Acceptance Criteria:
- `_actions.json` schema includes `tooling` with `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`.
- Each action item can include:
  - `executor` with allowed values: `codex`, `gemini`, `tm`, `manual`
  - `exec_prompt` (no code fences requirement captured)
  - optional `inputs` list
- Prompt A output includes these fields per item.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Extend `_actions.json` to include an executor plan per item”
- “`executor`: ‘codex’ | ‘gemini’ | ‘tm’ | ‘manual’”
- “`exec_prompt`: a short instruction string (no code fences; deterministic)”
```

```ticket T3
T3 Title:
- Add Action Pack “implement” mode (or Step 09) to dispatch executor commands based on `_actions.json`

Type:
- enhancement

Target Area:
- Stage-3 Action Pack template (runner-ingestible actionpack artifacts)

Summary:
- Action Packs must execute real tool commands to implement tasks; adding more `oracle` calls only produces more analysis/synthesis.
- Add an implement execution path that reads `_actions.json`, selects top N actions (`top_n`), and dispatches to the specified executor (e.g., `codex exec …`, `gemini -p …`) using per-item metadata.

In Scope:
- Add either:
  - `mode=implement`, or
  - a new Step 09 guarded by `if MODE == implement`
- Implement-mode behavior:
  - Read `<out_dir>/_actions.json`
  - Select the top N items (uses existing `top_n`)
  - Dispatch deterministically:
    - `codex exec …` for items with `executor=codex`
    - `gemini -p …` for items with `executor=gemini`
- Keep existing modes (`backlog|pipelines|autopilot`) intact.

Out of Scope:
- Changing oraclepack core logic for overrides/flag injection/validation (handled in T5)

Current Behavior (Actual):
- Implementation happens only when the Action Pack runs real tool commands (e.g., `task-master …`, `tm autopilot`).
- The Action Pack already runs non-oracle tools (`task-master …` and `tm autopilot`) and includes guarded branch checks for autopilot.

Expected Behavior:
- An implement mode/step exists that consumes `_actions.json` and runs executor-specific commands deterministically based on per-item metadata.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must read `<out_dir>/_actions.json`.
- Must use existing `top_n` to select items.
- Must dispatch based on `executor` field.
- Preserve deterministic behavior as described (stable ordering / fail-fast preflight referenced as desired properties).

Evidence:
- References: “Add an ‘implement’ mode (or a Step 09)… reads `<out_dir>/_actions.json`… selects the top N… dispatches deterministically…”

Open Items / Unknowns:
- The exact Action Pack template file path and step numbering constraints (not provided).
- How “top N” is currently computed/ordered (not provided).

Risks / Dependencies:
- Depends on T2 since implement mode reads `_actions.json` executor metadata.
- Safety defaults for tool execution are addressed in T4.

Acceptance Criteria:
- Action Pack supports `mode=implement` (or an equivalent Step 09 guarded behavior).
- Implement mode:
  - Reads `<out_dir>/_actions.json`
  - Selects top N via `top_n`
  - Runs `codex exec …` for `executor=codex`
  - Runs `gemini -p …` for `executor=gemini`
- Existing `backlog|pipelines|autopilot` behavior remains unchanged.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “implementation only happens when the Action Pack runs real tool commands”
- “Add an ‘implement’ mode (or a Step 09)… reads `<out_dir>/_actions.json`… selects the top N… dispatches deterministically”
```

```ticket T4
T4 Title:
- Enforce conservative safety defaults for executor dispatch (no auto-approve “yolo” behavior unless opted-in)

Type:
- chore

Target Area:
- Action Pack implement/dispatcher execution path (gemini/codex dispatch configuration)

Summary:
- The ticket explicitly calls out safety: executor tools may run commands or perform actions; defaults should be conservative.
[TRUNCATED]
```

.tickets/actions/Oraclepack Compatibility Issues.md
```
Parent Ticket:

* Title: Oraclepack Actionpack Compatibility: non-`oracle` tools execution, dispatcher/overrides gaps, and adding Codex/Gemini headless steps
* Summary:

  * There is confusion about whether oraclepack can run Action Packs that include non-`oracle` commands (e.g., `task-master` / `tm`, `codex`, `gemini`). The current behavior is that oraclepack executes each step as shell (`bash -lc ...`) and only applies oracle-specific injection/validation to commands that begin with `oracle`. The request also includes adding headless `gemini` + non-interactive `codex exec` automation into the placeholder steps of `ticket-action-pack.md`, and optionally extending this pattern to taskify-generated packs.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “injects flags into commands that begin with `oracle`… does not match `tm`, `codex`, `gemini`… won’t dispatch/wrap non-`oracle` commands” (per file)
* Global Constraints:

  * Action Pack is “oraclepack-ingestible” (single `bash` fence, `# NN)` steps) (per file)
  * Do not assume oraclepack overrides apply to non-`oracle` commands (per file)
* Global Environment:

  * Steps run via `bash -lc ...` in the project root; oraclepack does not change `WorkDir` to `out_dir` (per file)
  * ROI filter behavior may skip steps with no `ROI=` if a threshold > 0 is used (per file)
* Global Evidence:

  * References: `oracle_pack_and_taskify-skills.md`, `oraclepack-tui.md`, `ticket-action-pack.md` (per file)
  * Noted behaviors: oracle-only regex targeting; override validation runs `oracle --dry-run summary` on detected oracle invocations (per file)
  * Mentioned outputs: `.oraclepack/ticketify/_tickets_index.json`, `_actions.json`, `_actions.md`, `.taskmaster/docs/tickets_prd.md`, `.oraclepack/ticketify/tm-complexity.json`, `ticket-action-pack.state.json`, `ticket-action-pack.report.json` (per file)

Split Plan:

* Coverage Map:

  * “oraclepack runs each step as a bash script (`bash -lc <step script>`).” → T1
  * “injects flags into commands that begin with `oracle`… regex anchored to `^(\\s*)(oracle)\\b`.” → T2
  * “TUI override validation… runs `oracle --dry-run summary`… skips steps without oracle invocations.” → T2
  * “`tm` / `task-master` run directly… not routed through oracle.” → T1
  * “If you manually add `codex` / `gemini` lines… oraclepack will try to run them directly.” → T1
  * “If CLI isn’t installed/on PATH → step fails; if interactive → blocks.” → T3
  * “If you don’t add those commands… pack mainly uses `oracle` + Task Master; won’t ‘implement’ via Codex/Gemini.” → T1
  * `ticket-action-pack.md` Step 01 writes `.oraclepack/ticketify/_tickets_index.json` → Info-only
  * `ticket-action-pack.md` Step 02 writes `_actions.json` + `_actions.md` → Info-only
  * `ticket-action-pack.md` Step 03 writes `.taskmaster/docs/tickets_prd.md` → Info-only
  * `ticket-action-pack.md` Steps 05–07 run `task-master parse-prd`, `analyze-complexity`, `expand --all` and write `.oraclepack/ticketify/tm-complexity.json` → Info-only
  * “Steps 08–20 are placeholders/notes (echo guidance).” → T3
  * “Best insertion points… placeholder steps (09–13 and 16).” → T3
  * Step 09: Gemini headless selects next target, writes `.oraclepack/ticketify/next.json` → T3
  * Step 10: `codex exec` implements selected action, writes `.oraclepack/ticketify/codex-implement.md` → T3
  * Step 11: verification via `codex exec` and/or Gemini diff review, writes `.oraclepack/ticketify/codex-verify.md` and/or `.oraclepack/ticketify/gemini-review.json` → T3
  * Step 16: Gemini drafts PR body, writes `.oraclepack/ticketify/PR.md` → T3
  * “Optional… add an agent-mode to oraclepack-taskify packs… keep 20-step contract intact.” → T4
  * “Key constraint… overrides only target commands that begin with `oracle`; codex/gemini won’t inherit unless wrap/extend oraclepack.” → T2
  * Failure notes: missing `.tickets/`, missing `task-master` / provider keys, ROI filter gotcha → T1
* Dependencies:

  * T3 depends on T2 because codex/gemini steps will not participate in oraclepack override injection/validation unless oraclepack is extended beyond `oracle`-prefixed commands (per file).
  * T4 depends on T2 for the same reason (per file).
* Split Tickets:

```ticket T1
T# Title:
- Clarify current oraclepack Action Pack execution semantics (and common failure modes)

Type:
- docs

Target Area:
- oraclepack CLI/TUI user-facing documentation (exact file(s) not provided)

Summary:
- Document the current behavior that oraclepack executes Action Pack steps as `bash -lc ...` and only applies oracle-specific behavior to `oracle`-prefixed commands. Capture practical implications for running packs containing `task-master`/`tm`, `codex`, and `gemini`, including common failure modes and the ROI filter gotcha noted in the ticket content.

In Scope:
- Document that steps execute as shell (`bash -lc ...`) and whatever commands appear in the step body are executed.
- Document that non-`oracle` commands (`task-master`/`tm`, `codex`, `gemini`) run directly and are not routed through oracle.
- Document the “interactive CLI can block” and “missing binary on PATH fails the step” implications.
- Document `ticket-action-pack.md` likely failure points called out: missing `.tickets/`, missing `task-master`/provider configuration/API keys, ROI filter gotcha.
- Document that Steps 08–20 are placeholders unless replaced with real commands.

Out of Scope:
- Changing oraclepack dispatcher / override injection logic.
- Editing `ticket-action-pack.md` steps to add new automation.

Current Behavior (Actual):
- Confusion among reviewers/users about whether oraclepack “runs everything through oracle.”
- Running packs without `oracle ...` commands results in no oracle-specific override behavior being applied.

Expected Behavior:
- A clear, discoverable doc section explains what oraclepack does/does not do with non-`oracle` commands and how to interpret failures.

Reproduction Steps:
- Not provided.

Requirements / Constraints:
- Preserve the current semantics described in the ticket content (no implied change to execution model).

Evidence:
- “oraclepack runs each step as a bash script (`bash -lc <step script>`).”
- “If the CLI is interactive → it will block waiting for input.”
- “ROI filter gotcha… steps with no `ROI=`… may be skipped.”

Open Items / Unknowns:
- Where documentation should live (README, `oraclepack-tui.md`, or other) is not provided.
- Whether this should be shown in TUI help text vs repository docs is not provided.

Risks / Dependencies:
- Not provided.

Acceptance Criteria:
- Documentation explicitly states:
  - Steps execute via `bash -lc ...` and run the literal commands present.
  - Only `oracle`-prefixed commands receive oraclepack’s special handling.
  - Non-`oracle` tools (`tm`/`task-master`, `codex`, `gemini`) run directly (PATH/interactive caveats included).
  - The listed failure modes and ROI filter gotcha are described with practical guidance.
- Documentation includes a short “What to expect after running `ticket-action-pack.md`” section referencing the artifact paths named in the ticket content.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “oraclepack runs each step as a bash script (`bash -lc <step script>`).”
- “injects flags into commands that begin with `oracle`… does not match `tm`, `codex`, `gemini`.”
- “ROI filter gotcha… steps with no `ROI=`… may be skipped.”
```

```ticket T2
T# Title:
- Extend oraclepack override injection/validation beyond `oracle`-prefixed commands (dispatcher changes)

Type:
- enhancement

Target Area:
- oraclepack command detection + overrides/validation pipeline (regex/dispatch behavior; exact file paths not provided)

Summary:
- The current oraclepack behavior applies oracle-specific transforms only to commands that begin with `oracle` and the TUI override validation only targets oracle invocations. Implement “dispatcher changes” so that non-`oracle` commands (explicitly referenced: `tm`/`task-master`, `codex`, `gemini`) can participate in the same override/validation flow, or otherwise be handled explicitly as first-class command targets.

In Scope:
- Update command detection so it is not limited to `oracle` (currently anchored to `^(\\s*)(oracle)\\b` per the ticket text).
- Update override validation so it does not only run `oracle --dry-run summary` on detected oracle invocations and skip steps without oracle invocations.
- Ensure steps containing `tm`/`task-master`, `codex`, and/or `gemini` can be detected for dispatcher/validation purposes (as described in the ticket content).

Out of Scope:
- Adding new `codex exec` / `gemini` automation steps to specific packs (handled in T3).
- Changing Task Master’s behavior or requirements.

Current Behavior (Actual):
- Flag injection “only… injects flags into commands that begin with `oracle`… does not match `tm`, `task-master`, `codex`, `gemini`.”
- TUI override validation “only targets `oracle` commands… runs `oracle --dry-run summary`… steps without oracle invocations are skipped.”

Expected Behavior:
- Dispatcher/override handling is not limited to `oracle`-prefixed commands for the explicitly mentioned tool commands, so non-`oracle` steps are not silently excluded from override/validation.

Reproduction Steps:
- Not provided.

Requirements / Constraints:
- Must preserve existing `oracle` command behavior.
- Must address the limitation called out: overrides/validation currently only target `oracle` commands.

Evidence:
- “injects flags into commands that begin with `oracle` (regex… `^(\\s*)(oracle)\\b`).”
- “override validation… runs `oracle --dry-run summary`… steps without oracle invocations are skipped.”
- “codex/gemini won’t inherit oraclepack overrides unless you wrap them yourself or extend oraclepack.”

Open Items / Unknowns:
- Exact desired behavior for applying overrides to `tm`/`task-master`, `codex`, and `gemini` is not provided (which flags apply, how validation works).
- Whether the dispatcher should “interpret actions” vs only broaden prefix-based detection is not provided.

Risks / Dependencies:
- Risk: unclear spec for how overrides should apply to each non-`oracle` tool could lead to partial/incorrect behavior.

Acceptance Criteria:
- A pack step containing at least one of the referenced non-`oracle` command prefixes (`tm`/`task-master`, `codex`, `gemini`) is no longer automatically excluded from the override/validation pipeline solely due to not starting with `oracle`.
- Existing behavior for `oracle`-prefixed commands remains unchanged.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “injects flags into commands that begin with `oracle`… does not match `tm`, `codex`, `gemini`.”
- “override validation… only targets `oracle` commands… steps without oracle invocations are skipped.”
- “codex/gemini won’t inherit oraclepack overrides unless you wrap them yourself or extend oraclepack.”
```

```ticket T3
T# Title:
- Replace `ticket-action-pack.md` placeholder steps with headless Gemini + non-interactive Codex automation

Type:
- enhancement

Target Area:
- `ticket-action-pack.md` (ticketify Action Pack content/template; exact generator location not provided)

Summary:
- Steps 08–20 are described as placeholders/notes that only echo guidance. Replace specific placeholder steps (explicitly called out: 09–13 and 16) to add end-to-end automation using headless `gemini` and `codex exec`, producing machine-readable and human-readable artifacts under `.oraclepack/ticketify/`.

In Scope:
- Step 09: Add headless `gemini` selection that writes `.oraclepack/ticketify/next.json`.
- Step 10: Add non-interactive `codex exec` implementation that consumes `next.json` and writes `.oraclepack/ticketify/codex-implement.md`.
- Step 11: Add verification automation via `codex exec` and/or Gemini diff review:
  - `.oraclepack/ticketify/codex-verify.md` and/or `.oraclepack/ticketify/gemini-review.json`.
- Step 16: Add PR draft automation that writes `.oraclepack/ticketify/PR.md`.
[TRUNCATED]
```

.tickets/other/Oraclepack Pipeline Improvements.md
```
Title:

* Implement deterministic oraclepack pipeline improvements (strict validation, run manifests, resume/caching, Stage 3 “Actionizer”)

Summary:

* The current two-stage oraclepack workflow (Stage 1 pack generation → Stage 2 execution) is “weakly connected” and lacks deterministic handoff metadata, robust resume/retry, and an automated Stage 3 that converts 20 outputs into actionable engineering work.

    Oracle Pack Workflow Analysis

* This ticket proposes additive, backward-compatible enhancements to oraclepack and the Stage 1 generator prompts so runs are reproducible, CI-friendly, and produce machine-readable artifacts suitable for automation.

    Oracle Pack Workflow Analysis

Background / Context:

* Workflow context:

  * Stage 1: Codex skill or Gemini CLI slash command generates a single Markdown oracle question pack under `docs/*oracle-pack*.md`, following a strict oraclepack schema and containing exactly 20 `oracle ...` commands.

        Oracle Pack Workflow Analysis

  * Stage 2: oraclepack (Go wrapper around `@steipete/oracle`) executes the 20 commands and writes per-question outputs (via `--write-output`).

        Oracle Pack Workflow Analysis

  * Stage 3 is currently missing: outputs are not automatically turned into actionable implementation work.

        Oracle Pack Workflow Analysis

* Non-negotiable constraints:

  * No schema-breaking changes to the oraclepack Markdown pack schema without a backward-compatible migration path and validator-safe proof.

  * Automation must be deterministic and reproducible (no interactive steps in the critical path).

  * Stage 1 output must remain a single-pack deliverable that oraclepack can validate/run (no extra blocks/headers; no schema drift).

  * Prefer minimal file attachments per question; avoid broad globs unless unavoidable.

  * Optimize for longer runtimes with minimal human interaction (batching, resume/retry, caching, stable outputs, CI-friendly).

        Oracle Pack Workflow Analysis

Current Behavior (Actual):

* Stage 1 (generation) failure modes / friction points:

  * Packs can drift from the strict schema (extra fenced blocks, step-like headers, missing fields, wrong ordering, wrong count ≠ 20), causing ingestion/validation issues.

        Oracle Pack Workflow Analysis

  * Attachments may be bloated (broad globs, “just in case” files), increasing token cost and reducing signal-to-noise.

        Oracle Pack Workflow Analysis

  * ROI scoring can be inconsistent (unstable prioritization vs stated rationale).

        Oracle Pack Workflow Analysis

  * Coverage duplication across 20 questions (overlapping targets) wastes runs/budget.

        Oracle Pack Workflow Analysis

* Stage 2 (execution) failure modes / friction points:

  * Resume/retry semantics are weak (reruns may re-execute completed steps; partial failures require manual selection).

        Oracle Pack Workflow Analysis

  * Output determinism gaps: inconsistent output paths/slugs/out\_dir naming undermine CI diffs and Stage 3 discovery.

        Oracle Pack Workflow Analysis

  * Concurrency/rate limiting is not first-class (provider throttling/timeouts lead to nondeterministic failures).

        Oracle Pack Workflow Analysis

* Cross-stage handoff issues:

  * Missing traceability between pack ↔ outputs (no explicit manifest tying outputs to pack hash, git SHA, tool versions, provider/model settings).

        Oracle Pack Workflow Analysis

  * Stage 2 may be bypassed (pack executed “by hand”), losing wrapper state/report and consistent run directory.

        Oracle Pack Workflow Analysis

Expected Behavior:

* Stage 1 packs are always validator-safe and deterministic (single pack, exactly 20 oracle invocations, no schema drift).

    Oracle Pack Workflow Analysis

* Stage 2 produces stable, discoverable, machine-readable run artifacts that bind pack ↔ outputs and enable idempotent resume/rerun.

    Oracle Pack Workflow Analysis

* Stage 3 (“Actionizer”) exists and deterministically converts the 20 outputs into actionable engineering work artifacts (backlog + change plan + optional issue export), without duplicating work on reruns.

    Oracle Pack Workflow Analysis

* CI can run validate → run → actionize non-interactively with structured outputs and policy-driven exit codes.

    Oracle Pack Workflow Analysis

Requirements:

* Validation / linting (additive, backward-compatible):

  * Add `oraclepack validate --strict --json` that reports counts (steps=20, bash\_blocks=1, oracle\_invocations=20), ordering checks (ROI desc; ties effort asc), and required fields presence.

        Oracle Pack Workflow Analysis

* Deterministic run directory + manifest:

  * On `run`, create `.oraclepack/runs/<pack_id>/` and emit `run.json` + `steps.json`.

        Oracle Pack Workflow Analysis

  * `pack_id = YYYY-MM-DD__<gitshort>__<packhash8>`.

        Oracle Pack Workflow Analysis

  * `run.json` must include: `pack_path`, `pack_hash`, `git_sha`, `oraclepack_version`, `oracle_version`, `created_at`.

        Oracle Pack Workflow Analysis

  * `steps.json` must include: `step_id` (01..20), `reference`, `category`, `roi`, `command_hash`, `output_path`, `output_hash`, `status` (pending|ok|failed|skipped).

        Oracle Pack Workflow Analysis

* Resume/rerun semantics:

  * Make resume default: if `run.json` exists, skip steps whose output exists and matches recorded hash.

  * Support explicit overrides: `--rerun all|failed|01,03,07`.

        Oracle Pack Workflow Analysis

* Concurrency and rate limiting:

  * Add global `--max-parallel N` and optionally per-provider caps via config.

  * Implement exponential backoff + jitter on transient errors (e.g., 429/503) with a retry budget.

        Oracle Pack Workflow Analysis

* Deterministic caching (optional initially):

  * Implement caching keyed by `sha256(prompt + attached_file_hashes + oracle_flags + model)`, stored under `.oraclepack/cache/<invocation_key>.md`; rerun reuses cached outputs when key matches.

        Oracle Pack Workflow Analysis

* Stage 3 (“Actionizer”) design and artifacts:

  * Implement `oraclepack actionize --run-dir .oraclepack/runs/<pack_id>`.

  * Inputs: `run.json` + 20 outputs under `.oraclepack/runs/<pack_id>/outputs/`.

        Oracle Pack Workflow Analysis

  * Deterministic processing: normalize outputs → dedupe → categorize via fixed taxonomy → generate action tasks, including blocked/conflict handling.

        Oracle Pack Workflow Analysis

  * Outputs under `.oraclepack/runs/<pack_id>/actionizer/`:

    * `normalized.jsonl`, `backlog.md`, `change-plan.md`

    * Optional: `github-issues.json`, `taskmaster.json`.

            Oracle Pack Workflow Analysis

  * Idempotency: stable IDs derived from `pack_hash` (e.g., `OP3-<packhash8>-<issue_index>-<task_index>`), stable paths, byte-identical regeneration when inputs unchanged.

        Oracle Pack Workflow Analysis

* Stage 1 prompt/skill improvements (without breaking schema):

  * Embed structured mini-metadata inside each `-p` prompt text (not new pack headers), e.g., `QuestionId`, `Category`, `Reference`, `ExpectedArtifacts`.

        Oracle Pack Workflow Analysis

  * Enforce deterministic attachment minimization heuristics (reference file + 0–2 neighbors; avoid broad globs unless evidence demands).

        Oracle Pack Workflow Analysis

  * Standardize generator prompt across Codex skills and Gemini CLI commands using a single canonical prompt file in repo.

        Oracle Pack Workflow Analysis

* CI-native mode:

  * Provide `oraclepack run --ci --non-interactive --json-log` and `oraclepack actionize --ci`.

  * CI policy can fail build if validation fails, completion rate below threshold, or action yield below threshold (threshold values: Not provided).

        Oracle Pack Workflow Analysis

* Security/safety:

  * Path safety: prevent `--write-output` from escaping run dir (reject `..` traversal).

        Oracle Pack Workflow Analysis

Out of Scope:

* Breaking changes to the existing oraclepack Markdown pack schema (unless a backward-compatible migration path and validator-safe proof are provided).

    Oracle Pack Workflow Analysis

Reproduction Steps:

1. Generate a pack via Stage 1 and save to `docs/oracle-pack-YYYY-MM-DD.md`.

    Oracle Pack Workflow Analysis

2. Run `oraclepack validate docs/oracle-pack-YYYY-MM-DD.md` and observe schema drift / strictness gaps (exact current validator behavior: Unknown).

    Oracle Pack Workflow Analysis

3. Execute the pack, interrupt mid-run, rerun, and observe whether completed steps are skipped (current behavior: weak/unclear).

    Oracle Pack Workflow Analysis

4. Compare two runs on the same commit and observe output path/slug stability and traceability artifacts (manifest missing today).

    Oracle Pack Workflow Analysis

Environment:

* Tooling:

  * oraclepack (Go wrapper around `@steipete/oracle`).

        Oracle Pack Workflow Analysis

  * Stage 1 generators: Codex skills or Gemini CLI slash commands.

        Oracle Pack Workflow Analysis

* Repository/OS/versions: Unknown (git SHA, oraclepack version, oracle version, provider/model settings not provided in the conversation; also identified as missing traceability today).

    Oracle Pack Workflow Analysis

Evidence:

* Proposed stable artifact layout and handoff contract:

    Oracle Pack Workflow Analysis

  * `docs/oracle-pack-YYYY-MM-DD.md`

  * `.oraclepack/runs/<pack_id>/run.json`

  * `.oraclepack/runs/<pack_id>/steps.json`

  * `.oraclepack/runs/<pack_id>/outputs/01-<slug>.md … 20-<slug>.md`

  * `.oraclepack/runs/<pack_id>/actionizer/{normalized.jsonl, backlog.md, change-plan.md}`

* Proposed commands (some flag names explicitly “proposed where not already present”):

    Oracle Pack Workflow Analysis

  * `oraclepack validate --strict docs/oracle-pack-YYYY-MM-DD.md --json > .oraclepack/validate.json`

  * `oraclepack run docs/oracle-pack-YYYY-MM-DD.md --max-parallel 4 --resume --ci`

  * `oraclepack actionize --run-dir .oraclepack/runs/<pack_id> --ci`

* Example Stage 3 output record shape (JSONL line):

    Oracle Pack Workflow Analysis

Decisions / Agreements:

* Do not break the oraclepack Markdown pack schema; any change must be backward-compatible with a validator-safe proof.

    Oracle Pack Workflow Analysis

* Stage 3 (“Actionizer”) is required and should be implemented as a first-class oraclepack subcommand (`actionize`) producing deterministic artifacts with idempotent reruns.

    Oracle Pack Workflow Analysis

* Traceability and determinism should be achieved via additive sidecar files (e.g., `run.json`, `steps.json`) rather than pack schema changes.

    Oracle Pack Workflow Analysis

Open Items / Unknowns:

* Current oraclepack CLI surface area:

  * Whether `validate --strict`, `--json`, `run --ci`, `--resume`, `--json-log`, and `actionize` exist today vs need implementation (conversation notes some flags are “proposed”).

        Oracle Pack Workflow Analysis

* Current on-disk state/report formats and locations (“state lives today (intended): oraclepack state/report + per-step outputs”; exact paths not provided).

    Oracle Pack Workflow Analysis

* Threshold definitions for CI policy (“completion rate < threshold”, “action yield < threshold”): Not provided.

    Oracle Pack Workflow Analysis

* Exact strict pack schema invariants enforced today (beyond “strict output contract” and “exactly 20” requirement): Not provided in this conversation (referenced as external inputs).

    Oracle Pack Workflow Analysis

Risks / Dependencies:

* Risk: filesystem layout changes may affect existing users; mitigation is additive behavior that preserves current out\_dir behavior.

    Oracle Pack Workflow Analysis

* Risk: caching correctness depends on hashing all attached file contents; incomplete hashing risks “cache poisoning.”

    Oracle Pack Workflow Analysis

* Risk: provider throttling/timeouts require robust transient-error classification for backoff/retry behavior.

    Oracle Pack Workflow Analysis

* Dependency: Stage 3 quality depends on stable, parseable structure in per-question outputs; mitigated by deterministic normalization heuristics and improved Stage 1 prompt shaping.

Acceptance Criteria:

* Validation:

  * `oraclepack validate --strict --json` deterministically reports schema invariants (20 steps, 20 oracle invocations, schema drift detection) and can gate CI.

        Oracle Pack Workflow Analysis

* Run determinism and traceability:

  * Running a pack produces `.oraclepack/runs/<pack_id>/{run.json,steps.json,outputs/}` with stable `pack_id` and stable output naming.

  * `run.json` includes required metadata fields; `steps.json` includes required per-step fields and statuses.

        Oracle Pack Workflow Analysis

* Resume/rerun:

  * Interrupting a run mid-way and rerunning resumes without re-executing completed steps (validated via output hashes and statuses).

  * `--rerun failed|all|<step list>` behaves as specified.

        Oracle Pack Workflow Analysis

* Concurrency/rate limiting:

  * `--max-parallel N` bounds concurrency; transient failures (e.g., throttling/timeouts) are retried with backoff within a retry budget and recorded in step status.

        Oracle Pack Workflow Analysis

* Caching (if implemented):

  * Rerunning on unchanged inputs (same prompt, same attached file digests, same flags/model) results in zero provider calls and identical outputs.

[TRUNCATED]
```

.tickets/other/Oraclepack Prompt Generator.md
```
* Title: Create oraclepack-style prompt/skill generator for tickets and .tickets
* Summary:

  * Need a reusable prompt (and/or “skill” template) that can generate an oraclepack-style prompt/skill specifically for “tickets” and/or “.tickets”.
  * Must support the existing placeholder-driven wrapper pattern (e.g., `{user-idea}`, `{project-in-question}`, `{PAIN-POINT}`, `{REFERENCE-FILE}`, `{CAPABILITY}`, `{TARGET-AGENT}`, `{OPTIMIZE-PROMPT}`), including optional fields and “infer from context” behavior as described.
  * Also need guidance on what to change in the current skill and what other viable integration options exist (within the constraints already used in prior work).
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt: “prompt that can create an oraclepack prompt/skill but for our tickets and/or .tickets”
* Global Constraints:

  * Optional inputs may be omitted; proceed by inferring from context or requesting missing info within the generated prompt template (“Not always provided…”).
  * “Pain point” is optional; proceed without it if absent.
  * `{REFERENCE-FILE}` may be provided as additional constraints/spec content.
* Global Environment:

  * Unknown
* Global Evidence:

  * Existing wrapper pattern + MCP prompt/tool/resource publication precedent captured in: `/mnt/data/MCP server implementation.md`

Split Plan:

* Coverage Map:

  * Original item: “We need a prompt that can create an oraclepack prompt/skill but for our tickets and/or .tickets.”

    * Assigned Ticket ID: T2
  * Original item: “What could we do to our current skill…”

    * Assigned Ticket ID: T3
  * Original item: “…and/or what else are our options for this request?”

    * Assigned Ticket ID: T4
  * Original item: Wrapper placeholders + optionality rules (“Not always provided…”, “Our pain point…”, `{REFERENCE-FILE}`, `{TARGET-AGENT}`, `{CAPABILITY}`, `{OPTIMIZE-PROMPT}`)

    * Assigned Ticket ID: T1
  * Original item: “optimized prompt that will get the {TARGET-AGENT} to find us a solution for adding capabilities…”

    * Assigned Ticket ID: T2
* Dependencies:

  * T2 depends on T1 because the prompt/skill generator must align to the placeholder schema + optionality rules.
  * T3 depends on T2 because “current skill” changes should incorporate the finalized ticket prompt/skill template.
  * T5 depends on T2 and T3 because examples/validation need the final template and integration approach.

````ticket T1
T1 Title: Define ticket/.tickets prompt input schema and placeholder mapping
Type: docs
Target Area: Ticket input model (tickets and/or .tickets) + wrapper placeholders
Summary:
  - Define the canonical set of inputs and placeholders required to generate an oraclepack-style ticket prompt/skill.
  - Preserve the existing wrapper’s rules around optional inputs and context inference.
  - Provide a clear mapping between “tickets/.tickets” fields (if any) and wrapper placeholders without inventing unspecified fields.
In Scope:
  - Enumerate required vs optional placeholders: {user-idea}, {project-in-question}, {ADDITIONAL-INFORMATION}, {PAIN-POINT}, {REFERENCE-FILE}, {CAPABILITY}, {TARGET-AGENT}, {OPTIMIZE-PROMPT}.
  - Document handling rules explicitly stated: optional fields, “infer from context”, and behavior when pain point is absent.
  - Clarify what “tickets” vs “.tickets” means in this system using “Unknown/Not provided” where details are missing.
Out of Scope:
  - Defining new ticket fields beyond what is provided.
  - Implementing tooling or code changes (covered elsewhere).
Current Behavior (Actual):
  - Placeholder set and optionality rules exist in the wrapper pattern, but ticket/.tickets-specific mapping is not defined.
Expected Behavior:
  - A documented, stable mapping that the ticket prompt/skill generator can follow.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Do not add new placeholders or required fields beyond what is already used in the wrapper.
  - Preserve optionality rules: proceed safely when PAIN-POINT or additional info is absent.
Evidence:
  - Reference wrapper placeholders and prompt-engineer wrapper structure as precedent. (/mnt/data/MCP server implementation.md) :contentReference[oaicite:1]{index=1}
Open Items / Unknowns:
  - Exact structure/format of “tickets” and “.tickets” (not provided).
  - Whether “.tickets” is a file extension, folder convention, or schema name (not provided).
Risks / Dependencies:
  - Risk of mismatch between ticket data shape and placeholder mapping if .tickets format is not standardized.
Acceptance Criteria:
  - A single written spec exists that lists:
    - All placeholders and which are optional.
    - Rules for missing fields (“infer from context” as described).
    - How ticket/.tickets inputs populate placeholders (or explicitly “Unknown” where not provided).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Not always provided, inference from context…”
  - “Our pain point: {PAIN-POINT} … if not just continue…”
  - “```md {REFERENCE-FILE}.md”
````

```ticket T2
T2 Title: Author oraclepack-style prompt/skill template for ticket and .tickets generation
Type: enhancement
Target Area: Prompt/skill template content (oraclepack-style) for tickets/.tickets
Summary:
  - Create the actual prompt/skill template that produces an oraclepack-style prompt/skill when given a ticket or .tickets input.
  - The template must use the existing wrapper structure and placeholders, and must instruct the TARGET-AGENT to generate the desired capability for the project/tool in question.
  - Ensure the template explicitly supports optional inputs and reference-file injection as described.
In Scope:
  - Produce the “ticket prompt-engineer wrapper” template that mirrors the existing wrapper pattern but targets tickets/.tickets.
  - Include all placeholders: {user-idea}, {project-in-question}, {ADDITIONAL-INFORMATION}, {PAIN-POINT}, {REFERENCE-FILE}, {CAPABILITY}, {TARGET-AGENT}, {OPTIMIZE-PROMPT}.
  - Ensure the prompt text includes the “optimized prompt that will get the {TARGET-AGENT}…” requirement, scoped to tickets/.tickets.
Out of Scope:
  - Any new MCP server requirements, tools, or resource URI schemes not explicitly requested for tickets/.tickets.
Current Behavior (Actual):
  - There is no ticket/.tickets-specific oraclepack-style prompt/skill generator template.
Expected Behavior:
  - A single reusable prompt/skill template exists that can be filled with placeholders to drive a TARGET-AGENT to create ticket/.tickets capabilities.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must follow the wrapper’s optionality rules and placeholder usage.
  - Must treat {REFERENCE-FILE} content as “additional constraints/spec” when present.
Evidence:
  - Wrapper structure and placeholder set captured in existing reference prompt material. :contentReference[oaicite:2]{index=2}
Open Items / Unknowns:
  - Where this template will live (file path/naming) in the current repo/tooling (not provided).
Risks / Dependencies:
  - Depends on T1 for a stable placeholder-to-ticket mapping.
Acceptance Criteria:
  - Template includes:
    - All stated placeholders.
    - Explicit instruction to proceed when optional fields are missing.
    - A clearly stated “question to the prompt-engineer: {OPTIMIZE-PROMPT}” section.
  - Template is copy/paste ready and self-contained.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “create an oraclepack prompt/skill but for our tickets and/or .tickets”
  - “optimized prompt that will get the {TARGET-AGENT}… giving it {CAPABILITY}”
  - “Our question to the prompt-engineer: {OPTIMIZE-PROMPT}”
```

```ticket T3
T3 Title: Update current skill to support ticket/.tickets prompt-skill generation
Type: enhancement
Target Area: Existing “current skill” (location/name not provided)
Summary:
  - Identify changes required to the existing skill so it can generate or host the new tickets/.tickets oraclepack-style prompt/skill template.
  - Ensure the current skill can accept the ticket/.tickets inputs and populate the standardized placeholders.
  - Preserve existing behavior for non-ticket use cases (if any), since only ticket support is being added.
In Scope:
  - Incorporate the finalized template from T2 into the current skill workflow.
  - Add/adjust input handling so the current skill can be driven by “tickets and/or .tickets” as the source material.
  - Ensure optional inputs (pain point, additional information, reference file) remain optional in the workflow.
Out of Scope:
  - Designing a brand-new system if the current skill can be extended (unless extension is impossible; not provided).
Current Behavior (Actual):
  - Current skill does not explicitly support generating oraclepack-style prompts/skills for tickets/.tickets (per request).
Expected Behavior:
  - Current skill can produce the tickets/.tickets oraclepack-style prompt/skill using the same wrapper placeholder mechanism.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Do not remove or break existing skill behavior (implied by “current skill” extension request).
Evidence:
  - “What could we do to our current skill…”
Open Items / Unknowns:
  - Current skill name, file path, and execution context (not provided).
  - How tickets/.tickets are currently stored or passed into the system (not provided).
Risks / Dependencies:
  - Depends on T2 for the template content.
Acceptance Criteria:
  - Current skill supports a ticket/.tickets-driven flow that results in the T2 template with placeholders populated (or explicitly left blank when optional).
  - No regression to existing skill behaviors (validation method not provided; document what was exercised).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “What could we do to our current skill…”
  - “prompt… for our tickets and/or .tickets”
```

```ticket T4
T4 Title: Document integration options for delivering the tickets/.tickets prompt-skill capability
Type: docs
Target Area: Delivery/integration approach (within existing patterns)
Summary:
  - Provide a concise options write-up for how to deliver and reuse the tickets/.tickets prompt/skill generator, aligned to the existing approach patterns already in use.
  - Focus on the two explicitly requested dimensions: changes to the current skill and “other options” for fulfilling the request.
  - Avoid committing to new systems; frame as documented options with constraints and unknowns.
In Scope:
  - Option 1: Extend current skill to include tickets/.tickets support (ties to T3).
  - Option 2: Provide a standalone tickets/.tickets prompt/skill template artifact that can be consumed independently (ties to T2).
  - List constraints/unknowns impacting option choice (e.g., unknown .tickets format, unknown current-skill location).
Out of Scope:
  - Implementing the chosen option (handled by T3 and/or T2).
Current Behavior (Actual):
  - No documented approach exists for how tickets/.tickets prompt/skill generation should be delivered.
Expected Behavior:
  - A short decision-ready document exists describing the options and what each requires, without adding new requirements.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Options must stay within what’s already requested (modify current skill and/or alternative ways to package the prompt/skill).
Evidence:
  - “what else are our options for this request?”
Open Items / Unknowns:
  - Whether the user prefers a single consolidated skill vs multiple dedicated skills (not provided).
Risks / Dependencies:
  - Depends on T1/T2 clarity to accurately describe what each option would deliver.
Acceptance Criteria:
  - Document lists at least:
    - “Modify current skill” option (summary, prerequisites, impact).
    - “Standalone template” option (summary, prerequisites, impact).
    - Explicit unknowns that block a final choice.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “What could we do to our current skill…”
  - “…what else are our options for this request?”
```

````ticket T5
T5 Title: Add examples and validation checks for ticket/.tickets prompt-skill generation
Type: tests
Target Area: Examples + validation of generated prompt/skill output
Summary:
  - Provide concrete example inputs (ticket and/or .tickets) and the expected generated prompt/skill output shape for validation.
  - Ensure examples exercise optional fields (missing PAIN-POINT, missing ADDITIONAL-INFORMATION, with/without REFERENCE-FILE).
[TRUNCATED]
```

.tickets/other/Oraclepack Workflow Enhancement.md
```
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
```

.tickets/other/Verbose Payload Rendering TUI.md
```
Title:

* Add verbose payload rendering in TUI to display full prepared scripts/flags for oraclepack runs

Summary:

* The TUI should support a verbose mode that prints the full “prepared payload” being executed for oraclepack runs, including effective flags (post overrides and `--chatgpt-url` injection) and the entire script passed to execution.

    Verbose TUI Payload Rendering

* This is needed to verify exactly what payloads are being sent/executed during oraclepack runs and to support tests that confirm the rendered payload contents.

    Verbose TUI Payload Rendering

Background / Context:

* Proposed approach: add a reusable “prepared payload” layer to `internal/exec.Runner` (prepare prelude/step scripts after overrides + flag injection + sanitization), then have the TUI emit these prepared payload blocks to its log viewport immediately before execution.

    Verbose TUI Payload Rendering

* Files implicated by the proposal include `internal/exec/runner.go`, `internal/tui/tui.go`, `internal/cli/run.go`, plus new helpers/tests under `internal/tui/` and `internal/exec/`.

    Verbose TUI Payload Rendering

Current Behavior (Actual):

* The TUI does not provide a verbose rendering that shows the full prepared payload (full script + effective flags + extracted `oracle …` invocations) being executed for oraclepack runs.

    Verbose TUI Payload Rendering

Expected Behavior:

* When verbose payload logging is enabled, the TUI log viewport prints a payload block before each step runs that includes: effective oracle flags, extracted oracle invocations (full lines), and the full prepared script that will be executed.

    Verbose TUI Payload Rendering

* Verbose payload logging can be enabled via CLI flag (e.g., `--verbose-payload` with `-v`) and toggled in the TUI via a keybinding (proposed: `p`).

    Verbose TUI Payload Rendering

Requirements:

* Exec runner: expose “prepared payload” APIs:

  * `PreparePreludePayload(p *pack.Prelude) PreparedPreludePayload`

  * `PrepareStepPayload(s *pack.Step) PreparedStepPayload`

  * `RunPreparedPrelude(...)` / `RunPreparedStep(...)` to execute the prepared scripts.

        Verbose TUI Payload Rendering

* Prepared payload structures must include:

  * `Script` (sanitized, post injection),

  * `EffectiveFlags` (for steps),

  * `OracleInvocations` extracted from the prepared script,

  * sanitizer `Warnings`.

        Verbose TUI Payload Rendering

* TUI formatting helper:

  * Add `internal/tui/verbose_payload.go` to format payload blocks (header, effective flags, oracle invocations, then full script).

        Verbose TUI Payload Rendering

* TUI integration:

  * Add a `verbosePayload bool` toggle to the TUI model.

  * In the run flow, call `PrepareStepPayload` and, when enabled, push formatted payload lines into `logChan` before `RunPreparedStep`.

  * Add keybinding `p` to toggle `verbosePayload`.

        Verbose TUI Payload Rendering

* CLI wiring:

  * Add `--verbose-payload` / `-v` flag and pass it into `tui.NewModel(..., verbosePayload)`.

        Verbose TUI Payload Rendering

* Tests:

  * New `internal/exec/runner_payload_test.go` verifying prepared payload includes effective flags and injected oracle command text.

  * New `internal/tui/verbose_payload_test.go` verifying formatted lines include flags, invocation, and script content.

  * Update existing TUI tests to include the new `NewModel` arg.

        Verbose TUI Payload Rendering

Out of Scope:

* Not provided.

Reproduction Steps:

* Not provided.

Environment:

* Language/runtime: Go (per referenced `.go` files and packages).

    Verbose TUI Payload Rendering

* TUI framework: Bubble Tea (`tea.NewProgram(...)` referenced).

    Verbose TUI Payload Rendering

* OS / terminal / versions: Unknown.

Evidence:

* Proposed change list and implementation sketch in: /mnt/data/Verbose TUI Payload Rendering.md

    Verbose TUI Payload Rendering

* Proposed file tree changes:

  * `internal/exec/runner.go` (modify)

  * `internal/exec/runner_payload_test.go` (new)

  * `internal/tui/verbose_payload.go` (new)

  * `internal/tui/verbose_payload_test.go` (new)

  * `internal/tui/tui.go` (modify)

  * `internal/cli/run.go` (modify)

  * Update TUI tests to pass new model arg.

        Verbose TUI Payload Rendering

Decisions / Agreements:

* Adopt a “prepared payload” abstraction in `exec.Runner` to ensure the TUI logs exactly what will run after overrides, injection, and sanitization.

    Verbose TUI Payload Rendering

* Add both CLI flag control (`--verbose-payload` / `-v`) and an in-TUI toggle (proposed key: `p`).

    Verbose TUI Payload Rendering

Open Items / Unknowns:

* Exact existing TUI run flow for prelude execution (whether/where prelude runs in TUI) is not provided; proposal notes “if you also execute a prelude… do the same.”

    Verbose TUI Payload Rendering

* Exact current `NewModel(...)` signature call sites and all affected tests/files beyond those listed are not fully enumerated (some are referenced as examples).

    Verbose TUI Payload Rendering

Risks / Dependencies:

* Not provided.

Acceptance Criteria:

* Running the TUI with `--verbose-payload` causes each executed step to prepend a log block that includes:

  * “payload (step <id>)” header,

  * “effective oracle flags: …” line,

  * extracted “oracle invocations:” section (or explicit none found),

  * full “script:” content (not truncated),

  * “end payload” footer.

        Verbose TUI Payload Rendering

* Toggling `p` in the TUI flips payload logging on/off for subsequent step executions.

    Verbose TUI Payload Rendering

* `Runner.PrepareStepPayload` produces:

  * effective flags reflecting overrides and `--chatgpt-url`,

  * a prepared script containing the injected oracle invocation with those flags.

        Verbose TUI Payload Rendering

* New tests (`runner_payload_test.go`, `verbose_payload_test.go`) pass, and existing TUI tests compile and pass after updating `NewModel` call signature.

    Verbose TUI Payload Rendering

Priority & Severity (if inferable from text):

* Not provided.

Labels (optional):

* enhancement

* tui

* logging

* exec-runner

* cli

* testing

* go

---
```

.tickets/mcp/Expose Oraclepack as MCP.md
```
Parent Ticket:

* Title: Expose oraclepack as MCP tools (with Taskify Stage-2/Stage-3 helpers)
* Summary: Provide an MCP server that exposes `oraclepack` CLI capabilities (validate/list/run) plus helper tools for Stage-2 detection/validation and Stage-3 action-pack validation/execution/artifact summarization, with secure-by-default controls (allowlisted filesystem roots, execution gating, timeouts, truncation) and support for stdio + streamable-http transports.
* Source:

  * Link/ID (if present) or “Not provided”: /mnt/data/MCP tools for Oraclepack.md
  * Original ticket excerpt (≤25 words) capturing the overall theme: “Expose `oraclepack` … as **MCP tools**, so an agent can … run packs … validate Stage-2 … validate Stage-3 … summarize artifacts.”
* Global Constraints:

  * Support MCP transports: “stdio” and “streamable-http”.
  * Security defaults: “deny-by-default execution”, “allowlisted roots”, “stdout/stderr truncation and timeouts”.
* Global Environment:

  * Unknown
* Global Evidence:

  * MCP tool list and env var config list.
  * Reference implementation structure (README/requirements and Python modules).

Split Plan:

* Coverage Map:

  * “Expose `oraclepack` (validate/list/run) … as **MCP tools**” → T1
  * “run packs non-interactively (`--no-tui --yes --run-all`)” → T5
  * “validate Stage-2 outputs (01-*.md..20-*.md)” → T4
  * “validate Stage-3 Action Packs (single ```bash fence, step headers…)” → T7
  * “summarize Stage-3 artifacts (`_actions.json`, PRD, Task Master outputs, etc.)” → T7
  * “Tools: oraclepack_validate_pack / oraclepack_list_steps / oraclepack_run_pack …” → T5
  * “Tools: oraclepack_read_file …” → T5
  * “Tools: … taskify_detect_stage2 / taskify_validate_stage2 …” → T5
  * “Tools: … taskify_validate_action_pack / taskify_artifacts_summary …” → T5
  * “Tools: … taskify_run_action_pack …” → T5
  * “Transports: stdio … streamable-http …” → T6
  * “Tool annotations: readOnlyHint / destructiveHint / openWorldHint …” → T6
  * “Security defaults: ORACLEPACK_ENABLE_EXEC=1 gating …” → T2
  * “Security defaults: allowlisted filesystem roots …” → T2
  * “Security defaults: truncation and timeouts …” → T3
  * “Env vars: ORACLEPACK_ALLOWED_ROOTS / BIN / WORKDIR / ENABLE_EXEC / CHARACTER_LIMIT / MAX_READ_BYTES” → T2
  * “Typical Stage-3 usage: detect/validate → validate action pack → run → summarize” → Info-only
  * “Reference implementation tree (README, requirements.txt, modules list)” → Info-only
  * Links to MCP specs / python-sdk repo mentioned → Info-only
* Dependencies:

  * T5 depends on T2 because server tools must enforce allowed roots and execution gating.
  * T5 depends on T3 because `oraclepack_*run*` tools need subprocess execution with timeouts/truncation.
  * T5 depends on T4 because `oraclepack_taskify_*stage2*` tools call Stage-2 detection/validation.
  * T5 depends on T7 because `oraclepack_taskify_*action_pack*` tools call action-pack validation/summarization helpers.
  * T6 depends on T5 because annotations/transport-hardening apply to the MCP server surface.
* Split Tickets:

```ticket T1
T# Title: Scaffold oraclepack MCP server project (README + packaging entrypoints)
Type: chore
Target Area: oraclepack-mcp-server repo scaffolding (README.md, requirements.txt, __init__.py, __main__.py)
Summary:
- Create the MCP server project structure that exposes `oraclepack` + Taskify helpers as MCP tools, including installation and run instructions and the tool list.
- Ensure the package has an executable entrypoint to start the MCP server with selectable transport.
In Scope:
- Create/maintain project tree with:
  - README describing purpose, install, configuration env vars, run modes, tools list, and typical Stage-3 usage.
  - requirements.txt listing dependencies.
  - Python package layout with `oraclepack_mcp_server/__init__.py` and `oraclepack_mcp_server/__main__.py`.
- CLI args in `__main__.py` to accept `--transport` with choices `stdio` and `streamable-http`.
Out of Scope:
- Implementing tool logic (handled in other tickets).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- A runnable package that starts an MCP server with a chosen transport.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Must support `--transport` choices: `stdio`, `streamable-http`.
Evidence:
- Project tree and entrypoint snippet showing `choices=["stdio", "streamable-http"]`. :contentReference[oaicite:26]{index=26}
Open Items / Unknowns:
- Exact repository root / packaging approach (pip package vs repo-local module): Not provided.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Repository includes README.md, requirements.txt, and `oraclepack_mcp_server` package directory.
- [ ] Running `python -m oraclepack_mcp_server --transport stdio` is supported (starts server process).
- [ ] Running `python -m oraclepack_mcp_server --transport streamable-http` is supported (starts server process).
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “oraclepack-mcp-server (MCP wrapper for oraclepack + taskify helpers) … tree … README.md … requirements.txt … __main__.py” :contentReference[oaicite:27]{index=27}
- “choices=[‘stdio’, ‘streamable-http’] … mcp.run(transport=args.transport)” :contentReference[oaicite:28]{index=28}
```

```ticket T2
T# Title: Implement config + filesystem security controls (allowed roots, exec gating, max read bytes)
Type: chore
Target Area: oraclepack_mcp_server/config.py and oraclepack_mcp_server/security.py
Summary:
- Implement secure-by-default configuration for the MCP server, driven by environment variables, including allowlisted filesystem roots and explicit execution enablement.
- Provide safe path resolution under allowed roots and bounded file reads for tool operations.
In Scope:
- Env-driven config including:
  - `ORACLEPACK_ALLOWED_ROOTS` (colon-separated roots)
  - `ORACLEPACK_BIN`
  - `ORACLEPACK_WORKDIR`
  - `ORACLEPACK_ENABLE_EXEC`
  - `ORACLEPACK_CHARACTER_LIMIT`
  - `ORACLEPACK_MAX_READ_BYTES`
- Path allowlisting:
  - Resolve requested file paths and ensure they live under at least one allowed root.
  - Raise an explicit “path not allowed” error on violation.
- Safe file reads:
  - Read text/bytes with max size enforcement and “truncated” indicator.
Out of Scope:
- Subprocess execution and output truncation (handled in T3).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- Server loads config from env and enforces filesystem access boundaries for read tools.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Execution must be deny-by-default unless `ORACLEPACK_ENABLE_EXEC=1`.
- Filesystem access must be restricted to allowlisted roots.
Evidence:
- Env var list and semantics. :contentReference[oaicite:29]{index=29}
- Security guidance: “deny-by-default execution … allowlisted roots”. :contentReference[oaicite:30]{index=30}
Open Items / Unknowns:
- Whether Windows path separator support is required for `ORACLEPACK_ALLOWED_ROOTS`: Not provided.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Config loader reads the listed env vars and applies defaults as documented.
- [ ] Path resolution rejects paths outside allowed roots.
- [ ] File reads enforce max bytes and indicate truncation.
- [ ] Exec gating flag is available for run tools to check.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Environment variables: ORACLEPACK_ALLOWED_ROOTS … ORACLEPACK_ENABLE_EXEC … ORACLEPACK_MAX_READ_BYTES …” :contentReference[oaicite:31]{index=31}
- “Hard deny-by-default execution … Restrict filesystem access to allowlisted roots …” :contentReference[oaicite:32]{index=32}
```

```ticket T3
T# Title: Implement oraclepack subprocess runner with timeouts and stdout/stderr truncation
Type: chore
Target Area: oraclepack_mcp_server/oraclepack_cli.py (subprocess execution)
Summary:
- Provide a subprocess wrapper to invoke the `oraclepack` CLI with a hard timeout and bounded stdout/stderr capture to prevent wedging the MCP server.
- Return structured results including exit code, duration, and truncation indicators.
In Scope:
- Async subprocess execution wrapper (create subprocess, capture stdout/stderr).
- Timeout behavior:
  - Kill process on timeout and return an explicit “Timed out after {timeout}s” error result.
- Character-limit truncation for stdout/stderr based on configured limit.
Out of Scope:
- MCP tool registration (handled in T5).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- Running oraclepack commands yields deterministic, bounded outputs suitable for returning via MCP tools.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Must enforce “stdout/stderr truncation and timeouts”.
Evidence:
- Guidance: “Enforce stdout/stderr truncation and timeouts …”. :contentReference[oaicite:33]{index=33}
- Runner timeout error example snippet. :contentReference[oaicite:34]{index=34}
Open Items / Unknowns:
- Default timeout values per tool beyond examples (3600/7200) are not provided outside snippets.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Runner returns: ok, exit_code, duration_s, stdout, stderr, stdout_truncated, stderr_truncated.
- [ ] Timeout produces exit_code=124 (or equivalent) and includes a timeout message.
- [ ] Outputs are truncated to configured character limit and flags are set accordingly.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Enforce stdout/stderr truncation and timeouts so a pack can’t wedge the server process.” :contentReference[oaicite:35]{index=35}
- “Timed out after {timeout_s}s: …” return structure snippet. :contentReference[oaicite:36]{index=36}
```

```ticket T4
T# Title: Implement Taskify Stage-2 detection and validation (01-*.md..20-*.md + single-pack form)
Type: chore
Target Area: oraclepack_mcp_server/taskify.py (Stage-2 detection + validation)
Summary:
- Implement deterministic detection of Stage-2 outputs for agents, supporting both directory-form outputs (01-*.md..20-*.md) and a single-pack input file form.
- Provide validation that ensures exactly one match per prefix 01..20 and returns missing/ambiguous details.
In Scope:
- `validate_stage2_dir(out_dir)`:
  - For each prefix 01..20, glob `{pfx}-*.md`
  - Return missing patterns and ambiguous prefix matches.
- `detect_stage2(stage2_path, repo_root)`:
  - Support explicit dir path.
  - Support explicit file path with out_dir rules:
    - If under `docs/oracle-questions-YYYY-MM-DD/…`, use sibling `oracle-out` under that docs subtree; else default `repo_root/oracle-out`.
  - Support “auto” discovery (best-effort, deterministic ordering), including checking `repo_root/oracle-out`.
Out of Scope:
- Stage-3 action pack validation (handled in T7).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- Agents can resolve and validate Stage-2 outputs deterministically for downstream Stage-3 workflows.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- “Detect Stage-2 outputs (dir-form 01-*.md..20-*.md OR single-pack form)”.
- “Validate Stage-2 outputs (exactly one match per prefix 01..20)”.
Evidence:
- Requirements text for Stage-2 detection/validation. :contentReference[oaicite:37]{index=37}
- Validation logic snippet for 01..20 and ambiguous/missing. :contentReference[oaicite:38]{index=38}
- Out-dir rule snippet referencing `docs/oracle-questions-YYYY-MM-DD/…` → `oracle-out`. :contentReference[oaicite:39]{index=39}
Open Items / Unknowns:
- Full “auto discovery” search order beyond checking `repo_root/oracle-out`: Not provided in visible excerpts.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Validation returns ok=false with `missing` when any prefix has no matches.
- [ ] Validation returns ok=false with `ambiguous` when any prefix has >1 match.
- [ ] Validation returns ok=true only when exactly one match exists for every prefix 01..20.
- [ ] Detection supports explicit dir and explicit file resolution and produces an `out_dir`.
- [ ] Detection supports “auto” and returns deterministic results for the same filesystem state.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Detect Stage-2 outputs (dir-form 01-*.md..20-*.md OR single-pack form) … Validate … exactly one match per prefix 01..20.” :contentReference[oaicite:40]{index=40}
- “out_dir rules … if under docs/oracle-questions-YYYY-MM-DD/ … then …/oracle-out else oracle-out” :contentReference[oaicite:41]{index=41}
```

```ticket T5
T# Title: Implement MCP tools for oraclepack and taskify helper operations
Type: enhancement
Target Area: oraclepack_mcp_server/server.py (MCP tool registration + schemas + formatting)
Summary:
- Implement the MCP server surface that maps `oraclepack` CLI operations and Taskify helper functions into callable MCP tools.
[TRUNCATED]
```

.tickets/mcp/MCP Server for Oraclepack.md
```
Title:

* Add MCP server that exposes `oraclepack` + Taskify Stage-2/Stage-3 helpers as tools for agents

Summary:

* Agents need real-time access to `oraclepack` capabilities via MCP so they can validate, inspect, and run packs, then consume Taskify artifacts produced by the `oracle_pack_and_taskify-skills.md` workflow.
* Implement a secure-by-default Python MCP server that wraps the `oraclepack` CLI and adds deterministic helpers for Stage-2 detection/validation and Stage-3 Action Pack validation/execution + artifact summarization.

Background / Context:

* Request: “give access to agents/assistants the following oraclepack tool as a MCP tool… so they can perform actions using the artifacts generated from `oracle_pack_and_taskify-skills.md` … in real time.”
* Proposed reference implementation (from the conversation) is a Python project named `oraclepack-mcp-server` using FastMCP (MCP Python SDK), supporting `stdio` and `streamable-http` transports.
* Stage-2 outputs are expected to be 20 markdown files matching `01-*.md` … `20-*.md`; Stage-2 “single-pack form” needs out-dir resolution rules when the pack lives under `docs/oracle-questions-YYYY-MM-DD/...`.

Current Behavior (Actual):

* No MCP tool surface is available for `oraclepack` and Taskify workflows (agents cannot call validated tools to run/inspect packs and artifacts). (per user)

Expected Behavior:

* Agents can use MCP tools to:

  * Validate and inspect packs.
  * Run packs non-interactively to generate artifacts.
  * Detect/validate Stage-2 outputs and validate Stage-3 Action Packs before execution.
  * Summarize Stage-3 artifacts for immediate downstream consumption.

Requirements:

* MCP server implementation

  * Provide a Python MCP server project structure (e.g., `oraclepack-mcp-server/` with `oraclepack_mcp_server/` package).
  * Support transports:

    * `stdio`
    * `streamable-http`
* Tool surface (MCP tools)

  * `oraclepack_validate_pack`
  * `oraclepack_list_steps`
  * `oraclepack_run_pack` (execution gated)
  * `oraclepack_read_file`
  * `oraclepack_taskify_detect_stage2`
  * `oraclepack_taskify_validate_stage2`
  * `oraclepack_taskify_validate_action_pack`
  * `oraclepack_taskify_artifacts_summary`
  * `oraclepack_taskify_run_action_pack` (execution gated)
* Execution + safety controls

  * Deny-by-default execution; require `ORACLEPACK_ENABLE_EXEC=1` to enable “run” tools.
  * Restrict filesystem reads to allowlisted roots via `ORACLEPACK_ALLOWED_ROOTS` (colon-separated); block paths outside allowed roots.
  * Enforce timeouts and truncate stdout/stderr (`ORACLEPACK_CHARACTER_LIMIT`) and cap file read sizes (`ORACLEPACK_MAX_READ_BYTES`).
* Stage-2 reliability helpers

  * Validate Stage-2 directory contains exactly one match per prefix `01`..`20` (missing/ambiguous detection).
  * Deterministic Stage-2 detection:

    * Accept explicit dir or file.
    * If file is under `docs/oracle-questions-YYYY-MM-DD/...`, set out-dir to `docs/oracle-questions-YYYY-MM-DD/oracle-out`; otherwise default `repo_root/oracle-out`.
* Stage-3 reliability helpers

  * Validate Action Pack structure constraints before executing (e.g., “single ```bash fence, step headers”).
  * Summarize produced artifacts (examples cited: `_actions.json`, PRD path, Task Master outputs).
* Agent UX metadata

  * Apply MCP tool annotations:

    * validate/list/read: `readOnlyHint: true`
    * run: `destructiveHint: true`, `openWorldHint: true`

Out of Scope:

* Not provided.

Reproduction Steps:

* Not provided.

Environment:

* Language/runtime: Python (MCP server), wraps external `oraclepack` CLI.
* OS: Unknown
* Deployment: Unknown (local stdio vs HTTP service)
* MCP SDK version: Unknown (example uses `mcp>=1.0.0`, `pydantic>=2.0.0`).

Evidence:

* Conversation transcript + proposed reference implementation: `/mnt/data/MCP tools for Oraclepack.md`.
* Proposed env vars and tool list (deny-by-default exec, allowed roots, transports).
* Stage-2 directory validation (`01-*.md..20-*.md`) and Stage-2 out-dir resolution logic for `docs/oracle-questions-YYYY-MM-DD`.

Decisions / Agreements:

* Use a Python MCP server (FastMCP / MCP Python SDK) to expose `oraclepack` CLI + Taskify helpers.
* Support both `stdio` and `streamable-http` transports.
* Default-secure posture: execution gated by `ORACLEPACK_ENABLE_EXEC`, filesystem access constrained by allowlisted roots, truncation + timeout enforced.

Open Items / Unknowns:

* Where this MCP server should live (same repo as `oraclepack` vs separate repo) is not provided.
* Authentication / origin validation requirements for `streamable-http` deployment are mentioned conceptually but concrete requirements are not provided.
* Exact definition of “artifacts summary” contents/format beyond examples (`_actions.json`, PRD, Task Master outputs) is not provided.
* Whether `oraclepack_run_pack` must always use `--no-tui --yes --run-all` vs configurable flags is not provided (example suggests non-interactive flags).

Risks / Dependencies:

* Dependency on external `oraclepack` binary path/config (`ORACLEPACK_BIN`) and correct working directory (`ORACLEPACK_WORKDIR`).
* Security risk if exec is enabled without strict path/root controls and timeouts; mitigations are required as above.
* Stage-2 ambiguity risk when multiple `NN-*.md` match the same prefix; must report ambiguity deterministically and fail validation.

Acceptance Criteria:

* [ ] MCP server starts successfully in both `stdio` and `streamable-http` modes.
* [ ] All listed tools are exposed with the documented names.
* [ ] When `ORACLEPACK_ENABLE_EXEC!=1`, run tools refuse to execute and return a clear error; validate/list/read tools still work.
* [ ] `oraclepack_read_file` rejects paths outside `ORACLEPACK_ALLOWED_ROOTS`.
* [ ] Stage-2 validation enforces exactly one file per prefix `01`..`20` and returns `missing` and `ambiguous` sets when invalid.
* [ ] Stage-2 detection resolves out-dir correctly for both:

  * explicit dir inputs
  * single-pack file under `docs/oracle-questions-YYYY-MM-DD/...` → `docs/oracle-questions-YYYY-MM-DD/oracle-out`
  * otherwise default `repo_root/oracle-out`.
* [ ] stdout/stderr truncation and timeouts are enforced on CLI subprocess execution.
* [ ] Tool annotations are applied as specified for read-only vs destructive tools.

Priority & Severity (if inferable from text):

* Not provided.

Labels (optional):

* enhancement
* mcp
* oraclepack
* cli-wrapper
* taskify
* security
* tooling
```

.tickets/mcp/gaps-still-not-covered.md
```
## Gaps still not covered in the current oraclepack MCP proposal

### Transport + deployment correctness

* **`--transport streamable-http` is wired to the wrong FastMCP transport.** The proposal claims `transport="sse"` “maps to streamable-http”, but FastMCP supports Streamable HTTP explicitly and documents SSE as being superseded.  ([GitHub][1])
* **No production-grade HTTP hardening (auth, TLS posture, DNS-rebinding mitigations).** MCP security guidance explicitly warns about local HTTP servers (SSE/Streamable HTTP) without auth and recommends stdio or authenticated IPC/HTTP with mitigations. ([Model Context Protocol][2])
* **Dependency is unpinned despite a high-severity DNS rebinding advisory in the Python SDK.** The proposal uses `mcp[cli]>=0.1.0` (no minimum safe version). The advisory indicates affected versions and a patched release.  ([GitHub][3])

### Security model gaps (filesystem + execution)

* **Symlink escape is not addressed.** `validate_path()` normalizes with `abspath/normpath` and then `safe_read_file()` opens the path; this pattern typically allows “inside-root symlink → outside-root target” unless you resolve and check the realpath. No test covers symlink traversal.
* **Execution is only gated by a boolean env flag, without least-privilege scoping.** The server exposes “run pack” as open-world/destructive when enabled, but does not add per-tool scoping, allowlists, or authorization flows for HTTP mode.  ([Model Context Protocol][2])

### Parity gaps vs oraclepack TUI/runner workflows

* **No URL/project selection tooling exposed.** The TUI has explicit URL store + picker plumbing (the thing you need for “choose PRD-generator project URL”), but MCP doesn’t expose tools to list/select/manage those URLs.
* **No runtime overrides wizard parity.** The TUI supports an overrides flow (per-step flag add/remove, targeting, validation), but MCP doesn’t expose “get overrides / set overrides / validate overrides / apply and run”.
* **No structured access to run state/report artifacts.** MCP returns raw stdout/stderr strings and truncates them; it doesn’t provide first-class tools/resources for reading the oraclepack state/report outputs as structured objects.

### Execution control + UX gaps for agents

* **No step-level execution controls.** The MCP API offers `oraclepack_run_pack(... run_all=True)` but does not provide “run step N”, “run subset”, “resume”, “re-run failed only”, etc.
* **No streaming logs / cancellation.** `run_oraclepack()` waits for completion, then returns (with truncation). There’s no incremental progress, no cancellation hook, and timeout uses a hard kill only.

### Validation gaps (especially for “ticket-action-pack.md” style artifacts)

* **Action-pack validation in MCP is weaker than your existing validator script.** The MCP `validate_action_pack` logic is comparatively minimal, while the repo’s `validate-action-pack.sh` encodes stricter structural rules; MCP does not currently match those guardrails.

### Capability gaps (authoring/editing)

* **No “write/update pack” capability.** The server can read/list/validate/run packs, but cannot create/edit packs or write back transformed artifacts (which is often needed for “tickets → pack” workflows).

[1]: https://github.com/modelcontextprotocol/python-sdk?utm_source=chatgpt.com "GitHub - modelcontextprotocol/python-sdk: The official Python SDK for Model Context Protocol servers and clients"
[2]: https://modelcontextprotocol.io/specification/2025-11-25/basic/security_best_practices?utm_source=chatgpt.com "Security Best Practices - Model Context Protocol"
[3]: https://github.com/advisories/GHSA-9h52-p55h-vw2f?utm_source=chatgpt.com "Model Context Protocol (MCP) Python SDK does not enable DNS rebinding protection by default · CVE-2025-66416 · GitHub Advisory Database · GitHub"
```

.tickets/mcp/gaps_part2-mcp-builder.md
```
## Still needed to make “oraclepack as an MCP tool” production-ready

### 1) Fix transport + deployment mode choice (local vs remote)

* **Use the correct Streamable HTTP transport string** (`transport="streamable-http"`) and stop treating SSE as the Streamable HTTP equivalent. The Python SDK’s examples use `mcp.run(transport="streamable-http")`. ([GitHub][1])
* If you intend to run this as a **remote** MCP server (for multi-client / “real time” usage), implement the Streamable HTTP security requirements:

  * validate `Origin` header (DNS rebinding protection)
  * bind to `127.0.0.1` when local
  * require authentication ([Model Context Protocol][2])
* If you intend “agents/assistants” to run it **locally**, default to **stdio** and keep Streamable HTTP optional. (The MCP spec defines stdio + Streamable HTTP as the standard transports.) ([Model Context Protocol][2])

### 2) Bring `oraclepack_run_pack` up to parity with the Go CLI flags

Your current MCP tool only exposes `yes` and `run_all`.
But the CLI supports additional run-time controls (at least `--resume`, `--stop-on-fail`, ROI threshold/mode, plus the persistent `--oracle-bin` and `--out-dir`).
To avoid capability gaps (and ad-hoc “extra args” escape hatches), expose these explicitly in the tool schema.

### 3) Make Stage-2 auto-discovery match the **oraclepack-taskify** contract

The Stage 3 skill is strict about:

* deterministic discovery (lexicographic / ISO-date ordering; no mtimes)
* directory-form must contain **exactly one** `01-*.md` … `20-*.md`, else fail fast with explicit errors
  Also, the Action Pack template itself searches locations including `docs/oracle-out` and `docs/oracle-questions-*/…`.
  So the MCP-side “detect stage2” logic should:
* search the same ordered locations
* validate a candidate before returning it (not “first directory that exists”)
* prefer newest by lexicographic rules when multiple date-stamped runs exist

### 4) Tighten Action Pack validation to exactly match the skill’s validator

The skill’s validator requires:

* **exactly one** ```bash fence and **no other** fences
* sequential `# NN)` step headers inside the bash block
  If your Python validator is looser than `validate-action-pack.sh`, you’ll get drift (packs “validate” in MCP but fail in real usage).

### 5) Add “artifact-first” read tools for Stage-3 outputs (so assistants can act in real time)

Stage 3 produces canonical machine/human artifacts like:

* `<out_dir>/_actions.json`, `<out_dir>/_actions.md`
* `.taskmaster/docs/oracle-actions-prd.md`
* `<out_dir>/tm-complexity.json`
  To enable “agents/assistants” to use them immediately, add read-only tools like:
* list latest runs / outputs (by stable ordering)
* read + parse `_actions.json` (return structured JSON, not only text)
* read Task Master outputs (tasks.json location(s) you expect)

### 6) Operational hardening (especially if exec is enabled)

You already gate execution behind an env flag (`ORACLEPACK_ENABLE_EXEC`).
Still needed:

* enforce allowed roots not just for reads, but also for **where packs are allowed to write** (at minimum, validate/normalize `out_dir`)
* timeouts + output truncation + concurrency limits (oraclepack can run arbitrary bash steps)
* clear error taxonomy in tool responses (so clients can recover deterministically)

### 7) Client onboarding configs (so assistants can actually connect)

Depending on the target clients:

* **Codex**: document config via `~/.codex/config.toml` and/or `codex mcp` commands. ([OpenAI Developers][3])
* **Inspector**: document using the inspector to test/debug and export server configs (`mcp.json`). ([GitHub][4])
* **ChatGPT / Responses API remote MCP** (if that’s a goal): document remote server URL + auth patterns. ([OpenAI Platform][5])

### 8) Acceptance tests that lock the contracts

Minimum tests to prevent regressions:

* Stage-2 detection test matrix (oracle-out, docs/oracle-out, docs/oracle-questions-*/…, single-pack form)
* Action Pack validator parity with `validate-action-pack.sh`
* run-pack flag mapping correctness (`--resume`, ROI options, `--out-dir`, etc.)

---

## Minimum set to ship safely

1. Correct Streamable HTTP wiring (`transport="streamable-http"`) + security requirements if remote ([GitHub][1])
2. Expand `oraclepack_run_pack` to match the CLI surface
3. Stage-2 detection + Action Pack validation exactly match the oraclepack-taskify skill scripts
4. Add read tools for `_actions.json` / PRD / Task Master outputs so agents can “use artifacts in real time”

[1]: https://github.com/modelcontextprotocol/python-sdk "GitHub - modelcontextprotocol/python-sdk: The official Python SDK for Model Context Protocol servers and clients"
[2]: https://modelcontextprotocol.io/specification/2025-06-18/basic/transports "Transports - Model Context Protocol"
[3]: https://developers.openai.com/codex/mcp/?utm_source=chatgpt.com "Model Context Protocol"
[4]: https://github.com/modelcontextprotocol/inspector "GitHub - modelcontextprotocol/inspector: Visual testing tool for MCP servers"
[5]: https://platform.openai.com/docs/guides/tools-connectors-mcp?utm_source=chatgpt.com "Connectors and MCP servers | OpenAI API"
```

.tickets/mcp/oraclepack-MCP.md
```
## MCP surface for `oraclepack` (so agents can act on Taskify artifacts)

### What to expose as MCP tools

Map the existing `oraclepack` CLI capabilities (validate/list/run + flags like `--no-tui`, `--out-dir`, `--oracle-bin`) into MCP tools so an agent can run packs non-interactively and then inspect the resulting state/report/artifacts.

Add a small “taskify helper” layer to make the **Stage-2 → Stage-3** workflow deterministic for agents:

* **Detect Stage-2 outputs** (dir-form `01-*.md..20-*.md` OR single-pack form) using the ordered resolver rules described in your skill.
* **Validate Stage-2 outputs** (exactly one match per prefix 01..20).
* **Validate Stage-3 Action Pack** structure constraints (single ```bash fence, step headers `# NN)`, etc.) before executing anything.
* **Summarize Stage-3 artifacts** (e.g., `_actions.json`, `_actions.md`, PRD path, `tm-complexity.json`, pipelines doc) so agents can immediately consume them.

### Transport choices (real-time vs local)

* **stdio** is simplest for local agent runtimes; it requires MCP messages only on stdout (logs must go to stderr). ([Model Context Protocol][1])
* **Streamable HTTP** is better for “real-time” multi-client usage; implement Origin validation and bind to localhost + auth to avoid DNS rebinding and local-network abuse. ([Model Context Protocol][1])

### Tool UX metadata (important for agents)

Use MCP **tool annotations** so clients can correctly present approval UX:

* mark validate/list/read tools as `readOnlyHint: true`
* mark run tools as `destructiveHint: true`, `openWorldHint: true` (they can touch filesystem, run processes, etc.) ([Model Context Protocol][2])

### Security defaults (recommended)

* Hard **deny-by-default execution**: require an env flag (e.g. `ORACLEPACK_ENABLE_EXEC=1`) to enable “run” tools.
* Restrict filesystem access to **allowlisted roots** to avoid path traversal and accidental exfiltration of secrets.
* Enforce stdout/stderr truncation and timeouts so a pack can’t wedge the server process.

## Reference implementation

I provided a complete, drop-in **Python MCP server** (FastMCP from the official MCP Python SDK) that exposes:

* `oraclepack_validate_pack`, `oraclepack_list_steps`, `oraclepack_run_pack`
* `oraclepack_taskify_detect_stage2`, `oraclepack_taskify_validate_stage2`
* `oraclepack_taskify_validate_action_pack`, `oraclepack_taskify_artifacts_summary`
* `oraclepack_taskify_run_action_pack`

It supports **stdio** and **streamable-http** transports per the official SDK. ([GitHub][3])

The full multi-file project is in the canvas titled **“oraclepack-mcp-server”**.

[1]: https://modelcontextprotocol.io/specification/2025-03-26/basic/transports "Transports - Model Context Protocol"
[2]: https://modelcontextprotocol.io/legacy/concepts/tools "Tools - Model Context Protocol"
[3]: https://github.com/modelcontextprotocol/python-sdk "GitHub - modelcontextprotocol/python-sdk: The official Python SDK for Model Context Protocol servers and clients"
```

.tickets/mcp/oraclepack_mcp_server.md
```
# oraclepack-mcp-server (MCP wrapper for oraclepack + taskify helpers)

## tree

```text
oraclepack-mcp-server/
  README.md
  requirements.txt
  oraclepack_mcp_server/
    __init__.py
    __main__.py
    config.py
    security.py
    oraclepack_cli.py
    taskify.py
    server.py
```

---

```md
<!-- path: oraclepack-mcp-server/README.md -->
# oraclepack-mcp-server

Expose `oraclepack` (validate/list/run) plus a few Stage-3 (oraclepack-taskify) helpers as **MCP tools**, so an agent can:
- validate and inspect packs
- run packs non-interactively (`--no-tui --yes --run-all`) to produce artifacts
- validate Stage-2 outputs (01-*.md..20-*.md)
- validate Stage-3 Action Packs (single ```bash fence, step headers, etc.)
- summarize Stage-3 artifacts (`_actions.json`, PRD, Task Master outputs, etc.)

## Install

```bash
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```

## Configure (recommended)

Environment variables:

- `ORACLEPACK_ALLOWED_ROOTS` (default: current working directory)
  - Colon-separated list of allowed filesystem roots the server may read from.
  - Example: `ORACLEPACK_ALLOWED_ROOTS=/repo:/tmp/oracle-out`
- `ORACLEPACK_BIN` (default: `oraclepack`) – path to the oraclepack CLI
- `ORACLEPACK_WORKDIR` (default: current working directory)
  - Where packs are executed from (important for relative paths).
- `ORACLEPACK_ENABLE_EXEC` (default: `0`)
  - Must be `1` to enable `oraclepack_run_pack` and `oraclepack_taskify_run_action_pack`.
- `ORACLEPACK_CHARACTER_LIMIT` (default: `25000`) – truncate large stdout/stderr
- `ORACLEPACK_MAX_READ_BYTES` (default: `500000`) – max bytes read from a file

## Run (stdio)

```bash
# Stdio transport is the simplest local integration.
python -m oraclepack_mcp_server --transport stdio
```

## Run (Streamable HTTP)

```bash
python -m oraclepack_mcp_server --transport streamable-http
```

## Tools

- `oraclepack_validate_pack`
- `oraclepack_list_steps`
- `oraclepack_run_pack` (requires `ORACLEPACK_ENABLE_EXEC=1`)
- `oraclepack_read_file`
- `oraclepack_taskify_detect_stage2`
- `oraclepack_taskify_validate_stage2`
- `oraclepack_taskify_validate_action_pack`
- `oraclepack_taskify_artifacts_summary`
- `oraclepack_taskify_run_action_pack` (requires `ORACLEPACK_ENABLE_EXEC=1`)

## Typical Stage-3 usage

1) Detect/validate Stage-2 outputs (directory or single-pack)
2) Validate the Action Pack markdown
3) Run the Action Pack via `oraclepack run ...`
4) Summarize produced artifacts
```

```txt
# path: oraclepack-mcp-server/requirements.txt
mcp>=1.0.0
pydantic>=2.0.0
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/__init__.py
__all__ = []
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/config.py
from __future__ import annotations

import os
from dataclasses import dataclass
from pathlib import Path


def _truthy(value: str | None) -> bool:
    if value is None:
        return False
    return value.strip().lower() in {"1", "true", "yes", "y", "on"}


@dataclass(frozen=True)
class ServerConfig:
    allowed_roots: tuple[Path, ...]
    oraclepack_bin: str
    workdir: Path
    enable_exec: bool
    character_limit: int
    max_read_bytes: int


def load_config() -> ServerConfig:
    # Allowed roots: colon-separated. Default to CWD.
    roots_raw = os.environ.get("ORACLEPACK_ALLOWED_ROOTS")
    if roots_raw:
        roots = tuple(Path(p).expanduser().resolve() for p in roots_raw.split(":") if p.strip())
    else:
        roots = (Path.cwd().resolve(),)

    oraclepack_bin = os.environ.get("ORACLEPACK_BIN", "oraclepack")
    workdir = Path(os.environ.get("ORACLEPACK_WORKDIR", str(Path.cwd()))).expanduser().resolve()

    enable_exec = _truthy(os.environ.get("ORACLEPACK_ENABLE_EXEC", "0"))

    character_limit = int(os.environ.get("ORACLEPACK_CHARACTER_LIMIT", "25000"))
    max_read_bytes = int(os.environ.get("ORACLEPACK_MAX_READ_BYTES", "500000"))

    return ServerConfig(
        allowed_roots=roots,
        oraclepack_bin=oraclepack_bin,
        workdir=workdir,
        enable_exec=enable_exec,
        character_limit=character_limit,
        max_read_bytes=max_read_bytes,
    )
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/security.py
from __future__ import annotations

from pathlib import Path


class PathNotAllowedError(ValueError):
    pass


def resolve_under_roots(path: Path, allowed_roots: tuple[Path, ...]) -> Path:
    """Resolve a path and enforce it lives under at least one allowed root."""
    p = path.expanduser().resolve()

    for root in allowed_roots:
        r = root.expanduser().resolve()
        try:
            p.relative_to(r)
            return p
        except ValueError:
            continue

    raise PathNotAllowedError(
        f"Path not allowed (outside allowed roots). path={p} roots={[str(r) for r in allowed_roots]}"
    )


def safe_read_text(path: Path, max_bytes: int) -> tuple[str, bool]:
    """Read up to max_bytes from a file as UTF-8 (replace errors)."""
    data = path.read_bytes()
    truncated = False
    if len(data) > max_bytes:
        data = data[:max_bytes]
        truncated = True
    return data.decode("utf-8", errors="replace"), truncated


def safe_read_bytes(path: Path, max_bytes: int) -> tuple[bytes, bool]:
    data = path.read_bytes()
    truncated = False
    if len(data) > max_bytes:
        data = data[:max_bytes]
        truncated = True
    return data, truncated
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/oraclepack_cli.py
from __future__ import annotations

import asyncio
import os
import time
from dataclasses import dataclass
from pathlib import Path


@dataclass
class CmdResult:
    ok: bool
    exit_code: int
    duration_s: float
    stdout: str
    stderr: str
    stdout_truncated: bool
    stderr_truncated: bool


def _truncate(s: str, limit: int) -> tuple[str, bool]:
    if limit <= 0:
        return s, False
    if len(s) <= limit:
        return s, False
    return s[:limit], True


async def run_cmd(
    argv: list[str],
    cwd: Path,
    timeout_s: int,
    env: dict[str, str] | None,
    character_limit: int,
) -> CmdResult:
    start = time.time()

    proc = await asyncio.create_subprocess_exec(
        *argv,
        cwd=str(cwd),
        env=(os.environ | (env or {})),
        stdout=asyncio.subprocess.PIPE,
        stderr=asyncio.subprocess.PIPE,
    )

    try:
        out_b, err_b = await asyncio.wait_for(proc.communicate(), timeout=timeout_s)
    except asyncio.TimeoutError:
        proc.kill()
        await proc.communicate()
        duration = time.time() - start
        return CmdResult(
            ok=False,
            exit_code=124,
            duration_s=duration,
            stdout="",
            stderr=f"Timed out after {timeout_s}s: {' '.join(argv)}",
            stdout_truncated=False,
            stderr_truncated=False,
        )

    duration = time.time() - start
    out = out_b.decode("utf-8", errors="replace") if out_b else ""
    err = err_b.decode("utf-8", errors="replace") if err_b else ""

    out, out_tr = _truncate(out, character_limit)
    err, err_tr = _truncate(err, character_limit)

    exit_code = proc.returncode if proc.returncode is not None else 1
    return CmdResult(
        ok=(exit_code == 0),
        exit_code=exit_code,
        duration_s=duration,
        stdout=out,
        stderr=err,
        stdout_truncated=out_tr,
        stderr_truncated=err_tr,
    )
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/taskify.py
from __future__ import annotations

import re
from dataclasses import dataclass
from datetime import date
from pathlib import Path


@dataclass
class Stage2Resolution:
    kind: str  # "dir" | "file"
    stage2_path: Path
    out_dir: Path
    notes: list[str]


PREFIXES = [f"{i:02d}" for i in range(1, 21)]


def _is_iso_date(s: str) -> bool:
    # Minimal heuristic for YYYY-MM-DD
    return bool(re.fullmatch(r"\d{4}-\d{2}-\d{2}", s))


def validate_stage2_dir(out_dir: Path) -> dict:
    missing: list[str] = []
    ambiguous: dict[str, list[str]] = {}
    selected: dict[str, str] = {}

    for pfx in PREFIXES:
        matches = sorted(out_dir.glob(f"{pfx}-*.md"))
        if len(matches) == 0:
            missing.append(f"{pfx}-*.md")
        elif len(matches) > 1:
            ambiguous[pfx] = [m.name for m in matches]
        else:
            selected[pfx] = matches[0].name

    ok = (not missing) and (not ambiguous)
    return {
        "ok": ok,
        "out_dir": str(out_dir),
        "selected": selected,
        "missing": missing,
        "ambiguous": ambiguous,
    }


def _lexi_newest(paths: list[Path]) -> Path | None:
    # Deterministic: lexicographic max
    return sorted(paths, key=lambda p: p.name)[-1] if paths else None


def detect_stage2(stage2_path: str, repo_root: Path) -> Stage2Resolution:
    notes: list[str] = []

    if stage2_path != "auto":
        p = (repo_root / stage2_path).expanduser()
        if p.exists() and p.is_dir():
            return Stage2Resolution(kind="dir", stage2_path=p.resolve(), out_dir=p.resolve(), notes=["explicit dir"])
        if p.exists() and p.is_file():
            # out_dir rules from skill: if under docs/oracle-questions-YYYY-MM-DD/ then parent/oracle-out else oracle-out
            p_res = p.resolve()
            out = repo_root / "oracle-out"
            parts = list(p_res.parts)
            if "docs" in parts:
                try:
                    idx = parts.index("docs")
                    # docs/oracle-questions-YYYY-MM-DD/.../oracle-pack-YYYY-MM-DD.md
                    if idx + 1 < len(parts) and parts[idx + 1].startswith("oracle-questions-"):
                        out = Path(*parts[: idx + 2]) / "oracle-out"
                except ValueError:
                    pass
            return Stage2Resolution(kind="file", stage2_path=p_res, out_dir=out.resolve(), notes=["explicit file"])
        raise FileNotFoundError(f"stage2_path not found: {p}")

    # auto discovery (best-effort, deterministic ordering)
    searched: list[str] = []

    # 1) repo_root/oracle-out
    candidate = repo_root / "oracle-out"
    searched.append(str(candidate))
    if candidate.is_dir():
        v = validate_stage2_dir(candidate)
        if v["ok"]:
            notes.append("auto: selected repo_root/oracle-out")
            return Stage2Resolution(kind="dir", stage2_path=candidate.resolve(), out_dir=candidate.resolve(), notes=notes)

    # 2) docs/oracle-questions-YYYY-MM-DD/oracle-out (newest by lexicographic date suffix)
    docs = repo_root / "docs"
    if docs.is_dir():
        qdirs = [p for p in docs.glob("oracle-questions-*") if p.is_dir()]
        # deterministic: sort by name and take newest
        newest_q = _lexi_newest(qdirs)
        if newest_q:
            candidate = newest_q / "oracle-out"
            searched.append(str(candidate))
            if candidate.is_dir():
                v = validate_stage2_dir(candidate)
                if v["ok"]:
                    notes.append(f"auto: selected {candidate}")
                    return Stage2Resolution(kind="dir", stage2_path=candidate.resolve(), out_dir=candidate.resolve(), notes=notes)

    # 3) single-pack form (newer): look for docs/oracle-pack-*.md or docs/oraclepacks/oracle-pack-*.md
    file_candidates: list[Path] = []
    if docs.is_dir():
        file_candidates += list(docs.glob("oracle-pack-*.md"))
        file_candidates += list((docs / "oraclepacks").glob("oracle-pack-*.md"))
        file_candidates += list(docs.glob("oracle-questions-*/oraclepacks/oracle-pack-*.md"))

    newest_file = _lexi_newest(sorted([p for p in file_candidates if p.is_file()], key=lambda p: p.name))
    if newest_file:
        notes.append(f"auto: selected single-pack {newest_file}")
        out = repo_root / "oracle-out"
        # If under docs/oracle-questions-YYYY-MM-DD/..., default out_dir there.
        if "docs" in newest_file.parts:
            try:
                idx = newest_file.parts.index("docs")
                if idx + 1 < len(newest_file.parts) and newest_file.parts[idx + 1].startswith("oracle-questions-"):
                    out = Path(*newest_file.parts[: idx + 2]) / "oracle-out"
            except ValueError:
                pass
        return Stage2Resolution(kind="file", stage2_path=newest_file.resolve(), out_dir=out.resolve(), notes=notes)

    raise FileNotFoundError(
        "stage2_path=auto could not resolve Stage-2 outputs. Searched:\n- " + "\n- ".join(searched)
    )


def validate_action_pack(pack_path: Path) -> dict:
    text = pack_path.read_text(encoding="utf-8", errors="replace")

    bash_fence = re.findall(r"(?m)^\s*```bash\s*$", text)
    any_fence = re.findall(r"(?m)^\s*```\s*", text)

    errors: list[str] = []
    if len(bash_fence) != 1:
        errors.append(f"expected exactly one ```bash fence; found {len(bash_fence)}")
    if len(any_fence) != 2:
[TRUNCATED]
```

</source_code>