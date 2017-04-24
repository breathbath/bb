package observers

import (
	"testing"
	"github.com/breathbath/nn/observers"
	"github.com/breathbath/testing/assertions"
)

func TestObservingResultCreation(t *testing.T) {
	result := observers.NewObservingResult(1, "someName", 2.0)
	assertions.AssertEquals(1, result.NeuronId, t, "")
	assertions.AssertEquals("someName", result.Name, t, "")
	assertions.AssertEquals(2.0, result.CalculationResult, t, "")
}
