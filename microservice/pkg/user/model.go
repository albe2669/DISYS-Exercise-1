package user

import (
	"time"

	"gorm.io/gorm"
)

type CreateUser struct {
	Name       string `json:"name"`
	Mail       string `json:"mail"`
	Popularity int64  `json:"popularity,omitempty"`
	Role       string `json:"role"`
}

type User struct {
	CreateUser
	ID        uint64         `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func CreateToReal(user CreateUser) User {
	return User{CreateUser: user}
}
