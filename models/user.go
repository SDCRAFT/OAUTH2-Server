package models

import "time"

type User struct {
	ID            uint64    `json:"id" gorm:"primarykey,unique"`
	Name          string    `json:"name" gorm:"unique"`
	Email         string    `json:"email" gorm:"unique"`
	Password      string    `json:"password"`
	EmailVerified bool      `json:"email_verified"`
	MinecraftUUID string    `json:"uuid"`
	CreatedAt     time.Time `json:"create_at"`
	UpdatedAt     time.Time `json:"update_at"`
}

func NewUser(name string, email string, pw string) *User {
	return &User{
		Name:          name,
		Email:         email,
		Password:      pw,
		EmailVerified: false,
	}
}
