# Product Requirements Document (from tickets)

## Overview

This PRD is deterministically derived from repository ticket Markdown files.

## Goals

- Convert ticket intent into an actionable, trackable task plan.

## Non-goals

- Introducing undocumented scope beyond what is present in tickets.

## Ticket Inventory

- [Enable Action Packs Dispatch] Enable Action Packs Dispatch.md (.tickets/Enable Action Packs Dispatch.md)
- [Improving Oraclepack Workflow] Improving Oraclepack Workflow.md (.tickets/Improving Oraclepack Workflow.md)
- [Oraclepack Action Pack Integration] Oraclepack Action Pack Integration.md (.tickets/Oraclepack Action Pack Integration.md)
- [Oraclepack Action Pack Issue] Oraclepack Action Pack Issue.md (.tickets/Oraclepack Action Pack Issue.md)
- [Oraclepack Action Packs] Oraclepack Action Packs.md (.tickets/Oraclepack Action Packs.md)
- [Oraclepack Pipeline Improvements] Oraclepack Pipeline Improvements.md (.tickets/Oraclepack Pipeline Improvements.md)
- [Oraclepack Prompt Generator] Oraclepack Prompt Generator.md (.tickets/Oraclepack Prompt Generator.md)
- [Oraclepack Workflow Enhancement] Oraclepack Workflow Enhancement.md (.tickets/Oraclepack Workflow Enhancement.md)
- [Verbose Payload Rendering TUI] Verbose Payload Rendering TUI.md (.tickets/Verbose Payload Rendering TUI.md)
- [Expose Oraclepack as MCP] Expose Oraclepack as MCP.md (.tickets/mcp/Expose Oraclepack as MCP.md)
- [MCP Server for Oraclepack] MCP Server for Oraclepack.md (.tickets/mcp/MCP Server for Oraclepack.md)
- [oraclepack-MCP] oraclepack-MCP.md (.tickets/mcp/oraclepack-MCP.md)
- [oraclepack_mcp_server] oraclepack-mcp-server (MCP wrapper for oraclepack + taskify helpers) (.tickets/mcp/oraclepack_mcp_server.md)

## Requirements (by ticket)

### Enable Action Packs Dispatch: Enable Action Packs Dispatch.md

Title:

### Improving Oraclepack Workflow: Improving Oraclepack Workflow.md

Title:

### Oraclepack Action Pack Integration: Oraclepack Action Pack Integration.md

Parent Ticket:

### Oraclepack Action Pack Issue: Oraclepack Action Pack Issue.md

Parent Ticket:

### Oraclepack Action Packs: Oraclepack Action Packs.md

Parent Ticket:

### Oraclepack Pipeline Improvements: Oraclepack Pipeline Improvements.md

Title:

### Oraclepack Prompt Generator: Oraclepack Prompt Generator.md

Parent Ticket:

### Oraclepack Workflow Enhancement: Oraclepack Workflow Enhancement.md

Title:

### Verbose Payload Rendering TUI: Verbose Payload Rendering TUI.md

Title:

### Expose Oraclepack as MCP: Expose Oraclepack as MCP.md

Parent Ticket:

### MCP Server for Oraclepack: MCP Server for Oraclepack.md

Title:

### oraclepack-MCP: oraclepack-MCP.md

Map the existing `oraclepack` CLI capabilities (validate/list/run + flags like `--no-tui`, `--out-dir`, `--oracle-bin`) into MCP tools so an agent can run packs non-interactively and then inspect the resulting state/report/artifacts.

### oraclepack_mcp_server: oraclepack-mcp-server (MCP wrapper for oraclepack + taskify helpers)

```text oraclepack-mcp-server/ README.md requirements.txt oraclepack_mcp_server/ __init__.py __main__.py config.py security.py oraclepack_cli.py taskify.py server.py```

## Derived Actions

- Expose Oraclepack as MCP-A01: Repository includes README.md, requirements.txt, and `oraclepack_mcp_server` package directory. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A02: Running `python -m oraclepack_mcp_server --transport stdio` is supported (starts server process). (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A03: Running `python -m oraclepack_mcp_server --transport streamable-http` is supported (starts server process). (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A04: Config loader reads the listed env vars and applies defaults as documented. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A05: Path resolution rejects paths outside allowed roots. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A06: File reads enforce max bytes and indicate truncation. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A07: Exec gating flag is available for run tools to check. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A08: Runner returns: ok, exit_code, duration_s, stdout, stderr, stdout_truncated, stderr_truncated. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A09: Timeout produces exit_code=124 (or equivalent) and includes a timeout message. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A10: Outputs are truncated to configured character limit and flags are set accordingly. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A11: Validation returns ok=false with `missing` when any prefix has no matches. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A12: Validation returns ok=false with `ambiguous` when any prefix has >1 match. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A13: Validation returns ok=true only when exactly one match exists for every prefix 01..20. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A14: Detection supports explicit dir and explicit file resolution and produces an `out_dir`. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A15: Detection supports “auto” and returns deterministic results for the same filesystem state. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A16: All tools listed in the parent ticket are registered and callable. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A17: `oraclepack_read_file` enforces allowed roots and max read bytes. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A18: `oraclepack_run_pack` and `oraclepack_taskify_run_action_pack` refuse execution unless `ORACLEPACK_ENABLE_EXEC=1`. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A19: Response formatter supports markdown and json outputs and includes truncation indicators. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A20: Running with `--transport stdio` is supported and does not interleave logs on stdout. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A21: Running with `--transport streamable-http` includes Origin validation and uses localhost binding + authentication (mechanism documented/implemented). (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A22: Validate/list/read tools are annotated as read-only; run tools are annotated as destructive/open-world. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A23: Validation fails with a clear error when the action pack violates the “single bash fence” constraint. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A24: Validation fails with a clear error when step headers do not meet the stated expectations. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A25: Artifact summarizer reports presence/absence of the example artifacts and returns a readable summary. (ticket Expose Oraclepack as MCP)
- Expose Oraclepack as MCP-A26: Tool outputs are deterministic for the same filesystem state. (ticket Expose Oraclepack as MCP)
- MCP Server for Oraclepack-A01: MCP server starts successfully in both `stdio` and `streamable-http` modes. (ticket MCP Server for Oraclepack)
- MCP Server for Oraclepack-A02: All listed tools are exposed with the documented names. (ticket MCP Server for Oraclepack)
- MCP Server for Oraclepack-A03: When `ORACLEPACK_ENABLE_EXEC!=1`, run tools refuse to execute and return a clear error; validate/list/read tools still work. (ticket MCP Server for Oraclepack)
- MCP Server for Oraclepack-A04: `oraclepack_read_file` rejects paths outside `ORACLEPACK_ALLOWED_ROOTS`. (ticket MCP Server for Oraclepack)
- MCP Server for Oraclepack-A05: Stage-2 validation enforces exactly one file per prefix `01`..`20` and returns `missing` and `ambiguous` sets when invalid. (ticket MCP Server for Oraclepack)
- MCP Server for Oraclepack-A06: Stage-2 detection resolves out-dir correctly for both: (ticket MCP Server for Oraclepack)
- MCP Server for Oraclepack-A07: stdout/stderr truncation and timeouts are enforced on CLI subprocess execution. (ticket MCP Server for Oraclepack)
- MCP Server for Oraclepack-A08: Tool annotations are applied as specified for read-only vs destructive tools. (ticket MCP Server for Oraclepack)

## Acceptance Criteria

- All derived actions are implemented and validated per ticket intent.

## Open Questions

- TODO: fill unknowns discovered during implementation.
```