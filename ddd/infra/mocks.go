package infra

//EmailSenderMock is a mock of an `Emailsender`
type EmailSenderMock struct {
	Sendfn func(opts EmailOptions) error
}

//Send is a mock method that allows to be overrode for sending email
func (s EmailSenderMock) Send(opts EmailOptions) error {
	if s.Sendfn != nil {
		return s.Sendfn(opts)
	}
	return nil
}

//DatabaseProviderMock is a mock of a `DatabaseProvider`
type DatabaseProviderMock struct {
	GetByIDfn func(collection string, id string) ([]byte, error)
	Updatefn  func(collection string, entity interface{}) error
}

//GetByID is a mock method that allows to be overrode for getting
//an entity By collection and ID
func (d DatabaseProviderMock) GetByID(collection string, id string) ([]byte, error) {
	if d.GetByIDfn != nil {
		return d.GetByIDfn(collection, id)
	}
	return []byte("{}"), nil
}

//Update is a mock method that allows to be overrode for update
//an entity in the database
func (d DatabaseProviderMock) Update(collection string, entity interface{}) error {
	if d.Updatefn != nil {
		return d.Updatefn(collection, entity)
	}
	return nil
}
