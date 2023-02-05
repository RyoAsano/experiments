package mthd

import (
	"testing"

	"bitbucket.org/AsanoRyo/experiments/point"
	"bitbucket.org/AsanoRyo/experiments/vecfld"
)

func TestPtToPtMthdTo(t *testing.T) {
	init := point.New(1, 2)
	sec := point.New(3, 5)
	thir := point.New(-100, 24345)
	ter := point.New(445, 242)

	mthd := NewPtToPtMthd(init, sec, thir, ter)
	emptyDx := map[vecfld.VectorField]float64{}

	actSec, err := mthd.To(init, emptyDx)
	if err != nil || !point.CloseBtw(sec, actSec) {
		t.Fatalf("Failed to get second point.")
	}

	actThir, err := mthd.To(sec, emptyDx)
	if err != nil || !point.CloseBtw(thir, actThir) {
		t.Fatalf("Failed to get second point.")
	}

	actTer, err := mthd.To(thir, emptyDx)
	if err != nil || !point.CloseBtw(ter, actTer) {
		t.Fatalf("Failed to get second point.")
	}
}
