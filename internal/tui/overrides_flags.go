package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FlagItem struct {
	Flag       string
	Desc       string
	IsBaseline bool
	Selected   bool
}

func (i FlagItem) Title() string       { return i.Flag }
func (i FlagItem) Description() string { return i.Desc }
func (i FlagItem) FilterValue() string { return i.Flag }

type FlagsPickerModel struct {
	list list.Model
}

func NewFlagsPickerModel(baseline []string) FlagsPickerModel {
	baselineSet := make(map[string]bool, len(baseline))
	for _, f := range baseline {
		baselineSet[f] = true
	}

	curated := []FlagItem{
		{Flag: "--files-report", Desc: "Show per-file token usage"},
		{Flag: "--render", Desc: "Print assembled markdown bundle"},
		{Flag: "--render-plain", Desc: "Render markdown without ANSI"},
		{Flag: "--copy", Desc: "Copy assembled markdown bundle"},
		{Flag: "--wait", Desc: "Wait for background API runs"},
	}

	items := make([]list.Item, 0, len(curated))
	for _, c := range curated {
		c.IsBaseline = baselineSet[c.Flag]
		if c.IsBaseline {
			c.Selected = true
		}
		items = append(items, c)
	}

	delegate := newFlagsDelegate()
	l := list.New(items, delegate, 0, 0)
	l.Title = "Oracle Flags"
	l.SetFilteringEnabled(true)

	return FlagsPickerModel{list: l}
}

func (m FlagsPickerModel) Init() tea.Cmd {
	return nil
}

func (m FlagsPickerModel) Update(msg tea.Msg) (FlagsPickerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == " " {
			idx := m.list.Index()
			item, ok := m.list.SelectedItem().(FlagItem)
			if ok && !item.IsBaseline {
				item.Selected = !item.Selected
				_ = m.list.SetItem(idx, item)
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *FlagsPickerModel) SetSize(width, height int) {
	m.list.SetSize(width, height)
}

func (m FlagsPickerModel) View() string {
	return m.list.View()
}

func (m FlagsPickerModel) SelectedFlags() []string {
	var flags []string
	for _, item := range m.list.Items() {
		if fi, ok := item.(FlagItem); ok && fi.Selected && !fi.IsBaseline {
			flags = append(flags, fi.Flag)
		}
	}
	return flags
}

type flagsDelegate struct {
	list.DefaultDelegate
}

func newFlagsDelegate() flagsDelegate {
	d := list.NewDefaultDelegate()
	return flagsDelegate{DefaultDelegate: d}
}

func (d flagsDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	fi, ok := item.(FlagItem)
	if !ok {
		d.DefaultDelegate.Render(w, m, index, item)
		return
	}

	checked := fi.Selected || fi.IsBaseline
	marker := "[ ]"
	if checked {
		marker = "[x]"
	}
	if fi.IsBaseline {
		marker = "[*]"
	}

	label := fi.Flag
	if fi.Desc != "" {
		label = fmt.Sprintf("%s - %s", fi.Flag, fi.Desc)
	}
	if fi.IsBaseline {
		label = label + " (base)"
	}

	line := fmt.Sprintf("%s %s", marker, label)
	if index == m.Index() {
		line = d.Styles.SelectedTitle.Render(line)
	} else {
		line = d.Styles.NormalTitle.Render(line)
	}
	if fi.IsBaseline {
		line = lipgloss.NewStyle().Faint(true).Render(line)
	}

	fmt.Fprintln(w, line)
}
