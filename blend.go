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
	} else {
		panic(`Blend only available for type RGBA or NRGBA`)
	}
	if reflect.TypeOf(dst).Elem().Name() == reflect.TypeOf(image.RGBA{}).Name() {
		d := dst.(*image.RGBA)
		rowLen = d.Stride
		dstPix = d.Pix
	} else if reflect.TypeOf(dst).Elem().Name() == reflect.TypeOf(image.NRGBA{}).Name() {
		d := dst.(*image.NRGBA)
		rowLen = d.Stride
		dstPix = d.Pix
	} else {
		panic(`Blend only available for type RGBA or NRGBA`)
	}
	output := make([]uint8, 0, len(dstPix))
	// blending two images togather
	for i := 0; i < len(dstPix); i += 4 {
		r, g, b, a := blendFormula(color.RGBA{R: srcPix[i], G: srcPix[i+1], B: srcPix[i+2], A: srcPix[i+3]}, color.RGBA{R: dstPix[i], G: dstPix[i+1], B: dstPix[i+2], A: dstPix[i+3]})
		output = append(output, r, g, b, a)
	}
	return &image.RGBA{Stride: rowLen, Rect: dst.Bounds(), Pix: output}
}

func blendFormula(src color.Color, dst color.Color) (r, g, b, a uint8) {
	r1, g1, b1, a1 := src.RGBA()
	r2, g2, b2, a2 := dst.RGBA()

	alpha1 := float64(a1<<16) / float64(math.MaxUint32)
	alpha2 := float64(a2<<16) / float64(math.MaxUint32)

	outA := alpha1 + alpha2*(1-alpha1)
	outR := float64(r1) + float64(r2)*(1-alpha1)
	outG := float64(g1) + float64(g2)*(1-alpha1)
	outB := float64(b1) + float64(b2)*(1-alpha1)

	return uint8(outR), uint8(outG), uint8(outB), uint8(outA * math.MaxUint8)
}
