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
