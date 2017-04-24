package observers

import (
	"github.com/breathbath/nn/observers"
)

type ObserverMock struct {
	Results [] observers.ObservingResult
}
func (o *ObserverMock) AcceptResult(result observers.ObservingResult) {
	o.Results = append(o.Results, result)
}

func NewObserverMock() *ObserverMock {
	return &ObserverMock {[]observers.ObservingResult{}}
}