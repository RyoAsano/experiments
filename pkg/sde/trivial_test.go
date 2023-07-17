package sde

import (
	"testing"

	"github.com/AsanoRyo/stochastic_calculus/pkg/grd"
	"github.com/AsanoRyo/stochastic_calculus/pkg/point"
	"github.com/AsanoRyo/stochastic_calculus/pkg/stchprc"
)

func TestTrivial(t *testing.T) {
	grid := grd.NewEquiDistGrid(3, 1.0)
	// 3 dim constant process
	actInit := point.New(100, -2.4, 4.0)
	// 4 dim integrator process (which should have no effect)
	intr, err := stchprc.NewDeter(
		grid,
		point.New(0, 0, 0, 0),
		point.New(2, 3, 4, 5),
		point.New(3, 342, 3432, 543),
		point.New(342, 5546, 320, 2.5),
	)
	if err != nil {
		panic(err.Error())
	}
	z := NewTrivial(grid, actInit, intr)

	if z.Dim() != 3 {
		t.Fatalf("Wrong in the process' dimension.")
	}

	if !point.CloseBtw(z.InitPoint(), actInit) {
		t.Fatalf("Wrong in init pt setting.")
	}

	for dim := 1; dim <= 4; dim++ {
		vecFld, err := z.Integrand(dim)
		if err != nil {
			t.Fatalf("Failed to get vec fld: " + err.Error())
		}
		if dom, ran := vecFld.Dims(); dom != 3 || ran != 3 {
			t.Fatalf("Dimensions are wrong.")
		}
		for _, pt := range []point.Point{
			point.New(43, 354252, 324),
			point.New(-3, 2, 0),
		} {
			if actPt, err := vecFld.At(pt); err != nil || !point.CloseBtw(actPt, point.Origin(3)) {
				t.Fatalf("Failed to get a correct pt from vec fld: " + err.Error())
			}
		}
	}
}
