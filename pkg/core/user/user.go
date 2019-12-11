package user

import "time"

// User domain model.
type User struct {
	ID           string    `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	Password     string    `json:"password" db:"password"`
	RegisteredAt time.Time `json:"registeredAt" db:"registered_at"`
}

// NewUser creates a new instance of User domain model.
func NewUser(id string, username string, email string, password string, registeredAt time.Time) *User {
	return &User{
		ID:           id,
		Username:     username,
		Email:        email,
		Password:     password,
		RegisteredAt: registeredAt,
	}
}
