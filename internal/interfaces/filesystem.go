package interfaces

import (
	"io"
	"os"
)

// FileSystem provides an abstraction over file system operations.
// This interface enables dependency injection for testing.
type FileSystem interface {
	// MkdirAll creates a directory named path, along with any necessary parents.
	MkdirAll(path string, perm os.FileMode) error

	// Create creates or truncates the named file.
	Create(name string) (io.WriteCloser, error)

	// Stat returns a FileInfo describing the named file.
	Stat(name string) (os.FileInfo, error)

	// RemoveAll removes path and any children it contains.
	RemoveAll(path string) error

	// Getwd returns a rooted path name corresponding to the current directory.
	Getwd() (string, error)

	// Chmod changes the mode of the named file to mode.
	Chmod(name string, mode os.FileMode) error
}
