package fractal

import (
	"image/color"
	"math"
)

type GreyscaleColorModel struct {
	Threshold int
}

func (m GreyscaleColorModel) Color(result int) color.Color {
	return color.Gray16{Y: uint16(result * (1 << 16) / m.Threshold)}
}

func (m GreyscaleColorModel) ColorModel() color.Model {
	return color.Gray16Model
}

type HueColorModel struct {
	Threshold int
	HueRange
	BoldMode bool // rich shadows!
}

type HueRange struct {
	MinHue, MaxHue float64
}

var (
	BlueToYellow = HueRange{0.667, 0.167}
	RedToGreen   = HueRange{0, 0.333}
)

// convert h, s, v in [0, 1] to RGB
// https://en.wikipedia.org/wiki/HSL_and_HSV, of course
func hsv(h, s, v float64) color.Color {
	h *= 6
	c := s * v
	x := c * (1 - math.Abs(math.Mod(h, 2)-1))

	var r, g, b float64
	switch {
	case 0 <= h && h <= 1:
		r, g, b = c, x, 0
	case 1 < h && h <= 2:
		r, g, b = x, c, 0
	case 2 < h && h <= 3:
		r, g, b = 0, c, x
	case 3 < h && h <= 4:
		r, g, b = 0, x, c
	case 4 < h && h <= 5:
		r, g, b = x, 0, c
	case 5 < h && h <= 6:
		r, g, b = c, 0, x
	}

	m := v - c
	r, g, b = r+m, g+m, b+m

	const max = (1 << 16) - 1
	r, g, b = r*max, g*max, b*max
	return color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: max}
}

func (m HueColorModel) Color(result int) color.Color {
	if result == 0 {
		return color.Black
	}

	hueRange := math.Mod(math.Mod(m.MaxHue-m.MinHue, 1)+1, 1)

	fac := float64(result-1) / float64(m.Threshold-1) // range 0-1
	h := math.Mod(hueRange*fac+m.MinHue, 1)           // range MinHue-MaxHue

	v := 1.0
	if m.BoldMode {
		v = (fac - 0.4) / 0.6
		v *= v
	}

	return hsv(h, 1, v)
}

func (m HueColorModel) ColorModel() color.Model {
	return color.RGBA64Model
}
