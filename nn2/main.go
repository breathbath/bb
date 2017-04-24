package main

import (
	"os"
	_ "github.com/awalterschulze/gographviz"
	"github.com/breathbath/nn/nn2/input"
	"github.com/breathbath/nn/nn2/process"
	"github.com/breathbath/nn/nn2/result"
)

func main() {
	newMatrix := input.NewInputMatrix();
	newMatrix.AddVectors([]*input.Vector{
		input.NewVector(input.BIAS_SOURCE_NEURON_NUMBER, 3, 1),
		input.NewVector(input.BIAS_SOURCE_NEURON_NUMBER, 4, 1),
		input.NewVector(input.BIAS_SOURCE_NEURON_NUMBER, 5, 1),
		input.NewVector(input.BIAS_SOURCE_NEURON_NUMBER, 6, 1),
		input.NewVector(1, 3, 0.2),
		input.NewVector(1, 4, 0.5),
		input.NewVector(1, 5, 0.9),
		input.NewVector(2, 3, 0.8),
		input.NewVector(2, 4, 0.9),
		input.NewVector(2, 5, 0.4),
		input.NewVector(3, 6, 0.3),
		input.NewVector(4, 6, 0.9),
		input.NewVector(5, 6, 0.5),
	})

	resultMatrix := result.NewResultMatrix()
	processor := process.NewProcessor(resultMatrix, newMatrix)
	input1 := input.NewNetworkSignal(1, 1.0)
	input2 := input.NewNetworkSignal(2, 0)
	processor.ForwardPropagation(input1, input2)

	resultMatrix.PrintOutput()
	resultMatrix.PrintHistory()
//
//	g := gographviz.NewGraph()
//	g.SetName("G")
//	g.SetDir(true)
//	g.AddAttr("G", "rankdir", "TB")
//
//	g.AddNode("G", "input1", map[string]string{"shape":"point"})
//	g.AddNode("G", "input2", map[string]string{"shape":"point"})
//	g.AddNode("G", "output", map[string]string{"shape":"point"})
//
//	g.AddNode("G", "0", map[string]string{"label":"\"n:0\"", "style":"dotted", "shape":"ellipse"})
//	g.AddNode("G", "1", map[string]string{"label":"\"n:1\"", "style":"dotted", "shape":"ellipse"})
//	g.AddNode("G", "2", map[string]string{"label":"\"n:2, b:1\n res(i1*w1+i2*w2+b)=0*0.2+0.8*1=0.8\n af(1/(1+e^(-1*res))=1+e^-1*0.8=0.2\"", "shape":"ellipse"})
//	g.AddNode("G", "3", map[string]string{"label":"\"n:3\n b:1\"", "shape":"ellipse"})
//	g.AddNode("G", "4", map[string]string{"label":"\"n:4\n b:1\"", "shape":"ellipse"})
//	g.AddNode("G", "5", map[string]string{"label":"\"n:5\n b:1\"", "shape":"ellipse"})
//
//	g.AddEdge("input1", "0", true, map[string]string{"label":"1", "arrowhead":"empty"})
//	g.AddEdge("input2", "1", true, map[string]string{"label":"0", "arrowhead":"empty"})
//
//	g.AddEdge("0", "2", true, map[string]string{"label":"\"w: 0.8\n i:1\""})
//	g.AddEdge("0", "3", true, map[string]string{"label":"\"w: 0.4\n i:1\""})
//	g.AddEdge("0", "4", true, map[string]string{"label":"\"w: 0.3\n i:1\""})
//
//	g.AddEdge("1", "2", true, map[string]string{"label":"\"w: 0.2\n i:0\""})
//	g.AddEdge("1", "3", true, map[string]string{"label":"\"w: 0.9\n i:0\""})
//	g.AddEdge("1", "4", true, map[string]string{"label":"\"w: 0.5\n i:0\""})
//
//	g.AddEdge("2", "5", true, map[string]string{"label":"\"w: 0.3\n i:0\""})
//	g.AddEdge("3", "5", true, map[string]string{"label":"\"w: 0.5\n i:0\""})
//	g.AddEdge("4", "5", true, map[string]string{"label":"\"w: 0.9\n i:0\""})
//
//	g.AddEdge("5", "output", true, map[string]string{"label":"0.3", "arrowhead":"empty"})
//
//	g.AddAttr("G", "node [shape", "plaintext]")
//	legendString := `
//<<table border="0" cellpadding="2" cellspacing="0" cellborder="0">
//<tr><td align="left">n:node nr</td></tr>
//<tr><td align="left">w: weight</td></tr>
//<tr><td align="left">b: bias</td></tr>
//<tr><td align="left">i: input</td></tr>
//<tr><td align="left">res: result_function</td></tr>
//<tr><td align="left">af: activation_function</td></tr>
//</table>>`
//	subGraphAttrs := map[string]string{
//		"key [label": legendString + "]",
//	}
//	g.AddSubGraph("G", "cluster_01", subGraphAttrs)
//
//	paths := render.NewDefaultFilePaths()
//	err := render.NewDotFileGenerator(paths).GenerateDotFile(g)
//	errors.HandleError(err)
//
//	dotConverter := render.NewDotFileConverter(paths)
//
//	generatedFilePath, err := dotConverter.ConvertDotToImage()
//	errors.HandleError(err)
//	fmt.Println("Generated output file: " + generatedFilePath)
}

func isVerbose() bool {
	for _, param := range os.Args {
		if param == "vv" || param == "v" {
			return true
		}
	}
	return false
}
