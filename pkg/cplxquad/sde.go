package cplxquad

import (
	"bitbucket.org/AsanoRyo/experiments/pkg/bm"
	"bitbucket.org/AsanoRyo/experiments/pkg/grd"
	"bitbucket.org/AsanoRyo/experiments/pkg/point"
	"bitbucket.org/AsanoRyo/experiments/pkg/randgen"
	"bitbucket.org/AsanoRyo/experiments/pkg/sde"
	"bitbucket.org/AsanoRyo/experiments/pkg/stchprc"
	"bitbucket.org/AsanoRyo/experiments/pkg/vecfld"
)

func NewSDE(
	grid grd.Grid,
	x float64,
	y float64,
	gen randgen.RandGenerator,
) sde.SDE {
	drift, diffusion := NewVecFld()
	return &quad{
		grid:      grid,
		initPt:    point.New(x, y),
		drift:     drift,
		diffusion: diffusion,
		intr:      bm.New(grid, 1, gen, true),
	}
}

type quad struct {
	grid      grd.Grid
	initPt    point.Point
	drift     vecfld.VectorField
	diffusion vecfld.VectorField
	intr      stchprc.Process
}

var _ sde.SDE = (*quad)(nil)

func (q *quad) Grid() grd.Grid {
	return q.grid
}

func (q *quad) Dim() int {
	return 2
}

func (q *quad) InitPoint() point.Point {
	return q.initPt
}

func (q *quad) Integrator() stchprc.Process {
	return q.intr
}

func (q *quad) Integrand(dim int) (vecfld.VectorField, error) {
	if dim == 1 {
		return q.drift, nil
	} else if dim == 2 {
		return q.diffusion, nil
	} else {
		return nil, sde.DimOutOfRangeErr{SDE: q, GivenDim: dim}
	}
}
