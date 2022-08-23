package main

import (
	"fmt"
	"golang/database"
	"golang/pkg/mysql"
	"golang/routes"
	"net/http"

	"github.com/gorilla/handlers" // import this package ...
	"github.com/gorilla/mux"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))) // add this code

	// Setup allowed Header, Method, and Origin for CORS here ...
	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = "5000"
	fmt.Println("server running localhost:" + port)

	// Embed the setup allowed in 2 parameter here ...
	http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
