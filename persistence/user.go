package persistence

import (
	"github.com/jmoiron/sqlx"

	"github.com/joshuabezaleel/library-server/pkg/core/user"
)

type userRepository struct {
	DB *sqlx.DB
}

// NewUserRepository returns initialized implementations of the repository for
// User domain model.
func NewUserRepository(DB *sqlx.DB) user.Repository {
	return &userRepository{
		DB: DB,
	}
}

func (repo *userRepository) Save(user *user.User) (*user.User, error) {
	_, err := repo.DB.NamedExec("INSERT INTO users (id, student_id, role, username, email, password, total_fine, registered_at) VALUES (:id, :student_id, :role, :username, :email, :password, :total_fine, :registered_at)", user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *userRepository) Get(userID string) (*user.User, error) {
	user := user.User{}

	err := repo.DB.QueryRowx("SELECT * FROM users WHERE id=$1", userID).StructScan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) Update(user *user.User) (*user.User, error) {
	_, err := repo.DB.NamedExec("UPDATE users SET student_id=:student_id, role=:role, username=:username, email=:email, password=:password, total_fine=:total_fine WHERE id=:id", user)

	if err != nil {
		return nil, err
	}

	updatedUser, err := repo.Get(user.ID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (repo *userRepository) Delete(userID string) error {
	_, err := repo.DB.Exec("DELETE FROM users WHERE id=$1", userID)

	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) GetIDByUsername(username string) (string, error) {
	var userID string

	err := repo.DB.QueryRow("SELECT id FROM users WHERE username=$1", username).Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (repo *userRepository) CheckLibrarian(userID string) (string, error) {
	var role string

	err := repo.DB.QueryRow("SELECT role FROM users WHERE id=$1", userID).Scan(&role)
	if err != nil {
		return "", err
	}

	return role, nil
}
