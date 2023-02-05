package testutil

import (
	"math/rand"

	"bitbucket.org/AsanoRyo/experiments/pkg/randgen"
)

type DummyGen struct {
	Value float64
}

var _ randgen.RandGenerator = (*DummyGen)(nil)

func (gen DummyGen) Get() float64 {
	return gen.Value
}

func NewRand(seed int64) rand.Rand {
	return *rand.New(rand.NewSource(seed))
}

func NewStdNorm(seed int64) randgen.RandGenerator {
	return randgen.NewNorm(NewRand(seed), 0, 1)
}
