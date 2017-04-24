package vector

import (
	"fmt"
	"github.com/breathbath/nn/neuron"
	"github.com/breathbath/nn/log"
)

type Vector struct {
	From   neuron.NeuronBehaviour
	To     neuron.NeuronBehaviour
	Weight float64
	logger log.LogBehaviour
}

func NewVector(fromNeuron neuron.NeuronBehaviour, toNeuron neuron.NeuronBehaviour, weight float64, logger log.LogBehaviour) Vector {
	return Vector{
		From:fromNeuron,
		To: toNeuron,
		Weight:weight,
		logger:logger,
	}
}

func NewResultVector(fromNeuron neuron.NeuronBehaviour, weight float64, logger log.LogBehaviour) Vector {
	return Vector{
		From:fromNeuron,
		Weight:weight,
		logger:logger,
	}
}

func (v *Vector) TransferSignal(sig float64) {
	v.logger.Logf("Vector %s is transfering signal %.5f to neuron %s", v, sig, v.To)
	v.To.AcceptSignal(sig, v.Weight)
}

func (v *Vector) String() string {
	var toStr string
	if (v.To == nil) {
		toStr = "exit"
	} else {
		toStr = v.To.String()
	}
	return fmt.Sprintf("%s -> %s", v.From.String(), toStr)
}

func (v *Vector) GetTo() neuron.NeuronBehaviour {
	return v.To
}

func (v *Vector) GetFrom() neuron.NeuronBehaviour {
	return v.From
}

