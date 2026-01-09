package foundation

import "errors"

var (
	// ErrMissingBinary is returned when a required binary is not found on PATH.
	ErrMissingBinary = errors.New("missing binary")
	// ErrArtifactMissing is returned when an expected artifact is absent.
	ErrArtifactMissing = errors.New("artifact missing")
)
