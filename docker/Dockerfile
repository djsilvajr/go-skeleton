# ── Build stage ────────────────────────────────────────────────────────────────
FROM golang:1.22-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Generate swagger docs before building
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init -g cmd/api/main.go -o docs

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/api      ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/worker   ./cmd/worker
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/scheduler ./cmd/scheduler

# ── Runtime stage ───────────────────────────────────────────────────────────────
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/api       /app/api
COPY --from=builder /app/worker    /app/worker
COPY --from=builder /app/scheduler /app/scheduler

EXPOSE 8020

CMD ["/app/api"]
