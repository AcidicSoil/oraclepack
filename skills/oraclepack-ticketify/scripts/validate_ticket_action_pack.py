#!/usr/bin/env python3
"""
Validate a generated Ticket Action Pack Markdown file.

Checks:
- Exactly one fenced code block labeled 'bash'
- No other code fences
- Exactly 20 steps, numbered '# 01)'..'# 20)' within the bash block

Usage:
  python3 scripts/validate_ticket_action_pack.py path/to/ticket-action-pack.md
"""

from __future__ import annotations

import re
import sys
from pathlib import Path


FENCE_RE = re.compile(r"^```(\w+)?\s*$")
STEP_RE = re.compile(r"^#\s+(\d{2})\)\s*$")


def die(msg: str, code: int = 2) -> "None":
    sys.stderr.write(f"ERROR: {msg}\n")
    raise SystemExit(code)


def main() -> int:
    if len(sys.argv) != 2:
        die("expected exactly one argument: path to pack markdown")

    path = Path(sys.argv[1])
    if not path.is_file():
        die(f"file not found: {path}")

    raw = path.read_text(encoding="utf-8", errors="replace")
    if "{{" in raw and "}}" in raw:
        die("unresolved template placeholders ({{...}}) detected in pack")
    lines = raw.splitlines()

    fence_idxs: list[tuple[int, str]] = []
    for i, line in enumerate(lines):
        m = FENCE_RE.match(line)
        if m:
            fence_idxs.append((i, (m.group(1) or "").strip()))

    if len(fence_idxs) != 2:
        die(f"expected exactly 1 code fence pair (2 fence lines), found {len(fence_idxs)}")

    open_i, open_lang = fence_idxs[0]
    close_i, close_lang = fence_idxs[1]

    if open_lang != "bash":
        die(f"expected the only code fence to be ```bash, found ```{open_lang or '(none)'}")

    if close_lang != "":
        # closing fence should be ``` (no language)
        die("closing code fence must be plain ```")

    if close_i <= open_i:
        die("invalid fence ordering")

    bash_block = lines[open_i + 1 : close_i]

    step_nums: list[int] = []
    for line in bash_block:
        m = STEP_RE.match(line)
        if m:
            step_nums.append(int(m.group(1)))

    if len(step_nums) != 20:
        die(f"expected 20 step headers inside bash block, found {len(step_nums)}")

    expected = list(range(1, 21))
    if step_nums != expected:
        die(f"step numbering mismatch; expected {expected}, found {step_nums}")

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
