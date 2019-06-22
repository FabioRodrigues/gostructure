package infra

//DatabaseProvider is a generic provider responsible to make database operations
type DatabaseProvider interface {
	GetByID(collection string, id string) ([]byte, error)
	Update(collection string, entity interface{}) error
}

//EmailSender is a generic contract for sending emails
type EmailSender interface {
	Send(opts EmailOptions) error
}

//EmailOptions is just a `DTO` used for `EmailSender` to avoid receiving a lot of params
//necessary ti send an email
type EmailOptions struct {
	From   string
	To     []string
	Body   string
	IsHTML bool
}
