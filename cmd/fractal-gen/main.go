package main

import (
	"fmt"
	"image/png"
	"os"
	"runtime"

	"github.com/copperium/fractals"
	"github.com/integrii/flaggy"
)

const version = "0.0.0"

var (
	fractalType string
	left        = 0.0
	bottom      = 0.0
	width       = 1.0
	height      = 1.0
	threshold   = 1000.0
	paramStr    = "0+0i"
	iters       = 100
	imgWidth    = 1000
	colors      = "blue-to-yellow"
	boldMode    = false
)

func init() {
	flaggy.SetName("fractal-gen")
	flaggy.SetDescription("generate a fractal from the command line")
	flaggy.SetVersion(version)
	flaggy.AddPositionalValue(&fractalType, "fractal", 1, true, "Type of fractal: mandelbrot or julia")
	flaggy.Float64(&left, "l", "left", "Left bound of image in fractal")
	flaggy.Float64(&bottom, "b", "bottom", "Bottom bound of image in fractal")
	flaggy.Float64(&width, "W", "width", "Width of image in fractal")
	flaggy.Float64(&height, "H", "height", "Height of image in fractal")
	flaggy.Float64(&threshold, "t", "threshold", "Fractal computation threshold")
	flaggy.String(&paramStr, "p", "param", "Complex parameter for Julia fractal")
	flaggy.Int(&iters, "i", "iters", "Number of fractal iterations")
	flaggy.Int(&imgWidth, "w", "image-width", "Width of image (pixels)")
	flaggy.String(&colors, "c", "colors", "Color scheme: blue-to-yellow, red-to-green, or greyscale")
	flaggy.Bool(&boldMode, "bm", "bold-mode", "Enable bold mode!")
	flaggy.Parse()
}

func main() {
	var frac fractals.Fractal
	switch fractalType {
	case "mandelbrot":
		frac = &fractals.Mandelbrot{Threshold: threshold}
	case "julia":
		var realParam, imagParam float64
		_, err := fmt.Sscanf(paramStr, "%f%fi", &realParam, &imagParam)
		if err != nil {
			flaggy.ShowHelpAndExit("Invalid complex format")
		}
		frac = &fractals.Julia{Threshold: threshold, Param: complex(realParam, imagParam)}
	default:
		flaggy.ShowHelpAndExit("Unknown fractal type: options are mandelbrot, julia")
	}

	bounds := fractals.Rect{
		BottomLeft: &fractals.Point{X: left, Y: bottom},
		TopRight:   &fractals.Point{X: left + width, Y: bottom + height},
	}
	pixelSize := (bounds.TopRight.X - bounds.BottomLeft.X) / float64(imgWidth)

	var colorModel fractals.ColorModel
	switch colors {
	case "greyscale":
		colorModel = &fractals.GreyscaleColorModel{Threshold: iters}
	case "blue-to-yellow":
		colorModel = &fractals.HueColorModel{Threshold: iters, HueRange: fractals.BlueToYellow, BoldMode: boldMode}
	case "red-to-green":
		colorModel = &fractals.HueColorModel{Threshold: iters, HueRange: fractals.RedToGreen, BoldMode: boldMode}
	default:
		flaggy.ShowHelpAndExit("Unknown colors: options are greyscale, blue-to-yellow, red-to-green")
	}

	viz := fractals.Image{
		Fractal:       frac,
		Model:         colorModel,
		FractalBounds: bounds,
		Iters:         iters,
		PixelSize:     pixelSize,
	}
	cached := viz.ToCachedImage(runtime.NumCPU())

	err := png.Encode(os.Stdout, cached)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err.Error())
	}
}
