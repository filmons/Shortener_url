package services

import (
	"github.com/filmons/go-url-backend/application/dto"
	"github.com/filmons/go-url-backend/domain/user"
	"github.com/filmons/go-url-backend/application/mapper"
	// "golang.org/x/crypto/bcrypt"
)

// UserService provides user-related services.
type UserService struct {
    repository user.Repository
}

// NewUserService creates a new instance of UserService.
func NewUserService(repo user.Repository) *UserService {
    return &UserService{repository: repo}
}


func (s *UserService) GetAllUsers() ([]dto.UserDTO, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return mapper.ToUserDTOs(users), nil
}

// RegisterUser registers a new user.
// func (s *UserService) RegisterUser(u *user.User) error {
//     // Hash password
//     hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
//     if err != nil {
//         return err
//     }
//     // Assign hashed password to user
//     u.Password = string(hashedPassword)
//     // Call repository to create user
//     return s.repository.CreateUser(u)
// }
