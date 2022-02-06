package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET request responds with high level details of all recipes
func getRecipes(w http.ResponseWriter, r *http.Request) {
	var recipes []Recipe
	var user User
	params := mux.Vars(r)

	db.First(&user, params["userid"])

	db.Model(&user).Related(&recipes)
	json.NewEncoder(w).Encode(&recipes)

	if err != nil {
		return
	}
}

// //createRecipe adds a recipe from JSON received in request body
// func createRecipe(w http.ResponseWriter, r *http.Request) {
// 	var newRecipe Recipe
// 	_ = json.NewDecoder(r.Body).Decode(&newRecipe)

// 	recipes = append(recipes, newRecipe)

// 	db.Create(&newRecipe)

// 	json.NewEncoder(w).Encode(&newRecipe)

// 	if err != nil {
// 		return
// 	}
// }
