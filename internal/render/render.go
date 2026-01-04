package render

import (
	"sync"

	"github.com/charmbracelet/glamour"
	"github.com/user/oraclepack/internal/pack"
)

const (
	DefaultStyle = "dark"
	DefaultWidth = 80
)

type rendererKey struct {
	width int
	style string
}

var (
	rendererMu    sync.Mutex
	rendererCache = map[rendererKey]*glamour.TermRenderer{}
)

// RenderMarkdown renders markdown text as ANSI-styled text.
func RenderMarkdown(text string, width int, style string) (string, error) {
	if width <= 0 {
		width = DefaultWidth
	}
	if style == "" {
		style = DefaultStyle
	}

	r, err := rendererFor(width, style)
	if err != nil {
		return "", err
	}

	return r.Render(text)
}

// RenderStepCode renders a step's code block for preview.
func RenderStepCode(s pack.Step, width int, style string) (string, error) {
	md := "```bash\n" + s.Code + "\n```"
	return RenderMarkdown(md, width, style)
}

func rendererFor(width int, style string) (*glamour.TermRenderer, error) {
	key := rendererKey{width: width, style: style}

	rendererMu.Lock()
	r := rendererCache[key]
	rendererMu.Unlock()
	if r != nil {
		return r, nil
	}

	opts := []glamour.TermRendererOption{glamour.WithWordWrap(width)}
	if style == "auto" {
		opts = append(opts, glamour.WithAutoStyle())
	} else {
		opts = append(opts, glamour.WithStandardStyle(style))
	}

	r, err := glamour.NewTermRenderer(opts...)
	if err != nil {
		return nil, err
	}

	rendererMu.Lock()
	rendererCache[key] = r
	rendererMu.Unlock()
	return r, nil
}
