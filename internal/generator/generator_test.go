package generator

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/viveksharma/goforge/internal/mocks"
)

// TestNewGenerator tests the basic generator constructor
func TestNewGenerator(t *testing.T) {
	config := ProjectConfig{
		ProjectName: "test-api",
		ProjectPath: "/tmp/test-api",
		ModulePath:  "github.com/test/test-api",
		ServerType:  "fiber",
	}

	gen := NewGenerator(config)

	require.NotNil(t, gen)
	assert.Equal(t, "test-api", gen.config.ProjectName)
	assert.Equal(t, "/tmp/test-api", gen.config.ProjectPath)
	assert.Equal(t, "github.com/test/test-api", gen.config.ModulePath)
	assert.Equal(t, "fiber", gen.config.ServerType)
	assert.NotNil(t, gen.fs)
}

// TestNewGeneratorWithFS tests generator with custom filesystem
func TestNewGeneratorWithFS(t *testing.T) {
	config := ProjectConfig{
		ProjectName: "test-api",
		ProjectPath: "/tmp/test-api",
		ModulePath:  "github.com/test/test-api",
		ServerType:  "fiber",
	}

	mockFS := &mocks.MockFileSystem{}
	gen := NewGeneratorWithFS(config, mockFS)

	require.NotNil(t, gen)
	assert.Equal(t, mockFS, gen.fs)
}

// TestGetServerSpecificTemplates_Fiber tests Fiber template selection
func TestGetServerSpecificTemplates_Fiber(t *testing.T) {
	config := ProjectConfig{
		ServerType: "fiber",
	}

	gen := NewGeneratorWithFS(config, &mocks.MockFileSystem{})
	templates := gen.getServerSpecificTemplates()

	require.NotNil(t, templates)

	// Check for Fiber-specific templates
	assert.Contains(t, templates, "templates/internal/handler/health.go.tmpl")
	assert.Contains(t, templates, "templates/internal/handler/health_test.go.tmpl")
	assert.Contains(t, templates, "templates/internal/server/server.go.tmpl")
	assert.Contains(t, templates, "templates/internal/middleware/security.go.tmpl")
	assert.Contains(t, templates, "templates/internal/middleware/logger.go.tmpl")
	assert.Contains(t, templates, "templates/internal/middleware/recovery.go.tmpl")
	assert.Contains(t, templates, "templates/internal/middleware/metrics.go.tmpl")

	// Should map to correct output paths
	assert.Equal(t, "internal/handler/health.go", templates["templates/internal/handler/health.go.tmpl"])
	assert.Equal(t, "internal/server/server.go", templates["templates/internal/server/server.go.tmpl"])
}

// TestGetServerSpecificTemplates_Gin tests Gin template selection
func TestGetServerSpecificTemplates_Gin(t *testing.T) {
	config := ProjectConfig{
		ServerType: "gin",
	}

	gen := NewGeneratorWithFS(config, &mocks.MockFileSystem{})
	templates := gen.getServerSpecificTemplates()

	require.NotNil(t, templates)

	// Check for Gin-specific templates
	assert.Contains(t, templates, "templates/internal/handler/health_gin.go.tmpl")
	assert.Contains(t, templates, "templates/internal/handler/health_gin_test.go.tmpl")
	assert.Contains(t, templates, "templates/internal/server/server_gin.go.tmpl")
	assert.Contains(t, templates, "templates/internal/middleware/security_gin.go.tmpl")
	assert.Contains(t, templates, "templates/internal/middleware/logger_gin.go.tmpl")
	assert.Contains(t, templates, "templates/internal/middleware/recovery_gin.go.tmpl")
	assert.Contains(t, templates, "templates/internal/middleware/metrics_gin.go.tmpl")

	// Should map to correct output paths
	assert.Equal(t, "internal/handler/health.go", templates["templates/internal/handler/health_gin.go.tmpl"])
	assert.Equal(t, "internal/server/server.go", templates["templates/internal/server/server_gin.go.tmpl"])
}

// TestCreateDirectories tests directory creation
func TestCreateDirectories(t *testing.T) {
	mockFS := &mocks.MockFileSystem{
		MkdirAllFunc: func(path string, perm os.FileMode) error {
			return nil
		},
	}

	config := ProjectConfig{
		ProjectPath: "/tmp/test-project",
	}

	gen := NewGeneratorWithFS(config, mockFS)
	err := gen.createDirectories()

	require.NoError(t, err)

	// Verify all required directories were created
	expectedDirs := []string{
		"cmd/api",
		"internal/config",
		"internal/handler",
		"internal/middleware",
		"internal/server",
		"pkg/logger",
		"pkg/database",
		"pkg/cache",
		"pkg/migration",
		"migrations",
		"docs",
		"deployments",
		"scripts",
	}

	assert.Equal(t, len(expectedDirs), len(mockFS.MkdirAllCalls), "Should create all required directories")

	// Verify each directory was created with correct permissions
	for _, call := range mockFS.MkdirAllCalls {
		assert.Equal(t, os.FileMode(0755), call.Perm, "Directories should have 0755 permissions")

		// Check that path contains project path
		assert.Contains(t, call.Path, "/tmp/test-project")
	}
}

// TestCreateDirectories_Error tests error handling in directory creation
func TestCreateDirectories_Error(t *testing.T) {
	callCount := 0
	mockFS := &mocks.MockFileSystem{
		MkdirAllFunc: func(path string, perm os.FileMode) error {
			callCount++
			// Fail on second directory
			if callCount >= 2 {
				return errors.New("permission denied")
			}
			return nil
		},
	}

	config := ProjectConfig{
		ProjectPath: "/tmp/test-project",
	}

	gen := NewGeneratorWithFS(config, mockFS)
	err := gen.createDirectories()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create directory")
	assert.Contains(t, err.Error(), "permission denied")
}

// TestGenerate_Success tests full generation flow
func TestGenerate_Success(t *testing.T) {
	callCount := 0
	mockFS := &mocks.MockFileSystem{
		MkdirAllFunc: func(path string, perm os.FileMode) error {
			return nil
		},
		CreateFunc: func(name string) (io.WriteCloser, error) {
			return &mocks.MockWriteCloser{}, nil
		},
		ChmodFunc: func(name string, mode os.FileMode) error {
			callCount++
			return nil
		},
	}

	config := ProjectConfig{
		ProjectName: "test-api",
		ProjectPath: "/tmp/test-api",
		ModulePath:  "github.com/test/test-api",
		ServerType:  "fiber",
	}

	gen := NewGeneratorWithFS(config, mockFS)
	err := gen.Generate()

	require.NoError(t, err)

	// Verify directories were created
	assert.Greater(t, len(mockFS.MkdirAllCalls), 0, "Should create directories")

	// Verify files were created
	assert.Greater(t, len(mockFS.CreateCalls), 0, "Should create files")

	// Verify scripts were made executable
	assert.Greater(t, callCount, 0, "Should make scripts executable")
}

// TestGenerate_DirectoryCreationError tests error in directory creation
func TestGenerate_DirectoryCreationError(t *testing.T) {
	mockFS := &mocks.MockFileSystem{
		MkdirAllFunc: func(path string, perm os.FileMode) error {
			return errors.New("disk full")
		},
	}

	config := ProjectConfig{
		ProjectPath: "/tmp/test-api",
	}

	gen := NewGeneratorWithFS(config, mockFS)
	err := gen.Generate()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk full")
}

// TestGenerate_FileCreationError tests error in file creation
func TestGenerate_FileCreationError(t *testing.T) {
	mockFS := &mocks.MockFileSystem{
		MkdirAllFunc: func(path string, perm os.FileMode) error {
			return nil
		},
		CreateFunc: func(name string) (io.WriteCloser, error) {
			// Fail file creation
			return nil, errors.New("permission denied")
		},
	}

	config := ProjectConfig{
		ProjectName: "test-api",
		ProjectPath: "/tmp/test-api",
		ModulePath:  "github.com/test/test-api",
		ServerType:  "fiber",
	}

	gen := NewGeneratorWithFS(config, mockFS)
	err := gen.Generate()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create file")
}

// TestGenerateFile_ChmodError tests error in chmod
func TestGenerateFile_ChmodError(t *testing.T) {
	mockFS := &mocks.MockFileSystem{
		MkdirAllFunc: func(path string, perm os.FileMode) error {
			return nil
		},
		CreateFunc: func(name string) (io.WriteCloser, error) {
			return &mocks.MockWriteCloser{}, nil
		},
		ChmodFunc: func(name string, mode os.FileMode) error {
			return errors.New("chmod failed")
		},
	}

	config := ProjectConfig{
		ProjectName: "test-api",
		ProjectPath: "/tmp/test-api",
		ModulePath:  "github.com/test/test-api",
		ServerType:  "fiber",
	}

	gen := NewGeneratorWithFS(config, mockFS)
	err := gen.Generate()

	// Should fail on chmod error for script files
	require.Error(t, err)
	assert.Contains(t, err.Error(), "chmod")
}

// TestProjectConfig tests ProjectConfig struct
func TestProjectConfig(t *testing.T) {
	config := ProjectConfig{
		ProjectName: "my-api",
		ProjectPath: "/home/user/my-api",
		ModulePath:  "github.com/user/my-api",
		ServerType:  "gin",
	}

	assert.Equal(t, "my-api", config.ProjectName)
	assert.Equal(t, "/home/user/my-api", config.ProjectPath)
	assert.Equal(t, "github.com/user/my-api", config.ModulePath)
	assert.Equal(t, "gin", config.ServerType)
}
