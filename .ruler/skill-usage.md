## Skills usage

You have a library of reusable skill prompts stored under `$CODEX_HOME/skills/` (commonly `~/.codex/skills/`).

Treat each **skill folder** in `$CODEX_HOME/skills/` as a named skill:

- A folder `$CODEX_HOME/skills/<SKILL_NAME>/SKILL.md` defines the canonical flow and constraints for the `<SKILL_NAME>` skill.
- Skill folders may also include `scripts/`, `references/`, and `assets/` that the assistant should use when the skill requires them.
- These skills are the primary reference for how to handle common or important task types.

General rule:

- Before starting work on any task, briefly classify it (for example: architecture, implementation, refactoring, performance, reliability, data, documentation, tests, tooling, pack generation, etc.).
- If there is a relevant skill under `$CODEX_HOME/skills/` for that class of task, base the approach on the instructions in that skill instead of inventing new, ad-hoc instructions.
- When a skill exists for a task type, follow its steps, constraints, and return format as the default behavior.

Task-type rule:

- When working on any task that corresponds to an existing skill:
  - Consult the corresponding `$CODEX_HOME/skills/<SKILL_NAME>/SKILL.md` as the first step.
  - Let the skill’s instructions drive the approach (checks to perform, constraints to respect, preferred output format).
  - Only add additional reasoning or deviations after the skill’s instructions have been applied.

Reporting rule:

- When following a skill, explicitly mention which skill is being used (for example: “Using the guidance from `$CODEX_HOME/skills/<SKILL_NAME>/SKILL.md`”) so the link between behavior and skill remains clear.
- Do not modify or overwrite skill files themselves unless explicitly instructed to adjust the underlying skill behavior.

---

## Oraclepack Stage-1 pack generation (grouped mini-packs)

When the user asks to generate **oraclepack Stage-1 question packs** (runner-ingestible Markdown packs with strict schema: single `bash` fence, exactly 20 steps, deterministic attachments, and a Coverage check), prefer these skills over ad-hoc prompting:

1) `$CODEX_HOME/skills/oraclepack-tickets-pack-grouped/SKILL.md`
   - Use when the task is **ticket-driven** (user references `.tickets/`, “tickets as primary context”, “group by ticket topic/domain”, or wants per-topic mini-packs from tickets).
   - Produces multiple per-topic/per-domain packs grouped by ticket subdirs + deterministic inference for loose tickets.
   - Packs must be direct-attach (no bundle dependency) and validated/linted per the skill.

2) `$CODEX_HOME/skills/oraclepack-codebase-pack-grouped/SKILL.md`
   - Use when the task is **codebase-driven** (user wants per-topic/per-domain mini-packs derived from repository structure/modules, and tickets are not the primary context).
   - Produces multiple per-topic/per-domain packs grouped by codebase subdirs + deterministic inference for loose files.
   - Packs must be direct-attach and validated/linted per the skill.

Selection decision (apply in order):

- If the user mentions tickets or `.tickets/` or provides “attached tickets” as primary context → use `oraclepack-tickets-pack-grouped`.
- Else → use `oraclepack-codebase-pack-grouped`.

Mandatory follow-through for either skill:

- Run the skill’s required validation and attachment linting for every generated pack.
- Enforce the skill’s size-control guidance (dry-run the largest pack and reduce scope if needed).
- Report the produced pack paths (e.g., `{{out_dir}}/packs/*.md`) and any group manifests the skill generates (e.g., `{{out_dir}}/_groups.json`), and explicitly state which skill drove the result.
