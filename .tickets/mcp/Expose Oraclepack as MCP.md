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
- Provide consistent response formatting (markdown/json) and ensure run tools respect execution gating.
In Scope:
- Tool schemas/inputs covering:
  - `oraclepack_validate_pack`
  - `oraclepack_list_steps`
  - `oraclepack_run_pack` (gated)
  - `oraclepack_read_file`
  - `oraclepack_taskify_detect_stage2`
  - `oraclepack_taskify_validate_stage2`
  - `oraclepack_taskify_validate_action_pack`
  - `oraclepack_taskify_artifacts_summary`
  - `oraclepack_taskify_run_action_pack` (gated)
- Response formatting:
  - Support JSON and Markdown result formats (including stdout/stderr blocks and truncation notes).
- CLI arg mapping for oraclepack operations (including references to flags like `--no-tui`, `--out-dir`, `--oracle-bin` as inputs or internal argv composition).
- Ensure `ORACLEPACK_ENABLE_EXEC=1` gating is enforced for run tools.
Out of Scope:
- Implementing the underlying Stage-2 detection/validation logic (T4) and Stage-3 validation/summary logic (T7), except wiring them in.
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- MCP clients can invoke oraclepack and taskify workflows end-to-end via tools and receive bounded, formatted results.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Must expose the tool list as stated.
- Run tools must be disabled by default unless enabled via env flag.
Evidence:
- Tool list and gating note. :contentReference[oaicite:42]{index=42}
- Server schema snippets (e.g., ReadFileInput, Taskify*Input, timeout defaults). :contentReference[oaicite:43]{index=43}
Open Items / Unknowns:
- Exact argument surface for `oraclepack_validate_pack` / `oraclepack_list_steps` (flags and required params) is not fully specified in excerpts.
Risks / Dependencies:
- Depends on config/security/runner/taskify modules existing and being wired correctly.
Acceptance Criteria:
- [ ] All tools listed in the parent ticket are registered and callable.
- [ ] `oraclepack_read_file` enforces allowed roots and max read bytes.
- [ ] `oraclepack_run_pack` and `oraclepack_taskify_run_action_pack` refuse execution unless `ORACLEPACK_ENABLE_EXEC=1`.
- [ ] Response formatter supports markdown and json outputs and includes truncation indicators.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Tools — oraclepack_validate_pack … oraclepack_taskify_run_action_pack (requires ORACLEPACK_ENABLE_EXEC=1)” :contentReference[oaicite:44]{index=44}
- “Map … oraclepack CLI capabilities (validate/list/run + flags like --no-tui, --out-dir, --oracle-bin) into MCP tools” :contentReference[oaicite:45]{index=45}
```

```ticket T6
T# Title: Add MCP transport hardening guidance + tool UX annotations (readOnly/destructive/openWorld)
Type: enhancement
Target Area: MCP server transport configuration and tool metadata (server.py / deployment guidance)
Summary:
- Ensure the MCP server supports stdio and streamable-http in a way suitable for “real-time” usage, including the recommended security posture for HTTP transport.
- Add MCP tool annotations so clients can present appropriate approval UX for read-only vs execution tools.
In Scope:
- Transport support considerations:
  - stdio: ensure logs go to stderr (per guidance in ticket text).
  - streamable-http: implement recommended protections (“Origin validation”, “bind to localhost + auth”).
- Tool annotations:
  - Mark validate/list/read tools as `readOnlyHint: true`.
  - Mark run tools as `destructiveHint: true` and `openWorldHint: true`.
Out of Scope:
- Implementing tool business logic (T5).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- MCP clients see proper tool risk hints and HTTP transport is protected as recommended for local/real-time usage.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Must support “stdio” and “streamable-http”.
- Must provide tool annotations as described.
Evidence:
- Transport choices and HTTP security recommendations. :contentReference[oaicite:46]{index=46}
- Tool annotation guidance (readOnlyHint/destructiveHint/openWorldHint). :contentReference[oaicite:47]{index=47}
Open Items / Unknowns:
- Exact auth mechanism for streamable-http (token, mTLS, etc.): Not provided.
Risks / Dependencies:
- Depends on MCP SDK capabilities available in the chosen implementation.
Acceptance Criteria:
- [ ] Running with `--transport stdio` is supported and does not interleave logs on stdout.
- [ ] Running with `--transport streamable-http` includes Origin validation and uses localhost binding + authentication (mechanism documented/implemented).
- [ ] Validate/list/read tools are annotated as read-only; run tools are annotated as destructive/open-world.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Streamable HTTP … implement Origin validation and bind to localhost + auth to avoid DNS rebinding …” :contentReference[oaicite:48]{index=48}
- “mark validate/list/read tools as readOnlyHint … mark run tools as destructiveHint, openWorldHint …” :contentReference[oaicite:49]{index=49}
```

````ticket T7
T# Title: Implement Stage-3 Action Pack validation + artifact summarization helpers
Type: chore
Target Area: oraclepack_mcp_server/taskify.py (Stage-3 helpers) and wiring into server tools
Summary:
- Implement helper functions to validate Stage-3 “Action Pack” markdown constraints before execution and to summarize key Stage-3 artifacts produced by running action packs.
- These helpers support deterministic agent workflows around Taskify outputs.
In Scope:
- Action Pack validation logic:
  - Enforce “single ```bash fence” and “step headers” constraints as stated in the parent ticket.
- Artifacts summary logic:
  - Summarize outputs such as `_actions.json`, `_actions.md`, PRD path, `tm-complexity.json`, and “pipelines doc” when present.
- Integration points:
  - Provide outputs suitable for `oraclepack_taskify_validate_action_pack` and `oraclepack_taskify_artifacts_summary` tools (registered in T5).
Out of Scope:
- Stage-2 detection/validation (T4).
- Actual execution of action packs (tool wiring and subprocess invocation handled in T5/T3).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- Action packs can be validated for structural correctness prior to execution, and resulting artifacts can be summarized for quick agent consumption.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Validate Stage-3 Action Pack “single ```bash fence, step headers `# NN)`”.
- Summarize Stage-3 artifacts: `_actions.json`, PRD path, `tm-complexity.json`, pipelines doc, etc.
Evidence:
- Stage-3 validation requirement text. :contentReference[oaicite:50]{index=50}
- Artifact summary examples list. :contentReference[oaicite:51]{index=51}
Open Items / Unknowns:
- Exact, formal grammar for “step headers” beyond the example text: Not provided.
- Exact artifact filenames/paths beyond examples: Not provided.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Validation fails with a clear error when the action pack violates the “single bash fence” constraint.
- [ ] Validation fails with a clear error when step headers do not meet the stated expectations.
- [ ] Artifact summarizer reports presence/absence of the example artifacts and returns a readable summary.
- [ ] Tool outputs are deterministic for the same filesystem state.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Validate Stage-3 Action Pack … (single ```bash fence, step headers `# NN)` …)” :contentReference[oaicite:52]{index=52}
- “Summarize Stage-3 artifacts (e.g., `_actions.json`, `_actions.md`, PRD path, `tm-complexity.json`, pipelines doc)” :contentReference[oaicite:53]{index=53}
````
