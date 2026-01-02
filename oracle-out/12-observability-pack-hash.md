Direct answer (1–4 bullets, evidence-cited)

Yes—add a computed pack hash and version metadata. Right now ReportV1 has PackInfo with Name and Hash fields, but nothing in this file indicates the hash is computed/populated (it’s just a string field), and there’s no report schema version or tool/build version in ReportV1 to disambiguate how the report was produced. Evidence: PackInfo is only { Name string; Hash string } and ReportV1 is { Summary; PackInfo; Steps; GeneratedAt } with no Version field.

Hash should be stable and content-derived. Compute PackInfo.Hash as e.g. SHA-256 over the original pack source bytes (not post-parse/mutated representations), with CRLF→LF normalization so the same pack on Windows/Linux produces the same hash.

Add minimal “version metadata” for traceability: at least a report_version (schema) and an oraclepack_version (build/commit), and optionally pack_format_version if your pack format has one. This makes “same hash but different semantics due to a parser change” detectable.

Risks/unknowns (bullets)

What exactly is “the pack” to hash? Raw file bytes vs Pack.Source vs reconstructed text from parsed steps. If ROI stripping or other transforms mutate the source used for hashing, the identity becomes misleading.

Semantically equivalent packs can hash differently (whitespace/comment changes, different quoting, reordered metadata). That’s fine for identity but not for semantic equivalence comparisons—decide which you want.

Schema evolution: adding required version fields can break downstream consumers; keep new JSON fields optional (omitempty) or introduce an explicit schema/version field and document it.

Version provenance: if you want oraclepack_version, you need a single authoritative place to source it (ldflags, debug.ReadBuildInfo, or a version package).

Next smallest concrete experiment (1 action)

Implement ComputePackHash(source string) string (SHA-256 over normalized source), populate ReportV1.PackInfo.Hash in report generation, and add a unit test asserting: same source ⇒ same hash, one-byte change ⇒ different hash, and CRLF vs LF ⇒ same hash.

If evidence is insufficient, exact missing file/path pattern(s) to attach next

internal/report/generate.go (where ReportV1 is constructed/populated—needed to confirm whether PackInfo.Hash is currently set anywhere).

internal/pack/types.go (to confirm whether there is a Pack.Source/raw-source field you can hash without re-reading files).

Wherever pack files are read/normalized (often internal/pack/parser.go or a loader in internal/app/...) and wherever the build/version string is defined (commonly cmd/oraclepack/main.go or internal/version/*.go).
