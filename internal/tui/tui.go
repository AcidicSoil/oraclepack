package tui

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/state"
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
	pack        *pack.Pack
	runner      *exec.Runner
	state       *state.RunState
	
	width    int
	height   int
	running  bool
	runAll   bool // State for sequential execution
	currentIdx int
	autoRun  bool // Config to auto-start on init
	
	// Filtering state
	allSteps     []item // Store all items to support dynamic filtering
	roiThreshold float64
	roiMode      string
	isFiltering  bool

	err      error
	logLines []string
	logChan  chan string
}

func NewModel(p *pack.Pack, r *exec.Runner, s *state.RunState, roiThreshold float64, roiMode string, autoRun bool) Model {
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
	vp.SetContent("Press Enter to run selected, 'a' to run all filtered steps, 'f' to filter.")

	m := Model{
		list:         l,
		viewport:     vp,
		spinner:      sp,
		filterInput:  ti,
		pack:         p,
	runner:       r,
		state:        s,
		autoRun:      autoRun,
		allSteps:     allItems,
		roiThreshold: roiThreshold,
		roiMode:      roiMode,
		logChan:      make(chan string, 100),
	}

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
type FinishedMsg struct{ Err error; ID string }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// Handle global keys first if not filtering
	if !m.isFiltering {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "enter":
				if !m.running {
					i, ok := m.list.SelectedItem().(item)
					if ok {
						m.running = true
						m.runAll = false
						m.logLines = nil
						m.viewport.SetContent("Starting execution...")
						return m, tea.Batch(m.runStep(i.id), m.waitForLogs(), m.spinner.Tick)
					}
				}
			case "a":
				if !m.running && len(m.list.Items()) > 0 {
					m.running = true
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
					// Reset input to current threshold
					m.filterInput.SetValue(fmt.Sprintf("%.1f", m.roiThreshold))
					return m, textinput.Blink
				}
			}
		}
	} else {
		// Handling filter input
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				// Apply filter
				var val float64
				_, err := fmt.Sscanf(m.filterInput.Value(), "%f", &val)
				if err == nil {
					m.roiThreshold = val
					m = m.refreshList()
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

	switch msg := msg.(type) {
	case StartAutoRunMsg:
		if !m.running && len(m.list.Items()) > 0 {
			m.running = true
			m.runAll = true
			m.currentIdx = 0
			m.logLines = nil
			m.list.Select(0)
			// Trigger first step
			i := m.list.Items()[0].(item)
			m.viewport.SetContent(fmt.Sprintf("Auto-running all steps (Step 1/%d)...", len(m.list.Items())))
			return m, tea.Batch(m.runStep(i.id), m.waitForLogs(), m.spinner.Tick)
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetSize(msg.Width/3, msg.Height-4)
		m.viewport.Width = msg.Width - (msg.Width / 3) - 6
		m.viewport.Height = msg.Height - 4

	case LogMsg:
		m.logLines = append(m.logLines, string(msg))
		m.viewport.SetContent(strings.Join(m.logLines, "\n"))
		m.viewport.GotoBottom()
		return m, m.waitForLogs()

	case FinishedMsg:
		if msg.Err != nil {
			m.err = msg.Err
			m.logLines = append(m.logLines, fmt.Sprintf("\n❌ ERROR: %v", msg.Err))
			m.running = false
			m.runAll = false
		} else {
			m.logLines = append(m.logLines, "\n✅ SUCCESS")
			
			if m.runAll {
				m.currentIdx++
				if m.currentIdx < len(m.list.Items()) {
					m.list.Select(m.currentIdx)
					i := m.list.Items()[m.currentIdx].(item)
					m.logLines = append(m.logLines, fmt.Sprintf("\n--- Starting Step %d/%d ---\\n", m.currentIdx+1, len(m.list.Items())))
					return m, m.runStep(i.id)
				} else {
					m.logLines = append(m.logLines, "\n🏁 ALL STEPS COMPLETED")
					m.running = false
					m.runAll = false
				}
			} else {
				m.running = false
			}
		}
		m.viewport.SetContent(strings.Join(m.logLines, "\n"))
		m.viewport.GotoBottom()
		
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	if !m.running && !m.isFiltering {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.width == 0 {
		return "Initializing..."
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

	left := m.list.View()
	right := m.viewport.View()

	if m.running {
		status := "Running..."
		if m.runAll {
			status = fmt.Sprintf("Running All (%d/%d)...", m.currentIdx+1, len(m.list.Items()))
		}
		right = m.spinner.View() + " " + status + "\n" + right
	} else {
		// Show filter status if active
		filterStatus := ""
		if m.roiThreshold > 0 {
			modeSym := ">="
			if m.roiMode == "under" {
				modeSym = "<"
			}
			filterStatus = fmt.Sprintf(" [Filter: ROI %s %.1f]", modeSym, m.roiThreshold)
		}
		if filterStatus != "" {
			right = lipgloss.NewStyle().Foreground(lipgloss.Color("63")).Render(filterStatus) + "\n" + right
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, left, " | ", right)
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