package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type URLInputModel struct {
	input textinput.Model
	err   string
}

func NewURLInputModel() URLInputModel {
	ti := textinput.New()
	ti.Placeholder = "https://chat.openai.com/project/..."
	ti.CharLimit = 200
	ti.Width = 50

	return URLInputModel{input: ti}
}

func (m URLInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m URLInputModel) Update(msg tea.Msg) (URLInputModel, tea.Cmd) {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	m.err = ""
	if !m.IsValid() {
		m.err = "Invalid URL (must start with http:// or https://)"
	}
	return m, cmd
}

func (m URLInputModel) Value() string {
	return strings.TrimSpace(m.input.Value())
}

func (m URLInputModel) IsValid() bool {
	v := m.Value()
	if v == "" {
		return true
	}
	return strings.HasPrefix(v, "http://") || strings.HasPrefix(v, "https://")
}

func (m URLInputModel) View() string {
	body := m.input.View()
	if m.err != "" {
		body = body + "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render(m.err)
	}
	return body
}

func (m *URLInputModel) SetValue(v string) {
	m.input.SetValue(v)
}

func (m *URLInputModel) Focus() {
	m.input.Focus()
}

func (m *URLInputModel) Blur() {
	m.input.Blur()
}
