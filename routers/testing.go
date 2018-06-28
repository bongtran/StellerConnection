package routers

import (
	"github.com/gorilla/mux"
	"StellarConnection/controllers"
)

func SetTestingRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/testing/get", controllers.ShowSuccess).Methods("GET")
	router.HandleFunc("/testing", controllers.ShowError).Methods("POST")
	router.HandleFunc("/testing", controllers.ShowInternalError).Methods("PUT")
	return router
}
