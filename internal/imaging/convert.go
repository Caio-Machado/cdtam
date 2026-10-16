package imaging

import (
	"errors"
	"fmt"
	"image"
)

func BGRAToNRGBA(width, height int, pixels []byte) (image.Image, error) {
	if width == 0 || height == 0 {
		return nil, fmt.Errorf("invalid dimensions")
	}

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for i := 0; i+3 < len(pixels); i += 4 {
		b := pixels[i+0]
		g := pixels[i+1]
		r := pixels[i+2]

		img.Pix[i+0] = r
		img.Pix[i+1] = g
		img.Pix[i+2] = b
		img.Pix[i+3] = 255 // force alpha channel to 255
	}

	return img, nil
}

func NRGBAToBGRA(img image.Image) ([]byte, error) {
	nrgba, ok := img.(*image.NRGBA)
	if !ok {
		return nil, errors.New("image is not NRGBA")
	}
	imgSize := len(nrgba.Pix)
	slice := make([]byte, imgSize)
	for p := 0; p < imgSize; p += 4 {
		r := nrgba.Pix[p+0]
		g := nrgba.Pix[p+1]
		b := nrgba.Pix[p+2]

		slice[p+0] = b
		slice[p+1] = g
		slice[p+2] = r
		slice[p+3] = 255 // force alpha channel to 255
	}
	return slice, nil
}
