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
