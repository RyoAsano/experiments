package mthd

import (
	"testing"

	"bitbucket.org/AsanoRyo/stochastic_calculus/pkg/point"
)

func TestInject(t *testing.T) {
	m := NewPtToPtMthd([]point.Point{}...)
	s := 3.5

	var f = func(p point.Point) point.Point {
		return point.Mul(p, s)
	}

	newMthd := Inject(m, f)

	samplePt := point.New(1, 2)
	actPt := newMthd.Modify(samplePt)
	if !point.CloseBtw(actPt, point.New(s*1, s*2)) {
		t.Fatalf("Something wrong.")
	}
	// Try again.
	samplePt = point.New(2.4424, 1.343)
	actPt = newMthd.Modify(samplePt)
	if !point.CloseBtw(actPt, point.New(s*2.4424, s*1.343)) {
		t.Fatalf("Something wrong.")
	}
}
