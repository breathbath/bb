package neuron

import "github.com/breathbath/nn/log"
import (
	"github.com/breathbath/nn/observers"
)

type NeuronFactory struct {
	neuronCache map[int]NeuronBehaviour
	logger      log.LogBehaviour
	observer    observers.Observer
}

const TYPE_PERCEPTRON = "perceptron"
const TYPE_SIGMOID = "sigmoid"

var AllTypes map[string]string = map[string]string{
	TYPE_PERCEPTRON: TYPE_PERCEPTRON,
	TYPE_SIGMOID: TYPE_SIGMOID,
}

func NewNeuronFactory(logger log.LogBehaviour, observer observers.Observer) *NeuronFactory {
	return &NeuronFactory{neuronCache: make(map[int]NeuronBehaviour), logger:logger, observer:observer}
}

func (nf *NeuronFactory) CreateNeuron(bias float64, neuronNumber int, neuronType string) NeuronBehaviour {
	_, ok := nf.neuronCache[neuronNumber];
	if !ok {
		mappendNeuronType, ok := AllTypes[neuronType]
		if (!ok) {
			panic("Unknown neuron type " + neuronType)
		}
		if mappendNeuronType == TYPE_PERCEPTRON {
			nf.neuronCache[neuronNumber] = nf.createPerceptron(bias, neuronNumber)
		} else if mappendNeuronType == TYPE_SIGMOID {
			nf.neuronCache[neuronNumber] = NewSigmoid(bias, neuronNumber, nf.observer, nf.logger)
		}
	}
	return nf.neuronCache[neuronNumber]
}

func (nf *NeuronFactory) createPerceptron(bias float64, neuronNumber int) NeuronBehaviour {
	return NewPerceptron(bias, neuronNumber, nf.observer, nf.logger)
}
