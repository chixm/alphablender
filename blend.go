package alphablender

import (
	"image"
	"image/color"
	"math"
)

// Blend two images to One
// src image to blend
// dst background image
//
// src and dst do not need to be the same size. The output canvas matches
// dst's bounds; src is aligned to dst's origin and cropped/padded with
// transparency where it doesn't cover the full dst area.
func Blend(src image.Image, dst image.Image) *image.RGBA {
	if !isBlendable(src) || !isBlendable(dst) {
		panic(`Blend only available for type RGBA or NRGBA`)
	}

	dstBounds := dst.Bounds()
	srcBounds := src.Bounds()

	output := image.NewRGBA(dstBounds)
	// blending two images togather
	for y := dstBounds.Min.Y; y < dstBounds.Max.Y; y++ {
		for x := dstBounds.Min.X; x < dstBounds.Max.X; x++ {
			sp := image.Pt(srcBounds.Min.X+(x-dstBounds.Min.X), srcBounds.Min.Y+(y-dstBounds.Min.Y))

			var srcColor color.Color = color.RGBA{}
			if sp.In(srcBounds) {
				srcColor = src.At(sp.X, sp.Y)
			}

			r, g, b, a := blendFormula(srcColor, dst.At(x, y))
			output.Set(x, y, color.RGBA{R: r, G: g, B: b, A: a})
		}
	}
	return output
}

func isBlendable(img image.Image) bool {
	switch img.(type) {
	case *image.RGBA, *image.NRGBA:
		return true
	default:
		return false
	}
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
