package user

// Repository provides access to the User store.
type Repository interface {
	// CRUD operations.
	Save(user *User) (*User, error)
	Get(userID string) (*User, error)
	Update(user *User) (*User, error)
	Delete(userID string) error

	// Other operations.
}
