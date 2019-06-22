package entities

import (
	"fmt"
)

//Book ...
type Book struct {
	ID                string
	Title             string
	AvailableQuantity int
}

//NewBook ...
func NewBook(title string, availableQuantity int) Book {
	return Book{
		Title:             title,
		AvailableQuantity: availableQuantity,
	}
}

//Reserve ...
func (b *Book) Reserve() error {
	if b.AvailableQuantity == 0 {
		return fmt.Errorf("unable to reserve this book, there are no available books to reserve")
	}
	b.AvailableQuantity--
	return nil
}
