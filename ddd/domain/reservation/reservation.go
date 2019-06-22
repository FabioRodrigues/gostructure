package reservation

import (
	"github.com/fabiorodrigues/gostructure/ddd/domain"
	"github.com/fabiorodrigues/gostructure/ddd/infra"
)

//Service is a service of reservation in this case.
//it is used to handle everything needed to make operations about reservations.
//here is allowed to send email, instantiate everything it needs,
//call other services, send email, connect to databases etc.
type Service struct {
	bookRepo    domain.BookRepository
	emailSender infra.EmailSender
}

//NewService creates a new instance of a reservation service.
//it takes as params every necessary dependencies to work.
func NewService(
	bookRepo domain.BookRepository,
	emailSender infra.EmailSender) Service {
	return Service{
		bookRepo:    bookRepo,
		emailSender: emailSender,
	}
}

//Reserve is an exposed method that clearly reserve a book
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
