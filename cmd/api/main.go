package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/djsilvajr/go-skeleton/internal/config"
	"github.com/djsilvajr/go-skeleton/internal/infra/database"
	"github.com/djsilvajr/go-skeleton/internal/infra/redis"
	"github.com/djsilvajr/go-skeleton/internal/infra/tracer"
	"github.com/djsilvajr/go-skeleton/internal/router"
)

// @title           Go Skeleton API
// @version         1.0
// @description     API skeleton in Go — repository pattern, JWT auth, Redis, OpenTelemetry.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://github.com/djsilvajr/go-skeleton

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8020
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.Load()

	// Observability
	if cfg.OtelEnabled {
		shutdown, err := tracer.Init(cfg)
		if err != nil {
			log.Printf("tracer init warning: %v", err)
		} else {
			defer shutdown(context.Background())
		}
	}

	// Infrastructure
	db := database.Connect(cfg)
	rdb := redis.Connect(cfg)

	// Router
	r := router.Setup(cfg, db, rdb)

	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: r,
	}

	go func() {
		log.Printf("🚀 Server running on http://localhost:%s", cfg.AppPort)
		log.Printf("📄 Swagger: http://localhost:%s/api/documentation/index.html", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
