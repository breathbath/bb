package matrix

import (
	"github.com/breathbath/nn/log"
	"github.com/breathbath/nn/observers"
	"github.com/breathbath/nn/neuron"
	"github.com/breathbath/nn/net"
	"github.com/breathbath/nn/vector"
)

type Matrix []MatrixNode

func (m Matrix) ConvertToNet(logger log.LogBehaviour, resultSubscriber observers.Observer) net.Net {
	outputNeurons := make(map[int]MatrixNode)
	inputNeurons := make(map[int] neuron.NeuronBehaviour)

	neuronNet := net.NewNet(resultSubscriber, logger)
	neuronFactory := neuron.NewNeuronFactory(logger, neuronNet)

	var neuronFrom, neuronTo neuron.NeuronBehaviour
	for _, matrixItem := range m {
		outputNeurons[matrixItem.Coord.Y] = matrixItem

		delete(outputNeurons, matrixItem.Coord.X)
		delete(inputNeurons, matrixItem.Coord.Y)

		neuronFrom = neuronFactory.CreateNeuron(matrixItem.Bias, matrixItem.Coord.X, matrixItem.NodeType)
		neuronTo = neuronFactory.CreateNeuron(matrixItem.Bias, matrixItem.Coord.Y, matrixItem.NodeType)

		curVector := vector.NewVector(neuronFrom, neuronTo, matrixItem.Weight, logger)

		neuronTo.IncreaseExpectedCount()
		neuronNet.AddVector(&curVector)

		inputNeurons[matrixItem.Coord.X] = neuronFrom
	}

	for _, matrixItem := range outputNeurons {
		neuronFrom := neuronFactory.CreateNeuron(0, matrixItem.Coord.Y, matrixItem.NodeType)
		resVector := vector.NewResultVector(neuronFrom, 0, logger)
		neuronNet.AddVector(&resVector)
	}

	return neuronNet
}