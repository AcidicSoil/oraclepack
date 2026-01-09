package foundation

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigEnvOverrides(t *testing.T) {
	t.Setenv("ORACLEPACK_NAME", "env-name")
	t.Setenv("ORACLEPACK_RETRIES", "5")
	t.Setenv("ORACLEPACK_ENABLED", "true")
	t.Setenv("ORACLEPACK_THRESHOLD", "2.5")

	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	if err := os.WriteFile(path, []byte(`{"name":"json-name","retries":1,"enabled":false,"threshold":1.0}`), 0644); err != nil {
		t.Fatalf("write json: %v", err)
	}

	cfg, err := LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig: %v", err)
	}

	if cfg.Name != "env-name" || cfg.Retries != 5 || cfg.Enabled != true || cfg.Threshold != 2.5 {
		t.Fatalf("env overrides not applied: %+v", cfg)
	}
}

func TestLoadConfigJSONOnly(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	if err := os.WriteFile(path, []byte(`{"name":"json-name","retries":3,"enabled":true,"threshold":4.25}`), 0644); err != nil {
		t.Fatalf("write json: %v", err)
	}

	cfg, err := LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig: %v", err)
	}

	if cfg.Name != "json-name" || cfg.Retries != 3 || cfg.Enabled != true || cfg.Threshold != 4.25 {
		t.Fatalf("json load mismatch: %+v", cfg)
	}
}
