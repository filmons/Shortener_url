package url

// Repository defines the expected behaviour from a todo repository.
// type Repository interface {
// 	GetAll() ([]Url, error)
// }


// Repository defines the expected behavior from a URL repository.
type Repository interface {
    GetAll() ([]Url, error)
    GetByID(id uint) (*Url, error)
    Create(url *Url) error
    Update(id uint, url *Url) error
    Delete(id uint) error
}

