package cartdto

type CartRequest struct {
	ProductID     int   `json:"product_id"`
	TransactionID int   `json:"transaction_id"`
	Qty           int   `json:"qty" form:"qty"`
	SubAmount     int   `json:"subamount"`
	TopingID      []int `json:"toping_id" form:"toping_id" gorm:"-"`
	UserID        int   `json:"-" gorm:"type: int"`
}
