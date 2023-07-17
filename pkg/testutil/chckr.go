package testutil

import (
	"testing"

	"github.com/RyoAsano/stochastic_calculus/pkg/bm"
	"github.com/RyoAsano/stochastic_calculus/pkg/mathutil"
	"github.com/RyoAsano/stochastic_calculus/pkg/point"
	"github.com/RyoAsano/stochastic_calculus/pkg/sde"
	"github.com/RyoAsano/stochastic_calculus/pkg/stchprc"
	"github.com/RyoAsano/stochastic_calculus/pkg/vecfld"
)

func CheckTimeAndBMIntr(intr stchprc.Process, seed int64, t *testing.T) {
	bmPrc := bm.New(intr.Grid(), 1, NewStdNorm(seed), true)
	bmPath, _ := bmPrc.Realize()

	intrPath, err := intr.Realize()
	if err != nil {
		t.Fatalf("something wrong with realization impl.")
	}
	if intrPath.Dim() != 2 {
		t.Fatalf("Dimension is wrong.")
	}

	for k := 0; k < intr.Grid().Size(); k++ {
		ptExp, _ := bmPath.At(k)
		timeExp, _ := ptExp.Pr(1)
		bmExp, _ := ptExp.Pr(2)

		if actPt, err := intrPath.At(k); err != nil {
			t.Fatalf("Path is not filled entirely.")
		} else {
			// time process
			if timeAct, err := actPt.Pr(1); err != nil {
				t.Fatalf("Point's dimension is wrong.")
			} else if !mathutil.CloseBtw(timeExp, timeAct) {
				t.Fatalf("Something wrong with time process realization impl.")
			}
			// Brownian motion
			if bmAct, err := actPt.Pr(2); err != nil {
				t.Fatalf("Point's dimension is wrong.")
			} else if !mathutil.CloseBtw(bmExp, bmAct) {
				t.Fatalf("Something wrong with BM's realization impl.")
			}
		}
	}
}

type VecFldChecker struct {
	ExpectedFunc func([]float64) []float64
	Samples      []point.Point
}

func CheckVecFld(
	vecFld vecfld.VectorField,
	checker VecFldChecker,
	t *testing.T,
) {
	for _, pt := range checker.Samples {
		coord := make([]float64, pt.Dim())
		for k := range coord {
			coord[k], _ = pt.Pr(k + 1)
		}
		expPt := point.New(checker.ExpectedFunc(coord)...)

		actPt, err := vecFld.At(pt)
		if err != nil {
			t.Fatalf("Wrong with the vector field's definition.")
		}
		if !point.CloseBtw(expPt, actPt) {
			t.Fatalf("Values mismatched.")
		}

	}
}

func Check2DimSDE(
	targetSDE sde.SDE,
	expInitPt point.Point,
	driftChecker VecFldChecker,
	diffusionChecker VecFldChecker,
	seed int64,
	t *testing.T,
) {
	if !point.CloseBtw(targetSDE.InitPoint(), expInitPt) {
		t.Fatalf("Wrong with the initial point setting.")
	}

	driftVec, err := targetSDE.Integrand(1)
	if err != nil {
		t.Fatalf("Failed to get the drift: " + err.Error())
	}
	CheckVecFld(driftVec, driftChecker, t)

	diffVec, err := targetSDE.Integrand(2)
	if err != nil {
		t.Fatalf("Failed to get the drift: " + err.Error())
	}
	CheckVecFld(diffVec, diffusionChecker, t)
	CheckTimeAndBMIntr(targetSDE.Integrator(), seed, t)

}
