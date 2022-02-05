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

	db.First(&recipe, params["id"])

	if r.Method == "DELETE" {
		db.Delete(&recipe)
	} else if r.Method == "PATCH" {
		db.Model(&recipe).Update("deleted_at", "")
	}
}
