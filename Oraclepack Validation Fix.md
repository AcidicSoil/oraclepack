Title:

* Fix oraclepack pack validation failing when step headers use “—” instead of “)”

Summary:

* oraclepack’s pack validator fails with `invalid pack structure: at least one step is required` when a pack’s ` ```bash ... ``` ` block contains step headers formatted as `# 01 — ...` (em dash) rather than `# 01) ...`. As a result, the parser detects zero steps and validation fails. This appears to affect `docs/strategist-questions-oracle-pack-2025-12-31.md`, while the `2025-12-30` version used the recognized format.

    Error in Oraclepack Validation

Background / Context:

* Step detection in `internal/pack/parser.go` is driven by `stepHeaderRegex`, which requires a closing `)` after a two-digit step number (e.g., `# 01)`, `# 02)`).

    Error in Oraclepack Validation

* The newer pack file uses `# 01 — ROI=...` (em dash), which does not match the regex, leading to `len(p.Steps) == 0` and validation failure (per assistant).

    Error in Oraclepack Validation

Current Behavior (Actual):

* Validator returns: `invalid pack structure: at least one step is required`.

    Error in Oraclepack Validation

* Pack contains a bash code fence, but no steps are recognized because headers do not match `^#\s*(\d{2})\)`.

    Error in Oraclepack Validation

* If multiple bash fences exist, the parser may take the first one; if that first fence has unrecognized headers, the run still results in zero steps (per assistant).

    Error in Oraclepack Validation

Expected Behavior:

* Packs with clearly numbered step headers should validate successfully.

* Either:

  * The pack format should be consistently enforced/documented as `# NN)`; or

  * The parser should accept additional common step header variants like `# NN —` / `# NN -` (if intended).

        Error in Oraclepack Validation

Requirements:

* Ensure at least one runnable step is detected for valid packs.

* Provide one of the following resolutions:

  * Documentation/content fix: update pack step headers to include `)` (e.g., `# 01) ...`).

        Error in Oraclepack Validation

  * Code fix: expand `stepHeaderRegex` to accept `# 01 —` / `# 01 -` in addition to `# 01)`.

        Error in Oraclepack Validation

* If multiple bash fences are supported, ensure the correct fence is parsed or improve selection logic (risk noted; not confirmed).

    Error in Oraclepack Validation

Out of Scope:

* Not provided

Reproduction Steps:

1. Create/open a pack markdown file whose first ` ```bash ... ``` ` block contains headers like `# 01 — ROI=...` (em dash) rather than `# 01) ROI=...`.

    Error in Oraclepack Validation

2. Run oraclepack validation on the pack.

3. Observe error: `invalid pack structure: at least one step is required`.

    Error in Oraclepack Validation

Environment:

* OS: Unknown

* oraclepack version: Unknown

* Command used to validate: Unknown

Evidence:

* Error message: `invalid pack structure: at least one step is required`.

    Error in Oraclepack Validation

* Parser rule: `stepHeaderRegex = regexp.MustCompile(^#\s*(\d{2})\))` (requires `# 01)`).

    Error in Oraclepack Validation

* Format difference:

  * Failing: `# 01 — ROI=...` in `docs/strategist-questions-oracle-pack-2025-12-31.md`.

        Error in Oraclepack Validation

  * Working: `# 01) ROI=...` in `strategist-questions-oracle-pack-2025-12-30.md`.

        Error in Oraclepack Validation

Decisions / Agreements:

* Not provided (only options suggested: content change vs parser regex expansion).

    Error in Oraclepack Validation

Open Items / Unknowns:

* Whether supporting `# NN —` / `# NN -` is desired behavior or the pack format should be strict.

* Whether multiple bash fences are valid and, if so, how oraclepack should select the correct one.

* oraclepack version and exact validation command used.

Risks / Dependencies:

* Expanding the regex may introduce ambiguous parsing or false positives if other headings match the broadened pattern.

* Changing fence-selection logic (if applicable) could affect existing packs relying on current “first bash fence” behavior (per assistant).

    Error in Oraclepack Validation

Acceptance Criteria:

* A pack using `# 01) ...` headers validates successfully (baseline).

* If expanded-format support is implemented: a pack using `# 01 — ...` (and/or `# 01 - ...`) headers validates and produces ≥1 parsed step.

* If strict-format enforcement is chosen instead: validation error clearly indicates the required header format (e.g., `# NN)`), and docs/templates reflect it.

* Regression check: existing packs with `# NN)` continue to parse identically.

Priority & Severity (if inferable from text):

* Priority: Not provided

* Severity: Not provided

Labels (optional):

* bug

* parser

* validation

* markdown

* oraclepack

---
