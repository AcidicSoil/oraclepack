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
