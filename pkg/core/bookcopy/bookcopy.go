package bookcopy

import (
	"time"
)

// BookCopy domain model.
type BookCopy struct {
	ID      string    `json:"id" db:"id"`
	Barcode string    `json:"barcode" db:"barcode"`
	BookID  string    `json:"bookID" db:"book_id"`
	AddedAt time.Time `json:"addedAt" db:"added_at"`
}

// NewBookCopy creates a new instance of BookCopy domain model.
func NewBookCopy(id string, barcode string, bookID string, addedAt time.Time) *BookCopy {
	return &BookCopy{
		ID:      id,
		Barcode: barcode,
		BookID:  bookID,
		AddedAt: addedAt,
	}
}
