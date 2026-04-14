# Goforge

Forge production-ready Go applications with security, observability, and best practices built-in.

## 🚀 Installation

### Option 1: Go Install (Recommended)

```bash
go install github.com/viveksharma/goforge@latest
```

### Option 2: Download Binary

Download pre-built binaries from the [Releases](https://github.com/viveksharma/goforge/releases) page.

### Option 3: Install from Latest Main

```bash
git clone https://github.com/viveksharma/goforge.git
cd goforge
make install
```

This installs the latest development version to `$GOPATH/bin` (usually `~/go/bin`).

### Option 4: Build from Source

```bash
git clone https://github.com/viveksharma/goforge.git
cd goforge
go build -o goforge ./cmd/goforge
sudo mv goforge /usr/local/bin/
```

## 📖 Quick Start

### Basic Usage

Create a new project in seconds:

```bash
# Create with Fiber (high performance)
goforge create my-api -s fiber

# Create with Gin (feature-rich)
goforge create my-api -s gin
```

### Custom Module Path (Optional)

By default, your Go module will be named after your project (e.g., `module my-api`).

If you want a custom module path for GitHub/GitLab or organization projects:

```bash
# With custom module path
goforge create my-api -s fiber --module github.com/yourusername/my-api

# Short form
goforge create my-api -s fiber -m github.com/viveksharma/my-api
```

### Start Developing

```bash
cd my-api
make up
```

Your API is now running at `http://localhost:8080` 🚀

**Quick health check:**
```bash
curl http://localhost:8080/health/ready
```

## ✨ What You Get

Every generated project includes:

### 🎯 Framework Choice

- **Fiber**: Blazing-fast, Express-inspired framework with zero memory allocation router
- **Gin**: Feature-rich framework with excellent middleware ecosystem and proven track record
- **Same Quality**: Both options include identical security, observability, and production-ready features

### 🔒 Security First

- **Security Headers**: HSTS, CSP, X-Frame-Options, X-Content-Type-Options
- **Input Validation**: Path traversal protection, request validation
- **No Sensitive Logging**: Credentials never appear in logs
- **Panic Recovery**: Graceful error handling
- **Rate Limiting Ready**: Redis-backed rate limiting structure
- **Secure Defaults**: Non-root Docker user, read-only filesystem
- **Timeouts**: Request/connection timeouts prevent DoS

### 📊 Observability

- **Structured Logging**: JSON logs with correlation IDs (zap)
- **Health Checks**: Kubernetes-ready `/health/live` and `/health/ready`
- **Prometheus Metrics**: Built-in `/metrics` endpoint with request counters, duration histograms, and in-flight gauges
- **Request Tracing**: Unique request IDs for log correlation
- **Error Tracking**: Contextual error logging with stack traces

### 🏗️ Production Ready

- **Choice of Framework**: Fiber (high-performance) or Gin (feature-rich)
- **PostgreSQL**: Production-grade database with connection pooling
- **Redis**: Caching layer with connection management
- **Database Migrations**: Built-in golang-migrate support with up/down/version/force commands
- **Swagger/OpenAPI**: Auto-generated API documentation at `/swagger/index.html`
- **Docker Compose**: Zero-config local development
- **Graceful Shutdown**: Connection draining on SIGTERM/SIGINT
- **Environment Management**: Type-safe `.env` configuration

### 🔧 Developer Experience

- **Make Commands**: Common tasks via Makefile
- **Hot Reload Ready**: Easy integration with Air or CompileDaemon
- **Clean Architecture**: Separation of concerns (handler/service/repository pattern ready)
- **Dockerfile Included**: Multi-stage build with security best practices
- **Comprehensive README**: Documentation generated with every project

## 📦 Project Structure

```
your-project/
├── cmd/api/              # Application entry point
├── internal/
│   ├── config/           # Environment configuration
│   ├── handler/          # HTTP request handlers
│   ├── middleware/       # HTTP middleware (security, logging, recovery, metrics)
│   └── server/           # Server setup and routing
├── pkg/
│   ├── logger/           # Structured logging
│   ├── database/         # PostgreSQL client
│   ├── cache/            # Redis client
│   └── migration/        # Database migration runner
├── migrations/           # SQL migration files (up/down)
├── docs/                 # Swagger/OpenAPI documentation
├── deployments/
│   └── Dockerfile        # Multi-stage production build
├── docker-compose.yml    # Local development stack
├── Makefile              # Common commands
└── .env.example          # Environment variables template
```

## 🔒 Security Features

### Built-in Protections

1. **Path Traversal Prevention**: Project name validation prevents `../` attacks
2. **SQL Injection Protection**: Parameterized queries enforced
3. **XSS Protection**: Security headers set by default
4. **Clickjacking Protection**: X-Frame-Options: DENY
5. **MIME Sniffing Protection**: X-Content-Type-Options: nosniff
6. **Request Timeout Protection**: Read/write timeouts configured
7. **Body Size Limits**: 4MB default limit
8. **Non-Root Docker User**: Containers run as user 1000
9. **Read-Only Filesystem**: Docker containers use read-only root
10. **Secrets Management**: `.env` files never committed

### Security Best Practices

All generated code follows:

- OWASP Top 10 protection
- Principle of least privilege
- Defense in depth
- Secure defaults
- Input validation
- Output encoding (where applicable)

## 📚 Generated Project Commands

```bash
make up              # Start all services (API, PostgreSQL, Redis)
make down            # Stop all services
make logs            # View logs
make build           # Build the Go binary
make run             # Run without Docker
make test            # Run tests
make fmt             # Format code
make lint            # Run linter
make clean           # Clean up
make swagger         # Generate Swagger docs
make migrate-up      # Run all pending migrations
make migrate-down    # Revert last migration
make migrate-create  # Create migration (NAME=migration_name)
make migrate-version # Show current migration version
make migrate-force   # Force migration version (VERSION=1)
```

## 🛠️ Customization

Generated projects are fully customizable:

1. Add your business logic in `internal/handler/`
2. Create services in `internal/service/`
3. Add repositories in `internal/repository/`
4. Update routes in `internal/server/server.go`
5. Modify environment variables in `.env`

## 🔐 Environment Variables

Every project includes these environment variables:

| Variable       | Description                           | Default       |
| -------------- | ------------------------------------- | ------------- |
| `APP_ENV`      | Environment (development/production)  | `development` |
| `APP_PORT`     | HTTP server port                      | `8080`        |
| `DATABASE_URL` | PostgreSQL connection string          | Required      |
| `REDIS_URL`    | Redis connection string               | Required      |
| `LOG_LEVEL`    | Logging level (debug/info/warn/error) | `info`        |

## 🚢 Deployment

Generated projects are deployment-ready for:

- **Kubernetes**: Health check endpoints configured
- **Docker**: Multi-stage Dockerfile included
- **Cloud Run**: Listens on PORT environment variable
- **AWS ECS/Fargate**: 12-factor app compliant
- **Any platform**: Standard REST API

## 🧪 Testing

Every generated project is test-ready:

```bash
cd my-project
make test
```

Add your tests in:

- `internal/handler/*_test.go`
- `internal/service/*_test.go`
- `pkg/*_test.go`

## 📖 Examples

### Simple Project (Local Development)

For quick prototyping or local projects, use the simple syntax:

```bash
goforge create my-api -s fiber
# Result: go.mod → module my-api
```

### GitHub/GitLab Project

For projects you'll push to GitHub or GitLab, specify the full module path:

```bash
goforge create my-api -s fiber -m github.com/yourusername/my-api
# Result: go.mod → module github.com/yourusername/my-api
```

### Organization/Company Project

For company or organization projects:

```bash
goforge create payment-service -s gin -m gitlab.company.com/backend/payment-service
# Result: go.mod → module gitlab.company.com/backend/payment-service
```

### What Happens When You Create a Project

1. ✅ Creates project directory
2. ✅ Generates all project files with your chosen framework
3. ✅ Sets up PostgreSQL database with connection pooling
4. ✅ Configures Redis for caching
5. ✅ Adds health check endpoints (`/health/live`, `/health/ready`)
6. ✅ Sets up database migrations system
7. ✅ Generates Swagger/OpenAPI docs at `/swagger/index.html`
8. ✅ Adds Prometheus metrics at `/metrics`
9. ✅ Includes Docker Compose for one-command startup
10. ✅ Creates comprehensive README with all commands
11. ✅ Downloads dependencies automatically

### Start Coding Immediately

```bash
cd my-api
make up        # Starts PostgreSQL, Redis, and your API
make logs      # Watch the logs

# Test your API
curl http://localhost:8080/health/ready
curl http://localhost:8080/metrics

# View API docs
open http://localhost:8080/swagger/index.html
```

## 🔀 Choosing a Framework

### Use Fiber when:
- 🚀 You need **maximum performance** and minimal memory footprint
- 💚 You prefer **Express.js-like** syntax and patterns
- ⚡ Your application handles **high concurrent loads**
- 🏎️ You want the **fastest** request/response times

### Use Gin when:
- 🛡️ You prefer a **mature, battle-tested** framework
- 🔧 You need **extensive middleware** ecosystem
- ✅ You want **built-in validation** and binding
- 👥 Your team is **already familiar** with Gin

**Both frameworks generate identical project structure and features** - the only difference is the web framework implementation.

## 🎯 Module Path Guide

### When to Use Default (No `--module` flag)

```bash
goforge create my-project -s fiber
```

✅ **Perfect for:**
- Quick prototypes and experiments
- Learning projects
- Local-only development
- Tools and scripts that won't be shared

**Result:** `module my-project` (simple and clean)

### When to Use Custom Module Path

```bash
goforge create my-project -s fiber -m github.com/username/my-project
```

✅ **Perfect for:**
- Projects you'll push to GitHub/GitLab
- Open source projects
- Company/organization codebases
- Projects with internal imports

**Result:** `module github.com/username/my-project`

### Module Path Examples

```bash
# GitHub
goforge create api -s fiber -m github.com/viveksharma/api

# GitLab
goforge create service -s gin -m gitlab.com/myorg/service

# Self-hosted GitLab
goforge create payment -s fiber -m git.company.com/backend/payment

# Bitbucket
goforge create auth -s gin -m bitbucket.org/team/auth

# Company domain
goforge create users -s fiber -m go.company.com/services/users
```

## 🤝 Contributing

Contributions welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

1. Fork the repository
2. Create a feature branch
3. Add tests for new features
4. Ensure all tests pass
5. Submit a pull request

## 📝 License

MIT License - feel free to use this for any project.

## 🐛 Issues

Found a bug or have a feature request? Open an issue on GitHub.

## ⭐ Show Your Support

If this tool helped you, give it a star on GitHub!

---

**Made with ❤️ for the Go community**
