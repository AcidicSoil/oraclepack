package tui

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/render"
	"github.com/user/oraclepack/internal/state"
)

type ViewState int

const (
	ViewSteps ViewState = iota
	ViewRunning
	ViewDone
	ViewOverrides
	ViewStepPreview
)

type item struct {
	id     string
	title  string
	desc   string
	status state.Status
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	list        list.Model
	viewport    viewport.Model
	spinner     spinner.Model
	filterInput textinput.Model
	urlInput    URLInputModel
	urlPicker   URLPickerModel
	pack        *pack.Pack
	runner      *exec.Runner
	state       *state.RunState
	statePath   string

	width  int
	height int

	viewState     ViewState
	running       bool
	runAll        bool // State for sequential execution
	currentIdx    int
	autoRun       bool // Config to auto-start on init
	previewID     string
	previewWrap   bool
	previewNotice string

	// Filtering state
	allSteps     []item // Store all items to support dynamic filtering
	roiThreshold float64
	roiMode      string
	isFiltering  bool
	isEditingURL bool
	isPickingURL bool

	overridesFlow    OverridesFlowModel
	appliedOverrides *overrides.RuntimeOverrides
	chatGPTURL       string

	err      error
	logLines []string
	logChan  chan string
}

func NewModel(p *pack.Pack, r *exec.Runner, s *state.RunState, statePath string, roiThreshold float64, roiMode string, autoRun bool) Model {
	if s != nil {
		if s.ROIThreshold > 0 {
			roiThreshold = s.ROIThreshold
		}
		if s.ROIMode != "" {
			roiMode = s.ROIMode
		}
	}
	var allItems []item
	for _, step := range p.Steps {
		allItems = append(allItems, item{
			id:    step.ID,
			title: fmt.Sprintf("Step %s", step.ID),
			desc:  step.OriginalLine,
		})
	}

	ti := textinput.New()
	ti.Placeholder = "Enter ROI threshold (e.g. 2.5)"
	ti.CharLimit = 10
	ti.Width = 20

	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Oracle Pack Steps"

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	vp := viewport.New(0, 0)
	vp.SetContent("Press Enter to run selected, 'a' to run all filtered steps, 'f' to set ROI threshold, 'm' to toggle ROI mode, 'v' to view step, 'o' to configure overrides, 'u' for ChatGPT URL, 'U' to pick a saved URL.")

	projectPath := ProjectURLStorePath(statePath, p.Source)
	globalPath := GlobalURLStorePath()
	urlPicker := NewURLPickerModel(projectPath, globalPath)
	resolvedURL := r.ChatGPTURL
	if resolvedURL == "" {
		resolvedURL = urlPicker.DefaultURL()
	}
	if resolvedURL != "" {
		r.ChatGPTURL = resolvedURL
	}

	m := Model{
		list:          l,
		viewport:      vp,
		spinner:       sp,
		filterInput:   ti,
		urlInput:      NewURLInputModel(),
		urlPicker:     urlPicker,
		pack:          p,
		runner:        r,
		state:         s,
		statePath:     statePath,
		autoRun:       autoRun,
		allSteps:      allItems,
		roiThreshold:  roiThreshold,
		roiMode:       roiMode,
		logChan:       make(chan string, 100),
		viewState:     ViewSteps,
		overridesFlow: NewOverridesFlowModel(p.Steps, r.OracleFlags, RunnerOptionsFromRunner(r)),
		chatGPTURL:    resolvedURL,
		previewWrap:   true,
	}
	m.urlInput.SetValue(resolvedURL)
	m.urlInput.Blur()

	// Apply initial filter
	return m.refreshList()
}

func (m Model) refreshList() Model {
	var filtered []list.Item
	for _, it := range m.allSteps {
		// Find the original step to check ROI
		var step *pack.Step
		for _, s := range m.pack.Steps {
			if s.ID == it.id {
				step = &s
				break
			}
		}
		if step == nil {
			continue
		}

		if m.roiThreshold > 0 {
			if m.roiMode == "under" {
				if step.ROI >= m.roiThreshold {
					continue
				}
			} else {
				if step.ROI < m.roiThreshold {
					continue
				}
			}
		}
		filtered = append(filtered, it)
	}
	m.list.SetItems(filtered)
	return m
}

type StartAutoRunMsg struct{}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	if m.autoRun {
		cmds = append(cmds, func() tea.Msg { return StartAutoRunMsg{} })
	}
	cmds = append(cmds, textinput.Blink)
	return tea.Batch(cmds...)
}

type LogMsg string
type FinishedMsg struct {
	Err error
	ID  string
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// Global keys (Quit)
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	if msg, ok := msg.(OverridesStartedMsg); ok {
		_ = msg
		m.viewState = ViewOverrides
		m.overridesFlow = NewOverridesFlowModel(m.pack.Steps, m.runner.OracleFlags, RunnerOptionsFromRunner(m.runner))
		return m, nil
	}
	if msg, ok := msg.(OverridesAppliedMsg); ok {
		over := msg.Overrides
		m.appliedOverrides = &over
		if m.runner != nil {
			m.runner.Overrides = &over
		}
		m.viewState = ViewSteps
		return m, nil
	}
	if msg, ok := msg.(OverridesCancelledMsg); ok {
		_ = msg
		m.appliedOverrides = nil
		if m.runner != nil {
			m.runner.Overrides = nil
		}
		m.viewState = ViewSteps
		return m, nil
	}
	if msg, ok := msg.(URLPickedMsg); ok {
		m.chatGPTURL = msg.URL
		if m.runner != nil {
			m.runner.ChatGPTURL = m.chatGPTURL
		}
		m.urlInput.SetValue(m.chatGPTURL)
		m.isPickingURL = false
		return m, nil
	}
	if _, ok := msg.(URLPickerCancelledMsg); ok {
		m.isPickingURL = false
		return m, nil
	}

	if m.viewState == ViewOverrides {
		var cmd tea.Cmd
		m.overridesFlow, cmd = m.overridesFlow.Update(msg)
		return m, cmd
	}

	if m.viewState == ViewStepPreview {
		switch msg := msg.(type) {
		case clearPreviewNoticeMsg:
			m.previewNotice = ""
			return m, nil
		case tea.KeyMsg:
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "b", "esc":
				m.previewID = ""
				m.previewNotice = ""
				m.viewState = ViewSteps
				m.setListPreviewContent(m.selectedItemID())
				return m, nil
			case "t":
				m.previewWrap = !m.previewWrap
				m.viewport.SetContent(m.stepPreviewContent())
				return m, nil
			case "c":
				content := m.stepPlainTextFor(m.previewID)
				if err := copyToClipboard(content); err != nil {
					path, fallbackErr := writeClipboardFallback(content)
					if fallbackErr != nil {
						m.previewNotice = "Copy failed: " + err.Error()
					} else {
						m.previewNotice = "Copy failed; saved to " + path
					}
				} else {
					m.previewNotice = "Copied to clipboard"
				}
				return m, tea.Tick(2*time.Second, func(time.Time) tea.Msg {
					return clearPreviewNoticeMsg{}
				})
			}
		}
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}

	switch m.viewState {
	case ViewDone:
		if msg, ok := msg.(tea.KeyMsg); ok {
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "b":
				m.viewState = ViewSteps
				m.setListPreviewContent(m.selectedItemID())
				return m, nil
			case "n":
				m.resetState()
				return m, nil
			case "r":
				// Rerun selected step (if we have one selected in list)
				// Or rerun the whole sequence if that was the context?
				// Requirement says "rerun a step ('r')". Assuming selected step.
				// We need to transition to ViewSteps logic or trigger run directly.
				m.viewState = ViewSteps // Go back to steps view? Or Running?
				// To trigger run, we can fall through or simulate Enter.
				// Let's just switch to steps and let user press Enter, or trigger run immediately?
				// "trigger the execution logic for the specific failed step"
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.running = true
					m.viewState = ViewRunning
					m.logLines = nil
					m.viewport.SetContent("Re-running execution...")
					return m, tea.Batch(m.runStep(i.id), m.waitForLogs(), m.spinner.Tick)
				}
			}
		}
		// In Done view, we might still want to handle window size?
		if msg, ok := msg.(tea.WindowSizeMsg); ok {
			m.handleWindowSize(msg)
		}
		return m, nil
	}

	// Filter Input Mode
	if m.isFiltering {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				var val float64
				_, err := fmt.Sscanf(m.filterInput.Value(), "%f", &val)
				if err == nil {
					m.roiThreshold = val
					m = m.refreshList()
					m.persistFilterState()
				}
				m.isFiltering = false
				m.filterInput.Blur()
				return m, nil
			case "esc":
				m.isFiltering = false
				m.filterInput.Blur()
				return m, nil
			}
		}
		var cmd tea.Cmd
		m.filterInput, cmd = m.filterInput.Update(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	}

	// ChatGPT URL Input Mode
	if m.isEditingURL {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				if m.urlInput.IsValid() {
					m.chatGPTURL = m.urlInput.Value()
					if m.runner != nil {
						m.runner.ChatGPTURL = m.chatGPTURL
					}
					m.isEditingURL = false
					m.urlInput.Blur()
					return m, nil
				}
			case "esc":
				m.isEditingURL = false
				m.urlInput.Blur()
				return m, nil
			}
		}
		var cmd tea.Cmd
		m.urlInput, cmd = m.urlInput.Update(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	}

	// URL Picker Mode
	if m.isPickingURL {
		var cmd tea.Cmd
		m.urlPicker, cmd = m.urlPicker.Update(msg)
		return m, cmd
	}

	// Normal Steps View / Running
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			if !m.running {
				return m, tea.Quit
			}
		case "enter":
			if !m.running {
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.running = true
					m.viewState = ViewRunning
					m.runAll = false
					m.logLines = nil
					m.viewport.SetContent("Starting execution...")
					return m, tea.Batch(m.runStep(i.id), m.waitForLogs(), m.spinner.Tick)
				}
			}
		case "a":
			if !m.running && len(m.list.Items()) > 0 {
				m.running = true
				m.viewState = ViewRunning
				m.runAll = true
				m.currentIdx = 0
				m.logLines = nil
				m.list.Select(0)
				i := m.list.Items()[0].(item)
				m.viewport.SetContent(fmt.Sprintf("Starting sequential run (Step 1/%d)...", len(m.list.Items())))
				return m, tea.Batch(m.runStep(i.id), m.waitForLogs(), m.spinner.Tick)
			}
		case "f":
			if !m.running {
				m.isFiltering = true
				m.filterInput.Focus()
				m.filterInput.SetValue(fmt.Sprintf("%.1f", m.roiThreshold))
				return m, textinput.Blink
			}
		case "m":
			if !m.running {
				if m.roiMode == "under" {
					m.roiMode = "over"
				} else {
					m.roiMode = "under"
				}
				m = m.refreshList()
				m.persistFilterState()
				return m, nil
			}
		case "v":
			if !m.running {
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.previewID = i.id
					m.viewState = ViewStepPreview
					m.viewport.YOffset = 0
					m.viewport.SetContent(m.stepPreviewContent())
					return m, nil
				}
			}
		case "u":
			if !m.running {
				m.isEditingURL = true
				m.urlInput.SetValue(m.chatGPTURL)
				m.urlInput.Focus()
				return m, textinput.Blink
			}
		case "U":
			if !m.running {
				m.isPickingURL = true
				return m, nil
			}
		case "o":
			if !m.running {
				return m, func() tea.Msg { return OverridesStartedMsg{} }
			}
		}

	case StartAutoRunMsg:
		if !m.running && len(m.list.Items()) > 0 {
			m.running = true
			m.viewState = ViewRunning
			m.runAll = true
			m.currentIdx = 0
			m.logLines = nil
			m.list.Select(0)
			i := m.list.Items()[0].(item)
			m.viewport.SetContent(fmt.Sprintf("Auto-running all steps (Step 1/%d)...", len(m.list.Items())))
			return m, tea.Batch(m.runStep(i.id), m.waitForLogs(), m.spinner.Tick)
		}

	case tea.WindowSizeMsg:
		m.handleWindowSize(msg)

	case LogMsg:
		m.logLines = append(m.logLines, string(msg))
		m.viewport.SetContent(strings.Join(m.logLines, "\n"))
		m.viewport.GotoBottom()
		return m, m.waitForLogs()

	case FinishedMsg:
		m.recordWarnings()
		if msg.Err != nil {
			m.err = msg.Err
			m.logLines = append(m.logLines, fmt.Sprintf("\n❌ ERROR: %v", msg.Err))
			m.running = false
			m.runAll = false
			m.viewState = ViewDone // Or stay in steps? Requirement says ViewDone on completion?
			// If error, maybe stay on steps or go to done with error?
			// "Failed at step X" is a summary state.
			m.viewState = ViewDone
		} else {
			m.logLines = append(m.logLines, "\n✅ SUCCESS")

			if m.runAll {
				m.currentIdx++
				if m.currentIdx < len(m.list.Items()) {
					m.list.Select(m.currentIdx)
					i := m.list.Items()[m.currentIdx].(item)
					m.logLines = append(m.logLines, fmt.Sprintf("\n--- Starting Step %d/%d ---\n", m.currentIdx+1, len(m.list.Items())))
					return m, m.runStep(i.id)
				} else {
					m.logLines = append(m.logLines, "\n🏁 ALL STEPS COMPLETED")
					m.running = false
					m.runAll = false
					m.viewState = ViewDone
				}
			} else {
				m.running = false
				m.viewState = ViewDone // Single step done
			}
		}
		m.viewport.SetContent(strings.Join(m.logLines, "\n"))
		m.viewport.GotoBottom()

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	if !m.running && !m.isFiltering && m.viewState == ViewSteps {
		prevID := m.selectedItemID()
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
		newID := m.selectedItemID()
		if newID != "" && newID != prevID {
			m.viewport.YOffset = 0
			m.setListPreviewContent(newID)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) handleWindowSize(msg tea.WindowSizeMsg) {
	m.width = msg.Width
	m.height = msg.Height
	if m.viewState == ViewStepPreview {
		m.viewport.Width = msg.Width - 4
		m.viewport.Height = msg.Height - 6
		if m.viewport.Height < 1 {
			m.viewport.Height = 1
		}
		m.viewport.SetContent(m.stepPreviewContent())
		m.viewport.GotoTop()
		return
	}
	contentHeight := msg.Height - 5
	if contentHeight < 1 {
		contentHeight = 1
	}
	m.list.SetSize(msg.Width/3, contentHeight)
	m.viewport.Width = msg.Width - (msg.Width / 3) - 6
	m.viewport.Height = contentHeight
	if !m.running && m.viewState == ViewSteps {
		m.setListPreviewContent(m.selectedItemID())
	}
}

func (m *Model) resetState() {
	// Reset RunState
	m.state.StartTime = time.Now()
	m.state.StepStatuses = make(map[string]state.StepStatus)

	// Save cleared state to disk
	if m.statePath != "" {
		_ = state.SaveStateAtomic(m.statePath, m.state)
	}

	// Reset UI
	m.logLines = nil
	m.viewport.SetContent("State reset. Ready for new run.")
	m.list.Select(0)
	m.viewState = ViewSteps
	m.running = false
	m.runAll = false
	m.appliedOverrides = nil
	if m.runner != nil {
		m.runner.Overrides = nil
	}
}

func (m *Model) persistFilterState() {
	if m.state == nil || m.statePath == "" {
		return
	}
	m.state.ROIThreshold = m.roiThreshold
	m.state.ROIMode = m.roiMode
	_ = state.SaveStateAtomic(m.statePath, m.state)
}

func (m *Model) recordWarnings() {
	if m.state == nil || m.statePath == "" || m.runner == nil {
		return
	}
	warnings := m.runner.DrainWarnings()
	if len(warnings) == 0 {
		return
	}
	for _, w := range warnings {
		m.state.Warnings = append(m.state.Warnings, state.Warning{
			Scope:   w.Scope,
			StepID:  w.StepID,
			Line:    w.Line,
			Token:   w.Token,
			Message: w.Message,
		})
	}
	_ = state.SaveStateAtomic(m.statePath, m.state)
}

func (m *Model) setLogContent() {
	if len(m.logLines) == 0 {
		return
	}
	m.viewport.SetContent(strings.Join(m.logLines, "\n"))
	m.viewport.GotoBottom()
}

func (m *Model) stepPreviewContent() string {
	return m.stepPreviewContentFor(m.previewID)
}

func (m *Model) stepPreviewContentFor(id string) string {
	md, ok := m.stepMarkdownFor(id)
	if !ok {
		return md
	}
	width := m.previewRenderWidth()
	rendered, err := render.RenderMarkdown(md, width, "auto")
	if err != nil {
		return m.stepPlainTextFor(id)
	}
	return rendered
}

func (m *Model) stepMarkdownFor(id string) (string, bool) {
	if id == "" {
		return "No step selected.", false
	}
	step := m.stepForID(id)
	if step == nil {
		return "Step not found.", false
	}
	header := fmt.Sprintf("## Step %s\n%s\n\n", step.ID, step.OriginalLine)
	md := header + "```bash\n" + step.Code + "\n```\n"
	return md, true
}

func (m *Model) stepPlainTextFor(id string) string {
	if id == "" {
		return "No step selected."
	}
	step := m.stepForID(id)
	if step == nil {
		return "Step not found."
	}
	header := fmt.Sprintf("Step %s\n%s\n", step.ID, step.OriginalLine)
	return header + "\n" + step.Code
}

func (m *Model) stepForID(id string) *pack.Step {
	for i := range m.pack.Steps {
		if m.pack.Steps[i].ID == id {
			return &m.pack.Steps[i]
		}
	}
	return nil
}

func (m *Model) previewRenderWidth() int {
	width := m.viewport.Width
	if width <= 0 {
		width = render.DefaultWidth
	}
	if !m.previewWrap {
		if width < render.DefaultWidth {
			width = render.DefaultWidth
		}
		width = width * 4
	}
	return width
}

func (m *Model) selectedItemID() string {
	it, ok := m.list.SelectedItem().(item)
	if !ok {
		return ""
	}
	return it.id
}

func (m *Model) setListPreviewContent(id string) {
	if id == "" {
		m.viewport.SetContent("No step selected.")
		return
	}
	m.viewport.SetContent(m.stepPreviewContentFor(id))
	m.viewport.GotoTop()
}

type clearPreviewNoticeMsg struct{}

func (m Model) View() string {
	if m.width == 0 {
		return "Initializing..."
	}

	if m.viewState == ViewDone {
		return m.viewDone()
	}

	if m.viewState == ViewOverrides {
		return m.overridesFlow.View(m.width, m.height)
	}

	if m.isFiltering {
		return lipgloss.Place(m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			lipgloss.JoinVertical(lipgloss.Center,
				"Enter ROI Threshold:",
				m.filterInput.View(),
				"(Enter to apply, Esc to cancel)",
			),
		)
	}

	if m.isEditingURL {
		return lipgloss.Place(m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			lipgloss.JoinVertical(lipgloss.Center,
				"ChatGPT URL (browser mode):",
				m.urlInput.View(),
				"(Enter to apply, Esc to cancel)",
			),
		)
	}

	if m.isPickingURL {
		m.urlPicker.SetSize(m.width-4, m.height-4)
		return lipgloss.Place(m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			m.urlPicker.View(),
		)
	}

	if m.viewState == ViewStepPreview {
		m.viewport.Width = m.width - 4
		m.viewport.Height = m.height - 6
		title := lipgloss.NewStyle().Bold(true).Render("Step Preview")
		help := "[b] Back  [q] Quit  [t] Wrap  [c] Copy  (scroll with ↑↓ / PgUp/PgDn)"
		notice := ""
		if m.previewNotice != "" {
			notice = lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Render(m.previewNotice)
		}
		content := lipgloss.JoinVertical(lipgloss.Left,
			title,
			help,
			notice,
			"",
			m.viewport.View(),
		)
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
	}

	left := m.list.View()
	right := m.viewport.View()

	if m.running {
		status := "Running..."
		if m.runAll {
			status = fmt.Sprintf("Running All (%d/%d)...", m.currentIdx+1, len(m.list.Items()))
		}
		right = m.spinner.View() + " " + status + "\n" + right
	} else {
		filterStatus := ""
		if m.roiThreshold > 0 {
			modeSym := ">="
			if m.roiMode == "under" {
				modeSym = "<"
			}
			filterStatus = fmt.Sprintf(" [Filter: ROI %s %.1f]", modeSym, m.roiThreshold)
		}
		if filterStatus == "" {
			modeSym := ">="
			if m.roiMode == "under" {
				modeSym = "<"
			}
			filterStatus = fmt.Sprintf(" [Filter: ROI %s ∞]", modeSym)
		}
		overrideStatus := ""
		if m.appliedOverrides != nil {
			added := len(m.appliedOverrides.AddedFlags)
			removed := len(m.appliedOverrides.RemovedFlags)
			targeted := len(m.appliedOverrides.ApplyToSteps)
			overrideStatus = fmt.Sprintf(" [Overrides: +%d -%d steps:%d]", added, removed, targeted)
		}
		urlStatus := ""
		if m.chatGPTURL != "" {
			urlStatus = " [ChatGPT URL: set]"
		} else {
			urlStatus = " [ChatGPT URL: none]"
		}
		statusLine := strings.TrimSpace(filterStatus + overrideStatus + urlStatus)
		if statusLine != "" {
			right = lipgloss.NewStyle().Foreground(lipgloss.Color("63")).Render(statusLine) + "\n" + right
		}
	}

	help := m.stepsHelpBar(m.width)
	rightWidth := m.viewport.Width
	if rightWidth < 1 {
		rightWidth = 1
	}
	right = lipgloss.NewStyle().Width(rightWidth).Render(right)
	main := lipgloss.JoinHorizontal(lipgloss.Top, left, " | ", right)
	return lipgloss.JoinVertical(lipgloss.Left, main, help)
}

func (m Model) viewDone() string {
	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("42")).Render("Execution Complete")
	if m.err != nil {
		title = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("196")).Render("Execution Failed")
	}

	help := "[n] New Run  [r] Rerun  [b] Back to List  [q] Quit  [m] ROI Mode"

	// Show the log viewport in the done screen too? Or just a summary?
	// Requirement says "displays a summary".
	// But viewing the logs is useful.
	// I'll show the viewport in the center/bottom.

	content := lipgloss.JoinVertical(lipgloss.Center,
		title,
		"",
		m.viewport.View(),
		"",
		help,
	)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}

func (m Model) stepsHelpBar(width int) string {
	help := "[enter] run  [a] run all  [f] filter ROI  [m] ROI mode  [v] view  [o] overrides  [u] url  [U] url picker  [q] quit"
	if m.running {
		help = "[q] quit  [running] wait for completion"
	}
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	if width > 0 {
		style = style.Width(width)
	}
	return style.Render(help)
}

func (m Model) waitForLogs() tea.Cmd {
	return func() tea.Msg {
		line, ok := <-m.logChan
		if !ok {
			return nil
		}
		return LogMsg(line)
	}
}

func (m Model) runStep(id string) tea.Cmd {
	return func() tea.Msg {
		var step *pack.Step
		for _, s := range m.pack.Steps {
			if s.ID == id {
				step = &s
				break
			}
		}

		if step == nil {
			return FinishedMsg{Err: fmt.Errorf("step not found"), ID: id}
		}

		lw := &exec.LineWriter{
			Callback: func(line string) {
				m.logChan <- line
			},
		}

		err := m.runner.RunStep(context.Background(), step, lw)
		lw.Close()
		return FinishedMsg{Err: err, ID: id}
	}
}

func RunnerOptionsFromRunner(r *exec.Runner) exec.RunnerOptions {
	if r == nil {
		return exec.RunnerOptions{}
	}
	return exec.RunnerOptions{
		Shell:       r.Shell,
		WorkDir:     r.WorkDir,
		Env:         r.Env,
		OracleFlags: r.OracleFlags,
		Overrides:   r.Overrides,
		ChatGPTURL:  r.ChatGPTURL,
	}
}
