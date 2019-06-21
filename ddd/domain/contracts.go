package domain

import "github.com/fabiorodrigues/gostructure/ddd/domain/entities"

//ReservationService ...
type ReservationService interface {
	Reserve(bookID string) error
}

//BookRepository ...
type BookRepository interface {
	GetByID(ID string) (entities.Book, error)
	SaveOrUpdate(book entities.Book) error
}
