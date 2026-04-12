package adapters

import (
	"io"
	"os/exec"
)

// ExecCommander is the real implementation of Commander interface
// that executes actual system commands.
type ExecCommander struct{}

// NewExecCommander creates a new ExecCommander instance.
func NewExecCommander() *ExecCommander {
	return &ExecCommander{}
}

// Run executes a command with the given name and arguments.
func (e *ExecCommander) Run(name string, args []string, dir string, stdout, stderr io.Writer) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return cmd.Run()
}
