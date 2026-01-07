1) Direct answer (1–6 bullets, evidence-cited)

Primary public interface #1: CLI root command oraclepack (Cobra)

Inputs: persistent flags

--no-tui (bool, default false) to disable TUI

--oracle-bin (string, default "oracle") path to oracle binary

--out-dir / -o (string, default "") output directory for step execution

Outputs: process exit code is mapped via errors.ExitCode(err) when rootCmd.Execute() fails.

Defined in: internal/cli/root.go (root command + persistent flags + Execute error handling).

Primary public interface #2: CLI command validate [pack.md]

Inputs: one positional arg pack.md (cobra.ExactArgs(1)), passed as app.Config{PackPath: args[0]}

Outputs: on success prints Pack is valid.; on failure returns an error from a.LoadPack() (then surfaced by root Execute path).

Defined in: internal/cli/cmds.go (validate command + wiring into root).

Primary public interface #3: CLI command list [pack.md]

Inputs: one positional arg pack.md (cobra.ExactArgs(1)), passed as app.Config{PackPath: args[0]}

Outputs: prints one line per step: "<step.ID>: <step.OriginalLine>", iterating a.Pack.Steps.

Defined in: internal/cli/cmds.go (list command + wiring into root).

Primary public interface #4: “oracle-pack” Markdown contract (pack file)

Inputs (implied by CLI): a Markdown file path (pack.md) that can be loaded by a.LoadPack() from app.New(cfg) where cfg.PackPath is set.

Outputs (observable via list): the parsed pack must produce a.Pack.Steps, and each step must at least have fields ID and OriginalLine (since they’re printed).

Where defined in code: only the usage is shown in internal/cli/cmds.go; the actual contract/schema + parser are in the internal/app package (not included in the provided files).

2) Risks/unknowns (bullets)

Additional CLI commands/flags may exist in other internal/cli/*.go files; only validate and list are visible here.

The actual oracle-pack Markdown schema (required headings, step syntax, validation rules, parsing behavior, versioning) is not visible; we only know LoadPack() yields Pack.Steps[] with ID and OriginalLine.

Exit-code mapping and error taxonomy depend on internal/errors (only referenced, not shown).

--oracle-bin / --out-dir are declared at root-level but not used in the shown commands; actual execution interface likely exists in other commands not included.

3) Next smallest concrete experiment (exactly one action)

Run a repo-wide search for the contract types and parser entrypoints: rg -n "type Pack|type Step|LoadPack\\(|PackPath|OriginalLine" internal/ to locate the canonical oracle-pack Markdown schema definitions.

4) If evidence is insufficient, missing file/path pattern(s) to attach next

internal/app/*.go (especially files defining Config, New, LoadPack, and any Pack / Step structs)

internal/errors/*.go (to confirm public error*
