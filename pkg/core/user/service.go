package user

import (
	"time"

	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

// Service provides basic operations on User domain model.
type Service interface {
	// CRUD operations.
	Create(user *User) (*User, error)
	Get(userID string) (*User, error)
	Update(user *User) (*User, error)
	Delete(userID string) error

	// Other operations.

}

type service struct {
	userRepository Repository
}

// NewUserService creates an instance of the service for User domain model
// with all of the neccessary dependencies.
func NewUserService(userRepository Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) Create(user *User) (*User, error) {
	newUser := NewUser(newUserID(), user.StudentID, user.Username, user.Email, hashAndSalt([]byte(user.Password)), time.Now())

	return s.userRepository.Save(newUser)
}

func (s *service) Get(userID string) (*User, error) {
	return s.userRepository.Get(userID)
}

func (s *service) Update(user *User) (*User, error) {
	user.Password = hashAndSalt([]byte(user.Password))

	return s.userRepository.Update(user)
}

func (s *service) Delete(userID string) error {
	return s.userRepository.Delete(userID)
}

func newUserID() string {
	return ksuid.New().String()
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}
