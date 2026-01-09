package foundation

import (
	"fmt"
	"os"
)

// WriteAtomic writes data to path atomically by writing to a temp file and renaming.
func WriteAtomic(path string, data []byte, perm os.FileMode) error {
	tempPath := path + ".tmp"
	if err := os.WriteFile(tempPath, data, perm); err != nil {
		return fmt.Errorf("write temp file: %w", err)
	}
	if err := os.Rename(tempPath, path); err != nil {
		_ = os.Remove(tempPath)
		return fmt.Errorf("rename temp file: %w", err)
	}
	return nil
}
