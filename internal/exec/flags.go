package exec

import "strings"

// ApplyChatGPTURL ensures a single --chatgpt-url flag is present when url is set.
// It removes any existing --chatgpt-url/--browser-url flags and their values.
func ApplyChatGPTURL(flags []string, url string) []string {
	var out []string
	skipNext := false
	for _, f := range flags {
		if skipNext {
			skipNext = false
			continue
		}
		if f == "--chatgpt-url" || f == "--browser-url" {
			skipNext = true
			continue
		}
		if strings.HasPrefix(f, "--chatgpt-url=") || strings.HasPrefix(f, "--browser-url=") {
			continue
		}
		out = append(out, f)
	}
	if url != "" {
		out = append(out, "--chatgpt-url", url)
	}
	return out
}
