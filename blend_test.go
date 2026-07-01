package alphablender_test

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/chixm/alphablender"
)

func TestAlphaBlendImagesOfSameSize(t *testing.T) {
	back, err := os.Open(`./background.png`)
	if err != nil {
		t.Error(err)
		return
	}
	defer back.Close()
	star, err := os.Open(`./star.png`)
	if err != nil {
		t.Error(err)
		return
	}
	defer star.Close()

	backImage, err := png.Decode(back)
	if err != nil {
		t.Error(err)
		return
	}
	starImage, err := png.Decode(star)
	if err != nil {
		t.Error(err)
		return
	}
	blendedImage := alphablender.Blend(starImage, backImage)

	newImage, err := os.Create(`createdImage.png`)
	if err != nil {
		t.Error(err)
		return
	}
	defer newImage.Close()

	png.Encode(newImage, blendedImage)
}

func TestAlphaBlendImagesOfDifferentSizeFromFiles(t *testing.T) {
	back, err := os.Open(`./background_square.png`)
	if err != nil {
		t.Error(err)
		return
	}
	defer back.Close()
	star, err := os.Open(`./star.png`)
	if err != nil {
		t.Error(err)
		return
	}
	defer star.Close()

	backImage, err := png.Decode(back)
	if err != nil {
		t.Error(err)
		return
	}
	starImage, err := png.Decode(star)
	if err != nil {
		t.Error(err)
		return
	}
	blendedImage := alphablender.Blend(starImage, backImage)

	newImage, err := os.Create(`createdImage_square.png`)
	if err != nil {
		t.Error(err)
		return
	}
	defer newImage.Close()

	png.Encode(newImage, blendedImage)
}

func TestAlphaBlendImagesOfDifferentSize(t *testing.T) {
	// src is smaller than dst
	backImage := image.NewNRGBA(image.Rect(0, 0, 4, 4))
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			backImage.Set(x, y, color.NRGBA{R: 10, G: 20, B: 30, A: 255})
		}
	}

	starImage := image.NewRGBA(image.Rect(0, 0, 2, 2))
	for y := 0; y < 2; y++ {
		for x := 0; x < 2; x++ {
			starImage.Set(x, y, color.RGBA{R: 200, G: 0, B: 0, A: 255})
		}
	}

	blended := alphablender.Blend(starImage, backImage)

	if blended.Bounds() != backImage.Bounds() {
		t.Errorf("expected output bounds %v, got %v", backImage.Bounds(), blended.Bounds())
	}

	// covered by src: fully replaced by opaque src color (blendFormula's
	// float rounding can be off by one on the alpha channel)
	if got := blended.RGBAAt(0, 0); got.R != 200 || got.G != 0 || got.B != 0 || got.A < 254 {
		t.Errorf("expected src color at (0,0), got %v", got)
	}
	// outside src bounds: dst color shows through unchanged
	if got := blended.RGBAAt(3, 3); got.R != 10 || got.G != 20 || got.B != 30 || got.A < 254 {
		t.Errorf("expected dst color at (3,3), got %v", got)
	}
}
