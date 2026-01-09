package types

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/goccy/go-yaml"
)

func TestPackJSONRoundTrip(t *testing.T) {
	original := Pack{
		Prelude: Prelude{Code: "echo prelude"},
		Steps: []Step{
			{
				ID:           "01",
				Number:       1,
				Code:         "echo hello",
				OriginalLine: "# 01) Example",
				ROI:          3.2,
				Impact:       "High",
			},
		},
		Source:      "pack.md",
		OutDir:      "dist",
		WriteOutput: true,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("json marshal: %v", err)
	}

	var decoded Pack
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("json unmarshal: %v", err)
	}

	if !reflect.DeepEqual(original, decoded) {
		t.Fatalf("json round-trip mismatch: %#v != %#v", original, decoded)
	}
}

func TestPackYAMLRoundTrip(t *testing.T) {
	original := Pack{
		Prelude: Prelude{Code: "echo prelude"},
		Steps: []Step{
			{
				ID:           "02",
				Number:       2,
				Code:         "echo yaml",
				OriginalLine: "# 02) Example",
				ROI:          1.1,
				Impact:       "Low",
			},
		},
		Source:      "pack.yaml",
		OutDir:      "out",
		WriteOutput: false,
	}

	data, err := yaml.Marshal(original)
	if err != nil {
		t.Fatalf("yaml marshal: %v", err)
	}

	var decoded Pack
	if err := yaml.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("yaml unmarshal: %v", err)
	}

	if !reflect.DeepEqual(original, decoded) {
		t.Fatalf("yaml round-trip mismatch: %#v != %#v", original, decoded)
	}
}
