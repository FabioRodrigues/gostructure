package infra

//DatabaseProvider ...
type DatabaseProvider interface {
	GetByID(collection string, id string) []byte
}
