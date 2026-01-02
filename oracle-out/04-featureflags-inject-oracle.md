1. Direct answer (1–4 bullets, evidence-cited)

* **Yes: make `InjectFlags` explicitly idempotent (don’t re-add flags that are already present on the same `oracle` invocation).** As written, it *always* appends `flagStr` for any matching line (`if oracleCmdRegex.MatchString(line) { line += flagStr }`), so running injection twice or injecting into a line that already contains the same flags will duplicate them and can change semantics or break parsing for flags that shouldn’t repeat. Evidence: `line += flagStr` unconditional append in `InjectFlags`.
* **“Skip commented lines” for `# oracle ...` is already satisfied by the current regex**, because it only matches lines that begin (after optional whitespace) with `oracle` (`^(\s*)oracle(\s+|$)`), and `# oracle ...` does not match that anchor+token requirement. Evidence: `oracleCmdRegex = regexp.MustCompile(^(\s*)oracle(\s+|$))` and the single match gate `oracleCmdRegex.MatchString(line)`.
* **Still worth adding an explicit comment guard (`^\s*#`) anyway** as a defensive invariant (prevents future regex edits from accidentally matching, and makes intent obvious in code).

2. Risks/unknowns (bullets)

* **Flag representation ambiguity:** `flags []string` is joined with spaces, but it’s unclear whether values are encoded as `--k=v`, `--k v` (two slice elements), or quoted strings. A robust “already present” check depends on that convention. Evidence: `flagStr := " " + strings.Join(flags, " ")` shows `InjectFlags` treats `flags` as pre-tokenized strings, but not how they’re formed.
* **Shell parsing complexity:** lines can contain quotes, variable expansions, line continuations (`\`), multiple commands per line (`cmd1 && oracle ...`), or wrappers (`sudo oracle ...`). The current regex only matches `oracle` at the beginning of the line (after whitespace), so those cases are currently not injected (might be intended, might be a gap).
* **Dedup criteria:** if you only dedup “exact flags being injected” you get idempotence; if you instead “skip if *any* flags exist”, you’ll fail to inject required global flags on commands that already have unrelated flags.

3. Next smallest concrete experiment (1 action)

* Add a table-driven unit test (e.g., `internal/exec/inject_test.go`) that asserts **(a)** `# oracle ...` is unchanged, **(b)** `oracle --foo` only receives missing injected flags once, and **(c)** `InjectFlags(InjectFlags(script, flags), flags) == InjectFlags(script, flags)` (idempotence).

4. If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

* Call sites that construct and pass `flags` into `InjectFlags` (to determine token/value conventions): `internal/exec/*.go`, `cmd/oraclepack/**/*.go`, and any TUI/runner code that builds the flag list.
* Any existing tests or helpers around script/step rewriting: `internal/exec/*test*.go` and `internal/pack/*.go` if scripts are assembled there.
