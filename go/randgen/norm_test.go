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
	std_dev := 10.0
	gen_func := func() float64 {
		return 2.1
	}
	var expected float64 = 23 // = 10.0*2.1 + 2.0

	gen := normRandGenerator{
		mean:     mean,
		stdDev:   std_dev,
		gen_func: gen_func,
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
