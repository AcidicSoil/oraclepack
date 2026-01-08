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

REQUIRED_HEADER_KEYS = [
    "ROI",
    "impact",
    "confidence",
    "effort",
    "horizon",
    "category",
    "reference",
]


@dataclass(frozen=True)
class Step:
    n: str
    header_line: str
    block_lines: List[str]


def _fail(errors: List[str]) -> None:
    for e in errors:
        print(f"[ERROR] {e}", file=sys.stderr)
    sys.exit(1)


def _read_text(path: Path) -> str:
    return path.read_text(encoding="utf-8", errors="replace")


def _extract_single_bash_fence(lines: List[str]) -> Tuple[int, int, List[str], List[str]]:
    start = None
    for i, ln in enumerate(lines):
        if ln.strip().lower() == "```bash":
            if start is not None:
                raise ValueError("Multiple ```bash fences found; expected exactly one.")
            start = i

    if start is None:
        raise ValueError("No ```bash fence found; expected exactly one.")

    end = None
    for i in range(start + 1, len(lines)):
        if lines[i].strip() == "```":
            end = i
            break

    if end is None:
        raise ValueError("No closing ``` found for the ```bash fence.")

    fence_lines = [ln.rstrip("\n") for ln in lines[start + 1 : end]]
    outside_lines = [ln.rstrip("\n") for i, ln in enumerate(lines) if i < start or i > end]
    return start, end, fence_lines, outside_lines


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
        steps.append(Step(n=n, header_line=block[0], block_lines=block))
    return steps


def _parse_header_kv(header_line: str) -> Dict[str, str]:
    out: Dict[str, str] = {}
    tokens = header_line.strip().split()
    i = 0
    while i < len(tokens):
        tok = tokens[i]
        if "=" not in tok:
            i += 1
            continue
        key, val = tok.split("=", 1)
        if key in REQUIRED_HEADER_KEYS:
            if key == "category" and i + 1 < len(tokens):
                # Allow two-word categories like "background jobs" in headers.
                nxt = tokens[i + 1]
                if (val, nxt) in {
                    ("background", "jobs"),
                    ("UX", "flows"),
                    ("failure", "modes"),
                    ("feature", "flags"),
                }:
                    val = f"{val} {nxt}"
            out[key] = val
        i += 1
    return out


def _validate_header(step: Step, errors: List[str]) -> None:
    m = re.match(r"^#\s*(\d{2})\)\s+", step.header_line)
    if not m:
        errors.append(f"Step {step.n}: invalid header format (expected '# NN) ...').")
        return
    if m.group(1) != step.n:
        errors.append(f"Step {step.n}: header number mismatch (found {m.group(1)}).")

    kv = _parse_header_kv(step.header_line)
    for req in REQUIRED_HEADER_KEYS:
        if req not in kv:
            errors.append(f"Step {step.n}: header missing required token {req}=...")

    cat = kv.get("category")
    if cat is not None and cat not in ALLOWED_CATEGORIES:
        errors.append(
            f"Step {step.n}: category must be one of {ALLOWED_CATEGORIES}; got '{cat}'."
        )


def _validate_write_output(step: Step, errors: List[str]) -> None:
    joined = "\n".join(step.block_lines)

    m = re.search(r'(?<!\S)--write-output(?!\S)\s+"([^"]+)"', joined)
    if not m:
        errors.append(f"Step {step.n}: missing --write-output \"...\".")
        return

    out_path = m.group(1)
    if out_path.startswith("/") or out_path.startswith("~"):
        errors.append(f"Step {step.n}: --write-output must not be absolute: {out_path}")
        return

    p = PurePosixPath(out_path)
    if any(part == ".." for part in p.parts):
        errors.append(f"Step {step.n}: --write-output must not contain '..': {out_path}")

    if re.search(rf"(^|/){re.escape(step.n)}-", out_path) is None:
        errors.append(f"Step {step.n}: --write-output must include '{step.n}-' in filename: {out_path}")

    if not out_path.endswith(".md"):
        errors.append(f"Step {step.n}: --write-output must end with .md: {out_path}")


def _validate_ticket_bundle_reference(step: Step, errors: List[str]) -> None:
    joined = "\n".join(step.block_lines)

    if "_tickets_bundle" not in joined:
        errors.append(
            f"Step {step.n}: must reference the ticket bundle (expected '_tickets_bundle' in step block)."
        )

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
        kv = _parse_header_kv(st.header_line)
        cat = kv.get("category")
        if cat in counts:
            counts[cat].append(st.n)

    bad = []
    for cat, ids in counts.items():
        if len(ids) != 2:
            bad.append(f"{cat}={len(ids)} (steps={ids})")
    if bad:
        errors.append(
            "Category distribution must be exactly 2 steps per category (20 total). Problems: "
            + ", ".join(bad)
        )


def _validate_step_numbers(steps: List[Step], errors: List[str]) -> None:
    nums = [st.n for st in steps]
    if nums != [f"{i:02d}" for i in range(1, 21)]:
        errors.append(f"Step numbering must be 01..20 in order; got {nums}.")


def _validate_coverage_check(outside_lines: List[str], errors: List[str]) -> None:
    text = "\n".join(outside_lines)
    m = re.search(r"^##\s+Coverage check\s*$", text, flags=re.IGNORECASE | re.MULTILINE)
    if m is None:
        errors.append('Missing "## Coverage check" section (must be outside the bash fence).')
        return

    after = text[m.end() :]
    for cat in ALLOWED_CATEGORIES:
        pat = rf"^\s*[-*]\s+{re.escape(cat)}\s*:\s*(OK|Missing\([^)]*\))\s*$"
        if re.search(pat, after, flags=re.MULTILINE) is None:
            errors.append(f'Coverage check missing/invalid line for category: "{cat}"')


def _validate_bash_hazards(step: Step, errors: List[str]) -> None:
    lines = step.block_lines[1:]
    for i, ln in enumerate(lines):
        s = ln.strip()

        if s == "\\":
            errors.append(f"Step {step.n}: contains a bare '\\\\' line (orphan backslash).")

        if s.startswith("#"):
            j = i - 1
            while j >= 0 and not lines[j].strip():
                j -= 1
            if j >= 0 and lines[j].rstrip().endswith("\\"):
                errors.append(
                    f"Step {step.n}: comment line appears immediately after a '\\'-continued line (comment-in-continuation hazard)."
                )

        if s.startswith("-p "):
            j = i - 1
            while j >= 0 and not lines[j].strip():
                j -= 1
            if j < 0 or not lines[j].rstrip().endswith("\\"):
                errors.append(
                    f"Step {step.n}: '-p ...' line is not attached to a continued command (detached -p hazard)."
                )


def validate_pack(path: Path, require_bundle: bool) -> None:
    raw = _read_text(path)
    lines = raw.splitlines(True)

    errors: List[str] = []

    try:
        _start, _end, fence, outside = _extract_single_bash_fence(lines)
    except Exception as e:
        _fail([str(e)])

    steps = _parse_steps(fence)

    if len(steps) != 20:
        errors.append(f"Expected exactly 20 steps; found {len(steps)}.")
    else:
        _validate_step_numbers(steps, errors)

    for step in steps:
        _validate_header(step, errors)
        _validate_write_output(step, errors)
        _validate_answer_format(step, errors)
        _validate_bash_hazards(step, errors)
        if require_bundle:
            _validate_ticket_bundle_reference(step, errors)

    _validate_category_counts(steps, errors)
    _validate_coverage_check(outside, errors)

    if errors:
        _fail(errors)

    print("[OK] Pack validates against tickets Stage-1 contract.")


def main() -> None:
    p = argparse.ArgumentParser(description="Validate oraclepack Stage-1 ticket packs.")
    p.add_argument("pack_path", help="Path to the Markdown pack file")
    p.add_argument(
        "--mode",
        choices=["bundle", "direct"],
        default="direct",
        help="Validation mode: bundle requires _tickets_bundle attachments",
    )
    args = p.parse_args()

    path = Path(args.pack_path)
    if not path.exists():
        _fail([f"File not found: {path}"])

    validate_pack(path, require_bundle=args.mode == "bundle")


if __name__ == "__main__":
    main()
