package point

import (
	"math"
)

type point struct {
	dim   int
	coord []float64
}

var _ Point = (*point)(nil)

func (p point) Dim() int {
	return p.dim
}

func (p point) Pr(dim int) (float64, error) {
	if 1 <= dim && dim <= p.dim {
		return p.coord[dim-1], nil
	} else {
		return 0, DimOutOfRangeErr{givenDim: dim, accessedDim: p.dim}
	}
}

func (p point) Norm() float64 {
	squaredSum := .0
	for i := 0; i < p.dim; i++ {
		squaredSum += p.coord[i] * p.coord[i]
	}
	return math.Sqrt(squaredSum)
}
