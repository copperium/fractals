package mandelbrot

import (
	"github.com/copperium/fractals/fraclib"
	"math/cmplx"
)

type Mandelbrot struct {
	Threshold float64
}

func (m *Mandelbrot) At(point *fraclib.Point, iters int) int {
	c := point.Complex()
	var z complex128
	for i := 1; i <= iters; i++ {
		z = z*z + c
		if cmplx.Abs(z) > m.Threshold {
			return i
		}
	}
	return -1
}
