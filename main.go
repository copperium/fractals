package main

import (
	"fmt"
	"github.com/copperium/fractals/fraclib"
	"image/png"
	"os"
)

func main() {
	fractal := fraclib.Julia{Threshold: 1000, Param: -0.8 + 0.156i}
	iters := 100
	fracviz := fraclib.FractalImage{
		Model:   fraclib.ThresholdModel{Threshold: iters},
		Fractal: &fractal,
		FractalBounds: fraclib.Rect{
			BottomLeft: &fraclib.Point{X: -2, Y: -2},
			TopRight:   &fraclib.Point{X: 2, Y: 2},
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
