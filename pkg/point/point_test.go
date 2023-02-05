package point

import (
	"math"
	"testing"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/consts"
	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/mathutil"
)

func TestNew(t *testing.T) {
	pt := New(1.2, -2, 3)

	if pt.Dim() != 3 {
		t.Fatalf("Dimension is wrong.")
	}
	if !mathutil.CloseBtw(pt.Norm(), math.Sqrt(1.44+4+9)) {
		t.Fatalf("Norm is wrong.")
	}

	if _, err := pt.Pr(-1); err == nil {
		t.Fatalf("negative dimension shouldn't exist.")
	}

	if _, err := pt.Pr(0); err == nil {
		t.Fatalf("negative dimension shouldn't exist.")
	}

	if act, err := pt.Pr(1); err != nil {
		t.Fatalf("Falied to get a coord.")
	} else {
		if !mathutil.CloseBtw(act, 1.2) {
			t.Fatalf("Coord setting is wrong.")
		}
	}

	if act, err := pt.Pr(2); err != nil {
		t.Fatalf("Falied to get the second coord.")
	} else {
		if !mathutil.CloseBtw(act, -2) {
			t.Fatalf("Coord setting is wrong.")
		}
	}

	if act, err := pt.Pr(3); err != nil {
		t.Fatalf("Falied to get the third coord.")
	} else {
		if !mathutil.CloseBtw(act, 3) {
			t.Fatalf("Coord setting is wrong.")
		}
	}

	if _, err := pt.Pr(4); err == nil {
		t.Fatalf("4th dimension shouldn't exist.")
	}
}

func TestOrigin(t *testing.T) {
	pt := Origin(100)
	if pt.Dim() != 100 {
		t.Fatalf("Dimension is wrong.")
	}

	for dim := 1; dim <= 100; dim++ {
		actVal, err := pt.Pr(dim)
		if err != nil {
			t.Fatalf("Failed to get a coord.")
		}
		if !mathutil.CloseBtw(actVal, 0) {
			t.Fatalf("The coord must be zero.")
		}
	}
}

func TestAddAndDiff(t *testing.T) {
	coords := map[Point][]float64{}

	coord1 := []float64{1, 2, 3}
	pt1 := New(coord1...)
	coords[pt1] = coord1

	coord2 := []float64{-2, -4, -6}
	pt2 := New(coord2...)
	coords[pt2] = coord2

	sumedpt, err := Add(pt1, pt2)
	if err != nil {
		t.Fatalf("Failed to add: " + err.Error())
	}
	coords[sumedpt] = []float64{-1, -2, -3}

	diffpt, err := Diff(pt1, pt2)
	if err != nil {
		t.Fatalf("Failed to take diff: " + err.Error())
	}
	coords[diffpt] = []float64{3, 6, 9}

	if sumedpt.Dim() != 3 || diffpt.Dim() != 3 {
		t.Fatalf("Dimension is wrong.")
	}

	for pt, expCoord := range coords {
		for dim := 1; dim <= pt.Dim(); dim++ {
			actVal, err := pt.Pr(dim)
			if err != nil {
				t.Fatalf("Failed to get a value.")
			}
			expVal := expCoord[dim-1]
			if !mathutil.CloseBtw(actVal, expVal) {
				t.Fatalf("Coordinate setting is wrong.")
			}
		}
	}
}

func TestMul(t *testing.T) {
	coord := []float64{1, 0, -1}
	pt := New(coord...)
	scale := 2.4

	actPt := Mul(pt, scale)

	for dim := 1; dim <= 3; dim++ {
		actVal, err := actPt.Pr(dim)
		if err != nil {
			t.Fatalf("Failed to get a value.")
		}
		expVal := scale * coord[dim-1]
		if !mathutil.CloseBtw(actVal, expVal) {
			t.Fatalf("Value setting is wrong.")
		}
	}
}

func TestDist(t *testing.T) {
	pt := New(1, 2, 3)
	pt2 := New(0, 1, 2)
	dist, err := Dist(pt, pt2)
	if err != nil {
		t.Fatalf("Something is wrong.")
	}
	if !mathutil.CloseBtw(dist, math.Sqrt(3)) {
		t.Fatalf("Distance impl is wrong.")
	}
}

func TestCloseBtw(t *testing.T) {
	pt := New(1, 2)
	if !CloseBtw(pt, pt) {
		t.Fatalf("Any points should be close to themselves.")
	}
	pt2 := New(1+1.1*consts.CloseUpTo, 2+1.1*consts.CloseUpTo)
	if CloseBtw(pt, pt2) {
		t.Fatalf("The two points are not close to each other.")
	}
	pt3 := New(1+0.5*consts.CloseUpTo, 2+0.5*consts.CloseUpTo)
	if !CloseBtw(pt, pt3) {
		t.Fatalf("The two points are close to each other.")
	}
}
