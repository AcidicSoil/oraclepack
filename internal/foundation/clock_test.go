package foundation

import (
	"testing"
	"time"
)

func TestMockClock(t *testing.T) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	m := NewMockClock(start)
	if !m.Now().Equal(start) {
		t.Fatalf("expected %v, got %v", start, m.Now())
	}
	m.Advance(2 * time.Hour)
	want := start.Add(2 * time.Hour)
	if !m.Now().Equal(want) {
		t.Fatalf("expected %v, got %v", want, m.Now())
	}
}
