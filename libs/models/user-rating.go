package models

type BookRating struct {
	BookId int `json:"bookId"`
	UserId int `json:"userId"`
	Rating int `json:"rating"`
}

func (BookRating) TableName() string {
	return "books_ratings"
}
