//go:build !windows

package shell

import (
	"context"
	"testing"
)

// A command has no controlling terminal: under a gateway started from a
// window, an interactive zsh otherwise took the terminal and stayed stopped.
func TestACommandRunsInASessionOfItsOwn(t *testing.T) {
	cmd := Command(context.Background(), "true")
	if cmd.SysProcAttr == nil || !cmd.SysProcAttr.Setsid {
		t.Fatalf("no new session: %+v", cmd.SysProcAttr)
	}
}
