package point

import (
	"fmt"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/consts"
)

type Point interface {
	Dim() int
	Pr(dim int) (float64, error)
	Norm() float64
}

func New(coords ...float64) Point {
	return &point{dim: len(coords), coord: coords}
}

func Origin(dim int) Point {
	coord := make([]float64, dim)
	return New(coord...)
}

func Add(p Point, q Point) (Point, error) {
	if p.Dim() != q.Dim() {
		return nil, DimMismatchErr{p: p, q: q}
	}

	coord := make([]float64, p.Dim())
	for d := 1; d <= p.Dim(); d++ {
		p_coord, err := p.Pr(d)
		if err != nil {
			return nil, err
		}
		q_coord, err := q.Pr(d)
		if err != nil {
			return nil, err
		}
		coord[d-1] = p_coord + q_coord
	}
	return New(coord...), nil
}

func Mul(p Point, scale float64) Point {
	coord := make([]float64, p.Dim())
	for d := 1; d <= p.Dim(); d++ {
		val, _ := p.Pr(d)
		coord[d-1] = scale * val
	}
	return New(coord...)
}

func Diff(to Point, from Point) (Point, error) {
	return Add(to, Mul(from, -1.0))
}

func Dist(p Point, q Point) (float64, error) {
	if diff, err := Diff(p, q); err != nil {
		return 0, err
	} else {
		return diff.Norm(), nil
	}
}

func CloseBtw(p Point, q Point) bool {
	dist, err := Dist(p, q)
	if err != nil {
		panic("Something wrong: " + err.Error())
	}
	return dist < consts.CloseUpTo
}

type DimMismatchErr struct {
	p Point
	q Point
}

var _ error = (*DimMismatchErr)(nil)

func (err DimMismatchErr) Error() string {
	return fmt.Sprintf(
		"Dimension mismatched: %d vs %d",
		err.p.Dim(), err.q.Dim(),
	)
}

type DimOutOfRangeErr struct {
	givenDim    int
	accessedDim int
}

var _ error = (*DimOutOfRangeErr)(nil)

func (err DimOutOfRangeErr) Error() string {
	return fmt.Sprintf(
		"The dimension %d is out of the range [1, %d]",
		err.givenDim, err.accessedDim,
	)
}
