# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Database migration system with golang-migrate
- Swagger/OpenAPI documentation support
- Prometheus metrics middleware
- Test templates for generated projects
- Comprehensive test suite with 90%+ coverage
- CI/CD pipelines (GitHub Actions)
- Dependency injection architecture for testability

### Changed
- Made `--server` flag required (no default value)
- Automatic `go mod tidy` execution after project creation

### Fixed
- Import errors in generated projects

## [1.0.0] - 2026-04-12

### Added
- Initial release
- Support for Fiber and Gin web frameworks
- PostgreSQL database integration
- Redis cache integration
- Structured logging with Zap
- Docker and Docker Compose support
- Health check endpoints (liveness and readiness)
- Security middleware (headers, CORS, recovery)
- Request logging middleware
- Production-ready Dockerfile
- Makefile with common tasks
- Environment configuration with .env support

### Security
- Path traversal prevention in project names
- Reserved filename validation (Windows/Unix)
- Special character validation
- Secure default HTTP server settings
- Rate limiting and body size limits

[Unreleased]: https://github.com/yourusername/goforge/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/yourusername/goforge/releases/tag/v1.0.0
