package stchprc

import (
	"bitbucket.org/AsanoRyo/experiments/grd"
	"bitbucket.org/AsanoRyo/experiments/path"
)

type Process interface {
	Grid() grd.Grid
	Dim() int
	Realize() (path.Path, error)
}
