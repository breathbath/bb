package neuron

import (
	"github.com/breathbath/nn/observers"
	"github.com/breathbath/nn/log"
)


type Perceptron struct {
}

func (p Perceptron) CalculateForward (signalSum float64, bias float64) float64 {
	result := signalSum + bias
	if result > 0 {
		result = 1.0
	} else {
		result = 0
	}
	return result
}

func NewPerceptron(bias float64, matrixRowPosition int, outputSubscriber observers.Observer, logger log.LogBehaviour) NeuronBehaviour {
	return &Neuron{bias, matrixRowPosition, 0, 0, 0, outputSubscriber, logger, Perceptron{}}
}
