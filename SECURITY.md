# Security Policy

## 🔒 Security Commitment

The security of generated projects is a top priority. This document outlines the security features included in every generated project and how to report security vulnerabilities.

## 🛡️ Security Features

### 1. Application Security

#### HTTP Security Headers

All responses include security headers:

- `X-Content-Type-Options: nosniff` - Prevents MIME sniffing
- `X-Frame-Options: DENY` - Prevents clickjacking
- `X-XSS-Protection: 1; mode=block` - XSS protection for legacy browsers
- `Content-Security-Policy: default-src 'self'` - Restricts resource loading
- `Referrer-Policy: strict-origin-when-cross-origin` - Controls referrer info
- `Permissions-Policy` - Restricts browser features
- `Strict-Transport-Security` (production) - Forces HTTPS

#### Input Validation

- Project name validation prevents path traversal
- Request body size limits (4MB default)
- Query parameter validation
- Path parameter sanitization

#### SQL Injection Prevention

- Parameterized queries enforced via `database/sql`
- Context-based timeouts
- Connection pool limits
- Prepared statement support

#### Authentication & Authorization

- Ready for JWT/OAuth2 integration
- Request ID tracking for audit trails
- Structured logging for security events

### 2. Infrastructure Security

#### Docker Security

- **Non-root user**: Containers run as UID 1000
- **Read-only filesystem**: Root filesystem is read-only
- **Minimal base image**: Alpine Linux (smallest attack surface)
- **No unnecessary capabilities**: All capabilities dropped
- **Multi-stage builds**: Build dependencies not in final image
- **Health checks**: Automatic container health monitoring

#### Database Security

- Connection pooling with limits (prevents exhaustion)
- Connection timeouts (prevents hanging)
- SSL/TLS support ready (set `sslmode=require`)
- Secrets via environment variables (never hardcoded)
- Parameterized queries only

#### Redis Security

- Connection pooling with resource limits
- Authentication ready (set password in URL)
- TLS support ready (use `rediss://` scheme)
- Key expiration enforced (prevents memory exhaustion)
- ACL support ready (Redis 6+)

### 3. Operational Security

#### Logging

- **No sensitive data**: Credentials never logged
- **Structured JSON**: Machine-parseable logs
- **Correlation IDs**: Request tracking
- **Stack traces**: Error debugging without exposure
- **Sanitized paths**: Query strings removed

#### Secrets Management

- `.env` files in `.gitignore`
- `.env.example` for documentation only
- Environment variable validation
- No default production credentials

#### Error Handling

- Panic recovery middleware
- Generic error messages to clients
- Detailed errors in logs (with request ID)
- No stack traces exposed to users

#### Graceful Shutdown

- SIGTERM/SIGINT handling
- Connection draining (30s timeout)
- Database connection cleanup
- Redis connection cleanup

### 4. Dependency Security

Generated projects use:

- **go.sum**: Dependency checksums verified
- **Minimal dependencies**: Only essential packages
- **High-quality dependencies**:
  - Fiber (actively maintained)
  - PostgreSQL driver (official)
  - Redis client (official)
  - Zap logger (Uber production-tested)

## 🔍 Security Checklist

Before deploying, ensure you:

### Required

- [ ] Change default database credentials
- [ ] Set strong passwords (min 20 characters)
- [ ] Enable SSL/TLS for database (`sslmode=require`)
- [ ] Enable Redis AUTH with strong password
- [ ] Enable Redis TLS (`rediss://`)
- [ ] Use environment variables for secrets
- [ ] Enable HSTS header in production
- [ ] Review and adjust CORS allowed origins
- [ ] Set `APP_ENV=production`
- [ ] Review security headers for your use case

### Recommended

- [ ] Implement rate limiting
- [ ] Add authentication middleware
- [ ] Enable database audit logging
- [ ] Set up monitoring and alerts
- [ ] Implement backup strategy
- [ ] Use secrets manager (AWS Secrets Manager, Vault)
- [ ] Enable container scanning
- [ ] Set up log aggregation
- [ ] Implement request signing
- [ ] Add API key validation

### Advanced

- [ ] Enable mutual TLS (mTLS)
- [ ] Implement database encryption at rest
- [ ] Use Redis ACLs to limit commands
- [ ] Add WAF (Web Application Firewall)
- [ ] Enable DDoS protection
- [ ] Implement SIEM integration
- [ ] Add intrusion detection
- [ ] Perform security scanning
- [ ] Conduct penetration testing
- [ ] Implement zero-trust architecture

## 🚨 Reporting a Vulnerability

If you discover a security vulnerability in:

### The Generator Tool (`goforge`)

Please report via:

- Email: sharmavivek1709@gmail.com
- GitHub Security Advisory: Create a draft security advisory
- Do NOT create a public issue

### Generated Projects

Security of generated projects is the responsibility of the project owner. However, if you find a vulnerability in the generated code templates, please report it as above.

## 📋 Vulnerability Response

We follow this process:

1. **Acknowledgment**: Within 24 hours
2. **Assessment**: Within 72 hours
3. **Fix Development**: Priority-based timeline
4. **Disclosure**: Coordinated disclosure after fix
5. **Credit**: Security researchers credited (if desired)

### Severity Levels

| Severity     | Response Time | Examples                        |
| ------------ | ------------- | ------------------------------- |
| **Critical** | 24 hours      | RCE, SQL injection in templates |
| **High**     | 3 days        | Authentication bypass, XSS      |
| **Medium**   | 7 days        | Information disclosure          |
| **Low**      | 14 days       | Minor misconfigurations         |

## 🛠️ Security Updates

### Version Support

| Version  | Supported |
| -------- | --------- |
| Latest   | ✅ Yes    |
| < Latest | ❌ No     |

Always use the latest version: `go install github.com/viveksharma/goforge@latest`

### Security Advisories

Subscribe to security advisories:

- GitHub Watch → Custom → Security alerts
- Check releases for security fixes
- Star the repo to stay updated

## 🔐 Security Best Practices

### Development

1. Never commit `.env` files
2. Use separate credentials for dev/staging/prod
3. Rotate credentials regularly (90 days)
4. Use strong, unique passwords
5. Enable 2FA on all accounts
6. Keep dependencies updated

### Production

1. Use secrets manager (never env vars in orchestrator)
2. Enable all TLS/SSL connections
3. Use network segmentation
4. Implement least privilege access
5. Enable audit logging
6. Monitor security logs
7. Implement automated backups
8. Test disaster recovery
9. Conduct regular security audits
10. Have an incident response plan

### Docker

1. Scan images for vulnerabilities
2. Use specific image tags (not `latest`)
3. Keep base images updated
4. Minimize image size
5. Use multi-stage builds
6. Don't run as root
7. Use read-only filesystem
8. Limit resources (CPU, memory)

## 📚 Security Resources

- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [OWASP Cheat Sheet Series](https://cheatsheetseries.owasp.org/)
- [Go Security Policy](https://go.dev/security/)
- [CIS Docker Benchmark](https://www.cisecurity.org/benchmark/docker)
- [Fiber Security](https://docs.gofiber.io/guide/security)

## 📞 Contact

Security concerns: sharmavivek1709@gmail.com
General issues: GitHub Issues
Documentation: GitHub Wiki

---

**Security is a journey, not a destination. Stay vigilant! 🛡️**
