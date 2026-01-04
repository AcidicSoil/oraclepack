package render

import (
	"strings"
	"testing"
)

func TestRenderMarkdown(t *testing.T) {
	text := "# Hello\n**bold**"
	got, err := RenderMarkdown(text, 40, DefaultStyle)
	if err != nil {
		t.Fatalf("RenderMarkdown failed: %v", err)
	}

	// ANSI escape codes start with \x1b[
	if !strings.Contains(got, "\x1b[") {
		t.Errorf("expected ANSI codes in output, got: %q", got)
	}
}
