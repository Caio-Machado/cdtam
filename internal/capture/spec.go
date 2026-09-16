package capture

import (
	"github.com/caio-machado/cdtam/internal/output"
)

type CaptureTarget int

const (
	VirtualScreen CaptureTarget = iota
	Monitor
)

type CaptureSpec struct {
	Target    CaptureTarget
	MonitorID int
	Output    output.OutputSpec
}

func BuildSpec(output string, clipboard bool, monitorID int) CaptureSpec

func (cp CaptureSpec) Validate() error
