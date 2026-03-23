//go:build !windows

package cmd

import (
	"os"
	"syscall"
)

func daemonSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{Setsid: true}
}

func processExistsForPlatform(p *os.Process) bool {
	// Signal 0 checks if process exists without killing it.
	return p.Signal(syscall.Signal(0)) == nil
}

func terminateProcess(p *os.Process) error {
	return p.Signal(syscall.SIGTERM)
}
