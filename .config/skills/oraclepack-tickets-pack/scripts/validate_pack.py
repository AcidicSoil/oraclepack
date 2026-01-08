import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path, PurePosixPath
from typing import Dict, List, Tuple


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

# Required header tokens, in strict order.
HEADER_TOKEN_ORDER = [
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
    header_line_no: int  # 1-based within fence
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
        return path.read_text(encoding="utf-8", errors="replace")


def _extract_single_bash_fence(lines: List[str]) -> Tuple[int, int, List[str], List[str]]:
    """
    Enforces:
      - exactly one fenced code block labeled bash
      - no other fences anywhere
      - opening fence line must be exactly ```bash
      - closing fence line must be exactly ```
    """
    fence_locs = [i for i, ln in enumerate(lines) if re.match(r"^```", ln)]
    if len(fence_locs) != 2:
        # Show all fence-like lines to help debugging.
        details = []
        for i, ln in enumerate(lines):
            if re.match(r"^```", ln):
                details.append(f"line {i+1}: {ln.rstrip()}")
        raise ValueError(
            f"Expected exactly one fenced code block (2 fence lines), found {len(fence_locs)} fence line(s). "
            + ("Fences: " + "; ".join(details) if details else "")
        )

    open_i, close_i = fence_locs
    if lines[open_i].rstrip("\n") != "```bash":
        raise ValueError("Opening fence must be exactly ```bash on its own line (no spaces).")
    if lines[close_i].rstrip("\n") != "```":
        raise ValueError("Closing fence must be exactly ``` on its own line (no spaces).")
    if close_i <= open_i:
        raise ValueError("Closing fence appears before opening fence.")

    fence_lines = lines[open_i + 1 : close_i]
    outside_lines = lines[:open_i] + lines[close_i + 1 :]
    return open_i, close_i, fence_lines, outside_lines


def _parse_steps(fence_lines: List[str]) -> List[Step]:
    header_re = re.compile(r"^#\s*(\d{2})\)\s+")
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
                header_line_no=start_i + 1,
                header_line=block[0].rstrip("\n"),
                block_lines=[b.rstrip("\n") for b in block],
            )
        )
    return steps


def _header_token_positions(header: str) -> Dict[str, int]:
    pos: Dict[str, int] = {}
    for t in HEADER_TOKEN_ORDER:
        pos[t] = header.find(t)
    return pos


def _parse_category_value(header: str) -> str:
    if "category=" not in header:
        return ""
    after = header.split("category=", 1)[1]
    # Category ends at the start of " reference=" (strict contract).
    end = after.find(" reference=")
    if end == -1:
        # As a fallback, try other token starts, though contract expects reference= last.
        for token in [" ROI=", " impact=", " confidence=", " effort=", " horizon="]:
            p = after.find(token)
            if p != -1:
                end = p if end == -1 else min(end, p)
    if end == -1:
        cat = after.strip()
    else:
        cat = after[:end].strip()
    return cat


def _has_nonempty_scalar(header: str, key: str) -> bool:
    # scalar value ends at next whitespace
    m = re.search(rf"\b{re.escape(key)}=([^\s]+)", header)
    return bool(m and m.group(1).strip())


def _validate_header(step: Step, errors: List[str]) -> None:
    header = step.header_line

    # Strict start.
    if not re.match(rf"^#\s*{re.escape(step.n)}\)\s+", header):
        errors.append(f"Step {step.n}: header must start with '# {step.n})'. Got: {header}")

    # Tokens must appear in strict order.
    pos = _header_token_positions(header)
    for t, p in pos.items():
        if p == -1:
            errors.append(f"Step {step.n}: missing required token '{t}' in header: {header}")

    # Order check (only if all present).
    if all(p != -1 for p in pos.values()):
        last = -1
        for t in HEADER_TOKEN_ORDER:
            if pos[t] <= last:
                errors.append(
                    f"Step {step.n}: token '{t}' is out of order in header. "
                    f"Expected order: {' '.join(HEADER_TOKEN_ORDER)}. Got: {header}"
                )
                break
            last = pos[t]

    # Non-empty values.
    if not _has_nonempty_scalar(header, "ROI"):
        errors.append(f"Step {step.n}: missing/empty ROI= value in header: {header}")
    for k in ["impact", "confidence", "effort", "horizon", "reference"]:
        if not _has_nonempty_scalar(header, k):
            errors.append(f"Step {step.n}: missing/empty {k}= value in header: {header}")

    cat_val = _parse_category_value(header)
    if not cat_val:
        errors.append(f"Step {step.n}: missing/empty category= value in header: {header}")
    elif cat_val not in ALLOWED_CATEGORIES:
        errors.append(
            f"Step {step.n}: invalid category='{cat_val}'. Must be one of: {ALLOWED_CATEGORIES}. Header: {header}"
        )


def _validate_write_output(step: Step, errors: List[str]) -> None:
    joined = "\n".join(step.block_lines)
    # Strict: must use double quotes exactly: --write-output "<path>"
    m = re.search(r'--write-output\s+"([^"]+)"', joined)
    if not m:
        errors.append(f"Step {step.n}: missing --write-output \"...\" (double-quoted) in step block.")
        return

    out_path = m.group(1)

    # Disallow variable expansions in write paths.
    if "$" in out_path or "`" in out_path:
        errors.append(f"Step {step.n}: --write-output path must not contain shell expansions. Got: {out_path}")

    # Disallow absolute writes (and home shortcuts).
    if out_path.startswith("/") or out_path.startswith("~"):
        errors.append(f"Step {step.n}: --write-output path must be relative (no absolute/home paths). Got: {out_path}")

    # Disallow traversal.
    if re.search(r"(^|/)\.\.(/|$)", out_path):
        errors.append(f"Step {step.n}: --write-output path must not contain '..' traversal. Got: {out_path}")

    # Basic shape: <out_dir>/<nn>-<slug>.md
    if "/" not in out_path:
        errors.append(f"Step {step.n}: --write-output path must contain a directory component. Got: {out_path}")
        return

    filename = out_path.split("/")[-1]
    if not filename.startswith(f"{step.n}-"):
        errors.append(f"Step {step.n}: --write-output filename must start with '{step.n}-'. Got: {filename}")
    if not filename.endswith(".md"):
        errors.append(f"Step {step.n}: --write-output filename must end with '.md'. Got: {filename}")

    # Extra guard: ensure PurePosixPath doesn't include '..' (covers odd strings like 'a/../b').
    try:
        parts = PurePosixPath(out_path).parts
        if ".." in parts:
            errors.append(f"Step {step.n}: --write-output path contains '..' segment (unsafe). Got: {out_path}")
    except Exception:
        # Non-fatal; already handled by regex.
        pass


def _validate_ticket_bundle_reference(step: Step, errors: List[str]) -> None:
    joined = "\n".join(step.block_lines)

    # Require the bundle to be mentioned/attached.
    if "_tickets_bundle" not in joined:
        errors.append(
            f"Step {step.n}: must reference the ticket bundle (expected '_tickets_bundle' in step block)."
        )

    # Require a file attachment pointing to the bundle, double-quoted for stability.
    if re.search(r'(?<!\S)(-f|--file)(?!\S)\s+"[^"\n]*_tickets_bundle[^"\n]*"', joined) is None:
        errors.append(
            f"Step {step.n}: must attach the ticket bundle via -f/--file \"..._tickets_bundle...\"."
        )


def _validate_answer_format(step: Step, errors: List[str]) -> None:
    hay = "\n".join(step.block_lines).lower()
    required = [
        "answer format:",
        "direct answer",
        "risks/unknowns",
        "next smallest concrete experiment",
        "if evidence is insufficient",
        "missing file/path pattern",
    ]
    missing = [s for s in required if s not in hay]
    if missing:
        errors.append(f"Step {step.n}: prompt missing required Answer format components: {missing}")


def _validate_category_counts(steps: List[Step], errors: List[str]) -> None:
    counts: Dict[str, List[str]] = {c: [] for c in ALLOWED_CATEGORIES}
    for st in steps:
        cat = _parse_category_value(st.header_line)
        if cat in counts:
            counts[cat].append(st.n)

    bad = []
    for cat, ids in counts.items():
        if len(ids) != 2:
            bad.append(f"{cat}={len(ids)} (steps={ids})")
    if bad:
        errors.append(
            "Category distribution must be exactly 2 steps per category (20 total). Problems: " + ", ".join(bad)
        )


def _validate_coverage_check(outside_lines: List[str], errors: List[str]) -> None:
    text = "\n".join(outside_lines)
    m = re.search(r"^##\s+Coverage check\s*$", text, flags=re.IGNORECASE | re.MULTILINE)
    if m is None:
        errors.append('Missing "## Coverage check" section (must be outside the bash fence).')
        return

    after = text[m.end() :]
    for cat in ALLOWED_CATEGORIES:
        # Require a line like: "- <cat>: OK" OR "- <cat>: Missing(01,02)"
        pat = rf"^\s*[-*]\s+{re.escape(cat)}\s*:\s*(OK|Missing\([^)]*\))\s*$"
        if re.search(pat, after, flags=re.MULTILINE) is None:
            errors.append(f'Coverage check missing/invalid line for category: "{cat}"')


def validate_pack(path: Path) -> None:
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

    errors: List[str] = []
    for st in steps:
        _validate_header(st, errors)
        _validate_write_output(st, errors)
        _validate_ticket_bundle_reference(st, errors)
        _validate_answer_format(st, errors)

    _validate_category_counts(steps, errors)
    _validate_coverage_check(outside_lines, errors)

    if errors:
        _fail(errors)

    print("[OK] Pack validates against tickets Stage-1 contract.")

def main() -> None:
    p = argparse.ArgumentParser(
        description="Validate a ticket-driven oraclepack Stage-1 pack (single bash fence, 20 steps, strict headers/tokens, safe write paths, ticket bundle references, coverage check)."
    )
    p.add_argument("pack_path", help="Path to the Markdown pack file")
    args = p.parse_args()

    path = Path(args.pack_path)
    if not path.exists():
        _fail([f"File not found: {path}"])

    validate_pack(path)


if __name__ == "__main__":
    main()
