package capture

type CaptureTarget interface {
	Capture() (Screenshot, error)
}

type VirtualScreen struct{}

func (vs *VirtualScreen) Capture() (Screenshot, error) {
	return Screenshot{}, nil
}

func NewVirtualScreen() *VirtualScreen {
	return &VirtualScreen{}
}
