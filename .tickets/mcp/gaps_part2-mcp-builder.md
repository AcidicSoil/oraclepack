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
