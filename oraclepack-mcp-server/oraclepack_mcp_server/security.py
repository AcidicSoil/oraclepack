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