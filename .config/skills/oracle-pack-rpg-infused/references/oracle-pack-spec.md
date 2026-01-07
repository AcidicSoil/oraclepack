# Oracle Pack (RPG-infused) — Spec

## Strict output contract

1) Produce exactly one Markdown deliverable matching `assets/oracle-pack-template.md` structure exactly.
2) The deliverable must contain exactly ONE fenced bash block.
3) The bash block must contain EXACTLY 20 steps, numbered 01..20 sequentially.
4) Each step header line must match:

`# NN ROI=<roi> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>`

5) For each step, emit exactly one runnable oracle command:
   - command: `<oracle_cmd>` (default `oracle`)
   - include `<oracle_flags>` (default `--files-report`)
   - include deterministic output file:

`--write-output "<out_dir>/<nn>-<slug>.md"`

6) Compute ROI:

- Pick `impact`, `confidence`, `effort` each as one decimal in `[0.1..1.0]`.
- Compute `ROI = (impact * confidence) / effort` and round to one decimal.

7) Sort all 20 steps by ROI descending; tie-break by lower effort.

8) Include the coverage check section listing the 10 required categories exactly (OK or Missing(...)).

## Required categories (do not add/remove)

- contracts/interfaces
- invariants
- caching/state
- background jobs
- observability
- permissions
- migrations
- UX flows
- failure modes
- feature flags

## RPG infusion requirements (non-breaking)

A) Each strategist prompt (`-p`) MUST include this RPG block:

- FunctionalNode: capability + feature (WHAT)
- StructuralNode: module/file boundary + interface points (HOW)
- DependsOn: a list of prior step IDs, e.g. [01, 03]
- Phase: P0|P1|P2|P3

B) Dependencies MUST ONLY point backwards:
- DependsOn may reference only earlier step IDs.

C) Across the 20 steps, ensure coverage of BOTH semantics:
- At least 10 steps are functional-first (capability/feature discovery).
- At least 10 steps are structural-first (module/file/interface localization).

D) Ensure at least one step explicitly targets each RPG dimension:
- capabilities/features
- file/module boundaries
- public interfaces/contracts
- data flows/state
- dependency edges/topological build order

## Horizon mix (unchanged)

Exactly 12 Immediate and 8 Strategic.

## Attachment minimization (unchanged)

- Read index artifacts first; infer subsystem locations before sweeping.
- Use deterministic ck/ast-grep/fd queries if available; cap results deterministically.
- Attach only minimal evidence files needed for the single question.
- If reference is Unknown: attach only index files (if possible) and state the missing artifact pattern(s).

## Prompt body format inside -p (must include these sections)

Strategist question #NN
RPG:
FunctionalNode: <Capability>::<Feature> (WHAT)
StructuralNode: <module-or-dir> :: <public surface> (HOW)
DependsOn: [<prior step ids, optional>]
Phase: <P0|P1|P2|P3>

Reference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>
Category: <one of required categories>
Horizon: <Immediate|Strategic>
ROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)
Question: <question text>
Rationale: <exactly one sentence>
Smallest experiment today: <exactly one action>

Constraints: <constraints or None>
Non-goals: <non_goals or None>

Answer format:

1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

## Defaults

If args are missing, use:

- codebase_name=Unknown
- constraints=None
- non_goals=None
- team_size=Unknown
- deadline=Unknown
- out_dir=oracle-out
- oracle_cmd=oracle
- oracle_flags=--files-report
- extra_files=empty

## Final response requirement

In the assistant response:

- Print: `Output file: docs/oracle-pack-YYYY-MM-DD.md`
- Then print the exact same Markdown content (no extra commentary).
