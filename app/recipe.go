package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RecipeIngredient struct {
	Name     string
	Quantity int
	Unit     string
}

// GET request responds with JSON of recipe details including all steps associated and all ingredients associated per step
func getFullRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipe Recipe
	var steps []Step

	// determine the recipe being called
	db.First(&recipe, params["id"])

	// assign related steps
	db.Model(&recipe).Related(&steps)

	// for each step in the recipe, find and assign the ingredient details to the step
	// add the ingredient to the recipe's ingredient list
	var ingredients []Ingredient
	var allIngredients []Ingredient
	var stepsWithIngredients []Step

	for _, recipeStep := range steps {
		db.Model(&recipeStep).Related(&ingredients)
		recipeStep.Ingredients = ingredients

		allIngredients = append(allIngredients, ingredients...)
		stepsWithIngredients = append(stepsWithIngredients, recipeStep)
	}
	fmt.Println(allIngredients)
	recipe.Steps = stepsWithIngredients
	recipe.Ingredients = formatRecipeIngredients(allIngredients)
	recipe.NumIngredients = len(recipe.Ingredients)
	recipe.NumSteps = len(recipe.Steps)

	json.NewEncoder(w).Encode(&recipe)

	if err != nil {
		return
	}
}

func formatRecipeIngredients(allIngredients []Ingredient) []RecipeIngredient {
	var recipeIngredients []RecipeIngredient
	var n RecipeIngredient

	for _, ingredient := range allIngredients {
		// n.Name = ingredient.Ingredient

		i, found := Find(recipeIngredients, ingredient.Ingredient)

		if !found {
			n.Name = ingredient.Ingredient
			n.Quantity = ingredient.Amount
			n.Unit = ingredient.UnitMeasurement
			recipeIngredients = append(recipeIngredients, n)
		} else {
			recipeIngredients[i].Quantity += ingredient.Amount
		}

	}

	return recipeIngredients
}

func Find(slice []RecipeIngredient, val string) (int, bool) {
	for i, ingredient := range slice {
		if ingredient.Name == val {
			return i, true
		}
	}
	return -1, false
}

// func addQuantity() {
// 	for _, ingredient := range slice {
// 		if ingredient.Name == val {

// 		}
// 	}
// }
