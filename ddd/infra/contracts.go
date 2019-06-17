package infra

import (
	"github.com/fabiorodrigues/gostructure/ddd/domain/entities"
)

//DatabaseProvider ...
type DatabaseProvider interface {
	GetByID(collection string, id string) ([]byte, error)
	Update(collection string, entity interface{}) error
}

//BookRepository ...
type BookRepository interface {
	GetByID(ID string) (entities.Book, error)
	SaveOrUpdate(book entities.Book) error
}
