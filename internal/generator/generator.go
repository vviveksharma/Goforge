package generator

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/viveksharma/goforge/internal/adapters"
	"github.com/viveksharma/goforge/internal/interfaces"
)

//go:embed templates/*
var templatesFS embed.FS

type ProjectConfig struct {
	ProjectName string
	ProjectPath string
	ModulePath  string
	ServerType  string
}

type Generator struct {
	config ProjectConfig
	fs     interfaces.FileSystem
}

// NewGenerator creates a new Generator with real file system.
// This is a convenience wrapper for backward compatibility.
func NewGenerator(config ProjectConfig) *Generator {
	return NewGeneratorWithFS(config, adapters.NewOSFileSystem())
}

// NewGeneratorWithFS creates a new Generator with the provided file system.
// This enables dependency injection for testing.
func NewGeneratorWithFS(config ProjectConfig, fs interfaces.FileSystem) *Generator {
	return &Generator{
		config: config,
		fs:     fs,
	}
}

func (g *Generator) Generate() error {
	// Create directory structure
	if err := g.createDirectories(); err != nil {
		return err
	}

	// Generate files from templates
	if err := g.generateFiles(); err != nil {
		return err
	}

	return nil
}

func (g *Generator) createDirectories() error {
	dirs := []string{
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

	for _, dir := range dirs {
		path := filepath.Join(g.config.ProjectPath, dir)
		if err := g.fs.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	return nil
}

func (g *Generator) generateFiles() error {
	// Common files (same for all server types)
	commonFiles := map[string]string{
		"templates/go.mod.tmpl":                                        "go.mod",
		"templates/README.md.tmpl":                                     "README.md",
		"templates/Makefile.tmpl":                                      "Makefile",
		"templates/.env.example.tmpl":                                  ".env.example",
		"templates/.gitignore.tmpl":                                    ".gitignore",
		"templates/docker-compose.yml.tmpl":                            "docker-compose.yml",
		"templates/cmd/api/main.go.tmpl":                               "cmd/api/main.go",
		"templates/internal/config/config.go.tmpl":                     "internal/config/config.go",
		"templates/pkg/logger/logger.go.tmpl":                          "pkg/logger/logger.go",
		"templates/pkg/database/postgres.go.tmpl":                      "pkg/database/postgres.go",
		"templates/pkg/cache/redis.go.tmpl":                            "pkg/cache/redis.go",
		"templates/pkg/migration/migrate.go.tmpl":                      "pkg/migration/migrate.go",
		"templates/migrations/000001_create_users_table.up.sql.tmpl":   "migrations/000001_create_users_table.up.sql",
		"templates/migrations/000001_create_users_table.down.sql.tmpl": "migrations/000001_create_users_table.down.sql",
		"templates/docs/docs.go.tmpl":                                  "docs/docs.go",
		"templates/deployments/Dockerfile.tmpl":                        "deployments/Dockerfile",
		"templates/scripts/goswitch.tmpl":                              "scripts/goswitch",
	}

	// Server-specific files
	serverFiles := g.getServerSpecificTemplates()

	// Merge common and server-specific files
	allFiles := make(map[string]string)
	for k, v := range commonFiles {
		allFiles[k] = v
	}
	for k, v := range serverFiles {
		allFiles[k] = v
	}

	tmpl := template.New("project").Funcs(template.FuncMap{
		"toLower": strings.ToLower,
		"toUpper": strings.ToUpper,
		"eq":      func(a, b string) bool { return a == b },
	})

	for templatePath, outputPath := range allFiles {
		if err := g.generateFile(tmpl, templatePath, outputPath); err != nil {
			return err
		}
	}

	return nil
}

// getServerSpecificTemplates returns the template mappings based on server type
func (g *Generator) getServerSpecificTemplates() map[string]string {
	if g.config.ServerType == "gin" {
		return map[string]string{
			"templates/internal/handler/health_gin.go.tmpl":      "internal/handler/health.go",
			"templates/internal/handler/health_gin_test.go.tmpl": "internal/handler/health_test.go",
			"templates/internal/middleware/security_gin.go.tmpl": "internal/middleware/security.go",
			"templates/internal/middleware/logger_gin.go.tmpl":   "internal/middleware/logger.go",
			"templates/internal/middleware/recovery_gin.go.tmpl": "internal/middleware/recovery.go",
			"templates/internal/middleware/metrics_gin.go.tmpl":  "internal/middleware/metrics.go",
			"templates/internal/server/server_gin.go.tmpl":       "internal/server/server.go",
		}
	}

	// Default to Fiber
	return map[string]string{
		"templates/internal/handler/health.go.tmpl":      "internal/handler/health.go",
		"templates/internal/handler/health_test.go.tmpl": "internal/handler/health_test.go",
		"templates/internal/middleware/security.go.tmpl": "internal/middleware/security.go",
		"templates/internal/middleware/logger.go.tmpl":   "internal/middleware/logger.go",
		"templates/internal/middleware/recovery.go.tmpl": "internal/middleware/recovery.go",
		"templates/internal/middleware/metrics.go.tmpl":  "internal/middleware/metrics.go",
		"templates/internal/server/server.go.tmpl":       "internal/server/server.go",
	}
}

func (g *Generator) generateFile(tmpl *template.Template, templatePath, outputPath string) error {
	// Read template
	content, err := templatesFS.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template %s: %w", templatePath, err)
	}

	// Parse template
	t, err := tmpl.Parse(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
	}

	// Create output file
	outputFilePath := filepath.Join(g.config.ProjectPath, outputPath)
	file, err := g.fs.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outputPath, err)
	}
	defer file.Close()

	// Execute template
	if err := t.Execute(file, g.config); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", templatePath, err)
	}

	// Make shell scripts executable
	if strings.HasPrefix(outputPath, "scripts/") {
		if err := g.fs.Chmod(outputFilePath, 0755); err != nil {
			return fmt.Errorf("failed to make script executable %s: %w", outputPath, err)
		}
	}

	return nil
}
