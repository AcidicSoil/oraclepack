package tui

import (
	"os"
	"path/filepath"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/types"
)

func TestFilterLogic(t *testing.T) {
	// Setup pack with steps having different ROI
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", ROI: 1.0, OriginalLine: "Step 1"},
			{ID: "02", ROI: 5.0, OriginalLine: "Step 2"},
			{ID: "03", ROI: 10.0, OriginalLine: "Step 3"},
		},
	}
	r := exec.NewRunner(exec.RunnerOptions{})
	s := &state.RunState{}

	// Initialize model with no filter (threshold 0)
	m := NewModel(p, r, s, "", 0, "over", false, false, 0, false, "auto")

	if len(m.list.Items()) != 3 {
		t.Fatalf("expected 3 items initially, got %d", len(m.list.Items()))
	}

	// Apply filter: ROI >= 5.0
	m.roiThreshold = 5.0
	m.roiMode = "over"
	m = m.refreshList()

	if len(m.list.Items()) != 2 {
		t.Errorf("expected 2 items after filtering >= 5.0, got %d", len(m.list.Items()))
	}

	// Verify items are 02 and 03
	items := m.list.Items()
	if items[0].(item).id != "02" {
		t.Errorf("expected first item to be 02, got %s", items[0].(item).id)
	}
	if items[1].(item).id != "03" {
		t.Errorf("expected second item to be 03, got %s", items[1].(item).id)
	}

	// Apply filter: ROI < 5.0 ("under")
	m.roiThreshold = 5.0
	m.roiMode = "under"
	m = m.refreshList()

	if len(m.list.Items()) != 1 {
		t.Errorf("expected 1 item after filtering < 5.0, got %d", len(m.list.Items()))
	}
	if m.list.Items()[0].(item).id != "01" {
		t.Errorf("expected item to be 01, got %s", m.list.Items()[0].(item).id)
	}
}

func TestROIModeTogglePersists(t *testing.T) {
	dir := t.TempDir()
	statePath := filepath.Join(dir, "state.json")
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", ROI: 1.0, OriginalLine: "Step 1"},
		},
	}
	r := exec.NewRunner(exec.RunnerOptions{})
	s := &state.RunState{SchemaVersion: 1}

	m := NewModel(p, r, s, statePath, 0, "over", false, false, 0, false, "auto")

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("m")})
	m2 := updated.(Model)
	if m2.roiMode != "under" {
		t.Fatalf("expected roiMode to toggle to under, got %s", m2.roiMode)
	}

	loaded, err := state.LoadState(statePath)
	if err != nil {
		t.Fatalf("failed to load state: %v", err)
	}
	if loaded.ROIMode != "under" {
		t.Fatalf("expected persisted roiMode under, got %s", loaded.ROIMode)
	}

	if err := os.Remove(statePath); err != nil {
		t.Fatalf("failed to cleanup state file: %v", err)
	}
}
