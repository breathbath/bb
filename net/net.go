package net

import (
	"fmt"
	"github.com/breathbath/nn/vector"
	"github.com/breathbath/nn/observers"
	"github.com/breathbath/nn/log"
	"github.com/breathbath/nn/signal"
)

type Net struct {
	vectors          map[int][]vector.VectorBehaviour
	resultSubscriber observers.Observer
	logger           log.LogBehaviour
}

func NewNet(resultSubscriber observers.Observer, logger log.LogBehaviour) Net {
	return Net{
		vectors: make(map[int][]vector.VectorBehaviour),
		resultSubscriber: resultSubscriber,
		logger:logger,
	}
}


func (n Net) AcceptResult(result observers.ObservingResult) {
	n.logger.Logf("Network received result %f from neuron %s", result.CalculationResult, result.Name)
	vectors, ok := n.vectors[result.NeuronId]
	if !ok {
		panic("No vector for neuron " + result.Name)
	}

	for _, curVector := range vectors {
		if curVector.GetTo() == nil {
			n.logger.Logf("Found result vector %s. Will report result %.5f", curVector.String(), result.CalculationResult)
			n.resultSubscriber.AcceptResult(result)
		} else {
			n.logger.Logf("Transfer result %.5f to vector %s", result.CalculationResult, curVector.String())
			curVector.TransferSignal(result.CalculationResult)
		}
	}
}

func (n *Net) ForwardPropagation(inputs ...signal.InputSignal) {
	n.logger.Log("Starting network activity")
	for _, curSignal := range inputs {
		inputVectorGroup, ok := n.vectors[curSignal.NeuronNumber]
		if !ok {
			panic(fmt.Sprintf("Cannot find input vector by neuron number %d", curSignal.NeuronNumber))
		}
		for _, inputVector := range inputVectorGroup {
			n.logger.Logf("Found input vector %s. Transfering signal to it!", inputVector.String())
			inputVector.TransferSignal(curSignal.InputValue)
		}
	}
}

func (n *Net) AddVector(vector vector.VectorBehaviour) {
	n.vectors[vector.GetFrom().GetId()] = append(n.vectors[vector.GetFrom().GetId()], vector)
}