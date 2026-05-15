//go:build linux || darwin

package server

import (
	"os/exec"
	"syscall"
)

func setProcGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

func terminateProcessGroup(pid int) error {
	return syscall.Kill(-pid, syscall.SIGTERM)
}

func killProcessGroup(pid int) error {
	return syscall.Kill(-pid, syscall.SIGKILL)
}
