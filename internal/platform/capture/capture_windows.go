package platform

import (
	"errors"
	"fmt"
	"image"
	"runtime"
	"strconv"
	"strings"
	"unsafe"

	"github.com/caio-machado/cdtam/internal/imaging"
	"github.com/caio-machado/cdtam/internal/platform/screen"
	"github.com/caio-machado/cdtam/internal/sys/win32"
)

func CaptureVirtualScreen() (image.Image, error) {
	vs, err := getVirtualScreen()
	if err != nil {
		return &image.NRGBA{}, fmt.Errorf("capture virtual screen: %w", err)
	}

	img, err := captureBounds(vs.Bounds())
	if err != nil {
		return &image.NRGBA{}, fmt.Errorf("capture virtual screen: %w", err)
	}

	return img, nil
}

func CaptureMonitor(arg string) (image.Image, error) {
	target, err := resolveMonitor(arg)
	if err != nil {
		return nil, fmt.Errorf("capture monitor: %w", err)
	}

	img, err := captureBounds(target.Bounds)
	if err != nil {
		return &image.NRGBA{}, fmt.Errorf("capture monitor: %w", err)
	}

	return img, nil
}

func getVirtualScreen() (image.Rectangle, error) {
	x := win32.GetSystemMetrics(win32.SM_XVIRTUALSCREEN)
	y := win32.GetSystemMetrics(win32.SM_YVIRTUALSCREEN)
	width := win32.GetSystemMetrics(win32.SM_CXVIRTUALSCREEN)
	height := win32.GetSystemMetrics(win32.SM_CYVIRTUALSCREEN)
	if width <= 0 || height <= 0 {
		return image.Rectangle{}, fmt.Errorf("invalid virtual screen dimensions: width=%d height=%d", width, height)
	}

	return image.Rect(
		int(x),
		int(y),
		int(width),
		int(height),
	), nil
}

func resolveMonitor(arg string) (screen.Monitor, error) {
	monitors, err := screen.ListMonitors()
	if err != nil {
		return screen.Monitor{}, err
	}
	if len(monitors) == 0 {
		return screen.Monitor{}, errors.New("no monitors found")
	}

	arg = strings.TrimSpace(arg)
	if strings.EqualFold(arg, "primary") {
		return monitors[0], nil
	}

	if index, err := strconv.Atoi(arg); err == nil {
		if index < 0 || index >= len(monitors) {
			return screen.Monitor{}, fmt.Errorf("monitor index %d out of range [0..%d]", index, len(monitors)-1)
		}
		return monitors[index], nil
	}

	for _, m := range monitors {
		if strings.EqualFold(m.DeviceName, arg) ||
			strings.EqualFold(strings.TrimPrefix(m.DeviceName, `\\.\`), arg) {
			return m, nil
		}
	}

	return screen.Monitor{}, fmt.Errorf("monitor %q not found", arg)
}

func captureBounds(bounds image.Rectangle) (image.Image, error) {
	width, height := bounds.Dx(), bounds.Dy()
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("invalid bounds %v", bounds)
	}

	screenDC, err := win32.GetDC(0)
	if err != nil {
		return nil, err
	}
	defer win32.ReleaseDC(0, screenDC)

	memDC, err := win32.CreateCompatibleDC(screenDC)
	if err != nil {
		return nil, err
	}
	defer win32.DeleteDC(memDC)

	bitmap, err := win32.CreateCompatibleBitmap(screenDC, int32(width), int32(height))
	if err != nil {
		return nil, err
	}
	defer win32.DeleteObject(win32.HGDIOBJ(bitmap))

	old, err := win32.SelectObject(memDC, win32.HGDIOBJ(bitmap))
	if err != nil {
		return nil, err
	}
	defer win32.SelectObject(memDC, old)

	if err := win32.BitBlt(
		memDC,
		0,
		0,
		int32(width),
		int32(height),
		screenDC,
		int32(bounds.Min.X),
		int32(bounds.Min.Y),
		win32.ROP,
	); err != nil {
		return nil, err
	}

	if _, err := win32.SelectObject(memDC, old); err != nil {
		return nil, err
	}

	pixels := make([]byte, width*height*4)
	bmi := win32.NewBitmapInfoHeader(width, height)

	lines, err := win32.GetDIBits(
		memDC,
		bitmap,
		0,
		uint32(height),
		unsafe.Pointer(&pixels[0]),
		unsafe.Pointer(&bmi),
		win32.DIB_RGB_COLORS,
	)
	if err != nil {
		return nil, err
	}
	if lines != int32(height) {
		return nil, fmt.Errorf("capture scanned %d of %d lines", lines, height)
	}

	runtime.KeepAlive(pixels)
	runtime.KeepAlive(bmi)

	img, err := imaging.BGRAToNRGBA(width, height, pixels)
	if err != nil {
		return nil, err
	}

	return img, nil
}
