package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/infortech07/crudApi/src/controller"
)

func setPersonRoute(router *mux.Router) {
	subRoute := router.PathPrefix("/person/api").Subrouter()
	subRoute.HandleFunc("/all", controller.GetAll).Methods(http.MethodGet)
	subRoute.HandleFunc("/save", controller.Save).Methods(http.MethodPost)
	subRoute.HandleFunc("/delete/{id}", controller.Delete).Methods(http.MethodDelete)
	subRoute.HandleFunc("/find/{id}", controller.Get).Methods(http.MethodGet)
}
