<filetree>
Project Structure:
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
    return f"Detected Stage-2 directory: {out_dir} (Mode: {mode})"

@mcp.tool(annotations=ToolAnnotations(readOnlyHint=True))
async def oraclepack_taskify_validate_stage2(out_dir: str) -> str:
    """Validates Stage-2 directory structure (prefixes 01..20)."""
    resolved_dir = security.validate_path(out_dir)
    result = taskify.validate_stage2_dir(resolved_dir)
    if result.ok:
        return f"Stage-2 directory is valid. Found {len(result.valid_files)} files."
    
    output = ["Stage-2 validation failed:"]
    if result.missing:
        output.append(f"Missing prefixes: {', '.join(result.missing)}")
    if result.ambiguous:
        output.append("Ambiguous prefixes (multiple matches):")
        for pfx, matches in result.ambiguous.items():
            output.append(f"  {pfx}: {', '.join(matches)}")
            
    return "\n".join(output)

@mcp.tool(annotations=ToolAnnotations(readOnlyHint=True))
async def oraclepack_taskify_validate_action_pack(file_path: str) -> str:
    """Validates Stage-3 action pack structure."""
    resolved_path = security.validate_path(file_path)
    result = taskify.validate_action_pack(resolved_path)
    if result.ok:
        return f"Action pack is valid. Detected steps: {', '.join(result.steps)}"
    return f"Action pack validation failed: {result.error}"

@mcp.tool(annotations=ToolAnnotations(readOnlyHint=True))
async def oraclepack_taskify_artifacts_summary(out_dir: str) -> str:
    """Summarizes Stage-3 artifacts."""
    resolved_dir = security.validate_path(out_dir)
    summary = taskify.artifacts_summary(resolved_dir)
    
    output = [f"Artifacts Summary for {summary['out_dir']}:"]
    for name, info in summary["artifacts"].items():
        if info:
            output.append(f"- {name}: FOUND ({info['size']} bytes) at {info['path']}")
        else:
            output.append(f"- {name}: NOT FOUND")
            
    return "\n".join(output)

@mcp.tool(annotations=ToolAnnotations(destructiveHint=True, openWorldHint=True))
async def oraclepack_taskify_run_action_pack(file_path: str) -> str:
    """
    Runs a Stage-3 action pack. REQUIRES ORACLEPACK_ENABLE_EXEC=1.
    """
    # Simply wraps oraclepack_run_pack with action pack defaults
    return await oraclepack_run_pack(file_path, yes=True, run_all=True)

@mcp.tool(annotations=ToolAnnotations(readOnlyHint=True))
async def oraclepack_taskify_generate_prompt(file_path: str) -> str:
    """Generates instructions for an agent to run an action pack."""
    resolved_path = security.validate_path(file_path)
    result = taskify.validate_action_pack(resolved_path)
    if not result.ok:
        return f"Validation failed: {result.error}"
    return taskify.generate_agent_prompt(file_path, result.steps)
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
2. **Execution**: Use `oraclepack_taskify_run_action_pack` to execute the steps non-interactively.
3. **Artifacts**: After execution, use `oraclepack_taskify_artifacts_summary` to see produced files.
4. **Inspection**: Read `_actions.json` or `_actions.md` to understand the outcomes of each step.

## Security Note
Execution tools require `ORACLEPACK_ENABLE_EXEC=1` in the server environment.
"""
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

</source_code>