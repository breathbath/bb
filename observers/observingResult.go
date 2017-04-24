package observers

type ObservingResult struct {
	NeuronId int
	Name string
	CalculationResult float64
}

func NewObservingResult(neuronId int, name string, calculationResult float64) ObservingResult{
	return ObservingResult{neuronId, name, calculationResult}
}
