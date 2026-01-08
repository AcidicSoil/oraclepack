# Oraclepack MCP Server

MCP wrapper for the `oraclepack` Go CLI and Taskify helper tools.

## Features

- **CLI Wrapping**: Exposes `validate`, `list`, and `run` commands as MCP tools.
- **Taskify Helpers**: Stage-2 detection/validation and Stage-3 action-pack validation/summarization.
- **Security**: Allowlisted filesystem roots, execution gating, timeouts, and output truncation.
- **Transports**: Supports both `stdio` and `streamable-http`.

## Installation

```bash
pip install -r requirements.txt
```

## Configuration

The server is configured via environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `ORACLEPACK_BIN` | Path to the `oraclepack` binary | `oraclepack` |
| `ORACLEPACK_ALLOWED_ROOTS` | Colon-separated list of allowed filesystem roots | Current directory |
| `ORACLEPACK_WORKDIR` | Working directory for execution | Current directory |
| `ORACLEPACK_ENABLE_EXEC` | Set to `1` to enable execution tools | `0` (Disabled) |
| `ORACLEPACK_CHARACTER_LIMIT` | Max characters for tool outputs | `32000` |
| `ORACLEPACK_MAX_READ_BYTES` | Max bytes for file read operations | `65536` (64KB) |

## Connecting to Agents

### Claude Desktop

Add this to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "oraclepack": {
      "command": "python",
      "args": ["-m", "oraclepack_mcp_server", "--transport", "stdio"],
      "env": {
        "ORACLEPACK_BIN": "/path/to/your/oraclepack",
        "ORACLEPACK_ALLOWED_ROOTS": "/path/to/your/projects",
        "ORACLEPACK_ENABLE_EXEC": "1"
      }
    }
  }
}
```

### Gemini CLI

Add this to your `.mcp.json`:

```json
{
  "mcpServers": {
    "oraclepack": {
      "command": "python",
      "args": ["-m", "oraclepack_mcp_server", "--transport", "stdio"],
      "env": {
        "ORACLEPACK_BIN": "oraclepack",
        "ORACLEPACK_ALLOWED_ROOTS": ".",
        "ORACLEPACK_ENABLE_EXEC": "1"
      }
    }
  }
}
```

## Tools

- `oraclepack_read_file`: Read a file within allowed roots.
- `oraclepack_list_packs`: List available oracle packs (*.md).
- `oraclepack_validate_pack`: Validate an oracle pack.
- `oraclepack_list_steps`: List steps in an oracle pack.
- `oraclepack_run_pack`: Run an oracle pack (requires `ORACLEPACK_ENABLE_EXEC=1`).
- `oraclepack_taskify_detect_stage2`: Detect Stage-2 outputs.
- `oraclepack_taskify_validate_stage2`: Validate Stage-2 directory structure.
- `oraclepack_taskify_validate_action_pack`: Validate Stage-3 action pack structure.
- `oraclepack_taskify_artifacts_summary`: Summarize Stage-3 artifacts.
- `oraclepack_taskify_run_action_pack`: Run a Stage-3 action pack (requires `ORACLEPACK_ENABLE_EXEC=1`).
- `oraclepack_taskify_generate_prompt`: Generate instructions for an agent to run an action pack.