package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/filmons/go-url-backend/application/services"
    "github.com/filmons/go-url-backend/pkg/utils"
)

type UserController struct {
    Service *services.UserService
}

func NewUserController(s *services.UserService) *UserController {
    return &UserController{
        Service: s,
    }
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := c.Service.GetAllUsers()
    if err != nil {
        utils.Logger.Println("Error getting users:", err)
        http.Error(w, "Failed to get users", http.StatusInternalServerError)
        return
    }
    response, _ := json.Marshal(users)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}


// func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
//     var userDto dto.UserDTO
//     if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
//         http.Error(w, "Invalid request body", http.StatusBadRequest)
//         return
//     }

//     hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
//     if err != nil {
//         http.Error(w, "Failed to hash password", http.StatusInternalServerError)
//         return
//     }

//     user := &user.User{
//         Username:  userDto.Username,
//         Password:  string(hashedPassword),
//         Email:     userDto.Email,
//         CreatedAt: userDto.CreatedAt,
//     }

//     if err := uc.UserService.RegisterUser(user); err != nil {
//         http.Error(w, "Failed to register user", http.StatusInternalServerError)
//         return
//     }

//     w.WriteHeader(http.StatusCreated)
//     w.Write([]byte("User registered successfully"))
// }



// Login handles user login.
// func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
//     // Parse request body to get user login credentials
//     var loginData dto.LoginDTO
//     err := json.NewDecoder(r.Body).Decode(&loginData)
//     if err != nil {
//         http.Error(w, "Invalid request body", http.StatusBadRequest)
//         return
//     }

//     // Retrieve user from the database based on username or email
//     user, err := uc.UserService.GetUserByEmail(loginData.Username)
//     if err != nil {
//         http.Error(w, "Invalid username or password", http.StatusUnauthorized)
//         return
//     }

//     // Compare hashed password with the provided password
//     err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
//     if err != nil {
//         http.Error(w, "Invalid username or password", http.StatusUnauthorized)
//         return
//     }

//     // Generate JWT token upon successful login
// }
