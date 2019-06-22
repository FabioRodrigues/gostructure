package domain

import "github.com/fabiorodrigues/gostructure/ddd/domain/entities"

//MockBookRepository is a mock of `BookRepository`
type MockBookRepository struct {
	GetByIDfn      func(ID string) (entities.Book, error)
	SaveOrUpdatefn func(book entities.Book) error
}

//GetByID is a mock of the method `GetByID` from `BookRepository`.
//You can override it to you own implementation to simulate
//a new scenario
func (r MockBookRepository) GetByID(ID string) (entities.Book, error) {
	if r.GetByIDfn != nil {
		return r.GetByIDfn(ID)
	}
	return entities.CreateFakeBook(1), nil
}

//SaveOrUpdate is a mock of the method `SaveOrUpdate` from `BookRepository`.
//You can override it to you own implementation to simulate
//a new scenario
func (r MockBookRepository) SaveOrUpdate(book entities.Book) error {
	if r.SaveOrUpdatefn != nil {
		return r.SaveOrUpdatefn(book)
	}
	return nil
}
