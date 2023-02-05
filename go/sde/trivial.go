package sde

import (
	"bitbucket.org/AsanoRyo/experiments/grd"
	"bitbucket.org/AsanoRyo/experiments/point"
	"bitbucket.org/AsanoRyo/experiments/stchprc"
	"bitbucket.org/AsanoRyo/experiments/vecfld"
)

func NewTrivial(
	grid grd.Grid,
	initPt point.Point,
	intr stchprc.Process,
) SDE {
	zeroVecFlds := make([]vecfld.VectorField, intr.Dim())
	for k := 0; k < intr.Dim(); k++ {
		zeroVecFlds[k] = vecfld.NewZero(initPt.Dim(), initPt.Dim())
	}
	return constSDE{
		grid:        grid,
		initPt:      initPt,
		zeroVecFlds: zeroVecFlds,
		intr:        intr,
	}
}

type constSDE struct {
	grid        grd.Grid
	initPt      point.Point
	zeroVecFlds []vecfld.VectorField
	intr        stchprc.Process
}

var _ SDE = (*constSDE)(nil)

func (c constSDE) Grid() grd.Grid {
	return c.grid
}

func (c constSDE) Dim() int {
	return c.initPt.Dim()
}

func (c constSDE) InitPoint() point.Point {
	return c.initPt
}

func (c constSDE) Integrand(dim int) (vecfld.VectorField, error) {
	index := dim - 1
	if 0 <= index && index < len(c.zeroVecFlds) {
		return c.zeroVecFlds[index], nil
	} else {
		return nil, DimOutOfRangeErr{SDE: c, GivenDim: dim}
	}
}

func (c constSDE) Integrator() stchprc.Process {
	return c.intr
}
