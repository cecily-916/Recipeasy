package main

import "gorm.io/gorm"

// Each user has many access to specific recipes
// Can create their own recipe decks made up of both shared and personal recipes
// Can share recipes with other people
// Can see the recipes shared with them

type User struct {
	gorm.Model

	Name        string       `json:"name"`
	Username    string       `json:"username" sql:"unique"`
	Password    string       `json:"password"`
	Email       string       `json:"email"`
	Collections []Collection `json:"collections"`
	OwnRecipes  []Recipe     `json:"ownrecipes"`
	// Collections []Collection `json:"collections"`
	// SharedRecipes []Recipe     `json:"sharedrecipes"`
}

type Collection struct {
	gorm.Model

	Name    string    `json:"name"`
	UserID  int       `json:"userid"`
	Recipes []*Recipe `gorm:"many2many:recipe_collection"`
}
