package mathutil

import (
	"testing"

	"bitbucket.org/AsanoRyo/experiments/consts"
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
