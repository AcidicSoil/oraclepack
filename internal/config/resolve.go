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

// ResolveOutputRequireHeadings applies precedence: CLI flag > env var > default.
func ResolveOutputRequireHeadings(flagValue bool, flagSet bool) (bool, error) {
	if flagSet {
		return flagValue, nil
	}
	if val, ok := os.LookupEnv(EnvOutputRequireHeadings); ok {
		parsed, err := parseBoolish(val)
		if err != nil {
			return DefaultOutputRequireHeadings, fmt.Errorf("invalid %s: %w", EnvOutputRequireHeadings, err)
		}
		return parsed, nil
	}
	return DefaultOutputRequireHeadings, nil
}

// ResolveOutputChunkMode applies precedence: CLI flag > env var > default.
func ResolveOutputChunkMode(flagValue string, flagSet bool) (string, error) {
	if flagSet {
		return normalizeChunkMode(flagValue)
	}
	if val, ok := os.LookupEnv(EnvOutputChunkMode); ok {
		return normalizeChunkMode(val)
	}
	return normalizeChunkMode(DefaultOutputChunkMode)
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

func normalizeChunkMode(raw string) (string, error) {
	v := strings.ToLower(strings.TrimSpace(raw))
	switch v {
	case "auto", "single", "multi":
		return v, nil
	case "":
		return DefaultOutputChunkMode, nil
	default:
		return "", fmt.Errorf("invalid %s: expected auto|single|multi, got %q", EnvOutputChunkMode, raw)
	}
}
