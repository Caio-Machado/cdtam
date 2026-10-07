package capture

import "image"

type Screenshot struct {
	Image image.Image
}

func (s Screenshot) Valid() bool { return s.Image != nil }

func (s Screenshot) Bounds() image.Rectangle {
	if !s.Valid() {
		return image.Rectangle{}
	}
	return s.Image.Bounds()
}

func NewScreenshot(img image.Image) Screenshot {
	return Screenshot{
		Image: img,
	}
}
