package main

import (
	"fmt"
	"github.com/copperium/fractals/fractal"
	"image/png"
	"os"
	"runtime"
)

func main() {
	frac := fractal.Julia{Threshold: 1000, Param: -0.8 + 0.156i}
	//frac := fractal.Mandelbrot{Threshold: 1000}
	iters := 100
	viz := fractal.Image{
		Model:   fractal.HueColorModel{Threshold: iters, HueRange: fractal.RedToGreen},
		Fractal: &frac,
		FractalBounds: fractal.Rect{
			BottomLeft: &fractal.Point{X: -2, Y: -2},
			TopRight:   &fractal.Point{X: 2, Y: 2},
		},
		Iters:     iters,
		PixelSize: 0.001,
	}
	workers := runtime.NumCPU()
	cached := viz.ToCachedImage(workers)

	file, err := os.Create("julia.png")
	if err != nil {
		_ = fmt.Errorf(err.Error())
		return
	}
	defer file.Close()

	err = png.Encode(file, cached)
	if err != nil {
		_ = fmt.Errorf(err.Error())
	}
}
