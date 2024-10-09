package models

import "time"

type Request struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	BookName string `json:"book_name"`
	Comment  string `json:"comment"`

	CreatedAt time.Time `json:"created_at"`
}

func (Request) TableName() string {
	return "requests"
}
