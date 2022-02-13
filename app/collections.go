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

func getCollectionsPerUser(w http.ResponseWriter, r *http.Request) {
	var collections []Collection
	// var user User
	params := mux.Vars(r)

	// db.Where("userfront_id =?", params["userfrontid"]).First(&user)
	db.Where("user_id = ? ", params["userfrontid"]).Preload("Recipes").Find(&collections)

	// json.NewEncoder(w).Encode(&user)
	json.NewEncoder(w).Encode(&collections)

	if err != nil {
		return
	}
}

// func addRecipeToCollection(w http.ResponseWriter, r *http.Request) {

// }

func handleCollection(w http.ResponseWriter, r *http.Request) {

	var collection Collection
	params := mux.Vars(r)

	if r.Method == "DELETE" {
		db.Delete(&collection, params["collectionid"])
		json.NewEncoder(w).Encode("Deletion successful")

	} else if r.Method == "GET" {
		db.Preload("Recipes").First(&collection, params["collectionid"])
		json.NewEncoder(w).Encode(&collection)
	}
	if err != nil {
		return
	}
}

func handleRecipeCollection(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe
	var collection Collection
	params := mux.Vars(r)

	db.First(&recipe, params["recipeid"])
	db.First(&collection, params["collectionid"])

	if r.Method == "PATCH" {
		db.Model(&collection).Association("Recipes").Append([]Recipe{recipe})
		json.NewEncoder(w).Encode(&collection)
	} else if r.Method == "DELETE" {
		db.Model(&collection).Association("Recipes").Delete(recipe)
		json.NewEncoder(w).Encode(&collection)
	}
	if err != nil {
		return
	}
}

func getCollectionsPerRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe

	params := mux.Vars(r)

	db.Preload("Collections").First(&recipe, params["recipeid"])
	json.NewEncoder(w).Encode(&recipe)

	if err != nil {
		return
	}
}
