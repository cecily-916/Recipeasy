package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// handleCollections handles both POST and GET requests
// this request is used to populate the drop down menu of collections and display the associated recipes to the collection

func createCollection(w http.ResponseWriter, r *http.Request) {

	var newCollection Collection

	_ = json.NewDecoder(r.Body).Decode(&newCollection)

	db.Create(&newCollection)
	db.Save(&newCollection)
	json.NewEncoder(w).Encode(&newCollection)
}

func getCollections(w http.ResponseWriter, r *http.Request) {
	var collections []Collection

	params := mux.Vars(r)

	db.Where("user_id = ?", params["userid"]).Preload("Recipes").Find(&collections)

	json.NewEncoder(w).Encode(&collections)

	if err != nil {
		return
	}
}

// func addRecipeToCollection(w http.ResponseWriter, r *http.Request) {

// }
