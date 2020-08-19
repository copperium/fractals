package viz

import (
	"github.com/copperium/fractals/fraclib"
	"image"
	"image/color"
)

type FractalColorModel interface {
	Color(int) color.Color
	ColorModel() color.Model
}

type ThresholdModel struct {
	Threshold int
}

// greyscale for now
func (m ThresholdModel) Color(result int) color.Color {
	return color.Gray16{Y: uint16(result * (1 << 16) / m.Threshold)}
}

func (m ThresholdModel) ColorModel() color.Model {
	return color.Gray16Model
}

type FractalImage struct {
	Model         FractalColorModel
	Fractal       fraclib.Fractal
	FractalBounds fraclib.Rect
	Iters         int
	// numeric size per pixel
	PixelSize float64
}

func (i *FractalImage) ImageToFractalPoint(x, y int) fraclib.Point {
	return fraclib.Point{
		X: i.FractalBounds.BottomLeft.X + float64(x)*i.PixelSize,
		Y: i.FractalBounds.BottomLeft.Y + float64(y)*i.PixelSize,
	}
}

func (i *FractalImage) FractalToImagePoint(p *fraclib.Point) (x, y int) {
	x = int((p.X - i.FractalBounds.BottomLeft.X) / i.PixelSize)
	y = int((p.Y - i.FractalBounds.BottomLeft.Y) / i.PixelSize)
	return
}

func (i *FractalImage) ColorModel() color.Model {
	return i.Model.ColorModel()
}

func (i *FractalImage) Bounds() image.Rectangle {
	w := i.FractalBounds.TopRight.X - i.FractalBounds.BottomLeft.X
	h := i.FractalBounds.TopRight.Y - i.FractalBounds.BottomLeft.Y
	return image.Rect(0, 0, int(w/i.PixelSize), int(h/i.PixelSize))
}

func (i *FractalImage) At(x, y int) color.Color {
	point := i.ImageToFractalPoint(x, y)
	result := i.Fractal.At(&point, i.Iters)
	return i.Model.Color(result)
}

func (i *FractalImage) ToCachedImage() image.Image {
	img := image.NewRGBA(i.Bounds())
	results := make(chan fraclib.PointResult)
	numResults := fraclib.Compute(i.Fractal, i.FractalBounds, i.PixelSize, i.Iters, results)
	for j := 0; j < numResults; j++ {
		result := <-results
		x, y := i.FractalToImagePoint(result.Point)
		img.Set(x, y, i.Model.Color(result.Result))
	}
	return img
}
