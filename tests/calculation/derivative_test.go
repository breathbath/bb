package calculation

import (
	"testing"
	"github.com/breathbath/nn/calculation"
	"math"
)

func linearFunc (x float64) float64 {
	return  2.0 - x / 2.0
}

func squareFunc(x float64) float64 {
	return  math.Pow(x, 2)
}

func squareRoot(x float64) float64 {
	return math.Sqrt(x)
}

func sigmoidFunc (x float64) float64 {
	return 1.0 / (1.0 + math.Pow(math.E, -1.0 * x))
}


func TestDerivativeResults(t *testing.T) {
	var x float64

	x = 4.0
	linearDerivat := -0.5 //(2.0 - 1/2 * x)' = (2.0)' - 1/2(x)' = 0 - 1/2 * 1 = -1/2
	evaluateCalculationResult(x, linearDerivat, linearFunc, t)

	x = 4.0
	squareDerivat := 2.0 * x // (x^2)' = 2.0 * x
	evaluateCalculationResult(x, squareDerivat, squareFunc, t)

	x = 4.0
	squareRootDerivat := 1 / (2.0 * math.Sqrt(x)) //Vx = 1/(2. * Vx) see http://www.webmath.ru/poleznoe/formules_10_19.php
	evaluateCalculationResult(x, squareRootDerivat, squareRoot, t)

	x = 5.0
	sigmoidResult := sigmoidFunc(x)
	sigmoidDerivat := sigmoidResult * (1.0 - sigmoidResult)
	evaluateCalculationResult(x, sigmoidDerivat, sigmoidFunc, t)
}

func evaluateCalculationResult(x, expectedResult float64, functionToCheck calculation.DerivativeFunction, t *testing.T) {
	gotResult := calculation.CalculateDerivative(x, functionToCheck)
	if expectedResult -  gotResult > 1e-6 {
		t.Errorf("Derivative function should return %.50f but it returns %.50f, diff %.10f", expectedResult, gotResult, expectedResult -  gotResult)
	}
}
