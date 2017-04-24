package neuron

import (
	"github.com/breathbath/nn/net"
	"github.com/breathbath/nn/tests/log"
	"github.com/breathbath/nn/tests/matrix"
	observerTests "github.com/breathbath/nn/tests/observers"
	"github.com/breathbath/nn/signal"
	"github.com/breathbath/nn/observers"
	"fmt"
	"github.com/breathbath/testing/assertions"
	"testing"
)

var nandNet net.Net
var logger *log.LogMock
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

func initNet() {
	expectationResults = resultsCollection{}
	resultObserver = observerTests.NewObserverMock()

	logger = log.NewLogMock()

	nandMatix := matrix.NewLinearPerceptronMatrixStub()

	nandNet = nandMatix.ConvertToNet(logger, resultObserver)
}

/**
	this test expects to test a simple boolean nand gate as a base of any logic calculation, the use cases
	could be simplified to the following boolean rules (for carry bit see https://en.wikipedia.org/wiki/Bitwise_operation bit shifts)
	input1	input2	result	carryBit
	false	true	true	false
	true	false	true	false
	false	false	false	false
	true	true	false	true
 */
func TestNandResults(t *testing.T) {
	assertTriggeringResult(t, false, true, true, false)
	assertTriggeringResult(t, true, false, true, false)
	assertTriggeringResult(t, false, false, false, false)
	assertTriggeringResult(t, true, true, false, true)
}

func assertTriggeringResult(t *testing.T, input1 bool, input2 bool, expectedResult, expectedCarryBit bool) {
	initNet()

	leftNeuronId := 0
	rightNeuronId := 1

	carryBitOutputNeuronNr := 5
	resultOutputNeuronNr := 6

	inputSignal1 := signal.NewInputSignal(convertBoolToFloatInput(input1), leftNeuronId)
	inputSignal2 := signal.NewInputSignal(convertBoolToFloatInput(input2), rightNeuronId)

	nandNet.ForwardPropagation(inputSignal1, inputSignal2)

	addExpectation(carryBitOutputNeuronNr, convertBoolToFloatInput(expectedCarryBit))
	addExpectation(resultOutputNeuronNr, convertBoolToFloatInput(expectedResult))

	assertAllExpectationsAreMet(t)
}

func convertBoolToFloatInput(input bool) float64 {
	if (input) {
		return 1.0
	}
	return 0
}

func addExpectation(neuronId int, result float64) {
	expectationResults = append(expectationResults, observers.NewObservingResult(neuronId, "some", result))
}

func assertAllExpectationsAreMet(t *testing.T) {
	receivedResultsCount := len(resultObserver.Results)
	expectedResultsCount := len(expectationResults)
	message := fmt.Sprintf(
		"The expected results count %d isn't equal to the received results count %d.\n Expected results: %s, received results: %s",
		expectedResultsCount,
		receivedResultsCount,
		expectationResults,
		resultObserver.Results,
	)
	assertions.AssertEquals(expectedResultsCount, receivedResultsCount, t, message)

	notMetExpectations := resultsCollection{}
	for _, exRes := range expectationResults {
		foundResult, ok := findResultByExpectation(exRes)
		if !ok {
			notMetExpectations = append(notMetExpectations, foundResult)
		}
	}

	if len(notMetExpectations) > 0 {
		t.Errorf(
			"Expectations:\n%s were not met by the received results:\n%s",
			notMetExpectations,
			resultsCollection(resultObserver.Results),
		)
	}

}

func findResultByExpectation(expectation observers.ObservingResult) (observers.ObservingResult, bool) {
	result := observers.ObservingResult{}
	for _, curRes := range resultObserver.Results {
		if curRes.NeuronId == expectation.NeuronId && curRes.CalculationResult == expectation.CalculationResult {
			return curRes, true
		}
	}
	return result, false
}
