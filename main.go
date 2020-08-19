package main

import (
	"fmt"
	"github.com/copperium/fractals/fraclib"
	"github.com/copperium/fractals/mandelbrot"
)

func main() {
	fractal := mandelbrot.Mandelbrot{Threshold: 20}
	results := make(chan fraclib.PointResult)
	numPoints := fraclib.Compute(
		&fractal,
		fraclib.Rect{
			BottomLeft: &fraclib.Point{X: -10, Y: -10},
			TopRight:   &fraclib.Point{X: 10, Y: 10},
		}, 0.01, 10, results,
	)
	for i := 0; i < numPoints; i++ {
		pr := <-results
		fmt.Printf("%v: %v\n", pr.Point, pr.Result)
	}
}
