package sendgrid

import (
	"fmt"

	"github.com/fabiorodrigues/gostructure/ddd/infra"
)

//Client ...
type Client struct {
	secretKey string
}

//NewClient ...
func NewClient(secretKey string) Client {
	return Client{
		secretKey: secretKey,
	}
}

//Send ...
func (c Client) Send(opts infra.EmailOptions) error {
	for _, to := range opts.To {
		fmt.Println("\n\nsending email to", to)
	}
	return nil
}
