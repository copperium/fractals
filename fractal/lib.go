package fractal

import (
	"fmt"
	"github.com/copperium/fractals/bigcmplx"
	"math/big"
)

type Point struct {
	X, Y big.Rat
}

func NewPoint(x, y *big.Rat) *Point {
	return &Point{X: *x, Y: *y}
}

func (p *Point) Complex() *bigcmplx.Complex {
	return bigcmplx.NewComplex(&p.X, &p.Y)
}

func (p *Point) String() string {
	return fmt.Sprintf("(%s, %s)", p.X.String(), p.Y.String())
}

type Rect struct {
	BottomLeft *Point
	TopRight   *Point
}

type PointResult struct {
	Point  *Point
	Result int
}

type Fractal interface {
	At(point *Point, iters int) int
}

func computePoint(fractal Fractal, point *Point, iters int, results chan PointResult) {
	results <- PointResult{point, fractal.At(point, iters)}
}

func Compute(fractal Fractal, bounds Rect, precision *big.Rat, iters int, results chan PointResult) (numPoints int) {
	for x := bounds.BottomLeft.X; x.Cmp(&bounds.TopRight.X) < 0; x.Add(&x, precision) {
		for y := bounds.BottomLeft.Y; y.Cmp(&bounds.TopRight.Y) < 0; y.Add(&y, precision) {
			go computePoint(fractal, &Point{x, y}, iters, results)
			numPoints++
		}
	}
	return numPoints
}
