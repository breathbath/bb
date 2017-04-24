package neuron

import (
	"testing"
	"github.com/breathbath/nn/neuron"
	"github.com/breathbath/nn/tests/log"
	"github.com/breathbath/nn/tests/observers"
	"github.com/breathbath/testing/assertions"
)

func TestNeuronCreation(t *testing.T) {
	neuronFactory := createNeuronFactory()
	neurn := neuronFactory.CreateNeuron(10, 1, neuron.TYPE_PERCEPTRON)
	assertions.AssertEquals("(1)", neurn.String(), t, "")
	assertions.AssertEquals(1, neurn.GetId(), t, "")
}

func TestNeuronEquippedWithLoggerAndSubscriber(t *testing.T) {
	logger := log.NewLogMock()
	subscriber := observers.NewObserverMock()
	neuronFactory := neuron.NewNeuronFactory(logger, subscriber)
	neurn := neuronFactory.CreateNeuron(10, 1, neuron.TYPE_PERCEPTRON)
	neurn.AcceptSignal(10, 10)
	if !logger.HasMessages() {
		t.Error("Neuron should be created with the provided logger")
	}
	if len (subscriber.Results) <= 0 {
		t.Error("Neuron should be created with the provided subscriber")
	}
}

func TestFactoryNeuronCache(t *testing.T) {
	neuronFactory := createNeuronFactory()
	neurn1 := neuronFactory.CreateNeuron(10, 1, neuron.TYPE_PERCEPTRON)
	neurn2 := neuronFactory.CreateNeuron(9, 1, neuron.TYPE_PERCEPTRON)
	neurn3 := neuronFactory.CreateNeuron(9, 2, neuron.TYPE_PERCEPTRON)
	if neurn1.String() != neurn2.String() {
		t.Error("Neurons with the same matrix position should be fetched from cache rather than created")
	}
	if neurn2.String() == neurn3.String() {
		t.Error("Neurons with different matrix positions should be recreated")
	}
}

func createNeuronFactory() *neuron.NeuronFactory {
	logger := log.NewLogMock()
	subscriber := observers.NewObserverMock()
	neuronFactory := neuron.NewNeuronFactory(logger, subscriber)
	return neuronFactory
}

