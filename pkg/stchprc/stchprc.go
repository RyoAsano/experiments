package stchprc

import (
	"github.com/AsanoRyo/stochastic_calculus/pkg/grd"
	"github.com/AsanoRyo/stochastic_calculus/pkg/path"
)

type Process interface {
	Grid() grd.Grid
	Dim() int
	Realize() (path.Path, error)
}
