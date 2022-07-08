package book

import "time"

type Book struct {
	Id          int
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Book) TableName() string {
	return "books"
}
