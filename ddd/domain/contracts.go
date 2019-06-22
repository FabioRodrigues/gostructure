package domain

import "github.com/fabiorodrigues/gostructure/ddd/domain/entities"

//ReservationService is a service to handle reservations operations
type ReservationService interface {
	Reserve(bookID string) error
}

//BookRepository is a repository of book.
//If you prefer, use a database provider directly in your services
//ignoring repositories
//to simplify and make it more `go-like`
type BookRepository interface {
	GetByID(ID string) (entities.Book, error)
	SaveOrUpdate(book entities.Book) error
}
