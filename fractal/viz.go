package fractal

import (
	"image"
	"image/color"
	"math/big"
)

type ColorModel interface {
	Color(int) color.Color
	ColorModel() color.Model
}

type ThresholdColorModel struct {
	Threshold int
}

// greyscale for now
func (m ThresholdColorModel) Color(result int) color.Color {
	return color.Gray16{Y: uint16(result * (1 << 16) / m.Threshold)}
}

func (m ThresholdColorModel) ColorModel() color.Model {
	return color.Gray16Model
}

type Image struct {
	Model         ColorModel
	Fractal       Fractal
	FractalBounds Rect
	Iters         int
	// numeric size per pixel
	PixelSize *big.Rat
}

// ImageToFractalPoint returns the exact point in the fractal corresponding to the image pixel.
func (i *Image) ImageToFractalPoint(x, y int) *Point {
	var p Point
	p.X.Mul(i.PixelSize, big.NewRat(int64(x), 1))
	p.X.Add(&p.X, &i.FractalBounds.BottomLeft.X)
	p.Y.Mul(i.PixelSize, big.NewRat(int64(y), 1))
	p.Y.Add(&p.Y, &i.FractalBounds.BottomLeft.Y)
	return &p
}

// FractalToImagePoint returns the approximate image pixel corresponding to an exact fractal point.
func (i *Image) FractalToImagePoint(p *Point) (x, y int) {
	var rx, ry big.Rat
	rx.Sub(&p.X, &i.FractalBounds.BottomLeft.X)
	rx.Quo(&rx, i.PixelSize)
	ry.Sub(&p.Y, &i.FractalBounds.BottomLeft.Y)
	ry.Quo(&ry, i.PixelSize)
	fx, _ := rx.Float32()
	fy, _ := ry.Float32()
	return int(fx), int(fy)
}

func (i *Image) ColorModel() color.Model {
	return i.Model.ColorModel()
}

func (i *Image) Bounds() image.Rectangle {
	var w, h big.Rat
	w.Sub(&i.FractalBounds.TopRight.X, &i.FractalBounds.BottomLeft.X)
	h.Sub(&i.FractalBounds.TopRight.Y, &i.FractalBounds.BottomLeft.Y)
	w.Quo(&w, i.PixelSize)
	h.Quo(&h, i.PixelSize)
	fw, _ := w.Float32()
	fh, _ := h.Float32()
	return image.Rect(0, 0, int(fw), int(fh))
}

func (i *Image) At(x, y int) color.Color {
	point := i.ImageToFractalPoint(x, y)
	result := i.Fractal.At(point, i.Iters)
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
