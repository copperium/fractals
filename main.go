package main

import (
	"fmt"
	"github.com/copperium/fractals/bigcmplx"
	"github.com/copperium/fractals/fractal"
	"image/png"
	"math/big"
	"os"
)

func main() {
	frac := fractal.NewJulia(big.NewRat(1000, 1), bigcmplx.NewComplex(big.NewRat(-8, 10), big.NewRat(156, 1000)))
	iters := 100
	fracviz := fractal.Image{
		Model:   fractal.ThresholdColorModel{Threshold: iters},
		Fractal: frac,
		FractalBounds: fractal.Rect{
			BottomLeft: fractal.NewPoint(big.NewRat(-2, 1), big.NewRat(-2, 1)),
			TopRight:   fractal.NewPoint(big.NewRat(2, 1), big.NewRat(2, 1)),
		},
		Iters:     iters,
		PixelSize: big.NewRat(1, 1),
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
