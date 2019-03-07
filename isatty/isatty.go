// Copyright 2017 The Fuchsia Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux darwin

package isatty

import (
	"os"
	"syscall"
	"unsafe"
)

func IsTerminal() bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, os.Stdout.Fd(), ioctlTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}
