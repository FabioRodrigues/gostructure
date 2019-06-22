package reservation

import (
	"errors"
	"testing"

	"github.com/fabiorodrigues/gostructure/ddd/domain"
	"github.com/fabiorodrigues/gostructure/ddd/infra"
)

func TestReserve(t *testing.T) {
	t.Run("should return error when not", func(t *testing.T) {
		mockEmailSender := infra.EmailSenderMock{
			Sendfn: func(opts infra.EmailOptions) error {
				return errors.New("error when sending email")
			},
		}

		service := NewService(domain.MockBookRepository{}, mockEmailSender)
	})
}
