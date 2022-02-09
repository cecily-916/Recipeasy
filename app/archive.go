package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getArchive(w http.ResponseWriter, r *http.Request) {
	var recipes []Recipe

	db.Raw("SElECT * FROM recipes where deleted_at is not null").Scan(&recipes)
	fmt.Println(recipes)
	json.NewEncoder(w).Encode(&recipes)
}

func handleArchivedRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipe Recipe

	if r.Method == "DELETE" {
		db.Unscoped().Delete(&recipe, params["recipeid"])
		json.NewEncoder(w).Encode(&recipe)

	}
	if r.Method == "PATCH" {
		db.Unscoped().First(&recipe, params["recipeid"])
		db.Model(&recipe).Update("deleted_at", nil)

		json.NewEncoder(w).Encode(&recipe)
	}
	if err != nil {
		return
	}
}
