package models

import "time"

// User model struct
type Cart struct {
	ID            int             `json:"id" gorm:"PRIMARY_KEY"`
	ProductID     int             `json:"product_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product       ProductResponse `json:"product"`
	TransactionID int             `json:"transaction_id"`
	Qty           int             `json:"qty" form:"qty"`
	SubAmount     int             `json:"subamount"`
	Toping        []Toping        `json:"toping" gorm:"many2many:cart_topings; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TopingID      []int           `json:"-" form:"toping_id" gorm:"-"`
	CreatedAt     time.Time       `json:"-"`
	UpdatedAt     time.Time       `json:"-"`
}

type CartResponse struct {
	ProductID     int `json:"product_id"`
	TransactionID int `json:"transaction_id"`
	Qty           int `json:"qty" form:"qty"`
	SubAmount     int `json:"subamount"`
}

func (CartResponse) TableName() string {
	return "carts"
}
