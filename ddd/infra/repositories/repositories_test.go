package repositories

import (
	"errors"
	"testing"

	"github.com/fabiorodrigues/gostructure/ddd/domain/entities"
	"github.com/fabiorodrigues/gostructure/ddd/infra"
)

func TestBookRepository(t *testing.T) {
	t.Run("should return an error when receiving an error from databasedriver -getbyid", func(t *testing.T) {
		dbMock := infra.DatabaseProviderMock{
			GetByIDfn: func(collection string, id string) ([]byte, error) {
				return nil, errors.New("error when connecting to database")
			},
		}
		bookRepo := NewBook(dbMock)
		_, error := bookRepo.GetByID("123")

		if error == nil {
			t.Fail()
		}
	})

	t.Run("should return an error when receiving and error from databasedriver -getbyid", func(t *testing.T) {
		dbMock := infra.DatabaseProviderMock{
			GetByIDfn: func(collection string, id string) ([]byte, error) {
				return nil, errors.New("error when connecting to database")
			},
		}
		bookRepo := NewBook(dbMock)
		_, error := bookRepo.GetByID("123")

		if error == nil {
			t.Fail()
		}
	})

	t.Run("should return an error when receiving and error from databasedriver -update", func(t *testing.T) {
		dbMock := infra.DatabaseProviderMock{
			Updatefn: func(collection string, entity interface{}) error {
				return errors.New("error when connecting to database")
			},
		}
		bookRepo := NewBook(dbMock)
		if err := bookRepo.SaveOrUpdate(entities.CreateFakeBook(1)); err == nil {
			t.Fail()
		}

	})

}
