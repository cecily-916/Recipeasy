package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RecipeIngredient struct {
	Name     string
	Quantity float32
	Unit     string
}

// GET request gets list of recipes for user
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

// GET request responds with JSON of recipe details including all steps associated and all ingredients associated per step
func handleRecipe(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var recipe Recipe
	var steps []Step

	// determine the recipe being called
	db.First(&recipe, params["id"])

	if r.Method == "GET" {
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
	} else if r.Method == "DELETE" {
		db.Delete(&recipe)
	}
}

// POST new ingredient
func createRecipe(w http.ResponseWriter, r *http.Request) {
	var newRecipe Recipe

	_ = json.NewDecoder(r.Body).Decode(&newRecipe)

	db.Create(&newRecipe)
	db.Save(&newRecipe)
	json.NewEncoder(w).Encode(&newRecipe)
}

func formatRecipeIngredients(allIngredients []Ingredient) []RecipeIngredient {
	var recipeIngredients []RecipeIngredient
	var n RecipeIngredient

	for _, ingredient := range allIngredients {

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
