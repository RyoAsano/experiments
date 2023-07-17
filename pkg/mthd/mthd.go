package mthd

import (
	"github.com/RyoAsano/stochastic_calculus/pkg/grd"
	"github.com/RyoAsano/stochastic_calculus/pkg/path"
	"github.com/RyoAsano/stochastic_calculus/pkg/point"
	"github.com/RyoAsano/stochastic_calculus/pkg/sde"
	"github.com/RyoAsano/stochastic_calculus/pkg/stchprc"
	"github.com/RyoAsano/stochastic_calculus/pkg/vecfld"
)

type DX map[vecfld.VectorField]float64

type Method interface {
	To(from point.Point, dx DX) (point.Point, error)
	Apply(sde sde.SDE) stchprc.Process
	Modify(point point.Point) point.Point
}

type mthdAppliedPrc struct {
	mthd Method
	sde  sde.SDE
}

var _ stchprc.Process = (*mthdAppliedPrc)(nil)

func (prc *mthdAppliedPrc) Dim() int {
	return prc.sde.Dim()
}

func (prc *mthdAppliedPrc) Grid() grd.Grid {
	return prc.sde.Grid()
}

func (prc *mthdAppliedPrc) Realize() (path.Path, error) {
	intr, err := prc.sde.Integrator().Realize()
	if err != nil {
		return nil, err
	}

	var runningPt point.Point = prc.sde.InitPoint()
	pathGen := path.NewGenerator(prc.Grid(), prc.sde.Dim())

	pathGen.Set(0, prc.mthd.Modify(runningPt))

	for k := 1; k <= prc.Grid().Size(); k++ {
		// (dx1, dx2, ..., dxn)
		intrIncrs, err := intr.Incr(k-1, k)
		if err != nil {
			return nil, err
		}

		// (V1, dx1), (V2, dx2), ..., (Vn, dxn)
		incrs := map[vecfld.VectorField]float64{}
		for dim := 1; dim <= intr.Dim(); dim++ {
			vecFld, err := prc.sde.Integrand(dim)
			if err != nil {
				return nil, err
			}
			intrIncr, err := intrIncrs.Pr(dim)
			if err != nil {
				return nil, err
			}
			incrs[vecFld] = intrIncr
		}

		runningPt, err = prc.mthd.To(runningPt, incrs)
		if err != nil {
			return nil, err
		}
		pathGen.Set(k, prc.mthd.Modify(runningPt))
	}

	return pathGen.Generate()
}
