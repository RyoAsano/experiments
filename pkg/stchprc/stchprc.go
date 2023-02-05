package stchprc

import (
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/grd"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/path"
)

type Process interface {
	Grid() grd.Grid
	Dim() int
	Realize() (path.Path, error)
}
