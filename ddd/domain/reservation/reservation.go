package reservation

import (
	"github.com/fabiorodrigues/gostructure/ddd/infra"
)

//Service ...
type Service struct {
	bookRepo infra.BookRepository
}

//NewService ...
func NewService(bookRepo infra.BookRepository) Service {
	return Service{
		bookRepo: bookRepo,
	}
}

//Reserve ...
func (s Service) Reserve(bookID string) error {

	book, err := s.bookRepo.GetByID(bookID)
	if err != nil {
		return err
	}

	if err := book.Reserve(); err != nil {
		return err
	}

	if err := s.bookRepo.SaveOrUpdate(book); err != nil {
		return err
	}

	return nil
}
