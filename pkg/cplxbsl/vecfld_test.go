package cplxbsl

import (
	"math"
	"testing"

	"github.com/RyoAsano/stochastic_calculus/pkg/point"
	"github.com/RyoAsano/stochastic_calculus/pkg/testutil"
)

var driftChecker = testutil.VecFldChecker{
	ExpectedFunc: func(coord []float64) []float64 {
		x, y := coord[0], coord[1]
		return []float64{
			x / (math.Pow(x, 2) + math.Pow(y, 2)),
			-y / (math.Pow(x, 2) + math.Pow(y, 2)),
		}
	},
	Samples: []point.Point{
		point.New(1, 1.5),
		point.New(1, -1),
		point.New(-0.3, .4),
		point.New(-5, -.9),
	},
}

var diffusionChecker = testutil.VecFldChecker{
	ExpectedFunc: func(coord []float64) []float64 {
		return []float64{-1, 0}
	},
	Samples: []point.Point{
		point.New(1000, -2200),
		point.New(0, 0),
	},
}

func TestVectorFields(t *testing.T) {
	driftVecFld, diffVecFld := NewVecFld()
	testutil.CheckVecFld(driftVecFld, driftChecker, t)
	testutil.CheckVecFld(diffVecFld, diffusionChecker, t)
}
