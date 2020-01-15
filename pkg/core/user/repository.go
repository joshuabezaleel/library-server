package user

// Repository provides access to the User store.
type Repository interface {
	// CRUD operations.
	Save(user *User) (*User, error)
	Get(userID string) (*User, error)
	Update(user *User) (*User, error)
	Delete(userID string) error

	// Other operations.
	GetIDByUsername(username string) (string, error)
	GetRole(userID string) (string, error)
	AddFine(userID string, fine uint32) error
	GetTotalFine(userID string) (uint32, error)
}
