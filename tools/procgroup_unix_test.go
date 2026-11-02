//go:build !windows

package tools

import (
	"context"
	"testing"

	"memdoor/pkg/shell"
)

// The bash tool's own process group must keep the command's new session:
// Setpgid alone put an interactive zsh back on the gateway's terminal.
func TestTheBashToolKeepsTheCommandInASessionOfItsOwn(t *testing.T) {
	cmd := shell.Command(context.Background(), "true")
	ownProcessGroup(cmd)
	if !cmd.SysProcAttr.Setsid || cmd.SysProcAttr.Setpgid {
		t.Fatalf("want Setsid and no Setpgid, got %+v", cmd.SysProcAttr)
	}
}
