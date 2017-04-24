package vector

import (
	"github.com/breathbath/nn/neuron"
	NeuronTests "github.com/breathbath/nn/tests/neuron"
)

type VectorMock struct {
	Signals [] float64
	Name string
	To neuron.NeuronBehaviour
	From neuron.NeuronBehaviour
}

func NewVectorMockByNeurons(name string, to, from neuron.NeuronBehaviour) *VectorMock{
	return &VectorMock{[]float64{}, name, to, from}
}

func NewVectorMockByNeuronData(neuronNameFrom string, neuronIdFrom int, neuronNameTo string, neuronIdTo int, vectorName string) *VectorMock{
	fromNeuron := NeuronTests.NewNeuronMock(neuronNameFrom, neuronIdFrom)
	toNeuron := NeuronTests.NewNeuronMock(neuronNameTo, neuronIdTo)
	return NewVectorMockByNeurons(vectorName, toNeuron, fromNeuron)
}

func (vm *VectorMock) TransferSignal(sig float64) {
	vm.Signals = append(vm.Signals, sig)
}

func (vm *VectorMock) String() string {
	return vm.Name
}

func (vm *VectorMock) GetTo() neuron.NeuronBehaviour {
	return vm.To
}

func (vm *VectorMock) GetFrom() neuron.NeuronBehaviour {
	return vm.From
}
