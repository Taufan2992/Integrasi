package repositories

import (
	"golang/models"

	"gorm.io/gorm"
)

// Import "dumbmerch/models", "gorm.io/gorm"

// Declare UserRepository interface here ...
type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
}

// Create RepositoryUser function here ...
func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

// Create FindUsers method here ...
func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Profile").Find(&users).Error

	return users, err
}

// Create GetUser method here ...
func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Profile").Preload("Products").First(&user, ID).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Debug().Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
