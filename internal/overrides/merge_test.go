package overrides

import (
	"reflect"
	"testing"
)

func TestEffectiveFlags(t *testing.T) {
	tests := []struct {
		name      string
		overrides *RuntimeOverrides
		stepID    string
		baseline  []string
		want      []string
	}{
		{
			name:      "No overrides (nil)",
			overrides: nil,
			stepID:    "01",
			baseline:  []string{"--json"},
			want:      []string{"--json"},
		},
		{
			name: "Step not targeted",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"02": true},
				AddedFlags:   []string{"--verbose"},
			},
			stepID:   "01",
			baseline: []string{"--json"},
			want:     []string{"--json"},
		},
		{
			name: "Step targeted: Add flags",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"01": true},
				AddedFlags:   []string{"--verbose"},
			},
			stepID:   "01",
			baseline: []string{"--json"},
			want:     []string{"--json", "--verbose"},
		},
		{
			name: "Step targeted: Remove flags",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"01": true},
				RemovedFlags: []string{"--json"},
			},
			stepID:   "01",
			baseline: []string{"--json", "--other"},
			want:     []string{"--other"},
		},
		{
			name: "Step targeted: Add and Remove",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"01": true},
				AddedFlags:   []string{"--new"},
				RemovedFlags: []string{"--old"},
			},
			stepID:   "01",
			baseline: []string{"--old", "--keep"},
			want:     []string{"--keep", "--new"},
		},
		{
			name: "Step targeted: Inject ChatGPT URL",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"01": true},
				ChatGPTURL:   "https://chat.openai.com/share/123",
			},
			stepID:   "01",
			baseline: []string{"--json"},
			want:     []string{"--json", "--chatgpt-url", "https://chat.openai.com/share/123"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.overrides.EffectiveFlags(tt.stepID, tt.baseline)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EffectiveFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}
