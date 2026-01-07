# path: oraclepack-usecase-rpg/scripts/validate_oraclepack_pack.py
from __future__ import annotations

import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import Dict, List, Tuple


REQUIRED_CATEGORIES = [
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

BASH_FENCE_RE = re.compile(r"(?s)```bash\n(.*?)\n```")
STEP_HEADER_RE = re.compile(r"^#\s*(\d{2})(?:\)|[\s]+[—-])")
TOKEN_RE = {
    "roi": re.compile(r"\bROI=(\d+(?:\.\d+)?)\b"),
    "impact": re.compile(r"\bimpact=(\d+(?:\.\d+)?)\b"),
    "confidence": re.compile(r"\bconfidence=(\d+(?:\.\d+)?)\b"),
    "effort": re.compile(r"\beffort=(\d+(?:\.\d+)?)\b"),
    "horizon": re.compile(r"\bhorizon=(Immediate|Strategic)\b"),
    "category": re.compile(r"\bcategory=(contracts/interfaces|invariants|caching/state|background jobs|observability|permissions|migrations|UX flows|failure modes|feature flags)\b"),
    "reference": re.compile(r"\breference=([^\s]+)\b"),
}
ORACLE_INVOCATION_RE = re.compile(r"(?m)^\s*(oracle|\$oracle_cmd|\$\{oracle_cmd\})\b")
WRITE_OUTPUT_RE = re.compile(r'--write-output\s+["\\]?([^"\\]+)["\\]?')
PROMPT_FLAG_RE = re.compile(r"\s-\s*p\s|\s-p\s")


@dataclass
class Step:
    step_id: str  # "01"
    header_line: str
    body: str


def die(msg: str) -> None:
    print(f"ERROR: {msg}", file=sys.stderr)
    sys.exit(1)


def warn(msg: str) -> None:
    print(f"WARNING: {msg}", file=sys.stderr)


def parse_bash_block(md: str) -> str:
    matches = list(BASH_FENCE_RE.finditer(md))
    if len(matches) != 1:
        die(f"Expected exactly one fenced bash block starting with ```bash, found {len(matches)}.")
    return matches[0].group(1)


def split_steps(bash: str) -> List[Step]:
    lines = bash.splitlines(True)  # keep \n
    headers: List[Tuple[int, str, str]] = []  # (line_idx, step_id, header_line)
    for i, line in enumerate(lines):
        m = STEP_HEADER_RE.match(line.strip())
        if m:
            headers.append((i, m.group(1), line.rstrip("\n")))

    if len(headers) != 20:
        die(f"Expected exactly 20 step headers, found {len(headers)}.")

    # Ensure sequential 01..20 by appearance
    expected = [f"{i:02d}" for i in range(1, 21)]
    got = [sid for _, sid, _ in headers]
    if got != expected:
        die(f"Step IDs are not sequential 01..20 in order. Got: {got}")

    steps: List[Step] = []
    for idx, (line_i, sid, header_line) in enumerate(headers):
        start = line_i + 1
        end = headers[idx + 1][0] if idx + 1 < len(headers) else len(lines)
        body = "".join(lines[start:end]).strip("\n")
        steps.append(Step(step_id=sid, header_line=header_line, body=body))
        return steps


def parse_float(token: str, s: str, step_id: str) -> float:
    m = TOKEN_RE[token].search(s)
    if not m:
        die(f"Step {step_id}: missing token {token} in header.")
    return float(m.group(1))


def parse_str(token: str, s: str, step_id: str) -> str:
    m = TOKEN_RE[token].search(s)
    if not m:
        die(f"Step {step_id}: missing token {token} in header.")
    return m.group(1)


def validate_header_tokens(steps: List[Step]) -> Tuple[int, int, Dict[str, int]]:
    immediate = 0
    strategic = 0
    category_counts: Dict[str, int] = {c: 0 for c in REQUIRED_CATEGORIES}

    for st in steps:
        header = st.header_line

        roi = parse_float("roi", header, st.step_id)
        impact = parse_float("impact", header, st.step_id)
        confidence = parse_float("confidence", header, st.step_id)
        effort = parse_float("effort", header, st.step_id)
        horizon = parse_str("horizon", header, st.step_id)
        category = parse_str("category", header, st.step_id)
        _ = parse_str("reference", header, st.step_id)

        for name, val in [("impact", impact), ("confidence", confidence), ("effort", effort)]:
            if not (0.1 <= val <= 1.0):
                die(f"Step {st.step_id}: {name} must be in 0.1..1.0, got {val}.")
            # One decimal requirement (tolerate float formatting issues)
            if round(val, 1) != val:
                die(f"Step {st.step_id}: {name} must have one decimal, got {val}.")

        if effort == 0.0:
            die(f"Step {st.step_id}: effort must be non-zero.")

        expected_roi = (impact * confidence) / effort
        if abs(expected_roi - roi) > 0.02:
            die(
                f"Step {st.step_id}: ROI mismatch. Header ROI={roi}, expected {(impact*confidence)/effort:.2f} "
                f"from impact={impact}, confidence={confidence}, effort={effort}."
            )

        if horizon == "Immediate":
            immediate += 1
        elif horizon == "Strategic":
            strategic += 1
        else:
            die(f"Step {st.step_id}: invalid horizon={horizon}.")

        if category not in category_counts:
            die(f"Step {st.step_id}: category not in required set: {category}")
            category_counts[category] += 1

    return immediate, strategic, category_counts


def validate_step_bodies(steps: List[Step], out_dir_hint: str | None) -> None:
    for st in steps:
        body = st.body

        invocations = ORACLE_INVOCATION_RE.findall(body)
        if len(invocations) < 1:
            die(f"Step {st.step_id}: no oracle invocation detected in body.")
        if len(invocations) > 1:
            warn(f"Step {st.step_id}: multiple oracle-like invocations detected; expected one. Matches={invocations}")

        if "--write-output" not in body:
            die(f"Step {st.step_id}: missing --write-output in oracle invocation.")
        if "-p" not in body:
            die(f"Step {st.step_id}: missing -p prompt body in oracle invocation.")

        m = WRITE_OUTPUT_RE.search(body)
        if not m:
            die(f"Step {st.step_id}: could not parse --write-output path.")
        out_path = m.group(1)

        # Ensure step number is reflected in write-output file path
        if f"/{st.step_id}-" not in out_path and f"{st.step_id}-" not in out_path:
            die(
                f"Step {st.step_id}: --write-output path should include '{st.step_id}-<slug>'. Got: {out_path}"
            )

        # If out_dir provided in prelude, encourage usage (non-fatal)
        if out_dir_hint and out_dir_hint not in out_path and "$out_dir" not in out_path and "${out_dir}" not in out_path:
            warn(f"Step {st.step_id}: write-output path does not appear to use out_dir='{out_dir_hint}' (ok if intentional).")


def extract_out_dir_from_prelude(bash: str) -> str | None:
    # Match: out_dir="oracle-out" or out_dir='oracle-out' or out_dir=oracle-out
    m = re.search(r'(?m)^\s*out_dir\s*=\s*["\\]?([^"\\]+)["\\]?\s*$', bash)
    return m.group(1) if m else None


def validate_coverage(steps: List[Step], category_counts: Dict[str, int], md: str) -> None:
    missing = [c for c, n in category_counts.items() if n == 0]
    if missing:
        die(f"Missing required categories in step headers: {missing}")

    # Coverage check section is a strict requirement for this skill; verify presence.
    if re.search(r'(?im)^##\s*Coverage check\s*$', md) is None:
        warn("Coverage check section heading (## Coverage check) not found (skill expects it).")
        return

    # Light validation that each category appears in the coverage section.
    # (Do not over-parse; formatting may vary.)
    coverage_tail = md.split("## Coverage check", 1)[-1]
    for c in REQUIRED_CATEGORIES:
        if c not in coverage_tail:
            warn(f"Coverage check section does not mention category '{c}' (skill expects OK/Missing lines).")


def main() -> None:
    if len(sys.argv) != 2:
        die("Usage: python3 validate_oraclepack_pack.py <pack.md>")

    path = Path(sys.argv[1])
    if not path.exists():
        die(f"File not found: {path}")

    md = path.read_text(encoding="utf-8")
    bash = parse_bash_block(md)
    steps = split_steps(bash)

    immediate, strategic, category_counts = validate_header_tokens(steps)

    if immediate != 12 or strategic != 8:
        die(f"Horizon mix must be exactly 12 Immediate and 8 Strategic; got Immediate={immediate}, Strategic={strategic}")

    out_dir_hint = extract_out_dir_from_prelude(bash)
    if not out_dir_hint:
        warn('Prelude does not define out_dir="...". oraclepack can derive metadata from this if present.')

    validate_step_bodies(steps, out_dir_hint)
    validate_coverage(steps, category_counts, md)

    print("OK: Pack matches strict oraclepack-usecase-rpg validation rules.")


if __name__ == "__main__":
    main()
