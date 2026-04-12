package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/viveksharma/goforge/internal/adapters"
	"github.com/viveksharma/goforge/internal/generator"
	"github.com/viveksharma/goforge/internal/interfaces"
)

var (
	serverType string
)

// CreateOptions contains dependencies for the create command.
// This enables dependency injection for testing.
type CreateOptions struct {
	FS     interfaces.FileSystem
	CMD    interfaces.Commander
	Writer interfaces.Writer
}

var createCmd = &cobra.Command{
	Use:   "create [project-name]",
	Short: "Create a new Go project",
	Long:  `Create a new Go project with Fiber or Gin web framework, PostgreSQL, Redis, and production-ready features.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runCreate,
}

func init() {
	createCmd.Flags().StringVarP(&serverType, "server", "s", "", "Web framework to use (fiber or gin) [REQUIRED]")
	createCmd.MarkFlagRequired("server")
	rootCmd.AddCommand(createCmd)
}

// runCreate is the main entry point that uses real dependencies.
// This is a convenience wrapper for backward compatibility.
func runCreate(cmd *cobra.Command, args []string) error {
	return runCreateWithDeps(CreateOptions{
		FS:     adapters.NewOSFileSystem(),
		CMD:    adapters.NewExecCommander(),
		Writer: adapters.NewStdoutWriter(),
	}, cmd, args)
}

// runCreateWithDeps performs the actual project creation with injected dependencies.
// This function is designed to be testable with mock implementations.
func runCreateWithDeps(opts CreateOptions, cmd *cobra.Command, args []string) error {
	projectName := args[0]

	// Security: Validate project name to prevent path traversal
	if err := validateProjectName(projectName); err != nil {
		return fmt.Errorf("invalid project name: %w", err)
	}

	// Normalize and validate server type
	serverType = strings.ToLower(serverType)
	if err := validateServerType(serverType); err != nil {
		return err
	}

	// Get current directory
	currentDir, err := opts.FS.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	projectPath := filepath.Join(currentDir, projectName)

	// Check if directory already exists
	if _, err := opts.FS.Stat(projectPath); err == nil {
		return fmt.Errorf("directory '%s' already exists", projectName)
	}

	// Create project directory
	if err := opts.FS.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	opts.Writer.Printf("🚀 Creating project: %s\n", projectName)
	opts.Writer.Printf("📦 Server framework: %s\n", serverType)

	// Generate project
	config := generator.ProjectConfig{
		ProjectName: projectName,
		ProjectPath: projectPath,
		ModulePath:  fmt.Sprintf("github.com/yourusername/%s", projectName),
		ServerType:  serverType,
	}

	gen := generator.NewGeneratorWithFS(config, opts.FS)
	if err := gen.Generate(); err != nil {
		// Clean up on failure
		opts.FS.RemoveAll(projectPath)
		return fmt.Errorf("failed to generate project: %w", err)
	}

	// Run go mod tidy to download dependencies
	opts.Writer.Println("\n📦 Downloading dependencies...")
	if err := runGoModTidyWithDeps(opts, projectPath); err != nil {
		opts.Writer.Printf("⚠️  Warning: Failed to run go mod tidy: %v\n", err)
		opts.Writer.Println("   Please run 'go mod tidy' manually in the project directory")
	} else {
		opts.Writer.Println("✅ Dependencies downloaded successfully!")
	}

	opts.Writer.Println("\n✅ Project created successfully!")
	opts.Writer.Printf("\n📁 Next steps:\n")
	opts.Writer.Printf("   cd %s\n", projectName)
	opts.Writer.Printf("   make up\n")
	opts.Writer.Printf("   Visit http://localhost:8080/health/ready\n\n")

	return nil
}

// runGoModTidy executes go mod tidy in the project directory.
// This is a convenience wrapper for backward compatibility.
func runGoModTidy(projectPath string) error {
	return runGoModTidyWithDeps(CreateOptions{
		CMD: adapters.NewExecCommander(),
	}, projectPath)
}

// runGoModTidyWithDeps executes go mod tidy with injected command executor.
func runGoModTidyWithDeps(opts CreateOptions, projectPath string) error {
	return opts.CMD.Run("go", []string{"mod", "tidy"}, projectPath, os.Stdout, os.Stderr)
}

// validateProjectName ensures project name is safe and valid
func validateProjectName(name string) error {
	if name == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	// Prevent path traversal attacks
	if strings.Contains(name, "..") || strings.Contains(name, "/") || strings.Contains(name, "\\") {
		return fmt.Errorf("project name cannot contain path separators or '..'")
	}

	// Check for valid characters (alphanumeric, dash, underscore)
	validName := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	if !validName.MatchString(name) {
		return fmt.Errorf("project name can only contain letters, numbers, dashes, and underscores")
	}

	// Prevent reserved names
	reserved := []string{".", "..", "con", "prn", "aux", "nul", "com1", "com2", "com3", "com4", "lpt1", "lpt2"}
	for _, r := range reserved {
		if strings.EqualFold(name, r) {
			return fmt.Errorf("'%s' is a reserved name", name)
		}
	}

	return nil
}

// validateServerType ensures only supported server types are used
func validateServerType(server string) error {
	validServers := map[string]bool{
		"fiber": true,
		"gin":   true,
	}

	if !validServers[server] {
		return fmt.Errorf("invalid server type '%s': only 'fiber' and 'gin' are supported", server)
	}

	return nil
}
