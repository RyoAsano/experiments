package stchprc

import (
	"fmt"

	"bitbucket.org/AsanoRyo/experiments/grd"
	"bitbucket.org/AsanoRyo/experiments/path"
	"bitbucket.org/AsanoRyo/experiments/point"
)

func NewDeter(
	grid grd.Grid,
	initPt point.Point,
	dxs ...point.Point,
) (Process, error) {
	if grid.Size() != len(dxs) {
		return nil, fmt.Errorf("The size is inconsistent btw grid and dxs.")
	}

	pathGen := path.NewGenerator(grid, initPt.Dim())
	runningPt := initPt

	var err error
	pathGen.Set(0, runningPt)
	for k, dx := range dxs {
		runningPt, err = point.Add(runningPt, dx)
		if err != nil {
			return nil, err
		}
		pathGen.Set(k+1, runningPt)
	}

	pth, err := pathGen.Generate()
	if err != nil {
		return nil, err
	}
	return &oneStepDetPrc{
		grid: grid, initPt: initPt, pth: pth,
	}, nil
}

type oneStepDetPrc struct {
	grid   grd.Grid
	initPt point.Point
	pth    path.Path
}

var _ Process = (*oneStepDetPrc)(nil)

func (p *oneStepDetPrc) Grid() grd.Grid {
	return p.grid
}

func (p *oneStepDetPrc) Dim() int {
	return p.initPt.Dim()
}

func (p *oneStepDetPrc) Realize() (path.Path, error) {
	return p.pth, nil
}
