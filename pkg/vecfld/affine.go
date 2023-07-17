package vecfld

import (
	"github.com/AsanoRyo/stochastic_calculus/pkg/point"
)

func NewAffine(scale float64, adder point.Point) VectorField {
	return &affine{
		scale: scale,
		adder: adder,
	}
}

type affine struct {
	scale float64
	adder point.Point
}

var _ VectorField = (*affine)(nil)

func (v affine) Dims() (int, int) {
	return v.adder.Dim(), v.adder.Dim()
}

func (v affine) At(p point.Point) (point.Point, error) {
	if from, _ := v.Dims(); p.Dim() != from {
		return nil, OutOfDomainErr{Vecfld: v, GivenPoint: p}
	}
	return point.Add(point.Mul(p, v.scale), v.adder)
}
