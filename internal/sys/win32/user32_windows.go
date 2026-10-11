package win32

import (
	"fmt"

	"golang.org/x/sys/windows"
)

var libUser32 = windows.NewLazyDLL("user32.dll")

var (
	funcGetSystemMetrics = libUser32.NewProc("GetSystemMetrics")
	funcGetDC            = libUser32.NewProc("GetDC")
	funcReleaseDC        = libUser32.NewProc("ReleaseDC")
)

func GetSystemMetrics(nIndex int32) int32 {
	metric, _, _ := funcGetSystemMetrics.Call(uintptr(nIndex))
	return int32(metric)
}

func GetDC(hWnd HWND) (HDC, error) {
	r, _, err := funcGetDC.Call(uintptr(hWnd))
	if r == 0 {
		return 0, fmt.Errorf("GetDC: %w", err)
	}
	return HDC(r), nil
}

func ReleaseDC(hWnd HWND, hDc HDC) error {
	r, _, err := funcReleaseDC.Call(uintptr(hWnd), uintptr(hDc))
	if r == 0 {
		return fmt.Errorf("ReleasDC: %w", err)
	}

	return nil
}
