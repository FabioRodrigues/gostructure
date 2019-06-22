package reservation

import (
	"errors"
	"testing"

	"github.com/fabiorodrigues/gostructure/ddd/domain"
	"github.com/fabiorodrigues/gostructure/ddd/domain/entities"
	"github.com/fabiorodrigues/gostructure/ddd/infra"

	"github.com/ditointernet/go-assert"
)

func TestReserve(t *testing.T) {
	t.Run("should return error when got an error when getting book", func(t *testing.T) {
		mockBookRepo := domain.MockBookRepository{
			GetByIDfn: func(ID string) (entities.Book, error) {
				return entities.Book{}, errors.New("error getting book")
			},
		}

		service := NewService(mockBookRepo, infra.EmailSenderMock{})

		err := service.Reserve("123")
		assert.Equal(t, "error getting book", err.Error())
	})
	t.Run("should return error when not sending email", func(t *testing.T) {
		mockEmailSender := infra.EmailSenderMock{
			Sendfn: func(opts infra.EmailOptions) error {
				return errors.New("error when sending email")
			},
		}

		service := NewService(domain.MockBookRepository{}, mockEmailSender)
		err := service.Reserve("123")
		assert.Equal(t, "error when sending email", err.Error())
	})

	t.Run("should send email when reserving the book", func(t *testing.T) {

		emailSentCount := 0
		mockEmailSender := infra.EmailSenderMock{
			Sendfn: func(opts infra.EmailOptions) error {
				emailSentCount++
				return nil
			},
		}

		service := NewService(domain.MockBookRepository{}, mockEmailSender)
		service.Reserve("123")

		assert.Equal(t, 1, emailSentCount)
	})

	t.Run("should not return error when reserving successfully", func(t *testing.T) {

		service := NewService(domain.MockBookRepository{}, infra.EmailSenderMock{})
		if err := service.Reserve("123"); err != nil {
			t.Fail()
		}
	})

	t.Run("should set available quantity of a book to zero when reserving the last one", func(t *testing.T) {
		var availableQuantity int

		repo := domain.MockBookRepository{
			GetByIDfn: func(ID string) (entities.Book, error) {
				return entities.CreateFakeBook(1), nil
			},
			SaveOrUpdatefn: func(book entities.Book) error {
				availableQuantity = book.AvailableQuantity
				return nil
			},
		}

		service := NewService(repo, infra.EmailSenderMock{})
		service.Reserve("123")
		assert.Equal(t, 0, availableQuantity)
	})
}
