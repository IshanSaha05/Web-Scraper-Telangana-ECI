package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IshanSaha05/Web_Scrapper_Election/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Search Telengana State Election Results Constituency Wise.")
	fmt.Println("---------------------------------------------------------------------")

	router := mux.NewRouter()

	routes.Routes(router)

	fmt.Println("Starting the server at localhost:8080.")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println("Error encountered while starting the server.")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
