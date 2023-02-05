package timeprc

import (
	"bitbucket.org/AsanoRyo/experiments/pkg/grd"
	"bitbucket.org/AsanoRyo/experiments/pkg/path"
	"bitbucket.org/AsanoRyo/experiments/pkg/point"
	"bitbucket.org/AsanoRyo/experiments/pkg/stchprc"
)

func New(grid grd.Grid) stchprc.Process {
	return &timeProcess{grid: grid}
}

var _ stchprc.Process = (*timeProcess)(nil)

type timeProcess struct {
	grid grd.Grid
}

func (t *timeProcess) Dim() int {
	return 1
}

func (t *timeProcess) Grid() grd.Grid {
	return t.grid
}

func (t *timeProcess) Realize() (path.Path, error) {
	gen := path.NewGenerator(t.grid, t.Dim())
	for k := 0; k <= t.grid.Size(); k++ {
		gen.Set(k, point.New(t.grid.Get(k)))
	}
	return gen.Generate()
}
