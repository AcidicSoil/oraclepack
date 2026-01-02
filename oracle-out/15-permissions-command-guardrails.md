1) Direct answer (1–4 bullets, evidence-cited)

Yes—add an optional command policy (denylist first) aimed at “I’m about to run an untrusted pack” scenarios. The project explicitly executes arbitrary bash embedded in Markdown (“interactive bash scripts embedded within Markdown files”), so destructive steps are a core risk surface.

Make it opt-in with a clear “unsafe” escape hatch, not a silent always-on gate. The README positions oraclepack as a general workflow runner (not a sandbox), so a strict default policy would likely break legitimate packs; opt-in preserves existing behavior while offering guardrails.

Implement the filter at the last responsible moment: immediately before spawning the shell for each step (runtime), plus optionally at validate-time as warnings. Runtime is the only place you can reliably block execution; validate-time can only warn. This aligns with the current architecture: a CLI entrypoint calls internal CLI, which then drives execution.

Prefer a denylist of high-risk primitives + path/redirect patterns over a strict allowlist initially. An allowlist is hard to make usable for general bash; a denylist can target obvious footguns (e.g., rm -rf /, :(){ :|:& };:, dd if=, destructive chmod/chown on /, etc.) while you learn what users run. The README’s “one bash code block executed” means you’ll otherwise be parsing shell strings with many bypasses, so start coarse and transparent.

2) Risks/unknowns (bullets)

Bypass risk / “security theater”: string matching can be evaded via quoting, variable indirection, functions, heredocs, subshells, command/builtin, xargs, etc. A real solution would require parsing shell (AST) or executing in a constrained environment (container/VM), which is larger scope.

False positives / usability: a denylist can block legitimate automation (e.g., rm -rf dist/), creating friction unless policy is configurable per pack/project.

Cross-platform semantics: Windows + Git Bash/WSL wrappers change what “destructive” means and which paths are “system.” README already documents multiple Windows execution modes, increasing edge cases for path-based rules.

No evidence yet of where shell execution is centralized: I can’t confirm the best insertion point without the exec path (runner/app/tui execution code), so the “last responsible moment” recommendation is architectural rather than line-specific.

3) Next smallest concrete experiment (1 action)

Locate the single choke point where step code is turned into an exec.Command(...) (or equivalent) and add a failing unit test that proves you can block one step without affecting parsing. Concretely: search for exec.Command( and the -c / shell invocation usage and identify the function that runs a step; that is your filter insertion point.

4) If evidence is insufficient, exact missing file/path pattern(s) to attach next

To name the insertion point precisely (and suggest the smallest safe API surface), attach:

internal/exec/runner.go (or whatever file contains the step execution / shell spawn)

internal/app/run.go (wiring between CLI/TUI and runner)

internal/cli/run.go (flags + config plumbing)

internal/tui/tui.go (step-run flow; where per-step decisions are made)

Any existing config loading file (commonly internal/config/** or similar) if present
