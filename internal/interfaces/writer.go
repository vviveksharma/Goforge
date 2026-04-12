package interfaces

// Writer provides an abstraction over output operations.
// This interface enables dependency injection for testing.
type Writer interface {
	// Printf formats according to a format specifier and writes to standard output.
	Printf(format string, args ...interface{})

	// Println formats using the default formats and writes to standard output.
	Println(args ...interface{})
}
