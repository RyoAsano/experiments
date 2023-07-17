package path

import (
	"fmt"

	"github.com/AsanoRyo/stochastic_calculus/pkg/grd"
	"github.com/AsanoRyo/stochastic_calculus/pkg/point"
)

type DimMismatchErr struct {
	gen        *gen
	givenPoint point.Point
}

var _ error = (*DimMismatchErr)(nil)

func (err DimMismatchErr) Error() string {
	return fmt.Sprintf(
		"The point's dimension %d does not match the path's one %d",
		err.givenPoint.Dim(), err.gen.dim,
	)
}

type Generator interface {
	Set(index int, point point.Point) error
	Generate() (Path, error)
}

func NewGenerator(grid grd.Grid, dim int) Generator {
	return &gen{
		grid:   grid,
		dim:    dim,
		points: make([]point.Point, grid.Card()),
	}
}

type gen struct {
	grid   grd.Grid
	dim    int
	points []point.Point
}

var _ Generator = (*gen)(nil)

func (g *gen) Set(index int, pt point.Point) error {
	var err error
	if 0 <= index && index <= g.grid.Size() {
		if pt.Dim() != g.dim {
			return DimMismatchErr{gen: g, givenPoint: pt}
		}
		g.points[index] = pt
		err = nil
	} else {
		err = indexOutOfRangeErr(index, g.grid.Size()+1)
	}
	return err
}

func (g *gen) Generate() (Path, error) {
	for _, point := range g.points {
		if point == nil {
			return nil, notRealizedErr()
		}
	}
	return &path{grid: g.grid, points: g.points, dim: g.dim}, nil
}
