package models

import "github.com/google/uuid"

type BookGenre struct {
	BookId  uuid.UUID `json:"book_id"`
	GenreId uuid.UUID `json:"genre_id"`
}
