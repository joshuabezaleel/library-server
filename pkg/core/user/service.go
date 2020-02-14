package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	util "github.com/joshuabezaleel/library-server/pkg"
)

// Errors definition.
var (
	ErrCreateUser = errors.New("Error creating User")
	ErrGetUser    = errors.New("Error retrieving User")
	ErrUpdateUser = errors.New("Error updating User")
	ErrDeleteUser = errors.New("Error deleting User")

	ErrGetUserIDByUsername = errors.New("Error retrieving User ID")
	ErrGetRole             = errors.New("Error retrieving User's role")
	ErrAddFine             = errors.New("Error adding fine to User")
	ErrGetTotalFine        = errors.New("Error retrieving User's total fine")
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
	GetRole(username string) (string, error)
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
	var newUser *User

	if user.ID == "" {
		newUser = NewUser(util.NewID(), user.StudentID, user.Role, user.Username, user.Email, hashAndSalt(user.Password), user.TotalFine, time.Now())
	} else {
		newUser = NewUser(user.ID, user.StudentID, user.Role, user.Username, user.Email, hashAndSalt(user.Password), user.TotalFine, time.Now())
	}

	newUser, err := s.userRepository.Save(newUser)
	if err != nil {
		return nil, ErrCreateUser
	}

	return newUser, nil
}

func (s *service) Get(userID string) (*User, error) {
	user, err := s.userRepository.Get(userID)
	if err != nil {
		return nil, ErrGetUser
	}

	return user, nil
}

func (s *service) Update(user *User) (*User, error) {
	user.Password = hashAndSalt(user.Password)

	user, err := s.userRepository.Update(user)
	if err != nil {
		return nil, ErrUpdateUser
	}

	return user, nil
}

func (s *service) Delete(userID string) error {
	err := s.userRepository.Delete(userID)
	if err != nil {
		return ErrDeleteUser
	}

	return nil
}

func (s *service) GetUserIDByUsername(username string) (string, error) {
	userID, err := s.userRepository.GetIDByUsername(username)
	if err != nil {
		return "", ErrGetUserIDByUsername
	}

	return userID, nil
}

func (s *service) GetRole(username string) (string, error) {
	userID, err := s.GetUserIDByUsername(username)
	if err != nil {
		return "", ErrGetUserIDByUsername
	}

	role, err := s.userRepository.GetRole(userID)
	if err != nil {
		return "", ErrGetRole
	}

	return role, nil
}

func (s *service) AddFine(userID string, fine uint32) (uint32, error) {
	currentTotalFine, err := s.GetTotalFine(userID)
	if err != nil {
		return 0, ErrGetTotalFine
	}

	totalAddedFine := currentTotalFine + fine

	err = s.userRepository.AddFine(userID, totalAddedFine)
	if err != nil {
		return 0, ErrAddFine
	}

	return totalAddedFine, nil
}

func (s *service) GetTotalFine(userID string) (uint32, error) {
	totalFine, err := s.userRepository.GetTotalFine(userID)
	if err != nil {
		return 0, ErrGetTotalFine
	}

	return totalFine, nil
}

func hashAndSalt(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}
