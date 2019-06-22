package entities

import (
	"fmt"
)

//Book is an entity representing the books of the system
type Book struct {
	ID                string
	Title             string
	AvailableQuantity int
}

//NewBook is a constructor that has the mission of construct a valid book with
//every required fields
func NewBook(title string, availableQuantity int) Book {
	return Book{
		Title:             title,
		AvailableQuantity: availableQuantity,
	}
}

//Reserve is a method of the book that validates itself and return an error if
//it is not possible to reserve a book according to its rules
func (b *Book) Reserve() error {
	if b.AvailableQuantity == 0 {
		return fmt.Errorf("unable to reserve this book, there are no available books to reserve")
	}
	b.AvailableQuantity--
	return nil
}
