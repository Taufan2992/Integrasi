package models

import "time"

// User model struct
type Product struct {
	ID        int                  `json:"id" gorm:"PRIMARY_KEY"`
	Title     string               `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price     int                  `json:"price" form:"price" gorm:"type: int"`
	Image     string               `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID    int                  `json:"user_id" form:"user_id"`
	User      UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

type ProductUserResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price  int    `json:"price" form:"price" gorm:"type: int"`
	Image  string `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID int    `json:"user_id" form:"user_id"`
}

type ProductResponse struct {
	ID     int                  `json:"id"`
	Title  string               `json:"title"`
	Price  int                  `json:"price"`
	Image  string               `json:"image"`
	UserID int                  `json:"-"`
	User   UsersProfileResponse `json:"user"`
}

// type ProductTransactionResponse struct {
// 	ID       int      `json:"id"`
// 	Title    string   `json:"title"`
// 	Price    int      `json:"price"`
// 	Image    string   `json:"image"`
// 	UserID   int      `json:"-"`
// 	Toping   []Toping `json:"toping" gorm:"many2many:product_topings"`
// 	TopingID []int    `json:"-" form:"toping_id" gorm:"-"`
// }

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}

// func (ProductTransactionResponse) TableName() string {
// 	return "products"
// }
