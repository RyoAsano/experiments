package vecfld

import (
	"testing"

	"bitbucket.org/AsanoRyo/experiments/point"
)

func TestAffine(t *testing.T) {
	scale := -2.4
	adder := point.New(3, 4, -1)
	affine := NewAffine(scale, adder)

	if domDim, randim := affine.Dims(); domDim != 3 || randim != 3 {
		t.Fatalf("Dimension is wrong.")
	}

	for _, pt := range []point.Point{
		point.New(0, 0, 0),
		point.New(2, -4.5, -1),
		point.New(-2.4, 4.5, 100),
	} {
		expected, err := point.Add(point.Mul(pt, scale), adder)
		if err != nil {
			panic("Test config is wrong: " + err.Error())
		}

		actual, err := affine.At(pt)
		if err != nil {
			t.Fatalf("Failed to get a point.")
		}
		if !point.CloseBtw(actual, expected) {
			t.Fatalf("Unexpected value was set.")
		}
	}
}
