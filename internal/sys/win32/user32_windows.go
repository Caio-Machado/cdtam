package win32

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var libUser32 = windows.NewLazyDLL("user32.dll")

var (
	funcGetSystemMetrics    = libUser32.NewProc("GetSystemMetrics")
	funcGetDC               = libUser32.NewProc("GetDC")
	funcReleaseDC           = libUser32.NewProc("ReleaseDC")
	funcGetMonitorInfoW     = libUser32.NewProc("GetMonitorInfoW")
	funcEnumDisplayMonitors = libUser32.NewProc("EnumDisplayMonitors")
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

func GetMonitorInfoW(hMonitor HMONITOR, lpmi *MonitorInfoEx) error {
	r, _, err := funcGetMonitorInfoW.Call(uintptr(hMonitor), uintptr(unsafe.Pointer(lpmi)))
	if r == 0 {
		return fmt.Errorf("GetMonitorInfoW: %w", err)
	}
	return nil
}

//go:uintptrescapes
func EnumDisplayMonitors(hdc HDC, lprcClip uintptr, lpfnEnum uintptr, dwData uintptr) error {
	r, _, err := funcEnumDisplayMonitors.Call(
		uintptr(hdc),
		lprcClip,
		uintptr(lpfnEnum),
		dwData,
	)
	if r == 0 {
		return fmt.Errorf("EnumDisplayMonitors: %w", err)
	}
	return nil
}
