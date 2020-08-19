package fraclib

type Point struct {
	X, Y float64
}

type Rect struct {
	TopLeft     *Point
	BottomRight *Point
}

type PointResult struct {
	Point  *Point
	Result int
}

type Fractal interface {
	At(*Point) int
}

func computePoint(fractal Fractal, point *Point, results chan PointResult) {
	results <- PointResult{point, fractal.At(point)}
}

func Compute(fractal Fractal, bounds Rect, precision float64, results chan PointResult) {
	for x := bounds.TopLeft.X; x <= bounds.BottomRight.X; x += precision {
		for y := bounds.TopLeft.Y; y <= bounds.BottomRight.Y; y += precision {
			go computePoint(fractal, &Point{x, y}, results)
		}
	}
}
