package models

type BookPage struct {
	BookId  int    `json:"book_id"`
	Content string `json:"content,omitempty"`
}

func (BookPage) TableName() string {
	return "books_pages"
}
