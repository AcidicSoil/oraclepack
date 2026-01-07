1) Direct answer (1–8 bullets, evidence-cited)

State schema has an explicit version field but no migration logic is implemented. RunState includes SchemaVersion int \json:"schema_version"`ininternal/state/types.go, but internal/state/persist.go::LoadStatesimplyjson.Unmarshals into RunStateand returns it without checking or upgradingSchemaVersion`.

Evidence: internal/state/types.go (RunState.SchemaVersion) and internal/state/persist.go (LoadState unmarshal + return; no branching on version).

State evolution currently relies on Go JSON’s “additive compatibility” behavior. New optional fields like ROIThreshold, ROIMode, and Warnings are tagged omitempty in internal/state/types.go, meaning:

older state files missing these fields will still load (they default to zero values),

newer state files with extra fields would be ignored by older binaries (Go encoding/json ignores unknown fields by default).

Evidence: internal/state/types.go JSON tags on ROIThreshold, ROIMode, Warnings.

There is no tooling or “migration runner” in the shown code paths. The persistence layer only provides atomic write + read (SaveStateAtomic, LoadState) and does not do version stamping, backfills, or transforms.

Evidence: internal/state/persist.go contains only marshal/write/rename and read/unmarshal.

Practical implication: only additive changes are safe; renames/removals are breaking. Because no upgrade step exists, any future schema changes that rename/remove fields, change meaning, or change requiredness will either silently degrade (defaults) or break logic downstream.

Where migrations would be applied (natural chokepoint): internal/state/persist.go::LoadState (or the caller immediately after it) is the obvious place to:

read → unmarshal

check SchemaVersion against a CurrentSchemaVersion constant

run sequential migrations migrateV1toV2, migrateV2toV3, etc.

set SchemaVersion = CurrentSchemaVersion

optionally rewrite the upgraded file via SaveStateAtomic

Evidence for chokepoint existence: LoadState/SaveStateAtomic are the only provided persistence entry points in internal/state/persist.go.

Report format evolution cannot be confirmed from the provided snippets. Earlier context shows GenerateReport returns a *ReportV1, which suggests intent to version report structures by type name, but there’s no evidence (in the snippets here) of report persistence, report version fields, or report migrations.

Evidence: internal/report/generate.go shows GenerateReport(...) *ReportV1 (from prior snippet), but no internal/report/*.go definitions/persistence were provided here.

2) Risks/unknowns

Unknown whether SchemaVersion is ever set/non-zero (e.g., on initial state creation), and what the “current” version is meant to be (no constant or assignment shown).

Unknown whether there is any hidden migration handling in the caller (likely internal/app/run.go or similar) before/after LoadState.

Unknown report persistence/read path (where report files are written/read, whether a version field exists, whether multiple report versions are supported).

Unknown whether pack/config/overrides have versioning (separate from state schema; not shown in the provided files).

3) Next smallest concrete experiment (exactly one action)

Run and paste the output of:
rg -n "SchemaVersion|schema_version|CurrentSchema|migrat(e|ion)|ReportV[0-9]|report_version" internal/

4) If evidence is insufficient, exact missing file/path pattern(s) to attach next

internal/report/*.go (especially the file defining ReportV1 and any report write/read functions)

internal/app/run.go (or whichever file calls state.LoadState / state.SaveStateAtomic, and whi
