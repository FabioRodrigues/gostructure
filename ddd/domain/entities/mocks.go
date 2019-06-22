package entities

//CreateFakeBook ...
func CreateFakeBook(availableQuantity int) Book {
	return NewBook("Faketitle", availableQuantity)
}
