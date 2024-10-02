package models

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Login     string    `json:"login"`
	Password  string    `json:"-,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
}

func (User) TableName() string {
	return "users"
}
