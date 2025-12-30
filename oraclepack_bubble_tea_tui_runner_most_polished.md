# oraclepack — Bubble Tea TUI runner (most polished)

## What this provides

- Bubble Tea TUI (Charmbracelet):
  - Left: steps list with status (Pending/Running/Done/Failed/Skipped)
  - Right: preview pane (command + detected `--write-output` path)
  - Bottom: live log output while the step runs
  - Confirmation modal before running each step (y/n)
  - Keys: ↑/↓ select, Enter run, s skip, r retry, p toggle preview/log focus, q quit
- Parses your pack markdown:
  - extracts the first ```bash fenced block
  - splits it into a *prelude* + numbered steps (`# 01)`, `# 02)`, ...)
- Runs each step via `bash -lc` with strict mode.
  - To preserve pack behavior, the prelude is prepended to each step execution (safe default).

---

## Project layout

```text
oraclepack/
  go.mod
  main.go
```

---

## Files

```go
// path: oraclepack/go.mod
module oraclepack

go 1.22

require (
	github.com/charmbracelet/bubbles v0.18.0
	github.com/charmbracelet/bubbletea v0.26.6
	github.com/charmbracelet/lipgloss v0.11.0
)
```

```go
// path: oraclepack/main.go
package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// -----------------------------
// Parsing
// -----------------------------

type StepStatus string

const (
	StatusPending StepStatus = "pending"
	StatusRunning StepStatus = "running"
	StatusDone    StepStatus = "done"
	StatusFailed  StepStatus = "failed"
	StatusSkipped StepStatus = "skipped"
)

type Step struct {
	Num        int        `json:"num"`
	Header     string     `json:"header"`
	Body       string     `json:"body"`
	OutputPath string     `json:"output_path"`
	Status     StepStatus `json:"status"`
	ExitCode   int        `json:"exit_code"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type Pack struct {
	PackPath string `json:"pack_path"`
	Prelude  string `json:"prelude"`
	Steps    []Step `json:"steps"`
}

var (
	reFenceStart = regexp.MustCompile(`(?m)^```bash\s*$`)
	reFenceEnd   = regexp.MustCompile(`(?m)^```\s*$`)
	reStepHeader = regexp.MustCompile(`(?m)^#\s*([0-9]{2})\)`)
	reWriteOut1  = regexp.MustCompile(`--write-output\s+"([^"]+)"`)
	reWriteOut2  = regexp.MustCompile(`--write-output=([^\s]+)`)
	reOracleLine = regexp.MustCompile(`(?m)^oracle(\s+\\)?\s*$`)
	reOracleAny  = regexp.MustCompile(`(?m)^oracle\b`)
)

func extractFirstBashFence(md string) (string, error) {
	loc := reFenceStart.FindStringIndex(md)
	if loc == nil {
		return "", errors.New("no ```bash fenced block found")
	}
	rest := md[loc[1]:]
	end := reFenceEnd.FindStringIndex(rest)
	if end == nil {
		return "", errors.New("unterminated fenced block")
	}
	code := rest[:end[0]]
	return strings.TrimSpace(code) + "\n", nil
}

func parsePack(packPath string, injectFlags string) (Pack, error) {
	b, err := os.ReadFile(packPath)
	if err != nil {
		return Pack{}, err
	}
	code, err := extractFirstBashFence(string(b))
	if err != nil {
		return Pack{}, fmt.Errorf("%w: %s", err, packPath)
	}

	if strings.TrimSpace(injectFlags) != "" {
		code = injectOracleFlags(code, injectFlags)
	}

	// Split into prelude + step blocks
	matches := reStepHeader.FindAllStringSubmatchIndex(code, -1)
	if len(matches) == 0 {
		return Pack{}, errors.New("no step headers like '# 01)' found inside bash fence")
	}

	prelude := code[:matches[0][0]]
	steps := make([]Step, 0, len(matches))

	for i, m := range matches {
		numStr := code[m[2]:m[3]]
		num := atoi(numStr)
		headerLine := firstLine(code[m[0]:])

		bodyStart := m[1]
		bodyEnd := len(code)
		if i+1 < len(matches) {
			bodyEnd = matches[i+1][0]
		}
		body := strings.TrimSpace(code[bodyStart:bodyEnd]) + "\n"

		steps = append(steps, Step{
			Num:        num,
			Header:     strings.TrimSpace(headerLine),
			Body:       body,
			OutputPath: detectWriteOutput(body),
			Status:     StatusPending,
			ExitCode:   0,
			UpdatedAt:  time.Now(),
		})
	}

	// Ensure steps are sorted by Num
	sort.Slice(steps, func(i, j int) bool { return steps[i].Num < steps[j].Num })

	return Pack{PackPath: packPath, Prelude: prelude, Steps: steps}, nil
}

func injectOracleFlags(code string, injectFlags string) string {
	injectFlags = strings.TrimSpace(injectFlags)
	if injectFlags == "" {
		return code
	}

	// Insert after a bare `oracle` line or `oracle \` line.
	lines := strings.Split(code, "\n")
	for i, ln := range lines {
		trim := strings.TrimSpace(ln)
		if trim == "oracle" {
			lines[i] = "oracle " + injectFlags
			continue
		}
		if trim == "oracle \\" || trim == "oracle\\" || trim == "oracle\\" {
			// best-effort; rare formatting
			lines[i] = "oracle " + injectFlags + " \\" // preserve continuation
			continue
		}
		if strings.HasPrefix(trim, "oracle \") {
			// `oracle \`
			lines[i] = "oracle " + injectFlags + " \\" // keep continued command
			continue
		}
	}
	return strings.Join(lines, "\n")
}

func detectWriteOutput(body string) string {
	if m := reWriteOut1.FindStringSubmatch(body); m != nil {
		return m[1]
	}
	if m := reWriteOut2.FindStringSubmatch(body); m != nil {
		return m[1]
	}
	return ""
}

func firstLine(s string) string {
	i := strings.IndexByte(s, '\n')
	if i < 0 {
		return s
	}
	return s[:i]
}

func atoi(s string) int {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			break
		}
		n = n*10 + int(r-'0')
	}
	return n
}

// -----------------------------
// State persistence (optional)
// -----------------------------

type State struct {
	PackPath string          `json:"pack_path"`
	Steps    map[int]Step    `json:"steps"`
	SavedAt  time.Time       `json:"saved_at"`
	Meta     map[string]any  `json:"meta,omitempty"`
	Version  int             `json:"version"`
}

func loadState(path string) (*State, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var st State
	if err := json.Unmarshal(b, &st); err != nil {
		return nil, err
	}
	if st.Steps == nil {
		st.Steps = map[int]Step{}
	}
	return &st, nil
}

func saveState(path string, pack Pack) error {
	st := State{
		PackPath: pack.PackPath,
		Steps:    map[int]Step{},
		SavedAt:  time.Now(),
		Version:  1,
	}
	for _, s := range pack.Steps {
		st.Steps[s.Num] = s
	}

	b, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

func applyState(pack *Pack, st *State) {
	if st == nil {
		return
	}
	for i := range pack.Steps {
		if ss, ok := st.Steps[pack.Steps[i].Num]; ok {
			pack.Steps[i].Status = ss.Status
			pack.Steps[i].ExitCode = ss.ExitCode
			pack.Steps[i].UpdatedAt = ss.UpdatedAt
		}
	}
}

// -----------------------------
// Runner (streaming)
// -----------------------------

type runState struct {
	ctx    context.Context
	cancel context.CancelFunc
	outCh  chan string
	doneCh chan runResult
}

type runResult struct {
	exitCode int
	err      error
}

type msgOut struct{ chunk string }

type msgDone struct{ res runResult }

type msgStartRun struct{}

type msgAbort struct{}

func startRun(prelude, body, cwd string) *runState {
	ctx, cancel := context.WithCancel(context.Background())
	rs := &runState{
		ctx:    ctx,
		cancel: cancel,
		outCh:  make(chan string, 256),
		doneCh: make(chan runResult, 1),
	}

	script := buildScript(prelude, body)

	go func() {
		defer close(rs.outCh)

		cmd := exec.CommandContext(ctx, "bash", "-lc", script)
		if cwd != "" {
			cmd.Dir = cwd
		}

		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		if err := cmd.Start(); err != nil {
			rs.doneCh <- runResult{exitCode: 127, err: err}
			return
		}

		// Stream stdout/stderr concurrently
		stream := func(r io.Reader, prefix string) {
			sc := bufio.NewScanner(r)
			buf := make([]byte, 0, 64*1024)
			sc.Buffer(buf, 1024*1024)
			for sc.Scan() {
				line := sc.Text()
				rs.outCh <- prefix + line
			}
			if err := sc.Err(); err != nil {
				rs.outCh <- prefix + "<scanner error: " + err.Error() + ">"
			}
		}

		doneStream := make(chan struct{}, 2)
		go func() { stream(stdout, ""); doneStream <- struct{}{} }()
		go func() { stream(stderr, "[stderr] "); doneStream <- struct{}{} }()

		// Wait for process
		err := cmd.Wait()
		// Wait for both streams to finish draining
		<-doneStream
		<-doneStream

		exit := 0
		if err != nil {
			exit = exitCode(err)
		}
		rs.doneCh <- runResult{exitCode: exit, err: err}
	}()

	return rs
}

func buildScript(prelude, body string) string {
	// Run prelude for each step to preserve variable definitions like out_dir.
	// This is the safest default for packs that expect shell variables.
	var b bytes.Buffer
	b.WriteString("set -euo pipefail\n")
	b.WriteString("\n")
	b.WriteString(strings.TrimSpace(prelude))
	b.WriteString("\n\n")
	b.WriteString(strings.TrimSpace(body))
	b.WriteString("\n")
	return b.String()
}

func exitCode(err error) int {
	var ee *exec.ExitError
	if errors.As(err, &ee) {
		if ee.ProcessState != nil {
			return ee.ProcessState.ExitCode()
		}
	}
	// Context cancellation
	if errors.Is(err, context.Canceled) {
		return 130
	}
	return 1
}

func readOutCmd(rs *runState) tea.Cmd {
	return func() tea.Msg {
		chunk, ok := <-rs.outCh
		if !ok {
			return msgOut{chunk: ""}
		}
		return msgOut{chunk: chunk}
	}
}

func waitDoneCmd(rs *runState) tea.Cmd {
	return func() tea.Msg {
		res := <-rs.doneCh
		return msgDone{res: res}
	}
}

// -----------------------------
// TUI
// -----------------------------

type item struct{
	step *Step
}

func (it item) Title() string {
	return fmt.Sprintf("%02d  %s", it.step.Num, statusBadge(it.step.Status))
}

func (it item) Description() string {
	h := it.step.Header
	if h == "" {
		h = fmt.Sprintf("# %02d)", it.step.Num)
	}
	if it.step.OutputPath != "" {
		return strings.TrimSpace(h) + "  →  " + it.step.OutputPath
	}
	return strings.TrimSpace(h)
}

func (it item) FilterValue() string {
	return fmt.Sprintf("%02d %s %s", it.step.Num, it.step.Header, it.step.OutputPath)
}

func statusBadge(s StepStatus) string {
	switch s {
	case StatusPending:
		return "PENDING"
	case StatusRunning:
		return "RUNNING"
	case StatusDone:
		return "DONE"
	case StatusFailed:
		return "FAILED"
	case StatusSkipped:
		return "SKIPPED"
	default:
		return strings.ToUpper(string(s))
	}
}

type keyMap struct {
	Up       key.Binding
	Down     key.Binding
	Run      key.Binding
	Skip     key.Binding
	Retry    key.Binding
	Preview  key.Binding
	Abort    key.Binding
	Quit     key.Binding
	ConfirmY key.Binding
	ConfirmN key.Binding
	Help     key.Binding
}

func defaultKeyMap() keyMap {
	return keyMap{
		Up:       key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "up")),
		Down:     key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "down")),
		Run:      key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "run")),
		Skip:     key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "skip")),
		Retry:    key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "retry")),
		Preview:  key.NewBinding(key.WithKeys("p"), key.WithHelp("p", "toggle focus")),
		Abort:    key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "abort")),
		Quit:     key.NewBinding(key.WithKeys("q", "esc"), key.WithHelp("q/esc", "quit")),
		ConfirmY: key.NewBinding(key.WithKeys("y"), key.WithHelp("y", "confirm")),
		ConfirmN: key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "cancel")),
		Help:     key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
	}
}

type focusMode int

const (
	focusPreview focusMode = iota
	focusLog
)

type uiMode int

const (
	modeNormal uiMode = iota
	modeConfirm
)

type model struct {
	packPath  string
	statePath string
	cwd       string

	pack Pack
	lst  list.Model
	help help.Model
	keys keyMap
	spin spinner.Model

	vpPreview viewport.Model
	vpLog     viewport.Model
	focus     focusMode
	mode      uiMode

	run       *runState
	logLines  []string
	maxLog    int
	confirmIx int

	width  int
	height int

	errMsg string
}

func newModel(pack Pack, statePath, cwd string) model {
	items := make([]list.Item, 0, len(pack.Steps))
	for i := range pack.Steps {
		s := &pack.Steps[i]
		items = append(items, item{step: s})
	}

	lst := list.New(items, list.NewDefaultDelegate(), 0, 0)
	lst.Title = "Oracle Pack"
	lst.SetShowHelp(false)
	lst.DisableQuitKeybindings()
	lst.SetFilteringEnabled(true)

	sp := spinner.New()
	sp.Spinner = spinner.Line

	m := model{
		packPath:  pack.PackPath,
		statePath: statePath,
		cwd:       cwd,
		pack:      pack,
		lst:       lst,
		help:      help.New(),
		keys:      defaultKeyMap(),
		spin:      sp,
		vpPreview: viewport.New(0, 0),
		vpLog:     viewport.New(0, 0),
		focus:     focusPreview,
		mode:      modeNormal,
		maxLog:    5000,
		confirmIx: -1,
	}
	m.refreshPreview()
	m.refreshLog()
	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spin.Tick)
}

func (m model) selectedIndex() int {
	ix := m.lst.Index()
	if ix < 0 || ix >= len(m.pack.Steps) {
		return -1
	}
	return ix
}

func (m *model) refreshPreview() {
	ix := m.selectedIndex()
	if ix < 0 {
		m.vpPreview.SetContent("")
		return
	}
	s := m.pack.Steps[ix]

	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s\n", s.Header))
	b.WriteString(fmt.Sprintf("Status: %s\n", strings.ToUpper(string(s.Status))))
	if s.OutputPath != "" {
		b.WriteString(fmt.Sprintf("Write-output: %s\n", s.OutputPath))
	}
	if s.Status == StatusFailed {
		b.WriteString(fmt.Sprintf("Last exit code: %d\n", s.ExitCode))
	}
	b.WriteString("\n")
	b.WriteString("Command:\n")
	b.WriteString(strings.TrimSpace(s.Body))
	b.WriteString("\n")

	m.vpPreview.SetContent(b.String())
}

func (m *model) refreshLog() {
	m.vpLog.SetContent(strings.Join(m.logLines, "\n"))
}

func (m *model) appendLog(line string) {
	if line == "" {
		return
	}
	m.logLines = append(m.logLines, line)
	if len(m.logLines) > m.maxLog {
		m.logLines = m.logLines[len(m.logLines)-m.maxLog:]
	}
	m.refreshLog()
	m.vpLog.GotoBottom()
}

func (m *model) setStepStatus(ix int, st StepStatus, exit int) {
	m.pack.Steps[ix].Status = st
	m.pack.Steps[ix].ExitCode = exit
	m.pack.Steps[ix].UpdatedAt = time.Now()
	// refresh list item (re-wrap pointer)
	m.lst.SetItem(ix, item{step: &m.pack.Steps[ix]})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.layout(msg.Width, msg.Height)
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spin, cmd = m.spin.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Quit) {
			_ = m.persistState()
			return m, tea.Quit
		}

		if m.run != nil && key.Matches(msg, m.keys.Abort) {
			m.run.cancel()
			m.appendLog("<aborting…>")
			return m, nil
		}

		if key.Matches(msg, m.keys.Help) {
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		}

		// Confirm modal
		if m.mode == modeConfirm {
			if key.Matches(msg, m.keys.ConfirmY) {
				ix := m.confirmIx
				m.mode = modeNormal
				m.confirmIx = -1
				if ix >= 0 {
					return m, m.beginRun(ix)
				}
				return m, nil
			}
			if key.Matches(msg, m.keys.ConfirmN) {
				m.mode = modeNormal
				m.confirmIx = -1
				return m, nil
			}
			// allow skip from confirm
			if key.Matches(msg, m.keys.Skip) {
				ix := m.confirmIx
				m.mode = modeNormal
				m.confirmIx = -1
				if ix >= 0 {
					m.setStepStatus(ix, StatusSkipped, 0)
					m.refreshPreview()
					_ = m.persistState()
				}
				return m, nil
			}
			return m, nil
		}

		// Normal mode
		if key.Matches(msg, m.keys.Preview) {
			if m.focus == focusPreview {
				m.focus = focusLog
			} else {
				m.focus = focusPreview
			}
			return m, nil
		}

		if key.Matches(msg, m.keys.Run) {
			ix := m.selectedIndex()
			if ix >= 0 {
				// Only allow run if not currently running
				if m.run == nil {
					m.mode = modeConfirm
					m.confirmIx = ix
				}
			}
			return m, nil
		}

		if key.Matches(msg, m.keys.Skip) {
			ix := m.selectedIndex()
			if ix >= 0 && m.run == nil {
				m.setStepStatus(ix, StatusSkipped, 0)
				m.refreshPreview()
				_ = m.persistState()
			}
			return m, nil
		}

		if key.Matches(msg, m.keys.Retry) {
			ix := m.selectedIndex()
			if ix >= 0 && m.run == nil {
				// confirm retry
				m.mode = modeConfirm
				m.confirmIx = ix
			}
			return m, nil
		}

		// Delegate selection keys to list
		var cmd tea.Cmd
		m.lst, cmd = m.lst.Update(msg)
		m.refreshPreview()
		return m, cmd

	case msgOut:
		// If outCh closed, msgOut with empty chunk arrives; just stop re-issuing read cmd.
		if msg.chunk != "" {
			m.appendLog(msg.chunk)
			return m, readOutCmd(m.run)
		}
		return m, nil

	case msgDone:
		ix := m.selectedIndex()
		// Best-effort: mark the step that was running (we store it in list selection at start)
		runningIx := findRunningIndex(m.pack.Steps)
		if runningIx >= 0 {
			if msg.res.err == nil {
				m.setStepStatus(runningIx, StatusDone, 0)
				m.appendLog("<done>")
			} else {
				m.setStepStatus(runningIx, StatusFailed, msg.res.exitCode)
				m.appendLog(fmt.Sprintf("<failed: exit %d>", msg.res.exitCode))
			}
		}

		m.run = nil
		m.refreshPreview()
		_ = m.persistState()

		// Keep selection stable
		if ix >= 0 {
			m.lst.Select(ix)
		}
		return m, nil

	default:
		return m, nil
	}
}

func findRunningIndex(steps []Step) int {
	for i := range steps {
		if steps[i].Status == StatusRunning {
			return i
		}
	}
	return -1
}

func (m *model) beginRun(ix int) tea.Cmd {
	m.errMsg = ""
	m.logLines = nil
	m.refreshLog()

	m.setStepStatus(ix, StatusRunning, 0)
	m.refreshPreview()
	_ = m.persistState()

	m.run = startRun(m.pack.Prelude, m.pack.Steps[ix].Body, m.cwd)

	// Start streaming + wait for completion.
	return tea.Batch(
		readOutCmd(m.run),
		waitDoneCmd(m.run),
	)
}

func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	left := m.lst.View()

	rightTop := m.vpPreview.View()
	rightBottom := m.vpLog.View()

	// Focus styling
	border := lipgloss.RoundedBorder()
	box := func(title string, focused bool, content string) string {
		style := lipgloss.NewStyle().Border(border).Padding(0, 1)
		if focused {
			style = style.BorderForeground(lipgloss.Color("6"))
		}
		return style.Render(lipgloss.NewStyle().Bold(true).Render(title) + "\n" + content)
	}

	previewFocused := m.focus == focusPreview
	logFocused := m.focus == focusLog

	if m.run != nil {
		rightTop = "" + m.spin.View() + " running…\n\n" + rightTop
	}

	right := lipgloss.JoinVertical(
		lipgloss.Top,
		box("Preview", previewFocused, rightTop),
		box("Log", logFocused, rightBottom),
	)

	main := lipgloss.JoinHorizontal(lipgloss.Top, left, right)

	footer := m.footerView()

	out := lipgloss.JoinVertical(lipgloss.Left, main, footer)

	if m.mode == modeConfirm {
		out = overlayCentered(out, m.confirmView(), m.width, m.height)
	}

	return out
}

func (m model) footerView() string {
	// Show help keys; disable run keys when running.
	keys := []key.Binding{m.keys.Run, m.keys.Skip, m.keys.Retry, m.keys.Preview, m.keys.Help, m.keys.Quit}
	if m.run != nil {
		keys = []key.Binding{m.keys.Abort, m.keys.Preview, m.keys.Help, m.keys.Quit}
	}
	return lipgloss.NewStyle().Padding(0, 1).Render(m.help.ShortHelpView(keys))
}

func (m model) confirmView() string {
	ix := m.confirmIx
	if ix < 0 || ix >= len(m.pack.Steps) {
		return ""
	}
	s := m.pack.Steps[ix]

	var b strings.Builder
	b.WriteString(fmt.Sprintf("Run step %02d?\n\n", s.Num))
	b.WriteString(strings.TrimSpace(s.Header) + "\n")
	if s.OutputPath != "" {
		b.WriteString("\nWrite-output:\n  " + s.OutputPath + "\n")
	}
	b.WriteString("\nKeys: y=run  n=cancel  s=skip")

	style := lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).Padding(1, 2).Width(min(80, m.width-6))
	return style.Render(b.String())
}

func overlayCentered(bg, fg string, w, h int) string {
	// Minimal overlay: place fg after some top padding.
	fgLines := strings.Split(fg, "\n")
	fgH := len(fgLines)
	topPad := max(0, (h-fgH)/2)
	pad := strings.Repeat("\n", topPad)
	return bg + "\n" + pad + fg
}

func (m *model) layout(w, h int) {
	// Reserve 1 row for footer.
	hMain := max(1, h-1)

	leftW := min(42, max(28, w/3))
	rightW := max(20, w-leftW)

	m.lst.SetSize(leftW, hMain)

	// Right side split
	topH := max(8, hMain/2)
	botH := max(4, hMain-topH)

	m.vpPreview.Width = max(10, rightW-4)
	m.vpPreview.Height = max(4, topH-3)

	m.vpLog.Width = max(10, rightW-4)
	m.vpLog.Height = max(4, botH-3)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (m *model) persistState() error {
	if strings.TrimSpace(m.statePath) == "" {
		return nil
	}
	return saveState(m.statePath, m.pack)
}

// -----------------------------
// main
// -----------------------------

func main() {
	var (
		packPath  string
		cwd       string
		statePath string
		inject    string
		startAt   int
		noState   bool
	)

	flag.StringVar(&packPath, "pack", "", "Path to pack markdown file (contains a ```bash fence)")
	flag.StringVar(&cwd, "cwd", "", "Working directory to run steps from (defaults to current dir)")
	flag.StringVar(&statePath, "state", "", "Path to JSON state file (default: .oraclepack-state.json next to pack)")
	flag.StringVar(&inject, "inject", "", "Flags to inject after 'oracle' (e.g. '--browser-attachments always')")
	flag.IntVar(&startAt, "start", 1, "Select step number to start on (best-effort selection)")
	flag.BoolVar(&noState, "no-state", false, "Disable state read/write")
	flag.Parse()

	if packPath == "" {
		exitErr("--pack is required")
	}

	absPack, err := filepath.Abs(packPath)
	if err != nil {
		exitErr(err.Error())
	}
	packPath = absPack

	if cwd == "" {
		cwd, _ = os.Getwd()
	}

	if noState {
		statePath = ""
	} else if statePath == "" {
		statePath = filepath.Join(filepath.Dir(packPath), ".oraclepack-state.json")
	}

	pack, err := parsePack(packPath, inject)
	if err != nil {
		exitErr(err.Error())
	}

	// Load state if present
	var st *State
	if statePath != "" {
		if s, err := loadState(statePath); err == nil {
			st = s
		}
	}
	applyState(&pack, st)

	m := newModel(pack, statePath, cwd)
	m.packPath = packPath

	// Best-effort: select start step
	if startAt > 1 {
		m.selectByStepNum(startAt)
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		exitErr(err.Error())
	}
}

func (m *model) selectByStepNum(n int) {
	for i := range m.pack.Steps {
		if m.pack.Steps[i].Num >= n {
			m.lst.Select(i)
			m.refreshPreview()
			return
		}
	}
}

func exitErr(msg string) {
	fmt.Fprintln(os.Stderr, "error:", msg)
	os.Exit(1)
}

// -----------------------------
// Notes
// -----------------------------
// - This runner executes bash from the pack. Treat packs as trusted input.
// - It prepends the prelude to each step execution so shell variables (e.g. out_dir)
//   remain available even though each step runs in its own bash process.
// - If you want a stricter mode, change buildScript() to run the prelude once and
//   export variables; that requires pack conventions (export out_dir=...).
```

---

## Build & run

```bash
# from the oraclepack/ directory

go mod tidy

go build -o oraclepack

# run in your repo root so -f paths resolve
./oraclepack --pack /path/to/strategist-questions-oracle-pack-2025-12-30.md

# inject oracle flags into each command
./oraclepack --pack /path/to/strategist-questions-oracle-pack-2025-12-30.md \
  --inject '--browser-attachments always'

# resume selection from step 7
./oraclepack --pack /path/to/strategist-questions-oracle-pack-2025-12-30.md --start 7
```

---

## UX improvements you can add next (optional)

- A “Run next pending” key (e.g. `n`) that automatically advances after success.
- A “Dry-run” mode that never executes, only previews.
- Highlighting the exact `--write-output` path and whether it already exists.
- A “copy command” action (OSC52) for terminals that support it.
- A stricter parser for multiple fenced blocks, multiple pack formats, or nested code fences.

