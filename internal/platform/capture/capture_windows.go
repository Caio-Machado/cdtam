package platform

import (
	"errors"
	"fmt"
	"image"
	"runtime"
	"unsafe"

	"github.com/caio-machado/cdtam/internal/imaging"
	"github.com/caio-machado/cdtam/internal/sys/win32"
)

func CaptureVirtualScreen() (image.Image, error) {
	vsb, err := getVirtualScreenBounds()
	if err != nil {
		return &image.NRGBA{}, fmt.Errorf("capture virtual screen: %w", err)
	}

	hdc, err := win32.GetDC(0)
	if err != nil {
		return &image.NRGBA{}, fmt.Errorf("capture virtual screen: %w", err)
	}
	// TODO: Preciso de um valor valido para o primeiro parametro
	defer win32.ReleaseDC(0, hdc)

	cDCh, err := win32.CreateCompatibleDC(hdc)
	if err != nil {
		return &image.NRGBA{}, fmt.Errorf("capture virtual screen: %w", err)
	}
	defer win32.DeleteDC(cDCh)

	cBitmap, err := win32.CreateCompatibleBitmap(
		cDCh,
		int32(vsb.Bounds().Dx()),
		int32(vsb.Bounds().Dy()),
	)
	if err != nil {
		return &image.NRGBA{}, fmt.Errorf("capture virtual screen: %w", err)
	}
	defer win32.DeleteObject(win32.HGDIOBJ(cBitmap))

	oldObj, err := win32.SelectObject(cDCh, win32.HGDIOBJ(cBitmap))
	if err != nil {
		return &image.NRGBA{}, fmt.Errorf("capture virtual screen: %w", err)
	}
	defer win32.SelectObject(cDCh, oldObj)

	err = win32.BitBlt(
		cDCh,
		0,
		0,
		int32(vsb.Dx()),
		int32(vsb.Dy()),
		hdc,
		int32(vsb.Min.X),
		int32(vsb.Min.Y),
		win32.ROP,
	)
	if err != nil {
		return &image.NRGBA{}, fmt.Errorf("capture virtual screen: %w", err)
	}

	pixels := make([]byte, vsb.Dx()*vsb.Dy()*4)
	bmi := win32.NewBitmapInfoHeader(
		vsb.Dx(),
		vsb.Dy(),
	)

	diBits, err := win32.GetDIBits(
		cDCh,
		cBitmap,
		0,
		uint32(vsb.Dy()),
		unsafe.Pointer(&pixels[0]),
		unsafe.Pointer(&bmi),
		win32.DIB_RGB_COLORS,
	)
	if diBits != int32(vsb.Dy()) {
		return &image.NRGBA{}, errors.New("capture virtual screen: failed capture")
	}
	if err != nil {
		return &image.NRGBA{}, fmt.Errorf("capture virtual screen: %w", err)
	}

	runtime.KeepAlive(pixels)
	runtime.KeepAlive(bmi)

	img, err := imaging.BGRAToNRGBA(vsb.Dx(), vsb.Dy(), pixels)
	if err != nil {
		return &image.NRGBA{}, fmt.Errorf("capture virtual screen: %w", err)
	}

	return img, nil
}

func getVirtualScreenBounds() (image.Rectangle, error) {
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
