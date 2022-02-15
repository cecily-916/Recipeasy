package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET request responds with high level details of all recipes per user
func getUserData(w http.ResponseWriter, r *http.Request) {
	var user User

	params := mux.Vars(r)

	db.Where("userfront_id =?", params["userfrontid"]).Find(&user)
	// db.Preload("OwnRecipes.Steps.Ingredients").Preload("OwnRecipes.Collections").First(&user, params["userid"])

	json.NewEncoder(w).Encode(&user)

	if err != nil {
		return
	}
}
