package user

import (
	"errors"
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
	GetUserIDByUsername(username string) (string, error)
	CheckLibrarian(username string) (bool, error)
	AddFine(userID string, fine uint32) (uint32, error)
	GetTotalFine(userID string) (uint32, error)
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
	newUser := NewUser(newUserID(), user.StudentID, user.Role, user.Username, user.Email, hashAndSalt([]byte(user.Password)), user.TotalFine, time.Now())

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

func (s *service) GetUserIDByUsername(username string) (string, error) {
	return s.userRepository.GetIDByUsername(username)
}

func (s *service) CheckLibrarian(username string) (bool, error) {
	userID, err := s.GetUserIDByUsername(username)
	if err != nil {
		return false, err
	}

	role, err := s.userRepository.CheckLibrarian(userID)
	if err != nil {
		return false, err
	}

	if role != "librarian" {
		return false, errors.New("You are not authorized as a librarian to perform this action")
	}

	return true, nil
}

func (s *service) AddFine(userID string, fine uint32) (uint32, error) {
	currentTotalFine, err := s.GetTotalFine(userID)
	if err != nil {
		return 0, err
	}

	totalAddedFine := currentTotalFine + fine

	return totalAddedFine, s.userRepository.AddFine(userID, totalAddedFine)
}

func (s *service) GetTotalFine(userID string) (uint32, error) {
	return s.userRepository.GetTotalFine(userID)
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
