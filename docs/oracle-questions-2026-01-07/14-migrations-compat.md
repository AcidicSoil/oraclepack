1) Direct answer (backward/forward compatibility expectations; where ensured vs risky)

Run state is explicitly versioned; expectation is “migrate-by-version when needed.” RunState carries SchemaVersion int \json:"schema_version"`which is the primary hook for backward/forward compatibility across persisted state formats. Evidence:internal/state/types.go (RunState.SchemaVersion`).

State schema is designed to be additive (backward-compatible) via optional fields. Fields like ROIThreshold, ROIMode, and Warnings are tagged omitempty, implying they can be absent in older saved state and should default safely in newer code. Evidence: internal/state/types.go (ROIThreshold ...omitempty, ROIMode ...omitempty, Warnings ...omitempty).

Forward compatibility for state/report within Go likely relies on default encoding/json behavior (unknown fields ignored). The structs are plain JSON-tagged types with no visible strict decoding controls here; in Go, older binaries typically ignore new fields when unmarshalling (unless DisallowUnknownFields is used elsewhere). Evidence: presence of standard JSON tags across structs in internal/state/types.go and internal/report/types.go (no strictness shown).

Report schema is implicitly versioned by type name, not by an on-the-wire version field (riskier for evolvability). The report root type is ReportV1, but there is no schema_version/report_version JSON field on ReportV1. This suggests compatibility depends on file naming/conventions or external knowledge, not self-describing payloads. Evidence: internal/report/types.go (type ReportV1 struct { ... } with no version field).

Report is partially “consumer-friendly” via redundant duration fields, suggesting an evolution/compat strategy. Summary.TotalDuration (a time.Duration) is paired with TotalDurationMs int64, and StepReport.Duration is paired with DurationMs int64. This looks like an intentional compatibility measure for consumers that can’t/won’t parse Go time.Duration encoding. Evidence: internal/report/types.go (TotalDuration + TotalDurationMs, Duration + DurationMs).

Two distinct Warning types exist (state vs report), duplicating schema; compatibility depends on keeping them in lockstep. Both packages define Warning with the same JSON fields; this is currently compatible, but any future divergence becomes a schema drift risk. Evidence: internal/state/types.go type Warning struct { Scope, StepID, Line, Token, Message } and internal/report/types.go type Warning struct { ...same fields... }.

Map-based step status storage is tolerant of step additions but sensitive to step ID churn. RunState.StepStatuses map[string]StepStatus supports new step IDs appearing over time, but renames/removals can break continuity of resumed/inspected runs unless migration logic remaps keys. Evidence: internal/state/types.go (StepStatuses map[string]StepStatus).

2) Risks/unknowns

Unknown whether reads/writes are strict or tolerant. If any loader uses json.Decoder.DisallowUnknownFields(), forward compatibility (new fields read by old binaries) will break. Missing evidence: state/report load code.

No visible migration dispatcher. SchemaVersion exists, but we don’t see code that applies migrations on read (or stamps versions on write), so rollback/upgrade safety is unproven from the provided files alone.

Report payload is not self-describing. Without a report_version, consumers can’t robustly branch on schema, making forward/backward compatibility brittle as ReportV2 emerges.

G
