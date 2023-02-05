package sde

import (
	"fmt"

	"bitbucket.org/AsanoRyo/experiments/grd"
	"bitbucket.org/AsanoRyo/experiments/point"
	"bitbucket.org/AsanoRyo/experiments/stchprc"
	"bitbucket.org/AsanoRyo/experiments/vecfld"
)

type SDE interface {
	Grid() grd.Grid
	Dim() int
	InitPoint() point.Point
	Integrand(dim int) (vecfld.VectorField, error)
	Integrator() stchprc.Process
}

type DimOutOfRangeErr struct {
	SDE      SDE
	GivenDim int
}

var _ error = (*DimOutOfRangeErr)(nil)

func (err DimOutOfRangeErr) Error() string {
	return fmt.Sprintf(
		"No vector field for the dimension: %d; it must be in [1, %d]",
		err.GivenDim, err.SDE.Integrator().Dim(),
	)
}
