package models

import "time"

// User model struct
type Toping struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Title     string    `json:"title" gorm:"type: varchar(255)"`
	Price     int       `json:"price" gorm:"type: int"`
	Image     string    `json:"image" form:"image" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type TopingResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" gorm:"type: int"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
}

func (TopingResponse) TableName() string {
	return "topings"
}
