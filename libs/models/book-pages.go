package models

import "time"

type BookPage struct {
	Id         uint      `json:"id"`
	BookId     uint      `json:"book_id"`
	PageNumber uint      `json:"page_number"`
	Content    string    `json:"content,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
