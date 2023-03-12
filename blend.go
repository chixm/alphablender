package alphablender

import (
	"image"
	"image/color"
	"math"
)

// Blend two images to One
// src image to blend
// dst background image
func Blend(src *image.RGBA, dst *image.RGBA) *image.RGBA {
	output := []color.Color{}
	rowLen := dst.Stride
	for i := 0; i < len(dst.Pix); i += 4 {
		output[i] = blendFormula(color.RGBA{R: src.Pix[i], G: src.Pix[i+1], B: src.Pix[i+2], A: src.Pix[i+3]}, color.RGBA{R: dst.Pix[i], G: dst.Pix[i+1], B: dst.Pix[i+2], A: dst.Pix[i+3]})
	}
	newPix := []uint8{}
	for _, x := range output {
		r, g, b, a := x.RGBA()
		newPix = append(newPix, uint8(r), uint8(g), uint8(b), uint8(a))
	}
	fx := image.RGBA{Stride: rowLen, Rect: dst.Bounds(), Pix: newPix}
	return &fx
}

func blendFormula(src color.Color, dst color.Color) color.Color {
	r, g, b, a := src.RGBA()
	r2, g2, b2, a2 := dst.RGBA()

	fr := float64(r / math.MaxUint32)
	fg := float64(g / math.MaxUint32)
	fb := float64(b / math.MaxUint32)
	fa := float64(a / math.MaxUint32)

	fr2 := float64(r2 / math.MaxUint32)
	fg2 := float64(g2 / math.MaxUint32)
	fb2 := float64(b2 / math.MaxUint32)
	fa2 := float64(a2 / math.MaxUint32)

	outA := fa + fa2*(1-fa)
	outR := fr + fr2*(1-fr)
	outG := fg + fg2*(1-fg)
	outB := fb + fb2*(1-fb)

	return &Pic{r: outR, g: outG, b: outB, a: outA}
}
