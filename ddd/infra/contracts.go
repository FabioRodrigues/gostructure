package infra

//DatabaseProvider ...
type DatabaseProvider interface {
	GetByID(collection string, id string) ([]byte, error)
	Update(collection string, entity interface{}) error
}
