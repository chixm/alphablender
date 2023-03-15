package alphablender_test

import (
	"image/png"
	"os"
	"testing"

	"github.com/chixm/alphablender"
)

func TestAlphaBlendImagesOfSameSize(t *testing.T) {
	f, err := os.Open(`./background.png`)
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()
	star, err := os.Open(`./star.png`)
	if err != nil {
		t.Error(err)
		return
	}
	defer star.Close()

	backImage, err := png.Decode(f)
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
