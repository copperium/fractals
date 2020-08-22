package main

import (
	"fmt"
	"github.com/copperium/fractals/fractal"
	"image/png"
	"os"
)

func main() {
	frac := fractal.Julia{Threshold: 1000, Param: -0.8 + 0.156i}
	iters := 100
	fracviz := fractal.Image{
		Model:   fractal.ThresholdModel{Threshold: iters},
		Fractal: &frac,
		FractalBounds: fractal.Rect{
			BottomLeft: &fractal.Point{X: -2, Y: -2},
			TopRight:   &fractal.Point{X: 2, Y: 2},
		},
		Iters:     iters,
		PixelSize: 0.001,
	}

	file, err := os.Create("julia.png")
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
