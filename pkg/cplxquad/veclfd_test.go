package cplxquad

import (
	"testing"

	"bitbucket.org/AsanoRyo/experiments/pkg/point"
	"bitbucket.org/AsanoRyo/experiments/pkg/testutil"
)

var driftChecker = testutil.VecFldChecker{
	ExpectedFunc: func(coords []float64) []float64 {
		return []float64{0, 0}
	},
	Samples: []point.Point{
		point.New(0, 0),
		point.New(1000, -100),
	},
}

var diffusionChecker = testutil.VecFldChecker{
	ExpectedFunc: func(coords []float64) []float64 {
		x, y := coords[0], coords[1]
		return []float64{
			x*x - y*y,
			2 * x * y,
		}
	},
	Samples: []point.Point{
		point.New(0, 0),
		point.New(1, 1),
		point.New(-1, -1),
		point.New(2.3, 6.5),
		point.New(1.2, -3.1),
	},
}

func TestVecFld(t *testing.T) {
	driftVec, diffVec := NewVecFld()
	testutil.CheckVecFld(driftVec, driftChecker, t)
	testutil.CheckVecFld(diffVec, diffusionChecker, t)
}
