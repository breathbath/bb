package neuron

import (
	"testing"
	"github.com/breathbath/nn/tests/log"
	testObservers "github.com/breathbath/nn/tests/observers"
	"github.com/breathbath/nn/net"
	"github.com/breathbath/nn/tests/vector"
	"github.com/breathbath/nn/signal"
	"github.com/breathbath/testing/assertions"
	"github.com/breathbath/nn/tests/neuron"
	"github.com/breathbath/nn/observers"
)

var loggerMock *log.LogMock
var observerMock *testObservers.ObserverMock
var nnet net.Net
var inputVector, outputVector *vector.VectorMock

func initDependencies() {
	loggerMock = log.NewLogMock()
	observerMock = testObservers.NewObserverMock()

	nnet = net.NewNet(observerMock, loggerMock)

	inputVector = vector.NewVectorMockByNeuronData("neuronFrom1", 1, "neuronTo1", 2, "vectorName1")
	nnet.AddVector(inputVector)

	neuronFrom := neuron.NewNeuronMock("neuronFrom2", 3)
	outputVector = &vector.VectorMock{
		Signals: []float64{},
		Name: "outputVector",
		From: neuronFrom,
	}
	nnet.AddVector(outputVector)
}

func TestNetCalculation(t *testing.T) {
	initDependencies()
	signalToNeuron1 := signal.NewInputSignal(1.0, 1)
	signalToNeuron3 := signal.NewInputSignal(4.0, 3)
	nnet.ForwardPropagation(signalToNeuron1, signalToNeuron3)

	if (len(inputVector.Signals) == 0) {
		panic("Input vector should have received a signal")
	}

	signalReceivedByInputVector := inputVector.Signals[0]
	assertions.AssertEquals(1.0, signalReceivedByInputVector, t, "")

	signalReceivedByOutputVector := outputVector.Signals[0]
	assertions.AssertEquals(4.0, signalReceivedByOutputVector, t, "")
}

func TestNetNeuronObservation(t *testing.T) {
	initDependencies()
	resultToObserve := observers.NewObservingResult(3, "someName", 33.0)
	nnet.AcceptResult(resultToObserve)
	if len(observerMock.Results) == 0 {
		panic("Net should receive at least one observation result")
	}

	receivedResult := observerMock.Results[0]
	assertions.AssertEquals(33.0, receivedResult.CalculationResult, t, "")
	assertions.AssertEquals(3, receivedResult.NeuronId, t, "")
	assertions.AssertEquals("someName", receivedResult.Name, t, "")

	initDependencies()
	resultToObserve2 := observers.NewObservingResult(1, "someOtherName", 44.0)
	nnet.AcceptResult(resultToObserve2)
	if len(observerMock.Results) != 0 {
		panic("Net should not report final result")
	}

	if len(inputVector.Signals) == 0 {
		panic("Input vector should have received calculation result")
	}
	assertions.AssertEquals(44.0, inputVector.Signals[0], t, "")
}
