package randgen

import (
	"math"
	"math/rand"
	"testing"
)

func TestInterface(t *testing.T) {
}

func TestInitNormRandGen(t *testing.T) {
	r := *rand.New(rand.NewSource(99))
	gen := NewNorm(r, 1.0, 1.0)
	gen.Get()
}

func TestNormRandGen(t *testing.T) {

	mean := 2.0
	stdDev := 10.0
	genFunc := func() float64 {
		return 2.1
	}
	var expected float64 = 23 // = 10.0*2.1 + 2.0

	gen := normRandGenerator{
		mean:    mean,
		stdDev:  stdDev,
		genFunc: genFunc,
	}

	var actual float64

	actual = gen.Get()
	if math.Abs(expected-actual) > 1e-7 {
		t.Fatalf("Something wrong.")
	}
	// Again
	actual = gen.Get()
	if math.Abs(expected-actual) > 1e-7 {
		t.Fatalf("Something wrong.")
	}

}
