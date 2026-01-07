<!-- path: ~/.codex/skills/oracle-pack/assets/oracle-pack-template.md -->
# oracle strategist question pack

---

## parsed args

- codebase_name: <Unknown|value>
- constraints: <None|value>
- non_goals: <None|value>
- team_size: <Unknown|value>
- deadline: <Unknown|value>
- out_dir: <oracle-out|value>
- oracle_cmd: <oracle|value>
- oracle_flags: <--browser-attachments always --files-report|value>
- extra_files: <empty|value>

---

## commands (exactly 20; sorted by ROI desc; ties by lower effort)

```bash
# 01 — ROI=<..> impact=<..> confidence=<..> effort=<..> horizon=<Immediate|Strategic> category=<...> reference=<...>
<oracle_cmd> \
  <oracle_flags> \
  --write-output "<out_dir>/01-<slug>.md" \
  -p "Strategist question #01
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
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "<minimal evidence file 1>" \
  -f "<optional supporting file 2>" \
  <optional extra_files entries...>

# 02 ...
# ...
# 20 ...
```

---

## coverage check (must be satisfied)

*   contracts/interfaces: <OK|Missing (state missing artifact pattern)>

*   invariants: <OK|Missing (...)>

*   caching/state: <OK|Missing (...)>

*   background jobs: <OK|Missing (...)>

*   observability: <OK|Missing (...)>

*   permissions: <OK|Missing (...)>

*   migrations: <OK|Missing (...)>

*   UX flows: <OK|Missing (...)>

*   failure modes: <OK|Missing (...)>

*   feature flags: <OK|Missing (...)>
```