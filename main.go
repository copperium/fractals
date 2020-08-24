package main

import (
	"flag"
	"fmt"
	"github.com/copperium/fractals/fractal"
	"image/png"
	"os"
	"runtime"
)

func main() {
	fracType := flag.String("t", "mandelbrot", "Type of fractal: mandelbrot or julia")
	boundsStr := flag.String("b", "-1 -1 1 1", "Bounds: minx miny maxx maxy")
	width := flag.Int("w", 1000, "Width of resulting image")
	threshold := flag.Float64("threshold", 1000, "Fractal computation threshold")
	param := flag.String("param", "0+0i", "Complex parameter for Julia fractal")
	iters := flag.Int("iters", 100, "Fractal iterations")
	colors := flag.String("colors", "blue-to-yellow", "Color scheme: greyscale, blue-to-yellow, or red-to-green")
	boldMode := flag.Bool("bold-mode", false, "For colored schemes: enable bold mode")
	flag.Parse()

	var frac fractal.Fractal
	switch *fracType {
	case "mandelbrot":
		frac = &fractal.Mandelbrot{Threshold: *threshold}
	case "julia":
		var realParam, imagParam float64
		_, err := fmt.Sscanf(*param, "%f%fi", &realParam, &imagParam)
		if err != nil {
			flag.Usage()
			return
		}
		frac = &fractal.Julia{Threshold: *threshold, Param: complex(realParam, imagParam)}
	default:
		fmt.Println("Unknown fractal type:", *fracType)
		flag.Usage()
		return
	}

	var minx, miny, maxx, maxy float64
	_, err := fmt.Sscanf(*boundsStr, "%f %f %f %f", &minx, &miny, &maxx, &maxy)
	if err != nil {
		fmt.Println("Invalid bounds:", *boundsStr)
		flag.Usage()
		return
	}
	bounds := fractal.Rect{
		BottomLeft: &fractal.Point{X: minx, Y: miny},
		TopRight:   &fractal.Point{X: maxx, Y: maxy},
	}
	pixelSize := (bounds.TopRight.X - bounds.BottomLeft.X) / float64(*width)

	var colorModel fractal.ColorModel
	switch *colors {
	case "greyscale":
		colorModel = &fractal.GreyscaleColorModel{Threshold: *iters}
	case "blue-to-yellow":
		colorModel = &fractal.HueColorModel{Threshold: *iters, HueRange: fractal.BlueToYellow, BoldMode: *boldMode}
	case "red-to-green":
		colorModel = &fractal.HueColorModel{Threshold: *iters, HueRange: fractal.RedToGreen, BoldMode: *boldMode}
	default:
		fmt.Println("Unknown colors:", *colors)
		flag.Usage()
		return
	}

	viz := fractal.Image{
		Fractal:       frac,
		Model:         colorModel,
		FractalBounds: bounds,
		Iters:         *iters,
		PixelSize:     pixelSize,
	}
	cached := viz.ToCachedImage(runtime.NumCPU())

	err = png.Encode(os.Stdout, cached)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}
