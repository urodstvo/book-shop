package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Login     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Rating    uint
}
