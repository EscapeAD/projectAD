package repository

import "escapead/projectAD/backend/api/models"

// UserRepository - interface
type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindByID(uint32) (models.User, error)
	// Update(uint32) (uint64, error)
	// Delete(uint32) (uint64, error)
}
