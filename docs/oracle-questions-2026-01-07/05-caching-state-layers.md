Direct answer (1–8 bullets, evidence-cited)

Persisted run-state file (RunState JSON)

What it is: A single JSON document representing an oracle-pack run: schema version, pack hash, start time, per-step status map, ROI settings, and warnings. (internal/state/types.go: RunState{SchemaVersion, PackHash, StartTime, StepStatuses, ROIThreshold, ROIMode, Warnings})

Lifecycle (as implemented in this package): In-memory RunState is serialized to disk and later deserialized back into memory. (internal/state/persist.go: SaveStateAtomic, LoadState)

Read/write points: Writes via SaveStateAtomic(path, state); reads via LoadState(path). (internal/state/persist.go)

Per-step status tracking (StepStatus)

What it is: A per-step record with state machine values pending|running|success|failed|skipped, plus exit code, timestamps, and error string. (internal/state/types.go: type Status, constants; StepStatus{Status, ExitCode, StartedAt, EndedAt, Error})

Lifecycle (represented, not enforced here): A step typically transitions pending → running → (success|failed) or pending → skipped; timestamps support “running window” and duration computations. The code shown defines the model but not the transition logic. (internal/state/types.go)

Run warnings (non-fatal notes)

What it is: Structured warnings attached to run state, scoped and optionally tied to a specific step and line/token. (internal/state/types.go: Warning{Scope, StepID, Line, Token, Message}; RunState.Warnings)

Lifecycle / read/write: Persisted as part of RunState JSON alongside statuses. (internal/state/types.go; internal/state/persist.go)

Atomic persistence mechanism

What it does: Writes path + ".tmp" then renames to path, removing the temp file on rename failure. (internal/state/persist.go: tempPath := path + ".tmp", os.WriteFile, os.Rename, os.Remove)

Operational behavior: Output is pretty-printed JSON (json.MarshalIndent) which aids debugging but can be larger/slower than compact JSON. (internal/state/persist.go)

State “identity” / cache key fields

What it is: PackHash suggests the run state is intended to be tied to a specific pack content hash, likely to detect “pack changed vs. old state”. (internal/state/types.go: RunState.PackHash)

Lifecycle: The code provided stores the value but does not validate it during load; any reconciliation logic must be elsewhere. (internal/state/types.go; internal/state/persist.go)

Risks/unknowns (bullets)

Unknown where state is created/updated. The snippets define the data model and persistence, but not when StepStatuses are initialized, advanced, or saved (e.g., after each step, only at end, on error). That logic is likely in the executor/runner layer, not shown.

Missing “no state yet” behavior. LoadState returns an error when the state file does not exist ("state file not found"), which may force callers to treat “first run” as an error unless they special-case it. (internal/state/persist.go: os.IsNotExist(err))

Map key ambiguity. StepStatuses map[string]StepStatus does not guarantee ordering; the key format (step ID? label?) and collision rules are unknown without the parser/runner code that populates it. (internal/state/types.go)

Cross-platform atomicity. os.Rename(tempPath, path) may not overwrite an existing destination file on some p
