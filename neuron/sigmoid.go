package neuron

import (
	"github.com/breathbath/nn/observers"
	"github.com/breathbath/nn/log"
	"math"
)

type Sigmoid struct {
}

func (s Sigmoid) CalculateForward(signalSum float64, bias float64) float64 {
	return 1 / (1 + math.Exp(- signalSum - bias))
}

func NewSigmoid(bias float64, matrixRowPosition int, outputSubscriber observers.Observer, logger log.LogBehaviour) NeuronBehaviour {
	return &Neuron{bias, matrixRowPosition, 0, 0, 0, outputSubscriber, logger, Sigmoid{}}
}