package result

import "fmt"

type ResultItem struct {
	Column           int
	Sum              float64
	ActivationResult float64
	FinalResult      float64
}

func NewResultItem(res int) *ResultItem {
	return &ResultItem{Column:res}
}

func (this *ResultItem) String() string {
	return fmt.Sprintf("Column: %d, Sum: %.5f, ActResult: %.5f, Final: %.5f", this.Column, this.Sum, this.ActivationResult, this.FinalResult)
}
