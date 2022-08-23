package models

import "time"

// User model struct
type Transaction struct {
	ID        int64                `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount    int                  `json:"amount"`
	Status    string               `json:"status"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

type TransactionUserResponse struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	Amount int `json:"amount"`
}

type TransactionMail struct {
	ID        int64                `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID int                  `json:"product_id"`
	Product   ProductUserResponse  `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount    int                  `json:"amount"`
	BuyerID   int                  `json:"buyer_id"`
	Buyer     UsersProfileResponse `json:"buyer"`
	SellerID  int                  `json:"seller_id"`
	Seller    UsersProfileResponse `json:"seller"`
	Status    string               `json:"status"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

type TransactionResponse struct {
	UserID int `json:"user_id"`
	Amount int `json:"amount"`
}

func (TransactionUserResponse) TableName() string {
	return "transactions"
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
