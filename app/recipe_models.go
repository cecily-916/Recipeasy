package main

import "gorm.io/gorm"

// Each recipe has many steps
type Recipe struct {
	gorm.Model
	Title           string             `json:"title" sql:"size:255"`
	Description     string             `json:"description"`
	Rating          int                `json:"rating"`
	PrepTime        string             `json:"prepTime"`
	CookTime        string             `json:"cookTime"`
	Steps           []Step             `json:"steps" gorm:"constraint:OnDelete:CASCADE"` //one-to-many relationship
	Ingredients     []RecipeIngredient `json:"ingredients"`
	NumIngredients  int                `json:"numingredients"`
	NumSteps        int                `json:"numsteps"`
	MainImage       string             `json:"image"`
	Servings        int                `json:"servings"`
	OriginalCreator string             `json:"originalcreator"`
	UserID          int                `json:"user"`
	Collections     []*Collection      `gorm:"many2many:recipe_collection"`
}

// Quantities know which ingredient they are a part of but ingredients dont need to know the reverse
// type Ingredient struct {
// 	gorm.Model
// 	Name string `json:"name" sql:"AUTO_INCREMENT" gorm:"primary_key"`
// }

// // Each quantity has one ingredient and one unit of measurement
type Ingredient struct {
	gorm.Model

	StepID int
	// Step            Step
	Ingredient      string  `json:"ingredient"`
	UnitMeasurement string  `json:"unit"` //one-to-one relationship
	Amount          float32 `json:"amount"`
	IngredientOrder int     `json:"order"` //Order in step that this ingredient is used
	// Completed       bool    `sql:"DEFAULT: false"`
}

// Each step has many ingredients and belongs to only one recipe
type Step struct {
	gorm.Model

	Details      string       `json:"details"`
	ExtraDetails string       `json:"extradetails"`
	Ingredients  []Ingredient `json:"ingredients"` // one-to-many relationship
	// Recipe      Recipe       // one-to-one relationship
	RecipeID  int
	StepOrder int `json:"order"` //inputted by user
	// Completed bool `json:"completed" sql:"DEFAULT:false"`
}
