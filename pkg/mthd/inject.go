package mthd

import (
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/point"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/sde"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/stchprc"
)

type InjMap func(point.Point) point.Point

type injMthd struct {
	mthd   Method
	injMap InjMap
}

var _ Method = (*injMthd)(nil)

func (m *injMthd) Apply(sde sde.SDE) stchprc.Process {
	return &mthdAppliedPrc{mthd: m, sde: sde}
}

func (m *injMthd) To(from point.Point, dx DX) (point.Point, error) {
	return m.mthd.To(from, dx)
}

func (m *injMthd) Modify(point point.Point) point.Point {
	return m.injMap(point)
}

func Inject(method Method, injMap InjMap) Method {
	return &injMthd{mthd: method, injMap: injMap}
}
