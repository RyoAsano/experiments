package sample_sde1

import (
	"fmt"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/grd"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/point"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/sde"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/stchprc"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/vecfld"
)

// dX = V(X)dt + W(X)dWt,
// where X is 3-dim process and
// V(X) = s1*X, W(X) = s2*X
func New(
	grid grd.Grid,
	initPt point.Point,
	intr stchprc.Process,
	driftScale float64,
	diffuScale float64,
) (sde.SDE, error) {
	if initPt.Dim() != 3 {
		return nil, fmt.Errorf("Dimension of the initial point must be 3.")
	}
	if intr.Dim() != 2 {
		return nil, fmt.Errorf("Dimension of the integrator must be 2.")
	}
	return &sampleSDE1{
		grid:   grid,
		initPt: initPt,
		intr:   intr,
		drift: vecfld.NewAffine(
			driftScale,
			point.Origin(initPt.Dim()),
		),
		diffusion: vecfld.NewAffine(
			diffuScale,
			point.Origin(initPt.Dim()),
		),
	}, nil
}

type sampleSDE1 struct {
	grid      grd.Grid
	initPt    point.Point
	intr      stchprc.Process
	drift     vecfld.VectorField
	diffusion vecfld.VectorField
}

var _ sde.SDE = (*sampleSDE1)(nil)

func (s *sampleSDE1) Grid() grd.Grid {
	return s.grid
}

func (s *sampleSDE1) Dim() int {
	return 3
}

func (s *sampleSDE1) InitPoint() point.Point {
	return s.initPt
}

func (s *sampleSDE1) Integrand(dim int) (vecfld.VectorField, error) {
	switch dim {
	case 1:
		return s.drift, nil
	case 2:
		return s.diffusion, nil
	default:
		return nil, sde.DimOutOfRangeErr{SDE: s, GivenDim: dim}
	}
}

func (s *sampleSDE1) Integrator() stchprc.Process {
	return s.intr
}
