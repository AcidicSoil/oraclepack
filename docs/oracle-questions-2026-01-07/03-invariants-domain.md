1) Direct answer (1–8 bullets, evidence-cited)

Pack must contain a fenced ```bash ... ``` code block (the pack is parsed only from that block).
Enforced: Parse() searches bashFenceRegex and errors if none found (no bash code block found).
Not enforced / edge: if multiple bash blocks exist, only the first match is used (FindSubmatch), others are ignored.
Evidence: bashFenceRegex = ... "```bash\n(.*?)\n```" and if match == nil ... return ... ErrInvalidPack in internal/pack/parser.go.

Steps are defined only by Markdown-style headers inside the bash block, and must start with # + a 2-digit step id plus an accepted separator.
Enforced (implicitly by parsing): step boundaries are recognized only when stepHeaderRegex matches a line; everything else is either prelude (before first step) or step body.
Invariants implied: step IDs must be exactly two digits (\d{2}) and separator must be one of ), —, or - (with whitespace variants).
Evidence: stepHeaderRegex = regexp.MustCompile(^#\s*(\d{2})(?:\)|[\s]+[—-])) and the headerMatch := ...; if len(headerMatch) > 1 { ... } logic in internal/pack/parser.go.

All lines before the first step header become “Prelude”; all lines after step parsing begins are assigned to the current step’s Code.
Enforced: inSteps toggles true at first header; before that lines go to preludeLines; after that lines append to currentStep.Code.
Should be considered: any configuration intended for metadata extraction must appear in the prelude, or it won’t be discovered.
Evidence: var preludeLines []string; var inSteps bool + if inSteps { currentStep.Code += ... } else { preludeLines = append(...) } and pack.Prelude.Code = strings.Join(preludeLines, "\n") in internal/pack/parser.go.

At least one step is required for a valid pack.
Enforced: Validate() rejects packs with len(p.Steps) == 0. Note: Parse() itself does not call Validate()—caller must.
Evidence: if len(p.Steps) == 0 { return ... ErrInvalidPack } in internal/pack/parser.go.

Step numbers must be positive, unique, and sequential starting at 1 in the encountered order (i.e., the file order is the order).
Enforced: Validate() checks step.Number > 0, no duplicates (seen map), and step.Number == i+1. This implies the first step must be 01 (because 00 parses to 0 and fails).
Evidence: seen := make(map[int]bool) + duplicate check + if step.Number != i+1 { ... sequential starting from 1 ... } in internal/pack/parser.go.

Metadata invariants: out_dir=... and --write-output are derived only from the Prelude, with strict-ish patterns.
Enforced (extraction behavior): DeriveMetadata() looks only at p.Prelude.Code; out_dir must match ^out_dir=... with no spaces around =, and --write-output is a simple substring/regex presence test.
Not enforced: packs can place these in steps and still “work” at runtime, but metadata will be wrong; multiple out_dir= lines are not resolved explicitly (first match_
