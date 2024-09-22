package user

import (
	"fmt"

	"github.com/aormcuw/ecom/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	var user types.User

	// Use GORM to find the user by email
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, nil // Return other errors
	}

	return &user, nil // Return the user on success
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	var user types.User

	// Use GORM to find the user by id
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, nil // Return other errors
	}

	return &user, nil // Return the user on success
}

func (s *Store) CreateUser(user types.User) error {
	err := s.db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	return nil
}
