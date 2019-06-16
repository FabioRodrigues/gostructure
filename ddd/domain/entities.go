package domain

import "fmt"

//Book ...
type Book struct {
	ID                string
	Title             string
	AvailableQuantity int
}

//Reserve ...
func (b Book) Reserve() error {
	if b.AvailableQuantity == 0 {
		return fmt.Errorf("unable to reserve this book, there are no available books to reserve")
	}
	b.AvailableQuantity--
	return nil
}
