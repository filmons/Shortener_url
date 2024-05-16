package dto

import "time"

type UserDTO struct {
    ID        uint      `json:"id"`
    Username  string    `json:"username"`
    Password  string    `json:"password"`
    Email     string    `json:"email"` // Adding email to the DTO
    CreatedAt time.Time `json:"created_at"`
}
type LoginDTO struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
