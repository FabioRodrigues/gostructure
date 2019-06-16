package domain

//ReservationService ...
type ReservationService interface {
	Reserve(bookID string) error
}
