# Oraclepack Overview

**Oraclepack** is a TUI-driven runner for oracle-based interactive bash steps, accompanied by an MCP server and pack generation skills. It orchestrates the execution of "Oracle Packs" (Markdown files with structured bash steps), offering validation, execution, and reporting capabilities.

## Major Components

1.  **CLI Runner (Go)**:
    -   Located in `cmd/oraclepack` and `internal/`.
    -   Provides a polished TUI using [Bubble Tea](https://github.com/charmbracelet/bubbletea) (Evidence: `go.mod`, `internal/tui`).
    -   Executes bash steps, validates pack structure, and manages output.

2.  **MCP Server (Python)**:
    -   Located in `oraclepack-mcp-server`.
    -   Wraps the CLI to expose functionality via the Model Context Protocol (MCP).
    -   Tools include validation, inspection, and execution (requires explicit enable).

3.  **Pack Generation (Skills)**:
    -   Located in `skills/`.
    -   Scripts (Python) to generate "Stage-1" packs from codebase or tickets.
    -   Key skill: `oraclepack-codebase-pack-grouped` (Evidence: `skills/oraclepack-codebase-pack-grouped/SKILL.md`).

## Dependencies
-   **Go**: 1.24.0 (Evidence: `go.mod`).
-   **Python**: >=3.10 (Evidence: `oraclepack-mcp-server/pyproject.toml`).
-   **External**: `oracle` binary (runtime dependency for execution, referenced in `internal/cli/root.go` flags).
