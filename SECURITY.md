# Security Policy

## ⚠️ Development Configuration

This skeleton uses simplified configurations for local development.  
**Do not use these defaults in production.**

## Checklist before production deploy

- [ ] Move all credentials to environment variables (`.env`)
- [ ] Set a strong `JWT_SECRET` (min. 32 random characters)
- [ ] Set a strong `REDIS_PASSWORD`
- [ ] Set `APP_DEBUG=false`
- [ ] Configure HTTPS / TLS termination
- [ ] Review CORS settings in `middleware/logger.go`
- [ ] Protect or remove `/api/documentation` in production
- [ ] Configure MySQL with a non-root user and strong password
- [ ] Enable rate limiting per-route as needed
- [ ] Keep dependencies up to date (`go get -u ./...`)

## Reporting a Vulnerability

Open a GitHub issue marked **[SECURITY]** or contact the maintainer directly.
