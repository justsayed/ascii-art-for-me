package main

import (
	"os"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row, Col, Xpixel, Ypixel uint16
}

func getTerminalWidth() int {
	ws := &winsize{}
	ret, _, _ := syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdout.Fd(),
		syscall.TIOCGWINSZ,
		uintptr(unsafe.Pointer(ws)),
	)
	if ret == 0 {
		return int(ws.Col)
	}
	return 80
}
