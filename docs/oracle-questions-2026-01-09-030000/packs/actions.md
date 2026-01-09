# Oracle Pack — Unknown (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: Unknown
- out_dir: docs/oracle-questions-2026-01-09-030000/actions
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md
- ticket_max_files: 6
- group_name: actions
- group_slug: actions
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-09-030000/actions/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-09-030000/actions"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/02-contracts-interfaces-integration-points-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/02-contracts-interfaces-integration-points-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/02-contracts-interfaces-integration-points-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/02-contracts-interfaces-integration-points-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/03-invariants-invariant-map-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/03-invariants-invariant-map-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/03-invariants-invariant-map-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/03-invariants-invariant-map-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/04-invariants-validation-boundaries-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation); propose minimal validation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/04-invariants-validation-boundaries-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation); propose minimal validation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/04-invariants-validation-boundaries-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation); propose minimal validation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/04-invariants-validation-boundaries-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation); propose minimal validation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 05) ROI=4.4 impact=6 confidence=0.78 effort=2 horizon=NearTerm category=caching/state reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/05-caching-state-state-artifacts-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #05  (ticket-driven, group: actions)

Reference: actions
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.78, effort=2)

Question:
Using the attached tickets as the primary context, identify state/artifacts that must be produced and preserved; schema/format expectations; stability/back-compat requirements.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/05-caching-state-state-artifacts-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #05  (ticket-driven, group: actions)

Reference: actions
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.78, effort=2)

Question:
Using the attached tickets as the primary context, identify state/artifacts that must be produced and preserved; schema/format expectations; stability/back-compat requirements.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/05-caching-state-state-artifacts-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #05  (ticket-driven, group: actions)

Reference: actions
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.78, effort=2)

Question:
Using the attached tickets as the primary context, identify state/artifacts that must be produced and preserved; schema/format expectations; stability/back-compat requirements.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/05-caching-state-state-artifacts-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #05  (ticket-driven, group: actions)

Reference: actions
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.78, effort=2)

Question:
Using the attached tickets as the primary context, identify state/artifacts that must be produced and preserved; schema/format expectations; stability/back-compat requirements.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 06) ROI=4.2 impact=6 confidence=0.75 effort=2 horizon=NearTerm category=caching/state reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/06-caching-state-cache-keys-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #06  (ticket-driven, group: actions)

Reference: actions
Category: caching/state
Horizon: NearTerm
ROI: 4.2 (impact=6, confidence=0.75, effort=2)

Question:
Using the attached tickets as the primary context, identify any caching opportunities/risks (discovery caches, pack outputs, oracle outputs); define cache keys, invalidation, and correctness risks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/06-caching-state-cache-keys-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #06  (ticket-driven, group: actions)

Reference: actions
Category: caching/state
Horizon: NearTerm
ROI: 4.2 (impact=6, confidence=0.75, effort=2)

Question:
Using the attached tickets as the primary context, identify any caching opportunities/risks (discovery caches, pack outputs, oracle outputs); define cache keys, invalidation, and correctness risks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/06-caching-state-cache-keys-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #06  (ticket-driven, group: actions)

Reference: actions
Category: caching/state
Horizon: NearTerm
ROI: 4.2 (impact=6, confidence=0.75, effort=2)

Question:
Using the attached tickets as the primary context, identify any caching opportunities/risks (discovery caches, pack outputs, oracle outputs); define cache keys, invalidation, and correctness risks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/06-caching-state-cache-keys-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #06  (ticket-driven, group: actions)

Reference: actions
Category: caching/state
Horizon: NearTerm
ROI: 4.2 (impact=6, confidence=0.75, effort=2)

Question:
Using the attached tickets as the primary context, identify any caching opportunities/risks (discovery caches, pack outputs, oracle outputs); define cache keys, invalidation, and correctness risks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 07) ROI=4.3 impact=6 confidence=0.70 effort=2 horizon=MidTerm category=background jobs reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/07-background-jobs-job-model-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #07  (ticket-driven, group: actions)

Reference: actions
Category: background jobs
Horizon: MidTerm
ROI: 4.3 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, identify any background/async work implied (jobs, queues, long-running operations); define responsibilities and interfaces.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/07-background-jobs-job-model-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #07  (ticket-driven, group: actions)

Reference: actions
Category: background jobs
Horizon: MidTerm
ROI: 4.3 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, identify any background/async work implied (jobs, queues, long-running operations); define responsibilities and interfaces.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/07-background-jobs-job-model-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #07  (ticket-driven, group: actions)

Reference: actions
Category: background jobs
Horizon: MidTerm
ROI: 4.3 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, identify any background/async work implied (jobs, queues, long-running operations); define responsibilities and interfaces.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/07-background-jobs-job-model-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #07  (ticket-driven, group: actions)

Reference: actions
Category: background jobs
Horizon: MidTerm
ROI: 4.3 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, identify any background/async work implied (jobs, queues, long-running operations); define responsibilities and interfaces.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 08) ROI=4.0 impact=6 confidence=0.68 effort=3 horizon=MidTerm category=background jobs reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/08-background-jobs-queue-failure-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #08  (ticket-driven, group: actions)

Reference: actions
Category: background jobs
Horizon: MidTerm
ROI: 4.0 (impact=6, confidence=0.68, effort=3)

Question:
Using the attached tickets as the primary context, define how background failures are handled (retries, idempotency, poison messages); define observability hooks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/08-background-jobs-queue-failure-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #08  (ticket-driven, group: actions)

Reference: actions
Category: background jobs
Horizon: MidTerm
ROI: 4.0 (impact=6, confidence=0.68, effort=3)

Question:
Using the attached tickets as the primary context, define how background failures are handled (retries, idempotency, poison messages); define observability hooks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/08-background-jobs-queue-failure-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #08  (ticket-driven, group: actions)

Reference: actions
Category: background jobs
Horizon: MidTerm
ROI: 4.0 (impact=6, confidence=0.68, effort=3)

Question:
Using the attached tickets as the primary context, define how background failures are handled (retries, idempotency, poison messages); define observability hooks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/08-background-jobs-queue-failure-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #08  (ticket-driven, group: actions)

Reference: actions
Category: background jobs
Horizon: MidTerm
ROI: 4.0 (impact=6, confidence=0.68, effort=3)

Question:
Using the attached tickets as the primary context, define how background failures are handled (retries, idempotency, poison messages); define observability hooks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 09) ROI=4.7 impact=7 confidence=0.76 effort=1 horizon=Immediate category=observability reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/09-observability-logging-metrics-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #09  (ticket-driven, group: actions)

Reference: actions
Category: observability
Horizon: Immediate
ROI: 4.7 (impact=7, confidence=0.76, effort=1)

Question:
Using the attached tickets as the primary context, define what logging/metrics must exist to debug pack generation + step execution; propose minimal instrumentation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/09-observability-logging-metrics-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #09  (ticket-driven, group: actions)

Reference: actions
Category: observability
Horizon: Immediate
ROI: 4.7 (impact=7, confidence=0.76, effort=1)

Question:
Using the attached tickets as the primary context, define what logging/metrics must exist to debug pack generation + step execution; propose minimal instrumentation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/09-observability-logging-metrics-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #09  (ticket-driven, group: actions)

Reference: actions
Category: observability
Horizon: Immediate
ROI: 4.7 (impact=7, confidence=0.76, effort=1)

Question:
Using the attached tickets as the primary context, define what logging/metrics must exist to debug pack generation + step execution; propose minimal instrumentation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/09-observability-logging-metrics-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #09  (ticket-driven, group: actions)

Reference: actions
Category: observability
Horizon: Immediate
ROI: 4.7 (impact=7, confidence=0.76, effort=1)

Question:
Using the attached tickets as the primary context, define what logging/metrics must exist to debug pack generation + step execution; propose minimal instrumentation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 10) ROI=4.5 impact=7 confidence=0.74 effort=2 horizon=Immediate category=observability reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/10-observability-tracing-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #10  (ticket-driven, group: actions)

Reference: actions
Category: observability
Horizon: Immediate
ROI: 4.5 (impact=7, confidence=0.74, effort=2)

Question:
Using the attached tickets as the primary context, define tracing/correlation strategy across pack steps and downstream tools; identify required IDs and propagation.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/10-observability-tracing-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #10  (ticket-driven, group: actions)

Reference: actions
Category: observability
Horizon: Immediate
ROI: 4.5 (impact=7, confidence=0.74, effort=2)

Question:
Using the attached tickets as the primary context, define tracing/correlation strategy across pack steps and downstream tools; identify required IDs and propagation.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/10-observability-tracing-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #10  (ticket-driven, group: actions)

Reference: actions
Category: observability
Horizon: Immediate
ROI: 4.5 (impact=7, confidence=0.74, effort=2)

Question:
Using the attached tickets as the primary context, define tracing/correlation strategy across pack steps and downstream tools; identify required IDs and propagation.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/10-observability-tracing-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #10  (ticket-driven, group: actions)

Reference: actions
Category: observability
Horizon: Immediate
ROI: 4.5 (impact=7, confidence=0.74, effort=2)

Question:
Using the attached tickets as the primary context, define tracing/correlation strategy across pack steps and downstream tools; identify required IDs and propagation.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 11) ROI=4.1 impact=6 confidence=0.70 effort=2 horizon=NearTerm category=permissions reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/11-permissions-authz-gaps-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #11  (ticket-driven, group: actions)

Reference: actions
Category: permissions
Horizon: NearTerm
ROI: 4.1 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, identify permission/authz boundaries implied by tickets (file access, command execution, network); propose safe defaults.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/11-permissions-authz-gaps-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #11  (ticket-driven, group: actions)

Reference: actions
Category: permissions
Horizon: NearTerm
ROI: 4.1 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, identify permission/authz boundaries implied by tickets (file access, command execution, network); propose safe defaults.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/11-permissions-authz-gaps-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #11  (ticket-driven, group: actions)

Reference: actions
Category: permissions
Horizon: NearTerm
ROI: 4.1 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, identify permission/authz boundaries implied by tickets (file access, command execution, network); propose safe defaults.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/11-permissions-authz-gaps-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #11  (ticket-driven, group: actions)

Reference: actions
Category: permissions
Horizon: NearTerm
ROI: 4.1 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, identify permission/authz boundaries implied by tickets (file access, command execution, network); propose safe defaults.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 12) ROI=3.9 impact=6 confidence=0.68 effort=2 horizon=NearTerm category=permissions reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/12-permissions-secrets-config-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #12  (ticket-driven, group: actions)

Reference: actions
Category: permissions
Horizon: NearTerm
ROI: 3.9 (impact=6, confidence=0.68, effort=2)

Question:
Using the attached tickets as the primary context, identify secrets/config handling needs (API keys, tokens); propose secure config discovery and redaction.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/12-permissions-secrets-config-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #12  (ticket-driven, group: actions)

Reference: actions
Category: permissions
Horizon: NearTerm
ROI: 3.9 (impact=6, confidence=0.68, effort=2)

Question:
Using the attached tickets as the primary context, identify secrets/config handling needs (API keys, tokens); propose secure config discovery and redaction.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/12-permissions-secrets-config-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #12  (ticket-driven, group: actions)

Reference: actions
Category: permissions
Horizon: NearTerm
ROI: 3.9 (impact=6, confidence=0.68, effort=2)

Question:
Using the attached tickets as the primary context, identify secrets/config handling needs (API keys, tokens); propose secure config discovery and redaction.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/12-permissions-secrets-config-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #12  (ticket-driven, group: actions)

Reference: actions
Category: permissions
Horizon: NearTerm
ROI: 3.9 (impact=6, confidence=0.68, effort=2)

Question:
Using the attached tickets as the primary context, identify secrets/config handling needs (API keys, tokens); propose secure config discovery and redaction.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 13) ROI=3.8 impact=6 confidence=0.66 effort=3 horizon=MidTerm category=migrations reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/13-migrations-schema-migrations-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #13  (ticket-driven, group: actions)

Reference: actions
Category: migrations
Horizon: MidTerm
ROI: 3.8 (impact=6, confidence=0.66, effort=3)

Question:
Using the attached tickets as the primary context, identify any required migrations (schema/format/CLI flags); define migration strategy and compat approach.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/13-migrations-schema-migrations-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #13  (ticket-driven, group: actions)

Reference: actions
Category: migrations
Horizon: MidTerm
ROI: 3.8 (impact=6, confidence=0.66, effort=3)

Question:
Using the attached tickets as the primary context, identify any required migrations (schema/format/CLI flags); define migration strategy and compat approach.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/13-migrations-schema-migrations-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #13  (ticket-driven, group: actions)

Reference: actions
Category: migrations
Horizon: MidTerm
ROI: 3.8 (impact=6, confidence=0.66, effort=3)

Question:
Using the attached tickets as the primary context, identify any required migrations (schema/format/CLI flags); define migration strategy and compat approach.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/13-migrations-schema-migrations-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #13  (ticket-driven, group: actions)

Reference: actions
Category: migrations
Horizon: MidTerm
ROI: 3.8 (impact=6, confidence=0.66, effort=3)

Question:
Using the attached tickets as the primary context, identify any required migrations (schema/format/CLI flags); define migration strategy and compat approach.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 14) ROI=3.7 impact=6 confidence=0.64 effort=3 horizon=MidTerm category=migrations reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/14-migrations-backfill-plan-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #14  (ticket-driven, group: actions)

Reference: actions
Category: migrations
Horizon: MidTerm
ROI: 3.7 (impact=6, confidence=0.64, effort=3)

Question:
Using the attached tickets as the primary context, define any needed backfill/one-time transforms; estimate risks; define verification plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/14-migrations-backfill-plan-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #14  (ticket-driven, group: actions)

Reference: actions
Category: migrations
Horizon: MidTerm
ROI: 3.7 (impact=6, confidence=0.64, effort=3)

Question:
Using the attached tickets as the primary context, define any needed backfill/one-time transforms; estimate risks; define verification plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/14-migrations-backfill-plan-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #14  (ticket-driven, group: actions)

Reference: actions
Category: migrations
Horizon: MidTerm
ROI: 3.7 (impact=6, confidence=0.64, effort=3)

Question:
Using the attached tickets as the primary context, define any needed backfill/one-time transforms; estimate risks; define verification plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/14-migrations-backfill-plan-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #14  (ticket-driven, group: actions)

Reference: actions
Category: migrations
Horizon: MidTerm
ROI: 3.7 (impact=6, confidence=0.64, effort=3)

Question:
Using the attached tickets as the primary context, define any needed backfill/one-time transforms; estimate risks; define verification plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 15) ROI=4.6 impact=6 confidence=0.74 effort=1 horizon=Immediate category=UX flows reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/15-ux-flows-user-journeys-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #15  (ticket-driven, group: actions)

Reference: actions
Category: UX flows
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, identify UX/TUI workflows implied by tickets; define user journey states and expected outputs.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/15-ux-flows-user-journeys-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #15  (ticket-driven, group: actions)

Reference: actions
Category: UX flows
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, identify UX/TUI workflows implied by tickets; define user journey states and expected outputs.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/15-ux-flows-user-journeys-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #15  (ticket-driven, group: actions)

Reference: actions
Category: UX flows
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, identify UX/TUI workflows implied by tickets; define user journey states and expected outputs.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/15-ux-flows-user-journeys-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #15  (ticket-driven, group: actions)

Reference: actions
Category: UX flows
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, identify UX/TUI workflows implied by tickets; define user journey states and expected outputs.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 16) ROI=4.3 impact=6 confidence=0.72 effort=2 horizon=Immediate category=UX flows reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/16-ux-flows-edge-cases-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #16  (ticket-driven, group: actions)

Reference: actions
Category: UX flows
Horizon: Immediate
ROI: 4.3 (impact=6, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify edge cases in UX flows (cancel, resume, partial runs); define minimal UX behavior.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/16-ux-flows-edge-cases-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #16  (ticket-driven, group: actions)

Reference: actions
Category: UX flows
Horizon: Immediate
ROI: 4.3 (impact=6, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify edge cases in UX flows (cancel, resume, partial runs); define minimal UX behavior.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/16-ux-flows-edge-cases-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #16  (ticket-driven, group: actions)

Reference: actions
Category: UX flows
Horizon: Immediate
ROI: 4.3 (impact=6, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify edge cases in UX flows (cancel, resume, partial runs); define minimal UX behavior.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/16-ux-flows-edge-cases-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #16  (ticket-driven, group: actions)

Reference: actions
Category: UX flows
Horizon: Immediate
ROI: 4.3 (impact=6, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify edge cases in UX flows (cancel, resume, partial runs); define minimal UX behavior.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 17) ROI=4.9 impact=7 confidence=0.78 effort=1 horizon=Immediate category=failure modes reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/17-failure-modes-timeouts-retries-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #17  (ticket-driven, group: actions)

Reference: actions
Category: failure modes
Horizon: Immediate
ROI: 4.9 (impact=7, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, define timeouts/retries behavior for external calls; define failure classification and operator actions.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/17-failure-modes-timeouts-retries-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #17  (ticket-driven, group: actions)

Reference: actions
Category: failure modes
Horizon: Immediate
ROI: 4.9 (impact=7, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, define timeouts/retries behavior for external calls; define failure classification and operator actions.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/17-failure-modes-timeouts-retries-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #17  (ticket-driven, group: actions)

Reference: actions
Category: failure modes
Horizon: Immediate
ROI: 4.9 (impact=7, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, define timeouts/retries behavior for external calls; define failure classification and operator actions.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/17-failure-modes-timeouts-retries-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #17  (ticket-driven, group: actions)

Reference: actions
Category: failure modes
Horizon: Immediate
ROI: 4.9 (impact=7, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, define timeouts/retries behavior for external calls; define failure classification and operator actions.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 18) ROI=4.4 impact=7 confidence=0.74 effort=2 horizon=Immediate category=failure modes reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/18-failure-modes-rollback-plan-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #18  (ticket-driven, group: actions)

Reference: actions
Category: failure modes
Horizon: Immediate
ROI: 4.4 (impact=7, confidence=0.74, effort=2)

Question:
Using the attached tickets as the primary context, define rollback plan for partial runs and how to preserve artifacts; define 'safe to re-run' semantics.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/18-failure-modes-rollback-plan-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #18  (ticket-driven, group: actions)

Reference: actions
Category: failure modes
Horizon: Immediate
ROI: 4.4 (impact=7, confidence=0.74, effort=2)

Question:
Using the attached tickets as the primary context, define rollback plan for partial runs and how to preserve artifacts; define 'safe to re-run' semantics.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/18-failure-modes-rollback-plan-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #18  (ticket-driven, group: actions)

Reference: actions
Category: failure modes
Horizon: Immediate
ROI: 4.4 (impact=7, confidence=0.74, effort=2)

Question:
Using the attached tickets as the primary context, define rollback plan for partial runs and how to preserve artifacts; define 'safe to re-run' semantics.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/18-failure-modes-rollback-plan-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #18  (ticket-driven, group: actions)

Reference: actions
Category: failure modes
Horizon: Immediate
ROI: 4.4 (impact=7, confidence=0.74, effort=2)

Question:
Using the attached tickets as the primary context, define rollback plan for partial runs and how to preserve artifacts; define 'safe to re-run' semantics.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 19) ROI=4.0 impact=6 confidence=0.70 effort=2 horizon=NearTerm category=feature flags reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/19-feature-flags-flag-plan-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #19  (ticket-driven, group: actions)

Reference: actions
Category: feature flags
Horizon: NearTerm
ROI: 4.0 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, define feature-flag strategy for rollout (scopes, defaults, telemetry); ensure compat for existing users.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/19-feature-flags-flag-plan-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #19  (ticket-driven, group: actions)

Reference: actions
Category: feature flags
Horizon: NearTerm
ROI: 4.0 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, define feature-flag strategy for rollout (scopes, defaults, telemetry); ensure compat for existing users.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/19-feature-flags-flag-plan-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #19  (ticket-driven, group: actions)

Reference: actions
Category: feature flags
Horizon: NearTerm
ROI: 4.0 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, define feature-flag strategy for rollout (scopes, defaults, telemetry); ensure compat for existing users.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/19-feature-flags-flag-plan-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #19  (ticket-driven, group: actions)

Reference: actions
Category: feature flags
Horizon: NearTerm
ROI: 4.0 (impact=6, confidence=0.70, effort=2)

Question:
Using the attached tickets as the primary context, define feature-flag strategy for rollout (scopes, defaults, telemetry); ensure compat for existing users.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 20) ROI=3.8 impact=6 confidence=0.68 effort=2 horizon=NearTerm category=feature flags reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/20-feature-flags-compat-rollout-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #20  (ticket-driven, group: actions)

Reference: actions
Category: feature flags
Horizon: NearTerm
ROI: 3.8 (impact=6, confidence=0.68, effort=2)

Question:
Using the attached tickets as the primary context, define minimal compat-safe rollout plan and guardrails; include fallback behavior and monitoring gates.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/20-feature-flags-compat-rollout-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #20  (ticket-driven, group: actions)

Reference: actions
Category: feature flags
Horizon: NearTerm
ROI: 3.8 (impact=6, confidence=0.68, effort=2)

Question:
Using the attached tickets as the primary context, define minimal compat-safe rollout plan and guardrails; include fallback behavior and monitoring gates.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/20-feature-flags-compat-rollout-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #20  (ticket-driven, group: actions)

Reference: actions
Category: feature flags
Horizon: NearTerm
ROI: 3.8 (impact=6, confidence=0.68, effort=2)

Question:
Using the attached tickets as the primary context, define minimal compat-safe rollout plan and guardrails; include fallback behavior and monitoring gates.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/20-feature-flags-compat-rollout-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #20  (ticket-driven, group: actions)

Reference: actions
Category: feature flags
Horizon: NearTerm
ROI: 3.8 (impact=6, confidence=0.68, effort=2)

Question:
Using the attached tickets as the primary context, define minimal compat-safe rollout plan and guardrails; include fallback behavior and monitoring gates.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


```

## Coverage check
- contracts/interfaces: OK
- invariants: OK
- caching/state: OK
- background jobs: OK
- observability: OK
- permissions: OK
- migrations: OK
- UX flows: OK
- failure modes: OK
- feature flags: OK
