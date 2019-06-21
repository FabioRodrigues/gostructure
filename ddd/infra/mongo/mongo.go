package mongo

import (
	"encoding/json"
	"fmt"
)

//Client ...
type Client struct {
	address  string
	username string
	password string
}

//NewClient ...
func NewClient() Client {
	return Client{}
}

//GetByID ...
func (c Client) GetByID(collection string, id string) ([]byte, error) {
	fmt.Printf("mongo driver tried to retrieve collection %s with id %s", collection, id)
	return []byte(`{"id":"someuuid", "title": "Book sample", "availablequantity": 1}`), nil
}

//Update ...
func (c Client) Update(collection string, entity interface{}) error {
	item, _ := json.Marshal(entity)
	fmt.Printf("\nbook saved. %s", string(item))
	return nil
}
