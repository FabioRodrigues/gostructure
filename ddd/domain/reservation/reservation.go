package reservation

import (
	"github.com/fabiorodrigues/gostructure/ddd/domain"
	"github.com/fabiorodrigues/gostructure/ddd/infra"
)

//Service ...
type Service struct {
	bookRepo    domain.BookRepository
	emailSender infra.EmailSender
}

//NewService ...
func NewService(
	bookRepo domain.BookRepository,
	emailSender infra.EmailSender) Service {
	return Service{
		bookRepo:    bookRepo,
		emailSender: emailSender,
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

	return s.emailSender.Send(infra.EmailOptions{
		From:   "from@email.com",
		To:     []string{"to@email.com"},
		Body:   "Book is reserved. Please return it in 5 days",
		IsHTML: false,
	})

}
