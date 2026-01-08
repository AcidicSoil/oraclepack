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
