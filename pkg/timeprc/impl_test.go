package timeprc

import (
	"testing"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/grd"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/mathutil"
)

func TestTime(t *testing.T) {
	grid := grd.NewEquiDistGrid(10, 2.4)
	prc := New(grid)

	if prc.Dim() != 1 {
		t.Fatalf("Dimension must be 1.")
	}
	if prc.Grid() != grid {
		t.Fatalf("Grid setting is wrong.")
	}
	pth, err := prc.Realize()
	if err != nil {
		t.Fatalf("Failed to realize: " + err.Error())
	}
	for k := 0; k <= grid.Size(); k++ {
		expVal := grid.Get(k)
		if actPt, err := pth.At(k); err != nil {
			t.Fatalf("Failed to get a point from a path.")
		} else {
			actVal, err := actPt.Pr(1)
			if err != nil {
				t.Fatalf("Dimension is wrong: " + err.Error())
			}
			if !mathutil.CloseBtw(actVal, expVal) {
				t.Fatalf("Value setting is wrong.")
			}
		}
	}
}
