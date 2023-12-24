// just for mac + linux

//go:build darwin || linux
// +build darwin linux

package glog

import (
	"os"
	"syscall"
)

func redirectStderr(f *os.File) error {
	return syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
}
