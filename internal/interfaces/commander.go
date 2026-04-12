package interfaces

import "io"

// Commander provides an abstraction over command execution.
// This interface enables dependency injection for testing.
type Commander interface {
	// Run executes a command with the given name and arguments.
	// The command is executed in the specified directory.
	// Output is written to stdout and stderr.
	Run(name string, args []string, dir string, stdout, stderr io.Writer) error
}
