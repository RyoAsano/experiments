package vecfld

import (
	"testing"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/point"
)

func TestZeroVecFld(t *testing.T) {
	z := NewZero(4, 45)

	domainDim, rangeDim := z.Dims()
	if domainDim != 4 {
		t.Fatalf("Domain's dimension is wrong.")
	}
	if rangeDim != 45 {
		t.Fatalf("Range's dimension is wrong.")
	}

	for _, pt := range []point.Point{
		point.New(0, 0, 0, 0),
		point.New(10000, -233330, 343, 34242),
		point.New(-3432, -0.224, 43242, 43),
	} {
		actPt, err := z.At(pt)
		if err != nil {
			t.Fatalf("Failed to get pt: " + err.Error())
		}
		if !point.CloseBtw(actPt, point.Origin(45)) {
			t.Fatalf("Returned point is unexpected.")
		}
	}

	_, err := z.At(point.Origin(domainDim + 4))
	if err == nil {
		t.Fatalf("Domain error didn't occur as expected.")
	}
}
