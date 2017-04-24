package process

import (
	"github.com/breathbath/nn/nn2/result"
	"github.com/breathbath/nn/nn2/input"
)

type Processor struct {
	Results            *result.ResultMatrix
	InputMatrix        *input.InputMatrix
	activationFunction ActivationFunction
}

func NewProcessor(rm *result.ResultMatrix, ints *input.InputMatrix) *Processor {
	processor := Processor{rm, ints, Sigmoid{}}

	return &processor
}

func (this *Processor) ForwardPropagation(inputs ...input.NetworkSignal) {
	this.addInputsToMatrix(inputs)
	var sourceNeuronNr int
	var currInputValue, currWeight, activationResult float64
	var curSum float64
	var hasSingleInput bool
	for _, destinatinNeuronNr := range this.InputMatrix.GetDestinationKeys() {
		curSum = 0.0
		destinationVectors := this.InputMatrix.GetVectorsByDestinationNode(destinatinNeuronNr)
		hasSingleInput = len(destinationVectors) == 1

		for _, destinationVector := range destinationVectors {
			sourceNeuronNr = destinationVector.Address.SourceNeuronNr
			currWeight = destinationVector.Value

			currInput, ok := this.Results.GetCalculationResult(sourceNeuronNr)
			if ok {
				currInputValue = currInput.FinalResult
			} else {
				currInputValue = 1 //as x * 1 == x
			}

			curSum = currInputValue * currWeight + curSum
		}
		this.Results.AddSum(destinatinNeuronNr, curSum)

		//this condition identifies the input neurons, for such we don't use activation function
		if hasSingleInput && (sourceNeuronNr == input.INPUT_SOURCE_NEURON_NUMBER) {
			this.Results.AddFinalResult(destinatinNeuronNr, curSum, this.isOutputVector(destinatinNeuronNr))
		} else {
			activationResult = this.activationFunction.Activate(curSum)
			this.Results.AddActivationResult(destinatinNeuronNr, activationResult)
			this.Results.AddFinalResult(destinatinNeuronNr, activationResult, this.isOutputVector(destinatinNeuronNr))
		}

	}
}

func (this *Processor) isOutputVector(destinationNeuronNr int) bool {
	return !this.InputMatrix.HasSourceVector(destinationNeuronNr)
}

func (this *Processor) addInputsToMatrix(inputs[]input.NetworkSignal) {
	var newVector *input.Vector
	for _, inputItem := range inputs {
		newVector = input.NewVector(input.INPUT_SOURCE_NEURON_NUMBER, inputItem.NeuronNr, inputItem.Value)
		this.InputMatrix.AddVector(newVector)
	}
}