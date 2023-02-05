package main

import (
	"math/rand"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/bs"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/grd"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/mthd"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/randgen"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/stchprc"
)

func EulerMaruyamaBlackSholes(r rand.Rand, T float64, N int, x0 float64, mu float64, sigma float64) stchprc.Process {
	grid := grd.NewEquiDistGrid(N, T)
	normGen := randgen.NewNorm(r, 0, 1.0)
	blackScholes := bs.NewSDE(grid, normGen, x0, mu, sigma)
	eulerMaruyama := mthd.NewEulerMaruyama()
	// Applies Euler-Maruyama Method to Black Scholes SDE
	return eulerMaruyama.Apply(blackScholes)
}
