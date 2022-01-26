package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// recipe represents data about a recipe.
type recipe struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// recipes slice to seed recipe data.
var recipes = []recipe{
	{ID: "1", Title: "Cookies"},
	{ID: "2", Title: "Chicken Stew"},
	{ID: "3", Title: "Rice pudding"},
}

//getRecipes responds with the list of all recipes as JSON
func getRecipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, recipes)
}

//postRecipes adds a recipe from JSON received in request body
func postRecipes(c *gin.Context) {
	var newRecipe recipe

	if err := c.BindJSON(&newRecipe); err != nil {
		return
	}

	recipes = append(recipes, newRecipe)
	c.IndentedJSON(http.StatusCreated, newRecipe)
}

// getRecipeByID responds with the information about a specific recipe
// ID parameter sent by the client then returns the recipe that matches the ID

func getRecipeByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range recipes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipe not found"})
}
