package path

import (
	"errors"
	"fmt"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/grd"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/point"
)

type Path interface {
	Grid() grd.Grid
	Dim() int
	At(index int) (point.Point, error)
	Incr(from int, to int) (point.Point, error)
}

func notRealizedErr() error {
	return errors.New("Need to realize first.")
}

func indexOutOfRangeErr(index int, pathLen int) error {
	msg := fmt.Sprintf("%d is out of range: path's length is %d", index, pathLen)
	return errors.New(msg)
}

type path struct {
	grid   grd.Grid
	dim    int
	points []point.Point
}

var _ Path = (*path)(nil)

func (p *path) Grid() grd.Grid {
	return p.grid
}

func (p *path) Dim() int {
	return p.dim
}

func (p *path) At(index int) (point.Point, error) {
	// If the path is not fully realized we raise an error.
	if 0 <= index && index <= p.grid.Size() {
		return p.points[index], nil
	} else {
		return nil, indexOutOfRangeErr(index, p.grid.Size())
	}
}

func (p *path) Incr(from int, to int) (point.Point, error) {
	ptFrom, err := p.At(from)
	if err != nil {
		return nil, err
	}
	ptTo, err := p.At(to)
	if err != nil {
		return nil, err
	}
	diff, err := point.Diff(ptTo, ptFrom)
	if err != nil {
		return nil, err
	}
	return diff, nil
}
