package vecfld

import (
	"testing"

	"bitbucket.org/AsanoRyo/experiments/pkg/point"
)

func TestId(t *testing.T) {
	dim := 3
	id := NewId(dim)

	if domDim, ranDim := id.Dims(); domDim != dim || ranDim != dim {
		t.Fatalf("Dimension is wrong.")
	}

	for _, pt := range []point.Point{
		point.New(3.3, 4, 3),
		point.New(-23, 0.5, 3),
	} {
		if actPt, err := id.At(pt); err != nil || !point.CloseBtw(actPt, pt) {
			t.Fatalf("Impl is wrong.")
		}
	}

	_, err := id.At(point.Origin(dim + 1))
	if err == nil {
		t.Fatalf("Error didn't occur as expected.")
	}
}
