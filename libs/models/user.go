package models

import (
	"time"
)

type User struct {
	Id          int       `json:"id"`
	Login       string    `json:"login"`
	Password    string    `json:"password,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Rating      *uint     `json:"rating"`
	RatingCount uint      `json:"rating_count"`
}

func (User) TableName() string {
	return "users"
}
