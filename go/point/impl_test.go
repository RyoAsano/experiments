package point

import (
	"testing"

	"bitbucket.org/AsanoRyo/experiments/mathutil"
)

func TestImpl(t *testing.T) {
	pt := point{dim: 1, coord: []float64{2.4}}

	if pt.Dim() != 1 {
		t.Fatalf("Dimension is wrong.")
	}

	if actVal, err := pt.Pr(1); err != nil {
		t.Fatalf("Failed to get a value.")
	} else {
		if !mathutil.CloseBtw(actVal, 2.4) {
			t.Fatalf("Value setting is wrong.")
		}
	}

	if !mathutil.CloseBtw(pt.Norm(), 2.4) {
		t.Fatalf("Norm impl is wrong.")
	}

	if len(pt.coord) != 1 {
		t.Fatalf("The point's coord length is wrong.")
	}
	if val := pt.coord[0]; !mathutil.CloseBtw(val, 2.4) {
		t.Fatalf("Coordinate value seems changed.")
	}
}
