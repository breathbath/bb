package neuron

type NeuronBehaviour interface {
	AcceptSignal(sig float64, weight float64)
	String() string
	GetId() int
	IncreaseExpectedCount()
}
