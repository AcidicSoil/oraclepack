# Oraclepack Interfaces

## CLI Commands
Evidence: `internal/cli/root.go`, `internal/cli/cmds.go`
-   `oraclepack [flags]`: Root command.
    -   Flags: `--no-tui` (bool), `--oracle-bin` (string), `--out-dir` (string).
-   `oraclepack validate [pack.md]`: Validates the structure of an oracle pack.
-   `oraclepack list [pack.md]`: Lists steps in a pack.
-   `oraclepack run`: (Inferred) Executes the pack (referenced by MCP server args).

## MCP Tools
Evidence: `oraclepack-mcp-server/oraclepack_mcp_server/server.py`
-   `oraclepack_read_file(path)`: Safe file reading within allowed roots.
-   `oraclepack_list_packs(directory)`: Lists `*.md` files.
-   `oraclepack_validate_pack(pack_path)`: Wraps `oraclepack validate`.
-   `oraclepack_list_steps(pack_path)`: Wraps `oraclepack list`.
-   `oraclepack_run_pack(pack_path, yes, run_all)`: Wraps `oraclepack run --no-tui`. **Destructive**.
-   `oraclepack_taskify_detect_stage2(path)`: Detects Stage-2 output directories.
-   `oraclepack_taskify_validate_stage2(out_dir)`: Validates Stage-2 structure.
-   `oraclepack_taskify_validate_action_pack(file_path)`: Validates Stage-3 structure.
-   `oraclepack_taskify_run_action_pack(file_path)`: Runs a Stage-3 pack. **Destructive**.

## Pack Format (Stage 1)
Evidence: `skills/oraclepack-codebase-pack-grouped/SKILL.md`
-   **Schema**:
    -   Exactly one `bash` fenced block.
    -   Exactly 20 steps (01..20).
    -   Direct code attachments via `${code_args[@]}`.
    -   Final Coverage check.
