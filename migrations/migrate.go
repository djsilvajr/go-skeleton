package migrations

import (
	"log"

	"gorm.io/gorm"

	"github.com/djsilvajr/go-skeleton/internal/domain/user/model"
)

// Run executes auto-migrations for all registered models.
// Equivalent to `php artisan migrate`.
func Run(db *gorm.DB) {
	log.Println("🗄️  Running migrations...")

	if err := db.AutoMigrate(
		&model.User{},
		// register new models here
	); err != nil {
		log.Fatalf("migration error: %v", err)
	}

	log.Println("✅ Migrations complete")
}
