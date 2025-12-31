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
	list     list.Model
	viewport viewport.Model
	spinner  spinner.Model
	pack     *pack.Pack
	runner   *exec.Runner
	state    *state.RunState
	
	width    int
	height   int
	running  bool
	err      error
	logLines []string
	logChan  chan string
}

func NewModel(p *pack.Pack, r *exec.Runner, s *state.RunState, roiThreshold float64, roiMode string) Model {
	var items []list.Item
	for _, step := range p.Steps {
		// Filter by ROI
		if roiThreshold > 0 {
			if roiMode == "under" {
				if step.ROI >= roiThreshold {
					continue
				}
			} else {
				// Default to "over" (includes threshold)
				if step.ROI < roiThreshold {
					continue
				}
			}
		}

		items = append(items, item{
			id:    step.ID,
			title: fmt.Sprintf("Step %s", step.ID),
			desc:  step.OriginalLine,
		})
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Oracle Pack Steps"

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	vp := viewport.New(0, 0)
	vp.SetContent("Select a step and press Enter to run...")

	return Model{
		list:     l,
		viewport: vp,
		spinner:  sp,
		pack:     p,
		runner:   r,
		state:    s,
		logChan:  make(chan string, 100),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

type LogMsg string
type FinishedMsg struct{ Err error }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

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
					m.logLines = nil
					m.viewport.SetContent("Starting execution...")
					return m, tea.Batch(m.runStep(i.id), m.waitForLogs(), m.spinner.Tick)
				}
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
		} else {
			m.logLines = append(m.logLines, "\nSUCCESS")
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
		right = m.spinner.View() + " Running...\n" + right
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