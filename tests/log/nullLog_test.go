package log

import (
	"testing"
	"github.com/breathbath/nn/log"
)

func TestNullLog(t *testing.T) {
	nlog := log.NullLog{}
	nlog.Log("Some name")
	nlog.Logf("Somename", 1, 2, 3)
}
