package win32

import (
	"unsafe"
)

const (
	DIB_RGB_COLORS     = 0x00
	SM_XVIRTUALSCREEN  = 76
	SM_YVIRTUALSCREEN  = 77
	SM_CXVIRTUALSCREEN = 78
	SM_CYVIRTUALSCREEN = 79
	ROP                = 0x00CC0020
)

const (
	BiPlanes      = 1
	BiBitCount    = 32
	BiCompression = 0
)

type HWND uintptr
type HDC uintptr
type HGDIOBJ uintptr
type HBITMAP uintptr
type DWORD uint32
type HMONITOR uintptr

type Rect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type MonitorInfoEx struct {
	CbSize    uint32
	RcMonitor Rect
	RcWork    Rect
	DwFlags   uint32
	SzDevice  [32]uint16
}

type bitmapInfoHeader struct {
	BiSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

func NewBitmapInfoHeader(biWidth, biHeight int) bitmapInfoHeader {
	return bitmapInfoHeader{
		BiSize:        uint32(unsafe.Sizeof(bitmapInfoHeader{})),
		BiWidth:       int32(biWidth),
		BiHeight:      -int32(biHeight),
		BiPlanes:      BiPlanes,
		BiBitCount:    BiBitCount,
		BiCompression: BiCompression,
		BiSizeImage:   uint32(biWidth * biHeight * 4),
	}
}
