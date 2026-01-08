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
