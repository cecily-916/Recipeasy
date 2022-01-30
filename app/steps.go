package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var steps []Step
var step Step
var recipe Recipe

// GET all steps from a recipe
func handleStepsByRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db.First(&recipe, params["id"])
	db.Model(&recipe).Related(&steps)
	recipe.Steps = steps

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(&recipe)
	} else if r.Method == "DELETE" {
		db.Delete(&steps)
		json.NewEncoder(w).Encode("All steps deleted")
	} else if r.Method == "POST" {
		var newStep Step

		_ = json.NewDecoder(r.Body).Decode(&newStep)

		steps = append(steps, newStep)

		db.Create(&newStep)

		json.NewEncoder(w).Encode(&newStep)
	}
}

// // POST new step to a recipe
// func createStep(w http.ResponseWriter, r *http.Request) {

// // GET single step by stepID
func getStep(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db.First(&steps, params["stepID"])

	json.NewEncoder(w).Encode(&steps)

	if err != nil {
		return
	}

}

// DELETE single step by stepID
func deleteStep(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db.First(&step, params["stepID"])
	db.Delete(&step)

	// returns non-deleted ingredients
	db.Find(&steps)
	json.NewEncoder(w).Encode(&steps)
}

// // PATCH edits the step
// func editStep(w http.ResponseWriter, r *http.Request) {

// }
