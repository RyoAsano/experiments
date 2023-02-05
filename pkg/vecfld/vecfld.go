package vecfld

import (
	"fmt"

	"bitbucket.org/AsanoRyo/experiments/pkg/point"
)

type VectorField interface {
	Dims() (int, int)
	At(p point.Point) (point.Point, error)
}

type OutOfDomainErr struct {
	Vecfld     VectorField
	GivenPoint point.Point
}

func (err OutOfDomainErr) Error() string {
	return fmt.Sprint("The point does not lie in the vector field's domain.")
}
