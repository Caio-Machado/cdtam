package capture

import "image"

type Monitor struct {
	ID      int32
	Name    string
	Bounds  image.Rectangle
	Prymary bool
}

func DetectPrimaryMonitor() *Monitor

func DetectAllMonitors() []*Monitor
