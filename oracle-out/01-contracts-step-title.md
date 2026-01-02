1. Direct answer (1–4 bullets, evidence-cited)

* **Yes**: `Step` should expose a dedicated, non-mutating display field (e.g., `Title` / `DisplayTitle`) so ROI stripping and other normalization don’t require altering the raw header needed for reporting and round-tripping. Evidence: the current `Step` model only has `OriginalLine` (raw header) and `ROI` (parsed numeric), with **no separate place** to store a cleaned title; that structure pressures the parser to either mutate `OriginalLine` or drop information. (See `Step.OriginalLine` and `Step.ROI` in `internal/pack/types.go`.)
* Keep `OriginalLine` as **immutable canonical input**, and derive `Title` from it during parse. This preserves a stable “source of truth” while still supporting a user-friendly label for UI/reporting. (Same evidence: `OriginalLine` is explicitly “The header line,” implying it should remain raw in `internal/pack/types.go`.)

2. Risks/unknowns (bullets)

* Unknown where ROI stripping currently happens and whether it already mutates `OriginalLine` (parsing code not provided), so the scope of change may include parser logic, reporters, and any round-trip writer.
* You’ll need to define the contract precisely: what counts as `Title` (e.g., trim step number prefix? keep punctuation? preserve casing?), and how it behaves when ROI is absent or malformed.
* Adding a new field can break downstream consumers if `Step` is serialized/deserialized (JSON/TOML/YAML) or if code assumes “header text” is only `OriginalLine`.

3. Next smallest concrete experiment (1 action)

* Add a unit test around the step-header parse function that feeds a header containing ROI and asserts **both** are preserved: `OriginalLine` matches the exact input line byte-for-byte, while `Title` equals the ROI-stripped display value and `ROI` parses correctly (and does not affect `OriginalLine`). This single test will force the design and reveal where mutation is occurring.

4. If evidence is insufficient, exact missing file/path pattern(s) to attach next

* The step parsing implementation (one of): `internal/pack/parse.go`, `internal/pack/parser.go`, `internal/pack/pack.go`, or any file where step headers are tokenized/parsed.
* Any round-tripping/writer code (one of): `internal/pack/write*.go`, `internal/pack/render*.go`.
* Existing tests (if any): `internal/pack/*_test.go`.
