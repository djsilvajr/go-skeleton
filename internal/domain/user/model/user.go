package model

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

// User represents the users table.
// Mirrors the Laravel User model with soft-deletes and timestamps.
type User struct {
	ID        uint           `gorm:"primaryKey"                               json:"id"`
	Name      string         `gorm:"size:100;not null"                        json:"name"`
	Email     string         `gorm:"uniqueIndex;size:150;not null"            json:"email"`
	Password  string         `gorm:"not null"                                 json:"-"`
	Role      Role           `gorm:"type:enum('admin','user');default:'user'" json:"role"`
	CreatedAt time.Time      `                                    			  json:"createdAt"`
	UpdatedAt time.Time      `                                    			  json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"                        			  json:"deletedAt,omitempty"`
}
