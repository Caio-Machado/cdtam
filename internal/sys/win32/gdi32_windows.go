package win32

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var libGdi32 = windows.NewLazySystemDLL("gdi32.dll")

var (
	funcCreateCompatibleDC     = libGdi32.NewProc("CreateCompatibleDC")
	funcDeleteDC               = libGdi32.NewProc("DeleteDC")
	funcCreateCompatibleBitmap = libGdi32.NewProc("CreateCompatibleBitmap")
	funcSelectObject           = libGdi32.NewProc("SelectObject")
	funcDeleteObject           = libGdi32.NewProc("DeleteObject")
	funcBitBlt                 = libGdi32.NewProc("BitBlt")
	funcGetDIBits              = libGdi32.NewProc("GetDIBits")
)

func CreateCompatibleDC(hdc HDC) (HDC, error) {
	r, _, err := funcCreateCompatibleDC.Call(uintptr(hdc))
	if r == 0 {
		return 0, fmt.Errorf("CreateCompatibleDC: %w", err)
	}
	return HDC(r), nil
}

func DeleteDC(hdc HDC) error {
	r, _, err := funcDeleteDC.Call(uintptr(hdc))
	if r == 0 {
		return fmt.Errorf("DeleteDC: %w", err)
	}
	return nil
}

func CreateCompatibleBitmap(hdc HDC, cx, cy int32) (HBITMAP, error) {
	r, _, err := funcCreateCompatibleBitmap.Call(uintptr(hdc), uintptr(cx), uintptr(cy))
	if r == 0 {
		return 0, fmt.Errorf("CreateCompatibleBitmap: %w", err)
	}
	return HBITMAP(r), nil
}

func SelectObject(hdc HDC, h HGDIOBJ) (HGDIOBJ, error) {
	r, _, err := funcSelectObject.Call(uintptr(hdc), uintptr(h))
	if r == 0 {
		return 0, fmt.Errorf("SelectObject: %w", err)
	}
	return HGDIOBJ(r), nil
}

func DeleteObject(ho HGDIOBJ) error {
	r, _, err := funcDeleteObject.Call(uintptr(ho))
	if r == 0 {
		return fmt.Errorf("DeleteObject: %w", err)
	}
	return nil
}

func BitBlt(hdc HDC, x, y, cx, cy int32, hdcSrc HDC, x1, y1 int32, rop DWORD) error {
	r, _, err := funcBitBlt.Call(
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(hdcSrc),
		uintptr(x1),
		uintptr(y1),
		uintptr(rop),
	)
	if r == 0 {
		return fmt.Errorf("BitBlt: %w", err)
	}
	return nil
}

func GetDIBits(hdc HDC, hbm HBITMAP, start uint32, cLines uint32, lpvBits unsafe.Pointer, lpbmi unsafe.Pointer, usage uint32) (int32, error) {
	r, _, err := funcGetDIBits.Call(
		uintptr(hdc),
		uintptr(hbm),
		uintptr(start),
		uintptr(cLines),
		uintptr(lpvBits),
		uintptr(lpbmi),
		uintptr(usage),
	)
	if r == 0 {
		return 0, fmt.Errorf("GetDIBits: %w", err)
	}

	return int32(r), nil
}
