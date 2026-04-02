package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/djsilvajr/go-skeleton/internal/config"
	"github.com/djsilvajr/go-skeleton/internal/infra/database"
	"github.com/djsilvajr/go-skeleton/internal/scheduler"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)

	s := scheduler.New()

	// Register recurring tasks here. Mirror of Laravel Kernel::schedule().
	s.Add("cleanup_expired_tokens", 1*time.Hour, func(ctx context.Context) {
		log.Println("scheduler: cleaning up expired tokens")
		// db.Exec("DELETE FROM personal_access_tokens WHERE expires_at < NOW()")
		_ = db
	})

	s.Add("health_ping", 1*time.Minute, func(ctx context.Context) {
		log.Println("scheduler: ping — all good")
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	s.Start(ctx)
}
