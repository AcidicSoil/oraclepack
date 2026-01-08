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
