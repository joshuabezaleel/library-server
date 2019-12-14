package user

import "time"

// User domain model.
type User struct {
	ID           string    `json:"id" db:"id"`
	StudentID    string    `json:"studentID" db:"student_id"`
	Role         string    `json:"role" db:"role"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	Password     string    `json:"password" db:"password"`
	TotalFine    uint32    `json:"totalFine" db:"total_fine"`
	RegisteredAt time.Time `json:"registeredAt" db:"registered_at"`
}

// NewUser creates a new instance of User domain model.
func NewUser(id string, studentID string, role string, username string, email string, password string, totalFine uint32, registeredAt time.Time) *User {
	return &User{
		ID:           id,
		StudentID:    studentID,
		Role:         role,
		Username:     username,
		Email:        email,
		Password:     password,
		TotalFine:    totalFine,
		RegisteredAt: registeredAt,
	}
}
