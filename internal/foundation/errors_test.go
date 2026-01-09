package foundation

import (
	"errors"
	"testing"
)

func TestCommonErrors(t *testing.T) {
	if ErrMissingBinary == nil || ErrArtifactMissing == nil {
		t.Fatal("expected error variables to be initialized")
	}
	if !errors.Is(ErrMissingBinary, ErrMissingBinary) {
		t.Fatal("errors.Is failed for ErrMissingBinary")
	}
	if !errors.Is(ErrArtifactMissing, ErrArtifactMissing) {
		t.Fatal("errors.Is failed for ErrArtifactMissing")
	}
}
