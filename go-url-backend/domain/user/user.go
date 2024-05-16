

package user

import (
    "time"
)
// User represents the user domain model.
// type User struct {
//     ID         uint       `gorm:"primaryKey"` // Unique identifier for the user
//     Username   string     `gorm:"unique"`     // Username of the user, corresponds to the 'username' column in the database
//     Password   string                       // Password of the user, corresponds to the 'password_hash' column in the database
//     CreatedAt  time.Time  // Creation timestamp of the user, corresponds to the 'created_at' column in the database
// }
type User struct {
    ID        uint
    Username  string
    Password  string
    Email     string
    CreatedAt time.Time
}