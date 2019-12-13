package bookcopy

import (
	"time"
)

// BookCopy domain model.
type BookCopy struct {
	ID        string    `json:"id" db:"id"`
	Barcode   string    `json:"barcode" db:"barcode"`
	BookID    string    `json:"bookID" db:"book_id"`
	Condition string    `json:"condition" db:"condition"`
	AddedAt   time.Time `json:"addedAt" db:"added_at"`
}

// NewBookCopy creates a new instance of BookCopy domain model.
func NewBookCopy(id string, barcode string, bookID string, condition string, addedAt time.Time) *BookCopy {
	return &BookCopy{
		ID:        id,
		Barcode:   barcode,
		BookID:    bookID,
		Condition: condition,
		AddedAt:   addedAt,
	}
}
