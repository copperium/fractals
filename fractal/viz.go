package fractal

import (
	"image"
	"image/color"
)

type ColorModel interface {
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

type Image struct {
	Model         ColorModel
	Fractal       Fractal
	FractalBounds Rect
	Iters         int
	// numeric size per pixel
	PixelSize float64
}

func (i *Image) ImageToFractalPoint(x, y int) Point {
	return Point{
		X: i.FractalBounds.BottomLeft.X + float64(x)*i.PixelSize,
		Y: i.FractalBounds.BottomLeft.Y + float64(y)*i.PixelSize,
	}
}

func (i *Image) FractalToImagePoint(p *Point) (x, y int) {
	x = int((p.X - i.FractalBounds.BottomLeft.X) / i.PixelSize)
	y = int((p.Y - i.FractalBounds.BottomLeft.Y) / i.PixelSize)
	return
}

func (i *Image) ColorModel() color.Model {
	return i.Model.ColorModel()
}

func (i *Image) Bounds() image.Rectangle {
	w := i.FractalBounds.TopRight.X - i.FractalBounds.BottomLeft.X
	h := i.FractalBounds.TopRight.Y - i.FractalBounds.BottomLeft.Y
	return image.Rect(0, 0, int(w/i.PixelSize), int(h/i.PixelSize))
}

func (i *Image) At(x, y int) color.Color {
	point := i.ImageToFractalPoint(x, y)
	result := i.Fractal.At(&point, i.Iters)
	return i.Model.Color(result)
}

func (i *Image) ToCachedImage() image.Image {
	img := image.NewRGBA(i.Bounds())
	results := make(chan PointResult)
	numResults := Compute(i.Fractal, i.FractalBounds, i.PixelSize, i.Iters, results)
	for j := 0; j < numResults; j++ {
		result := <-results
		x, y := i.FractalToImagePoint(result.Point)
		img.Set(x, y, i.Model.Color(result.Result))
	}
	return img
}
