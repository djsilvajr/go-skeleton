package migrations

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/djsilvajr/go-skeleton/internal/domain/user/model"
)

// Seed populates the database with initial data.
// Equivalent to `php artisan db:seed`.
func Seed(db *gorm.DB) {
	log.Println("🌱 Seeding database...")

	hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	users := []model.User{
		{Name: "Admin", Email: "admin@example.com", Password: string(hash), Role: model.RoleAdmin},
		{Name: "Alice", Email: "alice@example.com", Password: string(hash), Role: model.RoleUser},
	}

	for _, u := range users {
		u := u
		db.Where(model.User{Email: u.Email}).FirstOrCreate(&u)
	}

	log.Println("✅ Seeding complete")
}
