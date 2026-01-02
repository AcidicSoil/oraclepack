1. Direct answer (1–4 bullets, evidence-cited)

* Yes—after `refreshList()` applies ROI filtering, clamp the list selection (and thus the list’s scroll/viewport) to a valid index to avoid “cursor points past end” navigation glitches. Evidence: `refreshList()` only calls `m.list.SetItems(filtered)` and returns, with no attempt to preserve/clamp selection.
* Without clamping, later interactions can observe an invalid/empty selection: `enter` relies on `m.list.SelectedItem()` returning an `item` (`i, ok := m.list.SelectedItem().(item)`), and “run all” paths assume `m.list.Items()[0]` exists after only checking `len(m.list.Items()) > 0`. If filtering removes the previously selected item, keyboard navigation and selection can desync or become `nil` until the user manually moves the cursor.
* Practical rule: after `SetItems`, if `len(items)==0`, clear selection state (or select 0 safely) and show “No steps match filter”; if `len(items)>0`, `idx := min(m.list.Index(), len(items)-1)` then `m.list.Select(idx)`.

2. Risks/unknowns (bullets)

* Bubble Tea’s `list.Model` may already clamp internally on `SetItems`, but that behavior is library-version-dependent; relying on it can regress silently on upgrades.
* Empty list semantics: some list implementations treat “no selection” differently (index 0 vs -1). If your UI assumes “selected item exists,” you should explicitly handle the empty state rather than depend on list internals.
* Current code has other selection-adjacent edge cases you may trip once you add clamping:

  * `resetState()` unconditionally does `m.list.Select(0)` even if the filtered list is empty.
  * Filtering changes do not adjust `m.currentIdx` (safe today because filtering is disabled while running, but worth keeping consistent).

3. Next smallest concrete experiment (1 action)

* Add a unit test that sets the selection to a high index, applies a filter that reduces items, and asserts `m.list.Index()` is clamped and `SelectedItem()` is valid when items exist (and safely handled when items are empty).

4. If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

* Sufficient for the decision. If you want a precise clamp implementation that matches your Bubble Tea version, attach: `go.mod` (Bubble Tea / Bubbles versions) and any TUI tests/utilities under `internal/tui/*_test.go`.
