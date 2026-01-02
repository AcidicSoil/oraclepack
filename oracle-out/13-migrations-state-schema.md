1. Direct answer (1–4 bullets, evidence-cited)

* **Yes—add a migration hook now (even if it’s a no-op today)** because `RunState` already carries an explicit `SchemaVersion int  \`json:"schema_version"``field, which is a strong signal you expect format evolution and need a controlled upgrade path to keep`--resume`working across versions. (internal/state/types.go:`RunState.SchemaVersion`)
* **Relying on Go’s default JSON behavior is not enough for “real” breaking changes**: missing fields zero-initialize (e.g., absent `schema_version` becomes `0`), and unknown fields are ignored, but **renames/type changes/semantic changes** will silently misload or lose information—especially around `StepStatuses map[string]StepStatus` and fields like `PackHash`. (internal/state/types.go: `RunState.StepStatuses`, `RunState.PackHash`)
* **Minimal migration can be cheap**: treat `schema_version==0` as legacy, normalize invariants (e.g., `StepStatuses != nil`), and bump to a `CurrentSchemaVersion`. This makes future breaking changes safer without committing to a heavy framework now. (internal/state/types.go: `RunState` layout)

2. Risks/unknowns (bullets)

* If `LoadState` currently unmarshals directly into `RunState` and proceeds without validation, a migration layer might mask corrupted/partial state unless you also add post-load validation.
* Without seeing resume logic, it’s unclear whether `PackHash` mismatch (empty vs expected) should hard-fail resume, soft-fail into fresh state, or trigger recomputation—migration needs to align with that policy.
* If you later introduce incompatible keying in `StepStatuses` (e.g., step index → step ID), migration will need additional pack context to map old keys correctly.

3. Next smallest concrete experiment (1 action)

* **Implement `MigrateRunState(raw []byte) (*RunState, error)` and call it from `LoadState`**, with one test proving that a legacy JSON blob **missing** `schema_version` (thus `0`) and possibly missing `step_statuses` upgrades to `SchemaVersion==CurrentSchemaVersion` and `StepStatuses` is non-nil.

4. If evidence is insufficient, attach next (exact missing file/path pattern(s))

* `internal/state/persist.go` (especially `LoadState` and how it handles `SchemaVersion`)
* `internal/cli/run.go` (how `--resume` decides to load/abort/fallback)
* Any state-related tests: `internal/state/*_test.go` or `**/*state*_test.go`
