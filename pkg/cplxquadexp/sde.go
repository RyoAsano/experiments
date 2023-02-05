package cplxquadexp

import (
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/bm"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/grd"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/point"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/randgen"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/sde"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/stchprc"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/vecfld"
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
