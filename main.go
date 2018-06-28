package main

import (
	"StellarConnection/common"
	"net/http"
	"StellarConnection/routers"
	"log"
)

func main() {
	common.StartUp()

	router := routers.InitRouters()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}

	log.Println("listerning...")

	server.ListenAndServe()
}
