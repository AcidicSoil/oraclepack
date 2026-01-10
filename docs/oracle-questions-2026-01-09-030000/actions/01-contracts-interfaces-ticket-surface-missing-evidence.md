Missing evidence

* `.tickets/actions/*.md` (full raw tickets, not summaries)
* `ticket-action-pack.md` (and its generator/template source): `**/*ticket-action-pack*.md`, `**/*ticketify*pack*.md`, `**/templates/**`, `**/assets/**`
* “taskify” Action Pack generator implementation (for proposed `mode=codex|gemini`): `**/*taskify*.*`, `**/taskify/**`, `**/oraclepack-taskify/**`
* Oraclepack CLI/TUI public surface (commands/flags/help) and wiring: `**/cmd/**`, `**/internal/cli/**`, `**/*root*.go`, `**/*commands*.go`
* Oraclepack dispatcher / command-detection / override-injection / validation pipeline (regex currently `^(\\s*)(oracle)\\b` per tickets): `**/*dispatch*.*`, `**/*detect*.*`, `**/*override*.*`, `**/*validate*.*`, `**/internal/**`
* Oraclepack TUI override-validation flow (to assess user-visible behavior/back-compat): `**/*tui*.*`, `**/internal/**tui**/**`
* Docs explicitly referenced by the tickets: `**/oraclepack-tui.md`, `**/oracle_pack_and_taskify-skills.md`
