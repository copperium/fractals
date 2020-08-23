package fractal

import (
	"github.com/copperium/fractals/bigcmplx"
	"math/big"
)

type Mandelbrot struct {
	squareThreshold big.Rat
}

func NewMandelbrot(threshold *big.Rat) *Mandelbrot {
	var m Mandelbrot
	m.squareThreshold.Mul(threshold, threshold)
	return &m
}

func (m *Mandelbrot) At(point *Point, iters int) int {
	c := point.Complex()
	var z, sqAbs bigcmplx.Complex
	for i := 1; i <= iters; i++ {
		z.Mul(&z, &z)
		z.Add(&z, c)
		sqAbs.SqAbs(&z)
		if sqAbs.Real().Cmp(&m.squareThreshold) > 0 {
			return i
		}
	}
	return 0
}
