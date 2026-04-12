package adapters

import "fmt"

// StdoutWriter is the real implementation of Writer interface
// that writes to standard output.
type StdoutWriter struct{}

// NewStdoutWriter creates a new StdoutWriter instance.
func NewStdoutWriter() *StdoutWriter {
	return &StdoutWriter{}
}

// Printf formats according to a format specifier and writes to standard output.
func (w *StdoutWriter) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Println formats using the default formats and writes to standard output.
func (w *StdoutWriter) Println(args ...interface{}) {
	fmt.Println(args...)
}
