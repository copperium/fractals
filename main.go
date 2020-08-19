package main

import (
	"fmt"
	"github.com/copperium/fractals/fraclib"
	"github.com/copperium/fractals/mandelbrot"
	"github.com/copperium/fractals/viz"
	"image/png"
	"os"
)

func main() {
	fractal := mandelbrot.Mandelbrot{Threshold: 1000}
	iters := 100
	fracviz := viz.FractalImage{
		Model:   viz.ThresholdModel{Threshold: iters},
		Fractal: &fractal,
		FractalBounds: fraclib.Rect{
			BottomLeft: &fraclib.Point{X: -2, Y: -1},
			TopRight:   &fraclib.Point{X: 1, Y: 1},
		},
		Iters:     iters,
		PixelSize: 0.001,
	}

	file, err := os.Create("mandelbrot.png")
	if err != nil {
		_ = fmt.Errorf(err.Error())
		return
	}
	defer file.Close()

	err = png.Encode(file, &fracviz)
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
}
