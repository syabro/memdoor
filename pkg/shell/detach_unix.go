//go:build !windows

package shell

import (
	"os/exec"
	"syscall"
)

// detach starts the command in a session of its own, with no controlling
// terminal. A gateway started from a window holds that window's terminal; an
// interactive zsh (-i) run under it tried to take the terminal, got SIGTTOU
// and stayed stopped for good. A new session is also a new process group, so
// a kill of the group reaches what the line started.
func detach(cmd *exec.Cmd) { cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true} }
