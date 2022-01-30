package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var ingredients []Ingredient
var ingredient Ingredient
var fullRecipe []Recipe

// GET all ingredients from a step
func handleIngredientsByStep(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db.First(&step, params["stepID"])
	db.Model(&step).Related(&ingredients)
	step.Ingredients = ingredients

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(&step)
	} else if r.Method == "DELETE" {
		db.Delete(&ingredients)
		json.NewEncoder(w).Encode("All ingredients deleted.")
	}
}

func getFullRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipe Recipe
	var steps []Step

	db.First(&recipe, params["id"])
	db.Model(&recipe).Related(&steps)
	recipe.Steps = steps

	var stepsWithIngredients []Step

	for _, recipeStep := range recipe.Steps {
		db.Model(&recipeStep).Related(&ingredients)
		recipeStep.Ingredients = ingredients
		stepsWithIngredients = append(stepsWithIngredients, recipeStep)
	}

	recipe.Steps = stepsWithIngredients

	json.NewEncoder(w).Encode(&recipe)

}

// POST new ingredient
func createIngredient(w http.ResponseWriter, r *http.Request) {
	var newIngredient Ingredient

	_ = json.NewDecoder(r.Body).Decode(&newIngredient)

	ingredients = append(ingredients, newIngredient)

	db.Create(&newIngredient)
	json.NewEncoder(w).Encode(&newIngredient)
}

// // GET ingredient
func getIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db.First(&ingredients, params["ingredientID"])

	json.NewEncoder(w).Encode(&ingredients)

	if err != nil {
		return
	}

}

// DELETE single ingredient
func deleteIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db.First(&ingredient, params["ingredientID"])
	db.Delete(&ingredient)

	// returns non-deleted ingredients
	db.Find(&ingredients)
	json.NewEncoder(w).Encode(&ingredients)
}

// // PATCH edits the ingredient
// func editIngredient(w http.ResponseWriter, r *http.Request) {

// }
