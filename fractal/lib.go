package fractal

import "fmt"

type Point struct {
	X, Y float64
}

func (p *Point) Complex() complex128 {
	return complex(p.X, p.Y)
}

func (p *Point) String() string {
	return fmt.Sprintf("(%g, %g)", p.X, p.Y)
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

func Compute(fractal Fractal, bounds Rect, precision float64, iters int, results chan PointResult) (numPoints int) {
	for x := bounds.BottomLeft.X; x < bounds.TopRight.X; x += precision {
		for y := bounds.BottomLeft.Y; y < bounds.TopRight.Y; y += precision {
			go computePoint(fractal, &Point{x, y}, iters, results)
			numPoints++
		}
	}
	return numPoints
}
