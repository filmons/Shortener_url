package services

import (
	"github.com/filmons/go-url-backend/application/dto"
	"github.com/filmons/go-url-backend/application/mapper"
	"github.com/filmons/go-url-backend/domain/url"
)

type UrlService struct {
	repository url.Repository
}



func NewUrlService(repo url.Repository) *UrlService {
	return &UrlService{
		repository: repo,
	}
}

func (s *UrlService) GetAllUrls() ([]dto.UrlDTO, error) {
	todos, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return mapper.ToUrlDTOs(todos), nil
}

func (s *UrlService) GetUrlByID(id uint) (*dto.UrlDTO, error) {
    url, err := s.repository.GetByID(id)
    if err != nil {
        return nil, err
    }
    // Convert the domain object to DTO and return its address
    urlDTO := mapper.ToUrlDTO(*url)
    return &urlDTO, nil
}


func (s *UrlService) CreateUrl(dto *dto.UrlDTO) error {
    url := mapper.FromUrlDTO(*dto) // Use the FromUrlDTO function to convert UrlDTO to url.Url
    return s.repository.Create(&url)
}


// func (s *UrlService) Update(id uint, url *url.Url) error {
//     // Call the repository's Update method with the ID and URL object
//     if err := s.repository.Update(id, url); err != nil {
//         return err
//     }
//     return nil
// }
func (s *UrlService) Update(id uint, dto *dto.UrlDTO) error {
    // Convert DTO to domain model within the service layer
    domainUrl := mapper.FromUrlDTO(*dto)
    
    // Call the repository's Update method with the ID and URL object
    if err := s.repository.Update(id, &domainUrl); err != nil {
        return err
    }
    return nil
}




func (s *UrlService) DeleteUrl(id uint) error {
    return s.repository.Delete(id)
}