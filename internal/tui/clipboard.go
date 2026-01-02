package tui

import (
	"os/exec"
	"runtime"
	"strings"
)

func copyToClipboard(content string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		if _, err := exec.LookPath("wl-copy"); err == nil {
			cmd = exec.Command("wl-copy")
		} else if _, err := exec.LookPath("xclip"); err == nil {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		} else if _, err := exec.LookPath("xsel"); err == nil {
			cmd = exec.Command("xsel", "--clipboard", "--input")
		} else {
			return err
		}
	case "windows":
		cmd = exec.Command("cmd", "/c", "clip")
	default:
		return exec.ErrNotFound
	}

	cmd.Stdin = strings.NewReader(content)
	return cmd.Run()
}
