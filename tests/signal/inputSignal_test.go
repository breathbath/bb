package signal

import (
	"testing"
	"github.com/breathbath/nn/signal"
)

func TestInputSignalCreation(t *testing.T) {
	sig := signal.NewInputSignal(1.0, 9)
	if sig.InputValue != 1.0 {
		t.Errorf("Unexpected input value %.5f, expecting 1.0", sig.InputValue)
	}
	if sig.NeuronNumber != 9 {
		t.Errorf("Unexpeted neuron number %d, expecting 9", sig.NeuronNumber)
	}
}