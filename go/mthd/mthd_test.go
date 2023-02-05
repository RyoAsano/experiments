package mthd

import (
	"testing"

	"bitbucket.org/AsanoRyo/experiments/grd"
	"bitbucket.org/AsanoRyo/experiments/point"
	"bitbucket.org/AsanoRyo/experiments/sde"
	"bitbucket.org/AsanoRyo/experiments/stchprc"
)

func TestRealizeInTo(t *testing.T) {
	initPt := point.New(1.2, 3.5)
	secPt := point.New(2.2, 5.34)
	thirPt := point.New(-4.5, 4.5)
	m := NewPtToPtMthd(initPt, secPt, thirPt)

	grid := grd.NewEquiDistGrid(2, 1.0)
	intr, err := stchprc.NewDeter(grid, point.Origin(1), point.Origin(1), point.Origin(1))
	if err != nil {
		panic(err.Error())
	}
	s := sde.NewTrivial(grid, initPt, intr)

	actPrc := m.Apply(s)
	pth, err := actPrc.Realize()
	if err != nil {
		t.Fatalf("Falied to realize: " + err.Error())
	}
	if actInitPt, err := pth.At(0); err != nil || !point.CloseBtw(actInitPt, initPt) {
		t.Fatalf("Point is wrong in realized path")
	}
	if actSecPt, err := pth.At(1); err != nil || !point.CloseBtw(actSecPt, secPt) {
		t.Fatalf("Point is wrong in realized path")
	}
	if actThirPt, err := pth.At(2); err != nil || !point.CloseBtw(actThirPt, thirPt) {
		t.Fatalf("Point is wrong in realized path")
	}
}

func TestRealizeInDX(t *testing.T) {
}
