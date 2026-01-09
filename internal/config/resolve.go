package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ResolveOutputVerify applies precedence: CLI flag > env var > default.
func ResolveOutputVerify(flagValue bool, flagSet bool) (bool, error) {
	if flagSet {
		return flagValue, nil
	}
	if val, ok := os.LookupEnv(EnvOutputVerify); ok {
		parsed, err := parseBoolish(val)
		if err != nil {
			return DefaultOutputVerify, fmt.Errorf("invalid %s: %w", EnvOutputVerify, err)
		}
		return parsed, nil
	}
	return DefaultOutputVerify, nil
}

// ResolveOutputRetries applies precedence: CLI flag > env var > default.
func ResolveOutputRetries(flagValue int, flagSet bool) (int, error) {
	if flagSet {
		return flagValue, nil
	}
	if val, ok := os.LookupEnv(EnvOutputRetries); ok {
		parsed, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			return DefaultOutputRetries, fmt.Errorf("invalid %s: %w", EnvOutputRetries, err)
		}
		return parsed, nil
	}
	return DefaultOutputRetries, nil
}

func parseBoolish(raw string) (bool, error) {
	v := strings.ToLower(strings.TrimSpace(raw))
	switch v {
	case "1", "true", "yes", "on":
		return true, nil
	case "0", "false", "no", "off":
		return false, nil
	default:
		return false, fmt.Errorf("expected boolean (true/false, 1/0, on/off), got %q", raw)
	}
}
