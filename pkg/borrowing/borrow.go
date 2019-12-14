package borrowing

import (
	"time"
)

// Borrow domain model.
type Borrow struct {
	ID         string    `json:"id" db:"id"`
	UserID     string    `json:"userID" db:"user_id"`
	BookCopyID string    `json:"bookCopyID" db:"bookcopy_id"`
	Fine       uint32    `json:"fine" db:"fine"`
	BorrowedAt time.Time `json:"borrowedAt" db:"borrowed_at"`
	DueDate    time.Time `json:"dueDate" db:"due_date"`
}

// NewBorrow creates a new instance of Borrow domain model.
func NewBorrow(id string, userID string, bookCopyID string, borrowedAt time.Time, dueDate time.Time, fine uint32) *Borrow {
	return &Borrow{
		ID:         id,
		UserID:     userID,
		BookCopyID: bookCopyID,
		Fine:       fine,
		BorrowedAt: borrowedAt,
		DueDate:    dueDate,
	}
}
