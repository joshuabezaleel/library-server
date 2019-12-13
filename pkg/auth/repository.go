package auth

// Repository provides access to the store by Auth service.
type Repository interface {
	GetPassword(username string) (string, error)
}
