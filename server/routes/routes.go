package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	TodoRoutes(r)
	AuthRoutes(r)
	UserRoutes(r)
	ProductRoutes(r)
	TopingRoutes(r)
	TransactionRoutes(r)
	CartRoutes(r)
}
