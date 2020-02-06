package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	util "github.com/joshuabezaleel/library-server/pkg"
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

	return s.userRepository.Save(newUser)
}

func (s *service) Get(userID string) (*User, error) {
	return s.userRepository.Get(userID)
}

func (s *service) Update(user *User) (*User, error) {
	user.Password = hashAndSalt(user.Password)

	return s.userRepository.Update(user)
}

func (s *service) Delete(userID string) error {
	return s.userRepository.Delete(userID)
}

func (s *service) GetUserIDByUsername(username string) (string, error) {
	return s.userRepository.GetIDByUsername(username)
}

func (s *service) GetRole(username string) (string, error) {
	userID, err := s.GetUserIDByUsername(username)
	if err != nil {
		return "", err
	}

	role, err := s.userRepository.GetRole(userID)
	if err != nil {
		return "", err
	}

	return role, nil
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

func hashAndSalt(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}
