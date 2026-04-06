# Architecture

## Structure

├── docker-compose.yml          # Compose     
├── Makefile                    # Some commands 
├── .env.example                # .env
├── docker/
│   ├── Dockerfile              # Docker builder
│   └── otel-collector.yml      # OpenTelemetry collector config
├── cmd/
│   ├── api/        main.go     # API Entrypoint
│   ├── worker/     main.go     # Queue worker Entrypoint 
│   ├── scheduler/  main.go     # Scheduler Entrypoint  
│   └── migrate/    main.go     # Migrations Entrypoint 
├── migrations/
│   ├── migrate.go              # AutoMigrate
│   └── seed.go                 # Seeders
└── internal/
    ├── config/                 # Load .env
    ├── router/                 # API Router
    ├── middleware/             # Auth JWT, AdminOnly, RateLimit, CORS, Logger
    ├── events/                 # Bus pub/sub
    ├── queue/                  # Worker + Dispatch using Redis
    ├── scheduler/              # Scheduler
    ├── infra/
    │   ├── database/           # connection GORM + MySQL
    │   ├── redis/              # connection + RedisService
    │   ├── mailer/             # SMTP
    │   └── tracer/             # OpenTelemetry
    └── domain/
        └── user/               # User domain example
            ├── model/          # Struct GORM
            ├── repository/     # Interface + GORM implementation
            ├── service/        # Call useCases
            |    └── useCase    # businessRule
            └── handler/        # HTTP handlers (Controllers)

## Pattern

Router → Domain → {DomainName} → handler → service → Interface / Repository

