package mapper

import (
    "github.com/filmons/go-url-backend/application/dto"
    "github.com/filmons/go-url-backend/domain/user"
)

// ToUserDTO maps a user domain model to a user DTO.
func ToUserDTO(u user.User) dto.UserDTO {
    return dto.UserDTO{
        ID:       u.ID,
        Username: u.Username,
        // You might want to exclude the password from the DTO for security reasons.
        // Password: u.Password,
        Password:  u.Password,
        Email:     u.Email,
        CreatedAt: u.CreatedAt,
    }
}
func ToUserDTOs(ts []user.User) []dto.UserDTO {
    dtos := make([]dto.UserDTO, len(ts))
    for i, u := range ts {
        dtos[i] = ToUserDTO(u)
    }
    return dtos
}