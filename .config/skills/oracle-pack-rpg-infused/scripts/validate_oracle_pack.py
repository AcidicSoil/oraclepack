#!/usr/bin/env python3
"""
Validate an oracle-pack (RPG-infused) markdown file against the strict contract.

Usage:
  scripts/validate_oracle_pack.py docs/oracle-pack-YYYY-MM-DD.md
"""

from __future__ import annotations

import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Optional, Tuple


REQUIRED_CATEGORIES: List[str] = [
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


@dataclass(frozen=True)
class Step:
    nn: int
    roi: float
    impact: float
    confidence: float
    effort: float
    horizon: str
    category: str
    reference: str
    cmd_line: str


def die(msg: str) -> None:
    print(f"[ERROR] {msg}", file=sys.stderr)
    raise SystemExit(1)


def ok(msg: str) -> None:
    print(f"[OK] {msg}")


def read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8")
    except Exception as e:
        die(f"Failed to read {path}: {e}")
        raise  # unreachable


def find_single_bash_block(md: str) -> str:
    blocks = re.findall(r"```bash\s*\n(.*?)\n```", md, flags=re.DOTALL)
    if len(blocks) != 1:
        die(f"Expected exactly 1 fenced bash block; found {len(blocks)}.")
    # Ensure no other fenced blocks exist
    fence_count = len(re.findall(r"```", md))
    if fence_count != 2:
        die(f"Expected exactly 1 fenced code block total (2 fences); found {fence_count} fences.")
    return blocks[0]


def parse_out_dir(md: str) -> str:
    m = re.search(r"(?m)^\s*-\s*out_dir:\s*([^\s]+)\s*$", md)
    if not m:
        return "oracle-out"
    return m.group(1).strip()


HEADER_RE = re.compile(
    r"""
    ^\n    #\s*\n    (?P<nn>\d{2})
    \s+
    ROI=(?P<roi>\d+(?:\.\d)?)
    \s+
    impact=(?P<impact>[01]\.\d)
    \s+
    confidence=(?P<confidence>[01]\.\d)
    \s+
    effort=(?P<effort>[01]\.\d)
    \s+
    horizon=(?P<horizon>Immediate|Strategic)
    \s+
    category=(?P<category>.+?)
    \s+
    reference=(?P<reference>.+)
    $
    """,
    flags=re.VERBOSE,
)


def parse_steps(bash: str) -> List[Step]:
    lines = [ln.rstrip("\n") for ln in bash.splitlines()]
    steps: List[Step] = []

    i = 0
    while i < len(lines):
        ln = lines[i].strip()
        if not ln:
            i += 1
            continue

        m = HEADER_RE.match(ln)
        if not m:
            die(f"Unexpected line in bash block (expected step header): {lines[i]}")
        nn = int(m.group("nn"))
        roi = float(m.group("roi"))
        impact = float(m.group("impact"))
        confidence = float(m.group("confidence")),
        effort = float(m.group("effort")),
        horizon = m.group("horizon")
        category = m.group("category").strip()
        reference = m.group("reference").strip()

        # Next non-empty, non-comment line must be the single oracle command
        j = i + 1
        cmd_line: Optional[str] = None
        while j < len(lines):
            nxt = lines[j].strip()
            if not nxt:
                j += 1
                continue
            if nxt.startswith("#"):
                die(f"Found comment line where oracle command line was expected after step {nn:02d}.")
            cmd_line = lines[j].strip()
            j += 1
            break
        if cmd_line is None:
            die(f"Missing oracle command line after step header {nn:02d}.")

        # Ensure there are no extra non-empty, non-comment lines before next header
        k = j
        while k < len(lines):
            peek = lines[k].strip()
            if not peek:
                k += 1
                continue
            if HEADER_RE.match(peek):
                break
            if peek.startswith("#"):
                die(f"Unexpected comment line inside step body for {nn:02d}; each step must be exactly two lines.")
            die(f"Unexpected extra command/content line inside step {nn:02d}: {lines[k]}")
        i = k

        steps.append(
            Step(
                nn=nn,
                roi=roi,
                impact=impact,
                confidence=confidence,
                effort=effort,
                horizon=horizon,
                category=category,
                reference=reference,
                cmd_line=cmd_line,
            )
        )

    return steps


def require_20_steps(steps: List[Step]) -> None:
    if len(steps) != 20:
        die(f"Expected exactly 20 steps; found {len(steps)}.")

    nns = [s.nn for s in steps]
    expected = list(range(1, 21))
    if nns != expected:
        die(f"Step numbers must be 01..20 in order; found: {[f'{n:02d}' for n in nns]}")


def validate_scores(steps: List[Step]) -> None:
    for s in steps:
        for name, v in [("impact", s.impact), ("confidence", s.confidence), ("effort", s.effort)]:
            if v < 0.1 or v > 1.0:
                die(f"Step {s.nn:02d} {name} must be in [0.1..1.0]; got {v}.")
            # One decimal check (string form after rounding to 1 decimal should match)
            if abs(v - round(v, 1)) > 1e-9:
                die(f"Step {s.nn:02d} {name} must have exactly one decimal; got {v}.")

        expected_roi = round((s.impact * s.confidence) / s.effort, 1)
        if abs(s.roi - expected_roi) > 0.05:
            die(
                f"Step {s.nn:02d} ROI mismatch: header ROI={s.roi} but computed ROI={expected_roi} "
                f"from (impact*confidence)/effort."
            )
        if abs(s.roi - round(s.roi, 1)) > 1e-9:
            die(f"Step {s.nn:02d} ROI must have exactly one decimal; got {s.roi}.")


def validate_sort_order(steps: List[Step]) -> None:
    for a, b in zip(steps, steps[1:]):
        if b.roi > a.roi + 1e-9:
            die(
                f"Steps must be sorted by ROI descending; step {a.nn:02d} ROI={a.roi} "
                f"precedes step {b.nn:02d} ROI={b.roi}."
            )
        if abs(b.roi - a.roi) < 1e-9 and b.effort < a.effort - 1e-9:
            die(
                f"ROI tie-break must prefer lower effort first; step {a.nn:02d} effort={a.effort} "
                f"precedes step {b.nn:02d} effort={b.effort} with equal ROI={a.roi}."
            )


def validate_horizon_mix(steps: List[Step]) -> None:
    immediate = sum(1 for s in steps if s.horizon == "Immediate")
    strategic = sum(1 for s in steps if s.horizon == "Strategic")
    if immediate != 12 or strategic != 8:
        die(f"Horizon mix must be exactly 12 Immediate and 8 Strategic; got {immediate} Immediate, {strategic} Strategic.")


def validate_categories(steps: List[Step]) -> None:
    allowed = set(REQUIRED_CATEGORIES)
    for s in steps:
        if s.category not in allowed:
            die(f"Step {s.nn:02d} category must be one of the required categories; got: {s.category!r}")

    covered = {s.category for s in steps}
    missing = [c for c in REQUIRED_CATEGORIES if c not in covered]
    if missing:
        die(f"All required categories must appear at least once across the 20 steps. Missing: {missing}")


def validate_commands(steps: List[Step], out_dir: str) -> None:
    for s in steps:
        if "--files-report" not in s.cmd_line:
            die(f"Step {s.nn:02d} oracle command must include --files-report (or oracle_flags equivalent).")

        # --write-output "<out_dir>/<nn>-<slug>.md"
        m = re.search(r'--write-output\s+"([^"]+)"', s.cmd_line)
        if not m:
            die(f"Step {s.nn:02d} oracle command missing --write-output \"...\".")
        out_path = m.group(1)
        expected_prefix = f"{out_dir}/{s.nn:02d}-"
        if not out_path.startswith(expected_prefix) or not out_path.endswith(".md"):
            die(
                f"Step {s.nn:02d} write-output path must start with {expected_prefix!r} and end with '.md'; got {out_path!r}."
            )
        slug = out_path[len(expected_prefix) : -3]
        if not re.fullmatch(r"[a-z0-9]+(?:-[a-z0-9]+)*", slug or ""):
            die(f"Step {s.nn:02d} slug must be kebab-case [a-z0-9-]; got {slug!r}.")

        if " -p " not in f" {s.cmd_line} " and not re.search(r"\s-p\s", s.cmd_line):
            die(f"Step {s.nn:02d} oracle command must include -p with the strategist prompt body.")

        # Minimal prompt-structure checks (best-effort; do not overfit quoting style)
        if f"Strategist question #{s.nn:02d}" not in s.cmd_line:
            die(f"Step {s.nn:02d} -p prompt must include 'Strategist question #{s.nn:02d}'.")
        for required_snip in ["RPG:", "FunctionalNode:", "StructuralNode:", "DependsOn:", "Phase:", "Answer format:"]:
            if required_snip not in s.cmd_line:
                die(f"Step {s.nn:02d} -p prompt missing required section marker: {required_snip}")


def validate_coverage_check_section(md: str) -> None:
    m = re.search(r"(?ms)^## Coverage check.*?\n(.*?)(?:\n## |\Z)", md)
    if not m:
        die("Missing '## Coverage check' section.")

    body = m.group(1).strip("\n")
    lines = [ln.strip() for ln in body.splitlines() if ln.strip()]
    items: List[Tuple[str, str]] = []
    for ln in lines:
        if not ln.startswith("- "):
            continue
        rest = ln[2:]
        if ":" not in rest:
            continue
        cat, status = [x.strip() for x in rest.split(":", 1)]
        items.append((cat, status))

    if not items:
        die("Coverage check section did not contain any '- category: status' items.")

    cats = [c for c, _ in items]
    if cats != REQUIRED_CATEGORIES:
        die(
            "Coverage check categories must appear exactly once each and in the required order.\n"
            f"Expected: {REQUIRED_CATEGORIES}\n"
            f"Found:    {cats}"
        )

    for cat, status in items:
        if status == "OK":
            continue
        if not status.startswith("Missing(") or not status.endswith(")"):
            die(f"Coverage check status for {cat!r} must be 'OK' or 'Missing(...)'; got {status!r}.")


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("markdown_path", help="Path to docs/oracle-pack-YYYY-MM-DD.md")
    args = ap.parse_args()

    path = Path(args.markdown_path)
    if not path.exists():
        die(f"File not found: {path}")

    md = read_text(path)

    bash = find_single_bash_block(md)
    out_dir = parse_out_dir(md)

    steps = parse_steps(bash)
    require_20_steps(steps)
    validate_scores(steps)
    validate_sort_order(steps)
    validate_horizon_mix(steps)
    validate_categories(steps)
    validate_commands(steps, out_dir)
    validate_coverage_check_section(md)

    ok("oracle-pack is valid.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
