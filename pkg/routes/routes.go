package routes

import (
	"github.com/IshanSaha05/Web-Scrapper-Telegana/pkg/controllers"
	"github.com/gorilla/mux"
)

var Routes = func(router *mux.Router) {
	router.HandleFunc("/api/v1/telengana-constituency/name/{constituency_name}", controllers.GetByConstituencyName).Methods("GET")
	router.HandleFunc("/api/v1/telengana-constituency/id/{constituency_id}", controllers.GetByConstituencyID).Methods("GET")
}
