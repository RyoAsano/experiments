package stchprc

import (
	"bitbucket.org/AsanoRyo/experiments/pkg/grd"
	"bitbucket.org/AsanoRyo/experiments/pkg/path"
)

type Process interface {
	Grid() grd.Grid
	Dim() int
	Realize() (path.Path, error)
}
