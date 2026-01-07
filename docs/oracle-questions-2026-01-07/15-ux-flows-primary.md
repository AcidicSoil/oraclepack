1) Direct answer (primary TUI user/operator flows, mapped to components + key state transitions)

App entry → TUI boot (CLI “run” command): internal/cli/run.go loads/initializes app (pack, runner, state), then constructs tui.Model via tui.NewModel(a.Pack, a.Runner, a.State, cfg.StatePath, cfg.ROIThreshold, cfg.ROIMode, runAll) and runs Bubble Tea (tea.NewProgram(..., tea.WithAltScreen())). The --run-all flag is passed as autoRun into the TUI model. (internal/cli/run.go: runCmd.RunE; internal/tui/tui.go: NewModel, Init)

Browse/select steps (default “Steps” screen): Primary UI is list.Model (left) + viewport.Model (right preview). Selection changes update preview content via setListPreviewContent(newID) when not running and viewState == ViewSteps. (internal/tui/tui.go: Model.View, Update list handling, setListPreviewContent, selectedItemID)

Run a single selected step (“Enter”): From ViewSteps, pressing Enter sets running=true, viewState=ViewRunning, runAll=false, clears logLines, and triggers runStep(stepID) + waitForLogs() + spinner.Tick. Completion arrives as FinishedMsg{Err,...}; success sets running=false, viewState=ViewDone; failure sets err, running=false, runAll=false, viewState=ViewDone. (internal/tui/tui.go: Update key "enter", runStep, waitForLogs, FinishedMsg handler)

Run all filtered steps sequentially (“a”) + auto-run on init: Pressing a in ViewSteps (when not running) sets running=true, viewState=ViewRunning, runAll=true, currentIdx=0, selects first filtered list item, then executes runStep(firstID). On each FinishedMsg success while runAll=true, it increments currentIdx, selects next item, and calls runStep(nextID) until exhausted, then sets running=false, runAll=false, viewState=ViewDone. Auto-run uses the same path via StartAutoRunMsg emitted from Init() when autoRun is true. (internal/tui/tui.go: Update key "a", StartAutoRunMsg handling, FinishedMsg sequential logic; Init)

ROI filtering flow (“f” and “m”):

f enters filter input modal: isFiltering=true, filterInput.Focus(), view switches to centered modal via View() (still viewState remains Steps, but UI is driven by isFiltering). Enter parses float, updates roiThreshold, calls refreshList(), and persists into state (m.state.ROIThreshold, m.state.ROIMode) via state.SaveStateAtomic. Esc cancels.

m toggles roiMode between "under" and "over", refreshes list, persists state.
This directly changes what “run all” iterates over because it uses m.list.Items() (filtered). (internal/tui/tui.go: Update keys "f", "m", refreshList, persistFilterState, View modal branches)

Step preview (“v”) with wrap/copy/back: v sets previewID, viewState=ViewStepPreview, resets viewport offset, and renders Markdown (render.RenderMarkdown) into the viewport. In preview: t toggles wrap (previewWrap) and re-renders; c copies step plain text (clipboard fallback writes to file); b/esc returns to ViewSteps and restores list preview. (internal/tui/tui.go: Update key "v", `ViewSte
