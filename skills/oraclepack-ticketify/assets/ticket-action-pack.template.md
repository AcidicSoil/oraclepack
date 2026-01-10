# Ticket Action Pack (deterministic)

This Action Pack is generated from ticket files and is intended to be runner-ingestible.

Config (resolved at generation time):
- tickets_dir: __TICKETS_DIR__
- ticket_glob: __TICKET_GLOB__
- include: __INCLUDE__
- exclude: __EXCLUDE__
- top_n: __TOP_N__
- out_dir: __OUT_DIR__
- prd_path: __PRD_PATH__
- mode: __MODE__

Only one code fence exists in this document (the `bash` block below).

```bash
# 01) Validate prerequisites and enumerate tickets (fail-fast)
set -euo pipefail
if [ ! -d "__TICKETS_DIR__" ]; then
  echo "ERROR: tickets_dir not found: __TICKETS_DIR__" >&2
  exit 2
fi
mkdir -p "__OUT_DIR__"
python3 - <<'PY'
import fnmatch, glob, hashlib, json, os, re, sys
from pathlib import Path

tickets_dir = "__TICKETS_DIR__"
ticket_glob = "__TICKET_GLOB__"
include_raw = "__INCLUDE__".strip()
exclude_raw = "__EXCLUDE__".strip()
top_n_raw = "__TOP_N__".strip()

def die(msg, code=2):
    sys.stderr.write("ERROR: " + msg + "\n")
    raise SystemExit(code)

try:
    top_n = int(top_n_raw)
except Exception:
    die(f"top_n is not an integer: {top_n_raw!r}")
if top_n < 1 or top_n > 200:
    die(f"top_n out of range (1..200): {top_n}")

repo_root = Path(os.getcwd()).resolve()

def norm_rel(p: Path) -> str:
    rel = p.resolve().relative_to(repo_root)
    return rel.as_posix()

patterns = [ticket_glob]
cands = []
for pat in patterns:
    cands.extend(glob.glob(str(Path(tickets_dir) / pat), recursive=True))

paths = []
for c in cands:
    p = Path(c)
    if p.is_file():
        paths.append(p)

if not paths:
    die(f"no files matched ticket_glob={ticket_glob!r} under {tickets_dir!r}")

rel_paths = [norm_rel(p) for p in paths]
rel_paths = sorted(set(rel_paths))

includes = [s.strip() for s in include_raw.split(",") if s.strip()] if include_raw and include_raw != "__INCLUDE__" else []
excludes = [s.strip() for s in exclude_raw.split(",") if s.strip()] if exclude_raw and exclude_raw != "__EXCLUDE__" else []

def match_any(path: str, pats: list[str]) -> bool:
    return any(fnmatch.fnmatch(path, pat) for pat in pats)

if includes:
    rel_paths = [rp for rp in rel_paths if match_any(rp, includes)]
if excludes:
    rel_paths = [rp for rp in rel_paths if not match_any(rp, excludes)]

if not rel_paths:
    die("no tickets remain after include/exclude filtering")

selected = rel_paths[:top_n]

out = {
    "version": 1,
    "tickets_dir": tickets_dir,
    "ticket_glob": ticket_glob,
    "include": includes,
    "exclude": excludes,
    "top_n": top_n,
    "selected_count": len(selected),
    "tickets": [],
}

frontmatter_re = re.compile(r"\A---\s*\n(.*?)\n---\s*\n", re.S)
kv_re = re.compile(r"^([A-Za-z0-9_-]+)\s*:\s*(.*?)\s*$")

def parse_frontmatter(text: str) -> dict:
    m = frontmatter_re.match(text)
    if not m:
        return {}
    block = m.group(1)
    data = {}
    for line in block.splitlines():
        line = line.strip()
        if not line or line.startswith("#"):
            continue
        km = kv_re.match(line)
        if not km:
            continue
        k, v = km.group(1), km.group(2)
        data[k.strip().lower()] = v.strip().strip('"').strip("'")
    return data

def derive_id(rel_path: str, fm: dict) -> str:
    for key in ("id", "ticket", "key"):
        if key in fm and fm[key]:
            return fm[key]
    # deterministic fallback: filename without extension
    return Path(rel_path).name.rsplit(".", 1)[0]

def derive_title(text: str, fm: dict, rel_path: str) -> str:
    if fm.get("title"):
        return fm["title"]
    for line in text.splitlines():
        if line.startswith("# "):
            return line[2:].strip()
    return Path(rel_path).name

def derive_summary(text: str) -> str:
    lines = []
    for line in text.splitlines():
        s = line.strip()
        if not s:
            if lines:
                break
            continue
        if s.startswith("#"):
            continue
        lines.append(s)
        if len(" ".join(lines)) >= 240:
            break
    return " ".join(lines).strip()

seen_ids = set()
for rp in selected:
    abs_path = (repo_root / rp)
    try:
        raw = abs_path.read_bytes()
    except Exception as e:
        die(f"unreadable ticket: {rp} ({e})")
    text = raw.decode("utf-8", errors="replace")
    sha = hashlib.sha256(raw).hexdigest()
    fm = parse_frontmatter(text)
    tid = derive_id(rp, fm)
    if tid in seen_ids:
        die(f"duplicate derived ticket_id: {tid} (path={rp})")
    seen_ids.add(tid)
    out["tickets"].append({
        "ticket_id": tid,
        "path": rp,
        "title": derive_title(text, fm, rp),
        "summary": derive_summary(text),
        "sha256": sha,
    })

index_path = Path("__OUT_DIR__") / "_tickets_index.json"
index_path.write_text(json.dumps(out, indent=2, sort_keys=True) + "\n", encoding="utf-8")
print(f"Wrote {index_path}")
PY

# 02) Derive canonical actions from tickets -> _actions.json and _actions.md (deterministic)
set -euo pipefail
python3 - <<'PY'
import hashlib, json, os, re, sys
from pathlib import Path

repo_root = Path(os.getcwd()).resolve()
out_dir = Path("__OUT_DIR__")
index_path = out_dir / "_tickets_index.json"
actions_json_path = out_dir / "_actions.json"
actions_md_path = out_dir / "_actions.md"

def die(msg, code=2):
    sys.stderr.write("ERROR: " + msg + "\n")
    raise SystemExit(code)

if not index_path.is_file():
    die(f"missing index file: {index_path}")

index = json.loads(index_path.read_text(encoding="utf-8"))
tickets = index.get("tickets", [])
if not tickets:
    die("index contains zero tickets")

task_item_re = re.compile(r"^\s*[-*]\s+\[( |x|X)\]\s+(.*\S)\s*$")
ac_heading_re = re.compile(r"^\s{0,3}#{2,6}\s+Acceptance Criteria\s*$", re.I)
bullet_re = re.compile(r"^\s*[-*]\s+(.*\S)\s*$")

actions = []
for t in tickets:
    rp = t["path"]
    tid = t["ticket_id"]
    abs_path = (repo_root / rp)
    raw = abs_path.read_text(encoding="utf-8", errors="replace").splitlines()
    in_ac = False
    for i, line in enumerate(raw, start=1):
        m = task_item_re.match(line)
        if m:
            actions.append({
                "ticket_id": tid,
                "ticket_path": rp,
                "source_line": i,
                "kind": "task",
                "text": m.group(2).strip(),
            })
            continue
        if ac_heading_re.match(line):
            in_ac = True
            continue
        if in_ac:
            if line.strip().startswith("#"):
                in_ac = False
                continue
            bm = bullet_re.match(line)
            if bm:
                actions.append({
                    "ticket_id": tid,
                    "ticket_path": rp,
                    "source_line": i,
                    "kind": "acceptance",
                    "text": bm.group(1).strip(),
                })

# stable ordering: by ticket_path then source_line then text
actions.sort(key=lambda a: (a["ticket_path"], a["source_line"], a["text"]))

# assign stable action_id per ticket
per_ticket = {}
for a in actions:
    per_ticket.setdefault(a["ticket_id"], []).append(a)

final = []
for tid in sorted(per_ticket.keys()):
    items = per_ticket[tid]
    for idx, a in enumerate(items, start=1):
        a2 = dict(a)
        a2["action_id"] = f"{tid}-A{idx:02d}"
        final.append(a2)

payload = {
    "version": 1,
    "out_dir": str(out_dir.as_posix()),
    "count": len(final),
    "actions": final,
}

actions_json_path.write_text(json.dumps(payload, indent=2, sort_keys=True) + "\n", encoding="utf-8")

lines = []
lines.append("# Actions (canonical, deterministic)")
lines.append("")
for a in final:
    lines.append(f"- [{a['action_id']}] ({a['ticket_id']}) {a['text']}")
actions_md_path.write_text("\n".join(lines).rstrip() + "\n", encoding="utf-8")

print(f"Wrote {actions_json_path}")
print(f"Wrote {actions_md_path}")
PY

# 03) Generate Task Master PRD (deterministic) -> __PRD_PATH__
set -euo pipefail
mkdir -p "$(python3 - <<'PY'
from pathlib import Path
print(Path("__PRD_PATH__").parent.as_posix())
PY
)"
python3 - <<'PY'
import json, os, sys
from pathlib import Path

out_dir = Path("__OUT_DIR__")
index_path = out_dir / "_tickets_index.json"
actions_path = out_dir / "_actions.json"
prd_path = Path("__PRD_PATH__")

def die(msg, code=2):
    sys.stderr.write("ERROR: " + msg + "\n")
    raise SystemExit(code)

if not index_path.is_file():
    die(f"missing index file: {index_path}")
if not actions_path.is_file():
    die(f"missing actions file: {actions_path}")

index = json.loads(index_path.read_text(encoding="utf-8"))
actions = json.loads(actions_path.read_text(encoding="utf-8"))
tickets = index.get("tickets", [])
acts = actions.get("actions", [])

lines = []
lines.append("# Product Requirements Document (from tickets)")
lines.append("")
lines.append("## Overview")
lines.append("This PRD is deterministically derived from repository ticket Markdown files.")
lines.append("")
lines.append("## Goals")
lines.append("- Convert ticket intent into an actionable, trackable task plan.")
lines.append("")
lines.append("## Non-goals")
lines.append("- Introducing undocumented scope beyond what is present in tickets.")
lines.append("")
lines.append("## Ticket Inventory")
for t in tickets:
    lines.append(f"- [{t['ticket_id']}] {t['title']} ({t['path']})")
lines.append("")
lines.append("## Requirements (by ticket)")
for t in tickets:
    lines.append(f"### {t['ticket_id']}: {t['title']}")
    if t.get("summary"):
        lines.append(t["summary"])
    lines.append("")
lines.append("## Derived Actions")
for a in acts:
    lines.append(f"- {a['action_id']}: {a['text']} (ticket {a['ticket_id']})")
lines.append("")
lines.append("## Acceptance Criteria")
lines.append("- All derived actions are implemented and validated per ticket intent.")
lines.append("")
lines.append("## Open Questions")
lines.append("- TODO: fill unknowns discovered during implementation.")
lines.append("")

prd_path.write_text("\n".join(lines).rstrip() + "\n", encoding="utf-8")
print(f"Wrote {prd_path}")
PY

# 04) (Mode-dependent) Generate pipelines/autopilot docs (deterministic, minimal)
set -euo pipefail
python3 - <<'PY'
from pathlib import Path

mode = "__MODE__".strip()
out_dir = Path("__OUT_DIR__")
out_dir.mkdir(parents=True, exist_ok=True)

if mode == "pipelines":
    p = out_dir / "pipelines.md"
    p.write_text(
        "# Pipelines (draft)\n\n" 
        "Deterministic stub. Replace with concrete pipeline definitions as needed.\n",
        encoding="utf-8",
    )
    print(f"Wrote {p}")
elif mode == "autopilot":
    p = out_dir / "autopilot.md"
    p.write_text(
        "# Autopilot (guarded)\n\n" 
        "This is a deterministic stub entrypoint.\n" 
        "Non-destructive by default. Requires explicit operator review before any execution.\n",
        encoding="utf-8",
    )
    print(f"Wrote {p}")
else:
    # backlog or unknown -> no-op
    pass
PY

# 05) Task Master: parse PRD (fail-fast if task-master missing)
set -euo pipefail
if ! command -v task-master >/dev/null 2>&1;
  then
  echo "ERROR: task-master not found on PATH; cannot parse PRD" >&2
  exit 3
fi
task-master parse-prd "__PRD_PATH__"

# 06) Task Master: analyze complexity -> __OUT_DIR__/tm-complexity.json (fallback if --output unsupported)
set -euo pipefail
if task-master analyze-complexity --output "__OUT_DIR__/tm-complexity.json";
  then
  echo "Wrote __OUT_DIR__/tm-complexity.json"
else
  task-master analyze-complexity > "__OUT_DIR__/tm-complexity.json"
  echo "Wrote __OUT_DIR__/tm-complexity.json (stdout fallback)"
fi

# 07) Task Master: expand all tasks
set -euo pipefail
task-master expand --all

# 08) Review generated artifacts (no-op guidance step)
set -euo pipefail
echo "Review: __OUT_DIR__/_tickets_index.json"
echo "Review: __OUT_DIR__/_actions.json"
echo "Review: __OUT_DIR__/_actions.md"
echo "Review: __PRD_PATH__"
echo "Review: __OUT_DIR__/tm-complexity.json"

# 09) Select next implementation target (no-op guidance step)
set -euo pipefail
echo "Select the next highest-value action from __OUT_DIR__/_actions.md and implement it."

# 10) Implementation step placeholder
set -euo pipefail
echo "TODO: Implement selected action(s). Ensure changes are scoped to ticket intent."

# 11) Verification step placeholder
set -euo pipefail
echo "TODO: Add/Run tests or checks relevant to implemented action(s)."

# 12) Documentation step placeholder
set -euo pipefail
echo "TODO: Update docs/notes linked to implemented action(s)."

# 13) Regression check placeholder
set -euo pipefail
echo "TODO: Run regressions appropriate for the repo."

# 14) Re-sync Task Master tasks placeholder
set -euo pipefail
echo "TODO: If tasks drifted, re-run task-master analyze/expand as needed (operator-driven)."

# 15) Validate determinism placeholder
set -euo pipefail
echo "TODO: Re-run steps 01-03 and confirm _tickets_index.json/_actions.json are stable for unchanged tickets."

# 16) Prepare PR/changeset placeholder
set -euo pipefail
echo "TODO: Prepare a PR with clear mapping to ticket IDs."

# 17) Final review placeholder
set -euo pipefail
echo "TODO: Perform final review against ticket acceptance criteria."

# 18) Release/merge placeholder
set -euo pipefail
echo "TODO: Merge/release per repo policy."

# 19) Post-merge validation placeholder
set -euo pipefail
echo "TODO: Validate post-merge behavior."

# 20) Close-out placeholder
set -euo pipefail
echo "TODO: Close tickets / update status and attach produced artifacts if required."
```