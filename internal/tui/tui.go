package tui

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
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
	pack        *pack.Pack
	runner      *exec.Runner
	state       *state.RunState
	
	width       int
	height      int
	running     bool
	autoAdvance bool // NEW: Tracks if we are running all steps sequentially
	err         error
	logLines    []string
	logChan     chan string
}

// NewModel creates a new TUI model.
func NewModel(p *pack.Pack, r *exec.Runner, s *state.RunState, autoAdvance bool) Model {
	items := make([]list.Item, len(p.Steps))
	for i, step := range p.Steps {
		items[i] = item{
			id:    step.ID,
			title: fmt.Sprintf("Step %s", step.ID),
			desc:  step.OriginalLine,
		}
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Oracle Pack Steps"

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	vp := viewport.New(0, 0)
	vp.SetContent("Select a step and press Enter to run...")

	return Model{
		list:        l,
		viewport:    vp,
		spinner:     sp,
		pack:        p,
		runner:      r,
		state:       s,
		autoAdvance: autoAdvance,
		logChan:     make(chan string, 100),
	}
}

func (m Model) Init() tea.Cmd {
	if m.autoAdvance {
		// If auto-advance is on at init, trigger the first step immediately
		// But wait, Update handles the logic better. 
		// We can return a specific command to trigger the run of the *first* selected item.
		// By default list selects the first item (index 0).
		
		// To be safe, let's just trigger a custom msg or rely on Update?
		// Update is only called on Msg.
		// Let's return a command that sends a "StartAutoRunMsg"
		return tea.Batch(m.spinner.Tick, func() tea.Msg { return StartAutoRunMsg{} })
	}
	return nil
}

type LogMsg string
type FinishedMsg struct{ Err error }
type StartAutoRunMsg struct{}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "a":
			// Toggle Auto Advance
			if !m.running {
				m.autoAdvance = true
				return m, func() tea.Msg { return StartAutoRunMsg{} }
			}
		case "enter":
			if !m.running {
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.running = true
					m.logLines = nil
					m.viewport.SetContent("Starting execution...")
					return m, tea.Batch(m.runStep(i.id), m.waitForLogs(), m.spinner.Tick)
				}
			}
		}

	case StartAutoRunMsg:
		if !m.running {
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.running = true
				m.autoAdvance = true // Ensure it's set
				m.logLines = nil
				m.viewport.SetContent(fmt.Sprintf("Auto-running Step %s...", i.id))
				return m, tea.Batch(m.runStep(i.id), m.waitForLogs(), m.spinner.Tick)
			}
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
		m.running = false
		if msg.Err != nil {
			m.err = msg.Err
			m.logLines = append(m.logLines, fmt.Sprintf("\nERROR: %v", msg.Err))
			m.autoAdvance = false // Stop on error
		} else {
			m.logLines = append(m.logLines, "\nSUCCESS")
			
			if m.autoAdvance {
				// Move to next item
				idx := m.list.Index()
				if idx < len(m.list.Items())-1 {
					m.list.Select(idx + 1)
					// Trigger run for next
					return m, func() tea.Msg { return StartAutoRunMsg{} }
				} else {
					m.autoAdvance = false // Done
					m.logLines = append(m.logLines, "\nAll steps completed.")
				}
			}
		}
		m.viewport.SetContent(strings.Join(m.logLines, "\n"))
		m.viewport.GotoBottom()
	
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	if !m.running {
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

	left := m.list.View()
	right := m.viewport.View()

	if m.running {
		status := "Running..."
		if m.autoAdvance {
			status = "Auto-Running..."
		}
		right = m.spinner.View() + " " + status + "\n" + right
	} else if m.autoAdvance {
		right = "Auto-Advance Active (Starting next step...)\n" + right
	} else {
		right = "[a] Run All | [Enter] Run Step | [q] Quit\n" + right
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, left, " | ", right)
}

func (m Model) waitForLogs() tea.Cmd {
	return func() tea.Msg {
		return LogMsg(<-m.logChan)
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
			return FinishedMsg{Err: fmt.Errorf("step not found")}
		}

		lw := &exec.LineWriter{
			Callback: func(line string) {
				m.logChan <- line
			},
		}
		
		err := m.runner.RunStep(context.Background(), step, lw)
		lw.Close()
		return FinishedMsg{Err: err}
	}
}