package render

import (
	"github.com/awalterschulze/gographviz"
	"github.com/breathbath/nn/matrix"
	"github.com/breathbath/nn/neuron"
)

type NetRenderer struct {

}

func NewNetRenderer() NetRenderer {
	return NetRenderer{}
}

func (this NetRenderer) generateGraph(m matrix.Matrix) *gographviz.Graph {
	g := gographviz.NewGraph()
	g.SetName("G")
	g.SetDir(true)
	g.AddAttr("G", "rankdir", "TB")

	notResultNeurons := make(map[int]int)
	notInputNeurons := make(map[int]int)
	allNeurons := make(map[int]matrix.MatrixNode)

	for _, matrixItem := range m {
		notResultNeurons[matrixItem.Coord.X] = 0
		notInputNeurons[matrixItem.Coord.Y] = 0

		allNeurons[matrixItem.Coord.X] = matrixItem
		allNeurons[matrixItem.Coord.Y] = matrixItem
	}

	for _, matrixItem := range m {

	}

	return g
}

func (this NetRenderer) isInputNeuron(m matrix.Matrix) {

}
