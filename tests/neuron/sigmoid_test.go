package neuron

import (
	"testing"
	"github.com/breathbath/nn/neuron"
	"github.com/breathbath/testing/assertions"
	testObservers "github.com/breathbath/nn/tests/observers"
	testLog "github.com/breathbath/nn/tests/log"
)

func TestSigmoidCreation(t *testing.T) {
	logger := testLog.NewLogMock()
	subscriber := testObservers.NewObserverMock()
	s := neuron.NewSigmoid(5, 2, subscriber, logger)
	assertions.AssertEquals(2, s.GetId(), t, "")
	assertions.AssertEquals("(2)", s.String(), t, "")
}

func TestSigmoidCalculation(t *testing.T) {
	p := neuron.Sigmoid{}
	result := p.CalculateForward(1.0, 2)
	assertions.AssertFloatEquals(0.952574, result, 0.000001, t)

	result = p.CalculateForward(1.0, -2)
	assertions.AssertFloatEquals(0.268941, result, 0.000001, t)
}
