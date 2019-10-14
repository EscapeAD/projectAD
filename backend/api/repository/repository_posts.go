package repository

import "escapead/projectAD/backend/api/models"

// PostsRepository - interface
type PostsRepository interface {
	Save(models.Post) (models.Post, error)
	FindAll() ([]models.Post, error)
	FindByID(uint32) (models.Post, error)
	Update(uint32, models.Post) (int64, error)
	Delete(uint32) (int64, error)
}
