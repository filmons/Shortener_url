
package mapper

import (
    "github.com/filmons/go-url-backend/application/dto"
    "github.com/filmons/go-url-backend/domain/url"
)

// ToUrlDTO maps Url domain model to UrlDTO.
func ToUrlDTO(t url.Url) dto.UrlDTO {
    return dto.UrlDTO{
        ID:           t.ID,
        OriginalURL:  t.OriginalURL,
        ShortCode:    t.ShortCode,
        CreatedAt: t.CreatedAt,
    }
}

// ToUrlDTOs maps slice of Urls to slice of UrlDTOs.
func ToUrlDTOs(ts []url.Url) []dto.UrlDTO {
    dtos := make([]dto.UrlDTO, len(ts))
    for i, t := range ts {
        dtos[i] = ToUrlDTO(t)
    }
    return dtos
}


func FromUrlDTO(dto dto.UrlDTO) url.Url {
    return url.Url{
        ID:           dto.ID,
        OriginalURL:  dto.OriginalURL,
        ShortCode:    dto.ShortCode,
        CreatedAt: dto.CreatedAt,
    }
}