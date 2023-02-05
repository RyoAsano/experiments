package cplxbsl

import (
	"bitbucket.org/AsanoRyo/experiments/pkg/point"
	"bitbucket.org/AsanoRyo/experiments/pkg/vecfld"
)

type drift struct{}

var _ vecfld.VectorField = (*drift)(nil)

func (bessel *drift) Dims() (int, int) {
	return 2, 2
}

func (bessel *drift) At(p point.Point) (point.Point, error) {
	xi, err := p.Pr(1)
	if err != nil {
		return nil, err
	}
	eta, err := p.Pr(2)
	if err != nil {
		return nil, err
	}
	return point.New(xi/(xi*xi+eta*eta), -eta/(xi*xi+eta*eta)), nil
}

type diffusion struct{}

var _ vecfld.VectorField = (*diffusion)(nil)

func (bessel *diffusion) Dims() (int, int) {
	return 2, 2
}

func (bessel *diffusion) At(p point.Point) (point.Point, error) {
	return point.New(-1, 0), nil
}

func NewVecFld() (vecfld.VectorField, vecfld.VectorField) {
	return &drift{}, &diffusion{}
}
