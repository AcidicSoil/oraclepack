1) Direct answer (validation boundaries + likely gaps)

Pack parse/structural validation happens in App.LoadPack(): it reads the pack file, calls pack.Parse(data), then p.Validate() and fails fast on any error (internal/app/app.go:LoadPack).

Boundary: “bytes on disk” → “in-memory Pack that is considered valid enough to run.”

Parse-time validation is minimal and mostly syntactic: pack.Parse() rejects packs without a fenced ```bash block (bashFenceRegex) and only recognizes steps whose headers match stepHeaderRegex (internal/pack/parser.go:Parse).

Gap: a Markdown pack can be “accepted” but silently miss intended steps if headers don’t match the regex (e.g., # 1 - ... vs # 01 - ...), leading to Validate() failing later with “at least one step is required” or producing an unexpected step set.

Pack semantic validation is narrow (step numbering only): Pack.Validate() enforces: at least one step, positive numbers, no duplicates, and sequential numbering starting from 1 in parse order (internal/pack/parser.go:Validate).

Gap: it does not validate step code presence/emptiness, command safety, required prelude directives, or cross-step dependencies—so many “bad” packs pass validation and fail at runtime.

Metadata extraction happens without validation: after parsing, DeriveMetadata() extracts out_dir=... and --write-output from the prelude using regexes (internal/pack/parser.go:DeriveMetadata).

Gap: extracted OutDir is not validated for emptiness, path traversal, or consistency with CLI overrides (it’s just a string captured by regex).

Config resolution validation is effectively “filesystem-only”: App.Prepare() resolves outDir by precedence CLI cfg.OutDir > pack p.OutDir > ".", then only validates by attempting os.MkdirAll(outDir, 0755) (internal/app/app.go:Prepare).

Gap: there’s no validation of other config fields here (PackPath, StatePath, ROIMode values, ROIThreshold range, etc.), so invalid config can survive until deeper runtime code.

State validation boundary is weak and can mask corruption: App.LoadState() on Resume tries state.LoadState(...); if it errors, it silently falls back to a fresh RunState{SchemaVersion:1, StepStatuses: map[...]} (internal/app/app.go:LoadState).

Most likely inconsistent-state gap: partial outputs on disk + corrupted/old state file ⇒ the system “resumes” from an empty state and may rerun steps unexpectedly or skip intended recovery paths.

Runtime checks shown are limited to environment shaping: Prepare() appends out_dir=<resolved> into Runner.Env but does not persist the resolved outDir back into Config/State, and does not de-dup/override existing out_dir entries (internal/app/app.go:Prepare).

Gap: scripts may see multiple out_dir values depending on how Runner.Env is constructed elsewhere; plus “where outputs went” is not recorded in state, increasing resume/report inconsistency risk.

Parser robustness gap that can cause silent truncation: pack.Parse() uses bufio.Scanner but does not call scanner.Err() after the scan loop, nor increase the scanner buffer (internal/pack/parser.go:Parse).

Likely inconsistent-state failure mode: long lines in bash blocks can exceed Scanner’s token lim
