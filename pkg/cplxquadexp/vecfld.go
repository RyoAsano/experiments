package cplxquadexp

import (
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/point"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/vecfld"
)

func NewVecFld() (vecfld.VectorField, vecfld.VectorField) {
	return &drift{}, &diffusion{}
}

type drift struct{}

var _ vecfld.VectorField = (*drift)(nil)

func (q *drift) Dims() (int, int) {
	return 2, 2
}

func (q *drift) At(p point.Point) (point.Point, error) {
	if from, _ := q.Dims(); p.Dim() != from {
		return nil, vecfld.OutOfDomainErr{Vecfld: q, GivenPoint: p}
	}
	x, err := p.Pr(1)
	if err != nil {
		return nil, err
	}
	return point.New(0, -2.0*x*x), nil
}

type diffusion struct{}

var _ vecfld.VectorField = (*diffusion)(nil)

func (q *diffusion) Dims() (int, int) {
	return 2, 2
}

func (q *diffusion) At(p point.Point) (point.Point, error) {
	if from, _ := q.Dims(); p.Dim() != from {
		return nil, vecfld.OutOfDomainErr{Vecfld: q, GivenPoint: p}
	}
	x, err := p.Pr(1)
	if err != nil {
		return nil, err
	}
	y, err := p.Pr(2)
	if err != nil {
		return nil, err
	}
	return point.New(x*x-y*y, 2.0*x), nil
}
