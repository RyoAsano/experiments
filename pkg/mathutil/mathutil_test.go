package mathutil

import (
	"testing"

	"github.com/AsanoRyo/stochastic_calculus/pkg/consts"
)

func TestCloseBtw(t *testing.T) {
	a := 1.0
	closeToA := 1.0 - 0.9*consts.CloseUpTo
	notCloseToA := 1.0 - 1.1*consts.CloseUpTo

	if !CloseBtw(a, closeToA) {
		t.Fatalf("Wrong with the pproximity.")
	}
	if CloseBtw(a, notCloseToA) {
		t.Fatalf("Wrong with the pproximity.")
	}
}
