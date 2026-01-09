package foundation

import "time"

// Clock abstracts time for deterministic testing.
type Clock interface {
	Now() time.Time
}

// RealClock uses the system clock.
type RealClock struct{}

// Now returns the current time.
func (RealClock) Now() time.Time { return time.Now() }

// MockClock returns a fixed time that can be advanced.
type MockClock struct {
	current time.Time
}

// NewMockClock initializes a mock clock with a starting time.
func NewMockClock(start time.Time) *MockClock {
	return &MockClock{current: start}
}

// Now returns the mock time.
func (m *MockClock) Now() time.Time { return m.current }

// Advance moves the mock time forward.
func (m *MockClock) Advance(d time.Duration) {
	m.current = m.current.Add(d)
}
