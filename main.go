package main

import (
	"github.com/gin-gonic/gin"
)

// Handler function sets up an association in which getRecipes handles requests to the /recipes endpoint path
func main() {
	router := gin.Default()

	// Recipe Routes
	router.GET("/recipes", getRecipes)
	router.POST("/recipes", postRecipes)
	router.GET("/recipes/:id", getRecipeByID)

	// User Routes
	router.GET("/users", getUsers)
	router.POST("/users", postUsers)

	router.Run("localhost:8080")

}
