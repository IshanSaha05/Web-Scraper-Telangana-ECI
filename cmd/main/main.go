package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IshanSaha05/Web-Scrapper-Telegana/pkg/config"
	"github.com/IshanSaha05/Web-Scrapper-Telegana/pkg/controllers"
	"github.com/IshanSaha05/Web-Scrapper-Telegana/pkg/mongodb"
	"github.com/IshanSaha05/Web-Scrapper-Telegana/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Search Telengana State Election Results Constituency Wise.")
	fmt.Println("---------------------------------------------------------------------")

	fmt.Println("Connecting to MongoDB")

	var mongoDBObject mongodb.MongoDB

	err := mongoDBObject.GetMongoClient()

	if err != nil {
		fmt.Println("Error while connecting to MongoDB Client.")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	err = mongoDBObject.GetMongoDatabase(config.DefaultDatabaseName)

	if err != nil {
		fmt.Println("Error while creating database with name \"", config.DefaultDatabaseName, "\".")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	controllers.GetMongoDBObject(&mongoDBObject)

	router := mux.NewRouter()

	routes.Routes(router)

	fmt.Println("Starting the server at localhost:8080.")
	err = http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println("Error encountered while starting the server.")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
