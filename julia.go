package fractals

import (
	"math/cmplx"
)

type Julia struct {
	Threshold float64
	Param     complex128
}

func (m *Julia) At(point *Point, iters int) int {
	z := point.Complex()
	for i := 1; i <= iters; i++ {
		z = z*z + m.Param
		if cmplx.Abs(z) > m.Threshold {
			return i
		}
	}
	return 0
}
