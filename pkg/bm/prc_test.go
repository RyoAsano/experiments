package bm

import (
	"math"
	"math/rand"
	"testing"

	"github.com/AsanoRyo/stochastic_calculus/pkg/grd"
	"github.com/AsanoRyo/stochastic_calculus/pkg/mathutil"
	"github.com/AsanoRyo/stochastic_calculus/pkg/randgen"
)

var (
	size         int      = 10
	terminal     float64  = 1.0
	grid         grd.Grid = grd.NewEquiDistGrid(size, terminal)
	mean, stddev float64  = 1.0, 1.0
)

func TestAttr(t *testing.T) {
	var expDim int = 14
	r := *rand.New(rand.NewSource(99))
	fedGen := randgen.NewNorm(r, mean, stddev)
	bm := New(grid, expDim, fedGen, false)

	if bm.Dim() != expDim {
		t.Fatalf("Dimension mismatched.")
	}
	if bm.Grid() != grid {
		t.Fatalf("Wrong with grid setting.")
	}
}

func TestRealize(t *testing.T) {
	r := *rand.New(rand.NewSource(99))
	fedGen := randgen.NewNorm(r, mean, stddev)
	bm := New(grid, 1, fedGen, true)
	p, err := bm.Realize()
	if err != nil {
		t.Fatalf("Failed to realize: " + err.Error())
	}

	if p.Dim() != 2 {
		t.Fatalf("Dimension is not as expected.")
	}

	delta := terminal / float64(size)
	r = *rand.New(rand.NewSource(99))
	gen := randgen.NewNorm(r, mean, stddev)

	var expected float64 = 0
	for k := 0; k <= bm.Grid().Size(); k++ {
		pt, err := p.At(k)
		if err != nil {
			t.Fatalf("Something wrong: " + err.Error())
		}
		if time, err := pt.Pr(1); err != nil {
			t.Fatalf("Failed to get time process.")
		} else if !mathutil.CloseBtw(time, bm.Grid().Get(k)) {
			t.Fatalf("Time process value is wrong.")
		}

		actual, err := pt.Pr(2)
		if err != nil {
			t.Fatalf("Something wrong: " + err.Error())
		}
		if !mathutil.CloseBtw(actual, expected) {
			t.Fatalf("Wrong with the realization impl.")
		}
		expected += gen.Get() * math.Sqrt(delta)
	}
}
