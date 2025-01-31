package models

import "github.com/google/uuid"

type User struct {
	ID            uuid.UUID `gorm:"<-:create;type:uuid;unique;primaryKey"`
	Name          string    `json:"name" gorm:"unique"`
	Email         string    `json:"email" gorm:"unique"`
	Password      string    `json:"password"`
	EmailVerified bool      `json:"email_verified"`
	MinecraftUUID string    `json:"uuid"`
	CreatedAt     uint64    `json:"create_at"`
	UpdatedAt     uint64    `json:"update_at"`
}

func NewUser(name string, email string, pw string) *User {
	return &User{
		ID:            uuid.New(),
		Name:          name,
		Email:         email,
		Password:      pw,
		EmailVerified: false,
	}
}
