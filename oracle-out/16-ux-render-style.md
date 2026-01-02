1. Direct answer (1–4 bullets, evidence-cited)

* Yes—`RenderMarkdown` should not force a dark theme; it currently hardcodes `glamour.WithStandardStyle("dark")`, which will reduce readability for users in light terminals and breaks “respect user environment” UX expectations. (Evidence: `RenderMarkdown` always passes `"dark"`.)
* Keep backward compatibility by defaulting to current behavior (or better: `auto`) while allowing an explicit override via config/CLI/env (e.g., `light|dark|auto`).
* Prefer `auto` as the default long-term: select light/dark based on terminal background detection when available, otherwise fall back to a safe default (likely current `"dark"` to avoid surprising existing users).

2. Risks/unknowns (bullets)

* Terminal background detection can be wrong or unavailable (SSH, tmux/screen, Windows Terminal quirks), so “auto” must have a deterministic fallback.
* Output-based tests can become brittle if Glamour changes style definitions; tests should compare against Glamour’s own renderer output rather than asserting specific ANSI codes.
* Unknown where user config is currently defined/loaded (no evidence here of a config package or CLI flag wiring), so “style option” placement may affect scope and complexity.

3. Next smallest concrete experiment (1 action)

* Add an `options` parameter with `Style string` (accept `light|dark|auto`) to the renderer construction path, and add a unit test that sets `Style:"light"` and asserts `RenderMarkdown(text, opts)` exactly matches a “baseline” output produced by directly calling `glamour.NewTermRenderer(glamour.WithStandardStyle("light"), glamour.WithWordWrap(80)).Render(text)`.

4. If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

* Any existing configuration/flags plumbing to determine where to surface the option:

  * `internal/config/**` (or equivalent)
  * `internal/cli/**` (flags/env parsing)
  * `internal/app/**` (wiring config into render calls)
  * Any existing TUI settings/state that might store user preferences (`internal/tui/**`).
