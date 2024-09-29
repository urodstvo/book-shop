package models

import "github.com/google/uuid"

type BookGenre struct {
	BookId  uuid.UUID
	GenreId uuid.UUID
}
