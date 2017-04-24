package matrix

import (
	"testing"
	"github.com/breathbath/nn/matrix"
)

func TestMatrixNodeCreation(t *testing.T) {
	matrix.NewMatrixNode(0, 0, 1.0, 2.0, "someType")
}
