package main

import (
	"github.com/djsilvajr/go-skeleton/internal/config"
	"github.com/djsilvajr/go-skeleton/internal/infra/database"
	"github.com/djsilvajr/go-skeleton/migrations"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)
	migrations.Run(db)
	migrations.Seed(db)
}
