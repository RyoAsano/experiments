package cplxquad

import (
	"github.com/AsanoRyo/stochastic_calculus/pkg/point"
	"github.com/AsanoRyo/stochastic_calculus/pkg/vecfld"
)

func NewVecFld() (vecfld.VectorField, vecfld.VectorField) {
	return &drift{}, &diffusion{}
}

type drift struct{}

var _ vecfld.VectorField = (*drift)(nil)

func (d *drift) Dims() (int, int) {
	return 2, 2
}

func (d *drift) At(p point.Point) (point.Point, error) {
	if from, _ := d.Dims(); p.Dim() != from {
		return nil, vecfld.OutOfDomainErr{Vecfld: d, GivenPoint: p}
	}
	return point.New(0, 0), nil
}

type diffusion struct{}

var _ vecfld.VectorField = (*diffusion)(nil)

func (diff *diffusion) Dims() (int, int) {
	return 2, 2
}

func (diff *diffusion) At(p point.Point) (point.Point, error) {
	if from, _ := diff.Dims(); p.Dim() != from {
		return nil, vecfld.OutOfDomainErr{Vecfld: diff, GivenPoint: p}
	}
	x, err := p.Pr(1)
	if err != nil {
		return nil, err
	}
	y, err := p.Pr(2)
	if err != nil {
		return nil, err
	}
	return point.New(x*x-y*y, 2.0*x*y), nil
}
