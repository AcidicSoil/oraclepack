Missing evidence

* `.tickets/actions/**/*.md` (including any actions “Parent Ticket Summary.md” / index files)
* The referenced pack(s): `**/ticket-action-pack.md`, `**/oracle_pack_and_taskify-skills.md`, `**/oraclepack-tui.md`
* The generator/template source for `ticket-action-pack.md`: `**/*ticket-action-pack*.md`, `**/*ticketify*pack*.md`, `**/templates/**`, `**/assets/**`
* Oraclepack command detection + override injection + validation implementation: `**/cmd/oraclepack/**`, `**/internal/**`, `**/*dispatch*.*`, `**/*override*.*`, `**/*validate*.*`
* Run state/report writers + schema definitions for the named artifacts: `**/*state*.go`, `**/*report*.go`, `**/*.state.json*`, `**/*.report.json*`, `**/schema/**`
* TUI surfaces for override/validation UX changes: `**/*tui*/**`, `**/*ui*/**`, `**/*bubbletea*/**`
