<!-- # path: oraclepack-usecase-rpg/references/oracle-pack-template.md -->
# Oracle Pack — {{codebase_name}} (RPG-infused)

## Parsed args
- codebase_name: {{codebase_name}}
- constraints: {{constraints}}
- non_goals: {{non_goals}}
- team_size: {{team_size}}
- deadline: {{deadline}}
- out_dir: {{out_dir}}
- oracle_cmd: {{oracle_cmd}}
- oracle_flags: {{oracle_flags}}
- extra_files: {{extra_files}}

## Commands
```bash
out_dir="{{out_dir}}"

# 01) ROI=... impact=... confidence=... effort=... horizon=Immediate category=contracts/interfaces reference=...
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/01-<slug>.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #01
RPG:
FunctionalNode: <Capability>::<Feature> (WHAT)
StructuralNode: <module-or-dir> :: <public surface> (HOW)
DependsOn: []
Phase: P0

Reference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>
Category: contracts/interfaces
Horizon: Immediate
ROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)
Question: <question text>
Rationale: <exactly one sentence>
Smallest experiment today: <exactly one action>

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ...
# ...
# 20) ...
```

## Coverage check

Mark each as `OK` or `Missing(<which step ids>)`:

*   contracts/interfaces: OK
*   invariants: OK
*   caching/state: OK
*   background jobs: OK
*   observability: OK
*   permissions: OK
*   migrations: OK
*   UX flows: OK
*   failure modes: OK
*   feature flags: OK

```