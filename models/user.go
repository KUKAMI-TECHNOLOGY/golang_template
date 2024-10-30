package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;" json:"id"`
	Username string    `gorm:"unique;not null" json:"username"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `json:"password"`
}

type Users struct {
	Users []User `json:"users"`
}

// BeforeCreate is a GORM hook that runs before creating a new User record.
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate a new UUID for the user ID
	user.ID = uuid.New()
	return
}
