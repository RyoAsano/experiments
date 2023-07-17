package stchprc

import (
	"testing"

	"github.com/RyoAsano/stochastic_calculus/pkg/grd"
	"github.com/RyoAsano/stochastic_calculus/pkg/point"
)

func TestDeter(t *testing.T) {
	expInit := point.New(1, -1)

	dx1 := point.New(2, -2)
	expSec := point.New(1+2, -1-2)

	dx2 := point.New(3, -3)
	expThir := point.New(1+2+3, -1-2-3)

	grid := grd.NewEquiDistGrid(2, 2.3)

	prc, err := NewDeter(grid, expInit, dx1, dx2)
	if err != nil {
		t.Fatalf("Wrong in init: " + err.Error())
	}
	pth, err := prc.Realize()
	if err != nil {
		t.Fatalf("Wrong in realize: " + err.Error())
	}

	if actInitPt, err := pth.At(0); err != nil || !point.CloseBtw(actInitPt, expInit) {
		t.Fatalf("Falied to get init pt: " + err.Error())
	}

	if actSecPt, err := pth.At(1); err != nil || !point.CloseBtw(actSecPt, expSec) {
		t.Fatalf("Falied to get second pt: " + err.Error())
	}

	if actThirPt, err := pth.At(2); err != nil || !point.CloseBtw(actThirPt, expThir) {
		t.Fatalf("Falied to get third pt: " + err.Error())
	}
}
