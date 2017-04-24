package signal

type InputSignal struct {
	InputValue   float64
	NeuronNumber int
}

func NewInputSignal(value float64, neuronNumber int) InputSignal {
	return InputSignal{value, neuronNumber}
}
