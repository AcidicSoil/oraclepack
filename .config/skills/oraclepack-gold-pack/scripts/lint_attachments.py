# path: oraclepack-gold-pack/scripts/lint_attachments.py
import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple


@dataclass
class Step:
    n: str
    header: str
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
        block = fence_lines[start_i:end_i]
        steps.append(Step(n=n, header=block[0], lines=block))
    return steps


def _count_attachments(step: Step) -> int:
    count = 0
    for ln in step.lines[1:]:
        s = ln.strip()
        if not s or s.startswith("#"):
            continue
        # Count occurrences in non-comment lines
        count += len(re.findall(r"(?<!\S)(-f|--file)(?!\S)", ln))
    return count


def _is_unknown_reference(step: Step) -> bool:
    # Step header token format contains reference=...
    m = re.search(r"\breference=([^\s]+)", step.header)
    if not m:
        return False
    val = m.group(1).strip()
    return val.lower() == "unknown"


def _has_missing_artifact_request(step: Step) -> bool:
    hay = "\n".join(step.lines).lower()
    # Accept several common phrasings, keep it simple and robust.
    patterns = [
        r"missing file/path pattern",
        r"missing file.*pattern",
        r"attach next",
        r"name the exact missing.*pattern",
    ]
    return any(re.search(p, hay) for p in patterns)

def lint(path: Path) -> None:
    raw = _read_text(path)
    lines = raw.splitlines(True)
    fence = _extract_bash_fence(lines)
    steps = _parse_steps(fence)

    errors: List[str] = []
    for step in steps:
        attachments = _count_attachments(step)
        if attachments > 2:
            errors.append(f"Step {step.n}: has {attachments} attachments; must be <= 2 (minimal attachments rule).")

        if _is_unknown_reference(step) and not _has_missing_artifact_request(step):
            errors.append(
                f"Step {step.n}: reference=Unknown but prompt does not request missing file/path pattern(s) to attach next."
            )

    if errors:
        for e in errors:
            print(f"[ERROR] {e}", file=sys.stderr)
        sys.exit(1)

    print("[OK] Attachment lint passed (<=2 attachments per step; Unknown-reference prompts request missing patterns).")

def main() -> None:
    p = argparse.ArgumentParser(description="Lint oraclepack Stage-1 gold pack attachments (<=2 per step) and Unknown-reference prompt behavior.")
    p.add_argument("pack_path", help="Path to the Markdown pack file")
    args = p.parse_args()

    path = Path(args.pack_path)
    if not path.exists():
        print(f"[ERROR] File not found: {path}", file=sys.stderr)
        sys.exit(1)

    lint(path)


if __name__ == "__main__":
    main()
