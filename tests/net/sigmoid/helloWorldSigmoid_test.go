package neuron

import (
	"github.com/breathbath/nn/net"
	"github.com/breathbath/nn/tests/matrix"
	observerTests "github.com/breathbath/nn/tests/observers"
	"github.com/breathbath/nn/signal"
	"testing"
	"github.com/breathbath/nn/observers"
	"fmt"
	"github.com/breathbath/nn/log"
)


var sigmoidNet net.Net
var logger log.LogBehaviour
var resultObserver *observerTests.ObserverMock
var expectationResults resultsCollection

type resultsCollection []observers.ObservingResult

func (r resultsCollection) String() string {
	stringToReturn := ""
	for _, res := range r {
		stringToReturn += fmt.Sprintf("{neuronId: %d, neuronName: %s, result: %f}\n", res.NeuronId, res.Name, res.CalculationResult)
	}
	return stringToReturn
}

func initSigmoidNet() {
	expectationResults = resultsCollection{}
	resultObserver = observerTests.NewObserverMock()

	logger = log.TtlLog{}

	sigmoidMatrix := matrix.NewSigmoidMatrixStub()

	sigmoidNet = sigmoidMatrix.ConvertToNet(logger, resultObserver)
}

func TestSigmoidNet(t *testing.T) {
	initSigmoidNet()

	inputSignal1 := signal.NewInputSignal(1.0, 0)
	inputSignal2 := signal.NewInputSignal(1.0, 1)

	sigmoidNet.ForwardPropagation(inputSignal1, inputSignal2)
}