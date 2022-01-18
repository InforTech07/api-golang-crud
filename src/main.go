package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/infortech07/crudApi/src/commons"
)

func main() {
	commons.Migrate()
	router := mux.NewRouter()
	routes.SetPeronsRoutes(router)

	srv := http.Server{
		Addr:    ":9000",
		Handler: router,
	}
	log.Println("sevidor ejecutando")
	log.Println(srv.ListenAndServe())
}
