package vector

import "github.com/breathbath/nn/neuron"

type VectorBehaviour interface {
	TransferSignal(sig float64)
	String() string
	GetTo() neuron.NeuronBehaviour
	GetFrom() neuron.NeuronBehaviour
}
