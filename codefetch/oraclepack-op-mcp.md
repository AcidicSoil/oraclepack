<filetree>
Project Structure:
├── internal
│   ├── app
│   │   ├── app.go
│   │   ├── app_test.go
│   │   ├── run.go
│   │   └── run_test.go
│   ├── artifacts
│   │   ├── contract.go
│   │   └── contract_test.go
│   ├── cli
│   │   ├── cmds.go
│   │   ├── root.go
│   │   ├── run.go
│   │   └── verify_outputs.go
│   ├── config
│   │   ├── defaults.go
│   │   └── resolve.go
│   ├── dispatch
│   │   ├── classify.go
│   │   └── classify_test.go
│   ├── errors
│   │   ├── errors.go
│   │   └── errors_test.go
│   ├── exec
│   │   ├── flags.go
│   │   ├── inject.go
│   │   ├── inject_test.go
│   │   ├── oracle_scan.go
│   │   ├── oracle_scan_test.go
│   │   ├── oracle_validate.go
│   │   ├── oracle_validate_test.go
│   │   ├── runner.go
│   │   ├── runner_test.go
│   │   ├── sanitize.go
│   │   ├── sanitize_test.go
│   │   └── stream.go
│   ├── foundation
│   │   ├── atomic.go
│   │   ├── atomic_test.go
│   │   ├── clock.go
│   │   ├── clock_test.go
│   │   ├── config.go
│   │   ├── config_test.go
│   │   ├── errors.go
│   │   └── errors_test.go
│   ├── overrides
│   │   ├── merge.go
│   │   ├── merge_test.go
│   │   └── types.go
│   ├── pack
│   │   ├── bash_syntax_validator.go
│   │   ├── bash_syntax_validator_test.go
│   │   ├── bash_tooling_checks.go
│   │   ├── metadata.go
│   │   ├── output_expectations.go
│   │   ├── output_expectations_test.go
│   │   ├── output_validator.go
│   │   ├── output_validator_test.go
│   │   ├── parser.go
│   │   ├── parser_test.go
│   │   └── verify_report.go
│   ├── render
│   │   ├── render.go
│   │   └── render_test.go
│   ├── report
│   │   ├── generate.go
│   │   ├── io.go
│   │   ├── io_test.go
│   │   ├── report_test.go
│   │   └── types.go
│   ├── shell
│   │   ├── detect.go
│   │   ├── detect_test.go
│   │   ├── engine.go
│   │   ├── engine_test.go
│   │   ├── runner.go
│   │   └── runner_test.go
│   ├── state
│   │   ├── io.go
│   │   ├── io_test.go
│   │   ├── persist.go
│   │   ├── state_test.go
│   │   └── types.go
│   ├── templates
│   │   ├── template_test.go
│   │   ├── ticket-action-pack.md
│   │   └── ticket_action_pack.go
│   ├── tools
│   │   ├── types.go
│   │   └── types_test.go
│   ├── tui
│   │   ├── clipboard.go
│   │   ├── filter_test.go
│   │   ├── overrides_confirm.go
│   │   ├── overrides_flags.go
│   │   ├── overrides_flow.go
│   │   ├── overrides_steps.go
│   │   ├── overrides_url.go
│   │   ├── preview_test.go
│   │   ├── tui.go
│   │   ├── tui_test.go
│   │   ├── url_picker.go
│   │   ├── url_store.go
│   │   └── url_store_test.go
│   ├── types
│   │   ├── pack.go
│   │   ├── pack_test.go
│   │   └── verification.go
│   └── validate
│       ├── artifact_gate.go
│       ├── artifact_gate_test.go
│       ├── composite.go
│       ├── composite_test.go
│       ├── oracle.go
│       ├── presence.go
│       ├── presence_test.go
│       ├── report.go
│       └── types.go
└── oraclepack-mcp-server
    ├── .pytest_cache
    │   ├── v
    │   │   └── cache
    │   │       ├── lastfailed
    │   │       └── nodeids
    │   └── CACHEDIR.TAG
    ├── oraclepack_mcp_server
    │   ├── __init__.py
    │   ├── __main__.py
    │   ├── config.py
    │   ├── oraclepack_cli.py
    │   ├── security.py
    │   ├── server.py
    │   └── taskify.py
    ├── oraclepack_mcp_server.egg-info
    │   ├── PKG-INFO
    │   ├── SOURCES.txt
    │   ├── dependency_links.txt
    │   ├── entry_points.txt
    │   ├── requires.txt
    │   └── top_level.txt
    ├── tests
    │   ├── test_cli.py
    │   ├── test_config.py
    │   ├── test_integration.py
    │   ├── test_security.py
    │   └── test_taskify.py
    ├── inspector.config.json
    ├── pyproject.toml
    └── requirements.txt

</filetree>

<source_code>
oraclepack-mcp-server/inspector.config.json
```
{
  "mcpServers": {
    "oraclepack": {
      "command": "/home/user/projects/temp/oraclepack/oraclepack-mcp-server/venv/bin/python",
      "args": ["-m", "oraclepack_mcp_server", "--transport", "stdio"],
      "env": {
        "ORACLEPACK_BIN": "oraclepack",
        "ORACLEPACK_ALLOWED_ROOTS": "/home/user/projects/temp/oraclepack",
        "ORACLEPACK_ENABLE_EXEC": "1"
      }
    }
  }
}
```

oraclepack-mcp-server/pyproject.toml
```
[build-system]
requires = ["setuptools>=61.0"]
build-backend = "setuptools.build_meta"

[project]
name = "oraclepack-mcp-server"
version = "0.1.0"
description = "MCP wrapper for oraclepack CLI"
authors = [
    { name = "Oraclepack Contributor" }
]
dependencies = [
    "mcp[cli]>=0.1.0",
    "pydantic-settings>=2.0.0",
    "pydantic>=2.0.0",
]
requires-python = ">=3.10"

[project.scripts]
oraclepack-mcp = "oraclepack_mcp_server.__main__:main"
```

oraclepack-mcp-server/requirements.txt
```
mcp[cli]
pydantic-settings
pydantic>=2.0
```

oraclepack-mcp-server/.pytest_cache/CACHEDIR.TAG
```
Signature: 8a477f597d28d172789f06886806bc55
# This file is a cache directory tag created by pytest.
# For information about cache directory tags, see:
#	https://bford.info/cachedir/spec.html
```

oraclepack-mcp-server/oraclepack_mcp_server.egg-info/PKG-INFO
```
Metadata-Version: 2.4
Name: oraclepack-mcp-server
Version: 0.1.0
Summary: MCP wrapper for oraclepack CLI
Author: Oraclepack Contributor
Requires-Python: >=3.10
Requires-Dist: mcp[cli]>=0.1.0
Requires-Dist: pydantic-settings>=2.0.0
Requires-Dist: pydantic>=2.0.0
```

oraclepack-mcp-server/oraclepack_mcp_server.egg-info/SOURCES.txt
```
README.md
pyproject.toml
oraclepack_mcp_server/__init__.py
oraclepack_mcp_server/__main__.py
oraclepack_mcp_server/config.py
oraclepack_mcp_server/oraclepack_cli.py
oraclepack_mcp_server/security.py
oraclepack_mcp_server/server.py
oraclepack_mcp_server/taskify.py
oraclepack_mcp_server.egg-info/PKG-INFO
oraclepack_mcp_server.egg-info/SOURCES.txt
oraclepack_mcp_server.egg-info/dependency_links.txt
oraclepack_mcp_server.egg-info/entry_points.txt
oraclepack_mcp_server.egg-info/requires.txt
oraclepack_mcp_server.egg-info/top_level.txt
tests/test_config.py
```

oraclepack-mcp-server/oraclepack_mcp_server.egg-info/dependency_links.txt
```

```

oraclepack-mcp-server/oraclepack_mcp_server.egg-info/entry_points.txt
```
[console_scripts]
oraclepack-mcp = oraclepack_mcp_server.__main__:main
```

oraclepack-mcp-server/oraclepack_mcp_server.egg-info/requires.txt
```
mcp[cli]>=0.1.0
pydantic-settings>=2.0.0
pydantic>=2.0.0
```

oraclepack-mcp-server/oraclepack_mcp_server.egg-info/top_level.txt
```
oraclepack_mcp_server
```

oraclepack-mcp-server/oraclepack_mcp_server/__init__.py
```
```

oraclepack-mcp-server/oraclepack_mcp_server/__main__.py
```
import argparse
import asyncio
from .server import mcp

def main():
    parser = argparse.ArgumentParser(description="Oraclepack MCP Server")
    parser.add_argument(
        "--transport", 
        choices=["stdio", "streamable-http"], 
        default="stdio",
        help="MCP transport to use (default: stdio)"
    )
    parser.add_argument(
        "--host", 
        default="localhost",
        help="Host to bind for streamable-http (default: localhost)"
    )
    parser.add_argument(
        "--port", 
        type=int, 
        default=8000,
        help="Port to bind for streamable-http (default: 8000)"
    )
    
    args = parser.parse_args()
    
    if args.transport == "stdio":
        mcp.run(transport="stdio")
    elif args.transport == "streamable-http":
        # FastMCP.run(transport="sse") is what maps to streamable-http in python SDK
        mcp.run(transport="sse", host=args.host, port=args.port)

if __name__ == "__main__":
    main()
```

oraclepack-mcp-server/oraclepack_mcp_server/config.py
```
import os
from pathlib import Path
from typing import List, Union, Any
from pydantic import Field, field_validator
from pydantic_settings import BaseSettings, SettingsConfigDict

class Settings(BaseSettings):
    model_config = SettingsConfigDict(
        env_prefix="ORACLEPACK_",
        env_file=".env",
        extra="ignore"
    )

    bin: str = Field(default="oraclepack", description="Path to the oraclepack binary")
    # Use Union to prevent pydantic-settings from forcing JSON decode
    allowed_roots: Any = Field(
        default_factory=lambda: [Path.cwd()],
        description="List of allowed filesystem roots"
    )
    workdir: Path = Field(default_factory=Path.cwd, description="Working directory for execution")
    enable_exec: bool = Field(default=False, description="Enable execution tools")
    max_read_bytes: int = Field(default=65536, description="Max bytes for file read operations")
    character_limit: int = Field(default=32000, description="Max characters for tool outputs")

    @field_validator("allowed_roots", mode="before")
    @classmethod
    def parse_allowed_roots(cls, v: Any) -> List[Path]:
        if isinstance(v, str):
            # Support both colon and comma separation
            delimiter = ":" if ":" in v else ","
            return [Path(p.strip()) for p in v.split(delimiter) if p.strip()]
        if isinstance(v, list):
            return [Path(p) if isinstance(p, (str, Path)) else p for p in v]
        return v

settings = Settings()
```

oraclepack-mcp-server/oraclepack_mcp_server/oraclepack_cli.py
```
import asyncio
import time
from dataclasses import dataclass
from typing import List, Optional
from .config import settings

@dataclass
class OraclepackResult:
    ok: bool
    exit_code: int
    duration_s: float
    stdout: str
    stderr: str
    stdout_truncated: bool
    stderr_truncated: bool
    error: Optional[str] = None

def truncate_output(text: str, limit: int) -> tuple[str, bool]:
    """Truncates text to limit and returns (truncated_text, is_truncated)."""
    if len(text) > limit:
        return text[:limit], True
    return text, False

async def run_oraclepack(args: List[str], timeout: float = 3600.0) -> OraclepackResult:
    """
    Runs the oraclepack CLI with the given arguments.
    Handles timeouts and output truncation.
    """
    start_time = time.time()
    
    cmd = [settings.bin] + args
    
    try:
        # Create the subprocess
        process = await asyncio.create_subprocess_exec(
            *cmd,
            stdout=asyncio.subprocess.PIPE,
            stderr=asyncio.subprocess.PIPE,
            cwd=settings.workdir
        )
        
        try:
            # Wait for completion with timeout
            stdout_bytes, stderr_bytes = await asyncio.wait_for(process.communicate(), timeout=timeout)
            exit_code = process.returncode
        except asyncio.TimeoutError:
            # Handle timeout
            process.kill()
            await process.wait() # Ensure process is cleaned up
            stdout_bytes, stderr_bytes = b"", b"Timed out after " + str(timeout).encode() + b"s"
            exit_code = 124 # Standard timeout exit code
            
    except Exception as e:
        duration = time.time() - start_time
        return OraclepackResult(
            ok=False,
            exit_code=-1,
            duration_s=duration,
            stdout="",
            stderr="",
            stdout_truncated=False,
            stderr_truncated=False,
            error=str(e)
        )

    duration = time.time() - start_time
    
    stdout_raw = stdout_bytes.decode("utf-8", errors="replace")
    stderr_raw = stderr_bytes.decode("utf-8", errors="replace")
    
    stdout, stdout_truncated = truncate_output(stdout_raw, settings.character_limit)
    stderr, stderr_truncated = truncate_output(stderr_raw, settings.character_limit)
    
    return OraclepackResult(
        ok=(exit_code == 0),
        exit_code=exit_code,
        duration_s=duration,
        stdout=stdout,
        stderr=stderr,
        stdout_truncated=stdout_truncated,
        stderr_truncated=stderr_truncated
    )
```

oraclepack-mcp-server/oraclepack_mcp_server/security.py
```
import os
from pathlib import Path
from typing import List, Optional
from .config import settings

class SecurityError(Exception):
    """Raised for security-related violations."""
    pass

def is_exec_enabled() -> bool:
    """Returns True if execution tools are explicitly enabled."""
    return settings.enable_exec

def validate_path(path: str | Path) -> Path:
    """
    Resolves a path and ensures it resides within one of the allowed roots.
    Returns the resolved Path if valid, otherwise raises SecurityError.
    """
    p = Path(path)
    # Always normalize the path to remove .. and other noise
    try:
        # resolve() is best but it follows symlinks and requires existence for full resolution on some platforms.
        # abspath + normpath is a good fallback for non-existent files.
        resolved_p = Path(os.path.abspath(os.path.normpath(p)))
    except Exception as e:
        raise SecurityError(f"Could not resolve path '{path}': {e}")

    # Check if resolved_p starts with any of the allowed roots
    is_allowed = False
    for root in settings.allowed_roots:
        try:
            resolved_root = Path(os.path.abspath(os.path.normpath(root)))
            # commonpath returns the common prefix. 
            # If resolved_p is under resolved_root, commonpath should be resolved_root.
            common = os.path.commonpath([str(resolved_root), str(resolved_p)])
            if common == str(resolved_root):
                is_allowed = True
                break
        except ValueError:
            # Different drives on Windows or other incompatibilities
            continue

    if not is_allowed:
        raise SecurityError(f"Access to path '{path}' is not allowed by ORACLEPACK_ALLOWED_ROOTS.")

    return resolved_p

def safe_read_file(path: str | Path) -> tuple[str, bool]:
    """
    Validates the path and reads its content up to max_read_bytes.
    Returns (content, truncated).
    """
    resolved_path = validate_path(path)
    
    if not resolved_path.exists():
        raise SecurityError(f"Path '{path}' does not exist.")
    if not resolved_path.is_file():
        raise SecurityError(f"Path '{path}' is not a file.")

    with open(resolved_path, "rb") as f:
        content_bytes = f.read(settings.max_read_bytes + 1)
        truncated = len(content_bytes) > settings.max_read_bytes
        content = content_bytes[:settings.max_read_bytes].decode("utf-8", errors="replace")
        return content, truncated
```

oraclepack-mcp-server/oraclepack_mcp_server/server.py
```
import logging
import sys
import os
from mcp.server.fastmcp import FastMCP
from mcp.types import ToolAnnotations
from .config import settings
from . import security
from . import oraclepack_cli
from . import taskify

# Configure logging to stderr to avoid interleaving with stdio transport
logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(name)s - %(levelname)s - %(message)s",
    stream=sys.stderr
)
logger = logging.getLogger("oraclepack-mcp-server")

# Initialize FastMCP
mcp = FastMCP("Oraclepack")

@mcp.tool(annotations=ToolAnnotations(readOnlyHint=True))
async def oraclepack_read_file(path: str) -> str:
    """
    Reads a file within allowed roots.
    Enforces ORACLEPACK_ALLOWED_ROOTS and ORACLEPACK_MAX_READ_BYTES.
    """
    content, truncated = security.safe_read_file(path)
    if truncated:
        return f"--- TRUNCATED (Max {settings.max_read_bytes} bytes) ---\n{content}"
    return content

@mcp.tool(annotations=ToolAnnotations(readOnlyHint=True))
async def oraclepack_list_packs(directory: str = ".") -> str:
    """Lists available oracle packs (*.md) in a directory."""
    resolved_dir = security.validate_path(directory)
    if not resolved_dir.is_dir():
        return f"Path '{directory}' is not a directory."
    
    packs = list(resolved_dir.glob("*.md"))
    if not packs:
        return f"No oracle packs found in '{directory}'."
    
    return "\n".join([p.name for p in packs])

@mcp.tool(annotations=ToolAnnotations(readOnlyHint=True))
async def oraclepack_validate_pack(pack_path: str) -> str:
    """Validates an oracle pack using the Go CLI."""
    resolved_path = security.validate_path(pack_path)
    result = await oraclepack_cli.run_oraclepack(["validate", str(resolved_path)])
    if not result.ok:
        return f"Validation failed:\n{result.stderr or result.error}"
    return "Pack is valid."

@mcp.tool(annotations=ToolAnnotations(readOnlyHint=True))
async def oraclepack_list_steps(pack_path: str) -> str:
    """Lists steps in an oracle pack."""
    resolved_path = security.validate_path(pack_path)
    result = await oraclepack_cli.run_oraclepack(["list", str(resolved_path)])
    if not result.ok:
        return f"Failed to list steps:\n{result.stderr or result.error}"
    return result.stdout

@mcp.tool(annotations=ToolAnnotations(destructiveHint=True, openWorldHint=True))
async def oraclepack_run_pack(pack_path: str, yes: bool = True, run_all: bool = True) -> str:
    """
    Runs an oracle pack. REQUIRES ORACLEPACK_ENABLE_EXEC=1.
    """
    if not security.is_exec_enabled():
        return "Execution is disabled. Set ORACLEPACK_ENABLE_EXEC=1 to enable."
    
    resolved_path = security.validate_path(pack_path)
    args = ["run", str(resolved_path), "--no-tui"]
    if yes: args.append("--yes")
    if run_all: args.append("--run-all")
    
    result = await oraclepack_cli.run_oraclepack(args)
    
    # Verbose Payload Rendering
    output = [f"# Execution Report: {pack_path}"]
    output.append(f"**Status**: {'✅ SUCCESS' if result.ok else '❌ FAILED'}")
    output.append(f"**Exit Code**: {result.exit_code}")
    output.append(f"**Duration**: {result.duration_s:.2f}s")
    
    if result.error:
        output.append(f"\n### Error\n{result.error}")
    
    if result.stdout:
        output.append("\n### Standard Output")
        output.append(f"```\n{result.stdout}\n```")
        if result.stdout_truncated:
            output.append("*Note: Output was truncated.*")
            
    if result.stderr:
        output.append("\n### Standard Error")
        output.append(f"```\n{result.stderr}\n```")
        if result.stderr_truncated:
            output.append("*Note: Error output was truncated.*")
            
    # Add artifact summary if successful or partially successful
    parent_dir = resolved_path.parent
    summary = taskify.artifacts_summary(parent_dir)
    output.append("\n### Artifacts Summary")
    for name, info in summary["artifacts"].items():
        if info:
            output.append(f"- {name}: FOUND")
        else:
            output.append(f"- {name}: NOT FOUND")
        
    return "\n".join(output)

@mcp.tool(annotations=ToolAnnotations(readOnlyHint=True))
async def oraclepack_taskify_detect_stage2(path: str = "auto") -> str:
    """Detects Stage-2 outputs."""
    out_dir, mode = taskify.detect_stage2(path, os.getcwd())
    if not out_dir:
        return f"Could not detect Stage-2 outputs in mode '{mode}'."
[TRUNCATED]
```

oraclepack-mcp-server/oraclepack_mcp_server/taskify.py
```
import os
import re
import glob
from pathlib import Path
from typing import List, Dict, Optional, Set, Tuple, Any
from dataclasses import dataclass, field

@dataclass
class Stage2ValidationResult:
    ok: bool
    missing: List[str] = field(default_factory=list)
    ambiguous: Dict[str, List[str]] = field(default_factory=dict)
    valid_files: List[str] = field(default_factory=list)

@dataclass
class ActionPackValidationResult:
    ok: bool
    error: Optional[str] = None
    steps: List[str] = field(default_factory=list)

def validate_stage2_dir(out_dir: str | Path) -> Stage2ValidationResult:
    """
    Enforces exactly one file per prefix 01..20.
    Returns Stage2ValidationResult with missing and ambiguous sets.
    """
    out_dir = Path(out_dir)
    missing = []
    ambiguous = {}
    valid_files = []
    
    for i in range(1, 21):
        pfx = f"{i:02d}"
        matches = list(out_dir.glob(f"{pfx}-*.md"))
        
        if not matches:
            missing.append(pfx)
        elif len(matches) > 1:
            ambiguous[pfx] = [m.name for m in matches]
        else:
            valid_files.append(matches[0].name)
            
    return Stage2ValidationResult(
        ok=(not missing and not ambiguous),
        missing=missing,
        ambiguous=ambiguous,
        valid_files=valid_files
    )

def detect_stage2(stage2_path: str, repo_root: str | Path) -> Tuple[Optional[Path], str]:
    """
    Resolves out_dir for Stage-2.
    Supports explicit dir, explicit file, and 'auto'.
    Returns (out_dir, mode).
    """
    repo_root = Path(repo_root)
    
    if stage2_path == "auto":
        candidates = [
            Path.cwd() / "oracle-out",
            repo_root / "oracle-out"
        ]
        
        docs_dir = repo_root / "docs"
        if docs_dir.exists():
            q_dirs = sorted(list(docs_dir.glob("oracle-questions-*")), reverse=True)
            for qd in q_dirs:
                candidates.append(qd / "oracle-out")
                
        for c in candidates:
            if c.exists() and c.is_dir():
                return c, "auto"
        return None, "auto"

    p = Path(stage2_path)
    if p.is_dir():
        return p, "explicit_dir"
    
    if p.is_file():
        parent = p.parent
        if (parent / "oracle-out").exists():
            return parent / "oracle-out", "explicit_file"
        return parent, "explicit_file"

    return None, "unknown"

def validate_action_pack(file_path: str | Path) -> ActionPackValidationResult:
    """
    Validates Stage-3 Action Pack constraints:
    - Single bash code fence
    - Step headers # NN)
    """
    try:
        content = Path(file_path).read_text()
    except Exception as e:
        return ActionPackValidationResult(ok=False, error=f"Read error: {e}")

    fences = re.findall(r"```bash", content)
    if len(fences) == 0:
        return ActionPackValidationResult(ok=False, error="No ```bash code fence found.")
    if len(fences) > 1:
        return ActionPackValidationResult(ok=False, error="Multiple ```bash code fences found. Only one is allowed.")

    steps = re.findall(r"^#\s*(\d+)\)", content, re.MULTILINE)
    if not steps:
        return ActionPackValidationResult(ok=False, error="No step headers (e.g. '# 01)') found.")

    return ActionPackValidationResult(ok=True, steps=steps)

def artifacts_summary(out_dir: str | Path) -> Dict[str, Any]:
    """
    Summarizes key Stage-3 artifacts.
    """
    out_dir = Path(out_dir)
    summary = {
        "out_dir": str(out_dir),
        "artifacts": {}
    }
    
    important_files = [
        "_actions.json",
        "_actions.md",
        "_tickets_index.json",
        "ticket-action-pack.md",
        "tm-complexity.json",
        "PRD.md",
        "prd.md",
        "tickets_prd.md"
    ]
    
    for f in important_files:
        found = False
        for search_path in [out_dir / f, out_dir.parent / f, out_dir.parent.parent / f]:
            if search_path.exists():
                summary["artifacts"][f] = {
                    "path": str(search_path),
                    "size": search_path.stat().st_size
                }
                found = True
                break
        if not found:
            summary["artifacts"][f] = None
            
    return summary

def generate_agent_prompt(pack_path: str, steps: List[str]) -> str:
    """
    Generates a prompt for an agent to run an action pack.
    """
    return f"""
# Oraclepack Action Pack Instructions

You are about to run an Oraclepack Action Pack: `{pack_path}`.
This pack contains {len(steps)} steps: {', '.join(steps)}.

## Recommended Workflow

1. **Verify**: Use `oraclepack_taskify_validate_action_pack` to ensure structure is correct.
[TRUNCATED]
```

oraclepack-mcp-server/tests/test_cli.py
```
import asyncio
import pytest
from unittest.mock import AsyncMock, patch, MagicMock
from oraclepack_mcp_server.oraclepack_cli import run_oraclepack, OraclepackResult
from oraclepack_mcp_server.config import settings

@pytest.mark.asyncio
async def test_run_oraclepack_success():
    # Mock process.communicate
    mock_stdout = b"success output"
    mock_stderr = b""
    
    with patch("asyncio.create_subprocess_exec", new_callable=AsyncMock) as mock_exec:
        mock_process = AsyncMock()
        mock_process.communicate.return_value = (mock_stdout, mock_stderr)
        mock_process.returncode = 0
        mock_exec.return_value = mock_process
        
        result = await run_oraclepack(["list", "pack.md"])
        
        assert result.ok is True
        assert result.exit_code == 0
        assert result.stdout == "success output"
        assert result.stdout_truncated is False

@pytest.mark.asyncio
async def test_run_oraclepack_timeout():
    with patch("asyncio.create_subprocess_exec", new_callable=AsyncMock) as mock_exec:
        mock_process = AsyncMock()
        
        async def slow_communicate():
            await asyncio.sleep(10)
            return (b"", b"")
            
        mock_process.communicate.side_effect = slow_communicate
        mock_exec.return_value = mock_process
        
        # Run with short timeout
        result = await run_oraclepack(["run", "pack.md"], timeout=0.1)
        
        assert result.ok is False
        assert result.exit_code == 124
        assert "Timed out" in result.stderr

@pytest.mark.asyncio
async def test_run_oraclepack_truncation(monkeypatch):
    monkeypatch.setattr(settings, "character_limit", 5)
    
    mock_stdout = b"1234567890"
    mock_stderr = b""
    
    with patch("asyncio.create_subprocess_exec", new_callable=AsyncMock) as mock_exec:
        mock_process = AsyncMock()
        mock_process.communicate.return_value = (mock_stdout, mock_stderr)
        mock_process.returncode = 0
        mock_exec.return_value = mock_process
        
        result = await run_oraclepack(["list"])
        
        assert result.stdout == "12345"
        assert result.stdout_truncated is True
```

oraclepack-mcp-server/tests/test_config.py
```
import os
import pytest
from pathlib import Path
from oraclepack_mcp_server.config import Settings

def test_default_config():
    # Clear env vars that might interfere
    for key in os.environ:
        if key.startswith("ORACLEPACK_"):
            del os.environ[key]
    
    # Reload settings or create a new instance
    # Note: the 'settings' object is already instantiated, so we test a new instance
    s = Settings()
    assert s.bin == "oraclepack"
    assert s.enable_exec is False
    assert s.character_limit == 32000
    assert Path.cwd() in s.allowed_roots

def test_env_override():
    os.environ["ORACLEPACK_BIN"] = "/custom/path/oraclepack"
    os.environ["ORACLEPACK_ENABLE_EXEC"] = "1"
    os.environ["ORACLEPACK_ALLOWED_ROOTS"] = "/tmp:/var/log"
    os.environ["ORACLEPACK_CHARACTER_LIMIT"] = "1000"
    
    s = Settings()
    assert s.bin == "/custom/path/oraclepack"
    assert s.enable_exec is True
    assert Path("/tmp") in s.allowed_roots
    assert Path("/var/log") in s.allowed_roots
    assert s.character_limit == 1000
    
    # Cleanup
    del os.environ["ORACLEPACK_BIN"]
    del os.environ["ORACLEPACK_ENABLE_EXEC"]
    del os.environ["ORACLEPACK_ALLOWED_ROOTS"]
    del os.environ["ORACLEPACK_CHARACTER_LIMIT"]
```

oraclepack-mcp-server/tests/test_integration.py
```
import asyncio
import pytest
import sys
from mcp.client.session import ClientSession
from mcp.client.stdio import stdio_client, StdioServerParameters
from mcp.types import CallToolRequestParams
import os
from pathlib import Path

@pytest.mark.asyncio
async def test_server_tools_list():
    """
    Spins up the server via stdio and checks if it can list its tools.
    """
    # Path to our package
    server_params = StdioServerParameters(
        command=sys.executable,
        args=["-m", "oraclepack_mcp_server", "--transport", "stdio"],
        env={**os.environ, "PYTHONPATH": str(Path(__file__).parent.parent)}
    )
    
    async with stdio_client(server_params) as (read, write):
        async with ClientSession(read, write) as session:
            await session.initialize()
            
            tools = await session.list_tools()
            tool_names = [t.name for t in tools.tools]
            
            assert "oraclepack_read_file" in tool_names
            assert "oraclepack_validate_pack" in tool_names
            assert "oraclepack_run_pack" in tool_names
            assert "oraclepack_taskify_detect_stage2" in tool_names
            assert "oraclepack_taskify_validate_stage2" in tool_names
            assert "oraclepack_taskify_validate_action_pack" in tool_names
            assert "oraclepack_taskify_artifacts_summary" in tool_names
            assert "oraclepack_taskify_generate_prompt" in tool_names

@pytest.mark.asyncio
async def test_oraclepack_read_file_unauthorized(tmp_path):
    """
    Verifies that the server enforces allowed roots.
    """
    root1 = tmp_path / "allowed"
    root1.mkdir()
    outside = tmp_path / "outside.txt"
    outside.touch()
    
    server_params = StdioServerParameters(
        command=sys.executable,
        args=["-m", "oraclepack_mcp_server", "--transport", "stdio"],
        env={
            **os.environ, 
            "PYTHONPATH": str(Path(__file__).parent.parent),
            "ORACLEPACK_ALLOWED_ROOTS": str(root1)
        }
    )
    
    async with stdio_client(server_params) as (read, write):
        async with ClientSession(read, write) as session:
            await session.initialize()
            
            # Try to read file outside allowed root
            result = await session.call_tool("oraclepack_read_file", {"path": str(outside)})
            
            # FastMCP returns the result even if there was an internal exception string returned
            assert "Access to path" in result.content[0].text
            assert "is not allowed" in result.content[0].text
```

oraclepack-mcp-server/tests/test_security.py
```
import os
import pytest
from pathlib import Path
from oraclepack_mcp_server.security import validate_path, is_exec_enabled, SecurityError, safe_read_file
from oraclepack_mcp_server.config import Settings, settings

def test_is_exec_enabled(monkeypatch):
    # Test with default (False)
    monkeypatch.setattr(settings, "enable_exec", False)
    assert is_exec_enabled() is False
    
    # Test with True
    monkeypatch.setattr(settings, "enable_exec", True)
    assert is_exec_enabled() is True

def test_validate_path_allowed(tmp_path, monkeypatch):
    # Setup tmp_path as an allowed root
    monkeypatch.setattr(settings, "allowed_roots", [tmp_path])
    
    # Path inside allowed root
    test_file = tmp_path / "test.txt"
    test_file.touch()
    
    assert validate_path(test_file) == test_file.resolve()
    assert validate_path(str(test_file)) == test_file.resolve()

def test_validate_path_denied(tmp_path, monkeypatch):
    # Setup allowed root
    root1 = tmp_path / "root1"
    root1.mkdir()
    monkeypatch.setattr(settings, "allowed_roots", [root1])
    
    # Path outside allowed root
    outside_file = tmp_path / "outside.txt"
    outside_file.touch()
    
    with pytest.raises(SecurityError, match="not allowed"):
        validate_path(outside_file)

def test_validate_path_traversal(tmp_path, monkeypatch):
    root1 = tmp_path / "root1"
    root1.mkdir()
    monkeypatch.setattr(settings, "allowed_roots", [root1])
    
    # Try to traverse out
    traversal_path = root1 / ".." / "outside.txt"
    
    with pytest.raises(SecurityError, match="not allowed"):
        validate_path(traversal_path)

def test_safe_read_file(tmp_path, monkeypatch):
    monkeypatch.setattr(settings, "allowed_roots", [tmp_path])
    monkeypatch.setattr(settings, "max_read_bytes", 10)
    
    test_file = tmp_path / "large.txt"
    test_file.write_text("0123456789ABCDE") # 15 chars
    
    content, truncated = safe_read_file(test_file)
    assert content == "0123456789"
    assert truncated is True
    
    small_file = tmp_path / "small.txt"
    small_file.write_text("hello")
    content, truncated = safe_read_file(small_file)
    assert content == "hello"
    assert truncated is False
```

oraclepack-mcp-server/tests/test_taskify.py
```
import pytest
from pathlib import Path
from oraclepack_mcp_server.taskify import validate_stage2_dir, detect_stage2, validate_action_pack

def test_validate_stage2_dir_ok(tmp_path):
    # Create 01..20 files
    for i in range(1, 21):
        (tmp_path / f"{i:02d}-test.md").touch()
    
    result = validate_stage2_dir(tmp_path)
    assert result.ok is True
    assert len(result.valid_files) == 20

def test_validate_stage2_dir_missing(tmp_path):
    # Missing 05 and 10
    for i in range(1, 21):
        if i in [5, 10]: continue
        (tmp_path / f"{i:02d}-test.md").touch()
    
    result = validate_stage2_dir(tmp_path)
    assert result.ok is False
    assert "05" in result.missing
    assert "10" in result.missing

def test_validate_stage2_dir_ambiguous(tmp_path):
    # Double 01
    (tmp_path / "01-a.md").touch()
    (tmp_path / "01-b.md").touch()
    for i in range(2, 21):
        (tmp_path / f"{i:02d}-test.md").touch()
        
    result = validate_stage2_dir(tmp_path)
    assert result.ok is False
    assert "01" in result.ambiguous
    assert len(result.ambiguous["01"]) == 2

def test_validate_action_pack_ok(tmp_path):
    pack_file = tmp_path / "pack.md"
    pack_file.write_text("""
# My Action Pack
# 01) Step One
```bash
echo hello
```
# 02) Step Two
""")
    result = validate_action_pack(pack_file)
    assert result.ok is True
    assert result.steps == ["01", "02"]

def test_validate_action_pack_multiple_fences(tmp_path):
    pack_file = tmp_path / "pack.md"
    pack_file.write_text("""
```bash
echo one
```
```bash
echo two
```
""")
    result = validate_action_pack(pack_file)
    assert result.ok is False
    assert "Multiple" in result.error

def test_detect_stage2_auto(tmp_path, monkeypatch):
    oracle_out = tmp_path / "oracle-out"
    oracle_out.mkdir()
    
    # Mock current working directory or just pass repo_root
    result_dir, mode = detect_stage2("auto", tmp_path)
    assert result_dir == oracle_out
    assert mode == "auto"
```

oraclepack-mcp-server/.pytest_cache/v/cache/lastfailed
```
{}
```

oraclepack-mcp-server/.pytest_cache/v/cache/nodeids
```
[
  "tests/test_cli.py::test_run_oraclepack_success",
  "tests/test_cli.py::test_run_oraclepack_timeout",
  "tests/test_cli.py::test_run_oraclepack_truncation",
  "tests/test_config.py::test_default_config",
  "tests/test_config.py::test_env_override",
  "tests/test_integration.py::test_oraclepack_read_file_unauthorized",
  "tests/test_integration.py::test_server_tools_list",
  "tests/test_security.py::test_is_exec_enabled",
  "tests/test_security.py::test_safe_read_file",
  "tests/test_security.py::test_validate_path_allowed",
  "tests/test_security.py::test_validate_path_denied",
  "tests/test_security.py::test_validate_path_traversal",
  "tests/test_taskify.py::test_detect_stage2_auto",
  "tests/test_taskify.py::test_validate_action_pack_multiple_fences",
  "tests/test_taskify.py::test_validate_action_pack_ok",
  "tests/test_taskify.py::test_validate_stage2_dir_ambiguous",
  "tests/test_taskify.py::test_validate_stage2_dir_missing",
  "tests/test_taskify.py::test_validate_stage2_dir_ok"
]
```

internal/app/app.go
```
package app

import (
	"fmt"
	"os"

	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/types"
)

// Config holds application-wide configuration.
type Config struct {
	PackPath              string
	StatePath             string
	ReportPath            string
	StopOnFail            bool
	Resume                bool
	Verbose               bool
	DryRun                bool
	OracleFlags           []string
	WorkDir               string
	OutDir                string // CLI override for output directory
	ROIThreshold          float64
	ROIMode               string // "over" or "under"
	OutputVerify          bool
	OutputRetries         int
	OutputRequireHeadings bool
	OutputChunkMode       string
}

// App orchestrates the execution flow.
type App struct {
	Config Config
	Pack   *types.Pack
	State  *state.RunState
	Runner *exec.Runner
}

// New creates a new application instance.
func New(cfg Config) *App {
	return &App{
		Config: cfg,
		Runner: exec.NewRunner(exec.RunnerOptions{
			WorkDir:     cfg.WorkDir,
			OracleFlags: cfg.OracleFlags,
		}),
	}
}

// LoadPack loads and validates the pack.
func (a *App) LoadPack() error {
	data, err := os.ReadFile(a.Config.PackPath)
	if err != nil {
		return err
	}

	p, err := pack.Parse(data)
	if err != nil {
		return err
	}

	if err := pack.Validate(p); err != nil {
		return err
	}

	a.Pack = p
	a.Pack.Source = a.Config.PackPath
	return nil
}

// LoadState loads or initializes the state.
func (a *App) LoadState() error {
	if a.Config.Resume {
		s, err := state.LoadState(a.Config.StatePath)
		if err == nil {
			a.State = s
			return nil
		}
	}

	a.State = &state.RunState{
		SchemaVersion: 1,
		StepStatuses:  make(map[string]state.StepStatus),
	}
	return nil
}

// Prepare resolves configuration and prepares the runtime environment.
func (a *App) Prepare() error {
	if a.Pack == nil {
		if err := a.LoadPack(); err != nil {
			return err
		}
	}

	// Resolve Output Directory
	// Precedence: CLI > Pack > Default (.)
	outDir := a.Config.OutDir
	if outDir == "" && a.Pack.OutDir != "" {
		outDir = a.Pack.OutDir
	}
	if outDir == "" {
		outDir = "."
	}

	// Provision Directory
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory %s: %w", outDir, err)
	}

	// Update Runner
	// We do NOT set WorkDir to outDir, so execution happens in the project root.
	// This preserves relative path resolution for -f flags.
	// a.Runner.WorkDir = outDir

	// Add out_dir to Env so scripts can reference it
	a.Runner.Env = append(a.Runner.Env, fmt.Sprintf("out_dir=%s", outDir))

	return nil
}
```

internal/app/app_test.go
```
package app

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"
)

func TestApp_RunPlain(t *testing.T) {
	steps := buildSteps(20, "echo")
	packContent := `
# Test Pack
` + "```" + `bash
` + steps + `
` + "```" + `
`
	packFile := "test.md"
	stateFile := "test_state.json"
	reportFile := "test_report.json"
	defer os.Remove(packFile)
	defer os.Remove(stateFile)
	defer os.Remove(reportFile)

	os.WriteFile(packFile, []byte(packContent), 0644)

	cfg := Config{
		PackPath:   packFile,
		StatePath:  stateFile,
		ReportPath: reportFile,
	}

	a := New(cfg)
	if err := a.Prepare(); err != nil {
		t.Fatalf("Prepare failed: %v", err)
	}
	if err := a.LoadState(); err != nil {
		t.Fatalf("LoadState failed: %v", err)
	}

	var out bytes.Buffer
	err := a.RunPlain(context.Background(), &out)
	if err != nil {
		t.Fatalf("RunPlain failed: %v", err)
	}

	output := out.String()
	if !contains(output, "step 1") || !contains(output, "step 2") {
		t.Errorf("output missing steps: %s", output)
	}

	if _, err := os.Stat(stateFile); os.IsNotExist(err) {
		t.Error("state file was not created")
	}

	if _, err := os.Stat(reportFile); os.IsNotExist(err) {
		t.Error("report file was not created")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || (len(substr) > 0 && (s[:len(substr)] == substr || contains(s[1:], substr))))
}

func buildSteps(count int, cmd string) string {
	var b bytes.Buffer
	for i := 1; i <= count; i++ {
		if i < 10 {
			b.WriteString("# 0")
		} else {
			b.WriteString("# ")
		}
		b.WriteString(fmt.Sprintf("%d)\n", i))
		b.WriteString(cmd)
		b.WriteString(fmt.Sprintf(" \"step %d\"\n", i))
	}
	return b.String()
}
```

internal/app/run.go
```
package app

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/report"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/types"
)

func (a *App) RunPlain(ctx context.Context, out io.Writer) error {
	// Assumes a.Prepare() and a.LoadState() have been called by the CLI entrypoint.
	if a.Pack == nil {
		return fmt.Errorf("pack not loaded")
	}
	if a.State == nil {
		return fmt.Errorf("state not loaded")
	}

	if a.State.StartTime.IsZero() {
		a.State.StartTime = time.Now()
	}

	fmt.Fprintf(out, "Running pack: %s\n", a.Config.PackPath)
	fmt.Fprintf(out, "Output directory: %s\n", a.Runner.WorkDir)

	// Prelude
	if a.Pack.Prelude.Code != "" {
		fmt.Fprintln(out, "Executing prelude...")
		err := a.Runner.RunPrelude(ctx, &a.Pack.Prelude, out)
		a.recordWarnings()
		if err != nil {
			return fmt.Errorf("prelude failed: %w", err)
		}
	}

	for _, step := range a.Pack.Steps {
		a.State.CurrentStep = step.Number
		// Filter by ROI
		if a.Config.ROIThreshold > 0 {
			if a.Config.ROIMode == "under" {
				// "under" is strictly less than
				if step.ROI >= a.Config.ROIThreshold {
					fmt.Fprintf(out, "Skipping step %s (ROI %.2f >= %.2f)\n", step.ID, step.ROI, a.Config.ROIThreshold)
					continue
				}
			} else {
				// "over" is greater than or equal to (3.3 or higher)
				if step.ROI < a.Config.ROIThreshold {
					fmt.Fprintf(out, "Skipping step %s (ROI %.2f < %.2f)\n", step.ID, step.ROI, a.Config.ROIThreshold)
					continue
				}
			}
		}

		// Check resume
		if s, ok := a.State.StepStatuses[step.ID]; ok && s.Status == state.StatusSuccess {
			fmt.Fprintf(out, "Skipping step %s (already succeeded)\n", step.ID)
			continue
		}

		fmt.Fprintf(out, "\n>>> Step %s: %s\n", step.ID, step.OriginalLine)

		status := state.StepStatus{
			Status:    state.StatusRunning,
			StartedAt: time.Now(),
		}
		a.State.StepStatuses[step.ID] = status
		a.saveState()

		// Execute
		err := a.runStepWithOutputVerification(ctx, &step, out)
		a.recordWarnings()

		status.EndedAt = time.Now()
		if err != nil {
			status.Status = state.StatusFailed
			status.Error = err.Error()
			a.State.StepStatuses[step.ID] = status
			a.saveState()

			if a.Config.StopOnFail {
				a.finalize(out)
				return err
			}
			continue
		}

		status.Status = state.StatusSuccess
		status.ExitCode = 0
		a.State.StepStatuses[step.ID] = status
		a.saveState()
	}

	a.finalize(out)
	return nil
}

func (a *App) runStepWithOutputVerification(ctx context.Context, step *types.Step, out io.Writer) error {
	retries := a.Config.OutputRetries
	if retries < 0 {
		retries = 0
	}
	for attempt := 0; attempt <= retries; attempt++ {
		err := a.Runner.RunStep(ctx, step, out)
		if err != nil {
			return err
		}
		if !a.Config.OutputVerify {
			return nil
		}
		outputFailures := pack.VerifyStepOutputs(step, a.Config.OutputRequireHeadings, a.Config.OutputChunkMode)
		if len(outputFailures) == 0 {
			return nil
		}
		var failures []string
		for _, failure := range outputFailures {
			if failure.Error != "" {
				failures = append(failures, fmt.Sprintf("%s error: %s", failure.Path, failure.Error))
				continue
			}
			if len(failure.MissingTokens) > 0 {
				failures = append(failures, fmt.Sprintf("%s missing: %s", failure.Path, strings.Join(failure.MissingTokens, ", ")))
			}
		}
		if len(failures) == 0 {
			return nil
		}
		if attempt == retries {
			return fmt.Errorf(
				"output verification failed for step %s: %s",
				step.ID,
				strings.Join(failures, "; "),
			)
		}
		fmt.Fprintf(out, "⚠ output verification failed for step %s (%s); re-running (%d/%d)...\n",
			step.ID, strings.Join(failures, "; "), attempt+1, retries)
	}
	return nil
}

func (a *App) recordWarnings() {
	if a.State == nil || a.Runner == nil {
		return
	}
	warnings := a.Runner.DrainWarnings()
	if len(warnings) == 0 {
		return
	}
	for _, w := range warnings {
		a.State.Warnings = append(a.State.Warnings, state.Warning{
			Scope:   w.Scope,
			StepID:  w.StepID,
			Line:    w.Line,
			Token:   w.Token,
			Message: w.Message,
		})
	}
[TRUNCATED]
```

internal/app/run_test.go
```
package app

import (
	"bytes"
	"context"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestApp_RunPlain_ROI(t *testing.T) {
	steps := buildROISteps()
	packContent := `
# ROI Test Pack
` + "```" + `bash
` + steps + `
` + "```" + `
`
	packFile := "roi_test.md"
	defer os.Remove(packFile)
	os.WriteFile(packFile, []byte(packContent), 0644)

	// Test Case 1: Filter OVER 3.3 (Should run 5.0 and 3.3)
	t.Run("Filter Over 3.3", func(t *testing.T) {
		var out bytes.Buffer
		cfg := Config{
			PackPath:     packFile,
			ROIThreshold: 3.3,
			ROIMode:      "over",
		}
		app := New(cfg)
		if err := app.Prepare(); err != nil {
			t.Fatalf("Prepare failed: %v", err)
		}
		if err := app.LoadState(); err != nil {
			t.Fatalf("LoadState failed: %v", err)
		}
		if err := app.RunPlain(context.Background(), &out); err != nil {
			t.Fatalf("RunPlain failed: %v", err)
		}
		output := out.String()
		if !strings.Contains(output, "Step 01") {
			t.Error("expected Step 01 (5.0) to run")
		}
		if !strings.Contains(output, "Step 02") {
			t.Error("expected Step 02 (3.3) to run (inclusive)")
		}
		if strings.Contains(output, "Step 03") && !strings.Contains(output, "Skipping step 03") {
			t.Error("expected Step 03 (1.0) to be skipped")
		}
	})

	// Test Case 2: Filter UNDER 3.3 (Should run 1.0 only)
	t.Run("Filter Under 3.3", func(t *testing.T) {
		var out bytes.Buffer
		cfg := Config{
			PackPath:     packFile,
			ROIThreshold: 3.3,
			ROIMode:      "under",
		}
		app := New(cfg)
		if err := app.Prepare(); err != nil {
			t.Fatalf("Prepare failed: %v", err)
		}
		if err := app.LoadState(); err != nil {
			t.Fatalf("LoadState failed: %v", err)
		}
		if err := app.RunPlain(context.Background(), &out); err != nil {
			t.Fatalf("RunPlain failed: %v", err)
		}
		output := out.String()
		if strings.Contains(output, "Step 01") && !strings.Contains(output, "Skipping step 01") {
			t.Error("expected Step 01 (5.0) to be skipped")
		}
		if strings.Contains(output, "Step 02") && !strings.Contains(output, "Skipping step 02") {
			t.Error("expected Step 02 (3.3) to be skipped (exclusive)")
		}
		if !strings.Contains(output, "Step 03") {
			t.Error("expected Step 03 (1.0) to run")
		}
	})
}

func buildROISteps() string {
	var b strings.Builder
	for i := 1; i <= 20; i++ {
		id := i
		if id < 10 {
			b.WriteString("# 0")
		} else {
			b.WriteString("# ")
		}
		b.WriteString(strconv.Itoa(id))
		if i == 1 {
			b.WriteString(") ROI=5.0\n")
			b.WriteString("echo \"high\"\n\n")
			continue
		}
		if i == 2 {
			b.WriteString(") ROI=3.3\n")
			b.WriteString("echo \"threshold\"\n\n")
			continue
		}
		if i == 3 {
			b.WriteString(") ROI=1.0\n")
			b.WriteString("echo \"low\"\n\n")
			continue
		}
		b.WriteString(")\n")
		b.WriteString("echo \"step ")
		b.WriteString(strconv.Itoa(id))
		b.WriteString("\"\n\n")
	}
	return b.String()
}
```

internal/cli/cmds.go
```
package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/oraclepack/internal/app"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/validate"
)

var validateCmd = &cobra.Command{
	Use:   "validate [pack.md]",
	Short: "Validate an oracle pack",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		out := cmd.OutOrStdout()
		data, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}
		p, err := pack.Parse(data)
		if err != nil {
			return err
		}
		if err := pack.Validate(p); err != nil {
			return err
		}
		findings, warning, err := pack.CheckPackScripts(p)
		if err != nil {
			return err
		}
		if warning != "" {
			fmt.Fprintf(out, "Warning: %s\n", warning)
		}
		if len(findings) > 0 {
			for _, finding := range findings {
				if finding.StepID != "" {
					fmt.Fprintf(out, "Step %s line %d: %s\n", finding.StepID, finding.Line, finding.Message)
				} else {
					fmt.Fprintf(out, "Line %d: %s\n", finding.Line, finding.Message)
				}
			}
			return fmt.Errorf("bash syntax validation failed")
		}
		cv := validate.CompositeValidator{}
		results := cv.ValidatePack(p)
		fmt.Fprintf(out, "Validated %d steps\n", len(results))
		for _, r := range results {
			fmt.Fprintf(out, "Step %s [%s] %s", r.StepID, r.ToolKind.Name(), r.Status)
			if r.Error != "" {
				fmt.Fprintf(out, " (%s)", r.Error)
			}
			fmt.Fprintln(out)
		}
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list [pack.md]",
	Short: "List steps in an oracle pack",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := app.Config{PackPath: args[0]}
		a := app.New(cfg)
		if err := a.LoadPack(); err != nil {
			return err
		}
		for _, s := range a.Pack.Steps {
			fmt.Printf("%s: %s\n", s.ID, s.OriginalLine)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(listCmd)
}
```

internal/cli/root.go
```
package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/oraclepack/internal/errors"
)

var (
	noTUI     bool
	oracleBin string
	outDir    string
)

var rootCmd = &cobra.Command{
	Use:   "oraclepack",
	Short: "Oracle Pack Runner",
	Long:  `A polished TUI-driven runner for oracle-based interactive bash steps.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(errors.ExitCode(err))
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&noTUI, "no-tui", false, "Disable the TUI and run in plain terminal mode")
	rootCmd.PersistentFlags().StringVar(&oracleBin, "oracle-bin", "oracle", "Path to the oracle binary")
	rootCmd.PersistentFlags().StringVarP(&outDir, "out-dir", "o", "", "Output directory for step execution")
}
```

internal/cli/run.go
```
package cli

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/oraclepack/internal/app"
	"github.com/user/oraclepack/internal/config"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/tui"
)

var (
	yes                   bool
	resume                bool
	stopOnFail            bool
	roiThreshold          float64
	roiMode               string
	runAll                bool
	outputVerify          bool
	outputRetries         int
	outputRequireHeadings bool
	outputChunkMode       string
)

var runCmd = &cobra.Command{
	Use:   "run [pack.md]",
	Short: "Run an oracle pack",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		packPath := args[0]

		// Setup paths
		base := strings.TrimSuffix(filepath.Base(packPath), filepath.Ext(packPath))
		statePath := base + ".state.json"
		reportPath := base + ".report.json"

		resolvedVerify, err := config.ResolveOutputVerify(outputVerify, cmd.Flags().Changed("output-verify"))
		if err != nil {
			return err
		}
		resolvedRetries, err := config.ResolveOutputRetries(outputRetries, cmd.Flags().Changed("output-retries"))
		if err != nil {
			return err
		}
		resolvedRequireHeadings, err := config.ResolveOutputRequireHeadings(outputRequireHeadings, cmd.Flags().Changed("output-require-headings"))
		if err != nil {
			return err
		}
		resolvedChunkMode, err := config.ResolveOutputChunkMode(outputChunkMode, cmd.Flags().Changed("output-chunk-mode"))
		if err != nil {
			return err
		}

		cfg := app.Config{
			PackPath:              packPath,
			StatePath:             statePath,
			ReportPath:            reportPath,
			Resume:                resume,
			StopOnFail:            stopOnFail,
			WorkDir:               ".",
			OutDir:                outDir,
			ROIThreshold:          roiThreshold,
			ROIMode:               roiMode,
			OutputVerify:          resolvedVerify,
			OutputRetries:         resolvedRetries,
			OutputRequireHeadings: resolvedRequireHeadings,
			OutputChunkMode:       resolvedChunkMode,
		}

		a := app.New(cfg)
		// Prepare the application (loads pack, resolves out_dir, provisions env)
		if err := a.Prepare(); err != nil {
			return err
		}

		if err := a.LoadState(); err != nil {
			return err
		}

		findings, warning, err := pack.CheckPackScripts(a.Pack)
		if err != nil {
			return err
		}
		if warning != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "Warning: %s\n", warning)
		}
		if len(findings) > 0 {
			for _, finding := range findings {
				if finding.StepID != "" {
					fmt.Fprintf(cmd.OutOrStdout(), "Step %s line %d: %s\n", finding.StepID, finding.Line, finding.Message)
				} else {
					fmt.Fprintf(cmd.OutOrStdout(), "Line %d: %s\n", finding.Line, finding.Message)
				}
			}
			return fmt.Errorf("bash syntax validation failed")
		}

		if noTUI {
			out := cmd.OutOrStdout()
			fmt.Fprintf(out, "[Selected] %s\n", packPath)
			fmt.Fprintln(out, "[Ready] Parsed and validated pack")
			err := a.RunPlain(context.Background(), out)
			if err != nil {
				fmt.Fprintf(out, "[Completed] Failed: %v\n", err)
				return err
			}
			fmt.Fprintln(out, "[Completed] Success")
			return nil
		}

		m := tui.NewModel(a.Pack, a.Runner, a.State, cfg.StatePath, cfg.ROIThreshold, cfg.ROIMode, runAll, cfg.OutputVerify, cfg.OutputRetries, cfg.OutputRequireHeadings, cfg.OutputChunkMode)
		p := tea.NewProgram(m, tea.WithAltScreen())
		_, err = p.Run()
		return err
	},
}

func init() {
	runCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Auto-approve all steps")
	runCmd.Flags().BoolVar(&resume, "resume", false, "Resume from last successful step")
	runCmd.Flags().BoolVar(&stopOnFail, "stop-on-fail", true, "Stop execution if a step fails")
	runCmd.Flags().Float64Var(&roiThreshold, "roi-threshold", 0.0, "Filter steps by ROI threshold")
	runCmd.Flags().StringVar(&roiMode, "roi-mode", "over", "ROI filter mode ('over' or 'under')")
	runCmd.Flags().BoolVar(&runAll, "run-all", false, "Automatically run all steps sequentially on start")
	runCmd.Flags().BoolVar(&outputVerify, "output-verify", config.DefaultOutputVerify, "Verify --write-output files contain required answer sections")
	runCmd.Flags().IntVar(&outputRetries, "output-retries", config.DefaultOutputRetries, "Retries for output verification failures")
	runCmd.Flags().BoolVar(&outputRequireHeadings, "output-require-headings", config.DefaultOutputRequireHeadings, "Require strict output headings when verifying outputs")
	runCmd.Flags().StringVar(&outputChunkMode, "output-chunk-mode", config.DefaultOutputChunkMode, "Output chunk verification mode: auto|single|multi")
	rootCmd.AddCommand(runCmd)
}
```

internal/cli/verify_outputs.go
```
package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/oraclepack/internal/config"
	"github.com/user/oraclepack/internal/pack"
)

var (
	verifyOutputsEnabled      bool
	verifyOutputsRequireHeads bool
	verifyOutputsChunkMode    string
)

var verifyOutputsCmd = &cobra.Command{
	Use:   "verify-outputs [pack.md]",
	Short: "Verify --write-output files without executing steps",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		out := cmd.OutOrStdout()
		data, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}

		p, err := pack.Parse(data)
		if err != nil {
			return err
		}
		if err := pack.Validate(p); err != nil {
			return err
		}

		verifyEnabled, err := config.ResolveOutputVerify(verifyOutputsEnabled, cmd.Flags().Changed("output-verify"))
		if err != nil {
			return err
		}
		if !verifyEnabled {
			fmt.Fprintln(out, "Output verification disabled (ORACLEPACK_OUTPUT_VERIFY=false).")
			return nil
		}
		requireHeadings, err := config.ResolveOutputRequireHeadings(verifyOutputsRequireHeads, cmd.Flags().Changed("output-require-headings"))
		if err != nil {
			return err
		}
		chunkMode, err := config.ResolveOutputChunkMode(verifyOutputsChunkMode, cmd.Flags().Changed("output-chunk-mode"))
		if err != nil {
			return err
		}

		report := pack.VerifyReport{
			TotalSteps: len(p.Steps),
		}

		for i := range p.Steps {
			step := &p.Steps[i]
			failures := pack.VerifyStepOutputs(step, requireHeadings, chunkMode)
			if len(failures) == 0 {
				continue
			}
			report.CheckedSteps++
			for _, failure := range failures {
				failure.StepID = step.ID
				report.Failures = append(report.Failures, failure)
			}
		}

		fmt.Fprint(out, pack.FormatVerifyReport(report))
		if len(report.Failures) > 0 {
			return fmt.Errorf("output verification failed")
		}
		return nil
	},
}

func init() {
	verifyOutputsCmd.Flags().BoolVar(&verifyOutputsEnabled, "output-verify", config.DefaultOutputVerify, "Verify --write-output files contain required answer sections")
	verifyOutputsCmd.Flags().BoolVar(&verifyOutputsRequireHeads, "output-require-headings", config.DefaultOutputRequireHeadings, "Require strict output headings when verifying outputs")
	verifyOutputsCmd.Flags().StringVar(&verifyOutputsChunkMode, "output-chunk-mode", config.DefaultOutputChunkMode, "Output chunk verification mode: auto|single|multi")
	rootCmd.AddCommand(verifyOutputsCmd)
}
```

internal/artifacts/contract.go
```
package artifacts

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/user/oraclepack/internal/foundation"
)

// Contract maps step IDs to required artifact paths.
type Contract map[string][]string

// DefaultContract returns the standard artifact contract.
func DefaultContract() Contract {
	base := ".oraclepack/ticketify"
	return Contract{
		"09": {filepath.Join(base, "next.json")},
		"10": {filepath.Join(base, "codex-implement.md")},
		"11": {filepath.Join(base, "codex-verify.md")},
		"12": {filepath.Join(base, "PR.md")},
	}
}

// EvaluateGates checks required artifacts for a given step.
func EvaluateGates(stepID string, contract Contract) error {
	paths, ok := contract[stepID]
	if !ok || len(paths) == 0 {
		return nil
	}
	var missing []string
	for _, p := range paths {
		info, err := os.Stat(p)
		if err != nil || info.IsDir() || info.Size() == 0 {
			missing = append(missing, p)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("%w: %v", foundation.ErrArtifactMissing, missing)
	}
	return nil
}
```

internal/artifacts/contract_test.go
```
package artifacts

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEvaluateGates(t *testing.T) {
	dir := t.TempDir()
	base := filepath.Join(dir, ".oraclepack", "ticketify")
	if err := os.MkdirAll(base, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	contract := Contract{
		"09": {filepath.Join(base, "next.json")},
	}

	// Missing file should error.
	if err := EvaluateGates("09", contract); err == nil {
		t.Fatal("expected missing artifact error")
	}

	// Create file and verify pass.
	path := filepath.Join(base, "next.json")
	if err := os.WriteFile(path, []byte("ok"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := EvaluateGates("09", contract); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
```

internal/config/defaults.go
```
package config

const (
	EnvOutputVerify          = "ORACLEPACK_OUTPUT_VERIFY"
	EnvOutputRetries         = "ORACLEPACK_OUTPUT_RETRIES"
	EnvOutputRequireHeadings = "ORACLEPACK_OUTPUT_REQUIRE_HEADINGS"
	EnvOutputChunkMode       = "ORACLEPACK_OUTPUT_CHUNK_MODE"
)

const (
	DefaultOutputVerify          = false
	DefaultOutputRetries         = 0
	DefaultOutputRequireHeadings = false
	DefaultOutputChunkMode       = "auto"
)
```

internal/config/resolve.go
```
package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ResolveOutputVerify applies precedence: CLI flag > env var > default.
func ResolveOutputVerify(flagValue bool, flagSet bool) (bool, error) {
	if flagSet {
		return flagValue, nil
	}
	if val, ok := os.LookupEnv(EnvOutputVerify); ok {
		parsed, err := parseBoolish(val)
		if err != nil {
			return DefaultOutputVerify, fmt.Errorf("invalid %s: %w", EnvOutputVerify, err)
		}
		return parsed, nil
	}
	return DefaultOutputVerify, nil
}

// ResolveOutputRetries applies precedence: CLI flag > env var > default.
func ResolveOutputRetries(flagValue int, flagSet bool) (int, error) {
	if flagSet {
		return flagValue, nil
	}
	if val, ok := os.LookupEnv(EnvOutputRetries); ok {
		parsed, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			return DefaultOutputRetries, fmt.Errorf("invalid %s: %w", EnvOutputRetries, err)
		}
		return parsed, nil
	}
	return DefaultOutputRetries, nil
}

// ResolveOutputRequireHeadings applies precedence: CLI flag > env var > default.
func ResolveOutputRequireHeadings(flagValue bool, flagSet bool) (bool, error) {
	if flagSet {
		return flagValue, nil
	}
	if val, ok := os.LookupEnv(EnvOutputRequireHeadings); ok {
		parsed, err := parseBoolish(val)
		if err != nil {
			return DefaultOutputRequireHeadings, fmt.Errorf("invalid %s: %w", EnvOutputRequireHeadings, err)
		}
		return parsed, nil
	}
	return DefaultOutputRequireHeadings, nil
}

// ResolveOutputChunkMode applies precedence: CLI flag > env var > default.
func ResolveOutputChunkMode(flagValue string, flagSet bool) (string, error) {
	if flagSet {
		return normalizeChunkMode(flagValue)
	}
	if val, ok := os.LookupEnv(EnvOutputChunkMode); ok {
		return normalizeChunkMode(val)
	}
	return normalizeChunkMode(DefaultOutputChunkMode)
}

func parseBoolish(raw string) (bool, error) {
	v := strings.ToLower(strings.TrimSpace(raw))
	switch v {
	case "1", "true", "yes", "on":
		return true, nil
	case "0", "false", "no", "off":
		return false, nil
	default:
		return false, fmt.Errorf("expected boolean (true/false, 1/0, on/off), got %q", raw)
	}
}

func normalizeChunkMode(raw string) (string, error) {
	v := strings.ToLower(strings.TrimSpace(raw))
	switch v {
	case "auto", "single", "multi":
		return v, nil
	case "":
		return DefaultOutputChunkMode, nil
	default:
		return "", fmt.Errorf("invalid %s: expected auto|single|multi, got %q", EnvOutputChunkMode, raw)
	}
}
```

internal/dispatch/classify.go
```
package dispatch

import (
	"regexp"
	"strings"

	"github.com/user/oraclepack/internal/tools"
)

var classifier = regexp.MustCompile(`^(\s*)(oracle|tm|task-master|codex|gemini)\b`)

// Classification describes a parsed command prefix.
type Classification struct {
	Kind    tools.ToolKind
	Prefix  string
	Command string
}

// Classify detects a supported tool prefix and returns the remaining command.
func Classify(line string) (Classification, bool) {
	m := classifier.FindStringSubmatch(line)
	if len(m) < 3 {
		return Classification{}, false
	}
	prefix := m[2]
	kind := toolKindFromPrefix(prefix)
	if kind == nil {
		return Classification{}, false
	}
	trimmed := strings.TrimSpace(line[len(m[1])+len(prefix):])
	return Classification{Kind: *kind, Prefix: prefix, Command: strings.TrimSpace(trimmed)}, true
}

func toolKindFromPrefix(prefix string) *tools.ToolKind {
	var kind tools.ToolKind
	switch prefix {
	case "oracle":
		kind = tools.ToolOracle
	case "tm":
		kind = tools.ToolTM
	case "task-master":
		kind = tools.ToolTaskMaster
	case "codex":
		kind = tools.ToolCodex
	case "gemini":
		kind = tools.ToolGemini
	default:
		return nil
	}
	return &kind
}
```

internal/dispatch/classify_test.go
```
package dispatch

import "testing"

func TestClassify(t *testing.T) {
	tests := []struct {
		line    string
		wantOK  bool
		wantCmd string
	}{
		{"oracle query \"hi\"", true, "query \"hi\""},
		{"  tm list", true, "list"},
		{"task-master next", true, "next"},
		{"codex exec \"x\"", true, "exec \"x\""},
		{"gemini run", true, "run"},
		{"echo hello", false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			got, ok := Classify(tt.line)
			if ok != tt.wantOK {
				t.Fatalf("expected ok=%v got %v", tt.wantOK, ok)
			}
			if ok && got.Command != tt.wantCmd {
				t.Fatalf("expected cmd %q got %q", tt.wantCmd, got.Command)
			}
		})
	}
}
```

internal/exec/flags.go
```
package exec

import "strings"

// ApplyChatGPTURL ensures a single --chatgpt-url flag is present when url is set.
// It removes any existing --chatgpt-url/--browser-url flags and their values.
func ApplyChatGPTURL(flags []string, url string) []string {
	var out []string
	skipNext := false
	for _, f := range flags {
		if skipNext {
			skipNext = false
			continue
		}
		if f == "--chatgpt-url" || f == "--browser-url" {
			skipNext = true
			continue
		}
		if strings.HasPrefix(f, "--chatgpt-url=") || strings.HasPrefix(f, "--browser-url=") {
			continue
		}
		out = append(out, f)
	}
	if url != "" {
		out = append(out, "--chatgpt-url", url)
	}
	return out
}
```

internal/exec/inject.go
```
package exec

import "strings"

// InjectFlags scans a script and appends flags to any 'oracle' command invocation.
func InjectFlags(script string, flags []string) string {
	if len(flags) == 0 {
		return script
	}

	flagStr := strings.Join(flags, " ")

	lines := strings.Split(script, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		insertIdx := oracleInsertIndex(line)
		if insertIdx == -1 {
			continue
		}

		lines[i] = insertFlagsInLine(line, insertIdx, flagStr)
	}

	return strings.Join(lines, "\n")
}

func oracleInsertIndex(line string) int {
	i := 0
	for i < len(line) && (line[i] == ' ' || line[i] == '\t') {
		i++
	}

	if !strings.HasPrefix(line[i:], "oracle") {
		return -1
	}

	end := i + len("oracle")
	if end < len(line) {
		next := line[end]
		if next != ' ' && next != '\t' {
			return -1
		}
	}

	return end
}

func insertFlagsInLine(line string, insertIdx int, flags string) string {
	prefix := line[:insertIdx]
	rest := line[insertIdx:]
	if rest == "" {
		return prefix + " " + flags
	}
	if rest[0] == ' ' || rest[0] == '\t' {
		return prefix + " " + flags + rest
	}
	return prefix + " " + flags + " " + rest
}
```

internal/exec/inject_test.go
```
package exec

import (
	"testing"
)

func TestInjectFlags(t *testing.T) {
	tests := []struct {
		name     string
		script   string
		flags    []string
		expected string
	}{
		{
			"simple injection",
			"oracle query 'hello'",
			[]string{"--verbose"},
			"oracle --verbose query 'hello'",
		},
		{
			"indented injection",
			"  oracle query 'hello'",
			[]string{"--verbose"},
			"  oracle --verbose query 'hello'",
		},
		{
			"no injection needed",
			"echo 'hello'",
			[]string{"--verbose"},
			"echo 'hello'",
		},
		{
			"multiple lines",
			"echo 'start'\noracle query\necho 'end'",
			[]string{"--debug"},
			"echo 'start'\noracle --debug query\necho 'end'",
		},
		{
			"multiline with continuation",
			"oracle \\\n  --json \\\n  --files",
			[]string{"--flag"},
			"oracle --flag \\\n  --json \\\n  --files",
		},
		{
			"multiline with args and continuation",
			"  oracle arg \\\n  --json",
			[]string{"--flag"},
			"  oracle --flag arg \\\n  --json",
		},
		{
			"commented command",
			"# oracle --json",
			[]string{"--verbose"},
			"# oracle --json",
		},
		{
			"oracle as part of word",
			"coracle query",
			[]string{"--verbose"},
			"coracle query",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InjectFlags(tt.script, tt.flags)
			if got != tt.expected {
				t.Errorf("InjectFlags() = %q, want %q", got, tt.expected)
			}
		})
	}
}
```

internal/exec/oracle_scan.go
```
package exec

import (
	"regexp"
	"strings"
)

var oracleCmdRegex = regexp.MustCompile(`^(\s*)(oracle)\b`)

// OracleInvocation represents a detected oracle command in a script.
type OracleInvocation struct {
	StartLine   int    // 0-based start line index
	EndLine     int    // 0-based end line index (inclusive)
	Raw         string // The full command string (joined if multi-line)
	Display     string // A trimmed version for UI display
	Indentation string // The leading whitespace
}

// ExtractOracleInvocations extracts oracle invocations from a script.
func ExtractOracleInvocations(script string) []OracleInvocation {
	var invocations []OracleInvocation
	lines := strings.Split(script, "\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)

		// Skip comments
		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		// Check for oracle command
		loc := oracleCmdRegex.FindStringSubmatchIndex(line)
		if loc != nil {
			startLine := i
			// Group 1 is the indentation
			indentation := line[loc[2]:loc[3]]

			var cmdBuilder strings.Builder
			cmdBuilder.WriteString(line)

			endLine := i
			// Handle line continuations
			// Check if line ends with backslash (ignoring trailing whitespace)
			for {
				if endLine+1 >= len(lines) {
					break
				}

				// Check current line for continuation
				currTrimmed := strings.TrimRight(lines[endLine], " \t")
				if !strings.HasSuffix(currTrimmed, "\\") {
					break
				}

				endLine++
				cmdBuilder.WriteString("\n")
				cmdBuilder.WriteString(lines[endLine])
			}

			raw := cmdBuilder.String()
			invocations = append(invocations, OracleInvocation{
				StartLine:   startLine,
				EndLine:     endLine,
				Raw:         raw,
				Display:     strings.TrimSpace(raw),
				Indentation: indentation,
			})

			i = endLine // Advance loop
		}
	}
	return invocations
}
```

internal/exec/oracle_scan_test.go
```
package exec

import (
	"reflect"
	"testing"
)

func TestExtractOracleInvocations(t *testing.T) {
	tests := []struct {
		name   string
		script string
		want   []OracleInvocation
	}{
		{
			name:   "Simple command",
			script: "oracle --json",
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 0, Raw: "oracle --json", Display: "oracle --json", Indentation: ""},
			},
		},
		{
			name:   "Indented command",
			script: "  oracle --json",
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 0, Raw: "  oracle --json", Display: "oracle --json", Indentation: "  "},
			},
		},
		{
			name: "Multiline command",
			script: `oracle \
  --json \
  --files`,
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 2, Raw: `oracle \
  --json \
  --files`, Display: `oracle \
  --json \
  --files`, Indentation: ""},
			},
		},
		{
			name: "Commented command",
			script: `# oracle --json
oracle --real`,
			want: []OracleInvocation{
				{StartLine: 1, EndLine: 1, Raw: "oracle --real", Display: "oracle --real", Indentation: ""},
			},
		},
		{
			name: "Multiple commands",
			script: `
echo start
oracle --one
echo mid
oracle --two
echo end
`,
			want: []OracleInvocation{
				{StartLine: 2, EndLine: 2, Raw: "oracle --one", Display: "oracle --one", Indentation: ""},
				{StartLine: 4, EndLine: 4, Raw: "oracle --two", Display: "oracle --two", Indentation: ""},
			},
		},
		{
			name:   "Oraclepack prefix (should not match)",
			script: "oraclepack run",
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractOracleInvocations(tt.script)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractOracleInvocations() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
```

internal/exec/oracle_validate.go
```
package exec

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/types"
)

// ValidationError captures a failed oracle validation for a step.
type ValidationError struct {
	StepID       string
	Command      string
	ErrorMessage string
}

// ValidateOverrides runs oracle --dry-run summary for targeted steps.
func ValidateOverrides(
	ctx context.Context,
	steps []types.Step,
	over *overrides.RuntimeOverrides,
	baseline []string,
	opts RunnerOptions,
) ([]ValidationError, error) {
	if over == nil || over.ApplyToSteps == nil {
		return nil, nil
	}

	shell := opts.Shell
	if shell == "" {
		shell = "/bin/bash"
	}
	env := append(os.Environ(), opts.Env...)

	var results []ValidationError
	for _, step := range steps {
		if !over.ApplyToSteps[step.ID] {
			continue
		}

		invocations := ExtractOracleInvocations(step.Code)
		if len(invocations) == 0 {
			continue
		}

		flags := over.EffectiveFlags(step.ID, baseline)
		flags = append(flags, "--dry-run", "summary")

		for _, inv := range invocations {
			cmdStr := InjectFlags(inv.Raw, flags)
			msg, err := execDryRun(ctx, shell, opts.WorkDir, env, cmdStr)
			if err == nil {
				continue
			}

			results = append(results, ValidationError{
				StepID:       step.ID,
				Command:      cmdStr,
				ErrorMessage: msg,
			})
		}
	}

	return results, nil
}

func execDryRun(ctx context.Context, shell, workDir string, env []string, command string) (string, error) {
	if pathVal := findEnvValue(env, "PATH"); pathVal != "" {
		command = "export PATH=" + shellQuote(pathVal) + "; " + command
	}

	cmd := exec.CommandContext(ctx, shell, "-lc", command)
	if workDir != "" {
		cmd.Dir = workDir
	}
	cmd.Env = env

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err == nil {
		return stdout.String(), nil
	}
	if stderr.Len() > 0 {
		return strings.TrimSpace(stderr.String()), err
	}
	if stdout.Len() > 0 {
		return strings.TrimSpace(stdout.String()), err
	}
	return err.Error(), err
}

func findEnvValue(env []string, key string) string {
	prefix := key + "="
	for _, entry := range env {
		if strings.HasPrefix(entry, prefix) {
			return strings.TrimPrefix(entry, prefix)
		}
	}
	return ""
}

func shellQuote(value string) string {
	if value == "" {
		return "''"
	}
	return "'" + strings.ReplaceAll(value, "'", "'\\''") + "'"
}
```

internal/exec/oracle_validate_test.go
```
package exec

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/types"
)

func TestValidateOverrides_Success(t *testing.T) {
	dir := t.TempDir()
	writeOracleStub(t, dir)

	steps := []types.Step{
		{ID: "01", Code: "oracle --ok"},
	}
	over := &overrides.RuntimeOverrides{
		ApplyToSteps: map[string]bool{"01": true},
	}

	_, err := ValidateOverrides(
		context.Background(),
		steps,
		over,
		[]string{"--base"},
		RunnerOptions{
			WorkDir: dir,
			Env:     []string{"PATH=" + dir + string(os.PathListSeparator) + os.Getenv("PATH")},
		},
	)
	if err != nil {
		t.Fatalf("ValidateOverrides failed: %v", err)
	}
}

func TestValidateOverrides_Error(t *testing.T) {
	dir := t.TempDir()
	writeOracleStub(t, dir)

	steps := []types.Step{
		{ID: "01", Code: "oracle --bad"},
	}
	over := &overrides.RuntimeOverrides{
		ApplyToSteps: map[string]bool{"01": true},
	}

	errs, err := ValidateOverrides(
		context.Background(),
		steps,
		over,
		nil,
		RunnerOptions{
			WorkDir: dir,
			Env:     []string{"PATH=" + dir + string(os.PathListSeparator) + os.Getenv("PATH")},
		},
	)
	if err != nil {
		t.Fatalf("ValidateOverrides failed: %v", err)
	}
	if len(errs) != 1 {
		t.Fatalf("expected 1 validation error, got %d", len(errs))
	}
	msg := errs[0].ErrorMessage
	if !strings.Contains(msg, "invalid flag") && !strings.Contains(msg, "unknown option") {
		t.Fatalf("unexpected error message: %q", msg)
	}
	if !strings.Contains(errs[0].Command, "--dry-run summary") {
		t.Fatalf("expected command to include --dry-run summary, got %q", errs[0].Command)
	}
}

func writeOracleStub(t *testing.T, dir string) {
	t.Helper()
	stub := `#!/bin/sh
has_dry=0
has_summary=0
for arg in "$@"; do
  if [ "$arg" = "--dry-run" ]; then has_dry=1; fi
  if [ "$arg" = "summary" ]; then has_summary=1; fi
  if [ "$arg" = "--bad" ]; then echo "invalid flag" 1>&2; exit 1; fi
done
if [ $has_dry -eq 0 ] || [ $has_summary -eq 0 ]; then
  echo "missing dry run" 1>&2
  exit 1
fi
exit 0
`
	path := filepath.Join(dir, "oracle")
	if err := os.WriteFile(path, []byte(stub), 0o755); err != nil {
		t.Fatalf("write oracle stub: %v", err)
	}
}
```

internal/exec/runner.go
```
package exec

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/user/oraclepack/internal/errors"
	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/types"
)

// Runner handles the execution of shell scripts.
type Runner struct {
	Shell       string
	WorkDir     string
	Env         []string
	OracleFlags []string
	Overrides   *overrides.RuntimeOverrides
	ChatGPTURL  string
	warnings    []SanitizeWarning
}

// RunnerOptions configures a Runner.
type RunnerOptions struct {
	Shell       string
	WorkDir     string
	Env         []string
	OracleFlags []string
	Overrides   *overrides.RuntimeOverrides
	ChatGPTURL  string
}

// NewRunner creates a new Runner with options.
func NewRunner(opts RunnerOptions) *Runner {
	shell := opts.Shell
	if shell == "" {
		shell = "/bin/bash"
	}

	return &Runner{
		Shell:       shell,
		WorkDir:     opts.WorkDir,
		Env:         append(os.Environ(), opts.Env...),
		OracleFlags: opts.OracleFlags,
		Overrides:   opts.Overrides,
		ChatGPTURL:  opts.ChatGPTURL,
	}
}

// RunPrelude executes the prelude code.
func (r *Runner) RunPrelude(ctx context.Context, p *types.Prelude, logWriter io.Writer) error {
	script, warnings := SanitizeScript(p.Code, "prelude", "")
	r.recordWarnings(warnings, logWriter)
	return r.run(ctx, script, logWriter)
}

// RunStep executes a single step's code.
func (r *Runner) RunStep(ctx context.Context, s *types.Step, logWriter io.Writer) error {
	flags := ApplyChatGPTURL(r.OracleFlags, r.ChatGPTURL)
	if r.Overrides != nil {
		flags = r.Overrides.EffectiveFlags(s.ID, r.OracleFlags)
		flags = ApplyChatGPTURL(flags, r.ChatGPTURL)
	}
	code := InjectFlags(s.Code, flags)
	script, warnings := SanitizeScript(code, "step", s.ID)
	r.recordWarnings(warnings, logWriter)
	return r.run(ctx, script, logWriter)
}

func (r *Runner) recordWarnings(warnings []SanitizeWarning, logWriter io.Writer) {
	if len(warnings) == 0 {
		return
	}
	for _, w := range warnings {
		r.warnings = append(r.warnings, w)
		if logWriter != nil {
			scope := w.Scope
			if scope == "" {
				scope = "script"
			}
			step := ""
			if w.StepID != "" {
				step = " step " + w.StepID
			}
			_, _ = fmt.Fprintf(logWriter, "⚠ oraclepack: sanitized label in %s%s line %d: %s\n", scope, step, w.Line, w.Token)
		}
	}
}

// DrainWarnings returns any sanitizer warnings collected since the last call.
func (r *Runner) DrainWarnings() []SanitizeWarning {
	if len(r.warnings) == 0 {
		return nil
	}
	out := make([]SanitizeWarning, len(r.warnings))
	copy(out, r.warnings)
	r.warnings = nil
	return out
}

func (r *Runner) run(ctx context.Context, script string, logWriter io.Writer) error {
	// We use bash -lc to ensure login shell (paths, aliases, etc)
	cmd := exec.CommandContext(ctx, r.Shell, "-lc", script)
	cmd.Dir = r.WorkDir
	cmd.Env = r.Env

	// Standardize stdout and stderr to the logWriter
	cmd.Stdout = logWriter
	cmd.Stderr = logWriter

	err := cmd.Run()
	if err != nil {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		return fmt.Errorf("%w: %v", errors.ErrExecutionFailed, err)
	}

	return nil
}
```

internal/exec/runner_test.go
```
package exec

import (
	"context"
	"strings"
	"testing"

	"github.com/user/oraclepack/internal/types"
)

func TestRunner_RunStep(t *testing.T) {
	r := NewRunner(RunnerOptions{})

	var lines []string
	lw := &LineWriter{
		Callback: func(line string) {
			lines = append(lines, line)
		},
	}

	step := &types.Step{
		Code: "echo 'hello world'",
	}

	err := r.RunStep(context.Background(), step, lw)
	if err != nil {
		t.Fatalf("RunStep failed: %v", err)
	}
	lw.Close()

	found := false
	for _, l := range lines {
		if strings.TrimSpace(l) == "hello world" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected 'hello world' in output, got: %v", lines)
	}
}

func TestRunner_ContextCancellation(t *testing.T) {
	r := NewRunner(RunnerOptions{})

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	step := &types.Step{
		Code: "sleep 10",
	}

	err := r.RunStep(ctx, step, nil)
	if err != context.Canceled {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}
```

internal/exec/sanitize.go
```
package exec

import (
	osexec "os/exec"
	"regexp"
	"strings"
)

// SanitizeWarning captures a label line that was converted to a safe echo.
type SanitizeWarning struct {
	Scope   string
	StepID  string
	Line    int
	Token   string
	Message string
}

var (
	labelTokenRegex   = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_-]*$`)
	heredocStartRegex = regexp.MustCompile(`<<-?\s*['"]?([A-Za-z0-9_]+)['"]?`)
)

var shellBuiltins = map[string]bool{
	"alias":    true,
	"bg":       true,
	"break":    true,
	"cd":       true,
	"command":  true,
	"continue": true,
	"declare":  true,
	"dirs":     true,
	"echo":     true,
	"eval":     true,
	"exec":     true,
	"exit":     true,
	"export":   true,
	"fg":       true,
	"getopts":  true,
	"hash":     true,
	"help":     true,
	"jobs":     true,
	"local":    true,
	"popd":     true,
	"printf":   true,
	"pushd":    true,
	"pwd":      true,
	"readonly": true,
	"return":   true,
	"set":      true,
	"shift":    true,
	"source":   true,
	"test":     true,
	"trap":     true,
	"true":     true,
	"type":     true,
	"ulimit":   true,
	"umask":    true,
	"unalias":  true,
	"unset":    true,
	"wait":     true,
	"false":    true,
}

var shellKeywords = map[string]bool{
	"case":     true,
	"do":       true,
	"done":     true,
	"elif":     true,
	"else":     true,
	"esac":     true,
	"fi":       true,
	"for":      true,
	"function": true,
	"if":       true,
	"in":       true,
	"select":   true,
	"then":     true,
	"time":     true,
	"until":    true,
	"while":    true,
}

// SanitizeScript converts bare label-like lines into safe echo statements.
func SanitizeScript(script, scope, stepID string) (string, []SanitizeWarning) {
	if script == "" {
		return script, nil
	}

	lines := strings.Split(script, "\n")
	var warnings []SanitizeWarning
	var heredocEnd string

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if heredocEnd != "" {
			if trimmed == heredocEnd {
				heredocEnd = ""
			}
			continue
		}
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		if end := heredocStartToken(trimmed); end != "" {
			heredocEnd = end
			continue
		}

		fields := strings.Fields(trimmed)
		if len(fields) != 1 {
			continue
		}
		token := fields[0]
		if !labelTokenRegex.MatchString(token) {
			continue
		}
		lower := strings.ToLower(token)
		if shellBuiltins[lower] || shellKeywords[lower] {
			continue
		}
		if _, err := osexec.LookPath(token); err == nil {
			continue
		}

		indent := line[:len(line)-len(strings.TrimLeft(line, " \t"))]
		lines[i] = indent + "echo \"" + token + "\""
		warnings = append(warnings, SanitizeWarning{
			Scope:   scope,
			StepID:  stepID,
			Line:    i + 1,
			Token:   token,
			Message: "Converted bare label to echo",
		})
	}

	return strings.Join(lines, "\n"), warnings
}

func heredocStartToken(line string) string {
	match := heredocStartRegex.FindStringSubmatch(line)
	if len(match) < 2 {
		return ""
	}
	return match[1]
}
```

internal/exec/sanitize_test.go
```
package exec

import "testing"

func TestSanitizeScript_LabelLine(t *testing.T) {
	input := "GenerateReport\noracle --help\n"
	got, warnings := SanitizeScript(input, "step", "01")
	if len(warnings) != 1 {
		t.Fatalf("expected 1 warning, got %d", len(warnings))
	}
	if warnings[0].Token != "GenerateReport" {
		t.Fatalf("expected token GenerateReport, got %s", warnings[0].Token)
	}
	wantPrefix := "echo \"GenerateReport\""
	if got[:len(wantPrefix)] != wantPrefix {
		t.Fatalf("expected sanitized line to start with %q, got %q", wantPrefix, got)
	}
}

func TestSanitizeScript_BuiltinUnchanged(t *testing.T) {
	input := "echo\n"
	got, warnings := SanitizeScript(input, "step", "01")
	if len(warnings) != 0 {
		t.Fatalf("expected no warnings, got %d", len(warnings))
	}
	if got != input {
		t.Fatalf("expected script unchanged, got %q", got)
	}
}

func TestSanitizeScript_HeredocUnchanged(t *testing.T) {
	input := "cat <<'EOF'\nGenerateReport\nEOF\n"
	got, warnings := SanitizeScript(input, "step", "01")
	if len(warnings) != 0 {
		t.Fatalf("expected no warnings, got %d", len(warnings))
	}
	if got != input {
		t.Fatalf("expected heredoc unchanged, got %q", got)
	}
}
```

internal/exec/stream.go
```
package exec

import (
	"io"
)

// LineWriter is an io.Writer that splits output into lines and calls a callback.
type LineWriter struct {
	Callback func(string)
	buffer   []byte
}

func (w *LineWriter) Write(p []byte) (n int, err error) {
	for _, b := range p {
		if b == '\n' {
			w.Callback(string(w.buffer))
			w.buffer = w.buffer[:0]
		} else {
			w.buffer = append(w.buffer, b)
		}
	}
	return len(p), nil
}

// Close flushes any remaining data in the buffer.
func (w *LineWriter) Close() error {
	if len(w.buffer) > 0 {
		w.Callback(string(w.buffer))
		w.buffer = w.buffer[:0]
	}
	return nil
}

// MultiWriter handles multiple writers efficiently.
func MultiWriter(writers ...io.Writer) io.Writer {
	return io.MultiWriter(writers...)
}
```

internal/errors/errors.go
```
package errors

import (
	"errors"
)

var (
	// ErrInvalidPack is returned when the Markdown pack is malformed.
	ErrInvalidPack = errors.New("invalid pack structure")
	// ErrExecutionFailed is returned when a shell command fails.
	ErrExecutionFailed = errors.New("execution failed")
	// ErrConfigInvalid is returned when CLI flags or environment variables are incorrect.
	ErrConfigInvalid = errors.New("invalid configuration")
)

// ExitCode returns the appropriate exit code for a given error.
func ExitCode(err error) int {
	if err == nil {
		return 0
	}

	if errors.Is(err, ErrConfigInvalid) {
		return 2
	}

	if errors.Is(err, ErrInvalidPack) {
		return 3
	}

	if errors.Is(err, ErrExecutionFailed) {
		return 4
	}

	return 1 // Generic error
}
```

internal/errors/errors_test.go
```
package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestExitCode(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected int
	}{
		{"nil error", nil, 0},
		{"generic error", errors.New("generic"), 1},
		{"invalid pack", ErrInvalidPack, 3},
		{"execution failed", ErrExecutionFailed, 4},
		{"config invalid", ErrConfigInvalid, 2},
		{"wrapped invalid pack", fmt.Errorf("wrap: %w", ErrInvalidPack), 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExitCode(tt.err); got != tt.expected {
				t.Errorf("ExitCode() = %v, want %v", got, tt.expected)
			}
		})
	}
}
```

internal/foundation/atomic.go
```
package foundation

import (
	"fmt"
	"os"
)

// WriteAtomic writes data to path atomically by writing to a temp file and renaming.
func WriteAtomic(path string, data []byte, perm os.FileMode) error {
	tempPath := path + ".tmp"
	if err := os.WriteFile(tempPath, data, perm); err != nil {
		return fmt.Errorf("write temp file: %w", err)
	}
	if err := os.Rename(tempPath, path); err != nil {
		_ = os.Remove(tempPath)
		return fmt.Errorf("rename temp file: %w", err)
	}
	return nil
}
```

internal/foundation/atomic_test.go
```
package foundation

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAtomic(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.json")
	if err := WriteAtomic(path, []byte("hello"), 0644); err != nil {
		t.Fatalf("WriteAtomic: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if string(data) != "hello" {
		t.Fatalf("unexpected contents: %q", string(data))
	}
}
```

internal/foundation/clock.go
```
package foundation

import "time"

// Clock abstracts time for deterministic testing.
type Clock interface {
	Now() time.Time
}

// RealClock uses the system clock.
type RealClock struct{}

// Now returns the current time.
func (RealClock) Now() time.Time { return time.Now() }

// MockClock returns a fixed time that can be advanced.
type MockClock struct {
	current time.Time
}

// NewMockClock initializes a mock clock with a starting time.
func NewMockClock(start time.Time) *MockClock {
	return &MockClock{current: start}
}

// Now returns the mock time.
func (m *MockClock) Now() time.Time { return m.current }

// Advance moves the mock time forward.
func (m *MockClock) Advance(d time.Duration) {
	m.current = m.current.Add(d)
}
```

internal/foundation/clock_test.go
```
package foundation

import (
	"testing"
	"time"
)

func TestMockClock(t *testing.T) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	m := NewMockClock(start)
	if !m.Now().Equal(start) {
		t.Fatalf("expected %v, got %v", start, m.Now())
	}
	m.Advance(2 * time.Hour)
	want := start.Add(2 * time.Hour)
	if !m.Now().Equal(want) {
		t.Fatalf("expected %v, got %v", want, m.Now())
	}
}
```

internal/foundation/config.go
```
package foundation

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Config holds runtime settings that can be loaded from JSON and environment variables.
// Env values always take precedence over JSON values.
type Config struct {
	Name      string  `json:"name" env:"ORACLEPACK_NAME"`
	Retries   int     `json:"retries" env:"ORACLEPACK_RETRIES"`
	Enabled   bool    `json:"enabled" env:"ORACLEPACK_ENABLED"`
	Threshold float64 `json:"threshold" env:"ORACLEPACK_THRESHOLD"`
}

// LoadConfig loads configuration from a JSON file and then applies environment overrides.
// If path is empty, JSON loading is skipped and only env overrides are applied.
func LoadConfig(path string) (Config, error) {
	var cfg Config
	if path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return Config{}, fmt.Errorf("read config: %w", err)
		}
		if err := json.Unmarshal(data, &cfg); err != nil {
			return Config{}, fmt.Errorf("parse config: %w", err)
		}
	}

	if v, ok := os.LookupEnv("ORACLEPACK_NAME"); ok {
		cfg.Name = v
	}
	if v, ok := os.LookupEnv("ORACLEPACK_RETRIES"); ok {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return Config{}, fmt.Errorf("parse ORACLEPACK_RETRIES: %w", err)
		}
		cfg.Retries = parsed
	}
	if v, ok := os.LookupEnv("ORACLEPACK_ENABLED"); ok {
		parsed, err := strconv.ParseBool(v)
		if err != nil {
			return Config{}, fmt.Errorf("parse ORACLEPACK_ENABLED: %w", err)
		}
		cfg.Enabled = parsed
	}
	if v, ok := os.LookupEnv("ORACLEPACK_THRESHOLD"); ok {
		parsed, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return Config{}, fmt.Errorf("parse ORACLEPACK_THRESHOLD: %w", err)
		}
		cfg.Threshold = parsed
	}

	return cfg, nil
}
```

internal/foundation/config_test.go
```
package foundation

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigEnvOverrides(t *testing.T) {
	t.Setenv("ORACLEPACK_NAME", "env-name")
	t.Setenv("ORACLEPACK_RETRIES", "5")
	t.Setenv("ORACLEPACK_ENABLED", "true")
	t.Setenv("ORACLEPACK_THRESHOLD", "2.5")

	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	if err := os.WriteFile(path, []byte(`{"name":"json-name","retries":1,"enabled":false,"threshold":1.0}`), 0644); err != nil {
		t.Fatalf("write json: %v", err)
	}

	cfg, err := LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig: %v", err)
	}

	if cfg.Name != "env-name" || cfg.Retries != 5 || cfg.Enabled != true || cfg.Threshold != 2.5 {
		t.Fatalf("env overrides not applied: %+v", cfg)
	}
}

func TestLoadConfigJSONOnly(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	if err := os.WriteFile(path, []byte(`{"name":"json-name","retries":3,"enabled":true,"threshold":4.25}`), 0644); err != nil {
		t.Fatalf("write json: %v", err)
	}

	cfg, err := LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig: %v", err)
	}

	if cfg.Name != "json-name" || cfg.Retries != 3 || cfg.Enabled != true || cfg.Threshold != 4.25 {
		t.Fatalf("json load mismatch: %+v", cfg)
	}
}
```

internal/foundation/errors.go
```
package foundation

import "errors"

var (
	// ErrMissingBinary is returned when a required binary is not found on PATH.
	ErrMissingBinary = errors.New("missing binary")
	// ErrArtifactMissing is returned when an expected artifact is absent.
	ErrArtifactMissing = errors.New("artifact missing")
)
```

internal/foundation/errors_test.go
```
package foundation

import (
	"errors"
	"testing"
)

func TestCommonErrors(t *testing.T) {
	if ErrMissingBinary == nil || ErrArtifactMissing == nil {
		t.Fatal("expected error variables to be initialized")
	}
	if !errors.Is(ErrMissingBinary, ErrMissingBinary) {
		t.Fatal("errors.Is failed for ErrMissingBinary")
	}
	if !errors.Is(ErrArtifactMissing, ErrArtifactMissing) {
		t.Fatal("errors.Is failed for ErrArtifactMissing")
	}
}
```

internal/overrides/merge.go
```
package overrides

// EffectiveFlags calculates the final flags for a step.
func (r *RuntimeOverrides) EffectiveFlags(stepID string, baseline []string) []string {
	if r == nil || r.ApplyToSteps == nil || !r.ApplyToSteps[stepID] {
		return baseline
	}

	var effective []string

	// Map for removed flags
	removed := make(map[string]bool)
	for _, f := range r.RemovedFlags {
		removed[f] = true
	}

	// Filter baseline
	for _, flag := range baseline {
		if !removed[flag] {
			effective = append(effective, flag)
		}
	}

	// Append added flags
	effective = append(effective, r.AddedFlags...)

	// Inject ChatGPTURL
	if r.ChatGPTURL != "" {
		effective = append(effective, "--chatgpt-url", r.ChatGPTURL)
	}

	return effective
}
```

internal/overrides/merge_test.go
```
package overrides

import (
	"reflect"
	"testing"
)

func TestEffectiveFlags(t *testing.T) {
	tests := []struct {
		name      string
		overrides *RuntimeOverrides
		stepID    string
		baseline  []string
		want      []string
	}{
		{
			name:      "No overrides (nil)",
			overrides: nil,
			stepID:    "01",
			baseline:  []string{"--json"},
			want:      []string{"--json"},
		},
		{
			name: "Step not targeted",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"02": true},
				AddedFlags:   []string{"--verbose"},
			},
			stepID:   "01",
			baseline: []string{"--json"},
			want:     []string{"--json"},
		},
		{
			name: "Step targeted: Add flags",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"01": true},
				AddedFlags:   []string{"--verbose"},
			},
			stepID:   "01",
			baseline: []string{"--json"},
			want:     []string{"--json", "--verbose"},
		},
		{
			name: "Step targeted: Remove flags",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"01": true},
				RemovedFlags: []string{"--json"},
			},
			stepID:   "01",
			baseline: []string{"--json", "--other"},
			want:     []string{"--other"},
		},
		{
			name: "Step targeted: Add and Remove",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"01": true},
				AddedFlags:   []string{"--new"},
				RemovedFlags: []string{"--old"},
			},
			stepID:   "01",
			baseline: []string{"--old", "--keep"},
			want:     []string{"--keep", "--new"},
		},
		{
			name: "Step targeted: Inject ChatGPT URL",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"01": true},
				ChatGPTURL:   "https://chat.openai.com/share/123",
			},
			stepID:   "01",
			baseline: []string{"--json"},
			want:     []string{"--json", "--chatgpt-url", "https://chat.openai.com/share/123"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.overrides.EffectiveFlags(tt.stepID, tt.baseline)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EffectiveFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

internal/overrides/types.go
```
package overrides

// RuntimeOverrides holds configuration for runtime flag modifications.
type RuntimeOverrides struct {
	AddedFlags   []string        // Flags to append (e.g., "--model=gpt-4")
	RemovedFlags []string        // Flags to remove (e.g., "--json")
	ChatGPTURL   string          // Optional URL to inject via --chatgpt-url
	ApplyToSteps map[string]bool // Set of step IDs to apply overrides to. If empty, applies to none.
}
```

internal/pack/bash_syntax_validator.go
```
package pack

import (
	"regexp"
	"strings"

	"github.com/user/oraclepack/internal/types"
)

var orphanFlagRegex = regexp.MustCompile(`^\s*(?:-p|--prompt)(?:\s+.+)?$`)

// FindOrphanedFlags detects lines that contain only flags like -p/--prompt
// without being part of a continued command.
func FindOrphanedFlags(script string) []types.SyntaxFinding {
	if script == "" {
		return nil
	}

	lines := strings.Split(script, "\n")
	var findings []types.SyntaxFinding
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		if !orphanFlagRegex.MatchString(line) {
			continue
		}

		if i > 0 && lineContinues(lines[i-1]) {
			continue
		}

		findings = append(findings, types.SyntaxFinding{
			Line:    i + 1,
			Token:   strings.Fields(trimmed)[0],
			Message: "Orphaned flag without a preceding command or line continuation",
		})
	}
	return findings
}

func lineContinues(line string) bool {
	trimmed := strings.TrimRight(line, " \t")
	return strings.HasSuffix(trimmed, "\\")
}

// CheckPackScripts validates prelude and step scripts for orphaned flags and bash syntax.
func CheckPackScripts(p *types.Pack) ([]types.SyntaxFinding, string, error) {
	if p == nil {
		return nil, "", nil
	}

	var findings []types.SyntaxFinding
	var warnings []string

	// Prelude orphaned flags
	for _, finding := range FindOrphanedFlags(p.Prelude.Code) {
		findings = append(findings, finding)
	}

	// Step orphaned flags and bash -n checks
	for _, step := range p.Steps {
		for _, finding := range FindOrphanedFlags(step.Code) {
			finding.StepID = step.ID
			findings = append(findings, finding)
		}

		syntaxFindings, warning, err := CheckBashSyntax(step.Code)
		if err != nil {
			return findings, warning, err
		}
		if warning != "" {
			warnings = append(warnings, warning)
		}
		for _, finding := range syntaxFindings {
			finding.StepID = step.ID
			findings = append(findings, finding)
		}
	}

	// Prelude bash -n check
	preludeFindings, warning, err := CheckBashSyntax(p.Prelude.Code)
	if err != nil {
		return findings, warning, err
	}
	if warning != "" {
		warnings = append(warnings, warning)
	}
	findings = append(findings, preludeFindings...)

	return findings, strings.Join(uniqueStrings(warnings), "; "), nil
}

func uniqueStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := make(map[string]bool, len(values))
	var out []string
	for _, value := range values {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		out = append(out, value)
	}
	return out
}
```

internal/pack/bash_syntax_validator_test.go
```
package pack

import (
	"os/exec"
	"testing"

	"github.com/user/oraclepack/internal/types"
)

func TestFindOrphanedFlags(t *testing.T) {
	script := "-p \"hello\"\n"
	findings := FindOrphanedFlags(script)
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(findings))
	}
	if findings[0].Line != 1 {
		t.Fatalf("expected line 1, got %d", findings[0].Line)
	}

	script = "oracle \\\n  -p \"ok\"\n"
	findings = FindOrphanedFlags(script)
	if len(findings) != 0 {
		t.Fatalf("expected no findings for continued line, got %d", len(findings))
	}
}

func TestCheckBashSyntax(t *testing.T) {
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("bash not available")
	}

	script := "if true\n  echo hello\n"
	findings, warning, err := CheckBashSyntax(script)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if warning != "" {
		t.Fatalf("expected no warning, got %s", warning)
	}
	if len(findings) == 0 {
		t.Fatal("expected syntax findings")
	}
	if findings[0].Line == 0 {
		t.Fatal("expected line number in syntax finding")
	}
}

func TestCheckPackScriptsReportsStepID(t *testing.T) {
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", Code: "-p \"oops\""},
		},
	}
	findings, _, err := CheckPackScripts(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(findings))
	}
	if findings[0].StepID != "01" {
		t.Fatalf("expected step ID 01, got %q", findings[0].StepID)
	}
}
```

internal/pack/bash_tooling_checks.go
```
package pack

import (
	"bytes"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/user/oraclepack/internal/types"
)

var bashLineRegex = regexp.MustCompile(`line\s+(\d+)`)

// CheckBashSyntax runs "bash -n" on the script and returns syntax findings.
// If bash is not found, it returns a warning and no findings.
func CheckBashSyntax(script string) ([]types.SyntaxFinding, string, error) {
	if strings.TrimSpace(script) == "" {
		return nil, "", nil
	}

	if _, err := exec.LookPath("bash"); err != nil {
		return nil, "bash not found on PATH; skipping bash -n syntax check", nil
	}

	cmd := exec.Command("bash", "-n")
	cmd.Stdin = strings.NewReader(script)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stderr

	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(stderr.String())
		line := 0
		if match := bashLineRegex.FindStringSubmatch(msg); len(match) > 1 {
			if parsed, parseErr := strconv.Atoi(match[1]); parseErr == nil {
				line = parsed
			}
		}
		return []types.SyntaxFinding{
			{
				Line:    line,
				Message: msg,
			},
		}, "", nil
	}

	return nil, "", nil
}
```

internal/pack/metadata.go
```
package pack

import (
	"fmt"

	"github.com/user/oraclepack/internal/errors"
	"github.com/user/oraclepack/internal/types"
)

// DeriveMetadata extracts configuration from the prelude.
func DeriveMetadata(p *types.Pack) {
	if p == nil {
		return
	}
	outDirMatch := outDirRegex.FindStringSubmatch(p.Prelude.Code)
	if len(outDirMatch) > 1 {
		p.OutDir = outDirMatch[1]
	}

	if writeOutputRegex.MatchString(p.Prelude.Code) {
		p.WriteOutput = true
	}
}

// Validate checks if the pack follows all rules.
func Validate(p *types.Pack) error {
	if p == nil {
		return fmt.Errorf("%w: pack is nil", errors.ErrInvalidPack)
	}
	if len(p.Steps) == 0 {
		return fmt.Errorf("%w: at least one step is required", errors.ErrInvalidPack)
	}
	if len(p.Steps) != 20 {
		return fmt.Errorf("%w: expected exactly 20 steps, got %d", errors.ErrInvalidPack, len(p.Steps))
	}

	seen := make(map[int]bool)
	for i, step := range p.Steps {
		if step.Number <= 0 {
			return fmt.Errorf("%w: invalid step number %d", errors.ErrInvalidPack, step.Number)
		}
		if seen[step.Number] {
			return fmt.Errorf("%w: duplicate step number %d", errors.ErrInvalidPack, step.Number)
		}
		seen[step.Number] = true

		// Ensure sequential starting from 1
		if step.Number != i+1 {
			return fmt.Errorf("%w: steps must be sequential starting from 1 (expected %d, got %d)", errors.ErrInvalidPack, i+1, step.Number)
		}
	}

	return nil
}
```

internal/pack/output_expectations.go
```
package pack

import (
	"regexp"
	"strings"

	"github.com/user/oraclepack/internal/types"
)

var (
	writeOutputPathRegex = regexp.MustCompile(`(?m)--write-output\s+["']?([^"'\s]+)["']?`)
	answerFormatRegex    = regexp.MustCompile(`(?i)answer\s+format`)
	directOnlyRegex      = regexp.MustCompile(`(?i)return\s+only[:\s]*direct\s+answer`)
)

// DetectOutputContract determines the expected response contract for a step.
func DetectOutputContract(step types.Step) types.OutputContract {
	if step.Code == "" {
		return types.OutputContractUnknown
	}

	hasAnswerFormat := answerFormatRegex.MatchString(step.Code)
	if !hasAnswerFormat {
		return types.OutputContractUnknown
	}

	if directOnlyRegex.MatchString(step.Code) {
		return types.OutputContractDirectAnswerOnly
	}

	return types.OutputContractAllSections
}

// StepOutputExpectations returns a map of output paths to required tokens.
// If no validation is needed, it returns nil.
func StepOutputExpectations(step *types.Step) map[string][]string {
	if step == nil {
		return nil
	}
	paths := ExtractWriteOutputPaths(step.Code)
	if len(paths) == 0 {
		return nil
	}

	if len(paths) > 1 {
		out := map[string][]string{}
		for _, path := range paths {
			switch {
			case strings.Contains(path, "-direct-answer"):
				out[path] = []string{"### Direct answer"}
			case strings.Contains(path, "-risks-unknowns"):
				out[path] = []string{"### Risks and unknowns"}
			case strings.Contains(path, "-next-experiment"):
				out[path] = []string{"### Next experiment"}
			case strings.Contains(path, "-missing-evidence"):
				out[path] = []string{"### Missing evidence"}
			}
		}
		if len(out) == 0 {
			return nil
		}
		return out
	}

	switch DetectOutputContract(*step) {
	case types.OutputContractDirectAnswerOnly:
		return map[string][]string{paths[0]: []string{"### Direct answer"}}
	case types.OutputContractAllSections:
		return map[string][]string{paths[0]: {
			"### Direct answer",
			"### Risks and unknowns",
			"### Next experiment",
			"### Missing evidence",
		}}
	default:
		return nil
	}
}

// StepOutputExpectationsWithMode returns expectations honoring chunk mode.
// chunkMode: auto (default), single (treat as single output), multi (force suffix mapping when multiple outputs).
func StepOutputExpectationsWithMode(step *types.Step, chunkMode string) map[string][]string {
	if step == nil {
		return nil
	}
	paths := ExtractWriteOutputPaths(step.Code)
	if len(paths) == 0 {
		return nil
	}
	mode := strings.ToLower(strings.TrimSpace(chunkMode))
	if mode == "" {
		mode = "auto"
	}

	switch mode {
	case "single":
		// Always treat as a single output (use first path).
		return expectationsForSingle(step, paths[0])
	case "multi":
		if len(paths) > 1 {
			return expectationsForSuffixes(paths)
		}
		return expectationsForSingle(step, paths[0])
	default: // auto
		if len(paths) > 1 {
			return expectationsForSuffixes(paths)
		}
		return expectationsForSingle(step, paths[0])
	}
}

func expectationsForSuffixes(paths []string) map[string][]string {
	out := map[string][]string{}
	for _, path := range paths {
		switch {
		case strings.Contains(path, "-direct-answer"):
			out[path] = []string{"### Direct answer"}
		case strings.Contains(path, "-risks-unknowns"):
			out[path] = []string{"### Risks and unknowns"}
		case strings.Contains(path, "-next-experiment"):
			out[path] = []string{"### Next experiment"}
		case strings.Contains(path, "-missing-evidence"):
			out[path] = []string{"### Missing evidence"}
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func expectationsForSingle(step *types.Step, path string) map[string][]string {
	switch DetectOutputContract(*step) {
	case types.OutputContractDirectAnswerOnly:
		return map[string][]string{path: []string{"### Direct answer"}}
	case types.OutputContractAllSections:
		return map[string][]string{path: {
			"### Direct answer",
			"### Risks and unknowns",
			"### Next experiment",
			"### Missing evidence",
		}}
	default:
		return nil
	}
}

// ExtractWriteOutputPaths returns all --write-output paths found in the step code.
func ExtractWriteOutputPaths(code string) []string {
	matches := writeOutputPathRegex.FindAllStringSubmatch(code, -1)
	if len(matches) == 0 {
		return nil
	}
	paths := make([]string, 0, len(matches))
	for _, m := range matches {
		if len(m) >= 2 {
			paths = append(paths, m[1])
		}
	}
	return paths
}
```

internal/pack/output_expectations_test.go
```
package pack

import (
	"reflect"
	"testing"

	"github.com/user/oraclepack/internal/types"
)

func TestDetectOutputContract(t *testing.T) {
	step := types.Step{
		Code: "Answer format:\nReturn only: Direct answer\n",
	}
	if got := DetectOutputContract(step); got != types.OutputContractDirectAnswerOnly {
		t.Fatalf("expected direct-answer-only, got %q", got)
	}

	step = types.Step{
		Code: "Answer format:\n### Direct answer\n### Risks and unknowns\n",
	}
	if got := DetectOutputContract(step); got != types.OutputContractAllSections {
		t.Fatalf("expected all-sections, got %q", got)
	}

	step = types.Step{Code: "no format here"}
	if got := DetectOutputContract(step); got != types.OutputContractUnknown {
		t.Fatalf("expected unknown, got %q", got)
	}
}

func TestStepOutputExpectations_SingleOutput(t *testing.T) {
	step := &types.Step{
		Code: `oracle -p "x" --write-output "out.txt"
Answer format:
Return only: Direct answer`,
	}
	expectations := StepOutputExpectations(step)
	want := map[string][]string{"out.txt": []string{"### Direct answer"}}
	if !reflect.DeepEqual(expectations, want) {
		t.Fatalf("unexpected expectations: %#v", expectations)
	}

	step = &types.Step{
		Code: `oracle -p "x" --write-output "out.txt"
Answer format:
### Direct answer
### Risks and unknowns
### Next experiment
### Missing evidence`,
	}
	expectations = StepOutputExpectations(step)
	want = map[string][]string{"out.txt": []string{
		"### Direct answer",
		"### Risks and unknowns",
		"### Next experiment",
		"### Missing evidence",
	}}
	if !reflect.DeepEqual(expectations, want) {
		t.Fatalf("unexpected expectations: %#v", expectations)
	}
}

func TestStepOutputExpectations_MultiOutput(t *testing.T) {
	step := &types.Step{
		Code: `oracle --write-output "out-direct-answer.md" \
  --write-output 'out-risks-unknowns.md' \
  --write-output out-next-experiment.md \
  --write-output out-missing-evidence.md`,
	}
	expectations := StepOutputExpectations(step)
	if len(expectations) != 4 {
		t.Fatalf("expected 4 expectations, got %d", len(expectations))
	}
	if got := expectations["out-direct-answer.md"]; len(got) != 1 || got[0] != "### Direct answer" {
		t.Fatalf("unexpected direct-answer tokens: %#v", got)
	}
	if got := expectations["out-risks-unknowns.md"]; len(got) != 1 || got[0] != "### Risks and unknowns" {
		t.Fatalf("unexpected risks-unknowns tokens: %#v", got)
	}
	if got := expectations["out-next-experiment.md"]; len(got) != 1 || got[0] != "### Next experiment" {
		t.Fatalf("unexpected next-experiment tokens: %#v", got)
	}
	if got := expectations["out-missing-evidence.md"]; len(got) != 1 || got[0] != "### Missing evidence" {
		t.Fatalf("unexpected missing-evidence tokens: %#v", got)
	}
}
```

internal/pack/output_validator.go
```
package pack

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/user/oraclepack/internal/types"
)

// ValidateOutputFile checks whether the output file contains the required answer sections.
// It returns ok=false with a populated OutputFailure when validation fails.
func ValidateOutputFile(path string, requiredTokens []string) (bool, types.OutputFailure) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, types.OutputFailure{
			Path:  path,
			Error: err.Error(),
		}
	}

	content := string(data)
	var missing []string
	for _, tok := range requiredTokens {
		if !containsToken(content, tok) {
			missing = append(missing, tok)
		}
	}

	if len(missing) > 0 {
		return false, types.OutputFailure{
			Path:          path,
			MissingTokens: missing,
		}
	}

	return true, types.OutputFailure{}
}

func containsToken(content, token string) bool {
	if strings.Contains(content, token) {
		return true
	}

	heading := strings.TrimSpace(strings.TrimPrefix(token, "###"))
	if heading == token {
		return false
	}

	alts := []string{heading}
	switch strings.ToLower(heading) {
	case "direct answer":
		alts = append(alts, "answer")
	case "risks and unknowns":
		alts = append(alts, "risks/unknowns", "risks & unknowns")
	case "next experiment":
		alts = append(alts, "next smallest concrete experiment")
	case "missing evidence":
		alts = append(alts, "if evidence is insufficient")
	}

	for _, alt := range alts {
		if alt == "" {
			continue
		}
		pat := `(?im)^\s*#{0,3}\s*` + regexp.QuoteMeta(alt) + `\b`
		if regexp.MustCompile(pat).MatchString(content) {
			return true
		}
	}

	return false
}

// VerifyStepOutputs validates output files for a step. If requireHeadings is false,
// it only checks that output files exist and are non-empty.
func VerifyStepOutputs(step *types.Step, requireHeadings bool, chunkMode string) []types.OutputFailure {
	if step == nil {
		return nil
	}
	paths := ExtractWriteOutputPaths(step.Code)
	if len(paths) == 0 {
		return nil
	}

	if !requireHeadings {
		paths = selectChunkPaths(paths, chunkMode)
		var failures []types.OutputFailure
		for _, path := range paths {
			info, err := os.Stat(path)
			if err != nil {
				failures = append(failures, types.OutputFailure{
					Path:  path,
					Error: err.Error(),
				})
				continue
			}
			if info.Size() == 0 {
				failures = append(failures, types.OutputFailure{
					Path:  path,
					Error: fmt.Sprintf("output file is empty: %s", path),
				})
			}
		}
		return failures
	}

	expectations := StepOutputExpectationsWithMode(step, chunkMode)
	if len(expectations) == 0 {
		return nil
	}
	var failures []types.OutputFailure
	for path, required := range expectations {
		ok, failure := ValidateOutputFile(path, required)
		if !ok {
			failures = append(failures, failure)
		}
	}
	return failures
}

func selectChunkPaths(paths []string, chunkMode string) []string {
	if len(paths) == 0 {
		return paths
	}
	mode := strings.ToLower(strings.TrimSpace(chunkMode))
	if mode == "" {
		mode = "auto"
	}
	switch mode {
	case "single":
		return []string{paths[0]}
	default:
		return paths
	}
}
```

internal/pack/output_validator_test.go
```
package pack

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/user/oraclepack/internal/types"
)

func TestValidateOutputFile_RelaxedHeadings(t *testing.T) {
	content := `Direct answer

Risks/unknowns

Next smallest concrete experiment

If evidence is insufficient`

	dir := t.TempDir()
	path := filepath.Join(dir, "out.md")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write output: %v", err)
	}

	ok, failure := ValidateOutputFile(path, []string{
		"### Direct answer",
		"### Risks and unknowns",
		"### Next experiment",
		"### Missing evidence",
	})
	if !ok {
		t.Fatalf("expected relaxed headings to pass, failure: %#v", failure)
	}

	step := &types.Step{Code: `oracle --write-output "` + path + `"`}
	failures := VerifyStepOutputs(step, true, "single")
	if len(failures) != 0 {
		t.Fatalf("expected no failures in single chunk mode, got %#v", failures)
	}
}
```

internal/pack/parser.go
```
package pack

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/user/oraclepack/internal/errors"
	"github.com/user/oraclepack/internal/types"
)

var (
	bashFenceRegex = regexp.MustCompile("(?s)```bash\n(.*?)\n```")
	// Updated regex to support ")", " —", and " -" separators
	stepHeaderRegex  = regexp.MustCompile(`^#\s*(\d{2})(?:\)|[\s]+[—-])`)
	roiRegex         = regexp.MustCompile(`ROI=(\d+(\.\d+)?)`)
	impactRegex      = regexp.MustCompile(`^#\s*Impact:\s*(.+)$`)
	outDirRegex      = regexp.MustCompile(`(?m)^out_dir=["']?([^"'\s]+)["']?`)
	writeOutputRegex = regexp.MustCompile(`(?m)--write-output`)
)

// Parse reads a Markdown content and returns a Pack.
func Parse(content []byte) (*types.Pack, error) {
	match := bashFenceRegex.FindSubmatch(content)
	if match == nil || len(match) < 2 {
		return nil, fmt.Errorf("%w: no bash code block found", errors.ErrInvalidPack)
	}

	bashCode := string(match[1])
	pack := &types.Pack{}

	scanner := bufio.NewScanner(strings.NewReader(bashCode))
	var currentStep *types.Step
	var preludeLines []string
	var inSteps bool

	for scanner.Scan() {
		line := scanner.Text()
		headerMatch := stepHeaderRegex.FindStringSubmatch(strings.TrimSpace(line))

		if len(headerMatch) > 1 {
			inSteps = true
			if currentStep != nil {
				applyStepMetadata(currentStep)
				pack.Steps = append(pack.Steps, *currentStep)
			}
			num, _ := strconv.Atoi(headerMatch[1])

			// Extract ROI if present
			var roi float64
			cleanedLine := line
			roiMatch := roiRegex.FindStringSubmatch(line)
			if len(roiMatch) > 1 {
				val, err := strconv.ParseFloat(roiMatch[1], 64)
				if err == nil {
					roi = val
					// Remove ROI tag from display title, but keep original line intact?
					// The task says "strip from Step.Title". Step struct currently has `OriginalLine`.
					// I'll assume OriginalLine is what is displayed, or I should add a Title field.
					// Looking at Step struct: ID, Number, Code, OriginalLine.
					// I'll remove it from OriginalLine for now or add a Title field.
					// The existing TUI uses OriginalLine as description.
					// Let's clean OriginalLine for display purposes or add a dedicated Title field.
					// Adding a dedicated Title field seems cleaner but requires struct change.
					// For now, I'll strip it from OriginalLine to match the prompt requirement "cleaner UI display".
					cleanedLine = strings.Replace(cleanedLine, roiMatch[0], "", 1)
					cleanedLine = strings.TrimSpace(cleanedLine)
					// Fix any double spaces or trailing separators if needed, but simple replace is a good start.
				}
			}

			currentStep = &types.Step{
				ID:           headerMatch[1],
				Number:       num,
				OriginalLine: cleanedLine,
				ROI:          roi,
			}
			continue
		}

		if inSteps {
			currentStep.Code += line + "\n"
		} else {
			preludeLines = append(preludeLines, line)
		}
	}

	if currentStep != nil {
		applyStepMetadata(currentStep)
		pack.Steps = append(pack.Steps, *currentStep)
	}

	pack.Prelude.Code = strings.Join(preludeLines, "\n")
	DeriveMetadata(pack)

	return pack, nil
}

func applyStepMetadata(step *types.Step) {
	if step == nil {
		return
	}
	lines := strings.Split(step.Code, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || !strings.HasPrefix(trimmed, "#") {
			continue
		}
		if step.ROI == 0 {
			if strings.HasPrefix(trimmed, "# ROI:") {
				val := strings.TrimSpace(strings.TrimPrefix(trimmed, "# ROI:"))
				if parsed, err := strconv.ParseFloat(val, 64); err == nil {
					step.ROI = parsed
				}
			}
		}
		if step.Impact == "" {
			if m := impactRegex.FindStringSubmatch(trimmed); len(m) > 1 {
				step.Impact = strings.TrimSpace(m[1])
			}
		}
	}
}
```

internal/pack/parser_test.go
```
package pack

import (
	"strconv"
	"strings"
	"testing"

	"github.com/user/oraclepack/internal/types"
)

func TestParse(t *testing.T) {
	steps := buildSteps(20, "echo")
	content := []byte(`
# My Pack
Some description.

` + "```" + `bash
out_dir="dist"
--write-output

` + steps + `
` + "```" + `
`)

	p, err := Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if p.OutDir != "dist" {
		t.Errorf("expected OutDir dist, got %s", p.OutDir)
	}

	if !p.WriteOutput {
		t.Errorf("expected WriteOutput true, got false")
	}

	if len(p.Steps) != 20 {
		t.Errorf("expected 20 steps, got %d", len(p.Steps))
	}

	if p.Steps[0].ID != "01" || p.Steps[0].Number != 1 {
		t.Errorf("step 1 mismatch: %+v", p.Steps[0])
	}

	if err := Validate(p); err != nil {
		t.Errorf("Validate failed: %v", err)
	}
}

func TestParseVariants(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{
			"em dash",
			`
` + "```" + `bash
# 01 — ROI=...
echo "step 1"
` + "```" + `
`,
		},
		{
			"hyphen",
			`
` + "```" + `bash
# 01 - ROI=...
echo "step 1"
` + "```" + `
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := Parse([]byte(tt.content))
			if err != nil {
				t.Fatalf("Parse failed: %v", err)
			}
			if len(p.Steps) != 1 {
				t.Errorf("expected 1 step, got %d", len(p.Steps))
			}
		})
	}
}

func TestParseROI(t *testing.T) {
	content := []byte(`
` + "```" + `bash
# 01) ROI=4.5 clean me
echo "high value"

# 02) ROI=0.5
echo "low value"

# 03) No ROI
echo "default"
` + "```" + `
`)

	p, err := Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(p.Steps) != 3 {
		t.Fatalf("expected 3 steps, got %d", len(p.Steps))
	}

	if p.Steps[0].ROI != 4.5 {
		t.Errorf("step 1 ROI mismatch: expected 4.5, got %f", p.Steps[0].ROI)
	}
	if strings.Contains(p.Steps[0].OriginalLine, "ROI=4.5") {
		t.Errorf("step 1 title was not cleaned: %q", p.Steps[0].OriginalLine)
	}

	if p.Steps[1].ROI != 0.5 {
		t.Errorf("step 2 ROI mismatch: expected 0.5, got %f", p.Steps[1].ROI)
	}

	if p.Steps[2].ROI != 0.0 {
		t.Errorf("step 3 ROI mismatch: expected 0.0, got %f", p.Steps[2].ROI)
	}
}

func TestValidateErrors(t *testing.T) {
	base := buildStepSlice(20)
	tests := []struct {
		name    string
		pack    *types.Pack
		wantErr string
	}{
		{
			"no steps",
			&types.Pack{},
			"at least one step is required",
		},
		{
			"duplicate steps",
			&types.Pack{
				Steps: func() []types.Step {
					steps := append([]types.Step(nil), base...)
					steps[1].Number = steps[0].Number
					steps[1].ID = steps[0].ID
					return steps
				}(),
			},
			"duplicate step number 1",
		},
		{
			"wrong step count",
			&types.Pack{
				Steps: []types.Step{
					{Number: 1, ID: "01"},
				},
			},
			"expected exactly 20 steps",
		},
		{
			"non-sequential",
			&types.Pack{
				Steps: func() []types.Step {
					steps := append([]types.Step(nil), base...)
					steps[1].Number = 3
					steps[1].ID = "03"
					return steps
				}(),
			},
			"steps must be sequential starting from 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.pack)
			if err == nil {
				t.Error("expected error, got nil")
			} else if !contains(err.Error(), tt.wantErr) {
				t.Errorf("expected error containing %q, got %q", tt.wantErr, err.Error())
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || (len(substr) > 0 && (s[:len(substr)] == substr || contains(s[1:], substr))))
}

func buildSteps(count int, cmd string) string {
[TRUNCATED]
```

internal/pack/verify_report.go
```
package pack

import (
	"fmt"
	"strings"

	"github.com/user/oraclepack/internal/types"
)

// VerifyReport captures output verification results across a pack.
type VerifyReport struct {
	TotalSteps   int
	CheckedSteps int
	Failures     []types.OutputFailure
}

// FormatVerifyReport renders a human-readable report.
func FormatVerifyReport(report VerifyReport) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Verified outputs for %d/%d steps\n", report.CheckedSteps, report.TotalSteps)
	if len(report.Failures) == 0 {
		b.WriteString("All required output tokens were found.\n")
		return b.String()
	}

	b.WriteString("Missing or invalid outputs:\n")
	for _, failure := range report.Failures {
		stepLabel := failure.StepID
		if stepLabel == "" {
			stepLabel = "unknown step"
		}
		fmt.Fprintf(&b, "- Step %s: %s", stepLabel, failure.Path)
		if failure.Error != "" {
			fmt.Fprintf(&b, " (error: %s)", failure.Error)
		} else if len(failure.MissingTokens) > 0 {
			fmt.Fprintf(&b, " missing %s", strings.Join(failure.MissingTokens, ", "))
		}
		b.WriteString("\n")
	}
	return b.String()
}
```

internal/render/render.go
```
package render

import (
	"sync"

	"github.com/charmbracelet/glamour"
	"github.com/user/oraclepack/internal/types"
)

const (
	DefaultStyle = "dark"
	DefaultWidth = 80
)

type rendererKey struct {
	width int
	style string
}

var (
	rendererMu    sync.Mutex
	rendererCache = map[rendererKey]*glamour.TermRenderer{}
)

// RenderMarkdown renders markdown text as ANSI-styled text.
func RenderMarkdown(text string, width int, style string) (string, error) {
	if width <= 0 {
		width = DefaultWidth
	}
	if style == "" {
		style = DefaultStyle
	}

	r, err := rendererFor(width, style)
	if err != nil {
		return "", err
	}

	return r.Render(text)
}

// RenderStepCode renders a step's code block for preview.
func RenderStepCode(s types.Step, width int, style string) (string, error) {
	md := "```bash\n" + s.Code + "\n```"
	return RenderMarkdown(md, width, style)
}

func rendererFor(width int, style string) (*glamour.TermRenderer, error) {
	key := rendererKey{width: width, style: style}

	rendererMu.Lock()
	r := rendererCache[key]
	rendererMu.Unlock()
	if r != nil {
		return r, nil
	}

	opts := []glamour.TermRendererOption{glamour.WithWordWrap(width)}
	if style == "auto" {
		opts = append(opts, glamour.WithAutoStyle())
	} else {
		opts = append(opts, glamour.WithStandardStyle(style))
	}

	r, err := glamour.NewTermRenderer(opts...)
	if err != nil {
		return nil, err
	}

	rendererMu.Lock()
	rendererCache[key] = r
	rendererMu.Unlock()
	return r, nil
}
```

internal/render/render_test.go
```
package render

import (
	"strings"
	"testing"
)

func TestRenderMarkdown(t *testing.T) {
	text := "# Hello\n**bold**"
	got, err := RenderMarkdown(text, 40, DefaultStyle)
	if err != nil {
		t.Fatalf("RenderMarkdown failed: %v", err)
	}

	// ANSI escape codes start with \x1b[
	if !strings.Contains(got, "\x1b[") {
		t.Errorf("expected ANSI codes in output, got: %q", got)
	}
}
```

internal/report/generate.go
```
package report

import (
	"time"

	"github.com/user/oraclepack/internal/state"
)

// GenerateReport creates a ReportV1 from a RunState.
func GenerateReport(s *state.RunState, packName string) *ReportV1 {
	report := &ReportV1{
		PackInfo: PackInfo{
			Name: packName,
			Hash: s.PackHash,
		},
		GeneratedAt: time.Now(),
		Steps:       []StepReport{},
	}

	var totalDuration time.Duration
	success, failure, skipped := 0, 0, 0

	for id, status := range s.StepStatuses {
		duration := status.EndedAt.Sub(status.StartedAt)
		if status.EndedAt.IsZero() || status.StartedAt.IsZero() {
			duration = 0
		}

		totalDuration += duration

		sr := StepReport{
			ID:         id,
			Status:     string(status.Status),
			ExitCode:   status.ExitCode,
			Duration:   duration,
			DurationMs: duration.Milliseconds(),
			Error:      status.Error,
		}
		report.Steps = append(report.Steps, sr)

		switch status.Status {
		case state.StatusSuccess:
			success++
		case state.StatusFailed:
			failure++
		case state.StatusSkipped:
			skipped++
		}
	}

	report.Summary = Summary{
		TotalSteps:      len(s.StepStatuses),
		SuccessCount:    success,
		FailureCount:    failure,
		SkippedCount:    skipped,
		TotalDuration:   totalDuration,
		TotalDurationMs: totalDuration.Milliseconds(),
	}

	if len(s.Warnings) > 0 {
		report.Warnings = make([]Warning, 0, len(s.Warnings))
		for _, w := range s.Warnings {
			report.Warnings = append(report.Warnings, Warning{
				Scope:   w.Scope,
				StepID:  w.StepID,
				Line:    w.Line,
				Token:   w.Token,
				Message: w.Message,
			})
		}
	}

	return report
}
```

internal/report/io.go
```
package report

import (
	"encoding/json"
	"fmt"
	"os"
)

// WriteReport writes a ReportV1 to disk.
func WriteReport(path string, rep *ReportV1) error {
	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal report: %w", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write report: %w", err)
	}
	return nil
}
```

internal/report/io_test.go
```
package report

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteReport(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "report.json")

	rep := &ReportV1{
		PackInfo: PackInfo{Name: "pack"},
		Summary:  Summary{TotalSteps: 1},
	}
	if err := WriteReport(path, rep); err != nil {
		t.Fatalf("WriteReport: %v", err)
	}

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected report file to exist: %v", err)
	}
}
```

internal/report/report_test.go
```
package report

import (
	"testing"
	"time"

	"github.com/user/oraclepack/internal/state"
)

func TestGenerateReport(t *testing.T) {
	s := &state.RunState{
		PackHash: "hash123",
		StepStatuses: map[string]state.StepStatus{
			"01": {
				Status:    state.StatusSuccess,
				StartedAt: time.Now().Add(-1 * time.Second),
				EndedAt:   time.Now(),
			},
		},
	}

	rep := GenerateReport(s, "my-pack")

	if rep.PackInfo.Name != "my-pack" {
		t.Errorf("expected name my-pack, got %s", rep.PackInfo.Name)
	}

	if rep.Summary.TotalSteps != 1 {
		t.Errorf("expected 1 total step, got %d", rep.Summary.TotalSteps)
	}

	if rep.Summary.SuccessCount != 1 {
		t.Errorf("expected 1 success, got %d", rep.Summary.SuccessCount)
	}
}
```

internal/report/types.go
```
package report

import (
	"time"
)

// ReportV1 represents the final machine-readable summary.
type ReportV1 struct {
	Summary     Summary      `json:"summary"`
	PackInfo    PackInfo     `json:"pack_info"`
	Steps       []StepReport `json:"steps"`
	Warnings    []Warning    `json:"warnings,omitempty"`
	GeneratedAt time.Time    `json:"generated_at"`
}

type Summary struct {
	TotalSteps      int           `json:"total_steps"`
	SuccessCount    int           `json:"success_count"`
	FailureCount    int           `json:"failure_count"`
	SkippedCount    int           `json:"skipped_count"`
	TotalDuration   time.Duration `json:"total_duration"`
	TotalDurationMs int64         `json:"total_duration_ms"`
}

type PackInfo struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

type StepReport struct {
	ID         string        `json:"id"`
	Status     string        `json:"status"`
	ExitCode   int           `json:"exit_code"`
	Duration   time.Duration `json:"duration"`
	DurationMs int64         `json:"duration_ms"`
	Error      string        `json:"error,omitempty"`
}

// Warning captures non-fatal execution notes surfaced during a run.
type Warning struct {
	Scope   string `json:"scope"`
	StepID  string `json:"step_id,omitempty"`
	Line    int    `json:"line"`
	Token   string `json:"token"`
	Message string `json:"message"`
}
```

internal/shell/detect.go
```
package shell

import "os/exec"

// DetectBinary checks PATH for a binary and returns its full path if found.
func DetectBinary(name string) (string, bool) {
	path, err := exec.LookPath(name)
	if err != nil {
		return "", false
	}
	return path, true
}
```

internal/shell/detect_test.go
```
package shell

import "testing"

func TestDetectBinary(t *testing.T) {
	if _, ok := DetectBinary("ls"); !ok {
		t.Fatalf("expected to find ls on PATH")
	}
	if _, ok := DetectBinary("definitely-not-a-real-binary-123"); ok {
		t.Fatalf("expected missing binary to return false")
	}
}
```

internal/shell/engine.go
```
package shell

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/user/oraclepack/internal/dispatch"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
	"github.com/user/oraclepack/internal/types"
)

// Engine executes pack steps in headless mode.
type Engine struct {
	Pack       *types.Pack
	State      *state.RunState
	StatePath  string
	StopOnFail bool
	Timeout    time.Duration
	Checker    tools.PresenceChecker
}

// Run executes all steps sequentially, updating state on disk.
func (e *Engine) Run(ctx context.Context) error {
	if e.Pack == nil {
		return nil
	}
	if e.State == nil {
		e.State = &state.RunState{
			SchemaVersion: 1,
			StepStatuses:  make(map[string]state.StepStatus),
		}
	}
	if e.State.StepStatuses == nil {
		e.State.StepStatuses = make(map[string]state.StepStatus)
	}

	for i := range e.Pack.Steps {
		step := e.Pack.Steps[i]
		e.State.CurrentStep = step.Number
		status := state.StepStatus{Status: state.StatusRunning, StartedAt: time.Now()}
		e.State.StepStatuses[step.ID] = status
		_ = state.WriteState(e.StatePath, e.State)

		kind := detectToolKind(&step)
		if shouldSkipForMissingTool(kind, e.Checker) {
			status.Status = state.StatusSkipped
			status.Error = "tool missing"
			status.EndedAt = time.Now()
			e.State.StepStatuses[step.ID] = status
			_ = state.WriteState(e.StatePath, e.State)
			continue
		}

		stepCtx := ctx
		if e.Timeout > 0 {
			var cancel context.CancelFunc
			stepCtx, cancel = context.WithTimeout(ctx, e.Timeout)
			defer cancel()
		}

		res, err := RunCommand(stepCtx, step.Code)
		if err == nil && res.ExitCode != 0 {
			err = fmt.Errorf("command failed with exit code %d", res.ExitCode)
		}
		status.EndedAt = time.Now()
		if err != nil {
			status.Status = state.StatusFailed
			status.Error = err.Error()
			e.State.StepStatuses[step.ID] = status
			_ = state.WriteState(e.StatePath, e.State)
			if e.StopOnFail {
				return err
			}
			continue
		}
		status.Status = state.StatusSuccess
		e.State.StepStatuses[step.ID] = status
		_ = state.WriteState(e.StatePath, e.State)
	}
	return nil
}

func detectToolKind(step *types.Step) tools.ToolKind {
	if step == nil {
		return tools.ToolUnknown
	}
	lines := strings.Split(step.Code, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		if cls, ok := dispatch.Classify(trimmed); ok {
			return cls.Kind
		}
		break
	}
	return tools.ToolUnknown
}

func shouldSkipForMissingTool(kind tools.ToolKind, checker tools.PresenceChecker) bool {
	if kind == tools.ToolUnknown {
		return false
	}
	meta, ok := tools.Metadata(kind)
	if !ok {
		return false
	}
	if checker == nil {
		_, found := DetectBinary(meta.Name)
		return !found
	}
	_, found := checker.DetectBinary(meta.Name)
	return !found
}
```

internal/shell/engine_test.go
```
package shell

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/types"
)

type fakeChecker struct {
	found map[string]bool
}

func (f fakeChecker) DetectBinary(name string) (string, bool) {
	return name, f.found[name]
}

func TestEngineSkipsMissingTool(t *testing.T) {
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", Number: 1, Code: "codex exec \"hi\""},
		},
	}
	dir := t.TempDir()
	engine := &Engine{
		Pack:      p,
		State:     &state.RunState{SchemaVersion: 1, StepStatuses: map[string]state.StepStatus{}},
		StatePath: filepath.Join(dir, "state.json"),
		Checker:   fakeChecker{found: map[string]bool{"codex": false}},
	}
	if err := engine.Run(context.Background()); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if engine.State.StepStatuses["01"].Status != state.StatusSkipped {
		t.Fatalf("expected skipped, got %s", engine.State.StepStatuses["01"].Status)
	}
}

func TestEngineFailsOnError(t *testing.T) {
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", Number: 1, Code: "exit 1"},
		},
	}
	engine := &Engine{
		Pack:       p,
		State:      &state.RunState{SchemaVersion: 1, StepStatuses: map[string]state.StepStatus{}},
		StopOnFail: true,
		Checker:    fakeChecker{found: map[string]bool{}},
	}
	if err := engine.Run(context.Background()); err == nil {
		t.Fatal("expected error, got nil")
	}
	if engine.State.StepStatuses["01"].Status != state.StatusFailed {
		t.Fatalf("expected failed, got %s", engine.State.StepStatuses["01"].Status)
	}
}
```

internal/shell/runner.go
```
package shell

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

// Result captures command output and exit code.
type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

// RunCommand executes a command with login shell semantics using /bin/bash -lc.
func RunCommand(ctx context.Context, cmd string) (Result, error) {
	c := exec.CommandContext(ctx, "/bin/bash", "-lc", cmd)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr

	err := c.Run()
	exitCode := 0
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			exitCode = ee.ExitCode()
		} else {
			return Result{}, fmt.Errorf("run command: %w", err)
		}
	}

	return Result{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: exitCode,
	}, nil
}
```

internal/shell/runner_test.go
```
package shell

import (
	"context"
	"strings"
	"testing"
)

func TestRunCommandLoginShell(t *testing.T) {
	res, err := RunCommand(context.Background(), "echo $PATH")
	if err != nil {
		t.Fatalf("RunCommand: %v", err)
	}
	if strings.TrimSpace(res.Stdout) == "" {
		t.Fatalf("expected PATH output, got empty stdout")
	}
	if res.ExitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", res.ExitCode)
	}
}
```

internal/state/io.go
```
package state

// Intentionally left without extra imports.

// WriteState writes RunState atomically to disk.
func WriteState(path string, state *RunState) error {
	return SaveStateAtomic(path, state)
}
```

internal/state/io_test.go
```
package state

import (
	"path/filepath"
	"testing"
	"time"
)

func TestWriteStateAndLoadState(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "state.json")

	input := &RunState{
		SchemaVersion: 1,
		PackHash:      "hash",
		StartTime:     time.Now(),
		CurrentStep:   2,
		StepStatuses: map[string]StepStatus{
			"01": {Status: StatusSuccess},
		},
	}

	if err := WriteState(path, input); err != nil {
		t.Fatalf("WriteState: %v", err)
	}

	out, err := LoadState(path)
	if err != nil {
		t.Fatalf("LoadState: %v", err)
	}

	if out.CurrentStep != 2 {
		t.Fatalf("expected CurrentStep 2, got %d", out.CurrentStep)
	}
	if out.StepStatuses["01"].Status != StatusSuccess {
		t.Fatalf("expected status success, got %s", out.StepStatuses["01"].Status)
	}
}
```

internal/state/persist.go
```
package state

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveStateAtomic saves the state to a file atomically.
func SaveStateAtomic(path string, state *RunState) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal state: %w", err)
	}

	tempPath := path + ".tmp"
	if err := os.WriteFile(tempPath, data, 0644); err != nil {
		return fmt.Errorf("write temp file: %w", err)
	}

	if err := os.Rename(tempPath, path); err != nil {
		os.Remove(tempPath)
		return fmt.Errorf("rename temp file: %w", err)
	}

	return nil
}

// LoadState loads the state from a file.
func LoadState(path string) (*RunState, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("state file not found: %w", err)
		}
		return nil, fmt.Errorf("read state file: %w", err)
	}

	var state RunState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("unmarshal state: %w", err)
	}

	return &state, nil
}
```

internal/state/state_test.go
```
package state

import (
	"os"
	"testing"
)

func TestStatePersistence(t *testing.T) {
	tmpFile := "test_state.json"
	defer os.Remove(tmpFile)

	s := &RunState{
		SchemaVersion: 1,
		PackHash:      "abc",
		StepStatuses: map[string]StepStatus{
			"01": {Status: StatusSuccess, ExitCode: 0},
		},
	}

	if err := SaveStateAtomic(tmpFile, s); err != nil {
		t.Fatalf("SaveStateAtomic failed: %v", err)
	}

	loaded, err := LoadState(tmpFile)
	if err != nil {
		t.Fatalf("LoadState failed: %v", err)
	}

	if loaded.PackHash != s.PackHash {
		t.Errorf("expected hash %s, got %s", s.PackHash, loaded.PackHash)
	}

	if loaded.StepStatuses["01"].Status != StatusSuccess {
		t.Errorf("expected status success, got %s", loaded.StepStatuses["01"].Status)
	}
}
```

internal/state/types.go
```
package state

import (
	"time"
)

type Status string

const (
	StatusPending Status = "pending"
	StatusRunning Status = "running"
	StatusSuccess Status = "success"
	StatusFailed  Status = "failed"
	StatusSkipped Status = "skipped"
)

// RunState tracks the execution progress of an oracle pack.
type RunState struct {
	SchemaVersion int                   `json:"schema_version"`
	PackHash      string                `json:"pack_hash"`
	StartTime     time.Time             `json:"start_time"`
	CurrentStep   int                   `json:"current_step,omitempty"`
	StepStatuses  map[string]StepStatus `json:"step_statuses"`
	ROIThreshold  float64               `json:"roi_threshold,omitempty"`
	ROIMode       string                `json:"roi_mode,omitempty"`
	Warnings      []Warning             `json:"warnings,omitempty"`
}

// StepStatus holds the outcome of an individual step.
type StepStatus struct {
	Status    Status    `json:"status"`
	ExitCode  int       `json:"exit_code"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	Error     string    `json:"error,omitempty"`
}

// Warning captures a non-fatal execution note (e.g., sanitized labels).
type Warning struct {
	Scope   string `json:"scope"`
	StepID  string `json:"step_id,omitempty"`
	Line    int    `json:"line"`
	Token   string `json:"token"`
	Message string `json:"message"`
}
```

internal/templates/template_test.go
```
package templates

import (
	"os"
	"testing"

	"github.com/user/oraclepack/internal/pack"
)

func TestRenderTicketActionPack(t *testing.T) {
	got := RenderTicketActionPack()
	if got == "" {
		t.Fatal("expected non-empty template")
	}

	// Golden comparison
	data, err := os.ReadFile("ticket-action-pack.md")
	if err != nil {
		t.Fatalf("read template: %v", err)
	}
	if string(data) != got {
		t.Fatalf("template mismatch with golden file")
	}

	// Ensure pack is parseable and validates 20-step contract.
	p, err := pack.Parse([]byte(got))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if err := pack.Validate(p); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}
	if len(p.Steps) != 20 {
		t.Fatalf("expected 20 steps, got %d", len(p.Steps))
	}
}
```

internal/templates/ticket-action-pack.md
```
# Ticket Action Pack

```bash
out_dir=".oraclepack/ticketify"
--write-output

# 01)
echo "Build tickets index"

# 02)
echo "Generate actions json"

# 03)
echo "Generate tickets PRD"

# 04)
echo "Prep taskmaster inputs"

# 05)
task-master parse-prd .taskmaster/docs/prd.md

# 06)
task-master analyze-complexity --research

# 07)
task-master expand --all --research

# 08)
echo "Prepare headless automation"

# 09)
if command -v gemini >/dev/null 2>&1; then
  gemini run "Select next tasks" --write-output ".oraclepack/ticketify/next.json"
else
  echo "Skipped: gemini missing"
fi

# 10)
if command -v codex >/dev/null 2>&1; then
  codex exec "Implement tasks" --write-output ".oraclepack/ticketify/codex-implement.md"
else
  echo "Skipped: codex missing"
fi

# 11)
if command -v codex >/dev/null 2>&1; then
  codex exec "Verify changes" --write-output ".oraclepack/ticketify/codex-verify.md"
else
  echo "Skipped: codex missing"
fi

# 12)
if command -v gemini >/dev/null 2>&1; then
  gemini run "Review outputs" --write-output ".oraclepack/ticketify/gemini-review.json"
else
  echo "Skipped: gemini missing"
fi

# 13)
if command -v codex >/dev/null 2>&1; then
  codex exec "Prepare fixes" --write-output ".oraclepack/ticketify/codex-fixes.md"
else
  echo "Skipped: codex missing"
fi

# 14)
echo "Summarize results"

# 15)
echo "Prepare release notes"

# 16)
if command -v codex >/dev/null 2>&1; then
  codex exec "Draft PR description" --write-output ".oraclepack/ticketify/PR.md"
else
  echo "Skipped: codex missing"
fi

# 17)
echo "Finalize checklist"

# 18)
echo "Post-run cleanup"

# 19)
echo "Audit artifacts"

# 20)
echo "Done"
```
```

internal/templates/ticket_action_pack.go
```
package templates

import _ "embed"

//go:embed ticket-action-pack.md
var ticketActionPack string

// RenderTicketActionPack returns the canonical ticket action pack template.
func RenderTicketActionPack() string {
	return ticketActionPack
}
```

internal/tools/types.go
```
package tools

// ToolKind identifies a supported tool prefix.
type ToolKind int

const (
	ToolUnknown ToolKind = iota
	ToolOracle
	ToolTM
	ToolTaskMaster
	ToolCodex
	ToolGemini
)

// ToolMetadata captures tool invocation details.
type ToolMetadata struct {
	Name string
	Args []string
}

var registry = map[ToolKind]ToolMetadata{
	ToolUnknown:    {Name: "unknown"},
	ToolOracle:     {Name: "oracle"},
	ToolTM:         {Name: "tm"},
	ToolTaskMaster: {Name: "task-master"},
	ToolCodex:      {Name: "codex", Args: []string{"exec"}},
	ToolGemini:     {Name: "gemini"},
}

// Metadata returns tool metadata if present.
func Metadata(kind ToolKind) (ToolMetadata, bool) {
	meta, ok := registry[kind]
	return meta, ok
}

// Name returns the canonical tool name.
func (k ToolKind) Name() string {
	if meta, ok := registry[k]; ok {
		return meta.Name
	}
	return "unknown"
}

// PresenceChecker abstracts binary detection.
type PresenceChecker interface {
	DetectBinary(name string) (string, bool)
}
```

internal/tools/types_test.go
```
package tools

import "testing"

func TestMetadataRegistry(t *testing.T) {
	meta, ok := Metadata(ToolCodex)
	if !ok {
		t.Fatalf("expected metadata for codex")
	}
	if meta.Name != "codex" {
		t.Fatalf("expected codex name, got %s", meta.Name)
	}
	if len(meta.Args) != 1 || meta.Args[0] != "exec" {
		t.Fatalf("expected codex exec args, got %+v", meta.Args)
	}
	if ToolOracle.Name() != "oracle" {
		t.Fatalf("expected oracle name, got %s", ToolOracle.Name())
	}
}
```

internal/tui/clipboard.go
```
package tui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func copyToClipboard(content string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		if _, err := exec.LookPath("wl-copy"); err == nil {
			cmd = exec.Command("wl-copy")
		} else if _, err := exec.LookPath("xclip"); err == nil {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		} else if _, err := exec.LookPath("xsel"); err == nil {
			cmd = exec.Command("xsel", "--clipboard", "--input")
		} else {
			return err
		}
	case "windows":
		cmd = exec.Command("cmd", "/c", "clip")
	default:
		return exec.ErrNotFound
	}

	cmd.Stdin = strings.NewReader(content)
	return cmd.Run()
}

func writeClipboardFallback(content string) (string, error) {
	file, err := os.CreateTemp("", "oraclepack-step-*.txt")
	if err != nil {
		return "", fmt.Errorf("create temp file: %w", err)
	}
	defer file.Close()
	if _, err := file.WriteString(content); err != nil {
		return "", fmt.Errorf("write temp file: %w", err)
	}
	return file.Name(), nil
}
```

internal/tui/filter_test.go
```
package tui

import (
	"os"
	"path/filepath"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/types"
)

func TestFilterLogic(t *testing.T) {
	// Setup pack with steps having different ROI
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", ROI: 1.0, OriginalLine: "Step 1"},
			{ID: "02", ROI: 5.0, OriginalLine: "Step 2"},
			{ID: "03", ROI: 10.0, OriginalLine: "Step 3"},
		},
	}
	r := exec.NewRunner(exec.RunnerOptions{})
	s := &state.RunState{}

	// Initialize model with no filter (threshold 0)
	m := NewModel(p, r, s, "", 0, "over", false, false, 0, false, "auto")

	if len(m.list.Items()) != 3 {
		t.Fatalf("expected 3 items initially, got %d", len(m.list.Items()))
	}

	// Apply filter: ROI >= 5.0
	m.roiThreshold = 5.0
	m.roiMode = "over"
	m = m.refreshList()

	if len(m.list.Items()) != 2 {
		t.Errorf("expected 2 items after filtering >= 5.0, got %d", len(m.list.Items()))
	}

	// Verify items are 02 and 03
	items := m.list.Items()
	if items[0].(item).id != "02" {
		t.Errorf("expected first item to be 02, got %s", items[0].(item).id)
	}
	if items[1].(item).id != "03" {
		t.Errorf("expected second item to be 03, got %s", items[1].(item).id)
	}

	// Apply filter: ROI < 5.0 ("under")
	m.roiThreshold = 5.0
	m.roiMode = "under"
	m = m.refreshList()

	if len(m.list.Items()) != 1 {
		t.Errorf("expected 1 item after filtering < 5.0, got %d", len(m.list.Items()))
	}
	if m.list.Items()[0].(item).id != "01" {
		t.Errorf("expected item to be 01, got %s", m.list.Items()[0].(item).id)
	}
}

func TestROIModeTogglePersists(t *testing.T) {
	dir := t.TempDir()
	statePath := filepath.Join(dir, "state.json")
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", ROI: 1.0, OriginalLine: "Step 1"},
		},
	}
	r := exec.NewRunner(exec.RunnerOptions{})
	s := &state.RunState{SchemaVersion: 1}

	m := NewModel(p, r, s, statePath, 0, "over", false, false, 0, false, "auto")

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("m")})
	m2 := updated.(Model)
	if m2.roiMode != "under" {
		t.Fatalf("expected roiMode to toggle to under, got %s", m2.roiMode)
	}

	loaded, err := state.LoadState(statePath)
	if err != nil {
		t.Fatalf("failed to load state: %v", err)
	}
	if loaded.ROIMode != "under" {
		t.Fatalf("expected persisted roiMode under, got %s", loaded.ROIMode)
	}

	if err := os.Remove(statePath); err != nil {
		t.Fatalf("failed to cleanup state file: %v", err)
	}
}
```

internal/tui/overrides_confirm.go
```
package tui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/overrides"
)

type ValidationResultMsg struct {
	Errors []exec.ValidationError
	Err    error
}

type OverridesConfirmModel struct {
	validating bool
	errMsg     string
	errors     []exec.ValidationError
}

func (m OverridesConfirmModel) View(over overrides.RuntimeOverrides, baseline []string) string {
	added := strings.Join(over.AddedFlags, ", ")
	if added == "" {
		added = "(none)"
	}
	removed := strings.Join(over.RemovedFlags, ", ")
	if removed == "" {
		removed = "(none)"
	}
	targeted := len(over.ApplyToSteps)
	targetList := formatTargetList(over.ApplyToSteps, 5)
	effective := effectiveFlagsSummary(over, baseline)
	lines := []string{
		"Summary:",
		fmt.Sprintf("Added flags: %s", added),
		fmt.Sprintf("Removed flags: %s", removed),
		fmt.Sprintf("Targeted steps: %d%s", targeted, targetList),
		fmt.Sprintf("Effective flags: %s", effective),
		"",
		"[Enter] Validate  [Esc] Cancel",
	}

	if m.validating {
		lines = append(lines, "", "Validating overrides...")
	}
	if m.errMsg != "" {
		lines = append(lines, "", "Validation failed:", m.errMsg)
	}
	if len(m.errors) > 0 {
		lines = append(lines, "", fmt.Sprintf("Validation errors (%d):", len(m.errors)))
		lines = append(lines, formatValidationErrors(m.errors, 6)...)
	}

	return strings.Join(lines, "\n")
}

func formatTargetList(targets map[string]bool, limit int) string {
	if len(targets) == 0 || limit <= 0 {
		return ""
	}
	ids := make([]string, 0, len(targets))
	for id := range targets {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	if len(ids) <= limit {
		return fmt.Sprintf(" (%s)", strings.Join(ids, ", "))
	}
	return fmt.Sprintf(" (%s, +%d more)", strings.Join(ids[:limit], ", "), len(ids)-limit)
}

func effectiveFlagsSummary(over overrides.RuntimeOverrides, baseline []string) string {
	if len(over.ApplyToSteps) == 0 {
		return "(no steps targeted)"
	}
	var first string
	for id := range over.ApplyToSteps {
		first = id
		break
	}
	flags := over.EffectiveFlags(first, baseline)
	if len(flags) == 0 {
		return "(none)"
	}
	return strings.Join(flags, " ")
}

func formatValidationErrors(errors []exec.ValidationError, limit int) []string {
	if limit <= 0 {
		return nil
	}
	lines := []string{}
	for i, err := range errors {
		if i >= limit {
			lines = append(lines, fmt.Sprintf("- (+%d more)", len(errors)-limit))
			break
		}
		msg := strings.TrimSpace(err.ErrorMessage)
		if msg == "" {
			msg = "(no error message)"
		}
		lines = append(lines, fmt.Sprintf("- Step %s: %s", err.StepID, firstLine(msg)))
	}
	return lines
}

func firstLine(msg string) string {
	if idx := strings.IndexByte(msg, '\n'); idx != -1 {
		return msg[:idx]
	}
	return msg
}
```

internal/tui/overrides_flags.go
```
package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FlagItem struct {
	Flag       string
	Desc       string
	IsBaseline bool
	Selected   bool
}

func (i FlagItem) Title() string       { return i.Flag }
func (i FlagItem) Description() string { return i.Desc }
func (i FlagItem) FilterValue() string { return i.Flag }

type FlagsPickerModel struct {
	list list.Model
}

func NewFlagsPickerModel(baseline []string) FlagsPickerModel {
	baselineSet := make(map[string]bool, len(baseline))
	for _, f := range baseline {
		baselineSet[f] = true
	}

	curated := []FlagItem{
		{Flag: "--files-report", Desc: "Show per-file token usage"},
		{Flag: "--render", Desc: "Print assembled markdown bundle"},
		{Flag: "--render-plain", Desc: "Render markdown without ANSI"},
		{Flag: "--copy", Desc: "Copy assembled markdown bundle"},
		{Flag: "--wait", Desc: "Wait for background API runs"},
	}

	items := make([]list.Item, 0, len(curated))
	for _, c := range curated {
		c.IsBaseline = baselineSet[c.Flag]
		if c.IsBaseline {
			c.Selected = true
		}
		items = append(items, c)
	}

	delegate := newFlagsDelegate()
	l := list.New(items, delegate, 0, 0)
	l.Title = "Oracle Flags"
	l.SetFilteringEnabled(true)

	return FlagsPickerModel{list: l}
}

func (m FlagsPickerModel) Init() tea.Cmd {
	return nil
}

func (m FlagsPickerModel) Update(msg tea.Msg) (FlagsPickerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == " " {
			idx := m.list.Index()
			item, ok := m.list.SelectedItem().(FlagItem)
			if ok && !item.IsBaseline {
				item.Selected = !item.Selected
				_ = m.list.SetItem(idx, item)
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *FlagsPickerModel) SetSize(width, height int) {
	m.list.SetSize(width, height)
}

func (m FlagsPickerModel) View() string {
	return m.list.View()
}

func (m FlagsPickerModel) SelectedFlags() []string {
	var flags []string
	for _, item := range m.list.Items() {
		if fi, ok := item.(FlagItem); ok && fi.Selected && !fi.IsBaseline {
			flags = append(flags, fi.Flag)
		}
	}
	return flags
}

type flagsDelegate struct {
	list.DefaultDelegate
}

func newFlagsDelegate() flagsDelegate {
	d := list.NewDefaultDelegate()
	return flagsDelegate{DefaultDelegate: d}
}

func (d flagsDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	fi, ok := item.(FlagItem)
	if !ok {
		d.DefaultDelegate.Render(w, m, index, item)
		return
	}

	checked := fi.Selected || fi.IsBaseline
	marker := "[ ]"
	if checked {
		marker = "[x]"
	}
	if fi.IsBaseline {
		marker = "[*]"
	}

	label := fi.Flag
	if fi.Desc != "" {
		label = fmt.Sprintf("%s - %s", fi.Flag, fi.Desc)
	}
	if fi.IsBaseline {
		label = label + " (base)"
	}

	line := fmt.Sprintf("%s %s", marker, label)
	if index == m.Index() {
		line = d.Styles.SelectedTitle.Render(line)
	} else {
		line = d.Styles.NormalTitle.Render(line)
	}
	if fi.IsBaseline {
		line = lipgloss.NewStyle().Faint(true).Render(line)
	}

	fmt.Fprintln(w, line)
}
```

internal/tui/overrides_flow.go
```
package tui

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/types"
)

type OverridesStep int

const (
	OverridesFlags OverridesStep = iota
	OverridesSteps
	OverridesConfirm
)

type OverridesStartedMsg struct{}

type OverridesAppliedMsg struct {
	Overrides overrides.RuntimeOverrides
}

type OverridesCancelledMsg struct{}

type OverridesFlowModel struct {
	step    OverridesStep
	flags   FlagsPickerModel
	steps   StepsPickerModel
	confirm OverridesConfirmModel

	packSteps        []types.Step
	baseline         []string
	runnerOpts       exec.RunnerOptions
	pendingOverrides overrides.RuntimeOverrides
}

func NewOverridesFlowModel(steps []types.Step, baseline []string, opts exec.RunnerOptions) OverridesFlowModel {
	return OverridesFlowModel{
		step:       OverridesFlags,
		flags:      NewFlagsPickerModel(nil),
		steps:      NewStepsPickerModel(steps),
		confirm:    OverridesConfirmModel{},
		packSteps:  steps,
		baseline:   exec.ApplyChatGPTURL(baseline, opts.ChatGPTURL),
		runnerOpts: opts,
	}
}

func (m OverridesFlowModel) Init() tea.Cmd {
	return nil
}

func (m OverridesFlowModel) Update(msg tea.Msg) (OverridesFlowModel, tea.Cmd) {
	var cmd tea.Cmd
	if m.step == OverridesFlags {
		m.flags, cmd = m.flags.Update(msg)
	}
	if m.step == OverridesSteps {
		m.steps, cmd = m.steps.Update(msg)
	}
	if m.step == OverridesConfirm {
		switch v := msg.(type) {
		case ValidationResultMsg:
			m.confirm.validating = false
			m.confirm.errors = v.Errors
			if v.Err != nil {
				m.confirm.errMsg = v.Err.Error()
				return m, nil
			}
			if len(v.Errors) > 0 {
				m.confirm.errMsg = fmt.Sprintf("%d validation errors detected.", len(v.Errors))
				return m, nil
			}
			return m, func() tea.Msg { return OverridesAppliedMsg{Overrides: m.pendingOverrides} }
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg { return OverridesCancelledMsg{} }
		case "shift+tab", "backspace":
			if m.step > OverridesFlags {
				m.step--
			}
		case "enter", "tab":
			if m.step == OverridesConfirm {
				if m.confirm.validating {
					return m, nil
				}
				m.pendingOverrides = m.currentOverrides()
				m.confirm.validating = true
				m.confirm.errMsg = ""
				m.confirm.errors = nil
				return m, m.validateCmd(m.pendingOverrides)
			}
			m.step++
		}
	}

	return m, cmd
}

func (m OverridesFlowModel) View(width, height int) string {
	title := lipgloss.NewStyle().Bold(true).Render("Overrides Wizard")
	step := fmt.Sprintf("Step %d/3", int(m.step)+1)
	body := fmt.Sprintf("Current step: %s\n\n[Enter] Next  [Esc] Cancel", overridesStepName(m.step))

	var content string
	if m.step == OverridesFlags {
		m.flags.SetSize(width-4, height-8)
		content = lipgloss.JoinVertical(lipgloss.Left,
			title,
			step,
			"",
			m.flags.View(),
			"",
			body,
		)
	} else if m.step == OverridesSteps {
		m.steps.SetSize(width-4, height-8)
		content = lipgloss.JoinVertical(lipgloss.Left,
			title,
			step,
			"",
			m.steps.View(),
			"",
			body,
		)
	} else if m.step == OverridesConfirm {
		content = lipgloss.JoinVertical(lipgloss.Left,
			title,
			step,
			"",
			m.confirm.View(m.currentOverrides(), m.baseline),
		)
	} else {
		content = lipgloss.JoinVertical(lipgloss.Left,
			title,
			step,
			"",
			body,
		)
	}

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
}

func (m OverridesFlowModel) currentOverrides() overrides.RuntimeOverrides {
	return overrides.RuntimeOverrides{
		AddedFlags:   m.flags.SelectedFlags(),
		RemovedFlags: nil,
		ApplyToSteps: m.steps.SelectedSteps(),
	}
}

func (m OverridesFlowModel) validateCmd(over overrides.RuntimeOverrides) tea.Cmd {
	return func() tea.Msg {
		errs, err := exec.ValidateOverrides(context.Background(), m.packSteps, &over, m.baseline, m.runnerOpts)
		return ValidationResultMsg{Errors: errs, Err: err}
	}
}

func overridesStepName(step OverridesStep) string {
	switch step {
	case OverridesFlags:
		return "Flags"
	case OverridesSteps:
		return "Target Steps"
	case OverridesConfirm:
		return "Confirm"
	default:
		return "Unknown"
	}
}
```

internal/tui/overrides_steps.go
```
package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/oraclepack/internal/types"
)

type StepItem struct {
	ID       string
	TitleTxt string
	DescTxt  string
	Selected bool
}

func (i StepItem) Title() string       { return i.TitleTxt }
func (i StepItem) Description() string { return i.DescTxt }
func (i StepItem) FilterValue() string { return i.TitleTxt }

type StepsPickerModel struct {
	list list.Model
}

func NewStepsPickerModel(steps []types.Step) StepsPickerModel {
	items := make([]list.Item, 0, len(steps))
	for _, s := range steps {
		items = append(items, StepItem{
			ID:       s.ID,
			TitleTxt: fmt.Sprintf("Step %s", s.ID),
			DescTxt:  s.OriginalLine,
			Selected: true,
		})
	}

	delegate := newStepsDelegate()
	l := list.New(items, delegate, 0, 0)
	l.Title = "Target Steps"
	l.SetFilteringEnabled(true)

	return StepsPickerModel{list: l}
}

func (m StepsPickerModel) Init() tea.Cmd {
	return nil
}

func (m StepsPickerModel) Update(msg tea.Msg) (StepsPickerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "a", "A":
			m = m.setAll(true)
			return m, nil
		case "i":
			m = m.invert()
			return m, nil
		case "n":
			m = m.setAll(false)
			return m, nil
		case " ":
			idx := m.list.Index()
			item, ok := m.list.SelectedItem().(StepItem)
			if ok {
				item.Selected = !item.Selected
				_ = m.list.SetItem(idx, item)
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *StepsPickerModel) SetSize(width, height int) {
	m.list.SetSize(width, height)
}

func (m StepsPickerModel) View() string {
	help := lipgloss.NewStyle().Faint(true).Render("[space] toggle  [a] all  [i] invert  [n] none")
	return m.list.View() + "\n" + help
}

func (m StepsPickerModel) SelectedSteps() map[string]bool {
	selected := make(map[string]bool)
	for _, item := range m.list.Items() {
		if si, ok := item.(StepItem); ok && si.Selected {
			selected[si.ID] = true
		}
	}
	return selected
}

func (m StepsPickerModel) setAll(value bool) StepsPickerModel {
	for idx, item := range m.list.Items() {
		si, ok := item.(StepItem)
		if !ok {
			continue
		}
		si.Selected = value
		_ = m.list.SetItem(idx, si)
	}
	return m
}

func (m StepsPickerModel) invert() StepsPickerModel {
	for idx, item := range m.list.Items() {
		si, ok := item.(StepItem)
		if !ok {
			continue
		}
		si.Selected = !si.Selected
		_ = m.list.SetItem(idx, si)
	}
	return m
}

type stepsDelegate struct {
	list.DefaultDelegate
}

func newStepsDelegate() stepsDelegate {
	d := list.NewDefaultDelegate()
	return stepsDelegate{DefaultDelegate: d}
}

func (d stepsDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	si, ok := item.(StepItem)
	if !ok {
		d.DefaultDelegate.Render(w, m, index, item)
		return
	}

	marker := "[ ]"
	if si.Selected {
		marker = "[x]"
	}

	label := si.TitleTxt
	if si.DescTxt != "" {
		label = fmt.Sprintf("%s - %s", si.TitleTxt, si.DescTxt)
	}

	line := fmt.Sprintf("%s %s", marker, label)
	if index == m.Index() {
		line = d.Styles.SelectedTitle.Render(line)
	} else {
		line = d.Styles.NormalTitle.Render(line)
	}

	fmt.Fprintln(w, line)
}
```

internal/tui/overrides_url.go
```
package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type URLInputModel struct {
	input textinput.Model
	err   string
}

func NewURLInputModel() URLInputModel {
	ti := textinput.New()
	ti.Placeholder = "https://chat.openai.com/project/..."
	ti.CharLimit = 200
	ti.Width = 50

	return URLInputModel{input: ti}
}

func (m URLInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m URLInputModel) Update(msg tea.Msg) (URLInputModel, tea.Cmd) {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	m.err = ""
	if !m.IsValid() {
		m.err = "Invalid URL (must start with http:// or https://)"
	}
	return m, cmd
}

func (m URLInputModel) Value() string {
	return strings.TrimSpace(m.input.Value())
}

func (m URLInputModel) IsValid() bool {
	v := m.Value()
	if v == "" {
		return true
	}
	return strings.HasPrefix(v, "http://") || strings.HasPrefix(v, "https://")
}

func (m URLInputModel) View() string {
	body := m.input.View()
	if m.err != "" {
		body = body + "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render(m.err)
	}
	return body
}

func (m *URLInputModel) SetValue(v string) {
	m.input.SetValue(v)
}

func (m *URLInputModel) Focus() {
	m.input.Focus()
}

func (m *URLInputModel) Blur() {
	m.input.Blur()
}
```

internal/tui/preview_test.go
```
package tui

import (
	"strings"
	"testing"

	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/types"
)

func TestStepPreviewContentUnwrapped(t *testing.T) {
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", OriginalLine: "Step 1", Code: "echo hello"},
		},
	}
	r := exec.NewRunner(exec.RunnerOptions{})
	s := &state.RunState{}
	m := NewModel(p, r, s, "", 0, "over", false, false, 0, false, "auto")
	m.width = 80
	m.previewID = "01"
	m.previewWrap = false

	content := m.stepPreviewContent()
	if !strings.Contains(content, "Step 01") {
		t.Fatalf("expected header to include step id, got %q", content)
	}
	if !strings.Contains(content, "echo hello") {
		t.Fatalf("expected content to include code, got %q", content)
	}
}
```

internal/tui/tui.go
```
package tui

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/render"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/types"
)

type ViewState int

const (
	ViewSteps ViewState = iota
	ViewRunning
	ViewDone
	ViewOverrides
	ViewStepPreview
)

type item struct {
	id     string
	title  string
	desc   string
	status state.Status
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	list        list.Model
	viewport    viewport.Model
	spinner     spinner.Model
	filterInput textinput.Model
	urlInput    URLInputModel
	urlPicker   URLPickerModel
	pack        *types.Pack
	runner      *exec.Runner
	state       *state.RunState
	statePath   string

	width  int
	height int

	viewState     ViewState
	running       bool
	runAll        bool // State for sequential execution
	currentIdx    int
	autoRun       bool // Config to auto-start on init
	previewID     string
	previewWrap   bool
	previewNotice string

	// Filtering state
	allSteps     []item // Store all items to support dynamic filtering
	roiThreshold float64
	roiMode      string
	isFiltering  bool
	isEditingURL bool
	isPickingURL bool

	overridesFlow         OverridesFlowModel
	appliedOverrides      *overrides.RuntimeOverrides
	chatGPTURL            string
	outputVerify          bool
	outputRetries         int
	outputRequireHeadings bool
	outputChunkMode       string

	err      error
	logLines []string
	logChan  chan string
}

func NewModel(p *types.Pack, r *exec.Runner, s *state.RunState, statePath string, roiThreshold float64, roiMode string, autoRun bool, outputVerify bool, outputRetries int, outputRequireHeadings bool, outputChunkMode string) Model {
	if s != nil {
		if s.ROIThreshold > 0 {
			roiThreshold = s.ROIThreshold
		}
		if s.ROIMode != "" {
			roiMode = s.ROIMode
		}
	}
	var allItems []item
	for _, step := range p.Steps {
		allItems = append(allItems, item{
			id:    step.ID,
			title: fmt.Sprintf("Step %s", step.ID),
			desc:  step.OriginalLine,
		})
	}

	ti := textinput.New()
	ti.Placeholder = "Enter ROI threshold (e.g. 2.5)"
	ti.CharLimit = 10
	ti.Width = 20

	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Oracle Pack Steps"

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	vp := viewport.New(0, 0)
	vp.SetContent("Press Enter to run selected, 'a' to run all filtered steps, 'f' to set ROI threshold, 'm' to toggle ROI mode, 'v' to view step, 'o' to configure overrides, 'u' for ChatGPT URL, 'U' to pick a saved URL.")

	projectPath := ProjectURLStorePath(statePath, p.Source)
	globalPath := GlobalURLStorePath()
	urlPicker := NewURLPickerModel(projectPath, globalPath)
	resolvedURL := r.ChatGPTURL
	if resolvedURL == "" {
		resolvedURL = urlPicker.DefaultURL()
	}
	if resolvedURL != "" {
		r.ChatGPTURL = resolvedURL
	}

	m := Model{
		list:                  l,
		viewport:              vp,
		spinner:               sp,
		filterInput:           ti,
		urlInput:              NewURLInputModel(),
		urlPicker:             urlPicker,
		pack:                  p,
		runner:                r,
		state:                 s,
		statePath:             statePath,
		autoRun:               autoRun,
		allSteps:              allItems,
		roiThreshold:          roiThreshold,
		roiMode:               roiMode,
		logChan:               make(chan string, 100),
		viewState:             ViewSteps,
		overridesFlow:         NewOverridesFlowModel(p.Steps, r.OracleFlags, RunnerOptionsFromRunner(r)),
		chatGPTURL:            resolvedURL,
		previewWrap:           true,
		outputVerify:          outputVerify,
		outputRetries:         outputRetries,
		outputRequireHeadings: outputRequireHeadings,
		outputChunkMode:       outputChunkMode,
	}
	m.urlInput.SetValue(resolvedURL)
	m.urlInput.Blur()

	// Apply initial filter
	return m.refreshList()
}

func (m Model) refreshList() Model {
	var filtered []list.Item
	for _, it := range m.allSteps {
		// Find the original step to check ROI
		var step *types.Step
		for _, s := range m.pack.Steps {
			if s.ID == it.id {
				step = &s
				break
			}
		}
		if step == nil {
			continue
		}

		if m.roiThreshold > 0 {
			if m.roiMode == "under" {
				if step.ROI >= m.roiThreshold {
					continue
				}
			} else {
				if step.ROI < m.roiThreshold {
					continue
				}
			}
		}
		filtered = append(filtered, it)
	}
	m.list.SetItems(filtered)
	return m
}

type StartAutoRunMsg struct{}

func (m Model) Init() tea.Cmd {
[TRUNCATED]
```

internal/tui/tui_test.go
```
package tui

import (
	"testing"

	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/types"
)

func TestInitAutoRun(t *testing.T) {
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", Number: 1, Code: "echo hello"},
		},
	}
	r := exec.NewRunner(exec.RunnerOptions{})
	s := &state.RunState{}

	// Test case 1: autoRun = true
	modelAuto := NewModel(p, r, s, "", 0, "over", true, false, 0, false, "auto")
	cmdAuto := modelAuto.Init()

	if cmdAuto == nil {
		t.Fatal("expected Init cmd to be non-nil when autoRun is true")
	}
	// Note: We can't easily assert the content of a Batch command in a unit test.

	// Test case 2: autoRun = false
	modelManual := NewModel(p, r, s, "", 0, "over", false, false, 0, false, "auto")
	// Even with autoRun false, we have textinput.Blink, so Init is not nil.
	cmdManual := modelManual.Init()
	if cmdManual == nil {
		t.Fatal("expected Init cmd to be non-nil due to textinput.Blink")
	}
}
```

internal/tui/url_picker.go
```
package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type URLPickedMsg struct {
	URL string
}

type URLPickerCancelledMsg struct{}

type urlItem struct {
	name  string
	url   string
	scope string
}

func (i urlItem) Title() string       { return i.name }
func (i urlItem) Description() string { return fmt.Sprintf("%s • %s", i.scope, i.url) }
func (i urlItem) FilterValue() string { return i.name }

type URLPickerModel struct {
	list list.Model

	projectPath string
	globalPath  string
	project     URLStore
	global      URLStore

	editing   bool
	editName  textinput.Model
	editURL   textinput.Model
	editScope string
	editIdx   int
	editIsNew bool

	errMsg string
}

func NewURLPickerModel(projectPath, globalPath string) URLPickerModel {
	project, _ := LoadURLStore(projectPath)
	global, _ := LoadURLStore(globalPath)

	items := makeURLItems(project, global)
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "ChatGPT Project URLs"
	l.SetFilteringEnabled(true)
	selectDefault(&l, project, global)

	name := textinput.New()
	name.Placeholder = "Name (e.g., Core Project)"
	name.CharLimit = 60
	name.Width = 40

	url := textinput.New()
	url.Placeholder = "https://chatgpt.com/g/.../project"
	url.CharLimit = 200
	url.Width = 60

	return URLPickerModel{
		list:        l,
		projectPath: projectPath,
		globalPath:  globalPath,
		project:     project,
		global:      global,
		editName:    name,
		editURL:     url,
	}
}

func (m URLPickerModel) Init() tea.Cmd {
	return nil
}

func (m URLPickerModel) Update(msg tea.Msg) (URLPickerModel, tea.Cmd) {
	if m.editing {
		return m.updateEdit(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg { return URLPickerCancelledMsg{} }
		case "enter":
			item, ok := m.list.SelectedItem().(urlItem)
			if !ok {
				return m, nil
			}
			m.touch(item)
			return m, func() tea.Msg { return URLPickedMsg{URL: item.url} }
		case "a":
			m.startEdit(urlScopeProject, "", "", true)
			return m, nil
		case "e":
			item, ok := m.list.SelectedItem().(urlItem)
			if !ok {
				return m, nil
			}
			m.startEdit(item.scope, item.name, item.url, false)
			return m, nil
		case "d":
			item, ok := m.list.SelectedItem().(urlItem)
			if !ok {
				return m, nil
			}
			m.delete(item)
			return m, nil
		case "s":
			item, ok := m.list.SelectedItem().(urlItem)
			if !ok {
				return m, nil
			}
			m.setDefault(item)
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *URLPickerModel) SetSize(width, height int) {
	m.list.SetSize(width, height-4)
}

func (m URLPickerModel) View() string {
	if m.editing {
		return m.editView()
	}

	help := lipgloss.NewStyle().Faint(true).Render("[enter] use  [a] add  [e] edit  [d] delete  [s] default  [esc] cancel")
	return m.list.View() + "\n" + help
}

func makeURLItems(project URLStore, global URLStore) []list.Item {
	var items []list.Item
	for _, it := range project.Items {
		items = append(items, urlItem{name: it.Name, url: it.URL, scope: urlScopeProject})
	}
	for _, it := range global.Items {
		items = append(items, urlItem{name: it.Name, url: it.URL, scope: urlScopeGlobal})
	}
	return items
}

func selectDefault(l *list.Model, project URLStore, global URLStore) {
	if l == nil {
		return
	}
	name, scope := defaultNameScope(project, global)
	if name == "" {
		return
	}
	for idx, item := range l.Items() {
		if it, ok := item.(urlItem); ok && it.name == name && it.scope == scope {
			l.Select(idx)
			return
		}
	}
}

func defaultNameScope(project URLStore, global URLStore) (string, string) {
	if project.Default != "" {
		return project.Default, urlScopeProject
	}
	if global.Default != "" {
		return global.Default, urlScopeGlobal
	}
	return "", ""
}

func (m URLPickerModel) DefaultURL() string {
	name, scope := defaultNameScope(m.project, m.global)
	if name == "" {
		return ""
	}
	store := m.storeFor(scope)
	if store == nil {
		return ""
	}
	for _, it := range store.Items {
		if it.Name == name {
			return it.URL
		}
	}
	return ""
}

func (m *URLPickerModel) refresh() {
	m.list.SetItems(makeURLItems(m.project, m.global))
	selectDefault(&m.list, m.project, m.global)
}

[TRUNCATED]
```

internal/tui/url_store.go
```
package tui

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	urlScopeProject = "project"
	urlScopeGlobal  = "global"
)

type URLItem struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	LastUsed string `json:"lastUsed,omitempty"`
}

type URLStore struct {
	Default string    `json:"default"`
	Items   []URLItem `json:"items"`
}

func LoadURLStore(path string) (URLStore, error) {
	if path == "" {
		return URLStore{}, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return URLStore{}, nil
		}
		return URLStore{}, err
	}
	var store URLStore
	if err := json.Unmarshal(data, &store); err != nil {
		return URLStore{}, err
	}
	return store, nil
}

func SaveURLStore(path string, store URLStore) error {
	if path == "" {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func ProjectURLStorePath(statePath, packSource string) string {
	if statePath != "" {
		base := strings.TrimSuffix(statePath, ".state.json")
		return base + ".chatgpt-urls.json"
	}
	if packSource == "" {
		return ""
	}
	return packSource + ".chatgpt-urls.json"
}

func GlobalURLStorePath() string {
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return ""
	}
	return filepath.Join(home, ".oraclepack", "chatgpt-urls.json")
}

func nowRFC3339() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func isValidURL(value string) bool {
	v := strings.TrimSpace(value)
	if v == "" {
		return false
	}
	return strings.HasPrefix(v, "http://") || strings.HasPrefix(v, "https://")
}
```

internal/tui/url_store_test.go
```
package tui

import (
	"path/filepath"
	"testing"
)

func TestURLStoreSaveLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "urls.json")

	store := URLStore{
		Default: "Primary",
		Items: []URLItem{{
			Name: "Primary",
			URL:  "https://chatgpt.com/g/primary",
		}},
	}

	if err := SaveURLStore(path, store); err != nil {
		t.Fatalf("failed to save store: %v", err)
	}

	loaded, err := LoadURLStore(path)
	if err != nil {
		t.Fatalf("failed to load store: %v", err)
	}

	if loaded.Default != store.Default {
		t.Fatalf("expected default %q, got %q", store.Default, loaded.Default)
	}
	if len(loaded.Items) != 1 || loaded.Items[0].URL != store.Items[0].URL {
		t.Fatalf("loaded items mismatch: %+v", loaded.Items)
	}
}

func TestURLPickerDefaultURLPrefersProject(t *testing.T) {
	dir := t.TempDir()
	projectPath := filepath.Join(dir, "project.json")
	globalPath := filepath.Join(dir, "global.json")

	project := URLStore{
		Default: "Project",
		Items: []URLItem{{
			Name: "Project",
			URL:  "https://chatgpt.com/g/project",
		}},
	}
	global := URLStore{
		Default: "Global",
		Items: []URLItem{{
			Name: "Global",
			URL:  "https://chatgpt.com/g/global",
		}},
	}

	if err := SaveURLStore(projectPath, project); err != nil {
		t.Fatalf("failed to save project store: %v", err)
	}
	if err := SaveURLStore(globalPath, global); err != nil {
		t.Fatalf("failed to save global store: %v", err)
	}

	picker := NewURLPickerModel(projectPath, globalPath)
	if got := picker.DefaultURL(); got != project.Items[0].URL {
		t.Fatalf("expected project default URL %q, got %q", project.Items[0].URL, got)
	}
}
```

internal/types/pack.go
```
package types

// Pack represents a parsed oracle pack.
type Pack struct {
	Prelude     Prelude `json:"prelude" yaml:"prelude"`
	Steps       []Step  `json:"steps" yaml:"steps"`
	Source      string  `json:"source,omitempty" yaml:"source,omitempty"`
	OutDir      string  `json:"out_dir,omitempty" yaml:"out_dir,omitempty"`
	WriteOutput bool    `json:"write_output" yaml:"write_output"`
}

// Prelude contains the shell code that runs before any steps.
type Prelude struct {
	Code string `json:"code" yaml:"code"`
}

// Step represents an individual executable step within the pack.
type Step struct {
	ID           string  `json:"id" yaml:"id"`                             // e.g., "01"
	Number       int     `json:"number" yaml:"number"`                     // e.g., 1
	Code         string  `json:"code" yaml:"code"`                         // The bash code
	OriginalLine string  `json:"original_line" yaml:"original_line"`       // The header line, e.g., "# 01)"
	ROI          float64 `json:"roi,omitempty" yaml:"roi,omitempty"`       // Return on Investment value extracted from header
	Impact       string  `json:"impact,omitempty" yaml:"impact,omitempty"` // Optional impact metadata extracted from step comments
}
```

internal/types/pack_test.go
```
package types

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/goccy/go-yaml"
)

func TestPackJSONRoundTrip(t *testing.T) {
	original := Pack{
		Prelude: Prelude{Code: "echo prelude"},
		Steps: []Step{
			{
				ID:           "01",
				Number:       1,
				Code:         "echo hello",
				OriginalLine: "# 01) Example",
				ROI:          3.2,
				Impact:       "High",
			},
		},
		Source:      "pack.md",
		OutDir:      "dist",
		WriteOutput: true,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("json marshal: %v", err)
	}

	var decoded Pack
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("json unmarshal: %v", err)
	}

	if !reflect.DeepEqual(original, decoded) {
		t.Fatalf("json round-trip mismatch: %#v != %#v", original, decoded)
	}
}

func TestPackYAMLRoundTrip(t *testing.T) {
	original := Pack{
		Prelude: Prelude{Code: "echo prelude"},
		Steps: []Step{
			{
				ID:           "02",
				Number:       2,
				Code:         "echo yaml",
				OriginalLine: "# 02) Example",
				ROI:          1.1,
				Impact:       "Low",
			},
		},
		Source:      "pack.yaml",
		OutDir:      "out",
		WriteOutput: false,
	}

	data, err := yaml.Marshal(original)
	if err != nil {
		t.Fatalf("yaml marshal: %v", err)
	}

	var decoded Pack
	if err := yaml.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("yaml unmarshal: %v", err)
	}

	if !reflect.DeepEqual(original, decoded) {
		t.Fatalf("yaml round-trip mismatch: %#v != %#v", original, decoded)
	}
}
```

internal/types/verification.go
```
package types

// OutputContract describes the expected response shape for a step.
type OutputContract string

const (
	OutputContractUnknown          OutputContract = ""
	OutputContractAllSections      OutputContract = "all_sections"
	OutputContractDirectAnswerOnly OutputContract = "direct_answer_only"
	OutputContractChunkedBySuffix  OutputContract = "chunked_by_suffix"
)

// OutputFailure captures a missing or invalid output artifact.
type OutputFailure struct {
	StepID        string   `json:"step_id,omitempty" yaml:"step_id,omitempty"`
	Path          string   `json:"path,omitempty" yaml:"path,omitempty"`
	MissingTokens []string `json:"missing_tokens,omitempty" yaml:"missing_tokens,omitempty"`
	Error         string   `json:"error,omitempty" yaml:"error,omitempty"`
}

// SyntaxFinding captures a structural or syntax issue in generated bash.
type SyntaxFinding struct {
	StepID  string `json:"step_id,omitempty" yaml:"step_id,omitempty"`
	Line    int    `json:"line" yaml:"line"`
	Token   string `json:"token,omitempty" yaml:"token,omitempty"`
	Message string `json:"message" yaml:"message"`
}
```

internal/validate/artifact_gate.go
```
package validate

import (
	"errors"

	"github.com/user/oraclepack/internal/artifacts"
	"github.com/user/oraclepack/internal/foundation"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
	"github.com/user/oraclepack/internal/types"
)

// ArtifactGateValidator checks expected artifacts after a step.
type ArtifactGateValidator struct {
	Contract artifacts.Contract
}

func (v ArtifactGateValidator) Validate(step *types.Step, kind tools.ToolKind, toolPresent bool) (state.Status, string) {
	if step == nil {
		return state.StatusSuccess, ""
	}
	contract := v.Contract
	if contract == nil {
		contract = artifacts.DefaultContract()
	}
	if !toolPresent {
		if _, ok := contract[step.ID]; ok {
			return state.StatusSkipped, "tool missing; skipping artifact gate"
		}
		return state.StatusSuccess, ""
	}
	err := artifacts.EvaluateGates(step.ID, contract)
	if err == nil {
		return state.StatusSuccess, ""
	}
	if errors.Is(err, foundation.ErrArtifactMissing) {
		return state.StatusFailed, err.Error()
	}
	return state.StatusFailed, err.Error()
}
```

internal/validate/artifact_gate_test.go
```
package validate

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/user/oraclepack/internal/artifacts"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
	"github.com/user/oraclepack/internal/types"
)

func TestArtifactGateValidator(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "next.json")
	contract := artifacts.Contract{"09": {path}}
	step := &types.Step{ID: "09"}
	v := ArtifactGateValidator{Contract: contract}

	status, _ := v.Validate(step, tools.ToolCodex, true)
	if status != state.StatusFailed {
		t.Fatalf("expected failed for missing artifact, got %s", status)
	}

	if err := os.WriteFile(path, []byte("ok"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	status, _ = v.Validate(step, tools.ToolCodex, true)
	if status != state.StatusSuccess {
		t.Fatalf("expected success, got %s", status)
	}

	status, _ = v.Validate(step, tools.ToolCodex, false)
	if status != state.StatusSkipped {
		t.Fatalf("expected skipped when tool missing, got %s", status)
	}
}
```

internal/validate/composite.go
```
package validate

import (
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/types"
)

// CompositeValidator coordinates multiple validators for all steps.
type CompositeValidator struct {
	ToolPresence ToolPresenceValidator
	OracleDryRun OracleDryRunValidator
	ArtifactGate ArtifactGateValidator
}

// ValidatePack validates all steps in a pack and returns per-step results.
func (v CompositeValidator) ValidatePack(p *types.Pack) []StepResult {
	if p == nil {
		return nil
	}
	results := make([]StepResult, 0, len(p.Steps))
	for i := range p.Steps {
		step := &p.Steps[i]
		kind := DetectToolKind(step)

		status, reason, toolPresent := v.ToolPresence.Validate(step, kind)
		if status == state.StatusSkipped {
			results = append(results, StepResult{
				StepID:   step.ID,
				ToolKind: kind,
				Status:   status,
				Error:    reason,
			})
			continue
		}

		if oracleStatus, oracleErr := v.OracleDryRun.Validate(step, kind); oracleStatus == state.StatusFailed {
			results = append(results, StepResult{
				StepID:   step.ID,
				ToolKind: kind,
				Status:   oracleStatus,
				Error:    oracleErr,
			})
			continue
		}

		gateStatus, gateErr := v.ArtifactGate.Validate(step, kind, toolPresent)
		finalStatus := gateStatus
		errMsg := gateErr
		if finalStatus == "" {
			finalStatus = state.StatusSuccess
		}
		results = append(results, StepResult{
			StepID:   step.ID,
			ToolKind: kind,
			Status:   finalStatus,
			Error:    errMsg,
		})
	}
	return results
}
```

internal/validate/composite_test.go
```
package validate

import (
	"testing"

	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/types"
)

func TestCompositeValidator(t *testing.T) {
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", Code: "codex exec \"hi\""},
		},
	}
	cv := CompositeValidator{
		ToolPresence: ToolPresenceValidator{Checker: fakeChecker{found: map[string]bool{"codex": false}}},
	}
	results := cv.ValidatePack(p)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Status != state.StatusSkipped {
		t.Fatalf("expected skipped, got %s", results[0].Status)
	}
}
```

internal/validate/oracle.go
```
package validate

import (
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
	"github.com/user/oraclepack/internal/types"
)

// OracleDryRunValidator is a placeholder for oracle-specific validation.
type OracleDryRunValidator struct{}

// Validate returns success for non-oracle tools and no-op for oracle until wired.
func (OracleDryRunValidator) Validate(step *types.Step, kind tools.ToolKind) (state.Status, string) {
	if kind != tools.ToolOracle {
		return state.StatusSuccess, ""
	}
	return state.StatusSuccess, ""
}
```

internal/validate/presence.go
```
package validate

import (
	"fmt"

	"github.com/user/oraclepack/internal/shell"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
	"github.com/user/oraclepack/internal/types"
)

// ToolPresenceValidator checks that tool binaries exist on PATH.
type ToolPresenceValidator struct {
	Checker tools.PresenceChecker
}

// Validate returns skipped if the tool is missing.
func (v ToolPresenceValidator) Validate(step *types.Step, kind tools.ToolKind) (state.Status, string, bool) {
	if kind == tools.ToolUnknown {
		return state.StatusSuccess, "", true
	}
	meta, ok := tools.Metadata(kind)
	if !ok {
		return state.StatusSuccess, "", true
	}
	checker := v.Checker
	if checker == nil {
		checker = shellChecker{}
	}
	if _, found := checker.DetectBinary(meta.Name); !found {
		return state.StatusSkipped, fmt.Sprintf("tool %s not found on PATH", meta.Name), false
	}
	return state.StatusSuccess, "", true
}

type shellChecker struct{}

func (shellChecker) DetectBinary(name string) (string, bool) {
	return shell.DetectBinary(name)
}
```

internal/validate/presence_test.go
```
package validate

import (
	"testing"

	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
	"github.com/user/oraclepack/internal/types"
)

type fakeChecker struct {
	found map[string]bool
}

func (f fakeChecker) DetectBinary(name string) (string, bool) {
	return name, f.found[name]
}

func TestToolPresenceValidator(t *testing.T) {
	step := &types.Step{ID: "01", Code: "codex exec \"hi\""}
	v := ToolPresenceValidator{Checker: fakeChecker{found: map[string]bool{"codex": false}}}
	status, reason, present := v.Validate(step, tools.ToolCodex)
	if status != state.StatusSkipped || present {
		t.Fatalf("expected skipped/missing, got status=%s present=%v", status, present)
	}
	if reason == "" {
		t.Fatalf("expected reason for missing tool")
	}
}
```

internal/validate/report.go
```
package validate

// ValidationReport captures results for all steps.
type ValidationReport struct {
	Steps []StepResult
}
```

internal/validate/types.go
```
package validate

import (
	"strings"

	"github.com/user/oraclepack/internal/dispatch"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
	"github.com/user/oraclepack/internal/types"
)

// StepResult captures validation output for a step.
type StepResult struct {
	StepID   string
	ToolKind tools.ToolKind
	Status   state.Status
	Error    string
}

// DetectToolKind scans a step for a known tool prefix.
func DetectToolKind(step *types.Step) tools.ToolKind {
	if step == nil {
		return tools.ToolUnknown
	}
	lines := strings.Split(step.Code, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		if cls, ok := dispatch.Classify(trimmed); ok {
			return cls.Kind
		}
		break
	}
	return tools.ToolUnknown
}
```

</source_code>