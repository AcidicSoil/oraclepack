package render

import (
	"github.com/charmbracelet/glamour"
	"github.com/user/oraclepack/internal/pack"
)

// RenderMarkdown renders markdown text as ANSI-styled text.
func RenderMarkdown(text string) (string, error) {
	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(80),
	)
	if err != nil {
		return "", err
	}

	return r.Render(text)
}

// RenderStepCode renders a step's code block for preview.
func RenderStepCode(s pack.Step) (string, error) {
	md := "```bash\n" + s.Code + "\n```"
	return RenderMarkdown(md)
}

