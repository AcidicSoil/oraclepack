Direct answer (1–8 bullets, evidence-cited)

A pack must contain a fenced Markdown bash block in the exact form bash … ; Parse hard-fails with “no bash code block found” if bashFenceRegex doesn’t match. Enforcement: internal/pack/parser.go:Parse (bashFenceRegex). Missing enforcement: if there are multiple bash fences, the regex match effectively selects only the first; there’s no “exactly one fence” check (even though the intended pack shape is “one bash block”). 

oraclepack-llms-full

Steps are recognized only by comment headers that match ^#\s*(\d{2})(?:\)|[\s]+[—-]) after TrimSpace; this implies (a) exactly two digits are required and (b) only specific separators are accepted (), space+—, space+-). Enforcement: internal/pack/parser.go:Parse (stepHeaderRegex). Missing enforcement: any line inside step bodies that happens to look like a step header (e.g., a comment # 03 - …) will start a new step; there’s no “must be at top-level step boundary” guard.

Step count is an invariant: exactly 20 steps are required. Enforcement: internal/pack/parser.go:Pack.Validate checks len(p.Steps) == 20 and also errors if zero steps. Missing enforcement: Parse itself does not validate; correctness depends on callers consistently invoking Validate. The intended pack contract also assumes 20 steps. 

oraclepack-llms-full

Step numbers must be unique and sequential starting from 1 in the order they appear. Enforcement: Pack.Validate rejects duplicates and enforces step.Number == i+1 (so “01..20 in order” is required, not just “contains 1..20”). Missing enforcement: there is no separate invariant that Step.ID (the captured two-digit string) must equal fmt("%02d", Number)—it’s implicitly assumed.

Prelude/steps split invariant: everything before the first step header is “prelude”; everything after is assigned to the current step until the next header; there is no “postlude”. Enforcement: Parse toggles inSteps on the first header and appends subsequent lines into currentStep.Code. Missing enforcement: trailing content after “step 20” is silently treated as part of step 20, which may surprise pack authors.

ROI metadata is parsed two ways: (a) from the step header via ROI=(\d+(\.\d+)?), and (b) as a fallback from a # ROI: comment inside the step body (only when step.ROI == 0). Enforcement: Parse + applyStepMetadata. Missing enforcement: an explicit header ROI=0 is indistinguishable from “ROI missing” and can be overwritten by an in-body # ROI:; there’s also no range/required-presence validation despite the pack contract expecting ROI tokens in headers. 

oraclepack-llms-full

“Impact” metadata is only parsed from a comment line # Impact: ... inside the step body (not from header tokens like impact=...). Enforcement: applyStepMetadata (impactRegex) populates Step.Impact if present. Missing enforcement: header impact=... is ignored by the parser even though the expected pack header format includes an impact=... token. 

oraclepack-llms-full

Derived pack metadata is best-effort and prelude-scoped: out_dir is extracted from the prelude using ^out_dir=..., and WriteOutput is set true if --write-output appears in the prelude. Enforcement: Pack.DeriveMetadata. Missing enforcement/likely bug: --write-output typically appears in step commands, not the prelude, so WriteOutput may remain false even when every step uses it; out_dir parsing also won’t handle paths with spaces (regex stops at whitespace). 

oraclepack-llms-full

Risks/unknowns (bullets)

Call-path uncertainty: from the provided excerpt, Parse does not call Validate; if any CLI/TUI path forgets to call Validate, malformed packs may execute partially. (Need the loader/runner call
