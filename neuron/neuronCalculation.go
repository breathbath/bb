package neuron

type NeuronCalculation interface {
	CalculateForward(signalSum float64, bias float64) float64
}
