package process

import "math"

type Sigmoid struct {

}

func (this Sigmoid) Activate(input float64) float64 {
	return 1 / (1 + math.Exp(-1 * input))
}