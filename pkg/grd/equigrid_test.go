package grd

import (
	"testing"

	"bitbucket.org/AsanoRyo/experiments/pkg/mathutil"
)

func TestEquiDistGrid(t *testing.T) {
	size := 10
	terminal := 2.5
	grid := NewEquiDistGrid(size, terminal)

	for k := 0; k <= size; k++ {
		actual := grid.Get(k)
		expected := float64(k) / float64(size) * terminal

		if !mathutil.CloseBtw(actual, expected) {
			t.Fatalf("Something wrong.")
		}
	}

	if !mathutil.CloseBtw(grid.Terminal(), terminal) {
		t.Fatalf("Terminal doesn't match.")
	}
}
