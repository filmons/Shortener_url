
package url
import (
    "time"
)

// Url represents the core domain model for a URL item.
type Url struct {
    ID            uint       `gorm:"primaryKey"`
    CreatedAt     time.Time
    UpdatedAt     time.Time
    OriginalURL   string     // Corresponds to the 'original_url' column in the database.
    ShortCode     string     // Corresponds to the 'short_code' column in the database.
}
