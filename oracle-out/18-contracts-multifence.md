1. Direct answer (1–4 bullets, evidence-cited)

* Yes—supporting **multiple bash fences requires either an explicit selector or a hard error on ambiguity**, because `Parse` currently uses `bashFenceRegex.FindSubmatch(content)`, which returns only the **first** match, so any example ` ```bash ``` ` block above the real pack will be silently chosen. Evidence: `bashFenceRegex = ...` and `match := bashFenceRegex.FindSubmatch(content)` in `internal/pack/parser.go:Parse`.
* Prefer **explicit selection with a backward-compatible default**: parse *all* bash fences, then:

  * if exactly one → use it (current behavior),
  * if multiple and exactly one is marked → use the marked one,
  * otherwise → return `ErrInvalidPack` with a message explaining how to mark/select. This avoids “surprising success” where the wrong block runs. Evidence: current behavior errors only when *no* bash block is found (`no bash code block found`), not when multiple exist.
* A practical “marker” that doesn’t require new Markdown syntax: require a sentinel comment immediately before the intended fence (e.g., `<!-- oraclepack -->`), or allow an info-string variant like <code>```bash oraclepack</code> and match that first. This keeps author intent unambiguous while allowing examples elsewhere.

2. Risks/unknowns (bullets)

* **Compatibility risk:** changing the default from “first fence” to “last fence” (or “error if multiple”) will break existing docs that happen to contain multiple bash fences; that’s why a marker + “error on ambiguous” is safer than changing implicit selection.
* **Parser/regex limitations:** the current fence regex is strict (` ```bash\n ... \n``` `). Real-world Markdown often includes `\r\n`, trailing spaces, or `bash {#id}`-style info strings; expanding matching may be needed once you support multiple fences.
* **UX surface unknown:** whether selection should be purely in-doc (marker) or also via CLI flag depends on CLI design (not shown here).

3. Next smallest concrete experiment (1 action)

* Add a single fixture/test with **two bash fences**: first is an “example” block, second is the real pack (with steps). Define expected behavior as: **fail with an ambiguity error unless one fence is marked** (or unless a selector flag is used). Implement using `FindAllSubmatchIndex` (or similar) so the test drives the multi-fence extraction.

4. If evidence is insufficient, exact missing file/path pattern(s) to attach next

* `internal/cli/*.go` (or wherever `run`/`parse` flags live) to see if a selector flag is feasible and consistent.
* `internal/pack/types.go` (Pack/Step structs) and any existing authoring docs/spec (e.g., `README.md` sections describing pack format) to ensure the marker/selector approach matches intended pack contracts.
