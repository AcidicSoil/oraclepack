package tui

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	urlScopeProject = "project"
	urlScopeGlobal  = "global"
)

type URLItem struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	LastUsed string `json:"lastUsed,omitempty"`
}

type URLStore struct {
	Default string    `json:"default"`
	Items   []URLItem `json:"items"`
}

func LoadURLStore(path string) (URLStore, error) {
	if path == "" {
		return URLStore{}, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return URLStore{}, nil
		}
		return URLStore{}, err
	}
	var store URLStore
	if err := json.Unmarshal(data, &store); err != nil {
		return URLStore{}, err
	}
	return store, nil
}

func SaveURLStore(path string, store URLStore) error {
	if path == "" {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func ProjectURLStorePath(statePath, packSource string) string {
	if statePath != "" {
		base := strings.TrimSuffix(statePath, ".state.json")
		return base + ".chatgpt-urls.json"
	}
	if packSource == "" {
		return ""
	}
	return packSource + ".chatgpt-urls.json"
}

func GlobalURLStorePath() string {
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return ""
	}
	return filepath.Join(home, ".oraclepack", "chatgpt-urls.json")
}

func nowRFC3339() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func isValidURL(value string) bool {
	v := strings.TrimSpace(value)
	if v == "" {
		return false
	}
	return strings.HasPrefix(v, "http://") || strings.HasPrefix(v, "https://")
}
