package main

import "gorm.io/gorm"

// Each recipe has many steps
type Recipe struct {
	gorm.Model
	Title string `json:"title" sql:"size:255;unique;index"`
	Steps []Step `json:"steps"` //one-to-many relationship
}

// Quantities know which ingredient they are a part of but ingredients dont need to know the reverse
// type Ingredient struct {
// 	gorm.Model
// 	Name string `json:"name" sql:"AUTO_INCREMENT" gorm:"primary_key"`
// }

// // Each quantity has one ingredient and one unit of measurement
type Quantity struct {
	gorm.Model

	StepID          int
	Step            Step
	Ingredient      string //one-to-one relationship
	UnitMeasurement string //one-to-one relationship
	Amount          int
	// Step            *Step // one-to-one relationship
}

// Each step has many quantities and belongs to only one recipe
type Step struct {
	gorm.Model

	Details    string     `json:"details"`
	Quantities []Quantity `json:"quantities"` // one-to-many relationship
	Recipe     Recipe     // one-to-one relationship
	RecipeID   int
	Completed  bool `sql:"DEFAULT:false"`
}
