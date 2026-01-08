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
