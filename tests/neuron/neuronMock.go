package neuron

type NeuronMock struct {
	Signals [] float64
	Weights [] float64
	Name    string
	Id      int
	Count   int
}

func NewNeuronMock(name string, id int) *NeuronMock {
	return &NeuronMock{[]float64{}, []float64{}, name, id, 0}
}

func (nm *NeuronMock) AcceptSignal(sig float64, weight float64) {
	nm.Signals = append(nm.Signals, sig)
	nm.Weights = append(nm.Weights, weight)
}

func (nm *NeuronMock) String() string {
	return nm.Name
}

func (nm *NeuronMock) GetId() int {
	return nm.Id
}

func (nm *NeuronMock) IncreaseExpectedCount() {
	nm.Count++
}
