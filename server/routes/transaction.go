package routes

import (
	"golang/handlers"
	"golang/pkg/middleware"
	"golang/pkg/mysql"
	"golang/repositories"

	"github.com/gorilla/mux"
)

// Create UserRoutes function here ...
func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transactionid", middleware.Auth(h.UpdateTransaction)).Methods("PATCH")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.DeleteTransaction)).Methods("DELETE")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}
