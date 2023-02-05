package grd

type equiDistGrid []float64

func (g *equiDistGrid) init(size int, terminal float64) {
	*g = make(equiDistGrid, size+1)
	for k := 1; k <= size; k++ {
		(*g)[k] = (*g)[k-1] + terminal/float64(size)
	}
}

func (g *equiDistGrid) Size() int {
	return len(*g) - 1
}

func (g *equiDistGrid) Card() int {
	return len(*g)
}

func (g *equiDistGrid) Get(k int) float64 {
	return (*g)[k]
}

func (g *equiDistGrid) Terminal() float64 {
	return g.Get(g.Size())
}
