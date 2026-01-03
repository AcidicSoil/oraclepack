Title:

* Integrate Glow-based Markdown rendering into oraclepack TUI viewport

Summary:

* Add `glow` (or its underlying Glamour rendering configuration) at the existing Markdown→ANSI rendering choke point so oraclepack can render Markdown inside its current TUI. Rendering should be injected into the right-hand viewport content and re-rendered on selection and window size changes. Avoid embedding Glow’s full Bubble Tea program since it wants to own the event loop.

    Glow Markdown Integration

Background / Context:

* The repo already centralizes Markdown rendering via `internal/render/render.go` (currently using Glamour).

* The request is: “Where’s the best spot to add `glow` in for its md rendering?” referencing Glow.

    Glow Markdown Integration

Current Behavior (Actual):

* Markdown rendering is handled via the existing renderer in `internal/render/render.go` (Glamour-based), without Glow-specific integration/configuration described.

    Glow Markdown Integration

Expected Behavior:

* Markdown should be rendered for display inside oraclepack’s existing TUI by rendering in `internal/render` and setting the rendered output into the viewport (rather than running Glow’s own UI program).

    Glow Markdown Integration

Requirements:

* Keep all Glow/Glamour rendering concerns isolated to `internal/render` (style selection, width, etc.).

    Glow Markdown Integration

* Expose a small rendering API such as `RenderMarkdown(text, width, style)` (or cache renderers keyed by width/style).

    Glow Markdown Integration

* Populate rendered content at the view boundary where the right-hand viewport is set (e.g., `m.viewport.SetContent(...)` in `internal/tui/tui.go`).

    Glow Markdown Integration

* Re-render on:

  * Selection changes (after `m.list.Update(msg)` when not running)

  * Window resizes (`handleWindowSize`), using viewport width for wrapping

        Glow Markdown Integration

* Do not embed Glow’s full Bubble Tea UI program for inline rendering (Glow’s Go surface is primarily `ui.NewProgram(...)` and expects to own the event loop).

    Glow Markdown Integration

* If “open in Glow” is desired, implement as a separate mode/command (or shell out), not inside the core TUI update loop.

    Glow Markdown Integration

Out of Scope:

* Embedding/running Glow’s full UI program inside oraclepack’s existing Bubble Tea loop.

    Glow Markdown Integration

Reproduction Steps:

* Not provided.

    Glow Markdown Integration

Environment:

* OS: Unknown

* oraclepack version/commit: Unknown

* Terminal: Unknown

    Glow Markdown Integration

Evidence:

* User reference: `oraclepack-all.md` (file context mentioned).

    Glow Markdown Integration

* Glow repository: `https://github.com/charmbracelet/glow`

    Glow Markdown Integration

* Glow Go UI surface mentioned: `https://pkg.go.dev/github.com/charmbracelet/glow/v2/ui`

    Glow Markdown Integration

* Target code points mentioned:

  * `internal/render/render.go`

  * `internal/tui/tui.go`

  * `m.viewport.SetContent(...)`

  * `handleWindowSize`

        Glow Markdown Integration

Decisions / Agreements:

* Integrate at `internal/render/render.go` as the rendering choke point (not by embedding Glow UI). (per assistant)

    Glow Markdown Integration

* Inject rendered content into the viewport and re-render on selection/resize. (per assistant)

    Glow Markdown Integration

Open Items / Unknowns:

* Whether the goal is strictly “better rendering inside the current TUI” vs. adding an additional “open in Glow” feature.

    Glow Markdown Integration

* Desired style/theme configuration requirements (e.g., which Glamour style, environment variables, config file).

    Glow Markdown Integration

* Performance expectations (caching strategy requirements).

    Glow Markdown Integration

Risks / Dependencies:

* Dependency: Glow’s primary Go interface is a Bubble Tea program; embedding it would conflict with oraclepack’s existing event loop.

    Glow Markdown Integration

* Rendering correctness/performance may require caching keyed by viewport width/style.

    Glow Markdown Integration

Acceptance Criteria:

* `internal/render` exposes a rendering function that accepts content and a width (and optionally style), returning ANSI-rendered text.

    Glow Markdown Integration

* The TUI sets viewport content using the renderer output at the viewport boundary (`m.viewport.SetContent(...)`).

    Glow Markdown Integration

* On window resize, viewport content is re-rendered using the updated viewport width.

    Glow Markdown Integration

* On selection change (when not running), viewport content is re-rendered for the newly selected item.

    Glow Markdown Integration

* No use of `glow/ui.NewProgram(...)` inside the main oraclepack TUI loop for inline rendering.

    Glow Markdown Integration

Priority & Severity (if inferable from text):

* Not provided.

Labels (optional):

* enhancement

* tui

* markdown-rendering

* glow

* glamour

---
