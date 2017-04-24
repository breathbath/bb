package matrix

type MatrixNode struct {
	Coord        Coord
	Weight, Bias float64
	NodeType     string
}

func NewMatrixNode(x, y int, weight, bias float64, nodeType string) MatrixNode {
	return MatrixNode{NewCoord(x, y), weight, bias, nodeType}
}


