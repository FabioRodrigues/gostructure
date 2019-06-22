package entities

import (
	"testing"

	"github.com/ditointernet/go-assert"
)

func TestBook(t *testing.T) {
	t.Run("should set available quantity to zero when reserving the last book", func(t *testing.T) {
		book := CreateFakeBook(1)
		book.Reserve()
		assert.Equal(t, book.AvailableQuantity, 0)
	})

	t.Run("should return an error when there are not available books to reserve", func(t *testing.T) {
		book := CreateFakeBook(0)

		if err := book.Reserve(); err == nil {
			t.Fail()
		}
	})

	t.Run("should set available quantity to five when there are six available books", func(t *testing.T) {
		book := CreateFakeBook(6)
		book.Reserve()
		assert.Equal(t, book.AvailableQuantity, 5)
	})
}
