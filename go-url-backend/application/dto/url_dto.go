package dto

import "time"

// UrlDTO is used to transfer URL data.
type UrlDTO struct {
	ID            uint   `json:"id"`
	OriginalURL   string `json:"original_url"`
	ShortCode     string `json:"short_code"`
	CreatedAt  time.Time `json:"created_at"`
}

