package routers

import (
	"github.com/gorilla/mux"
	"StellarConnection/controllers"
)

// SetUserRoutes registers routes for user entity
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	//router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	return router
}
