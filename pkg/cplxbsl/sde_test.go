package cplxbsl

import (
	"testing"

	"github.com/AsanoRyo/stochastic_calculus/pkg/grd"
	"github.com/AsanoRyo/stochastic_calculus/pkg/point"
	"github.com/AsanoRyo/stochastic_calculus/pkg/testutil"
)

func TestSDE(t *testing.T) {
	var terminal float64 = 1.0
	var size int = 10
	var x, y float64 = 1.0, 1.0
	var seed int64 = 99

	gen := testutil.NewStdNorm(seed)
	grid := grd.NewEquiDistGrid(size, terminal)

	theSDE := NewSDE(grid, gen, x, y)
	testutil.Check2DimSDE(theSDE, point.New(x, y), driftChecker, diffusionChecker, seed, t)
}
