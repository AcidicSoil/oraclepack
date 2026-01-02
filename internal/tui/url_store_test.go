package tui

import (
	"path/filepath"
	"testing"
)

func TestURLStoreSaveLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "urls.json")

	store := URLStore{
		Default: "Primary",
		Items: []URLItem{{
			Name: "Primary",
			URL:  "https://chatgpt.com/g/primary",
		}},
	}

	if err := SaveURLStore(path, store); err != nil {
		t.Fatalf("failed to save store: %v", err)
	}

	loaded, err := LoadURLStore(path)
	if err != nil {
		t.Fatalf("failed to load store: %v", err)
	}

	if loaded.Default != store.Default {
		t.Fatalf("expected default %q, got %q", store.Default, loaded.Default)
	}
	if len(loaded.Items) != 1 || loaded.Items[0].URL != store.Items[0].URL {
		t.Fatalf("loaded items mismatch: %+v", loaded.Items)
	}
}

func TestURLPickerDefaultURLPrefersProject(t *testing.T) {
	dir := t.TempDir()
	projectPath := filepath.Join(dir, "project.json")
	globalPath := filepath.Join(dir, "global.json")

	project := URLStore{
		Default: "Project",
		Items: []URLItem{{
			Name: "Project",
			URL:  "https://chatgpt.com/g/project",
		}},
	}
	global := URLStore{
		Default: "Global",
		Items: []URLItem{{
			Name: "Global",
			URL:  "https://chatgpt.com/g/global",
		}},
	}

	if err := SaveURLStore(projectPath, project); err != nil {
		t.Fatalf("failed to save project store: %v", err)
	}
	if err := SaveURLStore(globalPath, global); err != nil {
		t.Fatalf("failed to save global store: %v", err)
	}

	picker := NewURLPickerModel(projectPath, globalPath)
	if got := picker.DefaultURL(); got != project.Items[0].URL {
		t.Fatalf("expected project default URL %q, got %q", project.Items[0].URL, got)
	}
}
