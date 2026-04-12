package mocks

import (
	"io"
	"os"
	"sync"
	"time"
)

// MockFileSystem is a mock implementation of interfaces.FileSystem for testing.
type MockFileSystem struct {
	mu             sync.Mutex
	MkdirAllFunc   func(path string, perm os.FileMode) error
	CreateFunc     func(name string) (io.WriteCloser, error)
	StatFunc       func(name string) (os.FileInfo, error)
	RemoveAllFunc  func(path string) error
	GetwdFunc      func() (string, error)
	ChmodFunc      func(name string, mode os.FileMode) error
	MkdirAllCalls  []MkdirAllCall
	CreateCalls    []CreateCall
	StatCalls      []StatCall
	RemoveAllCalls []RemoveAllCall
	GetwdCalls     int
	ChmodCalls     []ChmodCall
}

type MkdirAllCall struct {
	Path string
	Perm os.FileMode
}

type CreateCall struct {
	Name string
}

type StatCall struct {
	Name string
}

type RemoveAllCall struct {
	Path string
}

type ChmodCall struct {
	Name string
	Mode os.FileMode
}

func (m *MockFileSystem) MkdirAll(path string, perm os.FileMode) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.MkdirAllCalls = append(m.MkdirAllCalls, MkdirAllCall{Path: path, Perm: perm})
	if m.MkdirAllFunc != nil {
		return m.MkdirAllFunc(path, perm)
	}
	return nil
}

func (m *MockFileSystem) Create(name string) (io.WriteCloser, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.CreateCalls = append(m.CreateCalls, CreateCall{Name: name})
	if m.CreateFunc != nil {
		return m.CreateFunc(name)
	}
	return &MockWriteCloser{}, nil
}

func (m *MockFileSystem) Stat(name string) (os.FileInfo, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.StatCalls = append(m.StatCalls, StatCall{Name: name})
	if m.StatFunc != nil {
		return m.StatFunc(name)
	}
	return nil, os.ErrNotExist
}

func (m *MockFileSystem) RemoveAll(path string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.RemoveAllCalls = append(m.RemoveAllCalls, RemoveAllCall{Path: path})
	if m.RemoveAllFunc != nil {
		return m.RemoveAllFunc(path)
	}
	return nil
}

func (m *MockFileSystem) Getwd() (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.GetwdCalls++
	if m.GetwdFunc != nil {
		return m.GetwdFunc()
	}
	return "/mock/current/dir", nil
}

func (m *MockFileSystem) Chmod(name string, mode os.FileMode) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ChmodCalls = append(m.ChmodCalls, ChmodCall{Name: name, Mode: mode})
	if m.ChmodFunc != nil {
		return m.ChmodFunc(name, mode)
	}
	return nil
}

type MockWriteCloser struct {
	WriteFunc func(p []byte) (n int, err error)
	CloseFunc func() error
	Data      []byte
}

func (m *MockWriteCloser) Write(p []byte) (n int, err error) {
	if m.WriteFunc != nil {
		return m.WriteFunc(p)
	}
	m.Data = append(m.Data, p...)
	return len(p), nil
}

func (m *MockWriteCloser) Close() error {
	if m.CloseFunc != nil {
		return m.CloseFunc()
	}
	return nil
}

type MockCommander struct {
	mu       sync.Mutex
	RunFunc  func(name string, args []string, dir string, stdout, stderr io.Writer) error
	RunCalls []CommandCall
}

type CommandCall struct {
	Name   string
	Args   []string
	Dir    string
	Stdout io.Writer
	Stderr io.Writer
}

func (m *MockCommander) Run(name string, args []string, dir string, stdout, stderr io.Writer) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.RunCalls = append(m.RunCalls, CommandCall{
		Name:   name,
		Args:   args,
		Dir:    dir,
		Stdout: stdout,
		Stderr: stderr,
	})
	if m.RunFunc != nil {
		return m.RunFunc(name, args, dir, stdout, stderr)
	}
	return nil
}

type MockWriter struct {
	mu           sync.Mutex
	PrintfFunc   func(format string, args ...interface{})
	PrintlnFunc  func(args ...interface{})
	PrintfCalls  []PrintfCall
	PrintlnCalls []PrintlnCall
	Output       []string
}

type PrintfCall struct {
	Format string
	Args   []interface{}
}

type PrintlnCall struct {
	Args []interface{}
}

func (m *MockWriter) Printf(format string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.PrintfCalls = append(m.PrintfCalls, PrintfCall{Format: format, Args: args})
	if m.PrintfFunc != nil {
		m.PrintfFunc(format, args...)
	}
}

func (m *MockWriter) Println(args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.PrintlnCalls = append(m.PrintlnCalls, PrintlnCall{Args: args})
	if m.PrintlnFunc != nil {
		m.PrintlnFunc(args...)
	}
}

// MockFileInfo is a mock implementation of os.FileInfo for testing.
type MockFileInfo struct {
	NameValue    string
	SizeValue    int64
	ModeValue    os.FileMode
	ModTimeValue time.Time
	IsDirValue   bool
	SysValue     interface{}
}

func (m *MockFileInfo) Name() string {
	if m.NameValue != "" {
		return m.NameValue
	}
	return "mockfile"
}

func (m *MockFileInfo) Size() int64 {
	return m.SizeValue
}

func (m *MockFileInfo) Mode() os.FileMode {
	if m.ModeValue != 0 {
		return m.ModeValue
	}
	return 0644
}

func (m *MockFileInfo) ModTime() time.Time {
	if !m.ModTimeValue.IsZero() {
		return m.ModTimeValue
	}
	return time.Now()
}

func (m *MockFileInfo) IsDir() bool {
	return m.IsDirValue
}

func (m *MockFileInfo) Sys() interface{} {
	return m.SysValue
}
