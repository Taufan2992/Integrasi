package models

import "time"

// User model struct
type Profile struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	UserID    int                  `json:"user_id"`
	Address   string               `json:"address" gorm:"type: varchar(255)"`
	Phone     string               `json:"phone" gorm:"type: varchar(255)"`
	User      UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

// for association relation with another table (user)
type ProfileResponse struct {
	Address string `json:"address"`
	Phone   string `json:"phone"`
	UserID  int    `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
