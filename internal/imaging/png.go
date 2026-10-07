package imaging

import (
	"image"
	"image/png"
	"io"
)

func EncodePNG(w io.Writer, img image.Image) error {
	return png.Encode(w, img)
}
