
package repository
import (
	"log"
	"github.com/filmons/go-url-backend/domain/url"
	"gorm.io/gorm"
)

type UrlRepository struct {
	DB *gorm.DB 
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{DB: db}
}


func (r *UrlRepository) GetAll() ([]url.Url, error) {
	var urls []url.Url
	result := r.DB.Find(&urls)
	if result.Error != nil {
		log.Printf("Failed to fetch urls: %v", result.Error)
		return nil, result.Error
	}
	log.Println("Fetched all urls successfully")
	return urls, nil
}


func (r *UrlRepository) GetByID(id uint) (*url.Url, error) {
    var url url.Url
    if err := r.DB.First(&url, id).Error; err != nil {
        return nil, err
    }
    return &url, nil
}

func (r *UrlRepository) Create(url *url.Url) error { // Change parameter to *url.Url
    result := r.DB.Create(url) // Pass the pointer directly
    if result.Error != nil {
        log.Printf("Failed to create url: %v", result.Error)
        return result.Error
    }
    log.Println("Created url successfully")
    return nil
}


// func (r *UrlRepository) Update(url *url.Url) error {
//     if err := r.DB.Save(url).Error; err != nil {
//         return err
//     }
//     return nil
// }
func (r *UrlRepository) Update(id uint, updatedUrl *url.Url) error {
    // Update the URL with the provided ID
    if err := r.DB.Model(&url.Url{}).Where("id = ?", id).Updates(updatedUrl).Error; err != nil {
        return err
    }
    return nil
}



func (r *UrlRepository) Delete(id uint) error {
    // Delete the URL with the provided ID
    if err := r.DB.Delete(&url.Url{}, id).Error; err != nil {
        return err
    }
    return nil
}
