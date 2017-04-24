package observers

import (
	"fmt"
)

type NetworkResultOutputter struct {
}

func (o NetworkResultOutputter) AcceptResult(result ObservingResult) {
	fmt.Printf("----->Result neuron %s gives output: %.5f<------\n", result.Name, result.CalculationResult)
}