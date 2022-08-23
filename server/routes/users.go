package routes

import (
	"golang/handlers"
	"golang/pkg/middleware"
	"golang/pkg/mysql"
	"golang/repositories"

	"github.com/gorilla/mux"
)

// Create UserRoutes function here ...
func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/user", h.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", middleware.Auth(middleware.UploadFile(h.UpdateUser))).Methods("PATCH")
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")

}
