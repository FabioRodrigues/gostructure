package sendgrid

import (
	"fmt"

	"github.com/fabiorodrigues/gostructure/ddd/infra"
)

//Client is a client of sendgrid responsible for sending email
//with `Sendgrid`
type Client struct {
	secretKey string
}

//NewClient instantiates a new client of `Sendgrid`
func NewClient(secretKey string) Client {
	return Client{
		secretKey: secretKey,
	}
}

//Send is responsible for sending email using `Sendgrid's` provider
func (c Client) Send(opts infra.EmailOptions) error {
	for _, to := range opts.To {
		fmt.Println("\n\nsending email to", to)
	}
	return nil
}
