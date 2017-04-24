package result

import "fmt"

type ResultMatrix struct {
	CalculationHistory map[int]*ResultItem
	Outputs            []*ResultItem
}

func NewResultMatrix() *ResultMatrix {
	outputs := []*ResultItem{}
	return &ResultMatrix{
		CalculationHistory: make(map[int]*ResultItem),
		Outputs: outputs,
	}
}

func (this *ResultMatrix) AddSum(NodeNr int, sum float64) {
	item := this.getOrCreateResultItem(NodeNr)
	item.Sum = sum
}

func (this *ResultMatrix) AddActivationResult(NodeNr int, actResult float64) {
	item := this.getOrCreateResultItem(NodeNr)
	item.ActivationResult = actResult
}

func (this *ResultMatrix) AddFinalResult(NodeNr int, finalResult float64, isOutput bool) {
	item := this.getOrCreateResultItem(NodeNr)
	item.FinalResult = finalResult
	if isOutput {
		this.Outputs = append(this.Outputs, item)
	}
}

func (this *ResultMatrix) getOrCreateResultItem(NodeNr int) *ResultItem {
	item, ok := this.GetCalculationResult(NodeNr)
	if !ok {
		item = NewResultItem(NodeNr)
		this.CalculationHistory[NodeNr] = item
	}

	return item
}

func (this *ResultMatrix) GetCalculationResult(NodeNr int) (*ResultItem, bool) {
	item, ok := this.CalculationHistory[NodeNr]
	return item, ok
}

func (this *ResultMatrix) GetOutputResult() []*ResultItem {
	return this.Outputs
}

func (this *ResultMatrix) PrintHistory() {
	for _, sum := range this.CalculationHistory {
		fmt.Println(sum.String())
	}
}

func (this *ResultMatrix) PrintOutput() {
	for _, outputResult := range this.Outputs {
		fmt.Println(outputResult.FinalResult)
	}
}