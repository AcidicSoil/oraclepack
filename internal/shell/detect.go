package shell

import "os/exec"

// DetectBinary checks PATH for a binary and returns its full path if found.
func DetectBinary(name string) (string, bool) {
	path, err := exec.LookPath(name)
	if err != nil {
		return "", false
	}
	return path, true
}
