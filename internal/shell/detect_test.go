package shell

import "testing"

func TestDetectBinary(t *testing.T) {
	if _, ok := DetectBinary("ls"); !ok {
		t.Fatalf("expected to find ls on PATH")
	}
	if _, ok := DetectBinary("definitely-not-a-real-binary-123"); ok {
		t.Fatalf("expected missing binary to return false")
	}
}
