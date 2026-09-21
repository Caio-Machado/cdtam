package output

import "image"

type ClipboardOutput struct{}

func (c ClipboardOutput) Write(img image.Image) error
