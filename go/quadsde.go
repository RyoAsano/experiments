package main

import (
	"math"
	"math/rand"

	"bitbucket.org/AsanoRyo/experiments/cplxquad"
	"bitbucket.org/AsanoRyo/experiments/cplxquadexp"
	"bitbucket.org/AsanoRyo/experiments/grd"
	"bitbucket.org/AsanoRyo/experiments/mthd"
	"bitbucket.org/AsanoRyo/experiments/point"
	"bitbucket.org/AsanoRyo/experiments/randgen"
	"bitbucket.org/AsanoRyo/experiments/stchprc"
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
