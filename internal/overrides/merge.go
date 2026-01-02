package overrides

// EffectiveFlags calculates the final flags for a step.
func (r *RuntimeOverrides) EffectiveFlags(stepID string, baseline []string) []string {
	if r == nil || r.ApplyToSteps == nil || !r.ApplyToSteps[stepID] {
		return baseline
	}

	var effective []string

	// Map for removed flags
	removed := make(map[string]bool)
	for _, f := range r.RemovedFlags {
		removed[f] = true
	}

	// Filter baseline
	for _, flag := range baseline {
		if !removed[flag] {
			effective = append(effective, flag)
		}
	}

	// Append added flags
	effective = append(effective, r.AddedFlags...)

	// Inject ChatGPTURL
	if r.ChatGPTURL != "" {
		effective = append(effective, "--chatgpt-url", r.ChatGPTURL)
	}

	return effective
}
