package screen

import (
	"fmt"
	"image"
	"runtime"
	"unsafe"

	"github.com/caio-machado/cdtam/internal/sys/win32"
	"golang.org/x/sys/windows"
)

type Monitor struct {
	Index      int
	DeviceName string
	Bounds     image.Rectangle
	WorkArea   image.Rectangle

	handle win32.HMONITOR
}

type enumState struct {
	monitors []Monitor
	err      error
}

var enumMonitorsCB = windows.NewCallback(enumMonitorsProc)

func enumMonitorsProc(hMonitor uintptr, _ uintptr, _ uintptr, dwData uintptr) uintptr {
	st := (*enumState)(unsafe.Pointer(dwData))
	handle := win32.HMONITOR(hMonitor)

	var info win32.MonitorInfoEx
	info.CbSize = uint32(unsafe.Sizeof(win32.MonitorInfoEx{}))
	if err := win32.GetMonitorInfoW(handle, &info); err != nil {
		if st.err == nil {
			st.err = err
		}
		return 1
	}

	var monitor Monitor
	monitor.handle = handle
	monitor.DeviceName = windows.UTF16ToString(info.SzDevice[:])
	monitor.Bounds = image.Rect(
		int(info.RcMonitor.Left),
		int(info.RcMonitor.Top),
		int(info.RcMonitor.Right),
		int(info.RcMonitor.Bottom),
	)
	monitor.WorkArea = image.Rect(
		int(info.RcWork.Left),
		int(info.RcWork.Top),
		int(info.RcWork.Right),
		int(info.RcWork.Bottom),
	)
	st.monitors = append(st.monitors, monitor)
	return 1
}

func ListMonitors() ([]Monitor, error) {
	st := &enumState{}
	err := win32.EnumDisplayMonitors(0, 0, enumMonitorsCB, uintptr(unsafe.Pointer(st)))
	runtime.KeepAlive(st)
	if err != nil {
		return []Monitor{}, fmt.Errorf("ListMonitors: %w", err)
	}

	return st.monitors, nil
}
