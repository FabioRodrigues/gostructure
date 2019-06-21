package infra

//DatabaseProvider ...
type DatabaseProvider interface {
	GetByID(collection string, id string) ([]byte, error)
	Update(collection string, entity interface{}) error
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
