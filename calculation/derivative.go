package calculation

import "math"

func CalculateDerivative(x float64, derivatFunc DerivativeFunction) float64 {
	deltaLimit := 1e-20
	deltaX := 0.01
	var derivat2, derivatDelta float64
	derivat1 := calculateDerivativeInX(x, deltaX, derivatFunc)
	for {
		deltaX = deltaX / 2
		derivat2 = calculateDerivativeInX(x, deltaX, derivatFunc)
		derivatDelta = math.Abs(derivat2 - derivat1)
		derivat1 = derivat2
		if derivatDelta < deltaLimit {
			break
		}
	}
	return derivat2
}

func calculateDerivativeInX(x, deltaX float64, derivatFunc DerivativeFunction) float64 {
	return (derivatFunc(x + deltaX)-derivatFunc(x))/deltaX
}
