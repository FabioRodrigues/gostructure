package reservation

import (
	"github.com/fabiorodrigues/gostructure/infra"
)

//Service ...
type Service struct {
	database infra.DatabaseProvider
}

//Reserve ...
func (s Service) Reserve(bookID string) error {
	
	book := database.GetByID(bookID)

	return nil
}
