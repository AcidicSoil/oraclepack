package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type URLPickedMsg struct {
	URL string
}

type URLPickerCancelledMsg struct{}

type urlItem struct {
	name  string
	url   string
	scope string
}

func (i urlItem) Title() string       { return i.name }
func (i urlItem) Description() string { return fmt.Sprintf("%s • %s", i.scope, i.url) }
func (i urlItem) FilterValue() string { return i.name }

type URLPickerModel struct {
	list list.Model

	projectPath string
	globalPath  string
	project     URLStore
	global      URLStore

	editing   bool
	editName  textinput.Model
	editURL   textinput.Model
	editScope string
	editIdx   int
	editIsNew bool

	errMsg string
}

func NewURLPickerModel(projectPath, globalPath string) URLPickerModel {
	project, _ := LoadURLStore(projectPath)
	global, _ := LoadURLStore(globalPath)

	items := makeURLItems(project, global)
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "ChatGPT Project URLs"
	l.SetFilteringEnabled(true)

	name := textinput.New()
	name.Placeholder = "Name (e.g., Core Project)"
	name.CharLimit = 60
	name.Width = 40

	url := textinput.New()
	url.Placeholder = "https://chatgpt.com/g/.../project"
	url.CharLimit = 200
	url.Width = 60

	return URLPickerModel{
		list:        l,
		projectPath: projectPath,
		globalPath:  globalPath,
		project:     project,
		global:      global,
		editName:    name,
		editURL:     url,
	}
}

func (m URLPickerModel) Init() tea.Cmd {
	return nil
}

func (m URLPickerModel) Update(msg tea.Msg) (URLPickerModel, tea.Cmd) {
	if m.editing {
		return m.updateEdit(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg { return URLPickerCancelledMsg{} }
		case "enter":
			item, ok := m.list.SelectedItem().(urlItem)
			if !ok {
				return m, nil
			}
			m.touch(item)
			return m, func() tea.Msg { return URLPickedMsg{URL: item.url} }
		case "a":
			m.startEdit(urlScopeProject, "", "", true)
			return m, nil
		case "e":
			item, ok := m.list.SelectedItem().(urlItem)
			if !ok {
				return m, nil
			}
			m.startEdit(item.scope, item.name, item.url, false)
			return m, nil
		case "d":
			item, ok := m.list.SelectedItem().(urlItem)
			if !ok {
				return m, nil
			}
			m.delete(item)
			return m, nil
		case "s":
			item, ok := m.list.SelectedItem().(urlItem)
			if !ok {
				return m, nil
			}
			m.setDefault(item)
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *URLPickerModel) SetSize(width, height int) {
	m.list.SetSize(width, height-4)
}

func (m URLPickerModel) View() string {
	if m.editing {
		return m.editView()
	}

	help := lipgloss.NewStyle().Faint(true).Render("[enter] use  [a] add  [e] edit  [d] delete  [s] default  [esc] cancel")
	return m.list.View() + "\n" + help
}

func makeURLItems(project URLStore, global URLStore) []list.Item {
	var items []list.Item
	for _, it := range project.Items {
		items = append(items, urlItem{name: it.Name, url: it.URL, scope: urlScopeProject})
	}
	for _, it := range global.Items {
		items = append(items, urlItem{name: it.Name, url: it.URL, scope: urlScopeGlobal})
	}
	return items
}

func (m *URLPickerModel) refresh() {
	m.list.SetItems(makeURLItems(m.project, m.global))
}

func (m *URLPickerModel) touch(item urlItem) {
	store := m.storeFor(item.scope)
	if store == nil {
		return
	}
	for i := range store.Items {
		if store.Items[i].Name == item.name {
			store.Items[i].LastUsed = nowRFC3339()
			break
		}
	}
	_ = m.saveStores()
}

func (m *URLPickerModel) delete(item urlItem) {
	store := m.storeFor(item.scope)
	if store == nil {
		return
	}
	var out []URLItem
	for _, it := range store.Items {
		if it.Name == item.name {
			continue
		}
		out = append(out, it)
	}
	store.Items = out
	if store.Default == item.name {
		store.Default = ""
	}
	_ = m.saveStores()
	m.refresh()
}

func (m *URLPickerModel) setDefault(item urlItem) {
	store := m.storeFor(item.scope)
	if store == nil {
		return
	}
	store.Default = item.name
	_ = m.saveStores()
}

func (m *URLPickerModel) startEdit(scope, name, url string, isNew bool) {
	m.editing = true
	m.editScope = scope
	m.editIsNew = isNew
	m.editName.SetValue(name)
	m.editURL.SetValue(url)
	m.editName.Focus()
	m.editURL.Blur()
	m.errMsg = ""
}

func (m URLPickerModel) editView() string {
	scopeLabel := fmt.Sprintf("Scope: %s (g=global, p=project)", m.editScope)
	if m.globalPath == "" && m.projectPath != "" {
		scopeLabel = "Scope: project"
		m.editScope = urlScopeProject
	}
	if m.projectPath == "" && m.globalPath != "" {
		scopeLabel = "Scope: global"
		m.editScope = urlScopeGlobal
	}
	lines := []string{
		"Add / Edit ChatGPT URL",
		scopeLabel,
		"",
		"Name:",
		m.editName.View(),
		"",
		"URL:",
		m.editURL.View(),
		"",
		"[tab] switch field  [enter] save  [esc] cancel",
	}
	if m.errMsg != "" {
		lines = append(lines, "", lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render(m.errMsg))
	}
	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func (m URLPickerModel) updateEdit(msg tea.Msg) (URLPickerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.editing = false
			return m, nil
		case "tab":
			if m.editName.Focused() {
				m.editName.Blur()
				m.editURL.Focus()
			} else {
				m.editURL.Blur()
				m.editName.Focus()
			}
			return m, nil
		case "g":
			if m.globalPath != "" {
				m.editScope = urlScopeGlobal
			}
		case "p":
			if m.projectPath != "" {
				m.editScope = urlScopeProject
			}
		case "enter":
			name := strings.TrimSpace(m.editName.Value())
			url := strings.TrimSpace(m.editURL.Value())
			if name == "" || !isValidURL(url) {
				m.errMsg = "Name and a valid URL are required."
				return m, nil
			}
			m.saveEdit(name, url)
			m.editing = false
			m.refresh()
			return m, nil
		}
	}

	var cmd tea.Cmd
	if m.editName.Focused() {
		m.editName, cmd = m.editName.Update(msg)
	} else {
		m.editURL, cmd = m.editURL.Update(msg)
	}
	return m, cmd
}

func (m *URLPickerModel) saveEdit(name, url string) {
	scope := m.editScope
	if scope == "" {
		scope = urlScopeProject
	}

	// remove from other store if scope changed
	m.removeByName(name, urlScopeProject)
	m.removeByName(name, urlScopeGlobal)

	store := m.storeFor(scope)
	if store == nil {
		return
	}
	updated := false
	for i := range store.Items {
		if store.Items[i].Name == name {
			store.Items[i].URL = url
			store.Items[i].LastUsed = nowRFC3339()
			updated = true
			break
		}
	}
	if !updated {
		store.Items = append(store.Items, URLItem{Name: name, URL: url, LastUsed: nowRFC3339()})
	}
	_ = m.saveStores()
}

func (m *URLPickerModel) removeByName(name, scope string) {
	store := m.storeFor(scope)
	if store == nil {
		return
	}
	var out []URLItem
	for _, it := range store.Items {
		if it.Name == name {
			continue
		}
		out = append(out, it)
	}
	store.Items = out
}

func (m *URLPickerModel) storeFor(scope string) *URLStore {
	switch scope {
	case urlScopeProject:
		if m.projectPath == "" {
			return nil
		}
		return &m.project
	case urlScopeGlobal:
		if m.globalPath == "" {
			return nil
		}
		return &m.global
	default:
		return nil
	}
}

func (m *URLPickerModel) saveStores() error {
	if err := SaveURLStore(m.projectPath, m.project); err != nil {
		return err
	}
	if err := SaveURLStore(m.globalPath, m.global); err != nil {
		return err
	}
	return nil
}
