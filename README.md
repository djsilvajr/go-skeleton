# Go Skeleton API

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00ACD7?logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker&logoColor=white)
![Auth](https://img.shields.io/badge/Auth-JWT-000000?logo=jsonwebtokens&logoColor=white)
![Cache](https://img.shields.io/badge/Cache-Redis-DC382D?logo=redis&logoColor=white)
![API](https://img.shields.io/badge/API-Swagger-85EA2D?logo=swagger&logoColor=black)
![Tracing](https://img.shields.io/badge/Tracing-Jaeger%20%2F%20OpenTelemetry-00A1C9?logo=opentelemetry&logoColor=white)
![Tests](https://img.shields.io/badge/Tests-Go%20Test-376ABD?logo=go&logoColor=white)

> Skeleton de API em Go — repository pattern, JWT, Redis, OpenTelemetry (Jaeger), queue worker, scheduler, eventos, rate limit e Swagger. Arquitetura espelhada no [Laravel-skeleton](https://github.com/djsilvajr/Laravel-skeleton), sem frontend.

---

## 🎯 Sobre este projeto

Base sólida para APIs RESTful em Go de médio porte, incluindo:

- ✅ **Arquitetura em camadas** — Handler → Service → Repository → Model
- ✅ **Repository pattern** com interfaces — troca de backend sem tocar na lógica de negócio
- ✅ **JWT Auth** — guards separados por rota, role-based (admin/user)
- ✅ **Rate limit** — por IP, configurável por rota
- ✅ **Redis** — cache + fila de jobs
- ✅ **Queue worker** — processo separado, handlers registráveis
- ✅ **Scheduler** — tarefas recorrentes em processo dedicado
- ✅ **Events** — pub/sub interno síncrono e assíncrono (fire-and-forget)
- ✅ **Swagger/OpenAPI** — gerado via anotações no código
- ✅ **Observabilidade** — OpenTelemetry → Jaeger (liga com uma env)
- ✅ **Testes unitários e de integração** — mocks hand-rolled, sem dependências externas
- ✅ **Ambiente dockerizado** — app + worker + scheduler + MySQL + Redis + Jaeger prontos

---

## 🗂️ Estrutura do repositório

```
/
├── docker-compose.yml          # Orquestração dos serviços
├── Makefile                    # Atalhos para comandos comuns
├── .env.example                # Variáveis de ambiente
├── docker/
│   ├── Dockerfile              # Multi-stage build
│   └── otel-collector.yml      # Config do OpenTelemetry collector
├── cmd/
│   ├── api/        main.go     # Entrypoint da API
│   ├── worker/     main.go     # Entrypoint do queue worker
│   ├── scheduler/  main.go     # Entrypoint do scheduler
│   └── migrate/    main.go     # Entrypoint das migrations
├── migrations/
│   ├── migrate.go              # AutoMigrate de todos os models
│   └── seed.go                 # Seeders (dados iniciais)
└── internal/
    ├── config/                 # Carrega variáveis de ambiente
    ├── router/                 # Rotas — equivalente ao routes/api.php
    ├── middleware/             # Auth JWT, AdminOnly, RateLimit, CORS, Logger
    ├── events/                 # Bus pub/sub interno
    ├── queue/                  # Worker + Dispatch via Redis
    ├── scheduler/              # Tarefas recorrentes
    ├── infra/
    │   ├── database/           # Conexão GORM + MySQL
    │   ├── redis/              # Conexão + RedisService
    │   ├── mailer/             # SMTP
    │   └── tracer/             # OpenTelemetry
    └── domain/
        └── user/               # Exemplo de domínio completo
            ├── model/          # Struct GORM
            ├── repository/     # Interface + implementação GORM
            ├── service/        # Lógica de negócio + erros sentinel
            └── handler/        # HTTP handlers (Controllers)
```

---

## 🛠️ Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/) v2

---

## ⚙️ Configuração inicial

```bash
cp .env.example .env
```

Suba os containers:

```bash
docker compose up -d
```

---

## 🗄️ Migrations e seed

```bash
docker exec -it go-skeleton-app sh
```

```bash
./migrate
```

Para popular o banco com dados iniciais, edite `migrations/seed.go` e chame `migrations.Seed(db)` no entrypoint de migrate.

Usuários criados por padrão:

| Email | Senha | Role |
|---|---|---|
| admin@example.com | password | admin |
| alice@example.com | password | user |

---

## ▶️ Acessando a aplicação

| Serviço | URL |
|---|---|
| API | http://localhost:8020/api/v1 |
| Health | http://localhost:8020/health |
| MySQL | localhost:3306 |
| Redis | localhost:6379 |
| Jaeger UI | http://localhost:16686 |

---

## 🔑 Rotas disponíveis

```
POST   /api/v1/auth/login       # Autenticação — retorna JWT
POST   /api/v1/auth/register    # Cadastro

GET    /api/v1/users            # Listar usuários       [JWT]
GET    /api/v1/users/:id        # Buscar por ID         [JWT]
POST   /api/v1/users            # Criar usuário         [JWT]
PUT    /api/v1/users/:id        # Atualizar usuário     [JWT]
DELETE /api/v1/users/:id        # Deletar usuário       [JWT + admin]

GET    /health                  # Health check
```

---

## 🧩 Comandos úteis

```bash
# Subir tudo
make up

# Rodar a API localmente (sem Docker)
make run

# Rodar worker localmente
make worker

# Rodar scheduler localmente
make scheduler

# Gerar docs Swagger (requer swag instalado)
make swagger

# Rodar todos os testes
make test

# Testes com cobertura
make test-cover

# Lint (requer golangci-lint)
make lint

# Limpar dependências
make tidy
```

---

## 🔄 Adicionando um novo domínio

Siga o mesmo padrão do domínio `user`:

```
internal/domain/produto/
    model/      produto.go          # struct GORM
    repository/ produto_repository.go   # interface + impl
    service/    produto_service.go  # lógica de negócio
                errors.go           # erros sentinel
    handler/    produto_handler.go  # HTTP handlers com anotações Swagger
```

Registre no router em `internal/router/router.go`:

```go
produtoRepo    := repository.NewProdutoRepository(db)
produtoSvc     := service.NewProdutoService(produtoRepo)
produtoHandler := handler.NewProdutoHandler(produtoSvc)

produtos := protected.Group("/produtos")
produtos.GET("",      produtoHandler.List)
produtos.GET("/:id",  produtoHandler.Show)
produtos.POST("",     produtoHandler.Store)
produtos.PUT("/:id",  produtoHandler.Update)
produtos.DELETE("/:id", middleware.AdminOnly(), produtoHandler.Destroy)
```

Registre a migration em `migrations/migrate.go`:

```go
db.AutoMigrate(&model.Produto{})
```

---

## 📬 Disparando um job para a fila

```go
queue.Dispatch(ctx, rdb, "send_welcome_email", map[string]string{
    "email": user.Email,
    "name":  user.Name,
})
```

Registre o handler no `cmd/worker/main.go`:

```go
worker.Register("send_welcome_email", func(ctx context.Context, payload json.RawMessage) error {
    // processar aqui
    return nil
})
```

---

## 📡 Disparando um evento

```go
events.DispatchAsync(events.UserCreated, map[string]any{"user_id": user.ID})
```

Registre um listener em qualquer `init()` ou no bootstrap:

```go
events.Listen(events.UserCreated, func(e events.Event) {
    log.Printf("usuário criado: %v", e.Payload)
})
```

---

## 🔭 Observabilidade

Por padrão o Jaeger fica desligado. Para ativar:

```env
OTEL_ENABLED=true
```

Acesse os traces em http://localhost:16686.

---

## ✅ Checklist rápido

- [ ] `cp .env.example .env`
- [ ] `docker compose up -d`
- [ ] `docker exec -it go-skeleton-app /app/migrate`
- [ ] Testar `POST /api/v1/auth/login` com `admin@example.com` / `password`
- [ ] Acessar Swagger em http://localhost:8020/api/documentation/index.html

Pronto 🎉

---

## 🔒 Checklist de segurança para produção

Veja [SECURITY.md](./SECURITY.md).

---

## 📄 License

MIT
