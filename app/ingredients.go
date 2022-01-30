package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var ingredients []Ingredient
var ingredient Ingredient

// GET all ingredients from a step
func handleIngredientsByStep(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db.First(&step, params["id"])
	db.Model(&step).Related(&ingredients)
	step.Ingredients = ingredients

	if r.Method == "Get" {
		json.NewEncoder(w).Encode(&ingredients)
	} else if r.Method == "DELETE" {
		db.Delete(&ingredients)
		json.NewEncoder(w).Encode("All ingredients deleted.")
	}
}

// POST new ingredient
func createIngredient(w http.ResponseWriter, r *http.Request) {
	var newIngredient Ingredient

	_ = json.NewDecoder(r.Body).Decode(&newIngredient)

	ingredients = append(ingredients, newIngredient)

	db.Create(&newIngredient)
	json.NewEncoder(w).Encode(&newIngredient)
}

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
