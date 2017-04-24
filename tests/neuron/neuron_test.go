package neuron

import (
	"testing"
	"github.com/breathbath/nn/tests/log"
	"github.com/breathbath/nn/tests/observers"
	"github.com/breathbath/nn/neuron"
)

type SignalMock struct {
	Sig, Weight float64
}

func TestNeuronTriggeringLessSignals(t *testing.T) {
	logger, observer := createLoggerObserver()
	neron := neuron.NewPerceptron(10.0, 1, observer, logger)
	neron.IncreaseExpectedCount()
	neron.IncreaseExpectedCount()
	neron.AcceptSignal(1.0, 1.0)
	if len(observer.Results) > 0 {
		t.Error("Neuron result observer should not be triggered as the singals count is less than expexted")
	}
}

func TestNeuronActivationWithMinusResult(t *testing.T) {
	createAndEvaluateNeuronResult(t, 3.0, 2, []SignalMock{{1.0, -2.0}, {1.0, -2.0}}, 1, 0, 1, "(1)")
}

func TestNeuronActivationWithPlusResult(t *testing.T) {
	createAndEvaluateNeuronResult(t, 3.0, 1, []SignalMock{{1.0, -2.0}}, 2, 1.0, 1, "(2)")
}

func createAndEvaluateNeuronResult(t *testing.T, bias float64, neuronSignalActivationCount int, signals []SignalMock, neuronNumberId int, expectedResult float64, expectedResultCount int, expectedName string) {
	logger, observer := createLoggerObserver()
	nron := neuron.NewPerceptron(bias, neuronNumberId, observer, logger)
	for i := 0; i < neuronSignalActivationCount; i++ {
		nron.IncreaseExpectedCount()
	}
	for _, s := range signals {
		nron.AcceptSignal(s.Sig, s.Weight)
	}

	if len(observer.Results) != expectedResultCount {
		t.Errorf("Neuron should give %d results, %d results given", expectedResultCount, len(observer.Results))
	}

	result := observer.Results[0]
	if result.CalculationResult != expectedResult {
		t.Errorf("Neuron calculation result %f should be equal to %f", result.CalculationResult, expectedResult)
	}

	if result.NeuronId != neuronNumberId {
		t.Errorf("Neuron number %d should be returned in the result but %d is returned", neuronNumberId, result.NeuronId)
	}

	if result.Name != expectedName {
		t.Errorf("Neuron name %s should be returned in the result but %s is returned", expectedName, result.Name)
	}

	if !logger.HasMessages() {
		t.Error("Neuron should call logger at least once")
	}

}

func createLoggerObserver() (*log.LogMock, *observers.ObserverMock) {
	return log.NewLogMock(), observers.NewObserverMock()
}
