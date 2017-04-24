package neuron

import (
	"testing"
	"github.com/breathbath/nn/neuron"
	"github.com/breathbath/nn/tests/log"
	"github.com/breathbath/nn/tests/observers"
	"github.com/breathbath/testing/assertions"
)

func TestPerceptronCreation(t *testing.T) {
	logger := log.NewLogMock()
	subscriber := observers.NewObserverMock()
	p := neuron.NewPerceptron(10, 1, subscriber, logger)
	assertions.AssertEquals(1, p.GetId(), t, "")
	assertions.AssertEquals("(1)", p.String(), t, "")
}

func TestPerceptronCalculation(t *testing.T) {
	p := neuron.Perceptron{}
	result := p.CalculateForward(1.0, 2)
	assertions.AssertEquals(1.0, result, t, "")

	result = p.CalculateForward(1.0, -2)
	assertions.AssertEquals(0.0, result, t, "")
}
