package adapters

import (
	"io"
	"os"
)

// OSFileSystem is the real implementation of FileSystem interface
// that performs actual file system operations.
type OSFileSystem struct{}

// NewOSFileSystem creates a new OSFileSystem instance.
func NewOSFileSystem() *OSFileSystem {
	return &OSFileSystem{}
}

// MkdirAll creates a directory named path, along with any necessary parents.
func (fs *OSFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// Create creates or truncates the named file.
func (fs *OSFileSystem) Create(name string) (io.WriteCloser, error) {
	return os.Create(name)
}

// Stat returns a FileInfo describing the named file.
func (fs *OSFileSystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

// RemoveAll removes path and any children it contains.
func (fs *OSFileSystem) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

// Getwd returns a rooted path name corresponding to the current directory.
func (fs *OSFileSystem) Getwd() (string, error) {
	return os.Getwd()
}

// Chmod changes the mode of the named file to mode.
func (fs *OSFileSystem) Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}
