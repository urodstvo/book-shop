package models

type BookGenre struct {
	BookId  int `json:"book_id"`
	GenreId int `json:"genre_id"`
}

func (BookGenre) TableName() string {
	return "books_genres"
}
