package productdto

type ProductRequest struct {
	Title    string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price    int    `json:"price" form:"price" gorm:"type: int"`
	Image    string `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID   int    `json:"user_id" gorm:"type: int"`
	TopingID int    `json:"toping_id" form:"toping_id" gorm:"type: int"`
}

type UpdateProductRequest struct {
	Title    string `json:"title"`
	Price    int    `json:"price"`
	Image    string `json:"image"`
	UserID   int    `json:"user_id"`
	TopingID int    `json:"toping_id"`
}
