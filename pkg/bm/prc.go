package bm

import (
	"math"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/grd"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/path"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/point"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/randgen"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/stchprc"
)

func New(
	grid grd.Grid,
	bmDim int,
	randgen randgen.RandGenerator,
	withTimeprc bool,
) stchprc.Process {
	totalDim := bmDim
	if withTimeprc {
		totalDim += 1
	}

	return &brownianMotion{
		grid:        grid,
		randNumGen:  randgen,
		bmDim:       bmDim,
		totalDim:    totalDim,
		withTimeprc: withTimeprc,
	}
}

type brownianMotion struct {
	randNumGen  randgen.RandGenerator
	grid        grd.Grid
	bmDim       int
	totalDim    int
	withTimeprc bool
}

var _ stchprc.Process = (*brownianMotion)(nil)

func (bm *brownianMotion) Dim() int {
	return bm.totalDim
}

func (bm *brownianMotion) Grid() grd.Grid {
	return bm.grid
}

func (bm *brownianMotion) Realize() (path.Path, error) {
	gen := path.NewGenerator(bm.grid, bm.totalDim)

	// Add the initial point.
	runningPt := point.Origin(bm.totalDim)
	gen.Set(0, runningPt)

	for k := 1; k <= bm.grid.Size(); k++ {
		// Get next coordinate
		nextCoord := make([]float64, bm.bmDim)
		for dim := 1; dim <= bm.bmDim; dim++ {
			runningVal, err := runningPt.Pr(dim)
			if err != nil {
				return nil, err
			}
			normRandVal := bm.randNumGen.Get()
			delta := bm.grid.Get(k) - bm.grid.Get(k-1)
			nextCoord[dim-1] = runningVal + normRandVal*math.Sqrt(delta)
		}
		// This line must be put here, NOT further;
		// o/w wrong nextCoord will be used.
		runningPt = point.New(nextCoord...)

		// Prepend time process if necessary
		if bm.withTimeprc {
			time := bm.grid.Get(k)
			nextCoord = append([]float64{time}, nextCoord...)
		}
		// Set a point at k
		if err := gen.Set(k, point.New(nextCoord...)); err != nil {
			return nil, err
		}
	}
	return gen.Generate()
}
