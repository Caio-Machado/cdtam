package output

import (
	"github.com/caio-machado/cdtam/internal/capture"
)

type OutputOpts struct {
	OutDir    string
	Clipboard bool
}

func (o OutputOpts) Resolve() ([]capture.Output, error) {
	var out []capture.Output

	if o.OutDir != "" {
		fileOutput, err := NewFileOutput(o.OutDir)
		if err != nil {
			return nil, err
		}
		out = append(out, fileOutput)
	}

	if o.Clipboard {
		out = append(out, ClipboardOutput{})
	}

	return out, nil
}

func NewOutputOpts(filePath string, clipboard bool) OutputOpts {
	return OutputOpts{
		OutDir:    filePath,
		Clipboard: clipboard,
	}
}
