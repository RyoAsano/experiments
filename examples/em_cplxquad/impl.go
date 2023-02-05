package main

import (
	"math"
	"math/rand"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/cplxquad"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/cplxquadexp"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/grd"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/mthd"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/point"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/randgen"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/stchprc"
)

func EulerMaruyamaQuadratic(r rand.Rand, T float64, N int, x float64, y float64) stchprc.Process {
	grid := grd.NewEquiDistGrid(N, T)
	normGen := randgen.NewNorm(r, 0, 1.0)
	eulerMaruyama := mthd.NewEulerMaruyama()
	// Applies Euler-Maruyama Method to Black Scholes SDE
	quadraticSDE := cplxquad.NewSDE(grid, x, y, normGen)
	return eulerMaruyama.Apply(quadraticSDE)
}

func EulerMaruyamaQuadraticImproved(r rand.Rand, T float64, N int, x float64, y float64) stchprc.Process {
	grid := grd.NewEquiDistGrid(N, T)
	normGen := randgen.NewNorm(r, 0, 1.0)
	eulerMaruyama := mthd.NewEulerMaruyama()
	// Applies Euler-Maruyama Method to Black Scholes SDE
	quadraticSDE := cplxquadexp.NewSDE(grid, x, normGen)

	expEulerMaruyama := mthd.Inject(
		eulerMaruyama,
		func(p point.Point) point.Point {
			X, _ := p.Pr(1)
			Y, _ := p.Pr(2)
			return point.New(X, y*math.Exp(Y))
		},
	)
	return expEulerMaruyama.Apply(quadraticSDE)
}
