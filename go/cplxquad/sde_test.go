package cplxquad

import (
	"testing"

	"bitbucket.org/AsanoRyo/experiments/grd"
	"bitbucket.org/AsanoRyo/experiments/point"
	"bitbucket.org/AsanoRyo/experiments/testutil"
)

var seed int64 = 99

func TestSDE(t *testing.T) {
	terminal := 1.0
	size := 10
	grid := grd.NewEquiDistGrid(size, terminal)
	x, y := 1.0, 1.0
	theSDE := NewSDE(grid, x, y, testutil.NewStdNorm(seed))

	testutil.Check2DimSDE(theSDE, point.New(x, y), driftChecker, diffusionChecker, seed, t)
}
