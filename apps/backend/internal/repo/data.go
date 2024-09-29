package repo

import "github.com/urodstvo/book-shop/libs/models"

type Database struct {
	Users      []models.User
	Carts      []models.Cart
	Books      []models.Book
	BookPages  []models.BookPage
	Genres     []models.Genre
	BookGenres []models.BookGenre
	Orders     []models.Order
	OrderBooks []models.OrderBook
	Payments   []models.Payment
}
