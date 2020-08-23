package fractal

import (
	"github.com/copperium/fractals/bigcmplx"
	"math/big"
)

type Julia struct {
	squareThreshold big.Rat
	param           *bigcmplx.Complex
}

func NewJulia(threshold *big.Rat, param *bigcmplx.Complex) *Julia {
	julia := &Julia{param: param}
	julia.squareThreshold.Mul(threshold, threshold)
	return julia
}

func (j *Julia) At(point *Point, iters int) int {
	// ...exponential growth in Int complexity...
	// what to do?
	z := point.Complex()
	var sqAbs bigcmplx.Complex
	for i := 1; i <= iters; i++ {
		z.Mul(z, z)
		z.Add(z, j.param)
		sqAbs.SqAbs(z)
		if z.Real().Cmp(&j.squareThreshold) > 0 {
			return i
		}
	}
	return 0
}
