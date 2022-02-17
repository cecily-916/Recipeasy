package main

import "github.com/jinzhu/gorm"

// Each recipe has many steps
type Recipe struct {
	gorm.Model
	Title       string `json:"title" sql:"size:255"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	PrepTime    string `json:"prepTime"`
	CookTime    string `json:"cookTime"`
	Steps       []Step `json:"steps" gorm:"constraint:OnDelete:CASCADE"` //one-to-many relationship
	Ingredients string `json:"ingredients"`
	// NumIngredients int           `json:"numingredients"`
	MainImage   string        `json:"image"`
	ImageID     string        `json:"imageid"`
	ImageDelete string        `json:"imagedelete"`
	Servings    string        `json:"servings"`
	Author      string        `json:"author"`
	Source      string        `json:"source"`
	UserID      int           `json:"user"`
	Collections []*Collection `gorm:"many2many:recipe_collection"`
}

// // Each quantity has one ingredient and one unit of measurement
type Ingredient struct {
	gorm.Model

	StepID int
	// Step            Step
	Ingredient      string  `json:"ingredient"`
	UnitMeasurement string  `json:"unit"` //one-to-one relationship
	AmountWhole     string  `json:"amountwhole"`
	AmountFrac      string  `json:"amountfrac"`
	Amount          float32 `json:"amount"`
	// Completed       bool    `sql:"DEFAULT: false"`
}

// Each step has many ingredients and belongs to only one recipe
type Step struct {
	gorm.Model

	Details      string       `json:"details"`
	Category     string       `json:"category"`
	ExtraDetails string       `json:"extradetails"`
	Image        string       `json:"image"`
	ImageID      string       `json:"imageid"`
	ImageDelete  string       `json:"imagedelete"`
	Ingredients  []Ingredient `json:"ingredients"` // one-to-many relationship
	// Recipe      Recipe       // one-to-one relationship
	RecipeID int
	// Completed bool `json:"completed" sql:"DEFAULT:false"`
}
