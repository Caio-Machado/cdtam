package output

import (
	"github.com/caio-machado/cdtam/internal/capture"
	"github.com/caio-machado/cdtam/internal/imaging"
	"github.com/caio-machado/cdtam/internal/platform"
)

type ClipboardOutput struct{}

func (c ClipboardOutput) Write(s capture.Screenshot) error {
	b := s.Image.Bounds()
	bgra, err := imaging.ImageToBGRA(s.Image)
	if err != nil {
		return err
	}

	if err := platform.WriteClipboard(b.Dx(), b.Dy(), bgra); err != nil {
		return err
	}

	return nil
}
