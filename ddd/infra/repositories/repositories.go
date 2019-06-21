package repositories

import (
	"encoding/json"

	"github.com/fabiorodrigues/gostructure/ddd/domain/entities"
	"github.com/fabiorodrigues/gostructure/ddd/infra"
)

//Book ...
type Book struct {
	db infra.DatabaseProvider
}

//NewBook ...
func NewBook(db infra.DatabaseProvider) Book {
	return Book{
		db: db,
	}
}

//GetByID ...
func (b Book) GetByID(ID string) (entities.Book, error) {
	result, err := b.db.GetByID("book", ID)
	if err != nil {
		return entities.Book{}, err
	}
	book := entities.Book{}
	if err := json.Unmarshal(result, &book); err != nil {
		return entities.Book{}, nil
	}
	return book, nil
}

//SaveOrUpdate ...
func (b Book) SaveOrUpdate(book entities.Book) error {
	return b.db.Update("book", book)
}
