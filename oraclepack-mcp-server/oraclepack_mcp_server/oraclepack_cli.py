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
