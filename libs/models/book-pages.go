package models

import "time"

type BookPage struct {
	Id         uint
	BookId     uint
	PageNumber uint
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
