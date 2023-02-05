package vecfld

import "bitbucket.org/AsanoRyo/stochastic_calculus/pkg/point"

func NewZero(domainDim int, rangeDim int) VectorField {
	return &zeroVecFld{domainDim: domainDim, rangeDim: rangeDim}
}

type zeroVecFld struct {
	domainDim int
	rangeDim  int
}

var _ VectorField = (*zeroVecFld)(nil)

func (z *zeroVecFld) Dims() (int, int) {
	return z.domainDim, z.rangeDim
}

func (z *zeroVecFld) At(p point.Point) (point.Point, error) {
	if p.Dim() != z.domainDim {
		return nil, OutOfDomainErr{Vecfld: z, GivenPoint: p}
	}
	return point.Origin(z.rangeDim), nil
}
