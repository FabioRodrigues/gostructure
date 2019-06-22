package entities

//CreateFakeBook is a func to create a mock od a book receiving an available quantity
//as initial value
func CreateFakeBook(availableQuantity int) Book {
	return NewBook("Faketitle", availableQuantity)
}
