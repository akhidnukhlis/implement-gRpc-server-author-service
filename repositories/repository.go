package repositories

import "github.com/jinzhu/gorm"

type Repository struct {
	Author Author
}

// NewRepository to setting services repositories
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Author: NewAuthor(db),
	}
}
