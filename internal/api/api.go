package api

import (
	"image/png"
	"runtime"

	"github.com/copperium/fractals"
	"github.com/gin-gonic/gin"
)

func testEndpoint(c *gin.Context) {
	const iters = 100
	viz := fractals.Image{
		Fractal: &fractals.Mandelbrot{Threshold: 1000},
		Model:   fractals.HueColorModel{Threshold: iters, HueRange: fractals.BlueToYellow},
		FractalBounds: fractals.Rect{
			BottomLeft: &fractals.Point{X: -1, Y: -1},
			TopRight:   &fractals.Point{X: 1, Y: 1},
		},
		Iters:     iters,
		PixelSize: 0.001,
	}
	image := viz.ToCachedImage(runtime.NumCPU())
	c.Writer.Header()["Content-Type"] = []string{"image/png"}
	err := png.Encode(c.Writer, image)
	if err != nil {
		panic(err)
	}
}

func Setup(addr ...string) error {
	router := gin.Default()
	router.GET("/test", testEndpoint)
	err := router.Run(addr...)
	return err
}
