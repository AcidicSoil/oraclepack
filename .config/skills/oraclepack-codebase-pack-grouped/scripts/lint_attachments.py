import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple


@dataclass
class Step:
    n: str
    lines: List[str]


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        return path.read_text(encoding="utf-8", errors="replace")


def _extract_bash_fence(lines: List[str]) -> List[str]:
    fence_idxs = [i for i, ln in enumerate(lines) if ln.startswith("```")]
    if len(fence_idxs) != 2:
        raise ValueError(f"Expected exactly one fenced block (2 fence lines). Found {len(fence_idxs)}.")
    open_i, close_i = fence_idxs
    if lines[open_i].rstrip("\n") != "```bash":
        raise ValueError("Opening fence must be exactly ```bash.")
    if lines[close_i].rstrip("\n") != "```":
        raise ValueError("Closing fence must be exactly ```.")
    return [ln.rstrip("\n") for ln in lines[open_i + 1 : close_i]]


def _parse_steps(fence_lines: List[str]) -> List[Step]:
    header_re = re.compile(r"^#\s*(\d{2})\)\s+")
    header_idxs: List[Tuple[int, str]] = []
    for i, ln in enumerate(fence_lines):
        m = header_re.match(ln)
        if m:
            header_idxs.append((i, m.group(1)))

    if not header_idxs:
        raise ValueError("No step headers found inside bash fence.")

    steps: List[Step] = []
    for idx, (start_i, n) in enumerate(header_idxs):
        end_i = header_idxs[idx + 1][0] if idx + 1 < len(header_idxs) else len(fence_lines)
        steps.append(Step(n=n, lines=fence_lines[start_i:end_i]))
    return steps


def lint(path: Path) -> None:
    raw = _read_text(path)
    lines = raw.splitlines(True)
    fence = _extract_bash_fence(lines)
    steps = _parse_steps(fence)

    errors: List[str] = []
    for step in steps:
        joined = "\n".join(step.lines)

        if "_tickets_bundle" in joined:
            errors.append(f"Step {step.n}: found '_tickets_bundle' reference (codebase packs must attach code files directly).")

        if re.search(r"mapfile\s+-t\s+__code_files\s+<\s+<\(", joined) is None:
            errors.append(f"Step {step.n}: missing mapfile code discovery stanza.")

        if re.search(r"code_args=\(\)", joined) is None or re.search(r"code_args\+\=\(\s*(-f|--file)\b", joined) is None:
            errors.append(f"Step {step.n}: missing code_args builder (code_args+=(-f \"$p\")).")

        if re.search(r"\$\{code_args\[@\]\}", joined) is None:
            errors.append(f"Step {step.n}: missing ${'{'}code_args[@]{'}'} usage in oracle invocation.")

    if errors:
        for e in errors:
            print(f"[ERROR] {e}", file=sys.stderr)
        sys.exit(1)

    print("[OK] Direct-code lint passed.")


def main() -> None:
    p = argparse.ArgumentParser(description="Lint codebase-driven Stage-1 packs (direct-code mode).")
    p.add_argument("pack_path", help="Path to the Markdown pack file")
    args = p.parse_args()

    path = Path(args.pack_path)
    if not path.exists():
        print(f"[ERROR] File not found: {path}", file=sys.stderr)
        sys.exit(1)

    lint(path)


if __name__ == "__main__":
    main()
