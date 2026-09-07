package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents the users table in PostgreSQL
type User struct {
	ID        string    `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName explicitly specifies the database table name
func (User) TableName() string {
	return "users"
}

// BeforeCreate GORM hook to generate UUID if not provided
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return nil
}
