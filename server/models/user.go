package models

import "time"

// User model struct
// buat model dengan draw io
type User struct {
	ID           int                   `json:"id" gorm:"primary_key:auto_increment"`
	Name         string                `json:"name" gorm:"type: varchar(255)"`
	Email        string                `json:"email" gorm:"type: varchar(255)"`
	Password     string                `json:"-" gorm:"type: varchar(255)"`
	Status       string                `json:"status" gorm:"type: varchar(255)"`
	Image        string                `json:"image" gorm:"type: varchar(255)"`
	Profile      ProfileResponse       `json:"profile"`
	Products     []ProductUserResponse `json:"products"`
	Transactions []TransactionUserResponse
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
