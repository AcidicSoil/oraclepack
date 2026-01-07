# path: oraclepack-gold-pack/scripts/validate_pack.py
import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple, Optional, Dict


ALLOWED_CATEGORIES = [
    "contracts/interfaces",
    "invariants",
    "caching/state",
    "background jobs",
    "observability",
    "permissions",
    "migrations",
    "UX flows",
    "failure modes",
    "feature flags",
]

REQUIRED_TOKENS = [
    "ROI=",
    "impact=",
    "confidence=",
    "effort=",
    "horizon=",
    "category=",
    "reference=",
]


@dataclass
class Step:
    n: str
    header_line_no: int
    header_line: str
    block_lines: List[str]


def _fail(errors: List[str]) -> None:
    for e in errors:
        print(f"[ERROR] {e}", file=sys.stderr)
    sys.exit(1)


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        # fallback
        return path.read_text(encoding="utf-8", errors="replace")


def _extract_single_bash_fence(lines: List[str]) -> Tuple[int, int, List[str], List[str]]:
    """
    Enforces:
      - exactly two fence lines in entire document
      - first fence line must be exactly ```bash
      - closing fence line must be exactly ```
    Returns: (open_idx, close_idx, fence_lines, outside_lines)
    """
    fence_idxs = [i for i, ln in enumerate(lines) if ln.startswith("```")]

    if len(fence_idxs) != 2:
        raise ValueError(
            f"Expected exactly 1 fenced code block (2 fence lines), found {len(fence_idxs)} fence line(s)."
        )

    open_idx, close_idx = fence_idxs
    if lines[open_idx].rstrip("\n") != "```bash":
        raise ValueError("Opening fence must be exactly ```bash (no extra tokens/spaces).")
    if lines[close_idx].rstrip("\n") != "```":
        raise ValueError("Closing fence must be exactly ``` (no extra tokens/spaces).")
    if close_idx <= open_idx:
        raise ValueError("Closing fence appears before opening fence.")

    fence_lines = lines[open_idx + 1 : close_idx]
    outside_lines = lines[:open_idx] + lines[close_idx + 1 :]
    return open_idx, close_idx, fence_lines, outside_lines


def _parse_steps(fence_lines: List[str]) -> List[Step]:
    header_re = re.compile(r"^#\s*(\d{2})\)?\s+")
    header_idxs: List[Tuple[int, str]] = []
    for i, ln in enumerate(fence_lines):
        m = header_re.match(ln)
        if m:
            header_idxs.append((i, m.group(1)))

    if len(header_idxs) != 20:
        raise ValueError(f"Expected exactly 20 step headers inside bash fence, found {len(header_idxs)}.")

    expected = [f"{i:02d}" for i in range(1, 21)]
    got = [n for _, n in header_idxs]
    if got != expected:
        raise ValueError(f"Step numbering must be sequential 01..20. Got: {got}")

    steps: List[Step] = []
    for idx, (start_i, n) in enumerate(header_idxs):
        end_i = header_idxs[idx + 1][0] if idx + 1 < len(header_idxs) else len(fence_lines)
        block = fence_lines[start_i:end_i]
        steps.append(
            Step(
                n=n,
                header_line_no=start_i + 1,  # 1-based within fence
                header_line=block[0].rstrip("\n"),
                block_lines=[b.rstrip("\n") for b in block],
            )
        )
    return steps


def _validate_header(step: Step, errors: List[str]) -> None:
    header = step.header_line

    for tok in REQUIRED_TOKENS:
        if tok not in header:
            errors.append(f"Step {step.n}: missing required token '{tok}' in header: {header}")

    # ensure each token has a value (non-empty, non-space)
    for key in ["ROI", "impact", "confidence", "effort", "horizon", "category", "reference"]:
        m = re.search(rf"\b{re.escape(key)}=([^\s]+)", header)
        if not m:
            errors.append(f"Step {step.n}: header missing/empty '{key}=' value: {header}")

    cat_m = re.search(r"\bcategory=([^\s]+(?:\s+[^\s]+)?)", header)
    # Category can contain a space (e.g., "background jobs", "UX flows", "failure modes", "feature flags")
    # The regex above captures up to two words; handle 2-word categories explicitly by matching allowed set.
    if cat_m:
        # Extract by checking allowed set presence after 'category='
        after = header.split("category=", 1)[1]
        # category value ends at next token start or end of line
        end = len(after)
        for token in [" reference=", " ROI=", " impact=", " confidence=", " effort=", " horizon="]:
            pos = after.find(token)
            if pos != -1:
                end = min(end, pos)
        cat_val = after[:end].strip()
        if cat_val not in ALLOWED_CATEGORIES:
            errors.append(
                f"Step {step.n}: invalid category='{cat_val}'. Must be one of: {ALLOWED_CATEGORIES}"
            )
    else:
        errors.append(f"Step {step.n}: could not parse category=... from header: {header}")


def _validate_write_output(step: Step, errors: List[str]) -> None:
    # Find first --write-output in the step block
    joined = "\n".join(step.block_lines)
    m = re.search(r"--write-output\s+(\"([^\"]+)\"|'([^']+)'|(\S+))", joined)
    if not m:
        errors.append(f"Step {step.n}: missing --write-output in step block.")
        return

    path = m.group(2) or m.group(3) or m.group(4) or ""
    if "/" not in path:
        errors.append(
            f"Step {step.n}: --write-output path must look like <out_dir>/{step.n}-<slug>.md; got: {path}"
        )
        return

    filename = path.split("/")[-1]
    if not filename.startswith(f"{step.n}-"):
        errors.append(
            f"Step {step.n}: --write-output filename must start with '{step.n}-'; got: {filename}"
        )
    if not filename.endswith(".md"):
        errors.append(f"Step {step.n}: --write-output filename must end with .md; got: {filename}")


def _validate_coverage_check(outside_lines: List[str], errors: List[str]) -> None:
    text = "\n".join(outside_lines)
    # Require a "Coverage check" heading (case-insensitive)
    if re.search(r"^##\s+Coverage check\s*$", text, flags=re.IGNORECASE | re.MULTILINE) is None:
        errors.append('Missing "## Coverage check" section (must be outside the bash fence).')
        return

    # Ensure every category appears as "<category>:` somewhere after the heading
    # Find slice after the heading
    m = re.search(r"^##\s+Coverage check\s*$", text, flags=re.IGNORECASE | re.MULTILINE)
    assert m is not None
    after = text[m.end() :]

    missing = []
    for cat in ALLOWED_CATEGORIES:
        # Escape special chars (/, etc.)
        if re.search(rf"^\s*[*-]\s+{re.escape(cat)}\s*:", after, flags=re.MULTILINE) is None:
            missing.append(cat)
    if missing:
        errors.append(f'Coverage check missing category lines for: {missing}')


def validate_pack(path: Path) -> None:
    errors: List[str] = []
    raw = _read_text(path)
    lines = raw.splitlines(True)

    try:
        _, _, fence_lines, outside_lines = _extract_single_bash_fence(lines)
    except ValueError as e:
        _fail([str(e)])

    try:
        steps = _parse_steps(fence_lines)
    except ValueError as e:
        _fail([str(e)])

    for step in steps:
        _validate_header(step, errors)
        _validate_write_output(step, errors)

    _validate_coverage_check(outside_lines, errors)

    if errors:
        _fail(errors)

    print("[OK] Pack validates against gold Stage-1 contract.")


def main() -> None:
    p = argparse.ArgumentParser(
        description="Validate an oraclepack Stage-1 gold pack (single bash fence, 20 steps, strict headers, coverage check)."
    )
    p.add_argument("pack_path", help="Path to the Markdown pack file")
    args = p.parse_args()

    path = Path(args.pack_path)
    if not path.exists():
        _fail([f"File not found: {path}"])

    validate_pack(path)


if __name__ == "__main__":
    main()
