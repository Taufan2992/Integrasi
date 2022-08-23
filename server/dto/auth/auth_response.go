package authdto

type LoginResponse struct {
	ID     int    `gorm:"type: varchar(255)" json:"id"`
	Name   string `gorm:"type: varchar(255)" json:"name"`
	Email  string `gorm:"type: varchar(255)" json:"email"`
	Token  string `gorm:"type: varchar(255)" json:"token"`
	Status string `gorm:"type: varchar(255)" json:"status"`
	Image  string `json:"image" gorm:"type: varchar(255)"`
}
