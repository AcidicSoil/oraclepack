package tools

import "testing"

func TestMetadataRegistry(t *testing.T) {
	meta, ok := Metadata(ToolCodex)
	if !ok {
		t.Fatalf("expected metadata for codex")
	}
	if meta.Name != "codex" {
		t.Fatalf("expected codex name, got %s", meta.Name)
	}
	if len(meta.Args) != 1 || meta.Args[0] != "exec" {
		t.Fatalf("expected codex exec args, got %+v", meta.Args)
	}
	if ToolOracle.Name() != "oracle" {
		t.Fatalf("expected oracle name, got %s", ToolOracle.Name())
	}
}
