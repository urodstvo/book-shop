package models

import "time"

type BookRating struct {
	BookId int `json:"book_id"`
	UserId int `json:"user_id"`
	Rating int `json:"rating"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (BookRating) TableName() string {
	return "books_ratings"
}
