package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IshanSaha05/Web-Scrapper-Telegana/pkg/mongodb"
	"github.com/IshanSaha05/Web-Scrapper-Telegana/pkg/services"
	"github.com/gorilla/mux"
)

var mongoDBObject *mongodb.MongoDB

func GetMongoDBObject(object *mongodb.MongoDB) {
	mongoDBObject = object
}

func GetByConstituencyName(w http.ResponseWriter, r *http.Request) {

	// Getting the constituency name.
	params := mux.Vars(r)
	constituencyName := params["constituency_name"]

	// Getting the link for the particular constiuency name.
	url, err := services.GetLink_ByName(constituencyName)

	if err != nil {
		fmt.Println("Error while getting the link for the constituency name: \"", constituencyName, "\"")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Fetching the particular constituency site for the results.
	response, err := services.GetSite(url)

	if err != nil {
		fmt.Println("Error while fetching the site: \"", url, "\"")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Parse and Print results.
	mongoDBObject.GetMongoCollection(constituencyName)
	services.ParseInsertDBPrintResults(response, mongoDBObject)

}

func GetByConstituencyID(w http.ResponseWriter, r *http.Request) {

	// Getting the constituency id.
	params := mux.Vars(r)
	constituencyID := params["constituency_id"]

	// Getting the link for the particular constiuency name.
	url, consconstituencyName, err := services.GetLink_ByID(constituencyID)
	if err != nil {
		fmt.Println("Error while getting the link for the constituency id: \"", constituencyID, "\"")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Fetching the particular constituency site for the results.
	response, err := services.GetSite(url)

	if err != nil {
		fmt.Println("Error while fetching the site: \"", url, "\"")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Parse and Print results.
	mongoDBObject.GetMongoCollection(consconstituencyName)
	services.ParseInsertDBPrintResults(response, mongoDBObject)

}
