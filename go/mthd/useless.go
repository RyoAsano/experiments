package mthd

import (
	"fmt"

	"bitbucket.org/AsanoRyo/experiments/point"
	"bitbucket.org/AsanoRyo/experiments/sde"
	"bitbucket.org/AsanoRyo/experiments/stchprc"
)

func NewPtToPtMthd(checkPoints ...point.Point) Method {
	return &ptToPtMthd{checkPoints: checkPoints}
}

type ptToPtMthd struct {
	checkPoints []point.Point
}

var _ Method = (*ptToPtMthd)(nil)

func (m *ptToPtMthd) To(from point.Point, dx DX) (point.Point, error) {
	for k, pt := range m.checkPoints {
		if point.CloseBtw(from, pt) {
			if k+1 < len(m.checkPoints) {
				return m.checkPoints[k+1], nil
			} else {
				return nil, fmt.Errorf("Out of range.")
			}
		}
	}
	return nil, fmt.Errorf("Failed to find next point.")
}

func (m *ptToPtMthd) Apply(s sde.SDE) stchprc.Process {
	return &mthdAppliedPrc{mthd: m, sde: s}
}

func (m *ptToPtMthd) Modify(p point.Point) point.Point {
	return p
}
