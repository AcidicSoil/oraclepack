# Prelude (allowed inside the single bash fence)
# - Creates out_dir deterministically
# - Builds ticket_bundle_path deterministically from ticket_root/ticket_glob OR ticket_paths
# - Uses lexicographic ordering only (no mtime/timestamps)

set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-08"

python3 - <<'PY'
from __future__ import annotations

import os
from pathlib import Path

CODEBASE_NAME = "Unknown"
OUT_DIR = "docs/oracle-questions-2026-01-08"
TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = "".strip()
BUNDLE_PATH = "docs/oracle-questions-2026-01-08/_tickets_bundle.md"

root = Path(TICKET_ROOT)

def read_text(p: Path) -> str:
    try:
        return p.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        return p.read_text(encoding="utf-8", errors="replace")

def title_from_md(text: str) -> str:
    for ln in text.splitlines():
        s = ln.strip()
        if s.startswith("#"):
            return s.lstrip("#").strip()[:160] or "Untitled"
    for ln in text.splitlines():
        s = ln.strip()
        if s:
            return s[:160]
    return "Untitled"

def select_key_sections(text: str) -> str:
    # Heuristic: if small, include all; else include common ticket sections + top excerpt.
    lines = text.splitlines()
    if len(text) <= 8000 and len(lines) <= 250:
        return text

    keep = []
    wanted = {"description", "context", "problem", "proposal", "solution", "acceptance criteria", "ac", "steps", "repro", "expected", "actual", "notes", "links"}
    i = 0
    while i < len(lines):
        ln = lines[i]
        s = ln.strip()
        if s.startswith("##"):
            hdr = s.lstrip("#").strip().lower()
            if hdr in wanted:
                keep.append(ln)
                i += 1
                # capture until next heading
                while i < len(lines) and not lines[i].lstrip().startswith("#"):
                    keep.append(lines[i])
                    i += 1
                continue
        i += 1

    # Fallback excerpt if no sections matched
    if not any(l.strip() for l in keep):
        excerpt = "\n".join(lines[:200])
        return excerpt + "\n\n[... truncated ...]\n"
    return "\n".join(keep) + "\n\n[... truncated ...]\n"

def resolve_ticket_files() -> list[Path]:
    if TICKET_PATHS:
        items = [p.strip() for p in TICKET_PATHS.split(",") if p.strip()]
        return sorted([Path(p) for p in items], key=lambda p: str(p))
    if not root.exists():
        return []
    return sorted(list(root.glob(TICKET_GLOB)), key=lambda p: str(p))

tickets = resolve_ticket_files()

bundle_lines = []
bundle_lines.append(f"# Tickets bundle — {CODEBASE_NAME}")
bundle_lines.append("")
bundle_lines.append("## Selection rules (deterministic)")
bundle_lines.append(f"- ticket_root: {TICKET_ROOT}")
bundle_lines.append(f"- ticket_glob: {TICKET_GLOB}")
bundle_lines.append(f"- ticket_paths: {TICKET_PATHS or '(none)'}")
bundle_lines.append("- ordering: lexicographic by path")
bundle_lines.append("")

if not tickets:
    bundle_lines.append("## WARNING")
    if TICKET_PATHS:
        bundle_lines.append("- No tickets were found from ticket_paths (check paths).")
    else:
        bundle_lines.append("- ticket_root missing or no tickets matched the glob.")
    bundle_lines.append("- This bundle is empty; Step 01 should request which ticket paths to attach next.")
    bundle_lines.append("")

for p in tickets:
    try:
        txt = read_text(p)
    except Exception as e:
        txt = f"[ERROR reading file: {e}]"
    title = title_from_md(txt)
    body = select_key_sections(txt)
    bundle_lines.append(f"## {title}")
    bundle_lines.append(f"- path: {p}")
    bundle_lines.append("")
    bundle_lines.append(body)
    bundle_lines.append("")
    bundle_lines.append("---")
    bundle_lines.append("")

out_path = Path(BUNDLE_PATH)
out_path.parent.mkdir(parents=True, exist_ok=True)
out_path.write_text("\n".join(bundle_lines).rstrip() + "\n", encoding="utf-8")

print(f"OK: wrote ticket bundle: {out_path}")
PY