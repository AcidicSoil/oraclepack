# Oraclepack File Map

## Go CLI (`.` and `internal/`)
-   `cmd/oraclepack/main.go`: Binary entrypoint.
-   `internal/cli/root.go`: Root Cobra command definition; global flags (`--no-tui`, `--oracle-bin`, `--out-dir`).
-   `internal/cli/cmds.go`: Subcommands definitions (`validate`, `list`).
-   `internal/cli/run.go`: Logic for the `run` subcommand (Inferred from existence in file list and MCP usage).
-   `internal/exec/`: Core logic for parsing and running steps.
    -   `oracle_scan.go`: Scans/parses oracle output.
    -   `runner.go`: Orchestrates step execution.
-   `internal/tui/`: User Interface logic (Bubble Tea).
    -   `tui.go`: Main TUI model.
    -   `url_picker.go`: Interactive URL selection.
-   `internal/pack/`: Markdown pack parsing logic (`parser.go`).
-   `internal/validate/`: Validation logic for packs (`composite.go`, `artifact_gate.go`).

## MCP Server (`oraclepack-mcp-server/`)
-   `oraclepack_mcp_server/__main__.py`: Server entrypoint; supports `stdio` and `streamable-http`.
-   `oraclepack_mcp_server/server.py`: Defines MCP tools (`oraclepack_validate_pack`, `oraclepack_run_pack`, etc.).
-   `oraclepack_mcp_server/oraclepack_cli.py`: Wrapper to invoke the Go CLI from Python.
-   `pyproject.toml`: Python project configuration.

## Skills (`skills/`)
-   `oraclepack-codebase-pack-grouped/`: Skill to generate codebase-driven packs.
    -   `SKILL.md`: Skill definition.
    -   `scripts/generate_grouped_packs.py`: Generation logic.
    -   `scripts/validate_pack.py`: Validation script.
