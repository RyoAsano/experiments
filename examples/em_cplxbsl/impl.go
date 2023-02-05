package main

import (
	"math/rand"

	"bitbucket.org/AsanoRyo/experiments/pkg/cplxbsl"
	"bitbucket.org/AsanoRyo/experiments/pkg/grd"
	"bitbucket.org/AsanoRyo/experiments/pkg/mthd"
	"bitbucket.org/AsanoRyo/experiments/pkg/randgen"
	"bitbucket.org/AsanoRyo/experiments/pkg/stchprc"
)

func EulerMaruyamaComplexBessel(r rand.Rand, T float64, N int, x float64, y float64) stchprc.Process {
	grid := grd.NewEquiDistGrid(N, T)
	normGen := randgen.NewNorm(r, 0, 1.0)
	eulerMaruyama := mthd.NewEulerMaruyama()
	// Applies Euler-Maruyama Method to Black Scholes SDE
	complexBesselSDE := cplxbsl.NewSDE(grid, normGen, x, y)
	return eulerMaruyama.Apply(complexBesselSDE)
}
