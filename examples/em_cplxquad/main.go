package main

import (
	"math/rand"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/stchprc"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/util"
)

func main() {
	var r rand.Rand = *rand.New(rand.NewSource(99))
	var T float64 = 1
	var N int = 10000000

	var process stchprc.Process
	var dir string

	// process, dir = EulerMaruyamaQuadratic(r, T, N, -50, -50), "quad"
	process, dir = EulerMaruyamaQuadraticImproved(r, T, N, 1, 1), "quad_impr"

	file := "path_n0p001_0p1.csv"

	// Generate a sample path
	if a_path, err := process.Realize(); err != nil {
		panic(err.Error())
	} else if err := util.OutputToCsv(a_path, dir, file); err != nil {
		panic(err.Error())
	}
}
