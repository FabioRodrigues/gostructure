package mongo

import (
	"encoding/json"
	"fmt"
)

//Client is a client of mongodb that is responsible to implement a `DatabaseProvider` interface.
//it could be easily changed to another database implementation like `postgres` just respecting
//the `DatabaseProvider` contract
type Client struct {
	address  string
	username string
	password string
}

//NewClient instantiates a new Client of `Mongodb`
func NewClient() Client {
	return Client{}
}

//GetByID is a method that returns a new document from database using its ID as a
//term for searching
func (c Client) GetByID(collection string, id string) ([]byte, error) {
	fmt.Printf("mongo driver tried to retrieve collection %s with id %s", collection, id)
	return []byte(`{"id":"someuuid", "title": "Book sample", "availablequantity": 1}`), nil
}

//Update is a method that updates an document in the database
func (c Client) Update(collection string, entity interface{}) error {
	item, _ := json.Marshal(entity)
	fmt.Printf("\nbook saved. %s", string(item))
	return nil
}
