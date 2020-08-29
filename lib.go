package fractals

import (
	"fmt"
	"sync"
)

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

func computePoint(fractal Fractal, iters int, points <-chan *Point, results chan<- *PointResult, wg *sync.WaitGroup) {
	for point := range points {
		result := &PointResult{point, fractal.At(point, iters)}
		results <- result
	}
	wg.Done()
}

func Compute(fractal Fractal, bounds Rect, precision float64, iters int, workers int, results chan<- *PointResult) {
	var wg sync.WaitGroup
	points := make(chan *Point, 1000)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go computePoint(fractal, iters, points, results, &wg)
	}
	for x := bounds.BottomLeft.X; x < bounds.TopRight.X; x += precision {
		for y := bounds.BottomLeft.Y; y < bounds.TopRight.Y; y += precision {
			points <- &Point{x, y}
		}
	}
	close(points)
	wg.Wait()
	close(results)
}
