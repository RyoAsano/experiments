package cplxquadexp

import (
	"bitbucket.org/AsanoRyo/experiments/bm"
	"bitbucket.org/AsanoRyo/experiments/grd"
	"bitbucket.org/AsanoRyo/experiments/point"
	"bitbucket.org/AsanoRyo/experiments/randgen"
	"bitbucket.org/AsanoRyo/experiments/sde"
	"bitbucket.org/AsanoRyo/experiments/stchprc"
	"bitbucket.org/AsanoRyo/experiments/vecfld"
)

func NewSDE(
	grid grd.Grid,
	x float64,
	gen randgen.RandGenerator,
) sde.SDE {
	drift, diffusion := NewVecFld()
	return &quadExp{
		grid:      grid,
		initPt:    point.New(x, .0),
		intr:      bm.New(grid, 1, gen, true),
		drift:     drift,
		diffusion: diffusion,
	}
}

type quadExp struct {
	grid      grd.Grid
	initPt    point.Point
	intr      stchprc.Process
	drift     vecfld.VectorField
	diffusion vecfld.VectorField
}

var _ sde.SDE = (*quadExp)(nil)

func (q *quadExp) Grid() grd.Grid {
	return q.grid
}

func (q *quadExp) Dim() int {
	return 2
}

func (q *quadExp) InitPoint() point.Point {
	return q.initPt
}

func (q *quadExp) Integrand(dim int) (vecfld.VectorField, error) {
	if dim == 1 {
		return q.drift, nil
	} else if dim == 2 {
		return q.diffusion, nil
	} else {
		return nil, sde.DimOutOfRangeErr{SDE: q, GivenDim: dim}
	}
}

func (q *quadExp) Integrator() stchprc.Process {
	return q.intr
}
