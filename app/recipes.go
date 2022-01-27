package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var recipes []Recipe

func getRecipes(w http.ResponseWriter, r *http.Request) {

	db.Find(&recipes)
	json.NewEncoder(w).Encode(&recipes)

	if err != nil {
		return
	}

}

//createRecipe adds a recipe from JSON received in request body
func createRecipe(w http.ResponseWriter, r *http.Request) {
	var newRecipe Recipe
	_ = json.NewDecoder(r.Body).Decode(&newRecipe)

	recipes = append(recipes, newRecipe)

	db.Create(&newRecipe)

	json.NewEncoder(w).Encode(&newRecipe)

	if err != nil {
		return
	}
}

// getRecipeByID responds with the information about a specific recipe
func getRecipeByID(w http.ResponseWriter, r *http.Request) {

	var recipe Recipe

	params := mux.Vars(r)

	db.First(&recipe, params["id"])
	json.NewEncoder(w).Encode(&recipe)

	if err != nil {
		return
	}
}
