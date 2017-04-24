package input

import "fmt"

const BIAS_SOURCE_NEURON_NUMBER = -1
const INPUT_SOURCE_NEURON_NUMBER = 0

type Coord struct {
	SourceNeuronNr, DestinationNeuronNr int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.SourceNeuronNr, c.DestinationNeuronNr)
}

func NewCoord(sourceNeuronNr, destinationNeuronNr int) Coord{
	return Coord{sourceNeuronNr, destinationNeuronNr}
}
