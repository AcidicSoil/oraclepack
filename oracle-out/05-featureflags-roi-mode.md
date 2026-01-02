Direct answer (1–4 bullets, evidence-cited)

Yes — validate --roi-mode to only allow over or under and return an error on anything else. In internal/cli/run.go, roiMode is currently just a free-form string flag with default "over" (StringVar(&roiMode, "roi-mode", "over", ...)) and is passed through unchanged into app.Config{ ROIMode: roiMode }, with no validation shown in RunE. This makes bad inputs silently propagate into filtering behavior.

Normalize case and whitespace before validating (e.g., accept OVER, Under) to avoid surprising user errors; but still hard-error on unknown values to prevent silently including/excluding steps incorrectly. Evidence: the flag help explicitly constrains values ("('over' or 'under')"), but code currently doesn’t enforce that constraint.

Risks/unknowns (bullets)

If downstream code currently treats unknown ROIMode as a default (e.g., “over”), introducing a hard error could be a behavior change for users relying on that accidental fallback.

It’s unclear whether the TUI layer (tui.NewModel(... cfg.ROIMode ...)) or app.Prepare()/pack filtering already validates or coerces ROIMode. If it does, CLI-level validation might be redundant but still beneficial for early failure and clearer error messages.

If you later add more modes, strict validation will require updating both CLI and downstream logic (which is desirable, but adds a small maintenance step).

Next smallest concrete experiment (1 action)

Add a CLI test that runs oraclepack run pack.md --roi-mode=bad and asserts RunE returns an error (non-nil), then add the minimal validation in RunE right before building cfg (lowercase + switch on over|under).

If evidence is insufficient, exact missing file/path pattern(s) to attach next

internal/app/* (where ROI filtering is applied): especially internal/app/app.go (or wherever Config.ROIMode is consumed) and any ROI filter helper.

internal/tui/* (whether TUI coerces/validates): model constructor and ROI filtering UI.

Any pack parsing / ROI annotation code that uses the mode (likely internal/pack/*).
