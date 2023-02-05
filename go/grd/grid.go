package grd

type Grid interface {
	Size() int
	Card() int
	Terminal() float64
	Get(k int) float64
}

func NewEquiDistGrid(size int, terminal float64) Grid {
	g := equiDistGrid{}
	g.init(size, terminal)
	return &g
}
