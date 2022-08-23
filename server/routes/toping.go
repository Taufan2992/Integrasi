package routes

import (
	"golang/handlers"
	"golang/pkg/middleware"
	"golang/pkg/mysql"
	"golang/repositories"

	"github.com/gorilla/mux"
)

// Create UserRoutes function here ...
func TopingRoutes(r *mux.Router) {
	topingRepository := repositories.RepositoryToping(mysql.DB)
	h := handlers.HandlerToping(topingRepository)

	r.HandleFunc("/topings", h.FindTopings).Methods("GET")
	r.HandleFunc("/toping/{id}", h.GetToping).Methods("GET")
	// r.HandleFunc("/toping", h.CreateToping).Methods("POST")
	r.HandleFunc("/toping", middleware.Auth(middleware.UploadFile(h.CreateToping))).Methods("POST")
	r.HandleFunc("/toping/{id}", middleware.Auth(middleware.UploadFile(h.UpdateToping))).Methods("PATCH")
	r.HandleFunc("/toping/{id}", middleware.Auth(h.DeleteToping)).Methods("DELETE")
}
