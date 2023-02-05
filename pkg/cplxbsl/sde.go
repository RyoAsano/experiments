package cplxbsl

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
	gen randgen.RandGenerator,
	x float64,
	y float64,
) sde.SDE {
	drift, diffusion := NewVecFld()
	return &complexBessel{
		initPt:    point.New(x, y),
		grid:      grid,
		drift:     drift,
		diffusion: diffusion,
		intr:      bm.New(grid, 1, gen, true),
	}
}

type complexBessel struct {
	initPt    point.Point
	grid      grd.Grid
	drift     vecfld.VectorField
	diffusion vecfld.VectorField
	intr      stchprc.Process
}

var _ sde.SDE = (*complexBessel)(nil)

func (bsl *complexBessel) Grid() grd.Grid {
	return bsl.grid
}

func (bsl *complexBessel) Dim() int {
	return 2
}

func (bsl *complexBessel) InitPoint() point.Point {
	return bsl.initPt
}

func (bsl *complexBessel) Integrand(dim int) (vecfld.VectorField, error) {
	if dim == 1 {
		return bsl.drift, nil
	} else if dim == 2 {
		return bsl.diffusion, nil
	} else {
		return nil, sde.DimOutOfRangeErr{SDE: bsl, GivenDim: dim}
	}
}

func (bsl *complexBessel) Integrator() stchprc.Process {
	return bsl.intr
}
