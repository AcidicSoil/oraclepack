package pack

import (
	"reflect"
	"testing"

	"github.com/user/oraclepack/internal/types"
)

func TestDetectOutputContract(t *testing.T) {
	step := types.Step{
		Code: "Answer format:\nReturn only: Direct answer\n",
	}
	if got := DetectOutputContract(step); got != types.OutputContractDirectAnswerOnly {
		t.Fatalf("expected direct-answer-only, got %q", got)
	}

	step = types.Step{
		Code: "Answer format:\n### Direct answer\n### Risks and unknowns\n",
	}
	if got := DetectOutputContract(step); got != types.OutputContractAllSections {
		t.Fatalf("expected all-sections, got %q", got)
	}

	step = types.Step{Code: "no format here"}
	if got := DetectOutputContract(step); got != types.OutputContractUnknown {
		t.Fatalf("expected unknown, got %q", got)
	}
}

func TestStepOutputExpectations_SingleOutput(t *testing.T) {
	step := &types.Step{
		Code: `oracle -p "x" --write-output "out.txt"
Answer format:
Return only: Direct answer`,
	}
	expectations := StepOutputExpectations(step)
	want := map[string][]string{"out.txt": []string{"### Direct answer"}}
	if !reflect.DeepEqual(expectations, want) {
		t.Fatalf("unexpected expectations: %#v", expectations)
	}

	step = &types.Step{
		Code: `oracle -p "x" --write-output "out.txt"
Answer format:
### Direct answer
### Risks and unknowns
### Next experiment
### Missing evidence`,
	}
	expectations = StepOutputExpectations(step)
	want = map[string][]string{"out.txt": []string{
		"### Direct answer",
		"### Risks and unknowns",
		"### Next experiment",
		"### Missing evidence",
	}}
	if !reflect.DeepEqual(expectations, want) {
		t.Fatalf("unexpected expectations: %#v", expectations)
	}
}

func TestStepOutputExpectations_MultiOutput(t *testing.T) {
	step := &types.Step{
		Code: `oracle --write-output "out-direct-answer.md" \
  --write-output 'out-risks-unknowns.md' \
  --write-output out-next-experiment.md \
  --write-output out-missing-evidence.md`,
	}
	expectations := StepOutputExpectations(step)
	if len(expectations) != 4 {
		t.Fatalf("expected 4 expectations, got %d", len(expectations))
	}
	if got := expectations["out-direct-answer.md"]; len(got) != 1 || got[0] != "### Direct answer" {
		t.Fatalf("unexpected direct-answer tokens: %#v", got)
	}
	if got := expectations["out-risks-unknowns.md"]; len(got) != 1 || got[0] != "### Risks and unknowns" {
		t.Fatalf("unexpected risks-unknowns tokens: %#v", got)
	}
	if got := expectations["out-next-experiment.md"]; len(got) != 1 || got[0] != "### Next experiment" {
		t.Fatalf("unexpected next-experiment tokens: %#v", got)
	}
	if got := expectations["out-missing-evidence.md"]; len(got) != 1 || got[0] != "### Missing evidence" {
		t.Fatalf("unexpected missing-evidence tokens: %#v", got)
	}
}
