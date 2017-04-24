package matrix

import (
	"github.com/breathbath/nn/matrix"
	"github.com/breathbath/nn/neuron"
)

func NewLinearPerceptronMatrixStub() matrix.Matrix {
	return matrix.Matrix{
		matrix.NewMatrixNode(0, 2, -2.0, 3.0, neuron.TYPE_PERCEPTRON),
		matrix.NewMatrixNode(0, 3, -2.0, 3.0, neuron.TYPE_PERCEPTRON),
		matrix.NewMatrixNode(1, 2, -2.0, 3.0, neuron.TYPE_PERCEPTRON),
		matrix.NewMatrixNode(1, 4, -2.0, 3.0, neuron.TYPE_PERCEPTRON),
		matrix.NewMatrixNode(2, 3, -2.0, 3.0, neuron.TYPE_PERCEPTRON),
		matrix.NewMatrixNode(2, 4, -2.0, 3.0, neuron.TYPE_PERCEPTRON),
		matrix.NewMatrixNode(2, 5, -4.0, 3.0, neuron.TYPE_PERCEPTRON),
		matrix.NewMatrixNode(3, 6, -2.0, 3.0, neuron.TYPE_PERCEPTRON),
		matrix.NewMatrixNode(4, 6, -2.0, 3.0, neuron.TYPE_PERCEPTRON),
	}
}

func NewSigmoidMatrixStub() matrix.Matrix {
	return matrix.Matrix{
		matrix.NewMatrixNode(0, 2, 0.8, 0, neuron.TYPE_SIGMOID),
		matrix.NewMatrixNode(0, 3, 0.4, 0, neuron.TYPE_SIGMOID),
		matrix.NewMatrixNode(0, 4, 0.3, 0, neuron.TYPE_SIGMOID),
		matrix.NewMatrixNode(1, 2, 0.2, 0, neuron.TYPE_SIGMOID),
		matrix.NewMatrixNode(1, 3, 0.9, 0, neuron.TYPE_SIGMOID),
		matrix.NewMatrixNode(1, 4, 0.5, 0, neuron.TYPE_SIGMOID),
		matrix.NewMatrixNode(2, 5, 0.3, 0, neuron.TYPE_SIGMOID),
		matrix.NewMatrixNode(3, 5, 0.5, 0, neuron.TYPE_SIGMOID),
		matrix.NewMatrixNode(4, 5, 0.9, 0, neuron.TYPE_SIGMOID),
	}
}
