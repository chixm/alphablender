package alphablender

type Pic struct {
	r float64
	g float64
	b float64
	a float64
}

func (m *Pic) RGBA() (r, g, b, a uint32) {
	return 0, 0, 0, 0
}
