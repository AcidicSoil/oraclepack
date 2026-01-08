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