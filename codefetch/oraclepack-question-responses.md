<filetree>
Project Structure:
└── docs
    └── oracle-questions-2026-01-09-030000
        ├── actions
        │   ├── 01-contracts-interfaces-ticket-surface-direct-answer.md
        │   ├── 01-contracts-interfaces-ticket-surface-missing-evidence.md
        │   ├── 01-contracts-interfaces-ticket-surface-next-experiment.md
        │   └── 01-contracts-interfaces-ticket-surface-risks-unknowns.md
        ├── mcp
        │   ├── 01-contracts-interfaces-ticket-surface-direct-answer.md
        │   ├── 01-contracts-interfaces-ticket-surface-missing-evidence.md
        │   ├── 01-contracts-interfaces-ticket-surface-next-experiment.md
        │   ├── 01-contracts-interfaces-ticket-surface-risks-unknowns.md
        │   ├── 02-contracts-interfaces-integration-points-direct-answer.md
        │   ├── 02-contracts-interfaces-integration-points-missing-evidence.md
        │   ├── 02-contracts-interfaces-integration-points-next-experiment.md
        │   ├── 02-contracts-interfaces-integration-points-risks-unknowns.md
        │   ├── 03-invariants-invariant-map-direct-answer.md
        │   ├── 03-invariants-invariant-map-missing-evidence.md
        │   ├── 03-invariants-invariant-map-next-experiment.md
        │   ├── 03-invariants-invariant-map-risks-unknowns.md
        │   ├── 04-invariants-validation-boundaries-direct-answer.md
        │   ├── 04-invariants-validation-boundaries-missing-evidence.md
        │   ├── 04-invariants-validation-boundaries-next-experiment.md
        │   └── 04-invariants-validation-boundaries-risks-unknowns.md
        ├── packs
        │   ├── actions.md
        │   ├── mcp.md
        │   ├── misc.md
        │   ├── other.md
        │   ├── prd-tui.md
        │   └── raw-exports.md
        ├── _groups.json
        └── manifest.json

</filetree>

<source_code>
docs/oracle-questions-2026-01-09-030000/_groups.json
```
{
  "PRD-TUI": [
    ".tickets/PRD-TUI/Oraclepack TUI Integration.md",
    ".tickets/PRD-TUI/PRD-generator URL routing.md"
  ],
  "actions": [
    ".tickets/actions/Enable Action Packs Dispatch.md",
    ".tickets/actions/Improving Oraclepack Workflow.md",
    ".tickets/actions/Oraclepack Action Pack Integration.md",
    ".tickets/actions/Oraclepack Action Pack Issue.md",
    ".tickets/actions/Oraclepack Action Packs.md",
    ".tickets/actions/Oraclepack Compatibility Issues.md"
  ],
  "mcp": [
    ".tickets/mcp/Expose Oraclepack as MCP.md",
    ".tickets/mcp/MCP Server for Oraclepack.md",
    ".tickets/mcp/gaps-still-not-covered.md",
    ".tickets/mcp/gaps_part2-mcp-builder.md",
    ".tickets/mcp/oraclepack-MCP.md",
    ".tickets/mcp/oraclepack_mcp_server.md"
  ],
  "misc": [
    ".tickets/Formalize LLM Decision Points.md",
    ".tickets/Oraclepack CLI MCP Parity.md",
    ".tickets/Oraclepack File Storage.md",
    ".tickets/Oraclepack Parity Automation.md",
    ".tickets/Oraclepack Schema Approach.md",
    ".tickets/Oraclepack bash fix.md",
    ".tickets/Oraclepack output verification issues.md",
    ".tickets/Oraclepack-CLI-agents.md",
    ".tickets/Publish OraclePack MCP.md"
  ],
  "other": [
    ".tickets/other/Oraclepack Pipeline Improvements.md",
    ".tickets/other/Oraclepack Prompt Generator.md",
    ".tickets/other/Oraclepack Workflow Enhancement.md",
    ".tickets/other/Verbose Payload Rendering TUI.md"
  ],
  "raw-exports": [
    ".tickets/raw-exports/Output verification failure.md"
  ]
}
```

docs/oracle-questions-2026-01-09-030000/manifest.json
```
{
  "groups": [
    {
      "attached_paths": [
        ".tickets/PRD-TUI/Oraclepack TUI Integration.md",
        ".tickets/PRD-TUI/PRD-generator URL routing.md"
      ],
      "group": "PRD-TUI",
      "original_tickets": [
        ".tickets/PRD-TUI/Oraclepack TUI Integration.md",
        ".tickets/PRD-TUI/PRD-generator URL routing.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-09-030000/prd-tui",
      "pack_path": "docs/oracle-questions-2026-01-09-030000/packs/prd-tui.md",
      "part": 1,
      "slug": "prd-tui"
    },
    {
      "attached_paths": [
        ".tickets/actions/Enable Action Packs Dispatch.md",
        ".tickets/actions/Improving Oraclepack Workflow.md",
        ".tickets/actions/Oraclepack Action Pack Integration.md",
        ".tickets/actions/Oraclepack Action Pack Issue.md",
        ".tickets/actions/Oraclepack Action Packs.md",
        ".tickets/actions/Oraclepack Compatibility Issues.md"
      ],
      "group": "actions",
      "original_tickets": [
        ".tickets/actions/Enable Action Packs Dispatch.md",
        ".tickets/actions/Improving Oraclepack Workflow.md",
        ".tickets/actions/Oraclepack Action Pack Integration.md",
        ".tickets/actions/Oraclepack Action Pack Issue.md",
        ".tickets/actions/Oraclepack Action Packs.md",
        ".tickets/actions/Oraclepack Compatibility Issues.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-09-030000/actions",
      "pack_path": "docs/oracle-questions-2026-01-09-030000/packs/actions.md",
      "part": 1,
      "slug": "actions"
    },
    {
      "attached_paths": [
        ".tickets/mcp/Expose Oraclepack as MCP.md",
        ".tickets/mcp/MCP Server for Oraclepack.md",
        ".tickets/mcp/gaps-still-not-covered.md",
        ".tickets/mcp/gaps_part2-mcp-builder.md",
        ".tickets/mcp/oraclepack-MCP.md",
        ".tickets/mcp/oraclepack_mcp_server.md"
      ],
      "group": "mcp",
      "original_tickets": [
        ".tickets/mcp/Expose Oraclepack as MCP.md",
        ".tickets/mcp/MCP Server for Oraclepack.md",
        ".tickets/mcp/gaps-still-not-covered.md",
        ".tickets/mcp/gaps_part2-mcp-builder.md",
        ".tickets/mcp/oraclepack-MCP.md",
        ".tickets/mcp/oraclepack_mcp_server.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-09-030000/mcp",
      "pack_path": "docs/oracle-questions-2026-01-09-030000/packs/mcp.md",
      "part": 1,
      "slug": "mcp"
    },
    {
      "attached_paths": [
        ".tickets/Formalize LLM Decision Points.md",
        ".tickets/Oraclepack CLI MCP Parity.md",
        ".tickets/Oraclepack File Storage.md",
        ".tickets/Oraclepack Parity Automation.md",
        ".tickets/Oraclepack Schema Approach.md",
        ".tickets/Oraclepack bash fix.md",
        ".tickets/Oraclepack output verification issues.md",
        ".tickets/Oraclepack-CLI-agents.md",
        ".tickets/Publish OraclePack MCP.md"
      ],
      "group": "misc",
      "original_tickets": [
        ".tickets/Formalize LLM Decision Points.md",
        ".tickets/Oraclepack CLI MCP Parity.md",
        ".tickets/Oraclepack File Storage.md",
        ".tickets/Oraclepack Parity Automation.md",
        ".tickets/Oraclepack Schema Approach.md",
        ".tickets/Oraclepack bash fix.md",
        ".tickets/Oraclepack output verification issues.md",
        ".tickets/Oraclepack-CLI-agents.md",
        ".tickets/Publish OraclePack MCP.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-09-030000/misc",
      "pack_path": "docs/oracle-questions-2026-01-09-030000/packs/misc.md",
      "part": 1,
      "slug": "misc"
    },
    {
      "attached_paths": [
        ".tickets/other/Oraclepack Pipeline Improvements.md",
        ".tickets/other/Oraclepack Prompt Generator.md",
        ".tickets/other/Oraclepack Workflow Enhancement.md",
        ".tickets/other/Verbose Payload Rendering TUI.md"
      ],
      "group": "other",
      "original_tickets": [
        ".tickets/other/Oraclepack Pipeline Improvements.md",
        ".tickets/other/Oraclepack Prompt Generator.md",
        ".tickets/other/Oraclepack Workflow Enhancement.md",
        ".tickets/other/Verbose Payload Rendering TUI.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-09-030000/other",
      "pack_path": "docs/oracle-questions-2026-01-09-030000/packs/other.md",
      "part": 1,
      "slug": "other"
    },
    {
      "attached_paths": [
        ".tickets/raw-exports/Output verification failure.md"
      ],
      "group": "raw-exports",
      "original_tickets": [
        ".tickets/raw-exports/Output verification failure.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-09-030000/raw-exports",
      "pack_path": "docs/oracle-questions-2026-01-09-030000/packs/raw-exports.md",
      "part": 1,
      "slug": "raw-exports"
    }
  ]
}
```

docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-direct-answer.md
```
Direct answer

Document (CLI/TUI user-facing) the current Action Pack execution contract: each step runs as shell via bash -lc ..., and oraclepack’s special handling (override injection/validation) applies only to commands beginning with oracle (non-oracle tools run directly). Evidence: Oraclepack_Compatibility_Issues.md (“oraclepack executes each step as shell (bash -lc ...) and only applies oracle-specific behavior to oracle-prefixed commands… non-oracle tools (tm/task-master, codex, gemini) run directly”). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Extend the dispatcher’s command-detection surface beyond oracle-prefixed commands (currently anchored to ^(\\s*)(oracle)\\b) so steps containing tm/task-master, codex, and/or gemini are no longer excluded from the override/validation pipeline solely due to prefix mismatch; preserve existing oracle behavior. Evidence: Oraclepack_Compatibility_Issues.md (T2: “injects flags into commands that begin with oracle… regex anchored to ^(\\s*)(oracle)\\b… does not match tm, codex, gemini… Must preserve existing oracle command behavior”). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Extend/adjust TUI override validation so it doesn’t only run oracle --dry-run summary for detected oracle invocations and skip steps without oracle commands (i.e., make override validation aware of non-oracle tool steps, per the dispatcher changes). Evidence: Oraclepack_Compatibility_Issues.md (“TUI override validation… runs oracle --dry-run summary… steps without oracle invocations are skipped”; T2 scope calls for updating override validation beyond oracle-only targeting). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Replace placeholder steps in ticket-action-pack.md (explicitly 09–13 and 16) with headless gemini + non-interactive codex exec automation, standardizing new output artifacts under .oraclepack/ticketify/ (e.g., next.json, codex-implement.md, codex-verify.md and/or gemini-review.json, PR.md). Back-compat constraint: the artifact paths are part of the step contract (don’t rename/move without versioning). Evidence: Oraclepack_Compatibility_Issues.md (T3 “Replace… placeholder steps… Step 09 writes .oraclepack/ticketify/next.json… Step 10 writes codex-implement.md… Step 11 writes codex-verify.md and/or gemini-review.json… Step 16 writes PR.md; “Must write the artifacts to the paths specified”). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Make tool-availability guards and “skip” behavior part of the public operational contract for these new non-oracle steps (e.g., command -v ... checks), and explicitly treat “interactive CLI blocks” as a first-class hazard; back-compat constraint: default runs should not start failing/hanging just because optional tools aren’t installed or are interactive. Evidence: Oraclepack_Compatibility_Issues.md (“If CLI isn’t installed/on PATH → step fails; if interactive → blocks”; T3 requires “command-availability guards and ‘skip’ behavior… to avoid hard failures when tools are missing”). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Add an “agent-mode” option to taskify-generated Action Packs (suggested mode=codex / mode=gemini) that swaps the existing autopilot entrypoint step with an agent-implementation step, without changing the schema/step-count contract; back-compat constraint: keep the “20-step contract intact” and keep existing default behavior unchanged unless mode is explicitly selected. Evidence: Oraclepack_Compatibility_Issues.md (T4 “Add ‘agent-mode’… mode=codex / mode=gemini… keep the 20-step contract intact by swapping… step slot”). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Preserve the Action Pack ingestibility contract as a hard interface boundary (single ```bash fence, # NN) step structure / fixed-step schema), even as new automation is inserted into placeholder slots. Evidence: Oraclepack_Compatibility_Issues.md (“Action Pack is ‘oraclepack-ingestible’ (single bash fence, # NN) steps)”; “keep the 20-step contract intact”).

Oraclepack_Compatibility_Issues
```

docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-missing-evidence.md
```
Missing evidence

* `.tickets/actions/**/*.md` (including any actions “Parent Ticket Summary.md” / index files)
* The referenced pack(s): `**/ticket-action-pack.md`, `**/oracle_pack_and_taskify-skills.md`, `**/oraclepack-tui.md`
* The generator/template source for `ticket-action-pack.md`: `**/*ticket-action-pack*.md`, `**/*ticketify*pack*.md`, `**/templates/**`, `**/assets/**`
* Oraclepack command detection + override injection + validation implementation: `**/cmd/oraclepack/**`, `**/internal/**`, `**/*dispatch*.*`, `**/*override*.*`, `**/*validate*.*`
* Run state/report writers + schema definitions for the named artifacts: `**/*state*.go`, `**/*report*.go`, `**/*.state.json*`, `**/*.report.json*`, `**/schema/**`
* TUI surfaces for override/validation UX changes: `**/*tui*/**`, `**/*ui*/**`, `**/*bubbletea*/**`
```

docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-next-experiment.md
```
Next smallest concrete experiment

* In the oraclepack repo root, locate the current oracle-only detection + validation callsites to scope the minimal dispatcher/validation expansion: `rg -n --hidden --glob '!.git/' -F -e '^(\\s*)(oracle)\\b' -e 'oracle --dry-run summary' -e 'override validation' -e 'injects flags into commands that begin with' .` 
```

docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-risks-unknowns.md
```
Risks/unknowns

The tickets do not specify which overrides/flags should apply to each non-oracle tool (tm/task-master, codex, gemini) or what “validation” means for them; implementing anything beyond broadened detection risks incorrect/partial behavior. 

Oraclepack_Compatibility_Issues

“Headless/non-interactive” invocation details for codex/gemini are not specified; interactivity is explicitly called out as a risk (runs can block), and missing binaries on PATH can hard-fail steps without guards/skip semantics. 

Oraclepack_Compatibility_Issues

The intended default for verification is underspecified: Step 11 could be codex exec, Gemini diff review, or both; similarly, whether Steps 12–13 should be modified is not stated. 

Oraclepack_Compatibility_Issues

The user-facing selection surface for taskify “agent-mode” is not specified (CLI flag vs TUI option vs config), and the exact step slot to swap is not specified. 

Oraclepack_Compatibility_Issues

Where the “clarify current execution semantics/failure modes” documentation should live (README vs oraclepack-tui.md vs TUI help text) is not provided, which affects discoverability and backward-compat messaging. 

Oraclepack_Compatibility_Issues

The tickets do not provide exact code locations for the current detection/override-validation machinery (they describe behavior like an oracle-anchored matcher and oracle-only validation, but not file paths), so the scope of required CLI/TUI/internal refactors is uncertain. 

Oraclepack_Compatibility_Issues

Inference: Several on-disk artifact paths are named as expected outputs (e.g., .oraclepack/ticketify/*, ticket-action-pack.state.json, ticket-action-pack.report.json), but the tickets don’t define schemas/versions for these artifacts, increasing drift risk for any downstream tooling that consumes them. 

Oraclepack_Compatibility_Issues

Reproduction steps are not provided across the split tickets, making regression testing of new/changed public behavior (CLI/TUI validation, pack execution outcomes) harder to anchor. 

Oraclepack_Compatibility_Issues
```

docs/oracle-questions-2026-01-09-030000/packs/actions.md
```
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
[TRUNCATED]
```

docs/oracle-questions-2026-01-09-030000/packs/mcp.md
```
# Oracle Pack — Unknown (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: Unknown
- out_dir: docs/oracle-questions-2026-01-09-030000/mcp
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md
- ticket_max_files: 6
- group_name: mcp
- group_slug: mcp
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-09-030000/mcp/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-09-030000/mcp"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=mcp

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md".strip()
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
  echo "WARNING: no tickets resolved for group 'mcp'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/01-contracts-interfaces-ticket-surface-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: mcp)

Reference: mcp
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/01-contracts-interfaces-ticket-surface-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: mcp)

Reference: mcp
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/01-contracts-interfaces-ticket-surface-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: mcp)

Reference: mcp
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/01-contracts-interfaces-ticket-surface-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: mcp)

Reference: mcp
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


# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=mcp

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md".strip()
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
  echo "WARNING: no tickets resolved for group 'mcp'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/02-contracts-interfaces-integration-points-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
[TRUNCATED]
```

docs/oracle-questions-2026-01-09-030000/packs/misc.md
```
# Oracle Pack — Unknown (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: Unknown
- out_dir: docs/oracle-questions-2026-01-09-030000/misc
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/Formalize LLM Decision Points.md,.tickets/Oraclepack CLI MCP Parity.md,.tickets/Oraclepack File Storage.md,.tickets/Oraclepack Parity Automation.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Oraclepack output verification issues.md,.tickets/Oraclepack-CLI-agents.md,.tickets/Publish OraclePack MCP.md
- ticket_max_files: 9
- group_name: misc
- group_slug: misc
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-09-030000/misc/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-09-030000/misc"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=misc

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/Formalize LLM Decision Points.md,.tickets/Oraclepack CLI MCP Parity.md,.tickets/Oraclepack File Storage.md,.tickets/Oraclepack Parity Automation.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Oraclepack output verification issues.md,.tickets/Oraclepack-CLI-agents.md,.tickets/Publish OraclePack MCP.md".strip()
MAX = int("9")


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
  echo "WARNING: no tickets resolved for group 'misc'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/01-contracts-interfaces-ticket-surface-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: misc)

Reference: misc
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/01-contracts-interfaces-ticket-surface-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: misc)

Reference: misc
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/01-contracts-interfaces-ticket-surface-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: misc)

Reference: misc
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/01-contracts-interfaces-ticket-surface-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: misc)

Reference: misc
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


# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=misc

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/Formalize LLM Decision Points.md,.tickets/Oraclepack CLI MCP Parity.md,.tickets/Oraclepack File Storage.md,.tickets/Oraclepack Parity Automation.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Oraclepack output verification issues.md,.tickets/Oraclepack-CLI-agents.md,.tickets/Publish OraclePack MCP.md".strip()
MAX = int("9")


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
  echo "WARNING: no tickets resolved for group 'misc'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
[TRUNCATED]
```

docs/oracle-questions-2026-01-09-030000/packs/other.md
```
# Oracle Pack — Unknown (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: Unknown
- out_dir: docs/oracle-questions-2026-01-09-030000/other
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md
- ticket_max_files: 4
- group_name: other
- group_slug: other
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-09-030000/other/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-09-030000/other"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=other

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md".strip()
MAX = int("4")


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
  echo "WARNING: no tickets resolved for group 'other'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/01-contracts-interfaces-ticket-surface-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: other)

Reference: other
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/01-contracts-interfaces-ticket-surface-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: other)

Reference: other
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/01-contracts-interfaces-ticket-surface-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: other)

Reference: other
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/01-contracts-interfaces-ticket-surface-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: other)

Reference: other
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


# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=other

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md".strip()
MAX = int("4")


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
  echo "WARNING: no tickets resolved for group 'other'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/02-contracts-interfaces-integration-points-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: other)

Reference: other
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

[TRUNCATED]
```

docs/oracle-questions-2026-01-09-030000/packs/prd-tui.md
```
# Oracle Pack — Unknown (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: Unknown
- out_dir: docs/oracle-questions-2026-01-09-030000/prd-tui
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md
- ticket_max_files: 2
- group_name: PRD-TUI
- group_slug: prd-tui
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-09-030000/prd-tui"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=prd-tui

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md".strip()
MAX = int("2")


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
  echo "WARNING: no tickets resolved for group 'PRD-TUI'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/01-contracts-interfaces-ticket-surface-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/01-contracts-interfaces-ticket-surface-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/01-contracts-interfaces-ticket-surface-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/01-contracts-interfaces-ticket-surface-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
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


# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=prd-tui

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md".strip()
MAX = int("2")


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
  echo "WARNING: no tickets resolved for group 'PRD-TUI'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/02-contracts-interfaces-integration-points-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

[TRUNCATED]
```

docs/oracle-questions-2026-01-09-030000/packs/raw-exports.md
```
# Oracle Pack — Unknown (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: Unknown
- out_dir: docs/oracle-questions-2026-01-09-030000/raw-exports
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/raw-exports/Output verification failure.md
- ticket_max_files: 1
- group_name: raw-exports
- group_slug: raw-exports
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-09-030000/raw-exports"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=raw-exports

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/raw-exports/Output verification failure.md".strip()
MAX = int("1")


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
  echo "WARNING: no tickets resolved for group 'raw-exports'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/01-contracts-interfaces-ticket-surface-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: raw-exports)

Reference: raw-exports
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/01-contracts-interfaces-ticket-surface-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: raw-exports)

Reference: raw-exports
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/01-contracts-interfaces-ticket-surface-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: raw-exports)

Reference: raw-exports
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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/01-contracts-interfaces-ticket-surface-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: raw-exports)

Reference: raw-exports
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


# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=raw-exports

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/raw-exports/Output verification failure.md".strip()
MAX = int("1")


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
  echo "WARNING: no tickets resolved for group 'raw-exports'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/02-contracts-interfaces-integration-points-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: raw-exports)

Reference: raw-exports
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
[TRUNCATED]
```

docs/oracle-questions-2026-01-09-030000/mcp/01-contracts-interfaces-ticket-surface-direct-answer.md
```
Direct answer

MCP_Server_for_Oraclepack

oraclepack_mcp_server
```

docs/oracle-questions-2026-01-09-030000/mcp/01-contracts-interfaces-ticket-surface-missing-evidence.md
```
## Risks/unknowns

* Streamable-HTTP transport correctness is uncertain across the “proposal” vs actual implementation: tickets call out that `--transport streamable-http` may be wired/mapped incorrectly (and that SSE is superseded), which can break non-stdio deployments. 
* Production streamable-HTTP posture is under-specified: no clear auth/TLS/DNS-rebinding mitigation story in the tickets’ “gaps” section, which is a deployment risk if HTTP mode is used beyond purely local/dev. 
* Dependency safety/pinning ambiguity: tickets flag an “unpinned” MCP Python SDK dependency in the presence of a DNS-rebinding advisory, but the current server’s `requirements.txt` shown in the code excerpt is only `mcp>=1.0.0` (still not pinned to an exact safe patch), so the intended minimum safe version is unclear.
* Remote execution exposure risk: “run” tools become available via a single boolean env gate (`ORACLEPACK_ENABLE_EXEC`), and the run tool is explicitly annotated as open-world/destructive, with no per-tool authorization/allowlisting described for HTTP mode.
* Filesystem escape hardening is a known concern in tickets (symlink traversal and lack of tests called out). While the server enforces `ORACLEPACK_ALLOWED_ROOTS`, the test coverage/guarantees for symlink + path traversal behavior remain an unknown from the tickets’ perspective.
* Operational fragility from external process dependency: MCP tools rely on an external `oraclepack` binary + correct working directory configuration, which can vary across MCP hosts and lead to confusing failures.
* Stage-2 ambiguity failure modes: tickets explicitly call out risk when multiple `NN-*.md` files match the same prefix; additionally, Stage-2 “auto” selection uses deterministic-but-heuristic lexicographic “newest” selection, which may select the wrong target in real repos.
* Artifacts summary semantics can be surprising: the implementation treats paths outside `ORACLEPACK_ALLOWED_ROOTS` as “missing: not allowed” rather than an error, which can mask config mistakes and complicate debugging.
* Distribution/public config surface is in flux: tickets propose multiple packaging/config approaches (PATH executable, `uvx`, `.mcpb`, registry publishing). Each path changes the client-facing `command/args` contract and can break existing MCP configs if not versioned/migrated. 
```

docs/oracle-questions-2026-01-09-030000/mcp/01-contracts-interfaces-ticket-surface-next-experiment.md
```
Direct answer

* Add a new public server entrypoint `python -m oraclepack_mcp_server --transport {stdio|streamable-http}` (default `stdio`) to expose oraclepack via MCP.  
* Expose a new MCP tool surface (public API contract) with these tool names: `oraclepack_validate_pack`, `oraclepack_list_steps`, `oraclepack_run_pack`, `oraclepack_read_file`, `oraclepack_taskify_detect_stage2`, `oraclepack_taskify_validate_stage2`, `oraclepack_taskify_validate_action_pack`, `oraclepack_taskify_artifacts_summary`, `oraclepack_taskify_run_action_pack`.  
* Introduce env-driven server configuration as a public contract: `ORACLEPACK_ALLOWED_ROOTS`, `ORACLEPACK_BIN`, `ORACLEPACK_WORKDIR`, `ORACLEPACK_ENABLE_EXEC`, `ORACLEPACK_CHARACTER_LIMIT`, `ORACLEPACK_MAX_READ_BYTES`; backwards-compat constraint is “secure by default” (exec disabled unless explicitly enabled; filesystem reads constrained to allowlisted roots).  
* Gate “run” tools behind `ORACLEPACK_ENABLE_EXEC=1` (deny-by-default execution); backwards-compat constraint is that enabling exec must remain an explicit opt-in to avoid silently expanding agent capabilities.  
* Add deterministic Stage-2 helper contracts: detect Stage-2 outputs from either a directory containing exactly one `01-*.md`…`20-*.md` per prefix or from a “single-pack form” file; enforce out-dir resolution rules for packs under `docs/oracle-questions-YYYY-MM-DD/...` (use sibling `oracle-out`, else `repo_root/oracle-out`).  
* Add Stage-3 helper contracts: validate Action Pack structure pre-execution (exactly one ```bash fence, step headers like `# 01) ...` sequential) and summarize canonical artifacts (e.g., `_actions.json`, `_actions.md`, PRD path, `tm-complexity.json`, pipelines doc). Backwards-compat constraint: MCP validation must not be looser than the existing validator expectations, or “passes in MCP, fails in real usage” drift occurs.   
* Expand/define the public input schema for `oraclepack_run_pack` beyond “run all” to include non-interactive flags and runtime controls (e.g., `no_tui`, `yes`, `run_all`, `resume`, `stop_on_fail`, ROI threshold/mode, `out_dir`, `timeout_s`), plus response-format selection (`markdown` vs `json`). Backwards-compat constraint: keep explicit, typed fields rather than an “extra args” escape hatch to avoid unstable tool contracts.  
* Attach MCP tool annotations as part of the contract/UX surface: validate/list/read tools should be marked read-only; run tools should be marked destructive/open-world (so clients can require approvals).   
* Transport semantics become part of the public surface: `streamable-http` must be wired as Streamable HTTP (not SSE-as-alias), and HTTP deployments require hardening (Origin validation/auth/local binding) to preserve the intended safety posture; backwards-compat constraint is that changing the transport string/behavior breaks client configs, so the chosen values must remain stable once shipped.   
```

docs/oracle-questions-2026-01-09-030000/mcp/01-contracts-interfaces-ticket-surface-risks-unknowns.md
```
Risks/unknownoraclepack-llms-fulloraclepack-llms-fulloraclepack-llms-full
```

docs/oracle-questions-2026-01-09-030000/mcp/02-contracts-interfaces-integration-points-direct-answer.md
```
Direct answer

oraclepack_mcp_server
```

docs/oracle-questions-2026-01-09-030000/mcp/02-contracts-interfaces-integration-points-missing-evidence.md
```
## Risks/unknowns

* Streamable HTTP transport may be mis-wired or inconsistently specified (proposal conflates SSE vs Streamable HTTP), risking client incompatibility for non-stdio deployments.
* Remote/HTTP deployment security posture is under-specified (auth/TLS/origin validation/DNS-rebinding mitigations), which is a material risk if integrations expect a reachable HTTP endpoint rather than local stdio.
* MCP Python SDK dependency pinning/patch-level requirements are unclear; tickets explicitly call out an advisory and “unpinned dependency” risk, while the reference requirements are only lower-bounded (e.g., `mcp>=...`).
* External integration configs (client-specific onboarding: Codex/Inspector/other MCP clients) are implied but not concretely locked down; changes to `command/args/cwd/env` conventions can break existing MCP configs.
* Execution safety relies primarily on a single env gate (`ORACLEPACK_ENABLE_EXEC`) with open-world/destructive semantics for run tools; least-privilege scoping (per-tool allowlists, auth flows for HTTP) is not defined.
* Filesystem boundary enforcement may be incomplete against symlink traversal (normalize/resolve + open patterns called out as risky); absence of explicit realpath-based checks/tests creates an escape/exfil risk.
* Failure/timeout behavior is only partially contract-defined: tools return bounded stdout/stderr with hard timeouts + truncation, but lack streaming progress, cancellation hooks, and structured “in-progress state” retrieval; long runs may feel wedged to clients.
* Tool schema parity with the underlying `oraclepack` CLI is uncertain (tickets call out missing run-time controls and “wizard parity” gaps), increasing the chance of ad-hoc “extra args” escape hatches or repeated schema churn.
* Stage-2 auto-discovery + validation determinism is a known sharp edge (ambiguous `NN-*.md` matches, differing search order expectations across tools/skills); rollout could break workflows that previously “worked by accident.”
* Stage-3 Action Pack validation parity is explicitly called out as weaker than existing shell validation scripts; mismatched validators create split-brain behavior (“valid” in MCP, “invalid” in real execution).
```

docs/oracle-questions-2026-01-09-030000/mcp/02-contracts-interfaces-integration-points-next-experiment.md
```
Next smallest

oraclepack_mcp_server
```

docs/oracle-questions-2026-01-09-030000/mcp/02-contracts-interfaces-integration-points-risks-unknowns.md
```
## Missing evidence

* None.
```

docs/oracle-questions-2026-01-09-030000/mcp/03-invariants-invariant-map-direct-answer.md
```
Next smallest concrete experiment

* Execute a timeout smoke-test against the proposed `run_cmd(...)` subprocess wrapper (expect `exit_code=124` and a `Timed out after 1s:` message in `stderr`):

  ````bash
  python - <<'PY'
  import asyncio
  from pathlib import Path
  from oraclepack_mcp_server.oraclepack_cli import run_cmd

  async def main():
      r = await run_cmd(["bash", "-lc", "sleep 2"], cwd=Path("."), timeout_s=1, env={}, character_limit=25000)
      print({"ok": r.ok, "exit_code": r.exit_code, "stderr": r.stderr})

  asyncio.run(main())
  PY
  ``` :contentReference[oaicite:0]{index=0}
  ````
```

docs/oracle-questions-2026-01-09-030000/mcp/03-invariants-invariant-map-missing-evidence.md
```
## Risks/unknowns

* “Action Pack” invariants are not fully specified (formal grammar for step headers, exact fence rules beyond examples), so validators can drift between MCP and the canonical `validate-action-pack.sh` guardrails.
* Stage-2 discovery/validation invariants are partly underspecified for “auto” mode (ordered search locations, tie-break rules across multiple date-stamped runs); inconsistent selection can break determinism and downstream Stage-3 assumptions.
* Filesystem boundary invariant (“must stay under allowed roots”) is vulnerable to symlink traversal unless realpath/TOCTOU-safe checks are enforced and tested; tickets explicitly call out symlink escape risk.
* Execution-safety invariants are incomplete: gating is a single env boolean, without least-privilege scoping (per-tool allowlists/authorization) and without concurrency/cancellation/streaming, which raises “wedged run” and unsafe execution risks.
* Output/observability invariants are weakly enforced: MCP tools primarily return truncated stdout/stderr and do not provide structured access to oraclepack run state/report artifacts, limiting reliable “resume/re-run failed only” workflows.
* “Write locations” invariants are unclear: reads are allowlisted, but the contract for where packs are permitted to write (e.g., validating/normalizing `out_dir` for all run paths) is called out as still needed for safe operation.
* Artifact inventory invariants for Stage-3 are only example-based (e.g., `_actions.json`, PRD path, `tm-complexity.json`, pipelines doc); the exact required set, naming, and stable locations are not fully defined.
* Transport invariants for “streamable-http” correctness and security are not locked down (SSE vs streamable-http mapping confusion; origin/auth hardening requirements), risking client breakage and insecure deployments.
* CLI surface parity invariants are uncertain: if MCP tool schemas expose only a subset of oraclepack runtime controls, users will rely on undefined behavior or demand ad-hoc “extra args” escape hatches, causing ongoing contract churn.
```

docs/oracle-questions-2026-01-09-030000/mcp/03-invariants-invariant-map-next-experiment.md
```
Next smallest

oraclepack_mcp_server

gaps-still-not-covered
```

docs/oracle-questions-2026-01-09-030000/mcp/03-invariants-invariant-map-risks-unknowns.md
```
## Missing evidence

* None.
```

docs/oracle-questions-2026-01-09-030000/mcp/04-invariants-validation-boundaries-direct-answer.md
```
Direct answeroraclepack-llms-fulloraclepack-llms-fulloraclepack-llms-fulloraclepack-llms-full
```

docs/oracle-questions-2026-01-09-030000/mcp/04-invariants-validation-boundaries-missing-evidence.md
```
## Risks/unknowns

* Validation parity drift risk: MCP-side “validate_action_pack” is called out as weaker than the repo’s `validate-action-pack.sh`, so “passes in MCP, fails in real usage” remains likely unless boundaries are unified.
* Stage-2 discovery/validation boundary ambiguity: “auto” discovery order and candidate selection rules are incomplete and can yield non-deterministic or surprising selections unless the detector validates candidates before returning them.
* Filesystem security validation boundary is incomplete against symlink traversal: current allowlisted-root checks rely on `resolve()/relative_to()` patterns that tickets explicitly flag as likely bypassable without realpath/TOCTOU-safe enforcement + tests.
* Pack execution gating boundary is coarse: a single env flag enables open-world/destructive “run” tools without least-privilege scoping (per-tool allowlists/authorization), especially problematic for HTTP transport.
* Transport validation boundary is unclear: tickets note `--transport streamable-http` wiring mistakes (SSE conflation) and missing production hardening requirements, which can invalidate any rollout assumptions for HTTP mode.
* Dependency/version boundary uncertainty: tickets explicitly flag an unpinned MCP Python SDK dependency in the presence of a DNS-rebinding advisory, but the reference requirements only specify broad minimums, leaving “known-safe” versions undefined.
* Observability/state validation boundary gap: MCP tools return truncated stdout/stderr but do not provide structured read APIs for oraclepack state/report artifacts, weakening deterministic “validate → run → inspect → recover” workflows.
* Ticket parsing / pack generation boundary is underspecified in the MCP surface: capability gaps explicitly note no “write/update pack” support, so it’s unclear where/when to validate generated packs derived from tickets (pre-write vs post-write) within this MCP design.
* Artifact contract validation boundary is fuzzy: “artifact summary” lists example files/paths but not a strict schema or required set, so validators and downstream tooling can diverge on what constitutes a “complete” Stage-3 output.
```

docs/oracle-questions-2026-01-09-030000/mcp/04-invariants-validation-boundaries-next-experiment.md
```
Direct answer

* Ticket parsing boundary: before any “ticketify” logic, validate `.tickets/` exists and ticket files are readable; then require the ticket index artifact to be produced (e.g., `.oraclepack/ticketify/_tickets_index.json`) before continuing downstream. 
* Pack generation boundary: generation must be “fail-fast” and deterministic—render packs without unresolved placeholders and ensure outputs are runner-ingestible (e.g., exact ` ```bash … ``` ` shape + “exactly 20 steps” contract), then immediately validate generated packs using the existing validator. 
* Pack validation boundary (core): `oraclepack validate` must enforce strict code-fence invariants—reject packs with multiple/non-bash fences or any extra fences (explicitly called out as needed to prevent “shape drift”). 
* Pack lint boundary (pre-execution): add/maintain a “Stage-1 lint” that rejects packs early unless they have exactly one ` ```bash ` fence, exactly 20 steps, and step IDs are strictly `01..20` sequential (plus optional header-token enforcement if enabled). 
* Stage-2 output boundary: validate the Stage-2 output directory contract as “exactly one file per prefix `01-*.md` … `20-*.md`” (missing/duplicate prefixes are hard failures); MCP already models this as `oraclepack_taskify_validate_stage2` backed by `validate_stage2_dir`.
* Stage-3 action pack boundary: validate the Action Pack markdown as “exactly one ` ```bash ` fence, no other fences, and sequential `# 01)` headers”; refuse execution if invalid (MCP enforces both lint + refusal-to-run on failure).
* Execution boundary (MCP): require explicit opt-in to execution (`ORACLEPACK_ENABLE_EXEC=1`) before any `run` tools are usable, so validation/inspection can remain safe-by-default.
* Minimal validation plan: implement CI gating for the “manifest-first workflow” as (1) validate manifest, (2) render Markdown deterministically, (3) run `oraclepack validate` on rendered packs, (4) optionally run oracle dry-run; additionally enforce filesystem boundaries in MCP by resolving all user-supplied paths under `ORACLEPACK_ALLOWED_ROOTS`.
```

docs/oracle-questions-2026-01-09-030000/mcp/04-invariants-validation-boundaries-risks-unknowns.md
```
## Missing evidence

* None.
```

</source_code>