package neuron

import (
	"fmt"
	"github.com/breathbath/nn/observers"
	"github.com/breathbath/nn/log"
)

type Neuron struct {
	Bias                float64
	MatrixRowPosition   int
	ExpectedInputsCount int
	callsCount          int
	signalSum           float64
	outputSubscriber    observers.Observer
	logger              log.LogBehaviour
	calculator NeuronCalculation
}

func (n *Neuron) AcceptSignal(sig float64, weight float64) {
	n.callsCount++
	n.signalSum += sig * weight
	n.logger.Logf("Saving signals sum %.5f", n.signalSum)
	n.logger.Logf("Neuron %s is deciding if it should fire or not", n)
	if n.callsCount >= n.ExpectedInputsCount {
		n.logger.Logf("Neuron has decided to fire as calls count %d >= expected count %d", n.callsCount, n.ExpectedInputsCount)
		result := n.calculator.CalculateForward(n.signalSum, n.Bias)
		n.logger.Logf("Neuron made a calculation: input signal sum %f, input bias %f, output result %f", n.signalSum, n.Bias, result)
		observingResult := observers.NewObservingResult(n.MatrixRowPosition, n.String(), result)
		n.outputSubscriber.AcceptResult(observingResult)
	} else {
		n.logger.Logf("Decided not to fire as calls count %d < expected count %d", n.callsCount, n.ExpectedInputsCount)
	}
}

func (n *Neuron) GetId() int {
	return n.MatrixRowPosition
}

func (n *Neuron) String() string {
	return fmt.Sprintf("(%d)", n.MatrixRowPosition)
}

func (n *Neuron) IncreaseExpectedCount() {
	n.ExpectedInputsCount++
}
