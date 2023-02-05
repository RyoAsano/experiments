package mathutil

import (
	"math"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/consts"
)

func CloseBtw(a float64, b float64) bool {
	return math.Abs(a-b) < consts.CloseUpTo
}
