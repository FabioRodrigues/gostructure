package infra

//EmailSenderMock ...
type EmailSenderMock struct {
	Sendfn func(opts EmailOptions) error
}

//Send ...
func (s EmailSenderMock) Send(opts EmailOptions) error {
	if s.Sendfn != nil {
		return s.Sendfn(opts)
	}
	return nil
}
