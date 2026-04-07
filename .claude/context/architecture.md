# 🏗️ Architecture

## 🎯 Purpose

Define a scalable, maintainable, and domain-driven architecture for the project.

---

## 📁 Project Structure

```
├── docker-compose.yml
├── Makefile
├── .env.example

├── docker/
│   ├── Dockerfile
│   └── otel-collector.yml

├── cmd/
│   ├── api/         # API entrypoint
│   ├── worker/      # Queue worker
│   ├── scheduler/   # Cron jobs
│   └── migrate/     # Migrations runner

├── migrations/
│   ├── migrate.go
│   └── seed.go

├── internal/
│   ├── config/        # Environment config
│   ├── router/        # Route definitions
│   ├── middleware/    # Auth, CORS, RateLimit, Logger
│   ├── events/        # Event bus (pub/sub)
│   ├── queue/         # Jobs + Redis
│   ├── scheduler/     # Scheduled tasks

│   ├── infra/
│   │   ├── database/  # GORM + MySQL
│   │   ├── redis/     # Redis connection
│   │   ├── mailer/    # SMTP
│   │   └── tracer/    # OpenTelemetry

│   └── domain/
│       └── user/      # Example domain
│           ├── model/        # GORM models
│           ├── repository/   # Interface + implementation
│           ├── service/      # Use case orchestration
│           │   └── usecase/  # Business rules
│           └── handler/      # HTTP layer (controllers)
```

---

## 🔁 Request Flow (STRICT)

```
Router → Middleware → Handler → Service → UseCase → Repository → Database
```

---

## 🧩 Layer Responsibilities

### Router

* Define routes
* Bind handlers
* No business logic

---

### Middleware

* Authentication (JWT)
* Authorization (AdminOnly)
* Rate limiting
* Logging
* CORS

---

### Handler (Controller)

* Parse request
* Validate input
* Call service
* Return response

❌ Must NOT contain business logic

---

### Service

* Orchestrates use cases
* Coordinates domain logic
* Handles transactions

---

### UseCase

* Contains business rules
* Pure domain logic
* No external dependencies

---

### Repository

* Interface definition
* Data persistence logic
* GORM implementation

---

### Infra

* External services (DB, Redis, SMTP, Tracing)
* Must NOT contain domain logic

---

## 🧠 Domain Rules

* Each domain is isolated
* Domains must NOT depend on each other directly
* Communication via:

  * Interfaces
  * Events

---

## 📦 Dependency Rule (CRITICAL)

Allowed direction:

```
Handler → Service → UseCase → Repository
```

Forbidden:

* Repository calling Service
* Domain depending on Infra directly
* Circular dependencies

---

## 🗂️ Naming Conventions

* Folders → lowercase
* Files → lowercase
* Structs → PascalCase
* JSON → camelCase

---

## 🚫 Forbidden

* Business logic in handlers
* Direct DB access outside repository
* Shared mutable state across domains
* Tight coupling between domains

---

## ✅ Goal

Create a clean, testable, scalable, and maintainable backend architecture.
