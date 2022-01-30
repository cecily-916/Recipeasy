package main

import "gorm.io/gorm"

// Each recipe has many steps
type Recipe struct {
	gorm.Model
	Title       string `json:"title" sql:"size:255;unique;index"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	PrepTime    string `json:"prep-time"`
	CookTime    string `json:"cook-time"`
	Steps       []Step `json:"steps"` //one-to-many relationship
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
	Ingredient      string `json:"ingredient"`       //one-to-one relationship
	UnitMeasurement string `json:"unit-measurement"` //one-to-one relationship
	Amount          int    `json:"amount"`
	QuantOrder      int    `json:"order"` //Order in step that this ingredient is used
	Completed       bool   `sql:"DEFAULT: false"`
	// Step            *Step // one-to-one relationship
}

// Each step has many ingredients and belongs to only one recipe
type Step struct {
	gorm.Model

	Details     string       `json:"details"`
	Ingredients []Ingredient `json:"ingredients"` // one-to-many relationship
	// Recipe      Recipe       // one-to-one relationship
	RecipeID  int
	StepOrder int  `json:"step"` //inputted by user
	Completed bool `json:"completed" sql:"DEFAULT:false"`
}
