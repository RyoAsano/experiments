package main

import (
	"math/rand"

	"github.com/AsanoRyo/stochastic_calculus/pkg/cplxbsl"
	"github.com/AsanoRyo/stochastic_calculus/pkg/grd"
	"github.com/AsanoRyo/stochastic_calculus/pkg/mthd"
	"github.com/AsanoRyo/stochastic_calculus/pkg/randgen"
	"github.com/AsanoRyo/stochastic_calculus/pkg/stchprc"
)

func EulerMaruyamaComplexBessel(r rand.Rand, T float64, N int, x float64, y float64) stchprc.Process {
	grid := grd.NewEquiDistGrid(N, T)
	normGen := randgen.NewNorm(r, 0, 1.0)
	eulerMaruyama := mthd.NewEulerMaruyama()
	// Applies Euler-Maruyama Method to Black Scholes SDE
	complexBesselSDE := cplxbsl.NewSDE(grid, normGen, x, y)
	return eulerMaruyama.Apply(complexBesselSDE)
}
