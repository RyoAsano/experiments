package randgen

import "math/rand"

type normRandGenerator struct {
	mean     float64
	stdDev   float64
	gen_func func() float64
}

func (gen *normRandGenerator) Get() float64 {
	return gen.gen_func()*gen.stdDev + gen.mean
}

var _ RandGenerator = (*normRandGenerator)(nil)

func NewNorm(r rand.Rand, mean float64, stdDev float64) RandGenerator {
	return &normRandGenerator{
		mean:     mean,
		stdDev:   stdDev,
		gen_func: r.NormFloat64,
	}
}
