package cplxquadexp

import (
	"testing"

	"github.com/AsanoRyo/stochastic_calculus/pkg/grd"
	"github.com/AsanoRyo/stochastic_calculus/pkg/point"
	"github.com/AsanoRyo/stochastic_calculus/pkg/testutil"
)

func TestSDE(t *testing.T) {
	terminal := 1.0
	size := 10
	grid := grd.NewEquiDistGrid(size, terminal)
	x := 3.0
	var seed int64 = 99

	theSDE := NewSDE(grid, x, testutil.NewStdNorm(seed))
	testutil.Check2DimSDE(theSDE, point.New(x, 0), driftChecker, diffusionChecker, seed, t)
}
