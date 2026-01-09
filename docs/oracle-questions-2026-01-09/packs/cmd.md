# Oracle Pack — oraclepack (Grouped Codebase Stage 1 — Direct Attach)

## Parsed args
- codebase_name: oraclepack
- out_dir: docs/oracle-questions-2026-01-09
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- code_root: .
- code_glob: **/*
- code_paths: 
- code_max_files: 600
- group_name: cmd
- group_slug: cmd
- group_mode: subdir+infer
- group_min_score: 0.1
- group_max_files: 200
- group_max_chars: 200000
- ignore_dirs: .archived,.mypy_cache,.cursor,.ruler,.rules,.taskmaster,codefetch,docs,skills,node_modules
- include_exts: .go,.md,.sh,.py,.json
- exclude_glob: node_modules/**,docs/oracle-questions-*/**,oraclepack.exe,**/*.state.json,**/*.report.json
- mode: codebase-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-09/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach code files directly (no bundle dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-09"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/01-contracts-interfaces-public-surface-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (codebase-driven, group: cmd)

Reference: cmd
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached code files as primary evidence, map the public surface area (CLI/TUI/API/interfaces/contracts). Call out backwards-compat constraints.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/01-contracts-interfaces-public-surface-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (codebase-driven, group: cmd)

Reference: cmd
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached code files as primary evidence, map the public surface area (CLI/TUI/API/interfaces/contracts). Call out backwards-compat constraints.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/01-contracts-interfaces-public-surface-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (codebase-driven, group: cmd)

Reference: cmd
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached code files as primary evidence, map the public surface area (CLI/TUI/API/interfaces/contracts). Call out backwards-compat constraints.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/01-contracts-interfaces-public-surface-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (codebase-driven, group: cmd)

Reference: cmd
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached code files as primary evidence, map the public surface area (CLI/TUI/API/interfaces/contracts). Call out backwards-compat constraints.

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


# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/02-contracts-interfaces-integrations-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (codebase-driven, group: cmd)

Reference: cmd
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached code files as primary evidence, identify external integrations implied by this area; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/02-contracts-interfaces-integrations-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (codebase-driven, group: cmd)

Reference: cmd
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached code files as primary evidence, identify external integrations implied by this area; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/02-contracts-interfaces-integrations-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (codebase-driven, group: cmd)

Reference: cmd
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached code files as primary evidence, identify external integrations implied by this area; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/02-contracts-interfaces-integrations-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (codebase-driven, group: cmd)

Reference: cmd
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached code files as primary evidence, identify external integrations implied by this area; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

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


# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/03-invariants-invariant-map-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (codebase-driven, group: cmd)

Reference: cmd
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached code files, map invariants and critical assumptions (data shape, ordering, idempotency, contracts). Identify the weakest or least-tested invariant.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/03-invariants-invariant-map-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (codebase-driven, group: cmd)

Reference: cmd
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached code files, map invariants and critical assumptions (data shape, ordering, idempotency, contracts). Identify the weakest or least-tested invariant.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/03-invariants-invariant-map-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (codebase-driven, group: cmd)

Reference: cmd
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached code files, map invariants and critical assumptions (data shape, ordering, idempotency, contracts). Identify the weakest or least-tested invariant.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/03-invariants-invariant-map-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (codebase-driven, group: cmd)

Reference: cmd
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached code files, map invariants and critical assumptions (data shape, ordering, idempotency, contracts). Identify the weakest or least-tested invariant.

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


# 04) ROI=4.7 impact=6 confidence=0.76 effort=1 horizon=NearTerm category=invariants reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/04-caching-state-reads-writes-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (codebase-driven, group: cmd)

Reference: cmd
Category: invariants
Horizon: NearTerm
ROI: 4.7 (impact=6, confidence=0.76, effort=1)

Question:
Using the attached code files, identify stateful reads/writes and any caches (in-memory, disk, external). Note invalidation boundaries and any silent staleness risks.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/04-caching-state-reads-writes-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (codebase-driven, group: cmd)

Reference: cmd
Category: invariants
Horizon: NearTerm
ROI: 4.7 (impact=6, confidence=0.76, effort=1)

Question:
Using the attached code files, identify stateful reads/writes and any caches (in-memory, disk, external). Note invalidation boundaries and any silent staleness risks.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/04-caching-state-reads-writes-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (codebase-driven, group: cmd)

Reference: cmd
Category: invariants
Horizon: NearTerm
ROI: 4.7 (impact=6, confidence=0.76, effort=1)

Question:
Using the attached code files, identify stateful reads/writes and any caches (in-memory, disk, external). Note invalidation boundaries and any silent staleness risks.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/04-caching-state-reads-writes-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (codebase-driven, group: cmd)

Reference: cmd
Category: invariants
Horizon: NearTerm
ROI: 4.7 (impact=6, confidence=0.76, effort=1)

Question:
Using the attached code files, identify stateful reads/writes and any caches (in-memory, disk, external). Note invalidation boundaries and any silent staleness risks.

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


# 05) ROI=4.4 impact=6 confidence=0.70 effort=1 horizon=NearTerm category=caching/state reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/05-background-jobs-queues-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #05  (codebase-driven, group: cmd)

Reference: cmd
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.70, effort=1)

Question:
Using the attached code files, list any background jobs/queues/cron tasks. Note retries, idempotency, and failure modes.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/05-background-jobs-queues-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #05  (codebase-driven, group: cmd)

Reference: cmd
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.70, effort=1)

Question:
Using the attached code files, list any background jobs/queues/cron tasks. Note retries, idempotency, and failure modes.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/05-background-jobs-queues-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #05  (codebase-driven, group: cmd)

Reference: cmd
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.70, effort=1)

Question:
Using the attached code files, list any background jobs/queues/cron tasks. Note retries, idempotency, and failure modes.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/05-background-jobs-queues-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #05  (codebase-driven, group: cmd)

Reference: cmd
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.70, effort=1)

Question:
Using the attached code files, list any background jobs/queues/cron tasks. Note retries, idempotency, and failure modes.

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


# 06) ROI=4.5 impact=6 confidence=0.75 effort=1 horizon=Immediate category=caching/state reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/06-observability-logging-metrics-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #06  (codebase-driven, group: cmd)

Reference: cmd
Category: caching/state
Horizon: Immediate
ROI: 4.5 (impact=6, confidence=0.75, effort=1)

Question:
Using the attached code files, identify logging/metrics/tracing in this area. Call out missing signals for debugging incidents.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/06-observability-logging-metrics-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #06  (codebase-driven, group: cmd)

Reference: cmd
Category: caching/state
Horizon: Immediate
ROI: 4.5 (impact=6, confidence=0.75, effort=1)

Question:
Using the attached code files, identify logging/metrics/tracing in this area. Call out missing signals for debugging incidents.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/06-observability-logging-metrics-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #06  (codebase-driven, group: cmd)

Reference: cmd
Category: caching/state
Horizon: Immediate
ROI: 4.5 (impact=6, confidence=0.75, effort=1)

Question:
Using the attached code files, identify logging/metrics/tracing in this area. Call out missing signals for debugging incidents.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/06-observability-logging-metrics-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #06  (codebase-driven, group: cmd)

Reference: cmd
Category: caching/state
Horizon: Immediate
ROI: 4.5 (impact=6, confidence=0.75, effort=1)

Question:
Using the attached code files, identify logging/metrics/tracing in this area. Call out missing signals for debugging incidents.

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


# 07) ROI=4.3 impact=6 confidence=0.68 effort=1 horizon=NearTerm category=background jobs reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/07-permissions-authz-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #07  (codebase-driven, group: cmd)

Reference: cmd
Category: background jobs
Horizon: NearTerm
ROI: 4.3 (impact=6, confidence=0.68, effort=1)

Question:
Using the attached code files, identify authorization and permission checks. Note any missing checks or implicit trust boundaries.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/07-permissions-authz-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #07  (codebase-driven, group: cmd)

Reference: cmd
Category: background jobs
Horizon: NearTerm
ROI: 4.3 (impact=6, confidence=0.68, effort=1)

Question:
Using the attached code files, identify authorization and permission checks. Note any missing checks or implicit trust boundaries.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/07-permissions-authz-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #07  (codebase-driven, group: cmd)

Reference: cmd
Category: background jobs
Horizon: NearTerm
ROI: 4.3 (impact=6, confidence=0.68, effort=1)

Question:
Using the attached code files, identify authorization and permission checks. Note any missing checks or implicit trust boundaries.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/07-permissions-authz-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #07  (codebase-driven, group: cmd)

Reference: cmd
Category: background jobs
Horizon: NearTerm
ROI: 4.3 (impact=6, confidence=0.68, effort=1)

Question:
Using the attached code files, identify authorization and permission checks. Note any missing checks or implicit trust boundaries.

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


# 08) ROI=4.2 impact=6 confidence=0.66 effort=1 horizon=NearTerm category=background jobs reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/08-migrations-backfills-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #08  (codebase-driven, group: cmd)

Reference: cmd
Category: background jobs
Horizon: NearTerm
ROI: 4.2 (impact=6, confidence=0.66, effort=1)

Question:
Using the attached code files, identify migrations/backfills/data-shape changes implied in this area. Note rollout risks.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/08-migrations-backfills-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #08  (codebase-driven, group: cmd)

Reference: cmd
Category: background jobs
Horizon: NearTerm
ROI: 4.2 (impact=6, confidence=0.66, effort=1)

Question:
Using the attached code files, identify migrations/backfills/data-shape changes implied in this area. Note rollout risks.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/08-migrations-backfills-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #08  (codebase-driven, group: cmd)

Reference: cmd
Category: background jobs
Horizon: NearTerm
ROI: 4.2 (impact=6, confidence=0.66, effort=1)

Question:
Using the attached code files, identify migrations/backfills/data-shape changes implied in this area. Note rollout risks.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/08-migrations-backfills-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #08  (codebase-driven, group: cmd)

Reference: cmd
Category: background jobs
Horizon: NearTerm
ROI: 4.2 (impact=6, confidence=0.66, effort=1)

Question:
Using the attached code files, identify migrations/backfills/data-shape changes implied in this area. Note rollout risks.

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


# 09) ROI=4.1 impact=6 confidence=0.64 effort=1 horizon=NearTerm category=observability reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/09-ux-flows-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #09  (codebase-driven, group: cmd)

Reference: cmd
Category: observability
Horizon: NearTerm
ROI: 4.1 (impact=6, confidence=0.64, effort=1)

Question:
Using the attached code files, describe the main user flows in this area. Note any fragile or confusing steps.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/09-ux-flows-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #09  (codebase-driven, group: cmd)

Reference: cmd
Category: observability
Horizon: NearTerm
ROI: 4.1 (impact=6, confidence=0.64, effort=1)

Question:
Using the attached code files, describe the main user flows in this area. Note any fragile or confusing steps.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/09-ux-flows-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #09  (codebase-driven, group: cmd)

Reference: cmd
Category: observability
Horizon: NearTerm
ROI: 4.1 (impact=6, confidence=0.64, effort=1)

Question:
Using the attached code files, describe the main user flows in this area. Note any fragile or confusing steps.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/09-ux-flows-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #09  (codebase-driven, group: cmd)

Reference: cmd
Category: observability
Horizon: NearTerm
ROI: 4.1 (impact=6, confidence=0.64, effort=1)

Question:
Using the attached code files, describe the main user flows in this area. Note any fragile or confusing steps.

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


# 10) ROI=4.0 impact=6 confidence=0.62 effort=1 horizon=NearTerm category=observability reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/10-failure-modes-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #10  (codebase-driven, group: cmd)

Reference: cmd
Category: observability
Horizon: NearTerm
ROI: 4.0 (impact=6, confidence=0.62, effort=1)

Question:
Using the attached code files, enumerate likely failure modes (network, data, validation, retries). Note missing handling.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/10-failure-modes-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #10  (codebase-driven, group: cmd)

Reference: cmd
Category: observability
Horizon: NearTerm
ROI: 4.0 (impact=6, confidence=0.62, effort=1)

Question:
Using the attached code files, enumerate likely failure modes (network, data, validation, retries). Note missing handling.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/10-failure-modes-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #10  (codebase-driven, group: cmd)

Reference: cmd
Category: observability
Horizon: NearTerm
ROI: 4.0 (impact=6, confidence=0.62, effort=1)

Question:
Using the attached code files, enumerate likely failure modes (network, data, validation, retries). Note missing handling.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/10-failure-modes-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #10  (codebase-driven, group: cmd)

Reference: cmd
Category: observability
Horizon: NearTerm
ROI: 4.0 (impact=6, confidence=0.62, effort=1)

Question:
Using the attached code files, enumerate likely failure modes (network, data, validation, retries). Note missing handling.

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


# 11) ROI=3.8 impact=6 confidence=0.60 effort=1 horizon=NearTerm category=permissions reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/11-feature-flags-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #11  (codebase-driven, group: cmd)

Reference: cmd
Category: permissions
Horizon: NearTerm
ROI: 3.8 (impact=6, confidence=0.60, effort=1)

Question:
Using the attached code files, identify any feature flags or config toggles. Note rollout/rollback behavior and gaps.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/11-feature-flags-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #11  (codebase-driven, group: cmd)

Reference: cmd
Category: permissions
Horizon: NearTerm
ROI: 3.8 (impact=6, confidence=0.60, effort=1)

Question:
Using the attached code files, identify any feature flags or config toggles. Note rollout/rollback behavior and gaps.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/11-feature-flags-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #11  (codebase-driven, group: cmd)

Reference: cmd
Category: permissions
Horizon: NearTerm
ROI: 3.8 (impact=6, confidence=0.60, effort=1)

Question:
Using the attached code files, identify any feature flags or config toggles. Note rollout/rollback behavior and gaps.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/11-feature-flags-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #11  (codebase-driven, group: cmd)

Reference: cmd
Category: permissions
Horizon: NearTerm
ROI: 3.8 (impact=6, confidence=0.60, effort=1)

Question:
Using the attached code files, identify any feature flags or config toggles. Note rollout/rollback behavior and gaps.

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


# 12) ROI=3.9 impact=6 confidence=0.62 effort=1 horizon=NearTerm category=permissions reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/12-caching-state-consistency-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #12  (codebase-driven, group: cmd)

Reference: cmd
Category: permissions
Horizon: NearTerm
ROI: 3.9 (impact=6, confidence=0.62, effort=1)

Question:
Using the attached code files, identify consistency boundaries (read-after-write, eventual vs strong). Note any mismatches across layers.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/12-caching-state-consistency-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #12  (codebase-driven, group: cmd)

Reference: cmd
Category: permissions
Horizon: NearTerm
ROI: 3.9 (impact=6, confidence=0.62, effort=1)

Question:
Using the attached code files, identify consistency boundaries (read-after-write, eventual vs strong). Note any mismatches across layers.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/12-caching-state-consistency-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #12  (codebase-driven, group: cmd)

Reference: cmd
Category: permissions
Horizon: NearTerm
ROI: 3.9 (impact=6, confidence=0.62, effort=1)

Question:
Using the attached code files, identify consistency boundaries (read-after-write, eventual vs strong). Note any mismatches across layers.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/12-caching-state-consistency-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #12  (codebase-driven, group: cmd)

Reference: cmd
Category: permissions
Horizon: NearTerm
ROI: 3.9 (impact=6, confidence=0.62, effort=1)

Question:
Using the attached code files, identify consistency boundaries (read-after-write, eventual vs strong). Note any mismatches across layers.

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


# 13) ROI=3.7 impact=6 confidence=0.58 effort=1 horizon=MidTerm category=migrations reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/13-observability-gaps-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #13  (codebase-driven, group: cmd)

Reference: cmd
Category: migrations
Horizon: MidTerm
ROI: 3.7 (impact=6, confidence=0.58, effort=1)

Question:
Using the attached code files, identify observability gaps that will block triage or SLA guarantees.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/13-observability-gaps-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #13  (codebase-driven, group: cmd)

Reference: cmd
Category: migrations
Horizon: MidTerm
ROI: 3.7 (impact=6, confidence=0.58, effort=1)

Question:
Using the attached code files, identify observability gaps that will block triage or SLA guarantees.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/13-observability-gaps-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #13  (codebase-driven, group: cmd)

Reference: cmd
Category: migrations
Horizon: MidTerm
ROI: 3.7 (impact=6, confidence=0.58, effort=1)

Question:
Using the attached code files, identify observability gaps that will block triage or SLA guarantees.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/13-observability-gaps-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #13  (codebase-driven, group: cmd)

Reference: cmd
Category: migrations
Horizon: MidTerm
ROI: 3.7 (impact=6, confidence=0.58, effort=1)

Question:
Using the attached code files, identify observability gaps that will block triage or SLA guarantees.

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


# 14) ROI=3.6 impact=6 confidence=0.56 effort=1 horizon=MidTerm category=migrations reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/14-permissions-gaps-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #14  (codebase-driven, group: cmd)

Reference: cmd
Category: migrations
Horizon: MidTerm
ROI: 3.6 (impact=6, confidence=0.56, effort=1)

Question:
Using the attached code files, identify authorization edge cases or privilege escalations to test.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/14-permissions-gaps-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #14  (codebase-driven, group: cmd)

Reference: cmd
Category: migrations
Horizon: MidTerm
ROI: 3.6 (impact=6, confidence=0.56, effort=1)

Question:
Using the attached code files, identify authorization edge cases or privilege escalations to test.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/14-permissions-gaps-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #14  (codebase-driven, group: cmd)

Reference: cmd
Category: migrations
Horizon: MidTerm
ROI: 3.6 (impact=6, confidence=0.56, effort=1)

Question:
Using the attached code files, identify authorization edge cases or privilege escalations to test.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/14-permissions-gaps-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #14  (codebase-driven, group: cmd)

Reference: cmd
Category: migrations
Horizon: MidTerm
ROI: 3.6 (impact=6, confidence=0.56, effort=1)

Question:
Using the attached code files, identify authorization edge cases or privilege escalations to test.

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


# 15) ROI=3.5 impact=6 confidence=0.54 effort=1 horizon=MidTerm category=UX flows reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/15-migrations-risk-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #15  (codebase-driven, group: cmd)

Reference: cmd
Category: UX flows
Horizon: MidTerm
ROI: 3.5 (impact=6, confidence=0.54, effort=1)

Question:
Using the attached code files, identify any migration risks, data backfill triggers, or state shape changes that require careful sequencing.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/15-migrations-risk-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #15  (codebase-driven, group: cmd)

Reference: cmd
Category: UX flows
Horizon: MidTerm
ROI: 3.5 (impact=6, confidence=0.54, effort=1)

Question:
Using the attached code files, identify any migration risks, data backfill triggers, or state shape changes that require careful sequencing.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/15-migrations-risk-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #15  (codebase-driven, group: cmd)

Reference: cmd
Category: UX flows
Horizon: MidTerm
ROI: 3.5 (impact=6, confidence=0.54, effort=1)

Question:
Using the attached code files, identify any migration risks, data backfill triggers, or state shape changes that require careful sequencing.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/15-migrations-risk-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #15  (codebase-driven, group: cmd)

Reference: cmd
Category: UX flows
Horizon: MidTerm
ROI: 3.5 (impact=6, confidence=0.54, effort=1)

Question:
Using the attached code files, identify any migration risks, data backfill triggers, or state shape changes that require careful sequencing.

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


# 16) ROI=3.4 impact=6 confidence=0.52 effort=1 horizon=MidTerm category=UX flows reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/16-ux-flow-gaps-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #16  (codebase-driven, group: cmd)

Reference: cmd
Category: UX flows
Horizon: MidTerm
ROI: 3.4 (impact=6, confidence=0.52, effort=1)

Question:
Using the attached code files, identify UX or developer flow bottlenecks; propose smallest flow test to validate.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/16-ux-flow-gaps-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #16  (codebase-driven, group: cmd)

Reference: cmd
Category: UX flows
Horizon: MidTerm
ROI: 3.4 (impact=6, confidence=0.52, effort=1)

Question:
Using the attached code files, identify UX or developer flow bottlenecks; propose smallest flow test to validate.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/16-ux-flow-gaps-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #16  (codebase-driven, group: cmd)

Reference: cmd
Category: UX flows
Horizon: MidTerm
ROI: 3.4 (impact=6, confidence=0.52, effort=1)

Question:
Using the attached code files, identify UX or developer flow bottlenecks; propose smallest flow test to validate.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/16-ux-flow-gaps-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #16  (codebase-driven, group: cmd)

Reference: cmd
Category: UX flows
Horizon: MidTerm
ROI: 3.4 (impact=6, confidence=0.52, effort=1)

Question:
Using the attached code files, identify UX or developer flow bottlenecks; propose smallest flow test to validate.

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


# 17) ROI=3.3 impact=6 confidence=0.50 effort=1 horizon=MidTerm category=failure modes reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/17-failure-modes-debt-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #17  (codebase-driven, group: cmd)

Reference: cmd
Category: failure modes
Horizon: MidTerm
ROI: 3.3 (impact=6, confidence=0.50, effort=1)

Question:
Using the attached code files, list failure handling debt or missing retries/rollbacks and rank by user impact.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/17-failure-modes-debt-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #17  (codebase-driven, group: cmd)

Reference: cmd
Category: failure modes
Horizon: MidTerm
ROI: 3.3 (impact=6, confidence=0.50, effort=1)

Question:
Using the attached code files, list failure handling debt or missing retries/rollbacks and rank by user impact.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/17-failure-modes-debt-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #17  (codebase-driven, group: cmd)

Reference: cmd
Category: failure modes
Horizon: MidTerm
ROI: 3.3 (impact=6, confidence=0.50, effort=1)

Question:
Using the attached code files, list failure handling debt or missing retries/rollbacks and rank by user impact.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/17-failure-modes-debt-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #17  (codebase-driven, group: cmd)

Reference: cmd
Category: failure modes
Horizon: MidTerm
ROI: 3.3 (impact=6, confidence=0.50, effort=1)

Question:
Using the attached code files, list failure handling debt or missing retries/rollbacks and rank by user impact.

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


# 18) ROI=3.2 impact=6 confidence=0.48 effort=1 horizon=LongTerm category=failure modes reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/18-feature-flags-roadmap-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #18  (codebase-driven, group: cmd)

Reference: cmd
Category: failure modes
Horizon: LongTerm
ROI: 3.2 (impact=6, confidence=0.48, effort=1)

Question:
Using the attached code files, identify where staged rollouts or flags should exist but do not.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/18-feature-flags-roadmap-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #18  (codebase-driven, group: cmd)

Reference: cmd
Category: failure modes
Horizon: LongTerm
ROI: 3.2 (impact=6, confidence=0.48, effort=1)

Question:
Using the attached code files, identify where staged rollouts or flags should exist but do not.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/18-feature-flags-roadmap-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #18  (codebase-driven, group: cmd)

Reference: cmd
Category: failure modes
Horizon: LongTerm
ROI: 3.2 (impact=6, confidence=0.48, effort=1)

Question:
Using the attached code files, identify where staged rollouts or flags should exist but do not.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/18-feature-flags-roadmap-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #18  (codebase-driven, group: cmd)

Reference: cmd
Category: failure modes
Horizon: LongTerm
ROI: 3.2 (impact=6, confidence=0.48, effort=1)

Question:
Using the attached code files, identify where staged rollouts or flags should exist but do not.

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


# 19) ROI=3.1 impact=6 confidence=0.46 effort=1 horizon=LongTerm category=feature flags reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/19-background-jobs-scale-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #19  (codebase-driven, group: cmd)

Reference: cmd
Category: feature flags
Horizon: LongTerm
ROI: 3.1 (impact=6, confidence=0.46, effort=1)

Question:
Using the attached code files, identify long-term scaling risks in background processing or async pipelines.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/19-background-jobs-scale-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #19  (codebase-driven, group: cmd)

Reference: cmd
Category: feature flags
Horizon: LongTerm
ROI: 3.1 (impact=6, confidence=0.46, effort=1)

Question:
Using the attached code files, identify long-term scaling risks in background processing or async pipelines.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/19-background-jobs-scale-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #19  (codebase-driven, group: cmd)

Reference: cmd
Category: feature flags
Horizon: LongTerm
ROI: 3.1 (impact=6, confidence=0.46, effort=1)

Question:
Using the attached code files, identify long-term scaling risks in background processing or async pipelines.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/19-background-jobs-scale-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #19  (codebase-driven, group: cmd)

Reference: cmd
Category: feature flags
Horizon: LongTerm
ROI: 3.1 (impact=6, confidence=0.46, effort=1)

Question:
Using the attached code files, identify long-term scaling risks in background processing or async pipelines.

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


# 20) ROI=3.0 impact=6 confidence=0.44 effort=1 horizon=LongTerm category=feature flags reference=cmd

# code files attached directly (deterministic; self-contained)
mapfile -t __code_files < <(python3 - <<'PY'
from __future__ import annotations
import json

files = json.loads(r'''[
  "cmd/oraclepack/main.go"
]''')
for p in files:
    print(p)
PY
)

code_args=()
for p in "${__code_files[@]}"; do
  code_args+=(-f "$p")
done

if [ "${#code_args[@]}" -eq 0 ]; then
  echo "WARNING: no code files resolved for group 'cmd'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/20-contracts-interfaces-roadmap-direct-answer.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #20  (codebase-driven, group: cmd)

Reference: cmd
Category: feature flags
Horizon: LongTerm
ROI: 3.0 (impact=6, confidence=0.44, effort=1)

Question:
Using the attached code files, identify longer-term public surface changes likely needed in this area.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/20-contracts-interfaces-roadmap-risks-unknowns.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #20  (codebase-driven, group: cmd)

Reference: cmd
Category: feature flags
Horizon: LongTerm
ROI: 3.0 (impact=6, confidence=0.44, effort=1)

Question:
Using the attached code files, identify longer-term public surface changes likely needed in this area.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/20-contracts-interfaces-roadmap-next-experiment.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #20  (codebase-driven, group: cmd)

Reference: cmd
Category: feature flags
Horizon: LongTerm
ROI: 3.0 (impact=6, confidence=0.44, effort=1)

Question:
Using the attached code files, identify longer-term public surface changes likely needed in this area.

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

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09/20-contracts-interfaces-roadmap-missing-evidence.md"   "${code_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #20  (codebase-driven, group: cmd)

Reference: cmd
Category: feature flags
Horizon: LongTerm
ROI: 3.0 (impact=6, confidence=0.44, effort=1)

Question:
Using the attached code files, identify longer-term public surface changes likely needed in this area.

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
