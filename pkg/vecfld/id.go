package vecfld

import "bitbucket.org/AsanoRyo/stochastic_calculus/pkg/point"

func NewId(dim int) VectorField {
	return &idVecFld{dim: dim}
}

type idVecFld struct {
	dim int
}

var _ VectorField = (*idVecFld)(nil)

func (id *idVecFld) Dims() (int, int) {
	return id.dim, id.dim
}

func (id *idVecFld) At(p point.Point) (point.Point, error) {
	if p.Dim() != id.dim {
		return nil, OutOfDomainErr{Vecfld: id, GivenPoint: p}
	} else {
		return p, nil
	}
}
