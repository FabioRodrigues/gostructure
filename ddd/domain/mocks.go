package domain

import "github.com/fabiorodrigues/gostructure/ddd/domain/entities"

//MockBookRepository ...
type MockBookRepository struct {
	GetByIDfn      func(ID string) (entities.Book, error)
	SaveOrUpdatefn func(book entities.Book) error
}

//GetByID ...
func (r MockBookRepository) GetByID(ID string) (entities.Book, error) {
	if r.GetByIDfn != nil {
		return r.GetByIDfn(ID)
	}
	return entities.CreateFakeBook(1), nil
}

func (r MockBookRepository) SaveOrUpdate(book entities.Book) error {
	if r.SaveOrUpdatefn != nil {
		return r.SaveOrUpdatefn(book)
	}
	return nil
}
