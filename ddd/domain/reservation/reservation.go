package reservation

import (
	"encoding/json"

	"github.com/fabiorodrigues/gostructure/ddd/domain"
	"github.com/fabiorodrigues/gostructure/ddd/infra"
)

//Service ...
type Service struct {
	database infra.DatabaseProvider
}

//Reserve ...
func (s Service) Reserve(bookID string) error {

	item, err := s.database.GetByID("book", bookID)
	if err != nil {
		return err
	}
	var book domain.Book
	err = json.Unmarshal(item, &book)
	if err != nil {
		return err
	}
	if err := book.Reserve(); err != nil {
		return err
	}

	if err := s.database.Update("book", book); err != nil {
		return err
	}

	return nil
}
