package vector

import (
	"testing"
	"github.com/breathbath/nn/vector"
	testNeurons "github.com/breathbath/nn/tests/neuron"
	"github.com/breathbath/nn/tests/log"
	"github.com/breathbath/testing/assertions"
)

var neuron1 *testNeurons.NeuronMock
var neuron2 *testNeurons.NeuronMock
var logger *log.LogMock

func initDependencies() {
	neuron1 = testNeurons.NewNeuronMock("n1", 1)
	neuron2 = testNeurons.NewNeuronMock("n2", 2)
	logger = log.NewLogMock()
}

func TestVectorCreation(t *testing.T) {
	initDependencies()
	v := vector.NewVector(neuron1, neuron2, 1, logger)
	assertions.AssertEquals(neuron1, v.GetFrom(), t, "")
	assertions.AssertEquals(neuron2, v.GetTo(), t, "")
	assertions.AssertEquals("n1 -> n2", v.String(), t, "")
	assertions.AssertEquals(1.0, v.Weight, t, "")
	v.TransferSignal(22)
	if !logger.HasMessages() {
		t.Error("Logger should be attached to the vector and be able to log messages")
	}
}

func TestResultVectorCreation(t *testing.T) {
	initDependencies()
	v := vector.NewResultVector(neuron1, 1, logger)
	assertions.AssertEquals(nil, v.To, t, "")
	assertions.AssertEquals(1.0, v.Weight, t, "")
}

func TestTransferringSignal(t *testing.T) {
	initDependencies()
	v := vector.NewVector(neuron1, neuron2, 1, logger)
	v.TransferSignal(11)
	assertions.AssertEquals(1, len(neuron2.Signals), t, "")
	assertions.AssertEquals(11.0, neuron2.Signals[0], t, "")
	assertions.AssertEquals(1, len(neuron2.Weights), t, "")
	assertions.AssertEquals(1.0, neuron2.Weights[0], t, "")
	assertions.AssertEquals(0, len(neuron1.Signals), t, "")
}
