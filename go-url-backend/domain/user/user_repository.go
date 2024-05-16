package user

// Repository defines the expected behaviour from a todo repository.
type Repository interface {
	GetAll() ([]User, error)
	// CreateUser(user *User) error
    // GetUserByEmail(email string) (*User, error)
}
