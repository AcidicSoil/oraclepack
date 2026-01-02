Direct answer (1–4 bullets, evidence-cited)

Yes—Validate() should reject steps with empty/whitespace-only Code. Today Validate() only enforces “at least one step,” uniqueness, and sequential numbering; it never checks step.Code, so a header immediately followed by the next header (or EOF) yields a step that passes validation but executes nothing. (Evidence: Validate() loop only checks step.Number, duplicates, and step.Number != i+1; Parse() only appends to currentStep.Code on non-header lines.)

Missing bash fence is already rejected by Parse(), but adding a redundant invariant in Validate() can still be reasonable if Validate() may be called on packs not produced by Parse(). (Evidence: Parse() returns ErrInvalidPack when bashFenceRegex.FindSubmatch(content) returns nil: "%w: no bash code block found".)

Duplicated bash fences should be rejected, but the right place is Parse() (or a “ValidateRaw(content)” variant), not the current Validate() signature. The current bashFenceRegex.FindSubmatch captures only the first block and silently ignores additional ```bash fences, and Validate() has no access to the original Markdown to detect duplicates. (Evidence: Parse() uses FindSubmatch, not FindAllSubmatch.)

Risks/unknowns (bullets)

Some legitimate steps might intentionally be “no-op” (comments only) or rely on prelude-only execution; a strict non-empty check may break those workflows unless you define “non-empty” as “contains at least one non-comment, non-whitespace command.”

If any downstream component already tolerates or depends on multiple ```bash blocks (e.g., “append mode”), rejecting duplicates would be a behavior change; you’ll want a clear contract.

It’s unclear whether Validate() is ever called on manually-constructed Pack values (not via Parse()); that affects whether missing/duplicate fence checks belong in Validate() or only in parsing.

Next smallest concrete experiment (1 action)

Add a unit test that parses a pack with two sequential step headers and no lines between them (empty body for step 1) and assert pack.Validate() returns errors.ErrInvalidPack; then implement strings.TrimSpace(step.Code) == "" rejection in Validate().

If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

internal/pack/types.go (to confirm Step/Pack fields and whether raw source is stored, e.g., Pack.Source or similar that could enable fence-dup detection in Validate()).

internal/pack/parser_test.go or internal/pack/*test.go (to align with existing testing patterns and error assertions).

Any runner/executor code that consumes Step.Code (e.g., internal/run/** or cmd/** paths) to confirm whether comment-only or empty steps are currently meaningful.
