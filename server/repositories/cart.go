package repositories

import (
	"golang/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCarts() ([]models.Cart, error)
	GetCart(ID int) (models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart, ID int) (models.Cart, error)
	FindCartTopings(ID []int) ([]models.Toping, error)
	FindCartsTransaction(TrxID int) ([]models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCarts() ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Toping").Preload("Product.User").Find(&carts).Error

	return carts, err
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Toping").Preload("Product").First(&cart, ID).Error

	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}

func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}

func (r *repository) FindCartTopings(ID []int) ([]models.Toping, error) {
	var topings []models.Toping
	err := r.db.Debug().Find(&topings, ID).Error

	return topings, err
}

func (r *repository) FindCartsTransaction(TrxID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product").Preload("Toping").Debug().Find(&carts, "transaction_id = ?", TrxID).Error

	return carts, err
}
