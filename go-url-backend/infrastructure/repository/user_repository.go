

package repository

import (
    "log"
    "github.com/filmons/go-url-backend/domain/user"
    "gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB 
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}
// CreateUser saves a new user to the database.
func (r *UserRepository) GetAll() ([]user.User, error) {
	var urls []user.User
	result := r.DB.Find(&urls)
	if result.Error != nil {
		log.Printf("Failed to fetch urls: %v", result.Error)
		return nil, result.Error
	}
	log.Println("Fetched all urls successfully")
	return urls, nil
}

// func (repo *UserRepository) CreateUser(u *user.User) error {
//     result := repo.db.Create(u)
//     return result.Error
// }
