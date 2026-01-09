package foundation

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Config holds runtime settings that can be loaded from JSON and environment variables.
// Env values always take precedence over JSON values.
type Config struct {
	Name      string  `json:"name" env:"ORACLEPACK_NAME"`
	Retries   int     `json:"retries" env:"ORACLEPACK_RETRIES"`
	Enabled   bool    `json:"enabled" env:"ORACLEPACK_ENABLED"`
	Threshold float64 `json:"threshold" env:"ORACLEPACK_THRESHOLD"`
}

// LoadConfig loads configuration from a JSON file and then applies environment overrides.
// If path is empty, JSON loading is skipped and only env overrides are applied.
func LoadConfig(path string) (Config, error) {
	var cfg Config
	if path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return Config{}, fmt.Errorf("read config: %w", err)
		}
		if err := json.Unmarshal(data, &cfg); err != nil {
			return Config{}, fmt.Errorf("parse config: %w", err)
		}
	}

	if v, ok := os.LookupEnv("ORACLEPACK_NAME"); ok {
		cfg.Name = v
	}
	if v, ok := os.LookupEnv("ORACLEPACK_RETRIES"); ok {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return Config{}, fmt.Errorf("parse ORACLEPACK_RETRIES: %w", err)
		}
		cfg.Retries = parsed
	}
	if v, ok := os.LookupEnv("ORACLEPACK_ENABLED"); ok {
		parsed, err := strconv.ParseBool(v)
		if err != nil {
			return Config{}, fmt.Errorf("parse ORACLEPACK_ENABLED: %w", err)
		}
		cfg.Enabled = parsed
	}
	if v, ok := os.LookupEnv("ORACLEPACK_THRESHOLD"); ok {
		parsed, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return Config{}, fmt.Errorf("parse ORACLEPACK_THRESHOLD: %w", err)
		}
		cfg.Threshold = parsed
	}

	return cfg, nil
}
