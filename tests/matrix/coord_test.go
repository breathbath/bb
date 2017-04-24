package matrix

import (
	"testing"
	"github.com/breathbath/nn/matrix"
)

func TestCoordStringer(t *testing.T) {
	coord := matrix.NewCoord(10, 20)
	stringResult := coord.String()
	if stringResult != "(10, 20)" {
		t.Errorf("Unexpected stringer result %s, expected %s", stringResult, "(10, 20)")
	}
}