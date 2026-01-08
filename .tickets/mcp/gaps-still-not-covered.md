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
