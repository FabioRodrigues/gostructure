package reservation

import (
	"errors"
	"testing"

	"github.com/fabiorodrigues/gostructure/ddd/domain"
	"github.com/fabiorodrigues/gostructure/ddd/domain/entities"
	"github.com/fabiorodrigues/gostructure/ddd/infra"
	"github.com/stretchr/testify/assert"
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
}
