package main

import (
	"os"
	"github.com/breathbath/nn/log"
	"github.com/breathbath/nn/matrix"
	"github.com/breathbath/nn/observers"
	"github.com/breathbath/nn/signal"
	"github.com/breathbath/nn/neuron"
	"github.com/awalterschulze/gographviz"
	"fmt"
	"github.com/breathbath/nn/render"
	"github.com/breathbath/nn/errors"
)

func main() {
	var logger log.LogBehaviour
	if isVerbose() {
		logger = log.TtlLog{}
	} else {
		logger = log.NullLog{}
	}

	currMatrix := matrix.Matrix{
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

	curNet := currMatrix.ConvertToNet(logger, observers.NetworkResultOutputter{})

	curNet.ForwardPropagation(signal.NewInputSignal(0.0, 0), signal.NewInputSignal(1.0, 1))

	g := gographviz.NewGraph()
	g.SetName("G")
	g.SetDir(true)
	g.AddAttr("G", "rankdir", "TB")

	g.AddNode("G", "input1", map[string]string{"shape":"point"})
	g.AddNode("G", "input2", map[string]string{"shape":"point"})
	g.AddNode("G", "output", map[string]string{"shape":"point"})

	g.AddNode("G", "0", map[string]string{"label":"\"n:0\"", "style":"dotted", "shape":"ellipse"})
	g.AddNode("G", "1", map[string]string{"label":"\"n:1\"", "style":"dotted", "shape":"ellipse"})
	g.AddNode("G", "2", map[string]string{"label":"\"n:2, b:1\n res(i1*w1+i2*w2+b)=0*0.2+0.8*1=0.8\n af(1/(1+e^(-1*res))=1+e^-1*0.8=0.2\"", "shape":"ellipse"})
	g.AddNode("G", "3", map[string]string{"label":"\"n:3\n b:1\"", "shape":"ellipse"})
	g.AddNode("G", "4", map[string]string{"label":"\"n:4\n b:1\"", "shape":"ellipse"})
	g.AddNode("G", "5", map[string]string{"label":"\"n:5\n b:1\"", "shape":"ellipse"})

	g.AddEdge("input1", "0", true, map[string]string{"label":"1", "arrowhead":"empty"})
	g.AddEdge("input2", "1", true, map[string]string{"label":"0", "arrowhead":"empty"})

	g.AddEdge("0", "2", true, map[string]string{"label":"\"w: 0.8\n i:1\""})
	g.AddEdge("0", "3", true, map[string]string{"label":"\"w: 0.4\n i:1\""})
	g.AddEdge("0", "4", true, map[string]string{"label":"\"w: 0.3\n i:1\""})

	g.AddEdge("1", "2", true, map[string]string{"label":"\"w: 0.2\n i:0\""})
	g.AddEdge("1", "3", true, map[string]string{"label":"\"w: 0.9\n i:0\""})
	g.AddEdge("1", "4", true, map[string]string{"label":"\"w: 0.5\n i:0\""})

	g.AddEdge("2", "5", true, map[string]string{"label":"\"w: 0.3\n i:0\""})
	g.AddEdge("3", "5", true, map[string]string{"label":"\"w: 0.5\n i:0\""})
	g.AddEdge("4", "5", true, map[string]string{"label":"\"w: 0.9\n i:0\""})

	g.AddEdge("5", "output", true, map[string]string{"label":"0.3", "arrowhead":"empty"})

	g.AddAttr("G", "node [shape", "plaintext]")
	legendString := `
<<table border="0" cellpadding="2" cellspacing="0" cellborder="0">
<tr><td align="left">n:node nr</td></tr>
<tr><td align="left">w: weight</td></tr>
<tr><td align="left">b: bias</td></tr>
<tr><td align="left">i: input</td></tr>
<tr><td align="left">res: result_function</td></tr>
<tr><td align="left">af: activation_function</td></tr>
</table>>`
	subGraphAttrs := map[string]string{
		"key [label": legendString + "]",
	}
	g.AddSubGraph("G", "cluster_01", subGraphAttrs)

	paths := render.NewDefaultFilePaths()
	err := render.NewDotFileGenerator(paths).GenerateDotFile(g)
	errors.HandleError(err)

	dotConverter := render.NewDotFileConverter(paths)

	generatedFilePath, err  := dotConverter.ConvertDotToImage()
	errors.HandleError(err)
	fmt.Println("Generated output file: " +  generatedFilePath)
}

func isVerbose() bool {
	for _, param := range os.Args {
		if param == "vv" || param == "v" {
			return true
		}
	}
	return false
}
