package input

type Vector struct {
	Address Coord
	Value float64
}

func NewVector(sourceNeuronNr, destinationNeuronNr int, value float64) *Vector {
	return &Vector{NewCoord(sourceNeuronNr, destinationNeuronNr), value}
}