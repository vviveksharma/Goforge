package cmd

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/viveksharma/goforge/internal/mocks"
)

// TestValidateProjectName tests project name validation logic
func TestValidateProjectName(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
		errorMsg  string
	}{
		// Valid names
		{
			name:      "valid simple name",
			input:     "myproject",
			wantError: false,
		},
		{
			name:      "valid name with dashes",
			input:     "my-project",
			wantError: false,
		},
		{
			name:      "valid name with underscores",
			input:     "my_project",
			wantError: false,
		},
		{
			name:      "valid name with numbers",
			input:     "project123",
			wantError: false,
		},
		{
			name:      "valid name starting with number",
			input:     "123project",
			wantError: false,
		},
		{
			name:      "valid mixed case",
			input:     "MyProject",
			wantError: false,
		},
		{
			name:      "valid complex name",
			input:     "My-Awesome_Project-123",
			wantError: false,
		},
		{
			name:      "valid single character",
			input:     "a",
			wantError: false,
		},

		// Invalid names - empty
		{
			name:      "empty name",
			input:     "",
			wantError: true,
			errorMsg:  "cannot be empty",
		},

		// Invalid names - path traversal
		{
			name:      "path traversal with ..",
			input:     "../etc",
			wantError: true,
			errorMsg:  "path separators",
		},
		{
			name:      "path traversal in middle",
			input:     "my../project",
			wantError: true,
			errorMsg:  "path separators",
		},
		{
			name:      "unix path separator",
			input:     "my/project",
			wantError: true,
			errorMsg:  "path separators",
		},
		{
			name:      "windows path separator",
			input:     "my\\project",
			wantError: true,
			errorMsg:  "path separators",
		},
		{
			name:      "absolute path unix",
			input:     "/tmp/project",
			wantError: true,
			errorMsg:  "path separators",
		},
		{
			name:      "absolute path windows",
			input:     "C:\\project",
			wantError: true,
			errorMsg:  "path separators",
		},

		// Invalid names - invalid characters
		{
			name:      "name with spaces",
			input:     "my project",
			wantError: true,
			errorMsg:  "can only contain letters",
		},
		{
			name:      "name with special char @",
			input:     "my@project",
			wantError: true,
			errorMsg:  "can only contain letters",
		},
		{
			name:      "name with special char #",
			input:     "my#project",
			wantError: true,
			errorMsg:  "can only contain letters",
		},
		{
			name:      "name with special char $",
			input:     "my$project",
			wantError: true,
			errorMsg:  "can only contain letters",
		},
		{
			name:      "name with special char %",
			input:     "my%project",
			wantError: true,
			errorMsg:  "can only contain letters",
		},
		{
			name:      "name with special char &",
			input:     "my&project",
			wantError: true,
			errorMsg:  "can only contain letters",
		},
		{
			name:      "name with asterisk",
			input:     "my*project",
			wantError: true,
			errorMsg:  "can only contain letters",
		},
		{
			name:      "name with parentheses",
			input:     "my(project)",
			wantError: true,
			errorMsg:  "can only contain letters",
		},
		{
			name:      "name with brackets",
			input:     "my[project]",
			wantError: true,
			errorMsg:  "can only contain letters",
		},
		{
			name:      "name with dot",
			input:     "my.project",
			wantError: true,
			errorMsg:  "can only contain letters",
		},
		{
			name:      "name with comma",
			input:     "my,project",
			wantError: true,
			errorMsg:  "can only contain letters",
		},

		// Invalid names - reserved names
		{
			name:      "reserved name: .",
			input:     ".",
			wantError: true,
			errorMsg:  "can only contain letters", // Caught by character check first
		},
		{
			name:      "reserved name: ..",
			input:     "..",
			wantError: true,
			errorMsg:  "path separators", // Caught by path check first
		},
		{
			name:      "reserved name: con (Windows)",
			input:     "con",
			wantError: true,
			errorMsg:  "reserved name",
		},
		{
			name:      "reserved name: CON (uppercase)",
			input:     "CON",
			wantError: true,
			errorMsg:  "reserved name",
		},
		{
			name:      "reserved name: prn",
			input:     "prn",
			wantError: true,
			errorMsg:  "reserved name",
		},
		{
			name:      "reserved name: aux",
			input:     "aux",
			wantError: true,
			errorMsg:  "reserved name",
		},
		{
			name:      "reserved name: nul",
			input:     "nul",
			wantError: true,
			errorMsg:  "reserved name",
		},
		{
			name:      "reserved name: com1",
			input:     "com1",
			wantError: true,
			errorMsg:  "reserved name",
		},
		{
			name:      "reserved name: lpt1",
			input:     "lpt1",
			wantError: true,
			errorMsg:  "reserved name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateProjectName(tt.input)

			if tt.wantError {
				assert.Error(t, err, "Expected error for input: %s", tt.input)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg,
						"Error message should contain: %s", tt.errorMsg)
				}
			} else {
				assert.NoError(t, err, "Expected no error for input: %s", tt.input)
			}
		})
	}
}

// TestValidateServerType tests server type validation logic
func TestValidateServerType(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
		errorMsg  string
	}{
		// Valid server types
		{
			name:      "valid fiber lowercase",
			input:     "fiber",
			wantError: false,
		},
		{
			name:      "valid gin lowercase",
			input:     "gin",
			wantError: false,
		},

		// Invalid server types
		{
			name:      "empty server type",
			input:     "",
			wantError: true,
			errorMsg:  "only 'fiber' and 'gin' are supported",
		},
		{
			name:      "invalid echo",
			input:     "echo",
			wantError: true,
			errorMsg:  "only 'fiber' and 'gin' are supported",
		},
		{
			name:      "invalid express",
			input:     "express",
			wantError: true,
			errorMsg:  "only 'fiber' and 'gin' are supported",
		},
		{
			name:      "invalid fastapi",
			input:     "fastapi",
			wantError: true,
			errorMsg:  "only 'fiber' and 'gin' are supported",
		},
		{
			name:      "invalid django",
			input:     "django",
			wantError: true,
			errorMsg:  "only 'fiber' and 'gin' are supported",
		},
		{
			name:      "invalid random string",
			input:     "notaframework",
			wantError: true,
			errorMsg:  "only 'fiber' and 'gin' are supported",
		},
		{
			name:      "uppercase FIBER should fail (validation expects lowercase input)",
			input:     "FIBER",
			wantError: true,
			errorMsg:  "only 'fiber' and 'gin' are supported",
		},
		{
			name:      "uppercase GIN should fail (validation expects lowercase input)",
			input:     "GIN",
			wantError: true,
			errorMsg:  "only 'fiber' and 'gin' are supported",
		},
		{
			name:      "mixed case Fiber",
			input:     "Fiber",
			wantError: true,
			errorMsg:  "only 'fiber' and 'gin' are supported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateServerType(tt.input)

			if tt.wantError {
				assert.Error(t, err, "Expected error for input: %s", tt.input)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg,
						"Error message should contain: %s", tt.errorMsg)
				}
			} else {
				assert.NoError(t, err, "Expected no error for input: %s", tt.input)
			}
		})
	}
}

// TestRunCreateWithDeps_Success tests successful project creation with mocks
func TestRunCreateWithDeps_Success(t *testing.T) {
	mockFS := &mocks.MockFileSystem{
		GetwdFunc: func() (string, error) {
			return "/home/user", nil
		},
		StatFunc: func(name string) (os.FileInfo, error) {
			return nil, os.ErrNotExist // Project doesn't exist
		},
		MkdirAllFunc: func(path string, perm os.FileMode) error {
			return nil
		},
		CreateFunc: func(name string) (io.WriteCloser, error) {
			return &mocks.MockWriteCloser{}, nil
		},
		ChmodFunc: func(name string, mode os.FileMode) error {
			return nil
		},
	}

	mockCMD := &mocks.MockCommander{
		RunFunc: func(name string, args []string, dir string, stdout, stderr io.Writer) error {
			return nil // Simulate successful go mod tidy
		},
	}

	mockWriter := &mocks.MockWriter{}

	opts := CreateOptions{
		FS:     mockFS,
		CMD:    mockCMD,
		Writer: mockWriter,
	}

	// Set server type flag
	serverType = "fiber"

	cmd := &cobra.Command{}
	err := runCreateWithDeps(opts, cmd, []string{"test-project"})

	require.NoError(t, err)

	// Verify file system calls
	assert.Equal(t, 1, mockFS.GetwdCalls, "Should call Getwd once")
	assert.Equal(t, 1, len(mockFS.StatCalls), "Should check if directory exists")
	assert.Greater(t, len(mockFS.MkdirAllCalls), 0, "Should create directories")

	// Verify command execution
	assert.Equal(t, 1, len(mockCMD.RunCalls), "Should run go mod tidy")
	assert.Equal(t, "go", mockCMD.RunCalls[0].Name)
	assert.Equal(t, []string{"mod", "tidy"}, mockCMD.RunCalls[0].Args)

	// Verify output
	assert.Greater(t, len(mockWriter.PrintfCalls), 0, "Should have output")
	assert.Greater(t, len(mockWriter.PrintlnCalls), 0, "Should have output")
}

// TestRunCreateWithDeps_DirectoryExists tests error when directory already exists
func TestRunCreateWithDeps_DirectoryExists(t *testing.T) {
	mockFS := &mocks.MockFileSystem{
		GetwdFunc: func() (string, error) {
			return "/home/user", nil
		},
		StatFunc: func(name string) (os.FileInfo, error) {
			// Simulate directory exists
			return &mocks.MockFileInfo{}, nil
		},
	}

	mockCMD := &mocks.MockCommander{}
	mockWriter := &mocks.MockWriter{}

	opts := CreateOptions{
		FS:     mockFS,
		CMD:    mockCMD,
		Writer: mockWriter,
	}

	serverType = "fiber"
	cmd := &cobra.Command{}
	err := runCreateWithDeps(opts, cmd, []string{"existing-project"})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")
}

// TestRunCreateWithDeps_InvalidProjectName tests validation errors
func TestRunCreateWithDeps_InvalidProjectName(t *testing.T) {
	mockFS := &mocks.MockFileSystem{}
	mockCMD := &mocks.MockCommander{}
	mockWriter := &mocks.MockWriter{}

	opts := CreateOptions{
		FS:     mockFS,
		CMD:    mockCMD,
		Writer: mockWriter,
	}

	serverType = "fiber"
	cmd := &cobra.Command{}

	tests := []struct {
		name        string
		projectName string
		errorMsg    string
	}{
		{
			name:        "path traversal",
			projectName: "../evil",
			errorMsg:    "invalid project name",
		},
		{
			name:        "invalid characters",
			projectName: "my@project",
			errorMsg:    "invalid project name",
		},
		{
			name:        "reserved name",
			projectName: "con",
			errorMsg:    "invalid project name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := runCreateWithDeps(opts, cmd, []string{tt.projectName})
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errorMsg)
		})
	}
}

// TestRunCreateWithDeps_InvalidServerType tests server type validation
func TestRunCreateWithDeps_InvalidServerType(t *testing.T) {
	mockFS := &mocks.MockFileSystem{
		GetwdFunc: func() (string, error) {
			return "/home/user", nil
		},
	}
	mockCMD := &mocks.MockCommander{}
	mockWriter := &mocks.MockWriter{}

	opts := CreateOptions{
		FS:     mockFS,
		CMD:    mockCMD,
		Writer: mockWriter,
	}

	serverType = "invalid"
	cmd := &cobra.Command{}
	err := runCreateWithDeps(opts, cmd, []string{"test-project"})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "only 'fiber' and 'gin' are supported")
}

// TestRunCreateWithDeps_GetwdError tests error handling for Getwd failure
func TestRunCreateWithDeps_GetwdError(t *testing.T) {
	mockFS := &mocks.MockFileSystem{
		GetwdFunc: func() (string, error) {
			return "", errors.New("permission denied")
		},
	}
	mockCMD := &mocks.MockCommander{}
	mockWriter := &mocks.MockWriter{}

	opts := CreateOptions{
		FS:     mockFS,
		CMD:    mockCMD,
		Writer: mockWriter,
	}

	serverType = "fiber"
	cmd := &cobra.Command{}
	err := runCreateWithDeps(opts, cmd, []string{"test-project"})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get current directory")
}

// TestRunCreateWithDeps_MkdirAllError tests error handling for directory creation failure
func TestRunCreateWithDeps_MkdirAllError(t *testing.T) {
	mockFS := &mocks.MockFileSystem{
		GetwdFunc: func() (string, error) {
			return "/home/user", nil
		},
		StatFunc: func(name string) (os.FileInfo, error) {
			return nil, os.ErrNotExist
		},
		MkdirAllFunc: func(path string, perm os.FileMode) error {
			return errors.New("permission denied")
		},
	}
	mockCMD := &mocks.MockCommander{}
	mockWriter := &mocks.MockWriter{}

	opts := CreateOptions{
		FS:     mockFS,
		CMD:    mockCMD,
		Writer: mockWriter,
	}

	serverType = "fiber"
	cmd := &cobra.Command{}
	err := runCreateWithDeps(opts, cmd, []string{"test-project"})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create project directory")
}

// TestRunCreateWithDeps_GeneratorError tests cleanup on generation failure
func TestRunCreateWithDeps_GeneratorError(t *testing.T) {
	mockFS := &mocks.MockFileSystem{
		GetwdFunc: func() (string, error) {
			return "/home/user", nil
		},
		StatFunc: func(name string) (os.FileInfo, error) {
			return nil, os.ErrNotExist
		},
		MkdirAllFunc: func(path string, perm os.FileMode) error {
			// Fail after creating project directory to test cleanup
			if path == "/home/user/test-project" {
				return nil
			}
			return errors.New("template error")
		},
		RemoveAllFunc: func(path string) error {
			return nil // Cleanup succeeds
		},
	}
	mockCMD := &mocks.MockCommander{}
	mockWriter := &mocks.MockWriter{}

	opts := CreateOptions{
		FS:     mockFS,
		CMD:    mockCMD,
		Writer: mockWriter,
	}

	serverType = "fiber"
	cmd := &cobra.Command{}
	err := runCreateWithDeps(opts, cmd, []string{"test-project"})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to generate project")

	// Verify cleanup was attempted
	assert.Greater(t, len(mockFS.RemoveAllCalls), 0, "Should attempt cleanup on failure")
}

// TestRunCreateWithDeps_GoModTidyError tests handling of go mod tidy failure
func TestRunCreateWithDeps_GoModTidyError(t *testing.T) {
	mockFS := &mocks.MockFileSystem{
		GetwdFunc: func() (string, error) {
			return "/home/user", nil
		},
		StatFunc: func(name string) (os.FileInfo, error) {
			return nil, os.ErrNotExist
		},
		MkdirAllFunc: func(path string, perm os.FileMode) error {
			return nil
		},
		CreateFunc: func(name string) (io.WriteCloser, error) {
			return &mocks.MockWriteCloser{}, nil
		},
		ChmodFunc: func(name string, mode os.FileMode) error {
			return nil
		},
	}

	mockCMD := &mocks.MockCommander{
		RunFunc: func(name string, args []string, dir string, stdout, stderr io.Writer) error {
			return errors.New("go command not found")
		},
	}

	mockWriter := &mocks.MockWriter{}

	opts := CreateOptions{
		FS:     mockFS,
		CMD:    mockCMD,
		Writer: mockWriter,
	}

	serverType = "fiber"
	cmd := &cobra.Command{}
	err := runCreateWithDeps(opts, cmd, []string{"test-project"})

	// Should NOT fail - go mod tidy error is warned but not fatal
	require.NoError(t, err)

	// Verify warning was printed
	foundWarning := false
	for _, call := range mockWriter.PrintfCalls {
		if contains(call.Format, "Warning") || contains(call.Format, "⚠️") {
			foundWarning = true
			break
		}
	}
	assert.True(t, foundWarning, "Should print warning about go mod tidy failure")
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr || len(s) > len(substr) &&
			(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
				indexOf(s, substr) != -1))
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
