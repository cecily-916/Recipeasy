package main

import "gorm.io/gorm"

// Each user has many access to specific recipes
// Can create their own recipe decks made up of both shared and personal recipes
// Can share recipes with other people
// Can see the recipes shared with them

type User struct {
	gorm.Model

	//LOOK INTO WHAT COMES OUT OF THE USERFRONT TOKENS!!
	Name               string `json:"name"`
	Username           string `json:"username" sql:"unique"`
	Password           string `json:"password"`
	Email              string `json:"email"`
	OwnRecipeIDs       []int  `json:"ownrecipes"`
	SharedRecipeIDs    []int  `json:"sharedrecipes"`
	FavoritedRecipeIDs []int  `json:"favorited"`
	Decks              []Deck `json:"decks"`
}

type Deck struct {
	gorm.Model

	Title             string `json:"title"`
	Owner             string `json:"owner"`
	RecipeIDs         []int  `json:"recipes"`
	SharedWithUserIDs []int  `json:"sharedwith"`
}
