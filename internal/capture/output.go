package capture

type Output interface {
	Write(s Screenshot) error
}
