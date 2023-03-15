package alphablender

import (
	"image"
	"image/color"
	"math"
	"reflect"
)

// Blend two images to One
// src image to blend
// dst background image
func Blend(src image.Image, dst image.Image) *image.RGBA {

	var rowLen int
	var srcPix []uint8
	var dstPix []uint8

	if reflect.TypeOf(src).Elem().Name() == reflect.TypeOf(image.RGBA{}).Name() {
		r := src.(*image.RGBA)
		srcPix = r.Pix
	} else if reflect.TypeOf(src).Elem().Name() == reflect.TypeOf(image.NRGBA{}).Name() {
		r := src.(*image.NRGBA)
		srcPix = r.Pix
	}
	if reflect.TypeOf(dst).Elem().Name() == reflect.TypeOf(image.RGBA{}).Name() {
		d := dst.(*image.RGBA)
		rowLen = d.Stride
		dstPix = d.Pix
	} else if reflect.TypeOf(dst).Elem().Name() == reflect.TypeOf(image.NRGBA{}).Name() {
		d := dst.(*image.NRGBA)
		rowLen = d.Stride
		dstPix = d.Pix
	}
	output := make([]color.Color, len(dstPix))
	// blending two images togather
	for i := 0; i < len(dstPix)-1; i += 4 {
		blended := blendFormula(color.RGBA{R: srcPix[i], G: srcPix[i+1], B: srcPix[i+2], A: srcPix[i+3]}, color.RGBA{R: dstPix[i], G: dstPix[i+1], B: dstPix[i+2], A: dstPix[i+3]})
		output = append(output, blended)
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
