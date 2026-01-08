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
