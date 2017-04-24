package input

type NetworkSignal struct {
	NeuronNr int
	Value    float64
}

func NewNetworkSignal(neuronNr int, value float64) NetworkSignal {
	return NetworkSignal{
		NeuronNr:neuronNr,
		Value:value,
	}
}