package transactiondto

type TransactionRequest struct {
	UserID int    `json:"user_id"`
	Status string `json:"status"`
	Amount int    `json:"amount"`
}

type UpdateTransactionRequest struct {
	UserID int    `json:"user_id"`
	Status string `json:"status"`
	Amount int    `json:"amount"`
}
