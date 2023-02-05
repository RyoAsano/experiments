package sample_sde1

import (
	"testing"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/grd"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/point"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/stchprc"
)

func Test(t *testing.T) {
	grid := grd.NewEquiDistGrid(2, 1.0)
	initPt := point.New(3, 1, 4)
	driftScale, diffuScale := 2.3, 4.3
	dummyPt := point.Origin(2)
	intr, err := stchprc.NewDeter(grid, dummyPt, dummyPt, dummyPt)
	if err != nil {
		panic(err.Error())
	}

	s, err := New(grid, initPt, intr, driftScale, diffuScale)
	if err != nil {
		t.Fatalf("Failed to get SDE: " + err.Error())
	}

	if s.Grid() != grid {
		t.Fatalf("Grid setting is wrong.")
	}
	if !point.CloseBtw(s.InitPoint(), initPt) {
		t.Fatalf("Initial point setting is wrong.")
	}
	if s.Integrator() != intr {
		t.Fatalf("Integrator's setting is wrong.")
	}

	drift, err := s.Integrand(1)
	if err != nil {
		t.Fatalf("Failed to get drift: " + err.Error())
	}
	diffu, err := s.Integrand(2)
	if err != nil {
		t.Fatalf("Failed to get diffusion: " + err.Error())
	}

	for _, pt := range []point.Point{
		point.New(1, 3, 4),
		point.New(3, 4, 4),
		point.New(-100, 3.5, 4),
	} {
		// drift
		if actPt, err := drift.At(pt); err != nil {
			t.Fatalf("Failed to ge point by drift: " + err.Error())
		} else {
			expPt := point.Mul(pt, driftScale)
			if !point.CloseBtw(actPt, expPt) {
				t.Fatalf("Someting wrong with drift impl")
			}
		}
		// diffusion
		if actPt, err := diffu.At(pt); err != nil {
			t.Fatalf("Failed to ge point by diffusion: " + err.Error())
		} else {
			expPt := point.Mul(pt, diffuScale)
			if !point.CloseBtw(actPt, expPt) {
				t.Fatalf("Someting wrong with diffusion impl")
			}
		}
	}
}
