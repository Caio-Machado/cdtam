package capture

import "errors"

func Run(target CaptureTarget, outs []Output) error {
	screenshot, err := target.Capture()
	if err != nil {
		return err
	}

	if !screenshot.Valid() {
		return errors.New("capture produced an empty screenshot")
	}

	for _, out := range outs {
		if err := out.Write(screenshot); err != nil {
			return err
		}
	}
	return nil
}
