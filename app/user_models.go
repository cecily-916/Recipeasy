package main

import "gorm.io/gorm"

// Each user has many access to specific recipes
// Can create their own recipe decks made up of both shared and personal recipes
// Can share recipes with other people
// Can see the recipes shared with them

type User struct {
	gorm.Model

	Name             string   `json:"name"`
	Username         string   `json:"username" sql:"unique"`
	Password         string   `json:"password"`
	Email            string   `json:"email"`
	OwnRecipes       []Recipe `json:"ownrecipes"`
	SharedRecipes    []Recipe `json:"sharedrecipes"`
	FavoritedRecipes []Recipe `json:"favorited"`
	Decks            []Deck   `json:"decks"`
}

type Deck struct {
	gorm.Model

	Title           string   `json:"title"`
	Owner           string   `json:"owner"`
	Recipes         []Recipe `json:"recipes"`
	SharedWithUsers []User   `json:"sharedwith"`
}
