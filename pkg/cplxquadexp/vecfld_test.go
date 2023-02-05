package cplxquadexp

import (
	"testing"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/point"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/testutil"
)

var driftChecker = testutil.VecFldChecker{
	ExpectedFunc: func(coords []float64) []float64 {
		x := coords[0]
		return []float64{
			0,
			-2.0 * x * x,
		}
	},
	Samples: []point.Point{
		point.New(0, 0),
		point.New(1.0, 1.0),
		point.New(-1.0, -1.0),
		point.New(4.0, -2.0),
		point.New(5.2, 2.0),
	},
}

var diffusionChecker = testutil.VecFldChecker{
	ExpectedFunc: func(coords []float64) []float64 {
		x, y := coords[0], coords[1]
		return []float64{
			x*x - y*y,
			2 * x,
		}
	},
	Samples: []point.Point{
		point.New(0, 0),
		point.New(1.0, 1.0),
		point.New(-1.0, -1.0),
		point.New(4.0, -2.0),
		point.New(5.2, 2.0),
	},
}

func TestVecFld(t *testing.T) {
	driftVec, diffVec := NewVecFld()
	testutil.CheckVecFld(driftVec, driftChecker, t)
	testutil.CheckVecFld(diffVec, diffusionChecker, t)
}
