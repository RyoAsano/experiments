package bs

import (
	"testing"

	"bitbucket.org/AsanoRyo/experiments/pkg/grd"
	"bitbucket.org/AsanoRyo/experiments/pkg/point"
	"bitbucket.org/AsanoRyo/experiments/pkg/testutil"
)

func TestSDE(t *testing.T) {
	size := 10
	terminal := 1.0
	grid := grd.NewEquiDistGrid(size, terminal)
	initpt := 2.0
	mu, sigma := 2.0, .5
	var seed int64 = 99
	bs := NewSDE(grid, testutil.NewStdNorm(seed), initpt, mu, sigma)

	driftCheker := testutil.VecFldChecker{
		ExpectedFunc: func(coords []float64) []float64 {
			return []float64{coords[0] * mu}
		},
		Samples: []point.Point{
			point.New(0),
			point.New(1),
			point.New(-3.4),
		},
	}

	diffusionChecker := testutil.VecFldChecker{
		ExpectedFunc: func(coords []float64) []float64 {
			return []float64{coords[0] * sigma}
		},
		Samples: []point.Point{
			point.New(0),
			point.New(1),
			point.New(-3.4),
		},
	}

	testutil.Check2DimSDE(bs, point.New(initpt), driftCheker, diffusionChecker, seed, t)
}
