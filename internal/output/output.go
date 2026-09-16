package output

import (
	"image"
)

type Output interface {
	Write(img image.Image) error
}

type OutputSpec struct {
	SaveFile  bool
	FilePath  string
	Clipboard bool
}

func (os OutputSpec) Validate() error

func (os OutputSpec) Resolve() ([]Output, error)

func NewOutputSpec(saveFile bool, filePath string, clipboard bool) OutputSpec {
	return OutputSpec{
		SaveFile:  saveFile,
		FilePath:  filePath,
		Clipboard: clipboard,
	}
}
