# alphablender
Alpha Blending packege for go

## very simple function to join two images togather
Function uses Alpha Blend for calculating pixels.
https://en.wikipedia.org/wiki/Alpha_compositing

## How to use
- Blend

just set blending image to first arg, background image to second arg.

example:
```
	blendedImage := alphablender.Blend(starImage, backImage)
```

starImage and backImage variable is image.Image type of Go.

see blend_test.go , describes how to blend two png file images togather.

You can blend two images togather

<nobr>
<img src="background.png" width="50px">
+
<img src="star.png" width="50px">
=
<img src="createdImage.png" width="50px">
</nobr>

# Limitation
- two images must be the same size.
- images must be RGBA or NRGBA type.

