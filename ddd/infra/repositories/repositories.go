package repositories

import (
	"encoding/json"

	"github.com/fabiorodrigues/gostructure/ddd/domain/entities"
	"github.com/fabiorodrigues/gostructure/ddd/infra"
)

//Book is a repository of `Book` entity
//Everything necessary to interact with database to save,
//edit and more related to a `Book` entity is concentrated here
type Book struct {
	db infra.DatabaseProvider
}

//NewBook instantiates a new instance of a `BookRepository`
func NewBook(db infra.DatabaseProvider) Book {
	return Book{
		db: db,
	}
}

//GetByID returns a `Book` that has the related ID
func (b Book) GetByID(ID string) (entities.Book, error) {
	result, err := b.db.GetByID("book", ID)
	if err != nil {
		return entities.Book{}, err
	}
	book := entities.Book{}
	if err := json.Unmarshal(result, &book); err != nil {
		return entities.Book{}, err
	}
	return book, nil
}

//SaveOrUpdate Saves or Updates an existing `Book`
func (b Book) SaveOrUpdate(book entities.Book) error {
	return b.db.Update("book", book)
}
