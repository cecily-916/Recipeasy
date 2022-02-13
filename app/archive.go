package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getArchive(w http.ResponseWriter, r *http.Request) {
	var recipes []Recipe
	// var user User
	params := mux.Vars(r)

	// db.First(&user, params["userid"]).Unscoped().Association("OwnRecipes").Where(map[string]interface{}{"deleted_at": nil}).Find(&recipes)
	db.Unscoped().Model(&recipes).Not(map[string]interface{}{"deleted_at": nil}).Where("user_id =?", params["userfrontid"]).Find(&recipes)
	json.NewEncoder(w).Encode(&recipes)

	if err != nil {
		return
	}
}

func handleArchivedRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipe Recipe

	if r.Method == "DELETE" {
		db.Unscoped().Delete(&recipe, params["recipeid"])
		json.NewEncoder(w).Encode("deletion successful")
	}
	if r.Method == "PUT" {
		db.Unscoped().First(&recipe, params["recipeid"])
		// db.Model(&recipe).Update("DeletedAt", nil)
		db.Unscoped().Model(&recipe).UpdateColumn("deleted_at", nil)
		json.NewEncoder(w).Encode(&recipe)
	}
	if err != nil {
		return
	}
}
