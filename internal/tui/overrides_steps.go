package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/oraclepack/internal/pack"
)

type StepItem struct {
	ID       string
	TitleTxt string
	DescTxt  string
	Selected bool
}

func (i StepItem) Title() string       { return i.TitleTxt }
func (i StepItem) Description() string { return i.DescTxt }
func (i StepItem) FilterValue() string { return i.TitleTxt }

type StepsPickerModel struct {
	list list.Model
}

func NewStepsPickerModel(steps []pack.Step) StepsPickerModel {
	items := make([]list.Item, 0, len(steps))
	for _, s := range steps {
		items = append(items, StepItem{
			ID:       s.ID,
			TitleTxt: fmt.Sprintf("Step %s", s.ID),
			DescTxt:  s.OriginalLine,
			Selected: true,
		})
	}

	delegate := newStepsDelegate()
	l := list.New(items, delegate, 0, 0)
	l.Title = "Target Steps"
	l.SetFilteringEnabled(true)

	return StepsPickerModel{list: l}
}

func (m StepsPickerModel) Init() tea.Cmd {
	return nil
}

func (m StepsPickerModel) Update(msg tea.Msg) (StepsPickerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == " " {
			idx := m.list.Index()
			item, ok := m.list.SelectedItem().(StepItem)
			if ok {
				item.Selected = !item.Selected
				_ = m.list.SetItem(idx, item)
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *StepsPickerModel) SetSize(width, height int) {
	m.list.SetSize(width, height)
}

func (m StepsPickerModel) View() string {
	return m.list.View()
}

func (m StepsPickerModel) SelectedSteps() map[string]bool {
	selected := make(map[string]bool)
	for _, item := range m.list.Items() {
		if si, ok := item.(StepItem); ok && si.Selected {
			selected[si.ID] = true
		}
	}
	return selected
}

type stepsDelegate struct {
	list.DefaultDelegate
}

func newStepsDelegate() stepsDelegate {
	d := list.NewDefaultDelegate()
	return stepsDelegate{DefaultDelegate: d}
}

func (d stepsDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	si, ok := item.(StepItem)
	if !ok {
		d.DefaultDelegate.Render(w, m, index, item)
		return
	}

	marker := "[ ]"
	if si.Selected {
		marker = "[x]"
	}

	label := si.TitleTxt
	if si.DescTxt != "" {
		label = fmt.Sprintf("%s - %s", si.TitleTxt, si.DescTxt)
	}

	line := fmt.Sprintf("%s %s", marker, label)
	if index == m.Index() {
		line = d.Styles.SelectedTitle.Render(line)
	} else {
		line = d.Styles.NormalTitle.Render(line)
	}

	fmt.Fprintln(w, line)
}
