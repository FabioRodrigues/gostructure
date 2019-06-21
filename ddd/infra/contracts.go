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

//EmailSender ...
type EmailSender interface {
	Send(opts EmailOptions) error
}

//EmailOptions ...
type EmailOptions struct {
	From   string
	To     []string
	Body   string
	IsHTML bool
}
