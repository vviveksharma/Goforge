# Contributing to GoForge

Thank you for your interest in contributing to GoForge! This document provides guidelines and instructions for contributing.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How Can I Contribute?](#how-can-i-contribute)
- [Development Setup](#development-setup)
- [Pull Request Process](#pull-request-process)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Documentation](#documentation)

## Code of Conduct

This project adheres to a Code of Conduct. By participating, you are expected to uphold this code. Please report unacceptable behavior to the project maintainers.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check existing issues to avoid duplicates. When creating a bug report, include:

- **Clear title and description**
- **Steps to reproduce** the issue
- **Expected behavior** vs **actual behavior**
- **Environment details** (OS, Go version, etc.)
- **Code samples** or **error messages**

### Suggesting Enhancements

Enhancement suggestions are welcome! Please provide:

- **Clear title and description**
- **Use case** for the enhancement
- **Expected behavior** and implementation ideas
- **Examples** from other projects (if applicable)

### Adding New Framework Support

To add support for a new web framework:

1. Create template files in `internal/generator/templates/`
2. Add framework-specific handlers, middleware, and server files
3. Update `getServerSpecificTemplates()` in `generator.go`
4. Update validation in `validateServerType()`
5. Add tests for the new framework
6. Update documentation

### Pull Requests

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests (`make test`)
5. Run linters (`make lint`)
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## Development Setup

### Prerequisites

- Go 1.26.2 or higher
- Git
- Make (optional, but recommended)

### Setup Steps

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/goforge.git
   cd goforge
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Build the project**
   ```bash
   make build
   # OR
   go build -o goforge ./cmd/goforge
   ```

4. **Run tests**
   ```bash
   make test
   # OR
   go test -v -race ./...
   ```

5. **Install locally**
   ```bash
   go install ./cmd/goforge
   ```

## Pull Request Process

1. **Update documentation** - Update README.md, code comments, and any relevant docs
2. **Add tests** - Ensure new code has appropriate test coverage
3. **Pass all checks** - Tests, linters, and CI must pass
4. **Keep commits clean** - Use clear, descriptive commit messages
5. **One feature per PR** - Keep PRs focused on a single feature or fix
6. **Update CHANGELOG** - Add an entry describing your changes

### PR Checklist

- [ ] Tests added/updated and passing
- [ ] Documentation updated
- [ ] Linters pass (`make lint`)
- [ ] No breaking changes (or clearly documented)
- [ ] CHANGELOG.md updated
- [ ] Commits are clean and well-described

## Coding Standards

### Go Style

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` for formatting
- Use `golangci-lint` for linting
- Write idiomatic Go code

### Code Organization

```
goforge/
├── cmd/goforge/          # CLI entry point
├── internal/             # Private application code
│   ├── cmd/             # CLI commands (create, version, etc.)
│   ├── generator/       # Project generation logic
│   ├── interfaces/      # Dependency injection interfaces
│   ├── adapters/        # Real implementations
│   └── mocks/           # Mock implementations for testing
```

### Naming Conventions

- **Packages**: lowercase, single word (e.g., `generator`, `handler`)
- **Files**: lowercase with underscores (e.g., `create_test.go`)
- **Variables**: camelCase (e.g., `projectName`)
- **Exported**: PascalCase (e.g., `NewGenerator`)
- **Interfaces**: noun or adjective (e.g., `FileSystem`, `Commander`)

### Error Handling

```go
// Good: Wrap errors with context
if err := doSomething(); err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}

// Good: Handle errors immediately
result, err := compute()
if err != nil {
    return err
}

// Avoid: Ignoring errors
_ = doSomething() // Only when truly safe to ignore
```

### Dependency Injection

All external dependencies (filesystem, commands, I/O) should use interfaces:

```go
// Define interface
type FileSystem interface {
    Create(name string) (WriteCloser, error)
    // ...
}

// Accept interface in constructors
func NewGenerator(fs FileSystem) *Generator {
    // ...
}
```

## Testing

### Running Tests

```bash
# All tests
make test

# With coverage
go test -v -race -coverprofile=coverage.out ./...

# View coverage
go tool cover -html=coverage.out

# Specific package
go test -v ./internal/cmd/...
```

### Writing Tests

```go
func TestFeature(t *testing.T) {
    // Arrange
    mockFS := &mocks.MockFileSystem{
        CreateFunc: func(name string) (WriteCloser, error) {
            return mockWriter, nil
        },
    }

    // Act
    result := DoSomething(mockFS)

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, expected, result)
}
```

### Test Coverage Goals

- **Minimum**: 80% overall coverage
- **Critical paths**: 100% coverage
- **Error handling**: All error paths tested
- **Edge cases**: Boundary conditions tested

## Documentation

### Code Comments

```go
// Package generator provides project generation functionality.
package generator

// Generator handles template processing and file generation.
type Generator struct {
    config ProjectConfig
}

// Generate creates a new project from templates.
// It validates the project name, creates directories, and generates files.
func (g *Generator) Generate() error {
    // Implementation
}
```

### README Updates

When adding features:

1. Update feature list
2. Add usage examples
3. Update command documentation
4. Add troubleshooting tips (if needed)

### CHANGELOG Format

```markdown
## [Version] - YYYY-MM-DD

### Added
- New feature description

### Changed
- Changed feature description

### Fixed
- Bug fix description

### Removed
- Removed feature description
```

## Template Development

### Adding New Templates

1. Create template file in `internal/generator/templates/`
2. Use `.tmpl` extension
3. Add to `commonFiles` or `getServerSpecificTemplates()`
4. Test with both Fiber and Gin (if applicable)

### Template Variables

Available variables in templates:

- `{{.ProjectName}}` - Project name
- `{{.ProjectPath}}` - Full project path
- `{{.ModulePath}}` - Go module path
- `{{.ServerType}}` - "fiber" or "gin"

### Template Functions

```go
// Available functions
{{toLower .ServerType}}  // Convert to lowercase
{{toUpper .ServerType}}  // Convert to uppercase
{{eq .ServerType "gin"}} // Equality check
```

## Release Process

1. Update version in `internal/cmd/version.go`
2. Update CHANGELOG.md
3. Create git tag: `git tag -a v1.x.x -m "Release v1.x.x"`
4. Push tag: `git push origin v1.x.x`
5. CI will create GitHub release automatically

## Questions?

- Open an issue for bug reports or feature requests
- Start a discussion for questions or ideas
- Check existing issues and discussions first

## License

By contributing, you agree that your contributions will be licensed under the same license as the project.

Thank you for contributing to GoForge! 🚀
