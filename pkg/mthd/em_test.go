package mthd

import (
	"testing"

	"github.com/RyoAsano/stochastic_calculus/pkg/grd"
	"github.com/RyoAsano/stochastic_calculus/pkg/point"
	"github.com/RyoAsano/stochastic_calculus/pkg/samplesde1"
	"github.com/RyoAsano/stochastic_calculus/pkg/stchprc"
	"github.com/RyoAsano/stochastic_calculus/pkg/vecfld"
)

func TestTo(t *testing.T) {
	emMthd := NewEulerMaruyama()
	dx := DX{
		vecfld.NewId(2): 1.0,
		vecfld.NewId(2): -3.0,
	}

	pt := point.New(2, 3)
	if actual, err := emMthd.To(pt, dx); err != nil {
		t.Fatalf("Failed to get next pt: " + err.Error())
	} else {
		expDx := point.New(-4, -6)
		expected, err := point.Add(pt, expDx)
		if err != nil {
			panic(err.Error())
		}
		if !point.CloseBtw(actual, expected) {
			t.Fatalf("Wrong iwth impl.")
		}
	}
}

func TestApply(t *testing.T) {
	grid := grd.NewEquiDistGrid(2, 1.0)
	// Integrator
	init := point.Origin(2)
	dx1 := point.New(1, -1)
	dx2 := point.New(2, -2)
	intr, err := stchprc.NewDeter(grid, init, dx1, dx2)
	if err != nil {
		panic(err.Error())
	}
	initPt := point.New(1, -2, 0.5)
	s, err := samplesde1.New(grid, initPt, intr, 2, 4)
	if err != nil {
		panic(err.Error())
	}
	// Now we have the SDE:
	// dY(t) = V(Y(t))dX1(t) + W(Y(t))dX2(t), Y(0)=(1,1,1)
	// where
	// V(y) = 2 (y1, y2, y3)
	// W(y) = 4 (y1, y2, y3)
	// dX1(t) = 1, dX2(t)=-1 for t=0
	// dX1(t) = 2, dX2(t)=-2 for t=1
	// Thus:
	// Y(0) = (1, -2, 0.5)
	// Y(1) = Y(0) + 2*Y(0)*1 + 4*Y(0)*(-1) = (-1, 2, -0.5)
	// Y(2) = Y(1) + 2*Y(1)*2 + 4*Y(1)*(-2) = (3, -6, -1.5)

	emMthd := NewEulerMaruyama()
	actPrc := emMthd.Apply(s)
	pth, err := actPrc.Realize()
	if err != nil {
		t.Fatalf("Failed to realize: " + err.Error())
	}

	if actInit, err := pth.At(0); err != nil {
		t.Fatalf("Failed to get init pt: " + err.Error())
	} else {
		if !point.CloseBtw(actInit, initPt) {
			t.Fatalf("Wrong with the initial point.")
		}
	}

	if actSec, err := pth.At(1); err != nil {
		t.Fatalf("Failed to get second pt: " + err.Error())
	} else {
		expSec := point.New(-1, 2, -0.5)
		if !point.CloseBtw(actSec, expSec) {
			t.Fatalf("Wrong with the second point.")
		}
	}

	if actThir, err := pth.At(2); err != nil {
		t.Fatalf("Failed to get third pt: " + err.Error())
	} else {
		expThir := point.New(3, -6, 1.5)
		if !point.CloseBtw(actThir, expThir) {
			t.Fatalf("Wrong with the third point.")
		}
	}
}
