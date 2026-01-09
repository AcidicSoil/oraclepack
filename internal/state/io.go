package state

// Intentionally left without extra imports.

// WriteState writes RunState atomically to disk.
func WriteState(path string, state *RunState) error {
	return SaveStateAtomic(path, state)
}
