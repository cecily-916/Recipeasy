package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var ingredients []Ingredient
var ingredient Ingredient

// GET all ingredients from a step
func getIngredientsByStep(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var step Step

	db.First(&step, params["id"])
	db.Model(&step).Related(&ingredients)
	step.Ingredients = ingredients
	json.NewEncoder(w).Encode(&ingredients)

}

// // POST new step to a recipe
// func createIngredient(w http.ResponseWriter, r *http.Request) {
// }

// // DELETE all ingredients from step
// func deleteIngredients(w http.ResponseWriter, r *http.Request) {

// }

// // GET single step by stepID
func getIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db.First(&ingredients, params["ingredientID"])

	json.NewEncoder(w).Encode(&ingredients)

	if err != nil {
		return
	}

}

// // DELETE single step by stepID
func deleteIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db.First(&ingredient, params["ingredientID"])
	db.Delete(&ingredient)

	// returns non-deleted ingredients
	db.Find(&ingredients)
	json.NewEncoder(w).Encode(&ingredients)
}

// // PATCH edits the step
// func editIngredient(w http.ResponseWriter, r *http.Request) {

// }
