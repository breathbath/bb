package matrix

import (
	"testing"
	"github.com/breathbath/nn/tests/log"
	"github.com/breathbath/nn/tests/observers"
	"github.com/breathbath/nn/signal"
)

func TestMatrixToNetConversion(t *testing.T) {
	loggerMock := log.NewLogMock()
	resultObserverMock := observers.NewObserverMock()
	matrix := NewLinearPerceptronMatrixStub()
	net := matrix.ConvertToNet(loggerMock, resultObserverMock)
	net.ForwardPropagation(signal.NewInputSignal(0.0, 0), signal.NewInputSignal(0.0, 1))
	if len(resultObserverMock.Results) != 2 {
		t.Errorf("Net should report %d reults but received %d results", 2, len(resultObserverMock.Results))
	}
	if len(loggerMock.Messages) <= 0 {
		t.Error("Logger should have been called multiple times")
	}
}
