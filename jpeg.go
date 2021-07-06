package tiff

import (
	"image"
	"image/draw"
)

func (d *decoder) decodeJPEG(dst image.Image, xmin, ymin, xmax, ymax int) {
	rMaxX := minInt(xmax, dst.Bounds().Max.X)
	rMaxY := minInt(ymax, dst.Bounds().Max.Y)

	var img draw.Image
	switch d.mode {
	case mGray, mGrayInvert:
		if d.bpp == 16 {
			img = dst.(*image.Gray16)
		} else {
			img = dst.(*image.Gray)
		}
	case mPaletted:
		img = dst.(*image.Paletted)
	case mRGB, mNRGBA, mRGBA:
		if d.bpp == 16 {
			img = dst.(*image.RGBA64)
		} else {
			img = dst.(*image.RGBA)
		}
	case mCMYK:
		// d.bpp must be 8
		img = dst.(*image.CMYK)
	}

	for y := 0; y+ymin < rMaxY; y++ {
		for x := 0; x+xmin < rMaxX; x++ {
			img.Set(x+xmin, y+ymin, d.tmp.At(x, y))
		}
	}
}
