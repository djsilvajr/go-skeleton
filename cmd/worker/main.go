package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/djsilvajr/go-skeleton/internal/config"
	"github.com/djsilvajr/go-skeleton/internal/infra/redis"
	"github.com/djsilvajr/go-skeleton/internal/queue"
)

func main() {
	cfg := config.Load()
	rdb := redis.Connect(cfg)

	worker := queue.NewWorker(rdb)

	// Register job handlers here.
	// Example: send welcome email after user creation.
	worker.Register("send_welcome_email", func(ctx context.Context, payload json.RawMessage) error {
		var data map[string]string
		if err := json.Unmarshal(payload, &data); err != nil {
			return err
		}
		log.Printf("📧 Sending welcome email to %s", data["email"])
		// wire mailer.Send() here
		return nil
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	worker.Run(ctx)
}
